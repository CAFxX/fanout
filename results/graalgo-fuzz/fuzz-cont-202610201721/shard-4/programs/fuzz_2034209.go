package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint }
type GenS3 struct { F0 uint; F1 int; F2 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int64; F1 uint16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int32; F1 uint16 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int8; F1 int16 }
func genFn0() int8 {
    switch "foo bar" {
    case "foo bar":
        var v4 string = "foo bar"
        _ = v4
    case "hello":
        for fi5 := 0; fi5 < 2; fi5++ {
            var v6 int16 = 42
            for fi7 := 0; fi7 < 20; fi7++ {
                if (((!(!(float64(-2.25) < float64(1.5)))) && true) && true) {
                    switch "foo bar" {
                    case "foo bar":
                        v6 = (v6 + v6)
                    case "hello":
                        if (578 >= 735) {
                            v6 = (v6 / (((v6 & (v6 << ((0) & 63)))) | 1))
                            v6 -= (^(v6))
                            switch fi5 {
                            case 128:
                                var v8 bool = false
                                _ = v8
                            case 937:
                                switch "a" {
                                case "\n\t":
                                    var v9 uint32 = 2147483648
                                    _ = v9
                                case "a":
                                    if (0 > 2147483647) {
                                        var v10 float64 = float64((1e10 / ((-2.25 / 1e10) - (float32(0.1) * (-2.25 - -2.25)))))
                                        _ = v10
                                    }
                                case "hello":
                                    v6 -= v6
                                default:
                                    v6 += (v6 | v6)
                                }
                            case 42:
                                v11 := GenS1{F0: ("hello" + "\n\t"), F1: 255}
                                _ = v11
                            default:
                                v12 := []int{128, -2147483648}
                                _ = v12
                            }
                        } else {
                            v6 += (-(v6))
                        }
                        if (-1 <= 127) {
                            v13 := []int{-1, 32767}
                            v14 := GenS2{F0: -9223372036854775808, F1: 0, F2: 42}
                            _ = v13
                            _ = v14
                        }
                    }
                    v15 := GenS2{F0: 9223372036854775807, F1: 65535, F2: 333}
                    _ = v15
                }
                var v16 uint32 = 527
                fmt.Println(fi5)
                v6 = 42
                v6 = (v6 % ((v6) | 1))
                _ = fi7
                _ = v16
            }
            fmt.Println(fi5, v6)
            _ = fi5
            _ = v6
        }
    case "a":
        var v17 uint64 = 2
        _ = v17
    default:
        var v18 uint32 = 1
        _ = v18
    }
    var v19 int64 = -9223372036854775808
    v19 = 42
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 int16 = 314
        var v22 int16 = (v21 | 11)
        if ((("\n\t" + "foo bar") + ("" + "a")) >= (("hello" + ("foo bar" + (("\n\t" + "a") + "\n\t"))) + ("a" + ("" + ((("hello" + "") + ((("hello" + "") + ("" + (("foo bar" + "\n\t") + ("hello" + "a")))) + "")) + ""))))) {
            for fi23 := 0; fi23 < 7; fi23++ {
                fmt.Println(v21, v22)
                switch ("hello" + "") {
                case "foo bar":
                    var v24 uint = 0
                    var v25 int = 42
                    var v26 bool = false
                    v27 := []int{}
                    if len(v27) > 0 {
                        v27[0] = (-(fi23))
                    }
                    if v26 {
                        v28 := GenS1{F0: "a", F1: 77}
                        v26 = v26
                        switch ("foo bar" + "\n\t") {
                        case "foo bar":
                            _ = len(v27)
                        case "hello":
                            if v26 {
                                if true {
                                    v22 += (^(v22))
                                    v29 := GenS1{F0: "hello", F1: 120}
                                    if len(v27) >= 2 {
                                        v27 = v27[1:len(v27)]
                                    }
                                    _ = v29
                                }
                                fmt.Println(v21, v22)
                                v25 = (v25 % ((700) | 1))
                            } else {
                                v30 := GenS3{F0: (^(v24)), F1: 2147483647, F2: 42}
                                var v31 int8 = 100
                                _ = v30
                                _ = v31
                            }
                        default:
                            fmt.Println(v21, fi20, v19)
                        }
                        var v32 int32 = -922
                        _ = v28
                        _ = v32
                    } else {
                        switch "foo bar" {
                        case "\n\t":
                            var v33 uint32 = 2
                            _ = v33
                        case "foo bar":
                            var v34 uint32 = 197
                            _ = v34
                        default:
                            v27 = append(v27, v25)
                        }
                    }
                    _ = v24
                    _ = v25
                    _ = v26
                    _ = v27
                case "":
                    v35 := []int{(fi20 & (fi23 | 127)), (fi23 * (-(fi20)))}
                    _ = v35
                case "hello":
                    var v36 float32 = 1e10
                    _ = v36
                default:
                    v21 += -32768
                }
                _ = fi23
            }
        }
        _ = fi20
        _ = v21
        _ = v22
    }
    _ = v19
    return -24
}

func main() {
    var v1 string = "a"
    v2 := make(map[int]string)
    v3 := []string{v1, v1, v1}
    var v37 uint = 836
    var v38 int = 0
    v1 = (("hello" + (v1 + (v1 + "a"))) + (v1 + (v1 + ("\n\t" + ""))))
    var v39 uint32 = 2147483647
    for fi40 := 0; fi40 < 4; fi40++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v37 = 18446744073709551615
        for fi41 := 0; fi41 < 17; fi41++ {
            v37 = (v37 / ((v37) | 1))
            var v42 uint16 = 0
            delete(v2, fi41)
            _ = fi41
            _ = v42
        }
        _ = fi40
    }
    _ = len(v3)
    for fi43 := 0; fi43 < 15; fi43++ {
        if (0.1 != 1e10) {
            v44 := GenS3{F0: v37, F1: -954, F2: 100}
            v45 := GenS1{F0: "a", F1: 127}
            var v46 bool = ((-32768 < 32767) || (1 < 4294967295))
            _ = v44
            _ = v45
            _ = v46
        }
        if true {
            var v47 uint64 = 401
            for fi48 := 0; fi48 < 18; fi48++ {
                var v49 int32 = 1
                switch v1 {
                case "hello":
                    if true {
                        fmt.Println(fi48, v38, fi48)
                        v50 := GenS4{F0: (v49 + (v49 - (v49 - (v49 + (v49 % ((v49) | 1))))))}
                        _ = v50
                    } else {
                        switch fi43 {
                        case 32768:
                            switch v1 {
                            case "a":
                                var v51 uint8 = 156
                                _ = v51
                            case "":
                                v38 = 2147483647
                            case "foo bar":
                                if (!(false || true)) {
                                    if (v47 < (v47 / ((v47) | 1))) {
                                        v47 -= v47
                                        if (664 < 0) {
                                            var v52 uint16 = 2
                                            _ = v52
                                        }
                                        var v53 int64 = 0
                                        _ = v53
                                    }
                                    if (-1 <= 1000) {
                                        if true {
                                            v54 := GenS5{F0: (1 < uint8(v39))}
                                            var v55 float32 = ((float32(-2.25) * -2.25) - -2.25)
                                            _ = v54
                                            _ = v55
                                        } else {
                                            if (("\n\t" + v1) != (("a" + "\n\t") + v1)) {
                                                switch fi43 {
                                                case 4294967295:
                                                    var v56 uint32 = v39
                                                    _ = v56
                                                case -2147483648:
                                                    v57 := GenS2{F0: 42, F1: 255, F2: (v37 & 2)}
                                                    _ = v57
                                                default:
                                                    if false {
                                                        v58 := GenS1{F0: v1, F1: 0}
                                                        v59 := GenS6{F0: -697, F1: 2}
                                                        v3 = append(v3, v1)
                                                        v60 := v2[int(v47)]
                                                        v2[(fi48 + fi48)] = "\n\t"
                                                        _ = v58
                                                        _ = v59
                                                        _ = v60
                                                    }
                                                }
                                            } else {
                                                v61 := GenS7{F0: v38}
                                                v47 -= v47
                                                _ = v61
                                            }
                                            v62 := GenS4{F0: 2147483647}
                                            v1 = ((v1 + (v1 + "a")) + v1)
                                            _ = v62
                                        }
                                        v49 = v49
                                    } else {
                                        v63 := GenS4{F0: v49}
                                        if true {
                                            v64 := GenS8{F0: (v49 + v49), F1: 1}
                                            _ = v64
                                        }
                                        v3 = append(v3, ((v1 + "\n\t") + v1))
                                        _ = v63
                                    }
                                    fmt.Println(v49)
                                } else {
                                    fmt.Println(fi48)
                                }
                            }
                        case 128:
                            v3 = append(v3, v1)
                            v39 += (v39 ^ (v39 * (v39 + v39)))
                            fmt.Println(v49, fi48, fi48)
                            if len(v3) > 0 {
                                v3[0] = "foo bar"
                            }
                            v65 := GenS9{F0: -9223372036854775808}
                            v49 += v49
                            _ = v65
                        case -2147483648:
                            var v66 uint = v37
                            _ = v66
                        }
                    }
                }
                var v67 uint8 = 0
                _ = fi48
                _ = v49
                _ = v67
            }
            _ = v47
        }
        _ = fi43
    }
    var v68 uint16 = 404
    fmt.Println(v38, v1, v68)
    v69 := v2[-9223372036854775808]
    v70 := GenS2{F0: -9223372036854775808, F1: 1, F2: v37}
    v39 += (v39 ^ uint32(v68))
    v71 := GenS10{F0: 25, F1: 385}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v68
    _ = v69
    _ = v70
    _ = v71
}
