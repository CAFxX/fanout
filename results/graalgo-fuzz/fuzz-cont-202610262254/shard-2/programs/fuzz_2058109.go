package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 string; F1 int64; F2 int8 }
type GenS3 struct { F0 uint8; F1 int8; F2 uint32 }
func genFn0() uint {
    var v5 int = 451
    switch "foo bar" {
    case "a":
        v6 := GenS1{F0: "hello", F1: 255}
        _ = v6
    case "foo bar":
        v7 := []int{}
        _ = v7
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int64 = -1
        _ = fi8
        _ = v9
    }
    var v10 uint = 889
    var v11 uint = (^(v10))
    _ = v10
    _ = v11
    _ = v5
    return v11
}

func main() {
    v1 := make(map[int]string)
    var v2 int32 = 140
    v3 := make(map[string]string)
    var v4 float64 = (((0.1 * (((1.5 + float64(((1.5 * -2.25) + 1.5))) + 0.1) * 0.0)) - 0.1) + ((float64(1.5) / -2.25) * float64((1.5 / 3.14159))))
    var v12 int64 = 224
    switch 9223372036854775807 {
    case 1:
        if (true && false) {
            v13 := GenS2{F0: "foo bar", F1: (v12 & (v12 << ((2) & 63))), F2: 1}
            switch -562 {
            case -9223372036854775808:
                if false {
                    var v14 bool = false
                    _ = v14
                } else {
                    _ = len(v3)
                    v4 = float64((((1e10 / 3.14159) * -2.25) - ((3.14159 + 1.5) * (float32((v4 * 1e10)) - 0.0))))
                    v4 -= float64(0.0)
                }
            }
            _ = v13
        }
    }
    v3["foo bar"] = "\n\t"
    if (!false) {
        if ((255 > 18446744073709551615) && true) {
            for fi15 := 0; fi15 < 5; fi15++ {
                var v16 uint64 = 803
                _ = fi15
                _ = v16
            }
        }
        var v17 int16 = 1000
        _ = v17
    } else {
        v18 := GenS1{F0: "\n\t", F1: 905}
        if true {
            v19 := GenS1{F0: "", F1: 2}
            v4 = v4
            v20 := GenS3{F0: 76, F1: 58, F2: 1}
            _ = v19
            _ = v20
        } else {
            v12 = 630
            v12 += (v12 >> ((4294967295) & 63))
            v21 := []int{32767}
            _ = v21
        }
        v2 = v2
        _ = v18
    }
    fmt.Println(v2, v12)
    v22 := v1[-180]
    for fi23 := 0; fi23 < 14; fi23++ {
        fmt.Println(v12, v12)
        _ = fi23
    }
    if (-2.25 <= (v4 + (3.14159 / v4))) {
        var v24 int64 = 1
        switch "" {
        case "\n\t":
            v24 += 0
        case "foo bar":
            fmt.Println(v24, v24, v2)
        default:
            v3[v22] = v22
        }
        _ = v24
    } else {
        var v25 float64 = v4
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
