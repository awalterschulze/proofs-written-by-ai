Section Example.

(*

Chat GPT 4

Please only answer with code and DO NOT provide any explanation after or before. Also do not use the auto tactic. Also give each hypothesis a descriptive but short name as you introduce it.

Using the Coq Proof Assistant version 8.16 can you prove the following theorem:

Variables P Q R T : Prop.
Theorem imp_trans : (P -> Q) -> (Q -> R) -> P -> R.

*)

Variables P Q R T : Prop.
Theorem imp_trans : (P -> Q) -> (Q -> R) -> P -> R.
Proof.
intros PQ QR p.
apply QR, PQ, p.
Qed.

End Example.