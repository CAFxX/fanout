package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 int }
func genFn0(p0 uint16, p1 int, p2 uint64) uint64 {
    var v4 uint8 = 128
    if (-1 < -1) {
        if (!((("hello" + "") >= ("foo bar" + ("" + ""))) || ((1e10 / 1e10) <= 3.14159))) {
            if false {
                fmt.Println(p0, p0)
                var v5 int32 = -321
                _ = v5
            } else {
                if true {
                    fmt.Println(p2, p0)
                    fmt.Println(v4, v4)
                    var v6 int16 = 42
                    _ = v6
                }
            }
            var v7 int64 = -9223372036854775808
            _ = v7
        }
    }
    p0 -= 255
    if (42 >= p2) {
        p0 = (p0 + 644)
        v8 := []int{(-(p1)), p1, 243, (p1 & (p1 * p1))}
        _ = v8
    } else {
        switch "a" {
        case "hello":
            var v9 float32 = 1.5
            _ = v9
        case "foo bar":
            if (false || (0.0 > 1e10)) {
                switch "" {
                case "hello":
                    v10 := []int{2147483647, p1, p1}
                    _ = v10
                case "":
                    v11 := []int{}
                    _ = v11
                }
                v12 := GenS1{F0: 1, F1: uint(v4)}
                _ = v12
            }
        default:
            for fi13 := 0; fi13 < 16; fi13++ {
                var v14 int8 = 0
                _ = fi13
                _ = v14
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}

func main() {
    v1 := []uint32{2}
    var v2 int8 = 42
    var v3 string = "hello"
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 uint = 2
        _ = fi15
        _ = v16
    }
    fmt.Println(v2)
    var v17 float32 = 0.0
    v1 = append(v1, 42)
    var v18 uint = 18446744073709551615
    v3 = "\n\t"
    var v19 int32 = 42
    switch 128 {
    case 4294967295:
        var v20 int = 32767
        _ = v20
    case -246:
        for fi21 := 0; fi21 < 3; fi21++ {
            if false {
                v22 := GenS2{F0: (!true)}
                var v23 uint16 = 65535
                _ = v22
                _ = v23
            }
            var v24 int8 = v2
            if (!false) {
                var v25 uint16 = 1
                var v26 uint32 = 1
                v2 = (v2 * (v24 % (((v24 >> (((v18 + v18)) & 63))) | 1)))
                _ = v25
                _ = v26
            }
            _ = fi21
            _ = v24
        }
    case -2147483648:
        v1 = append(v1, 0)
    default:
        v1 = append(v1, 65535)
    }
    if (2147483647 == 194) {
        v18 -= 0
        v27 := GenS3{F0: false, F1: 1}
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
