package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := []int{}
    v2 := []uint32{}
    var v3 float64 = float64(float32((((0.1 + 0.0) / 1e10) / 1.0)))
    fmt.Println(v3)
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 int8 = 42
        switch "" {
        case "hello":
            _ = len(v1)
        default:
            for fi6 := 0; fi6 < 7; fi6++ {
                v7 := GenS1{F0: 42}
                if false {
                    switch (^(fi6)) {
                    case 1:
                        v3 = v3
                    case -685:
                        v1 = append(v1, (fi6 | (fi4 ^ fi4)))
                        if (193 <= 52) {
                            var v8 uint16 = 1
                            _ = v8
                        } else {
                            v5 += (v5 | v5)
                        }
                        v3 += v3
                    default:
                        if (42 > -32768) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            fmt.Println(v5, v3)
                            var v9 string = "hello"
                            _ = len(v2)
                            switch "a" {
                            case "foo bar":
                                var v10 uint64 = 795
                                _ = v10
                            default:
                                var v11 uint = 1
                                _ = v11
                            }
                            _ = v9
                        }
                    }
                    v12 := GenS1{F0: 0}
                    _ = v12
                }
                _ = fi6
                _ = v7
            }
        }
        _ = fi4
        _ = v5
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        for fi14 := 0; fi14 < 20; fi14++ {
            if len(v2) > 0 {
                v2[0] = 2
            }
            _ = fi14
        }
        var v15 int64 = 1
        _ = fi13
        _ = v15
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        v17 := GenS1{F0: 151}
        _ = fi16
        _ = v17
    }
    var v18 int16 = 188
    v18 -= (v18 ^ (v18 & -325))
    switch 0 {
    case -435:
        v3 -= float64(float32(float64((0.0 + (1.5 - -2.25)))))
    }
    v19 := GenS1{F0: 2147483647}
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 uint32 = 4294967295
        var v22 int = (fi20 >> ((4294967295) & 63))
        var v23 uint = 894
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        v3 += v3
        var v25 string = "a"
        _ = fi24
        _ = v25
    }
    v3 = (((1e10 * v3) * (((0.1 / (v3 / v3)) - ((1e10 / (0.0 + 1.5)) / (1.5 * (((v3 + 3.14159) + 1.5) / ((float64(-2.25) / 1.0) / 0.1))))) - v3)) - 0.1)
    switch "foo bar" {
    case "foo bar":
        var v26 int16 = v18
        _ = v26
    case "":
        v27 := GenS2{F0: 128}
        _ = v27
    }
    v3 -= v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
