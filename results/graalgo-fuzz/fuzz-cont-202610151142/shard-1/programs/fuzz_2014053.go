package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 uint64 }
func genFn0(p0 int8, p1 bool, p2 uint32) int16 {
    v6 := GenS1{F0: 413}
    v7 := GenS2{F0: 908, F1: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -32768
}
func genFn1() int {
    switch ("a" + ("\n\t" + "foo bar")) {
    case "\n\t":
        var v8 float32 = 1.5
        _ = v8
    default:
        var v9 float32 = float32(1e10)
        _ = v9
    }
    var v10 float64 = (0.1 + float64((((float32((0.1 + 1.5)) + (1e10 * (3.14159 - 0.0))) - float32(0.0)) - float32((3.14159 * -2.25)))))
    _ = v10
    return 128
}

func main() {
    v1 := []string{"foo bar", "foo bar"}
    var v2 int8 = -32
    var v3 int64 = 42
    var v4 float32 = (0.1 - (float32(0.0) * 1e10))
    var v5 int32 = -216
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v5, v2)
    var v11 uint8 = 127
    _ = len(v1)
    for fi12 := 0; fi12 < 11; fi12++ {
        v3 -= (^(v3))
        v3 -= -1
        v5 = v5
        _ = fi12
    }
    v4 = float32(float64((1e10 * v4)))
    v13 := GenS1{F0: 42}
    v3 = (-(v3))
    switch ("hello" + "\n\t") {
    case "hello":
        v11 += (^(v11))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
