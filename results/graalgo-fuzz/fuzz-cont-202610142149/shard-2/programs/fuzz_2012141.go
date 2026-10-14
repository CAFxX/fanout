package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint64; F1 string; F2 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 int; F2 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64; F1 string; F2 int16 }

func main() {
    var v1 int64 = 42
    v2 := make(map[uint32]bool)
    var v3 uint64 = 0
    var v4 uint16 = 0
    v5 := []string{}
    v6 := []string{("a" + "hello")}
    v5 = append(v5, "\n\t")
    var v7 uint = 4294967295
    var v8 uint = v7
    if (1 >= 1) {
        v9 := GenS1{F0: uint8(v3), F1: ("\n\t" + "\n\t")}
        for fi10 := 0; fi10 < 15; fi10++ {
            v7 = v7
            _ = fi10
        }
        switch "foo bar" {
        case "\n\t":
            v11 := GenS1{F0: 247, F1: ("a" + (("\n\t" + "foo bar") + (("hello" + ("hello" + ("\n\t" + "hello"))) + "foo bar")))}
            _ = v11
        case "hello":
            v12 := GenS2{F0: (^(v3)), F1: ("" + ("" + "hello")), F2: -128}
            _ = v12
        case "a":
            for fi13 := 0; fi13 < 2; fi13++ {
                if true {
                    v2[1] = ((v1 >> (((^(v7))) & 63)) > v1)
                    var v14 int32 = 1
                    _ = v14
                } else {
                    var v15 int32 = 0
                    fmt.Println(v7)
                    fmt.Println(v7)
                    _ = v15
                }
                _ = fi13
            }
        default:
            for fi16 := 0; fi16 < 9; fi16++ {
                v3 = v3
                for fi17 := 0; fi17 < 6; fi17++ {
                    var v18 string = "a"
                    if true {
                        var v19 int8 = 100
                        _ = v19
                    }
                    if ((!true) && true) {
                        v4 = v4
                        if true {
                            fmt.Println(v18)
                            switch ("hello" + ((("" + ("hello" + v18)) + (v18 + v18)) + (v18 + ("a" + v18)))) {
                            case "":
                                if len(v6) > 0 {
                                    v6[0] = v18
                                }
                            case "foo bar":
                                var v20 uint8 = 128
                                _ = v20
                            }
                            v2[2] = (0 >= 23)
                            v21 := GenS3{F0: v4}
                            _ = v21
                        } else {
                            if false {
                                switch "\n\t" {
                                case "\n\t":
                                    fmt.Println(v4)
                                    v4 = (v4 | (v4 ^ v4))
                                default:
                                    v22 := GenS2{F0: v3, F1: (v18 + "hello"), F2: -108}
                                    _ = v22
                                }
                            } else {
                                if (!(42 < -140)) {
                                    var v23 float32 = 1e10
                                    var v24 bool = true
                                    var v25 float32 = ((((0.1 + float32(1.5)) - v23) / (v23 * 1.5)) + -2.25)
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                                v18 = "foo bar"
                                var v26 int = fi17
                                _ = v26
                            }
                            v7 = (^(v8))
                        }
                        if (!(1 < -128)) {
                            var v27 string = v18
                            _ = v27
                        }
                    }
                    _ = fi17
                    _ = v18
                }
                var v28 int16 = 32767
                _ = fi16
                _ = v28
            }
        }
        _ = v9
    }
    for fi29 := 0; fi29 < 12; fi29++ {
        if false {
            for fi30 := 0; fi30 < 17; fi30++ {
                v4 -= v4
                if false {
                    v3 += v3
                    var v31 string = ("hello" + (("a" + ("\n\t" + "hello")) + ""))
                    v32 := GenS1{F0: 1, F1: v31}
                    _ = v31
                    _ = v32
                }
                if (!false) {
                    v3 = (v3 - v3)
                } else {
                    v4 = v4
                    v3 = (v3 % ((v3) | 1))
                    v3 += v3
                }
                _ = fi30
            }
        }
        if false {
            fmt.Println(v3, v3, v4)
        } else {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            for fi33 := 0; fi33 < 4; fi33++ {
                if (true || false) {
                    if (!(true || false)) {
                        v34 := GenS4{F0: 1, F1: -1, F2: -1}
                        v7 -= (v7 | v7)
                        var v35 int = (-(fi33))
                        _ = v34
                        _ = v35
                    } else {
                        _ = len(v2)
                        var v36 int8 = 100
                        var v37 int8 = int8(fi33)
                        _ = v36
                        _ = v37
                    }
                    v38 := GenS3{F0: (^(v4))}
                    _ = v38
                }
                _ = fi33
            }
            var v39 uint16 = (v4 * uint16(v3))
            _ = v39
        }
        var v40 string = "\n\t"
        _ = fi29
        _ = v40
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v8 -= (v8 + v7)
    v1 -= 0
    if false {
        v7 = 42
    }
    var v41 uint16 = v4
    var v42 float64 = 1e10
    var v43 uint16 = v4
    if true {
        var v44 int16 = 32767
        switch ("hello" + "a") {
        case "hello":
            var v45 int64 = (v1 % ((-346) | 1))
            _ = v45
        case "a":
            switch 4294967295 {
            case -787:
                var v46 uint16 = v41
                _ = v46
            case 32767:
                if (1 >= 709) {
                    switch "hello" {
                    case "":
                        switch 127 {
                        case 128:
                            var v47 uint8 = 126
                            _ = v47
                        }
                    case "hello":
                        v2[91] = false
                    }
                }
            default:
                if false {
                    var v48 bool = true
                    _ = v48
                } else {
                    switch (("hello" + "foo bar") + ("\n\t" + (("hello" + "") + ("\n\t" + "hello")))) {
                    case "":
                        var v49 bool = false
                        _ = v49
                    }
                    if len(v5) > 0 {
                        v5[0] = "a"
                    }
                    for fi50 := 0; fi50 < 1; fi50++ {
                        for fi51 := 0; fi51 < 13; fi51++ {
                            var v52 string = ("a" + "\n\t")
                            _ = fi51
                            _ = v52
                        }
                        _ = fi50
                    }
                }
            }
        }
        _ = v44
    }
    for fi53 := 0; fi53 < 11; fi53++ {
        v3 = 18446744073709551615
        _ = fi53
    }
    fmt.Println(v1)
    if false {
        var v54 int64 = (^(v1))
        switch "" {
        case "foo bar":
            var v55 int = 1
            _ = v55
        case "hello":
            var v56 uint32 = 2147483647
            _ = v56
        case "\n\t":
            var v57 float64 = v42
            _ = v57
        }
        var v58 int = 4294967295
        _ = v54
        _ = v58
    } else {
        if true {
            if true {
                var v59 float32 = (-2.25 + 0.1)
                fmt.Println(v1)
                _ = v59
            } else {
                _ = len(v2)
                v60 := GenS5{F0: v43}
                _ = v60
            }
            switch ((((("\n\t" + "a") + "\n\t") + "foo bar") + "foo bar") + ("\n\t" + ((("foo bar" + "foo bar") + "a") + ("hello" + "")))) {
            case "":
                switch 42 {
                case 32767:
                    var v61 int32 = int32(v3)
                    _ = v61
                case 2147483647:
                    if (2147483648 != 1) {
                        for fi62 := 0; fi62 < 3; fi62++ {
                            v42 += (v42 - v42)
                            var v63 uint = (v8 + v8)
                            v64 := GenS6{F0: uint64(v63), F1: "foo bar", F2: -32768}
                            _ = fi62
                            _ = v63
                            _ = v64
                        }
                        delete(v2, 2147483647)
                        v41 -= (v43 ^ v4)
                    } else {
                        v41 += (^(v41))
                        v65 := GenS2{F0: (v3 << ((v7) & 63)), F1: ("" + ("foo bar" + ("" + ("hello" + ("" + ("foo bar" + "")))))), F2: 3}
                        if true {
                            v3 = (v3 * (v3 & (v3 ^ uint64(v7))))
                        }
                        _ = v65
                    }
                case 0:
                    fmt.Println(v3)
                }
            default:
                v66 := GenS6{F0: v3, F1: ("hello" + "foo bar"), F2: 1000}
                _ = v66
            }
        } else {
            v4 = 65535
        }
        for fi67 := 0; fi67 < 3; fi67++ {
            switch (fi67 | fi67) {
            case 32767:
                if false {
                    v8 = 2
                    var v68 string = "a"
                    _ = v68
                } else {
                    var v69 bool = true
                    if v69 {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    _ = v69
                }
            case 940:
                v1 = (v1 / (((v1 * int64(v41))) | 1))
            }
            var v70 string = ""
            _ = fi67
            _ = v70
        }
        var v71 uint = 255
        _ = v71
    }
    var v72 uint16 = (v41 / (((v43 * v43)) | 1))
    for fi73 := 0; fi73 < 12; fi73++ {
        for fi74 := 0; fi74 < 19; fi74++ {
            var v75 uint = v8
            _ = fi74
            _ = v75
        }
        v76 := v2[42]
        v4 = 899
        _ = fi73
        _ = v76
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v72:", v72)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v72
    _ = v8
}
