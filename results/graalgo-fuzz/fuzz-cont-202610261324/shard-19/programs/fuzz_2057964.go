package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8 }
func genFn0() int {
    v7 := []int{-2147483648, -2147483648, 32767, 32768}
    switch "hello" {
    case "a":
        if (!(1.5 >= 0.1)) {
            var v8 int64 = 1
            for fi9 := 0; fi9 < 5; fi9++ {
                for fi10 := 0; fi10 < 8; fi10++ {
                    switch fi10 {
                    case -1:
                        fmt.Println(fi9, fi9)
                    case 127:
                        v8 = v8
                    case 32768:
                        v8 -= v8
                    }
                    fmt.Println(v8, fi10, v8)
                    _ = fi10
                }
                _ = fi9
            }
            _ = v8
        }
    case "":
        var v11 uint32 = 234
        _ = v11
    default:
        if len(v7) > 0 {
            v7[0] = 87
        }
    }
    _ = v7
    return 4294967295
}
func genFn1() string {
    if true {
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 string = "foo bar"
            v14 := GenS1{F0: 2}
            if (false || false) {
                var v15 int8 = 42
                var v16 int16 = -117
                switch fi12 {
                case 425:
                    var v17 int16 = 32767
                    _ = v17
                case -673:
                    v18 := GenS1{F0: 255}
                    _ = v18
                }
                _ = v15
                _ = v16
            } else {
                v13 = v13
                var v19 string = "\n\t"
                v13 = v19
                _ = v19
            }
            _ = fi12
            _ = v13
            _ = v14
        }
    } else {
        var v20 int = 127
        for fi21 := 0; fi21 < 15; fi21++ {
            for fi22 := 0; fi22 < 11; fi22++ {
                var v23 int = fi21
                var v24 uint16 = 255
                switch ("" + ((("" + ("hello" + ("foo bar" + ""))) + "") + "\n\t")) {
                case "":
                    if false {
                        v25 := []int{v20, 32767, fi22}
                        if len(v25) >= 2 {
                            v25 = v25[1:len(v25)]
                        }
                        switch (((("" + "\n\t") + "\n\t") + "") + "a") {
                        case "foo bar":
                            fmt.Println(fi22, v24, v23)
                        default:
                            v26 := GenS1{F0: 0}
                            _ = v26
                        }
                        _ = v25
                    }
                case "\n\t":
                    v23 -= (v23 << ((255) & 63))
                case "foo bar":
                    var v27 float32 = -2.25
                    _ = v27
                default:
                    switch (v23 >> ((2) & 63)) {
                    case 128:
                        fmt.Println(fi22, v20, fi22)
                    }
                }
                if false {
                    var v28 int32 = 225
                    if false {
                        v24 += 674
                    } else {
                        v23 = fi21
                    }
                    _ = v28
                }
                v29 := GenS2{F0: 42, F1: ((((("" + "foo bar") + (("a" + "hello") + ("a" + ""))) + "a") <= "a") && (uint8(v24) > 127))}
                v30 := []int{fi21, (v20 & (^(fi22)))}
                fmt.Println(v23, fi22, v23)
                v20 -= 32768
                fmt.Println(fi22, fi22, fi21)
                var v31 int64 = -408
                _ = fi22
                _ = v23
                _ = v24
                _ = v29
                _ = v30
                _ = v31
            }
            fmt.Println(fi21, v20)
            _ = fi21
        }
        if false {
            v32 := []int{2147483647, v20, v20, 1}
            switch genFn0() {
            case -1:
                for fi33 := 0; fi33 < 14; fi33++ {
                    if false {
                        switch (("hello" + ("" + ("a" + "\n\t"))) + ("foo bar" + "\n\t")) {
                        case "\n\t":
                            if (false || true) {
                                var v34 uint = 255
                                fmt.Println(v20, fi33)
                                switch (v20 & (-(fi33))) {
                                case -940:
                                    for fi35 := 0; fi35 < 9; fi35++ {
                                        v36 := GenS2{F0: 1, F1: ((1 > 0) && (!false))}
                                        v34 -= (v34 % (((v34 & (v34 ^ (^(v34))))) | 1))
                                        var v37 bool = (1 <= 42)
                                        _ = fi35
                                        _ = v36
                                        _ = v37
                                    }
                                case -9223372036854775808:
                                    if len(v32) >= 2 {
                                        v32 = v32[1:len(v32)]
                                    }
                                case 1:
                                    for fi38 := 0; fi38 < 9; fi38++ {
                                        v32 = append(v32, fi33)
                                        var v39 uint64 = 4294967295
                                        _ = fi38
                                        _ = v39
                                    }
                                default:
                                    v34 = 42
                                }
                                _ = v34
                            }
                        case "a":
                            for fi40 := 0; fi40 < 14; fi40++ {
                                fmt.Println(fi33)
                                var v41 string = ("hello" + ((("" + "a") + ("\n\t" + "a")) + "foo bar"))
                                _ = fi40
                                _ = v41
                            }
                        case "hello":
                            switch "a" {
                            case "foo bar":
                                if false {
                                    switch 128 {
                                    case 904:
                                        v20 -= (v20 ^ -688)
                                    default:
                                        var v42 float32 = (1e10 + 1.5)
                                        _ = v42
                                    }
                                }
                            case "\n\t":
                                if false {
                                    var v43 int8 = 100
                                    if ((!(true || false)) && false) {
                                        var v44 float64 = 3.14159
                                        var v45 float64 = v44
                                        v46 := GenS3{F0: 4294967295}
                                        _ = v44
                                        _ = v45
                                        _ = v46
                                    }
                                    v20 = (fi33 % ((int(v43)) | 1))
                                    _ = v43
                                }
                            case "hello":
                                var v47 int = v20
                                _ = v47
                            }
                        }
                    }
                    for fi48 := 0; fi48 < 4; fi48++ {
                        if true {
                            var v49 string = "a"
                            _ = v49
                        }
                        v20 -= fi33
                        var v50 uint8 = 255
                        v32 = append(v32, genFn0())
                        _ = fi48
                        _ = v50
                    }
                    _ = fi33
                }
            case 127:
                if true {
                    switch (v20 / ((v20) | 1)) {
                    case 42:
                        var v51 int = genFn0()
                        _ = v51
                    case 32767:
                        v20 = v20
                    default:
                        switch ("" + "hello") {
                        case "foo bar":
                            v52 := GenS2{F0: 42, F1: true}
                            _ = v52
                        default:
                            var v53 uint32 = 290
                            _ = v53
                        }
                    }
                    if (2 < 0) {
                        for fi54 := 0; fi54 < 2; fi54++ {
                            if len(v32) > 0 {
                                v32[0] = (fi54 >> ((18446744073709551615) & 63))
                            }
                            _ = fi54
                        }
                        v20 = (v20 | v20)
                    } else {
                        fmt.Println(v20)
                    }
                }
            case -536:
                v20 = (v20 / ((v20) | 1))
            default:
                var v55 uint8 = 181
                _ = v55
            }
            _ = v32
        }
        _ = v20
    }
    v56 := []int{}
    _ = v56
    return ("" + "foo bar")
}

func main() {
    var v1 int8 = 0
    v2 := []string{(("a" + "a") + ("\n\t" + "hello")), ((("" + "foo bar") + "") + "hello")}
    var v3 bool = true
    v4 := []int{-934, 127, 1, 128}
    v5 := make(map[string]uint32)
    v6 := []int{712}
    v3 = true
    var v57 uint32 = 2147483648
    v4 = append(v4, 32767)
    var v58 uint = 62
    v59 := GenS3{F0: v57}
    v60 := GenS4{F0: 1}
    for fi61 := 0; fi61 < 14; fi61++ {
        v62 := GenS1{F0: 18446744073709551615}
        fmt.Println(v1, fi61, v1)
        _ = fi61
        _ = v62
    }
    for fi63 := 0; fi63 < 11; fi63++ {
        var v64 float32 = 3.14159
        _ = fi63
        _ = v64
    }
    fmt.Println(v1, v58, v3)
    switch 0 {
    case -462:
        switch ("foo bar" + "") {
        case "foo bar":
            v65 := GenS3{F0: v57}
            _ = v65
        }
    }
    v3 = (v3 && (((0.1 - 0.1) - (1.5 * float64(((-2.25 / 1.0) / 3.14159)))) != 0.0))
    var v66 uint64 = 1
    for fi67 := 0; fi67 < 8; fi67++ {
        var v68 float64 = -2.25
        var v69 uint8 = 0
        _ = fi67
        _ = v68
        _ = v69
    }
    if (v57 <= (v57 << ((v58) & 63))) {
        if v3 {
            fmt.Println(v3)
            fmt.Println(v66, v3)
            fmt.Println(v66, v57)
        } else {
            fmt.Println(v3, v57)
            switch 837 {
            case 9223372036854775807:
                v57 = (v57 + (v57 ^ 2147483647))
            default:
                fmt.Println(v66)
            }
            v2 = append(v2, "a")
        }
    } else {
        var v70 uint16 = 1
        v71 := v5["a"]
        var v72 int8 = (v1 - (v1 % ((37) | 1)))
        _ = v70
        _ = v71
        _ = v72
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v66
}
