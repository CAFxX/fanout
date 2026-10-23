package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 int8; F1 uint64 }
func genFn0(p0 int16, p1 bool, p2 uint32) uint32 {
    var v7 int16 = (p0 & p0)
    p1 = (!false)
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := []int{(fi9 - fi8), (fi8 % ((-352) | 1)), 32767, (^(fi9))}
            _ = fi9
            _ = v10
        }
        if p1 {
            for fi11 := 0; fi11 < 8; fi11++ {
                p0 += 42
                p2 -= p2
                p0 = p0
                fmt.Println(p2)
                _ = fi11
            }
            fmt.Println(p0, v7, v7)
            for fi12 := 0; fi12 < 3; fi12++ {
                var v13 uint32 = p2
                if (127 == 42) {
                    v14 := []int{}
                    v15 := GenS1{F0: -127, F1: 0}
                    _ = v14
                    _ = v15
                }
                fmt.Println(p2, v7, v13)
                _ = fi12
                _ = v13
            }
        }
        v16 := GenS1{F0: 0, F1: 59}
        _ = fi8
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p2
}
func genFn1(p0 int16, p1 int16, p2 uint64) uint32 {
    p0 -= (p0 % ((p1) | 1))
    for fi17 := 0; fi17 < 7; fi17++ {
        for fi18 := 0; fi18 < 3; fi18++ {
            v19 := []int{(-(fi17)), (fi18 % ((-627) | 1)), fi17, (fi17 % (((fi18 ^ fi18)) | 1))}
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    p0 += 32767
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 uint8 = 42
        if (42 < (p0 << ((1) & 63))) {
            for fi22 := 0; fi22 < 2; fi22++ {
                v23 := GenS2{F0: p2}
                v24 := []int{(fi20 | (fi22 ^ 29)), (fi22 % ((fi20) | 1)), fi20, fi22}
                var v25 float32 = 3.14159
                if true {
                    v21 -= 1
                    if len(v24) > 0 {
                        v24[0] = (fi20 % ((fi22) | 1))
                    }
                    p1 = 1000
                }
                v26 := GenS1{F0: int32(fi20), F1: 219}
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
                _ = v26
            }
            v27 := GenS1{F0: 0, F1: (v21 - 2)}
            v28 := GenS1{F0: 2147483647, F1: 255}
            _ = v27
            _ = v28
        }
        v29 := GenS2{F0: p2}
        _ = fi20
        _ = v21
        _ = v29
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 uint8 = 51
    var v2 int16 = -1
    var v3 string = ("a" + "\n\t")
    var v4 uint = 18446744073709551615
    v5 := make(map[string]string)
    var v6 int = 445
    v30 := GenS1{F0: 1, F1: v1}
    v6 = -9223372036854775808
    fmt.Println(v3, v3)
    var v31 uint16 = 0
    v32 := GenS3{F0: 598, F1: (42 == -687)}
    v31 -= (^(v31))
    v33 := []int{(v6 & (v6 / (((^(v6))) | 1)))}
    if true {
        var v34 int64 = 42
        v35 := GenS4{F0: 42, F1: 4294967295}
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
