package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint }
func genFn0(p0 int32, p1 string, p2 uint8) int {
    v6 := GenS1{F0: (-(p0)), F1: -128}
    v7 := GenS1{F0: (p0 / ((146) | 1)), F1: 26}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -854
}
func genFn1(p0 int16) bool {
    var v8 string = "hello"
    v9 := []int{127, -1}
    if len(v9) > 0 {
        v9[0] = 1
    }
    _ = p0
    _ = v8
    _ = v9
    return true
}

func main() {
    v1 := []uint32{834, 811, 691, 794}
    var v2 bool = (false || (1 == 333))
    v3 := make(map[uint32]int)
    var v4 float32 = 3.14159
    v5 := make(map[uint32]uint32)
    var v10 int64 = 1
    var v11 uint32 = 0
    var v12 int16 = -1
    for fi13 := 0; fi13 < 2; fi13++ {
        if true {
            v2 = v2
            for fi14 := 0; fi14 < 3; fi14++ {
                v12 = v12
                v12 -= v12
                _ = fi14
            }
        }
        var v15 uint8 = 31
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi13
        _ = v15
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        fmt.Println(v11, v10, v12)
        var v17 int64 = (v10 * (v10 << ((42) & 63)))
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: 0, F1: 1}
    var v19 uint = 1
    v20 := GenS1{F0: 910, F1: 127}
    if v2 {
        v21 := GenS1{F0: -2147483648, F1: 105}
        for fi22 := 0; fi22 < 1; fi22++ {
            v1 = append(v1, v11)
            var v23 float32 = v4
            _ = fi22
            _ = v23
        }
        _ = v21
    } else {
        delete(v5, v11)
    }
    if false {
        v4 -= v4
        v24 := GenS2{F0: v11, F1: v12}
        for fi25 := 0; fi25 < 4; fi25++ {
            var v26 uint16 = 0
            _ = fi25
            _ = v26
        }
        _ = v24
    }
    if v2 {
        v12 -= (v12 - 42)
    } else {
        var v27 int8 = 0
        _ = v27
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        v11 = v11
        _ = fi28
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi29 := 0; fi29 < 8; fi29++ {
        var v30 int32 = 402
        v31 := GenS3{F0: v19}
        for fi32 := 0; fi32 < 18; fi32++ {
            v11 -= v11
            v30 = (v30 >> ((2) & 63))
            _ = fi32
        }
        _ = fi29
        _ = v30
        _ = v31
    }
    switch int(v10) {
    case 1:
        _ = len(v1)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
