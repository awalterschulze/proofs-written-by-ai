From Hammer Require Import Tactics.
Section Example.

Variables P Q R T: Prop.
Theorem apply_example:
  (Q -> R -> T) -> (P -> Q) -> P -> R -> T.
Proof.
sauto.
Qed.


End Example.