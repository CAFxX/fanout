package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 uint64; F1 uint8 }
type GenS8 struct { F0 uint32; F1 uint; F2 uint8 }
func genFn0(p0 uint32, p1 uint16) uint8 {
    p1 = (p1 % (((p1 / (((p1 ^ (p1 - 2))) | 1))) | 1))
    v6 := GenS1{F0: p1}
    var v7 int16 = 175
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 127
}
func genFn1(p0 string) uint8 {
    var v8 uint16 = 65535
    v8 = 596
    if true {
        v8 = v8
        v9 := GenS2{F0: -70, F1: 1, F2: 1}
        _ = v9
    } else {
        fmt.Println(p0, v8)
        v10 := []int{4294967295, -9223372036854775808, 127, 0}
        _ = v10
    }
    switch 42 {
    case -668:
        v11 := GenS2{F0: 42, F1: v8, F2: 18446744073709551615}
        _ = v11
    case 2147483647:
        fmt.Println(v8, v8)
    case 42:
        if (((-25 >= -128) || false) && false) {
            if true {
                fmt.Println(v8)
            }
            v12 := GenS1{F0: (v8 & 32767)}
            if true {
                for fi13 := 0; fi13 < 9; fi13++ {
                    var v14 bool = false
                    fmt.Println(v8, v14, fi13)
                    _ = fi13
                    _ = v14
                }
            } else {
                fmt.Println(p0, p0)
                for fi15 := 0; fi15 < 6; fi15++ {
                    if false {
                        v16 := []int{}
                        _ = v16
                    } else {
                        p0 = ("a" + "foo bar")
                    }
                    for fi17 := 0; fi17 < 5; fi17++ {
                        fmt.Println(v8, v8, fi17)
                        fmt.Println(v8, v8)
                        _ = fi17
                    }
                    _ = fi15
                }
                p0 = (p0 + ("\n\t" + (p0 + ((((p0 + (((("a" + ("" + ("" + p0))) + p0) + p0) + p0)) + ((p0 + p0) + p0)) + (("" + p0) + ("" + p0))) + "a"))))
            }
            _ = v12
        } else {
            var v18 bool = true
            v19 := []int{128, -668, 127, -2147483648}
            _ = v18
            _ = v19
        }
    default:
        for fi20 := 0; fi20 < 6; fi20++ {
            fmt.Println(v8)
            switch 0 {
            case 42:
                for fi21 := 0; fi21 < 12; fi21++ {
                    var v22 int32 = 2147483647
                    _ = fi21
                    _ = v22
                }
            default:
                v8 -= 177
            }
            v23 := []int{fi20, fi20}
            _ = fi20
            _ = v23
        }
    }
    v24 := []int{128}
    _ = p0
    _ = v24
    _ = v8
    return genFn0(65535, (^(v8)))
}

func main() {
    var v1 uint64 = 4294967295
    var v2 string = ("\n\t" + (("\n\t" + (("foo bar" + "hello") + ("" + "foo bar"))) + ""))
    var v3 uint8 = 42
    v4 := make(map[string]uint32)
    v5 := make(map[int]bool)
    for fi25 := 0; fi25 < 8; fi25++ {
        fmt.Println(fi25, v2, v1)
        v26 := GenS3{F0: 0, F1: 2}
        _ = fi25
        _ = v26
    }
    if false {
        var v27 uint64 = v1
        _ = v27
    }
    _ = len(v5)
    var v28 int16 = 1
    switch 42 {
    case 128:
        for fi29 := 0; fi29 < 16; fi29++ {
            switch fi29 {
            case 4294967295:
                v1 -= (v1 ^ v1)
            default:
                var v30 float32 = (0.1 / float32((1.5 * 3.14159)))
                _ = v30
            }
            var v31 uint = 18446744073709551615
            _ = fi29
            _ = v31
        }
    case 9223372036854775807:
        switch v2 {
        case "":
            if (0 == 18446744073709551615) {
                if (2147483648 == 208) {
                    var v32 string = (v2 + (v2 + v2))
                    _ = v32
                } else {
                    var v33 uint64 = (v1 & v1)
                    v33 = v33
                    var v34 int16 = v28
                    _ = v33
                    _ = v34
                }
                for fi35 := 0; fi35 < 9; fi35++ {
                    var v36 float64 = 1.5
                    var v37 int32 = -1
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                var v38 int = 128
                _ = v38
            }
        case "foo bar":
            for fi39 := 0; fi39 < 18; fi39++ {
                fmt.Println(fi39, fi39, v2)
                _ = fi39
            }
        case "\n\t":
            v28 += (v28 * (v28 ^ v28))
        default:
            v28 -= (-(v28))
        }
    }
    switch (("hello" + ("hello" + v2)) + v2) {
    case "a":
        v40 := []int{0, -2147483648}
        _ = v40
    case "":
        if true {
            for fi41 := 0; fi41 < 3; fi41++ {
                fmt.Println(v1, v28)
                v42 := GenS2{F0: int64(v1), F1: 0, F2: 0}
                var v43 float64 = 1.5
                _ = fi41
                _ = v42
                _ = v43
            }
            if (-616 <= -1) {
                var v44 float32 = (3.14159 / (float32((0.0 + 3.14159)) * float32(((1e10 + float64(((0.1 / 1.0) * -2.25))) - 1e10))))
                _ = v44
            } else {
                v45 := []int{42, -1, 32768, 46}
                _ = v45
            }
            v1 -= v1
        } else {
            v46 := GenS2{F0: -1, F1: 506, F2: (v1 - (^(v1)))}
            var v47 int32 = -728
            _ = v46
            _ = v47
        }
    }
    v48 := GenS4{F0: v2, F1: (genFn1((v2 + v2)) >= genFn0(631, 678))}
    v3 = genFn1(v2)
    var v49 bool = false
    if (v49 && (((!true) || (!((!v49) || (false && (!v49))))) || ((^(v28)) >= v28))) {
        for fi50 := 0; fi50 < 19; fi50++ {
            fmt.Println(v28, v2)
            for fi51 := 0; fi51 < 16; fi51++ {
                if (2 == 32767) {
                    fmt.Println(v2, fi51, v49)
                }
                _ = fi51
            }
            _ = fi50
        }
        var v52 uint32 = 0
        var v53 float64 = 3.14159
        _ = v52
        _ = v53
    }
    var v54 bool = (((v49 || (32768 >= 42)) || ((true && true) && v49)) && ((v28 < 42) && v49))
    v28 = (v28 >> ((255) & 63))
    v55 := GenS5{F0: 563}
    v2 = v2
    switch "foo bar" {
    case "a":
        if false {
            for fi56 := 0; fi56 < 12; fi56++ {
                switch (fi56 >> ((255) & 63)) {
                case -9223372036854775808:
                    v57 := GenS6{F0: ((v2 + (v2 + (v2 + "a"))) + "\n\t"), F1: 2147483648}
                    _ = v57
                case -336:
                    v58 := GenS1{F0: 32767}
                    _ = v58
                case 32768:
                    if false {
                        if (2 != 707) {
                            v2 = ("\n\t" + "")
                        } else {
                            switch "a" {
                            case "hello":
                                if (0.0 > 0.1) {
                                    for fi59 := 0; fi59 < 8; fi59++ {
                                        var v60 uint8 = (v3 - (^(v3)))
                                        v3 = v60
                                        v3 += v3
                                        _ = fi59
                                        _ = v60
                                    }
                                    v61 := []int{}
                                    _ = v61
                                }
                            default:
                                var v62 int = (fi56 * fi56)
                                _ = v62
                            }
                            var v63 int = (^(fi56))
                            v54 = (genFn1(v2) == 235)
                            _ = v63
                        }
                        switch ((v2 + (v2 + "a")) + v2) {
                        case "":
                            switch -358 {
                            case 32768:
                                _ = len(v4)
                            }
                        case "hello":
                            _ = len(v4)
                        case "foo bar":
                            for fi64 := 0; fi64 < 17; fi64++ {
                                if (65535 >= 754) {
                                    var v65 uint = 0
                                    _ = v65
                                }
                                var v66 int16 = 1
                                _ = fi64
                                _ = v66
                            }
                        }
                        v67 := []int{}
                        _ = v67
                    }
                }
                fmt.Println(v2, v28)
                _ = fi56
            }
            switch -9223372036854775808 {
            case 916:
                if v49 {
                    fmt.Println(v1, v1, v3)
                    var v68 float64 = 1e10
                    _ = v68
                }
            case -850:
                for fi69 := 0; fi69 < 9; fi69++ {
                    fmt.Println(v49)
                    var v70 bool = ("" >= (("a" + "hello") + ""))
                    _ = fi69
                    _ = v70
                }
            }
            v1 = v1
        }
    case "foo bar":
        v71 := GenS7{F0: v1, F1: (v3 | 42)}
        _ = v71
    case "hello":
        var v72 uint32 = 2
        _ = v72
    default:
        v73 := []int{-75, -544}
        _ = v73
    }
    if v54 {
        v74 := GenS8{F0: 4294967295, F1: 861, F2: v3}
        var v75 bool = (2147483647 == -95)
        if (!true) {
            v2 = "\n\t"
            var v76 uint = 822
            _ = v76
        }
        _ = v74
        _ = v75
    } else {
        var v77 float64 = float64(1e10)
        var v78 uint64 = 1
        _ = v77
        _ = v78
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v54
    _ = v55
}
