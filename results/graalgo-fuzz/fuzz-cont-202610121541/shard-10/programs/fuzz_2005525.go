package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 int32; F1 uint8 }

func main() {
    var v1 uint16 = 1
    v2 := []int{-510, 128}
    v3 := []string{("" + "a"), ("foo bar" + ("a" + ((("" + "a") + "foo bar") + ""))), "a"}
    for fi4 := 0; fi4 < 16; fi4++ {
        var v5 int8 = 42
        fmt.Println(v1, v1)
        if (151 <= -1) {
            for fi6 := 0; fi6 < 6; fi6++ {
                var v7 uint8 = 128
                v7 += v7
                var v8 float32 = 1.5
                _ = fi6
                _ = v7
                _ = v8
            }
            var v9 bool = ((0 < 18446744073709551615) && false)
            v10 := GenS1{F0: -1}
            _ = v10
            _ = v9
        }
        _ = fi4
        _ = v5
    }
    v11 := GenS1{F0: -278}
    var v12 int = 128
    v1 = (v1 % ((691) | 1))
    v12 = 42
    v13 := GenS2{F0: 2147483648}
    _ = len(v3)
    v12 += v12
    v1 = v1
    for fi14 := 0; fi14 < 2; fi14++ {
        v15 := GenS3{F0: "a", F1: 427}
        _ = fi14
        _ = v15
    }
    switch v12 {
    case 128:
        _ = len(v2)
    case -1:
        v2 = append(v2, (-(v12)))
    case -2147483648:
        for fi16 := 0; fi16 < 2; fi16++ {
            var v17 bool = (-128 <= 127)
            v2 = append(v2, 0)
            var v18 uint16 = v1
            _ = fi16
            _ = v17
            _ = v18
        }
    default:
        v12 += v12
    }
    v12 += (v12 - (v12 << ((42) & 63)))
    v19 := GenS1{F0: -1}
    var v20 bool = false
    v21 := GenS4{F0: 42, F1: 128}
    v12 = (v12 + v12)
    var v22 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
