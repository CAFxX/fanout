package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string }
func genFn0(p0 int8, p1 uint32) uint8 {
    if true {
        if (("\n\t" + (((((("a" + "a") + ("a" + (((("" + ("" + ("" + "a"))) + "hello") + (("hello" + ("foo bar" + ("" + "foo bar"))) + "foo bar")) + "foo bar"))) + "hello") + ("a" + "a")) + "\n\t") + "hello")) > (("\n\t" + "a") + "\n\t")) {
            var v6 int16 = -437
            var v7 int64 = int64(p0)
            var v8 uint16 = 332
            _ = v6
            _ = v7
            _ = v8
        }
    }
    v9 := GenS1{F0: -2147483648}
    var v10 uint32 = p1
    var v11 int = 4294967295
    var v12 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 128
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = true
    v3 := make(map[uint32]uint32)
    v4 := []uint32{}
    var v5 uint = 42
    _ = len(v4)
    v13 := GenS2{F0: 32767, F1: v2}
    var v14 uint8 = 2
    fmt.Println(v2, v5, v5)
    v14 += v14
    if v2 {
        v15 := GenS3{F0: "a"}
        var v16 float64 = ((0.0 / 3.14159) / 0.1)
        if (v5 > 255) {
            v17 := GenS4{F0: "foo bar"}
            for fi18 := 0; fi18 < 4; fi18++ {
                switch "foo bar" {
                case "foo bar":
                    if (!(!(!(v2 && v2)))) {
                        switch fi18 {
                        case 42:
                            fmt.Println(fi18, v2, v5)
                        case 32767:
                            switch "a" {
                            case "foo bar":
                                v16 += v16
                            case "":
                                var v19 uint8 = genFn0(0, 65535)
                                _ = v19
                            }
                        default:
                            v16 += v16
                        }
                    } else {
                        fmt.Println(v5, v2)
                        fmt.Println(v14, v2)
                        v4 = append(v4, 42)
                    }
                default:
                    var v20 uint64 = 4294967295
                    _ = v20
                }
                var v21 uint32 = 2147483647
                v21 = (v21 * v21)
                _ = fi18
                _ = v21
            }
            _ = v17
        } else {
            v2 = (!(v2 || v2))
            var v22 uint16 = 65535
            _ = v22
        }
        _ = v15
        _ = v16
    }
    var v23 int64 = 9223372036854775807
    v23 += v23
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
