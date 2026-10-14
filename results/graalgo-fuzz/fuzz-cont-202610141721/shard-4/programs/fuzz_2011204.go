package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 uint; F1 uint }

func main() {
    var v1 uint32 = 1
    var v2 int = 710
    v3 := make(map[uint32]int)
    fmt.Println(v1)
    var v4 int32 = 1
    if ((v4 % ((-2147483648) | 1)) > v4) {
        var v5 uint32 = 4294967295
        _ = v5
    }
    fmt.Println(v2)
    fmt.Println(v1)
    v6 := GenS1{F0: (-(v4)), F1: v4}
    v2 = v2
    if ((0.1 - 1e10) != ((0.1 + 3.14159) + 0.0)) {
        v1 += 2
    }
    fmt.Println(v4)
    for fi7 := 0; fi7 < 19; fi7++ {
        fmt.Println(v4, v1, v1)
        v8 := []int{v2, -949, (fi7 + (v2 * 1)), (v2 & 128)}
        var v9 float32 = (0.0 - -2.25)
        _ = fi7
        _ = v8
        _ = v9
    }
    if (false || false) {
        v10 := GenS2{F0: -32768, F1: 42}
        _ = v10
    } else {
        v2 += (v2 / ((v2) | 1))
        v2 -= (^(v2))
        if false {
            switch "hello" {
            case "a":
                v2 = (v2 | v2)
            case "":
                if (65535 >= 0) {
                    v11 := []int{v2, v2, 32768, v2}
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    v12 := GenS3{F0: -1, F1: 42}
                    _ = v11
                    _ = v12
                }
            case "\n\t":
                for fi13 := 0; fi13 < 8; fi13++ {
                    for fi14 := 0; fi14 < 15; fi14++ {
                        v15 := GenS4{F0: 317, F1: 255}
                        v1 = v1
                        _ = fi14
                        _ = v15
                    }
                    _ = fi13
                }
            }
            var v16 int32 = (v4 ^ (v4 | (v4 - (v4 | 518))))
            fmt.Println(v4, v2, v2)
            _ = v16
        }
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
