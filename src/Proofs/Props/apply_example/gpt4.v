Section Example.

(*

## Chat GPT Prompt

Please only answer with code and DO NOT provide any explanation after or before. Also DO NOT use the auto tactic. Also give each hypothesis a descriptive but short name as you introduce it. Also you MUST copy and paste the code I provided EXACTLY as I wrote it as the start of your proof. Also you MUST NOT import any libraries.

Using the Coq Proof Assistant version 8.16 can you prove the following theorem:

Variables P Q R T: Prop.
Theorem apply_example:
  (Q -> R -> T) -> (P -> Q) -> P -> R -> T.

*)

Variable P Q R T: Prop.
Theorem apply_example:
  (Q -> R -> T) -> (P -> Q) -> P -> R -> T.
Proof.
intros H1 H2 H3 H4.
apply H1; [apply H2 | ]; assumption.
Qed.


End Example.