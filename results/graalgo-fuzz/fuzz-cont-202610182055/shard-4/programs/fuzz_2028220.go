package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS6 struct { F0 bool; F1 bool }
type GenS7 struct { F0 bool }

func main() {
    var v1 int16 = 42
    var v2 bool = (4294967295 > 9223372036854775807)
    var v3 int32 = -822
    v4 := make(map[string]uint32)
    v5 := []string{"\n\t", "\n\t"}
    delete(v4, "a")
    v6 := GenS1{F0: 2}
    _ = len(v4)
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            v5 = append(v5, "\n\t")
            _ = fi8
        }
        if (v2 && (!v2)) {
            if ((true && (!false)) || v2) {
                switch fi7 {
                case 127:
                    v9 := GenS2{F0: 128, F1: 32767, F2: 2}
                    _ = v9
                case 128:
                    if len(v5) > 0 {
                        v5[0] = "\n\t"
                    }
                }
                v10 := GenS2{F0: 0, F1: 0, F2: 1}
                v11 := GenS3{F0: "foo bar", F1: 65535}
                _ = v10
                _ = v11
            }
            var v12 uint16 = 104
            v13 := GenS2{F0: 255, F1: 255, F2: (v12 % ((255) | 1))}
            _ = v12
            _ = v13
        }
        var v14 uint64 = 575
        _ = fi7
        _ = v14
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        if v2 {
            v3 += 306
            var v16 int64 = 9223372036854775807
            _ = v16
        } else {
            var v17 string = "a"
            var v18 int32 = (-(v3))
            _ = v17
            _ = v18
        }
        _ = fi15
    }
    v19 := GenS3{F0: "a", F1: 1}
    fmt.Println(v2)
    if (0 <= 9223372036854775807) {
        v20 := GenS4{F0: 866}
        if v2 {
            switch 117 {
            case 127:
                var v21 int16 = v1
                _ = v21
            case -765:
                var v22 uint32 = 2147483648
                _ = v22
            case -9223372036854775808:
                switch 9223372036854775807 {
                case -59:
                    var v23 uint8 = 255
                    _ = v23
                case -950:
                    v24 := GenS2{F0: 2, F1: 2, F2: 656}
                    _ = v24
                default:
                    switch "\n\t" {
                    case "hello":
                        var v25 uint32 = 2147483647
                        _ = v25
                    }
                }
            }
        } else {
            var v26 uint16 = 0
            switch 2147483647 {
            case 4294967295:
                if v2 {
                    var v27 int64 = -1
                    for fi28 := 0; fi28 < 7; fi28++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        if len(v5) > 0 {
                            v5[0] = ("\n\t" + (("hello" + "foo bar") + ("a" + "\n\t")))
                        }
                        v29 := GenS1{F0: 127}
                        _ = fi28
                        _ = v29
                    }
                    _ = v27
                }
            }
            if (((v2 && (-140 >= v1)) && true) && (v2 || ((v3 >= -1) && v2))) {
                var v30 float64 = 0.1
                _ = v30
            } else {
                var v31 uint32 = 2147483648
                _ = v31
            }
            _ = v26
        }
        _ = v20
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        fmt.Println(v3, v3)
        fmt.Println(v3, v2)
    }
    v3 = v3
    v32 := GenS5{F0: 65535, F1: 65535, F2: 0}
    if v2 {
        if (v3 == 174) {
            for fi33 := 0; fi33 < 2; fi33++ {
                switch "\n\t" {
                case "":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case "foo bar":
                    v1 += (v1 & (v1 * v1))
                }
                var v34 uint16 = 32767
                fmt.Println(v3, v2, v3)
                _ = fi33
                _ = v34
            }
        }
        for fi35 := 0; fi35 < 20; fi35++ {
            var v36 string = "\n\t"
            var v37 int8 = 127
            v36 = (v36 + "foo bar")
            _ = fi35
            _ = v36
            _ = v37
        }
        for fi38 := 0; fi38 < 12; fi38++ {
            var v39 uint32 = 346
            fmt.Println(v2)
            v39 -= 0
            _ = fi38
            _ = v39
        }
    }
    fmt.Println(v1, v3)
    v40 := GenS6{F0: v2, F1: true}
    v4["\n\t"] = 2147483648
    var v41 bool = (-2.25 == (1e10 - 0.1))
    if v41 {
        var v42 float32 = (((float32(3.14159) / -2.25) + 0.1) / 1e10)
        fmt.Println(v2)
        _ = v42
    }
    if true {
        v1 = (-(v1))
        var v43 string = ("hello" + (("\n\t" + "a") + "\n\t"))
        _ = v43
    } else {
        v44 := GenS6{F0: ((414 >= 2) && ((((("a" + ("a" + ("" + "\n\t"))) + (("\n\t" + "foo bar") + ("a" + ("" + ((("a" + "foo bar") + "a") + ""))))) + "\n\t") == ("hello" + "foo bar")) || ((v1 >= v1) && v41))), F1: true}
        if (0 == 0) {
            if (2147483647 > 32768) {
                fmt.Println(v2, v41, v2)
                switch -866 {
                case 42:
                    if false {
                        v4["\n\t"] = 2147483647
                        for fi45 := 0; fi45 < 13; fi45++ {
                            for fi46 := 0; fi46 < 18; fi46++ {
                                var v47 float64 = 0.1
                                if v41 {
                                    var v48 uint16 = 42
                                    v49 := GenS5{F0: 42, F1: 98, F2: 18446744073709551615}
                                    if len(v5) > 0 {
                                        v5[0] = ""
                                    }
                                    _ = v48
                                    _ = v49
                                } else {
                                    v2 = (false || (!(fi45 == fi45)))
                                    if len(v5) > 0 {
                                        v5[0] = ""
                                    }
                                }
                                v2 = false
                                if (v1 > int16(fi46)) {
                                    var v50 uint8 = 1
                                    _ = v50
                                }
                                _ = fi46
                                _ = v47
                            }
                            _ = fi45
                        }
                    } else {
                        _ = len(v4)
                        var v51 int8 = -1
                        _ = v51
                    }
                case 1:
                    _ = len(v5)
                default:
                    delete(v4, ("\n\t" + ("a" + "a")))
                }
            } else {
                for fi52 := 0; fi52 < 1; fi52++ {
                    _ = len(v4)
                    fmt.Println(v2)
                    fmt.Println(v3, v1)
                    _ = fi52
                }
                switch "\n\t" {
                case "\n\t":
                    var v53 uint64 = 42
                    _ = v53
                }
                for fi54 := 0; fi54 < 11; fi54++ {
                    var v55 uint8 = 42
                    v5 = append(v5, "foo bar")
                    if (0 >= 32767) {
                        if ("" == "foo bar") {
                            fmt.Println(v2)
                            v56 := GenS6{F0: false, F1: (v3 > 0)}
                            if ((-2.25 - -2.25) != 1.5) {
                                v1 += (v1 & v1)
                                v3 = v3
                                var v57 uint64 = 1
                                _ = v57
                            }
                            _ = v56
                        }
                    }
                    _ = fi54
                    _ = v55
                }
            }
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v44
    }
    switch 306 {
    case 32768:
        switch 0 {
        case 837:
            for fi58 := 0; fi58 < 10; fi58++ {
                v59 := GenS7{F0: (!(255 <= 0))}
                _ = fi58
                _ = v59
            }
        default:
            v41 = (!(-410 <= 9223372036854775807))
        }
    case 2147483647:
        var v60 int = -1
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
