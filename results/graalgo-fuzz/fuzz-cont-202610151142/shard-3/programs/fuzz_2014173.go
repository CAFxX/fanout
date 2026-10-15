package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int16; F2 int16 }

func main() {
    var v1 uint16 = 605
    var v2 int16 = -1
    v3 := []int{127, 797}
    var v4 uint16 = 65535
    var v5 uint16 = v4
    var v6 string = ""
    v7 := GenS1{F0: 42}
    v8 := GenS2{F0: 1, F1: (v2 % (((-(v2))) | 1)), F2: 42}
    var v9 float64 = ((((0.0 / 0.1) + float64(1e10)) + 3.14159) / 0.1)
    var v10 uint8 = 1
    v9 = ((1.5 - v9) * (v9 - ((v9 + float64(3.14159)) + (v9 + (3.14159 + (1.5 * v9))))))
    if true {
        switch "foo bar" {
        case "":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "foo bar":
            fmt.Println(v9, v6)
        default:
            v2 += v2
        }
        switch -1 {
        case -9223372036854775808:
            v10 += 128
        }
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi11
    }
    v12 := GenS2{F0: 18446744073709551615, F1: v2, F2: -1}
    v13 := GenS2{F0: 801, F1: (v2 ^ 1), F2: (v2 >> ((676) & 63))}
    var v14 int = 590
    for fi15 := 0; fi15 < 20; fi15++ {
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 bool = (((!(!false)) && true) && ((v10 & v10) < (^(v10))))
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    fmt.Println(v10, v4)
    var v18 bool = true
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
