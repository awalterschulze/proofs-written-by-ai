package main

import (
	"bytes"
	"os"
	"path"
	"strings"
)

var AllAIs = []string{"human", "auto", "copilot", "crush", "gpt4", "hammer", "sauto"}

type Proof struct {
	AI   string
	Done *bool
}

type Theorem struct {
	Heading string
	Name    string
	Link    *string
	Code    string
	Proofs  []Proof
}

func parseTheoremName(code string) string {
	lines := strings.Split(code, "\n")
	name := ""
	for _, line := range lines {
		if strings.HasPrefix(line, "Theorem ") || strings.HasPrefix(line, "Lemma") {
			noPrefix := strings.TrimPrefix(strings.TrimPrefix(line, "Theorem "), "Lemma ")
			index := strings.IndexAny(noPrefix, ":{(")
			name = strings.TrimSpace(noPrefix[0:index])
		}
	}
	return name
}

func parse(dir string) (string, []Theorem) {
	filename := path.Join(dir, "Readme.md")
	theorems := []Theorem{}
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	theoremStrings := strings.Split(string(data), "\n## ")
	sectionHeading := strings.TrimSpace(strings.Split(theoremStrings[0], "\n")[0])
	theoremStrings = theoremStrings[1:]
	for _, theoremString := range theoremStrings {
		headAndRest := strings.Split(theoremString, "```coq")
		theoremHeading := strings.Split(headAndRest[0], "\n")[0]
		codeAndRest := strings.Split(headAndRest[1], "```")
		code := strings.TrimSpace(codeAndRest[0])
		theoremName := parseTheoremName(code)
		var link *string
		var proofs []Proof = nil
		if strings.HasPrefix(theoremHeading, "[") {
			linky := strings.Split(strings.Split(theoremHeading, "(")[1], ")")[0]
			link = &linky
			theoremHeading = strings.Split(strings.Split(theoremHeading, "[")[1], "]")[0]
			for _, ai := range AllAIs {
				proofFilename := path.Join(path.Join(dir, theoremName), ai+".v")
				if _, err := os.Stat(path.Join(path.Join(dir, theoremName), ai+".v")); err == nil {
					data, err := os.ReadFile(proofFilename)
					if err != nil {
						panic(err)
					}
					var done *bool
					if bytes.Contains(data, []byte("Fail")) || bytes.Contains(data, []byte("Abort")) {
						d := false
						done = &d
					} else if bytes.Contains(data, []byte("Admitted.")) {
						done = nil
					} else {
						d := true
						done = &d
					}
					proofs = append(proofs, Proof{
						AI:   ai,
						Done: done,
					})
				}
			}
		}
		theorems = append(theorems, Theorem{
			Heading: theoremHeading,
			Name:    theoremName,
			Link:    link,
			Code:    code,
			Proofs:  proofs,
		})
	}
	return sectionHeading, theorems
}

func writeCodeFile(code, src, dst string) {
	data, err := os.ReadFile(src)
	if err != nil {
		panic(err)
	}
	s := string(data)
	code += "\n(* TODO *)\nAdmitted.\n"
	s = strings.Replace(s, "(*TODO:CODE*)", code, 1)
	if err := os.WriteFile(dst, []byte(s), os.FileMode(int(0664))); err != nil {
		panic(err)
	}
}

func writeCodeFiles(src, dst string, theorem Theorem) {
	filenames := []string{}
	for _, ai := range AllAIs {
		filenames = append(filenames, ai+".v")
	}
	for _, filename := range filenames {
		writeCodeFile(theorem.Code, path.Join(src, filename), path.Join(dst, filename))
	}
}

func writeReadme(dir string, theorem Theorem) {
	filename := path.Join(dir, "Readme.md")
	readmeBytes, err := os.ReadFile("./src/template/Readme.md")
	if err != nil {
		panic(err)
	}
	readme := string(readmeBytes)
	readme = strings.Replace(readme, "TODO:TITLE", theorem.Heading, 1)
	readme = strings.Replace(readme, "TODO:CODE", theorem.Code, 1)
	if err := os.WriteFile(filename, []byte(readme), os.FileMode(int(0664))); err != nil {
		panic(err)
	}
}

func writeFiles(dir string, theorems []Theorem) {
	for _, theorem := range theorems {
		dstdir := path.Join(dir, theorem.Name)
		if _, err := os.Stat(dstdir); err != nil {
			if err := os.Mkdir(dstdir, os.FileMode(int(0775))); err != nil {
				panic(err)
			}
			writeCodeFiles("./src/template/", dstdir, theorem)
			writeReadme(dstdir, theorem)
		}
	}
}

func rewriteReadme(filename string, heading string, theorems []Theorem) {
	readmeStrings := []string{}
	p := func(s string) {
		readmeStrings = append(readmeStrings, s)
	}
	p("# " + heading + "\n\n")
	for _, theorem := range theorems {
		if theorem.Link != nil {
			p("## [" + theorem.Heading + "](" + *theorem.Link + ")\n\n")
		} else {
			p("## [" + theorem.Heading + "](./" + theorem.Name + ")\n\n")
		}
		p("```coq\n" + theorem.Code + "\n```\n\n")
	}
	if err := os.WriteFile(filename, []byte(strings.Join(readmeStrings, "")), os.FileMode(int(0664))); err != nil {
		panic(err)
	}
}

func writeCoqProjectFile(dir string, theorems []Theorem, filename string) {
	ss := []string{"-R src Example\n\nsrc/Crush/Crush.v\n\n"}
	p := func(s string) {
		ss = append(ss, s)
	}
	for _, ai := range AllAIs {
		p("src/template/" + ai + ".v\n")
	}
	p("\n")
	for _, theorem := range theorems {
		for _, proof := range theorem.Proofs {
			p(path.Join(path.Join(dir, theorem.Name), proof.AI+".v") + "\n")
		}
		p("\n")
	}
	project := strings.Join(ss, "")
	if err := os.WriteFile(filename, []byte(project), os.FileMode(int(0664))); err != nil {
		panic(err)
	}
}

func updateStatuses(dir string, theorems []Theorem) {
	for _, theorem := range theorems {
		theoremDir := path.Join(dir, theorem.Name)
		readmeFilename := path.Join(theoremDir, "Readme.md")

		readmeBytes, err := os.ReadFile(readmeFilename)
		if err != nil {
			panic(err)
		}
		readme := string(readmeBytes)
		readmes := strings.Split(readme, "## Status")
		prefix := readmes[0]
		readmes = strings.Split(readmes[1], "## Script")
		suffix := readmes[1]
		newTables := []string{"| human  | copilot | auto   | crush  | sauto  | hammer | gpt4   |\n"}
		newTables = append(newTables, "| ---    | ---     | ---    | ---    | ---    | ---    | ---    |\n")
		for _, proof := range theorem.Proofs {
			newTables = append(newTables, "| [")
			if proof.Done == nil {
				newTables = append(newTables, "🔄")
			} else if *proof.Done {
				newTables = append(newTables, "✅")
			} else {
				newTables = append(newTables, "❌")
			}
			newTables = append(newTables, "](./"+proof.AI+".v) ")
		}
		newTables = append(newTables, "|")
		newTable := strings.Join(newTables, "")
		readme = prefix + "## Status\n\n" + newTable + "\n\n## Script" + suffix
		if err := os.WriteFile(readmeFilename, []byte(readme), os.FileMode(int(0664))); err != nil {
			panic(err)
		}
	}
}

func generate(dir string) {
	heading, theorems := parse(dir)
	writeFiles(dir, theorems)
	updateStatuses(dir, theorems)
	filename := path.Join(dir, "Readme.md")
	rewriteReadme(filename, heading, theorems)
	writeCoqProjectFile(dir, theorems, "_CoqProject")
}

func main() {
	generate("./src/Proofs/Props/")
}
