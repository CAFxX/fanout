package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 uint; F2 int32 }
type GenS5 struct { F0 uint8; F1 int16; F2 uint }
func genFn0(p0 uint32, p1 uint, p2 int64) int16 {
    v4 := GenS1{F0: 266}
    fmt.Println(p1)
    p1 += (p1 / ((p1) | 1))
    switch "a" {
    case "\n\t":
        p0 = (p0 / ((p0) | 1))
    case "foo bar":
        p2 = p2
    default:
        v5 := []int{42, 32767}
        _ = v5
    }
    var v6 int8 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return -82
}

func main() {
    var v1 float64 = 0.1
    var v2 uint32 = 4294967295
    var v3 int16 = 0
    if false {
        fmt.Println(v2, v2)
        var v7 int16 = (v3 % (((v3 + -133)) | 1))
        _ = v7
    }
    fmt.Println(v3, v2, v2)
    v2 = (v2 + (v2 | 633))
    fmt.Println(v1, v2, v1)
    v8 := []int{-1, 128, 42, -466}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    switch 32768 {
    case 266:
        var v9 float64 = -2.25
        _ = v9
    default:
        if true {
            fmt.Println(v3, v3, v1)
            switch -2147483648 {
            case 2147483647:
                v3 = v3
            }
            switch "\n\t" {
            case "foo bar":
                for fi10 := 0; fi10 < 18; fi10++ {
                    v11 := GenS1{F0: 0}
                    _ = fi10
                    _ = v11
                }
            case "a":
                var v12 string = ((("foo bar" + "\n\t") + "foo bar") + "")
                _ = v12
            default:
                if false {
                    fmt.Println(v1, v1, v3)
                    switch "\n\t" {
                    case "hello":
                        var v13 bool = (!((!false) && true))
                        _ = v13
                    case "\n\t":
                        fmt.Println(v3)
                    default:
                        if ((v2 + (v2 % ((v2) | 1))) >= (v2 >> ((1) & 63))) {
                            if true {
                                v14 := GenS1{F0: 65535}
                                for fi15 := 0; fi15 < 1; fi15++ {
                                    v16 := GenS2{F0: (v3 << ((339) & 63))}
                                    v3 = genFn0(2, 0, 42)
                                    _ = fi15
                                    _ = v16
                                }
                                var v17 bool = false
                                _ = v14
                                _ = v17
                            } else {
                                v1 += (v1 + (((v1 + (0.1 * v1)) - 3.14159) / v1))
                                v1 -= float64(((0.0 / 1.5) * float32(float64(3.14159))))
                                _ = len(v8)
                            }
                        }
                    }
                    v1 = (v1 - (v1 / v1))
                } else {
                    switch -1 {
                    case 127:
                        var v18 int64 = -360
                        _ = v18
                    case 42:
                        var v19 string = ""
                        _ = v19
                    default:
                        v20 := GenS3{F0: uint(v2)}
                        _ = v20
                    }
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    var v21 int32 = -286
                    _ = v21
                }
            }
        }
    }
    if len(v8) > 0 {
        v8[0] = 1
    }
    if (((9223372036854775807 < 743) && (74 == v2)) && true) {
        if false {
            for fi22 := 0; fi22 < 18; fi22++ {
                fmt.Println(v3, v2, v1)
                _ = fi22
            }
        }
        v1 = v1
    }
    switch -9223372036854775808 {
    case 42:
        v1 -= ((v1 - v1) / float64(float32(v1)))
    case 9223372036854775807:
        if (false || (1.5 == 3.14159)) {
            v23 := GenS1{F0: 597}
            _ = v23
        } else {
            var v24 uint = 407
            _ = v24
        }
    case 82:
        if true {
            if false {
                var v25 int8 = -111
                for fi26 := 0; fi26 < 15; fi26++ {
                    v1 -= (1e10 / float64((1.5 + -2.25)))
                    switch (^(fi26)) {
                    case 128:
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                    }
                    v27 := GenS3{F0: 976}
                    _ = fi26
                    _ = v27
                }
                v28 := GenS4{F0: 65535, F1: 1, F2: 2147483647}
                _ = v25
                _ = v28
            }
        } else {
            v29 := GenS5{F0: 2, F1: v3, F2: 4294967295}
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
}
