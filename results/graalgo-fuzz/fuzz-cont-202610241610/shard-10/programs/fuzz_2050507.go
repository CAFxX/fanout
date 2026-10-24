package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 uint; F1 string; F2 uint }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 uint8; F1 int32; F2 int64 }
type GenS6 struct { F0 bool; F1 bool }
type GenS7 struct { F0 int64; F1 bool; F2 int64 }
type GenS8 struct { F0 uint; F1 uint64; F2 int8 }
type GenS9 struct { F0 uint32; F1 int16 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 bool; F1 uint64 }
func genFn0(p0 int64, p1 int64) int64 {
    p1 = 42
    v4 := []int{1, 32767}
    var v5 int32 = -1
    for fi6 := 0; fi6 < 10; fi6++ {
        if false {
            var v7 uint8 = 83
            var v8 int16 = -32768
            if true {
                if false {
                    switch (fi6 + (fi6 & fi6)) {
                    case 141:
                        var v9 uint16 = 137
                        _ = v9
                    case -2147483648:
                        v10 := GenS1{F0: 42}
                        _ = v10
                    case 32767:
                        v11 := GenS1{F0: 2147483648}
                        _ = v11
                    default:
                        if (true || (true || true)) {
                            switch "" {
                            case "hello":
                                if false {
                                    var v12 uint = 689
                                    v12 = v12
                                    v12 = v12
                                    _ = v12
                                }
                            }
                        } else {
                            v7 = 0
                        }
                    }
                }
            }
            _ = v7
            _ = v8
        }
        var v13 int32 = -2147483648
        for fi14 := 0; fi14 < 12; fi14++ {
            v15 := GenS2{F0: 9223372036854775807, F1: 2}
            var v16 int64 = p1
            switch fi6 {
            case -2147483648:
                v16 += p1
            default:
                var v17 uint32 = 42
                _ = v17
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi6
        _ = v13
    }
    v5 -= -2147483648
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return p1
}
func genFn1() int {
    for fi18 := 0; fi18 < 20; fi18++ {
        if true {
            v19 := GenS2{F0: -935, F1: 155}
            var v20 uint64 = 1
            v21 := GenS3{F0: 4294967295, F1: "hello", F2: 1}
            _ = v19
            _ = v20
            _ = v21
        } else {
            var v22 float64 = ((-2.25 + ((float64(0.0) * (0.0 + float64(0.0))) - 1e10)) - 3.14159)
            _ = v22
        }
        _ = fi18
    }
    if (false && (true || false)) {
        var v23 string = "hello"
        if true {
            v24 := []int{}
            _ = v24
        }
        _ = v23
    }
    switch 1 {
    case -646:
        var v25 int32 = -733
        _ = v25
    case 2147483647:
        if false {
            v26 := GenS1{F0: 42}
            _ = v26
        } else {
            if (((456 >= 255) && ((float32(3.14159) * -2.25) != ((1.5 / (1.5 * (float32(3.14159) / -2.25))) - 3.14159))) || (!false)) {
                var v27 uint16 = 277
                _ = v27
            }
            switch -2147483648 {
            case 42:
                v28 := []int{}
                _ = v28
            case 1:
                var v29 uint16 = 640
                _ = v29
            case 127:
                v30 := GenS2{F0: -819, F1: 1}
                _ = v30
            default:
                var v31 float64 = float64(0.0)
                _ = v31
            }
            for fi32 := 0; fi32 < 15; fi32++ {
                var v33 uint64 = 4294967295
                _ = fi32
                _ = v33
            }
        }
    }
    v34 := GenS1{F0: 2}
    _ = v34
    return 0
}

func main() {
    v1 := []int{9223372036854775807, -1}
    v2 := make(map[string]string)
    v3 := make(map[int]string)
    var v35 float32 = -2.25
    v36 := GenS4{F0: -128, F1: 1}
    switch "foo bar" {
    case "\n\t":
        v37 := GenS1{F0: 2147483647}
        _ = v37
    case "a":
        var v38 uint32 = 65535
        _ = v38
    case "":
        v35 -= ((float32((1.5 + (-2.25 * 3.14159))) + v35) + -2.25)
    default:
        v35 += 1e10
    }
    for fi39 := 0; fi39 < 13; fi39++ {
        v40 := GenS1{F0: 65535}
        v35 -= v35
        _ = fi39
        _ = v40
    }
    if false {
        var v41 int32 = -1
        _ = v41
    } else {
        var v42 float64 = 1.5
        _ = v42
    }
    v35 = ((v35 + float32(3.14159)) / v35)
    if false {
        switch "\n\t" {
        case "":
            if (-128 == 42) {
                if true {
                    var v43 int64 = -9223372036854775808
                    var v44 int16 = -32768
                    var v45 uint16 = 65535
                    _ = v43
                    _ = v44
                    _ = v45
                } else {
                    fmt.Println(v35)
                }
                var v46 int64 = -838
                var v47 uint8 = 2
                _ = v46
                _ = v47
            }
        case "\n\t":
            v35 -= v35
        }
        var v48 int64 = 0
        if true {
            for fi49 := 0; fi49 < 12; fi49++ {
                var v50 bool = ("foo bar" >= ("foo bar" + (("hello" + (("\n\t" + "a") + "a")) + (("\n\t" + (("\n\t" + (("hello" + ("hello" + ("" + ""))) + ("\n\t" + ""))) + ("" + "foo bar"))) + (("\n\t" + (("" + ((("" + (("" + "hello") + ((((("hello" + ("foo bar" + "hello")) + "foo bar") + ("" + ("hello" + ""))) + "\n\t") + ""))) + "") + "\n\t")) + (("\n\t" + ("\n\t" + "hello")) + "hello"))) + "\n\t")))))
                _ = fi49
                _ = v50
            }
        } else {
            switch "" {
            case "foo bar":
                fmt.Println(v48, v35)
            default:
                for fi51 := 0; fi51 < 19; fi51++ {
                    var v52 int = (fi51 - (fi51 % ((fi51) | 1)))
                    switch genFn1() {
                    case 1:
                        v53 := GenS5{F0: 171, F1: 1, F2: v48}
                        _ = v53
                    default:
                        for fi54 := 0; fi54 < 7; fi54++ {
                            switch "foo bar" {
                            case "\n\t":
                                fmt.Println(v52, v35)
                            case "foo bar":
                                _ = len(v2)
                            default:
                                v35 += 0.1
                            }
                            _ = len(v3)
                            _ = fi54
                        }
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi51
                    _ = v52
                }
            }
        }
        _ = v48
    } else {
        fmt.Println(v35)
        var v55 uint16 = 159
        if true {
            fmt.Println(v35)
        }
        _ = v55
    }
    for fi56 := 0; fi56 < 20; fi56++ {
        var v57 float64 = (1e10 - 1e10)
        _ = fi56
        _ = v57
    }
    delete(v2, ("\n\t" + "a"))
    for fi58 := 0; fi58 < 10; fi58++ {
        v59 := GenS6{F0: (!(false || true)), F1: (499 < 128)}
        v35 += float32(3.14159)
        _ = fi58
        _ = v59
    }
    fmt.Println(v35)
    v35 += 0.1
    var v60 uint64 = 577
    if true {
        for fi61 := 0; fi61 < 1; fi61++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if false {
                for fi62 := 0; fi62 < 11; fi62++ {
                    switch ("\n\t" + "\n\t") {
                    case "a":
                        var v63 int = (fi61 * fi61)
                        _ = v63
                    case "\n\t":
                        var v64 float64 = 0.0
                        v65 := GenS7{F0: 469, F1: (fi61 >= (fi61 * 2147483647)), F2: -9223372036854775808}
                        _ = v64
                        _ = v65
                    case "hello":
                        v66 := v3[fi61]
                        _ = v66
                    default:
                        v67 := GenS8{F0: 0, F1: (v60 % ((761) | 1)), F2: -69}
                        _ = v67
                    }
                    var v68 bool = (!true)
                    _ = fi62
                    _ = v68
                }
            }
            _ = fi61
        }
        switch "a" {
        case "":
            if false {
                fmt.Println(v35, v35)
            }
        default:
            v35 = v35
        }
    }
    var v69 int64 = 1
    v35 -= (0.0 / v35)
    v2[(("hello" + "") + ("a" + (((("\n\t" + ((("foo bar" + ("" + ("" + "hello"))) + "a") + ("\n\t" + ""))) + "foo bar") + "hello") + "\n\t")))] = "foo bar"
    switch -845 {
    case 893:
        v69 = -9223372036854775808
    case 27:
        switch ("hello" + "foo bar") {
        case "foo bar":
            v70 := GenS9{F0: 2147483648, F1: -1}
            _ = v70
        case "\n\t":
            v71 := GenS10{F0: (^(v60))}
            _ = v71
        case "hello":
            fmt.Println(v60, v69, v69)
        default:
            for fi72 := 0; fi72 < 7; fi72++ {
                if false {
                    var v73 int = -953
                    _ = v73
                } else {
                    for fi74 := 0; fi74 < 13; fi74++ {
                        switch "\n\t" {
                        case "\n\t":
                            v69 = v69
                        default:
                            var v75 int64 = v69
                            _ = v75
                        }
                        var v76 float32 = -2.25
                        _ = fi74
                        _ = v76
                    }
                    v77 := GenS7{F0: (v69 + (v69 & v69)), F1: (0 == 42), F2: v69}
                    v78 := GenS3{F0: 720, F1: "hello", F2: 255}
                    _ = v77
                    _ = v78
                }
                v79 := GenS11{F0: ((((float32(1.5) < v35) && true) && ((v60 * 347) >= (v60 / ((910) | 1)))) && (true && true)), F1: v60}
                _ = fi72
                _ = v79
            }
        }
    default:
        v80 := GenS8{F0: 42, F1: v60, F2: -1}
        _ = v80
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v60:", v60)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v60
    _ = v69
}
