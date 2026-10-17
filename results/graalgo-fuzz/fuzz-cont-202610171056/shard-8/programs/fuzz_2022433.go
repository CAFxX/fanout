package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := []uint32{42}
    v2 := make(map[uint32]bool)
    var v3 uint64 = 18446744073709551615
    v4 := []uint32{4294967295, 2147483647, 1}
    v5 := []string{"hello", ("hello" + ""), ("" + "a")}
    var v6 uint8 = 1
    v6 = (v6 % ((v6) | 1))
    var v7 int16 = 32767
    var v8 uint8 = (v6 % ((v6) | 1))
    v3 = (^(v3))
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := v2[65535]
        fmt.Println(v10, v10, v3)
        var v11 uint64 = uint64(v8)
        _ = fi9
        _ = v10
        _ = v11
    }
    v6 += (v6 + (v6 | (^(v8))))
    v12 := GenS1{F0: -6}
    switch "\n\t" {
    case "a":
        var v13 uint16 = 556
        _ = v13
    case "hello":
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 uint = 879
            var v16 float32 = 1e10
            if (-2.25 != 1e10) {
                if false {
                    if false {
                        v17 := GenS2{F0: 1, F1: (^(v6))}
                        _ = v17
                    } else {
                        var v18 float64 = (0.1 - 0.1)
                        for fi19 := 0; fi19 < 2; fi19++ {
                            var v20 int32 = -402
                            _ = fi19
                            _ = v20
                        }
                        fmt.Println(v6)
                        _ = v18
                    }
                    v21 := GenS2{F0: -2147483648, F1: v6}
                    _ = v21
                }
            }
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    switch ("hello" + "") {
    case "foo bar":
        if ((v8 + v8) > v6) {
            var v22 uint = 0
            var v23 uint16 = 42
            v24 := GenS3{F0: 42}
            _ = v22
            _ = v23
            _ = v24
        }
    case "hello":
        var v25 int8 = 0
        _ = v25
    }
    v26 := GenS1{F0: -128}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
