package main

import "bufio"

// Структуры персонажей
type Warrior struct {
	Name      string
	Health    int
	MaxHP     int
	MinDamage int
	MaxDamage int
	Lvl       int
	XP        int
	Gold      int
	Tired     bool // усталость сохраняется МЕЖДУ боями, сбрасывается только "Быстрым выпадом"
	BonusCrit int  // бонус крита тоже переносится в следующий бой
}

type Mage struct {
	Name      string
	Health    int
	MaxHP     int
	Mana      int
	MaxMana   int
	MinDamage int
	MaxDamage int
	Lvl       int
	XP        int
	Gold      int
	FreezeNext bool
}

type Paladin struct {
	Name      string
	Health    int
	MaxHP     int
	Shield    int
	MaxShield int
	MinDamage int
	MaxDamage int
	Lvl       int
	XP        int
	Gold      int
}

type Rogue struct {
	Name      string
	Health    int
	MaxHP     int
	Energy    int
	MinDamage int
	MaxDamage int
	Lvl       int
	XP        int
	Gold      int
	BonusCrit int
}

type Hunter struct {
	Name      string
	Health    int
	MaxHP     int
	PetHp     int
	Lvl       int
	XP        int
	Gold      int
	MinDamage int
	MaxDamage int
}

// Структура монстра из твоего скриншота
type Monster struct {
	Kind     string
	Hp       int
	MinPower int
	MaxPower int
	DropXP   int
	GoldDrop int
}

// Интерфейсы
type Target interface {
	GetName() string
}

type Player interface {
	Target
	ShowStats()
	GetHealth() int
	GetMana() int
	SetHealth(hp int)
	UseMana(mana int) bool              // Проверка и трата маны для мага
	GetDamage(reader *bufio.Reader) int // Расчёт урона (рандом внутри класса с выбором приемов)
	AddXP(xp int, nextLvlXP int) bool   // Возвращает true, если апнулся левел
	AddGold(gold int)
	LevelUp() int
	GetGold() int
	SpendGold(amount int) bool
	FullHeal()
	UpgradeWeapon(bonusDmg int)
}
