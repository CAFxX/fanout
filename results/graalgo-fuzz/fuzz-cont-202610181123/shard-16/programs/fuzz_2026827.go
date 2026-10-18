package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int; F1 uint16; F2 int64 }
type GenS4 struct { F0 int32; F1 uint32; F2 int32 }
func genFn0(p0 bool, p1 uint16) uint16 {
    var v5 int64 = -9223372036854775808
    v6 := []int{263, 727}
    p1 -= 32767
    v7 := GenS1{F0: (32767 < 4294967295)}
    var v8 int16 = 32767
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return (p1 | p1)
}
func genFn1(p0 uint16, p1 int32) uint8 {
    v9 := []int{9223372036854775807, -2147483648, 649}
    for fi10 := 0; fi10 < 20; fi10++ {
        fmt.Println(p0, p0)
        v11 := GenS1{F0: (!(!(1 >= 0)))}
        v12 := GenS1{F0: false}
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v9
    return 42
}

func main() {
    v1 := []uint32{}
    var v2 float64 = -2.25
    var v3 int8 = 19
    var v4 float64 = 3.14159
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (1 == 127) {
        if (1000 < 0) {
            fmt.Println(v4, v4)
            if ("a" > (("" + "\n\t") + "hello")) {
                v3 += v3
                fmt.Println(v3)
            }
        } else {
            v1 = append(v1, 65535)
            if (4294967295 >= 42) {
                for fi13 := 0; fi13 < 1; fi13++ {
                    for fi14 := 0; fi14 < 10; fi14++ {
                        var v15 int64 = -200
                        _ = len(v1)
                        v15 -= v15
                        v1 = append(v1, 4294967295)
                        if ((v2 / float64(-2.25)) < ((v4 - 1.5) * 0.0)) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        v2 = float64(((1e10 / float32((((v4 / (((1e10 * (v2 * 1.5)) + float64((((-2.25 * 0.1) - -2.25) * 0.1))) + 1e10)) - (v4 + (0.0 / 1.0))) + float64((1e10 / 0.1))))) - 0.0))
                        _ = fi14
                        _ = v15
                    }
                    v3 += (v3 << ((1) & 63))
                    _ = fi13
                }
                v16 := GenS2{F0: 42, F1: -645}
                for fi17 := 0; fi17 < 15; fi17++ {
                    switch "\n\t" {
                    case "":
                        v18 := GenS2{F0: 1, F1: 1}
                        _ = v18
                    case "hello":
                        v4 = v4
                    }
                    v4 -= v4
                    v2 = v2
                    _ = fi17
                }
                _ = v16
            }
            _ = len(v1)
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch -55 {
    case 859:
        for fi19 := 0; fi19 < 4; fi19++ {
            v2 = v2
            _ = len(v1)
            _ = fi19
        }
    default:
        for fi20 := 0; fi20 < 4; fi20++ {
            if true {
                v21 := GenS1{F0: ((42 > v3) && (49 != 127))}
                v22 := GenS1{F0: false}
                var v23 int64 = -1
                _ = v21
                _ = v22
                _ = v23
            } else {
                v4 -= 1e10
                var v24 uint16 = 899
                var v25 int64 = -1
                _ = v24
                _ = v25
            }
            _ = len(v1)
            _ = fi20
        }
    }
    v26 := GenS3{F0: 32768, F1: genFn0(((!(-219 > 32767)) && false), 0), F2: 42}
    if len(v1) > 0 {
        v1[0] = 1
    }
    for fi27 := 0; fi27 < 4; fi27++ {
        _ = len(v1)
        _ = fi27
    }
    v2 = v4
    for fi28 := 0; fi28 < 2; fi28++ {
        v3 += v3
        v29 := GenS4{F0: -2147483648, F1: 65535, F2: -2147483648}
        _ = fi28
        _ = v29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
