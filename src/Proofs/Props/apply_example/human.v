Section Example.

Variables P Q R T: Prop.
Theorem apply_example:
  (Q -> R -> T) -> (P -> Q) -> P -> R -> T.
Proof.
intros QRT PQ p r.
apply QRT.
- apply PQ.
  exact p.
- exact r.
Qed.


End Example.