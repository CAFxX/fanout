package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 int; F1 int; F2 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32; F1 uint64; F2 int }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int8; F1 int64; F2 uint }
type GenS9 struct { F0 bool; F1 uint32 }
func genFn0(p0 uint8, p1 int64, p2 int64) uint32 {
    if true {
        fmt.Println(p0)
        if true {
            p0 -= (p0 >> ((65535) & 63))
        } else {
            v6 := GenS1{F0: 2147483647, F1: (p1 << ((18) & 63))}
            var v7 uint8 = (^(p0))
            if false {
                v8 := []int{-9223372036854775808, 32767, int(p1), 1}
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = v8
            }
            _ = v6
            _ = v7
        }
    } else {
        p2 = 277
        var v9 float32 = (float32(3.14159) - 3.14159)
        var v10 int16 = 32767
        _ = v10
        _ = v9
    }
    if (true && true) {
        var v11 uint8 = p0
        _ = v11
    } else {
        var v12 bool = false
        var v13 uint16 = 2
        _ = v12
        _ = v13
    }
    var v14 bool = (!true)
    var v15 int16 = -32768
    fmt.Println(p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 703
}
func genFn1(p0 uint) int8 {
    v16 := GenS1{F0: 42, F1: 42}
    fmt.Println(p0)
    p0 = 18446744073709551615
    _ = p0
    _ = v16
    return 127
}

func main() {
    v1 := []string{}
    var v2 uint32 = 42
    var v3 uint32 = 42
    v4 := make(map[string]string)
    var v5 uint = 1
    v1 = append(v1, "a")
    for fi17 := 0; fi17 < 7; fi17++ {
        v18 := GenS1{F0: -732, F1: 42}
        _ = fi17
        _ = v18
    }
    if true {
        v19 := GenS2{F0: -1}
        var v20 bool = true
        var v21 uint32 = (^(v3))
        _ = v19
        _ = v20
        _ = v21
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v22 int = 9223372036854775807
    var v23 int32 = 0
    var v24 float64 = 1.5
    for fi25 := 0; fi25 < 4; fi25++ {
        v2 = (v2 * v2)
        var v26 int64 = -445
        _ = fi25
        _ = v26
    }
    var v27 float32 = 3.14159
    fmt.Println(v23, v2, v2)
    for fi28 := 0; fi28 < 2; fi28++ {
        switch v22 {
        case 816:
            switch fi28 {
            case 468:
                for fi29 := 0; fi29 < 11; fi29++ {
                    v30 := GenS1{F0: (-(v23)), F1: -987}
                    if true {
                        var v31 int = (fi29 & 1)
                        if false {
                            v32 := GenS2{F0: (-(fi28))}
                            fmt.Println(v2, v5, v2)
                            v33 := GenS3{F0: 0, F1: 127}
                            _ = v32
                            _ = v33
                        } else {
                            v3 += v2
                            if false {
                                if true {
                                    v34 := GenS1{F0: v23, F1: 0}
                                    _ = v34
                                } else {
                                    v35 := GenS4{F0: (fi29 ^ fi28), F1: 9223372036854775807, F2: 1}
                                    var v36 bool = (550 > int32(v5))
                                    var v37 int64 = 1
                                    if true {
                                        var v38 uint64 = 18446744073709551615
                                        switch "a" {
                                        case "a":
                                            var v39 int64 = v37
                                            _ = v39
                                        case "foo bar":
                                            v22 -= v31
                                        default:
                                            if v36 {
                                                switch "" {
                                                case "":
                                                    var v40 bool = (!((v38 - v38) < (^(v38))))
                                                    _ = v40
                                                case "a":
                                                    var v41 string = "hello"
                                                    _ = v41
                                                case "foo bar":
                                                    v42 := GenS1{F0: int32(v37), F1: v37}
                                                    _ = v42
                                                default:
                                                    v3 += v3
                                                }
                                            } else {
                                                v43 := GenS2{F0: (^(v31))}
                                                fmt.Println(v31, v22, v22)
                                                var v44 uint32 = v3
                                                if ((((v5 + v5) > (^(v5))) || false) || v36) {
                                                    switch "" {
                                                    case "":
                                                        var v45 float64 = 1.5
                                                        _ = v45
                                                    case "a":
                                                        if len(v1) > 0 {
                                                            v1[0] = ((("hello" + ("" + "a")) + ("hello" + "")) + ("a" + "a"))
                                                        }
                                                    default:
                                                        v2 = (v44 | v3)
                                                    }
                                                } else {
                                                    var v46 int8 = 1
                                                    v47 := GenS5{F0: 128}
                                                    delete(v4, (("foo bar" + "foo bar") + "\n\t"))
                                                    _ = v46
                                                    _ = v47
                                                }
                                                if (-32768 < -593) {
                                                    if v36 {
                                                        v27 += v27
                                                    }
                                                    if true {
                                                        var v48 uint8 = 2
                                                        v1 = append(v1, "hello")
                                                        _ = v48
                                                    }
                                                } else {
                                                    switch (fi29 / ((fi29) | 1)) {
                                                    case -2147483648:
                                                        switch "" {
                                                        case "\n\t":
                                                            v36 = (0 >= -226)
                                                        case "":
                                                            v31 -= int(v38)
                                                        }
                                                    case 2147483647:
                                                        v49 := GenS6{F0: (v2 - (v2 & genFn0(128, v37, v37))), F1: 966, F2: 0}
                                                        _ = v49
                                                    case -394:
                                                        var v50 string = ("" + "\n\t")
                                                        _ = v50
                                                    }
                                                    var v51 uint64 = v38
                                                    _ = v51
                                                }
                                                v38 = (v38 * v38)
                                                _ = v43
                                                _ = v44
                                            }
                                        }
                                        switch "a" {
                                        case "a":
                                            fmt.Println(v22, v5)
                                        case "foo bar":
                                            v36 = ((v2 & 937) < 177)
                                        case "hello":
                                            fmt.Println(v38, v22, v23)
                                            switch ("" + ("\n\t" + "\n\t")) {
                                            case "hello":
                                                var v52 float32 = v27
                                                _ = v52
                                            case "foo bar":
                                                var v53 string = (("a" + "foo bar") + "\n\t")
                                                _ = v53
                                            }
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v38
                                    } else {
                                        if ((v2 % (((v3 / (((v3 & genFn0(42, int64(fi29), (v37 / ((v37) | 1))))) | 1))) | 1)) >= (v2 - (v2 ^ v3))) {
                                            var v54 uint = v5
                                            v55 := GenS7{F0: 255}
                                            if v36 {
                                                v23 = (v23 % (((v23 + 1)) | 1))
                                                fmt.Println(v22, v22, fi29)
                                            }
                                            _ = v54
                                            _ = v55
                                        } else {
                                            switch "a" {
                                            case "\n\t":
                                                v56 := v4["hello"]
                                                _ = v56
                                            }
                                            var v57 float64 = v24
                                            var v58 float64 = float64(((-2.25 - v27) + float32(((v57 / 1e10) - 1e10))))
                                            _ = v57
                                            _ = v58
                                        }
                                    }
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                }
                                switch (v22 & fi28) {
                                case 2147483647:
                                    switch "hello" {
                                    case "a":
                                        var v59 uint16 = 65535
                                        _ = v59
                                    }
                                case 163:
                                    v5 += 2
                                case 792:
                                    var v60 uint16 = 1
                                    _ = v60
                                }
                            } else {
                                switch fi28 {
                                case 32767:
                                    fmt.Println(v3, v31)
                                default:
                                    if true {
                                        if true {
                                            fmt.Println(v23, v3)
                                        }
                                    } else {
                                        var v61 uint32 = v3
                                        v22 = (fi29 + (-(v22)))
                                        fmt.Println(v5, fi29, v22)
                                        v5 = v5
                                        if (float32(3.14159) != ((v27 + (v27 * (v27 * 3.14159))) - 3.14159)) {
                                            v62 := GenS8{F0: -1, F1: -839, F2: v5}
                                            _ = v62
                                        } else {
                                            _ = len(v1)
                                            switch (v31 + fi29) {
                                            case 1:
                                                switch ("foo bar" + "hello") {
                                                case "hello":
                                                    v31 -= int(v23)
                                                case "a":
                                                    v23 += v23
                                                default:
                                                    var v63 bool = (42 == 0)
                                                    _ = v63
                                                }
                                            case -2147483648:
                                                v31 -= v22
                                            case 32767:
                                                v64 := GenS6{F0: 1, F1: 4294967295, F2: (v22 ^ 32767)}
                                                _ = v64
                                            }
                                            if true {
                                                fmt.Println(v22, v61, v2)
                                            } else {
                                                if true {
                                                    v3 = (^(v2))
                                                    v2 += 1
                                                    switch "hello" {
                                                    case "":
                                                        fmt.Println(v22)
                                                    }
                                                    fmt.Println(fi29)
                                                    v61 += v3
                                                    _ = len(v1)
                                                    var v65 float64 = v24
                                                    _ = v65
                                                }
                                            }
                                            var v66 uint8 = 146
                                            var v67 int16 = -242
                                            _ = v66
                                            _ = v67
                                        }
                                        v68 := GenS9{F0: (1 > int(v2)), F1: (v61 & (^(v2)))}
                                        _ = v61
                                        _ = v68
                                    }
                                }
                                v3 = genFn0(127, -1, -836)
                                if false {
                                    fmt.Println(v5, v24, v31)
                                    switch "" {
                                    case "":
                                        v27 = (((1e10 / 3.14159) * v27) * 3.14159)
                                    case "hello":
                                        delete(v4, "foo bar")
                                    default:
                                        var v69 uint8 = uint8(v5)
                                        _ = v69
                                    }
                                    var v70 int16 = -258
                                    _ = v70
                                }
                            }
                        }
                        _ = v31
                    }
                    var v71 int64 = 1
                    _ = fi29
                    _ = v30
                    _ = v71
                }
            case 128:
                v5 = (v5 % (((v5 | 731)) | 1))
            case 4294967295:
                var v72 float64 = 0.0
                _ = v72
            }
        case 127:
            var v73 int = v22
            _ = v73
        case 4294967295:
            v1 = append(v1, "foo bar")
        default:
            var v74 uint16 = 255
            _ = v74
        }
        v75 := GenS1{F0: (v23 % (((v23 - v23)) | 1)), F1: 9223372036854775807}
        delete(v4, "hello")
        _ = fi28
        _ = v75
    }
    var v76 uint16 = 255
    var v77 int = -172
    fmt.Println(v23, v76, v77)
    v23 += (^(v23))
    v3 = v2
    for fi78 := 0; fi78 < 15; fi78++ {
        switch v22 {
        case -1:
            switch -999 {
            case 1:
                fmt.Println(v77, v23)
            }
        case 127:
            v3 = (v3 ^ v3)
        case -9223372036854775808:
            if (42 <= 265) {
                v77 += -124
                v76 = 65535
            }
        }
        var v79 float64 = (v24 * v24)
        v2 = v3
        _ = fi78
        _ = v79
    }
    var v80 int32 = v23
    var v81 uint16 = (v76 % (((^(v76))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v76
    _ = v77
    _ = v80
    _ = v81
}
