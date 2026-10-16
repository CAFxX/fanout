package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 bool; F1 int32 }
func genFn0(p0 uint32) int8 {
    fmt.Println(p0)
    var v4 int16 = 42
    _ = p0
    _ = v4
    return -1
}

func main() {
    var v1 int16 = 326
    var v2 uint64 = 2
    var v3 float32 = ((-2.25 + 0.0) - (-2.25 * (float32(1e10) + 3.14159)))
    v1 += (v1 | 482)
    fmt.Println(v3, v1)
    v5 := GenS1{F0: -147, F1: 2147483648}
    v1 = v1
    switch 127 {
    case -427:
        v6 := GenS1{F0: 139, F1: 1}
        _ = v6
    default:
        v7 := GenS2{F0: false, F1: 2147483647}
        _ = v7
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 int16 = -32768
        fmt.Println(v3)
        fmt.Println(fi8)
        _ = fi8
        _ = v9
    }
    var v10 float64 = float64(1e10)
    var v11 string = ""
    v11 = v11
    v2 = v2
    fmt.Println(v11, v2, v2)
    for fi12 := 0; fi12 < 8; fi12++ {
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 uint = 18446744073709551615
            v3 = v3
            var v15 int32 = -301
            if (float32(v10) > float32(1.5)) {
                v16 := []int{fi13, (fi12 + fi12), (fi13 | fi13)}
                _ = v16
            }
            switch "hello" {
            case "\n\t":
                if (743 != 110) {
                    v3 += v3
                    switch 9223372036854775807 {
                    case 42:
                        v17 := []int{(fi13 ^ (fi12 ^ (fi12 - fi12))), 32768, (fi12 / ((fi13) | 1))}
                        _ = v17
                    case 1:
                        var v18 float64 = v10
                        _ = v18
                    case 128:
                        v15 = (v15 - (v15 - v15))
                    default:
                        var v19 string = (((v11 + (("a" + "foo bar") + "")) + v11) + "hello")
                        v1 -= 32767
                        _ = v19
                    }
                    v20 := []int{(fi12 * fi12), (-(fi12))}
                    _ = v20
                } else {
                    v2 -= 2
                    v21 := []int{(fi13 ^ (fi12 + (^(fi12)))), int(v1), (fi13 ^ fi13), fi13}
                    _ = v21
                }
            case "":
                var v22 int8 = genFn0(2147483647)
                _ = v22
            case "hello":
                v11 = "\n\t"
            default:
                var v23 uint32 = 1
                _ = v23
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        var v24 string = v11
        v25 := []int{2147483647, (fi12 >> ((42) & 63))}
        _ = fi12
        _ = v24
        _ = v25
    }
    v26 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v3
    _ = v5
}
