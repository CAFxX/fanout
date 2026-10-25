package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint64; F2 string }

func main() {
    var v1 uint8 = 255
    var v2 int64 = 645
    v3 := []string{}
    var v4 int32 = 0
    switch -9223372036854775808 {
    case 587:
        if (32767 <= 200) {
            fmt.Println(v2, v4)
        } else {
            if (!false) {
                v3 = append(v3, "\n\t")
                switch "hello" {
                case "foo bar":
                    var v5 uint32 = 65535
                    _ = v5
                default:
                    _ = len(v3)
                }
            }
            v6 := GenS1{F0: "foo bar", F1: 880, F2: 1}
            _ = v6
        }
    case 4294967295:
        v3 = append(v3, "")
    case 9223372036854775807:
        if (100 == 71) {
            var v7 uint = 258
            v1 = 0
            var v8 int8 = 52
            _ = v7
            _ = v8
        }
    }
    if ((3.14159 * 1e10) > ((1e10 / (3.14159 / float32(1e10))) / (0.1 - 0.0))) {
        fmt.Println(v1, v2)
        var v9 int = -2147483648
        _ = v9
    } else {
        var v10 int16 = 623
        if (false || false) {
            switch 762 {
            case 83:
                v10 = v10
            }
            v11 := GenS1{F0: "a", F1: 255, F2: (v10 % ((v10) | 1))}
            _ = v11
        } else {
            v12 := GenS2{F0: ("foo bar" + "\n\t")}
            _ = v12
        }
        _ = v10
    }
    v1 = v1
    v4 = v4
    if (!(true && (!true))) {
        var v13 int8 = -128
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 float64 = 1e10
            switch "a" {
            case "":
                _ = len(v3)
            case "a":
                var v16 uint16 = 570
                _ = v16
            default:
                var v17 string = "foo bar"
                _ = v17
            }
            _ = fi14
            _ = v15
        }
        var v18 int32 = v4
        _ = v13
        _ = v18
    } else {
        if false {
            var v19 float32 = (1e10 - 0.1)
            _ = v19
        } else {
            v4 -= 42
            v20 := GenS3{F0: 65535, F1: 2, F2: "\n\t"}
            _ = v20
        }
    }
    var v21 int64 = (v2 ^ (v2 + (v2 & v2)))
    v3 = append(v3, "a")
    var v22 int64 = v2
    v4 += v4
    var v23 float64 = (1e10 / 1.5)
    for fi24 := 0; fi24 < 10; fi24++ {
        if len(v3) > 0 {
            v3[0] = ("\n\t" + "hello")
        }
        v21 = (v21 % ((v2) | 1))
        fmt.Println(fi24, v1, v2)
        _ = fi24
    }
    v4 -= (v4 & int32(v22))
    var v25 uint8 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
}
