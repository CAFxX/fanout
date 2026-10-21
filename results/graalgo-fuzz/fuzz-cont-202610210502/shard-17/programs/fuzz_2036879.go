package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint16; F1 int64; F2 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int; F1 uint16; F2 uint16 }
func genFn0(p0 string, p1 uint16, p2 string) uint32 {
    fmt.Println(p0, p1, p1)
    if true {
        v4 := GenS1{F0: 0, F1: (2147483648 != 186)}
        _ = v4
    } else {
        for fi5 := 0; fi5 < 6; fi5++ {
            var v6 uint16 = 65535
            _ = fi5
            _ = v6
        }
    }
    p0 = "foo bar"
    switch p0 {
    case "a":
        var v7 string = p2
        _ = v7
    case "":
        if false {
            v8 := GenS1{F0: -1, F1: (!((((100 == 100) || false) && ((false && true) && true)) && false))}
            fmt.Println(p2, p2, p0)
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 uint64 = 42
                _ = fi9
                _ = v10
            }
            _ = v8
        }
    case "\n\t":
        if ((p2 + "hello") == (p0 + (p2 + "a"))) {
            p0 = p0
            switch 42 {
            case 0:
                var v11 int = -1
                _ = v11
            }
            p0 = p0
        }
    }
    v12 := GenS1{F0: -89, F1: (!(2 == 255))}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 157
}
func genFn1(p0 int) uint {
    v13 := []int{}
    if false {
        var v14 string = (("foo bar" + "a") + "")
        _ = v14
    } else {
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
    }
    var v15 int16 = -33
    for fi16 := 0; fi16 < 19; fi16++ {
        v17 := GenS1{F0: (v15 * (v15 | -541)), F1: (0 <= -9223372036854775808)}
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = v13
    _ = v15
    return 984
}

func main() {
    v1 := []int{2147483647, 42}
    var v2 bool = true
    var v3 uint64 = 0
    if (!false) {
        var v18 string = ""
        v3 += v3
        _ = v18
    }
    v19 := GenS1{F0: 32767, F1: ((32767 <= -32768) && true)}
    v20 := GenS1{F0: 42, F1: v2}
    var v21 int = 42
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 int8 = -63
        _ = fi22
        _ = v23
    }
    if true {
        if (false && true) {
            var v24 uint8 = 127
            for fi25 := 0; fi25 < 6; fi25++ {
                if len(v1) > 0 {
                    v1[0] = (v21 % (((fi25 * 1)) | 1))
                }
                var v26 string = (("hello" + "a") + "a")
                v27 := GenS2{F0: 2, F1: -580, F2: 32767}
                _ = fi25
                _ = v26
                _ = v27
            }
            if false {
                v28 := GenS3{F0: -9223372036854775808}
                if (!(true || v2)) {
                    fmt.Println(v3, v21, v24)
                }
                switch (^(v21)) {
                case 4294967295:
                    var v29 uint32 = 42
                    _ = v29
                case 32768:
                    v2 = ((v2 && v2) || true)
                default:
                    if len(v1) > 0 {
                        v1[0] = v21
                    }
                }
                _ = v28
            }
            _ = v24
        }
        v3 -= v3
    } else {
        v2 = (genFn0(("hello" + "foo bar"), 42, "") == 410)
    }
    var v30 string = ("hello" + ("\n\t" + "hello"))
    if len(v1) > 0 {
        v1[0] = 32767
    }
    if (((!(v2 || (4294967295 != 395))) && ((0 != 1) || (("a" + v30) == "\n\t"))) || true) {
        v30 = v30
        for fi31 := 0; fi31 < 15; fi31++ {
            v3 = (v3 ^ v3)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if v2 {
                v3 = (v3 & (v3 >> ((1) & 63)))
            } else {
                _ = len(v1)
            }
            _ = fi31
        }
        fmt.Println(v30, v21, v21)
    } else {
        if (100 != -1) {
            var v32 uint32 = 556
            _ = v32
        } else {
            v33 := GenS4{F0: 42}
            _ = v33
        }
        var v34 int8 = -128
        var v35 float64 = -2.25
        _ = v34
        _ = v35
    }
    if (4294967295 != v21) {
        v36 := GenS5{F0: -57}
        if (v2 || (42 >= 199)) {
            switch (v21 * 127) {
            case 4294967295:
                fmt.Println(v30, v21)
            }
        } else {
            for fi37 := 0; fi37 < 10; fi37++ {
                var v38 uint32 = genFn0(("hello" + "\n\t"), 255, "foo bar")
                v38 -= (v38 * (v38 >> ((65535) & 63)))
                _ = fi37
                _ = v38
            }
            v21 -= v21
        }
        _ = v36
    } else {
        v3 = v3
        if len(v1) > 0 {
            v1[0] = v21
        }
        if len(v1) > 0 {
            v1[0] = 127
        }
    }
    v3 += 4294967295
    if (v3 == 42) {
        if v2 {
            if v2 {
                v1 = append(v1, v21)
                var v39 uint32 = 4294967295
                for fi40 := 0; fi40 < 18; fi40++ {
                    v3 -= (v3 >> ((255) & 63))
                    var v41 bool = false
                    _ = fi40
                    _ = v41
                }
                _ = v39
            } else {
                var v42 uint32 = 1
                _ = v42
            }
            v2 = true
        }
        if false {
            if len(v1) > 0 {
                v1[0] = v21
            }
            var v43 uint8 = 64
            v44 := GenS1{F0: 307, F1: false}
            _ = v43
            _ = v44
        } else {
            fmt.Println(v21)
            if true {
                var v45 float64 = float64(float32((0.1 + 0.1)))
                _ = v45
            } else {
                v30 = "hello"
                v46 := GenS1{F0: 32767, F1: v2}
                var v47 float64 = 1e10
                _ = v46
                _ = v47
            }
        }
    }
    v30 = v30
    v3 -= v3
    if (false && v2) {
        v2 = (100 == 100)
        if v2 {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            if (1000 <= 1000) {
                v3 -= (v3 * 42)
                fmt.Println(v21, v3, v21)
            }
        }
        if ((2147483647 > -1) && v2) {
            if (((127 != 128) || v2) || v2) {
                switch v21 {
                case 496:
                    v3 = (v3 % (((v3 + 42)) | 1))
                }
                switch 145 {
                case 9223372036854775807:
                    if (((0.1 * 0.0) * (float64((0.0 / (1e10 / 0.1))) + (float64(0.1) + -2.25))) != ((3.14159 / float64(-2.25)) / 3.14159)) {
                        var v48 int64 = 42
                        _ = v48
                    } else {
                        v3 += (v3 - (v3 / ((v3) | 1)))
                    }
                case -9223372036854775808:
                    v2 = v2
                case 1:
                    v2 = v2
                }
                _ = len(v1)
            } else {
                var v49 bool = true
                v50 := GenS4{F0: 128}
                v51 := GenS1{F0: 1000, F1: v49}
                _ = v49
                _ = v50
                _ = v51
            }
        } else {
            for fi52 := 0; fi52 < 4; fi52++ {
                if false {
                    var v53 int = (fi52 - fi52)
                    for fi54 := 0; fi54 < 8; fi54++ {
                        var v55 float32 = ((-2.25 + float32(1.5)) - (0.1 - 1e10))
                        v56 := GenS5{F0: -1}
                        switch v30 {
                        case "foo bar":
                            switch (fi52 << ((65535) & 63)) {
                            case 9223372036854775807:
                                if (fi52 < v21) {
                                    v53 -= (v21 * 0)
                                    var v57 uint8 = 128
                                    _ = v57
                                }
                            case 2147483647:
                                v1 = append(v1, (v53 + (fi54 >> ((195) & 63))))
                            case 127:
                                switch (fi52 - fi52) {
                                case 42:
                                    var v58 int16 = 1000
                                    _ = v58
                                case 2147483647:
                                    v2 = v2
                                }
                            }
                        case "\n\t":
                            v2 = (v30 < "hello")
                        default:
                            switch fi54 {
                            case 2147483647:
                                var v59 uint = 4294967295
                                _ = v59
                            default:
                                v60 := GenS6{F0: fi54, F1: 739, F2: 1}
                                _ = v60
                            }
                        }
                        _ = fi54
                        _ = v55
                        _ = v56
                    }
                    switch v53 {
                    case 379:
                        for fi61 := 0; fi61 < 16; fi61++ {
                            fmt.Println(v2)
                            v2 = (0.1 <= 0.0)
                            _ = len(v1)
                            v21 = v21
                            _ = fi61
                        }
                    case 32768:
                        if v2 {
                            var v62 uint = 728
                            for fi63 := 0; fi63 < 14; fi63++ {
                                var v64 int16 = -257
                                _ = fi63
                                _ = v64
                            }
                            _ = v62
                        }
                    default:
                        if true {
                            var v65 uint32 = 537
                            _ = v65
                        }
                    }
                    _ = v53
                }
                _ = fi52
            }
            fmt.Println(v21, v3, v3)
        }
    } else {
        var v66 float32 = 3.14159
        var v67 float32 = v66
        _ = v66
        _ = v67
    }
    for fi68 := 0; fi68 < 8; fi68++ {
        if (!v2) {
            for fi69 := 0; fi69 < 10; fi69++ {
                v70 := GenS6{F0: 32767, F1: 32767, F2: 0}
                var v71 int = (v21 / (((fi69 | 684)) | 1))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi69
                _ = v70
                _ = v71
            }
        } else {
            switch v30 {
            case "":
                var v72 bool = v2
                _ = v72
            case "hello":
                if len(v1) > 0 {
                    v1[0] = (v21 | (v21 << ((2) & 63)))
                }
            case "a":
                v30 = (v30 + "a")
            }
            v1 = append(v1, fi68)
        }
        v21 = v21
        for fi73 := 0; fi73 < 3; fi73++ {
            v21 = (fi68 - fi68)
            _ = fi73
        }
        _ = fi68
    }
    v2 = true
    v74 := GenS5{F0: 596}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v74
}
