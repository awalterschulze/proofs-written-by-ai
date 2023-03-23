# Apply example

<a href="https://www.youtube.com/watch?v=GBppco4psFk&list=PLYwF9EIrl42TTu08EiK2hPYb2iXOR-OQt&index=2" target="_blank">
 <img src="https://img.youtube.com/vi/GBppco4psFk/maxres1.jpg" alt="Watch the video" width="480" border="10" />
</a>

## Status

| human  | copilot | auto   | crush  | sauto  | hammer | gpt4   |
| ---    | ---     | ---    | ---    | ---    | ---    | ---    |
| [✅](./human.v) | [✅](./auto.v) | [✅](./copilot.v) | [✅](./crush.v) | [✅](./gpt4.v) | [✅](./hammer.v) | [✅](./sauto.v) |

## Script

We want to prove the following statement:

```coq
Variables P Q R T: Prop.
Theorem apply_example: (Q -> R -> T) -> (P -> Q) -> P -> R -> T.
```

If Q and R implies T
and P implies Q
then given P and R
we should be able to show T

