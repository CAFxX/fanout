package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 bool; F1 uint16; F2 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 int64; F1 string; F2 int8 }
type GenS8 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint64; F1 string }
type GenS11 struct { F0 uint; F1 int }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{-675, 417, 402}
    var v3 int32 = 344
    var v4 uint32 = 4294967295
    var v5 int32 = (v3 >> ((2) & 63))
    for fi6 := 0; fi6 < 6; fi6++ {
        switch (fi6 % ((int(v3)) | 1)) {
        case -990:
            v7 := GenS1{F0: 255}
            _ = v7
        case 32768:
            v8 := GenS2{F0: 1, F1: (fi6 ^ fi6)}
            _ = v8
        }
        var v9 float32 = 0.0
        _ = fi6
        _ = v9
    }
    v10 := GenS3{F0: 1, F1: -32768}
    v11 := GenS2{F0: 2, F1: 32768}
    v3 -= v5
    fmt.Println(v4, v4, v3)
    var v12 int16 = 307
    if ((-1 > 127) || false) {
        v13 := GenS2{F0: 4294967295, F1: 0}
        fmt.Println(v5, v5, v4)
        _ = v13
    } else {
        v5 = v5
    }
    v3 = v5
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := GenS1{F0: 32767}
        v5 -= v3
        v5 -= -1
        _ = fi14
        _ = v15
    }
    v4 -= (v4 | v4)
    v16 := GenS4{F0: ((v12 / ((v12) | 1)) == (v12 ^ v12)), F1: 1, F2: (("hello" + "a") + "hello")}
    if (-128 < -106) {
        if true {
            for fi17 := 0; fi17 < 14; fi17++ {
                var v18 float64 = (1e10 - 3.14159)
                var v19 float32 = 3.14159
                _ = fi17
                _ = v18
                _ = v19
            }
            for fi20 := 0; fi20 < 20; fi20++ {
                for fi21 := 0; fi21 < 11; fi21++ {
                    v22 := GenS5{F0: 201}
                    if true {
                        fmt.Println(v4)
                        var v23 int16 = v12
                        switch "hello" {
                        case "\n\t":
                            v12 -= (v12 | 32767)
                        case "foo bar":
                            v24 := GenS3{F0: 0, F1: v12}
                            _ = v24
                        default:
                            if ((v23 % ((166) | 1)) != -1) {
                                switch "a" {
                                case "\n\t":
                                    switch "" {
                                    case "hello":
                                        if false {
                                            v4 = (v4 << ((4294967295) & 63))
                                            _ = len(v1)
                                            v3 -= 42
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        } else {
                                            v25 := GenS1{F0: 65535}
                                            if (42 != 649) {
                                                var v26 bool = false
                                                if false {
                                                    var v27 bool = v26
                                                    fmt.Println(v12, v3)
                                                    var v28 float32 = 1.5
                                                    _ = v27
                                                    _ = v28
                                                } else {
                                                    if (v26 && v26) {
                                                        v12 -= v23
                                                        if (341 <= 255) {
                                                            v2 = append(v2, (fi21 & 0))
                                                            var v29 uint8 = 1
                                                            _ = v29
                                                        } else {
                                                            v23 = v23
                                                        }
                                                    } else {
                                                        switch "foo bar" {
                                                        case "a":
                                                            v23 = v23
                                                        case "hello":
                                                            v30 := GenS1{F0: 0}
                                                            _ = v30
                                                        default:
                                                            if true {
                                                                switch (fi21 | (fi21 & 531)) {
                                                                case 1:
                                                                    v12 = v12
                                                                case -2147483648:
                                                                    var v31 uint16 = 1
                                                                    _ = v31
                                                                }
                                                                if v26 {
                                                                    delete(v1, v4)
                                                                    if (74 == fi20) {
                                                                        v32 := GenS6{F0: v4, F1: 65535}
                                                                        switch "" {
                                                                        case "hello":
                                                                            var v33 bool = (v26 || v26)
                                                                            _ = v33
                                                                        case "\n\t":
                                                                            var v34 uint = 0
                                                                            _ = v34
                                                                        case "foo bar":
                                                                            switch (fi20 % ((127) | 1)) {
                                                                            case -596:
                                                                                v35 := GenS2{F0: 42, F1: (fi21 & (fi20 * fi20))}
                                                                                _ = v35
                                                                            }
                                                                        }
                                                                        delete(v1, (^(v4)))
                                                                        _ = v32
                                                                    } else {
                                                                        v36 := GenS7{F0: -350, F1: "", F2: 1}
                                                                        var v37 int = -2147483648
                                                                        if v26 {
                                                                            v4 = v4
                                                                            v5 += v5
                                                                        }
                                                                        _ = v36
                                                                        _ = v37
                                                                    }
                                                                } else {
                                                                    v38 := GenS8{F0: 4, F1: v4, F2: (v5 ^ v3)}
                                                                    v26 = v26
                                                                    _ = v38
                                                                }
                                                                if v26 {
                                                                    switch fi21 {
                                                                    case -9223372036854775808:
                                                                        v3 = 1
                                                                    default:
                                                                        var v39 uint64 = 4294967295
                                                                        _ = v39
                                                                    }
                                                                    v4 -= (v4 ^ v4)
                                                                    var v40 string = (("hello" + "") + "foo bar")
                                                                    _ = v40
                                                                }
                                                            }
                                                        }
                                                        v3 = -1
                                                        v5 = (v3 << ((379) & 63))
                                                    }
                                                }
                                                if v26 {
                                                    v3 -= (v5 % (((v5 / (((v3 >> ((977) & 63))) | 1))) | 1))
                                                    var v41 int32 = -1
                                                    _ = len(v2)
                                                    _ = v41
                                                } else {
                                                    switch "hello" {
                                                    case "\n\t":
                                                        v2 = append(v2, (fi20 % ((fi21) | 1)))
                                                    case "":
                                                        fmt.Println(v26)
                                                    case "foo bar":
                                                        var v42 uint16 = 640
                                                        var v43 int16 = int16(v4)
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                }
                                                _ = v26
                                            }
                                            v5 = v5
                                            _ = v25
                                        }
                                    case "foo bar":
                                        fmt.Println(fi20, v5, fi21)
                                    case "\n\t":
                                        v3 = 42
                                    default:
                                        switch ("a" + "") {
                                        case "\n\t":
                                            v5 += (v3 ^ (v5 + -2147483648))
                                        case "foo bar":
                                            var v44 uint = 4294967295
                                            _ = v44
                                        case "a":
                                            v45 := GenS2{F0: 0, F1: fi21}
                                            _ = v45
                                        }
                                    }
                                default:
                                    var v46 bool = true
                                    _ = v46
                                }
                                var v47 string = "foo bar"
                                _ = v47
                            } else {
                                v1[v4] = 1
                                _ = len(v2)
                            }
                        }
                        _ = v23
                    }
                    if (!(fi21 >= fi21)) {
                        var v48 int8 = -128
                        _ = v48
                    } else {
                        if false {
                            v12 -= v12
                            if true {
                                fmt.Println(v12, v3)
                            }
                            v4 += (v4 ^ v4)
                        } else {
                            if (!false) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v49 uint = 42
                                _ = v49
                            } else {
                                v12 += v12
                                if ((v12 - v12) < (v12 >> ((255) & 63))) {
                                    var v50 int32 = (-(v5))
                                    switch (("\n\t" + ("\n\t" + ("foo bar" + "\n\t"))) + ("\n\t" + "foo bar")) {
                                    case "hello":
                                        var v51 int = (fi21 / (((fi21 / ((fi21) | 1))) | 1))
                                        v2 = append(v2, fi20)
                                        var v52 bool = true
                                        _ = v51
                                        _ = v52
                                    case "\n\t":
                                        if len(v2) > 0 {
                                            v2[0] = fi21
                                        }
                                    case "foo bar":
                                        fmt.Println(fi20, v3, v50)
                                    default:
                                        v5 -= (v50 >> ((726) & 63))
                                    }
                                    v5 += (v5 / ((v3) | 1))
                                    _ = v50
                                } else {
                                    var v53 uint = 2
                                    if (!true) {
                                        v54 := GenS9{F0: 1}
                                        _ = v54
                                    }
                                    _ = v53
                                }
                            }
                            v12 = -827
                            switch "" {
                            case "a":
                                if (true || false) {
                                    fmt.Println(v3)
                                    var v55 uint16 = 904
                                    _ = v55
                                }
                            case "\n\t":
                                var v56 uint8 = 255
                                _ = v56
                            default:
                                var v57 int = 32768
                                _ = v57
                            }
                        }
                    }
                    var v58 string = ("\n\t" + "foo bar")
                    _ = fi21
                    _ = v22
                    _ = v58
                }
                _ = fi20
            }
        } else {
            var v59 uint = 2
            for fi60 := 0; fi60 < 6; fi60++ {
                v61 := GenS10{F0: 0, F1: "\n\t"}
                _ = fi60
                _ = v61
            }
            var v62 bool = true
            _ = v59
            _ = v62
        }
        v4 -= v4
    } else {
        var v63 uint64 = 0
        _ = v63
    }
    for fi64 := 0; fi64 < 17; fi64++ {
        v5 -= (v3 >> ((65535) & 63))
        var v65 uint = 4294967295
        _ = fi64
        _ = v65
    }
    if true {
        v3 += (v5 % (((-(v5))) | 1))
    } else {
        v12 += v12
    }
    var v66 int8 = -122
    v67 := GenS11{F0: 255, F1: -361}
    v2 = append(v2, 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v66
    _ = v67
}
