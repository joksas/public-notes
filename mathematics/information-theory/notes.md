---
title: "Information Theory"
---

These are my notes for information theory based mostly on the textbook "Elements of Information Theory" by Thomas M. Cover and Joy A. Thomas.

# Chapter 2

## 2.1

**Entropy** of a discrete random variable $X$ is
\begin{equation}
  H(x) \equiv - \sum_{x \in \mathcal{X}} p(x) \log p(x)
\end{equation}

**Expectation** of a random variable $X$ is
\begin{equation}
  E_p g(X) \equiv \sum_{x \in \mathcal{X}} g(x) p(x)
\end{equation}

## 2.2

**Joint entropy** of a pair of discrete random variables $(X, Y)$ with a joint distribution $p(x, y)$ is
\begin{equation}
  H(X, Y) \equiv - \sum_{x \in \mathcal{X}} \sum_{y \in \mathcal{Y}} p(x, y) \log p(x, y) = -E \log p(X, Y)
\end{equation}

**Conditional entropy** is
\begin{align}
  H(Y | X) &\equiv \sum_{x \in \mathcal{X}} p(x) H(Y | X = x) \\
	   &= -\sum_{x \in \mathcal{X}} p(x) \sum_{y \in \mathcal{Y}} p(y|x) \log p(y|x) \\
	   &= -\sum_{x \in \mathcal{X}} \sum_{y \in \mathcal{Y}} p(x, y) \log p(y|x)
\end{align}

**Chain rule** involving conditional entropy is
\begin{equation} \label{eq:chain-rule-conditional-entropy}
  H(X, Y) = H(X) + H(Y|X)
\end{equation}

## 2.3

**Relative entropy** or **Kullback-Leibler distance** is
\begin{align}
  D(p || q) &\equiv \sum_{x \in \mathcal{X}} p(x) \log \frac{p(x)}{q(x)} \\
	    &= E_p \log \frac{p(X)}{q(X)}
\end{align}

* Relative entropy is a measure of the distance between two distributions. If we knew distribution $p$, we could construct code with average description length $H(p)$.
* If we used code for distribution $q$, the average description length for $p$ would be $H(p) + D(p || q)$.

Relative entropy is not a true distance because it is not symmetric and does not satisfy triangle inequality. Nevertheless, it is often useful to think of relative entropy as "distance".

**Mutual information** is the relative entropy between the joint distribution and the product distribution:
\begin{equation}
  I(X; Y) \equiv D \left( p(x, y) || p(x) p(y) \right)
\end{equation}

## 2.4

We can show that
\begin{equation}
  I(X; Y) = H(X) - H(X|Y) = H(Y) - H(Y|X)
\end{equation}

Mutual information is the reduction in uncertainty of one variable due to the knowledge of the other variable.

Also
\begin{equation}
  I(X; Y) = H(X) + H(Y) - H(X, Y)
\end{equation}
and
\begin{equation}
  I(X; X) = H(X)
\end{equation}

## 2.5

By repeatedly applying Equation \eqref{eq:chain-rule-conditional-entropy} we can show that
\begin{equation} \label{eq:chain-rule-conditional-entropy-multivar}
  H(X_1, X_2, \dots, X_n) = \sum_{i=1}^n H(X_i|X_{i-1}, \dots, X_1)
\end{equation}

We can define **conditional mutual information** as the reduction in the uncertainty of $X$ due to knowledge of $Y$ when $Z$ is given:
\begin{equation}
  I(X; Y | Z) \equiv H(X | Z) - H(X | Y, Z)
\end{equation}

Mutual information satisfies chain-rule-like relation, similar to the one with entropy in Equation \eqref{eq:chain-rule-conditional-entropy-multivar}:
\begin{equation} \label{eq:chain-rule-conditional-mutual-information-multivar}
  I(X_1, X_2, \dots, X_n ; Y) = \sum_{i=1}^n I(X_i ; Y |X_{i-1}, \dots, X_1)
\end{equation}

**Conditional relative entropy** is
\begin{equation}
  D( p(y|x) || q(y|x) ) \equiv \sum_{x} p(x) \sum_{y} p(y|x) \log \frac{p(y|x)}{q(y|x)}
\end{equation}

It can be shown that
\begin{equation}
  D( p(x, y) || q(x, y) ) = D( p(x) || q(x) ) + D( p(y|x) || q(y|x) )
\end{equation}

## 2.6

A function is **convex** if it lies above all its tangents, and **concave** if it lies below all its tangents. Alternatively, one may say that convex function lies below all its chords, and concave function lies above all its chords.

The latter definition may be stated using a formula. For a function $f$ to be convex over an interval $(a, b)$, the following must hold for all $x_1, x_2 \in (a, b)$ and all $\lambda \in [0, 1]$:
\begin{equation} \label{eq:convex-function-definition}
  f(\lambda x_1 + (1-\lambda) x_2) \leq \lambda f(x_1) + (1-\lambda) f(x_2)
\end{equation}

The RHS of Equation \eqref{eq:convex-function-definition} represents the chord between points $(x_1, f(x_1))$ and $(x_2, f(x_2))$, and the LHS represents the value of the function $f$ when $x$ in the interval $[x_1, x_2]$ is plugged in.

If function $f$ has a non-negative second derivative over an interval, then it is convex over that interval.

**Jensen's inequality** states that for convex function $f$ we have
\begin{equation}
  E f(X) \geq f(EX)
\end{equation}

This inequality can be used to prove a number of theorems:
\begin{equation}
  D(p||q) \geq 0
\end{equation}

\begin{equation}
  I(X; Y) \geq 0
\end{equation}

\begin{equation}
  D( p(y|x) || q(y|x) ) \geq 0
\end{equation}

\begin{equation}
  I( X; Y | Z) \geq 0
\end{equation}

Any random variable with range $\mathcal{X}$ has entropy no greater than $\log |\mathcal{X}|$, where $|\mathcal{X}|$ is the number of elements in $X$.

We can show this by assuming that $u(x) = \frac{1}{|\mathcal{X}|}$ is uniform distribution:
\begin{align}
  D(p || u) &= \sum_{x} p(x) \log \frac{p(x)}{u(x)} \\
	    &= \sum_{x} p(x) \log \left( p(x) |\mathcal{X}| \right) \\
	    &= \sum_{x} p(x) \log (p(x)) + \sum_{x} p(x) \log |\mathcal{X}| \\
	    &= -H(x) + \log |\mathcal{X}| \\
	    &\geq 0 \\
  \Rightarrow H(x) &\leq \log |\mathcal{X}|
\end{align}

Thus we get maximum entropy only when there is uniform distribution over the range $\mathcal{X}$.

**Conditioning reduces entropy**:
\begin{equation}
  H(X|Y) \leq H(X)
\end{equation}

This essentially says that knowing another variable $Y$ can only reduce the uncertainty in $X$. However, that is the case only *on average*---we might have a situation where $H(X|Y=y) > H(X)$, however it will always be the case that $\sum\limits_{y} p(y) H(X|Y=y) \leq H(X)$.

We can also prove that
\begin{equation}
  H(X_1, X_2, \dots, X_n) \leq \sum_{i=1}^{n} H(X_i)
\end{equation}
with equality if and only if $X_i$ are independent.
