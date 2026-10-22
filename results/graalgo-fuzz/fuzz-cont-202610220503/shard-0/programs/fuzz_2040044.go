package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 int8 }
func genFn0(p0 int16, p1 int16) string {
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 int16 = -1
        _ = fi6
        _ = v7
    }
    v8 := []int{768}
    for fi9 := 0; fi9 < 3; fi9++ {
        p1 -= (p1 + p1)
        _ = fi9
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v8
    return "a"
}
func genFn1(p0 int) uint {
    p0 += 132
    p0 = p0
    _ = p0
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 0
    v3 := []int{}
    var v4 bool = (1 > 663)
    var v5 uint = 1
    v2 = 726
    switch ("a" + ("\n\t" + "hello")) {
    case "hello":
        var v10 float64 = -2.25
        _ = v10
    case "a":
        v11 := GenS1{F0: true, F1: 128}
        _ = v11
    case "\n\t":
        for fi12 := 0; fi12 < 2; fi12++ {
            if v4 {
                var v13 int32 = 42
                if (9223372036854775807 >= 1) {
                    for fi14 := 0; fi14 < 1; fi14++ {
                        v15 := GenS2{F0: -128}
                        v3 = append(v3, fi14)
                        var v16 uint16 = 409
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    var v17 float32 = float32(1e10)
                    _ = v17
                } else {
                    var v18 uint = 18446744073709551615
                    v5 = v18
                    if ("\n\t" == "a") {
                        _ = len(v3)
                        if v4 {
                            fmt.Println(v13)
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v5 = (v5 & (^(v5)))
                        }
                    } else {
                        for fi19 := 0; fi19 < 15; fi19++ {
                            v20 := GenS3{F0: 0, F1: uint8(v2), F2: 1}
                            v3 = append(v3, fi19)
                            var v21 uint8 = 197
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                    }
                    _ = v18
                }
                fmt.Println(v4, v2)
                _ = v13
            }
            _ = fi12
        }
    default:
        v1[2147483647] = v4
    }
    v2 -= v2
    var v22 bool = ((^(v5)) == v5)
    v22 = v4
    v1[4294967295] = (v22 || v22)
    switch "\n\t" {
    case "foo bar":
        var v23 int16 = -21
        _ = v23
    case "a":
        v1[2] = v22
    case "hello":
        var v24 bool = (2147483647 != 2147483648)
        _ = v24
    default:
        for fi25 := 0; fi25 < 13; fi25++ {
            v3 = append(v3, (^(fi25)))
            for fi26 := 0; fi26 < 3; fi26++ {
                fmt.Println(v2, v2, v4)
                _ = fi26
            }
            _ = fi25
        }
    }
    fmt.Println(v4, v4)
    v2 = 645
    v4 = v4
    switch "hello" {
    case "foo bar":
        for fi27 := 0; fi27 < 19; fi27++ {
            v28 := GenS1{F0: (!((2 > 4294967295) && (uint8(v2) >= 128))), F1: 128}
            v2 = uint64(v5)
            v3 = append(v3, (fi27 & 4294967295))
            _ = fi27
            _ = v28
        }
    }
    v29 := GenS4{F0: 0, F1: 2147483648}
    var v30 uint16 = uint16(v2)
    if false {
        var v31 float64 = 1e10
        fmt.Println(v31, v5)
        _ = v31
    }
    if v22 {
        switch "hello" {
        case "foo bar":
            fmt.Println(v5)
        default:
            v32 := GenS5{F0: 0}
            _ = v32
        }
        for fi33 := 0; fi33 < 11; fi33++ {
            v1[65535] = ((v5 * 255) < (v5 << (((v5 & v5)) & 63)))
            _ = fi33
        }
    }
    v1[42] = (-2.25 != 3.14159)
    v4 = v4
    switch "\n\t" {
    case "":
        var v34 int8 = -22
        _ = v34
    case "hello":
        if (1 != -9223372036854775808) {
            var v35 bool = v22
            _ = v35
        }
    default:
        var v36 int64 = 405
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
