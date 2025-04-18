package weightconv

// KgToLbs converts a kilograms weight to pounds.
func KgToLbs(kg Kilogram) Pound {
	return Pound(kg / 0.45359237)
}

// LbsToKg converts a pounds weight to kilograms.
func LbsToKg(lbs Pound) Kilogram {
	return Kilogram(lbs * 0.45359237)
}
