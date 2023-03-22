Section Example.

Variables P Q R T : Prop.
Theorem imp_trans : (P -> Q) -> (Q -> R) -> P -> R.
Proof.
intros PQ QR p.
exact (QR (PQ p)).
Qed.

Theorem imp_trans' : (P -> Q) -> (Q -> R) -> P -> R.
Proof.
intros PQ QR p.
apply QR.
apply PQ.
assumption.
Qed.


End Example.