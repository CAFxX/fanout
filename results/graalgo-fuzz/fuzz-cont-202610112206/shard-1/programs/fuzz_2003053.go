package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 int32 }
func genFn0(p0 string, p1 string) int16 {
    v7 := GenS1{F0: -9223372036854775808, F1: false}
    p1 = p1
    if false {
        var v8 float32 = ((0.1 * 0.0) * (1.5 * (-2.25 + ((((3.14159 / 1.0) / -2.25) - (0.1 - -2.25)) - (0.1 / 1e10)))))
        v9 := GenS2{F0: false}
        _ = v8
        _ = v9
    }
    p1 = p1
    v10 := GenS2{F0: (1 <= 2)}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return -677
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{}
    v3 := make(map[uint32]string)
    var v4 int = 128
    v5 := []int{(v4 | -9223372036854775808), (^(v4)), (v4 << ((65535) & 63))}
    var v6 int = 1
    var v11 float32 = 0.0
    v12 := GenS3{F0: 2147483648}
    var v13 uint16 = 0
    fmt.Println(v13)
    v6 -= v4
    v4 = (v4 - (v6 * (v4 * 1)))
    for fi14 := 0; fi14 < 3; fi14++ {
        v15 := GenS3{F0: 1}
        switch "hello" {
        case "hello":
            fmt.Println(v13)
        }
        switch "\n\t" {
        case "hello":
            if true {
                _ = len(v1)
                _ = len(v1)
            } else {
                v4 = (v4 / (((fi14 << ((42) & 63))) | 1))
            }
        case "foo bar":
            var v16 uint64 = uint64(v13)
            _ = v16
        case "a":
            if len(v5) > 0 {
                v5[0] = 9223372036854775807
            }
        default:
            if ((((1.5 * float64(v11)) - 3.14159) * 1e10) >= float64(1.5)) {
                if (!((0.1 < 3.14159) && (!(!true)))) {
                    fmt.Println(v6, v4, v6)
                } else {
                    v4 = (-(fi14))
                }
            } else {
                for fi17 := 0; fi17 < 14; fi17++ {
                    switch 341 {
                    case 127:
                        v18 := GenS4{F0: 1}
                        _ = v18
                    case 76:
                        var v19 int64 = 0
                        _ = v19
                    case 0:
                        v13 += v13
                    }
                    v20 := GenS5{F0: 55, F1: 1}
                    var v21 uint8 = 2
                    _ = fi17
                    _ = v20
                    _ = v21
                }
                v5 = append(v5, fi14)
            }
        }
        _ = fi14
        _ = v15
    }
    if (true || true) {
        if (0 <= 60) {
            var v22 uint64 = 0
            _ = v22
        }
        var v23 uint64 = 18446744073709551615
        v11 -= (v11 + (0.0 / float32(3.14159)))
        _ = v23
    }
    var v24 uint64 = 828
    var v25 uint16 = (v13 >> ((0) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
