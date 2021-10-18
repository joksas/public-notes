# Group theory

## Groups

Group $G$ is a set with an operation $*$ which

* is closed, i.e.\ $\forall a, b \in G\colon a * b \in G$
* has an identity $i$, i.e.\ $\forall a \in G\colon i*a = a*i = a$
* is associative, i.e. $\forall a, b, c \in G\colon a * (b * c)  = (a * b) * c$
* every element $a$ has an inverse $a^{-1}$, i.e.\ $\forall a \in G\colon a * a^{-1} = a^{-1} * a = e$

Abelian or commutative group is a group where the underlying operation works in a commutative manner, i.e.
\begin{equation}
  \forall a, b \in G\colon a * b = b * a
\end{equation}

## Rings

Ring is an abelian group whose (primary) operation $+$ is called addition, and which has a secondary binary operation $\cdot$ called multiplication which

* is closed
* is associative
* is distributive over the addition operation, i.e.\ $\forall a, b, c \in G\colon a \cdot (b + c) = a \cdot b + a \cdot c$

Most definitions of a ring also require multiplication to have an identity.
