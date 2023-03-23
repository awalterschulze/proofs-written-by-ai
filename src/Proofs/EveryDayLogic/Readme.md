
## Less than or Equal the Successor of the Successor

```coq
Theorem le_i_SSi: forall i:nat, i <= S (S i).
```

## Less than or Equal is Transitive

```coq
Theorem le_trans: forall n m p: nat, n <= m -> m <= p -> n <= p.
```

## Less than or Equal is less than or equal to multiplying by a constant

```coq
Theorem mult_le_compat_l: forall n m p: nat, n <= m -> p*n <= p*m.
```

## Less than or Equal is less than or equal to multiplying by a constant on the right side

```coq
Theorem mult_le_r: forall n m p: nat, n <= p -> n*m <= p*m.
```

## Less than or Equal is less than or equal to multiplying by less than or equal

```coq
Theorem le_mult_mult: forall a b c d: nat, a <= c -> b <= d -> a*b <= c*d.
```

## Less than or Equal multiplying by zero

```coq
Theorem le_0_mult : forall n p: nat, 0*n <= 0*p
```

## Less than or Equal multiplying by zero on the right side

```coq
Theorem le_0_mult : forall n p: nat, n*0 <= p*0
```

## Less than is less than successor

```coq
Theorem lt_S : forall n p: nat, n < p -> n <- S p.
```

## Absurd Negation

```coq
Theorem absurd: forall P Q: Prop, P -> ~P -> Q.
```

## Double Negation

```coq
Theorem double_neg_i: forall P: Prop, P -> ~~P.
```

## Contrap

```coq
Theorem contrap: forall A B: Prop, (A -> B) -> ~B -> ~A.
```

## Not False

```coq
Theorem not_false: ~False.
```

## Not Not Not is Not

```coq
Theorem tiple_not_is_not: forall P: Prop, ~ ~ ~ P -> ~ P.
```

## Not Not Not and P is anything

```coq
Theorem triple_not_and_p_is_anything: forall P Q: Prop, ~ ~ ~ P -> P -> Q.
```

## P implies Q and not Q is not P

```coq
Theorem p_imp_q_and_not_q_is_not_p: forall P Q: Prop,
  (P -> Q -> ~Q -> ~P.
```

## P implies Q and P implies not Q and P is anything

```coq
Theorem p_imp_q_p_imp_not_q_and_p_is_anything:
  forall P Q R: Prop,
  (P -> Q) -> (P -> ~Q) -> P -> R.
```

