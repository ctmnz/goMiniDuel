package main

import (
	"fmt"
	"math/rand"
	"time"
)


type hero struct {
	Name string
	Atributes heroAtributes
}

type heroAtributes struct {
	Strenght int
	Dexterity int
	Vitality int
	CritChance int
}


func generateRandomHero(name string) *hero {
	heroName := name
	rand.Seed( time.Now().UTC().UnixNano())
	heroAtrStrenght := 1+(rand.Intn(99))	
	heroAtrDexterity := 1+(rand.Intn(99))
	heroVitlity := 1+(rand.Intn(99))
	heroCritChance := 1+(rand.Intn(99))
	heroAtr := heroAtributes{Strenght: heroAtrStrenght, Dexterity: heroAtrDexterity, Vitality: heroVitlity, CritChance: heroCritChance}
	return &hero{Name: heroName, Atributes: heroAtr}
}


func showHeroInfo(hero *hero) {
	fmt.Println("========================")
	fmt.Println("# Name: ", hero.Name)
	fmt.Println("# Strenght: ", hero.Atributes.Strenght)
	fmt.Println("# Dexterity: ", hero.Atributes.Dexterity)
	fmt.Println("# Vitality: ", hero.Atributes.Vitality)
	fmt.Println("# CritChance: ", hero.Atributes.CritChance)
}


func hit(strenght int, critchance int) (int, bool) {
	rand.Seed( time.Now().UTC().UnixNano())
	iscrit := false
	dmgMin := 2
	dmgMax := 8

	dmgRatio := dmgMax - dmgMin

	damage := (dmgMin + (dmgRatio/((strenght - rand.Intn(strenght)))))
	if ((rand.Intn(100)) < critchance) {
		iscrit = true
		damage = (damage * 2)
	} else {
		iscrit = false
	}
	return damage, iscrit
}

func duel(h1 *hero, h2 *hero) string {	
	// h1 hit
	for {
	hitdmgP1, hitcritP1 := hit(h1.Atributes.Strenght, h1.Atributes.CritChance)
	
	h2.Atributes.Vitality = h2.Atributes.Vitality - hitdmgP1
	
	critmsgP1 := ""
	if(hitcritP1) {critmsgP1 = "[CRITICAL]"}

	fmt.Println(h1.Name, "Hits", h2.Name, "(",h2.Atributes.Vitality ,")", " with the sword ", critmsgP1, " and took ", hitdmgP1 , " HP" )
	if (h2.Atributes.Vitality <= 0) {
		return h1.Name
	}


        // h2 hit
        hitdmgP2, hitcritP2 := hit(h2.Atributes.Strenght, h2.Atributes.CritChance)
        h1.Atributes.Vitality = h1.Atributes.Vitality - hitdmgP2
        
	critmsgP2 := ""
	if(hitcritP2) {critmsgP2 = "[CRITICAL]"}
	fmt.Println(h2.Name, "Hits", h1.Name, "(",h2.Atributes.Vitality ,")", " with the sword ", critmsgP2, " and took ", hitdmgP1 , " HP" )

        if (h1.Atributes.Vitality <= 0) {
                return h2.Name
        }
	
	}
	
	return "false"

}

func main() {
	h1 := generateRandomHero("Kuhan")
	h2 := generateRandomHero("Odoum")
	showHeroInfo(h1)
	showHeroInfo(h2)
	fmt.Println("The winner is: ", duel(h1, h2))

}








