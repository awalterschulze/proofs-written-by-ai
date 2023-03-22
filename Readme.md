# Proofs written by AI

We evaluate how well AI can write proofs.

Unlike code, we don't have to check proofs, since the compiler can do that for us.
This means that if a generated proof is passed as correct by the compiler, no intervention is needed by a human.

Each folder contains a theorem, while each file in that folder contains an attempt by an AI model to prove the theorem.

## Setup

- Install [Coq 8.16](https://github.com/coq/platform/releases/tag/2022.09.1)
- Install [CoqHammer](https://coqhammer.github.io/)

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