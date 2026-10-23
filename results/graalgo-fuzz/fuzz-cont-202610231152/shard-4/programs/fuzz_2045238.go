package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint8 }
func genFn0() uint8 {
    v6 := []int{}
    _ = len(v6)
    _ = len(v6)
    var v7 uint16 = 2
    fmt.Println(v7)
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int32, p1 int32) uint16 {
    v8 := GenS1{F0: -510}
    p0 += p1
    var v9 float64 = (-2.25 + 3.14159)
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 float32 = 0.0
    var v2 int32 = 0
    var v3 int16 = 1000
    var v4 uint64 = 4294967295
    var v5 uint32 = 2147483647
    for fi10 := 0; fi10 < 20; fi10++ {
        fmt.Println(v4, v2, fi10)
        switch "foo bar" {
        case "\n\t":
            v2 = -1
        case "hello":
            var v11 uint32 = v5
            _ = v11
        }
        _ = fi10
    }
    v12 := []int{}
    var v13 uint16 = genFn1((-(v2)), (v2 << ((1) & 63)))
    switch "a" {
    case "foo bar":
        for fi14 := 0; fi14 < 20; fi14++ {
            v3 = v3
            v13 = genFn1((v2 >> ((681) & 63)), int32(v13))
            if false {
                var v15 int16 = v3
                _ = v15
            } else {
                var v16 uint8 = genFn0()
                _ = v16
            }
            _ = fi14
        }
    default:
        switch 2147483647 {
        case -1:
            v17 := GenS2{F0: 54, F1: v5, F2: 42}
            _ = v17
        default:
            v2 = (-(v2))
        }
    }
    v3 = v3
    var v18 bool = false
    fmt.Println(v13, v1, v5)
    switch 9223372036854775807 {
    case 947:
        switch 504 {
        case 32768:
            v19 := GenS1{F0: 560}
            _ = v19
        case -729:
            if ((v18 || v18) || (v4 >= 1)) {
                if ((v5 / (((v5 >> ((uint(v13)) & 63))) | 1)) != (v5 * (^(v5)))) {
                    v20 := GenS1{F0: int16(v5)}
                    switch "hello" {
                    case "\n\t":
                        v21 := GenS1{F0: (v3 | -32768)}
                        _ = v21
                    case "hello":
                        if v18 {
                            if (!(1.5 >= (((0.0 * 0.0) + v1) * ((0.1 - v1) * -2.25)))) {
                                v18 = false
                                v13 = 0
                                fmt.Println(v13, v2, v5)
                            }
                            for fi22 := 0; fi22 < 16; fi22++ {
                                switch fi22 {
                                case 632:
                                    fmt.Println(v3)
                                default:
                                    v1 = v1
                                }
                                _ = fi22
                            }
                        }
                    }
                    _ = v20
                }
                if true {
                    var v23 int32 = v2
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 7; fi24++ {
                        var v25 float32 = v1
                        _ = fi24
                        _ = v25
                    }
                }
            } else {
                for fi26 := 0; fi26 < 11; fi26++ {
                    for fi27 := 0; fi27 < 17; fi27++ {
                        var v28 uint = 4294967295
                        switch "foo bar" {
                        case "\n\t":
                            var v29 string = ""
                            _ = v29
                        case "":
                            v28 -= v28
                        case "foo bar":
                            var v30 int16 = v3
                            _ = v30
                        }
                        if len(v12) > 0 {
                            v12[0] = (fi26 | (fi27 ^ (^(fi26))))
                        }
                        _ = fi27
                        _ = v28
                    }
                    _ = fi26
                }
                var v31 bool = v18
                v18 = ((!v31) && v18)
                _ = v31
            }
        case 127:
            v4 += 18446744073709551615
        }
    default:
        fmt.Println(v2)
    }
    var v32 uint32 = (v5 >> ((162) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
