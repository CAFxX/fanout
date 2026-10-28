package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int8; F1 string; F2 int8 }
type GenS3 struct { F0 int8; F1 uint64; F2 int }

func main() {
    var v1 uint8 = 2
    var v2 int8 = 1
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 2
    v5 := []int{2147483647, 42, 0}
    if (255 <= 42) {
        v6 := GenS1{F0: 4294967295, F1: 830, F2: 632}
        if true {
            v7 := GenS2{F0: v2, F1: "", F2: v2}
            _ = v7
        } else {
            v5 = append(v5, 127)
        }
        for fi8 := 0; fi8 < 11; fi8++ {
            if true {
                if len(v5) > 0 {
                    v5[0] = fi8
                }
            } else {
                _ = len(v3)
                v4 = v4
            }
            v2 -= 100
            v2 -= (v2 + v2)
            _ = fi8
        }
        _ = v6
    }
    if (1000 > 32767) {
        if (!(-9223372036854775808 >= 42)) {
            switch "" {
            case "hello":
                var v9 string = ("a" + "hello")
                _ = v9
            default:
                for fi10 := 0; fi10 < 10; fi10++ {
                    v4 = (v4 - (v4 % ((2147483648) | 1)))
                    _ = fi10
                }
            }
            _ = len(v3)
            fmt.Println(v2, v2, v1)
        } else {
            var v11 int32 = int32(v1)
            _ = v11
        }
        var v12 uint = 0
        v2 += v2
        _ = v12
    }
    v13 := GenS1{F0: uint64(v1), F1: 42, F2: 2}
    v2 = (v2 + v2)
    v4 += (v4 ^ 2147483648)
    v14 := GenS3{F0: -28, F1: 2, F2: 127}
    v15 := GenS3{F0: v2, F1: 402, F2: -422}
    fmt.Println(v2, v1)
    var v16 int = 641
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
