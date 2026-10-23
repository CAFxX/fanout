package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int }

func main() {
    var v1 int16 = -165
    var v2 int64 = -9223372036854775808
    v3 := []int{-729, 610, 127}
    v4 := make(map[uint32]string)
    var v5 int16 = v1
    fmt.Println(v5, v1, v1)
    v5 += (v5 - (v1 ^ -106))
    v1 -= v5
    var v6 int32 = 0
    if len(v3) > 0 {
        v3[0] = 9223372036854775807
    }
    fmt.Println(v5, v1, v5)
    v6 -= v6
    fmt.Println(v6, v2, v1)
    switch ("\n\t" + ("foo bar" + "\n\t")) {
    case "foo bar":
        var v7 bool = (!false)
        _ = v7
    case "":
        var v8 int32 = (v6 ^ int32(v5))
        _ = v8
    default:
        v1 -= (v5 / ((v5) | 1))
    }
    var v9 uint64 = 42
    v10 := GenS1{F0: 32767, F1: (v5 | 1000), F2: 689}
    v9 -= 811
    if (((-2.25 * 0.1) * (1.5 - 3.14159)) > -2.25) {
        v2 = (v2 / ((v2) | 1))
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 int = fi11
            _ = fi11
            _ = v12
        }
    } else {
        if ("\n\t" <= "a") {
            switch 281 {
            case 597:
                var v13 uint = 91
                _ = v13
            case 4294967295:
                v14 := GenS1{F0: v1, F1: (v5 % ((v1) | 1)), F2: -202}
                _ = v14
            case 127:
                if true {
                    v15 := GenS1{F0: 0, F1: v1, F2: -761}
                    _ = len(v4)
                    var v16 int8 = 42
                    _ = v15
                    _ = v16
                }
            }
        } else {
            v17 := GenS1{F0: v5, F1: v1, F2: 32767}
            var v18 uint = 0
            v1 -= (v1 | -428)
            _ = v17
            _ = v18
        }
    }
    if false {
        v1 = (v1 - v1)
        var v19 uint32 = 429
        _ = v19
    } else {
        _ = len(v3)
    }
    if false {
        v5 = -545
    } else {
        v20 := GenS1{F0: 1000, F1: -1, F2: int(v6)}
        var v21 int64 = -9223372036854775808
        _ = v20
        _ = v21
    }
    v9 = v9
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
