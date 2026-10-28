package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 bool }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint; F1 uint32; F2 int }

func main() {
    var v1 uint = 361
    var v2 string = "foo bar"
    var v3 int = 128
    v4 := make(map[int]bool)
    var v5 int8 = -109
    v2 = ((v2 + "") + v2)
    fmt.Println(v3)
    fmt.Println(v3, v2)
    v6 := GenS1{F0: false, F1: v5, F2: true}
    v3 += v3
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := []int{(v3 >> ((v1) & 63)), (v3 | 0), fi7}
        var v9 int = fi7
        switch -9223372036854775808 {
        case 42:
            for fi10 := 0; fi10 < 8; fi10++ {
                _ = len(v8)
                fmt.Println(fi7, v5)
                var v11 int8 = (v5 | -128)
                if len(v8) > 0 {
                    v8[0] = (v3 << (((v1 >> (((v1 / ((0) | 1))) & 63))) & 63))
                }
                _ = fi10
                _ = v11
            }
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    v12 := []int{v3, v3}
    v13 := GenS2{F0: 164, F1: 2}
    for fi14 := 0; fi14 < 9; fi14++ {
        v15 := GenS3{F0: 65535, F1: 539, F2: (v3 % ((v3) | 1))}
        _ = fi14
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
