package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 string }
type GenS2 struct { F0 int; F1 int16; F2 string }
type GenS3 struct { F0 int64; F1 uint16 }
func genFn0() bool {
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 int16 = -113
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS1{F0: 2147483648, F1: "\n\t", F2: "foo bar"}
        if true {
            if (false && true) {
                for fi11 := 0; fi11 < 4; fi11++ {
                    var v12 int8 = 63
                    _ = fi11
                    _ = v12
                }
                var v13 int32 = -2147483648
                _ = v13
            }
            var v14 uint32 = 4294967295
            _ = v14
        }
        _ = fi9
        _ = v10
    }
    var v15 uint = 365
    _ = v15
    return false
}
func genFn1(p0 string) uint16 {
    p0 = p0
    if false {
        var v16 int32 = 1
        if (2147483647 > 410) {
            v17 := []int{-70, -324, 211}
            for fi18 := 0; fi18 < 8; fi18++ {
                for fi19 := 0; fi19 < 16; fi19++ {
                    v20 := GenS1{F0: 0, F1: p0, F2: p0}
                    if false {
                        fmt.Println(v16, fi18, fi18)
                        var v21 uint = 613
                        _ = v21
                    }
                    v17 = append(v17, 128)
                    _ = fi19
                    _ = v20
                }
                v22 := GenS1{F0: 65535, F1: "a", F2: "hello"}
                _ = fi18
                _ = v22
            }
            _ = v17
        } else {
            v16 = v16
            var v23 int8 = 1
            switch 0 {
            case 0:
                var v24 int16 = 1000
                _ = v24
            default:
                p0 = "foo bar"
            }
            _ = v23
        }
        if false {
            for fi25 := 0; fi25 < 10; fi25++ {
                v26 := []int{fi25, fi25, fi25}
                v26 = append(v26, (^(fi25)))
                _ = fi25
                _ = v26
            }
            for fi27 := 0; fi27 < 2; fi27++ {
                v28 := []int{32768}
                var v29 int = fi27
                for fi30 := 0; fi30 < 17; fi30++ {
                    v16 += v16
                    _ = fi30
                }
                _ = fi27
                _ = v28
                _ = v29
            }
        }
        _ = v16
    }
    v31 := []int{-2147483648, -343, 119}
    p0 = (p0 + p0)
    _ = p0
    _ = v31
    return 0
}

func main() {
    v1 := []int{}
    v2 := []uint32{2147483647}
    v3 := []uint32{859, 42, 158}
    var v4 bool = (false && true)
    v5 := []string{("a" + "\n\t")}
    v6 := []int{32768, 4294967295}
    v4 = ((v4 && v4) || v4)
    switch 127 {
    case 4294967295:
        var v32 int32 = 42
        _ = v32
    case 9223372036854775807:
        var v33 float64 = 1.5
        _ = v33
    default:
        var v34 float64 = ((float64(((1.5 * -2.25) - (1e10 + 0.1))) * 1e10) - 1.5)
        _ = v34
    }
    var v35 bool = (267 == 128)
    _ = len(v2)
    var v36 int = 4294967295
    for fi37 := 0; fi37 < 14; fi37++ {
        v35 = false
        v38 := GenS1{F0: 635, F1: (((("foo bar" + (("a" + "a") + "hello")) + "foo bar") + "") + "a"), F2: "foo bar"}
        _ = fi37
        _ = v38
    }
    v6 = append(v6, 127)
    var v39 uint = 255
    if len(v1) > 0 {
        v1[0] = v36
    }
    for fi40 := 0; fi40 < 3; fi40++ {
        v35 = false
        if (!v4) {
            if v4 {
                fmt.Println(v4)
            }
        }
        var v41 string = "\n\t"
        _ = fi40
        _ = v41
    }
    for fi42 := 0; fi42 < 12; fi42++ {
        if false {
            v36 -= (fi42 >> (((^(v39))) & 63))
        } else {
            var v43 float32 = -2.25
            for fi44 := 0; fi44 < 10; fi44++ {
                fmt.Println(fi44, v36, fi44)
                if len(v5) > 0 {
                    v5[0] = "foo bar"
                }
                var v45 uint32 = 1
                _ = fi44
                _ = v45
            }
            var v46 int = fi42
            _ = v43
            _ = v46
        }
        if (1.5 >= (3.14159 / float64((1.5 - (float32((1.5 / (1.5 / 3.14159))) + 1.5))))) {
            v4 = v35
            fmt.Println(fi42, v35, v4)
        } else {
            var v47 float64 = 3.14159
            for fi48 := 0; fi48 < 7; fi48++ {
                _ = len(v1)
                if (2 <= 392) {
                    _ = len(v3)
                    var v49 uint32 = 357
                    _ = v49
                }
                _ = fi48
            }
            v50 := GenS1{F0: 6, F1: "hello", F2: "a"}
            _ = v47
            _ = v50
        }
        _ = fi42
    }
    var v51 uint32 = 904
    var v52 float32 = 1.5
    if (true || true) {
        if len(v3) > 0 {
            v3[0] = v51
        }
        if (0 > 42) {
            for fi53 := 0; fi53 < 20; fi53++ {
                var v54 bool = false
                _ = fi53
                _ = v54
            }
            switch v36 {
            case 128:
                v55 := GenS1{F0: v51, F1: ("hello" + "a"), F2: ("foo bar" + (("" + ("hello" + "foo bar")) + "\n\t"))}
                _ = v55
            default:
                for fi56 := 0; fi56 < 8; fi56++ {
                    v57 := GenS1{F0: (v51 + v51), F1: "", F2: ((((("foo bar" + "") + "a") + "") + "") + "")}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    for fi58 := 0; fi58 < 18; fi58++ {
                        v59 := GenS2{F0: fi56, F1: -992, F2: "\n\t"}
                        var v60 int64 = 1
                        _ = fi58
                        _ = v59
                        _ = v60
                    }
                    _ = fi56
                    _ = v57
                }
            }
            if true {
                var v61 int64 = 42
                var v62 string = ""
                _ = v61
                _ = v62
            } else {
                v4 = v35
                if v35 {
                    var v63 uint16 = genFn1("foo bar")
                    _ = v63
                } else {
                    for fi64 := 0; fi64 < 11; fi64++ {
                        var v65 float32 = (0.0 + v52)
                        for fi66 := 0; fi66 < 18; fi66++ {
                            var v67 uint8 = 1
                            switch 4294967295 {
                            case 837:
                                v67 = v67
                            }
                            switch ((("foo bar" + "") + "foo bar") + ((("\n\t" + "") + "") + "foo bar")) {
                            case "a":
                                v35 = (801 != -32768)
                            case "foo bar":
                                v65 = (float32(float64(-2.25)) * v52)
                            default:
                                var v68 int16 = 1000
                                _ = v68
                            }
                            var v69 float32 = 0.1
                            var v70 int32 = 0
                            v65 = (v69 / 1.5)
                            v65 = (v52 + (v65 * (v52 - 0.0)))
                            v71 := GenS3{F0: -1, F1: 1}
                            _ = fi66
                            _ = v67
                            _ = v69
                            _ = v70
                            _ = v71
                        }
                        _ = fi64
                        _ = v65
                    }
                }
                var v72 uint32 = (v51 - v51)
                _ = v72
            }
        } else {
            v73 := GenS2{F0: 9223372036854775807, F1: -384, F2: ("" + "hello")}
            v74 := GenS2{F0: v36, F1: 0, F2: ((("\n\t" + "foo bar") + "a") + "\n\t")}
            fmt.Println(v35, v36)
            _ = v73
            _ = v74
        }
        var v75 int8 = 42
        _ = v75
    } else {
        v36 = v36
        for fi76 := 0; fi76 < 6; fi76++ {
            v77 := GenS3{F0: 0, F1: 255}
            if (4294967295 > v39) {
                var v78 uint64 = 1
                if len(v6) > 0 {
                    v6[0] = (fi76 & (fi76 + v36))
                }
                _ = v78
            } else {
                fmt.Println(v4)
                _ = len(v6)
            }
            var v79 uint = v39
            _ = fi76
            _ = v77
            _ = v79
        }
        for fi80 := 0; fi80 < 3; fi80++ {
            if v4 {
                var v81 uint16 = genFn1("foo bar")
                _ = v81
            }
            for fi82 := 0; fi82 < 11; fi82++ {
                if v4 {
                    v83 := GenS1{F0: 226, F1: "\n\t", F2: "a"}
                    v36 = fi80
                    _ = v83
                } else {
                    v35 = v4
                }
                _ = fi82
            }
            v4 = (!true)
            _ = fi80
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
