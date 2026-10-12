package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int8; F2 uint }
type GenS4 struct { F0 int8; F1 int32; F2 int8 }
type GenS5 struct { F0 bool; F1 int32 }
func genFn0() uint8 {
    var v6 int64 = -412
    v7 := []int{32767, 2147483647, 289, -141}
    if true {
        v8 := GenS1{F0: "", F1: 18446744073709551615}
        v6 += v6
        _ = v8
    }
    v6 -= (v6 | (v6 ^ v6))
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 int = fi9
        fmt.Println(v10, v6, fi9)
        var v11 float32 = (((1.5 + 1.5) / 1.5) / 1.0)
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = v6
    _ = v7
    return 255
}

func main() {
    var v1 int8 = 100
    v2 := []int{781, 2147483647, 42, 9223372036854775807}
    var v3 float32 = float32(0.1)
    v4 := []int{9223372036854775807, 686}
    v5 := make(map[string]string)
    var v12 uint64 = 433
    if true {
        var v13 string = "hello"
        v14 := GenS1{F0: v13, F1: v12}
        var v15 float64 = float64((1.5 / v3))
        _ = v13
        _ = v14
        _ = v15
    } else {
        fmt.Println(v12, v12, v12)
        fmt.Println(v3, v3, v12)
        var v16 int = 0
        _ = v16
    }
    v17 := GenS2{F0: 1}
    var v18 uint8 = 76
    if true {
        switch "" {
        case "foo bar":
            v19 := GenS3{F0: 32767, F1: v1, F2: 535}
            _ = v19
        default:
            if false {
                v3 += v3
                var v20 int32 = -2147483648
                _ = v20
            }
        }
        var v21 int = 624
        if (true && true) {
            var v22 uint = 2
            _ = v22
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(v18, v12)
            if (!false) {
                switch "foo bar" {
                case "foo bar":
                    var v23 int8 = (v1 * int8(v21))
                    _ = v23
                default:
                    v24 := GenS3{F0: 538, F1: v1, F2: 2}
                    _ = v24
                }
            }
        }
        _ = v21
    } else {
        for fi25 := 0; fi25 < 11; fi25++ {
            v12 -= (^(v12))
            var v26 float32 = v3
            v27 := GenS4{F0: -71, F1: 2147483647, F2: 39}
            _ = fi25
            _ = v26
            _ = v27
        }
        if true {
            switch -9223372036854775808 {
            case 0:
                var v28 int = 632
                _ = v28
            case 128:
                v3 = (v3 / v3)
            }
        }
    }
    if (!(!(1e10 > (float64(v3) - 3.14159)))) {
        switch 128 {
        case 1:
            switch ("foo bar" + "\n\t") {
            case "hello":
                v3 += v3
            case "a":
                v29 := GenS5{F0: (1000 == 18446744073709551615), F1: 0}
                _ = v29
            }
        }
    }
    v30 := v5[""]
    fmt.Println(v18, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
