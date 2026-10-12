package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
func genFn0() uint {
    var v5 int8 = -82
    if false {
        var v6 int8 = (v5 * (v5 / (((v5 / ((42) | 1))) | 1)))
        v5 += v5
        _ = v6
    }
    _ = v5
    return 1
}
func genFn1(p0 uint8) int8 {
    switch 9223372036854775807 {
    case 128:
        p0 += (p0 + p0)
    case -714:
        fmt.Println(p0)
    case 0:
        var v7 uint16 = 255
        _ = v7
    }
    fmt.Println(p0)
    if true {
        var v8 uint64 = 42
        _ = v8
    }
    v9 := []int{614}
    fmt.Println(p0)
    _ = p0
    _ = v9
    return 127
}

func main() {
    var v1 string = ("foo bar" + "")
    var v2 bool = false
    var v3 uint = 255
    var v4 uint = v3
    v2 = (1e10 <= float32(((float64(0.1) - 1e10) - 1e10)))
    fmt.Println(v1)
    v10 := []int{4294967295, -9223372036854775808, 2147483647, -2147483648}
    if v2 {
        var v11 string = v1
        v11 = (("foo bar" + v11) + (("" + ("hello" + "")) + (v11 + "")))
        v2 = false
        _ = v11
    }
    var v12 uint32 = 65535
    fmt.Println(v3)
    fmt.Println(v12, v2)
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 15; fi14++ {
            v3 = (v3 % ((v3) | 1))
            _ = fi14
        }
        _ = fi13
    }
    var v15 uint64 = uint64(v12)
    var v16 int8 = 127
    var v17 uint32 = v12
    var v18 uint8 = 127
    if ((v12 ^ 682) >= uint32(v4)) {
        var v19 int64 = -9223372036854775808
        var v20 uint64 = 81
        _ = v19
        _ = v20
    } else {
        v10 = append(v10, -402)
    }
    v21 := GenS1{F0: v12, F1: true}
    v18 = v18
    v22 := GenS1{F0: v12, F1: ((v2 && v2) && (v1 > ((("" + ("foo bar" + ("\n\t" + "a"))) + (v1 + v1)) + "foo bar")))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
