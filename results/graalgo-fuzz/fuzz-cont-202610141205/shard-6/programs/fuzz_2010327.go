package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 bool; F2 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 uint; F2 uint }
type GenS6 struct { F0 int32; F1 int8; F2 int8 }
func genFn0(p0 int32, p1 uint64, p2 int32) bool {
    v7 := GenS1{F0: 255}
    v8 := []int{-2147483648, 127, 4294967295}
    for fi9 := 0; fi9 < 15; fi9++ {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        switch fi9 {
        case 127:
            var v10 uint = 42
            _ = v10
        case -705:
            _ = len(v8)
        case -220:
            for fi11 := 0; fi11 < 10; fi11++ {
                fmt.Println(p0, p1)
                _ = fi11
            }
        }
        _ = fi9
    }
    var v12 uint16 = 42
    fmt.Println(v12)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v7
    _ = v8
    return (65535 < 4294967295)
}
func genFn1() uint16 {
    var v13 uint16 = 892
    var v14 int16 = -829
    fmt.Println(v14, v13)
    v14 = 0
    if false {
        var v15 string = ((("a" + "a") + "") + "a")
        if (v15 <= v15) {
            v13 = v13
            fmt.Println(v15, v14, v14)
        } else {
            fmt.Println(v14)
        }
        _ = v15
    }
    _ = v13
    _ = v14
    return 42
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{}
    var v3 uint = 255
    var v4 int64 = -1
    var v5 int16 = -153
    var v6 uint64 = 42
    fmt.Println(v4, v6)
    v6 += v6
    var v16 float64 = (1.5 * 0.0)
    if true {
        if false {
            v6 -= (v6 / ((2) | 1))
        } else {
            if true {
                v17 := GenS1{F0: 127}
                _ = v17
            } else {
                _ = len(v1)
                v3 -= (v3 ^ (v3 - 2))
            }
        }
        var v18 float32 = 3.14159
        _ = v18
    }
    _ = len(v2)
    fmt.Println(v3)
    switch 9223372036854775807 {
    case 32767:
        for fi19 := 0; fi19 < 2; fi19++ {
            for fi20 := 0; fi20 < 13; fi20++ {
                v4 += (v4 << (((v3 / (((v3 % (((v3 + 0)) | 1))) | 1))) & 63))
                _ = fi20
            }
            if false {
                v21 := GenS2{F0: 42}
                var v22 string = ""
                fmt.Println(v16, v22)
                _ = v21
                _ = v22
            } else {
                v4 = v4
            }
            for fi23 := 0; fi23 < 18; fi23++ {
                switch (fi23 & fi19) {
                case 500:
                    if ((-(v4)) <= int64(fi19)) {
                        var v24 float32 = 0.1
                        v4 -= (v4 - (v4 >> ((v3) & 63)))
                        _ = v24
                    } else {
                        if (!(((v3 + v3) >= (v3 | (v3 | v3))) || (293 > 16))) {
                            v25 := GenS1{F0: 128}
                            var v26 bool = (!(-9223372036854775808 < 0))
                            _ = v25
                            _ = v26
                        }
                        var v27 int = 2147483647
                        _ = v27
                    }
                case -647:
                    if true {
                        var v28 int32 = 326
                        _ = len(v1)
                        _ = v28
                    } else {
                        var v29 float32 = 1.5
                        v30 := GenS3{F0: fi19, F1: (488 != genFn1()), F2: 1}
                        _ = v29
                        _ = v30
                    }
                }
                if false {
                    if false {
                        delete(v1, (("foo bar" + ("\n\t" + "\n\t")) + ""))
                        fmt.Println(v4, v16)
                    } else {
                        var v31 bool = false
                        _ = v31
                    }
                    switch ("a" + (("" + ("\n\t" + "\n\t")) + "foo bar")) {
                    case "hello":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "foo bar":
                        var v32 int16 = 317
                        _ = v32
                    case "":
                        fmt.Println(v4)
                    }
                }
                var v33 int8 = 0
                _ = fi23
                _ = v33
            }
            _ = fi19
        }
    case -579:
        var v34 uint8 = 2
        _ = v34
    case -1:
        switch 32767 {
        case 1:
            for fi35 := 0; fi35 < 6; fi35++ {
                if (1000 > v5) {
                    var v36 string = ("a" + ("hello" + "foo bar"))
                    _ = v36
                }
                var v37 float64 = float64(0.0)
                _ = fi35
                _ = v37
            }
        case -356:
            var v38 int32 = -367
            _ = v38
        }
    default:
        var v39 int8 = 42
        _ = v39
    }
    for fi40 := 0; fi40 < 20; fi40++ {
        fmt.Println(fi40, fi40)
        for fi41 := 0; fi41 < 3; fi41++ {
            if true {
                var v42 int64 = 9223372036854775807
                _ = v42
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v16 -= v16
                v43 := GenS1{F0: 1}
                _ = v43
            }
            var v44 uint = v3
            v3 = v44
            fmt.Println(v44, fi41, v5)
            switch (fi40 >> ((1) & 63)) {
            case 4294967295:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v45 int8 = 0
                var v46 uint = (v44 << ((v44) & 63))
                _ = v45
                _ = v46
            }
            _ = fi41
            _ = v44
        }
        _ = fi40
    }
    switch 9223372036854775807 {
    case -1:
        fmt.Println(v4)
    default:
        if (412 <= (^(v6))) {
            if (((("\n\t" + "") + "foo bar") + "\n\t") > ("\n\t" + "a")) {
                v47 := GenS1{F0: 42}
                for fi48 := 0; fi48 < 19; fi48++ {
                    v49 := GenS2{F0: 4294967295}
                    v2 = append(v2, (fi48 >> ((v3) & 63)))
                    v50 := GenS4{F0: (!false)}
                    _ = fi48
                    _ = v49
                    _ = v50
                }
                _ = v47
            }
            if true {
                v51 := v1["hello"]
                v52 := GenS3{F0: (v51 * (^(v51))), F1: (false || ((v51 >= -9223372036854775808) && (v16 != (v16 / v16)))), F2: 2}
                var v53 uint64 = (v6 + v6)
                _ = v51
                _ = v52
                _ = v53
            }
            for fi54 := 0; fi54 < 15; fi54++ {
                if (-353 <= -765) {
                    for fi55 := 0; fi55 < 14; fi55++ {
                        var v56 uint32 = 853
                        _ = fi55
                        _ = v56
                    }
                    fmt.Println(v6)
                    v57 := GenS2{F0: 65535}
                    _ = v57
                } else {
                    var v58 uint64 = (v6 << (((v3 / ((65535) | 1))) & 63))
                    _ = v58
                }
                if (true && true) {
                    var v59 uint16 = 1
                    if true {
                        var v60 uint32 = 2
                        v60 -= v60
                        _ = v60
                    }
                    _ = v59
                } else {
                    if (0.0 <= (-2.25 - 1e10)) {
                        var v61 bool = (false || (("foo bar" + "a") <= ("" + ("a" + "a"))))
                        _ = v61
                    } else {
                        if ((v3 | 0) < 992) {
                            v3 = (v3 & 2)
                            switch "" {
                            case "":
                                v62 := GenS3{F0: fi54, F1: (150 < 42), F2: 473}
                                _ = v62
                            }
                        }
                    }
                    var v63 int = -2147483648
                    v3 = v3
                    _ = v63
                }
                _ = fi54
            }
        } else {
            fmt.Println(v4, v3, v16)
        }
    }
    _ = len(v1)
    v4 += v4
    switch -9223372036854775808 {
    case -594:
        v64 := GenS5{F0: 127, F1: (v3 * (v3 + v3)), F2: (^(v3))}
        _ = v64
    case 722:
        for fi65 := 0; fi65 < 10; fi65++ {
            v6 = uint64(v3)
            v6 = 405
            _ = fi65
        }
    case -860:
        if (uint32(v6) > 1) {
            var v66 uint8 = 128
            _ = v66
        } else {
            var v67 bool = true
            for fi68 := 0; fi68 < 11; fi68++ {
                v3 = v3
                _ = len(v2)
                v6 = v6
                _ = fi68
            }
            _ = v67
        }
    default:
        delete(v1, ("hello" + "foo bar"))
    }
    v16 = (float64(0.0) / (1e10 / (-2.25 / (1e10 * v16))))
    if (true && true) {
        v5 = (v5 ^ (v5 - (v5 / (((v5 - v5)) | 1))))
        if (true && false) {
            v3 += v3
            switch "" {
            case "\n\t":
                v6 = (v6 - (v6 - (v6 ^ (v6 + (v6 ^ v6)))))
            case "":
                v3 += 0
            }
        } else {
            var v69 int8 = 0
            fmt.Println(v69)
            fmt.Println(v6, v69, v5)
            _ = v69
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi70 := 0; fi70 < 10; fi70++ {
            v2 = append(v2, (fi70 + (fi70 * fi70)))
            _ = fi70
        }
        if (814 < 1) {
            var v71 uint64 = v6
            _ = len(v1)
            for fi72 := 0; fi72 < 10; fi72++ {
                v3 = 2
                _ = fi72
            }
            _ = v71
        }
    }
    for fi73 := 0; fi73 < 19; fi73++ {
        if (179 < 1) {
            v74 := GenS5{F0: 128, F1: 507, F2: (v3 << ((v3) & 63))}
            fmt.Println(v4, v5)
            v6 += (v6 & v6)
            _ = v74
        }
        v75 := GenS6{F0: -2147483648, F1: 1, F2: 1}
        var v76 uint8 = 42
        _ = fi73
        _ = v75
        _ = v76
    }
    fmt.Println(v3)
    v6 += 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
