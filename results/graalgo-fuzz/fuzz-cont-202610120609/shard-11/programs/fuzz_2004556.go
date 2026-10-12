package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 int64; F1 int16; F2 uint16 }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint, p1 uint8) uint16 {
    var v5 float64 = 3.14159
    v6 := GenS1{F0: 0, F1: -1, F2: 127}
    if false {
        var v7 uint8 = (p1 % (((p1 >> ((p0) & 63))) | 1))
        p0 += (p0 << ((p0) & 63))
        v8 := GenS2{F0: (v7 << ((p0) & 63))}
        _ = v7
        _ = v8
    }
    switch 127 {
    case 155:
        var v9 uint = p0
        _ = v9
    case 32767:
        v5 -= (0.1 - v5)
    case 0:
        p0 = (p0 >> ((0) & 63))
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return uint16(p0)
}
func genFn1(p0 bool, p1 int) int {
    var v10 int64 = -9223372036854775808
    if (((!(-761 > 1)) || (!p0)) && p0) {
        var v11 int = (^(p1))
        fmt.Println(p1)
        _ = v11
    }
    var v12 uint64 = 1
    _ = p0
    _ = p1
    _ = v10
    _ = v12
    return p1
}

func main() {
    v1 := []int{0, -557}
    v2 := []uint32{}
    var v3 uint8 = 106
    var v4 int32 = 42
    var v13 int32 = (v4 / (((-(v4))) | 1))
    if (!("" <= ("" + "\n\t"))) {
        var v14 uint8 = v3
        if true {
            var v15 uint64 = 154
            v13 -= 0
            v16 := GenS3{F0: (v3 - (v3 ^ (v14 ^ 128))), F1: ""}
            _ = v15
            _ = v16
        } else {
            fmt.Println(v3, v14, v4)
        }
        _ = v14
    } else {
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v4 -= v4
    v2 = append(v2, 2)
    fmt.Println(v3, v3, v3)
    switch -626 {
    case -208:
        if true {
            v13 += (-(v13))
        } else {
            fmt.Println(v4, v13)
            v2 = append(v2, 0)
        }
    case 32767:
        switch "a" {
        case "\n\t":
            v17 := GenS2{F0: 1}
            _ = v17
        case "":
            if true {
                v13 += v4
                v4 -= v4
            }
        }
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        v3 = v3
        for fi19 := 0; fi19 < 12; fi19++ {
            v4 = (-(v4))
            fmt.Println(fi18, v4, fi19)
            _ = fi19
        }
        _ = fi18
    }
    _ = len(v1)
    v20 := GenS3{F0: 127, F1: "foo bar"}
    switch "a" {
    case "\n\t":
        if len(v1) > 0 {
            v1[0] = 42
        }
    }
    v13 = (v4 ^ v13)
    if true {
        v21 := GenS1{F0: 42, F1: 9223372036854775807, F2: -431}
        _ = v21
    } else {
        v3 = v3
        switch ("a" + "foo bar") {
        case "":
            if true {
                v4 = 25
                if len(v2) > 0 {
                    v2[0] = 342
                }
                v2 = append(v2, 975)
            } else {
                if true {
                    v22 := GenS4{F0: 9223372036854775807, F1: 0, F2: 255}
                    if ((v13 & 0) == v4) {
                        var v23 int16 = 1
                        _ = v23
                    }
                    _ = v22
                } else {
                    fmt.Println(v3, v4, v13)
                    if true {
                        v1 = append(v1, 42)
                        if (((!(v3 <= 1)) || false) || true) {
                            fmt.Println(v4, v3)
                            if (0 >= -512) {
                                if len(v2) > 0 {
                                    v2[0] = 65535
                                }
                                for fi24 := 0; fi24 < 16; fi24++ {
                                    if false {
                                        if (886 != v4) {
                                            if false {
                                                v3 += (^(v3))
                                            } else {
                                                if false {
                                                    var v25 bool = false
                                                    switch (fi24 + fi24) {
                                                    case 2147483647:
                                                        fmt.Println(v13)
                                                    }
                                                    _ = len(v1)
                                                    _ = v25
                                                }
                                                v3 -= (v3 >> ((18446744073709551615) & 63))
                                                v4 -= v4
                                            }
                                        }
                                        var v26 uint8 = v3
                                        _ = v26
                                    }
                                    fmt.Println(v4, v3)
                                    _ = fi24
                                }
                                _ = len(v2)
                            } else {
                                for fi27 := 0; fi27 < 9; fi27++ {
                                    var v28 uint = 18446744073709551615
                                    fmt.Println(v3, v4)
                                    _ = fi27
                                    _ = v28
                                }
                            }
                            v2 = append(v2, 212)
                        }
                        v29 := GenS5{F0: 100}
                        _ = v29
                    }
                    fmt.Println(v3)
                }
                var v30 bool = (((-2.25 * -2.25) + 3.14159) < 0.1)
                v30 = (42 >= 2147483647)
                _ = v30
            }
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "foo bar":
            fmt.Println(v3, v3)
        }
        if len(v2) > 0 {
            v2[0] = 544
        }
    }
    switch "foo bar" {
    case "a":
        var v31 int32 = (v13 << ((2) & 63))
        _ = v31
    case "hello":
        switch "foo bar" {
        case "hello":
            var v32 uint = 0
            _ = v32
        case "a":
            v3 += (v3 - v3)
        default:
            _ = len(v1)
        }
    case "\n\t":
        var v33 string = ("a" + ("a" + ("" + (("foo bar" + "a") + "\n\t"))))
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
