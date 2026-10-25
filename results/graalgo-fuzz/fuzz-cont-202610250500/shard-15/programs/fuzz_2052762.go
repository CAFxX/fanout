package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 string; F1 uint8; F2 uint16 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 int64; F1 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 42
    v3 := make(map[string]uint32)
    var v4 uint64 = 476
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 uint64 = v4
        v4 -= v4
        var v7 float64 = 3.14159
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        v1["hello"] = fi8
        if false {
            switch "hello" {
            case "foo bar":
                switch fi8 {
                case 128:
                    var v9 uint32 = 2
                    _ = v9
                case 42:
                    for fi10 := 0; fi10 < 9; fi10++ {
                        if true {
                            if (18446744073709551615 == 18446744073709551615) {
                                v11 := GenS1{F0: v4, F1: false}
                                var v12 int = fi10
                                _ = v11
                                _ = v12
                            }
                            fmt.Println(fi10, fi10)
                            v4 = v4
                            v13 := GenS1{F0: (v4 | 27), F1: (!false)}
                            v4 = 42
                            v14 := []int{(-(fi8)), (fi8 << ((254) & 63)), fi10}
                            if (v4 <= (^(v4))) {
                                if (1e10 != 0.0) {
                                    v2 = 0
                                } else {
                                    _ = len(v1)
                                }
                            }
                            _ = v13
                            _ = v14
                        }
                        if false {
                            fmt.Println(v2)
                        } else {
                            fmt.Println(fi10)
                            fmt.Println(v2)
                            v2 = v2
                        }
                        switch ("a" + "\n\t") {
                        case "foo bar":
                            switch "\n\t" {
                            case "hello":
                                switch "\n\t" {
                                case "a":
                                    v2 = (v2 / ((v2) | 1))
                                case "":
                                    var v15 int32 = 0
                                    delete(v3, "")
                                    _ = v15
                                case "hello":
                                    v16 := GenS1{F0: v4, F1: (((false && false) || (1 != 18446744073709551615)) && true)}
                                    _ = v16
                                default:
                                    switch fi8 {
                                    case 9223372036854775807:
                                        var v17 bool = false
                                        _ = v17
                                    }
                                }
                            case "a":
                                v2 = (v2 + v2)
                                var v18 bool = true
                                _ = v18
                            }
                        case "hello":
                            if (0 <= v4) {
                                v19 := []int{(fi8 | (fi8 - fi8)), 128, fi10}
                                fmt.Println(fi10, v4)
                                delete(v1, "a")
                                fmt.Println(v4, v2)
                                _ = v19
                            }
                        default:
                            var v20 uint8 = 61
                            _ = v20
                        }
                        _ = fi10
                    }
                default:
                    var v21 uint32 = 0
                    _ = v21
                }
            case "hello":
                for fi22 := 0; fi22 < 6; fi22++ {
                    v23 := []int{(-(fi8)), 0}
                    v2 = (v2 | v2)
                    _ = fi22
                    _ = v23
                }
            }
            for fi24 := 0; fi24 < 17; fi24++ {
                if (true && (!false)) {
                    var v25 int8 = 127
                    if true {
                        v4 += (v4 >> ((18446744073709551615) & 63))
                        v26 := []int{fi8, 9223372036854775807, fi24}
                        switch "hello" {
                        case "foo bar":
                            if true {
                                v4 = (^(v4))
                            }
                            v27 := GenS1{F0: v4, F1: true}
                            if true {
                                var v28 uint = 0
                                fmt.Println(v25)
                                var v29 int64 = -1
                                _ = v28
                                _ = v29
                            }
                            _ = v27
                        case "\n\t":
                            if true {
                                v30 := GenS1{F0: (^(v4)), F1: ((((v2 <= v2) && false) || true) || (true && true))}
                                var v31 int32 = -464
                                _ = v30
                                _ = v31
                            } else {
                                var v32 float32 = 1e10
                                switch (-(fi24)) {
                                case -2147483648:
                                    switch fi24 {
                                    case -2147483648:
                                        fmt.Println(fi24, v32, v4)
                                    case -532:
                                        fmt.Println(v2, fi8)
                                    default:
                                        v4 = v4
                                    }
                                case 841:
                                    var v33 float32 = 0.1
                                    _ = v33
                                default:
                                    fmt.Println(fi8, v32)
                                    if (!(fi8 > fi24)) {
                                        v4 = (v4 ^ 42)
                                        v25 += -45
                                        if false {
                                            var v34 int8 = v25
                                            var v35 uint = 2
                                            v35 = v35
                                            _ = v34
                                            _ = v35
                                        }
                                    }
                                }
                                fmt.Println(v4, v25, fi8)
                                v36 := GenS1{F0: v4, F1: false}
                                _ = v32
                                _ = v36
                            }
                        }
                        _ = v26
                    } else {
                        var v37 uint8 = 255
                        _ = v37
                    }
                    v38 := GenS1{F0: v4, F1: (false && (837 <= -941))}
                    _ = v25
                    _ = v38
                } else {
                    if false {
                        var v39 bool = (255 == 2)
                        var v40 int = (fi24 / (((fi24 / ((fi24) | 1))) | 1))
                        _ = v39
                        _ = v40
                    }
                    fmt.Println(fi8, fi8, v2)
                }
                v4 = 455
                _ = fi24
            }
        }
        _ = fi8
    }
    v4 += v4
    fmt.Println(v4, v2)
    for fi41 := 0; fi41 < 18; fi41++ {
        if ("" < "hello") {
            v4 += 758
            v2 = v2
            if (!(true && false)) {
                switch "\n\t" {
                case "a":
                    v42 := []int{140, (fi41 % (((fi41 + (fi41 - -891))) | 1)), (fi41 | fi41), (-(fi41))}
                    _ = v42
                case "foo bar":
                    var v43 int32 = 1
                    _ = v43
                case "\n\t":
                    var v44 int = 32767
                    _ = v44
                default:
                    switch 4294967295 {
                    case 4294967295:
                        var v45 int64 = -9223372036854775808
                        _ = v45
                    case 1:
                        if ((87 <= 100) && ("" <= "")) {
                            for fi46 := 0; fi46 < 14; fi46++ {
                                v47 := GenS1{F0: v4, F1: (32767 >= 32767)}
                                v2 = (^(v2))
                                var v48 string = ("hello" + ((("\n\t" + "hello") + (((("foo bar" + ("hello" + (((("" + "\n\t") + "\n\t") + ("a" + "a")) + "hello"))) + "hello") + "a") + ("a" + ("\n\t" + (("foo bar" + ("foo bar" + "")) + (("foo bar" + "a") + ("" + "foo bar"))))))) + ((("hello" + "a") + "hello") + "foo bar")))
                                v49 := GenS2{F0: 70, F1: (-(fi41))}
                                _ = fi46
                                _ = v47
                                _ = v48
                                _ = v49
                            }
                        }
                    case 42:
                        for fi50 := 0; fi50 < 20; fi50++ {
                            if false {
                                v51 := []int{885, fi41}
                                _ = v51
                            } else {
                                switch "hello" {
                                case "":
                                    v4 = v4
                                case "foo bar":
                                    v52 := GenS1{F0: v4, F1: true}
                                    _ = v52
                                case "a":
                                    var v53 float32 = 1.5
                                    _ = v53
                                default:
                                    if true {
                                        var v54 float32 = 1e10
                                        v55 := []int{}
                                        v2 += (v2 >> ((uint(v2)) & 63))
                                        _ = v54
                                        _ = v55
                                    }
                                }
                            }
                            _ = fi50
                        }
                    default:
                        if (("foo bar" < "foo bar") && (false || false)) {
                            v2 += v2
                        } else {
                            var v56 int64 = 1
                            v1[("foo bar" + "\n\t")] = (^(fi41))
                            v57 := GenS1{F0: (^(v4)), F1: true}
                            _ = v56
                            _ = v57
                        }
                    }
                }
                switch "hello" {
                case "foo bar":
                    fmt.Println(fi41)
                case "a":
                    var v58 int64 = -18
                    _ = v58
                case "":
                    switch "a" {
                    case "":
                        if false {
                            v4 = v4
                            v59 := GenS1{F0: 18446744073709551615, F1: false}
                            var v60 bool = true
                            _ = v59
                            _ = v60
                        }
                    case "hello":
                        _ = len(v3)
                    case "a":
                        if (!false) {
                            fmt.Println(v4)
                            v2 = v2
                            var v61 int32 = -98
                            _ = v61
                        } else {
                            v62 := GenS1{F0: 18446744073709551615, F1: (false && (true || false))}
                            v63 := GenS3{F0: "", F1: 127, F2: (v2 ^ 454)}
                            _ = v62
                            _ = v63
                        }
                    }
                }
            } else {
                for fi64 := 0; fi64 < 5; fi64++ {
                    v4 = (v4 >> ((255) & 63))
                    v65 := GenS1{F0: 224, F1: true}
                    _ = fi64
                    _ = v65
                }
                var v66 int16 = -857
                _ = v66
            }
        } else {
            switch "foo bar" {
            case "foo bar":
                v67 := []int{fi41, (^(fi41)), (fi41 / (((^(fi41))) | 1))}
                _ = v67
            }
            var v68 uint64 = (^(v4))
            var v69 int = (fi41 * fi41)
            _ = v68
            _ = v69
        }
        v4 = v4
        var v70 float32 = 0.0
        _ = fi41
        _ = v70
    }
    v71 := GenS4{F0: 255, F1: 32767}
    var v72 int = 83
    if (65535 == v2) {
        switch "" {
        case "a":
            v73 := GenS4{F0: 128, F1: int16(v72)}
            _ = v73
        }
    } else {
        v74 := GenS3{F0: (("a" + (("\n\t" + "") + "foo bar")) + "\n\t"), F1: 128, F2: v2}
        switch ("a" + "\n\t") {
        case "\n\t":
            if true {
                fmt.Println(v2, v2, v4)
            } else {
                v75 := GenS4{F0: 2, F1: -893}
                v76 := GenS4{F0: 128, F1: -134}
                v77 := GenS3{F0: "hello", F1: uint8(v4), F2: (v2 - (^(v2)))}
                _ = v75
                _ = v76
                _ = v77
            }
        case "a":
            v4 += 0
        case "foo bar":
            v4 = v4
        }
        var v78 int8 = 53
        _ = v74
        _ = v78
    }
    for fi79 := 0; fi79 < 12; fi79++ {
        fmt.Println(v4, v72, fi79)
        var v80 uint64 = v4
        var v81 uint = 519
        _ = fi79
        _ = v80
        _ = v81
    }
    delete(v3, (("" + "a") + ""))
    fmt.Println(v2, v2)
    if (922 == 0) {
        v2 -= v2
        v82 := GenS5{F0: 1, F1: -751}
        v2 -= v2
        _ = v82
    } else {
        v83 := GenS2{F0: 0, F1: 1}
        _ = v83
    }
    var v84 uint = 0
    v85 := []int{(v72 / (((^(v72))) | 1)), (v72 >> ((v84) & 63))}
    var v86 uint16 = 99
    if ((!(108 > 1000)) || true) {
        for fi87 := 0; fi87 < 19; fi87++ {
            v84 += (v84 << ((255) & 63))
            var v88 int64 = 42
            if len(v85) >= 2 {
                v85 = v85[1:len(v85)]
            }
            _ = fi87
            _ = v88
        }
        v85 = append(v85, -1)
        for fi89 := 0; fi89 < 19; fi89++ {
            var v90 uint16 = (v86 % (((v86 << ((341) & 63))) | 1))
            _ = fi89
            _ = v90
        }
    }
    v4 = (v4 | 319)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v72:", v72)
    fmt.Println("v84:", v84)
    fmt.Println("slicelen v85:", len(v85), v85)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v71
    _ = v72
    _ = v84
    _ = v85
    _ = v86
}
