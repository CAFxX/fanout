package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64; F1 string; F2 int8 }
type GenS6 struct { F0 uint16; F1 uint; F2 int64 }
func genFn0() uint16 {
    v4 := GenS1{F0: 1}
    switch -1 {
    case 604:
        switch "a" {
        case "a":
            var v5 int8 = -1
            _ = v5
        case "\n\t":
            var v6 int = 4294967295
            _ = v6
        }
    case 864:
        for fi7 := 0; fi7 < 15; fi7++ {
            fmt.Println(fi7)
            var v8 int32 = 1
            var v9 bool = (656 > 351)
            _ = fi7
            _ = v8
            _ = v9
        }
    default:
        var v10 uint = 743
        _ = v10
    }
    var v11 uint64 = 42
    _ = v11
    _ = v4
    return uint16(v11)
}
func genFn1(p0 int) int8 {
    v12 := []int{}
    var v13 bool = (1.5 == float32(((float64(((float32(0.1) / (3.14159 / 1.5)) * float32((3.14159 + float64(1.5))))) - 0.0) / (float64((((0.0 * ((0.1 * 0.0) - 1e10)) / 1e10) - 1e10)) / 0.1))))
    var v14 uint8 = 96
    for fi15 := 0; fi15 < 14; fi15++ {
        fmt.Println(p0)
        switch (("a" + ("foo bar" + ("hello" + "foo bar"))) + "hello") {
        case "":
            switch (-(fi15)) {
            case 128:
                var v16 int16 = 32767
                _ = v16
            case 127:
                switch "foo bar" {
                case "":
                    v14 = v14
                case "\n\t":
                    for fi17 := 0; fi17 < 4; fi17++ {
                        if (!v13) {
                            v13 = (-9223372036854775808 != 1)
                            if (-2147483648 > int32(v14)) {
                                fmt.Println(v13, v13)
                            }
                        } else {
                            fmt.Println(v14, fi17, v14)
                            p0 += (fi17 << ((1) & 63))
                        }
                        if len(v12) > 0 {
                            v12[0] = (p0 << ((4294967295) & 63))
                        }
                        var v18 int16 = 42
                        _ = fi17
                        _ = v18
                    }
                default:
                    if v13 {
                        fmt.Println(v13)
                    }
                }
            }
        }
        _ = fi15
    }
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return 42
}

func main() {
    v1 := []uint32{265, 65535, 111}
    var v2 uint8 = 1
    var v3 float32 = (float32(((3.14159 - float64(((-2.25 + (0.1 - 1e10)) - -2.25))) * (-2.25 - ((-2.25 + 0.1) - 3.14159)))) + 3.14159)
    if (((^(v2)) == v2) && false) {
        switch (("hello" + "foo bar") + (("a" + "") + (("\n\t" + ((("\n\t" + ("\n\t" + (("foo bar" + ((("a" + "") + ("foo bar" + "a")) + "a")) + ("\n\t" + "hello")))) + "a") + "a")) + ("foo bar" + "hello")))) {
        case "foo bar":
            if (!(0 == 1)) {
                switch "hello" {
                case "":
                    var v19 uint = 115
                    _ = v19
                default:
                    if (4294967295 < 867) {
                        v20 := GenS2{F0: ("" + "foo bar"), F1: 4294967295}
                        switch "foo bar" {
                        case "foo bar":
                            if len(v1) > 0 {
                                v1[0] = 42
                            }
                        case "hello":
                            for fi21 := 0; fi21 < 8; fi21++ {
                                v1 = append(v1, 0)
                                v3 -= 0.0
                                v3 = -2.25
                                _ = fi21
                            }
                        default:
                            v1 = append(v1, 2147483648)
                        }
                        _ = v20
                    } else {
                        switch ("\n\t" + ("a" + ("foo bar" + "a"))) {
                        case "a":
                            _ = len(v1)
                        case "\n\t":
                            var v22 uint8 = v2
                            _ = v22
                        case "hello":
                            v2 = v2
                        default:
                            switch "a" {
                            case "\n\t":
                                v2 = (v2 & (^(v2)))
                            case "":
                                v2 = v2
                            default:
                                _ = len(v1)
                            }
                        }
                        v23 := GenS2{F0: "a", F1: 0}
                        _ = v23
                    }
                }
            } else {
                if false {
                    var v24 uint16 = genFn0()
                    _ = v24
                }
                v2 -= v2
            }
        case "a":
            fmt.Println(v2, v2)
        default:
            switch "" {
            case "foo bar":
                v2 += 47
            }
        }
        v3 -= (0.0 - v3)
    } else {
        var v25 int64 = -9223372036854775808
        switch 1 {
        case -1:
            var v26 uint32 = 2147483648
            _ = v26
        case -792:
            v2 += (v2 & (v2 + 0))
        default:
            var v27 uint = 18446744073709551615
            _ = v27
        }
        _ = v25
    }
    v3 = -2.25
    var v28 float64 = (0.0 / float64(v3))
    v1 = append(v1, 708)
    var v29 int = -333
    var v30 int64 = -1
    v31 := GenS3{F0: (v29 - v29), F1: 354}
    if (234 >= 343) {
        var v32 uint64 = 866
        v33 := GenS1{F0: v32}
        _ = v32
        _ = v33
    }
    if true {
        var v34 int32 = -600
        _ = v34
    }
    v1 = append(v1, 65535)
    _ = len(v1)
    var v35 bool = (805 > 1)
    switch (((("foo bar" + "\n\t") + (("" + ("foo bar" + ("hello" + ""))) + "hello")) + "\n\t") + ("" + ("a" + "foo bar"))) {
    case "\n\t":
        v29 = (v29 ^ 0)
    }
    if ((v30 & (^(v30))) != v30) {
        v35 = (!(v35 || v35))
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v2 = 0
    } else {
        switch (v29 << ((42) & 63)) {
        case -980:
            var v36 uint8 = v2
            _ = v36
        case -337:
            v2 += v2
        case 128:
            v28 = 0.0
        default:
            var v37 float32 = 1.5
            _ = v37
        }
        v30 -= (v30 - (v30 / ((v30) | 1)))
        v38 := GenS4{F0: -1}
        _ = v38
    }
    for fi39 := 0; fi39 < 16; fi39++ {
        switch ("" + "foo bar") {
        case "foo bar":
            v29 += (v29 & 9223372036854775807)
        case "hello":
            v40 := GenS5{F0: 810, F1: (((("\n\t" + ("hello" + "a")) + (((("\n\t" + "") + "hello") + ("foo bar" + ("hello" + "\n\t"))) + "foo bar")) + ("foo bar" + "\n\t")) + ("" + "hello")), F2: 1}
            _ = v40
        case "a":
            v35 = (32767 == 1000)
        }
        _ = fi39
    }
    for fi41 := 0; fi41 < 1; fi41++ {
        for fi42 := 0; fi42 < 1; fi42++ {
            if (!(v3 == v3)) {
                var v43 string = "\n\t"
                if (!false) {
                    switch "a" {
                    case "":
                        v2 = (v2 >> ((373) & 63))
                        var v44 float32 = float32((((3.14159 + (v28 / -2.25)) - v28) + ((v28 / v28) - 1.5)))
                        v45 := GenS4{F0: 127}
                        _ = v44
                        _ = v45
                    case "hello":
                        fmt.Println(v29)
                    case "foo bar":
                        v3 += v3
                    default:
                        v1 = append(v1, 2)
                    }
                } else {
                    v1 = append(v1, 686)
                    switch fi42 {
                    case 1:
                        v29 -= (fi42 % (((fi41 + v29)) | 1))
                    default:
                        _ = len(v1)
                    }
                    fmt.Println(v30)
                    fmt.Println(v28)
                }
                var v46 int16 = 1
                _ = v43
                _ = v46
            } else {
                var v47 uint64 = 510
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v47
            }
            switch (fi42 - (fi41 * (v29 / ((v29) | 1)))) {
            case 42:
                v48 := GenS3{F0: (v29 - v29), F1: 362}
                _ = v48
            case 32768:
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            case 32767:
                fmt.Println(fi42, fi42, fi42)
            default:
                switch fi41 {
                case 127:
                    var v49 int64 = v30
                    _ = v49
                default:
                    if v35 {
                        fmt.Println(fi41)
                        v50 := GenS4{F0: genFn1((v29 ^ int(v30)))}
                        _ = v50
                    } else {
                        v51 := GenS6{F0: 255, F1: 65535, F2: v30}
                        fmt.Println(fi42, v30)
                        _ = v51
                    }
                }
            }
            v35 = (v35 && v35)
            _ = fi42
        }
        _ = fi41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
}
