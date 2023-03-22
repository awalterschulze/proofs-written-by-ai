Section Example.

Variables P Q R T : Prop.
Theorem imp_trans : (P -> Q) -> (Q -> R) -> P -> R.
Proof.
intros H1 H2 H3.
apply H2.
apply H1.
apply H3.
Qed.

End Example.