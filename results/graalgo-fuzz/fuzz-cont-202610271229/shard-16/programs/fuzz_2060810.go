package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
func genFn0(p0 string) int16 {
    var v6 int64 = 1
    p0 = (("foo bar" + p0) + p0)
    switch 42 {
    case 2147483647:
        v7 := []int{0}
        _ = v7
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 float32 = (1e10 + 0.0)
        var v10 uint = 278
        fmt.Println(v6, v10)
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = v6
    return 135
}

func main() {
    var v1 int8 = 127
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    var v4 float32 = 1e10
    var v5 int64 = 42
    if false {
        v11 := v2[908]
        switch -1 {
        case 399:
            fmt.Println(v1)
        }
        _ = v11
    }
    switch int(v5) {
    case 9223372036854775807:
        delete(v2, 65535)
    default:
        if (2147483648 <= 985) {
            fmt.Println(v4, v1)
            var v12 string = ""
            fmt.Println(v1)
            _ = v12
        }
    }
    if (!false) {
        v4 = v4
        fmt.Println(v5, v5)
    } else {
        delete(v3, "")
    }
    var v13 uint64 = 42
    v5 = (-(v5))
    var v14 string = ("a" + "hello")
    switch -2147483648 {
    case 128:
        v15 := GenS1{F0: v13, F1: 9223372036854775807}
        _ = v15
    default:
        v2[144] = v14
    }
    var v16 uint32 = 2147483648
    v17 := GenS1{F0: 42, F1: 42}
    var v18 uint16 = 32767
    if true {
        for fi19 := 0; fi19 < 1; fi19++ {
            var v20 uint32 = v16
            v21 := []int{fi19}
            _ = fi19
            _ = v20
            _ = v21
        }
        for fi22 := 0; fi22 < 12; fi22++ {
            var v23 float64 = 3.14159
            _ = fi22
            _ = v23
        }
    } else {
        fmt.Println(v16)
        v18 += (v18 / ((v18) | 1))
    }
    var v24 uint64 = v13
    v14 = (v14 + v14)
    fmt.Println(v4, v24, v1)
    var v25 int64 = (-(v5))
    v26 := []int{127, -136}
    var v27 uint64 = v24
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
