package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint; F2 int }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8; F1 int32 }
func genFn0() int8 {
    switch 210 {
    case -71:
        v4 := GenS1{F0: 302}
        _ = v4
    }
    var v5 int = -979
    v6 := GenS2{F0: 4294967295, F1: 1}
    var v7 bool = true
    for fi8 := 0; fi8 < 17; fi8++ {
        v5 += (v5 / ((fi8) | 1))
        _ = fi8
    }
    _ = v5
    _ = v6
    _ = v7
    return 127
}
func genFn1(p0 uint32, p1 uint8, p2 int16) uint64 {
    for fi9 := 0; fi9 < 17; fi9++ {
        p0 += 65535
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := GenS3{F0: 2, F1: 18446744073709551615, F2: fi9}
            _ = fi10
            _ = v11
        }
        switch "\n\t" {
        case "foo bar":
            for fi12 := 0; fi12 < 16; fi12++ {
                p1 -= (p1 + (p1 ^ 128))
                _ = fi12
            }
        case "":
            for fi13 := 0; fi13 < 9; fi13++ {
                p2 = p2
                v14 := GenS4{F0: 420}
                _ = fi13
                _ = v14
            }
        default:
            p2 = (p2 & (p2 + (p2 % (((p2 % ((p2) | 1))) | 1))))
        }
        _ = fi9
    }
    v15 := []int{32767}
    var v16 uint64 = 18446744073709551615
    if len(v15) > 0 {
        v15[0] = 4294967295
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return v16
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 140
    v3 := []int{4294967295, 42, 32768}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1[0] = "hello"
    v2 += 42
    var v17 int64 = -593
    var v18 uint64 = 0
    v3 = append(v3, 128)
    var v19 int = 2147483647
    if ((347 >= -2147483648) && true) {
        switch "" {
        case "a":
            v2 = v2
        }
        fmt.Println(v19, v18)
        for fi20 := 0; fi20 < 13; fi20++ {
            v3 = append(v3, (fi20 / ((v19) | 1)))
            for fi21 := 0; fi21 < 18; fi21++ {
                switch ("a" + ("foo bar" + ("a" + "\n\t"))) {
                case "hello":
                    if (755 <= uint(v2)) {
                        if ((v19 - 127) <= (fi20 ^ fi21)) {
                            var v22 uint16 = uint16(v18)
                            fmt.Println(v22, fi21)
                            if true {
                                if (0.0 > -2.25) {
                                    v17 = v17
                                    var v23 uint = 65535
                                    _ = v23
                                } else {
                                    var v24 int8 = 0
                                    fmt.Println(fi20, v17)
                                    if (1 < 0) {
                                        v25 := GenS2{F0: 0, F1: 42}
                                        switch (-(v19)) {
                                        case -9223372036854775808:
                                            if len(v3) > 0 {
                                                v3[0] = fi20
                                            }
                                        default:
                                            var v26 uint8 = 128
                                            _ = v26
                                        }
                                        _ = v25
                                    }
                                    _ = v24
                                }
                                if false {
                                    v3 = append(v3, 865)
                                    v27 := GenS2{F0: 287, F1: v17}
                                    var v28 int32 = 1
                                    _ = v27
                                    _ = v28
                                }
                            } else {
                                var v29 uint64 = (v18 | 42)
                                v30 := GenS2{F0: 0, F1: (v17 * v17)}
                                _ = v29
                                _ = v30
                            }
                            _ = v22
                        }
                        if (-113 <= 0) {
                            fmt.Println(fi21, v17, v18)
                            var v31 uint64 = 0
                            fmt.Println(fi20, v31)
                            var v32 bool = (-57 != -1)
                            _ = v31
                            _ = v32
                        }
                    } else {
                        v18 = genFn1(65535, (v2 + (v2 * (v2 % ((v2) | 1)))), -1)
                    }
                    var v33 float64 = 0.1
                    _ = v33
                case "foo bar":
                    v1[9] = "hello"
                }
                _ = fi21
            }
            delete(v1, 974)
            _ = fi20
        }
    }
    v34 := GenS5{F0: 0, F1: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v34
}
