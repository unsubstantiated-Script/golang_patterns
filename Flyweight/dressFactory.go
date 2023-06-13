package Flyweight

import "fmt"

const (
	//Clothing types
	TerroristSkinType        = "tSkin"
	CounterTerroristSkinType = "ctSkin"
)

var (
	skinFactorySingleInstance = &SkinFactory{
		skinMap: make(map[string]Skin),
	}
)

type SkinFactory struct {
	skinMap map[string]Skin
}

func (s *SkinFactory) getSkinByType(skinType string) (Skin, error) {
	if s.skinMap[skinType] != nil {
		return s.skinMap[skinType], nil
	}

	if skinType == TerroristSkinType {
		s.skinMap[skinType] = newTerroristSkin()
		return s.skinMap[skinType], nil
	}
	if skinType == CounterTerroristSkinType {
		s.skinMap[skinType] = newCounterTerroristSkin()
		return s.skinMap[skinType], nil
	}
	return nil, fmt.Errorf("wrong skin type passed")
}

func getSkinFactorySingleInstance() *SkinFactory {
	return skinFactorySingleInstance
}
