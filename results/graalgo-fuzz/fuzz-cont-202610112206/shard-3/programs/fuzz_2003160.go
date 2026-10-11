package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int }
type GenS2 struct { F0 bool; F1 int32; F2 bool }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int32; F1 uint16; F2 string }
func genFn0() int {
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            switch "foo bar" {
            case "foo bar":
                switch fi7 {
                case -1:
                    var v8 float32 = 3.14159
                    _ = v8
                case 127:
                    var v9 float64 = 0.1
                    _ = v9
                }
            case "\n\t":
                var v10 int = (^(fi7))
                _ = v10
            default:
                v11 := GenS1{F0: -1, F1: 0, F2: (fi6 ^ 127)}
                _ = v11
            }
            var v12 int16 = 1
            switch fi7 {
            case 0:
                v13 := GenS1{F0: (v12 & (v12 & (v12 >> ((346) & 63)))), F1: 0, F2: fi6}
                var v14 int8 = -128
                switch fi6 {
                case -39:
                    fmt.Println(fi7)
                default:
                    v12 += (v12 + (v12 + (v12 << ((42) & 63))))
                    v15 := GenS2{F0: false, F1: -913, F2: (42 == 406)}
                    if ((!(!false)) || ((true && (false || (!(0 < 883)))) || true)) {
                        v16 := []int{1, 0}
                        if (!((!("" == "hello")) && false)) {
                            _ = len(v16)
                            v17 := GenS3{F0: "hello", F1: 0}
                            _ = v17
                        }
                        v18 := GenS2{F0: false, F1: 664, F2: (true || (((fi7 >> ((911) & 63)) < -9223372036854775808) || true))}
                        _ = v16
                        _ = v18
                    }
                    _ = v15
                }
                _ = v13
                _ = v14
            case 32767:
                var v19 float32 = -2.25
                _ = v19
            case -1:
                if (true || false) {
                    v20 := GenS3{F0: "hello", F1: 65535}
                    _ = v20
                } else {
                    var v21 int16 = (-(v12))
                    v22 := GenS2{F0: false, F1: 294, F2: (42 > 65535)}
                    _ = v21
                    _ = v22
                }
            default:
                v23 := []int{(^(fi7)), (fi6 >> ((42) & 63)), (fi6 >> ((1) & 63)), fi7}
                _ = v23
            }
            if (489 == 0) {
                fmt.Println(fi6, v12, fi7)
            } else {
                var v24 uint64 = 4294967295
                switch "\n\t" {
                case "a":
                    v12 = v12
                case "":
                    v24 -= 325
                case "foo bar":
                    var v25 int64 = 42
                    _ = v25
                default:
                    v26 := []int{(-(fi6)), fi6}
                    _ = v26
                }
                switch (fi7 ^ -254) {
                case -1:
                    fmt.Println(fi7, fi6)
                case 335:
                    var v27 uint32 = 544
                    _ = v27
                case 32768:
                    v28 := []int{(^(fi7))}
                    _ = v28
                default:
                    if true {
                        var v29 uint = 2
                        v12 = (v12 % ((v12) | 1))
                        switch (fi7 & 322) {
                        case 310:
                            var v30 uint8 = 2
                            _ = v30
                        case 1:
                            var v31 uint64 = (v24 + 0)
                            _ = v31
                        case 2147483647:
                            var v32 bool = false
                            _ = v32
                        }
                        fmt.Println(v12, v12, fi7)
                        v29 = 2
                        _ = v29
                    }
                }
                _ = v24
            }
            switch (-(fi7)) {
            case 615:
                v12 += -739
            default:
                var v33 uint8 = 2
                _ = v33
            }
            _ = fi7
            _ = v12
        }
        v34 := GenS3{F0: "a", F1: 1}
        _ = fi6
        _ = v34
    }
    var v35 uint64 = 188
    if ((1e10 / (0.1 / 1e10)) != -2.25) {
        v35 = (v35 ^ (v35 >> ((42) & 63)))
        var v36 int32 = 1
        _ = v36
    }
    if true {
        switch int(v35) {
        case 9223372036854775807:
            v37 := GenS3{F0: "\n\t", F1: 65535}
            _ = v37
        default:
            v38 := []int{}
            _ = v38
        }
    } else {
        if false {
            var v39 uint32 = 660
            v40 := []int{0, 32767, -2147483648}
            v39 = v39
            _ = v39
            _ = v40
        } else {
            v41 := []int{0, 127, 32768, 534}
            v41 = append(v41, 1)
            _ = v41
        }
    }
    _ = v35
    return -393
}

func main() {
    v1 := []string{("foo bar" + "a"), ("a" + "foo bar"), ("a" + ""), ("foo bar" + ("foo bar" + (("hello" + "") + "")))}
    v2 := make(map[string]uint32)
    var v3 int64 = -1
    v4 := []string{"foo bar", "", "a"}
    v5 := make(map[uint32]uint32)
    if (false || (0 != 165)) {
        fmt.Println(v3)
        switch "hello" {
        case "a":
            for fi42 := 0; fi42 < 15; fi42++ {
                var v43 uint = 4294967295
                for fi44 := 0; fi44 < 12; fi44++ {
                    if true {
                        v45 := GenS4{F0: -2147483648, F1: 255, F2: "hello"}
                        _ = v45
                    }
                    _ = fi44
                }
                if true {
                    v43 = (v43 * 729)
                    v3 = v3
                    switch fi42 {
                    case 2147483647:
                        fmt.Println(fi42)
                    }
                }
                _ = fi42
                _ = v43
            }
        }
    } else {
        v3 = (^(v3))
        fmt.Println(v3)
        if false {
            _ = len(v1)
        }
    }
    fmt.Println(v3)
    v46 := GenS2{F0: false, F1: 42, F2: (0 < 0)}
    var v47 uint64 = 275
    var v48 float64 = 1e10
    delete(v2, "hello")
    var v49 int32 = -246
    switch 9223372036854775807 {
    case 128:
        var v50 float32 = 0.0
        _ = v50
    case 127:
        v49 -= 1
    case 1:
        var v51 int = genFn0()
        _ = v51
    default:
        delete(v5, 1)
    }
    for fi52 := 0; fi52 < 14; fi52++ {
        for fi53 := 0; fi53 < 14; fi53++ {
            var v54 uint8 = 255
            _ = fi53
            _ = v54
        }
        if true {
            for fi55 := 0; fi55 < 12; fi55++ {
                switch genFn0() {
                case 815:
                    switch ("" + "\n\t") {
                    case "a":
                        if (v47 < v47) {
                            fmt.Println(v48, fi52)
                            var v56 bool = false
                            _ = v56
                        }
                    }
                case 2147483647:
                    switch (fi52 | (-(fi55))) {
                    case 32767:
                        v57 := v2["a"]
                        _ = v57
                    }
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                var v58 string = "foo bar"
                _ = fi55
                _ = v58
            }
        }
        _ = fi52
    }
    for fi59 := 0; fi59 < 10; fi59++ {
        var v60 bool = true
        v47 = 4294967295
        _ = fi59
        _ = v60
    }
    for fi61 := 0; fi61 < 5; fi61++ {
        if (!false) {
            _ = len(v1)
            v62 := GenS4{F0: (v49 - v49), F1: 1, F2: "\n\t"}
            v63 := GenS3{F0: "", F1: 1}
            _ = v62
            _ = v63
        }
        _ = fi61
    }
    if (168 == 18446744073709551615) {
        v48 -= v48
        v3 = 9223372036854775807
        v64 := GenS4{F0: 2147483647, F1: 1, F2: ("a" + ("" + "\n\t"))}
        _ = v64
    } else {
        for fi65 := 0; fi65 < 10; fi65++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v47 = v47
            _ = fi65
        }
        if len(v4) > 0 {
            v4[0] = ""
        }
        var v66 int16 = 0
        _ = v66
    }
    v48 -= 0.1
    var v67 float64 = ((((v48 * (v48 * v48)) + (1e10 * v48)) - v48) / float64(1e10))
    var v68 int64 = 1
    var v69 float32 = float32(v67)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v67
    _ = v68
    _ = v69
}
