package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint }
func genFn0(p0 string, p1 uint) uint32 {
    var v6 uint16 = 255
    var v7 int = 128
    fmt.Println(v7, p0, p0)
    fmt.Println(v7)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 1
}
func genFn1() int {
    v8 := []int{-9223372036854775808, 1, -467}
    if true {
        var v9 string = "a"
        var v10 int = 32768
        _ = v10
        _ = v9
    } else {
        var v11 int64 = -1
        if (255 > 38) {
            var v12 int = 42
            var v13 int = (v12 - (v12 + (v12 % ((v12) | 1))))
            v14 := GenS1{F0: genFn0("a", 0), F1: 0, F2: 1}
            _ = v12
            _ = v13
            _ = v14
        } else {
            var v15 string = ""
            _ = v15
        }
        v16 := GenS1{F0: 65535, F1: 18446744073709551615, F2: v11}
        _ = v11
        _ = v16
    }
    var v17 uint64 = 0
    fmt.Println(v17)
    _ = v17
    _ = v8
    return 2147483647
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 1
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]bool)
    var v5 uint8 = 128
    fmt.Println(v5, v2)
    v18 := v1[1]
    if v18 {
        v19 := v4[genFn0("\n\t", 4294967295)]
        if true {
            v18 = (-2.25 < 0.0)
        }
        var v20 int16 = 32767
        _ = v19
        _ = v20
    } else {
        var v21 string = "a"
        for fi22 := 0; fi22 < 6; fi22++ {
            var v23 float64 = (3.14159 / 3.14159)
            _ = fi22
            _ = v23
        }
        _ = v21
    }
    var v24 int32 = 2147483647
    fmt.Println(v24, v2, v24)
    v5 += v5
    switch "\n\t" {
    case "\n\t":
        var v25 int64 = 42
        _ = v25
    case "foo bar":
        v26 := v4[1]
        _ = v26
    }
    if true {
        var v27 uint32 = 0
        v18 = v18
        for fi28 := 0; fi28 < 16; fi28++ {
            fmt.Println(v24, v27, v2)
            var v29 int = genFn1()
            fmt.Println(v18)
            _ = fi28
            _ = v29
        }
        _ = v27
    } else {
        v30 := GenS2{F0: 255}
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
