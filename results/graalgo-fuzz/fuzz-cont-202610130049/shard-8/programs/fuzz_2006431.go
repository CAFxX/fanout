package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int8; F1 int32 }
type GenS8 struct { F0 uint8; F1 uint32; F2 uint }
type GenS9 struct { F0 uint16; F1 uint64 }
type GenS10 struct { F0 uint32; F1 string; F2 string }
type GenS11 struct { F0 int8 }

func main() {
    var v1 string = "hello"
    v2 := make(map[string]int)
    var v3 uint32 = 2147483648
    v4 := []uint32{42, (^(v3)), v3, 0}
    v5 := GenS1{F0: 9223372036854775807, F1: 32767}
    switch (v1 + v1) {
    case "foo bar":
        if (2147483647 > 0) {
            if ((v3 >= v3) && ((true && (251 <= 65535)) && false)) {
                v3 += (^(v3))
                if (!(9223372036854775807 == 42)) {
                    var v6 bool = true
                    v3 -= v3
                    if (-33 >= 1) {
                        var v7 uint16 = 42
                        v6 = (-147 > 876)
                        _ = v7
                    }
                    _ = v6
                } else {
                    if (1 >= 42) {
                        for fi8 := 0; fi8 < 5; fi8++ {
                            fmt.Println(v3, v3)
                            _ = fi8
                        }
                    } else {
                        switch v1 {
                        case "a":
                            switch -863 {
                            case 4294967295:
                                var v9 uint16 = 1
                                _ = v9
                            case 1:
                                var v10 int32 = 42
                                _ = v10
                            case -44:
                                v1 = "hello"
                            default:
                                switch v1 {
                                case "hello":
                                    v1 = "foo bar"
                                case "":
                                    v3 = v3
                                case "foo bar":
                                    for fi11 := 0; fi11 < 19; fi11++ {
                                        for fi12 := 0; fi12 < 5; fi12++ {
                                            switch (fi12 << ((65535) & 63)) {
                                            case -1:
                                                switch (((v1 + "") + v1) + v1) {
                                                case "\n\t":
                                                    v3 = (v3 & v3)
                                                case "":
                                                    var v13 int8 = 115
                                                    _ = v13
                                                case "foo bar":
                                                    v3 = 485
                                                    v1 = v1
                                                    v3 -= (v3 ^ v3)
                                                    if false {
                                                        v14 := GenS1{F0: (fi11 | fi11), F1: 1}
                                                        v3 = 1
                                                        _ = v14
                                                    }
                                                    if (!(4294967295 <= 0)) {
                                                        if (!true) {
                                                            fmt.Println(fi11, v1)
                                                            if (2 < 0) {
                                                                v15 := v2[""]
                                                                var v16 float64 = 3.14159
                                                                _ = v15
                                                                _ = v16
                                                            }
                                                        } else {
                                                            var v17 string = v1
                                                            if false {
                                                                var v18 uint16 = 255
                                                                v19 := GenS1{F0: (fi11 << ((0) & 63)), F1: 0}
                                                                if false {
                                                                    switch (^(fi12)) {
                                                                    case 32768:
                                                                        v3 = v3
                                                                        var v20 float32 = (1.5 - 0.0)
                                                                        var v21 uint8 = 22
                                                                        var v22 string = (("\n\t" + v17) + ("foo bar" + v1))
                                                                        _ = v20
                                                                        _ = v21
                                                                        _ = v22
                                                                    case 524:
                                                                        v18 = (^(v18))
                                                                    case 9223372036854775807:
                                                                        var v23 uint8 = 127
                                                                        _ = v23
                                                                    default:
                                                                        if false {
                                                                            v24 := v2[((("" + v17) + v1) + v1)]
                                                                            v3 += (v3 >> ((2) & 63))
                                                                            _ = v24
                                                                        } else {
                                                                            v18 = v18
                                                                        }
                                                                    }
                                                                    v25 := GenS1{F0: (fi11 * (fi12 | 2147483647)), F1: v18}
                                                                    _ = v25
                                                                }
                                                                _ = v18
                                                                _ = v19
                                                            }
                                                            _ = v17
                                                        }
                                                        fmt.Println(fi12, v1, v3)
                                                    } else {
                                                        var v26 int16 = 1
                                                        _ = v26
                                                    }
                                                    if (v1 >= v1) {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        v27 := GenS2{F0: v3}
                                                        v3 = 0
                                                        v28 := GenS2{F0: (^(v3))}
                                                        var v29 uint8 = 128
                                                        _ = v27
                                                        _ = v28
                                                        _ = v29
                                                    } else {
                                                        var v30 float64 = 0.1
                                                        v3 = (v3 - v3)
                                                        if len(v4) > 0 {
                                                            v4[0] = (v3 + (v3 / ((v3) | 1)))
                                                        }
                                                        _ = v30
                                                    }
                                                    v2[v1] = (fi11 % ((fi11) | 1))
                                                    if true {
                                                        var v31 int = 384
                                                        v31 = v31
                                                        _ = v31
                                                    }
                                                default:
                                                    switch v1 {
                                                    case "\n\t":
                                                        var v32 int8 = -128
                                                        _ = v32
                                                    case "hello":
                                                        var v33 int = fi11
                                                        _ = v33
                                                    case "a":
                                                        fmt.Println(fi12, v1)
                                                    }
                                                }
                                            case -9223372036854775808:
                                                v34 := GenS2{F0: 65535}
                                                _ = v34
                                            case 42:
                                                v3 -= v3
                                            }
                                            _ = fi12
                                        }
                                        var v35 float32 = ((0.0 - (1e10 - ((float32(-2.25) + ((1e10 - 3.14159) / (1.5 / (3.14159 / 1.0)))) + (1e10 * 0.1)))) - (1e10 + 0.0))
                                        _ = fi11
                                        _ = v35
                                    }
                                }
                            }
                        case "hello":
                            v3 = (v3 ^ (^(v3)))
                        case "foo bar":
                            var v36 bool = false
                            _ = v36
                        default:
                            if true {
                                v1 = (v1 + (v1 + v1))
                            } else {
                                var v37 uint32 = (v3 + v3)
                                v4 = append(v4, (v3 << ((65535) & 63)))
                                _ = v37
                            }
                        }
                        v38 := GenS3{F0: true}
                        v1 = "\n\t"
                        _ = v38
                    }
                }
                var v39 uint = 42
                _ = v39
            } else {
                switch 105 {
                case -273:
                    v1 = v1
                case 333:
                    v40 := GenS2{F0: (^(v3))}
                    _ = v40
                case 127:
                    v41 := GenS4{F0: 2, F1: (^(v3))}
                    _ = v41
                }
                fmt.Println(v1)
            }
        } else {
            var v42 bool = (!true)
            if (1 != 42) {
                var v43 int16 = 1
                v4 = append(v4, v3)
                _ = v43
            }
            _ = v42
        }
    default:
        v44 := GenS5{F0: 506}
        _ = v44
    }
    if (-1 == 100) {
        var v45 int32 = 2147483647
        for fi46 := 0; fi46 < 9; fi46++ {
            var v47 uint64 = 713
            v48 := GenS6{F0: (v3 / ((v3) | 1))}
            _ = fi46
            _ = v47
            _ = v48
        }
        _ = v45
    } else {
        v49 := GenS7{F0: 1, F1: -371}
        var v50 int = -1
        _ = v49
        _ = v50
    }
    if true {
        switch 0 {
        case 32767:
            switch ("" + "foo bar") {
            case "a":
                v51 := GenS4{F0: 567, F1: (v3 - v3)}
                _ = v51
            case "\n\t":
                for fi52 := 0; fi52 < 3; fi52++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi52
                }
            case "hello":
                _ = len(v2)
            default:
                v3 -= (v3 + (v3 & v3))
            }
        case 128:
            if true {
                var v53 int8 = 0
                for fi54 := 0; fi54 < 1; fi54++ {
                    switch (v1 + "\n\t") {
                    case "\n\t":
                        var v55 bool = true
                        _ = v55
                    default:
                        for fi56 := 0; fi56 < 4; fi56++ {
                            if true {
                                v1 = v1
                                var v57 int32 = 983
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                var v58 string = v1
                                var v59 int = (fi56 - (-(fi54)))
                                var v60 string = (v1 + "\n\t")
                                if false {
                                    var v61 int64 = -1
                                    _ = v61
                                }
                                _ = v57
                                _ = v58
                                _ = v59
                                _ = v60
                            }
                            var v62 int64 = 42
                            _ = fi56
                            _ = v62
                        }
                    }
                    _ = fi54
                }
                fmt.Println(v3, v53)
                _ = v53
            } else {
                fmt.Println(v3)
                v1 = v1
            }
        case 42:
            if (!(0 < -9223372036854775808)) {
                var v63 uint16 = 0
                v1 = (v1 + v1)
                fmt.Println(v1)
                _ = v63
            }
        default:
            switch (v1 + "") {
            case "hello":
                for fi64 := 0; fi64 < 11; fi64++ {
                    for fi65 := 0; fi65 < 18; fi65++ {
                        _ = len(v4)
                        v1 = v1
                        _ = fi65
                    }
                    for fi66 := 0; fi66 < 7; fi66++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v67 uint32 = v3
                        if (false && true) {
                            switch fi64 {
                            case 2147483647:
                                v68 := GenS8{F0: 2, F1: (v3 & v67), F2: 2}
                                _ = v68
                            case 32768:
                                v69 := GenS9{F0: 1, F1: 643}
                                _ = v69
                            case 9223372036854775807:
                                var v70 int32 = 1
                                _ = v70
                            }
                            fmt.Println(v3, fi64, v1)
                        }
                        v71 := GenS8{F0: 170, F1: 42, F2: 878}
                        v67 = v67
                        _ = len(v4)
                        _ = fi66
                        _ = v67
                        _ = v71
                    }
                    _ = fi64
                }
            case "a":
                if len(v4) > 0 {
                    v4[0] = v3
                }
            default:
                var v72 uint32 = 65535
                _ = v72
            }
        }
        if (3.14159 <= 1.5) {
            v3 += 1
            var v73 float64 = (float64(1e10) + 1.5)
            _ = v73
        } else {
            v3 = (v3 / (((v3 ^ 2)) | 1))
            var v74 bool = true
            v1 = v1
            _ = v74
        }
        for fi75 := 0; fi75 < 15; fi75++ {
            v4 = append(v4, (v3 << ((255) & 63)))
            if false {
                for fi76 := 0; fi76 < 14; fi76++ {
                    var v77 int64 = 0
                    v78 := GenS10{F0: v3, F1: v1, F2: ""}
                    _ = fi76
                    _ = v77
                    _ = v78
                }
                if (!true) {
                    if (-1 != -583) {
                        for fi79 := 0; fi79 < 5; fi79++ {
                            v80 := v2["\n\t"]
                            _ = fi79
                            _ = v80
                        }
                        var v81 string = ((v1 + v1) + (v1 + "hello"))
                        _ = v81
                    }
                    if (0 <= 0) {
                        var v82 string = v1
                        _ = v82
                    } else {
                        switch ((v1 + ((v1 + v1) + v1)) + v1) {
                        case "hello":
                            var v83 float64 = (1e10 * 3.14159)
                            _ = v83
                        case "a":
                            _ = len(v4)
                        }
                        if true {
                            v2[((v1 + (v1 + v1)) + (("\n\t" + v1) + "\n\t"))] = (-(fi75))
                        }
                    }
                } else {
                    var v84 uint = 2
                    fmt.Println(v1, v1)
                    switch 1 {
                    case -401:
                        v85 := v2[""]
                        _ = v85
                    }
                    _ = v84
                }
                if false {
                    v86 := GenS1{F0: fi75, F1: 778}
                    for fi87 := 0; fi87 < 1; fi87++ {
                        var v88 uint = 742
                        v89 := GenS11{F0: -1}
                        _ = fi87
                        _ = v88
                        _ = v89
                    }
                    _ = v86
                }
            } else {
                if (-1 > 0) {
                    fmt.Println(v1, v1, v1)
                    v3 -= 1
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                var v90 uint64 = 0
                var v91 float32 = 1.5
                _ = v90
                _ = v91
            }
            v1 = v1
            _ = fi75
        }
    }
    v4 = append(v4, v3)
    fmt.Println(v3, v3)
    switch -769 {
    case -9223372036854775808:
        for fi92 := 0; fi92 < 1; fi92++ {
            if true {
                v93 := GenS3{F0: false}
                fmt.Println(fi92)
                _ = v93
            } else {
                v3 += v3
                delete(v2, v1)
            }
            fmt.Println(v3, v3, v3)
            v1 = v1
            _ = fi92
        }
    case -2147483648:
        var v94 int64 = -967
        _ = v94
    case 4294967295:
        for fi95 := 0; fi95 < 16; fi95++ {
            v1 = v1
            fmt.Println(v1, v3, v3)
            _ = fi95
        }
    default:
        switch v1 {
        case "\n\t":
            v1 = v1
        case "foo bar":
            var v96 uint8 = 47
            _ = v96
        case "hello":
            v2[(v1 + v1)] = 162
        }
    }
    v3 = (v3 % (((v3 ^ 834)) | 1))
    switch v1 {
    case "\n\t":
        var v97 uint16 = 32767
        _ = v97
    case "":
        var v98 bool = false
        _ = v98
    default:
        for fi99 := 0; fi99 < 1; fi99++ {
            fmt.Println(fi99, fi99, fi99)
            v3 = 2147483647
            if false {
                var v100 uint8 = 2
                _ = v100
            }
            _ = fi99
        }
    }
    if (-1 > 588) {
        var v101 uint8 = 0
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v102 bool = true
        _ = v101
        _ = v102
    } else {
        v103 := v2["a"]
        var v104 uint = 2
        v105 := v2[(v1 + ((v1 + ((((v1 + v1) + v1) + "hello") + "")) + (v1 + ("foo bar" + "foo bar"))))]
        _ = v103
        _ = v104
        _ = v105
    }
    for fi106 := 0; fi106 < 2; fi106++ {
        v107 := v2["hello"]
        v3 -= (v3 << ((4294967295) & 63))
        _ = fi106
        _ = v107
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
