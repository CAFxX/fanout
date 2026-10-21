package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 int8; F2 string }

func main() {
    var v1 float32 = (0.0 - (0.0 - 1e10))
    v2 := []uint32{897}
    v3 := make(map[uint32]string)
    var v4 float32 = (1e10 + v1)
    v5 := []uint32{}
    v6 := make(map[uint32]bool)
    v1 -= v1
    delete(v6, 42)
    fmt.Println(v4, v4)
    var v7 float64 = 1.5
    if (!true) {
        if (2147483647 >= 456) {
            fmt.Println(v4, v4, v4)
            v1 += 3.14159
            v4 = (float32(1.5) / 1e10)
        } else {
            v4 -= 1e10
            if (!(!true)) {
                v1 += (v1 + -2.25)
                fmt.Println(v1, v7, v7)
                v8 := GenS1{F0: 10}
                _ = v8
            } else {
                for fi9 := 0; fi9 < 8; fi9++ {
                    var v10 uint16 = 65535
                    var v11 bool = true
                    v12 := GenS2{F0: -1, F1: 42, F2: ("foo bar" + "foo bar")}
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
                v13 := GenS2{F0: -474, F1: -1, F2: "foo bar"}
                _ = v13
            }
            v1 -= v1
        }
        var v14 string = ("a" + "")
        _ = v14
    } else {
        var v15 int8 = -128
        _ = v15
    }
    var v16 float64 = 0.1
    fmt.Println(v7)
    var v17 uint16 = 768
    v7 = (0.1 + v7)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
