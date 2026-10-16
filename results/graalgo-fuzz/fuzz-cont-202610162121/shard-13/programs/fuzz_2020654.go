package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 1e10
    var v3 uint32 = 129
    v4 := []string{("\n\t" + "a"), "foo bar"}
    var v5 uint32 = v3
    for fi6 := 0; fi6 < 12; fi6++ {
        if (!true) {
            v2 = ((((0.1 + v2) / v2) * v2) - v2)
            var v7 uint8 = 127
            v2 = ((((v2 * v2) / (((v2 - (0.0 + v2)) * v2) / -2.25)) - ((v2 * v2) + 3.14159)) / v2)
            _ = v7
        }
        v8 := GenS1{F0: -9223372036854775808, F1: 879, F2: 1000}
        if false {
            fmt.Println(v5, v2, v2)
        } else {
            var v9 float32 = 0.0
            v10 := GenS1{F0: 0, F1: 18446744073709551615, F2: -1}
            _ = v10
            _ = v9
        }
        _ = fi6
        _ = v8
    }
    v2 -= -2.25
    fmt.Println(v5, v3, v5)
    switch 2147483647 {
    case 128:
        v5 = 65535
    case -247:
        v2 += v2
    case -390:
        v4 = append(v4, "\n\t")
    default:
        var v11 int16 = 32767
        _ = v11
    }
    var v12 int32 = -2147483648
    if (!(9223372036854775807 <= 42)) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v13 int = 433
    fmt.Println(v5, v5, v13)
    v2 += float32(-2.25)
    if ("a" <= "hello") {
        var v14 string = "hello"
        v15 := GenS2{F0: -1}
        v1["foo bar"] = v3
        _ = v14
        _ = v15
    } else {
        v2 = 0.1
        if (!(2 >= 18446744073709551615)) {
            v16 := GenS3{F0: "\n\t", F1: -1}
            _ = v16
        }
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 uint8 = 42
            fmt.Println(v13, fi17)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi17
            _ = v18
        }
    }
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
