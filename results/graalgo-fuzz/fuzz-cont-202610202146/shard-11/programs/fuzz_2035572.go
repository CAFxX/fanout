package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 float64 = -2.25
    v2 := []int{}
    v3 := []int{}
    for fi4 := 0; fi4 < 12; fi4++ {
        fmt.Println(fi4, fi4)
        if (true && (!(-46 == 1))) {
            var v5 int32 = 2147483647
            var v6 uint16 = 42
            for fi7 := 0; fi7 < 18; fi7++ {
                v8 := GenS1{F0: 175}
                var v9 uint8 = 0
                var v10 float64 = float64(0.1)
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
            _ = v5
            _ = v6
        }
        switch fi4 {
        case -2147483648:
            if (2147483647 > 1) {
                if (true && true) {
                    v3 = append(v3, 2147483647)
                    var v11 uint = 2
                    _ = v11
                } else {
                    fmt.Println(fi4, fi4)
                    for fi12 := 0; fi12 < 9; fi12++ {
                        switch "\n\t" {
                        case "a":
                            v1 -= v1
                        }
                        v1 -= (1.5 + v1)
                        _ = fi12
                    }
                }
            }
        case -1:
            v1 += v1
        default:
            var v13 int8 = 127
            _ = v13
        }
        _ = fi4
    }
    v3 = append(v3, 32767)
    if (519 != 42) {
        for fi14 := 0; fi14 < 15; fi14++ {
            v15 := GenS1{F0: 4294967295}
            v16 := GenS1{F0: 2}
            var v17 int8 = int8(fi14)
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        v18 := GenS1{F0: 2}
        _ = v18
    }
    v3 = append(v3, 32768)
    var v19 uint8 = 1
    if false {
        v19 = v19
        var v20 float64 = ((float64(float32(0.1)) / v1) * v1)
        _ = v20
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        v19 = 2
        v3 = append(v3, fi21)
        _ = fi21
    }
    v3 = append(v3, -959)
    for fi22 := 0; fi22 < 13; fi22++ {
        for fi23 := 0; fi23 < 20; fi23++ {
            v19 += v19
            _ = fi23
        }
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
}
