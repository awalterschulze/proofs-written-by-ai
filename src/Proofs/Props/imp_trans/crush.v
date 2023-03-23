Require Import Crush.
Section Example.

Variables P Q R T : Prop.
Theorem imp_trans : (P -> Q) -> (Q -> R) -> P -> R.
crush.
Qed.


End Example.