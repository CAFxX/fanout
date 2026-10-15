package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 bool; F1 int8; F2 int16 }
type GenS5 struct { F0 int16; F1 uint; F2 uint32 }
type GenS6 struct { F0 int8 }
func genFn0(p0 int) int8 {
    fmt.Println(p0)
    v5 := GenS1{F0: 42}
    _ = p0
    _ = v5
    return -128
}
func genFn1(p0 string) string {
    v6 := []int{}
    if true {
        if false {
            p0 = ((p0 + p0) + ((p0 + "\n\t") + ((p0 + (("\n\t" + p0) + "\n\t")) + "")))
        }
    }
    fmt.Println(p0)
    if (false || true) {
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 float32 = (float32(1.5) - -2.25)
            v8 -= ((((-2.25 + v8) * v8) / v8) / v8)
            var v9 uint = 42
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        switch 1 {
        case 0:
            v6 = append(v6, 9223372036854775807)
        case -176:
            for fi10 := 0; fi10 < 7; fi10++ {
                v11 := GenS1{F0: 2147483647}
                _ = fi10
                _ = v11
            }
        default:
            p0 = p0
        }
        var v12 uint32 = 4294967295
        _ = v12
    }
    p0 = ""
    _ = p0
    _ = v6
    return p0
}

func main() {
    var v1 int64 = 0
    var v2 uint16 = 656
    var v3 int32 = 2147483647
    v4 := make(map[string]bool)
    for fi13 := 0; fi13 < 16; fi13++ {
        if (!((!(1 > 4294967295)) && true)) {
            fmt.Println(v1, fi13)
            var v14 int16 = 97
            if true {
                for fi15 := 0; fi15 < 4; fi15++ {
                    if ("hello" == "foo bar") {
                        if (255 <= 42) {
                            v16 := GenS1{F0: v3}
                            v3 = (v3 + (v3 ^ (v3 * (v3 % ((v3) | 1)))))
                            if (!false) {
                                v14 += v14
                                var v17 int16 = (v14 * (v14 / (((v14 | v14)) | 1)))
                                _ = v17
                            }
                            _ = v16
                        }
                        if true {
                            if true {
                                fmt.Println(v1)
                                switch (fi13 | 381) {
                                case 0:
                                    var v18 float64 = 1.5
                                    _ = v18
                                case 128:
                                    v19 := GenS1{F0: (v3 & -1)}
                                    _ = v19
                                default:
                                    v20 := v4[""]
                                    _ = v20
                                }
                            }
                            var v21 float32 = 3.14159
                            var v22 uint32 = 4294967295
                            _ = v21
                            _ = v22
                        }
                    }
                    _ = fi15
                }
                var v23 int16 = (v14 / (((v14 ^ 1000)) | 1))
                v14 -= v23
                _ = v23
            }
            _ = v14
        }
        _ = fi13
    }
    v24 := GenS1{F0: (^(v3))}
    var v25 uint8 = 127
    var v26 uint32 = 1
    v25 -= (^(v25))
    v27 := GenS2{F0: v2, F1: v2}
    switch "a" {
    case "hello":
        v2 = v2
    case "foo bar":
        switch (("hello" + ("foo bar" + "")) + "hello") {
        case "hello":
            fmt.Println(v2)
        case "":
            for fi28 := 0; fi28 < 4; fi28++ {
                var v29 float32 = 0.1
                switch -1 {
                case 1:
                    v3 = 2147483647
                }
                v3 -= v3
                _ = fi28
                _ = v29
            }
        case "\n\t":
            for fi30 := 0; fi30 < 16; fi30++ {
                v31 := v4["foo bar"]
                for fi32 := 0; fi32 < 17; fi32++ {
                    var v33 uint32 = v26
                    _ = fi32
                    _ = v33
                }
                _ = fi30
                _ = v31
            }
        default:
            v25 = 128
        }
    }
    var v34 string = "hello"
    for fi35 := 0; fi35 < 4; fi35++ {
        v36 := []int{fi35}
        if false {
            switch fi35 {
            case -1:
                switch (fi35 >> ((0) & 63)) {
                case -894:
                    if len(v36) >= 2 {
                        v36 = v36[1:len(v36)]
                    }
                case 748:
                    var v37 int64 = (v1 << ((1) & 63))
                    _ = v37
                case -562:
                    if len(v36) > 0 {
                        v36[0] = fi35
                    }
                default:
                    if ((v1 != v1) && (-32768 >= -1)) {
                        switch (fi35 * (fi35 * fi35)) {
                        case 4294967295:
                            v38 := GenS3{F0: false, F1: "\n\t"}
                            _ = v38
                        case -9223372036854775808:
                            v36 = append(v36, fi35)
                        case -2147483648:
                            v25 = (^(v25))
                        }
                        v39 := GenS2{F0: 455, F1: v2}
                        v40 := GenS4{F0: true, F1: 127, F2: 1000}
                        _ = v39
                        _ = v40
                    } else {
                        var v41 int32 = v3
                        if false {
                            v26 = v26
                        } else {
                            fmt.Println(fi35, v34)
                        }
                        _ = v41
                    }
                }
            default:
                var v42 uint64 = 717
                _ = v42
            }
            v1 += (v1 & v1)
            var v43 int = 128
            _ = v43
        } else {
            var v44 int = fi35
            _ = v44
        }
        var v45 float64 = (0.0 - 1.5)
        _ = fi35
        _ = v36
        _ = v45
    }
    var v46 float64 = 0.0
    v25 = 0
    var v47 bool = (false || false)
    fmt.Println(v26, v3, v47)
    var v48 float32 = (1.5 * 0.1)
    switch -137 {
    case -9223372036854775808:
        var v49 uint64 = 2
        _ = v49
    case -728:
        fmt.Println(v46, v2)
    case 0:
        if v47 {
            if (v34 == "hello") {
                switch -9223372036854775808 {
                case 32767:
                    v50 := GenS4{F0: (v1 <= (v1 / ((9223372036854775807) | 1))), F1: 100, F2: -32768}
                    _ = v50
                case 101:
                    fmt.Println(v48, v47, v2)
                }
                v51 := GenS1{F0: v3}
                for fi52 := 0; fi52 < 2; fi52++ {
                    fmt.Println(v48, v47)
                    for fi53 := 0; fi53 < 10; fi53++ {
                        var v54 uint = 18446744073709551615
                        v26 -= 0
                        _ = fi53
                        _ = v54
                    }
                    v55 := GenS1{F0: 0}
                    _ = fi52
                    _ = v55
                }
                _ = v51
            } else {
                var v56 float32 = (v48 * (v48 / (0.0 - (v48 * 0.1))))
                var v57 uint8 = 1
                _ = v56
                _ = v57
            }
        } else {
            switch 32768 {
            case 128:
                var v58 uint16 = 32767
                _ = v58
            case 441:
                var v59 int32 = v3
                _ = v59
            case 42:
                switch 32768 {
                case 127:
                    for fi60 := 0; fi60 < 14; fi60++ {
                        var v61 uint32 = (v26 * (v26 & v26))
                        _ = fi60
                        _ = v61
                    }
                case -9223372036854775808:
                    v62 := GenS5{F0: 798, F1: 65535, F2: (v26 >> ((255) & 63))}
                    _ = v62
                case 9223372036854775807:
                    v34 = v34
                default:
                    if (!(!(!(-1 < -1)))) {
                        v63 := v4["foo bar"]
                        if (42 == v2) {
                            var v64 string = "\n\t"
                            switch v34 {
                            case "foo bar":
                                for fi65 := 0; fi65 < 8; fi65++ {
                                    v66 := GenS4{F0: (933 < 113), F1: 100, F2: 994}
                                    v67 := GenS4{F0: (!v63), F1: 42, F2: 1000}
                                    _ = fi65
                                    _ = v66
                                    _ = v67
                                }
                            default:
                                v68 := GenS2{F0: v2, F1: 0}
                                _ = v68
                            }
                            if v47 {
                                var v69 uint = 0
                                _ = v69
                            } else {
                                v70 := v4["foo bar"]
                                if (!v63) {
                                    for fi71 := 0; fi71 < 9; fi71++ {
                                        v72 := GenS4{F0: (false && (0.1 < (v48 / (v48 - v48)))), F1: 127, F2: 42}
                                        if v70 {
                                            v73 := GenS5{F0: -32768, F1: 18446744073709551615, F2: (v26 / ((v26) | 1))}
                                            _ = v73
                                        }
                                        _ = fi71
                                        _ = v72
                                    }
                                    var v74 string = "\n\t"
                                    var v75 int8 = 0
                                    _ = v74
                                    _ = v75
                                }
                                _ = v70
                            }
                            _ = v64
                        } else {
                            var v76 string = ("" + v34)
                            v25 = (v25 - v25)
                            var v77 uint8 = (v25 - (v25 % ((v25) | 1)))
                            _ = v76
                            _ = v77
                        }
                        var v78 int8 = 42
                        _ = v63
                        _ = v78
                    } else {
                        fmt.Println(v1)
                    }
                }
            }
            v3 += v3
            fmt.Println(v2, v1, v47)
        }
    }
    v79 := GenS6{F0: 100}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v79
}
