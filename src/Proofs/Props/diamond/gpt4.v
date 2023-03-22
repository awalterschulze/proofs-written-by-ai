Section Example.

(*

## Chat GPT Prompt

Please only answer with code and DO NOT provide any explanation after or before. Also do not use the auto tactic. Also give each hypothesis a descriptive but short name as you introduce it.

Using the Coq Proof Assistant version 8.16 can you prove the following theorem:

Variables P Q R S : Prop.
Lemma diamond : (P -> Q) -> (P -> R) -> (Q -> R -> S) -> P -> S.

*)

Variables P Q R S : Prop.
Lemma diamond : (P -> Q) -> (P -> R) -> (Q -> R -> S) -> P -> S.
(* TODO *)
Admitted.


End Example.