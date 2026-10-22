package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS10 struct { F0 uint; F1 uint32 }
func genFn0(p0 int64) uint16 {
    p0 = -4
    var v4 int32 = 570
    var v5 uint8 = 128
    v6 := []int{-901, 0, 128, -9223372036854775808}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 32767
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]int)
    v3 := []int{}
    for fi7 := 0; fi7 < 16; fi7++ {
        v2[65535] = 1
        var v8 uint32 = 2147483647
        var v9 int32 = 203
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint8 = 236
    v11 := GenS1{F0: -192, F1: 370}
    var v12 uint8 = 27
    v13 := GenS2{F0: 0}
    if ((v10 | (^(v12))) < (^(v10))) {
        var v14 float64 = 0.0
        _ = len(v1)
        delete(v2, 0)
        _ = v14
    }
    switch (("\n\t" + ("a" + "")) + "") {
    case "a":
        if ((3.14159 * (1e10 - 0.0)) <= 0.0) {
            var v15 int16 = -32768
            v16 := GenS3{F0: 2147483647}
            for fi17 := 0; fi17 < 5; fi17++ {
                var v18 uint8 = 42
                var v19 int8 = 1
                switch int(v15) {
                case 9223372036854775807:
                    if (!true) {
                        if (0.0 != 0.0) {
                            fmt.Println(v15, v10, v19)
                            if len(v3) > 0 {
                                v3[0] = -1
                            }
                            for fi20 := 0; fi20 < 20; fi20++ {
                                if (v18 > (v10 >> ((870) & 63))) {
                                    if (!(-214 < 1)) {
                                        if (0 >= int64(v15)) {
                                            v10 = v10
                                        }
                                        v19 = v19
                                        var v21 uint = 255
                                        _ = v21
                                    } else {
                                        v22 := GenS4{F0: true, F1: v18}
                                        _ = v22
                                    }
                                    var v23 uint8 = 39
                                    v10 += v18
                                    v15 = 1
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v23
                                } else {
                                    fmt.Println(fi20, v15)
                                    if (535 == 18446744073709551615) {
                                        if true {
                                            fmt.Println(v15, fi17, fi20)
                                            fmt.Println(v15, fi20)
                                        }
                                    }
                                    v24 := GenS2{F0: 0}
                                    _ = v24
                                }
                                v25 := GenS5{F0: v18, F1: (^(v10))}
                                _ = fi20
                                _ = v25
                            }
                        } else {
                            var v26 int = int(v10)
                            switch v26 {
                            case 32768:
                                var v27 uint64 = 18446744073709551615
                                _ = v27
                            case 0:
                                var v28 int64 = -9223372036854775808
                                _ = v28
                            case 2147483647:
                                for fi29 := 0; fi29 < 13; fi29++ {
                                    v30 := GenS3{F0: -2147483648}
                                    switch "a" {
                                    case "a":
                                        if true {
                                            var v31 string = "a"
                                            var v32 uint32 = uint32(v18)
                                            switch "foo bar" {
                                            case "":
                                                if len(v3) > 0 {
                                                    v3[0] = (fi17 ^ (fi29 % ((v26) | 1)))
                                                }
                                            case "\n\t":
                                                v12 += v12
                                            default:
                                                switch v26 {
                                                case 128:
                                                    fmt.Println(fi17)
                                                default:
                                                    v31 = v31
                                                }
                                            }
                                            _ = v31
                                            _ = v32
                                        } else {
                                            v26 -= fi29
                                            var v33 int = -2147483648
                                            if (73 != 42) {
                                                switch v26 {
                                                case 42:
                                                    delete(v1, ((("hello" + "\n\t") + (("a" + "a") + ((("" + "\n\t") + ("" + (("\n\t" + ("foo bar" + "hello")) + (("\n\t" + "foo bar") + ("hello" + ""))))) + "hello"))) + ("\n\t" + ("a" + "foo bar"))))
                                                default:
                                                    if (2 >= 0) {
                                                        switch (("\n\t" + "\n\t") + "") {
                                                        case "a":
                                                            v34 := GenS4{F0: true, F1: (v12 / ((2) | 1))}
                                                            _ = v34
                                                        case "":
                                                            var v35 int8 = (-(v19))
                                                            _ = v35
                                                        }
                                                        v36 := GenS2{F0: uint16(v12)}
                                                        v19 += -1
                                                        _ = v36
                                                    }
                                                }
                                            }
                                            v3 = append(v3, (v26 + fi17))
                                            _ = v33
                                        }
                                    default:
                                        if (128 != v12) {
                                            if true {
                                                v10 -= 128
                                                var v37 bool = ((v15 + 32767) == v15)
                                                _ = v37
                                            }
                                        } else {
                                            var v38 uint32 = 42
                                            if true {
                                                v26 = fi29
                                                var v39 float32 = ((0.1 + -2.25) - 1.5)
                                                v40 := GenS6{F0: 9223372036854775807, F1: ((v18 - 2) < 0)}
                                                _ = v39
                                                _ = v40
                                            }
                                            fmt.Println(v19)
                                            _ = v38
                                        }
                                    }
                                    v19 = v19
                                    v10 += v12
                                    v10 -= 127
                                    _ = fi29
                                    _ = v30
                                }
                            default:
                                fmt.Println(v18)
                            }
                            v41 := GenS7{F0: (uint16(v18) <= 65535)}
                            _ = v26
                            _ = v41
                        }
                        _ = len(v3)
                        if true {
                            var v42 int8 = 42
                            _ = v42
                        }
                    }
                default:
                    if (!true) {
                        var v43 float32 = (1.5 + ((1.5 / 1.0) + (1e10 - 1.5)))
                        for fi44 := 0; fi44 < 7; fi44++ {
                            fmt.Println(v12)
                            v12 = (v12 & (v10 | (v12 / (((^(v18))) | 1))))
                            _ = fi44
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v43
                    }
                }
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = v15
            _ = v16
        } else {
            v3 = append(v3, 127)
            v45 := GenS8{F0: true}
            _ = v45
        }
    default:
        var v46 int16 = 951
        _ = v46
    }
    v47 := GenS2{F0: 0}
    if true {
        v48 := GenS7{F0: false}
        v49 := GenS3{F0: 1}
        fmt.Println(v10, v10)
        _ = v48
        _ = v49
    } else {
        v12 = (v10 << ((919) & 63))
        if (true || true) {
            for fi50 := 0; fi50 < 11; fi50++ {
                if false {
                    var v51 int8 = 100
                    v52 := GenS2{F0: 292}
                    _ = v51
                    _ = v52
                }
                var v53 uint64 = 688
                switch 747 {
                case -9223372036854775808:
                    var v54 uint16 = 450
                    _ = v54
                case 32767:
                    var v55 int8 = 127
                    _ = v55
                }
                _ = fi50
                _ = v53
            }
            v56 := GenS9{F0: (v12 >> ((350) & 63)), F1: 2, F2: 4294967295}
            for fi57 := 0; fi57 < 4; fi57++ {
                v58 := GenS8{F0: false}
                v59 := GenS10{F0: 42, F1: 65535}
                for fi60 := 0; fi60 < 5; fi60++ {
                    if (!(-901 != -991)) {
                        fmt.Println(v10, v12)
                    }
                    switch (fi60 % (((^(fi60))) | 1)) {
                    case 128:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case -359:
                        var v61 uint8 = (v10 % (((v12 | (v10 / ((v10) | 1)))) | 1))
                        _ = v61
                    }
                    fmt.Println(fi57)
                    _ = fi60
                }
                _ = fi57
                _ = v58
                _ = v59
            }
            _ = v56
        }
    }
    var v62 uint32 = 42
    v3 = append(v3, 4294967295)
    _ = len(v3)
    if true {
        if true {
            v10 -= (v10 & (v10 ^ (v12 + v10)))
            var v63 int = 32768
            _ = v63
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v47
    _ = v62
}
