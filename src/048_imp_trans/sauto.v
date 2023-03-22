From Hammer Require Import Tactics.
Section Example.

Variables P Q R T : Prop.
Theorem imp_trans : (P -> Q) -> (Q -> R) -> P -> R.
sauto.
Qed.

End Example.