package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 string }
func genFn0() uint16 {
    v6 := GenS1{F0: 173}
    var v7 float64 = (-2.25 + ((3.14159 - 1.5) / (float64(((1e10 / float32(1e10)) + 0.1)) * 1.5)))
    _ = v6
    _ = v7
    return 1
}
func genFn1() int8 {
    var v8 uint64 = 42
    var v9 bool = (-1 == -1)
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 uint64 = 4294967295
    var v2 float64 = 0.1
    var v3 string = ("" + (("hello" + ("hello" + "a")) + "foo bar"))
    v4 := []string{("\n\t" + v3), "", "\n\t"}
    v5 := []uint32{183}
    var v10 float64 = 0.0
    var v11 uint = 558
    fmt.Println(v3)
    v12 := GenS1{F0: 584}
    var v13 uint16 = 255
    v14 := GenS2{F0: 1, F1: v3}
    v10 = ((v10 * v2) - float64(-2.25))
    if (0.1 > ((0.0 - v2) * (float64((1.5 * (float32((v2 / v2)) * (1e10 * 1.5)))) + v2))) {
        v1 -= v1
        var v15 int16 = 0
        _ = v15
    }
    switch ("" + "foo bar") {
    case "":
        var v16 bool = (true || true)
        _ = v16
    default:
        var v17 int64 = -9223372036854775808
        _ = v17
    }
    var v18 int64 = 0
    var v19 uint32 = 2147483648
    var v20 int64 = (^(v18))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
