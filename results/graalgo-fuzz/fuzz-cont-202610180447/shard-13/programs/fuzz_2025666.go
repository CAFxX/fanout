package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 int32; F1 int16; F2 int64 }
type GenS6 struct { F0 int; F1 int32 }
type GenS7 struct { F0 int8 }
func genFn0(p0 uint8, p1 string, p2 uint8) uint8 {
    switch p1 {
    case "":
        p0 -= 128
    case "\n\t":
        if true {
            var v4 int64 = -909
            v5 := GenS1{F0: (p2 ^ (p2 >> ((255) & 63))), F1: 138}
            _ = v4
            _ = v5
        }
    case "a":
        switch 813 {
        case 32767:
            v6 := GenS1{F0: p0, F1: p0}
            _ = v6
        default:
            if ((-9223372036854775808 > 1) && (p2 <= p0)) {
                v7 := GenS2{F0: 2147483647}
                _ = v7
            } else {
                for fi8 := 0; fi8 < 12; fi8++ {
                    p1 = p1
                    p0 = (p0 * p2)
                    v9 := []int{fi8, (fi8 | fi8), (fi8 - fi8), -2147483648}
                    _ = fi8
                    _ = v9
                }
                p2 = (^(p2))
                switch p1 {
                case "foo bar":
                    for fi10 := 0; fi10 < 5; fi10++ {
                        var v11 uint8 = (^(p0))
                        for fi12 := 0; fi12 < 16; fi12++ {
                            fmt.Println(fi10, fi12)
                            _ = fi12
                        }
                        p0 = v11
                        _ = fi10
                        _ = v11
                    }
                case "\n\t":
                    v13 := GenS1{F0: p2, F1: (p0 + p0)}
                    _ = v13
                case "a":
                    for fi14 := 0; fi14 < 5; fi14++ {
                        if false {
                            if (true && ((!(2 <= 2147483647)) && ((2147483647 < 1) || (1 >= 18446744073709551615)))) {
                                fmt.Println(p0)
                                if (42 == 993) {
                                    p0 += (^(p0))
                                }
                                v15 := []int{(fi14 >> ((285) & 63))}
                                _ = v15
                            } else {
                                var v16 float32 = float32(0.0)
                                _ = v16
                            }
                            v17 := []int{(fi14 & (-(fi14))), (fi14 | fi14)}
                            fmt.Println(p2)
                            _ = v17
                        } else {
                            v18 := GenS3{F0: 894}
                            _ = v18
                        }
                        _ = fi14
                    }
                default:
                    fmt.Println(p2, p1)
                }
            }
        }
    }
    v19 := []int{9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 0
}
func genFn1(p0 int) uint64 {
    switch ("hello" + "") {
    case "\n\t":
        var v20 uint32 = 2
        _ = v20
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        v22 := GenS3{F0: 65535}
        p0 += (p0 ^ (fi21 >> ((42) & 63)))
        if false {
            for fi23 := 0; fi23 < 11; fi23++ {
                fmt.Println(p0, fi23)
                v24 := GenS4{F0: false, F1: 42}
                p0 -= (fi21 >> ((65535) & 63))
                fmt.Println(fi23, fi21)
                _ = fi23
                _ = v24
            }
            p0 -= (fi21 >> ((18446744073709551615) & 63))
            v25 := []int{fi21, 9223372036854775807, p0}
            _ = v25
        }
        _ = fi21
        _ = v22
    }
    _ = p0
    return 0
}

func main() {
    var v1 int16 = 1
    var v2 uint = 946
    v3 := []int{664, 9223372036854775807, 4294967295}
    var v26 uint = (^(v2))
    fmt.Println(v2, v26, v1)
    var v27 int16 = (v1 / (((v1 ^ v1)) | 1))
    v26 = v26
    if (uint16(v2) > 1) {
        if false {
            v1 = v1
            var v28 int64 = -1
            switch "\n\t" {
            case "foo bar":
                v29 := GenS5{F0: 1, F1: 0, F2: (v28 / (((v28 - (v28 + (v28 * (v28 % ((v28) | 1)))))) | 1))}
                _ = v29
            }
            _ = v28
        } else {
            for fi30 := 0; fi30 < 18; fi30++ {
                v3 = append(v3, 9223372036854775807)
                _ = fi30
            }
            for fi31 := 0; fi31 < 18; fi31++ {
                if true {
                    v2 += 0
                    if len(v3) > 0 {
                        v3[0] = fi31
                    }
                }
                v3 = append(v3, fi31)
                _ = fi31
            }
            if (-2.25 > float64(-2.25)) {
                var v32 uint = (^(v26))
                _ = v32
            }
        }
    } else {
        fmt.Println(v2)
        v26 += (v26 & (v2 % (((v26 >> ((v2) & 63))) | 1)))
        var v33 uint = (v2 - 4294967295)
        _ = v33
    }
    var v34 int = 32768
    var v35 int8 = 1
    if true {
        fmt.Println(v26, v1, v35)
    } else {
        if (!false) {
            v26 = 1
            var v36 float64 = 1e10
            _ = v36
        }
        v1 = (v27 / (((v27 + v1)) | 1))
    }
    var v37 bool = true
    v37 = v37
    if true {
        for fi38 := 0; fi38 < 1; fi38++ {
            fmt.Println(v34, v2, v34)
            v34 -= 750
            _ = fi38
        }
    } else {
        v1 = (v1 - 976)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v35 = v35
    }
    fmt.Println(v37)
    for fi39 := 0; fi39 < 19; fi39++ {
        switch int(v35) {
        case 42:
            v40 := GenS6{F0: (fi39 - fi39), F1: -1}
            _ = v40
        case 32767:
            var v41 bool = v37
            _ = v41
        case 128:
            if (1.5 == 1.5) {
                var v42 string = ""
                v43 := GenS7{F0: (-(v35))}
                _ = v42
                _ = v43
            }
        default:
            switch fi39 {
            case 1:
                v26 = (v26 % (((v2 + 2)) | 1))
            }
        }
        for fi44 := 0; fi44 < 13; fi44++ {
            var v45 float32 = 0.0
            v45 -= -2.25
            v37 = true
            _ = fi44
            _ = v45
        }
        _ = fi39
    }
    switch "" {
    case "":
        var v46 bool = v37
        _ = v46
    }
    switch "a" {
    case "a":
        v47 := GenS2{F0: v34}
        _ = v47
    case "hello":
        if ((2 == 678) && (127 <= genFn0(231, (("foo bar" + "hello") + "\n\t"), 42))) {
            _ = len(v3)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v34 += int(v26)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v37
}
