package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint64; F2 uint32 }

func main() {
    var v1 int8 = 0
    v2 := []string{"a", "\n\t"}
    var v3 uint16 = 42
    var v4 int64 = -703
    if (((v1 == v1) || false) && ((2 > 158) || ((!(4294967295 != 4294967295)) && true))) {
        v4 += 0
    }
    _ = len(v2)
    var v5 int16 = 351
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 uint64 = 4294967295
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := GenS1{F0: v3, F1: v3, F2: (v3 >> ((190) & 63))}
            switch "foo bar" {
            case "hello":
                switch fi6 {
                case 2147483647:
                    v10 := GenS1{F0: (v3 + v3), F1: 42, F2: v3}
                    _ = v10
                }
                v11 := GenS1{F0: 744, F1: (v3 | (v3 / ((v3) | 1))), F2: (v3 & 2)}
                _ = v11
            case "\n\t":
                var v12 bool = ((v7 + (v7 & 401)) <= v7)
                _ = v12
            case "foo bar":
                v1 += (-(v1))
                var v13 float64 = (0.1 - float64(3.14159))
                _ = v13
            default:
                v4 = 42
            }
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v14 int32 = 42
    v3 -= v3
    var v15 int = 302
    var v16 float64 = 0.1
    v4 -= v4
    var v17 int8 = 127
    v16 += 3.14159
    if (v1 == v17) {
        switch 9223372036854775807 {
        case -971:
            v17 -= (v17 * (v1 + v17))
        case 32767:
            v16 += 0.1
        case 2147483647:
            var v18 string = "foo bar"
            _ = v18
        }
        fmt.Println(v1, v3)
    } else {
        var v19 int64 = v4
        fmt.Println(v17, v19, v3)
        for fi20 := 0; fi20 < 12; fi20++ {
            fmt.Println(v15, v14, v3)
            _ = fi20
        }
        _ = v19
    }
    if false {
        v21 := GenS2{F0: false, F1: 18446744073709551615, F2: 1}
        for fi22 := 0; fi22 < 4; fi22++ {
            var v23 int16 = v5
            for fi24 := 0; fi24 < 1; fi24++ {
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
                var v25 int64 = int64(v23)
                _ = fi24
                _ = v25
            }
            v14 = (v14 * v14)
            _ = fi22
            _ = v23
        }
        _ = v21
    }
    var v26 uint16 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
