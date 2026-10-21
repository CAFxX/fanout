package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int16; F2 int32 }
func genFn0(p0 uint32, p1 int) int16 {
    if (("\n\t" + "foo bar") < "hello") {
        switch 1 {
        case 32768:
            v4 := []int{p1, 0}
            _ = v4
        case 487:
            var v5 uint8 = 1
            _ = v5
        }
        if (42 > 32767) {
            var v6 uint16 = 65535
            _ = v6
        }
        var v7 string = ("hello" + ("a" + "foo bar"))
        _ = v7
    }
    switch (p1 << ((557) & 63)) {
    case 32768:
        var v8 float32 = ((0.1 - ((3.14159 - 0.1) - 0.0)) + (0.1 + 1.5))
        _ = v8
    case -173:
        p0 = p0
    }
    v9 := GenS1{F0: 2147483647, F1: 789}
    _ = p0
    _ = p1
    _ = v9
    return -32768
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[int]string)
    v3 := make(map[uint32]bool)
    fmt.Println(v1)
    v10 := []int{9223372036854775807, 2147483647, 633, 32767}
    fmt.Println(v1)
    var v11 uint16 = 255
    var v12 uint = (v1 | 202)
    var v13 int8 = -128
    if ((!true) && (("hello" >= "a") || false)) {
        for fi14 := 0; fi14 < 12; fi14++ {
            fmt.Println(v12, v1)
            v13 = v13
            _ = fi14
        }
        var v15 float64 = float64(3.14159)
        v16 := GenS2{F0: 65535, F1: 42, F2: -1}
        _ = v15
        _ = v16
    } else {
        v17 := GenS2{F0: 0, F1: 0, F2: -891}
        switch ("a" + "\n\t") {
        case "":
            v1 += (^(v12))
        }
        _ = v17
    }
    v3[4294967295] = false
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
