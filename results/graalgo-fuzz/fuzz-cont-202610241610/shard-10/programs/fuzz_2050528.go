package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 string; F1 uint8; F2 uint8 }
type GenS5 struct { F0 int }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = -144
    var v3 float64 = 1.5
    var v4 uint64 = 974
    var v5 int32 = -2147483648
    v6 := GenS1{F0: 147, F1: 42}
    fmt.Println(v5)
    v7 := GenS2{F0: 0, F1: 42}
    var v8 uint = 42
    var v9 int8 = -90
    var v10 int64 = 674
    v11 := GenS1{F0: uint8(v8), F1: 30}
    delete(v1, 2147483647)
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := GenS3{F0: v10, F1: (v8 + (v8 / (((v8 * 2)) | 1)))}
        var v14 int = (fi12 % ((32767) | 1))
        switch (v14 << (((^(v8))) & 63)) {
        case 0:
            v15 := GenS4{F0: ("" + "foo bar"), F1: 255, F2: 128}
            _ = v15
        case 2147483647:
            v16 := []int{(fi12 << (((^(v8))) & 63)), (fi12 >> (((v8 << ((v8) & 63))) & 63)), v14}
            _ = v16
        default:
            for fi17 := 0; fi17 < 3; fi17++ {
                v2 = (v2 % (((^(v2))) | 1))
                v2 = (v2 << ((v8) & 63))
                _ = fi17
            }
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    v5 = v5
    fmt.Println(v3, v9, v4)
    if (2 <= 2147483648) {
        v3 -= (v3 / (v3 + (float64(-2.25) - 0.1)))
        var v18 string = "a"
        _ = v18
    }
    v19 := GenS5{F0: -844}
    v20 := []int{1, 4294967295}
    var v21 int16 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
