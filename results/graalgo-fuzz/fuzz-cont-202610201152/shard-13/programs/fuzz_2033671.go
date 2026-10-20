package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint8 = 1
    var v2 uint64 = 178
    v3 := make(map[string]int)
    var v4 int = 42
    var v5 int64 = 42
    v6 := GenS1{F0: (v1 * 225)}
    if true {
        v2 = 847
        var v7 float64 = (-2.25 * (1.5 - 1e10))
        for fi8 := 0; fi8 < 11; fi8++ {
            fmt.Println(v1, v4)
            v4 += (v4 ^ 9223372036854775807)
            _ = fi8
        }
        _ = v7
    } else {
        for fi9 := 0; fi9 < 7; fi9++ {
            v4 += (v4 % ((fi9) | 1))
            _ = fi9
        }
        if (722 > 1) {
            v5 -= v5
            v5 -= (-(v5))
        }
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 float64 = float64(0.1)
            _ = fi10
            _ = v11
        }
    }
    var v12 float32 = 3.14159
    var v13 int16 = 0
    v14 := GenS1{F0: v1}
    v13 = v13
    if true {
        var v15 string = "foo bar"
        if (true || (!true)) {
            if false {
                for fi16 := 0; fi16 < 11; fi16++ {
                    var v17 uint8 = (^(v1))
                    switch 42 {
                    case 2147483647:
                        v5 += v5
                    case -447:
                        var v18 uint16 = 65535
                        _ = v18
                    default:
                        var v19 int16 = v13
                        _ = v19
                    }
                    _ = fi16
                    _ = v17
                }
                var v20 int8 = -128
                _ = v20
            }
        }
        _ = v15
    }
    var v21 bool = (2147483647 <= 0)
    for fi22 := 0; fi22 < 17; fi22++ {
        switch fi22 {
        case -2147483648:
            delete(v3, "a")
        }
        var v23 uint8 = (v1 / ((v1) | 1))
        if (v21 && true) {
            for fi24 := 0; fi24 < 3; fi24++ {
                if (v4 != fi22) {
                    fmt.Println(v4, v13)
                    v4 += (v4 / (((fi22 << ((18446744073709551615) & 63))) | 1))
                    v25 := v3["\n\t"]
                    _ = v25
                } else {
                    var v26 uint8 = 137
                    v27 := GenS2{F0: 773}
                    var v28 uint64 = (v2 | (v2 ^ 4294967295))
                    _ = v26
                    _ = v27
                    _ = v28
                }
                _ = fi24
            }
        }
        _ = fi22
        _ = v23
    }
    if (!v21) {
        for fi29 := 0; fi29 < 4; fi29++ {
            v4 = (v4 / (((v4 % ((fi29) | 1))) | 1))
            var v30 int32 = -2147483648
            _ = fi29
            _ = v30
        }
        if v21 {
            delete(v3, "")
            for fi31 := 0; fi31 < 16; fi31++ {
                var v32 uint16 = uint16(v2)
                v21 = true
                v12 += (1.5 * v12)
                _ = fi31
                _ = v32
            }
        }
        fmt.Println(v4, v13, v21)
    }
    v33 := GenS2{F0: 0}
    fmt.Println(v13, v4, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
