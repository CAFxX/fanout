package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 bool }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int64; F1 int16 }
func genFn0(p0 bool) string {
    v4 := []int{32768, -473, 0}
    p0 = (991 < 65535)
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 uint = 409
        var v7 int64 = -1
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = v4
    return "a"
}
func genFn1(p0 int8, p1 string, p2 bool) int {
    var v8 float32 = 0.0
    v9 := GenS1{F0: -9223372036854775808, F1: 2, F2: p2}
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 12; fi11++ {
            switch fi11 {
            case 4294967295:
                v12 := GenS2{F0: 65535, F1: fi10}
                var v13 float64 = 3.14159
                _ = v12
                _ = v13
            default:
                var v14 int8 = (-(p0))
                _ = v14
            }
            var v15 uint64 = 4294967295
            _ = fi11
            _ = v15
        }
        _ = fi10
    }
    v16 := []int{-500, -252, -9223372036854775808}
    if true {
        v8 = v8
        p0 += 42
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v8
    _ = v9
    return -789
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 42
    var v3 uint = 4294967295
    v3 = (v3 * (v3 | (v3 + uint(v2))))
    for fi17 := 0; fi17 < 18; fi17++ {
        for fi18 := 0; fi18 < 20; fi18++ {
            switch "foo bar" {
            case "":
                v19 := GenS2{F0: 42, F1: fi17}
                delete(v1, 292)
                v20 := GenS1{F0: 42, F1: v2, F2: (2 < 4294967295)}
                v21 := GenS2{F0: 1, F1: fi17}
                _ = v19
                _ = v20
                _ = v21
            case "a":
                var v22 float32 = 3.14159
                _ = v22
            default:
                var v23 string = "hello"
                _ = v23
            }
            if true {
                var v24 uint = 389
                v3 += 18446744073709551615
                _ = v24
            }
            _ = fi18
        }
        fmt.Println(fi17, v2, v3)
        for fi25 := 0; fi25 < 20; fi25++ {
            v3 += (v3 / (((v3 >> ((496) & 63))) | 1))
            _ = fi25
        }
        _ = fi17
    }
    delete(v1, v2)
    v26 := []int{-9223372036854775808, genFn1(8, "\n\t", false), 555, 9223372036854775807}
    v3 += v3
    _ = len(v26)
    if (false && false) {
        for fi27 := 0; fi27 < 13; fi27++ {
            switch "" {
            case "hello":
                v3 = 477
            case "a":
                var v28 int16 = 1000
                _ = v28
            default:
                if false {
                    v29 := GenS1{F0: 797, F1: (v2 >> ((v3) & 63)), F2: (!false)}
                    switch fi27 {
                    case 2147483647:
                        v3 -= v3
                    }
                    _ = v29
                }
            }
            var v30 uint = (v3 | v3)
            v31 := GenS3{F0: 1, F1: 42}
            _ = fi27
            _ = v30
            _ = v31
        }
        v2 -= v2
        v1[v2] = v2
    }
    if (670 >= -976) {
        _ = len(v1)
    }
    v3 -= (v3 / ((v3) | 1))
    v32 := v1[(v2 ^ v2)]
    fmt.Println(v3)
    v2 = 42
    if (0.0 >= float32(1e10)) {
        v2 = 625
    }
    if false {
        v33 := GenS4{F0: -9223372036854775808, F1: 42}
        var v34 int = 2147483647
        _ = v33
        _ = v34
    } else {
        var v35 float32 = (3.14159 / 1.0)
        _ = v35
    }
    if len(v26) > 0 {
        v26[0] = 128
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
}
