package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 uint64; F2 string }
type GenS5 struct { F0 int16; F1 int32; F2 string }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 uint64; F1 uint; F2 int8 }
type GenS8 struct { F0 int16; F1 bool; F2 uint16 }

func main() {
    var v1 uint32 = 2
    var v2 int = -722
    var v3 uint = 0
    var v4 uint16 = 32767
    v5 := make(map[uint32]string)
    v6 := GenS1{F0: 42, F1: 42, F2: v4}
    var v7 float64 = 3.14159
    if false {
        switch (v2 + v2) {
        case 701:
            for fi8 := 0; fi8 < 1; fi8++ {
                if false {
                    var v9 uint = 42
                    v10 := GenS1{F0: 430, F1: 0, F2: 319}
                    v3 = (^(v9))
                    _ = v10
                    _ = v9
                } else {
                    v4 -= 42
                    fmt.Println(fi8, v4)
                }
                var v11 string = "a"
                _ = fi8
                _ = v11
            }
        case 1:
            fmt.Println(v1, v7)
        default:
            v7 -= 1.5
        }
        var v12 int16 = 1000
        var v13 uint = (v3 << ((2) & 63))
        _ = v12
        _ = v13
    }
    switch 32768 {
    case 32767:
        var v14 int8 = -1
        _ = v14
    case 151:
        for fi15 := 0; fi15 < 18; fi15++ {
            switch v2 {
            case -824:
                var v16 uint8 = 42
                _ = v16
            case 32767:
                switch ("\n\t" + "foo bar") {
                case "hello":
                    if false {
                        v7 = v7
                        v3 += v3
                    }
                }
            default:
                fmt.Println(v3)
            }
            fmt.Println(v1)
            v2 += -918
            _ = fi15
        }
    default:
        if (128 > 127) {
            fmt.Println(v2, v4, v4)
        } else {
            v17 := GenS2{F0: v1, F1: (float32(v7) == 3.14159)}
            v18 := GenS2{F0: v1, F1: true}
            fmt.Println(v2)
            _ = v17
            _ = v18
        }
    }
    var v19 uint64 = 341
    var v20 uint16 = uint16(v1)
    var v21 uint64 = (^(v19))
    v22 := GenS2{F0: (v1 * v1), F1: false}
    switch (v2 | v2) {
    case 9223372036854775807:
        v7 -= 1e10
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        v21 += uint64(v1)
        _ = fi23
    }
    switch -2147483648 {
    case -178:
        for fi24 := 0; fi24 < 17; fi24++ {
            if (255 < 127) {
                var v25 int32 = 0
                if (v1 == (v1 >> (((^(v3))) & 63))) {
                    v26 := GenS1{F0: 2, F1: -33, F2: (v20 * uint16(v1))}
                    v27 := GenS3{F0: 1}
                    if false {
                        for fi28 := 0; fi28 < 3; fi28++ {
                            if (true || false) {
                                v7 -= float64((((0.1 / (float32((float64(1.5) / v7)) * -2.25)) * -2.25) - float32(v7)))
                            } else {
                                if (!true) {
                                    if ((false && true) || false) {
                                        v29 := GenS3{F0: 0}
                                        var v30 int32 = (v25 >> ((v3) & 63))
                                        fmt.Println(v3, v4, v2)
                                        _ = v29
                                        _ = v30
                                    } else {
                                        fmt.Println(v4, v20)
                                    }
                                    switch 1 {
                                    case -2147483648:
                                        v19 = v21
                                    case 716:
                                        if (42 <= 244) {
                                            var v31 bool = false
                                            v3 -= 0
                                            _ = v31
                                        } else {
                                            var v32 string = ((("hello" + "\n\t") + "foo bar") + "\n\t")
                                            v33 := GenS2{F0: uint32(v4), F1: ((v25 <= v25) && false)}
                                            v25 = 42
                                            _ = v32
                                            _ = v33
                                        }
                                    case 128:
                                        v34 := v5[v1]
                                        v19 = v19
                                        fmt.Println(v2)
                                        _ = v34
                                    default:
                                        var v35 int64 = 693
                                        _ = v35
                                    }
                                } else {
                                    if (v1 >= 65535) {
                                        v36 := GenS4{F0: (v2 ^ -729), F1: uint64(v3), F2: ("foo bar" + "hello")}
                                        if false {
                                            v7 += ((v7 - v7) - v7)
                                            var v37 uint64 = 18446744073709551615
                                            v38 := GenS4{F0: 128, F1: (v37 & v37), F2: ("\n\t" + "foo bar")}
                                            _ = v37
                                            _ = v38
                                        } else {
                                            fmt.Println(v21, v20, v20)
                                            v39 := GenS1{F0: v19, F1: -1, F2: (v4 & v20)}
                                            if false {
                                                v40 := GenS4{F0: fi28, F1: 62, F2: "a"}
                                                _ = v40
                                            } else {
                                                switch "a" {
                                                case "\n\t":
                                                    v41 := GenS3{F0: 1000}
                                                    _ = v41
                                                case "hello":
                                                    if false {
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            var v42 uint32 = (v1 + (v1 / (((^(v1))) | 1)))
                                                            var v43 int32 = 0
                                                            switch "" {
                                                            case "a":
                                                                v44 := []int{127, 128}
                                                                var v45 float64 = (float64(0.1) / v7)
                                                                var v46 bool = (v3 != 696)
                                                                _ = v44
                                                                _ = v45
                                                                _ = v46
                                                            case "":
                                                                delete(v5, v42)
                                                            default:
                                                                v42 += (^(v1))
                                                            }
                                                            _ = v42
                                                            _ = v43
                                                        }
                                                        if (("\n\t" + "\n\t") >= "hello") {
                                                            v1 += 149
                                                        } else {
                                                            v3 += v3
                                                            v47 := GenS1{F0: 42, F1: 100, F2: 42}
                                                            var v48 uint8 = 1
                                                            _ = v47
                                                            _ = v48
                                                        }
                                                    } else {
                                                        v2 = (^(fi28))
                                                        var v49 float64 = (v7 - float64(1.5))
                                                        _ = v49
                                                    }
                                                default:
                                                    v25 += v25
                                                }
                                                var v50 float32 = 0.0
                                                _ = v50
                                            }
                                            v51 := []int{369, v2, fi24, -9223372036854775808}
                                            v52 := GenS5{F0: 42, F1: -1, F2: "hello"}
                                            fmt.Println(v20, v20, fi24)
                                            if (((3.14159 + v7) < -2.25) && (2147483647 > fi28)) {
                                                v53 := GenS2{F0: v1, F1: true}
                                                var v54 string = ("foo bar" + "hello")
                                                if false {
                                                    v4 = 1
                                                    v55 := GenS4{F0: v2, F1: v21, F2: v54}
                                                    v25 -= v25
                                                    _ = v55
                                                }
                                                _ = v53
                                                _ = v54
                                            }
                                            _ = v39
                                            _ = v51
                                            _ = v52
                                        }
                                        if true {
                                            var v56 int = 888
                                            if true {
                                                if (v1 >= 0) {
                                                    v7 = (1.5 * v7)
                                                    fmt.Println(fi24)
                                                } else {
                                                    v19 += v21
                                                }
                                            } else {
                                                if false {
                                                    switch 32767 {
                                                    case 128:
                                                        v1 -= (v1 & (v1 >> ((uint(v19)) & 63)))
                                                    default:
                                                        fmt.Println(v21, fi24)
                                                    }
                                                    var v57 uint32 = (v1 & v1)
                                                    var v58 float32 = 1.5
                                                    _ = v57
                                                    _ = v58
                                                }
                                                fmt.Println(v2, v1)
                                            }
                                            _ = v56
                                        }
                                        _ = v36
                                    } else {
                                        var v59 uint8 = 255
                                        _ = v59
                                    }
                                }
                                if (65535 > v20) {
                                    switch 2147483647 {
                                    case -9223372036854775808:
                                        var v60 bool = false
                                        _ = v60
                                    case 32767:
                                        var v61 uint = v3
                                        _ = v61
                                    }
                                    v62 := GenS5{F0: 32767, F1: 1, F2: "\n\t"}
                                    _ = v62
                                } else {
                                    v2 = 127
                                }
                                if false {
                                    v3 = 255
                                }
                                v63 := v5[v1]
                                _ = v63
                            }
                            v21 = (v19 | 394)
                            var v64 uint32 = (v1 >> ((v3) & 63))
                            var v65 uint16 = v20
                            if (true || false) {
                                v19 += v21
                                var v66 uint32 = v64
                                _ = v66
                            } else {
                                v67 := GenS1{F0: v19, F1: 87, F2: v4}
                                _ = v67
                            }
                            v68 := []int{fi28, fi28, fi24}
                            v69 := GenS5{F0: -1, F1: (-(v25)), F2: ("" + "hello")}
                            v70 := GenS5{F0: 1000, F1: (v25 | (v25 / (((-(v25))) | 1))), F2: ("foo bar" + "")}
                            delete(v5, 241)
                            _ = fi28
                            _ = v64
                            _ = v65
                            _ = v68
                            _ = v69
                            _ = v70
                        }
                        for fi71 := 0; fi71 < 18; fi71++ {
                            switch (("a" + "\n\t") + "a") {
                            case "a":
                                switch "" {
                                case "":
                                    var v72 int8 = 127
                                    _ = v72
                                case "hello":
                                    switch ("hello" + ("hello" + "hello")) {
                                    case "hello":
                                        switch "a" {
                                        case "":
                                            switch (^(v2)) {
                                            case 570:
                                                if true {
                                                    v73 := []int{0}
                                                    _ = v73
                                                } else {
                                                    if false {
                                                        v74 := GenS2{F0: v1, F1: true}
                                                        fmt.Println(v3, fi71, v7)
                                                        _ = v74
                                                    } else {
                                                        v75 := GenS1{F0: (v21 - (v19 >> (((v3 % (((v3 << (((^(v3))) & 63))) | 1))) & 63))), F1: -41, F2: (v4 * (v20 / ((v20) | 1)))}
                                                        if true {
                                                            var v76 float64 = v7
                                                            v77 := GenS6{F0: (v3 ^ (^(v3))), F1: v1}
                                                            var v78 float64 = (1.5 * v76)
                                                            _ = v76
                                                            _ = v77
                                                            _ = v78
                                                        } else {
                                                            _ = len(v5)
                                                        }
                                                        var v79 uint8 = 128
                                                        _ = v75
                                                        _ = v79
                                                    }
                                                }
                                            case 457:
                                                v80 := GenS7{F0: v21, F1: (^(v3)), F2: -72}
                                                _ = v80
                                            case -2147483648:
                                                v19 = (v19 * 0)
                                            }
                                        case "foo bar":
                                            var v81 int16 = 1
                                            _ = v81
                                        case "a":
                                            switch (("hello" + "\n\t") + "foo bar") {
                                            case "foo bar":
                                                switch 295 {
                                                case 127:
                                                    v5[4294967295] = "\n\t"
                                                }
                                            }
                                        }
                                    case "foo bar":
                                        v82 := []int{}
                                        switch "\n\t" {
                                        case "a":
                                            switch fi24 {
                                            case 128:
                                                fmt.Println(v2, v2)
                                            default:
                                                var v83 uint64 = (v19 + (v21 - v19))
                                                _ = v83
                                            }
                                        case "":
                                            fmt.Println(v21)
                                        }
                                        if false {
                                            v20 = v4
                                            v2 += (fi24 >> ((v3) & 63))
                                            switch "a" {
                                            case "":
                                                var v84 int16 = 32767
                                                _ = v84
                                            case "a":
                                                v7 += 0.0
                                            case "\n\t":
                                                v85 := GenS1{F0: 18446744073709551615, F1: -128, F2: 255}
                                                _ = v85
                                            default:
                                                if len(v82) > 0 {
                                                    v82[0] = -1
                                                }
                                                fmt.Println(v25, v7)
                                            }
                                        }
                                        var v86 float32 = 1.5
                                        _ = v82
                                        _ = v86
                                    }
                                default:
                                    v4 -= (v4 & 255)
                                }
                            case "hello":
                                switch int(v4) {
                                case -93:
                                    delete(v5, (v1 - 2))
                                }
                            default:
                                v87 := GenS8{F0: 0, F1: false, F2: uint16(v3)}
                                _ = v87
                            }
                            v19 = (^(v19))
                            _ = fi71
                        }
                    }
                    _ = v26
                    _ = v27
                }
                _ = v25
            }
            v1 = v1
            fmt.Println(fi24)
            _ = fi24
        }
    case -9223372036854775808:
        v88 := []int{(v2 % (((^(v2))) | 1)), v2, (^(v2))}
        _ = v88
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
