package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint32; F1 uint32 }
func genFn0() bool {
    var v6 uint = 0
    var v7 uint64 = 281
    _ = v6
    _ = v7
    return true
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[uint32]string)
    var v3 int8 = -125
    var v4 int8 = v3
    var v5 uint32 = 914
    switch "\n\t" {
    case "foo bar":
        var v8 uint64 = 18446744073709551615
        _ = v8
    case "":
        var v9 int64 = -472
        _ = v9
    default:
        v10 := []int{32768}
        _ = v10
    }
    v11 := GenS1{F0: 1, F1: 2}
    var v12 int16 = -102
    delete(v2, v5)
    for fi13 := 0; fi13 < 12; fi13++ {
        delete(v2, v5)
        v5 = v5
        var v14 float32 = (((1.5 + (1e10 * -2.25)) - float32((0.1 * 1.5))) + 1e10)
        _ = fi13
        _ = v14
    }
    if true {
        v5 = v5
        v5 = (v5 ^ 65535)
    } else {
        v15 := GenS1{F0: (v3 - (v4 - -1)), F1: 135}
        v16 := GenS1{F0: v4, F1: 65535}
        v3 += (v3 - 127)
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        v1 = 0
        v1 = v1
        for fi18 := 0; fi18 < 13; fi18++ {
            v19 := GenS1{F0: v4, F1: 1}
            fmt.Println(v5, v5)
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        var v21 bool = false
        v3 = (v4 - v4)
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        v23 := GenS2{F0: v5, F1: v5}
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
