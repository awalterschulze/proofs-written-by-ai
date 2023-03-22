# Proofs written by AI

We evaluate how well AI can write proofs.

Unlike code, we don't have to check proofs, since the compiler can do that for us.
This means that if a generated proof is passed as correct by the compiler, no intervention is needed by a human.

Inside the Proofs folder each folder contains a theorem, while each file in that folder contains an attempt by a different AI model to prove the theorem.

## Proof Irrelevance

> There is a slight difference: when developing a program for a specification A, two programs t and t' may not be considered completely equivalent. If A is the specification for a sorting algorithm, if t implements bubblesort, and t' implements quicksort, efficiency considerations make them very different When considering proofs on the other hand, two proofs p and p' of a proposition P play exactly the same role: only their existence matters and is important to ensure the truth of P. This possibility to interchange proofs of a given proposition is called _proof irrelevance_. - Coq'Art

## Proofs

### Coq Art Book

Exercises and Examples from the Coq Art book:

 - Chapter 3: [Propositions and Proofs](./src/Proofs/Props)

## Setup

- Install [Coq 8.16](https://github.com/coq/platform/releases/tag/2022.09.1)
- Install [CoqHammer](https://coqhammer.github.io/)

### Generate Theorems from Readmes

[Golang](https://golang.org/) is required to run `$ go run generate.go`.

This script reads and updates several `Readme.md` files in the project, so be careful when you edit them.
It also parses these `Readme.md` files for new theorems to generate folders and uses the `template` folder as input for the files to generate in these folders.

This means when you add a theorem to a `Readme.md` and run `$ go run generate.go` a folder will be generated for that theorem and a file for each AI model. The input `Readme.md` will also be updated with links to the newly generated folders. It will also regenerate the `_CoqProject` file to only contain the files contained in these theorem folders. It will also update the Status section of the `Readme.md` in each theorem folder, checking whether the theorem has a `Fail`, `Admitted` or `Abort` keyword.

If you want to add a new theorem, simply add your theorem to a `Readme.md`:
1. Give it a heading starting with hashtag hashtag space name (with no link)
2. Add the Coq code for the theorem inside triple backticks with a `coq` annotation to specify the code highlighting.
3. Run `$ go run generate.go`
4. Proof some theorems
5. Run `$ go run generate.go` to update the status
6. Run `$ make` to check your proofs

If you want to add a new folder of proofs:
1. Create a new folder inside the Proofs folder
2. Create a `Readme.md` inside that folder
3. In `generate.go` add a line: `generate("<path to your folder>")`
4. Add at least one theorem as specified above
5. Add the folder to the list of proofs, above in this Readme.

### Clean

Cleans all files, including `.aux` files:

```
$ make cleanall
```

### Regenerate Makefile

Coq version upgrade requires regenerating the Makefile.
This requires having `coq_makefile` in your path.

Here is an example of how to add `coq_makefile` to your PATH:

```
$ export PATH=/Applications/Coq-Platform~8.16~2022.09.app/Contents/Resources/bin:${PATH}
```

Then we can use `coq_makefile` to regenerate the `Makefile`.

```
$ coq_makefile -f _CoqProject -o Makefile
```

## References

- [Coq Art book](https://www.labri.fr/perso/casteran/CoqArt/) - source of theorems and exercises to prove.
- [CoqHammer](https://coqhammer.github.io/) - Machine Learning tool used to solve theorems using tactics such as `sauto` and `hammer`.
- [Github Copilot](https://github.com/features/copilot) - AI pair programmer
- [Chat GPT](https://openai.com/blog/chatgpt) - Large Language Model.
- [Crush](http://adam.chlipala.net/cpdt/) - Smart Coq tactic.