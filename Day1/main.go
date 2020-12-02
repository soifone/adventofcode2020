// Day 1
// --- Part One ---
// After saving Christmas five years in a row, you've decided to take a vacation at a nice resort on a tropical island.
// Surely, Christmas will go on without you.
//
// The tropical island has its own currency and is entirely cash-only. The gold coins used there have a little picture
// of a starfish; the locals just call them stars. None of the currency exchanges seem to have heard of them, but
// somehow, you'll need to find fifty of these coins by the time you arrive so you can pay the deposit on your room.
//
// To save your vacation, you need to get all fifty stars by December 25th.
//
// Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second
// puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!
//
// Before you leave, the Elves in accounting just need you to fix your expense report (your puzzle input); apparently,
// something isn't quite adding up.
//
// Specifically, they need you to find the two entries that sum to 2020 and then multiply those two numbers together.
//
// For example, suppose your expense report contained the following:
//
// 1721
// 979
// 366
// 299
// 675
// 1456
// In this list, the two entries that sum to 2020 are 1721 and 299. Multiplying them together produces
// 1721 * 299 = 514579, so the correct answer is 514579.
//
// Of course, your expense report is much larger. Find the two entries that sum to 2020; what do you get if you
// multiply them together?
//
// --- Part Two ---
// The Elves in accounting are thankful for your help; one of them even offers you a starfish coin they had left over
// from a past vacation. They offer you a second one if you can find three numbers in your expense report that meet
// the same criteria.
//
// Using the above example again, the three entries that sum to 2020 are 979, 366, and 675. Multiplying them together
// produces the answer, 241861950.
//
// In your expense report, what is the product of the three entries that sum to 2020?
// link https://adventofcode.com/2020/day/1

package main

import (
	"fmt"
)

var items = []int{
	1844,
	1123,
	1490,
	1478,
	1108,
	1120,
	1594,
	1101,
	1831,
	1146,
	1084,
	1535,
	1016,
	1722,
	1388,
	1188,
	1351,
	1477,
	1215,
	1678,
	1159,
	1558,
	1581,
	1400,
	1550,
	1306,
	1852,
	1745,
	1224,
	1896,
	1596,
	1005,
	1499,
	1797,
	976,
	1777,
	1129,
	1601,
	1058,
	1510,
	1704,
	1818,
	1795,
	1364,
	1276,
	1362,
	1801,
	1985,
	1421,
	1311,
	1855,
	1977,
	1613,
	1951,
	2001,
	1327,
	1872,
	1517,
	1040,
	1692,
	1595,
	1769,
	1956,
	1763,
	1470,
	1898,
	1366,
	1443,
	312,
	1685,
	1014,
	1409,
	1717,
	1105,
	1290,
	1703,
	1732,
	1939,
	1790,
	1164,
	1225,
	1474,
	1713,
	1413,
	1906,
	1488,
	1931,
	1702,
	1848,
	1668,
	1737,
	1614,
	1719,
	1647,
	1171,
	1508,
	1035,
	1133,
	1179,
	1180,
	1472,
	1414,
	1659,
	1820,
	1544,
	1697,
	1175,
	1033,
	1805,
	1576,
	1196,
	1597,
	1739,
	1291,
	1103,
	1457,
	1514,
	1462,
	1619,
	1683,
	1338,
	1604,
	1372,
	1653,
	16,
	1725,
	1181,
	1952,
	1201,
	1531,
	1505,
	1912,
	1527,
	1853,
	1905,
	1833,
	1913,
	1131,
	1069,
	1537,
	1754,
	1551,
	1435,
	1914,
	1093,
	1273,
	1726,
	1094,
	1439,
	1689,
	1607,
	1646,
	1588,
	1698,
	1539,
	1493,
	1352,
	1163,
	1482,
	1044,
	1523,
	1142,
	1253,
	1250,
	1986,
	1049,
	1330,
	1219,
	1162,
	1088,
	1100,
	1532,
	1727,
	1761,
	1107,
	1916,
	1220,
	1319,
	1098,
	1431,
	1260,
	1642,
	1269,
	1102,
	1432,
	1267,
	1824,
	1658,
	1149,
	1024,
	1963,
	1498,
	1904,
	1812,
	1600,
	1773,
	1283,
	1463,
	1776,
	1406,
}

func main() {
	r, err := checkForSumPairAndMultiply(items, 2020)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Pair multiplied are %d\n", r)

	r, err = checkForTripletsSumAndMultiply(items, 2020)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Triplet multiplied are %d\n", r)
}

// checkForSumPairAndMultiply takes a list of ints
// checks for 2 items in the list, that will match
// the sum result when added.
//
// After found such a pair stop and multiple return.
//
// 		items: 1, 2, 3, 4, 5
//		sum: 3
//		found: 1 + 2 = 3
//		return 1 * 2
func checkForSumPairAndMultiply(items []int, sum int) (int, error) {
	m := make(map[int]bool, len(items))
	for _, item := range items {
		d := sum - item
		if m[item] {
			return d * item, nil
		}

		m[d] = true
	}

	return 0, fmt.Errorf("no pair found")
}

// checkForTripletsSumAndMultiply takes a list of ints
// checks for 3 items in the list, that will match
// the sum result when added.
//
// After found such a pair stop and multiple return.
//
// 		items: 1, 2, 3, 4, 5
//		sum: 6
//		found: 1 + 2 + 3 = 6
//		return 1 * 2 * 3
func checkForTripletsSumAndMultiply(items []int, sum int) (int, error) {
	m := make(map[int]bool, len(items))
	for _, item := range items {
		for c := range m {
			d := sum - (c + item)
			if m[d] {
				return c * item * d, nil
			}
		}

		m[item] = true
	}

	return 0, fmt.Errorf("no triplet found found")
}
