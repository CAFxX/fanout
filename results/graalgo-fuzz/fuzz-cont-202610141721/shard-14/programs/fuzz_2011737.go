package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint64 }
func genFn0() bool {
    if false {
        var v6 uint8 = 1
        _ = v6
    }
    v7 := GenS1{F0: 2, F1: 42}
    var v8 int8 = -49
    _ = v7
    _ = v8
    return (65535 <= 255)
}
func genFn1(p0 uint16) int {
    fmt.Println(p0)
    switch "" {
    case "hello":
        var v9 uint8 = 42
        _ = v9
    case "":
        v10 := []int{87, 127, 2147483647, 2147483647}
        _ = v10
    default:
        var v11 int32 = 2147483647
        _ = v11
    }
    var v12 uint = 22
    v13 := GenS2{F0: 4294967295}
    _ = p0
    _ = v12
    _ = v13
    return -219
}

func main() {
    v1 := []string{"a", ("foo bar" + ("a" + ("foo bar" + "a"))), "\n\t", "hello"}
    v2 := []uint32{2, 42}
    var v3 uint64 = 0
    v4 := make(map[uint32]string)
    var v5 uint16 = 32767
    switch 127 {
    case 128:
        fmt.Println(v5)
    default:
        var v14 float64 = 0.1
        _ = v14
    }
    _ = len(v2)
    if ((v3 + v3) < 1) {
        v5 -= v5
        v3 -= v3
    }
    v3 += v3
    v5 = (v5 - v5)
    var v15 uint64 = (v3 - 0)
    if ((((0.1 * float32((1e10 * (-2.25 + float64(0.0))))) / -2.25) / ((3.14159 + 0.0) - ((0.1 / (float32((0.1 + 1.5)) + (((1.5 * 0.1) - (0.1 + 0.1)) * 0.1))) + 1.5))) != 3.14159) {
        fmt.Println(v3, v3, v5)
        for fi16 := 0; fi16 < 8; fi16++ {
            v5 = (v5 << ((185) & 63))
            _ = fi16
        }
        for fi17 := 0; fi17 < 10; fi17++ {
            v15 = v3
            _ = fi17
        }
    }
    var v18 int64 = -528
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
