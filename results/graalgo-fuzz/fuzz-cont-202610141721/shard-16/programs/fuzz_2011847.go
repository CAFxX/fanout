package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 int16; F1 string }
func genFn0(p0 int, p1 int) uint16 {
    var v4 float32 = -2.25
    var v5 int64 = -9223372036854775808
    switch "" {
    case "foo bar":
        fmt.Println(v5)
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 32767
}
func genFn1(p0 int16) uint32 {
    v6 := []int{42, -14, 32767}
    p0 -= (p0 >> ((4294967295) & 63))
    _ = len(v6)
    v7 := GenS1{F0: true}
    _ = p0
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 uint = 1
    v2 := make(map[int]bool)
    v3 := []uint32{uint32(v1)}
    v1 = 0
    var v8 uint = v1
    v9 := GenS1{F0: false}
    var v10 uint16 = 42
    switch "\n\t" {
    case "a":
        _ = len(v2)
    case "\n\t":
        fmt.Println(v1, v10, v10)
    case "foo bar":
        v8 -= (^(v1))
    default:
        v11 := GenS2{F0: 42, F1: (false || (448 >= -2147483648))}
        _ = v11
    }
    switch -947 {
    case -452:
        for fi12 := 0; fi12 < 19; fi12++ {
            if false {
                if ("\n\t" > ("" + "hello")) {
                    if true {
                        v13 := GenS2{F0: 32767, F1: (-2147483648 < -2147483648)}
                        if ((!((2 <= uint32(v1)) && true)) || ("foo bar" > "")) {
                            var v14 uint16 = (v10 % ((1) | 1))
                            _ = v14
                        } else {
                            if true {
                                v1 = v1
                            }
                            v15 := v2[(fi12 - (fi12 >> ((v8) & 63)))]
                            var v16 int16 = 506
                            _ = v15
                            _ = v16
                        }
                        _ = v13
                    } else {
                        fmt.Println(v10, v10, v10)
                    }
                    for fi17 := 0; fi17 < 8; fi17++ {
                        _ = len(v2)
                        if false {
                            var v18 uint8 = 255
                            if len(v3) > 0 {
                                v3[0] = 673
                            }
                            v19 := GenS3{F0: 2, F1: true}
                            _ = v18
                            _ = v19
                        }
                        _ = fi17
                    }
                } else {
                    for fi20 := 0; fi20 < 20; fi20++ {
                        var v21 uint64 = 511
                        _ = fi20
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 12; fi22++ {
                        v23 := GenS4{F0: -1, F1: 4294967295}
                        v3 = append(v3, 2)
                        v24 := v2[(fi12 ^ fi22)]
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    fmt.Println(fi12)
                }
                if ((v8 < (v8 - (^(v8)))) || true) {
                    switch "" {
                    case "":
                        v25 := GenS5{F0: -820, F1: ("foo bar" + "a")}
                        _ = v25
                    case "\n\t":
                        var v26 int = fi12
                        _ = v26
                    case "hello":
                        v27 := v2[int(v10)]
                        _ = v27
                    default:
                        if len(v3) > 0 {
                            v3[0] = 755
                        }
                    }
                    v10 = v10
                }
            }
            var v28 bool = false
            if v28 {
                v8 -= (v8 - 0)
                v8 = (v1 & v8)
                _ = len(v2)
            }
            _ = fi12
            _ = v28
        }
    }
    v1 = (v1 % ((65535) | 1))
    fmt.Println(v10, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
