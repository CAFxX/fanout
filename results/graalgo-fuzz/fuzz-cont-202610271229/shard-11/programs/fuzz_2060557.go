package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int) int64 {
    var v4 uint8 = 127
    for fi5 := 0; fi5 < 7; fi5++ {
        fmt.Println(v4, fi5, v4)
        var v6 uint8 = (^(v4))
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        if false {
            v8 := GenS1{F0: 760, F1: "a"}
            if true {
                fmt.Println(v4, p0)
                v4 -= 128
                for fi9 := 0; fi9 < 15; fi9++ {
                    fmt.Println(v4)
                    v10 := []int{fi9, fi7, -2147483648}
                    var v11 uint16 = 32767
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            } else {
                for fi12 := 0; fi12 < 8; fi12++ {
                    var v13 int16 = 42
                    v14 := []int{(fi7 << ((1) & 63))}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
            _ = v8
        } else {
            v15 := GenS1{F0: -9223372036854775808, F1: ("hello" + "\n\t")}
            v16 := GenS1{F0: 552, F1: "foo bar"}
            _ = v15
            _ = v16
        }
        var v17 float32 = 0.1
        v18 := GenS2{F0: 127, F1: 816}
        _ = fi7
        _ = v17
        _ = v18
    }
    if false {
        for fi19 := 0; fi19 < 6; fi19++ {
            if false {
                for fi20 := 0; fi20 < 5; fi20++ {
                    v21 := GenS1{F0: 158, F1: "\n\t"}
                    v4 += v4
                    _ = fi20
                    _ = v21
                }
            }
            _ = fi19
        }
    }
    _ = p0
    _ = v4
    return -9223372036854775808
}
func genFn1(p0 uint) uint64 {
    v22 := []int{}
    v22 = append(v22, int(p0))
    p0 = 18446744073709551615
    p0 += p0
    _ = p0
    _ = v22
    return 212
}

func main() {
    var v1 uint = 887
    var v2 float32 = 1e10
    var v3 float64 = 1.5
    if false {
        var v23 uint8 = 127
        v23 = 128
        var v24 int8 = 0
        _ = v23
        _ = v24
    }
    var v25 int8 = -128
    v3 = ((v3 * v3) / (v3 + v3))
    if ((v25 * (^(v25))) > (v25 ^ -128)) {
        v25 -= (v25 / ((v25) | 1))
    }
    for fi26 := 0; fi26 < 14; fi26++ {
        for fi27 := 0; fi27 < 2; fi27++ {
            v25 = -40
            v28 := GenS3{F0: 42}
            if (((v25 * v25) >= v25) || (370 <= 32767)) {
                v29 := GenS3{F0: -2147483648}
                v25 = (v25 | (v25 - -92))
                _ = v29
            } else {
                switch fi26 {
                case 0:
                    v30 := GenS1{F0: 9223372036854775807, F1: "hello"}
                    _ = v30
                case 32767:
                    if (-1 <= 0) {
                        v31 := GenS2{F0: (v25 << ((4294967295) & 63)), F1: 0}
                        _ = v31
                    }
                default:
                    var v32 int16 = -501
                    _ = v32
                }
            }
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    v1 = v1
    fmt.Println(v25, v1)
    var v33 int8 = v25
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v33
}
