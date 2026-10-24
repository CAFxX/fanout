package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 int8; F1 int; F2 int32 }
type GenS7 struct { F0 string; F1 uint64 }
func genFn0(p0 uint64) bool {
    p0 -= 223
    if false {
        p0 -= p0
    } else {
        var v5 int8 = -4
        v6 := GenS1{F0: 65535, F1: true}
        v7 := GenS2{F0: (42 > -2147483648)}
        _ = v5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        switch (fi8 / (((fi8 << ((65535) & 63))) | 1)) {
        case 4294967295:
            p0 = 833
        case -9223372036854775808:
            if true {
                v9 := []int{128, (fi8 & fi8), -1}
                p0 -= (p0 % ((2) | 1))
                if true {
                    switch fi8 {
                    case 9223372036854775807:
                        for fi10 := 0; fi10 < 5; fi10++ {
                            switch "a" {
                            case "":
                                p0 = p0
                            case "a":
                                p0 = p0
                            }
                            p0 += 187
                            _ = len(v9)
                            _ = fi10
                        }
                    case 0:
                        p0 -= 1
                    case 4294967295:
                        if false {
                            p0 = p0
                            switch fi8 {
                            case 1:
                                if false {
                                    var v11 uint16 = 42
                                    if false {
                                        v9 = append(v9, -7)
                                        v11 = (v11 - v11)
                                    } else {
                                        var v12 int = 128
                                        for fi13 := 0; fi13 < 2; fi13++ {
                                            var v14 uint32 = 596
                                            var v15 string = "a"
                                            if (!false) {
                                                var v16 int8 = -72
                                                v14 = v14
                                                _ = v16
                                            } else {
                                                if false {
                                                    if (2147483647 >= -698) {
                                                        fmt.Println(v15, v14)
                                                        if true {
                                                            v12 = fi13
                                                            if len(v9) > 0 {
                                                                v9[0] = fi13
                                                            }
                                                        } else {
                                                            var v17 bool = (!true)
                                                            v12 += (-(fi8))
                                                            _ = v17
                                                        }
                                                        v18 := GenS3{F0: 2, F1: 130}
                                                        _ = v18
                                                    }
                                                }
                                            }
                                            _ = fi13
                                            _ = v14
                                            _ = v15
                                        }
                                        _ = v12
                                    }
                                    _ = v11
                                }
                            }
                        } else {
                            _ = len(v9)
                            var v19 uint16 = 155
                            _ = v19
                        }
                    }
                    switch 9223372036854775807 {
                    case 32768:
                        p0 += p0
                    case 324:
                        switch (^(fi8)) {
                        case -65:
                            if len(v9) > 0 {
                                v9[0] = 503
                            }
                        }
                    case -2147483648:
                        var v20 uint32 = 1
                        _ = v20
                    }
                }
                _ = v9
            } else {
                for fi21 := 0; fi21 < 9; fi21++ {
                    v22 := []int{4294967295, (^(fi8)), -1, fi8}
                    if false {
                        _ = len(v22)
                        var v23 float64 = ((-2.25 / 1.5) + 1e10)
                        var v24 uint8 = 1
                        _ = v23
                        _ = v24
                    }
                    _ = fi21
                    _ = v22
                }
            }
        default:
            v25 := []int{-893, (fi8 - fi8), (-(fi8)), fi8}
            _ = v25
        }
        p0 += (p0 + (p0 * p0))
        _ = fi8
    }
    _ = p0
    return false
}
func genFn1(p0 string, p1 uint8) int8 {
    var v26 string = ((p0 + p0) + "hello")
    p1 += (p1 << ((18446744073709551615) & 63))
    v26 = ("hello" + p0)
    v27 := []int{-2147483648, -946, 1}
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    return 42
}

func main() {
    var v1 uint = 2
    var v2 uint32 = 42
    var v3 float32 = (0.1 - ((((((3.14159 * 1.5) / 1.0) * float32(((1e10 * (0.0 + 0.1)) + ((-2.25 - (float64(0.1) / 1e10)) * 0.1)))) * 0.1) - 1.5) + (-2.25 * (((float32(1e10) + (float32(float64(1.5)) / (3.14159 - 1.5))) + float32((1e10 + 1.5))) + (3.14159 / 1.0)))))
    v4 := []string{"hello", "hello"}
    if len(v4) > 0 {
        v4[0] = ""
    }
    for fi28 := 0; fi28 < 8; fi28++ {
        fmt.Println(v3)
        for fi29 := 0; fi29 < 9; fi29++ {
            switch "hello" {
            case "hello":
                if ((fi28 << ((v1) & 63)) > fi28) {
                    v4 = append(v4, "foo bar")
                }
            default:
                var v30 int64 = -1
                _ = v30
            }
            v1 = v1
            v1 = v1
            fmt.Println(fi28)
            _ = fi29
        }
        v3 = v3
        _ = fi28
    }
    v31 := GenS1{F0: 1, F1: true}
    switch 127 {
    case -250:
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
    default:
        v32 := GenS4{F0: 227, F1: uint16(v1), F2: -2147483648}
        _ = v32
    }
    _ = len(v4)
    v33 := GenS4{F0: 255, F1: 11, F2: -442}
    for fi34 := 0; fi34 < 1; fi34++ {
        if true {
            var v35 uint32 = 4294967295
            _ = v35
        } else {
            if false {
                v2 -= v2
                fmt.Println(fi34, v1)
            }
        }
        _ = fi34
    }
    for fi36 := 0; fi36 < 6; fi36++ {
        var v37 float64 = 0.0
        v38 := GenS4{F0: 1, F1: 32767, F2: 576}
        _ = fi36
        _ = v37
        _ = v38
    }
    if false {
        if false {
            v39 := GenS5{F0: 94, F1: 4294967295}
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
            _ = v39
        }
    }
    v3 += v3
    var v40 uint32 = v2
    switch "a" {
    case "":
        switch ("hello" + ("\n\t" + "foo bar")) {
        case "hello":
            var v41 int16 = -1
            _ = v41
        default:
            for fi42 := 0; fi42 < 5; fi42++ {
                v43 := GenS6{F0: genFn1("foo bar", 255), F1: (fi42 ^ 42), F2: 2147483647}
                _ = fi42
                _ = v43
            }
        }
    case "\n\t":
        for fi44 := 0; fi44 < 7; fi44++ {
            for fi45 := 0; fi45 < 19; fi45++ {
                v3 -= ((v3 * v3) * 3.14159)
                fmt.Println(v3)
                var v46 float32 = (1e10 * v3)
                _ = fi45
                _ = v46
            }
            _ = fi44
        }
    }
    for fi47 := 0; fi47 < 17; fi47++ {
        v48 := GenS7{F0: "\n\t", F1: 1}
        _ = fi47
        _ = v48
    }
    v49 := GenS1{F0: v40, F1: (!false)}
    v40 = (v40 | v40)
    for fi50 := 0; fi50 < 13; fi50++ {
        for fi51 := 0; fi51 < 6; fi51++ {
            fmt.Println(fi51, v1)
            _ = len(v4)
            _ = fi51
        }
        switch (fi50 * (fi50 * (fi50 & (fi50 % ((127) | 1))))) {
        case 32767:
            switch fi50 {
            case -9223372036854775808:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 0:
                v40 += (v2 & 146)
            }
        case 4294967295:
            if false {
                var v52 int32 = 309
                _ = v52
            } else {
                for fi53 := 0; fi53 < 14; fi53++ {
                    var v54 uint64 = uint64(v2)
                    _ = fi53
                    _ = v54
                }
                fmt.Println(v40)
                _ = len(v4)
            }
        default:
            v2 = 2147483648
        }
        v1 = v1
        _ = fi50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v4
    _ = v40
    _ = v49
}
