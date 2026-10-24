package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 string; F1 uint; F2 uint16 }
type GenS4 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS5 struct { F0 string; F1 int16 }

func main() {
    var v1 bool = (2147483647 != 413)
    v2 := []uint32{0, 65535, 763, 0}
    var v3 int8 = 100
    v4 := make(map[string]int)
    for fi5 := 0; fi5 < 2; fi5++ {
        v1 = v1
        _ = fi5
    }
    if v1 {
        switch 51 {
        case 2147483647:
            for fi6 := 0; fi6 < 4; fi6++ {
                if v1 {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    for fi7 := 0; fi7 < 8; fi7++ {
                        v8 := GenS1{F0: 2, F1: 32767, F2: v3}
                        _ = fi7
                        _ = v8
                    }
                } else {
                    for fi9 := 0; fi9 < 16; fi9++ {
                        v1 = v1
                        var v10 string = "\n\t"
                        _ = fi9
                        _ = v10
                    }
                }
                _ = fi6
            }
        case 128:
            v4["hello"] = 127
        case 32768:
            v1 = (1e10 >= (1.5 + 1e10))
        default:
            for fi11 := 0; fi11 < 9; fi11++ {
                v12 := GenS1{F0: 74, F1: 32767, F2: v3}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                fmt.Println(fi11, v3)
                _ = fi11
                _ = v12
            }
        }
    }
    var v13 float64 = 0.1
    for fi14 := 0; fi14 < 2; fi14++ {
        if v1 {
            var v15 float64 = float64(3.14159)
            var v16 int = (fi14 % ((fi14) | 1))
            _ = len(v2)
            _ = v15
            _ = v16
        } else {
            switch "foo bar" {
            case "":
                for fi17 := 0; fi17 < 10; fi17++ {
                    switch ((("" + ("hello" + ("foo bar" + ("\n\t" + "hello")))) + "") + "") {
                    case "a":
                        fmt.Println(v1)
                    case "":
                        v1 = v1
                    }
                    var v18 float64 = 3.14159
                    _ = fi17
                    _ = v18
                }
            case "\n\t":
                fmt.Println(v1, v3)
            }
            v19 := GenS1{F0: 165, F1: 864, F2: v3}
            v3 = v3
            _ = v19
        }
        _ = fi14
    }
    if v1 {
        var v20 string = "hello"
        _ = v20
    }
    var v21 bool = true
    if false {
        switch ("" + "foo bar") {
        case "\n\t":
            v21 = true
        case "a":
            v22 := v4[("foo bar" + ("hello" + ((((("\n\t" + "") + ("a" + "")) + "foo bar") + "hello") + "\n\t")))]
            _ = v22
        case "foo bar":
            for fi23 := 0; fi23 < 9; fi23++ {
                var v24 float64 = v13
                if len(v2) > 0 {
                    v2[0] = 525
                }
                _ = fi23
                _ = v24
            }
        default:
            for fi25 := 0; fi25 < 12; fi25++ {
                if (1000 <= 1) {
                    _ = len(v2)
                    v26 := GenS1{F0: 0, F1: 308, F2: (v3 ^ (-(v3)))}
                    _ = v26
                } else {
                    _ = len(v4)
                    v13 += v13
                    var v27 int64 = 42
                    _ = v27
                }
                var v28 int64 = 0
                fmt.Println(fi25, v1, v3)
                _ = fi25
                _ = v28
            }
        }
        for fi29 := 0; fi29 < 15; fi29++ {
            var v30 uint64 = 0
            switch "hello" {
            case "":
                for fi31 := 0; fi31 < 14; fi31++ {
                    var v32 uint64 = v30
                    switch (-(fi29)) {
                    case 0:
                        v33 := GenS1{F0: 255, F1: 293, F2: (^(v3))}
                        _ = v33
                    }
                    v13 -= v13
                    _ = fi31
                    _ = v32
                }
            case "\n\t":
                v30 -= v30
            default:
                var v34 uint8 = 185
                _ = v34
            }
            for fi35 := 0; fi35 < 11; fi35++ {
                if v1 {
                    v36 := GenS1{F0: 0, F1: 32767, F2: v3}
                    if (int16(v30) == 397) {
                        v3 += (v3 & v3)
                        switch (-(fi35)) {
                        case 127:
                            v37 := GenS2{F0: -235, F1: -425}
                            _ = v37
                        case -807:
                            var v38 float64 = (((1e10 - 1e10) / v13) + v13)
                            _ = v38
                        case 4294967295:
                            fmt.Println(fi29, v13)
                        default:
                            v13 = float64(0.1)
                        }
                    } else {
                        if ((719 < 2) || v1) {
                            switch (("\n\t" + "\n\t") + ("" + ("" + (("\n\t" + ("" + "a")) + "")))) {
                            case "\n\t":
                                if ((float64(-2.25) / (((v13 - v13) / 0.1) / 1.5)) > 0.1) {
                                    var v39 bool = v1
                                    _ = v39
                                }
                            case "hello":
                                v13 += 0.1
                                if false {
                                    v40 := v4["\n\t"]
                                    fmt.Println(v1, fi35)
                                    v30 += 939
                                    _ = v40
                                } else {
                                    v4["\n\t"] = (fi29 ^ fi29)
                                    fmt.Println(fi29)
                                }
                            default:
                                _ = len(v2)
                            }
                            v30 = 42
                            switch fi35 {
                            case 42:
                                fmt.Println(fi29, v21)
                            default:
                                v21 = v1
                            }
                        }
                        v41 := GenS3{F0: "hello", F1: 0, F2: 255}
                        if v1 {
                            if true {
                                v3 -= v3
                                _ = len(v4)
                            }
                            v21 = v1
                            if v1 {
                                var v42 float32 = (3.14159 * -2.25)
                                v1 = v21
                                _ = v42
                            } else {
                                if (!((^(v30)) == v30)) {
                                    v30 -= 0
                                    v1 = ((73 < 736) && (127 >= 0))
                                } else {
                                    v43 := GenS3{F0: "", F1: 26, F2: 65535}
                                    v44 := GenS4{F0: (!v21), F1: 2, F2: 65535}
                                    var v45 int64 = 1
                                    fmt.Println(v3)
                                    switch "a" {
                                    case "foo bar":
                                        switch fi29 {
                                        case 32768:
                                            if v21 {
                                                v46 := GenS2{F0: -2147483648, F1: 0}
                                                fmt.Println(v45, fi29)
                                                v13 = 1.5
                                                _ = v46
                                            } else {
                                                var v47 uint = 255
                                                if (("\n\t" + ("" + (((((("foo bar" + "foo bar") + "a") + "hello") + "foo bar") + (("a" + "hello") + "")) + ("" + "foo bar")))) >= "foo bar") {
                                                    switch "" {
                                                    case "\n\t":
                                                        var v48 bool = true
                                                        _ = v48
                                                    }
                                                    var v49 uint32 = 588
                                                    var v50 int = (fi35 - (fi29 | (fi35 & (fi35 & (^(fi35))))))
                                                    _ = v49
                                                    _ = v50
                                                }
                                                _ = v47
                                            }
                                        default:
                                            fmt.Println(v21)
                                        }
                                    case "":
                                        switch "a" {
                                        case "a":
                                            if v21 {
                                                if v1 {
                                                    var v51 int32 = 692
                                                    var v52 int64 = (v45 | (v45 / ((v45) | 1)))
                                                    var v53 uint16 = 495
                                                    if true {
                                                        v30 += 0
                                                        switch "" {
                                                        case "a":
                                                            var v54 uint16 = 65535
                                                            fmt.Println(v54, v51, v30)
                                                            fmt.Println(v3, v54, v52)
                                                            _ = len(v2)
                                                            _ = v54
                                                        case "\n\t":
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        default:
                                                            v45 = -1
                                                        }
                                                    } else {
                                                        fmt.Println(v30, v30)
                                                        var v55 uint16 = (v53 / ((v53) | 1))
                                                        _ = v55
                                                    }
                                                    _ = len(v4)
                                                    if v1 {
                                                        var v56 uint16 = v53
                                                        switch fi35 {
                                                        case 0:
                                                            var v57 uint32 = 1
                                                            _ = v57
                                                        case -2147483648:
                                                            var v58 uint = 184
                                                            _ = v58
                                                        default:
                                                            if v1 {
                                                                if (945 <= 1) {
                                                                    var v59 uint8 = 147
                                                                    if true {
                                                                        if len(v2) > 0 {
                                                                            v2[0] = 2147483647
                                                                        }
                                                                    } else {
                                                                        var v60 int16 = 1
                                                                        var v61 uint16 = 31
                                                                        if v21 {
                                                                            fmt.Println(v21, v56, v61)
                                                                            v13 = 1e10
                                                                            var v62 uint8 = uint8(v30)
                                                                            var v63 int64 = (v52 | v52)
                                                                            switch fi29 {
                                                                            case -9223372036854775808:
                                                                                if v1 {
                                                                                    switch (^(fi35)) {
                                                                                    case -147:
                                                                                        v52 = -652
                                                                                    }
                                                                                    var v64 uint16 = v56
                                                                                    if (((int8(v62) < 127) && false) || ((v51 >= int32(fi35)) || ((^(v51)) >= (v51 << ((255) & 63))))) {
                                                                                        v65 := GenS5{F0: "a", F1: 42}
                                                                                        var v66 bool = false
                                                                                        v67 := v4[((("foo bar" + (("a" + "foo bar") + ("" + "foo bar"))) + "hello") + ("a" + "\n\t"))]
                                                                                        _ = v65
                                                                                        _ = v66
                                                                                        _ = v67
                                                                                    }
                                                                                    if v21 {
                                                                                        v59 = (v62 * v59)
                                                                                    }
                                                                                    v1 = v21
                                                                                    _ = v64
                                                                                }
                                                                            }
                                                                            fmt.Println(fi29, fi29, fi35)
                                                                            _ = v62
                                                                            _ = v63
                                                                        } else {
                                                                            v1 = (v45 < 804)
                                                                            var v68 uint64 = v30
                                                                            v59 = v59
                                                                            _ = v68
                                                                        }
                                                                        _ = v60
                                                                        _ = v61
                                                                    }
                                                                    _ = v59
                                                                } else {
                                                                    v45 = (v52 + (v52 / ((v52) | 1)))
                                                                }
                                                            }
                                                        }
                                                        _ = v56
                                                    }
                                                    if (v1 || true) {
                                                        fmt.Println(fi29, v3)
                                                    } else {
                                                        var v69 uint8 = 42
                                                        _ = v69
                                                    }
                                                    v70 := GenS1{F0: 42, F1: v53, F2: v3}
                                                    _ = v51
                                                    _ = v52
                                                    _ = v53
                                                    _ = v70
                                                }
                                            }
                                        case "\n\t":
                                            v13 = 1e10
                                        default:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                    default:
                                        var v71 float64 = v13
                                        _ = v71
                                    }
                                    _ = v43
                                    _ = v44
                                    _ = v45
                                }
                            }
                        } else {
                            var v72 uint64 = 18446744073709551615
                            _ = v72
                        }
                        _ = v41
                    }
                    v21 = v1
                    _ = v36
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(fi35)
                    case "a":
                        v4["a"] = 0
                    case "":
                        var v73 uint32 = 181
                        _ = v73
                    default:
                        _ = len(v4)
                    }
                    var v74 int = fi29
                    if v1 {
                        if v1 {
                            v21 = false
                        } else {
                            v3 = (v3 % ((-1) | 1))
                            fmt.Println(v3)
                            _ = len(v4)
                        }
                        if false {
                            fmt.Println(fi29)
                        }
                        if true {
                            v75 := GenS2{F0: -551, F1: -1}
                            _ = v75
                        } else {
                            var v76 int8 = (v3 ^ v3)
                            v76 -= (-(v3))
                            var v77 uint = 0
                            var v78 uint8 = 204
                            _ = v76
                            _ = v77
                            _ = v78
                        }
                    } else {
                        if (1 >= 0) {
                            v1 = v21
                            v21 = (!v21)
                            v13 -= (v13 / v13)
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            switch (fi35 % ((v74) | 1)) {
                            case 1:
                                var v79 uint32 = 0
                                _ = v79
                            }
                        }
                        var v80 float32 = 0.1
                        fmt.Println(v3, v74, v21)
                        switch "" {
                        case "foo bar":
                            _ = len(v2)
                        }
                        var v81 uint32 = 2147483648
                        _ = v80
                        _ = v81
                    }
                    _ = v74
                }
                _ = fi35
            }
            _ = fi29
            _ = v30
        }
        v13 = (float64((1.5 * 0.0)) - v13)
    }
    if true {
        var v82 uint16 = 924
        _ = v82
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
