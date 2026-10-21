package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int }

func main() {
    v1 := []string{("hello" + ""), (("a" + "a") + ""), "\n\t", "foo bar"}
    var v2 int16 = 1000
    v3 := make(map[int]int)
    var v4 int8 = 100
    var v5 int = 286
    var v6 int32 = 186
    v5 = v5
    switch (v5 * (v5 ^ (v5 - int(v2)))) {
    case 716:
        v3[v5] = v5
    case 32768:
        if (false && ((!(false || false)) && (v6 <= v6))) {
            fmt.Println(v4, v6)
            for fi7 := 0; fi7 < 17; fi7++ {
                if len(v1) > 0 {
                    v1[0] = ("hello" + "foo bar")
                }
                var v8 int = (fi7 >> ((255) & 63))
                _ = fi7
                _ = v8
            }
            v9 := GenS1{F0: 0}
            _ = v9
        }
    case 42:
        var v10 int8 = (v4 ^ v4)
        _ = v10
    default:
        var v11 int64 = 424
        _ = v11
    }
    v12 := GenS2{F0: v6, F1: (v6 % ((v6) | 1))}
    _ = len(v1)
    v5 -= (-(v5))
    v5 -= v5
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := GenS1{F0: -9223372036854775808}
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 int = (v5 << ((18446744073709551615) & 63))
            switch "\n\t" {
            case "foo bar":
                fmt.Println(fi15, v6, v6)
            case "\n\t":
                v17 := GenS3{F0: v5}
                _ = v17
            }
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    var v18 int32 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
