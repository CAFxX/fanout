package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int8 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 string; F1 uint32; F2 int8 }
type GenS5 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 bool; F1 uint64 }
type GenS8 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS9 struct { F0 uint32; F1 int64 }
type GenS10 struct { F0 int16 }
func genFn0() int8 {
    switch 9223372036854775807 {
    case -528:
        v4 := GenS1{F0: -1, F1: 308, F2: 0}
        _ = v4
    case -259:
        v5 := GenS2{F0: false, F1: 307}
        _ = v5
    }
    v6 := []int{-967, 283, 869, 42}
    var v7 uint = 4294967295
    var v8 int8 = 127
    _ = v6
    _ = v7
    _ = v8
    return 50
}
func genFn1() int64 {
    switch ("" + "\n\t") {
    case "a":
        v9 := []int{818, 4294967295}
        _ = v9
    case "\n\t":
        var v10 uint = 18446744073709551615
        _ = v10
    }
    v11 := []int{32767, -1, -49}
    if ((-2.25 >= -2.25) && false) {
        var v12 int8 = 1
        v11 = append(v11, 2147483647)
        var v13 uint16 = 42
        _ = v12
        _ = v13
    } else {
        var v14 float32 = (((3.14159 / (0.1 + 0.1)) + (0.1 / (float32(-2.25) - 3.14159))) + (((float32(3.14159) + 1.5) * (((float32(-2.25) * 0.1) * float32((0.0 - 1e10))) - 0.0)) / float32((0.0 + (3.14159 / 1.0)))))
        _ = len(v11)
        v14 -= ((((((-2.25 / (v14 - v14)) + v14) - 1.5) + ((1e10 / 3.14159) - 1.5)) * 1.5) + ((-2.25 - float32(1e10)) + ((0.0 * (-2.25 + 3.14159)) - v14)))
        _ = v14
    }
    switch "foo bar" {
    case "foo bar":
        var v15 int32 = -217
        _ = v15
    default:
        for fi16 := 0; fi16 < 10; fi16++ {
            var v17 float64 = 1.5
            var v18 bool = true
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    _ = v11
    return 1
}

func main() {
    v1 := []int{}
    v2 := make(map[int]uint32)
    v3 := make(map[string]bool)
    v19 := GenS1{F0: -151, F1: 4294967295, F2: -128}
    var v20 int = -886
    var v21 uint8 = 255
    if false {
        v21 -= v21
        if true {
            var v22 string = "hello"
            if false {
                v22 = (v22 + ("\n\t" + (((v22 + (v22 + v22)) + v22) + v22)))
            }
            v22 = v22
            _ = v22
        } else {
            fmt.Println(v21, v20)
            fmt.Println(v21, v20)
        }
        fmt.Println(v21, v21)
    }
    var v23 int32 = -2147483648
    switch v20 {
    case -57:
        v2[v20] = 2147483647
    case 32768:
        if false {
            var v24 float64 = (-2.25 - 1e10)
            for fi25 := 0; fi25 < 1; fi25++ {
                var v26 int16 = 32767
                _ = fi25
                _ = v26
            }
            if ((v20 / ((v20) | 1)) == v20) {
                for fi27 := 0; fi27 < 15; fi27++ {
                    v28 := GenS3{F0: 42, F1: (v21 >> ((255) & 63))}
                    var v29 uint64 = 0
                    switch v20 {
                    case -309:
                        v1 = append(v1, 1)
                    }
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            } else {
                v21 = (v21 & 1)
            }
            _ = v24
        }
    }
    var v30 float32 = ((1.5 - 1.5) * 1.5)
    if (false || false) {
        if (("" >= "hello") && false) {
            var v31 int32 = -1
            if (!true) {
                var v32 int32 = 42
                _ = v32
            } else {
                v33 := v3[("\n\t" + "\n\t")]
                v34 := GenS4{F0: "foo bar", F1: 2, F2: -128}
                if (((("hello" + "a") + "\n\t") + "a") >= "hello") {
                    var v35 int32 = 119
                    _ = v35
                } else {
                    v21 = v21
                    var v36 float64 = ((-2.25 - (-2.25 / 1e10)) * 1.5)
                    switch ("hello" + (("hello" + "hello") + "")) {
                    case "":
                        for fi37 := 0; fi37 < 4; fi37++ {
                            if (0 > -1) {
                                switch "" {
                                case "":
                                    var v38 float64 = ((3.14159 - 0.0) - v36)
                                    _ = v38
                                case "\n\t":
                                    if ((!v33) || (v33 || (42 > 4294967295))) {
                                        if ((597 == 1) && (-1 < 42)) {
                                            fmt.Println(v31, v21, v21)
                                        }
                                        v30 = (v30 / -2.25)
                                    } else {
                                        var v39 uint64 = 42
                                        switch ("foo bar" + "\n\t") {
                                        case "a":
                                            v31 = v23
                                        case "\n\t":
                                            _ = len(v1)
                                        }
                                        for fi40 := 0; fi40 < 7; fi40++ {
                                            var v41 float64 = v36
                                            _ = fi40
                                            _ = v41
                                        }
                                        _ = v39
                                    }
                                case "foo bar":
                                    if (65535 < 366) {
                                        switch 4294967295 {
                                        case 1:
                                            var v42 uint8 = (v21 * v21)
                                            _ = v42
                                        case 770:
                                            v43 := v2[(v20 * (v20 * v20))]
                                            _ = v43
                                        case 42:
                                            for fi44 := 0; fi44 < 1; fi44++ {
                                                fmt.Println(v36, v20, fi44)
                                                v36 = v36
                                                v45 := GenS5{F0: (v21 & v21), F1: 0, F2: 0}
                                                _ = fi44
                                                _ = v45
                                            }
                                        default:
                                            for fi46 := 0; fi46 < 16; fi46++ {
                                                v2[v20] = 911
                                                var v47 int32 = v31
                                                _ = fi46
                                                _ = v47
                                            }
                                        }
                                    }
                                }
                                for fi48 := 0; fi48 < 6; fi48++ {
                                    v31 += v31
                                    fmt.Println(v30)
                                    v1 = append(v1, (fi37 >> ((uint(v21)) & 63)))
                                    _ = fi48
                                }
                            }
                            v31 += (-(v23))
                            v33 = ((((v36 - ((float64(0.0) * 3.14159) + v36)) <= 1.5) || (v20 >= fi37)) && (!(v33 || v33)))
                            _ = fi37
                        }
                    }
                    _ = v36
                }
                _ = v33
                _ = v34
            }
            _ = v31
        } else {
            v23 = (v23 ^ v23)
        }
    } else {
        for fi49 := 0; fi49 < 3; fi49++ {
            delete(v3, "\n\t")
            for fi50 := 0; fi50 < 18; fi50++ {
                v51 := GenS4{F0: ((("a" + ("" + "foo bar")) + ("hello" + "\n\t")) + ("" + "")), F1: 1, F2: 48}
                v52 := GenS6{F0: 32767}
                _ = fi50
                _ = v51
                _ = v52
            }
            v23 = 585
            _ = fi49
        }
    }
    for fi53 := 0; fi53 < 15; fi53++ {
        switch fi53 {
        case 0:
            fmt.Println(v21)
        case 128:
            v20 = (v20 >> ((774) & 63))
        case -446:
            v54 := GenS7{F0: (18446744073709551615 != 4294967295), F1: 4294967295}
            _ = v54
        }
        switch ("\n\t" + "") {
        case "\n\t":
            v20 -= v20
        case "a":
            var v55 uint64 = 1
            _ = v55
        case "":
            if true {
                switch fi53 {
                case 127:
                    if (32767 <= 42) {
                        for fi56 := 0; fi56 < 13; fi56++ {
                            var v57 uint16 = 416
                            v58 := GenS8{F0: v21, F1: 179, F2: 0}
                            switch (fi56 % ((291) | 1)) {
                            case 4294967295:
                                fmt.Println(v20)
                            case 32767:
                                v23 = v23
                            case 127:
                                v59 := GenS8{F0: v21, F1: 852, F2: 0}
                                _ = v59
                            }
                            delete(v3, "\n\t")
                            var v60 uint = 18446744073709551615
                            _ = fi56
                            _ = v57
                            _ = v58
                            _ = v60
                        }
                    }
                case 42:
                    v23 -= 287
                default:
                    fmt.Println(v30, fi53)
                }
            }
        }
        for fi61 := 0; fi61 < 20; fi61++ {
            if (-775 < 32767) {
                var v62 int = -1
                var v63 uint16 = 1
                var v64 int32 = (v23 - 42)
                _ = v62
                _ = v63
                _ = v64
            } else {
                v30 -= (1.5 - (v30 - float32((1e10 * (0.0 - (0.1 - ((0.0 + 1e10) / ((-2.25 / 3.14159) - 1.5))))))))
                v30 = (float32((float64((3.14159 + v30)) - ((1e10 + ((1e10 / ((1.5 + 0.0) - 0.0)) + 1e10)) * ((-2.25 * 1.5) * (0.0 / 1.0))))) * v30)
                if (!((-539 < 1) || true)) {
                    if (-1 != 11) {
                        v65 := GenS1{F0: (-(fi61)), F1: (fi53 * 786), F2: -119}
                        _ = v65
                    }
                    if false {
                        switch -9223372036854775808 {
                        case 0:
                            var v66 int32 = 42
                            fmt.Println(fi61)
                            _ = v66
                        case -25:
                            v1 = append(v1, (fi53 / ((868) | 1)))
                        case 187:
                            if ((v21 + (v21 - 128)) != (v21 | v21)) {
                                v21 -= (v21 << ((42) & 63))
                            } else {
                                if false {
                                    switch fi53 {
                                    case -870:
                                        v67 := GenS9{F0: 1, F1: int64(v21)}
                                        _ = v67
                                    default:
                                        v23 = 883
                                    }
                                }
                                switch "a" {
                                case "":
                                    switch fi61 {
                                    case 42:
                                        fmt.Println(v20)
                                        v20 = fi61
                                        v21 = 193
                                    default:
                                        if false {
                                            v20 -= fi61
                                            fmt.Println(fi53, v20, v20)
                                            v23 -= 1
                                        } else {
                                            fmt.Println(v23, v30, fi61)
                                            var v68 uint = 442
                                            _ = v68
                                        }
                                    }
                                default:
                                    var v69 float64 = float64(v30)
                                    _ = v69
                                }
                            }
                        default:
                            fmt.Println(fi53)
                        }
                        var v70 string = "foo bar"
                        _ = v70
                    } else {
                        var v71 bool = ((2147483647 == 65535) && false)
                        var v72 uint16 = 378
                        if (1 != uint64(v21)) {
                            v20 += int(v72)
                        } else {
                            fmt.Println(v30, v20)
                        }
                        var v73 float64 = 0.1
                        var v74 uint8 = uint8(v72)
                        _ = v71
                        _ = v72
                        _ = v73
                        _ = v74
                    }
                }
            }
            if (((2 != 4294967295) && (255 >= 310)) || true) {
                fmt.Println(v21)
            }
            v23 = v23
            _ = fi61
        }
        _ = fi53
    }
    _ = len(v1)
    if (478 < 1) {
        v75 := GenS8{F0: v21, F1: 42, F2: 1}
        var v76 int64 = 261
        v77 := GenS10{F0: -32768}
        _ = v75
        _ = v76
        _ = v77
    } else {
        v21 = v21
        fmt.Println(v23, v21, v23)
        if len(v1) > 0 {
            v1[0] = (^(v20))
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v30
}
