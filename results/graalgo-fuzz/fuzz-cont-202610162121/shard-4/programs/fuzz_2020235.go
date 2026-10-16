package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int16 }
type GenS2 struct { F0 int32; F1 uint32 }
func genFn0(p0 uint32) string {
    if (!true) {
        p0 -= (p0 ^ p0)
        var v4 uint = 42
        _ = v4
    } else {
        p0 -= p0
        v5 := GenS1{F0: 2147483647, F1: 1, F2: 32767}
        v6 := []int{612}
        _ = v5
        _ = v6
    }
    var v7 uint8 = 42
    v8 := []int{127}
    fmt.Println(v7)
    fmt.Println(p0, v7)
    _ = p0
    _ = v7
    _ = v8
    return "a"
}
func genFn1() int16 {
    var v9 uint16 = 387
    v10 := []int{0, 2147483647, -2147483648}
    v9 = (^(v9))
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    if (4294967295 > 384) {
        v9 += v9
        fmt.Println(v9)
        v9 += v9
    } else {
        var v11 int8 = -1
        if true {
            var v12 string = "\n\t"
            var v13 float64 = (3.14159 + (float64(-2.25) - -2.25))
            _ = v12
            _ = v13
        }
        _ = v11
    }
    _ = v10
    _ = v9
    return 540
}

func main() {
    var v1 uint16 = 0
    var v2 int16 = 32767
    var v3 int64 = -1
    v14 := GenS2{F0: 824, F1: 4294967295}
    var v15 int64 = v3
    v1 += v1
    v2 += -32768
    var v16 float64 = (0.1 / 1.5)
    if (65535 == uint32(v1)) {
        fmt.Println(v3, v16)
    }
    v2 = v2
    var v17 uint16 = (v1 >> ((877) & 63))
    v2 = genFn1()
    v18 := GenS2{F0: 2147483647, F1: 2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
