package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint16; F1 int16; F2 string }

func main() {
    var v1 float32 = (0.0 * 0.1)
    var v2 int32 = 2147483647
    v3 := make(map[string]string)
    v4 := make(map[int]uint32)
    v4[-2147483648] = 256
    fmt.Println(v1, v1)
    fmt.Println(v2, v2)
    v5 := GenS1{F0: 0}
    v2 += v2
    switch 0 {
    case -9223372036854775808:
        v6 := GenS2{F0: -9223372036854775808, F1: 4294967295, F2: 0}
        _ = v6
    case -2147483648:
        v1 = 0.1
    default:
        v2 += -1
    }
    var v7 int8 = -83
    v8 := GenS3{F0: 968, F1: 32767, F2: ("foo bar" + "\n\t")}
    if false {
        var v9 uint64 = 2
        v9 = v9
        _ = v9
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        fmt.Println(v7)
        _ = fi10
    }
    if (!(288 >= 559)) {
        if ("foo bar" != "") {
            var v11 uint = 42
            for fi12 := 0; fi12 < 10; fi12++ {
                v1 = (v1 / v1)
                v2 -= v2
                _ = fi12
            }
            _ = v11
        }
        v7 -= v7
        v2 = (v2 + 1)
    }
    if true {
        var v13 uint16 = 0
        _ = v13
    } else {
        v4[-1] = 2147483647
        if (42 == 2147483648) {
            v7 = 124
        } else {
            v14 := []int{}
            _ = v14
        }
        fmt.Println(v7, v1)
    }
    switch ("foo bar" + "hello") {
    case "hello":
        for fi15 := 0; fi15 < 10; fi15++ {
            v16 := GenS1{F0: 18446744073709551615}
            _ = fi15
            _ = v16
        }
    default:
        v1 -= (1.5 - ((((((3.14159 / (((v1 / (v1 * float32(((float64(0.1) / 1.0) / -2.25)))) * (float32((1e10 - 0.1)) * 0.0)) + (v1 / v1))) + float32(0.0)) + v1) - v1) + 0.0) / v1))
    }
    var v17 uint = 255
    v1 -= ((v1 + 0.0) * v1)
    var v18 int64 = 0
    var v19 uint16 = 32767
    v20 := GenS3{F0: v19, F1: 42, F2: ""}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
