package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS2 struct { F0 int16; F1 int16 }

func main() {
    v1 := []uint32{318, 2147483647, 42}
    var v2 int32 = 42
    v3 := make(map[string]uint32)
    var v4 uint8 = 2
    var v5 uint16 = 32767
    for fi6 := 0; fi6 < 16; fi6++ {
        if false {
            if len(v1) > 0 {
                v1[0] = 1
            }
            if (4294967295 < 18446744073709551615) {
                v2 = 0
                for fi7 := 0; fi7 < 2; fi7++ {
                    _ = len(v3)
                    var v8 float64 = (-2.25 - 0.0)
                    _ = fi7
                    _ = v8
                }
                if len(v1) > 0 {
                    v1[0] = 0
                }
            }
        }
        _ = fi6
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 float32 = -2.25
        if (false && (!((^(v2)) == v2))) {
            v5 += v5
            var v11 uint = 50
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    if (2147483647 >= 42) {
        var v12 int8 = -1
        fmt.Println(v12)
        switch 128 {
        case 122:
            switch "\n\t" {
            case "":
                v5 += 306
            case "foo bar":
                var v13 uint64 = 789
                _ = v13
            default:
                v4 = v4
            }
        case 932:
            var v14 bool = (false || false)
            _ = v14
        case -912:
            for fi15 := 0; fi15 < 18; fi15++ {
                for fi16 := 0; fi16 < 6; fi16++ {
                    if (false && true) {
                        var v17 int8 = v12
                        v3["\n\t"] = 907
                        var v18 uint32 = 2
                        _ = v17
                        _ = v18
                    }
                    var v19 int = fi16
                    _ = fi16
                    _ = v19
                }
                for fi20 := 0; fi20 < 13; fi20++ {
                    v21 := GenS1{F0: 1, F1: 922, F2: 1}
                    v12 = (v12 | v12)
                    _ = fi20
                    _ = v21
                }
                _ = fi15
            }
        }
        _ = v12
    } else {
        if false {
            v22 := GenS1{F0: 0, F1: 326, F2: -1}
            switch 42 {
            case 127:
                v23 := GenS2{F0: 1, F1: 404}
                _ = v23
            case -525:
                _ = len(v1)
            default:
                var v24 string = ""
                _ = v24
            }
            _ = v22
        }
        var v25 uint8 = v4
        _ = v25
    }
    if ((false || (-263 >= -9223372036854775808)) || ((v5 <= (v5 | v5)) || ((float32(0.1) / 1.0) > (float32(-2.25) / ((3.14159 - -2.25) / 3.14159))))) {
        v5 -= v5
    } else {
        v2 = -2147483648
        var v26 float64 = ((0.0 * 3.14159) + (1.5 - 1.5))
        v1 = append(v1, 944)
        _ = v26
    }
    var v27 uint8 = (v4 * (^(v4)))
    v27 = (v27 >> ((0) & 63))
    if ("\n\t" > ("hello" + "a")) {
        v4 += v4
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
