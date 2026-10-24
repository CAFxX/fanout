package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint8 }
type GenS2 struct { F0 int8; F1 uint }
func genFn0(p0 uint16, p1 int64, p2 uint32) uint {
    v5 := GenS1{F0: 42, F1: (false || (1 > 0)), F2: 1}
    fmt.Println(p1, p1)
    v6 := []int{-2147483648}
    p0 = (^(p0))
    fmt.Println(p1, p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 uint, p1 int32, p2 string) uint8 {
    var v7 uint = 1
    for fi8 := 0; fi8 < 5; fi8++ {
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 string = (p2 + "\n\t")
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: 2, F1: (!true), F2: 2}
        p0 += (p0 << ((v7) & 63))
        _ = fi8
        _ = v11
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        v7 = 4294967295
        _ = fi12
    }
    var v13 int = 0
    v14 := []int{(v13 >> ((genFn0(782, -737, 65535)) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v7
    return 128
}

func main() {
    var v1 float64 = 1.5
    var v2 int32 = 550
    var v3 uint8 = 0
    var v4 uint16 = 279
    var v15 bool = ((v3 + (v3 | 255)) != (v3 | genFn1(42, 652, "foo bar")))
    v16 := []int{0, -2147483648, int(v4)}
    var v17 uint64 = 162
    v15 = v15
    v18 := GenS1{F0: 1, F1: ((479 != 0) && true), F2: genFn1(331, v2, "")}
    for fi19 := 0; fi19 < 7; fi19++ {
        v16 = append(v16, fi19)
        _ = fi19
    }
    var v20 uint = 42
    var v21 uint64 = (v17 >> (((v20 >> ((v20) & 63))) & 63))
    v22 := GenS2{F0: 86, F1: genFn0((^(v4)), 771, 166)}
    v4 -= 42
    var v23 int16 = -884
    for fi24 := 0; fi24 < 13; fi24++ {
        fmt.Println(v17)
        var v25 uint16 = v4
        v25 = v4
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        switch ("" + "\n\t") {
        case "a":
            if ((false || false) && v15) {
                var v27 int16 = (v23 << ((v20) & 63))
                fmt.Println(fi26, v21)
                if true {
                    if v15 {
                        switch "\n\t" {
                        case "a":
                            if true {
                                var v28 float64 = (v1 / (((1.5 + v1) / (float64(3.14159) - v1)) + (0.0 - ((v1 - float64((1e10 - 1e10))) * 1.5))))
                                var v29 uint8 = (^(v3))
                                v2 = v2
                                _ = v28
                                _ = v29
                            } else {
                                var v30 bool = v15
                                v31 := GenS2{F0: 42, F1: (v20 ^ 4294967295)}
                                for fi32 := 0; fi32 < 18; fi32++ {
                                    var v33 uint16 = (v4 << (((^(v20))) & 63))
                                    switch "a" {
                                    case "\n\t":
                                        if len(v16) >= 2 {
                                            v16 = v16[1:len(v16)]
                                        }
                                    default:
                                        if true {
                                            v27 -= int16(fi26)
                                        } else {
                                            _ = len(v16)
                                        }
                                    }
                                    _ = fi32
                                    _ = v33
                                }
                                _ = v30
                                _ = v31
                            }
                        case "hello":
                            v23 -= -32768
                        default:
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                        }
                    } else {
                        if len(v16) > 0 {
                            v16[0] = fi26
                        }
                    }
                } else {
                    if false {
                        var v34 uint8 = v3
                        for fi35 := 0; fi35 < 1; fi35++ {
                            fmt.Println(v20)
                            if v15 {
                                v15 = (v34 >= uint8(v20))
                                if len(v16) > 0 {
                                    v16[0] = 848
                                }
                                _ = len(v16)
                                fmt.Println(fi35)
                            }
                            fmt.Println(v34, fi26)
                            _ = fi35
                        }
                        _ = v34
                    }
                    if len(v16) > 0 {
                        v16[0] = (-(fi26))
                    }
                    if len(v16) > 0 {
                        v16[0] = 2147483647
                    }
                }
                _ = v27
            }
        case "\n\t":
            v36 := GenS2{F0: 45, F1: 1}
            _ = v36
        case "hello":
            if ((!(!false)) && (v20 > (^(v20)))) {
                v21 = (v17 & (^(v17)))
            }
        }
        var v37 uint8 = genFn1(4294967295, (v2 << ((uint(v17)) & 63)), ("hello" + ("" + "")))
        _ = fi26
        _ = v37
    }
    fmt.Println(v17, v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
