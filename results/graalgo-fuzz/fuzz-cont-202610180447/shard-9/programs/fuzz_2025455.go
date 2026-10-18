package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 string }
type GenS2 struct { F0 int; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint; F1 string; F2 uint64 }
type GenS4 struct { F0 string; F1 uint16; F2 bool }
type GenS5 struct { F0 int8; F1 int32; F2 uint }
func genFn0(p0 int32, p1 int8, p2 int32) uint16 {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 9; fi5++ {
        p1 = p1
        v6 := GenS1{F0: 2147483647, F1: ((fi5 - fi5) >= (fi5 / ((fi5) | 1))), F2: "foo bar"}
        var v7 uint = 2
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := []int{42, 9223372036854775807, -2147483648, 127}
    _ = len(v8)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 32767
}
func genFn1(p0 int32, p1 string) uint64 {
    v9 := GenS2{F0: -596, F1: 42, F2: 0}
    switch 2147483647 {
    case -1:
        var v10 uint32 = 845
        _ = v10
    case -157:
        switch p1 {
        case "\n\t":
            fmt.Println(p0, p0)
        case "":
            p1 = (p1 + (p1 + (p1 + (p1 + "a"))))
        default:
            p0 = (p0 / ((p0) | 1))
        }
    case 317:
        var v11 int16 = 1000
        _ = v11
    default:
        var v12 uint32 = 840
        _ = v12
    }
    p0 += (^(p0))
    for fi13 := 0; fi13 < 11; fi13++ {
        v14 := GenS3{F0: 2, F1: "\n\t", F2: 42}
        _ = fi13
        _ = v14
    }
    var v15 uint16 = genFn0((p0 >> ((255) & 63)), 32, 0)
    _ = p0
    _ = p1
    _ = v15
    _ = v9
    return 42
}

func main() {
    var v1 int64 = 247
    v2 := make(map[int]uint32)
    var v3 uint64 = 18446744073709551615
    var v4 string = "\n\t"
    v16 := GenS2{F0: 717, F1: 518, F2: 2}
    switch "foo bar" {
    case "a":
        fmt.Println(v3, v1)
    case "":
        var v17 uint = 1
        _ = v17
    case "\n\t":
        fmt.Println(v4)
    default:
        v4 = v4
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        var v19 uint = uint(v3)
        v20 := GenS4{F0: "hello", F1: genFn0(-2147483648, -97, 42), F2: ((1.5 < -2.25) || false)}
        _ = fi18
        _ = v19
        _ = v20
    }
    switch v4 {
    case "\n\t":
        for fi21 := 0; fi21 < 1; fi21++ {
            if false {
                if true {
                    var v22 uint16 = 42
                    _ = v22
                }
                fmt.Println(v4)
                if ((false && true) && false) {
                    v23 := GenS5{F0: 118, F1: -219, F2: 42}
                    _ = v23
                } else {
                    var v24 int8 = 8
                    _ = v24
                }
            }
            delete(v2, (fi21 | fi21))
            var v25 int32 = -2147483648
            _ = fi21
            _ = v25
        }
    case "foo bar":
        var v26 int = -821
        _ = v26
    case "hello":
        var v27 uint32 = 42
        _ = v27
    }
    delete(v2, 32768)
    v1 -= (v1 * (v1 + v1))
    switch 127 {
    case 32768:
        v28 := []int{2147483647, -9223372036854775808, -447}
        _ = v28
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        var v30 uint32 = 42
        _ = fi29
        _ = v30
    }
    v31 := GenS1{F0: 2, F1: ((0.0 / ((float64(0.0) + (0.1 + 1e10)) / 1e10)) <= 1.5), F2: v4}
    v1 -= 9223372036854775807
    switch v4 {
    case "hello":
        v1 = (v1 >> ((18446744073709551615) & 63))
    case "foo bar":
        var v32 bool = true
        _ = v32
    }
    var v33 bool = ((v1 ^ v1) >= v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v4
}
