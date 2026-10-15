package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int16, p1 bool, p2 bool) int {
    var v5 int64 = 42
    fmt.Println(p1, p1, p2)
    if (9223372036854775807 <= 42) {
        if p2 {
            if p1 {
                fmt.Println(p1, p2)
                for fi6 := 0; fi6 < 18; fi6++ {
                    p0 = (p0 & p0)
                    _ = fi6
                }
            }
            v7 := GenS1{F0: v5}
            _ = v7
        }
        var v8 float64 = 0.1
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 32767
}
func genFn1(p0 uint32, p1 int32) uint8 {
    var v9 int16 = 42
    p1 = (p1 % ((p1) | 1))
    var v10 uint8 = 42
    p1 = (-(p1))
    p1 -= 1
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return v10
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 0
    var v3 int16 = -32768
    v4 := make(map[uint32]string)
    switch genFn0(866, (!true), (1e10 <= float32((((3.14159 * (float64(1.5) / 1.0)) + (1.5 + 1e10)) * 1e10)))) {
    case 42:
        v3 = v3
    default:
        switch ("a" + "a") {
        case "hello":
            if false {
                _ = len(v4)
                fmt.Println(v2, v3)
            }
        case "a":
            v2 += 2
        case "\n\t":
            switch "hello" {
            case "\n\t":
                v11 := v1[42]
                _ = v11
            default:
                var v12 float32 = 1.5
                _ = v12
            }
        default:
            for fi13 := 0; fi13 < 1; fi13++ {
                v3 -= 42
                _ = fi13
            }
        }
    }
    fmt.Println(v3, v2)
    var v14 uint32 = 591
    v15 := v4[v14]
    var v16 int64 = 9223372036854775807
    fmt.Println(v16, v16, v16)
    for fi17 := 0; fi17 < 1; fi17++ {
        v14 = v2
        _ = fi17
    }
    var v18 uint16 = 541
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
