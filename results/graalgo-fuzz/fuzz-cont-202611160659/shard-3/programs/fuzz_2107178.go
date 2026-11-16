package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int }
type GenS2 struct { F0 uint16; F1 uint64; F2 string }
type GenS3 struct { F0 uint32; F1 int16; F2 string }
type GenS4 struct { F0 int16; F1 uint; F2 bool }
type GenS5 struct { F0 int16; F1 int; F2 int32 }
func genFn0() int {
    v5 := GenS1{F0: 32767, F1: "", F2: 42}
    v6 := []int{-1}
    for fi7 := 0; fi7 < 14; fi7++ {
        v6 = append(v6, (fi7 & (fi7 | 1)))
        if true {
            if ((9223372036854775807 > 670) || (255 <= 55)) {
                var v8 int32 = -96
                var v9 bool = true
                var v10 uint8 = 128
                _ = v10
                _ = v8
                _ = v9
            } else {
                var v11 float32 = (1e10 - 0.1)
                if true {
                    v11 += v11
                    var v12 uint16 = 255
                    _ = v12
                }
                v11 += v11
                _ = v11
            }
            switch fi7 {
            case -2147483648:
                for fi13 := 0; fi13 < 11; fi13++ {
                    if false {
                        if true {
                            _ = len(v6)
                        }
                        fmt.Println(fi7, fi13)
                        var v14 float32 = 0.0
                        v6 = append(v6, 128)
                        _ = v14
                    } else {
                        fmt.Println(fi13)
                        var v15 int64 = 999
                        switch fi7 {
                        case 248:
                            _ = len(v6)
                        }
                        _ = v15
                    }
                    _ = fi13
                }
            }
        }
        _ = fi7
    }
    var v16 int32 = -1
    v16 += (v16 + v16)
    _ = v16
    _ = v5
    _ = v6
    return 32768
}
func genFn1(p0 uint8) int16 {
    for fi17 := 0; fi17 < 18; fi17++ {
        p0 = 165
        p0 = p0
        _ = fi17
    }
    if (1 != (p0 | p0)) {
        var v18 string = "hello"
        _ = v18
    } else {
        v19 := GenS1{F0: -533, F1: "a", F2: -9223372036854775808}
        _ = v19
    }
    v20 := GenS1{F0: -2147483648, F1: "hello", F2: -750}
    _ = p0
    _ = v20
    return -307
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{777, -759}
    var v3 uint = 1
    v4 := []uint32{uint32(v3), 854, 2}
    if false {
        for fi21 := 0; fi21 < 14; fi21++ {
            v3 = (v3 | v3)
            for fi22 := 0; fi22 < 19; fi22++ {
                v3 += v3
                switch fi21 {
                case 555:
                    if true {
                        if false {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v3 += v3
                        }
                        var v23 int16 = 1
                        if (-2.25 <= 0.1) {
                            v2 = append(v2, 467)
                            fmt.Println(fi22, v3)
                            v23 = v23
                        } else {
                            v23 += genFn1(128)
                        }
                        _ = v23
                    } else {
                        var v24 int = 9223372036854775807
                        v24 -= fi21
                        _ = v24
                    }
                case -176:
                    v25 := GenS1{F0: (fi22 >> ((0) & 63)), F1: "foo bar", F2: genFn0()}
                    _ = v25
                }
                if true {
                    var v26 int = 9223372036854775807
                    v27 := GenS1{F0: (^(v26)), F1: (("foo bar" + "") + "hello"), F2: (v26 / (((^(v26))) | 1))}
                    v28 := GenS2{F0: 0, F1: 0, F2: ("\n\t" + (("foo bar" + ((("a" + "foo bar") + ("foo bar" + (("" + "a") + "a"))) + ("foo bar" + "\n\t"))) + "\n\t"))}
                    _ = v26
                    _ = v27
                    _ = v28
                } else {
                    var v29 int64 = -311
                    if len(v4) > 0 {
                        v4[0] = 510
                    }
                    _ = len(v2)
                    _ = v29
                }
                _ = fi22
            }
            _ = fi21
        }
        var v30 int16 = -32768
        switch "\n\t" {
        case "":
            for fi31 := 0; fi31 < 7; fi31++ {
                var v32 int32 = 262
                _ = fi31
                _ = v32
            }
        }
        _ = v30
    } else {
        if (false && (!(255 <= 42))) {
            v3 = 790
            switch "a" {
            case "foo bar":
                var v33 uint32 = 4294967295
                _ = v33
            case "hello":
                for fi34 := 0; fi34 < 10; fi34++ {
                    for fi35 := 0; fi35 < 4; fi35++ {
                        if (46 > -32768) {
                            var v36 bool = (false && false)
                            v37 := v1["\n\t"]
                            v37 -= (v37 | (v37 >> ((v3) & 63)))
                            _ = v36
                            _ = v37
                        }
                        var v38 uint = v3
                        _ = fi35
                        _ = v38
                    }
                    if true {
                        if (1 < genFn1(1)) {
                            v3 = (^(v3))
                            v39 := GenS1{F0: fi34, F1: (("" + "a") + "a"), F2: fi34}
                            _ = v39
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v3 = (v3 >> ((v3) & 63))
                        }
                        v3 = v3
                        v40 := GenS3{F0: 42, F1: -1, F2: ("\n\t" + "")}
                        _ = v40
                    }
                    _ = fi34
                }
            default:
                for fi41 := 0; fi41 < 6; fi41++ {
                    for fi42 := 0; fi42 < 7; fi42++ {
                        v3 += v3
                        var v43 uint16 = 2
                        fmt.Println(v3, fi42, fi41)
                        _ = fi42
                        _ = v43
                    }
                    var v44 bool = true
                    fmt.Println(v44, v44)
                    _ = fi41
                    _ = v44
                }
            }
            var v45 float32 = 0.0
            _ = v45
        }
        var v46 int = 746
        _ = v46
    }
    var v47 int16 = 1
    v48 := GenS1{F0: -2147483648, F1: "\n\t", F2: 127}
    if true {
        var v49 int8 = -85
        v1[("hello" + "hello")] = 2147483648
        if ("foo bar" > ("" + "")) {
            fmt.Println(v49)
            _ = len(v1)
        } else {
            if true {
                switch 42 {
                case 4294967295:
                    v50 := GenS2{F0: 285, F1: 2, F2: ("\n\t" + (("" + ("" + "\n\t")) + ("foo bar" + "hello")))}
                    _ = v50
                case 128:
                    fmt.Println(v3)
                case 32767:
                    if (0 == -735) {
                        v51 := GenS4{F0: (v47 >> (((^(v3))) & 63)), F1: 4294967295, F2: ((-536 <= 2147483647) && ((127 != 0) && ((128 > 2) || false)))}
                        for fi52 := 0; fi52 < 6; fi52++ {
                            for fi53 := 0; fi53 < 20; fi53++ {
                                var v54 uint = v3
                                _ = fi53
                                _ = v54
                            }
                            fmt.Println(v47, v47)
                            _ = fi52
                        }
                        fmt.Println(v47)
                        _ = v51
                    }
                }
                switch 32767 {
                case 32767:
                    var v55 bool = false
                    _ = v55
                case 9223372036854775807:
                    var v56 uint8 = 2
                    _ = v56
                case -74:
                    for fi57 := 0; fi57 < 16; fi57++ {
                        var v58 int = int(v3)
                        _ = fi57
                        _ = v58
                    }
                default:
                    v47 = int16(v3)
                }
            }
        }
        _ = v49
    }
    for fi59 := 0; fi59 < 1; fi59++ {
        var v60 int64 = -501
        v3 += v3
        _ = fi59
        _ = v60
    }
    for fi61 := 0; fi61 < 10; fi61++ {
        fmt.Println(v3, v3)
        var v62 int = genFn0()
        _ = fi61
        _ = v62
    }
    for fi63 := 0; fi63 < 14; fi63++ {
        v47 = v47
        switch "\n\t" {
        case "":
            v47 = -374
        case "foo bar":
            for fi64 := 0; fi64 < 19; fi64++ {
                v47 = (-(v47))
                var v65 int32 = 2147483647
                _ = fi64
                _ = v65
            }
        case "\n\t":
            v66 := GenS5{F0: v47, F1: genFn0(), F2: -746}
            _ = v66
        }
        var v67 int8 = 127
        _ = fi63
        _ = v67
    }
    _ = len(v1)
    v47 = genFn1(84)
    v68 := GenS1{F0: 73, F1: ("a" + "a"), F2: -898}
    v47 -= (v47 & (v47 * 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v68
}
