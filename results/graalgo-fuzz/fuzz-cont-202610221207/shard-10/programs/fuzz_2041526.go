package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int32) uint16 {
    p0 = (p0 | (p0 << ((1) & 63)))
    var v6 int32 = (p0 / ((p0) | 1))
    _ = p0
    _ = v6
    return 32767
}
func genFn1(p0 int8, p1 bool, p2 uint16) uint {
    fmt.Println(p0)
    v7 := []int{533, 1000}
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: 1, F1: fi8, F2: 4294967295}
        p1 = ((0.0 * (((1.5 + (0.1 * 3.14159)) - 1.5) - 1e10)) >= (0.0 - -2.25))
        var v10 int16 = 32767
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 string = ("foo bar" + "")
        var v13 int16 = 475
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[int]uint32)
    var v3 int16 = 0
    v4 := make(map[string]int)
    v5 := make(map[string]string)
    fmt.Println(v1)
    v3 = (v3 << ((genFn1(1, false, 32767)) & 63))
    switch "\n\t" {
    case "":
        var v14 uint8 = 0
        _ = v14
    }
    v15 := v4[("\n\t" + (("\n\t" + ("a" + "foo bar")) + "foo bar"))]
    _ = len(v5)
    v3 -= 730
    var v16 uint = 69
    if false {
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 int32 = 42
            var v19 uint = v16
            for fi20 := 0; fi20 < 2; fi20++ {
                var v21 uint = 0
                if (9223372036854775807 > -608) {
                    var v22 int64 = 524
                    v23 := []int{fi20, (fi17 % ((718) | 1)), fi17}
                    _ = v22
                    _ = v23
                } else {
                    v18 = v18
                    if (!false) {
                        v24 := GenS2{F0: (v15 ^ (fi17 & (-(fi20)))), F1: -86}
                        var v25 float64 = v1
                        var v26 uint16 = 0
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                }
                v15 = v15
                var v27 bool = (v1 < v1)
                fmt.Println(v1, fi17, v18)
                var v28 uint16 = 387
                v21 = (v19 / (((v21 - genFn1(0, true, genFn0((v18 ^ v18))))) | 1))
                var v29 uint16 = v28
                _ = fi20
                _ = v21
                _ = v27
                _ = v28
                _ = v29
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        v3 = (v3 & int16(v16))
    } else {
        fmt.Println(v1)
        var v30 int8 = 101
        v31 := GenS2{F0: v15, F1: v30}
        _ = v30
        _ = v31
    }
    v32 := GenS3{F0: 20}
    var v33 float64 = ((v1 + v1) - v1)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
