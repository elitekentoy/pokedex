package pokeapi

type Pokemon struct {
	Id                     int               `json:"id"`
	BaseExperience         int               `json:"base_experience"`
	Height                 int               `json:"height"`
	Order                  int               `json:"order"`
	Weight                 int               `json:"weight"`
	Name                   string            `json:"name"`
	LocationAreaEncounters string            `json:"location_area_encounters"`
	IsDefault              bool              `json:"is_default"`
	Species                BaseDescription   `json:"species"`
	Sprites                Sprites           `json:"sprites"`
	Cries                  Cries             `json:"cries"`
	Stats                  []Stats           `json:"stats"`
	Abilities              []Abilities       `json:"abilities"`
	Forms                  []BaseDescription `json:"forms"`
	GameIndeces            []GameIndex       `json:"game_indices"`
	HeldItems              []ItemHeld        `json:"held_items"`
	Moves                  []PokemonMove     `json:"moves"`
	Types                  []PokemonTypes    `json:"types"`
	PastTypes              []PastPokemonType `json:"past_types"`
}

type Abilities struct {
	IsHidden bool            `json:"is_hidden"`
	Slot     int             `json:"slot"`
	Ability  BaseDescription `json:"ability"`
}

type GameIndex struct {
	Index   int             `json:"game_index"`
	Version BaseDescription `json:"version"`
}

type ItemHeld struct {
	Item           BaseDescription `json:"item"`
	VersionDetails []ItemVersion   `json:"version_details"`
}

type ItemVersion struct {
	Rarity  int             `json:"rarity"`
	Version BaseDescription `json:"version"`
}

type PokemonMove struct {
	Move         BaseDescription `json:"move"`
	VersionGroup []MoveVersion   `json:"version_group_details"`
}

type MoveVersion struct {
	LevelLearnedAt int             `json:"level_learned_at"`
	VersionGroup   BaseDescription `json:"version_group"`
	LearnMethod    BaseDescription `json:"move_learn_method"`
}

type Sprites struct {
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  string `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}

type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type Stats struct {
	BaseStat int             `json:"base_stat"`
	effort   int             `json:"effort"`
	Stat     BaseDescription `json:"stat"`
}

type PokemonTypes struct {
	Slot int             `json:"slot"`
	Type BaseDescription `json:"type"`
}

type PastPokemonType struct {
	Generation BaseDescription `json:"generation"`
	Types      []PokemonTypes  `json:"types"`
}
