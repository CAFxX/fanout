package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int8; F1 int64; F2 int }

func main() {
    v1 := []int{}
    var v2 string = "a"
    var v3 int64 = 0
    if true {
        _ = len(v1)
        v2 = v2
    } else {
        v4 := GenS1{F0: -317, F1: (0 < 572)}
        fmt.Println(v3, v2)
        _ = v4
    }
    v3 += (v3 << ((1) & 63))
    if (float64(1e10) >= 0.1) {
        var v5 int64 = (v3 % (((v3 ^ v3)) | 1))
        v3 = 643
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v5
    } else {
        var v6 string = v2
        for fi7 := 0; fi7 < 6; fi7++ {
            v6 = ("foo bar" + "a")
            _ = fi7
        }
        v8 := GenS1{F0: 9223372036854775807, F1: false}
        _ = v6
        _ = v8
    }
    var v9 int32 = 103
    var v10 string = ((v2 + (((v2 + "\n\t") + v2) + (v2 + v2))) + "")
    v10 = ((v2 + v2) + "")
    v9 -= v9
    v3 = 9223372036854775807
    for fi11 := 0; fi11 < 3; fi11++ {
        v2 = v10
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 uint = 18446744073709551615
            if true {
                v9 = 0
                _ = len(v1)
                switch v2 {
                case "foo bar":
                    fmt.Println(v3, fi11, fi12)
                    fmt.Println(v2, v10, v9)
                    v9 -= 837
                default:
                    var v14 float64 = 1e10
                    _ = v14
                }
                _ = len(v1)
                v15 := GenS2{F0: -78, F1: v3, F2: fi11}
                _ = len(v1)
                _ = v15
            }
            v13 -= 65535
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v9
}
