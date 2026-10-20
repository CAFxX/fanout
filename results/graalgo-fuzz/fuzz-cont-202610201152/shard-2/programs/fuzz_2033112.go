package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int32; F2 int64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 0
    var v3 int32 = -722
    v4 := make(map[int]uint32)
    v5 := GenS1{F0: 0, F1: 2147483648}
    var v6 uint64 = 42
    if true {
        v7 := []int{2147483647, 1}
        v6 -= v6
        for fi8 := 0; fi8 < 19; fi8++ {
            for fi9 := 0; fi9 < 3; fi9++ {
                v3 += 97
                switch (fi9 | (fi9 & (fi8 << ((1) & 63)))) {
                case 128:
                    v6 = (v6 + (v6 / ((v6) | 1)))
                case 127:
                    v10 := GenS2{F0: 912, F1: v3, F2: 42}
                    _ = v10
                case 32768:
                    var v11 uint = 4294967295
                    _ = v11
                default:
                    fmt.Println(fi9, fi8)
                }
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = fi9
            }
            _ = fi8
        }
        _ = v7
    }
    v6 = (^(v6))
    _ = len(v1)
    v2 = (v2 + (v2 << ((166) & 63)))
    if ("a" >= "") {
        v12 := []int{933, -2147483648, 14}
        v3 = v3
        for fi13 := 0; fi13 < 9; fi13++ {
            v6 = (v6 | v6)
            _ = fi13
        }
        _ = v12
    }
    v14 := v4[-289]
    var v15 uint32 = v14
    v6 = (v6 & 770)
    v15 += 1
    if ((^(v2)) < (v2 << ((42) & 63))) {
        if (!(false && true)) {
            var v16 bool = ("foo bar" < "foo bar")
            fmt.Println(v14, v2, v16)
            _ = v16
        } else {
            switch ("hello" + ("\n\t" + "")) {
            case "foo bar":
                v17 := []int{1, 862}
                _ = v17
            }
            _ = len(v1)
        }
        v18 := []int{0, 347}
        v6 = 18446744073709551615
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
