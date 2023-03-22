# Implies is transitive

## Status

| human  | copilot | auto   | crush  | sauto  | hammer | gpt4   |
| ---    | ---     | ---    | ---    | ---    | ---    | ---    |
| [✅](./human.v) | [🔄](./auto.v) | [🔄](./copilot.v) | [🔄](./crush.v) | [🔄](./gpt4.v) | [🔄](./hammer.v) | [🔄](./sauto.v) |

## Script

We want to prove the following statement:

```coq
Variables P Q R T : Prop.
Theorem imp_trans : (P -> Q) -> (Q -> R) -> P -> R.
```

We want to prove that if P implies Q and Q implies R then P implies R.

Using intros we can transform the goal, which is statement below the line into multiple hypothesis that are above the line and a smaller goal below the line.

We can then pass P to the PQ function and the resulting Q to the QR function to get R.

We can also prove this alternatively using reverse reasoning.

If we have an hypothesis that goes from Q to R and our goal is R, then we really only need to prove Q.

If we have an hypothesis that goes from P to Q and our goal is Q, then we really only need to prove P.

We have an hypothesis P, so our prove is complete.

## Reference

This theorem was taken from the English version of the [Coq Art book](https://www.labri.fr/perso/casteran/CoqArt/) on page 48.
