package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint8 }
type GenS2 struct { F0 int32; F1 bool; F2 int64 }

func main() {
    var v1 uint64 = 42
    var v2 float64 = 3.14159
    var v3 uint8 = 194
    switch ((((("\n\t" + "") + "a") + (("" + (("\n\t" + "") + (("a" + ((("" + "foo bar") + ((("" + "a") + (("a" + ("foo bar" + "")) + "")) + "a")) + "")) + "a"))) + "")) + ("foo bar" + "foo bar")) + "a") {
    case "a":
        v3 -= 127
    case "hello":
        v4 := GenS1{F0: 1, F1: int16(v1), F2: (v3 - v3)}
        _ = v4
    case "foo bar":
        if (true && ((255 < 2) && (false && (42 != 4294967295)))) {
            for fi5 := 0; fi5 < 17; fi5++ {
                fmt.Println(fi5, v2)
                v3 += 255
                var v6 uint32 = 1
                _ = fi5
                _ = v6
            }
        } else {
            fmt.Println(v3)
            var v7 int64 = 806
            _ = v7
        }
    }
    var v8 int8 = -1
    var v9 uint16 = 939
    var v10 int8 = v8
    var v11 uint16 = 1
    v12 := []int{32767, 55}
    v9 += v9
    var v13 float32 = 1e10
    v10 += v8
    if (677 > 32768) {
        switch 2147483647 {
        case 4294967295:
            fmt.Println(v11, v1)
        case -475:
            var v14 int16 = 42
            _ = v14
        default:
            v15 := GenS1{F0: 9223372036854775807, F1: 42, F2: 128}
            _ = v15
        }
        v1 -= 42
    } else {
        v9 += 707
        var v16 bool = false
        var v17 float64 = (v2 / v2)
        _ = v16
        _ = v17
    }
    v2 -= (v2 / (v2 + 1.5))
    for fi18 := 0; fi18 < 4; fi18++ {
        fmt.Println(fi18, v2)
        _ = fi18
    }
    fmt.Println(v13, v3)
    v3 = (^(v3))
    v1 = v1
    switch 4294967295 {
    case 32767:
        v9 = v11
    case -632:
        for fi19 := 0; fi19 < 19; fi19++ {
            var v20 float32 = ((((((v13 - v13) / (v13 / 1.0)) + ((1e10 * (float32(0.0) * float32(v2))) / v13)) / v13) * v13) + (v13 + 0.0))
            v21 := GenS2{F0: 0, F1: ((true && true) && (42 > -424)), F2: 1}
            if len(v12) > 0 {
                v12[0] = (^(fi19))
            }
            _ = fi19
            _ = v20
            _ = v21
        }
    case -9223372036854775808:
        for fi22 := 0; fi22 < 6; fi22++ {
            var v23 uint = 42
            fmt.Println(v3, v2)
            _ = fi22
            _ = v23
        }
    }
    v3 = 223
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
