package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 uint32 }
func genFn0(p0 uint8) int8 {
    v4 := []int{-676, 42}
    v5 := GenS1{F0: 1, F1: 1}
    v6 := GenS1{F0: 549, F1: -678}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return -128
}

func main() {
    var v1 uint8 = 255
    var v2 bool = (32768 != -517)
    var v3 uint16 = 1
    v7 := []int{127, 1}
    var v8 int8 = 1
    switch 212 {
    case 42:
        var v9 uint8 = v1
        _ = v9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        for fi11 := 0; fi11 < 15; fi11++ {
            v2 = (4294967295 <= fi11)
            v12 := GenS2{F0: (v1 >> ((997) & 63)), F1: 42}
            v8 = (v8 - 1)
            _ = fi11
            _ = v12
        }
        var v13 float32 = (float32((3.14159 + 1.5)) - 0.1)
        _ = fi10
        _ = v13
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        v3 += (v3 % ((v3) | 1))
        v2 = v2
        _ = fi14
    }
    if v2 {
        v15 := GenS3{F0: v2}
        switch 4294967295 {
        case -2147483648:
            v8 = genFn0((v1 / ((0) | 1)))
        default:
            v16 := GenS4{F0: v8, F1: 42}
            _ = v16
        }
        for fi17 := 0; fi17 < 9; fi17++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            v2 = (v1 > uint8(v3))
            _ = fi17
        }
        _ = v15
    } else {
        var v18 uint64 = 4294967295
        if v2 {
            if true {
                if true {
                    for fi19 := 0; fi19 < 19; fi19++ {
                        if len(v7) > 0 {
                            v7[0] = (-(fi19))
                        }
                        _ = fi19
                    }
                    for fi20 := 0; fi20 < 18; fi20++ {
                        for fi21 := 0; fi21 < 1; fi21++ {
                            var v22 int64 = 9223372036854775807
                            v8 += (v8 ^ genFn0(v1))
                            _ = fi21
                            _ = v22
                        }
                        _ = fi20
                    }
                } else {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    if len(v7) > 0 {
                        v7[0] = 4294967295
                    }
                }
                _ = len(v7)
                v7 = append(v7, 2147483647)
            } else {
                v1 = v1
                v7 = append(v7, 172)
                fmt.Println(v2, v1, v1)
            }
            var v23 float64 = 0.0
            _ = v23
        }
        var v24 bool = false
        _ = v18
        _ = v24
    }
    v25 := GenS2{F0: (^(v1)), F1: 950}
    v7 = append(v7, 817)
    for fi26 := 0; fi26 < 13; fi26++ {
        var v27 int8 = int8(v3)
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v7
    _ = v8
}
