# CoqArt - Chapter 3 : Propositions and Proofs

## [Implies is transitive](./imp_trans)

```coq
Variables P Q R T : Prop.
Theorem imp_trans : (P -> Q) -> (Q -> R) -> P -> R.
```

## [Apply example](./apply_example)

```coq
Variables P Q R T: Prop.
Theorem apply_example: (Q -> R -> T) -> (P -> Q) -> P -> R -> T.
```

## [Implies is distributive](./imp_dist)

```coq
Variables P Q R : Prop.
Theorem imp_dist : (P -> Q -> R) -> (P -> Q) -> (P -> R).
```

## [Implies identity](./id_P)

```coq
Variables P : Prop.
Lemma id_P : P -> P.
```

## [Implies identity p p](./id_PP)

```coq
Variables P : Prop.
Lemma id_PP : (P -> P) -> P -> P.
```

## [Implies permutation](./imp_perm)

```coq
Variables P Q R : Prop.
Lemma imp_perm : (P -> Q -> R) -> Q -> P -> R.
```

## [Implies ignore Q](./ignore_Q)

```coq
Variables P Q R : Prop.
Lemma ignore_Q : (P -> R) -> P -> Q -> R.
```

## [Implies delta](./delta_imp)

```coq
Variables P Q : Prop.
Lemma delta_imp : (P -> P -> Q) -> P -> Q.
```

## [Implies delta right](./delta_impR)

```coq
Variables P Q : Prop.
Lemma delta_impR : (P -> Q) -> P -> P -> Q.
```

## [Implies diamond](./diamond)

```coq
Variables P Q R S : Prop.
Lemma diamond : (P -> Q) -> (P -> R) -> (Q -> R -> S) -> P -> S.
```

## [Implies weak Peirce](./weak_peirce)

```coq
Variables P Q : Prop.
Lemma weak_peirce : ((((P -> Q) -> P) -> P) -> Q) -> Q.
```

