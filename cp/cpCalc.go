package main

import "math"

// https://db.pokemongohub.net/tools/cp-calculator


func statFrac(n int) float64 {
  return float64(n) / float64(8)
}

// FLOOR(((Attack + Attack IV) * SQRT(Defense + Defense IV) * SQRT(Stamina + Stamina IV) * (CPM_AT_LEVEL(Level) ^ 2)) / 10)
func CPCalc(stats []float64, lvl int, atkIV int, defIV int, staIV int) int {
// HP, Atk, Def, SpAtk, SpDef, Spd
  higherAtk := math.Max(stats[1], stats[3])
  lowerAtk := math.Min(stats[1], stats[3])

  higherDef := math.Max(stats[2], stats[4])
  lowerDef := math.Min(stats[2], stats[4])


  scaledAttack := math.Round(2 * (statFrac(7) * higherAtk + statFrac(1) * lowerAtk))

  scaledDefense := math.Round(2 * (statFrac(5) * higherDef + statFrac(3) * lowerDef))

  speedMod := 1 + (stats[5] - 75) / 500

  
  baseAttack := math.Round(scaledAttack * speedMod)
  baseDefense := math.Round(scaledDefense * speedMod)
  baseStamina := math.Floor(stats[0] * 1.75 + 50)

  cp := math.Max(10, math.Floor(((baseAttack + float64(atkIV)) * math.Sqrt(baseDefense + float64(defIV)) * math.Pow(baseStamina + float64(staIV), 0.5) * math.Pow(cpm[lvl - 1], 2)) / 10))

  return int(cp)
}