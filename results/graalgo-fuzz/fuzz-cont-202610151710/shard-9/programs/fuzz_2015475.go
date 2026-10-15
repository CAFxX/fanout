package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int32; F2 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int; F1 int16; F2 uint32 }
type GenS6 struct { F0 bool; F1 int; F2 int64 }
type GenS7 struct { F0 uint64 }
func genFn0(p0 uint16) int64 {
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := GenS1{F0: 0, F1: 4294967295}
        var v7 uint64 = 42
        _ = fi5
        _ = v6
        _ = v7
    }
    switch ("hello" + "") {
    case "":
        if ((!false) && true) {
            var v8 uint8 = 42
            fmt.Println(v8)
            switch 32768 {
            case -2147483648:
                var v9 int16 = 1000
                _ = v9
            }
            _ = v8
        }
    case "foo bar":
        var v10 int8 = -128
        _ = v10
    case "a":
        if false {
            for fi11 := 0; fi11 < 7; fi11++ {
                switch (fi11 - (fi11 / ((fi11) | 1))) {
                case 127:
                    var v12 uint = 764
                    _ = v12
                case 466:
                    for fi13 := 0; fi13 < 9; fi13++ {
                        var v14 uint32 = 42
                        _ = fi13
                        _ = v14
                    }
                }
                switch ("foo bar" + ("hello" + "hello")) {
                case "":
                    p0 = (p0 + p0)
                case "foo bar":
                    v15 := GenS1{F0: -128, F1: 2}
                    _ = v15
                case "hello":
                    p0 = p0
                }
                p0 = (p0 / ((p0) | 1))
                _ = fi11
            }
            var v16 uint16 = p0
            if (42 < 4294967295) {
                v17 := []int{-849}
                var v18 int32 = 954
                _ = v17
                _ = v18
            }
            _ = v16
        }
    default:
        if true {
            switch "foo bar" {
            case "a":
                if true {
                    switch 427 {
                    case 1:
                        p0 = p0
                    case -965:
                        var v19 uint64 = 4294967295
                        _ = v19
                    }
                    switch "\n\t" {
                    case "a":
                        fmt.Println(p0)
                    default:
                        v20 := []int{33, 42, 297, 2147483647}
                        _ = v20
                    }
                }
            case "foo bar":
                p0 += 2
            default:
                var v21 uint64 = 491
                _ = v21
            }
        }
    }
    var v22 bool = (!(2 > 255))
    var v23 uint64 = uint64(p0)
    var v24 bool = (127 <= 0)
    _ = p0
    _ = v22
    _ = v23
    _ = v24
    return 0
}
func genFn1(p0 int64) uint16 {
    v25 := GenS1{F0: -1, F1: 4294967295}
    for fi26 := 0; fi26 < 9; fi26++ {
        v27 := GenS1{F0: 42, F1: 311}
        p0 = p0
        _ = fi26
        _ = v27
    }
    switch "\n\t" {
    case "hello":
        if (42 >= 834) {
            p0 += genFn0(32767)
            var v28 int64 = (p0 % ((p0) | 1))
            _ = v28
        } else {
            p0 = (p0 | (^(p0)))
            for fi29 := 0; fi29 < 13; fi29++ {
                for fi30 := 0; fi30 < 1; fi30++ {
                    var v31 int32 = 1
                    switch "a" {
                    case "":
                        v32 := []int{-1}
                        _ = v32
                    case "a":
                        fmt.Println(fi30)
                    }
                    _ = fi30
                    _ = v31
                }
                _ = fi29
            }
            v33 := []int{4294967295, 128, 1, 4294967295}
            _ = v33
        }
    case "":
        switch "a" {
        case "":
            if false {
                var v34 int16 = -696
                for fi35 := 0; fi35 < 16; fi35++ {
                    v36 := GenS1{F0: 127, F1: 130}
                    _ = fi35
                    _ = v36
                }
                p0 += (p0 >> ((4294967295) & 63))
                _ = v34
            } else {
                var v37 int8 = 42
                _ = v37
            }
        case "hello":
            var v38 string = "hello"
            _ = v38
        case "a":
            if true {
                p0 = p0
                var v39 uint8 = 133
                v40 := GenS2{F0: 255}
                _ = v39
                _ = v40
            }
        }
    case "\n\t":
        var v41 string = (("hello" + ("a" + ("a" + "\n\t"))) + "hello")
        _ = v41
    }
    for fi42 := 0; fi42 < 5; fi42++ {
        p0 = 1
        p0 = genFn0(754)
        _ = fi42
    }
    _ = p0
    _ = v25
    return 831
}

func main() {
    var v1 uint16 = 255
    var v2 bool = true
    var v3 uint16 = (v1 % ((v1) | 1))
    var v4 uint16 = 437
    v43 := GenS1{F0: 100, F1: 548}
    v44 := []int{109, -113}
    var v45 uint64 = 20
    _ = len(v44)
    v44 = append(v44, 0)
    v46 := GenS3{F0: (v3 | v3), F1: 927, F2: ""}
    switch ("" + "a") {
    case "":
        if false {
            v47 := GenS4{F0: -551}
            _ = v47
        }
    }
    v48 := GenS1{F0: -128, F1: 2}
    switch 128 {
    case 4294967295:
        switch -195 {
        case 32768:
            v49 := GenS1{F0: 0, F1: 0}
            _ = v49
        default:
            v50 := GenS1{F0: 42, F1: 18446744073709551615}
            _ = v50
        }
    }
    v3 = genFn1(9223372036854775807)
    v51 := GenS2{F0: 0}
    v3 = (^(v1))
    switch 2147483647 {
    case -520:
        if ((-2147483648 > 746) || ((1e10 + 1e10) >= float64(0.0))) {
            switch "a" {
            case "a":
                for fi52 := 0; fi52 < 19; fi52++ {
                    v4 = (^(v4))
                    v53 := GenS3{F0: genFn1(genFn0((v3 / (((^(v3))) | 1)))), F1: -1, F2: ("hello" + "foo bar")}
                    v44 = append(v44, (fi52 ^ fi52))
                    _ = fi52
                    _ = v53
                }
            case "\n\t":
                var v54 float64 = (0.1 - 0.0)
                _ = v54
            }
            fmt.Println(v1, v45)
            if (127 > 2) {
                var v55 string = ("a" + ((("foo bar" + "a") + ("" + "")) + "foo bar"))
                _ = v55
            }
        } else {
            if v2 {
                var v56 int32 = 530
                if v2 {
                    var v57 int8 = 42
                    _ = v57
                } else {
                    fmt.Println(v4, v3, v3)
                    var v58 int32 = 248
                    _ = v58
                }
                _ = v56
            }
            if (uint64(v1) > (^(v45))) {
                if ((127 <= 255) || v2) {
                    for fi59 := 0; fi59 < 14; fi59++ {
                        switch ("a" + (("hello" + ("\n\t" + "a")) + "foo bar")) {
                        case "hello":
                            for fi60 := 0; fi60 < 12; fi60++ {
                                v61 := GenS5{F0: (fi59 + (fi60 >> ((4294967295) & 63))), F1: -32768, F2: 2147483647}
                                v1 -= 384
                                _ = fi60
                                _ = v61
                            }
                        case "a":
                            if (("a" + "") < "a") {
                                var v62 int16 = 1
                                v2 = v2
                                v63 := GenS6{F0: (true || false), F1: fi59, F2: -824}
                                _ = v62
                                _ = v63
                            }
                        case "":
                            var v64 float32 = 3.14159
                            _ = v64
                        }
                        _ = fi59
                    }
                } else {
                    if len(v44) > 0 {
                        v44[0] = -957
                    }
                    switch "\n\t" {
                    case "hello":
                        for fi65 := 0; fi65 < 20; fi65++ {
                            v45 = 505
                            fmt.Println(v4)
                            v1 = v3
                            _ = fi65
                        }
                    case "foo bar":
                        if len(v44) >= 2 {
                            v44 = v44[1:len(v44)]
                        }
                    }
                }
                if false {
                    for fi66 := 0; fi66 < 3; fi66++ {
                        var v67 float32 = (-2.25 / 1.0)
                        fmt.Println(fi66, v67)
                        v4 = (v1 % (((v1 / ((genFn1(9223372036854775807)) | 1))) | 1))
                        _ = fi66
                        _ = v67
                    }
                }
            } else {
                v44 = append(v44, 42)
                v44 = append(v44, 9223372036854775807)
                switch "hello" {
                case "foo bar":
                    v2 = true
                case "a":
                    v68 := GenS7{F0: uint64(v3)}
                    _ = v68
                }
            }
            fmt.Println(v1)
        }
    default:
        if v2 {
            v69 := GenS1{F0: -128, F1: 543}
            v45 += v45
            _ = v69
        } else {
            if len(v44) >= 2 {
                v44 = v44[1:len(v44)]
            }
            var v70 int32 = 42
            v2 = ((0 == 128) && (!(!((0.1 / (1e10 / 1.0)) < (((3.14159 - 0.1) - float32(float64((1e10 / float32((3.14159 / float64((0.0 + float32(1e10))))))))) + -2.25)))))
            _ = v70
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v48
    _ = v51
}
