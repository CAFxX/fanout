package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 int; F2 int8 }
func genFn0() int8 {
    v4 := []int{-1, 315, 4294967295, -9223372036854775808}
    var v5 uint8 = 1
    _ = len(v4)
    v5 = 0
    var v6 int8 = 0
    _ = v4
    _ = v5
    _ = v6
    return v6
}
func genFn1() uint8 {
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := []int{(fi7 << ((2) & 63))}
        var v9 uint16 = 732
        _ = fi7
        _ = v8
        _ = v9
    }
    switch 2147483647 {
    case 32768:
        switch 128 {
        case -1:
            var v10 int32 = -108
            _ = v10
        case 9223372036854775807:
            var v11 int32 = 42
            _ = v11
        case 1:
            switch "\n\t" {
            case "a":
                var v12 int32 = -547
                _ = v12
            case "hello":
                v13 := []int{}
                _ = v13
            case "":
                var v14 string = "foo bar"
                _ = v14
            }
        }
    case 9223372036854775807:
        switch -871 {
        case 2147483647:
            for fi15 := 0; fi15 < 9; fi15++ {
                for fi16 := 0; fi16 < 1; fi16++ {
                    var v17 uint16 = 32767
                    _ = fi16
                    _ = v17
                }
                switch (("\n\t" + "a") + ("\n\t" + ("hello" + "a"))) {
                case "a":
                    fmt.Println(fi15)
                case "foo bar":
                    fmt.Println(fi15)
                default:
                    for fi18 := 0; fi18 < 10; fi18++ {
                        var v19 float64 = 1e10
                        var v20 uint = 42
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                }
                _ = fi15
            }
        case -926:
            for fi21 := 0; fi21 < 9; fi21++ {
                var v22 string = "a"
                _ = fi21
                _ = v22
            }
        default:
            var v23 uint64 = 228
            _ = v23
        }
    case 137:
        if false {
            var v24 float64 = 1.5
            for fi25 := 0; fi25 < 1; fi25++ {
                v26 := GenS1{F0: 575}
                var v27 int8 = -1
                _ = fi25
                _ = v26
                _ = v27
            }
            var v28 int8 = 127
            _ = v24
            _ = v28
        } else {
            var v29 float32 = 0.0
            _ = v29
        }
    default:
        if (!true) {
            var v30 int8 = 68
            var v31 uint8 = 128
            _ = v30
            _ = v31
        } else {
            var v32 uint64 = 4294967295
            v33 := []int{1, 4294967295, 42, -498}
            _ = v32
            _ = v33
        }
    }
    v34 := []int{223, -159, 9223372036854775807, 248}
    _ = v34
    return 42
}

func main() {
    var v1 float32 = float32((((0.0 * (((float64(((0.1 * 1.5) - 0.1)) / 1e10) - 3.14159) - (1.5 + float64(0.0)))) / 0.1) / 1.0))
    var v2 uint32 = 1
    var v3 float64 = (3.14159 / 3.14159)
    fmt.Println(v1, v3, v3)
    v3 -= 1e10
    v2 = (v2 / ((2) | 1))
    v2 = 1
    switch "" {
    case "a":
        for fi35 := 0; fi35 < 15; fi35++ {
            switch (fi35 - fi35) {
            case -1:
                v2 = v2
            default:
                v36 := GenS1{F0: -669}
                _ = v36
            }
            for fi37 := 0; fi37 < 15; fi37++ {
                v1 += (3.14159 / v1)
                var v38 uint64 = 42
                _ = fi37
                _ = v38
            }
            v1 = 0.1
            _ = fi35
        }
    case "foo bar":
        if true {
            fmt.Println(v2)
            v2 = (v2 | v2)
            switch "\n\t" {
            case "\n\t":
                var v39 int64 = -1
                _ = v39
            case "foo bar":
                v3 -= v3
            case "":
                if ((((1e10 / float32(v3)) + 1.5) + (0.1 * float32(0.1))) < 3.14159) {
                    fmt.Println(v1, v2)
                }
            default:
                var v40 bool = true
                _ = v40
            }
        } else {
            for fi41 := 0; fi41 < 2; fi41++ {
                var v42 uint32 = v2
                v43 := []int{(fi41 >> ((2) & 63)), (fi41 & 950)}
                _ = fi41
                _ = v42
                _ = v43
            }
            if false {
                v44 := GenS1{F0: -989}
                _ = v44
            } else {
                if (1 != 1) {
                    var v45 string = "hello"
                    v46 := GenS2{F0: true, F1: ((v45 + "") + (v45 + v45))}
                    _ = v45
                    _ = v46
                } else {
                    if true {
                        var v47 float64 = v3
                        v48 := GenS2{F0: true, F1: ("foo bar" + ("a" + "a"))}
                        _ = v47
                        _ = v48
                    } else {
                        v2 = v2
                    }
                    if false {
                        v3 -= v3
                        switch -427 {
                        case 0:
                            v2 = 65535
                        case 4294967295:
                            var v49 int16 = 1000
                            _ = v49
                        default:
                            v50 := []int{-1, 128, -9223372036854775808}
                            _ = v50
                        }
                    }
                    if false {
                        var v51 int8 = genFn0()
                        v3 += 0.0
                        var v52 string = "\n\t"
                        _ = v51
                        _ = v52
                    }
                }
                switch 128 {
                case -514:
                    var v53 int = -9223372036854775808
                    _ = v53
                case -1:
                    switch 42 {
                    case 32767:
                        fmt.Println(v3, v1, v2)
                    case 778:
                        v54 := GenS2{F0: ((float32((v3 + v3)) + (1e10 + (v1 - v1))) < float32(v3)), F1: ""}
                        _ = v54
                    case 728:
                        fmt.Println(v2)
                    default:
                        v55 := GenS1{F0: -2147483648}
                        _ = v55
                    }
                case 1:
                    switch 42 {
                    case 511:
                        v56 := GenS3{F0: 4294967295}
                        _ = v56
                    default:
                        v57 := GenS4{F0: v2, F1: 127, F2: -85}
                        _ = v57
                    }
                default:
                    v3 = 3.14159
                }
                if (-1 < -418) {
                    var v58 int8 = 100
                    switch "\n\t" {
                    case "foo bar":
                        var v59 int = 2147483647
                        _ = v59
                    case "\n\t":
                        switch "a" {
                        case "hello":
                            if false {
                                switch ("a" + "\n\t") {
                                case "\n\t":
                                    switch -869 {
                                    case -9223372036854775808:
                                        var v60 uint64 = 0
                                        _ = v60
                                    }
                                }
                                v61 := []int{9223372036854775807}
                                fmt.Println(v1)
                                _ = v61
                            }
                        default:
                            v62 := []int{}
                            _ = v62
                        }
                    case "hello":
                        for fi63 := 0; fi63 < 10; fi63++ {
                            v64 := []int{(fi63 - (fi63 << ((188) & 63))), fi63, (fi63 % (((fi63 ^ (fi63 * (fi63 >> ((441) & 63))))) | 1)), fi63}
                            _ = fi63
                            _ = v64
                        }
                    }
                    var v65 uint = 4294967295
                    _ = v58
                    _ = v65
                } else {
                    for fi66 := 0; fi66 < 7; fi66++ {
                        var v67 string = (("" + "a") + "\n\t")
                        _ = fi66
                        _ = v67
                    }
                    for fi68 := 0; fi68 < 19; fi68++ {
                        v3 = float64(v1)
                        for fi69 := 0; fi69 < 11; fi69++ {
                            if (((false && (true && (((v2 << ((255) & 63)) <= (v2 >> ((502) & 63))) && false))) && false) || false) {
                                var v70 uint64 = 4294967295
                                _ = v70
                            }
                            fmt.Println(v3, fi69, fi68)
                            _ = fi69
                        }
                        v1 = (0.0 * ((0.0 - ((1e10 + (3.14159 - v1)) + ((1.5 - (0.0 + (v1 * 0.0))) * v1))) + (v1 + float32(((v3 * (((v3 * ((1.5 / float64(((-2.25 - 0.0) + v1))) / float64(v1))) + v3) + (v3 * 1.5))) + (0.0 / (0.1 + v3)))))))
                        _ = fi68
                    }
                    var v71 uint64 = 0
                    _ = v71
                }
            }
        }
    default:
        for fi72 := 0; fi72 < 11; fi72++ {
            var v73 string = "\n\t"
            _ = fi72
            _ = v73
        }
    }
    switch ("" + ("hello" + ("hello" + ""))) {
    case "foo bar":
        switch 32768 {
        case 4294967295:
            fmt.Println(v1)
        case 42:
            v2 += v2
        case -1:
            switch ("a" + ((((("a" + "\n\t") + (("hello" + ("a" + ("foo bar" + "foo bar"))) + "hello")) + ((("" + "foo bar") + (("\n\t" + "") + "hello")) + "foo bar")) + "a") + (("a" + "hello") + ""))) {
            case "\n\t":
                v3 -= (v3 + (0.1 + 0.1))
            case "foo bar":
                var v74 int64 = 941
                _ = v74
            case "a":
                v3 = v3
            default:
                v75 := []int{-28}
                _ = v75
            }
        default:
            var v76 uint64 = 0
            _ = v76
        }
    default:
        var v77 int16 = 1
        _ = v77
    }
    switch ("hello" + "a") {
    case "foo bar":
        if false {
            var v78 int16 = -32768
            var v79 int16 = (v78 ^ v78)
            v80 := []int{9223372036854775807, -698, -956}
            _ = v78
            _ = v79
            _ = v80
        } else {
            v81 := []int{}
            v3 = (0.1 + float64(1e10))
            var v82 uint32 = v2
            _ = v81
            _ = v82
        }
    case "hello":
        var v83 int8 = 35
        _ = v83
    }
    for fi84 := 0; fi84 < 15; fi84++ {
        var v85 float64 = float64(v1)
        _ = fi84
        _ = v85
    }
    v86 := []int{-9223372036854775808, -52, 1, -939}
    var v87 string = "\n\t"
    if (true && true) {
        v2 = (^(v2))
        fmt.Println(v1, v2)
    }
    v87 = v87
    for fi88 := 0; fi88 < 10; fi88++ {
        v1 = 1e10
        _ = fi88
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v86:", len(v86), v86)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v86
    _ = v87
}
