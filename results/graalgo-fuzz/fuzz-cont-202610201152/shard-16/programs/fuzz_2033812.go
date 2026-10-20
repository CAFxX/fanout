package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() uint8 {
    v6 := []int{}
    switch -790 {
    case 2147483647:
        var v7 uint32 = 0
        _ = v7
    case 386:
        switch "foo bar" {
        case "\n\t":
            var v8 uint16 = 65535
            _ = v8
        case "a":
            if len(v6) > 0 {
                v6[0] = 32768
            }
        }
    default:
        _ = len(v6)
    }
    var v9 int8 = 42
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = v6
    _ = v9
    return 122
}
func genFn1(p0 int64, p1 int8, p2 uint8) bool {
    v10 := []int{}
    v11 := GenS1{F0: genFn0()}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return false
}

func main() {
    var v1 uint8 = 42
    var v2 float32 = 1.5
    var v3 int64 = 0
    var v4 float32 = v2
    var v5 uint32 = 4294967295
    var v12 bool = false
    var v13 int16 = -32768
    v14 := []int{-105}
    var v15 uint32 = (v5 - (v5 ^ v5))
    var v16 int16 = (v13 >> ((249) & 63))
    fmt.Println(v15, v3, v15)
    for fi17 := 0; fi17 < 4; fi17++ {
        for fi18 := 0; fi18 < 7; fi18++ {
            var v19 float32 = (float32(0.0) - 3.14159)
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    fmt.Println(v2, v13, v2)
    v15 = v15
    _ = len(v14)
    var v20 int8 = -36
    v12 = (false && true)
    var v21 uint64 = 1
    for fi22 := 0; fi22 < 9; fi22++ {
        fmt.Println(v16)
        v23 := GenS1{F0: uint8(v5)}
        var v24 float64 = (-2.25 + (0.1 / (3.14159 * float64(-2.25))))
        _ = fi22
        _ = v23
        _ = v24
    }
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
