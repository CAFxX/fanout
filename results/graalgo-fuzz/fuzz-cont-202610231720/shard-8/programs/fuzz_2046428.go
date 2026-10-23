package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int16; F1 int32 }
func genFn0(p0 uint16, p1 int64, p2 string) uint64 {
    if (!false) {
        fmt.Println(p0, p1)
        p0 = (p0 / ((p0) | 1))
    }
    switch "\n\t" {
    case "\n\t":
        var v6 bool = (-795 <= -84)
        _ = v6
    default:
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 uint16 = p0
            _ = fi7
            _ = v8
        }
    }
    v9 := []int{-641, 42, 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 666
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{282, 42, 4294967295, 4294967295}
    var v3 string = ("hello" + (("foo bar" + "") + (("hello" + ("hello" + (("a" + "hello") + "hello"))) + ("" + "hello"))))
    var v4 uint64 = 1
    var v5 int = 4294967295
    if true {
        v10 := GenS1{F0: false, F1: 1}
        var v11 float32 = ((3.14159 * (0.1 * 0.0)) + 3.14159)
        _ = v10
        _ = v11
    }
    fmt.Println(v5, v4, v5)
    v3 = ""
    fmt.Println(v4)
    v3 = "\n\t"
    if false {
        var v12 uint = 465
        switch v3 {
        case "hello":
            v12 = 944
        case "foo bar":
            fmt.Println(v3)
        }
        v13 := GenS1{F0: false, F1: 7}
        _ = v12
        _ = v13
    }
    var v14 int16 = -32768
    v2 = append(v2, 4294967295)
    var v15 string = (v3 + (v3 + (((v3 + "a") + (("foo bar" + "\n\t") + v3)) + "foo bar")))
    for fi16 := 0; fi16 < 16; fi16++ {
        _ = len(v2)
        v5 += fi16
        for fi17 := 0; fi17 < 3; fi17++ {
            var v18 uint8 = 127
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    if true {
        _ = len(v2)
        v19 := GenS2{F0: (v14 >> ((4294967295) & 63)), F1: 42}
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
