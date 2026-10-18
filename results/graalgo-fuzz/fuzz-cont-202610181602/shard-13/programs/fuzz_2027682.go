package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int32 }
type GenS2 struct { F0 string; F1 int32; F2 uint16 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 int64; F1 uint64; F2 bool }
type GenS5 struct { F0 int8; F1 uint16; F2 string }
type GenS6 struct { F0 string; F1 uint32; F2 bool }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS9 struct { F0 uint32; F1 int }
func genFn0() uint64 {
    v7 := []int{-89, 128}
    v8 := GenS1{F0: ("foo bar" + "foo bar"), F1: "", F2: 0}
    switch "a" {
    case "hello":
        var v9 uint32 = 2
        _ = v9
    case "a":
        switch 0 {
        case -680:
            if ((2 >= 0) && false) {
                var v10 float32 = (1.5 * 1e10)
                if (2 > 4294967295) {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                } else {
                    v11 := GenS1{F0: "hello", F1: ("\n\t" + "\n\t"), F2: 1}
                    var v12 int64 = 1
                    for fi13 := 0; fi13 < 4; fi13++ {
                        v14 := GenS1{F0: ("" + ("hello" + "")), F1: "\n\t", F2: 0}
                        _ = fi13
                        _ = v14
                    }
                    _ = v11
                    _ = v12
                }
                _ = v10
            } else {
                if (true || true) {
                    var v15 float32 = 1.5
                    _ = v15
                } else {
                    v7 = append(v7, -71)
                    v16 := GenS1{F0: "foo bar", F1: "foo bar", F2: -1}
                    _ = v16
                }
                for fi17 := 0; fi17 < 19; fi17++ {
                    if len(v7) > 0 {
                        v7[0] = 0
                    }
                    var v18 uint = 42
                    _ = fi17
                    _ = v18
                }
            }
        }
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 float32 = 0.0
        _ = fi19
        _ = v20
    }
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := []int{9223372036854775807}
    var v2 int = 127
    var v3 bool = false
    var v4 uint8 = 42
    var v5 uint16 = 1
    v6 := make(map[int]string)
    v1 = append(v1, int(v4))
    if v3 {
        fmt.Println(v5, v5)
        v4 = (v4 << ((0) & 63))
        switch (("" + "a") + ("" + ("\n\t" + ("" + "a")))) {
        case "a":
            for fi21 := 0; fi21 < 7; fi21++ {
                v22 := v6[(^(v2))]
                var v23 float64 = 1.5
                _ = fi21
                _ = v22
                _ = v23
            }
        default:
            var v24 float32 = -2.25
            _ = v24
        }
    } else {
        var v25 int64 = 333
        _ = v25
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        switch -9223372036854775808 {
        case 127:
            if v3 {
                v2 += 424
                var v27 uint64 = 785
                _ = v27
            } else {
                delete(v6, int(v5))
            }
        case 0:
            for fi28 := 0; fi28 < 1; fi28++ {
                v3 = v3
                fmt.Println(v4, v2)
                var v29 uint = 255
                _ = fi28
                _ = v29
            }
        }
        _ = fi26
    }
    v30 := v6[v2]
    switch v2 {
    case -137:
        switch (^(v2)) {
        case 1:
            var v31 uint16 = v5
            _ = v31
        default:
            v32 := GenS2{F0: (v30 + ((v30 + v30) + ("" + ("" + v30)))), F1: -2147483648, F2: v5}
            _ = v32
        }
    case 32767:
        var v33 float64 = 0.0
        _ = v33
    case 42:
        v34 := GenS1{F0: v30, F1: v30, F2: 42}
        _ = v34
    }
    fmt.Println(v30, v3)
    switch ((v30 + "\n\t") + ("foo bar" + (v30 + v30))) {
    case "":
        v35 := v6[v2]
        _ = v35
    case "hello":
        var v36 uint64 = genFn0()
        _ = v36
    case "a":
        fmt.Println(v30, v30, v2)
    }
    if true {
        for fi37 := 0; fi37 < 14; fi37++ {
            for fi38 := 0; fi38 < 18; fi38++ {
                if (!(((2 <= 42) && v3) && (((v4 - v4) < (v4 - (v4 % (((v4 / ((128) | 1))) | 1)))) || (v3 && (-784 > -1))))) {
                    if v3 {
                        var v39 float64 = 1.5
                        _ = v39
                    } else {
                        switch "" {
                        case "hello":
                            switch 900 {
                            case -1:
                                var v40 int8 = 1
                                _ = v40
                            case -570:
                                if v3 {
                                    var v41 float32 = ((0.1 + (3.14159 / (1.5 - -2.25))) + (3.14159 - 0.0))
                                    v41 -= float32(1.5)
                                    if (false || v3) {
                                        v3 = (1 == -435)
                                        v42 := GenS2{F0: "hello", F1: -521, F2: v5}
                                        _ = v42
                                    } else {
                                        var v43 uint64 = 0
                                        if (!(!((v2 == 128) && (!v3)))) {
                                            v30 = "a"
                                            v44 := GenS2{F0: "a", F1: -263, F2: 2}
                                            if (v3 || v3) {
                                                switch v30 {
                                                case "":
                                                    var v45 uint32 = uint32(v4)
                                                    _ = v45
                                                }
                                                var v46 string = ((v30 + ("a" + v30)) + (v30 + v30))
                                                _ = v46
                                            }
                                            v47 := GenS3{F0: -1, F1: 188}
                                            fmt.Println(v30, fi37, v5)
                                            _ = v44
                                            _ = v47
                                        }
                                        v48 := GenS4{F0: 0, F1: (v43 + (v43 % ((genFn0()) | 1))), F2: (1 <= -86)}
                                        _ = v43
                                        _ = v48
                                    }
                                    _ = v41
                                }
                            case 9223372036854775807:
                                v4 += 42
                            default:
                                var v49 uint64 = genFn0()
                                _ = v49
                            }
                        case "foo bar":
                            v50 := v6[(fi37 << ((2) & 63))]
                            _ = v50
                        default:
                            var v51 bool = (!false)
                            _ = v51
                        }
                    }
                    var v52 string = (v30 + (v30 + v30))
                    var v53 uint32 = 65535
                    v54 := GenS1{F0: (("\n\t" + "") + v30), F1: v52, F2: int32(v4)}
                    v1 = append(v1, fi38)
                    _ = v52
                    _ = v53
                    _ = v54
                } else {
                    var v55 int = 1
                    var v56 uint16 = 691
                    _ = v55
                    _ = v56
                }
                _ = fi38
            }
            _ = fi37
        }
        var v57 uint8 = v4
        for fi58 := 0; fi58 < 16; fi58++ {
            if v3 {
                if len(v1) > 0 {
                    v1[0] = 0
                }
                for fi59 := 0; fi59 < 3; fi59++ {
                    var v60 float64 = 1e10
                    var v61 float32 = 0.1
                    _ = fi59
                    _ = v60
                    _ = v61
                }
                var v62 uint32 = 2147483648
                _ = v62
            } else {
                switch v2 {
                case 9223372036854775807:
                    v63 := GenS5{F0: 1, F1: 42, F2: v30}
                    _ = v63
                }
                fmt.Println(v3)
            }
            _ = fi58
        }
        _ = v57
    } else {
        fmt.Println(v4, v5)
    }
    if (1 >= 0) {
        v2 -= 0
        for fi64 := 0; fi64 < 19; fi64++ {
            if (-2.25 >= -2.25) {
                v30 = ("hello" + v30)
            } else {
                fmt.Println(v2, fi64)
                v65 := GenS6{F0: "a", F1: 65535, F2: false}
                var v66 int64 = 0
                _ = v65
                _ = v66
            }
            v4 = v4
            var v67 int16 = -770
            _ = fi64
            _ = v67
        }
    } else {
        switch v2 {
        case 1:
            if false {
                _ = len(v6)
                if v3 {
                    if v3 {
                        var v68 uint16 = v5
                        _ = v68
                    } else {
                        switch "foo bar" {
                        case "":
                            switch (v2 * (-(v2))) {
                            case 32767:
                                for fi69 := 0; fi69 < 2; fi69++ {
                                    v70 := GenS1{F0: ((v30 + (v30 + "foo bar")) + v30), F1: "a", F2: -1}
                                    _ = fi69
                                    _ = v70
                                }
                            case 127:
                                if (42 > 42) {
                                    for fi71 := 0; fi71 < 4; fi71++ {
                                        for fi72 := 0; fi72 < 12; fi72++ {
                                            v1 = append(v1, (fi71 << ((42) & 63)))
                                            _ = fi72
                                        }
                                        v73 := GenS4{F0: 0, F1: 1, F2: true}
                                        switch v30 {
                                        case "foo bar":
                                            var v74 int = (v2 ^ v2)
                                            _ = v74
                                        case "hello":
                                            if (v5 == 678) {
                                                var v75 bool = (v3 || (v3 || (1 != 1000)))
                                                _ = v75
                                            }
                                        default:
                                            v76 := GenS7{F0: 323}
                                            _ = v76
                                        }
                                        _ = fi71
                                        _ = v73
                                    }
                                    fmt.Println(v4, v30)
                                } else {
                                    switch "\n\t" {
                                    case "a":
                                        if (1 != 0) {
                                            var v77 uint16 = (^(v5))
                                            var v78 uint = 65535
                                            _ = v77
                                            _ = v78
                                        }
                                    case "":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case "\n\t":
                                        if v3 {
                                            var v79 uint64 = 2
                                            _ = v79
                                        }
                                    default:
                                        v5 -= 0
                                    }
                                    fmt.Println(v4, v4)
                                    switch (v2 % ((127) | 1)) {
                                    case -30:
                                        v2 -= v2
                                    case 128:
                                        switch 32768 {
                                        case -187:
                                            _ = len(v6)
                                        }
                                    }
                                }
                            case 2147483647:
                                switch (v2 - (v2 ^ (v2 ^ v2))) {
                                case -847:
                                    switch v30 {
                                    case "\n\t":
                                        var v80 int64 = -9223372036854775808
                                        _ = v80
                                    case "foo bar":
                                        _ = len(v6)
                                    default:
                                        if false {
                                            var v81 float32 = ((0.0 / 3.14159) - -2.25)
                                            v5 = 255
                                            switch "a" {
                                            case "\n\t":
                                                var v82 int16 = -32768
                                                _ = v82
                                            case "a":
                                                for fi83 := 0; fi83 < 16; fi83++ {
                                                    var v84 int8 = -128
                                                    _ = fi83
                                                    _ = v84
                                                }
                                            case "hello":
                                                v81 += v81
                                            }
                                            _ = v81
                                        }
                                    }
                                default:
                                    v2 = v2
                                }
                            }
                        case "foo bar":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case "\n\t":
                            for fi85 := 0; fi85 < 6; fi85++ {
                                switch fi85 {
                                case 32768:
                                    var v86 float64 = 0.1
                                    _ = v86
                                default:
                                    var v87 int = 128
                                    _ = v87
                                }
                                _ = fi85
                            }
                        default:
                            var v88 uint8 = (v4 ^ v4)
                            _ = v88
                        }
                    }
                    v89 := GenS5{F0: 127, F1: v5, F2: (v30 + v30)}
                    _ = v89
                } else {
                    v2 -= -1
                    v90 := GenS8{F0: (v4 << ((949) & 63)), F1: -90, F2: 65535}
                    _ = v90
                }
                v91 := GenS9{F0: 0, F1: 150}
                _ = v91
            }
        case -208:
            for fi92 := 0; fi92 < 9; fi92++ {
                if true {
                    v93 := GenS4{F0: -9223372036854775808, F1: 22, F2: (0 < 101)}
                    v6[(-(fi92))] = (v30 + "foo bar")
                    _ = v93
                } else {
                    if ((uint64(v5) < 0) || (-10 > 127)) {
                        fmt.Println(fi92, v5)
                    } else {
                        var v94 float32 = (1.5 * -2.25)
                        _ = v94
                    }
                    if len(v1) > 0 {
                        v1[0] = v2
                    }
                }
                if len(v1) > 0 {
                    v1[0] = fi92
                }
                _ = fi92
            }
        }
        switch "" {
        case "\n\t":
            v4 = (^(v4))
        default:
            var v95 float64 = ((1.5 * float64(float32((0.0 + (1.5 + 1e10))))) * 0.1)
            _ = v95
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
