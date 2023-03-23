Section Example.

Variables P Q R T: Prop.
Theorem apply_example:
  (Q -> R -> T) -> (P -> Q) -> P -> R -> T.
Proof.
auto.
Qed.


End Example.