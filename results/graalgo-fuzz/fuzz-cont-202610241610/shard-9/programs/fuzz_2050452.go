package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint32; F1 bool }
func genFn0(p0 string) int {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 int = fi6
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    _ = p0
    return 282
}
func genFn1(p0 int16, p1 int32) string {
    if true {
        var v8 int = 42
        _ = v8
    } else {
        p0 += (p0 + p0)
        fmt.Println(p0, p1)
    }
    var v9 uint32 = 1
    if true {
        v10 := []int{1}
        fmt.Println(p1, p1)
        _ = v10
    } else {
        p0 = (^(p0))
        v11 := []int{}
        _ = v11
    }
    var v12 int64 = 733
    _ = p0
    _ = p1
    _ = v12
    _ = v9
    return (("\n\t" + ((((("\n\t" + "") + "") + ("hello" + "")) + ("a" + ("\n\t" + ("" + "hello")))) + "foo bar")) + "hello")
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{83, 1, -687, -2147483648}
    v3 := []uint32{}
    var v4 bool = (0 > 810)
    v5 := make(map[string]bool)
    for fi13 := 0; fi13 < 14; fi13++ {
        for fi14 := 0; fi14 < 16; fi14++ {
            v4 = v4
            var v15 int16 = 1
            v16 := GenS1{F0: 1, F1: 65535, F2: -9223372036854775808}
            fmt.Println(v15, fi14)
            var v17 int8 = -128
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        switch 107 {
        case 32767:
            v18 := GenS1{F0: 42, F1: 1, F2: -1}
            _ = v18
        case 128:
            v4 = v4
        }
        _ = fi13
    }
    v4 = v4
    var v19 float32 = float32(-2.25)
    var v20 uint8 = 1
    var v21 uint = 4294967295
    v22 := GenS1{F0: 2147483647, F1: 2, F2: -9223372036854775808}
    if false {
        fmt.Println(v19, v20, v4)
        var v23 string = "hello"
        _ = v23
    }
    v24 := GenS2{F0: 1, F1: false}
    fmt.Println(v19, v19, v21)
    v21 = 0
    for fi25 := 0; fi25 < 9; fi25++ {
        fmt.Println(fi25)
        _ = fi25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
