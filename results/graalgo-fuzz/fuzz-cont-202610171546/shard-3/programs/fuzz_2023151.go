package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int32; F2 uint }
func genFn0(p0 uint64, p1 uint16) uint16 {
    fmt.Println(p0, p0)
    if (-937 < 1) {
        var v5 uint16 = (^(p1))
        if (18446744073709551615 == 255) {
            var v6 uint32 = 0
            switch "" {
            case "":
                if true {
                    p0 = 2
                    var v7 int8 = 42
                    _ = v7
                }
            case "foo bar":
                v8 := []int{31, 0, 4294967295, 793}
                _ = v8
            default:
                var v9 uint16 = 2
                _ = v9
            }
            var v10 uint64 = p0
            _ = v10
            _ = v6
        } else {
            for fi11 := 0; fi11 < 12; fi11++ {
                v12 := GenS1{F0: 2147483647}
                var v13 bool = true
                _ = fi11
                _ = v12
                _ = v13
            }
            var v14 int16 = -32768
            _ = v14
        }
        var v15 int32 = 2147483647
        _ = v15
        _ = v5
    }
    var v16 int32 = 2147483647
    if false {
        v17 := GenS2{F0: 127, F1: 42, F2: 0}
        _ = v17
    } else {
        p0 -= p0
    }
    _ = p0
    _ = p1
    _ = v16
    return (^(p1))
}
func genFn1(p0 uint64, p1 int32) uint16 {
    p1 -= (p1 - p1)
    v18 := []int{-1, 498, 4294967295}
    _ = p0
    _ = p1
    _ = v18
    return 32767
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{-142, 42, -2147483648}
    var v3 int64 = 1
    v4 := []uint32{2147483648, 2147483647, 2147483647}
    v3 = 42
    v19 := GenS1{F0: -659}
    var v20 uint64 = 0
    delete(v1, 2147483648)
    var v21 bool = (((((0.0 * 1.5) / 1.0) + (0.1 * 1.5)) - float64((0.0 + (3.14159 + 3.14159)))) > (-2.25 + 1.5))
    var v22 int32 = 0
    switch ("\n\t" + "a") {
    case "hello":
        var v23 uint = 0
        _ = v23
    case "foo bar":
        var v24 int16 = 0
        _ = v24
    case "\n\t":
        var v25 uint32 = 2
        _ = v25
    }
    v3 = 559
    for fi26 := 0; fi26 < 4; fi26++ {
        var v27 float64 = 3.14159
        _ = fi26
        _ = v27
    }
    if true {
        var v28 int8 = -1
        _ = v28
    }
    var v29 uint = 18446744073709551615
    _ = len(v1)
    if len(v4) > 0 {
        v4[0] = 891
    }
    var v30 string = "hello"
    fmt.Println(v29, v3, v30)
    v22 -= v22
    var v31 float64 = 0.0
    if (-2147483648 > 1) {
        delete(v1, 256)
        var v32 int32 = (^(v22))
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
