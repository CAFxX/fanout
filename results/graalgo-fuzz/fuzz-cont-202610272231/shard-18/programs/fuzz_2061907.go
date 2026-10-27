package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 int32; F1 uint32; F2 bool }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 int8 }

func main() {
    v1 := []uint32{2}
    v2 := make(map[string]bool)
    var v3 float32 = -2.25
    v4 := make(map[uint32]string)
    var v5 uint = 586
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: -1, F1: 32767, F2: 9}
        var v8 float64 = 1e10
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 float32 = v3
            _ = fi9
            _ = v10
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    if false {
        fmt.Println(v3)
        fmt.Println(v3)
        for fi11 := 0; fi11 < 20; fi11++ {
            if ((9223372036854775807 > 0) || (!false)) {
                v3 = v3
                var v12 string = ("hello" + "a")
                _ = v12
            } else {
                for fi13 := 0; fi13 < 6; fi13++ {
                    var v14 bool = (0 >= 0)
                    _ = fi13
                    _ = v14
                }
                for fi15 := 0; fi15 < 11; fi15++ {
                    var v16 bool = (true || (1e10 >= -2.25))
                    if (!(1 <= 2)) {
                        v17 := GenS1{F0: -2147483648, F1: 427, F2: 42}
                        var v18 uint32 = 2
                        _ = v17
                        _ = v18
                    } else {
                        if (v3 == v3) {
                            v3 += v3
                            v2["foo bar"] = false
                        }
                    }
                    v3 = (1.5 * (((-2.25 * v3) / v3) - ((v3 - float32(-2.25)) - (1e10 * (v3 / v3)))))
                    _ = fi15
                    _ = v16
                }
                v5 = v5
            }
            _ = fi11
        }
    }
    var v19 int16 = 1000
    var v20 uint16 = 32767
    for fi21 := 0; fi21 < 6; fi21++ {
        for fi22 := 0; fi22 < 20; fi22++ {
            v3 = v3
            var v23 int16 = v19
            _ = fi22
            _ = v23
        }
        v5 -= (v5 + v5)
        if true {
            fmt.Println(v20, v20, v20)
            for fi24 := 0; fi24 < 15; fi24++ {
                if false {
                    v25 := v4[42]
                    var v26 uint = v5
                    _ = v25
                    _ = v26
                } else {
                    v5 = 907
                }
                v27 := GenS1{F0: -2147483648, F1: 255, F2: -117}
                var v28 uint = v5
                _ = fi24
                _ = v27
                _ = v28
            }
            var v29 int32 = -2147483648
            _ = v29
        } else {
            v3 += 1e10
            v20 = 2
            fmt.Println(v5, v19)
        }
        _ = fi21
    }
    if (int32(v20) >= -830) {
        v2[("" + "a")] = false
        v19 = (v19 + v19)
    } else {
        switch ("" + (((((("foo bar" + "a") + ("a" + "\n\t")) + "foo bar") + "\n\t") + "") + (("a" + "foo bar") + "\n\t"))) {
        case "a":
            var v30 uint32 = 0
            _ = v30
        case "\n\t":
            v31 := GenS1{F0: 0, F1: v20, F2: 42}
            _ = v31
        case "foo bar":
            v3 -= v3
        }
    }
    for fi32 := 0; fi32 < 14; fi32++ {
        for fi33 := 0; fi33 < 8; fi33++ {
            var v34 uint8 = 0
            v19 = 42
            v34 -= v34
            var v35 float32 = v3
            if true {
                v20 += 65535
                var v36 int8 = 81
                v37 := GenS1{F0: 1, F1: 255, F2: (v36 * v36)}
                if (v36 >= (v36 * (v36 - (v36 << ((v5) & 63))))) {
                    v5 -= (v5 / ((65535) | 1))
                    v5 += 0
                    v34 = v34
                    var v38 uint16 = uint16(v34)
                    v5 = v5
                    _ = v38
                } else {
                    v39 := GenS2{F0: 32767}
                    v40 := GenS2{F0: int(v34)}
                    v41 := GenS1{F0: 42, F1: v20, F2: v36}
                    v5 = v5
                    if false {
                        v20 -= 65535
                    } else {
                        fmt.Println(v35)
                    }
                    _ = v39
                    _ = v40
                    _ = v41
                }
                v19 = (v19 | v19)
                _ = v36
                _ = v37
            } else {
                v42 := GenS3{F0: ("\n\t" + ("" + "foo bar")), F1: 9223372036854775807}
                fmt.Println(v19)
                _ = v42
            }
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = fi32
    }
    if ((v20 >> (((^(v5))) & 63)) > (v20 * v20)) {
        switch 352 {
        case -414:
            var v43 uint8 = 2
            _ = v43
        case 32768:
            v44 := GenS4{F0: 0, F1: 42, F2: false}
            _ = v44
        case -223:
            v45 := GenS4{F0: 0, F1: 407, F2: (!false)}
            _ = v45
        default:
            v19 += v19
        }
    } else {
        switch ("hello" + "") {
        case "hello":
            for fi46 := 0; fi46 < 6; fi46++ {
                if (9223372036854775807 <= 42) {
                    if true {
                        var v47 int16 = 32767
                        var v48 bool = false
                        _ = v47
                        _ = v48
                    }
                    for fi49 := 0; fi49 < 20; fi49++ {
                        v50 := GenS5{F0: 986, F1: 761}
                        if true {
                            var v51 bool = false
                            v52 := GenS4{F0: 0, F1: 2147483647, F2: (!(1 == 100))}
                            _ = v51
                            _ = v52
                        }
                        _ = fi49
                        _ = v50
                    }
                    if (!false) {
                        if (0 >= 0) {
                            v3 += v3
                            if false {
                                var v53 string = ""
                                _ = v53
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = 42
                                }
                                var v54 float64 = 3.14159
                                _ = v54
                            }
                            if (!false) {
                                var v55 uint8 = 1
                                _ = v55
                            }
                        } else {
                            v3 -= v3
                        }
                        v20 -= v20
                    } else {
                        v19 = (^(v19))
                        if true {
                            var v56 string = ("hello" + ((("foo bar" + ((("\n\t" + (("\n\t" + "a") + "a")) + "\n\t") + ("foo bar" + ("hello" + ("" + "foo bar"))))) + ("foo bar" + "\n\t")) + ("\n\t" + "a")))
                            var v57 string = (v56 + (("\n\t" + v56) + "\n\t"))
                            _ = v56
                            _ = v57
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v58 uint = v5
                            _ = v58
                        }
                        var v59 int16 = v19
                        _ = v59
                    }
                } else {
                    var v60 int64 = -561
                    if false {
                        v61 := GenS5{F0: 273, F1: 1}
                        if (true || true) {
                            for fi62 := 0; fi62 < 2; fi62++ {
                                if true {
                                    if false {
                                        fmt.Println(v5)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    if (true || true) {
                                        if (-128 == 0) {
                                            v20 = (v20 << (((v5 + v5)) & 63))
                                            fmt.Println(fi62, fi46)
                                            var v63 int = (fi62 ^ fi46)
                                            _ = v63
                                        }
                                    }
                                    switch "" {
                                    case "foo bar":
                                        var v64 uint16 = (v20 % ((v20) | 1))
                                        _ = v64
                                    }
                                }
                                var v65 float64 = 1.5
                                var v66 int = 301
                                _ = fi62
                                _ = v65
                                _ = v66
                            }
                            _ = len(v4)
                            var v67 uint32 = 2147483647
                            _ = v67
                        } else {
                            for fi68 := 0; fi68 < 6; fi68++ {
                                fmt.Println(v19, v19)
                                _ = fi68
                            }
                        }
                        switch 1 {
                        case 32767:
                            switch "foo bar" {
                            case "\n\t":
                                if false {
                                    v3 = 1e10
                                    for fi69 := 0; fi69 < 17; fi69++ {
                                        var v70 float32 = v3
                                        _ = fi69
                                        _ = v70
                                    }
                                } else {
                                    if (575 != (v20 & v20)) {
                                        var v71 bool = false
                                        switch (("\n\t" + "hello") + ("foo bar" + "hello")) {
                                        case "":
                                            var v72 uint64 = 42
                                            _ = v72
                                        default:
                                            if ((^(v5)) <= (v5 & (v5 % ((v5) | 1)))) {
                                                v73 := GenS6{F0: -1}
                                                _ = v73
                                            } else {
                                                v74 := GenS6{F0: 127}
                                                v1 = append(v1, 2)
                                                _ = v74
                                            }
                                        }
                                        var v75 uint = 4294967295
                                        _ = v71
                                        _ = v75
                                    } else {
                                        v76 := GenS1{F0: 2147483647, F1: (v20 - (v20 % ((v20) | 1))), F2: -53}
                                        for fi77 := 0; fi77 < 5; fi77++ {
                                            switch "a" {
                                            case "":
                                                _ = len(v4)
                                            case "a":
                                                v2["foo bar"] = ((v60 * -1) >= (v60 - (v60 << ((v5) & 63))))
                                            }
                                            _ = fi77
                                        }
                                        _ = len(v2)
                                        _ = v76
                                    }
                                }
                            case "hello":
                                var v78 int = (-(fi46))
                                _ = v78
                            case "a":
                                fmt.Println(v20, v19, v60)
                            default:
                                fmt.Println(v20, fi46)
                            }
                        case -224:
                            fmt.Println(v5, fi46)
                        case 0:
                            v20 = (v20 + v20)
                        default:
                            switch (fi46 << (((^(v5))) & 63)) {
                            case -357:
                                var v79 bool = true
                                _ = v79
                            default:
                                var v80 uint = v5
                                _ = v80
                            }
                        }
                        _ = v61
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        for fi81 := 0; fi81 < 15; fi81++ {
                            if false {
                                v3 = v3
                                var v82 uint8 = 57
                                switch "" {
                                case "hello":
                                    if len(v1) > 0 {
                                        v1[0] = 2
                                    }
                                default:
                                    v83 := GenS4{F0: -1, F1: 42, F2: (v20 <= (v20 + v20))}
                                    _ = v83
                                }
                                _ = v82
                            }
                            switch 32767 {
                            case 127:
                                _ = len(v2)
                            default:
                                var v84 int32 = 0
                                _ = v84
                            }
                            _ = fi81
                        }
                    }
                    _ = v60
                }
                _ = fi46
            }
        case "foo bar":
            var v85 int16 = (v19 & 109)
            _ = v85
        default:
            fmt.Println(v19, v3, v20)
        }
    }
    var v86 int = 2147483647
    if true {
        switch "foo bar" {
        case "":
            for fi87 := 0; fi87 < 2; fi87++ {
                v3 -= v3
                var v88 float64 = (((1.5 * float64(((0.1 / 1e10) / v3))) * float64((float32(1e10) * (float32((3.14159 / 1.0)) / float32(float64(v3)))))) - 1.5)
                _ = fi87
                _ = v88
            }
        case "\n\t":
            var v89 int16 = v19
            _ = v89
        case "foo bar":
            var v90 string = "foo bar"
            _ = v90
        }
    }
    if (uint32(v20) >= 4294967295) {
        if len(v1) > 0 {
            v1[0] = 738
        }
        for fi91 := 0; fi91 < 17; fi91++ {
            v5 -= (v5 + v5)
            v86 = v86
            fmt.Println(v86, v86)
            _ = fi91
        }
    }
    if false {
        for fi92 := 0; fi92 < 15; fi92++ {
            v93 := GenS6{F0: -1}
            var v94 int64 = 0
            var v95 uint8 = 5
            _ = fi92
            _ = v93
            _ = v94
            _ = v95
        }
        var v96 uint64 = 0
        v20 -= v20
        _ = v96
    } else {
        var v97 uint32 = 2147483647
        fmt.Println(v86)
        _ = v97
    }
    v86 -= (v86 % (((v86 / ((int(v19)) | 1))) | 1))
    v86 = 128
    fmt.Println(v5, v86)
    for fi98 := 0; fi98 < 9; fi98++ {
        switch (fi98 + int(v19)) {
        case -9223372036854775808:
            v19 -= v19
        case 9223372036854775807:
            var v99 bool = ((true || (((false && false) && true) || (!(!false)))) && false)
            _ = v99
        case -2147483648:
            v86 = v86
        }
        v20 -= v20
        _ = fi98
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v86
}
