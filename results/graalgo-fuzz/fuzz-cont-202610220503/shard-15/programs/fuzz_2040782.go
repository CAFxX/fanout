package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 bool; F1 int64; F2 int8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int32; F1 int8 }

func main() {
    var v1 bool = (1 >= -866)
    var v2 string = "hello"
    v3 := []uint32{42}
    v4 := []int{824, -1}
    v5 := make(map[uint32]uint32)
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := GenS1{F0: 127, F1: -2147483648}
        _ = fi6
        _ = v7
    }
    var v8 uint = 65535
    switch -1 {
    case -882:
        for fi9 := 0; fi9 < 16; fi9++ {
            v1 = v1
            _ = fi9
        }
    }
    if v1 {
        v10 := GenS1{F0: 255, F1: -2147483648}
        v1 = (!v1)
        v2 = v2
        _ = v10
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        if v1 {
            v1 = v1
        }
        v12 := GenS2{F0: 1}
        var v13 uint64 = 4294967295
        _ = fi11
        _ = v12
        _ = v13
    }
    v8 += v8
    var v14 int = 336
    var v15 uint16 = 42
    fmt.Println(v15)
    v16 := GenS1{F0: 246, F1: 42}
    for fi17 := 0; fi17 < 3; fi17++ {
        fmt.Println(v8, fi17, v1)
        var v18 uint16 = 65535
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi17
        _ = v18
    }
    v19 := GenS2{F0: (v15 - v15)}
    v14 += v14
    _ = len(v3)
    for fi20 := 0; fi20 < 8; fi20++ {
        for fi21 := 0; fi21 < 4; fi21++ {
            v8 = (v8 ^ 4294967295)
            var v22 int8 = 100
            v22 -= (v22 * (v22 - (v22 + -128)))
            var v23 uint8 = uint8(v8)
            _ = fi21
            _ = v22
            _ = v23
        }
        var v24 int16 = 0
        fmt.Println(v8, v2)
        _ = fi20
        _ = v24
    }
    if v1 {
        v1 = v1
    }
    var v25 int64 = 1
    switch v14 {
    case 128:
        if v1 {
            switch v2 {
            case "\n\t":
                v1 = (247 >= 558)
            case "":
                switch v2 {
                case "hello":
                    for fi26 := 0; fi26 < 14; fi26++ {
                        if len(v4) > 0 {
                            v4[0] = fi26
                        }
                        if false {
                            for fi27 := 0; fi27 < 3; fi27++ {
                                switch -1 {
                                case 4294967295:
                                    if (fi27 <= fi26) {
                                        var v28 float64 = -2.25
                                        if v1 {
                                            delete(v5, 981)
                                        }
                                        var v29 float64 = ((v28 + 3.14159) - v28)
                                        _ = v28
                                        _ = v29
                                    } else {
                                        switch v2 {
                                        case "":
                                            v1 = v1
                                        case "foo bar":
                                            v8 = (^(v8))
                                        }
                                        v30 := v5[42]
                                        var v31 int = (v14 / ((-1) | 1))
                                        fmt.Println(v14, v15, v14)
                                        var v32 int = v14
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                    }
                                case -13:
                                    if v1 {
                                        if (v1 || v1) {
                                            v8 += v8
                                            switch ((v2 + v2) + (v2 + v2)) {
                                            case "a":
                                                switch (v14 * fi26) {
                                                case 148:
                                                    switch (("hello" + (((v2 + "") + v2) + v2)) + v2) {
                                                    case "a":
                                                        if v1 {
                                                            switch fi26 {
                                                            case 127:
                                                                switch (-(fi27)) {
                                                                case 206:
                                                                    if v1 {
                                                                        var v33 uint64 = 0
                                                                        v8 -= v8
                                                                        v34 := GenS2{F0: v15}
                                                                        _ = v33
                                                                        _ = v34
                                                                    } else {
                                                                        if v1 {
                                                                            switch fi26 {
                                                                            case -662:
                                                                                v15 -= v15
                                                                            default:
                                                                                v35 := GenS3{F0: fi26, F1: 911}
                                                                                _ = v35
                                                                            }
                                                                            var v36 int8 = 42
                                                                            switch (v2 + ("foo bar" + ("\n\t" + v2))) {
                                                                            case "a":
                                                                                v37 := GenS2{F0: (v15 - (v15 / (((v15 - 1)) | 1)))}
                                                                                _ = v37
                                                                            case "\n\t":
                                                                                var v38 string = (v2 + "\n\t")
                                                                                _ = v38
                                                                            }
                                                                            _ = v36
                                                                        } else {
                                                                            switch (fi26 ^ (v14 - int(v25))) {
                                                                            case 32768:
                                                                                v1 = v1
                                                                            }
                                                                            v39 := GenS4{F0: v1, F1: (v25 & (v25 << (((v8 & 126)) & 63))), F2: 127}
                                                                            v40 := v5[1]
                                                                            _ = v39
                                                                            _ = v40
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                case -1:
                                                    v14 = -2147483648
                                                default:
                                                    switch v2 {
                                                    case "\n\t":
                                                        fmt.Println(v14)
                                                    default:
                                                        fmt.Println(fi27)
                                                    }
                                                }
                                            default:
                                                switch (fi27 % ((fi26) | 1)) {
                                                case -1:
                                                    _ = len(v5)
                                                case 1:
                                                    v41 := GenS4{F0: ((!(true || ((!(v1 && ((v15 >> ((v8) & 63)) != 65535))) || v1))) && ((v1 && v1) && (v15 == v15))), F1: v25, F2: 42}
                                                    _ = v41
                                                default:
                                                    var v42 int16 = 42
                                                    _ = v42
                                                }
                                            }
                                            switch v14 {
                                            case -2147483648:
                                                var v43 int16 = -46
                                                _ = v43
                                            case 963:
                                                fmt.Println(v15)
                                                v2 = "foo bar"
                                                if ((v1 || true) || (0.0 < 1e10)) {
                                                    var v44 int64 = 0
                                                    var v45 uint = 65535
                                                    var v46 uint = (v45 ^ (v8 & (v8 / ((v8) | 1))))
                                                    _ = v44
                                                    _ = v45
                                                    _ = v46
                                                } else {
                                                    fmt.Println(v14, v15, v2)
                                                    var v47 int64 = (^(v25))
                                                    var v48 int32 = -805
                                                    var v49 uint64 = 2
                                                    _ = v47
                                                    _ = v48
                                                    _ = v49
                                                }
                                            case -376:
                                                v50 := GenS3{F0: fi26, F1: 793}
                                                _ = v50
                                            }
                                        } else {
                                            fmt.Println(v15, v25)
                                            v25 = v25
                                            if (!(42 > -2147483648)) {
                                                var v51 int64 = (v25 & 232)
                                                _ = v51
                                            }
                                            v52 := GenS2{F0: v15}
                                            var v53 string = "\n\t"
                                            var v54 bool = v1
                                            v53 = "\n\t"
                                            var v55 int = (fi27 - (v14 * 0))
                                            switch "hello" {
                                            case "":
                                                var v56 uint8 = 2
                                                _ = v56
                                            case "foo bar":
                                                switch (v55 & (fi27 + -162)) {
                                                case 960:
                                                    fmt.Println(v2)
                                                }
                                            case "a":
                                                fmt.Println(v55)
                                            }
                                            _ = v52
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                        }
                                        v25 = v25
                                        if v1 {
                                            if (v1 && (!false)) {
                                                var v57 int = (fi27 + (^(v14)))
                                                _ = v57
                                            } else {
                                                v58 := GenS5{F0: int(v15)}
                                                switch v2 {
                                                case "a":
                                                    if len(v4) > 0 {
                                                        v4[0] = 0
                                                    }
                                                case "\n\t":
                                                    v59 := GenS6{F0: -30, F1: -108}
                                                    _ = v59
                                                case "hello":
                                                    fmt.Println(fi27)
                                                }
                                                _ = v58
                                            }
                                            _ = len(v3)
                                        }
                                    }
                                }
                                var v60 int8 = 1
                                var v61 float64 = 1e10
                                _ = fi27
                                _ = v60
                                _ = v61
                            }
                            fmt.Println(fi26, fi26, v25)
                        } else {
                            if (-1 != (^(v14))) {
                                v14 += 127
                                v8 += v8
                            }
                        }
                        var v62 bool = v1
                        _ = fi26
                        _ = v62
                    }
                default:
                    for fi63 := 0; fi63 < 7; fi63++ {
                        switch ("a" + "foo bar") {
                        case "foo bar":
                            for fi64 := 0; fi64 < 17; fi64++ {
                                v25 -= v25
                                var v65 float64 = (0.1 * -2.25)
                                fmt.Println(v15, v8, v8)
                                _ = fi64
                                _ = v65
                            }
                        case "a":
                            v15 -= v15
                        default:
                            for fi66 := 0; fi66 < 16; fi66++ {
                                v15 -= v15
                                var v67 uint32 = 0
                                _ = fi66
                                _ = v67
                            }
                        }
                        for fi68 := 0; fi68 < 18; fi68++ {
                            fmt.Println(v14)
                            var v69 int = v14
                            _ = fi68
                            _ = v69
                        }
                        _ = fi63
                    }
                }
            default:
                _ = len(v4)
            }
            for fi70 := 0; fi70 < 13; fi70++ {
                var v71 int64 = v25
                v1 = (v1 && (false && v1))
                v8 = (v8 + v8)
                _ = fi70
                _ = v71
            }
        }
    case 32767:
        fmt.Println(v14)
    case 32768:
        v72 := v5[4294967295]
        _ = v72
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
