package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int32 }
type GenS2 struct { F0 bool; F1 int; F2 bool }
type GenS3 struct { F0 string }

func main() {
    var v1 uint = 904
    var v2 int = -929
    var v3 float32 = 1.5
    var v4 uint32 = 2147483648
    var v5 int8 = -1
    v6 := make(map[int]bool)
    fmt.Println(v2, v5, v4)
    fmt.Println(v4, v2)
    v7 := GenS1{F0: true, F1: uint(v4), F2: -2147483648}
    fmt.Println(v3, v5)
    var v8 string = "hello"
    v4 = (^(v4))
    v9 := []int{(v2 - (^(v2))), v2, (v2 >> (((v1 << ((v1) & 63))) & 63)), (-(v2))}
    v10 := GenS2{F0: (-911 >= 42), F1: v2, F2: (7 > (^(v5)))}
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := GenS3{F0: ("" + ("a" + ((v8 + v8) + v8)))}
        for fi13 := 0; fi13 < 20; fi13++ {
            v4 = v4
            var v14 int8 = (v5 % ((v5) | 1))
            _ = fi13
            _ = v14
        }
        v4 = (^(v4))
        _ = fi11
        _ = v12
    }
    var v15 float64 = float64(v3)
    v4 -= (v4 & v4)
    for fi16 := 0; fi16 < 14; fi16++ {
        v6[-237] = true
        for fi17 := 0; fi17 < 6; fi17++ {
            switch (v2 % ((fi16) | 1)) {
            case 9223372036854775807:
                v18 := GenS1{F0: true, F1: v1, F2: 42}
                _ = v18
            default:
                switch fi16 {
                case 128:
                    fmt.Println(fi17, v2)
                default:
                    var v19 uint16 = 418
                    _ = v19
                }
            }
            v2 = 238
            v8 = v8
            _ = fi17
        }
        fmt.Println(v5, v8)
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
