package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = 1.5
    v3 := make(map[uint32]uint32)
    v4 := []int{-246, 128, 127, 38}
    var v5 int8 = 42
    v5 -= v5
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch 42 {
    case 42:
        v5 = 74
    default:
        v2 += v2
    }
    if false {
        var v6 int = -1
        _ = v6
    } else {
        var v7 int64 = 626
        var v8 bool = ((3.14159 >= 1e10) && true)
        if v8 {
            if len(v4) > 0 {
                v4[0] = 32767
            }
            var v9 int8 = 42
            if ((18446744073709551615 <= 0) || false) {
                fmt.Println(v2)
            }
            _ = v9
        }
        _ = v7
        _ = v8
    }
    fmt.Println(v5, v2)
    v3[0] = 2
    for fi10 := 0; fi10 < 14; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := GenS1{F0: 685}
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 15; fi13++ {
            _ = len(v3)
            var v14 uint16 = 85
            v5 += (v5 ^ v5)
            _ = fi13
            _ = v14
        }
        _ = fi10
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
