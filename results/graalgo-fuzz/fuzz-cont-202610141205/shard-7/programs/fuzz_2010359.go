package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint16 }
type GenS2 struct { F0 uint8; F1 bool; F2 uint16 }

func main() {
    v1 := []int{-1}
    var v2 int32 = 99
    var v3 string = ("a" + "")
    var v4 uint16 = 1
    v3 = (v3 + ((v3 + v3) + v3))
    v4 += (^(v4))
    var v5 bool = true
    if (((true || (1e10 == ((((1e10 / (((3.14159 * (((-2.25 + 1.5) / ((3.14159 * 3.14159) - 3.14159)) / 1e10)) / 1.0) * (0.0 - 1e10))) / float64((0.1 + (0.0 - float32(0.0))))) * -2.25) / (1e10 / 1.0)))) && (811 >= 65535)) && v5) {
        fmt.Println(v3, v3, v5)
        v2 += v2
    } else {
        fmt.Println(v5, v3)
    }
    for fi6 := 0; fi6 < 1; fi6++ {
        if (v5 && (3.14159 >= ((3.14159 / -2.25) / -2.25))) {
            switch (fi6 / (((-(fi6))) | 1)) {
            case -177:
                var v7 int8 = 100
                _ = v7
            case 32768:
                var v8 uint8 = 1
                _ = v8
            default:
                var v9 int32 = (v2 % (((v2 + (v2 ^ v2))) | 1))
                _ = v9
            }
            var v10 bool = v5
            _ = v10
        } else {
            v1 = append(v1, -939)
            v11 := GenS1{F0: 42, F1: "\n\t", F2: 32767}
            _ = v11
        }
        _ = fi6
    }
    v12 := GenS1{F0: 69, F1: v3, F2: (v4 - v4)}
    var v13 int64 = int64(v2)
    for fi14 := 0; fi14 < 11; fi14++ {
        if v5 {
            if v5 {
                v15 := GenS1{F0: 42, F1: v3, F2: 2}
                fmt.Println(v13, v2)
                _ = v15
            } else {
                switch "hello" {
                case "foo bar":
                    v16 := GenS2{F0: 2, F1: (-1 <= (v13 + v13)), F2: v4}
                    _ = v16
                case "\n\t":
                    v5 = (!((0.1 + float64(0.1)) <= (float64((1.5 - 0.0)) / 0.1)))
                case "a":
                    var v17 uint = 0
                    _ = v17
                }
            }
        }
        _ = fi14
    }
    if ((!(float32(3.14159) >= 1.5)) || ((-(v13)) != (^(v13)))) {
        v18 := GenS1{F0: 127, F1: v3, F2: 32767}
        _ = v18
    } else {
        fmt.Println(v5, v2, v2)
        v19 := GenS2{F0: 127, F1: v5, F2: 1}
        var v20 float64 = (1e10 - 1.5)
        _ = v19
        _ = v20
    }
    var v21 uint16 = (v4 % ((v4) | 1))
    for fi22 := 0; fi22 < 13; fi22++ {
        if (1.5 == 1.5) {
            var v23 int64 = v13
            _ = v23
        }
        var v24 int16 = 32767
        _ = fi22
        _ = v24
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v25 int64 = v13
    var v26 float64 = ((((1e10 - (0.1 - (float64(0.0) / 3.14159))) + -2.25) - 0.1) / 0.1)
    v13 = -9223372036854775808
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
