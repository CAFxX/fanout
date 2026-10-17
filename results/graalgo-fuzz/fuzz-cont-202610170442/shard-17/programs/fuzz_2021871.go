package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint32 }
func genFn0(p0 int, p1 int16) int16 {
    var v7 uint64 = 0
    v8 := GenS1{F0: 18446744073709551615}
    var v9 uint = uint(v7)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return p1
}

func main() {
    var v1 uint16 = 1
    var v2 uint = 255
    var v3 int8 = 0
    v4 := []uint32{65535, 58, 65535, 4294967295}
    var v5 string = "\n\t"
    v6 := make(map[string]string)
    var v10 uint64 = 42
    v4 = append(v4, 1)
    v1 = v1
    fmt.Println(v10, v5, v10)
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 float64 = 1e10
        _ = fi11
        _ = v12
    }
    fmt.Println(v1)
    for fi13 := 0; fi13 < 16; fi13++ {
        v10 -= v10
        v5 = "hello"
        _ = fi13
    }
    v14 := GenS2{F0: (!(!false)), F1: 2147483648}
    var v15 uint32 = 543
    if (v5 >= ("hello" + v5)) {
        for fi16 := 0; fi16 < 16; fi16++ {
            var v17 int32 = -3
            var v18 int = (fi16 >> (((v2 - 4294967295)) & 63))
            _ = fi16
            _ = v17
            _ = v18
        }
    } else {
        var v19 uint32 = (v15 / ((v15) | 1))
        v1 -= v1
        for fi20 := 0; fi20 < 8; fi20++ {
            _ = len(v4)
            _ = fi20
        }
        _ = v19
    }
    v5 = ("a" + v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
