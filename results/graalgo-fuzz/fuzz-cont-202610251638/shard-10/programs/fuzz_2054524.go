package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int16, p1 int, p2 int) int64 {
    p1 = (-(p1))
    v5 := []int{(^(p1)), p1, 0, p2}
    var v6 uint8 = 128
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 int64 = -9223372036854775808
            var v10 string = "a"
            v11 := GenS1{F0: -2147483648}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -900
}
func genFn1() int8 {
    var v12 int8 = 42
    if false {
        var v13 uint = 114
        var v14 int32 = 1
        var v15 float64 = 1.5
        _ = v13
        _ = v14
        _ = v15
    }
    _ = v12
    return v12
}

func main() {
    var v1 int32 = -2147483648
    var v2 int32 = v1
    var v3 int64 = -1
    v4 := make(map[int]string)
    if (!true) {
        fmt.Println(v2, v2)
    }
    var v16 int = -176
    v1 = v2
    var v17 uint64 = 1
    v17 += (v17 >> ((220) & 63))
    var v18 float64 = 1e10
    fmt.Println(v2, v18)
    var v19 uint16 = 42
    var v20 float32 = 3.14159
    var v21 string = (("" + ("\n\t" + "a")) + (("a" + "") + ("" + "")))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
