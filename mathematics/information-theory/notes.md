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
  \begin{split}
    H(Y | X) &\equiv \sum_{x \in \mathcal{X}} p(x) H(Y | X = x) \\
	     &= -\sum_{x \in \mathcal{X}} p(x) \sum_{y \in \mathcal{Y}} p(y|x) \log p(y|x) \\
	     &= -\sum_{x \in \mathcal{X}} \sum_{y \in \mathcal{Y}} p(x, y) \log p(y|x)
  \end{split}
\end{align}

**Chain rule** involving conditional entropy is
\begin{equation} \label{eq:chain-rule-conditional-entropy}
  H(X, Y) = H(X) + H(Y|X)
\end{equation}

## 2.3

**Relative entropy** or **Kullback-Leibler distance** is
\begin{align}
  \begin{split}
    D(p || q) &\equiv \sum_{x \in \mathcal{X}} p(x) \log \frac{p(x)}{q(x)} \\
	      &= E_p \log \frac{p(X)}{q(X)}
  \end{split}
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
  \begin{split}
    D(p || u) &= \sum_{x} p(x) \log \frac{p(x)}{u(x)} \\
	      &= \sum_{x} p(x) \log \left( p(x) |\mathcal{X}| \right) \\
	      &= \sum_{x} p(x) \log (p(x)) + \sum_{x} p(x) \log |\mathcal{X}| \\
	      &= -H(x) + \log |\mathcal{X}| \\
	      &\geq 0 \\
    \Rightarrow H(x) &\leq \log |\mathcal{X}|
  \end{split}
\end{align}

Thus we get maximum entropy only when there is uniform distribution over the range $\mathcal{X}$.

**Conditioning reduces entropy**:
\begin{equation} \label{eq:conditioning}
  H(X|Y) \leq H(X)
\end{equation}

This essentially says that knowing another variable $Y$ can only reduce the uncertainty in $X$. However, that is the case only *on average*---we might have a situation where $H(X|Y=y) > H(X)$, however it will always be the case that $\sum\limits_{y} p(y) H(X|Y=y) \leq H(X)$.

We can also prove that
\begin{equation}
  H(X_1, X_2, \dots, X_n) \leq \sum_{i=1}^{n} H(X_i)
\end{equation}
with equality if and only if $X_i$ are independent.

## 2.7

**Log sum inequality** states that for non-negative numbers $a_1, a_2, \dots, a_n$ and $b_1, b_2, \dots, b_n$, we have
\begin{equation} \label{eq:log-sum-inequality}
  \sum_{i=1}^{n} a_i \log \frac{a_i}{b_i} \geq a \log \frac{a}{b}
\end{equation}
where $a \equiv \sum\limits_{i=1}^{n} a_i$ and $b \equiv \sum\limits_{i=1}^{n} b_i$.

We can prove this by introducing function $f(x)= x \log x$ (which is convex since $f''(x) > 0$):
\begin{align}
  \begin{split}
    \sum_{i=1}^{n} a_i \log \frac{a_i}{b_i} &= \sum_{i=1}^{n} b_i f \left( \frac{a_i}{b_i} \right) \\
					    &= b \sum_{i=1}^{n} \frac{b_i}{b} f \left( \frac{a_i}{b_i} \right) \\
					    &\geq b f \left( \sum_{i=1}^{n} \frac{b_i}{b} \frac{a_i}{b_i} \right) \\
					    &= b f \left( \frac{1}{b} \sum_{i=1}^{n} a_i \right) \\
					    &= b f \left( \frac{a}{b} \right) \\
					    &= a \log \frac{a}{b}
  \end{split}
\end{align}

We can use Equation \ref{eq:log-sum-inequality} to prove a number of theorems in information theory. For example, we can use it to show that relative entropy is always non-negative:
\begin{align}
  \begin{split}
    D(p||q) &= \sum_{x \in \mathcal{X}} p(x) \log \frac{p(x)}{q(x)} \\
	    &\geq \sum_{x \in \mathcal{X}} p(x) \log \frac{\sum\limits_{x \in \mathcal{X}} p(x)}{\sum\limits_{x \in \mathcal{X}} q(x)} \\
	    &= 1 \cdot \log \frac{1}{1} \\
	    &= 0
  \end{split}
\end{align}

We can also show that **relative entropy** $D(p||q)$ **is convex**.

By extension, because entropy $H(p)$ can be expressed as $\log |\mathcal{X}| - D(p||u)$, we deduce that **entropy** $H(p)$ **is a concave function**.

Suppose that $(X, Y) \sim p(x, y) = p(y|x) p(x)$. Mutual information $I(X; Y)$ is a concave function of $p(x)$ for fixed $p(y|x)$ and a convex function of $p(y|x)$ for fixed $p(x)$.

## 2.8

$X$, $Y$, and $Z$ form a Markov chain $X \to Y \to Z$ if
\begin{equation}
  p(x, y, z) = p(z|y) p(y|x) p(x)
\end{equation}

Several important remarks:
\begin{itemize}
  \item $X \to Y \to Z$ if and only if there is conditional independence of $X$ and $Z$, given $Y$, i.e. $p( x, z | y ) = p(x|y) p(z|y)$.
  \item $X \to Y \to Z$ implies $Z \to Y \to X$.
  \item If $Z = f(Y)$, then $X \to Y \to Z$.
\end{itemize}

**Data processing inequality** states that if $X \to Y \to Z$, then
\begin{equation}
  I(X; Y) \geq I(X; Z)
\end{equation}

Similarly, one can show that
\begin{equation}
  I(Y; Z) \geq I(X; Z)
\end{equation}

Importantly, if $Z = g(Y)$, then $I(X; Y) \geq I(X; g(Y))$, thus functions of $Y$ cannot increase information about $X$.

Also
\begin{equation}
  I(X; Y|Z) \leq I(X; Y)
\end{equation}
thus dependence of $X$ and $Y$ can only be decreased by the observation of a "downstream" variable $Z$.

## 2.9

Suppose we have a family of probability mass functions $\{f_\theta(x)\}$ indexed by $\theta$. Let $X$ be a sample from a distribution in this family. Let $T(X)$ be any statistic (function of the sample). Then $\theta \to X \to T(X)$ and so
\begin{equation*}
  I(\theta; T(X)) \leq I(\theta; X)
\end{equation*}

If equality holds, then no information is lost, i.e. there is nothing else we can deduce about $\theta$ from $X$ that $T(X)$ doesn't already tell us.

We call such $T(X)$ **sufficient statistics**; formally, they are statistics for which $\theta \to T(X) \to X$, or equivalently
\begin{equation}
  I(\theta; X) = I(\theta; T(X))
\end{equation}

Following from the example [here](https://www.youtube.com/watch?v=5j4E2FRR384), suppose that we toss a coin multiple times, and each individual toss can be modelled using Bernoulli distribution:
\begin{equation*}
  X_i \sim \mathrm{Ber}(\theta) \quad \text{with PMF } f(X_i, \theta) = \theta^{X_i} (1 - \theta)^{1-X_i}, \text{ where } X_i \in \{0, 1\}
\end{equation*}

Given $N$ tosses, we would have
\begin{align*}
  \begin{split}
    f_N(X, \theta) &= \prod_{i=1}^{N} \theta^{X_i} (1 - \theta)^{1-X_i} \\
		   &= \theta^{\sum_{i=1}^N X_i} (1 - \theta)^{N - \sum_{i=1}^N X_i} \\
		   &= \theta^{T(X)} (1 - \theta)^{N - T(X)}
  \end{split}
\end{align*}

So, if we know $N$ in advance, $T(X) = \sum\limits_{i=1}^{N} X_i$ is a sufficient statistic. Nothing else from $X$, such as the order of 1's and 0's would provide additional information about $\theta$.

An alternative way to prove this is alluded to in the textbook and provided in detail [here](https://www.youtube.com/watch?v=G5pLCoPXr9g). We can show that given $T(X)$, all sequences having that many 1's are equally likely and independent of $\theta$ (thus $X$ could not provide any additional insight into $\theta$ beyond what $T(X)$ could):
\begin{align*}
  \begin{split}
    \Pr \left\{ X = x \middle| \sum\limits_{i=1}^{N} X_i = T \right\} &= \frac{\Pr \left\{ X = x, \sum\limits_{i=1}^{N} X_i = T \right\} }{\Pr \left\{ \sum\limits_{i=1}^{N} X_i = T \right\} } \\
	&= \frac{\Pr \left\{ X = x \right\} }{\Pr \left\{ \sum\limits_{i=1}^{N} X_i = k \right\} } \\
	&= \frac{\theta^{T} (1 - \theta)^{N - T}}{\binom{N}{T} \theta^{T} (1 - \theta)^{N - T}} \\
	&= \frac{1}{\binom{N}{T}}
  \end{split}
\end{align*}

$T(X)$ is a **minimal sufficient statistic** if it is a function of every other sufficient statistic $U(X)$, or equivalently
\begin{equation}
  \theta \to T(X) \to U(X) \to X
\end{equation}

We can say that minimal sufficient statistic maximally compresses the information about $\theta$ in the sample $X$. For example, for a normal distribution, a tuple $(m_o, m_s)$ containing the means of odd and even samples, respectively, is a sufficient statistic, but not a minimal sufficient statistic. Sample mean can be extracted from that tuple, but not the other way around.

## 2.10

**Fano's ineqaulity** states that for any estimator $\hat{X}$ such that $X \to Y \to \hat{X}$, w have
\begin{equation} \label{eq:fano-inequality}
  H(P_e) + P_e \log |\mathcal{X}| \geq H(X | \hat{X})
\end{equation}
where $P_e \equiv \Pr(X \neq \hat{X})$.

To prove this, we can introduce error random variable
\begin{equation*}
  E \equiv
  \begin{cases}
    0 &\text{if } \hat{X} = X \\
    1 &\text{if } \hat{X} \neq X
  \end{cases}
\end{equation*}

Then we expand $H(E, X| \hat{X})$ in two different ways:
\begin{align*}
  H(E, X| \hat{X}) &= H(X|\hat{X}) + H(E|X, \hat{X}) \\
		   &= H(E|\hat{X}) + H(X|E, \hat{X})
\end{align*}

\begin{itemize}
  \item $H(E|X, \hat{X}) = 0$ because $E$ is a function of just $X$ and $\hat{X}$.
  \item $H(E|\hat{X}) \leq H(E) = H(P_e)$, as shown in Equation \eqref{eq:conditioning}.
  \item $H(X|E, \hat{X}) = H(X|\hat{X}, E=0) \Pr(E=0) + H(X|\hat{X}, E=1) \Pr(E=1)$
    \begin{itemize}
      \item[$\ast$] $H(X|\hat{X}, E=0) \Pr(E=0) = 0(1 - P_e) = 0$ because if error is zero, then $\hat{X} = X$ and so $H(X|\hat{X}) = 0$
      \item[$\ast$] $H(X|\hat{X}, E=1) \Pr(E=1) \leq \log |\mathcal{X}| P_e$ and 2) because $H(X)$ (and so $H(X|\hat{X})$) is upper bounded by $\log |\mathcal{X}|$.
      \item[$\ast$] Therefore $H(X|E, \hat{X}) \leq P_e \log |\mathcal{X}|$.
    \end{itemize}
\end{itemize}

Combining these results, we get Equation \eqref{eq:fano-inequality}.

Because $X \to Y \to \hat{X}$, we have that $H(X|\hat{X}) \geq H(X|Y)$ and so
\begin{equation}
  H(P_e) + P_e \log |\mathcal{X}| \geq H(X | \hat{X}) \geq H(X|Y)
\end{equation}

We also have that if $X$ and $\hat{X}$ are independent and identically distributed, then
\begin{equation}
  \Pr(X = \hat{X}) \geq 2^{-H(X)}
\end{equation}

# Chapter 3

If we have $X_1, X_2, \dots, X_n$ i.i.d.\ random variables, then not every sequence will be equally likely. However, for large $n$, the sequences that are *actually observed* are almost equally likely. Specifically, $p(X_1, X_2, \dots, X_n)$ is close to $2^{-nH}$ with high probability:
\begin{equation} \label{eq:p-of-typical-set}
  \Pr \left\{ p(X_1, X_2, \dots, X_n) = 2^{-n(H \pm \epsilon)} \right\} \approx 1
\end{equation}

This is another way of saying that if, for example, we have a binary variable, then the number of 1's is close to $np$ with high probability, and all such sequences have roughly the same probability.

## 3.1

Expression in Equation \eqref{eq:p-of-typical-set} comes from a theorem referred to as **asymptotic equipartition theory (AEP)**. If $X_1, X_2, \dots, X_n$ are i.i.d.\ $\sim p(x)$, then
\begin{equation}
  -\frac{1}{n} \log p(X_1, X_2, \dots, X_n) \to H(X)
\end{equation}

We can prove this using the weak law of large numbers:
\begin{align}
  \begin{split}
    -\frac{1}{n} \log p(X_1, X_2, \dots, X_n) &= -\frac{1}{n} \log \prod_{i} p(X_i) \\
	&= -\frac{1}{n} \sum_{i} \log p(X_i) \\
	&\to -E \log p(X) \\
	&= H(X)
  \end{split}
\end{align}

**Typical set** $A_{\epsilon}^{(n)}$ w.r.t.\ $p(x)$ is the set of sequences $(x_1, x_2, \dots, x_n) \in \mathcal{X}^n$ with the property
\begin{equation}
  2^{-n(H(X) + \epsilon)} \leq p(x_1, x_2, \dots, x_n) \leq 2^{-n(H(X) - \epsilon)}
\end{equation}

It has the following properties:
\begin{enumerate}
  \item If $(x_1, x_2, \dots, x_n) \in A_{\epsilon}^{(n)}$, then $H(X) - \epsilon \leq -\frac{1}{n} p(x_1, x_2, \dots, x_n) \leq H(X) + \epsilon$.
  \item $\Pr \big\{ A_{\epsilon}^{(n)} \big\} > 1 - \epsilon$ for sufficiently large $n$.
  \item $\big| A_{\epsilon}^{(n)} \big| \leq 2^{n(H(X) + \epsilon)}$, where $|A|$ is the number of elements in $A$.
  \item $\big| A_{\epsilon}^{(n)} \big| \geq (1 - \epsilon) 2^{n(H(X) - \epsilon)}$ for sufficiently large $n$.
\end{enumerate}

## 3.2

Sequences in typical set $A_{\epsilon}^{(n)}$ may be represented using $n(H + \epsilon) + 1$ bits (extra bit needed because $n(H + \epsilon)$ may not be an integer). We prefix all such sequences with a "0", thus making the total length $n(H + \epsilon) + 2$ bits.

Sequences in non-typical set $A_{\epsilon}^{(n)c}$ may be represented using $n \log |\mathcal{X}| + 1$ bits, and we may prepend them with a "1", thus making the total length $n \log |\mathcal{X}| + 2$ bits.

\begin{enumerate}
  \item This code is one-to-one and easily decodable.
  \item Although we used brute-force enumeration for $A_{\epsilon}^{(n)c}$, it still yields efficient description.
  \item Typical sequences have short descriptions of length $\sim nH$.
\end{enumerate}

Let $l(x^n)$ be the length of the codeword corresponding to $x^n$ denoting sequence $x_1, x_2, \dots, x_n$. If $n$ is sufficiently large, $\Pr \big\{ A_{\epsilon}^{(n)} \big\} \geq 1 - \epsilon$ and the expected length of the codeword is
\begin{align}
  \begin{split}
    E(l(X^n)) &= \sum_{x^n} p(x^n) l(x^n) \\
	      &= \sum_{x^n \in A_{\epsilon}^{(n)}} p(x^n) l(x^n) + \sum_{x^n \in A_{\epsilon}^{(n)c}} p(x^n) l(x^n) \\
	      &\leq \sum_{x^n \in A_{\epsilon}^{(n)}} p(x^n) (n(H + \epsilon) + 2) + \sum_{x^n \in A_{\epsilon}^{(n)c}} p(x^n) (n \log |\mathcal{X}| + 2) \\
	      &= \Pr \big\{ A_{\epsilon}^{(n)} \big\} (n(H + \epsilon) + 2) + \Pr \big\{ A_{\epsilon}^{(n)c} \big\} (n \log |\mathcal{X}| + 2) \\
	      &\leq 1 (n(H + \epsilon) + 2) + \epsilon (n \log |\mathcal{X}| + 2) \\
	      &= n(H + \epsilon) + 2 + \epsilon n \log |\mathcal{X}| + 2 \epsilon \\
	      &\leq n(H + \epsilon) + \epsilon n \log |\mathcal{X}| + 2 \\
	      &= n(H + \epsilon')
  \end{split}
\end{align}
where $\epsilon' \equiv \epsilon + \epsilon \log |\mathcal{X}| + \frac{2}{n}$ can be made arbitrarily small by an appropriate choice of $\epsilon$ followed by an appropriate choice of $n$.

Therefore, for sufficiently large $n$, we have
\begin{equation}
  E \left[ \frac{1}{n} l(X^n) \right] \leq H(X) + \epsilon
\end{equation}

Thus, on average, sequences $X^n$ can be represented using $n H(X)$ bits.

## 3.3

Let $B_{\delta}^{(n)} \in \mathcal{X}^n$ be the smallest set with
\begin{equation}
  \Pr \big\{ B_{\delta}^{(n)} \big\} \geq 1 - \delta
\end{equation}

It can be proven that when $X_1, X_2, \dots, X_n$ is i.i.d.\ $\sim p(x)$, then for $\delta < \frac{1}{2}$ and any $\delta' > 0$, we have
\begin{equation} \label{eq:high-probability-set}
  \frac{1}{n} \log \big| B_{\delta}^{(n)} \big| > H - \delta' \quad \text{for sufficiently large $n$}
\end{equation}

This essentially says that $B_{\delta}^{(n)}$ must have at least $2^{nH}$ elements, to first order in the exponent. Because $A_{\epsilon}^{(n)}$ has $2^{n(H \pm \epsilon)}$, *the typical set is about the same size as the smallest high-probability set*.

We will introduce new notation:
\begin{equation}
  a_n \doteq b_n \Leftrightarrow \lim_{n \to \infty} \frac{1}{n} \log \frac{a_n}{b_n} = 0
\end{equation}
This implies that $a_n$ and $b_n$ are equal to the first order in the exponent.

Equation \eqref{eq:high-probability-set} can now be restated by saying that if $\delta_n \to 0$ and $\epsilon_n \to 0$, then
\begin{equation}
  \big| B_{\delta_n}^{(n)} \big| \doteq \big| A_{\epsilon_n}^{(n)} \big| \doteq 2^{nH}
\end{equation}
