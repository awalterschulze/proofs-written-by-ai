# **TODO**

We want to prove the following statement:

```coq
Variables P Q R T.
Theorem apply_example: (Q -> R -> T) -> (P -> Q) -> P -> R -> T.
```

If Q and R implies T
and P implies Q
then given P and R
we should be able to show T



## Reference

This theorem was taken from the English version of the [Coq Art book](https://www.labri.fr/perso/casteran/CoqArt/) on page 54.
