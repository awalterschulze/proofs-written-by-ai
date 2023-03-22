# Implies diamond

## Status

| human  | copilot | auto   | crush  | sauto  | hammer | gpt4   |
| ---    | ---     | ---    | ---    | ---    | ---    | ---    |
| [🔄](./human.v) | [🔄](./auto.v) | [🔄](./copilot.v) | [🔄](./crush.v) | [🔄](./gpt4.v) | [🔄](./hammer.v) | [🔄](./sauto.v) |

## Script

We want to prove the following statement:

```coq
Variables P Q R S : Prop.
Lemma diamond : (P -> Q) -> (P -> R) -> (Q -> R -> S) -> P -> S.
```


