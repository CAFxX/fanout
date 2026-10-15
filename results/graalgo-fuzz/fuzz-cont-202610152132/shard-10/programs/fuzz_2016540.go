package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int32 }
type GenS2 struct { F0 int32; F1 int32; F2 string }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32; F1 int16; F2 int8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int16; F1 int64 }

func main() {
    var v1 bool = (2 < 0)
    v2 := []uint32{}
    v3 := make(map[int]int)
    v4 := make(map[string]bool)
    v5 := make(map[int]uint32)
    var v6 uint32 = 4294967295
    v6 = v6
    v6 = (v6 * v6)
    if false {
        v6 -= 58
        for fi7 := 0; fi7 < 6; fi7++ {
            if v1 {
                fmt.Println(v1, v1, v6)
                if v1 {
                    v8 := v5[fi7]
                    v1 = v1
                    if (-9223372036854775808 == 9223372036854775807) {
                        if v1 {
                            v1 = ((0 < 65535) || ((v6 / ((v6) | 1)) < 768))
                            v1 = (-709 > 1)
                        }
                        for fi9 := 0; fi9 < 11; fi9++ {
                            var v10 uint8 = 128
                            var v11 int64 = 0
                            if len(v2) > 0 {
                                v2[0] = (v8 / (((v8 ^ v6)) | 1))
                            }
                            _ = fi9
                            _ = v10
                            _ = v11
                        }
                    } else {
                        v2 = append(v2, (v8 % (((^(v8))) | 1)))
                        var v12 uint64 = uint64(v6)
                        _ = v12
                    }
                    _ = v8
                } else {
                    if false {
                        v6 += (v6 - 2147483648)
                    } else {
                        v6 = 2147483647
                        var v13 string = ((("a" + "\n\t") + "foo bar") + "hello")
                        _ = v13
                    }
                    v14 := GenS1{F0: 22, F1: ((v6 << ((18446744073709551615) & 63)) <= 209), F2: 891}
                    _ = v14
                }
            }
            var v15 uint64 = 4294967295
            for fi16 := 0; fi16 < 12; fi16++ {
                if (((0.1 * 1.5) >= 0.1) && v1) {
                    var v17 float64 = (float64(1.5) / 1.0)
                    var v18 float64 = float64(0.0)
                    v15 = 1
                    _ = v17
                    _ = v18
                }
                var v19 float64 = 3.14159
                _ = fi16
                _ = v19
            }
            _ = fi7
            _ = v15
        }
        if false {
            for fi20 := 0; fi20 < 17; fi20++ {
                v1 = ((!(!(1.5 != (0.0 * 1.5)))) || (2 == 32767))
                _ = fi20
            }
            var v21 float64 = ((1.5 / 1.0) + 1.5)
            if v1 {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v6 += (v6 + 299)
            }
            _ = v21
        }
    }
    switch ("\n\t" + ("\n\t" + "a")) {
    case "hello":
        if true {
            var v22 uint64 = 819
            v22 = v22
            if (false && ((v22 ^ (v22 % (((v22 | v22)) | 1))) == v22)) {
                if ((!(!((!(!(4 < -32768))) || true))) || false) {
                    for fi23 := 0; fi23 < 3; fi23++ {
                        if len(v2) > 0 {
                            v2[0] = v6
                        }
                        for fi24 := 0; fi24 < 1; fi24++ {
                            switch (fi23 * fi23) {
                            case 2147483647:
                                v22 += (v22 % ((1) | 1))
                            case 1:
                                var v25 int64 = -9223372036854775808
                                _ = v25
                            default:
                                fmt.Println(fi24)
                            }
                            _ = fi24
                        }
                        if (39 <= 872) {
                            if (true || (127 < 42)) {
                                var v26 string = "hello"
                                var v27 int8 = 42
                                for fi28 := 0; fi28 < 7; fi28++ {
                                    v22 = v22
                                    v6 = (v6 >> ((302) & 63))
                                    v22 = v22
                                    _ = fi28
                                }
                                _ = v26
                                _ = v27
                            } else {
                                var v29 string = ("\n\t" + (("\n\t" + ("\n\t" + "\n\t")) + ("foo bar" + "")))
                                switch (v29 + "a") {
                                case "":
                                    var v30 int64 = 513
                                    _ = v30
                                }
                                _ = v29
                            }
                        }
                        _ = fi23
                    }
                    var v31 uint64 = v22
                    _ = v31
                } else {
                    for fi32 := 0; fi32 < 2; fi32++ {
                        var v33 uint64 = (^(v22))
                        var v34 bool = ((v6 % ((2147483647) | 1)) >= 2147483648)
                        v1 = ((v6 << ((0) & 63)) != (^(v6)))
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                    var v35 int = -958
                    var v36 int16 = 32767
                    _ = v35
                    _ = v36
                }
            } else {
                var v37 int = -1
                for fi38 := 0; fi38 < 3; fi38++ {
                    v39 := v4[("hello" + "\n\t")]
                    v40 := GenS2{F0: -957, F1: 2147483647, F2: "foo bar"}
                    var v41 bool = v39
                    _ = fi38
                    _ = v39
                    _ = v40
                    _ = v41
                }
                _ = v37
            }
            _ = v22
        }
    case "a":
        _ = len(v2)
    case "foo bar":
        for fi42 := 0; fi42 < 20; fi42++ {
            var v43 uint32 = 2147483648
            var v44 int = 42
            _ = fi42
            _ = v43
            _ = v44
        }
    default:
        var v45 uint32 = 2147483648
        _ = v45
    }
    var v46 uint32 = 2
    var v47 uint16 = 242
    var v48 int32 = -2147483648
    if v1 {
        v46 -= v46
    } else {
        switch 560 {
        case -1:
            if (v1 && true) {
                if ((v47 >> ((685) & 63)) >= (v47 >> ((18446744073709551615) & 63))) {
                    v49 := GenS3{F0: 2, F1: 128}
                    if (!(695 != 9223372036854775807)) {
                        switch 0 {
                        case -9223372036854775808:
                            delete(v5, 127)
                        case 903:
                            v6 = v46
                        case 9223372036854775807:
                            var v50 int32 = v48
                            _ = v50
                        }
                        v48 -= v48
                    }
                    _ = v49
                } else {
                    _ = len(v3)
                    fmt.Println(v47, v46, v1)
                    if false {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if len(v2) > 0 {
                            v2[0] = (v6 * 2147483647)
                        }
                        fmt.Println(v46, v47)
                    }
                }
                v6 -= (v46 / ((v6) | 1))
                for fi51 := 0; fi51 < 19; fi51++ {
                    v52 := GenS4{F0: fi51}
                    if true {
                        switch (fi51 << ((590) & 63)) {
                        case -2147483648:
                            v53 := GenS2{F0: (^(v48)), F1: v48, F2: "a"}
                            _ = v53
                        case -411:
                            if (4294967295 == 42) {
                                v47 = (v47 ^ 586)
                                if false {
                                    if v1 {
                                        v47 = (v47 ^ v47)
                                        v6 += (v6 | 65535)
                                    } else {
                                        _ = len(v2)
                                        v54 := GenS2{F0: v48, F1: (v48 - (v48 % ((v48) | 1))), F2: "foo bar"}
                                        _ = v54
                                    }
                                    v47 = 65535
                                } else {
                                    fmt.Println(v46, fi51, fi51)
                                    var v55 uint16 = v47
                                    _ = v55
                                }
                            }
                        case 0:
                            var v56 int8 = 1
                            _ = v56
                        }
                        v57 := GenS5{F0: v47}
                        v5[fi51] = v46
                        _ = v57
                    } else {
                        for fi58 := 0; fi58 < 1; fi58++ {
                            v59 := GenS6{F0: (^(v46)), F1: -32768, F2: 1}
                            if v1 {
                                switch fi58 {
                                case 1:
                                    v60 := GenS7{F0: ("a" + "a")}
                                    _ = v60
                                case 32768:
                                    var v61 uint16 = (v47 + (v47 & 2))
                                    _ = v61
                                case 32767:
                                    if (v47 >= v47) {
                                        v2 = append(v2, v46)
                                        var v62 int = (-(fi51))
                                        _ = v62
                                    }
                                default:
                                    switch (fi51 - fi58) {
                                    case -160:
                                        var v63 uint16 = 1
                                        _ = v63
                                    case 491:
                                        v64 := GenS8{F0: 1, F1: 9223372036854775807}
                                        _ = v64
                                    }
                                }
                            }
                            _ = fi58
                            _ = v59
                        }
                    }
                    _ = fi51
                    _ = v52
                }
            } else {
                v65 := v4[""]
                v1 = v1
                fmt.Println(v6)
                _ = v65
            }
        default:
            fmt.Println(v48, v48, v6)
        }
        delete(v3, 127)
    }
    v48 += (-(v48))
    switch "" {
    case "foo bar":
        var v66 bool = true
        _ = v66
    }
    v46 = (v6 + (v6 - 0))
    var v67 uint16 = (v47 / ((v47) | 1))
    v4["foo bar"] = v1
    for fi68 := 0; fi68 < 19; fi68++ {
        switch -9223372036854775808 {
        case 414:
            _ = len(v2)
        }
        v2 = append(v2, (v46 % ((v46) | 1)))
        var v69 int = fi68
        _ = fi68
        _ = v69
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
    _ = v67
}
