Section Example.

Variables P Q R : Prop.
Theorem imp_dist : (P -> Q -> R) -> (P -> Q) -> (P -> R).
Proof.
intros PQR PQ p.
remember (PQ p) as q.
remember (PQR p q) as r.
assumption.
Qed.

Theorem imp_dist' : (P -> Q -> R) -> (P -> Q) -> (P -> R).
Proof.
intros PQR PQ p.
apply PQR.
- assumption.
- apply PQ.
  assumption.
Qed.


End Example.