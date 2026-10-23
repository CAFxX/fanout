package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS2 struct { F0 string; F1 int8; F2 string }

func main() {
    var v1 int32 = 628
    v2 := []string{}
    v3 := make(map[int]bool)
    var v4 float32 = float32(0.1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v5 int8 = 0
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS1{F0: 0, F1: 32767, F2: v5}
        v3[fi6] = ((!true) && true)
        var v8 int64 = -423
        _ = fi6
        _ = v7
        _ = v8
    }
    switch "\n\t" {
    case "hello":
        if (42 > 18446744073709551615) {
            v5 = 50
            v9 := GenS1{F0: 65535, F1: 0, F2: v5}
            v10 := GenS1{F0: 42, F1: 32767, F2: (v5 & v5)}
            _ = v10
            _ = v9
        } else {
            var v11 int16 = 373
            v12 := GenS2{F0: "hello", F1: (v5 & v5), F2: ""}
            if (-428 > (v11 << ((18446744073709551615) & 63))) {
                var v13 float64 = 3.14159
                _ = v13
            }
            _ = v11
            _ = v12
        }
    case "\n\t":
        v1 = v1
    default:
        v5 = v5
    }
    v1 += (v1 & (^(v1)))
    fmt.Println(v1)
    delete(v3, 1)
    fmt.Println(v1, v5)
    var v14 uint8 = 62
    for fi15 := 0; fi15 < 7; fi15++ {
        v1 -= int32(v5)
        v1 -= v1
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
