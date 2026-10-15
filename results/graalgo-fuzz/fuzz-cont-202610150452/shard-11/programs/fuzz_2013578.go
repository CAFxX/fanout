package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 bool; F2 int8 }
type GenS3 struct { F0 uint8; F1 int16; F2 int8 }

func main() {
    var v1 int64 = 232
    var v2 int = 32767
    v3 := make(map[string]string)
    var v4 int8 = 127
    v1 += (^(v1))
    fmt.Println(v1)
    for fi5 := 0; fi5 < 17; fi5++ {
        v6 := GenS1{F0: 0}
        for fi7 := 0; fi7 < 19; fi7++ {
            switch -9223372036854775808 {
            case -9223372036854775808:
                if (1 > 65535) {
                    v8 := []int{fi5, (v2 % (((^(v2))) | 1)), fi5, 2147483647}
                    fmt.Println(v4, v4)
                    _ = v8
                }
            case -613:
                var v9 int16 = 32767
                _ = v9
            case 10:
                var v10 uint8 = 2
                _ = v10
            default:
                if (v4 != 1) {
                    v11 := GenS1{F0: 42}
                    _ = v11
                } else {
                    delete(v3, "")
                }
            }
            _ = fi7
        }
        if (42 <= 539) {
            v4 -= -128
        }
        _ = fi5
        _ = v6
    }
    switch "" {
    case "":
        v12 := []int{(-(v2)), (v2 % ((v2) | 1)), v2, (v2 / ((197) | 1))}
        _ = v12
    case "a":
        v13 := GenS1{F0: 41}
        _ = v13
    case "\n\t":
        var v14 float32 = -2.25
        _ = v14
    default:
        v15 := GenS2{F0: "", F1: false, F2: (v4 >> ((65535) & 63))}
        _ = v15
    }
    fmt.Println(v2, v2)
    var v16 int = 42
    v17 := GenS1{F0: 2}
    v18 := GenS3{F0: 127, F1: -1, F2: int8(v16)}
    _ = len(v3)
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
