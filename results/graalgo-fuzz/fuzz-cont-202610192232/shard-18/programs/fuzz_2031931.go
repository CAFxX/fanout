package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int64; F2 int64 }
type GenS3 struct { F0 string; F1 uint32 }
func genFn0(p0 bool) int32 {
    var v7 uint16 = 2
    switch "a" {
    case "foo bar":
        v8 := []int{128, -170, 32768, 32767}
        _ = v8
    }
    _ = p0
    _ = v7
    return 42
}

func main() {
    var v1 uint16 = 286
    var v2 uint16 = 170
    var v3 float64 = 0.1
    v4 := make(map[uint32]int)
    v5 := []int{-82, 128}
    var v6 float64 = 1e10
    v1 -= (^(v2))
    v9 := GenS1{F0: 128, F1: (^(v2))}
    var v10 int8 = 100
    v11 := GenS2{F0: 255, F1: -9223372036854775808, F2: -9223372036854775808}
    if true {
        for fi12 := 0; fi12 < 13; fi12++ {
            for fi13 := 0; fi13 < 12; fi13++ {
                switch (fi13 | (fi13 - fi12)) {
                case -327:
                    var v14 uint32 = 2
                    _ = v14
                default:
                    v1 -= (v2 - v2)
                }
                delete(v4, 763)
                _ = fi13
            }
            _ = fi12
        }
        v15 := v4[2147483648]
        _ = v15
    } else {
        switch "foo bar" {
        case "foo bar":
            _ = len(v5)
        case "":
            v3 = (3.14159 / (3.14159 - (v6 - (v6 / (v3 * (v3 / v3))))))
        }
        for fi16 := 0; fi16 < 18; fi16++ {
            _ = len(v4)
            _ = fi16
        }
    }
    fmt.Println(v6, v3, v10)
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 int = (fi17 & (fi17 >> ((1) & 63)))
        v5 = append(v5, fi17)
        for fi19 := 0; fi19 < 1; fi19++ {
            v20 := GenS1{F0: 1, F1: v1}
            _ = fi19
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    v21 := GenS3{F0: "foo bar", F1: 42}
    var v22 int16 = -1
    v22 = (v22 << ((1) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
