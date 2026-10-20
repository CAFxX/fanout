package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint }
type GenS2 struct { F0 int; F1 uint16; F2 int32 }
type GenS3 struct { F0 int16; F1 int; F2 uint64 }
type GenS4 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS5 struct { F0 int16; F1 int64; F2 int }
type GenS6 struct { F0 int; F1 int16; F2 int16 }
type GenS7 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS8 struct { F0 int64; F1 int64; F2 uint }
type GenS9 struct { F0 bool; F1 string }
type GenS10 struct { F0 uint64 }
func genFn0() int {
    v6 := []int{-845}
    v7 := GenS1{F0: -128, F1: 2, F2: 65535}
    v8 := GenS2{F0: 1, F1: 32767, F2: 42}
    var v9 uint16 = 32767
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 int64 = 42
        switch fi10 {
        case 1:
            v11 += v11
        case -489:
            fmt.Println(v9, v11)
        case 9223372036854775807:
            for fi12 := 0; fi12 < 19; fi12++ {
                v13 := GenS2{F0: (^(fi10)), F1: (^(v9)), F2: -2147483648}
                _ = fi12
                _ = v13
            }
        }
        var v14 uint64 = 2
        _ = fi10
        _ = v11
        _ = v14
    }
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return -9223372036854775808
}
func genFn1(p0 uint8, p1 uint64, p2 uint) int32 {
    if (!(false && false)) {
        p0 = (p0 << (((p2 << ((18446744073709551615) & 63))) & 63))
        v15 := []int{-410}
        _ = v15
    } else {
        v16 := GenS2{F0: int(p2), F1: 724, F2: 1}
        v17 := GenS3{F0: int16(p0), F1: 128, F2: (p1 % (((p1 ^ p1)) | 1))}
        p1 += 719
        _ = v16
        _ = v17
    }
    var v18 int64 = -1
    var v19 int32 = 42
    var v20 float32 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    _ = v20
    return v19
}

func main() {
    var v1 int16 = -1
    var v2 int64 = 9223372036854775807
    v3 := []string{}
    var v4 uint = 65535
    var v5 int8 = 127
    v5 = (v5 << ((v4) & 63))
    switch 0 {
    case -1:
        fmt.Println(v1, v5)
    case 0:
        v21 := GenS2{F0: -464, F1: 241, F2: -2147483648}
        _ = v21
    case -9223372036854775808:
        v22 := GenS2{F0: 0, F1: 65535, F2: -518}
        _ = v22
    default:
        _ = len(v3)
    }
    v2 -= v2
    v2 += v2
    v23 := GenS2{F0: -304, F1: 2, F2: -1}
    switch "foo bar" {
    case "":
        var v24 float64 = (0.0 - float64((1.5 + ((0.1 + 1.5) * (float32((3.14159 / 1.0)) * ((float32((float64((1.5 / (((((0.0 / 3.14159) / (3.14159 / 1.0)) + (((float32((0.1 / (-2.25 / 1e10))) * 1e10) + -2.25) - -2.25)) * 3.14159) + 3.14159))) + 1.5)) - ((1e10 * (0.1 - 3.14159)) + 1e10)) / (0.0 + -2.25)))))))
        _ = v24
    default:
        var v25 uint16 = 32767
        _ = v25
    }
    v2 = (v2 / ((v2) | 1))
    v26 := GenS1{F0: 1, F1: 0, F2: v4}
    if (32768 < 127) {
        fmt.Println(v2, v1)
        v4 += v4
        if false {
            v27 := GenS4{F0: 857, F1: 409, F2: 275}
            _ = v27
        }
    }
    v1 = (v1 | 42)
    v1 += v1
    if (!false) {
        if true {
            if (2 == 42) {
                v4 -= (^(v4))
                if (false || false) {
                    v28 := GenS5{F0: (v1 % ((1000) | 1)), F1: -546, F2: -9223372036854775808}
                    var v29 int16 = int16(v2)
                    _ = v28
                    _ = v29
                } else {
                    var v30 uint16 = 1
                    _ = v30
                }
            }
            v1 -= (v1 - -1)
            var v31 uint = (v4 << ((519) & 63))
            _ = v31
        } else {
            for fi32 := 0; fi32 < 11; fi32++ {
                var v33 uint64 = 42
                _ = fi32
                _ = v33
            }
        }
        switch "a" {
        case "":
            var v34 int64 = v2
            _ = v34
        case "\n\t":
            for fi35 := 0; fi35 < 5; fi35++ {
                var v36 uint64 = 519
                _ = fi35
                _ = v36
            }
        default:
            if true {
                switch ("foo bar" + "hello") {
                case "hello":
                    var v37 string = "hello"
                    _ = v37
                case "a":
                    fmt.Println(v5, v1)
                default:
                    fmt.Println(v1, v2)
                }
            } else {
                if (!true) {
                    v2 += v2
                    v1 = (v1 ^ int16(v5))
                }
            }
        }
        v4 += (v4 & 0)
    } else {
        switch 42 {
        case 1:
            if len(v3) > 0 {
                v3[0] = (("" + (("\n\t" + "foo bar") + (("a" + "foo bar") + ("a" + (((("hello" + ("\n\t" + ((("" + "hello") + "a") + (("a" + ("\n\t" + "")) + "")))) + "foo bar") + "hello") + "\n\t"))))) + ("a" + "\n\t"))
            }
        case 127:
            for fi38 := 0; fi38 < 8; fi38++ {
                fmt.Println(v2, v4, v1)
                v3 = append(v3, "\n\t")
                _ = fi38
            }
        }
        v3 = append(v3, "")
    }
    if true {
        fmt.Println(v5)
        v2 = (-(v2))
        switch (("a" + "") + ("hello" + (("a" + "foo bar") + "hello"))) {
        case "hello":
            fmt.Println(v4, v2, v1)
        default:
            var v39 float32 = 1e10
            _ = v39
        }
    }
    v40 := GenS6{F0: 128, F1: int16(v4), F2: v1}
    v41 := GenS6{F0: 564, F1: v1, F2: (v1 << ((1) & 63))}
    for fi42 := 0; fi42 < 5; fi42++ {
        for fi43 := 0; fi43 < 7; fi43++ {
            if false {
                v5 = v5
                switch "" {
                case "hello":
                    switch (fi43 % (((fi43 % ((9223372036854775807) | 1))) | 1)) {
                    case 878:
                        if (((3.14159 - (0.0 * 1.5)) * ((1.5 / 0.1) / ((((1e10 * (0.1 - 1e10)) + (((float64(0.1) + 0.0) * (0.1 / 1e10)) / 1.0)) * 3.14159) * -2.25))) >= -2.25) {
                            var v44 bool = true
                            v45 := GenS7{F0: uint32(v4), F1: 42, F2: 65535}
                            _ = v44
                            _ = v45
                        }
                    case 4294967295:
                        if false {
                            fmt.Println(v5, fi43, v2)
                        } else {
                            v2 = v2
                        }
                    }
                case "\n\t":
                    v46 := GenS8{F0: (v2 - v2), F1: (v2 ^ (v2 + (v2 - (v2 - (v2 * (v2 / (((v2 - (v2 * v2))) | 1))))))), F2: 18446744073709551615}
                    _ = v46
                case "foo bar":
                    var v47 uint32 = 103
                    _ = v47
                }
            } else {
                v1 -= (v1 >> ((v4) & 63))
            }
            var v48 bool = (!false)
            _ = fi43
            _ = v48
        }
        switch 1 {
        case 32767:
            v4 -= v4
        case -9223372036854775808:
            _ = len(v3)
        case -2147483648:
            var v49 uint32 = 479
            _ = v49
        }
        if true {
            var v50 string = "foo bar"
            v51 := GenS9{F0: true, F1: (v50 + v50)}
            fmt.Println(v2, fi42)
            _ = v50
            _ = v51
        } else {
            v52 := GenS10{F0: 0}
            switch (^(fi42)) {
            case 1:
                _ = len(v3)
            case -9223372036854775808:
                switch "foo bar" {
                case "hello":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "":
                    _ = len(v3)
                default:
                    v5 -= v5
                }
            case -1:
                var v53 int64 = 1
                _ = v53
            }
            _ = v52
        }
        _ = fi42
    }
    if (!true) {
        var v54 int64 = v2
        _ = len(v3)
        v55 := GenS3{F0: (v1 + (v1 / ((1000) | 1))), F1: 109, F2: 24}
        _ = v54
        _ = v55
    } else {
        v3 = append(v3, "a")
        fmt.Println(v5, v2)
        if ((v4 / (((^(v4))) | 1)) < v4) {
            v56 := GenS4{F0: 42, F1: 426, F2: 4294967295}
            v2 = v2
            switch "foo bar" {
            case "":
                var v57 uint64 = 0
                _ = v57
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = v56
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
