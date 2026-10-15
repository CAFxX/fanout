package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint; F2 uint8 }
func genFn0(p0 uint32, p1 bool, p2 bool) int8 {
    for fi7 := 0; fi7 < 13; fi7++ {
        p2 = false
        _ = fi7
    }
    p1 = (524 >= 128)
    fmt.Println(p1, p2)
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 int8, p1 int16, p2 int8) uint32 {
    if (p2 == p2) {
        var v8 uint64 = 1
        fmt.Println(p0)
        v8 -= v8
        _ = v8
    } else {
        v9 := []int{127, 9223372036854775807, -2147483648, 32767}
        var v10 float64 = 0.1
        _ = v10
        _ = v9
    }
    p1 -= p1
    v11 := GenS1{F0: (float32(float64((0.1 + 3.14159))) >= (0.0 + 0.1))}
    var v12 int64 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 533
}

func main() {
    v1 := []string{"\n\t"}
    v2 := []string{}
    var v3 uint16 = 32767
    v4 := make(map[uint32]bool)
    v5 := make(map[string]bool)
    var v6 uint16 = (v3 / (((v3 % ((v3) | 1))) | 1))
    if (65535 < 42) {
        var v13 int16 = 42
        switch "a" {
        case "a":
            switch "hello" {
            case "a":
                v14 := GenS2{F0: true}
                _ = v14
            }
        case "\n\t":
            v13 += v13
        case "hello":
            v15 := GenS2{F0: ("a" >= "foo bar")}
            _ = v15
        default:
            if true {
                fmt.Println(v13, v13, v3)
                v3 -= (^(v3))
                var v16 bool = (-1 == 9223372036854775807)
                _ = v16
            }
        }
        _ = v13
    }
    v6 -= 32767
    _ = len(v2)
    var v17 int16 = 1000
    v18 := GenS1{F0: true}
    fmt.Println(v17, v17, v3)
    v3 += 2
    v3 -= 865
    for fi19 := 0; fi19 < 2; fi19++ {
        v17 = v17
        v20 := GenS1{F0: (false || (!((1e10 / 1.0) < float32(1e10))))}
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        var v22 uint8 = 42
        _ = fi21
        _ = v22
    }
    fmt.Println(v6)
    delete(v4, 4294967295)
    if false {
        v6 = (v6 ^ (v6 << ((909) & 63)))
        v3 = 2
        var v23 int8 = genFn0(uint32(v6), (65535 >= 255), (255 != 800))
        _ = v23
    }
    var v24 uint = 859
    v3 += (v6 | 2)
    fmt.Println(v3)
    v25 := GenS3{F0: 1, F1: (v24 | (v24 % (((v24 | (v24 & v24))) | 1))), F2: 234}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
