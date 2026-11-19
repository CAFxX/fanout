package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int32 }

func main() {
    var v1 int = 0
    v2 := []uint32{0, 253, 4294967295}
    v3 := make(map[int]bool)
    var v4 bool = (1.5 == 1.5)
    var v5 float32 = float32((0.0 - float64(float32(1.5))))
    var v6 string = (("a" + ("foo bar" + ("" + "a"))) + "a")
    v4 = (-2.25 > 0.1)
    var v7 uint64 = 1
    var v8 bool = (127 <= 87)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v9 float32 = v5
    switch (v1 % ((4294967295) | 1)) {
    case -1:
        var v10 int8 = 15
        _ = v10
    case 0:
        _ = len(v2)
    case 518:
        delete(v3, v1)
    default:
        v8 = v4
    }
    _ = len(v3)
    var v11 uint8 = 210
    if (42 == -747) {
        if v8 {
            var v12 int16 = -32768
            _ = v12
        } else {
            if (!(true || v4)) {
                switch v6 {
                case "foo bar":
                    var v13 int16 = -1
                    _ = v13
                case "a":
                    delete(v3, v1)
                case "hello":
                    switch "hello" {
                    case "a":
                        fmt.Println(v6, v6)
                    case "":
                        var v14 int8 = 1
                        _ = v14
                    default:
                        if (!v4) {
                            var v15 int8 = int8(v11)
                            _ = v15
                        } else {
                            v5 = ((v5 * float32((-2.25 * 1e10))) - v5)
                            v16 := GenS1{F0: -1, F1: 2147483647, F2: 1}
                            _ = v16
                        }
                    }
                }
            } else {
                v7 = v7
            }
        }
    }
    v5 = v9
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    v17 := GenS1{F0: 42, F1: uint32(v11), F2: 0}
    delete(v3, (v1 - (^(v1))))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
