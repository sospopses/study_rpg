package main

import (
	"bufio"
	"fmt"
	"math/rand"
)

// Вспомогательная функция для генерации случайного урона в диапазоне
func randomDmg(min, max int) int {
	if max-min <= 0 {
		return min
	}
	return rand.Intn(max-min+1) + min
}

// Проверка на критический удар
func isCrit(chance int) bool {
	return rand.Intn(100) < chance
}

// ==========================================
// МЕТОДЫ ВОИНА (Warrior)
// ==========================================

func (w *Warrior) GetName() string {
	return w.Name
}

func (w *Warrior) ShowStats() {
	fmt.Printf("[Воин] Имя: %s | Здоровье: %d/%d | Урон: %d-%d | Лвл: %d | Опыт: %d | Золото: %d",
		w.Name, w.Health, w.MaxHP, w.MinDamage, w.MaxDamage, w.Lvl, w.XP, w.Gold)
}

func (w *Warrior) GetHealth() int {
	return w.Health
}

func (w *Warrior) GetMana() int {
	return 0
}

func (w *Warrior) SetHealth(hp int) {
	w.Health = hp
	if w.Health > w.MaxHP {
		w.Health = w.MaxHP
	}
}

func (w *Warrior) UseMana(mana int) bool {
	return false
}

func (w *Warrior) AddGold(gold int) {
	w.Gold += gold
}

func (w *Warrior) GetDamage(reader *bufio.Reader) int {
	for {
		fmt.Println("\n--- ХОД ВОИНА ---")
		if w.Tired {
			fmt.Println("Вы устали! Следующий Размашистый удар нанесет урон ВАМ!")
		}
		fmt.Println("1. Быстрый выпад (70% урона, снимает усталость, 100% точность)")
		fmt.Printf("2. Мощный удар (100%% урон, дает +20%% к шансу крита на след. ход). Текущий бонус: +%d%%\n", w.BonusCrit)
		fmt.Println("3. Размашистый удар (180% урон, 25% базовый крит. Вызывает усталость!)")
		fmt.Print("Выбери прием: ")

		choice := readInt(reader)

		baseDmg := randomDmg(w.MinDamage, w.MaxDamage)

		switch choice {
		case 1:
			w.Tired = false
			dmg := int(float64(baseDmg) * 0.7)
			if dmg < 1 {
				dmg = 1
			}
			critChance := 15 + w.BonusCrit
			w.BonusCrit = 0
			if isCrit(critChance) {
				dmg = int(float64(dmg) * 1.5)
				fmt.Printf("КРИТ! %s молниеносно протыкает врага на %d урона!\n", w.Name, dmg)
			} else {
				fmt.Printf("%s делает быстрый выпад на %d урона. Усталость снята!\n", w.Name, dmg)
			}
			pause_midle()
			return dmg

		case 2:
			dmg := baseDmg
			critChance := 15 + w.BonusCrit
			w.BonusCrit = 20

			if isCrit(critChance) {
				dmg = int(float64(dmg) * 1.5)
				fmt.Printf("КРИТ! %s обрушивает тяжелый меч на %d урона!\n", w.Name, dmg)
			} else {
				fmt.Printf("%s наносит мощный удар на %d урона и готовится к следующей атаке!\n", w.Name, dmg)
			}
			pause_midle()
			return dmg

		case 3:
			if w.Tired {
				selfDmg := int(float64(baseDmg) * 0.5)
				w.Health -= selfDmg
				if w.Health < 1 {
					w.Health = 1
				}
				reducedDmg := int(float64(baseDmg) * 0.3)
				fmt.Printf("Из-за усталости %s теряет равновесие! Вы нанесли себе %d урона, а врагу всего %d!\n", w.Name, selfDmg, reducedDmg)
				pause_midle()
				return reducedDmg
			}

			w.Tired = true
			dmg := int(float64(baseDmg) * 1.8)
			critChance := 25 + w.BonusCrit
			w.BonusCrit = 0

			if isCrit(critChance) {
				dmg *= 2
				fmt.Printf("СОКРУШИТЕЛЬНЫЙ КРИТ! %s разрубает врага на %d урона!\n", w.Name, dmg)
			} else {
				fmt.Printf("%s совершает тяжелый размашистый удар на %d урона!\n", w.Name, dmg)
			}
			pause_midle()
			return dmg

		default:
			fmt.Println("Неверный выбор! Соберись!")
		}
	}
}

func (w *Warrior) AddXP(xp int, nextLvlXP int) bool {
	w.XP += xp
	if w.XP >= nextLvlXP {
		w.XP -= nextLvlXP
		return true
	}
	return false
}

func (w *Warrior) LevelUp() int {
	w.Lvl++
	w.MaxHP = int(float64(w.MaxHP) * 1.3)
	w.MinDamage = int(float64(w.MinDamage) * 1.4)
	w.MaxDamage = int(float64(w.MaxDamage) * 1.4)
	w.Health = int(float64(w.Health) * 1.3)
	return w.MaxHP
}

func (w *Warrior) GetLevel() int {
	return w.Lvl
}

func (w *Warrior) GetGold() int {
	return w.Gold
}

func (w *Warrior) SpendGold(amount int) bool {
	if w.Gold >= amount {
		w.Gold -= amount
		return true
	}
	return false
}

func (w *Warrior) FullHeal() {
	w.Health = w.MaxHP
}

func (w *Warrior) UpgradeWeapon(bonusDmg int) {
	w.MinDamage += bonusDmg
	w.MaxDamage += bonusDmg
}

// ==========================================
// МЕТОДЫ МАГА (Mage)
// ==========================================

func (m *Mage) GetName() string {
	return m.Name
}

func (m *Mage) ShowStats() {
	fmt.Printf("[Маг] Имя: %s | Здоровье: %d/%d | Мана: %d/%d | Урон: %d-%d | Лвл: %d | Опыт: %d | Золото: %d",
		m.Name, m.Health, m.MaxHP, m.Mana, m.MaxMana, m.MinDamage, m.MaxDamage, m.Lvl, m.XP, m.Gold)
}

func (m *Mage) GetHealth() int {
	return m.Health
}

func (m *Mage) GetMana() int {
	return m.Mana
}

func (m *Mage) SetHealth(hp int) {
	m.Health = hp
	if m.Health > m.MaxHP {
		m.Health = m.MaxHP
	}
}

func (m *Mage) UseMana(mana int) bool {
	if m.Mana >= mana {
		m.Mana -= mana
		return true
	}
	return false
}

func (m *Mage) AddGold(gold int) {
	m.Gold += gold
}

func (m *Mage) GetDamage(reader *bufio.Reader) int {
	for {
		fmt.Println("\n--- ХОД МАГА ---")
		if m.FreezeNext {
			fmt.Println("❄️ Враг заморожен! Следующее заклинание гарантированно КРИТАНЕТ!")
		}
		fmt.Println("1. Удар посохом (70% урона, восстанавливает +25 маны)")
		fmt.Println("2. Ледяная стрела (120% урона, стоит 15 маны, 20% шанс заморозить врага)")
		fmt.Printf("3. Огненный шар (250%% урона, стоит 40 маны!). Мана: %d/%d\n", m.Mana, m.MaxMana)
		fmt.Print("Выбери заклинание: ")

		choice := readInt(reader)

		baseDmg := randomDmg(m.MinDamage, m.MaxDamage)

		switch choice {
		case 1:
			m.Mana += 25
			if m.Mana > m.MaxMana {
				m.Mana = m.MaxMana
			}
			dmg := int(float64(baseDmg) * 0.7)
			if dmg < 1 {
				dmg = 1
			}
			m.FreezeNext = false
			fmt.Printf("%s бьет врага посохом на %d урона и впитывает ману! (Мана: %d)\n", m.Name, dmg, m.Mana)
			pause_midle()
			return dmg

		case 2:
			if !m.UseMana(15) {
				fmt.Println("Недостаточно маны для Ледяной стрелы!")
				continue
			}
			dmg := int(float64(baseDmg) * 1.2)

			if m.FreezeNext || isCrit(15) {
				dmg *= 2
				m.FreezeNext = false
				fmt.Printf("КРИТ! Ледяная стрела разрывает цель на %d урона!\n", dmg)
			} else {
				if rand.Intn(100) < 20 {
					m.FreezeNext = true
					fmt.Printf("Ледяная стрела наносит %d урона и КРИСТАЛЛИЗУЕТ врага!\n", dmg)
				} else {
					fmt.Printf("Ледяная стрела наносит %d урона.\n", dmg)
				}
			}
			pause_midle()
			return dmg

		case 3:
			if !m.UseMana(40) {
				fmt.Println("Не хватает маны на Файербол! Сделай пару ударов посохом.")
				continue
			}
			dmg := int(float64(baseDmg) * 2.5)

			if m.FreezeNext || isCrit(10) {
				dmg *= 2
				m.FreezeNext = false
				fmt.Printf("АДСКИЙ КРИТ! Огненный шар испепеляет врага на %d урона!\n", dmg)
			} else {
				fmt.Printf("%s запускает Огненный шар на %d урона!\n", m.Name, dmg)
			}
			pause_midle()
			return dmg

		default:
			fmt.Println("Ты запутался в заклинаниях!")
		}
	}
}

func (m *Mage) AddXP(xp int, nextLvlXP int) bool {
	m.XP += xp
	if m.XP >= nextLvlXP {
		m.XP -= nextLvlXP
		return true
	}
	return false
}

func (m *Mage) LevelUp() int {
	m.Lvl++
	m.MaxHP = int(float64(m.MaxHP) * 1.3)
	m.MaxMana = int(float64(m.MaxMana) * 1.5)
	m.MinDamage = int(float64(m.MinDamage) * 1.4)
	m.MaxDamage = int(float64(m.MaxDamage) * 1.4)
	m.Health = int(float64(m.Health) * 1.3)
	m.Mana = int(float64(m.Mana) * 1.5)
	return m.MaxHP
}

func (m *Mage) GetLevel() int {
	return m.Lvl
}

func (m *Mage) GetGold() int {
	return m.Gold
}

func (m *Mage) SpendGold(amount int) bool {
	if m.Gold >= amount {
		m.Gold -= amount
		return true
	}
	return false
}

func (m *Mage) FullHeal() {
	m.Health = m.MaxHP
	m.Mana = m.MaxMana
}

func (m *Mage) UpgradeWeapon(bonusDmg int) {
	m.MinDamage += bonusDmg
	m.MaxDamage += bonusDmg
}

// ==========================================
// МЕТОДЫ ПАЛАДИНА (Paladin)
// ==========================================

func (p *Paladin) GetName() string {
	return p.Name
}

func (p *Paladin) ShowStats() {
	fmt.Printf("[Паладин] Имя: %s | Здоровье: %d/%d | Щит: %d | Урон: %d-%d | Лвл: %d | Опыт: %d | Золото: %d",
		p.Name, p.Health, p.MaxHP, p.Shield, p.MinDamage, p.MaxDamage, p.Lvl, p.XP, p.Gold)
}

func (p *Paladin) GetHealth() int {
	return p.Health
}

func (p *Paladin) GetMana() int {
	return 0
}

func (p *Paladin) SetHealth(hp int) { //сразу с вычетом щита
	if hp < p.Health {
		damage := p.Health - hp
		if p.Shield > 0 {
			if p.Shield >= damage {
				p.Shield -= damage
				fmt.Printf("🛡️ Щит паладина заблокировал весь урон! (Остаток щита: %d)\n", p.Shield)
				return
			} else {
				damage -= p.Shield
				fmt.Printf("🛡️ Щит паладина полностью разрушен! Пропущено %d урона по здоровью.\n", damage)
				p.Shield = 0
			}
		}
		p.Health -= damage
	} else {
		p.Health = hp
		if p.Health > p.MaxHP {
			p.Health = p.MaxHP
		}
	}
}

func (p *Paladin) UseMana(mana int) bool {
	return false
}

func (p *Paladin) AddGold(gold int) {
	p.Gold += gold
}

func (p *Paladin) GetDamage(reader *bufio.Reader) int {
	for {
		fmt.Println("\n--- ХОД ПАЛАДИНА ---")
		fmt.Println("1. Удар Света (100% урон, лечит себя на 30% от нанесенного урона, крит 10%)")
		fmt.Printf("2. Удар щитом (60%% урон, восстанавливает +15 ед. Щита). Твой щит: %d\n", p.Shield)
		fmt.Println("3. Священный пламень (220% урон, ТРЕБУЕТ 20 ед. Щита, шанс крита 50%!)")
		fmt.Print("Выбери прием: ")

		choice := readInt(reader)

		baseDmg := randomDmg(p.MinDamage, p.MaxDamage)

		switch choice {
		case 1:
			dmg := baseDmg
			isCritStrike := isCrit(10)
			if isCritStrike {
				dmg = int(float64(dmg) * 1.5)
			}

			heal := int(float64(dmg) * 0.3)
			if heal < 1 {
				heal = 1
			}
			p.Health += heal
			if p.Health > p.MaxHP {
				p.Health = p.MaxHP
			}

			if isCritStrike {
				fmt.Printf("КРИТ СВЕТА! %s наносит %d урона и исцеляется на %d ХП! (ХП: %d/%d)\n", p.Name, dmg, heal, p.Health, p.MaxHP)
			} else {
				fmt.Printf("%s бьет булавой на %d урона и исцеляется на %d ХП! (ХП: %d/%d)\n", p.Name, dmg, heal, p.Health, p.MaxHP)
			}
			pause_midle()
			return dmg

		case 2:
			p.Shield += 15
			if p.Shield > p.MaxShield {
				p.Shield = p.MaxShield
			}
			dmg := int(float64(baseDmg) * 0.6)
			if dmg < 1 {
				dmg = 1
			}
			fmt.Printf("%s делает сильный выпад щитом на %d урона и восстанавливает броню (+15 к щиту)!\n", p.Name, dmg)
			pause_midle()
			return dmg

		case 3:
			if p.Shield < 20 {
				fmt.Println("❌ Недостаточно щита для проведения Священного пламени! Накапливай прочность щитом.")
				continue
			}
			p.Shield -= 20
			dmg := int(float64(baseDmg) * 2.2)

			if isCrit(50) {
				dmg = int(float64(dmg) * 1.8)
				fmt.Printf("НЕБЕСНАЯ КАРА! %s испепеляет врага на %d священного урона!\n", p.Name, dmg)
			} else {
				fmt.Printf("Паладин обрушивает Священный пламень на %d урона, потратив 20 щита!\n", dmg)
			}
			pause_midle()
			return dmg

		default:
			fmt.Println("Выбери достойное праведника действие!")
		}
	}
}

func (p *Paladin) AddXP(xp int, nextLvlXP int) bool {
	p.XP += xp
	if p.XP >= nextLvlXP {
		p.XP -= nextLvlXP
		return true
	}
	return false
}

func (p *Paladin) LevelUp() int {
	p.Lvl++
	p.MaxHP = int(float64(p.MaxHP) * 1.3)
	p.MaxShield = int(float64(p.MaxShield) * 1.2)
	p.Shield = int(float64(p.Shield) * 1.2)
	p.MinDamage = int(float64(p.MinDamage) * 1.4)
	p.MaxDamage = int(float64(p.MaxDamage) * 1.4)
	p.Health = int(float64(p.Health) * 1.3)
	return p.MaxHP
}

func (p *Paladin) GetLevel() int {
	return p.Lvl
}

func (p *Paladin) GetGold() int {
	return p.Gold
}

func (p *Paladin) SpendGold(amount int) bool {
	if p.Gold >= amount {
		p.Gold -= amount
		return true
	}
	return false
}

func (p *Paladin) FullHeal() {
	p.Health = p.MaxHP
	p.Shield = p.MaxShield
}

func (p *Paladin) UpgradeWeapon(bonusDmg int) {
	p.MinDamage += bonusDmg
	p.MaxDamage += bonusDmg
}

// ==========================================
// МЕТОДЫ ПЛУТА (Rogue)
// ==========================================

func (r *Rogue) GetName() string {
	return r.Name
}

func (r *Rogue) ShowStats() {
	fmt.Printf("[Плут] Имя: %s | Здоровье: %d/%d | Энергия: %d | Урон: %d-%d | Лвл: %d | Опыт: %d | Золото: %d",
		r.Name, r.Health, r.MaxHP, r.Energy, r.MinDamage, r.MaxDamage, r.Lvl, r.XP, r.Gold)
}

func (r *Rogue) GetHealth() int {
	return r.Health
}

func (r *Rogue) GetMana() int {
	return 0
}

func (r *Rogue) SetHealth(hp int) {
	r.Health = hp
	if r.Health > r.MaxHP {
		r.Health = r.MaxHP
	}
}

func (r *Rogue) UseMana(mana int) bool {
	return false
}

func (r *Rogue) AddGold(gold int) {
	r.Gold += gold
}

func (r *Rogue) GetDamage(reader *bufio.Reader) int {
	for {
		fmt.Println("\n--- ХОД ПЛУТА ---")
		fmt.Printf("1. Подготовка (80%% урона, дает +30 Энергии и +25%% к шансу крита на след. ход). Бонус: +%d%%\n", r.BonusCrit)
		fmt.Println("2. Потрошение (140% урона, требует 35 Энергии. При КРИТЕ возвращает 15 Энергии)")
		fmt.Printf("3. Удар в спину (220%% урона, требует 60 Энергии! 100%% КРИТ, если ХП < 50%%). Твоя энергия: %d\n", r.Energy)
		fmt.Print("Выбери маневр: ")

		choice := readInt(reader)

		baseDmg := randomDmg(r.MinDamage, r.MaxDamage)

		switch choice {
		case 1:
			r.Energy += 30
			if r.Energy > r.MaxEnergy {
				r.Energy = r.MaxEnergy
			}
			dmg := int(float64(baseDmg) * 0.8)
			critChance := 30 + r.BonusCrit
			r.BonusCrit = 25

			if isCrit(critChance) {
				dmg *= 2
				fmt.Printf("КРИТ! %s бьет под колено на %d урона и готовится к серии!\n", r.Name, dmg)
			} else {
				fmt.Printf("%s делает быстрый укол на %d урона и концентрирует силы (Энергия: %d)\n", r.Name, dmg, r.Energy)
			}
			pause_midle()
			return dmg

		case 2:
			if r.Energy < 35 {
				fmt.Println("Маловато энергии для Потрошения! Сделай Подготовку.")
				continue
			}
			r.Energy -= 35
			dmg := int(float64(baseDmg) * 1.4)
			critChance := 30 + r.BonusCrit
			r.BonusCrit = 0

			if isCrit(critChance) {
				dmg *= 2
				r.Energy += 15
				fmt.Printf("КРИТ-ПОТРОШЕНИЕ! %s наносит %d урона и возвращает 15 Энергии! (Энергия: %d)\n", r.Name, dmg, r.Energy)
			} else {
				fmt.Printf("%s проводит серию ударов на %d урона! (Осталось энергии: %d)\n", r.Name, dmg, r.Energy)
			}
			pause_midle()
			return dmg

		case 3:
			if r.Energy < 60 {
				fmt.Println("Недостаточно энергии для Ультимейта (нужно 60)! Копи силы.")
				continue
			}
			r.Energy -= 60
			dmg := int(float64(baseDmg) * 2.2)
			critChance := 30 + r.BonusCrit
			r.BonusCrit = 0

			isGuaranteed := r.Health < (r.MaxHP / 2)
			if isGuaranteed || isCrit(critChance) {
				dmg *= 2
				if isGuaranteed {
					fmt.Printf("ОТЧАЯННЫЙ УДАР ИЗ ТЕНИ! %s, превозмогая боль, вонзает клинки на %d урона!\n", r.Name, dmg)
				} else {
					fmt.Printf("КРИТ ИЗ ТЕНИ! %s стирает врага на %d урона!\n", r.Name, dmg)
				}
			} else {
				fmt.Printf("%s наносит сильный удар из тени на %d урона!\n", r.Name, dmg)
			}
			pause_midle()
			return dmg

		default:
			fmt.Println("Двигайся бесшумно и делай правильный выбор!")
		}
	}
}

func (r *Rogue) AddXP(xp int, nextLvlXP int) bool {
	r.XP += xp
	if r.XP >= nextLvlXP {
		r.XP -= nextLvlXP
		return true
	}
	return false
}

func (r *Rogue) LevelUp() int {
	r.Lvl++
	r.MaxHP = int(float64(r.MaxHP) * 1.3)
	r.MaxEnergy = int(float64(r.MaxEnergy) * 1.2)
	r.Energy = int(float64(r.Energy) * 1.2)
	r.MinDamage = int(float64(r.MinDamage) * 1.4)
	r.MaxDamage = int(float64(r.MaxDamage) * 1.4)
	r.Health = int(float64(r.Health) * 1.3)
	return r.MaxHP
}

func (r *Rogue) GetLevel() int {
	return r.Lvl
}

func (r *Rogue) GetGold() int {
	return r.Gold
}

func (r *Rogue) SpendGold(amount int) bool {
	if r.Gold >= amount {
		r.Gold -= amount
		return true
	}
	return false
}

func (r *Rogue) FullHeal() {
	r.Health = r.MaxHP
	r.Energy = r.MaxEnergy
}

func (r *Rogue) UpgradeWeapon(bonusDmg int) {
	r.MinDamage += bonusDmg
	r.MaxDamage += bonusDmg
}

// ==========================================
// МЕТОДЫ ОХОТНИКА (Hunter)
// ==========================================

func (h *Hunter) GetName() string {
	return h.Name
}

func (h *Hunter) ShowStats() {
	fmt.Printf("[Охотник] Имя: %s | Здоровье: %d/%d | ХП Пета: %d | Урон: %d-%d | Лвл: %d | Опыт: %d/ | Золото: %d",
		h.Name, h.Health, h.MaxHP, h.PetHp, h.MinDamage, h.MaxDamage, h.Lvl, h.XP, h.Gold)
}

func (h *Hunter) GetHealth() int {
	return h.Health
}

func (h *Hunter) GetMana() int {
	return 0
}

func (h *Hunter) SetHealth(hp int) {
	h.Health = hp
	if h.Health > h.MaxHP {
		h.Health = h.MaxHP
	}
}

func (h *Hunter) UseMana(mana int) bool {
	return false
}

func (h *Hunter) AddGold(gold int) {
	h.Gold += gold
}

func (h *Hunter) GetDamage(reader *bufio.Reader) int {
	for {
		fmt.Println("\n--- ХОД ОХОТНИКА ---")
		fmt.Printf("1. Выстрел кобры (100%% урона, восстанавливает +20 ХП питомцу). ХП пета: %d\n", h.PetHp)
		fmt.Println("2. Прицельный выстрел (150% урона охотника, повышенный шанс крита 40%, пет отдыхает)")
		fmt.Println("3. Команда «Взять!» (Пет наносит 250% урона, но теряет 25 ХП в схватке!)")
		fmt.Print("Выбери приказ: ")

		choice := readInt(reader)

		baseDmg := randomDmg(h.MinDamage, h.MaxDamage)

		switch choice {
		case 1:
			dmg := baseDmg
			if isCrit(20) {
				dmg = int(float64(dmg) * 1.7)
				fmt.Printf("КРИТ! Выстрел кобры нанес %d урона!\n", dmg)
			} else {
				fmt.Printf("%s стреляет из лука на %d урона.\n", h.Name, dmg)
			}

			if h.PetHp > 0 {
				h.PetHp += 20
				fmt.Printf("Охотник перевязывает раны питомцу! (ХП пета: %d)\n", h.PetHp)
			}
			pause_midle()
			return dmg

		case 2:
			dmg := int(float64(baseDmg) * 1.5)
			if isCrit(40) {
				dmg = int(float64(dmg) * 1.7)
				fmt.Printf("СНАЙПЕРСКИЙ КРИТ! %s поражает уязвимую точку на %d урона!\n", h.Name, dmg)
			} else {
				fmt.Printf("%s делает точный прицельный выстрел на %d урона!\n", h.Name, dmg)
			}
			pause_midle()
			return dmg

		case 3:
			if h.PetHp <= 0 {
				fmt.Println("Твой верный питомец без сознания! Сначала подлечи его Выстрелом кобры.")
				continue
			}

			h.PetHp -= 25
			dmg := int(float64(baseDmg) * 2.5)

			if isCrit(20) {
				dmg = int(float64(dmg) * 1.7)
				fmt.Printf("КРИТИЧЕСКИЙ КУСЬ! Зверь рвет врага на %d урона!\n", dmg)
			} else {
				fmt.Printf("Зверь яростно атакует врага на %d урона!\n", dmg)
			}

			fmt.Println("Питомец получает 25 урона в пылу драки!")
			if h.PetHp <= 0 {
				h.PetHp = 0
				fmt.Println("Твой питомец потерял сознание и больше не может атаковать!")
			}
			pause_midle()
			return dmg

		default:
			fmt.Println("Отдай понятную зверю команду!")
		}
	}
}

func (h *Hunter) AddXP(xp int, nextLvlXP int) bool {
	h.XP += xp
	if h.XP >= nextLvlXP {
		h.XP -= nextLvlXP
		return true
	}
	return false
}

func (h *Hunter) LevelUp() int {
	h.Lvl++
	h.MaxHP = int(float64(h.MaxHP) * 1.3)
	h.MaxPetHp = int(float64(h.MaxPetHp) * 1.3)
	h.PetHp = int(float64(h.PetHp) * 1.3)
	h.MinDamage = int(float64(h.MinDamage) * 1.4)
	h.MaxDamage = int(float64(h.MaxDamage) * 1.4)
	h.Health = int(float64(h.Health) * 1.3)
	return h.MaxHP
}

func (h *Hunter) GetLevel() int {
	return h.Lvl
}

func (h *Hunter) GetGold() int {
	return h.Gold
}

func (h *Hunter) SpendGold(amount int) bool {
	if h.Gold >= amount {
		h.Gold -= amount
		return true
	}
	return false
}

func (h *Hunter) FullHeal() {
	h.Health = h.MaxHP
	h.PetHp = h.MaxPetHp
}

func (h *Hunter) UpgradeWeapon(bonusDmg int) {
	h.MinDamage += bonusDmg
	h.MaxDamage += bonusDmg
}
