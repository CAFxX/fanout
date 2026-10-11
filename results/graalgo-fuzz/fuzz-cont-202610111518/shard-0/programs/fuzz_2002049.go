package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{728}
    v3 := make(map[uint32]int)
    var v4 float64 = -2.25
    v4 = float64((1.5 * (float32((v4 * v4)) + 0.0)))
    if len(v2) > 0 {
        v2[0] = 418
    }
    fmt.Println(v4)
    var v5 uint32 = 2
    fmt.Println(v4, v5)
    if (-9223372036854775808 == -327) {
        switch ("a" + (("hello" + ("hello" + "\n\t")) + "hello")) {
        case "hello":
            v6 := GenS1{F0: 2, F1: 18446744073709551615}
            _ = v6
        }
    } else {
        for fi7 := 0; fi7 < 11; fi7++ {
            v8 := v1[fi7]
            _ = fi7
            _ = v8
        }
    }
    fmt.Println(v4, v5)
    v5 += 65535
    v2 = append(v2, (v5 % ((v5) | 1)))
    if (!(true || (!true))) {
        fmt.Println(v5, v5)
    }
    if true {
        if false {
            var v9 uint = 1
            v4 -= float64((-2.25 / (float32(0.1) * 1.5)))
            switch -652 {
            case -883:
                var v10 int32 = 42
                _ = v10
            case 32767:
                v11 := GenS1{F0: 42, F1: 588}
                _ = v11
            case 1:
                fmt.Println(v5, v9)
            default:
                if len(v2) > 0 {
                    v2[0] = 2
                }
            }
            _ = v9
        } else {
            fmt.Println(v4, v5)
        }
        var v12 float32 = -2.25
        v5 = 714
        _ = v12
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
