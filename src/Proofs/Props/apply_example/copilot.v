Section Example.

Variables P Q R T: Prop.
Theorem apply_example:
  (Q -> R -> T) -> (P -> Q) -> P -> R -> T.
Proof.
  intros H1 H2 H3 H4.
  apply H1.
    apply H2.
      apply H3.
    apply H4.
  Qed.


End Example.