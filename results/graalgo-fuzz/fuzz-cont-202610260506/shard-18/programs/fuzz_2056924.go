package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
func genFn0(p0 int32, p1 uint16) int32 {
    p0 = (p0 * (p0 & 508))
    fmt.Println(p0)
    p0 -= 42
    _ = p0
    _ = p1
    return (-(p0))
}

func main() {
    var v1 uint32 = 42
    var v2 string = "\n\t"
    var v3 float64 = 0.1
    var v4 int64 = 1
    var v5 uint64 = 42
    v6 := []int{}
    v7 := GenS1{F0: (v4 - -993), F1: -127}
    v1 += 4294967295
    for fi8 := 0; fi8 < 18; fi8++ {
        v5 = 96
        if (!(!((!(!true)) && ((65535 < 453) || false)))) {
            for fi9 := 0; fi9 < 15; fi9++ {
                if (1 != 18446744073709551615) {
                    v1 = 2147483648
                } else {
                    v2 = ""
                }
                var v10 uint = 0
                if true {
                    var v11 uint64 = v5
                    _ = v11
                } else {
                    v2 = v2
                    if true {
                        switch -29 {
                        case -9223372036854775808:
                            if (float64(1.5) == 1e10) {
                                var v12 uint64 = (v5 - v5)
                                var v13 int = fi9
                                _ = v12
                                _ = v13
                            } else {
                                v2 = "foo bar"
                            }
                        case 128:
                            var v14 uint32 = 65535
                            _ = v14
                        case -892:
                            var v15 uint64 = v5
                            _ = v15
                        default:
                            switch v2 {
                            case "":
                                if (fi8 < (fi8 / ((fi8) | 1))) {
                                    fmt.Println(v10, v5, v1)
                                    _ = len(v6)
                                    fmt.Println(v4, fi8, v1)
                                }
                            default:
                                var v16 uint32 = 2147483647
                                _ = v16
                            }
                        }
                    }
                    v1 += (v1 ^ 42)
                    var v17 float32 = 1e10
                    _ = v17
                }
                _ = fi9
                _ = v10
            }
            v2 = v2
        } else {
            if (!true) {
                var v18 string = "a"
                fmt.Println(v2, v4)
                if false {
                    v19 := GenS1{F0: v4, F1: -128}
                    for fi20 := 0; fi20 < 11; fi20++ {
                        var v21 uint = 2
                        if true {
                            var v22 bool = (true || ((!false) && true))
                            v22 = true
                            _ = v22
                        } else {
                            var v23 int16 = 258
                            v23 = (^(v23))
                            _ = v23
                        }
                        fmt.Println(v1, v5, v18)
                        _ = fi20
                        _ = v21
                    }
                    _ = v19
                } else {
                    for fi24 := 0; fi24 < 17; fi24++ {
                        if false {
                            if (!(v5 > 18446744073709551615)) {
                                if len(v6) > 0 {
                                    v6[0] = fi24
                                }
                                var v25 string = v18
                                _ = v25
                            }
                        }
                        var v26 bool = (!(v5 >= (v5 | 13)))
                        _ = fi24
                        _ = v26
                    }
                    _ = len(v6)
                    var v27 int32 = genFn0(genFn0(-2147483648, 1), 978)
                    _ = v27
                }
                _ = v18
            }
        }
        for fi28 := 0; fi28 < 20; fi28++ {
            if ((v5 << ((195) & 63)) < 18446744073709551615) {
                v29 := GenS1{F0: (v4 >> ((42) & 63)), F1: 127}
                var v30 int8 = -128
                v31 := GenS1{F0: (-(v4)), F1: (v30 % ((v30) | 1))}
                _ = v29
                _ = v30
                _ = v31
            }
            var v32 uint64 = v5
            _ = fi28
            _ = v32
        }
        _ = fi8
    }
    if len(v6) > 0 {
        v6[0] = 32767
    }
    var v33 uint64 = (v5 >> ((2) & 63))
    if (v5 < v33) {
        v33 += (v33 | 373)
        var v34 uint64 = (v5 - (v33 % (((v5 << ((4294967295) & 63))) | 1)))
        _ = v34
    } else {
        for fi35 := 0; fi35 < 1; fi35++ {
            var v36 uint32 = uint32(v5)
            _ = fi35
            _ = v36
        }
        if false {
            v3 = (v3 + ((v3 - ((v3 * float64(0.1)) - v3)) + float64((float32(0.0) + (1.5 / 0.1)))))
            v4 = 1
            var v37 float32 = (3.14159 * 3.14159)
            _ = v37
        }
    }
    v6 = append(v6, 42)
    var v38 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
