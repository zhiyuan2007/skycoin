package visor

import "github.com/skycoin/skycoin/src/coin"

const (
	// Maximum supply of skycoins
	MaxCoinSupply uint64 = 28e8 // 100,000,000 million

	// Number of distribution addresses
	DistributionAddressesTotal uint64 = 280

	DistributionAddressInitialBalance uint64 = MaxCoinSupply / DistributionAddressesTotal

	// Initial number of unlocked addresses
	InitialUnlockedCount uint64 = 25

	// Number of addresses to unlock per unlock time interval
	UnlockAddressRate uint64 = 5

	// Unlock time interval, measured in seconds
	// Once the InitialUnlockedCount is exhausted,
	// UnlockAddressRate addresses will be unlocked per UnlockTimeInterval
	UnlockTimeInterval uint64 = 60 * 60 * 24 * 365 // 1 year
)

func init() {
	if MaxCoinSupply%DistributionAddressesTotal != 0 {
		panic("MaxCoinSupply should be perfectly divisible by DistributionAddressesTotal")
	}
}

// Returns a copy of the hardcoded distribution addresses array.
// Each address has 1,000,000 coins. There are 100 addresses.
func GetDistributionAddresses() []string {
	addrs := make([]string, len(distributionAddresses))
	for i := range distributionAddresses {
		addrs[i] = distributionAddresses[i]
	}
	return addrs
}

// Returns distribution addresses that are unlocked, i.e. they have spendable outputs
func GetUnlockedDistributionAddresses() []string {
	// The first InitialUnlockedCount (25) addresses are unlocked by default.
	// Subsequent addresses will be unlocked at a rate of UnlockAddressRate (5) per year,
	// after the InitialUnlockedCount (25) addresses have no remaining balance.
	// The unlock timer will be enabled manually once the
	// InitialUnlockedCount (25) addresses are distributed.

	// NOTE: To have automatic unlocking, transaction verification would have
	// to be handled in visor rather than in coin.Transactions.Visor(), because
	// the coin package is agnostic to the state of the blockchain and cannot reference it.
	// Instead of automatic unlocking, we can hardcode the timestamp at which the first 30%
	// is distributed, then compute the unlocked addresses easily here.

	addrs := make([]string, InitialUnlockedCount)
	for i := range distributionAddresses[:InitialUnlockedCount] {
		addrs[i] = distributionAddresses[i]
	}
	return addrs
}

// Returns distribution addresses that are locked, i.e. they have unspendable outputs
func GetLockedDistributionAddresses() []string {
	// TODO -- once we reach 30% distribution, we can hardcode the
	// initial timestamp for releasing more coins
	addrs := make([]string, DistributionAddressesTotal-InitialUnlockedCount)
	for i := range distributionAddresses[InitialUnlockedCount:] {
		addrs[i] = distributionAddresses[InitialUnlockedCount+uint64(i)]
	}
	return addrs
}

// Returns true if the transaction spends locked outputs
func TransactionIsLocked(inUxs coin.UxArray) bool {
	lockedAddrs := GetLockedDistributionAddresses()
	lockedAddrsMap := make(map[string]struct{})
	for _, a := range lockedAddrs {
		lockedAddrsMap[a] = struct{}{}
	}

	for _, o := range inUxs {
		uxAddr := o.Body.Address.String()
		if _, ok := lockedAddrsMap[uxAddr]; ok {
			return true
		}
	}

	return false
}

var distributionAddresses = [DistributionAddressesTotal]string{
	"2eUvJEQWRkhwK7h7i6pbirZH6xSdTCvZfcd",
"YBNf5zLG3REjBMPqAXgcw6yTe5dkTWgrkx",
"fyJ5uzA1ngLP4SRkpACdVgAMBAMHKzPxoB",
"PQnYEj3BYrxMURW5CaeNc16NwEYbGwxD4A",
"iW7AyhfoZ1TUTz3fB9LvgCP6TSax4rLGPg",
"Pe79XLw6bVxhNPTrXjntP1F9YvUwftqcAi",
"2M8Z1CkT8A9HNKLokkaUue5yfj26XL3zQdu",
"2UygEscSqYbMk7jvcnezojanraE6yij57LA",
"9vQZH3eXCXCyGby4VU1aqzGzUkrrjDe4WQ",
"9KtDPZPzp57F9Xr79Lbko1zUGL36PNFSG8",
"2JoFWDt3s652sazEbnKc3U1n9PpwiP7VmKK",
"2ZQAv6A992wamjhYQdQ9JfHixvhLSMyZ3Xk",
"2MpNBbmTGr3tntuAYFk49685agCPVB8KcCs",
"cChFLCcEXJkbodtBqaU2YVmDWtkeXSoXMt",
"2NaxY7EVD47MT2Q8E5bp2qyVYTpq2Wov3gv",
"2A6KPwSRXDJZ3xANoUR88fGSgoqLvq8dQmJ",
"FF8EcekJ3ssP63ApbfCttyt3cY2XoDmU5w",
"oqpMGSL2z3Ai67BYnAmWpYgBYdLWsfJEKK",
"wLwxFiwRPj6xH2yGGjW9Rf1B553aAjJ7MT",
"XFuQ4P7V3e1X5M4wuCUU8MSks88Spduz3K",
"8L1Qzd8USZgHC4VpLFC5ii6APtWGkqR4b2",
"7MVQNvhxVTveckRFTPN6UBN14Ku37etmmZ",
"9F6UVC4v7XSSBUU1CnRyXn1iNBhubwxoXM",
"2LXyzm8rdBke2r3hkSCLGPQPvitGbyjRBe7",
"24wxAVvb9zGEHnEXa9YGm4391MgvYBNA7QH",
"29dkTeGRi1u88sSNJhwrDYeazbuiBnpPn8R",
"2JP6nnWHFpcfYcsJGqMcgVhCMuYVKFtKqgY",
"2PJejwQLT1qJdYXmQSM26Es8XQjJmK4ZUe5",
"2RRMfhCZTN1rWphegxd2zCdkXEpnttREW95",
"2j8GjvJxKH5ZdEsm3wL3DnZMe6GUia9bvQf",
"EtrV7pqWQGc3yfpsLy5tRF9GHziFDjDHm8",
"2cYtVkZaue7XYbm85ijSpWVY4fxVD1uUxAi",
"f3zxgZDFoJetfWp3ikYzCJDR9M4UhVkW41",
"2edZdMCBPjmuHnRVjFTkCGGu7VyorryYjtt",
"4Cu9FRncNXQLspi7xNUEpVdU3aRpsPpZXh",
"2HmrA2h9RxwrdLdGAp4RnXieu5Hn6wnFE4R",
"im2ProCs2ApNiBcrhzkb4yGy8qecEtKZNR",
"h5fFrscpsEy4hM38y84kcU5zLUaoCU983k",
"2TLtawHDCXHD6Tgq8Q3bssjEEDoWX6yBsDU",
"psC8o4NBMrGHhJQLM15aA5xWKWHfzY74E2",
"23G2rnHGHQ81d8azJxynQQtUhphpGbJB3r2",
"2EHasA95JbMH1NUBPSfvgonax2ibt5zzD4n",
"2QkddHMW5g5RL9c8vNL1AWuVvT8NjXrLLB",
"2mm27VNFEbjCPgKTToh7SZdATTBqde2k2Qn",
"2XqBFjwyNRXWxZdKJNAQntq5Lq9ZPzffVxu",
"2A1uUGZycH52FV1824m5ATnqgEPvxBTd5KQ",
"WyBAgu6BetP7VU2Jt89CswkHMYt2YjakBJ",
"2CahDHnFgS3bs9nZvkygWqtGJQdt67ALchN",
"2PjbsjwLNpXEuozWHYKQHZ1YU4qkj4TMkVL",
"p5BuZNT9jBmQ2hH3oPFiUPsvHs78HdKHUT",
"2beHJ7Zw7pxbwG8RtPvkRoDkKiFWY7Swh5W",
"eBadce84eyDirw2fU1qNb1NLn6P2Kbg7jf",
"2e41swD2UMSU5mwjbEP5j1pwKRLXWj3bgHv",
"2gsmxwUvejfen1PPwLhfg4LQjDJ9ye61cSQ",
"vCZqNyduw4svXnzwNuTVUExvksW7aHBaXt",
"DyxDoTjuo5HNwau1uwTLwz1DNeSKQXT2Gp",
"25zaKqhRXYYCeW1d3EGRP5NYZbeLuqbU7ZK",
"2fYAMrXV4jXkQhcvCBoMGr9rG9BvqavLzVx",
"kpFNimaktYinrM7DGhNpWAW6emmDvpNbzG",
"JjbZgcVTJVB5ZLmeEjBFHQywPpY8bJ7evR",
"gXE4tbNVoVE4vPsrEDeakAefcRL5ovqoBa",
"2jsVs6judvCAGjYJyiXGW4muhVeAZcBwBKU",
"n7ZXoEyN9Ny22gdRTCcaVqLp2sMqhpnWD3",
"2Uh6wHLFupwe6Bt3bJzwcEQW1PNtdYskmvw",
"28frRpTEhRne4TnkA6SQz7miakfrBnUz9kh",
"2T84qCYoJEgRX6oLqUwYP44RxFrdR6T8KRB",
"2CFtRhJmhRWjJPYCrFeUZVEPbtpzjwo4Ciw",
"27D7u9NnoZ77ymhJxmFHDBJ6bCbd7Yg2eGP",
"4g3HCsMSJCUEhhmwoeyWZPdcCrhYYomzia",
"2cy2Bq5nX917nA5Cucti9fWxZdaQYD1gXeF",
"6vv2GPMBonKPZGKBVBvmxP1dQ9mkJe1cvJ",
"2dUTP5hSkKk29aw7iJTgb5wMe26r5dXB5Wd",
"d3nZFSFVBhoY2zouj41KZ8QoVEm7FnondX",
"2RL6STqWifBeHrkmnGvRygNH2hpJAKujC6b",
"ZXjXehSCmVye1qzGafmj18gJXrD1nvr3PG",
"2iekXHcEtEx6kJTbST1eFUmiGZpJkZ8ikcR",
"o5biG16HwBT9DLYdP8mNyU5jwr8NZddWNN",
"2VimEV5X2c29gxb46WZ4Jhni5nxZL9TCkWJ",
"96xSYvPKKecYJ4k2obrwq5GEPsiwHcbhkD",
"sddnscSs8FH4Y18qdjH1fqeqFbwFbaendc",
"YVPKpYvNsq5K7qRTeeARcMbUh9evzg236i",
"u6rZVw7BdDJn6vrFxSWcwpxq3arTpsLXVK",
"44h5tkv1hocoDkq4dkZ724zzmbtNckESqh",
"6FN5eQhd9BffBMZuHwQvMJsDargUd1uHUn",
"oFWigdejMsJnFiLBTebL3t52DUSyu8x2Cu",
"EMLkBU69SG1EmqRpbSqSpbsjNJSiUDBE7s",
"2UkixssRi9qi7xD84cmD5Ahw7npAvbSpkJf",
"PUTn1JHcaVtLdeJgsN6FjpCTHPwee9w1Cb",
"nvcrenJFJ43KZVKLgcusP8MGSvH6WFhxb1",
"YhWQqbVFfJJX1Mj7nz9f9f2m5eR9Acrtcu",
"2DgvgNF82CMKtXbSKjz6VjMb6mevHbFaBcq",
"RhfmhjMLuQV3rD5SHvLWdHoKS2ZV9LnZoX",
"2Fu6QWvBUiorTmJPsRSFaT1oh4d35fpsHYg",
"2QZJCzuYRugMq4AZRAPvFpLQX6BviVexnxy",
"N8bLJfRBLYA7psCaAsqbDYeRW1jV76SSrS",
"2bH8PqRvRfbxDvxJSKykHBZ4p8yDQ7iUULS",
"2MUL3eCAKFA9K6qLGZgKgnDeXxi9JLrbhFg",
"EG3RNZw7xpEQqDTYb8aqBC6GjtqVQNw8zq",
"2WYJdECnQNoowanGyqPkC8SrYtKdaiimj3f",
"onk5iudaJBZZirnjnRchvg35Bpgjocda5v",
"2J4RZWsEpbebQMG5oqezs6STXVdpXswLdfF",
"jDXDZHN1rGv1KnsE3nMuohEt3ejZLvPK14",
"2mWZfeX17Kw7VikCioTWyXLoS2AmfK6dYha",
"7cuPjSJUh7SrnYVK9YqQJihCGFeLeMAqnd",
"2Jarjfyhg1xihn7mFSydP6ehtAAFwARwvDD",
"2Pivhy8nqWz8hv2T9NUnUUf8MyPKSAin8HU",
"s2AebLUiC3wNMw2MPALDP9RB72r5NguttM",
"2XZsRhEgVxNtd5az7Ldt9aW53mNQ3dggFV9",
"bg2p4SChtbtGvtXcJhyAGY8fq2TydSVH2b",
"2NkSs4epueWQCoJZDRPi1VjzymXXSpjUKHk",
"hZzFNz7m2AdZT27RyYYQxVKqbhv7bCGhua",
"bCq8hteeZCRK9UZofedkY6hb53e88Wk3Hs",
"raycKTSfRJH9GWEhmTG1YozbshirHrM595",
"C5DG6Qh8EkmmfhRF9byjnW6pkdbw3bAmeU",
"8fZ1GY4QA8Rk3pL6MzZu4H2QZjgjnN4NAq",
"veaYVD4cvaAug5JoTZBGSnUX87qqjqorus",
"2chHJFLJtbr8cgBGBTu9pzmvpbMEsVAg4sK",
"2P1ZPhqpYVaG585xpHk68n8JWbiJSfMWosb",
"Hu4XXu4ir43qKEgTtsnNtifTab3piRhxj2",
"2G7ahNsWjfR9NKEpuytzu8pLE6LJqiUWEcD",
"2HZLCpyUVZLxHrWRpJV5bSWAzi29JcdRJ83",
"F7zUb7kXhzPqMW94mYUhYRUjXHfkgU5qL3",
"22RXxqFhRZyS5wFhSF8K9ningZRZkPeHSDp",
"2iDzi7hmaPuifsFCVtyE9UEFrvX2rTsbrMn",
"BjCmxbjbndXiUDXP51kgLW7d6A2KT6pvRF",
"4WaVojXKfX2EBanJtHxECEYE29ajW7VWxV",
"c37Gk8Hc7nNrTwXBj7rAWsqCyz9JaNSK2",
"1GMryH6qvJnTpdmMpeGcXZ3wPvtrDktnsi",
"2ZFoeVmhFpM56Ekq6vvXN6ehYMYXSGZGDbh",
"eqk74TWP3LAGButbvqUJbRTEyQvzyBUH15",
"bTNDh7oRoYHbcQ8PUms2mwtw42cEvzm4tg",
"2gAsCwcbU3SNpv1GB3SB5WY1HYABGZEjSyJ",
"2JJ3RjghDgRAWEGmquQLcCKDNcwdygxKShN",
"7myLsX2Ev6NRMHEHbbWFf6JjK418NjkXMm",
"PMRLhiN9kfqPHCLP6ah8w3tkQBTE1dEbai",
"27Pf8TZJBJhjQNxhSonTfuj5aBfAfkxShF3",
"2T5VrGHjLJ8R6XBBvhuVRifLBwBcEZMcmWS",
"Q8jTdKKwCL2hQgSWea4YkKHh6M9yqCmfH2",
"2m5QjFqMroAG8JdkhLnGnY5uJcE8yZsc7aD",
"okv5r6dbL3Y2ap5a7RNwknfcZht8GWwHcr",
"2DkhKXKzspDBFF6St69Ew1EW3NGN4rTaBcY",
"27CoBkfoCDhaJUWSLeHpsC3ddJULqieeYz2",
"jL45E6T42tjU2iE2hjB8JZciPzsqnANG2t",
"6AgyxRKskW9ZSPuBKU8ZnCrMnTig9NqnmP",
"cQ45WH8Yastkr3EP75SAxa9E2WiWpgvgoK",
"Qt5CYphMjDpwGsA7KjFJ4fGwxFPD78E3v9",
"qWQcJ89Q43qJTyZPPumTnvkGhhebmdsg8A",
"2KM7LyWnNLPSeJNY5KEjeichbs7psdTXgfG",
"cArTAPxRXXYrSkGupgCp9c4Q9ZyEp2DpG5",
"9kznGeWHiCri5wC5FwJDvJC1WZ6BMW1KU1",
"2XpH9iW9hsr4YmA9R7fzVoeP78yiYkmisHt",
"9Qm53uDCQEdiSYdGBzptmXR8SWNChFoJ7k",
"iq446vsZhSsfqBgccz6mUVARu4bq8LF6Fn",
"ZnF1YtJ2UAnf45QS1azTKtsq9ynGVrLhog",
"WGsSocm1ogKYg5AsKx48ggqu3a8eKbJhBP",
"yNBYe4pBnd4R3uXtiQ1gi8Xiv5xVRrprxV",
"2Q18WhR96bADhvduZMUFin33cFswhXAubpG",
"gxcfmDpJmNkfD2tSjhKKyrReUd89kmxuBR",
"Hjvb9gdeZ6Md1BMHAPAknpgoKv1numuK69",
"2LFxAa5B8cuvs56MV8rykvyZKboQbdmHQd1",
"a1D7aSPpx48dHbedhJUNobYTSE4bMQbgkv",
"wHTubHsbL7GTAquSCqRwqiV5UH6J3jTwwd",
"scW28zD4JPa18gqC3yQRaGJfrtEWb5Gv7x",
"zjby8UumoMPSmVfofnqZNpJXgiBpPMaDCK",
"QyvdBSa1TFxJTC6eHjzCuMQJcWPy96RdH6",
"EaJ2SbTBstMjXHccPkGb7tzgjsAqKpCVvq",
"21RqkXeVQwAtbtpBH2Q5rMvHpC5i4Yjxm1f",
"2Hi6EoNoNGymAujUuYTbvADMzBM6L1mbA4Y",
"2ChHf33ZWtQNsYKbKkthk4LnwriTfJjajS",
"2HEq8TGFx95oZUuDEX14yK7auSRcWCQoQ6b",
"2Bhn69yDVrk8V61z7upEW3vh1GSsXMdCjE4",
"2hxRhvmsRBsFLYcaiaXn5brpLX2aBdKDaz6",
"2N6QPjo6FVLRdSV1z9n1QLiUZLeDBeZ8Zn2",
"8U61vNvzfEV8kp5vWMXu8jV4ukuqkhryzz",
"LkwrVNeH2MUvz54z561mLyyNo2vM5ePHaV",
"J6TbqQhjBmd1xhatRxjMcdtHeEVNAMR5VG",
"2HHAqdf2YS6AU2EnUxmWkMfzWMUhvUAy2Ah",
"mdEV2LK59fzC94nHMju9GBUBmC9xHxBJAP",
"2bW9PP6LhKmLYXTAorQXMLsAGUiT1o1ksQK",
"JRwFf6dYxxYby82GeqbUG8uqudoCcUjR9d",
"9VJRezsV1mNkA7qLRSnanqQhW6j28n356W",
"2Wia5nLxzgwT6oAdVyv9Dt79o5gb5t7ET97",
"EYsxvh5nmRKjginPKuEofyvtTfRwrgGMuH",
"SupWJ5q3QktKW4i1Aat6DantP9iDPU6QdF",
"8bQdM7ypE2DQZnuSd7Ui1Vde11vUnY1VQz",
"CJv16y93qP33Rkheqw3LF6QFEKHokgZuuF",
"2d9ivUERN9LR1iDei2EaU2V8b2EqKUhjzZJ",
"YQUWd72sYq7q7eM5dvYzMa4UG6PPbUjnd1",
"HeQHgoNcUf68CYXijGC2H6Rtij45TUBwQ7",
"wWaexWaCR7Wg8u1BtYaXMVQTWhQFxHJAZn",
"8imUx3EvLKZABubciZjfFSQqFp2axsGSJv",
"tgggLXwNAC5VkkMLghyTRNbc49nLDGHc8P",
"2BBdmvGPoEoJtggudAALPPwNYWcN82xbnXw",
"at3nH59uFrXv7PyqiKeAp2sYfaVMmtPwMK",
"2WQ8PZtQ2HfJgKfztAbeynzwP8kYUgAAjaT",
"uRuhwyG5V4EuD1LzV3m3JVuCSgLxFEKzJL",
"2dCBte3qWybhbpmUHScQADNB3Q2nF7tDznL",
"29EtDFfCFPV17BoKAtrNaJUePxCRhxDDRQN",
"2hKFZzkUaRam3V1XXdkEG3aTfyfVP4dyak9",
"TY3CBpXXx11TcahqLuq6UjxWrthXPxQtsa",
"f63WKULKZJrD4YsFDozKynwWE12fpXRyQy",
"2PFN8SkigCLbvfLaHi659dn4mBYjNsz962N",
"2ie4Aat75Y4zJdWkXoqxHuSNHFFppzHL2zq",
"2eNFJTHAeP7jEsmLmby8ozsF9FkoDDccEdD",
"2LxDovYaPbKgGx15BhLbSuLhqH8utWHNsRw",
"gHCUzuY76rF3qC2ZcAwX8JLoRZH6HsRdmG",
"2EkV14HtXGJSnQEWvRJvwBEdDtzUGNdmHuN",
"zd3pk3wF9PA2q4LzL6bHHSjFVSNcXGiLZ9",
"aucz4pr7shBet6Rpg3v4uW3fYGPgRuF1FD",
"9zkor4GWUic2UYsA1UtivCfCaC7uWavRkS",
"bLDnMH8zgCNazzxeBZtHzjWEgpbp8yyjhp",
"2TYrsYsobWEAjnf1L8zEMB1J9n9HBHzmjsJ",
"N3QRmyvTs6Nw1ypfGw5jpXtxVtRLiRSNcW",
"2P4UusxANNc2FhPCs8JjeBAqAzxNdakR8qe",
"n9zTARCCdYAyiAFJ9gvG4KuxQ47rXX5agA",
"6Lb6scKzot7JyV4VpQmb1EKmZZouWg3zYD",
"2BxF1gkqzhCZKeuSZwss71v7a7o1fq1z4un",
"bZ4wiGcDmjv6PjB5S4utGBU4ux4pnKT6Sy",
"2mavF3PbFhMKxdnBkiqyKcbrsWGtqnrLBEM",
"9ad6nA19WyGenZiu96i9fNTWeaci8Jh2qy",
"KxcMGjidVuFc7BmTBY1y3pfSH2UDMUkUGD",
"28HQbW7qwrTqF7Z9xGsXQVAijCjDpTeZEG2",
"KcYxY3hzQLnwqQ76PwfH9LmQpEvtoaz4Nz",
"ZbPRsHuYsB4KMtRCG2qn8gd34MuzweEvw3",
"2Wr76wMg81iJzXiXGLzdJC9wqzrtfiFu9Tc",
"2hAxP3GYQytcfh6tTuSMa8DXnD4GvgKPP8B",
"Cb6NT7Q6PoEWqyyYtLEPmQofPmgyRjRAnT",
"cSKAgDGNFa5sn4UV65PPtF59UN6jvT4ox3",
"2593QAmyck4XAe8noSRgYxnoF4vapmSmerq",
"CXjFbYPCvaS9vZXg46rC2EeXbZQhWxvxSa",
"2NX5kqcUvUi5qSTChrXWW3oq8SVeSSDuRPk",
"qkWNpGcyE6azpjwEvZY6Aw5wxrJpJdrvdP",
"G4XfuaGyLr3WYAWShj5yQpq9SMm7tjGDS7",
"yNnjX1ufQ5KnkfYFhwXiTWFymtCKSuz4yt",
"j2TbuZy1rtwPvgdWjPgmSsxunLvD9SMkXE",
"TLF9MX4ucqUKWAahonio9BFd96uHVjgXHi",
"kDqCkFXjfyFZhBvmpvExKyf73aqhY3RPjx",
"TTyYiRhFXu7PMfzNgGbhU3v735CJwvgYtF",
"2WNpHUKQzDbfzEzVGXDoUm3ap2GZVXxDXgr",
"2SRfQbiWEMBAav6LufN8tsxtfhamiBFd4YZ",
"2gNX63EKDXs6JDvriCuZgLSMGZNHHgeN8Up",
"2NFJRLn9MwCPx11DRVo1twrAZJXvr8PBGLb",
"tDZJSYMoBj4LtcHWKSZoL3rA9sg1vuydRL",
"BQoCLgobiKsVygEfT3YZBEm9hrF2WF7ivo",
"2HsLWDHWKayEmkvBvRWpfAMA3HBWPK8DsHd",
"2fY45XocpvUtB41qBYVYNnyeZbfANqoK3rz",
"YAGMLjbHdiBJzPkB4Gr7y5hJGwjfaMESNr",
"2bVwgdCXV34MBY7sZbHdrHjeek5mpXJCL6h",
"2D13KRCtv6GpRAz3M3v5QsUrRgYLiDt8CLu",
"277U76G97uQnVgjrk49NFgWdpjiSgiHQJCa",
"MuRMEyLZ6tUyF2mVqdhn8kpC8cqbRf3iah",
"2GdYtoPXLG7JA3NtDbZeWrCQQXAMERMnKey",
"2SrvAQGWK84c9JSkUp7zDYXRdrrotpaLqC3",
"2C5nAeQ5HsJBBgewDLUQsFMcvQbhX7sXrBv",
"pr3vNsfytBafdRGNA6XDXNjV1oB8UNtkYF",
"vZWcSzmDKwEG4s2DrWEL6rzzDnyWN8ZB57",
"N22YVGkty3G1NSAiUsxFS9XvRnQaGMNSw7",
"r4fwFR2NLF3kcbTqkoSdVMqdhbPC415BW5",
"2UPgitR3MayCEbyhgozgCkZFBL7zomGC7xB",
"21pDonsYTQ2xa8cgBrP7dH7XDgNZb4VjZgQ",
"2VJiqhUYbcSYaLN6bTq4P54KHD1XSwzjKuS",
"ptfxS8xCjUdSP6izmqovwA5fAM7TpX6Tmg",
"DxtTAwkoVEPGgtV2RT3j9cxYVAHGcSCrZH",
"24mnbdrZN9dPpkF4LE8BM3JbY8v4LSPHaV3",
"2EWhUsBcZ3va5Gn7Jh4SYgJVjSajPjzDUiG",
"2ajAo78tTBz6u2vQ92VDeqMqsjcENu8gm4U",
"XbpYNaqHQtzJ3TTVZKJzuWcMH46z6F4aCE",
"axwvuMMQhCgtLRPLFwKjVUCky7E4TMnL9r",
"uu9ZRtTzqV2XJQfx2XdP5vY59LbkC15YCg",
"22H5NzwKQ5HFFXijqtcPpJnAv3xJfKvAqZD",
"tKFFoc4YjgbPw45wLh8kdupKJ9SMnjE5wN",
"2Bnu6GDmfc4omgeyjEAjqiUwHGzM7PB8vr3",
"kHZVaoksq28BLJmiNsweXHzqPyj7u41Q2p",
"e6ML77rss1Wpw8QtJ8TXZCY6q8CVoPnLC2",
"2Ri7bPmgjhHa4kq8K389ew3ifth3H8pSabs",
"Ryx6Frz8cAaSESywpYLv9KjgZ5twAtY5Lb",
"2k51mTgkKYRTGnsCXy6Dr59dRqtYUcidLzT",
"2GYiDyHdr4nZVLB8P6w68b2ikKVDuFhZWRM",
"MURdRTPkZZmNDvvWHnJzz9PUMT9UpjPWp6",
"29FtHdtoLKZe7dHRNKd2zrQGr1Z3pJALi8V",
}
