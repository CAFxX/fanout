package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 int16 }
type GenS7 struct { F0 int32; F1 bool; F2 uint64 }
type GenS8 struct { F0 int8; F1 uint32 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 string; F1 bool }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 string; F1 int64; F2 uint64 }
type GenS13 struct { F0 uint32; F1 uint32; F2 uint }
type GenS14 struct { F0 uint64; F1 string; F2 int8 }
type GenS15 struct { F0 int64; F1 int; F2 int32 }
type GenS16 struct { F0 bool; F1 string; F2 int }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[uint32]uint32)
    v3 := []int{127, 711}
    var v4 bool = ((0.1 <= ((0.1 - (0.1 * 1.5)) / 1.0)) && false)
    v5 := GenS1{F0: 18446744073709551615, F1: 1}
    _ = len(v3)
    var v6 uint64 = 1
    var v7 uint32 = 0
    if v4 {
        v3 = append(v3, 42)
    } else {
        for fi8 := 0; fi8 < 2; fi8++ {
            switch ("\n\t" + "a") {
            case "hello":
                if v4 {
                    v9 := GenS1{F0: uint(v6), F1: fi8}
                    _ = v9
                }
            default:
                for fi10 := 0; fi10 < 2; fi10++ {
                    switch (fi10 << ((255) & 63)) {
                    case -1:
                        if true {
                            if (0 > int64(fi10)) {
                                fmt.Println(v6, fi10, v7)
                                var v11 uint16 = 540
                                switch "a" {
                                case "hello":
                                    v4 = (!(-933 == v1))
                                    if (4294967295 != 18446744073709551615) {
                                        v7 -= 42
                                    }
                                    v12 := GenS2{F0: -1}
                                    if v4 {
                                        v4 = v4
                                        var v13 int8 = -43
                                        v13 += v13
                                        _ = v13
                                    } else {
                                        if (!(true && v4)) {
                                            switch fi10 {
                                            case -9223372036854775808:
                                                if v4 {
                                                    fmt.Println(v6, v6)
                                                    v4 = v4
                                                } else {
                                                    v14 := GenS3{F0: 0, F1: 42, F2: (^(v11))}
                                                    _ = v14
                                                }
                                                v6 -= v6
                                                var v15 uint = 733
                                                _ = v15
                                            case 665:
                                                v16 := GenS4{F0: 2, F1: (v6 > v6)}
                                                _ = v16
                                            }
                                            switch "hello" {
                                            case "":
                                                var v17 int32 = int32(v7)
                                                _ = v17
                                            case "foo bar":
                                                v18 := GenS3{F0: v7, F1: 36, F2: (v11 & (v11 >> ((415) & 63)))}
                                                _ = v18
                                            case "hello":
                                                v19 := GenS1{F0: uint(v6), F1: fi10}
                                                _ = v19
                                            default:
                                                switch (fi10 + (fi10 | -190)) {
                                                case -2147483648:
                                                    switch (("foo bar" + "\n\t") + "\n\t") {
                                                    case "":
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    case "a":
                                                        if true {
                                                            var v20 uint32 = v7
                                                            _ = v20
                                                        }
                                                    default:
                                                        fmt.Println(v4)
                                                    }
                                                }
                                            }
                                            v7 += (v7 + v7)
                                            if false {
                                                delete(v2, 42)
                                                fmt.Println(v1, fi10, fi8)
                                                v21 := GenS1{F0: 515, F1: fi10}
                                                _ = v21
                                            }
                                        } else {
                                            v7 = (v7 ^ (v7 % (((v7 << ((42) & 63))) | 1)))
                                            v7 += 0
                                            v4 = ((!((v7 - v7) != (^(v7)))) || v4)
                                        }
                                    }
                                    v22 := GenS3{F0: 2, F1: 127, F2: (^(v11))}
                                    _ = v12
                                    _ = v22
                                case "a":
                                    switch "a" {
                                    case "a":
                                        switch "" {
                                        case "":
                                            _ = len(v2)
                                        case "hello":
                                            v4 = (100 < -92)
                                            if (v4 || ((0 == v11) || v4)) {
                                                var v23 int64 = 786
                                                _ = v23
                                            } else {
                                                v24 := GenS5{F0: -2147483648}
                                                v25 := GenS6{F0: 42, F1: -1}
                                                _ = v24
                                                _ = v25
                                            }
                                            _ = len(v2)
                                        default:
                                            if v4 {
                                                var v26 int = 32767
                                                v27 := GenS4{F0: 1, F1: v4}
                                                if (3.14159 <= (0.0 * (3.14159 * 1e10))) {
                                                    v28 := GenS3{F0: v7, F1: -76, F2: (v11 >> ((255) & 63))}
                                                    v29 := GenS7{F0: v1, F1: ((32767 != -5) && (!(!(false || false)))), F2: v6}
                                                    var v30 float64 = -2.25
                                                    _ = v28
                                                    _ = v29
                                                    _ = v30
                                                } else {
                                                    v7 = (v7 + (v7 * (v7 ^ (v7 * v7))))
                                                }
                                                var v31 int64 = int64(v26)
                                                _ = len(v2)
                                                v7 -= v7
                                                _ = v26
                                                _ = v27
                                                _ = v31
                                            }
                                        }
                                    case "foo bar":
                                        var v32 uint64 = (v6 % ((0) | 1))
                                        _ = v32
                                    case "\n\t":
                                        switch "" {
                                        case "hello":
                                            var v33 float64 = 3.14159
                                            _ = v33
                                        case "foo bar":
                                            v4 = ((v4 || (v4 || v4)) || (false || false))
                                        }
                                    default:
                                        if (!(v4 && (!(2 <= 128)))) {
                                            if (!("a" == "a")) {
                                                v34 := GenS8{F0: -1, F1: v7}
                                                v35 := GenS9{F0: 255}
                                                _ = v34
                                                _ = v35
                                            } else {
                                                fmt.Println(v11)
                                                v4 = v4
                                            }
                                        } else {
                                            v36 := GenS10{F0: "a", F1: (v4 || true)}
                                            switch fi8 {
                                            case 42:
                                                _ = len(v2)
                                            default:
                                                switch fi8 {
                                                case 128:
                                                    v37 := GenS11{F0: 2}
                                                    _ = v37
                                                case -707:
                                                    var v38 int32 = (-(v1))
                                                    _ = v38
                                                }
                                            }
                                            _ = v36
                                        }
                                    }
                                case "":
                                    fmt.Println(fi8)
                                }
                                _ = v11
                            } else {
                                var v39 float64 = 3.14159
                                v40 := GenS10{F0: "foo bar", F1: true}
                                var v41 int16 = 1000
                                _ = v39
                                _ = v40
                                _ = v41
                            }
                            v1 += 1
                        } else {
                            v42 := GenS12{F0: "\n\t", F1: 0, F2: 1}
                            v2[58] = 2147483648
                            _ = v42
                        }
                        v7 = 656
                        var v43 uint32 = (v7 / (((v7 & (v7 % (((v7 >> ((140) & 63))) | 1)))) | 1))
                        var v44 int16 = -310
                        _ = v43
                        _ = v44
                    case -9223372036854775808:
                        var v45 int16 = -1
                        _ = v45
                    case -402:
                        delete(v2, v7)
                    default:
                        v46 := GenS13{F0: v7, F1: (^(v7)), F2: 0}
                        fmt.Println(fi8, v7)
                        v4 = (("\n\t" + "\n\t") < "")
                        _ = v46
                    }
                    fmt.Println(v6, v1, fi10)
                    var v47 uint8 = 127
                    var v48 int64 = 42
                    switch ("hello" + "foo bar") {
                    case "a":
                        var v49 int8 = 127
                        _ = v49
                    default:
                        var v50 bool = (v4 || v4)
                        _ = v50
                    }
                    v7 = 2147483647
                    _ = fi10
                    _ = v47
                    _ = v48
                }
            }
            if false {
                if false {
                    delete(v2, (^(v7)))
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(fi8)
                    case "hello":
                        for fi51 := 0; fi51 < 11; fi51++ {
                            v1 += 1
                            _ = fi51
                        }
                    }
                }
                v1 = v1
                fmt.Println(v7, v6)
            } else {
                if v4 {
                    for fi52 := 0; fi52 < 14; fi52++ {
                        if v4 {
                            if (0 < 42) {
                                v4 = (!((!v4) || v4))
                            }
                        }
                        if (-1 <= 487) {
                            v3 = append(v3, -487)
                        } else {
                            v4 = v4
                            if false {
                                switch ((("hello" + ("" + "a")) + "") + "\n\t") {
                                case "foo bar":
                                    v53 := GenS8{F0: 85, F1: (v7 % ((v7) | 1))}
                                    _ = v53
                                case "a":
                                    v54 := GenS5{F0: v1}
                                    _ = v54
                                }
                                v55 := GenS12{F0: "hello", F1: 1, F2: (v6 ^ uint64(v7))}
                                var v56 int32 = int32(fi52)
                                _ = v55
                                _ = v56
                            }
                            v57 := GenS14{F0: 922, F1: "\n\t", F2: 42}
                            _ = v57
                        }
                        _ = fi52
                    }
                    var v58 int16 = 42
                    _ = v58
                }
            }
            v7 -= v7
            _ = fi8
        }
    }
    v3 = append(v3, -988)
    for fi59 := 0; fi59 < 18; fi59++ {
        var v60 uint64 = v6
        for fi61 := 0; fi61 < 14; fi61++ {
            var v62 float64 = 1e10
            _ = fi61
            _ = v62
        }
        for fi63 := 0; fi63 < 8; fi63++ {
            v7 -= v7
            v1 = (v1 % ((-812) | 1))
            v64 := GenS15{F0: -6, F1: (fi63 + 1), F2: -227}
            _ = fi63
            _ = v64
        }
        _ = fi59
        _ = v60
    }
    var v65 int8 = -1
    var v66 int32 = (v1 / ((v1) | 1))
    var v67 bool = (910 < 4294967295)
    fmt.Println(v7, v7)
    if ("" == "hello") {
        v67 = v67
        var v68 string = ("foo bar" + "\n\t")
        var v69 int32 = v1
        _ = v68
        _ = v69
    } else {
        var v70 int = 115
        _ = v70
    }
    for fi71 := 0; fi71 < 1; fi71++ {
        v66 -= v66
        v72 := GenS16{F0: (!v67), F1: "hello", F2: (-(fi71))}
        _ = fi71
        _ = v72
    }
    var v73 float32 = (-2.25 - 0.1)
    var v74 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v7:", v7)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
    _ = v67
    _ = v7
    _ = v73
    _ = v74
}
