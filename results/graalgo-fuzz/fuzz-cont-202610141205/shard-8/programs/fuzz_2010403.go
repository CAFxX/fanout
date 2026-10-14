package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint }
type GenS2 struct { F0 uint32; F1 int32; F2 int32 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 uint8; F1 int64 }
func genFn0(p0 uint64) uint {
    if false {
        var v7 int32 = 42
        _ = v7
    } else {
        if false {
            if (false && true) {
                var v8 int8 = 100
                v8 -= 17
                v9 := []int{0, 9223372036854775807, 32768}
                _ = v8
                _ = v9
            } else {
                if (!(0 >= -79)) {
                    v10 := []int{127, 128, -1, 128}
                    switch "\n\t" {
                    case "hello":
                        p0 -= 986
                    case "a":
                        if ((-688 == -32768) || true) {
                            var v11 int32 = 0
                            _ = v11
                        }
                    case "foo bar":
                        p0 = 42
                    default:
                        var v12 int8 = 42
                        _ = v12
                    }
                    _ = v10
                } else {
                    p0 = 2
                    var v13 int = 128
                    _ = v13
                }
                v14 := GenS1{F0: 65535, F1: (p0 / (((p0 ^ p0)) | 1)), F2: 255}
                _ = v14
            }
            v15 := []int{42, 127, -1, 32768}
            if (-125 > -852) {
                if len(v15) > 0 {
                    v15[0] = 4294967295
                }
            } else {
                var v16 uint64 = (^(p0))
                _ = v16
            }
            _ = v15
        } else {
            p0 = p0
            fmt.Println(p0)
        }
        v17 := GenS2{F0: 2, F1: -2147483648, F2: 358}
        _ = v17
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        var v19 int32 = 2147483647
        fmt.Println(fi18, fi18)
        _ = fi18
        _ = v19
    }
    _ = p0
    return 65535
}

func main() {
    var v1 uint64 = 42
    v2 := []string{"foo bar", "\n\t", ("" + ""), "hello"}
    v3 := make(map[uint32]int)
    v4 := []string{("\n\t" + "hello")}
    var v5 uint8 = 169
    var v6 uint32 = 42
    v20 := GenS1{F0: v6, F1: v1, F2: 4294967295}
    for fi21 := 0; fi21 < 19; fi21++ {
        v22 := GenS3{F0: (false && false), F1: 42}
        if true {
            v6 = (v6 / ((v6) | 1))
            v5 = (v5 | 61)
        } else {
            v5 -= 255
            var v23 uint8 = v5
            _ = v23
        }
        if (!true) {
            var v24 int = (-(fi21))
            if false {
                v6 -= 1
                if ((-32768 == 0) || ((((false || false) || (-688 == 1)) || (true && false)) && false)) {
                    fmt.Println(v24, v1, fi21)
                    var v25 int = (fi21 % ((fi21) | 1))
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 20; fi26++ {
                        v24 += fi26
                        var v27 int64 = -239
                        switch "foo bar" {
                        case "\n\t":
                            if false {
                                var v28 uint = 2
                                if len(v2) > 0 {
                                    v2[0] = "a"
                                }
                                _ = v28
                            } else {
                                var v29 uint32 = (v6 * (v6 >> ((444) & 63)))
                                var v30 bool = (false || true)
                                _ = v29
                                _ = v30
                            }
                        default:
                            fmt.Println(v24)
                        }
                        _ = fi26
                        _ = v27
                    }
                    var v31 uint32 = (^(v6))
                    if false {
                        if true {
                            v32 := GenS4{F0: 0, F1: (v6 << ((4294967295) & 63))}
                            _ = v32
                        }
                        switch (("foo bar" + "") + "a") {
                        case "foo bar":
                            for fi33 := 0; fi33 < 3; fi33++ {
                                v34 := GenS3{F0: true, F1: 67}
                                v3[(^(v6))] = (fi33 % ((127) | 1))
                                _ = fi33
                                _ = v34
                            }
                        case "\n\t":
                            for fi35 := 0; fi35 < 18; fi35++ {
                                v6 += (v6 % ((v31) | 1))
                                _ = fi35
                            }
                        default:
                            fmt.Println(fi21, v31)
                        }
                        v5 = (v5 | (^(v5)))
                    }
                    _ = v31
                }
                fmt.Println(fi21, v5, v1)
            }
            for fi36 := 0; fi36 < 17; fi36++ {
                v37 := GenS1{F0: v6, F1: v1, F2: 599}
                var v38 int = 32768
                _ = fi36
                _ = v37
                _ = v38
            }
            _ = v24
        } else {
            v39 := GenS5{F0: (v5 << ((836) & 63)), F1: -1}
            _ = v39
        }
        _ = fi21
        _ = v22
    }
    fmt.Println(v1, v1, v1)
    var v40 string = ""
    var v41 int16 = 841
    v41 = v41
    fmt.Println(v41, v6, v6)
    v40 = (v40 + v40)
    var v42 uint32 = uint32(v5)
    var v43 float32 = (3.14159 - ((((-2.25 - 3.14159) - (float32((-2.25 + 3.14159)) / (0.1 / (float32(float64(0.0)) + (0.1 + 0.0))))) - 0.0) * (-2.25 - 0.1)))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v44 uint = 255
    v45 := GenS1{F0: v6, F1: (^(v1)), F2: uint(v5)}
    if (true || (true && true)) {
        v3[v42] = -995
        fmt.Println(v44)
    }
    var v46 float64 = (1e10 - (3.14159 * (-2.25 + -2.25)))
    _ = len(v4)
    v47 := GenS3{F0: (v5 != 0), F1: 100}
    var v48 float32 = v43
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
