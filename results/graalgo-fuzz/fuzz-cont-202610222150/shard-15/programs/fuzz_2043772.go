package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
func genFn0(p0 uint, p1 uint, p2 int8) string {
    var v6 uint16 = 42
    var v7 uint = p0
    var v8 int64 = 1
    v9 := GenS1{F0: -32768, F1: uint(v6)}
    v10 := GenS1{F0: 259, F1: (p0 / (((p1 & 255)) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return ""
}

func main() {
    var v1 int32 = -931
    var v2 int64 = 9223372036854775807
    var v3 float32 = 1e10
    var v4 int64 = (v2 / (((v2 * v2)) | 1))
    var v5 int = -2147483648
    var v11 int16 = -32768
    v12 := GenS1{F0: v11, F1: 65535}
    if (!false) {
        if (2147483647 <= 65535) {
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 int64 = v4
                v15 := []int{fi13, v5, fi13}
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        for fi16 := 0; fi16 < 15; fi16++ {
            if false {
                v1 += (v1 * v1)
                var v17 uint16 = 838
                _ = v17
            } else {
                var v18 uint64 = 18446744073709551615
                switch "\n\t" {
                case "a":
                    switch "hello" {
                    case "hello":
                        v19 := []int{fi16, -1}
                        _ = v19
                    case "foo bar":
                        v20 := []int{32768}
                        _ = v20
                    }
                default:
                    fmt.Println(v3)
                }
                v21 := GenS1{F0: (^(v11)), F1: 530}
                _ = v18
                _ = v21
            }
            v11 = (-(v11))
            var v22 uint8 = 85
            _ = fi16
            _ = v22
        }
    } else {
        var v23 int8 = 93
        _ = v23
    }
    var v24 float32 = float32((0.0 - -2.25))
    v11 -= 0
    v25 := GenS1{F0: v11, F1: 554}
    var v26 int8 = 1
    v27 := []int{0, v5, -2147483648, v5}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
