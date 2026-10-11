package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint8 }
func genFn0() uint32 {
    var v4 int64 = -217
    fmt.Println(v4)
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := []int{42, (fi5 * (fi5 * fi5))}
        v4 = v4
        _ = fi5
        _ = v6
    }
    var v7 uint64 = 0
    v8 := GenS1{F0: 2, F1: false, F2: 1}
    _ = v4
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint = 4294967295
    var v2 bool = ((1 < 1000) || ((-2.25 - (0.1 * 1.5)) != 3.14159))
    v3 := []int{-2147483648, 4294967295}
    fmt.Println(v1, v2)
    v9 := GenS1{F0: 1, F1: ((1 > 955) && (391 != 4294967295)), F2: 37}
    if (((0.0 + 1e10) - 1.5) == 0.1) {
        v10 := GenS1{F0: 255, F1: ("hello" == "foo bar"), F2: 80}
        switch 2147483647 {
        case 32767:
            var v11 int8 = 75
            _ = v11
        case 42:
            v12 := GenS1{F0: 758, F1: false, F2: 2}
            _ = v12
        case -9223372036854775808:
            switch 4294967295 {
            case 1:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case -829:
                v2 = v2
            }
        }
        _ = v10
    }
    var v13 int16 = 1000
    var v14 uint32 = genFn0()
    v13 = v13
    v14 -= v14
    switch 32768 {
    case 0:
        var v15 float64 = (float64(1.5) * float64(1.5))
        _ = v15
    default:
        var v16 uint32 = (v14 % ((2147483647) | 1))
        _ = v16
    }
    v3 = append(v3, 32767)
    var v17 int32 = -548
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}
