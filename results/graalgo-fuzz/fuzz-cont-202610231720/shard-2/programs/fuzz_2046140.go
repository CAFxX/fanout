package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint32; F1 int16; F2 string }
type GenS4 struct { F0 int; F1 string; F2 uint }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 uint32; F1 uint8 }
type GenS7 struct { F0 bool; F1 bool; F2 string }
type GenS8 struct { F0 int16; F1 uint64; F2 int32 }
func genFn0() int16 {
    var v5 int16 = 1
    if (!(!true)) {
        var v6 uint8 = 255
        v7 := GenS1{F0: (v5 & (v5 % (((v5 - v5)) | 1))), F1: 2}
        _ = v6
        _ = v7
    }
    v5 = 0
    _ = v5
    return (v5 / ((v5) | 1))
}

func main() {
    v1 := []uint32{1, 586}
    v2 := make(map[string]bool)
    var v3 bool = ("foo bar" == "hello")
    var v4 int32 = -302
    v3 = (!v3)
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 int8 = 92
        var v10 int8 = 42
        _ = fi8
        _ = v10
        _ = v9
    }
    if (1 <= 85) {
        for fi11 := 0; fi11 < 7; fi11++ {
            v12 := GenS1{F0: -32768, F1: 42}
            if v3 {
                delete(v2, "")
            }
            var v13 uint32 = 228
            _ = fi11
            _ = v12
            _ = v13
        }
        delete(v2, "\n\t")
    }
    var v14 bool = (42 > 281)
    switch "hello" {
    case "\n\t":
        if ("foo bar" == (((("\n\t" + ("\n\t" + (("foo bar" + ((("a" + "a") + "hello") + "\n\t")) + "foo bar"))) + "hello") + (("a" + "a") + ("\n\t" + (("hello" + "") + "")))) + ("\n\t" + "hello"))) {
            v3 = false
            for fi15 := 0; fi15 < 12; fi15++ {
                v16 := v2[("" + ("\n\t" + "hello"))]
                _ = fi15
                _ = v16
            }
            if len(v1) > 0 {
                v1[0] = 0
            }
        }
    case "hello":
        for fi17 := 0; fi17 < 20; fi17++ {
            v14 = (v4 <= -2147483648)
            if false {
                v14 = v3
            } else {
                v18 := GenS2{F0: (v4 * v4), F1: 701}
                fmt.Println(v14, fi17)
                v19 := GenS1{F0: 1, F1: 0}
                _ = v18
                _ = v19
            }
            _ = fi17
        }
    case "foo bar":
        fmt.Println(v14, v4)
    default:
        var v20 int8 = 127
        _ = v20
    }
    v3 = false
    for fi21 := 0; fi21 < 18; fi21++ {
        if ((!v14) && (128 < 0)) {
            v3 = (845 >= 2)
        } else {
            fmt.Println(v3, v14)
            for fi22 := 0; fi22 < 20; fi22++ {
                if false {
                    switch 128 {
                    case 9223372036854775807:
                        var v23 int64 = 1
                        v24 := GenS3{F0: 42, F1: -32768, F2: ""}
                        _ = v23
                        _ = v24
                    }
                    v14 = (!v3)
                    var v25 uint64 = 4294967295
                    _ = v25
                } else {
                    v3 = (true && v14)
                    if v3 {
                        if (65535 < 1) {
                            if len(v1) > 0 {
                                v1[0] = 2147483648
                            }
                            var v26 uint = 255
                            _ = v26
                        } else {
                            fmt.Println(v4, v4)
                        }
                        v27 := GenS2{F0: v4, F1: 544}
                        var v28 uint = 18446744073709551615
                        _ = v27
                        _ = v28
                    } else {
                        var v29 int64 = 1
                        _ = v29
                    }
                }
                _ = fi22
            }
            v4 -= (v4 - (v4 % ((0) | 1)))
        }
        for fi30 := 0; fi30 < 13; fi30++ {
            var v31 int32 = v4
            if (2 <= 42) {
                var v32 int8 = -1
                v31 = (v4 % ((2147483647) | 1))
                v4 -= (-(v4))
                _ = v32
            } else {
                if v3 {
                    var v33 float32 = (0.0 + -2.25)
                    if (!false) {
                        if false {
                            var v34 int8 = -1
                            _ = v34
                        } else {
                            var v35 int16 = -69
                            var v36 float64 = (3.14159 + (3.14159 / ((0.1 - 1.5) / ((((float64((v33 * 1e10)) - 3.14159) + float64(3.14159)) / (1e10 * 1.5)) - (float64(3.14159) + (0.1 + 1e10))))))
                            fmt.Println(v33, v31)
                            _ = v35
                            _ = v36
                        }
                        switch ("" + ((("" + "foo bar") + "foo bar") + ("\n\t" + ""))) {
                        case "foo bar":
                            var v37 int16 = -32768
                            _ = v37
                        case "hello":
                            v38 := GenS4{F0: -1, F1: "a", F2: 764}
                            _ = v38
                        case "":
                            switch (fi30 | -2147483648) {
                            case 1:
                                var v39 uint8 = 0
                                _ = v39
                            case -9223372036854775808:
                                v4 = (v31 + (v31 % (((v31 / (((v31 % (((^(v31))) | 1))) | 1))) | 1)))
                            }
                        }
                        v4 += 1
                    }
                    v40 := GenS5{F0: 4294967295, F1: fi30}
                    switch (("\n\t" + "") + (("hello" + "\n\t") + ("foo bar" + (("hello" + "hello") + ("" + ("\n\t" + "\n\t")))))) {
                    case "\n\t":
                        if (true || v3) {
                            if v3 {
                                fmt.Println(v31, v4)
                                fmt.Println(v3, v4, v31)
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v31 = (v31 % ((v4) | 1))
                                if false {
                                    var v41 uint = 65535
                                    if (117 > v41) {
                                        fmt.Println(v3, v31, v41)
                                        var v42 int16 = -32768
                                        if (v3 || v14) {
                                            _ = len(v1)
                                            switch (fi30 >> (((v41 - v41)) & 63)) {
                                            case 397:
                                                var v43 int = (fi30 - (fi30 / (((fi21 & fi21)) | 1)))
                                                _ = v43
                                            case -2147483648:
                                                v4 = v31
                                            }
                                        }
                                        _ = v42
                                    }
                                    var v44 float32 = v33
                                    _ = v41
                                    _ = v44
                                }
                            }
                        }
                    default:
                        fmt.Println(v3, fi21, v33)
                    }
                    v4 -= (v4 / ((1) | 1))
                    v31 = v31
                    _ = v33
                    _ = v40
                } else {
                    if len(v1) > 0 {
                        v1[0] = 4294967295
                    }
                    v3 = false
                }
                if len(v1) > 0 {
                    v1[0] = 277
                }
            }
            _ = fi30
            _ = v31
        }
        _ = fi21
    }
    for fi45 := 0; fi45 < 10; fi45++ {
        var v46 uint32 = 40
        v47 := GenS1{F0: 42, F1: 2}
        _ = fi45
        _ = v46
        _ = v47
    }
    if (2 == 255) {
        v14 = (32767 > 9223372036854775807)
        fmt.Println(v4)
        switch 9223372036854775807 {
        case -686:
            var v48 uint16 = 0
            _ = v48
        case 128:
            v49 := GenS1{F0: -1, F1: 42}
            _ = v49
        default:
            for fi50 := 0; fi50 < 8; fi50++ {
                if (437 > 80) {
                    v51 := GenS6{F0: 98, F1: 129}
                    _ = v51
                } else {
                    for fi52 := 0; fi52 < 15; fi52++ {
                        v53 := GenS7{F0: (!(v3 || v3)), F1: (622 == 65535), F2: "a"}
                        if false {
                            switch fi52 {
                            case 42:
                                v4 += v4
                            default:
                                if v14 {
                                    fmt.Println(fi52)
                                    v3 = v14
                                    v14 = v14
                                }
                            }
                            if (100 <= 118) {
                                var v54 int64 = -655
                                _ = len(v1)
                                _ = v54
                            }
                            var v55 uint16 = 32767
                            _ = v55
                        } else {
                            var v56 int = (^(fi52))
                            _ = v56
                        }
                        var v57 int = -9223372036854775808
                        _ = fi52
                        _ = v53
                        _ = v57
                    }
                    var v58 int8 = 127
                    v59 := GenS8{F0: -32768, F1: 18446744073709551615, F2: v4}
                    _ = v58
                    _ = v59
                }
                fmt.Println(v3, v14)
                if len(v1) > 0 {
                    v1[0] = 2
                }
                _ = fi50
            }
        }
    }
    var v60 uint = 4294967295
    fmt.Println(v4)
    switch "foo bar" {
    case "a":
        if false {
            fmt.Println(v14)
            switch -1 {
            case 42:
                if ((!(188 == 42)) || ((((-2.25 * 1.5) > 0.0) && (!false)) || v3)) {
                    var v61 int32 = v4
                    for fi62 := 0; fi62 < 9; fi62++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi62
                    }
                    _ = v61
                }
            case 554:
                fmt.Println(v60, v60)
            case 0:
                var v63 uint16 = 1
                _ = v63
            }
        } else {
            v64 := GenS4{F0: 32767, F1: "\n\t", F2: 18446744073709551615}
            fmt.Println(v14)
            _ = v64
        }
    case "hello":
        switch -612 {
        case 32767:
            for fi65 := 0; fi65 < 8; fi65++ {
                var v66 float32 = (1.5 + 0.0)
                for fi67 := 0; fi67 < 16; fi67++ {
                    fmt.Println(v60, fi67, v3)
                    v66 = 1.5
                    v2["a"] = ((-(v4)) < 0)
                    _ = fi67
                }
                _ = fi65
                _ = v66
            }
        }
    default:
        v60 += v60
    }
    v68 := GenS1{F0: genFn0(), F1: 42}
    v60 -= 18446744073709551615
    v14 = v3
    var v69 uint64 = 156
    var v70 uint16 = 53
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v60:", v60)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v60
    _ = v68
    _ = v69
    _ = v70
}
