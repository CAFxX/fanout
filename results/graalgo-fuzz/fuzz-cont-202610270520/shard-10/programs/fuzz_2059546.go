package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 string; F1 int32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{-182, 42, 42}
    var v3 uint32 = 901
    v3 += v3
    for fi4 := 0; fi4 < 9; fi4++ {
        var v5 int16 = 1
        var v6 int32 = -2147483648
        _ = fi4
        _ = v5
        _ = v6
    }
    fmt.Println(v3)
    v3 = 2147483647
    var v7 float64 = -2.25
    v8 := GenS1{F0: 160, F1: true}
    v7 += ((float64(1e10) + v7) - float64(-2.25))
    for fi9 := 0; fi9 < 15; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            fmt.Println(fi9, fi10)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if (2 < 1) {
                v11 := GenS1{F0: 42, F1: (!false)}
                var v12 bool = ((-869 == 2147483647) || (631 > -9223372036854775808))
                _ = v11
                _ = v12
            }
            _ = fi10
        }
        if true {
            var v13 int8 = 54
            for fi14 := 0; fi14 < 16; fi14++ {
                fmt.Println(v3, fi9, v7)
                var v15 string = "a"
                v16 := GenS2{F0: v15, F1: 888}
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = v13
        } else {
            fmt.Println(v7, fi9, v7)
            _ = len(v2)
        }
        _ = fi9
    }
    v3 -= v3
    var v17 int16 = -654
    delete(v1, 128)
    switch 738 {
    case 127:
        for fi18 := 0; fi18 < 1; fi18++ {
            if false {
                v3 += v3
                v19 := GenS2{F0: "a", F1: -835}
                delete(v1, 42)
                _ = v19
            } else {
                v20 := GenS1{F0: -414, F1: (false && (((-18 >= 100) && true) || false))}
                var v21 int8 = 113
                var v22 float64 = float64(((1.5 / 3.14159) + 1.5))
                _ = v20
                _ = v21
                _ = v22
            }
            var v23 int8 = 42
            fmt.Println(v3)
            _ = fi18
            _ = v23
        }
    default:
        v7 = v7
    }
    if len(v2) > 0 {
        v2[0] = 588
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
