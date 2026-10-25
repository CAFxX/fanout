package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int16, p1 uint8) int32 {
    v5 := []int{-2147483648, -267, 0, 128}
    _ = len(v5)
    _ = p0
    _ = p1
    _ = v5
    return -1
}
func genFn1() int32 {
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 string = "hello"
        var v8 uint16 = 2
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int8 = -1
    v10 := []int{-202}
    var v11 int = -50
    v10 = append(v10, v11)
    _ = v10
    _ = v11
    _ = v9
    return 42
}

func main() {
    var v1 float32 = 1.5
    v2 := make(map[int]int)
    var v3 int64 = 910
    var v4 string = "foo bar"
    var v12 int32 = genFn0(-555, 255)
    var v13 int = 9223372036854775807
    fmt.Println(v4)
    if (!(true && (120 <= 4294967295))) {
        v3 = v3
    }
    var v14 uint16 = 626
    var v15 int16 = -32768
    for fi16 := 0; fi16 < 19; fi16++ {
        fmt.Println(v12, v15, v14)
        v17 := GenS1{F0: (v3 << ((857) & 63))}
        _ = fi16
        _ = v17
    }
    fmt.Println(v13, v12, v12)
    var v18 int16 = (v15 ^ -790)
    v15 -= (v15 / (((-(v15))) | 1))
    fmt.Println(v18, v15)
    v19 := []int{-867}
    fmt.Println(v4)
    if ((v14 >> ((65535) & 63)) != (v14 ^ 0)) {
        v20 := GenS2{F0: 0}
        _ = v20
    }
    v15 = v18
    if (v1 > ((3.14159 - 1e10) + float32(0.1))) {
        var v21 uint32 = 1
        _ = v21
    }
    fmt.Println(v4, v12)
    v1 = float32(float64(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
