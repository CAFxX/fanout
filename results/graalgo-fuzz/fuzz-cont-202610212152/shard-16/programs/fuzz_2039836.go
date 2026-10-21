package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
func genFn0(p0 int32, p1 int) uint32 {
    p0 -= (p0 + (-(p0)))
    v6 := []int{-9223372036854775808, (p1 >> ((42) & 63)), p1}
    var v7 uint64 = 2
    var v8 uint16 = 42
    if (((0.0 - 1.5) * 1e10) >= ((1.5 - (((1e10 * 3.14159) - 3.14159) / 1e10)) / 1.0)) {
        v9 := GenS1{F0: 1, F1: ("" + ("hello" + ("hello" + "hello")))}
        _ = v9
    } else {
        var v10 string = "foo bar"
        var v11 int8 = 42
        var v12 int8 = v11
        _ = v10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 2147483648
}
func genFn1(p0 uint16) uint8 {
    v13 := []int{-9223372036854775808, 42, 4294967295}
    fmt.Println(p0)
    _ = p0
    _ = v13
    return 42
}

func main() {
    v1 := []string{"hello"}
    var v2 string = ((("hello" + (("\n\t" + ("\n\t" + "foo bar")) + "")) + "foo bar") + (("foo bar" + "a") + ("" + "hello")))
    var v3 float64 = 3.14159
    var v4 int16 = -1
    var v5 int32 = 42
    var v14 uint16 = 368
    var v15 int32 = v5
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 bool = false
        var v18 uint8 = 165
        _ = fi16
        _ = v17
        _ = v18
    }
    v1 = append(v1, v2)
    var v19 float32 = 0.1
    _ = len(v1)
    var v20 float64 = (float64(v19) / v3)
    v2 = v2
    for fi21 := 0; fi21 < 19; fi21++ {
        fmt.Println(v14)
        _ = fi21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
