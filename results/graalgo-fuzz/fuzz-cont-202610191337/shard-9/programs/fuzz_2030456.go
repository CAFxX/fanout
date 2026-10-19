package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int8, p1 int32, p2 uint8) int8 {
    p1 += p1
    for fi4 := 0; fi4 < 7; fi4++ {
        for fi5 := 0; fi5 < 6; fi5++ {
            p2 -= (p2 - 42)
            _ = fi5
        }
        if (!false) {
            fmt.Println(p0, fi4, p2)
        } else {
            p0 = (p0 << ((0) & 63))
            p0 = -102
        }
        _ = fi4
    }
    var v6 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p0 * (p0 - p0))
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[uint32]bool)
    v3 := make(map[string]uint32)
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: 0}
        _ = fi7
        _ = v8
    }
    switch 2147483647 {
    case 4294967295:
        v9 := GenS1{F0: 42}
        _ = v9
    case 127:
        switch ("a" + "") {
        case "\n\t":
            if ((-740 > -1) && (!(true || true))) {
                for fi10 := 0; fi10 < 6; fi10++ {
                    if false {
                        switch ("hello" + "foo bar") {
                        case "\n\t":
                            var v11 uint = 4294967295
                            _ = v11
                        default:
                            var v12 uint8 = 255
                            _ = v12
                        }
                    } else {
                        var v13 uint8 = 182
                        for fi14 := 0; fi14 < 18; fi14++ {
                            v13 += (^(v13))
                            _ = fi14
                        }
                        for fi15 := 0; fi15 < 16; fi15++ {
                            fmt.Println(v13, v1, fi10)
                            var v16 int16 = -1
                            _ = fi15
                            _ = v16
                        }
                        _ = v13
                    }
                    _ = fi10
                }
            } else {
                v17 := []int{-1, 42, 588, 127}
                _ = v17
            }
        }
    case 32768:
        fmt.Println(v1)
    }
    var v18 int64 = -9223372036854775808
    if false {
        if (((0.1 + 0.1) * ((1e10 - 0.0) / float64((((-2.25 / -2.25) + (float32((3.14159 * float64(((((-2.25 * 0.1) - 0.1) * 1e10) / 1.5)))) / ((3.14159 * 0.1) + 3.14159))) - 3.14159)))) >= 1.5) {
            v1 = v1
        }
    } else {
        fmt.Println(v1)
        switch 4294967295 {
        case 2147483647:
            switch "foo bar" {
            case "a":
                v19 := []int{32767, -819}
                _ = v19
            }
        }
        for fi20 := 0; fi20 < 15; fi20++ {
            var v21 uint = 65535
            v1 += v1
            var v22 int8 = genFn0(127, int32(v18), 1)
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    v23 := []int{}
    _ = len(v2)
    var v24 string = "\n\t"
    var v25 int = 127
    v23 = append(v23, (v25 / ((912) | 1)))
    v1 -= (v1 / (((v1 << ((2) & 63))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
