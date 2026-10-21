package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 int; F2 int }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 uint16; F1 int64; F2 bool }
type GenS6 struct { F0 string; F1 int32; F2 uint }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int16; F1 int32; F2 uint32 }
type GenS9 struct { F0 uint8 }
func genFn0() uint64 {
    var v5 int8 = 100
    var v6 uint16 = 255
    fmt.Println(v6, v5)
    v5 += (v5 + (v5 << ((2) & 63)))
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[uint32]bool)
    var v3 uint16 = 34
    var v4 int8 = -109
    fmt.Println(v4)
    switch 462 {
    case -1:
        var v7 float64 = ((0.1 * (1e10 / ((((1e10 * (((float64((0.0 * 1.5)) * 3.14159) / 3.14159) * 0.1)) * ((float64(1.5) * -2.25) - 3.14159)) * -2.25) + (1e10 + (1.5 - ((((-2.25 + ((float64(-2.25) + 1e10) - 1e10)) + 0.0) / 1e10) - 0.1)))))) / 1.0)
        _ = v7
    case -779:
        var v8 bool = (((-1 >= 4294967295) && (uint8(v1) >= 2)) && true)
        _ = v8
    case 4294967295:
        if ((!(true && false)) && ((42 == 42) && (-487 == 1000))) {
            switch 1 {
            case 0:
                switch int(v4) {
                case 0:
                    for fi9 := 0; fi9 < 10; fi9++ {
                        var v10 int32 = -1
                        _ = fi9
                        _ = v10
                    }
                case -2147483648:
                    v11 := GenS1{F0: 676}
                    _ = v11
                default:
                    switch 2147483647 {
                    case 0:
                        for fi12 := 0; fi12 < 2; fi12++ {
                            v13 := GenS1{F0: genFn0()}
                            switch fi12 {
                            case 9223372036854775807:
                                v4 = v4
                            case 570:
                                v14 := GenS2{F0: -63, F1: 42, F2: fi12}
                                _ = v14
                            case -2147483648:
                                switch ("foo bar" + "") {
                                case "a":
                                    v1 = (v1 | v1)
                                case "":
                                    switch (fi12 << ((0) & 63)) {
                                    case -622:
                                        v3 += 342
                                    case -1:
                                        v15 := GenS3{F0: ("foo bar" + "foo bar"), F1: (^(v3))}
                                        _ = v15
                                    case -389:
                                        v16 := GenS4{F0: v4, F1: (v4 | (v4 << ((18446744073709551615) & 63)))}
                                        _ = v16
                                    default:
                                        fmt.Println(v1, fi12, v4)
                                    }
                                default:
                                    v17 := GenS5{F0: 0, F1: -466, F2: true}
                                    _ = v17
                                }
                            }
                            var v18 int = (fi12 + fi12)
                            _ = fi12
                            _ = v13
                            _ = v18
                        }
                    }
                }
            case -481:
                for fi19 := 0; fi19 < 17; fi19++ {
                    var v20 uint32 = uint32(v3)
                    for fi21 := 0; fi21 < 4; fi21++ {
                        var v22 int32 = 42
                        switch "" {
                        case "foo bar":
                            switch (fi21 * 127) {
                            case 127:
                                var v23 bool = false
                                _ = v23
                            case -485:
                                var v24 int16 = 823
                                _ = v24
                            }
                        case "":
                            fmt.Println(v4, v4, fi19)
                        case "hello":
                            v25 := []int{-764, 2147483647, (-(fi19)), (fi19 | fi21)}
                            _ = v25
                        }
                        var v26 uint64 = 2
                        _ = fi21
                        _ = v22
                        _ = v26
                    }
                    for fi27 := 0; fi27 < 11; fi27++ {
                        switch (fi27 & (fi27 + (-(fi27)))) {
                        case 32768:
                            switch (("hello" + (("" + "") + "\n\t")) + "foo bar") {
                            case "a":
                                var v28 int64 = 1
                                _ = v28
                            case "foo bar":
                                switch fi19 {
                                case 4294967295:
                                    if (!true) {
                                        v29 := GenS6{F0: (((("" + (("\n\t" + ((((("\n\t" + "") + "\n\t") + "") + "hello") + "hello")) + "hello")) + "hello") + ("foo bar" + ("hello" + ""))) + ("hello" + (("" + "foo bar") + "hello"))), F1: 42, F2: 2}
                                        if false {
                                            switch ("a" + ("\n\t" + ("" + ("hello" + ((("hello" + (("a" + ("hello" + ("hello" + "\n\t"))) + "a")) + "") + "\n\t"))))) {
                                            case "a":
                                                if (true || (false || (false || true))) {
                                                    var v30 uint8 = 169
                                                    v31 := GenS4{F0: v4, F1: (v4 >> ((2) & 63))}
                                                    v30 = v30
                                                    _ = v30
                                                    _ = v31
                                                } else {
                                                    var v32 int64 = 604
                                                    v1 = v20
                                                    v3 = (v3 | (^(v3)))
                                                    var v33 uint64 = 42
                                                    _ = v32
                                                    _ = v33
                                                }
                                            case "":
                                                fmt.Println(v20, v1, fi19)
                                            default:
                                                if true {
                                                    v34 := []int{(-(fi19)), 439}
                                                    v3 = (v3 << ((255) & 63))
                                                    fmt.Println(fi27, v1)
                                                    _ = v34
                                                } else {
                                                    switch "hello" {
                                                    case "\n\t":
                                                        v1 = (v1 - (v20 + v20))
                                                    }
                                                }
                                            }
                                            var v35 int32 = -844
                                            var v36 float32 = (3.14159 + 0.1)
                                            _ = v35
                                            _ = v36
                                        } else {
                                            if false {
                                                var v37 float32 = (1e10 * 0.0)
                                                if (false || true) {
                                                    v20 -= 479
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v38 uint32 = (v1 | (v1 & 4294967295))
                                                        _ = v38
                                                    case "a":
                                                        v4 = v4
                                                    }
                                                    switch fi27 {
                                                    case 127:
                                                        v3 = (^(v3))
                                                    case 42:
                                                        v39 := []int{32768, (^(fi27)), (fi27 >> ((226) & 63))}
                                                        _ = v39
                                                    case 831:
                                                        v40 := GenS5{F0: v3, F1: -9223372036854775808, F2: (false || (1000 <= 1000))}
                                                        _ = v40
                                                    }
                                                } else {
                                                    v41 := v2[2147483648]
                                                    switch fi27 {
                                                    case -1:
                                                        _ = len(v2)
                                                        fmt.Println(fi19)
                                                    case 9223372036854775807:
                                                        v3 = 42
                                                    default:
                                                        v42 := GenS7{F0: v4}
                                                        var v43 float64 = -2.25
                                                        v44 := GenS5{F0: (v3 + v3), F1: 42, F2: (1 != 18446744073709551615)}
                                                        if ((v43 + (v43 + 1.5)) <= 0.0) {
                                                            fmt.Println(v3, fi27, fi27)
                                                            var v45 uint32 = 967
                                                            _ = v45
                                                        }
                                                        var v46 int8 = 0
                                                        _ = v42
                                                        _ = v43
                                                        _ = v44
                                                        _ = v46
                                                    }
                                                    var v47 uint8 = 0
                                                    _ = v41
                                                    _ = v47
                                                }
                                                if true {
                                                    var v48 string = "a"
                                                    var v49 int32 = 410
                                                    v37 -= 3.14159
                                                    v20 -= (v1 * (v1 % (((v20 + v20)) | 1)))
                                                    _ = len(v2)
                                                    _ = v48
                                                    _ = v49
                                                }
                                                v50 := GenS8{F0: 145, F1: 2147483647, F2: (^(v1))}
                                                _ = v37
                                                _ = v50
                                            }
                                        }
                                        v3 -= (v3 >> ((92) & 63))
                                        _ = v29
                                    } else {
                                        var v51 uint = 454
                                        v20 = (^(v20))
                                        v3 += v3
                                        _ = v51
                                    }
                                }
                                switch "foo bar" {
                                case "":
                                    fmt.Println(fi27, v1)
                                }
                            case "":
                                v52 := GenS7{F0: (v4 << ((65535) & 63))}
                                var v53 uint32 = 4294967295
                                v53 = (v1 % ((1) | 1))
                                var v54 uint64 = 2
                                _ = v52
                                _ = v53
                                _ = v54
                            default:
                                v55 := []int{-2147483648}
                                _ = v55
                            }
                        default:
                            var v56 string = ""
                            _ = v56
                        }
                        v1 -= v1
                        v57 := GenS9{F0: 127}
                        _ = fi27
                        _ = v57
                    }
                    _ = fi19
                    _ = v20
                }
            case 4294967295:
                for fi58 := 0; fi58 < 7; fi58++ {
                    delete(v2, v1)
                    _ = fi58
                }
            default:
                if false {
                    var v59 float32 = 3.14159
                    _ = v59
                } else {
                    fmt.Println(v4)
                    if true {
                        switch "foo bar" {
                        case "foo bar":
                            switch 600 {
                            case 2147483647:
                                var v60 int64 = -1
                                _ = v60
                            case 1:
                                if (127 < 0) {
                                    if true {
                                        v61 := GenS6{F0: (("hello" + (("foo bar" + "hello") + "\n\t")) + "\n\t"), F1: 1, F2: 267}
                                        fmt.Println(v1)
                                        v62 := GenS1{F0: 722}
                                        _ = v61
                                        _ = v62
                                    }
                                } else {
                                    var v63 int32 = -1
                                    for fi64 := 0; fi64 < 18; fi64++ {
                                        for fi65 := 0; fi65 < 16; fi65++ {
                                            v4 += v4
                                            delete(v2, uint32(v3))
                                            _ = fi65
                                        }
                                        var v66 int32 = v63
                                        var v67 int8 = 127
                                        _ = fi64
                                        _ = v66
                                        _ = v67
                                    }
                                    var v68 uint8 = 28
                                    _ = v63
                                    _ = v68
                                }
                            case -1:
                                v69 := []int{-28, 400, -2147483648}
                                _ = v69
                            default:
                                var v70 uint32 = (^(v1))
                                _ = v70
                            }
                        case "a":
                            for fi71 := 0; fi71 < 7; fi71++ {
                                var v72 uint32 = 764
                                v1 = v72
                                v73 := v2[v72]
                                _ = fi71
                                _ = v72
                                _ = v73
                            }
                        case "hello":
                            for fi74 := 0; fi74 < 11; fi74++ {
                                var v75 int32 = -1
                                v3 = 0
                                _ = fi74
                                _ = v75
                            }
                        }
                    }
                }
            }
            v3 -= (^(v3))
            v76 := []int{}
            _ = v76
        } else {
            fmt.Println(v3, v3)
        }
    default:
        var v77 uint64 = 42
        _ = v77
    }
    var v78 uint16 = (v3 * (v3 + v3))
    var v79 int64 = 738
    v4 = v4
    var v80 int8 = (v4 & -128)
    switch 9223372036854775807 {
    case 32767:
        var v81 uint16 = (^(v3))
        _ = v81
    case 128:
        if (-228 < -2147483648) {
            fmt.Println(v4, v4)
            var v82 int64 = (^(v79))
            _ = v82
        }
    default:
        v80 += int8(v79)
    }
    v83 := GenS8{F0: -32768, F1: 1, F2: v1}
    var v84 bool = true
    v80 = (v80 - v4)
    v85 := v2[v1]
    v3 += (v78 >> ((0) & 63))
    var v86 float32 = (-2.25 / 1.5)
    v79 = v79
    var v87 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v78
    _ = v79
    _ = v80
    _ = v83
    _ = v84
    _ = v85
    _ = v86
    _ = v87
}
