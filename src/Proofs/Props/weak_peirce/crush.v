Require Import Crush.
Section Example.

Variables P Q : Prop.
Lemma weak_peirce : ((((P -> Q) -> P) -> P) -> Q) -> Q.
(* TODO *)
Admitted.


End Example.