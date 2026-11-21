package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint; F1 uint }
func genFn0(p0 uint16, p1 int32, p2 int16) int {
    v7 := []int{225, 2147483647, -9223372036854775808, int(p2)}
    v8 := GenS1{F0: "foo bar", F1: (128 > 234)}
    for fi9 := 0; fi9 < 16; fi9++ {
        p2 = -1
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 32767
}
func genFn1(p0 int32, p1 int64, p2 uint8) int8 {
    v10 := GenS1{F0: "", F1: (368 == 2147483647)}
    fmt.Println(p2, p0)
    fmt.Println(p2, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return -128
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 2
    v3 := []int{-9223372036854775808, 9223372036854775807, -9223372036854775808}
    var v4 uint32 = 480
    var v5 float32 = (-2.25 / ((((3.14159 + (-2.25 / 0.1)) - 0.1) + 1.5) * -2.25))
    var v6 float32 = float32((3.14159 / float64(v5)))
    fmt.Println(v2)
    v6 -= (v5 * v5)
    v6 -= float32(1e10)
    var v11 string = (("a" + ((((("foo bar" + "hello") + "hello") + "") + "") + ("a" + ("a" + ("hello" + "\n\t"))))) + "\n\t")
    v11 = v11
    var v12 float32 = float32(1.5)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v13 int32 = 42
    v13 += v13
    var v14 float64 = 1.5
    if (42 < 85) {
        v15 := GenS1{F0: v11, F1: ((!(false || (v13 != 0))) && true)}
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        switch (fi16 - fi16) {
        case 128:
            fmt.Println(v2, v13)
        }
        var v17 uint = 263
        v18 := GenS2{F0: (v17 ^ v17), F1: v17}
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 string = (v11 + "hello")
    var v20 bool = (false || true)
    for fi21 := 0; fi21 < 15; fi21++ {
        switch (fi21 & (fi21 ^ (fi21 << ((0) & 63)))) {
        case 9223372036854775807:
            v14 += -2.25
        }
        _ = fi21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
