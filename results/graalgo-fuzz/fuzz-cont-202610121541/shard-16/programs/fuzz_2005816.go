package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
func genFn0(p0 int16) uint32 {
    switch "foo bar" {
    case "a":
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 bool = (false && true)
            var v8 int64 = -1
            _ = fi6
            _ = v7
            _ = v8
        }
    default:
        var v9 uint8 = 128
        _ = v9
    }
    if (!true) {
        if true {
            if ((p0 >> ((42) & 63)) >= p0) {
                p0 -= p0
            } else {
                switch 32768 {
                case 128:
                    p0 = 32767
                default:
                    switch 9223372036854775807 {
                    case -101:
                        p0 -= p0
                    case 4294967295:
                        v10 := []int{32767}
                        _ = v10
                    }
                }
            }
            fmt.Println(p0)
        } else {
            p0 -= (p0 ^ 549)
            v11 := GenS1{F0: 2147483647, F1: 0}
            _ = v11
        }
        fmt.Println(p0)
    }
    _ = p0
    return 299
}

func main() {
    var v1 uint16 = 45
    var v2 uint64 = 977
    var v3 bool = true
    var v4 uint64 = (v2 << ((584) & 63))
    var v5 uint16 = v1
    v1 -= v1
    for fi12 := 0; fi12 < 17; fi12++ {
        v2 = 18446744073709551615
        _ = fi12
    }
    v3 = (!(2 != 127))
    switch "foo bar" {
    case "a":
        v13 := []int{127, 569, -826, 508}
        _ = v13
    default:
        if true {
            if v3 {
                v14 := []int{32768}
                v2 = (v4 * (v4 << ((uint(v5)) & 63)))
                _ = v14
            }
        }
    }
    v15 := GenS1{F0: 2147483647, F1: -1}
    switch ("" + (("\n\t" + "a") + "hello")) {
    case "hello":
        v16 := GenS1{F0: 0, F1: -707}
        _ = v16
    case "a":
        for fi17 := 0; fi17 < 13; fi17++ {
            if (!(fi17 > fi17)) {
                var v18 int64 = -834
                v19 := GenS1{F0: 2147483647, F1: 649}
                v4 = (^(v4))
                _ = v18
                _ = v19
            }
            v1 -= (v5 / (((v1 * (v1 / ((v5) | 1)))) | 1))
            _ = fi17
        }
    case "":
        for fi20 := 0; fi20 < 18; fi20++ {
            var v21 int64 = 9223372036854775807
            _ = fi20
            _ = v21
        }
    }
    v22 := []int{428}
    if (32767 != 1000) {
        var v23 uint64 = v2
        _ = v23
    }
    v4 -= (^(v4))
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    var v24 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
