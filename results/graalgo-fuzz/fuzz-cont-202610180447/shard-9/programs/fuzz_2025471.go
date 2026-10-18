package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
func genFn0(p0 int16, p1 int16, p2 int16) uint8 {
    v7 := GenS1{F0: 127, F1: (!true)}
    v8 := []int{169, 139, 42, 186}
    var v9 uint8 = 127
    if true {
        p2 = p2
        var v10 uint32 = 65535
        var v11 uint = 2
        _ = v10
        _ = v11
    } else {
        p2 = -1
        var v12 int = 0
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return (v9 | v9)
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = -32768
    var v3 uint8 = 128
    v4 := make(map[uint32]int)
    var v5 string = "a"
    var v6 int32 = 42
    var v13 int16 = 0
    v13 = (v2 | v13)
    var v14 uint16 = 360
    for fi15 := 0; fi15 < 9; fi15++ {
        v6 -= 493
        var v16 float32 = float32((1e10 * ((float64((0.0 * (1e10 + -2.25))) * 3.14159) / float64(1e10))))
        _ = fi15
        _ = v16
    }
    v14 -= v14
    var v17 int = 2147483647
    var v18 int8 = 42
    var v19 int32 = -537
    var v20 int8 = v18
    v20 = (v20 & v18)
    fmt.Println(v13, v18, v17)
    if (true || false) {
        var v21 uint64 = 597
        _ = v21
    }
    if false {
        if false {
            v22 := []int{}
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
            _ = v22
        } else {
            fmt.Println(v19, v3, v5)
        }
        for fi23 := 0; fi23 < 9; fi23++ {
            var v24 int8 = v20
            _ = fi23
            _ = v24
        }
        var v25 uint32 = 42
        _ = v25
    } else {
        if true {
            v26 := []int{(v17 >> ((0) & 63))}
            _ = v26
        }
        var v27 int = (v17 ^ (^(v17)))
        delete(v4, 4294967295)
        _ = v27
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        var v29 uint = 65535
        v1["hello"] = ("\n\t" + v5)
        _ = fi28
        _ = v29
    }
    v30 := []int{}
    switch v17 {
    case 1:
        v3 -= (v3 >> ((uint(v3)) & 63))
    case 32768:
        v17 -= (v17 + (v17 + -75))
    case 2147483647:
        for fi31 := 0; fi31 < 7; fi31++ {
            var v32 int8 = 100
            _ = len(v1)
            _ = fi31
            _ = v32
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
