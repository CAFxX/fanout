package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 int }
func genFn0(p0 int16) uint64 {
    v5 := []int{735, 1, 0}
    if ((true || true) && false) {
        fmt.Println(p0)
        var v6 int = -9223372036854775808
        if true {
            fmt.Println(p0)
        }
        _ = v6
    } else {
        if (0 == 2147483647) {
            var v7 int64 = 9223372036854775807
            if false {
                p0 = (p0 - p0)
            } else {
                p0 += (p0 * (p0 | p0))
            }
            _ = v7
        }
    }
    _ = p0
    _ = v5
    return 1
}

func main() {
    var v1 uint16 = 2
    var v2 int16 = 1
    v3 := make(map[uint32]uint32)
    var v4 int = 2147483647
    fmt.Println(v1, v1)
    v8 := []int{v4, v4, 2147483647, 505}
    var v9 float64 = 0.0
    var v10 uint64 = 2
    v1 = 977
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        if true {
            v12 := GenS1{F0: ("hello" + ""), F1: -1, F2: (("" + "") + "hello")}
            v4 += v4
            for fi13 := 0; fi13 < 16; fi13++ {
                if (false && false) {
                    v14 := GenS1{F0: "foo bar", F1: 42, F2: ""}
                    fmt.Println(fi13, v1, v10)
                    var v15 uint64 = v10
                    _ = v14
                    _ = v15
                } else {
                    var v16 uint64 = (v10 * v10)
                    fmt.Println(v1)
                    var v17 int8 = 1
                    delete(v3, 2147483647)
                    if (v16 > genFn0((v2 * (v2 * v2)))) {
                        v18 := GenS2{F0: fi11}
                        v3[0] = 2147483648
                        var v19 uint8 = 98
                        _ = v18
                        _ = v19
                    }
                    _ = v16
                    _ = v17
                }
                var v20 uint64 = v10
                _ = fi13
                _ = v20
            }
            _ = v12
        }
        for fi21 := 0; fi21 < 3; fi21++ {
            v22 := GenS1{F0: "hello", F1: 0, F2: "a"}
            v10 += (v10 - (^(v10)))
            var v23 float32 = (((0.0 + 0.0) * 1e10) + 1.5)
            _ = fi21
            _ = v22
            _ = v23
        }
        if ((-(v2)) < v2) {
            var v24 int8 = 42
            _ = len(v8)
            v4 -= 32768
            _ = v24
        } else {
            v25 := GenS1{F0: "", F1: -9223372036854775808, F2: ("" + ("" + "foo bar"))}
            _ = v25
        }
        _ = fi11
    }
    switch "\n\t" {
    case "foo bar":
        var v26 int64 = 42
        _ = v26
    case "hello":
        var v27 uint8 = 255
        _ = v27
    default:
        _ = len(v3)
    }
    if true {
        var v28 uint16 = (^(v1))
        v29 := GenS3{F0: 2147483647, F1: -622}
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
