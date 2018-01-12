package visor

import "github.com/skycoin/skycoin/src/coin"

const (
	// Maximum supply of skycoins
	MaxCoinSupply uint64 = 1e8 // 100,000,000 million

	// Number of distribution addresses
	DistributionAddressesTotal uint64 = 100

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
    "2Pu9b56hQQuCuaANmD4rjgZgBu7gToTFmpG",
    "2KdydCsEbvPedbkTUL8pULULQCA9UMJwKVM",
    "2dDzP785qgcy4aeRkNcrLHX4ak9U3STSoQ3",
    "23QwVn8Xev4kD82pXxqKsYJWtyZyuCBvP7P",
    "2J3727vEHerTzfHjPuTFdGmqT1iVziFA2W7",
    "cVZ4CGod8aENJJ77ifZgqmR9bsFML5WHDQ",
    "2aCRVW4rQqCKgx63WA1CrmwnoVT8o7oQpr8",
    "2dkCVrdF9xeuKQm6h3eXzsoHfjRue1RXNXq",
    "VcEVfAsfsMhdKsuR6EJbpvwJA1rPp9omsR",
    "2Np938RWZa96r6gm7PUGa5GDmz6ZTeRC77D",
    "d7jVrMrNz7QYzV1rHqWfR1bVgnPvd8eb8E",
    "2UdeJEsPhUonvEVnjmAkpsmc6hY4MRuRYiN",
    "ESSfPuBuchvGh3cLBhESeh3L1yAGhScYgp",
    "ESdndE2CTqVMVcENwfGe5kZ1QoX6tmXTw8",
    "2F1Zp3Kps43Bt6KrTKWp83FWHSnUbvKLK3R",
    "dnXmEebDmUUPVAeUUZYDkPdPJzxVvjepsv",
    "DkWHgqZcyw74wYtNuGzXMZE6vAd36BNR3r",
    "AvYQv4TxDEh1DCmFXbJbacCcHxBhriy53s",
    "2M2u1Xc9zocVfVL4dB2nJu2Ri9K9Y5tN1EE",
    "2ajuuwmaB4FLycXvRmrjQ4RjmLtZwUqGJhP",
    "2kHTHzJXYeFXgBFHHR1egm4u97ecgNiuUZf",
    "Hw39pCX9jmU8b827oFZH93grE6BQRPcmt2",
    "64sbxuWSw1ArZrDHPfLVouaYycCzsDhjKu",
    "DTfuMxbcJTW4s9NgdXoPnFBBpGHPGRBmWe",
    "2kfJsCYgQEhhy487Pzi4ULdQANHdZTk4wsK",
    "26oEJsgQiSX5fLV4E7k2VmNQ8U7PDKw8sUe",
    "kJYUSkwd8oSRxJHcbqXUqSpWpNxNqUECDD",
    "2A5atMu7zsZEeVpZNpJxPZ6nLpaAxU7QNt4",
    "iHve5p9TWo5R75PaPrDLi7tscHbpr2YDr8",
    "w9tWEkoVaDDJJuaAHm4uRmJW9iAvfSNsfn",
    "2VTt6yyEPR8qPr87bk2w155rFHDiScvuoS3",
    "Boa4yAoT2pS6RMVn1XNvDRfjThBFo9QsNo",
    "2SoEv4oLxGpjVnBJuTpeS99idb2JBAZy2xs",
    "2bTHZ66SGb9qBgQm3tgm6kC5BHktF8rhNox",
    "xoRL6qaeS5g9qw2DbSwaP4MsoTNLiF3jP3",
    "uPKFoHN1YZXHeFTQPHxZ7jNBaJfmuz6kVz",
    "2NEVvy739Mibi6PvZ4sBzawX5Nmsip4FDV4",
    "yb3uh64f4ZU9UVekPMeJLmwpnghHCR8njY",
    "4tiwzBCuQsV3w15ptrh53JP2Aqx1BZbTWh",
    "LeA4W3gRdm92ETYJECnVg5cyYFYEzUb1yP",
    "tbKn8UDQ8mBJHzy25B6zWhuwd4z17L7Yqa",
    "bA35xsWJ755J9cPDM8eiBh2Kgi6ZVQW6YL",
    "2L1AkRMtC7ejdQHvK7wFgDt5m1VFt74pMNV",
    "JFjyfJpHTb1zhHdQUV6GUHmZs2LDtoH7YY",
    "4rUzNhR6Ec1m4Ekj8nEdfmFJoQaBddp7Af",
    "2GZRAQMCijxkshiUNiYBKw6Nnd1ijeGB11e",
    "2SKtyQ5q7mwwWJKGNS7K8o5f4sBrLp8YyeT",
    "2kH7ifjyhqfJRKFJYJiUsraCxPaNHvHY4zm",
    "cXZfLvsQkmSez91yaEAUwJc4M68MVGY6ix",
    "MYKMywRpz2Ym8rXovhfPJWo5EjCRZtjNEL",
    "2Vnxt1Sci7uTS3uGvu78RjQCzveDrXSKEVn",
    "2YwrtGT62S6rPGpyEVT1zwGbX1mPZhTnMxB",
    "4DGvFPAvHD3UN78Jtkv4EGtK4kypsaqpDF",
    "pJELiYE1ctGFd395o1mTu6kHCm4o2RsCNp",
    "2cYhtvkTBNbHajCUY8HkqqWWfrSDrEgc3tg",
    "2CioGgAXcySsMLAsy8ozr1iNbXMscEJsAxq",
    "28ULayukdk3b1iyKv5QpBJ4zqZrEU3LN7Ga",
    "6Qg1ohtEEwKsH2UBZn7hQMztZMatyZpajR",
    "2aiKwg7pVe5Lzb1UMUMy8e3dTnBdroCBdBY",
    "4SQ5DqpJ1CkpcbiqLEmzE2Fo8v44VDDf5t",
    "2Vmj5mKfyE9xxNYVkYKk3qNQevCk2h7CdmB",
    "QCRzv12FtWfyxTye6vPfoCkN3XtizR6Ztj",
    "e4XtZC53SmdRsx5LV3BrJS3f7Xi2Rf9d7j",
    "TGKfEggK76c6WrYnaf2bgjuNNzd6Y71BY4",
    "2Y6CJSKsvw35gYurEKkvzYL77LepJDWVeAo",
    "uL3aEANiTRCk2Y7DBr9dAjir9SEtEjuGNt",
    "36dwq54zUZXFeTxBzDjCgLyo5QRgzcNAfJ",
    "h3rJc5otrgvR8oaHtiRoLhtGKv51mhjACa",
    "bcbcYVk2mJZpV2kHMPYBWUpWGYuXDWxvTV",
    "6UrxrbMp8psEPtVaKbQnYGKKxHjvNo2fix",
    "CaaECd1HB6xMMNrroR2dmBEi8eiRzvX3dZ",
    "DgzWfsV5aMmwqKeqjQ481AAkphwwJt8AE2",
    "2AWwsRPyRNN5KjLyhGj6bSDyUj5R3rs3GjM",
    "na8BKnrzNhei1156efNZ2LTjXhDR8fgaXA",
    "42JyAkUYfPqFzk8f6ABUGDyiCFF6pERmUE",
    "2j2gFmrgcxWHpb2uoDFWCx2XyTQ9bu8XYay",
    "GS4LWmqSKzZ46JyVSTsJ5zAVeEdaGLxrgC",
    "num2ppLn3UXNFymGqGsB8ypf9L1KEoedEy",
    "MAnWaTNRvN6pQky16f1FFrvEwWe8byZumB",
    "2LK9gjjQitMkAfaYzSsWdaALtJFhJkwExrU",
    "2aKudGDYD5iaC5ZUQCHYEY5ULpvd3viYCdj",
    "5JYvS7Mx9GaUS6vYnAc6AMzvMVjFoSRz3A",
    "PooaS7xH2UVM8UzDHD96Rs3tr8TLMuzqGv",
    "XkTwMvZpoh8vmKEreg6gmhKsCRC8ngJztS",
    "2hbEZ1zDhWuQEKCcxcTeTP7MbKQh3RbDxhP",
    "2VHVsWRWv4jrT6X1tBbdkGzD7Yxr1sXmxDM",
    "ch73x6xoQSBX1KBpDgpR6FQwNze44DbF8V",
    "q6rKc15z2uJEsgfctVyxmnjZMfaVb6q4ga",
    "2Y5psVayUH9MpCt9c8LTWk81jRQ7S7zK92T",
    "2X8gY2DJ7F2JgjYLyrfKxsunJD3WQzdbu2t",
    "26fNdi3nq7ttwBA5vZZo3aQAMKG3igeEJbT",
    "u6LHNqcCie2JwA6YUBTdUrv8NnfuP6yyTt",
    "21SEt5JaG62fq1TrBP6ojHAkiJ2KRfPJJfM",
    "GQwS38GbgDF8sgtYcipQJiGBTv2LUyHfrn",
    "WjzSzBa321wsmYmr5RCMyAJ7vFAubM2bpw",
    "T1pWpmaL8J92z4PUkXbPUKd1oHDLGdR78L",
    "2SowYhhvZ7hopEQWM4xjGMSb4tbaWszPbXS",
    "YhfdaAMa4nB54uyA65LrQz9QtdRkvanaV4",
    "2h7qQRHnMwpF9N4pCm45MrVWxRCSzwrxoty",
    "29gkXRsQFfmmedMS3ZaPXyZLnsH5XyiMmci",
}
