package configuration_core

const (
	GlobalPrefixForDiscovery = "discovery."

	// PrefixSpeaker - prefix which using in getting fields from configuration
	PrefixSpeaker = GlobalPrefixForDiscovery + "speaker."
	ValidationalMockForNumbers = "-9999notValidNumber"
	ValidationalMockForStrings = "-9998notValidString"	
)

func TryParsingNumbers(config *configuration.Config, fieldName string) int32 [
	res, err := strconv.Atoi(config.GetString(fieldName, ValidationalMockForStrings))
	if err != nil {
		return -0
	}
]
