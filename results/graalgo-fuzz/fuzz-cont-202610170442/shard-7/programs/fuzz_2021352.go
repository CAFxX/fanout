package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint32; F1 uint64 }

func main() {
    var v1 float64 = 1.5
    v2 := make(map[string]uint32)
    var v3 uint16 = 791
    switch 2147483647 {
    case -1:
        v3 = 690
    case 0:
        var v4 int8 = 1
        _ = v4
    case -2147483648:
        if (195 > 225) {
            v1 += (v1 + (((float64(1e10) * 0.0) + v1) + float64((0.1 + (float32(v1) * 1.5)))))
            v5 := []int{9223372036854775807, -9223372036854775808, 261}
            _ = v5
        } else {
            if (42 == 970) {
                var v6 uint8 = 255
                var v7 int = 542
                if true {
                    var v8 uint8 = (v6 & v6)
                    var v9 uint8 = v8
                    _ = v8
                    _ = v9
                }
                _ = v6
                _ = v7
            }
            if false {
                v10 := v2["\n\t"]
                _ = v10
            } else {
                switch -2147483648 {
                case 32768:
                    if true {
                        v11 := GenS1{F0: 0, F1: false}
                        _ = v11
                    } else {
                        var v12 uint32 = 42
                        v3 = 26
                        _ = v12
                    }
                }
                var v13 uint32 = 494
                _ = v13
            }
            v14 := []int{-173, -1}
            _ = v14
        }
    default:
        if (1 <= (^(v3))) {
            fmt.Println(v1, v3)
        } else {
            var v15 uint = 18446744073709551615
            _ = v15
        }
    }
    var v16 uint32 = 4294967295
    _ = len(v2)
    v16 -= (^(v16))
    if (2 != 784) {
        var v17 bool = ((v3 / (((v3 | 255)) | 1)) > v3)
        fmt.Println(v16)
        delete(v2, "")
        _ = v17
    }
    v18 := []int{127, 708, 761}
    v16 += (^(v16))
    fmt.Println(v16, v16)
    var v19 uint16 = (v3 + v3)
    switch "" {
    case "":
        var v20 int8 = 127
        _ = v20
    case "hello":
        if len(v18) > 0 {
            v18[0] = -2147483648
        }
    case "foo bar":
        v21 := v2[(("" + "hello") + ("a" + ""))]
        _ = v21
    default:
        var v22 uint16 = 32767
        _ = v22
    }
    switch 9223372036854775807 {
    case 127:
        var v23 string = ("a" + "a")
        _ = v23
    default:
        v24 := GenS1{F0: 100, F1: true}
        _ = v24
    }
    v1 -= v1
    v25 := GenS2{F0: v16, F1: uint64(v3)}
    fmt.Println(v1, v19)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
}
