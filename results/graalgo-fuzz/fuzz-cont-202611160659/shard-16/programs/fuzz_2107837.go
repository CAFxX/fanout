package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint64 }
func genFn0() uint8 {
    switch 9223372036854775807 {
    case -351:
        for fi4 := 0; fi4 < 6; fi4++ {
            v5 := []int{}
            _ = fi4
            _ = v5
        }
    case 786:
        var v6 string = "foo bar"
        _ = v6
    case 101:
        var v7 int8 = -1
        _ = v7
    }
    var v8 bool = (921 < 762)
    fmt.Println(v8)
    v9 := []int{96, -2147483648, -975}
    var v10 uint64 = 18446744073709551615
    _ = v10
    _ = v8
    _ = v9
    return 128
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = 42
    v3 := []uint32{0, 431, 760}
    if false {
        v11 := GenS1{F0: 32767, F1: -1, F2: -879}
        _ = v11
    } else {
        var v12 int16 = 32767
        _ = v12
    }
    v2 = v2
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 int64 = -928
        _ = fi13
        _ = v14
    }
    v2 -= (v2 | 1)
    var v15 uint = 2
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 float32 = 1e10
        _ = fi16
        _ = v17
    }
    var v18 uint = 2
    if true {
        v19 := v1[1]
        var v20 uint16 = 462
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 7; fi21++ {
            v15 += v15
            fmt.Println(v18)
            _ = fi21
        }
    }
    v15 = v18
    var v22 uint16 = 0
    var v23 bool = true
    switch int(v15) {
    case 1:
        fmt.Println(v15, v15, v23)
    case 128:
        switch 701 {
        case -9223372036854775808:
            var v24 uint8 = 2
            _ = v24
        case -998:
            v2 += v2
        case 4294967295:
            v25 := GenS2{F0: genFn0(), F1: 469}
            _ = v25
        }
    case -1:
        v22 = v22
    default:
        if v23 {
            if len(v3) > 0 {
                v3[0] = 0
            }
            v2 += v2
        } else {
            _ = len(v3)
            v18 += v18
        }
    }
    v26 := GenS2{F0: 1, F1: 42}
    for fi27 := 0; fi27 < 3; fi27++ {
        for fi28 := 0; fi28 < 15; fi28++ {
            var v29 uint = (v18 / ((v18) | 1))
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    if len(v3) > 0 {
        v3[0] = 545
    }
    var v30 uint64 = 0
    if (42 >= v30) {
        var v31 uint32 = 42
        v32 := GenS1{F0: 1000, F1: -332, F2: -1000}
        v2 += v2
        _ = v31
        _ = v32
    }
    v30 = v30
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v30
}
