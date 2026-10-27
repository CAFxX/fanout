package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint32) uint {
    v5 := []int{}
    p0 -= 4294967295
    _ = p0
    _ = v5
    return 255
}
func genFn1() uint {
    var v6 int8 = -128
    if (32767 != -2147483648) {
        v7 := GenS1{F0: 255, F1: 42}
        _ = v7
    } else {
        v8 := GenS1{F0: 224, F1: -462}
        _ = v8
    }
    v6 = -103
    v9 := GenS2{F0: genFn0(0)}
    _ = v6
    _ = v9
    return 65535
}

func main() {
    var v1 string = "hello"
    var v2 int16 = -716
    var v3 bool = (v1 >= "foo bar")
    var v4 int = 42
    var v10 string = v1
    var v11 int64 = 0
    var v12 int16 = v2
    var v13 float32 = float32(0.0)
    v14 := GenS1{F0: 107, F1: -1}
    if true {
        for fi15 := 0; fi15 < 3; fi15++ {
            v16 := GenS2{F0: 65535}
            v2 = v12
            v17 := []int{}
            _ = fi15
            _ = v16
            _ = v17
        }
        v18 := []int{}
        _ = v18
    } else {
        var v19 uint = 4294967295
        var v20 int16 = v2
        _ = v19
        _ = v20
    }
    fmt.Println(v1)
    v21 := GenS2{F0: 0}
    v22 := []int{32767, (v4 % (((-(v4))) | 1)), v4, 32767}
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 int32 = -425
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        v12 -= v2
        for fi26 := 0; fi26 < 4; fi26++ {
            fmt.Println(v13, v13)
            v27 := GenS1{F0: 128, F1: 0}
            if len(v22) > 0 {
                v22[0] = (-(fi25))
            }
            v28 := GenS2{F0: 65535}
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
            _ = fi26
            _ = v27
            _ = v28
        }
        var v29 string = "a"
        _ = fi25
        _ = v29
    }
    v12 += 1000
    var v30 int16 = 430
    v2 -= v30
    v31 := GenS2{F0: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
