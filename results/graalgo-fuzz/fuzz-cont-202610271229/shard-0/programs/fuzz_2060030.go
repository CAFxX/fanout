package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 string }
func genFn0(p0 uint32) uint16 {
    var v5 uint64 = 4294967295
    var v6 float32 = 0.0
    v5 += v5
    _ = p0
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 int16, p1 string) string {
    v7 := []int{330}
    if true {
        v8 := GenS1{F0: p1}
        _ = v8
    } else {
        v9 := GenS2{F0: 1, F1: (!(!true))}
        if true {
            v10 := GenS1{F0: p1}
            var v11 uint16 = 35
            _ = v10
            _ = v11
        }
        _ = v9
    }
    switch p1 {
    case "\n\t":
        p0 -= 1000
    case "":
        var v12 int = -704
        _ = v12
    case "hello":
        p1 = "foo bar"
    }
    _ = len(v7)
    _ = p0
    _ = p1
    _ = v7
    return p1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = false
    var v3 int8 = 42
    var v4 int16 = 0
    v13 := GenS1{F0: ("a" + "a")}
    for fi14 := 0; fi14 < 2; fi14++ {
        fmt.Println(v4, v4)
        _ = fi14
    }
    v15 := GenS3{F0: 0}
    v4 -= -32768
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 bool = (0.0 > float32(float64(3.14159)))
        if v17 {
            fmt.Println(v2)
            v3 = (v3 | 100)
        }
        _ = fi16
        _ = v17
    }
    v2 = (42 < v4)
    v18 := GenS1{F0: "hello"}
    v3 = -128
    v19 := GenS4{F0: (-2147483648 >= -315), F1: "\n\t"}
    if false {
        v20 := []int{-9223372036854775808, 2147483647, -9223372036854775808, 32767}
        var v21 float32 = float32(((((float64((-2.25 / -2.25)) - (-2.25 * 1.5)) * (0.0 / 0.1)) * 1.5) * 1e10))
        v20 = append(v20, 315)
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        switch fi22 {
        case 32768:
            var v23 int32 = 1
            _ = v23
        default:
            v24 := []int{}
            _ = v24
        }
        _ = fi22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
