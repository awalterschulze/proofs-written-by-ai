# Implies is Distributive

We want to prove the following statement:

```coq
Variables P Q R T : Prop.
Theorem imp_dist : (P -> Q -> R) -> (P -> Q) -> (P -> R).
```

We want to prove that
if P and Q implies R
and P implies Q
then P implies R.



## Reference

This theorem was taken from the English version of the [Coq Art book](https://www.labri.fr/perso/casteran/CoqArt/) on page 54.
