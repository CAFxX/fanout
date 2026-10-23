package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 bool; F1 int; F2 bool }
func genFn0() uint {
    v4 := []int{}
    switch 709 {
    case -206:
        var v5 uint = 204
        _ = v5
    }
    switch -9223372036854775808 {
    case 81:
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 uint32 = 2147483648
            switch (fi6 & (fi6 | 0)) {
            case 32768:
                fmt.Println(fi6, v7)
            case 469:
                var v8 int16 = -36
                _ = v8
            case 9223372036854775807:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = fi6
            _ = v7
        }
    case 32767:
        var v9 float32 = ((0.1 / 0.1) / -2.25)
        _ = v9
    case 9223372036854775807:
        if ((1 <= 2147483647) || true) {
            v10 := GenS1{F0: (true || (0 != 2147483647))}
            var v11 int16 = 42
            _ = v10
            _ = v11
        } else {
            var v12 bool = (!(0 <= -677))
            v13 := GenS2{F0: -1, F1: ""}
            _ = len(v4)
            _ = v12
            _ = v13
        }
    default:
        var v14 string = "foo bar"
        _ = v14
    }
    var v15 uint = 541
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 uint = (v15 ^ (^(v15)))
        fmt.Println(v15, fi16, fi16)
        _ = fi16
        _ = v17
    }
    _ = v15
    _ = v4
    return v15
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 2
    var v3 uint8 = (v2 / ((v2) | 1))
    switch ((("hello" + "") + "") + "\n\t") {
    case "foo bar":
        v2 -= (v3 << ((841) & 63))
    case "a":
        v18 := GenS1{F0: (((true && (2147483647 != 2147483648)) && (!(!(uint(v3) < genFn0())))) || (genFn0() != 3))}
        _ = v18
    case "\n\t":
        var v19 int = 2147483647
        _ = v19
    default:
        v2 = 174
    }
    v20 := GenS3{F0: (42 != 0), F1: 32768, F2: (2 > 876)}
    fmt.Println(v3, v2)
    _ = len(v1)
    v3 += 128
    var v21 uint64 = 4294967295
    var v22 uint32 = 2
    v3 = v2
    var v23 float64 = 3.14159
    v23 = ((v23 / (v23 / v23)) + float64(0.0))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
