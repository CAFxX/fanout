package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32; F1 int64; F2 uint16 }
type GenS6 struct { F0 int32; F1 bool }

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[string]string)
    v3 := make(map[string]uint32)
    fmt.Println(v1)
    v4 := []int{(v1 | v1), -152}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "\n\t" {
    case "":
        v1 = (v1 * (v1 ^ v1))
    case "foo bar":
        v1 = 32767
    case "a":
        v1 -= v1
    }
    var v5 bool = false
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v1, v5)
    if v5 {
        if v5 {
            v1 -= -2147483648
            if v5 {
                v6 := GenS1{F0: (("" + "\n\t") + ""), F1: ((false && v5) && v5)}
                for fi7 := 0; fi7 < 6; fi7++ {
                    for fi8 := 0; fi8 < 6; fi8++ {
                        v1 = 32767
                        switch "foo bar" {
                        case "":
                            var v9 int8 = -84
                            _ = v9
                        case "foo bar":
                            var v10 uint = 0
                            _ = v10
                        default:
                            if v5 {
                                var v11 uint64 = 1
                                _ = v11
                            }
                        }
                        _ = fi8
                    }
                    switch "a" {
                    case "\n\t":
                        var v12 uint8 = 42
                        _ = v12
                    case "a":
                        var v13 uint16 = 552
                        _ = v13
                    case "foo bar":
                        v5 = v5
                    default:
                        var v14 uint32 = 42
                        _ = v14
                    }
                    if v5 {
                        var v15 uint32 = 2147483648
                        fmt.Println(v1)
                        _ = v15
                    } else {
                        v5 = (42 < -427)
                    }
                    _ = fi7
                }
                var v16 string = "a"
                _ = v16
                _ = v6
            }
            v2["a"] = "\n\t"
        } else {
            v17 := v2[("" + ("" + ((("\n\t" + ("a" + "a")) + "hello") + "")))]
            var v18 int8 = int8(v1)
            _ = v17
            _ = v18
        }
        var v19 float32 = 1.5
        _ = v19
    } else {
        if v5 {
            var v20 int = (-(v1))
            _ = v20
        }
        switch "hello" {
        case "":
            if false {
                switch "\n\t" {
                case "foo bar":
                    v5 = v5
                }
            } else {
                v1 -= v1
            }
        default:
            if len(v4) > 0 {
                v4[0] = v1
            }
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v21 := GenS1{F0: ("foo bar" + ("\n\t" + "foo bar")), F1: v5}
    for fi22 := 0; fi22 < 17; fi22++ {
        if (v5 && (v5 || (128 == 255))) {
            var v23 int64 = -1
            fmt.Println(fi22)
            _ = v23
        } else {
            if (41 != 0) {
                for fi24 := 0; fi24 < 20; fi24++ {
                    v5 = (0 == 1)
                    var v25 string = ("a" + "\n\t")
                    v26 := GenS2{F0: 1, F1: 42, F2: 666}
                    switch "a" {
                    case "hello":
                        v5 = v5
                    default:
                        v3[(v25 + (("" + v25) + ("" + v25)))] = 2147483648
                        v27 := v3["hello"]
                        var v28 uint32 = 65535
                        _ = v27
                        _ = v28
                    }
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            } else {
                for fi29 := 0; fi29 < 16; fi29++ {
                    if (("\n\t" < "") && (0 > 18446744073709551615)) {
                        var v30 uint16 = 2
                        _ = v30
                    } else {
                        v31 := GenS3{F0: 542}
                        _ = v31
                    }
                    switch fi22 {
                    case 4294967295:
                        var v32 int16 = 97
                        _ = v32
                    case 127:
                        fmt.Println(v5, fi22)
                    case 32768:
                        v33 := GenS3{F0: 1}
                        _ = v33
                    }
                    _ = fi29
                }
            }
            for fi34 := 0; fi34 < 5; fi34++ {
                var v35 float32 = (0.0 / (1e10 - -2.25))
                _ = fi34
                _ = v35
            }
            switch "" {
            case "":
                for fi36 := 0; fi36 < 10; fi36++ {
                    fmt.Println(fi36, fi22)
                    v5 = (255 != 211)
                    _ = fi36
                }
            case "foo bar":
                v37 := v2[""]
                _ = v37
            case "a":
                for fi38 := 0; fi38 < 5; fi38++ {
                    v39 := GenS2{F0: 13, F1: 1, F2: 42}
                    _ = fi38
                    _ = v39
                }
            default:
                if v5 {
                    var v40 int32 = 509
                    _ = v40
                } else {
                    v41 := GenS4{F0: 437}
                    _ = v41
                }
            }
        }
        v1 += (fi22 ^ v1)
        _ = fi22
    }
    if v5 {
        var v42 string = ("a" + (("" + "\n\t") + ""))
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v42
    } else {
        if (0 > 211) {
            if false {
                if false {
                    switch ("a" + "\n\t") {
                    case "\n\t":
                        for fi43 := 0; fi43 < 6; fi43++ {
                            v1 = (v1 ^ v1)
                            _ = fi43
                        }
                    case "foo bar":
                        for fi44 := 0; fi44 < 17; fi44++ {
                            var v45 uint8 = 68
                            _ = fi44
                            _ = v45
                        }
                    default:
                        v5 = v5
                    }
                    v4 = append(v4, v1)
                    _ = len(v4)
                }
                for fi46 := 0; fi46 < 2; fi46++ {
                    v47 := v2[("\n\t" + "a")]
                    var v48 string = "foo bar"
                    for fi49 := 0; fi49 < 16; fi49++ {
                        delete(v2, v47)
                        _ = fi49
                    }
                    _ = fi46
                    _ = v47
                    _ = v48
                }
                v50 := GenS4{F0: 2147483648}
                _ = v50
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v51 bool = (-620 == 32767)
                _ = v51
            }
            for fi52 := 0; fi52 < 9; fi52++ {
                v53 := GenS2{F0: 18446744073709551615, F1: 139, F2: 32767}
                _ = fi52
                _ = v53
            }
            v54 := GenS5{F0: 1, F1: 0, F2: 255}
            _ = v54
        } else {
            delete(v3, ("hello" + ("" + "")))
            switch (("a" + "\n\t") + "hello") {
            case "a":
                v5 = (false && (0.0 >= (0.0 / -2.25)))
            case "\n\t":
                if v5 {
                    v55 := GenS2{F0: 2, F1: 58, F2: 0}
                    v56 := GenS3{F0: 4294967295}
                    var v57 int32 = -1
                    _ = v55
                    _ = v56
                    _ = v57
                } else {
                    v5 = v5
                }
            }
            for fi58 := 0; fi58 < 18; fi58++ {
                if (!v5) {
                    var v59 uint64 = 2
                    _ = v59
                } else {
                    var v60 string = "\n\t"
                    var v61 int8 = 127
                    var v62 uint = 2
                    _ = v60
                    _ = v61
                    _ = v62
                }
                for fi63 := 0; fi63 < 12; fi63++ {
                    v5 = v5
                    fmt.Println(v5, fi58, v1)
                    _ = fi63
                }
                _ = fi58
            }
        }
    }
    for fi64 := 0; fi64 < 2; fi64++ {
        v65 := GenS6{F0: 42, F1: (!v5)}
        _ = fi64
        _ = v65
    }
    if (("a" + "\n\t") >= "a") {
        _ = len(v4)
    }
    for fi66 := 0; fi66 < 18; fi66++ {
        if (1 == 42) {
            if (70 <= 43) {
                v1 += fi66
                if len(v4) > 0 {
                    v4[0] = 986
                }
            } else {
                v1 -= v1
            }
            var v67 int16 = 42
            _ = v67
        } else {
            v5 = v5
        }
        if (!v5) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if v5 {
                fmt.Println(fi66, v5, v5)
                var v68 int32 = 1
                _ = v68
            } else {
                v5 = (!v5)
                var v69 uint64 = 4294967295
                _ = v69
            }
        }
        switch "\n\t" {
        case "foo bar":
            v5 = (0 <= -1)
        case "hello":
            var v70 int16 = 32767
            _ = v70
        }
        _ = fi66
    }
    switch (^(v1)) {
    case -9223372036854775808:
        _ = len(v3)
    case -2147483648:
        _ = len(v4)
    case 0:
        var v71 int16 = 772
        _ = v71
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
