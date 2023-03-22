# Implies weak Peirce

## Status

| human  | copilot | auto   | crush  | sauto  | hammer | gpt4   |
| ---    | ---     | ---    | ---    | ---    | ---    | ---    |
| [🔄](./human.v) | [🔄](./auto.v) | [🔄](./copilot.v) | [🔄](./crush.v) | [🔄](./gpt4.v) | [🔄](./hammer.v) | [🔄](./sauto.v) |

## Script

We want to prove the following statement:

```coq
Variables P Q : Prop.
Lemma weak_peirce : ((((P -> Q) -> P) -> P) -> Q) -> Q.
```


