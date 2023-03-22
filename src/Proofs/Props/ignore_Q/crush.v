Require Import Crush.
Section Example.

Variables P Q R : Prop.
Lemma ignore_Q : (P -> R) -> P -> Q -> R.
(* TODO *)
Admitted.


End Example.