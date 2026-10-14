package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int16 }
type GenS2 struct { F0 int64; F1 int32; F2 uint32 }
type GenS3 struct { F0 int8; F1 string; F2 uint }

func main() {
    v1 := make(map[int]int)
    var v2 int64 = -1
    v3 := []uint32{1, 1}
    v4 := make(map[int]uint32)
    var v5 int32 = 1
    v5 = (v5 & v5)
    for fi6 := 0; fi6 < 18; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            v5 = v5
            switch ("a" + (("" + "hello") + "hello")) {
            case "hello":
                var v8 uint16 = 2
                _ = v8
            case "":
                switch fi7 {
                case -2147483648:
                    v9 := GenS1{F0: "hello", F1: 42, F2: 1}
                    _ = v9
                case 127:
                    if len(v3) > 0 {
                        v3[0] = 964
                    }
                }
            case "foo bar":
                v10 := GenS2{F0: 9223372036854775807, F1: 429, F2: 2}
                _ = v10
            default:
                if true {
                    v5 = -2147483648
                    v11 := GenS3{F0: -128, F1: "", F2: 0}
                    fmt.Println(fi6)
                    _ = v11
                } else {
                    v5 += (v5 & (v5 % ((1) | 1)))
                }
            }
            v5 = -899
            _ = fi7
        }
        _ = fi6
    }
    v12 := GenS2{F0: (v2 << ((108) & 63)), F1: v5, F2: 42}
    v5 -= (v5 ^ (v5 & (v5 - 42)))
    var v13 uint = 0
    fmt.Println(v2)
    fmt.Println(v13, v5)
    var v14 int = 0
    v2 -= v2
    v5 = (v5 | 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
