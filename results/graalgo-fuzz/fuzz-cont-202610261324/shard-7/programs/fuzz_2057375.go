package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 int64 }
func genFn0() int64 {
    var v4 float64 = (1.5 - -2.25)
    var v5 string = (("foo bar" + "\n\t") + "")
    _ = v4
    _ = v5
    return 108
}

func main() {
    v1 := []string{"hello", ((("hello" + (("foo bar" + "\n\t") + ("" + ""))) + (("foo bar" + "hello") + "\n\t")) + "hello"), ("\n\t" + ("foo bar" + "")), (("\n\t" + "\n\t") + "foo bar")}
    v2 := make(map[uint32]int)
    var v3 uint32 = 494
    if (0.0 != (0.1 * (-2.25 / 0.1))) {
        switch 32767 {
        case -27:
            switch "\n\t" {
            case "":
                for fi6 := 0; fi6 < 11; fi6++ {
                    for fi7 := 0; fi7 < 17; fi7++ {
                        fmt.Println(fi6)
                        var v8 float64 = float64(-2.25)
                        _ = len(v1)
                        _ = fi7
                        _ = v8
                    }
                    if len(v1) > 0 {
                        v1[0] = ("\n\t" + ("a" + ""))
                    }
                    for fi9 := 0; fi9 < 2; fi9++ {
                        v3 = v3
                        fmt.Println(fi6, fi9)
                        var v10 int32 = 0
                        var v11 int64 = 929
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                    _ = fi6
                }
            case "a":
                if true {
                    v3 += v3
                } else {
                    v12 := GenS1{F0: 1}
                    var v13 float64 = (((1.5 + float64(float32(1.5))) + -2.25) + 0.0)
                    _ = v12
                    _ = v13
                }
            case "hello":
                if (((0.0 == 0.1) || (uint(v3) > 0)) || (true && ((false && (42 != 42)) || false))) {
                    v14 := GenS1{F0: -32768}
                    _ = v14
                } else {
                    v3 = v3
                }
            default:
                if false {
                    if (0 != 36) {
                        for fi15 := 0; fi15 < 16; fi15++ {
                            switch "a" {
                            case "hello":
                                if (870 >= 42) {
                                    v3 = v3
                                }
                            case "foo bar":
                                fmt.Println(v3)
                            default:
                                v16 := GenS2{F0: 18}
                                _ = v16
                            }
                            _ = fi15
                        }
                    }
                    if ((65535 < 255) && false) {
                        switch "" {
                        case "":
                            for fi17 := 0; fi17 < 19; fi17++ {
                                var v18 int16 = -644
                                _ = fi17
                                _ = v18
                            }
                        case "\n\t":
                            var v19 uint64 = 417
                            _ = v19
                        case "foo bar":
                            for fi20 := 0; fi20 < 20; fi20++ {
                                for fi21 := 0; fi21 < 3; fi21++ {
                                    switch ("hello" + "foo bar") {
                                    case "a":
                                        fmt.Println(v3)
                                    }
                                    v3 = (^(v3))
                                    _ = fi21
                                }
                                _ = fi20
                            }
                        }
                        if true {
                            var v22 float32 = float32(((1e10 / (1e10 - 3.14159)) * (0.0 * (1.5 * float64(0.1)))))
                            _ = v22
                        } else {
                            v3 -= 509
                            var v23 float32 = -2.25
                            switch "" {
                            case "a":
                                for fi24 := 0; fi24 < 16; fi24++ {
                                    var v25 uint64 = 42
                                    v1 = append(v1, (("" + "foo bar") + ("" + "")))
                                    _ = fi24
                                    _ = v25
                                }
                            case "\n\t":
                                switch int(v3) {
                                case 127:
                                    if len(v1) > 0 {
                                        v1[0] = "foo bar"
                                    }
                                case 42:
                                    fmt.Println(v3, v23)
                                default:
                                    fmt.Println(v23, v23)
                                }
                            case "":
                                for fi26 := 0; fi26 < 13; fi26++ {
                                    v27 := GenS2{F0: 255}
                                    var v28 uint64 = 239
                                    _ = fi26
                                    _ = v27
                                    _ = v28
                                }
                            default:
                                v29 := GenS3{F0: (v3 | v3), F1: true}
                                _ = v29
                            }
                            _ = v23
                        }
                    }
                    v3 -= v3
                } else {
                    fmt.Println(v3)
                    v30 := GenS3{F0: v3, F1: false}
                    v3 = v3
                    _ = v30
                }
            }
        case -701:
            v31 := GenS1{F0: -32768}
            _ = v31
        default:
            var v32 uint = 99
            _ = v32
        }
        v3 -= 2147483648
    } else {
        v3 += (v3 & v3)
        v33 := GenS2{F0: 127}
        _ = v33
    }
    if len(v1) > 0 {
        v1[0] = ((("" + ("\n\t" + (("a" + "foo bar") + ""))) + "a") + "")
    }
    for fi34 := 0; fi34 < 6; fi34++ {
        switch "hello" {
        case "":
            v3 -= v3
        }
        for fi35 := 0; fi35 < 5; fi35++ {
            if (!(false || true)) {
                if true {
                    fmt.Println(fi35, v3, fi35)
                } else {
                    _ = len(v2)
                }
            } else {
                v36 := GenS3{F0: v3, F1: ((true && ((1 == 0) && (755 == 0))) && (!false))}
                switch (fi35 >> ((1) & 63)) {
                case 32767:
                    v3 -= v3
                    v3 += v3
                case 32768:
                    v37 := GenS1{F0: 1000}
                    _ = v37
                case 0:
                    var v38 uint = 18446744073709551615
                    _ = v38
                }
                v39 := GenS4{F0: -9223372036854775808, F1: "hello"}
                _ = v36
                _ = v39
            }
            var v40 string = ""
            _ = fi35
            _ = v40
        }
        _ = fi34
    }
    for fi41 := 0; fi41 < 18; fi41++ {
        v42 := GenS2{F0: 228}
        if true {
            switch fi41 {
            case 894:
                v43 := GenS2{F0: 1}
                _ = v43
            case 2147483647:
                var v44 int = 4294967295
                _ = v44
            default:
                for fi45 := 0; fi45 < 4; fi45++ {
                    switch "foo bar" {
                    case "a":
                        _ = len(v2)
                    case "foo bar":
                        fmt.Println(fi41)
                        fmt.Println(fi45)
                        v3 += (v3 & (v3 / ((0) | 1)))
                    case "hello":
                        v3 -= (^(v3))
                    }
                    v46 := GenS1{F0: -191}
                    var v47 float32 = 3.14159
                    _ = fi45
                    _ = v46
                    _ = v47
                }
            }
            var v48 float32 = ((0.0 - (float32((3.14159 - -2.25)) + 1e10)) * 0.0)
            _ = v48
        } else {
            if false {
                switch "a" {
                case "":
                    v3 += 680
                case "foo bar":
                    if (!false) {
                        v3 += v3
                        fmt.Println(fi41)
                        if len(v1) > 0 {
                            v1[0] = "hello"
                        }
                    }
                case "hello":
                    v49 := GenS2{F0: 42}
                    _ = v49
                }
                fmt.Println(v3)
                for fi50 := 0; fi50 < 12; fi50++ {
                    _ = len(v2)
                    if (18446744073709551615 > 675) {
                        v51 := GenS5{F0: 42}
                        _ = v51
                    } else {
                        var v52 float32 = (((1e10 - (float32((float64((1.5 / (3.14159 + float32(-2.25)))) - (1e10 - 0.0))) + 1.5)) + 0.0) * float32(((0.0 / (1e10 * 3.14159)) / (float64(-2.25) + (1e10 * 1e10)))))
                        v53 := v2[v3]
                        var v54 uint8 = 21
                        var v55 int32 = -469
                        _ = v52
                        _ = v53
                        _ = v54
                        _ = v55
                    }
                    var v56 int64 = 0
                    _ = fi50
                    _ = v56
                }
            } else {
                fmt.Println(v3, v3)
                v3 = v3
                v3 = (v3 >> ((0) & 63))
            }
            switch (("\n\t" + "\n\t") + "\n\t") {
            case "foo bar":
                fmt.Println(v3)
            default:
                _ = len(v2)
            }
            v3 += 26
        }
        _ = fi41
        _ = v42
    }
    for fi57 := 0; fi57 < 11; fi57++ {
        for fi58 := 0; fi58 < 9; fi58++ {
            var v59 bool = (true && (1 != 2147483647))
            v59 = ("" < "a")
            switch "" {
            case "a":
                fmt.Println(v59, v59)
            default:
                v3 -= v3
            }
            _ = fi58
            _ = v59
        }
        var v60 float64 = (0.1 * 1.5)
        var v61 int = fi57
        _ = fi57
        _ = v60
        _ = v61
    }
    var v62 float64 = 0.0
    fmt.Println(v3, v62)
    fmt.Println(v62, v3)
    switch "a" {
    case "hello":
        var v63 int16 = 42
        _ = v63
    case "":
        v3 = 4294967295
    case "foo bar":
        v64 := v2[(v3 + (v3 >> ((0) & 63)))]
        _ = v64
    default:
        var v65 bool = true
        _ = v65
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v62
}
