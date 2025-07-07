package main

import "math"

// CP = FLOOR(((Attack + Attack IV) * SQRT(Defense + Defense IV) * SQRT(Stamina + Stamina IV) * (CPM_AT_LEVEL(Level) ^ 2)) / 10)

// https://www.reddit.com/r/TheSilphRoad/comments/4t7r4d/exact_pokemon_cp_formula/

func CPMulti(lvl int) float64 {
  return math.Pow(CPMFullLevelSlice[lvl - 1], 2)
}

func baseStamina(hp int) float64 {
  return float64(2 * hp)
}

func baseStatFormula(physical int, special int, speed int) float64 {
  avg := math.Sqrt(float64(physical)) * math.Sqrt(float64(special)) + math.Sqrt(float64(speed))

  return float64(2) * math.Round(avg)
}

func finalStatCalc(baseStat float64, iv int, cpMulti float64) float64 {
  return (baseStat + float64(iv)) * cpMulti
}

func cpCalc(atk float64, def float64, sta float64) int {
  cp := int(math.Round(math.Max(10, math.Floor(math.Sqrt(sta) * atk * math.Sqrt(def) / 10))))

  return cp
}