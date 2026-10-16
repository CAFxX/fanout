package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint8 }
func genFn0(p0 uint8, p1 uint64, p2 uint32) uint64 {
    if true {
        var v6 int32 = -984
        if ((uint64(p2) > p1) || false) {
            for fi7 := 0; fi7 < 17; fi7++ {
                var v8 uint64 = 42
                for fi9 := 0; fi9 < 12; fi9++ {
                    var v10 int16 = -689
                    v11 := []int{-230, (fi7 * (fi9 + fi9))}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                _ = fi7
                _ = v8
            }
            var v12 uint64 = 0
            _ = v12
        }
        _ = v6
    }
    var v13 uint64 = (p1 & (^(p1)))
    var v14 int16 = 1000
    if false {
        v15 := GenS1{F0: p0, F1: uint16(v13)}
        p0 = (p0 + p0)
        _ = v15
    }
    switch "\n\t" {
    case "foo bar":
        switch "a" {
        case "":
            var v16 uint64 = (v13 / (((^(v13))) | 1))
            _ = v16
        default:
            v17 := GenS2{F0: 1, F1: (^(p0))}
            _ = v17
        }
    case "":
        var v18 uint32 = (^(p2))
        _ = v18
    case "hello":
        v19 := GenS2{F0: (^(p1)), F1: (p0 % (((p0 - p0)) | 1))}
        _ = v19
    default:
        v20 := GenS1{F0: (p0 & (p0 / (((p0 % (((p0 << ((6) & 63))) | 1))) | 1))), F1: 65535}
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return (^(p1))
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    var v3 string = "\n\t"
    var v4 float64 = (((float64((((0.0 + 0.0) * 0.1) * -2.25)) - 1.5) - 0.0) * 1.5)
    var v5 float64 = (v4 * 0.1)
    var v21 int8 = 127
    fmt.Println(v4, v4)
    if true {
        if (2147483647 > -202) {
            for fi22 := 0; fi22 < 17; fi22++ {
                v3 = "\n\t"
                v23 := []int{}
                var v24 int32 = -2147483648
                _ = fi22
                _ = v23
                _ = v24
            }
        }
        var v25 uint64 = 18446744073709551615
        switch 128 {
        case 1:
            _ = len(v2)
        }
        _ = v25
    } else {
        var v26 string = v3
        v27 := []int{}
        _ = v26
        _ = v27
    }
    fmt.Println(v21)
    switch "\n\t" {
    case "":
        delete(v2, 0)
    }
    var v28 int32 = -2147483648
    v3 = ("foo bar" + v3)
    v29 := []int{2147483647, 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
