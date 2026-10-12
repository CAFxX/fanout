package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8; F1 uint64; F2 string }
type GenS7 struct { F0 string; F1 string; F2 uint }
type GenS8 struct { F0 uint; F1 uint32 }
func genFn0() int {
    for fi4 := 0; fi4 < 12; fi4++ {
        var v5 int16 = 635
        var v6 uint64 = 4294967295
        if true {
            switch (fi4 | int(v5)) {
            case -1:
                v7 := []int{}
                _ = v7
            case 42:
                v8 := []int{32767}
                _ = v8
            }
            fmt.Println(fi4, fi4, v5)
            for fi9 := 0; fi9 < 13; fi9++ {
                var v10 uint64 = 0
                if true {
                    if (0 != 255) {
                        var v11 uint = 765
                        switch 32768 {
                        case 127:
                            switch fi9 {
                            case 2147483647:
                                var v12 uint = v11
                                _ = v12
                            case 1:
                                switch "foo bar" {
                                case "a":
                                    if false {
                                        var v13 string = ("foo bar" + "foo bar")
                                        v14 := GenS1{F0: -565}
                                        v10 -= 4294967295
                                        var v15 uint64 = v6
                                        if (((-2.25 + 0.1) - (0.1 / (0.1 - ((1.5 + (float64(float32(0.0)) + (-2.25 / 3.14159))) * 1.5)))) != (1e10 + (-2.25 * -2.25))) {
                                            var v16 string = "hello"
                                            _ = v16
                                        }
                                        _ = v13
                                        _ = v14
                                        _ = v15
                                    }
                                case "foo bar":
                                    v17 := []int{-9223372036854775808}
                                    _ = v17
                                default:
                                    fmt.Println(v6, fi4)
                                }
                            case -9223372036854775808:
                                switch ("" + "hello") {
                                case "hello":
                                    v18 := []int{(fi9 * (fi4 & (^(fi9)))), 0, (fi9 + fi4)}
                                    _ = v18
                                default:
                                    var v19 uint64 = (v10 ^ (v10 >> ((v11) & 63)))
                                    _ = v19
                                }
                            }
                        case 1:
                            var v20 int32 = 2147483647
                            _ = v20
                        default:
                            v21 := GenS2{F0: (!(!(!((^(v11)) > 65535)))), F1: (v11 >> ((v11) & 63))}
                            _ = v21
                        }
                        _ = v11
                    }
                    if true {
                        var v22 float64 = -2.25
                        switch "hello" {
                        case "foo bar":
                            if false {
                                var v23 string = "\n\t"
                                _ = v23
                            } else {
                                var v24 float64 = v22
                                v22 -= ((0.0 + (1e10 * 1.5)) + 3.14159)
                                _ = v24
                            }
                        }
                        if (!(42 < 0)) {
                            if true {
                                v10 -= 102
                                var v25 int32 = -2147483648
                                _ = v25
                            } else {
                                v22 -= (-2.25 + 0.0)
                                var v26 float32 = 1e10
                                v27 := GenS3{F0: "hello"}
                                var v28 uint32 = 65535
                                _ = v26
                                _ = v27
                                _ = v28
                            }
                        } else {
                            v29 := []int{fi9}
                            var v30 int8 = 100
                            _ = v29
                            _ = v30
                        }
                        _ = v22
                    }
                    var v31 float64 = 1.5
                    v6 = v6
                    switch (-(fi4)) {
                    case 4294967295:
                        v32 := []int{(fi4 | fi4), (fi4 & (fi9 * fi4)), (^(fi9)), 2147483647}
                        _ = v32
                    case 2147483647:
                        var v33 uint64 = v10
                        _ = v33
                    case 128:
                        switch ("\n\t" + (("foo bar" + "a") + ("\n\t" + ("a" + ("a" + ""))))) {
                        case "\n\t":
                            fmt.Println(v10, fi4)
                        case "hello":
                            v10 += v10
                        case "a":
                            v6 += v10
                        }
                    }
                    switch (^(fi4)) {
                    case 1:
                        v6 -= v10
                    case -1:
                        fmt.Println(v31, v5, fi9)
                    case 868:
                        v34 := []int{fi4, fi9, fi9}
                        _ = v34
                    }
                    _ = v31
                }
                _ = fi9
                _ = v10
            }
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi35 := 0; fi35 < 7; fi35++ {
        for fi36 := 0; fi36 < 12; fi36++ {
            var v37 int64 = 132
            _ = fi36
            _ = v37
        }
        _ = fi35
    }
    return 9223372036854775807
}
func genFn1(p0 int32) int64 {
    v38 := []int{4294967295}
    for fi39 := 0; fi39 < 9; fi39++ {
        var v40 bool = false
        v40 = true
        for fi41 := 0; fi41 < 15; fi41++ {
            v42 := GenS1{F0: 0}
            if len(v38) > 0 {
                v38[0] = genFn0()
            }
            var v43 float64 = 0.0
            _ = fi41
            _ = v42
            _ = v43
        }
        _ = fi39
        _ = v40
    }
    p0 -= 0
    _ = len(v38)
    _ = len(v38)
    _ = p0
    _ = v38
    return -553
}

func main() {
    var v1 int16 = 1
    var v2 int64 = -1
    v3 := []string{(("foo bar" + ("hello" + "")) + ("\n\t" + ("foo bar" + ("\n\t" + "a"))))}
    v2 = -240
    fmt.Println(v2, v2)
    v44 := GenS4{F0: -2147483648}
    var v45 int64 = (^(v2))
    if (v45 >= -951) {
        for fi46 := 0; fi46 < 16; fi46++ {
            var v47 float32 = 0.1
            switch fi46 {
            case 261:
                for fi48 := 0; fi48 < 1; fi48++ {
                    if (-500 <= (fi48 | (^(fi46)))) {
                        switch fi46 {
                        case -9223372036854775808:
                            if true {
                                var v49 uint64 = 1
                                var v50 int32 = -1
                                _ = v49
                                _ = v50
                            }
                        case 1:
                            var v51 bool = true
                            _ = v51
                        }
                        v1 += (v1 - 42)
                        v2 = v2
                    } else {
                        var v52 uint64 = 188
                        _ = v52
                    }
                    _ = fi48
                }
            case -9223372036854775808:
                if (2 >= 0) {
                    for fi53 := 0; fi53 < 3; fi53++ {
                        v3 = append(v3, "a")
                        if false {
                            switch (fi46 << ((1) & 63)) {
                            case -385:
                                v54 := GenS2{F0: (!(!(161 != genFn0()))), F1: 0}
                                v55 := GenS1{F0: (v45 / (((v45 / (((v45 - v45)) | 1))) | 1))}
                                v3 = append(v3, "\n\t")
                                _ = v54
                                _ = v55
                            case 0:
                                if false {
                                    _ = len(v3)
                                    v1 -= 1
                                } else {
                                    fmt.Println(fi46, v47)
                                    v2 -= 9223372036854775807
                                }
                            }
                            v56 := GenS5{F0: false}
                            _ = v56
                        }
                        v45 = (v2 + genFn1(1))
                        _ = fi53
                    }
                    fmt.Println(fi46, v2)
                    for fi57 := 0; fi57 < 10; fi57++ {
                        fmt.Println(v1)
                        v58 := GenS4{F0: -682}
                        if true {
                            switch fi46 {
                            case -9223372036854775808:
                                v2 += v45
                            }
                        } else {
                            v59 := GenS6{F0: 42, F1: 0, F2: ""}
                            fmt.Println(fi46, fi46, v45)
                            _ = v59
                        }
                        _ = fi57
                        _ = v58
                    }
                } else {
                    v60 := GenS7{F0: "\n\t", F1: "hello", F2: 830}
                    _ = v60
                }
            }
            _ = fi46
            _ = v47
        }
        var v61 int16 = -1
        for fi62 := 0; fi62 < 19; fi62++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v63 uint64 = 18446744073709551615
            _ = fi62
            _ = v63
        }
        _ = v61
    }
    for fi64 := 0; fi64 < 20; fi64++ {
        v3 = append(v3, "hello")
        var v65 int8 = -1
        _ = fi64
        _ = v65
    }
    v66 := GenS8{F0: 689, F1: 896}
    for fi67 := 0; fi67 < 10; fi67++ {
        switch ("hello" + "foo bar") {
        case "foo bar":
            v3 = append(v3, "hello")
        }
        v45 += (v45 + v45)
        var v68 uint64 = 1
        _ = fi67
        _ = v68
    }
    v45 += (v45 ^ v45)
    if (!(148 != 202)) {
        var v69 uint16 = 255
        for fi70 := 0; fi70 < 18; fi70++ {
            var v71 uint8 = 128
            _ = fi70
            _ = v71
        }
        _ = v69
    }
    if (!false) {
        var v72 int8 = -128
        v45 += v2
        _ = v72
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v66
}
