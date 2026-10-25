package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint32; F1 uint8; F2 uint }
type GenS7 struct { F0 int16; F1 uint; F2 int16 }
type GenS8 struct { F0 uint32; F1 uint32 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 string; F1 int64 }
type GenS11 struct { F0 int64; F1 uint8; F2 string }
type GenS12 struct { F0 bool }
func genFn0() bool {
    if true {
        if false {
            var v5 uint32 = 2
            v5 = (v5 * v5)
            _ = v5
        }
        for fi6 := 0; fi6 < 5; fi6++ {
            if true {
                var v7 float64 = -2.25
                _ = v7
            }
            var v8 int8 = -1
            var v9 bool = (!(v8 == int8(fi6)))
            _ = fi6
            _ = v8
            _ = v9
        }
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 float64 = (0.1 / float64(1e10))
            v11 -= (v11 / ((v11 - v11) / (v11 + v11)))
            _ = fi10
            _ = v11
        }
    } else {
        v12 := []int{}
        _ = v12
    }
    switch ("foo bar" + "hello") {
    case "":
        v13 := []int{4294967295, 770, 2147483647}
        _ = v13
    case "foo bar":
        for fi14 := 0; fi14 < 18; fi14++ {
            for fi15 := 0; fi15 < 7; fi15++ {
                if false {
                    if true {
                        v16 := GenS1{F0: 65535, F1: 103, F2: 576}
                        _ = v16
                    }
                    var v17 uint8 = 128
                    v18 := []int{(fi15 << ((255) & 63))}
                    if len(v18) > 0 {
                        v18[0] = (fi14 - fi15)
                    }
                    if false {
                        v17 = (v17 & 27)
                        var v19 uint32 = 420
                        _ = v19
                    } else {
                        fmt.Println(fi15, v17, v17)
                    }
                    var v20 uint16 = 460
                    v21 := GenS1{F0: 1, F1: 42, F2: (v20 << ((2) & 63))}
                    _ = v17
                    _ = v18
                    _ = v20
                    _ = v21
                }
                _ = fi15
            }
            _ = fi14
        }
    case "\n\t":
        for fi22 := 0; fi22 < 2; fi22++ {
            var v23 string = ("" + "hello")
            _ = fi22
            _ = v23
        }
    }
    var v24 uint64 = 18446744073709551615
    v24 -= v24
    for fi25 := 0; fi25 < 5; fi25++ {
        v26 := GenS2{F0: (0 != -1)}
        v24 = v24
        _ = fi25
        _ = v26
    }
    _ = v24
    return ((3.14159 / (3.14159 - (0.0 + 1.5))) != (1.5 + ((float32(0.0) / 1.5) * 1e10)))
}
func genFn1() bool {
    var v27 int64 = -9223372036854775808
    v27 = -1
    if (1 < 2) {
        switch 32768 {
        case 4294967295:
            v28 := []int{-1, 42}
            _ = v28
        case 991:
            v27 += -775
        case 2147483647:
            if false {
                if (128 != -948) {
                    v27 = 1
                    var v29 int64 = 42
                    v27 += (v27 & (v29 ^ v27))
                    _ = v29
                } else {
                    for fi30 := 0; fi30 < 18; fi30++ {
                        v31 := GenS3{F0: ("a" + (("foo bar" + ("hello" + "\n\t")) + ("foo bar" + (("foo bar" + "\n\t") + "\n\t"))))}
                        v32 := []int{1}
                        v32 = append(v32, fi30)
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                }
            } else {
                v33 := []int{}
                var v34 uint16 = 989
                _ = v33
                _ = v34
            }
        default:
            v35 := GenS1{F0: 941, F1: 42, F2: 65535}
            _ = v35
        }
        var v36 bool = (((-2.25 / 3.14159) * -2.25) >= (float64(((1.5 + 0.0) * 0.1)) / 0.1))
        v37 := []int{-2147483648, 9223372036854775807}
        _ = v36
        _ = v37
    } else {
        v38 := []int{}
        v27 = -9223372036854775808
        if len(v38) > 0 {
            v38[0] = 2147483647
        }
        _ = v38
    }
    v27 -= v27
    if false {
        fmt.Println(v27)
    }
    _ = v27
    return (!false)
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 109
    var v3 int64 = 1
    var v4 int32 = -1
    fmt.Println(v3, v2, v4)
    for fi39 := 0; fi39 < 6; fi39++ {
        switch (-(fi39)) {
        case -9223372036854775808:
            v40 := GenS4{F0: -989, F1: (^(v4))}
            _ = v40
        case 127:
            if true {
                var v41 float32 = 0.1
                v42 := GenS5{F0: 42}
                _ = v41
                _ = v42
            } else {
                var v43 uint8 = 128
                v2 = v2
                if true {
                    for fi44 := 0; fi44 < 4; fi44++ {
                        v45 := GenS6{F0: 0, F1: (v43 - v43), F2: 1}
                        var v46 uint16 = 893
                        v47 := GenS6{F0: 2147483648, F1: v43, F2: 856}
                        _ = fi44
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                    if false {
                        v4 = (v4 / ((v4) | 1))
                    } else {
                        v48 := []int{(fi39 % ((9223372036854775807) | 1))}
                        if (1.5 >= 1e10) {
                            fmt.Println(fi39)
                            if (!true) {
                                v43 = (v43 + v43)
                                for fi49 := 0; fi49 < 5; fi49++ {
                                    var v50 int8 = -128
                                    if true {
                                        var v51 int8 = -1
                                        _ = v51
                                    } else {
                                        v50 = v2
                                        var v52 string = "a"
                                        fmt.Println(v3, v4)
                                        _ = v52
                                    }
                                    _ = fi49
                                    _ = v50
                                }
                            }
                        }
                        for fi53 := 0; fi53 < 1; fi53++ {
                            v1[(fi53 & 127)] = ("hello" + "hello")
                            if true {
                                switch ("a" + ("hello" + "a")) {
                                case "":
                                    v54 := GenS2{F0: (((1e10 + ((0.1 + 0.1) * 3.14159)) / 0.1) <= 1.5)}
                                    _ = v54
                                default:
                                    switch fi53 {
                                    case -732:
                                        var v55 uint32 = 0
                                        _ = v55
                                    case 2147483647:
                                        var v56 uint16 = 255
                                        if len(v48) >= 2 {
                                            v48 = v48[1:len(v48)]
                                        }
                                        v43 += (v43 + v43)
                                        _ = v56
                                    }
                                }
                                v43 -= v43
                                if ((255 <= 457) && true) {
                                    fmt.Println(v43, fi39, v2)
                                    v4 = (v4 + (v4 - (v4 << ((84) & 63))))
                                    v57 := GenS7{F0: int16(v43), F1: 4294967295, F2: 1}
                                    _ = v57
                                }
                                if ("\n\t" != "foo bar") {
                                    switch ("foo bar" + (("a" + "a") + "a")) {
                                    case "\n\t":
                                        switch "a" {
                                        case "foo bar":
                                            if (!true) {
                                                v43 = (^(v43))
                                            } else {
                                                var v58 uint8 = 9
                                                var v59 bool = ((int8(v58) < (v2 * -1)) && (false && false))
                                                _ = v58
                                                _ = v59
                                            }
                                        case "\n\t":
                                            v3 = (v3 % (((v3 - 9223372036854775807)) | 1))
                                        case "hello":
                                            fmt.Println(v43, v4, fi39)
                                        default:
                                            var v60 int64 = v3
                                            _ = v60
                                        }
                                    case "hello":
                                        var v61 uint8 = v43
                                        var v62 string = ("" + "")
                                        _ = v61
                                        _ = v62
                                    case "foo bar":
                                        switch fi39 {
                                        case -620:
                                            fmt.Println(v3, fi53)
                                        default:
                                            v3 = v3
                                            v63 := GenS8{F0: 777, F1: 42}
                                            _ = v63
                                        }
                                    default:
                                        v2 = 0
                                    }
                                    var v64 int = fi39
                                    v48 = append(v48, (fi53 + fi53))
                                    _ = v64
                                } else {
                                    v65 := GenS9{F0: 100}
                                    _ = v65
                                }
                                switch fi39 {
                                case -956:
                                    fmt.Println(v3)
                                case 4294967295:
                                    var v66 uint8 = (v43 / ((v43) | 1))
                                    _ = v66
                                case 2147483647:
                                    fmt.Println(fi53, v4, v43)
                                }
                            }
                            _ = fi53
                        }
                        _ = v48
                    }
                    var v67 int32 = 1
                    _ = v67
                } else {
                    var v68 int = 2147483647
                    _ = v68
                }
                _ = v43
            }
        case 32767:
            v2 = v2
        default:
            if (!(true && true)) {
                for fi69 := 0; fi69 < 14; fi69++ {
                    v70 := v1[fi39]
                    var v71 int16 = 1
                    _ = fi69
                    _ = v70
                    _ = v71
                }
            } else {
                v2 = 100
                var v72 float32 = 0.1
                _ = v72
            }
        }
        v3 = int64(v2)
        _ = fi39
    }
    fmt.Println(v3, v2, v4)
    if (316 == 1) {
        v73 := GenS10{F0: ("\n\t" + "hello"), F1: 9223372036854775807}
        v4 += v4
        var v74 float64 = 0.1
        _ = v73
        _ = v74
    } else {
        v75 := []int{}
        var v76 uint64 = 5
        switch 128 {
        case -1:
            fmt.Println(v2, v76, v3)
        case 538:
            switch "hello" {
            case "foo bar":
                for fi77 := 0; fi77 < 14; fi77++ {
                    var v78 float32 = ((float32(float64(1e10)) - 0.0) + 1.5)
                    _ = fi77
                    _ = v78
                }
            }
        case 324:
            for fi79 := 0; fi79 < 17; fi79++ {
                for fi80 := 0; fi80 < 5; fi80++ {
                    var v81 uint8 = 93
                    _ = fi80
                    _ = v81
                }
                _ = fi79
            }
        }
        _ = v75
        _ = v76
    }
    var v82 int32 = v4
    v83 := GenS5{F0: 32767}
    fmt.Println(v82)
    var v84 int8 = (^(v2))
    switch -2147483648 {
    case 2147483647:
        var v85 float64 = (1.5 / (1e10 * (3.14159 + float64((0.1 + (3.14159 + 1e10))))))
        _ = v85
    default:
        for fi86 := 0; fi86 < 19; fi86++ {
            v87 := GenS11{F0: v3, F1: 128, F2: "hello"}
            v88 := GenS12{F0: false}
            v84 = v84
            _ = fi86
            _ = v87
            _ = v88
        }
    }
    switch int(v2) {
    case 32768:
        var v89 uint16 = 255
        _ = v89
    default:
        var v90 int64 = 549
        _ = v90
    }
    var v91 int64 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v82:", v82)
    fmt.Println("v84:", v84)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v82
    _ = v83
    _ = v84
    _ = v91
}
