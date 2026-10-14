package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 string; F1 int64; F2 uint8 }
type GenS4 struct { F0 string }
func genFn0(p0 uint8) uint32 {
    v6 := []int{}
    v7 := GenS1{F0: 238}
    for fi8 := 0; fi8 < 15; fi8++ {
        if len(v6) > 0 {
            v6[0] = (fi8 & fi8)
        }
        _ = fi8
    }
    p0 = p0
    _ = p0
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 string = ("" + "hello")
    var v2 uint64 = 18446744073709551615
    var v3 int64 = 293
    var v4 bool = (128 > -9223372036854775808)
    var v5 int16 = 304
    for fi9 := 0; fi9 < 20; fi9++ {
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 uint32 = 65535
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    v1 = (v1 + ((v1 + ("" + v1)) + v1))
    switch 0 {
    case 42:
        for fi12 := 0; fi12 < 6; fi12++ {
            for fi13 := 0; fi13 < 8; fi13++ {
                v4 = v4
                switch ("foo bar" + "\n\t") {
                case "\n\t":
                    if ((v3 / ((v3) | 1)) <= (v3 ^ (^(v3)))) {
                        v5 = v5
                    }
                case "foo bar":
                    fmt.Println(fi13, v2)
                case "hello":
                    v14 := []int{(fi12 | fi12), fi13, (fi12 ^ 9223372036854775807)}
                    _ = v14
                }
                _ = fi13
            }
            var v15 float32 = 3.14159
            v1 = v1
            _ = fi12
            _ = v15
        }
    case -2147483648:
        for fi16 := 0; fi16 < 20; fi16++ {
            fmt.Println(v5, v3, v4)
            _ = fi16
        }
    default:
        var v17 float32 = 1e10
        _ = v17
    }
    switch "hello" {
    case "foo bar":
        var v18 float32 = (-2.25 - 3.14159)
        _ = v18
    }
    v19 := GenS2{F0: 128, F1: 1}
    fmt.Println(v3, v2, v2)
    v20 := GenS3{F0: "foo bar", F1: v3, F2: 13}
    v4 = ((!true) && (255 == 1))
    v21 := GenS2{F0: 127, F1: 65535}
    v22 := GenS4{F0: v1}
    v23 := []int{148, 32768, 127}
    _ = len(v23)
    v3 -= 42
    var v24 int32 = 0
    v3 += 1
    fmt.Println(v4, v5)
    v3 = (v3 >> ((0) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
