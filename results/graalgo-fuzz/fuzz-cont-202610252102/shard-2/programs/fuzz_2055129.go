package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 string; F2 uint8 }
type GenS7 struct { F0 int8; F1 uint16 }

func main() {
    var v1 uint32 = 42
    var v2 bool = true
    v3 := []uint32{2147483648, (v1 ^ 241), v1, v1}
    v4 := make(map[string]bool)
    var v5 bool = (1 <= 780)
    v6 := make(map[string]string)
    for fi7 := 0; fi7 < 11; fi7++ {
        v2 = (134 >= 255)
        switch (fi7 + (fi7 | (fi7 - fi7))) {
        case -675:
            if v5 {
                if v2 {
                    fmt.Println(v5)
                    v3 = append(v3, v1)
                }
                v8 := GenS1{F0: 956, F1: 65535, F2: 128}
                if (0 != 4) {
                    var v9 int8 = -128
                    _ = v9
                }
                _ = v8
            } else {
                v2 = (!false)
                var v10 uint64 = 307
                if true {
                    v11 := GenS2{F0: -1}
                    v12 := GenS3{F0: -1, F1: 32767, F2: uint16(v1)}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v11
                    _ = v12
                } else {
                    var v13 uint = 42
                    v14 := GenS3{F0: int(v10), F1: 101, F2: 65535}
                    _ = v13
                    _ = v14
                }
                _ = v10
            }
        case -2147483648:
            var v15 int16 = 954
            _ = v15
        default:
            if v2 {
                var v16 float32 = -2.25
                v17 := v6[("\n\t" + ("" + ((("\n\t" + "a") + "") + "")))]
                _ = len(v3)
                _ = v16
                _ = v17
            }
        }
        _ = fi7
    }
    switch "hello" {
    case "foo bar":
        switch 42 {
        case 976:
            switch 23 {
            case 302:
                delete(v6, "")
            case 2147483647:
                for fi18 := 0; fi18 < 3; fi18++ {
                    var v19 int8 = -128
                    _ = fi18
                    _ = v19
                }
            }
        case 128:
            v5 = (2 > 692)
        case 351:
            var v20 uint = 0
            _ = v20
        default:
            fmt.Println(v5)
        }
    }
    if true {
        v2 = (v5 || true)
        var v21 int = 1
        _ = len(v3)
        _ = v21
    } else {
        v22 := GenS4{F0: 127}
        var v23 uint32 = v1
        _ = v22
        _ = v23
    }
    if (("" + ("\n\t" + "a")) > "foo bar") {
        if (v1 <= 1) {
            switch ("a" + "") {
            case "a":
                v5 = (1 < 664)
            default:
                v24 := GenS4{F0: -1}
                _ = v24
            }
            v25 := GenS4{F0: -1}
            _ = len(v3)
            _ = v25
        }
        v26 := GenS5{F0: 42}
        _ = v26
    } else {
        _ = len(v6)
        var v27 string = ("hello" + ("hello" + (("hello" + "hello") + "\n\t")))
        switch -419 {
        case 32768:
            v5 = ((645 < 0) || (v1 == (v1 + v1)))
        case 9223372036854775807:
            v5 = (!v2)
        case -9223372036854775808:
            if true {
                v28 := GenS2{F0: 939}
                _ = v28
            }
        default:
            for fi29 := 0; fi29 < 7; fi29++ {
                var v30 uint16 = 255
                _ = fi29
                _ = v30
            }
        }
        _ = v27
    }
    v31 := GenS5{F0: 0}
    for fi32 := 0; fi32 < 10; fi32++ {
        var v33 uint32 = (v1 / (((v1 / ((v1) | 1))) | 1))
        v34 := GenS6{F0: 255, F1: "", F2: 42}
        _ = fi32
        _ = v33
        _ = v34
    }
    if v2 {
        if (!v5) {
            fmt.Println(v2, v5)
            for fi35 := 0; fi35 < 6; fi35++ {
                var v36 int16 = 0
                switch ("a" + ("foo bar" + "hello")) {
                case "":
                    switch "" {
                    case "\n\t":
                        if (!v5) {
                            for fi37 := 0; fi37 < 13; fi37++ {
                                var v38 uint32 = (v1 & v1)
                                var v39 float64 = (float64(((1e10 - 1e10) / 3.14159)) * 1e10)
                                _ = fi37
                                _ = v38
                                _ = v39
                            }
                        } else {
                            v36 += (v36 ^ v36)
                        }
                    case "foo bar":
                        switch "\n\t" {
                        case "hello":
                            var v40 bool = ((!false) && (42 > -1))
                            _ = v40
                        default:
                            var v41 bool = v2
                            _ = v41
                        }
                    default:
                        switch "hello" {
                        case "":
                            if v2 {
                                fmt.Println(v5)
                                v36 = v36
                            } else {
                                fmt.Println(v5, v5)
                                delete(v4, ("a" + "foo bar"))
                            }
                        default:
                            v42 := GenS5{F0: 143}
                            _ = v42
                        }
                    }
                default:
                    var v43 uint = 65535
                    _ = v43
                }
                var v44 bool = true
                _ = fi35
                _ = v36
                _ = v44
            }
        } else {
            switch "a" {
            case "foo bar":
                _ = len(v4)
            }
        }
        v1 = (v1 * (v1 % ((65535) | 1)))
        var v45 uint = 443
        _ = v45
    } else {
        fmt.Println(v2, v1)
        for fi46 := 0; fi46 < 4; fi46++ {
            for fi47 := 0; fi47 < 13; fi47++ {
                v5 = (v2 && (1.5 == float32(1.5)))
                var v48 int8 = int8(fi47)
                var v49 uint8 = 0
                _ = fi47
                _ = v48
                _ = v49
            }
            v1 += (v1 ^ v1)
            var v50 float32 = ((0.0 - 1e10) / 3.14159)
            _ = fi46
            _ = v50
        }
        _ = len(v3)
    }
    for fi51 := 0; fi51 < 17; fi51++ {
        _ = len(v6)
        v1 = v1
        for fi52 := 0; fi52 < 13; fi52++ {
            switch (^(fi51)) {
            case -1:
                fmt.Println(v5, v5, fi52)
            case 32767:
                if (-337 < 171) {
                    var v53 uint = 18446744073709551615
                    delete(v4, "")
                    var v54 int16 = 32767
                    if (2 <= 2) {
                        var v55 uint64 = 2
                        _ = v55
                    } else {
                        var v56 uint32 = v1
                        var v57 int64 = 919
                        var v58 uint = v53
                        _ = v56
                        _ = v57
                        _ = v58
                    }
                    v1 -= v1
                    v53 += v53
                    _ = v53
                    _ = v54
                }
            default:
                v59 := GenS5{F0: 133}
                _ = v59
            }
            v60 := v4["a"]
            if v2 {
                var v61 int8 = 127
                _ = v61
            } else {
                switch "\n\t" {
                case "foo bar":
                    _ = len(v4)
                case "\n\t":
                    v1 = v1
                case "a":
                    if v2 {
                        if (0.1 > ((1.5 + (3.14159 * 3.14159)) + (0.1 + 0.1))) {
                            var v62 int64 = 42
                            _ = v62
                        } else {
                            var v63 float64 = 1.5
                            v64 := GenS4{F0: fi51}
                            _ = v63
                            _ = v64
                        }
                        var v65 int64 = 9223372036854775807
                        fmt.Println(v65, v60, v5)
                        v66 := GenS2{F0: v65}
                        v67 := GenS6{F0: 42, F1: "\n\t", F2: 1}
                        _ = len(v4)
                        var v68 uint64 = 4294967295
                        _ = v65
                        _ = v66
                        _ = v67
                        _ = v68
                    }
                }
            }
            _ = fi52
            _ = v60
        }
        _ = fi51
    }
    v69 := GenS1{F0: 0, F1: 2, F2: 42}
    v5 = (!v5)
    if (982 < 59) {
        v2 = v2
    } else {
        var v70 int64 = 9223372036854775807
        var v71 int32 = -266
        _ = v70
        _ = v71
    }
    delete(v6, "hello")
    v6["\n\t"] = (("a" + "") + "a")
    if (!true) {
        if v5 {
            for fi72 := 0; fi72 < 5; fi72++ {
                if (132 >= 175) {
                    var v73 uint16 = 2
                    for fi74 := 0; fi74 < 12; fi74++ {
                        v2 = true
                        switch "a" {
                        case "":
                            var v75 uint16 = v73
                            _ = v75
                        case "foo bar":
                            v6[("a" + "\n\t")] = "hello"
                        default:
                            fmt.Println(v2, fi72)
                        }
                        _ = fi74
                    }
                    _ = v73
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if v2 {
                        v4["foo bar"] = ((v2 || (-32768 >= 32767)) && (-1 < 1))
                        var v76 int32 = 2147483647
                        switch 0 {
                        case 2147483647:
                            for fi77 := 0; fi77 < 18; fi77++ {
                                switch (fi77 | fi72) {
                                case 32768:
                                    fmt.Println(v2, v5, v1)
                                case -1:
                                    fmt.Println(fi77, v76)
                                }
                                _ = fi77
                            }
                        case 127:
                            v5 = v5
                        case -2147483648:
                            v78 := GenS7{F0: 85, F1: 0}
                            _ = v78
                        }
                        _ = v76
                    } else {
                        v2 = false
                        v5 = ("foo bar" < "\n\t")
                    }
                    v1 -= v1
                }
                var v79 int8 = -107
                _ = fi72
                _ = v79
            }
            v5 = (255 != 0)
        } else {
            if (-321 == -745) {
                for fi80 := 0; fi80 < 11; fi80++ {
                    var v81 uint16 = uint16(v1)
                    v3 = append(v3, v1)
                    v81 = 65535
                    _ = fi80
                    _ = v81
                }
            }
        }
    }
    var v82 uint64 = 1
    v82 = v82
    var v83 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v82
    _ = v83
}
