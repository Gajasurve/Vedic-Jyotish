//I, Gajanand, bow to Goddess Saraswati for wisdom and guidance.
package main

import (
        "fmt"
        "math"
        "strconv"
        "strings"
)

func main() {
        var lagnaInput string
        fmt.Print("Enter Lagna (e.g., Li,14,14,9 or 14,14,9): ")
        fmt.Scanln(&lagnaInput)

        lagnaSign, lagnaDegrees, lagnaMinutes, lagnaSeconds, err := parseLagnaInput(lagnaInput)
        if err != nil {
                fmt.Println("Error:", err)
                return
        }

        lagnaDecimalDegrees := calculateDecimalDegrees(lagnaSign, lagnaDegrees, lagnaMinutes, lagnaSeconds)
        kundaDegrees := calculateKundaDegrees(lagnaDecimalDegrees)

        fmt.Println("\n--- Calculation Steps ---")
        fmt.Printf("1. Input Lagna: %s %.0f°%.0f'%.0f''\n", lagnaSign, lagnaDegrees, lagnaMinutes, lagnaSeconds)
        fmt.Printf("2. Lagna Decimal Degrees: %.6f\n", lagnaDecimalDegrees)

        // Print possible signs
        possibleSigns := getPossibleSigns(lagnaSign)
        fmt.Printf("Possible Kunda placements for this Lagna: %s\n", strings.Join(possibleSigns, ", "))

        kundaSign := getZodiacSign(kundaDegrees)
        fmt.Printf("3. Kunda Degrees: %.6f (%s)\n", kundaDegrees, kundaSign)

        calculateAndPrintAdjustments(lagnaSign, kundaDegrees)
        // fmt.Printf("here")
}

func parseLagnaInput(input string) (string, float64, float64, float64, error) {
        parts := strings.Split(input, ",")
        if len(parts) < 3 || len(parts) > 4 {
                return "", 0, 0, 0, fmt.Errorf("invalid input format. Please use 'Lagna,degrees,minutes,seconds' (e.g., Li,14,14,9 or 14,14,9)")
        }

        lagnaSign := ""
        degreeIndex := 0
        if len(parts) == 4 {
                lagnaSign = strings.ToUpper(parts[0])
                degreeIndex = 1
        }

        degrees, err := strconv.ParseFloat(parts[degreeIndex], 64)
        if err != nil {
                return "", 0, 0, 0, err
        }

        minutes, err := strconv.ParseFloat(parts[degreeIndex+1], 64)
        if err != nil {
                return "", 0, 0, 0, err
        }

        seconds, err := strconv.ParseFloat(parts[degreeIndex+2], 64)
        if err != nil {
                return "", 0, 0, 0, err
        }

        return lagnaSign, degrees, minutes, seconds, nil
}

func calculateDecimalDegrees(lagnaSign string, degrees, minutes, seconds float64) float64 {
        signOffset := 0.0
        switch lagnaSign {
        case "AR":
                signOffset = 0.0
        case "TA":
                signOffset = 30.0
        case "GE":
                signOffset = 60.0
        case "CN":
                signOffset = 90.0
        case "LE":
                signOffset = 120.0
        case "VI":
                signOffset = 150.0
        case "LI":
                signOffset = 180.0
        case "SC":
                signOffset = 210.0
        case "SA":
                signOffset = 240.0
        case "CP":
                signOffset = 270.0
        case "AQ":
                signOffset = 300.0
        case "PI":
                signOffset = 330.0
        default:
                signOffset = 0 // if no sign, assume Aries, but this needs to be changed for more robust functionality.
        }

        return signOffset + degrees + (minutes / 60.0) + (seconds / 3600.0)
}

func calculateKundaDegrees(decimalDegrees float64) float64 {
        kundaDegrees := decimalDegrees * 81.0
        return math.Mod(kundaDegrees, 360.0)
}

func calculateAndPrintAdjustments(lagnaSign string, kundaDegrees float64) {
        lagnaOffset := getSignOffset(lagnaSign)
        trinesAndSeventh := getTrinesAndSeventh(lagnaOffset)

        fmt.Println("\n--- Rectification Adjustments ---")
        fmt.Println("Target Sign | Degree Adjustment | Time Adjustment")
        fmt.Println("------------|-------------------|-----------------")

        isKundaCorrect := false // Initialize to false

        for sign, target := range trinesAndSeventh {
                targetRangeStart := target - 15
                targetRangeEnd := target + 15

                // Adjust for wrapping around 360 degrees
                if targetRangeStart < 0 {
                        targetRangeStart += 360
                }
                if targetRangeEnd > 360 {
                        targetRangeEnd -= 360
                }

                // Check if kunda is within the target range
                if (kundaDegrees >= targetRangeStart && kundaDegrees <= targetRangeEnd) ||
                        (targetRangeStart > targetRangeEnd && (kundaDegrees >= targetRangeStart || kundaDegrees <= targetRangeEnd)) {
                        isKundaCorrect = true
                        continue // Skip to the next target sign
                }

                difference := target - kundaDegrees
                if difference < 0 {
                        difference += 360
                }

                lagnaAdjustment := difference / 81.0
                minutes := lagnaAdjustment * 60.0
                seconds := (minutes - math.Floor(minutes)) * 60.0

                fmt.Printf("%-12s| %-17.4f| %-17.0f minutes and %.0f seconds\n", sign, lagnaAdjustment, math.Floor(minutes), math.Floor(seconds))
        }

        if isKundaCorrect {
                fmt.Println("\nKunda is correctly placed within the target signs.")
        } else {
                fmt.Println("\nKunda is not correctly placed and needs rectification. Try Adjusting D-81 qith possible lagnas or choose the closet degree in above table")
        }
}

func getSignOffset(sign string) float64 {
        switch sign {
        case "AR":
                return 0.0
        case "TA":
                return 30.0
        case "GE":
                return 60.0
        case "CN":
                return 90.0
        case "LE":
                return 120.0
        case "VI":
                return 150.0
        case "LI":
                return 180.0
        case "SC":
                return 210.0
        case "SA":
                return 240.0
        case "CP":
                return 270.0
        case "AQ":
                return 300.0
        case "PI":
                return 330.0
        default:
                return 0.0
        }
}

func getTrinesAndSeventh(lagnaOffset float64) map[string]float64 {
        signs := []string{"AR", "TA", "GE", "CN", "LE", "VI", "LI", "SC", "SA", "CP", "AQ", "PI"}
        signOffsets := make(map[string]float64)
        for i, sign := range signs {
                signOffsets[sign] = float64(i * 30)
        }

        trinesAndSeventh := make(map[string]float64)
        lagnaIndex := int(lagnaOffset / 30)
        trinesAndSeventh[signs[lagnaIndex]] = signOffsets[signs[lagnaIndex]] + 15 // midpoint
        trinesAndSeventh[signs[(lagnaIndex+4)%12]] = signOffsets[signs[(lagnaIndex+4)%12]] + 15
        trinesAndSeventh[signs[(lagnaIndex+8)%12]] = signOffsets[signs[(lagnaIndex+8)%12]] + 15
        trinesAndSeventh[signs[(lagnaIndex+6)%12]] = signOffsets[signs[(lagnaIndex+6)%12]] + 15

        return trinesAndSeventh
}

func getZodiacSign(degrees float64) string {
        signs := []string{"AR", "TA", "GE", "CN", "LE", "VI", "LI", "SC", "SA", "CP", "AQ", "PI"}
        signIndex := int(degrees / 30)
        return signs[signIndex%12]
}

func getPossibleSigns(lagnaSign string) []string {
        lagnaOffset := getSignOffset(lagnaSign)
        trinesAndSeventh := getTrinesAndSeventh(lagnaOffset)

        possibleSigns := make([]string, 0, len(trinesAndSeventh))
        for sign := range trinesAndSeventh {
                possibleSigns = append(possibleSigns, sign)
        }

        return possibleSigns
}

// Composed this spending some days at Vajrayogini Temple, Nagarkot, Nepal. 
