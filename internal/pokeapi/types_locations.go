package pokeapi

type RespShallowLocations struct {
	Count    int
	Next     *string
	Previous *string
	Results  []LocationArea
}

type LocationArea struct {
	Name string
	Url  string
}

type RespFullLocations struct {
	id                   int                   `json:"id"`
	GameIndex            int                   `json:"game_index"`
	Name                 string                `json:"name"`
	Location             LocationArea          `json:"location"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Names                []LocationLanguage    `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	EncounterMethod BaseDescription `json:"encounter_method"`
	VersionDetails  []VersionDetail `json:"version_details"`
}

type VersionDetail struct {
	Rate    int             `json:"rate"`
	Version BaseDescription `json:"version"`
}

type LocationLanguage struct {
	Language BaseDescription `json:"language"`
	Name     string          `json:"name"`
}

type PokemonEncounter struct {
	Pokemon                 BaseDescription           `json:"pokemon"`
	EncounterVersionDetails []EncounterVersionDetails `json:"version_details"`
}

type EncounterVersionDetails struct {
	EncounterDetails []EncounterDetail `json:"encounter_details"`
	MaxChance        int               `json:"max_chance"`
	Version          BaseDescription   `json:"version"`
}

type EncounterDetail struct {
	Chance int `json:"chance"`
	// yes ConditionValue
	MaxLevel int             `json:"max_level"`
	Method   BaseDescription `json:"method"`
	Minlevel int             `json:"min_level"`
}
