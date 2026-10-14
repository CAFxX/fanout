package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 bool, p1 int32, p2 int32) uint32 {
    p0 = p0
    if p0 {
        var v5 uint64 = 1
        v6 := GenS1{F0: "hello", F1: (p1 * 0)}
        v7 := []int{314, 762}
        _ = v5
        _ = v6
        _ = v7
    } else {
        v8 := []int{32768}
        if p0 {
            _ = len(v8)
            for fi9 := 0; fi9 < 12; fi9++ {
                switch -801 {
                case 9223372036854775807:
                    v10 := GenS2{F0: 128}
                    _ = v10
                default:
                    var v11 float64 = float64(1.5)
                    _ = v11
                }
                if len(v8) > 0 {
                    v8[0] = fi9
                }
                for fi12 := 0; fi12 < 2; fi12++ {
                    fmt.Println(fi9, fi9, fi9)
                    var v13 uint16 = 1
                    var v14 int32 = -2147483648
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                _ = fi9
            }
        } else {
            p2 = 582
            var v15 string = "\n\t"
            _ = v15
        }
        if ("a" <= "a") {
            switch 42 {
            case 913:
                v8 = append(v8, 32768)
            case 0:
                var v16 int = -827
                _ = v16
            case 453:
                for fi17 := 0; fi17 < 5; fi17++ {
                    switch "a" {
                    case "hello":
                        _ = len(v8)
                    }
                    if len(v8) > 0 {
                        v8[0] = (fi17 / ((fi17) | 1))
                    }
                    _ = fi17
                }
            }
            var v18 float64 = 0.0
            var v19 uint32 = 2147483647
            _ = v18
            _ = v19
        }
        _ = v8
    }
    v20 := []int{418, -753, 32767}
    var v21 float64 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return 98
}

func main() {
    v1 := []uint32{0, 2147483647, 2, 2147483648}
    var v2 uint16 = 255
    var v3 uint64 = 18446744073709551615
    var v4 int64 = 0
    var v22 int = 0
    fmt.Println(v22, v3, v4)
    v23 := GenS2{F0: 73}
    v24 := GenS2{F0: 42}
    switch "\n\t" {
    case "":
        var v25 int16 = 0
        _ = v25
    case "foo bar":
        v26 := GenS1{F0: "hello", F1: 42}
        _ = v26
    }
    v4 -= (v4 * (v4 - v4))
    var v27 int = v22
    v3 = 796
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
}
