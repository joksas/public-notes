---
bibliography: bibliography.bib
nocite: '@*'
---

# Group theory

## Groups

Group $G$ is a set with an operation $*$ which

* is closed, i.e.\ $\forall a, b \in G\colon a * b \in G$
* has an identity $i$, i.e.\ $\forall a \in G\colon e*a = a*e = a$
* is associative, i.e. $\forall a, b, c \in G\colon a * (b * c)  = (a * b) * c$
* every element $a$ has an inverse $a^{-1}$, i.e.\ $\forall a \in G\colon a * a^{-1} = a^{-1} * a = e$

Abelian or commutative group is a group where the underlying operation works in a commutative manner, i.e.
\begin{equation}
  \forall a, b \in G\colon a * b = b * a
\end{equation}

In additive groups, identity element is denoted by $0$, while in multiplicative groups it is denoted by $1$.

Order of a group is the number of elements in it.

## Subgroups

Subgroup $H$ is a subset of group $G$ and is itself a group. It is denoted by
\begin{equation}
  H \leq G
\end{equation}

Any group $G$ is also a subgroup of itself. If $H$ is a subgroup of $G$ but not the same as $G$, it is called a proper subgroup. It is denoted by
\begin{equation}
  H < G
\end{equation}

## Cayley tables

Features:

* if identity is the first element in the row and column headers, the first row and column are just repeats of the headers
* every column contains the identity (because every element in group has an inverse)
* for abelian groups, the table is symmetric around the diagonal
* no duplicate elements in rows or columns

*Proof for last statement:*

Assume there's a row $a$ with duplicate elements $z$ corresponding to two columns $x$ and $y$. Then
\begin{align*}
  a * x &= a * y \\
  a^{-1} * (a * x) &= a^{-1} * (a * y) \\
  (a^{-1} * a) * x &= (a^{-1} * a) * y \\
  e * x &= e * y \\
  x &= y
\end{align*}
But $x$ and $y$ are different---a contradiction. Thus, rows don't have duplicate elements. Analogous proof for showing no duplicate elements in columns.

Using Cayley tables, we can show that there is one group of order 1, one group of order 2, one group of order 3, and 2 (distinct) groups of order 4.

## Cosets and Lagrange's theorem

**Left coset** is defined as
\begin{equation}
  g_i H \equiv \{ g_i * h \text{ for all } h \in H\}
\end{equation}
where $H \leq G$ and $g_i \in G$.

Similarly, **right coset** is defined as
\begin{equation}
  H g_i \equiv \{ h * g_i \text{ for all } h \in H\}
\end{equation}
where $H \leq G$ and $g_i \in G$.

Every group $G$ has at least two subgroups:

1. $G$
2. Trivial group $\{e\}$

Lagrange's theorem:
\begin{equation}
  H \leq G \implies |H| \text{ divides } |G|
\end{equation}

For example, if $|G| = 15$, then $|H| \in \{1, 3, 5, 15\}$. However, Lagrange's theorem doesn't say that subgroups of such order *necessarily* exist.

It can be shown that
\begin{equation}
  g_1 \notin H \implies H \cap g_1 H = \varnothing
\end{equation}

We can also pick $g_2 \in G$ and show that
\begin{equation}
  g_2 \notin H \quad \text{and} \quad g_2 \notin g_1 H \implies H \cap g_2 H = \varnothing \quad \text{and} \quad g_1 H \cap g_2 H = \varnothing
\end{equation}

Repeating this process until all elements in $G$ are covered splits $G$ into non-overlapping left cosets $H, g_1 H, g_2 H, \ldots, g_n H$, each of size $|H|$. Thus, $|H|$ divides $|G|$.

## Congruence classes

For integers modulo $n$ we may separate the integers into $n$ sets depending on the remainder $r$. We can denote these sets as $[r]_n$. For example, for $n = 3$, we will have
\begin{align*}
  [0]_3 &= \{\ldots, -6, -3, 0, 3, 6, \ldots\} \\
  [1]_3 &= \{\ldots, -5, -2, 1, 4, 7, \ldots\} \\
  [2]_3 &= \{\ldots, -4, -1, 2, 5, 8, \ldots\}
\end{align*}

In general, we will have
\begin{equation}
  [r]_n = \{z \in \mathbb{Z} | a - z = kn, k \in \mathbb{Z}\}
\end{equation}

These sets are called **congruence classes**. These classes form a group under addition (where addition is defined as adding any two elements from any of the classes). $[0]_n$ is always the identity element.

When two integers $a$ and $b$ are in the same congruence class, we write
\begin{equation}
  a \equiv b \pmod{n}
\end{equation}
and say "$a$ is congruent to $b$ mod $n$".

## Normal groups and quotient groups

Integers $\mathbb{Z}$ under addition have infinite number of subgroups (multiples of $n$): $\mathbb{Z}, 2\mathbb{Z}, 3\mathbb{Z}, 4\mathbb{Z}, \ldots$

If we look at $3\mathbb{Z}$, we may form three sets of numbers:
\begin{align*}
  3\mathbb{Z} &= \{\ldots, -6, -3, 0, 3, 6, \ldots\} \\
  1 + 3\mathbb{Z} &= \{\ldots, -5, -2, 1, 4, 7, \ldots\} \\
  2 + 3\mathbb{Z} &= \{\ldots, -4, -1, 2, 5, 8, \ldots\}
\end{align*}
$1 + 3\mathbb{Z}$ and $2 + 3\mathbb{Z}$ are left cosets\footnote{$3\mathbb{Z}$ may also be thought of as a left coset $0 + 3\mathbb{Z}$.} but not groups because they

* are not closed under addition
* don't have inverses
* don't have identity elements

They also don't overlap with $3\mathbb{Z}$ but combined together with it cover all numbers in $\mathbb{Z}$. When cosets of subgroup $N$ form a group we call that it a **quotient group** and we refer to $N$ as a **normal subgroup**. To denote that $N$ is normal subgroup of $G$, we write $N \trianglelefteq G$.

When cosets of subgroup $N$ of group $G$ themselves form a group, we have
\begin{equation}
  x * y \in (xN) * (yN) \quad \text{or equivalently} \quad (xN) * (yN) = xyN
\end{equation}

We can show that
\begin{equation}
  (xN) * (yN) = xyN \iff y^{-1} N y = N \quad \text{for any $y \in G$}
\end{equation}

## Cyclic groups

Cyclic group is a group that is generated by a single element $a$, i.e.\ every element is a power of $a$ is a power of $a$ if the group is multiplicative or a multiple of $a$ if the group is additive.

Multiplicative generator $\langle g \rangle$ is given by
\begin{equation}
  \langle g \rangle \equiv  \{\ldots, x^{-2}, x^{-1}, 1, x^1, x^2, \ldots\}
\end{equation}

Additive generator $\langle g \rangle$ is given by
\begin{equation}
  \langle g \rangle \equiv  \{\ldots, -2x, -x, 0, x, 2x, \ldots\}
\end{equation}

Integers $\mathbb{Z}$ are a cyclic group under addition because
\begin{equation}
  \langle 1 \rangle = \{\ldots, -2 \times 1, -1 \times 1, 0, 1 \times 1, 2 \times 1, \ldots\}
\end{equation}

## Rings

Ring is an abelian group whose (primary) operation $+$ is called addition, and which has a secondary binary operation $\cdot$ called multiplication which

* is closed
* is associative
* is distributive over the addition operation, i.e.\ $\forall a, b, c \in G\colon a \cdot (b + c) = a \cdot b + a \cdot c$

Most definitions of a ring also require multiplication to have an identity.

## Bibliography
