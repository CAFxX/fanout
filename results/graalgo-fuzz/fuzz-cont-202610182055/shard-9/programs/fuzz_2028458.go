package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int64 }
func genFn0(p0 int, p1 bool) uint64 {
    var v7 int64 = -9223372036854775808
    if (int16(v7) >= 936) {
        var v8 uint64 = 4294967295
        v7 -= v7
        _ = v8
    }
    fmt.Println(p0, p1)
    var v9 float32 = (((3.14159 - 0.1) + 0.0) + 1.5)
    v7 = v7
    _ = p0
    _ = p1
    _ = v7
    _ = v9
    return 2
}
func genFn1() bool {
    var v10 int = 1
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 int = v10
        var v13 uint64 = 2
        var v14 uint64 = genFn0(fi11, false)
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 int8 = -128
    if false {
        v16 := []int{(v10 * v10), v10, v10}
        for fi17 := 0; fi17 < 5; fi17++ {
            if false {
                v15 += v15
                var v18 string = "a"
                _ = v18
            }
            _ = fi17
        }
        _ = v16
    }
    _ = v10
    _ = v15
    return (128 < 1)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 4294967295
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 65535
    var v5 int16 = 32767
    v6 := []string{("hello" + "\n\t")}
    switch "foo bar" {
    case "hello":
        v4 = v4
    }
    var v19 uint64 = 228
    var v20 uint32 = 1
    fmt.Println(v2, v19)
    v21 := GenS1{F0: ((^(v5)) == 0), F1: ("a" + "")}
    for fi22 := 0; fi22 < 14; fi22++ {
        switch -1 {
        case 797:
            v19 = (^(v19))
        case 9223372036854775807:
            v23 := GenS1{F0: true, F1: "a"}
            _ = v23
        default:
            fmt.Println(v5)
        }
        var v24 uint64 = (v19 % (((v19 & (v19 ^ 732))) | 1))
        var v25 uint16 = 0
        _ = fi22
        _ = v24
        _ = v25
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v5 -= v5
    for fi26 := 0; fi26 < 3; fi26++ {
        fmt.Println(v19)
        v2 = 1
        _ = fi26
    }
    if false {
        v6 = append(v6, "a")
    }
    v27 := GenS2{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
