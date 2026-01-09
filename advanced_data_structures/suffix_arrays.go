package advanceddatastructures

// buildSuffixArray(s):
//     n = length of s
//     sa = [0..n-1]
//     rank[i] = s[i]
//     tempRank = array of size n

//     for k = 1; k < n; k *= 2:
//         sort sa by (rank[i], rank[i + k])

//         tempRank[sa[0]] = 0
//         for i = 1 to n-1:
//             tempRank[sa[i]] =
//                 tempRank[sa[i-1]] +
//                 ((rank[sa[i]] != rank[sa[i-1]]) OR
//                  (rank[sa[i]+k] != rank[sa[i-1]+k]))

//         rank = tempRank
//         if rank[sa[n-1]] == n-1:
//             break

//     return sa

import (
	"sort"
)

func buildSuffixArray(s string) []int {
	n := len(s)
	sa := make([]int, n)
	rank := make([]int, n)
	tmp := make([]int, n)

	for i := 0; i < n; i++ {
		sa[i] = i
		rank[i] = int(s[i])
	}

	for k := 1; k < n; k *= 2 {
		sort.Slice(sa, func(i, j int) bool {
			if rank[sa[i]] != rank[sa[j]] {
				return rank[sa[i]] < rank[sa[j]]
			}
			ri, rj := -1, -1
			if sa[i]+k < n {
				ri = rank[sa[i]+k]
			}
			if sa[j]+k < n {
				rj = rank[sa[j]+k]
			}
			return ri < rj
		})

		tmp[sa[0]] = 0
		for i := 1; i < n; i++ {
			prev, curr := sa[i-1], sa[i]
			tmp[curr] = tmp[prev]
			if rank[prev] != rank[curr] ||
				(prev+k < n && curr+k < n && rank[prev+k] != rank[curr+k]) ||
				(prev+k >= n && curr+k < n) ||
				(prev+k < n && curr+k >= n) {
				tmp[curr]++
			}
		}

		copy(rank, tmp)
		if rank[sa[n-1]] == n-1 {
			break
		}
	}

	return sa
}