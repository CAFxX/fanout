package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint }
type GenS2 struct { F0 int32; F1 uint; F2 uint64 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int8; F1 uint; F2 int16 }

func main() {
    v1 := []string{}
    v2 := make(map[string]bool)
    var v3 float64 = 3.14159
    v4 := []uint32{42}
    if true {
        switch "\n\t" {
        case "\n\t":
            var v5 string = "foo bar"
            _ = v5
        default:
            fmt.Println(v3)
        }
        if false {
            switch 1 {
            case -458:
                switch 127 {
                case 9223372036854775807:
                    var v6 int16 = 1
                    _ = v6
                case -880:
                    v3 -= (((v3 / v3) - (float64(0.1) + (v3 * v3))) - ((0.0 - 1.5) / v3))
                }
            default:
                v7 := GenS1{F0: 2, F1: 323, F2: 65535}
                _ = v7
            }
            fmt.Println(v3)
        }
        fmt.Println(v3)
    }
    fmt.Println(v3)
    v3 += 0.0
    v3 += 3.14159
    v8 := GenS2{F0: -2147483648, F1: 776, F2: 353}
    v3 -= (v3 + (0.1 + 3.14159))
    var v9 int64 = 1
    if false {
        fmt.Println(v3)
        v3 = ((float64(3.14159) - (v3 / v3)) + ((((v3 * v3) / 3.14159) / (float64((-2.25 / -2.25)) + (v3 * 0.1))) / ((1.5 - 0.1) / float64(3.14159))))
        v10 := GenS3{F0: 1, F1: 1}
        _ = v10
    } else {
        var v11 float64 = v3
        var v12 int64 = (^(v9))
        for fi13 := 0; fi13 < 9; fi13++ {
            for fi14 := 0; fi14 < 11; fi14++ {
                switch (("hello" + "") + ("a" + ("" + "a"))) {
                case "\n\t":
                    v15 := GenS4{F0: 100, F1: 4294967295, F2: 32767}
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v15
                case "a":
                    v16 := GenS4{F0: -123, F1: 0, F2: -32768}
                    _ = v16
                case "":
                    var v17 uint64 = 6
                    _ = v17
                }
                _ = fi14
            }
            fmt.Println(v3, v9)
            _ = fi13
        }
        _ = v11
        _ = v12
    }
    v18 := v2["foo bar"]
    if v18 {
        if (v18 && v18) {
            v3 = float64(float32(1.5))
            v9 += v9
        }
        delete(v2, ("hello" + "hello"))
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        v18 = (42 > 2147483647)
        var v20 uint64 = 208
        _ = fi19
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
