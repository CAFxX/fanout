package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint; F2 uint32 }
type GenS4 struct { F0 uint64; F1 uint64; F2 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = (0.1 / ((-2.25 / 1.5) + (1e10 / (((0.1 + 1e10) / ((3.14159 + (0.0 / 1.0)) * 0.1)) / 1e10))))
    v3 := []uint32{}
    var v4 uint16 = 58
    var v5 uint8 = 1
    v4 -= 267
    var v6 float64 = v2
    switch -2147483648 {
    case -2147483648:
        fmt.Println(v4, v2, v6)
    case 32767:
        fmt.Println(v5, v2, v6)
    default:
        if (32767 > 32767) {
            if false {
                v4 = (^(v4))
                v2 += (v2 / v2)
            }
        } else {
            var v7 uint = 965
            for fi8 := 0; fi8 < 8; fi8++ {
                if (127 != 42) {
                    v9 := GenS1{F0: "\n\t", F1: "\n\t", F2: 341}
                    _ = v9
                }
                v10 := GenS1{F0: "\n\t", F1: "", F2: 0}
                _ = fi8
                _ = v10
            }
            _ = v7
        }
    }
    switch "hello" {
    case "":
        for fi11 := 0; fi11 < 8; fi11++ {
            v3 = append(v3, 2147483648)
            var v12 uint = 42
            var v13 int64 = -9223372036854775808
            _ = fi11
            _ = v12
            _ = v13
        }
    case "\n\t":
        fmt.Println(v4, v2)
    case "hello":
        var v14 float64 = 1.5
        _ = v14
    }
    v15 := GenS1{F0: ("" + "a"), F1: "\n\t", F2: 1}
    for fi16 := 0; fi16 < 14; fi16++ {
        v4 -= v4
        v17 := GenS2{F0: -767}
        v4 = (^(v4))
        _ = fi16
        _ = v17
    }
    v18 := GenS3{F0: 127, F1: 18446744073709551615, F2: 714}
    if (true || (false || (float32(((1.5 - (0.1 - v6)) - (v2 / ((v6 + 1.5) * (-2.25 + (v6 + (1e10 * v6))))))) < ((1.5 / 1.0) * 0.1)))) {
        _ = len(v3)
        if (!(!(!false))) {
            v5 = (v5 & (^(v5)))
        }
    } else {
        v19 := GenS1{F0: "", F1: "a", F2: 42}
        var v20 string = ((("a" + "") + "hello") + "hello")
        v21 := GenS4{F0: 619, F1: 42, F2: (v4 & v4)}
        _ = v19
        _ = v20
        _ = v21
    }
    v2 = (((((0.1 - v2) - v6) + (((v2 / 1.0) + (-2.25 / (v6 / 1.5))) - v2)) - v2) + (v2 - (((float64(0.0) + float64(float32(float64(1e10)))) * float64(0.1)) - 3.14159)))
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 float64 = (v6 - float64(3.14159))
        var v24 uint8 = 6
        if (!false) {
            var v25 float32 = -2.25
            _ = v25
        }
        _ = fi22
        _ = v23
        _ = v24
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v26 uint = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
