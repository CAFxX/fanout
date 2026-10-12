package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint32 }
func genFn0(p0 string, p1 uint32) uint32 {
    p1 += (p1 ^ p1)
    var v6 uint64 = 4294967295
    p0 = p0
    v7 := []int{988, 575}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1(p0 bool) uint32 {
    fmt.Println(p0)
    v8 := []int{-1, 2147483647, 0}
    _ = p0
    _ = v8
    return 393
}

func main() {
    v1 := []string{}
    v2 := []int{32767, 2147483647, -1, 4294967295}
    v3 := []string{}
    v4 := []int{-2147483648, -118, 207, -321}
    var v5 float64 = float64(((1e10 / (-2.25 * 1.5)) / 1e10))
    if len(v4) > 0 {
        v4[0] = 42
    }
    fmt.Println(v5)
    var v9 uint16 = 2
    for fi10 := 0; fi10 < 8; fi10++ {
        v9 -= 0
        var v11 string = "\n\t"
        _ = fi10
        _ = v11
    }
    var v12 int = 1
    v9 += v9
    switch v12 {
    case 42:
        switch v12 {
        case -1:
            v5 += v5
        default:
            var v13 int64 = 0
            _ = v13
        }
    case 9223372036854775807:
        if false {
            v12 = (v12 | (v12 | (v12 * 32768)))
        }
    }
    v12 += 32768
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v12 += (-(fi14))
        var v15 uint = 18446744073709551615
        _ = fi14
        _ = v15
    }
    var v16 uint64 = 1
    if ((v16 / ((v16) | 1)) <= v16) {
        for fi17 := 0; fi17 < 13; fi17++ {
            if true {
                if (!false) {
                    var v18 int8 = 9
                    var v19 uint = 1
                    if false {
                        v20 := GenS1{F0: 2147483648, F1: v16, F2: 785}
                        _ = v20
                    }
                    _ = v18
                    _ = v19
                }
            } else {
                if len(v4) > 0 {
                    v4[0] = fi17
                }
                v21 := GenS1{F0: 2147483648, F1: 222, F2: 2147483647}
                _ = v21
            }
            var v22 uint8 = 250
            if false {
                var v23 uint16 = (v9 % ((uint16(v22)) | 1))
                _ = v23
            }
            _ = fi17
            _ = v22
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
