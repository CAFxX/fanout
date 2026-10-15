package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 bool; F1 int64; F2 uint64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{(("\n\t" + "foo bar") + "\n\t"), "\n\t"}
    var v3 uint = 139
    var v4 uint32 = 0
    var v5 int32 = 1
    var v6 int32 = (v5 * (-(v5)))
    for fi7 := 0; fi7 < 6; fi7++ {
        fmt.Println(v5, fi7, v6)
        fmt.Println(v3, v6)
        _ = fi7
    }
    v8 := GenS1{F0: 1}
    v2 = append(v2, "")
    for fi9 := 0; fi9 < 1; fi9++ {
        v4 = v4
        v10 := GenS2{F0: 1, F1: -2147483648}
        _ = fi9
        _ = v10
    }
    v6 = v6
    var v11 uint64 = uint64(v3)
    if true {
        var v12 uint = v3
        v4 = (v4 + v4)
        for fi13 := 0; fi13 < 20; fi13++ {
            if (-63 >= 42) {
                v1[(v4 & v4)] = (fi13 * (fi13 % ((2147483647) | 1)))
                if ((v3 + (^(v3))) != v12) {
                    delete(v1, v4)
                    v14 := GenS2{F0: 65535, F1: (v6 >> (((v12 << (((v12 + (v3 | v3))) & 63))) & 63))}
                    var v15 int8 = -1
                    _ = v14
                    _ = v15
                }
            }
            v16 := GenS3{F0: false, F1: -9223372036854775808, F2: v11}
            if (107 < -1) {
                var v17 uint64 = (^(v11))
                fmt.Println(v17)
                _ = v17
            }
            _ = fi13
            _ = v16
        }
        _ = v12
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
