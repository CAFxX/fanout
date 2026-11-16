package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 uint8; F2 int64 }
func genFn0(p0 string, p1 bool, p2 uint16) int32 {
    var v6 uint = 1
    v7 := GenS1{F0: (v6 & (^(v6)))}
    p2 = (p2 | p2)
    v8 := GenS1{F0: 5}
    var v9 int16 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return -2147483648
}
func genFn1(p0 int32, p1 int32) string {
    var v10 float64 = 1.5
    var v11 int8 = -128
    var v12 uint = 65535
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return (("\n\t" + "\n\t") + "hello")
}

func main() {
    var v1 float32 = 1e10
    v2 := make(map[int]bool)
    var v3 int16 = -623
    var v4 string = "hello"
    var v5 float32 = v1
    switch (v4 + v4) {
    case "a":
        fmt.Println(v5, v5)
    case "":
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 int32 = int32(fi13)
            _ = fi13
            _ = v14
        }
    default:
        v15 := GenS2{F0: 2147483647, F1: 127, F2: 9223372036854775807}
        _ = v15
    }
    v4 = "a"
    var v16 uint16 = 32767
    fmt.Println(v1)
    for fi17 := 0; fi17 < 6; fi17++ {
        fmt.Println(v3, v5)
        v18 := []int{}
        _ = fi17
        _ = v18
    }
    fmt.Println(v16)
    var v19 uint16 = v16
    _ = len(v2)
    var v20 uint64 = 0
    if (false && true) {
        v21 := v2[-1]
        v3 = 0
        _ = v21
    }
    if false {
        v1 = (float32(0.1) - v5)
        _ = len(v2)
        v1 = (1e10 + v1)
    }
    var v22 int16 = 0
    var v23 uint = 42
    switch (("" + v4) + "foo bar") {
    case "\n\t":
        fmt.Println(v1)
    case "hello":
        if true {
            v24 := []int{9223372036854775807}
            var v25 int8 = -128
            var v26 uint16 = (v16 | v19)
            _ = v24
            _ = v25
            _ = v26
        }
    }
    fmt.Println(v5)
    var v27 int32 = -23
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
