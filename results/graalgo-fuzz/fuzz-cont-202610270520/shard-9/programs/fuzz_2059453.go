package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 int16; F1 int }
func genFn0(p0 bool, p1 int, p2 string) int8 {
    switch 127 {
    case 933:
        var v5 int = p1
        _ = v5
    }
    v6 := GenS1{F0: p0, F1: 2}
    var v7 uint32 = 683
    switch (("hello" + p2) + p2) {
    case "\n\t":
        fmt.Println(p2, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint32, p1 uint8) uint16 {
    v8 := GenS2{F0: 42, F1: 18446744073709551615, F2: 18446744073709551615}
    v9 := []int{128}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 255
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{-497}
    v3 := []uint32{42, 2147483647, 438}
    v4 := []uint32{42, 2}
    var v10 string = "foo bar"
    switch (("" + ((v10 + ((v10 + v10) + v10)) + "hello")) + ((v10 + "\n\t") + (v10 + ("hello" + "")))) {
    case "":
        var v11 int = 128
        _ = v11
    case "foo bar":
        v10 = "\n\t"
    }
    if false {
        _ = len(v2)
        v12 := GenS3{F0: 9223372036854775807, F1: 42}
        v13 := GenS4{F0: 1, F1: -860}
        _ = v12
        _ = v13
    }
    var v14 float64 = ((float64(1e10) + 1.5) * ((0.1 * 0.1) - (((1e10 - (1.5 - (-2.25 - float64(-2.25)))) - 3.14159) - 1.5)))
    v10 = ("hello" + "\n\t")
    fmt.Println(v10, v14)
    v15 := GenS2{F0: 0, F1: 241, F2: 252}
    v14 -= (((v14 / (v14 - (float64(1.5) * v14))) + 0.0) / v14)
    v10 = ((("foo bar" + (v10 + v10)) + ("a" + (v10 + (("" + v10) + ("hello" + "a"))))) + v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
