package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 uint16 }
func genFn0(p0 int, p1 int64, p2 uint16) uint64 {
    var v7 bool = (!(((false && (!false)) || (18446744073709551615 < 1)) && false))
    var v8 int64 = p1
    p1 += p1
    v9 := GenS1{F0: "hello", F1: (p0 | (-(p0))), F2: 33}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 int, p1 int64) uint8 {
    if (-2.25 == 1e10) {
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 float32 = 0.0
            var v12 uint16 = 1
            for fi13 := 0; fi13 < 11; fi13++ {
                p0 += (fi13 & (-(fi10)))
                _ = fi13
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        if (!false) {
            v14 := []int{p0}
            for fi15 := 0; fi15 < 6; fi15++ {
                if true {
                    p1 = (p1 + p1)
                }
                var v16 float64 = (1e10 * ((-2.25 / 0.1) - 0.0))
                _ = fi15
                _ = v16
            }
            _ = v14
        } else {
            v17 := []int{}
            v17 = append(v17, p0)
            var v18 uint8 = 239
            _ = v17
            _ = v18
        }
    }
    p0 -= p0
    _ = p0
    _ = p1
    return 255
}

func main() {
    var v1 uint16 = 382
    v2 := []int{42, -1}
    v3 := make(map[int]uint32)
    var v4 string = "foo bar"
    var v5 bool = false
    var v6 uint16 = v1
    v19 := GenS2{F0: v5}
    for fi20 := 0; fi20 < 7; fi20++ {
        v2 = append(v2, fi20)
        var v21 uint = 934
        v22 := GenS1{F0: (((v4 + v4) + (("foo bar" + v4) + v4)) + ((v4 + "a") + v4)), F1: (fi20 << ((v21) & 63)), F2: 42}
        _ = fi20
        _ = v21
        _ = v22
    }
    delete(v3, -899)
    var v23 string = (("a" + (("a" + "foo bar") + (v4 + "hello"))) + v4)
    v1 += v1
    var v24 bool = v5
    v25 := v3[9223372036854775807]
    fmt.Println(v25)
    switch -1 {
    case -2147483648:
        v26 := GenS3{F0: -66, F1: (^(v1))}
        _ = v26
    default:
        v25 = v25
    }
    v3[32767] = 65535
    v25 += v25
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
