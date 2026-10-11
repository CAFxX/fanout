package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 uint32; F1 uint16; F2 bool }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint8; F1 int }
type GenS9 struct { F0 int16; F1 int16 }
type GenS10 struct { F0 int8; F1 int8; F2 string }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint16 }
func genFn0() int16 {
    for fi5 := 0; fi5 < 12; fi5++ {
        if false {
            for fi6 := 0; fi6 < 9; fi6++ {
                fmt.Println(fi6, fi5)
                _ = fi6
            }
            fmt.Println(fi5)
            for fi7 := 0; fi7 < 3; fi7++ {
                switch "foo bar" {
                case "foo bar":
                    var v8 int16 = -994
                    _ = v8
                case "\n\t":
                    if (107 <= 2) {
                        v9 := GenS1{F0: "\n\t"}
                        var v10 float32 = -2.25
                        _ = v10
                        _ = v9
                    }
                case "a":
                    fmt.Println(fi5)
                }
                _ = fi7
            }
        }
        _ = fi5
    }
    var v11 uint16 = 65535
    v12 := []int{42, 2147483647}
    _ = v11
    _ = v12
    return -1
}

func main() {
    v1 := []string{"foo bar"}
    v2 := make(map[string]bool)
    v3 := make(map[uint32]uint32)
    var v4 bool = (!(-1 != -23))
    switch ("hello" + ("\n\t" + "hello")) {
    case "\n\t":
        if (((0.0 - ((1e10 + float64(1e10)) * (1.5 * float64(1e10)))) + (3.14159 + 0.1)) < (3.14159 / 3.14159)) {
            fmt.Println(v4)
            v4 = false
        }
    case "hello":
        v4 = (false || (1e10 < 1.5))
    }
    v13 := GenS1{F0: "hello"}
    var v14 int = 588
    v4 = (!("\n\t" != ("a" + ("hello" + ("\n\t" + "foo bar")))))
    v14 += 9223372036854775807
    v15 := GenS1{F0: "\n\t"}
    var v16 uint8 = 255
    if v4 {
        v17 := GenS1{F0: ((("foo bar" + "") + "\n\t") + "")}
        for fi18 := 0; fi18 < 19; fi18++ {
            v19 := GenS2{F0: 1, F1: 42}
            _ = fi18
            _ = v19
        }
        var v20 int64 = 1
        _ = v17
        _ = v20
    } else {
        var v21 uint64 = 2
        _ = v21
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        var v23 uint8 = 42
        var v24 string = "hello"
        _ = fi22
        _ = v23
        _ = v24
    }
    if (1e10 < 3.14159) {
        switch v14 {
        case 26:
            v25 := GenS2{F0: -1, F1: -1}
            _ = v25
        case 31:
            var v26 uint = 2
            _ = v26
        default:
            for fi27 := 0; fi27 < 16; fi27++ {
                v16 -= v16
                _ = fi27
            }
        }
        if v4 {
            v28 := GenS3{F0: 271}
            var v29 uint16 = 152
            _ = v28
            _ = v29
        } else {
            v1 = append(v1, ("foo bar" + "foo bar"))
            v4 = v4
            v30 := GenS3{F0: 65535}
            _ = v30
        }
    } else {
        var v31 uint64 = 42
        _ = v31
    }
    if v4 {
        var v32 int = 42
        v33 := GenS4{F0: -32768}
        _ = len(v3)
        _ = v32
        _ = v33
    } else {
        for fi34 := 0; fi34 < 13; fi34++ {
            v35 := GenS2{F0: 100, F1: 0}
            v36 := GenS3{F0: 65535}
            _ = fi34
            _ = v35
            _ = v36
        }
        fmt.Println(v16, v14, v14)
    }
    var v37 int32 = 1
    v16 -= (v16 - v16)
    _ = len(v1)
    if (!v4) {
        delete(v3, 691)
        delete(v3, 0)
        if (127 == -6) {
            v38 := GenS4{F0: 32767}
            fmt.Println(v14, v16, v16)
            if false {
                fmt.Println(v37, v37, v4)
                v1 = append(v1, "\n\t")
                if true {
                    fmt.Println(v16, v14, v14)
                    v4 = v4
                    v4 = (!false)
                } else {
                    var v39 uint16 = 255
                    v14 -= v14
                    switch (v14 | v14) {
                    case 128:
                        v40 := GenS5{F0: int(v37), F1: -9223372036854775808}
                        _ = v40
                    case -9223372036854775808:
                        v16 += uint8(v39)
                    case 32768:
                        for fi41 := 0; fi41 < 13; fi41++ {
                            for fi42 := 0; fi42 < 1; fi42++ {
                                var v43 int8 = 8
                                _ = fi42
                                _ = v43
                            }
                            _ = fi41
                        }
                    default:
                        if false {
                            var v44 int = (v14 + (v14 << ((255) & 63)))
                            var v45 float64 = (1e10 + 0.0)
                            v46 := GenS1{F0: "a"}
                            _ = v44
                            _ = v45
                            _ = v46
                        } else {
                            v4 = (true && true)
                            if (((v4 && v4) || (v4 || (uint64(v16) >= 18446744073709551615))) && true) {
                                fmt.Println(v4, v16, v39)
                                for fi47 := 0; fi47 < 19; fi47++ {
                                    v48 := GenS4{F0: -917}
                                    var v49 uint = 2
                                    fmt.Println(v37, v16)
                                    _ = fi47
                                    _ = v48
                                    _ = v49
                                }
                                for fi50 := 0; fi50 < 3; fi50++ {
                                    for fi51 := 0; fi51 < 12; fi51++ {
                                        fmt.Println(v4, fi51, v16)
                                        _ = fi51
                                    }
                                    _ = fi50
                                }
                            }
                        }
                    }
                    _ = v39
                }
            }
            _ = v38
        }
    } else {
        for fi52 := 0; fi52 < 15; fi52++ {
            v53 := v2["\n\t"]
            _ = len(v2)
            v54 := GenS6{F0: 1, F1: 1, F2: v4}
            _ = fi52
            _ = v53
            _ = v54
        }
        v55 := GenS7{F0: (v37 % ((1) | 1))}
        _ = v55
    }
    for fi56 := 0; fi56 < 17; fi56++ {
        switch v14 {
        case -9223372036854775808:
            for fi57 := 0; fi57 < 16; fi57++ {
                var v58 int64 = -1
                _ = fi57
                _ = v58
            }
        case -15:
            switch 127 {
            case 4294967295:
                fmt.Println(v16)
            case 32768:
                if (int16(fi56) >= -50) {
                    fmt.Println(v14, fi56, fi56)
                    v59 := GenS1{F0: "a"}
                    v37 = (v37 << ((65535) & 63))
                    _ = v59
                }
            }
        case 4294967295:
            switch "a" {
            case "hello":
                switch "a" {
                case "hello":
                    switch "" {
                    case "\n\t":
                        if (9223372036854775807 >= 0) {
                            if v4 {
                                v60 := GenS7{F0: -2147483648}
                                _ = v60
                            } else {
                                var v61 bool = false
                                _ = v61
                            }
                            var v62 float64 = 1.5
                            delete(v2, (("hello" + "foo bar") + ("foo bar" + "foo bar")))
                            _ = v62
                        }
                    case "foo bar":
                        v16 -= v16
                    case "":
                        var v63 uint8 = v16
                        _ = v63
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                case "":
                    _ = len(v1)
                case "a":
                    v64 := GenS8{F0: 255, F1: fi56}
                    _ = v64
                }
            default:
                for fi65 := 0; fi65 < 7; fi65++ {
                    var v66 uint32 = 431
                    v37 = v37
                    v67 := GenS9{F0: -32768, F1: 0}
                    var v68 uint16 = 255
                    if v4 {
                        if len(v1) > 0 {
                            v1[0] = "\n\t"
                        }
                        v69 := GenS3{F0: 65535}
                        fmt.Println(v16)
                        _ = v69
                    } else {
                        switch fi65 {
                        case -891:
                            v70 := GenS10{F0: -20, F1: -27, F2: ("hello" + ("\n\t" + "a"))}
                            v66 = 1
                            _ = v70
                        case 445:
                            v14 = v14
                        case -9223372036854775808:
                            switch ("a" + "\n\t") {
                            case "hello":
                                _ = len(v2)
                            default:
                                v71 := GenS11{F0: -199}
                                _ = v71
                            }
                        default:
                            v66 = 608
                            if v4 {
                                v4 = v4
                                v3[v66] = (v66 | 65535)
                            } else {
                                if false {
                                    fmt.Println(v4, v14)
                                } else {
                                    if v4 {
                                        var v72 uint16 = (v68 / (((^(v68))) | 1))
                                        fmt.Println(v14, v66)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v72
                                    }
                                    if (v4 && (false || (("hello" + ("hello" + "hello")) != ("" + "a")))) {
                                        fmt.Println(v68)
                                        if len(v1) > 0 {
                                            v1[0] = "a"
                                        }
                                    } else {
                                        v73 := GenS6{F0: (v66 * 464), F1: v68, F2: v4}
                                        fmt.Println(v16)
                                        fmt.Println(fi56, fi65)
                                        v14 = v14
                                        _ = v73
                                    }
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v74 int8 = 42
                                _ = v74
                            }
                            fmt.Println(v4)
                            v4 = true
                        }
                    }
                    _ = fi65
                    _ = v66
                    _ = v67
                    _ = v68
                }
            }
        default:
            switch 127 {
            case 1:
                if (!(1000 != -1)) {
                    _ = len(v2)
                    for fi75 := 0; fi75 < 20; fi75++ {
                        switch (fi75 << ((0) & 63)) {
                        case 570:
                            v1 = append(v1, "\n\t")
                        }
                        var v76 uint32 = 4294967295
                        _ = fi75
                        _ = v76
                    }
                }
            }
        }
        v77 := GenS12{F0: 32767}
        for fi78 := 0; fi78 < 4; fi78++ {
            fmt.Println(v4, v16)
            _ = fi78
        }
        _ = fi56
        _ = v77
    }
    v16 += v16
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v37
    _ = v4
}
