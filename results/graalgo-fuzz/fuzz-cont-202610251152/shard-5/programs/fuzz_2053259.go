package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
func genFn0() string {
    if true {
        v7 := []int{0, -961}
        var v8 int16 = -1
        _ = v7
        _ = v8
    } else {
        v9 := GenS1{F0: 1, F1: 2}
        var v10 float64 = 1.5
        var v11 uint = 0
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 float64 = (3.14159 + 1.5)
    if true {
        for fi13 := 0; fi13 < 9; fi13++ {
            v12 -= 0.1
            fmt.Println(fi13)
            v14 := GenS1{F0: 42, F1: 255}
            _ = fi13
            _ = v14
        }
    } else {
        v15 := []int{176, 0}
        fmt.Println(v12)
        fmt.Println(v12)
        _ = v15
    }
    v16 := GenS1{F0: 128, F1: 1}
    var v17 uint64 = 1
    _ = v12
    _ = v16
    _ = v17
    return (("foo bar" + "foo bar") + "hello")
}

func main() {
    var v1 float64 = (0.0 / (float64(0.1) / float64(((1e10 + (3.14159 / (3.14159 + 3.14159))) + 3.14159))))
    var v2 int16 = -32768
    v3 := []uint32{1, 2, 2}
    var v4 int32 = 1
    v5 := make(map[string]bool)
    v6 := make(map[uint32]uint32)
    for fi18 := 0; fi18 < 12; fi18++ {
        var v19 uint16 = 32767
        if len(v3) > 0 {
            v3[0] = 2147483648
        }
        _ = fi18
        _ = v19
    }
    fmt.Println(v2, v4)
    if (!true) {
        var v20 uint32 = 2147483647
        var v21 float64 = (v1 - (((v1 * (v1 * (v1 - v1))) / (3.14159 - (0.1 * v1))) / (v1 + (1e10 + 1.5))))
        _ = v20
        _ = v21
    }
    v1 -= ((3.14159 / (((3.14159 - 1.5) / (v1 / 0.1)) / ((v1 - float64(1e10)) + v1))) - ((0.1 / v1) * -2.25))
    fmt.Println(v4)
    if (0 >= 42) {
        v1 += 0.1
        v5[("" + ("\n\t" + "a"))] = (("" + "foo bar") >= "hello")
    } else {
        delete(v5, (("a" + "foo bar") + "\n\t"))
        v4 = (v4 ^ (v4 - v4))
        if (!false) {
            delete(v6, 488)
            fmt.Println(v4, v1, v2)
            var v22 string = ((("" + "hello") + "a") + "a")
            _ = v22
        }
    }
    switch 1 {
    case 4294967295:
        var v23 int16 = v2
        _ = v23
    default:
        v24 := GenS1{F0: 2, F1: 0}
        _ = v24
    }
    v2 -= 42
    var v25 uint16 = 2
    v25 += v25
    var v26 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
