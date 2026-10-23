package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int8 }
type GenS2 struct { F0 bool; F1 uint; F2 uint32 }
type GenS3 struct { F0 int8; F1 uint16; F2 int16 }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 uint64; F1 uint32 }
type GenS7 struct { F0 int; F1 int }
type GenS8 struct { F0 uint16 }
func genFn0(p0 int32, p1 uint16, p2 uint64) int64 {
    if true {
        v5 := GenS1{F0: (p2 - (p2 - 18446744073709551615)), F1: 1, F2: -69}
        if false {
            var v6 uint = 240
            for fi7 := 0; fi7 < 3; fi7++ {
                v8 := GenS2{F0: true, F1: v6, F2: 42}
                if true {
                    var v9 uint8 = 0
                    v10 := GenS2{F0: (-219 <= -9223372036854775808), F1: v6, F2: 0}
                    _ = v10
                    _ = v9
                } else {
                    if false {
                        var v11 int = fi7
                        v12 := GenS2{F0: ((p1 / ((p1) | 1)) <= 2), F1: (^(v6)), F2: 65535}
                        v13 := []int{(fi7 % (((^(fi7))) | 1)), (fi7 + (v11 % (((-(fi7))) | 1))), (fi7 + 2147483647), (fi7 - 2147483647)}
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                }
                for fi14 := 0; fi14 < 5; fi14++ {
                    var v15 int = (fi7 << (((v6 >> (((v6 << ((0) & 63))) & 63))) & 63))
                    if true {
                        if false {
                            var v16 uint16 = p1
                            switch ((("a" + ("" + "\n\t")) + "a") + "foo bar") {
                            case "":
                                fmt.Println(p2)
                            case "hello":
                                if true {
                                    if (-128 == -1) {
                                        var v17 int32 = p0
                                        fmt.Println(p1)
                                        fmt.Println(v6, p2, p1)
                                        _ = v17
                                    } else {
                                        fmt.Println(p0, v15, v6)
                                        var v18 uint16 = v16
                                        _ = v18
                                    }
                                    var v19 uint64 = p2
                                    switch "hello" {
                                    case "\n\t":
                                        var v20 uint16 = 147
                                        _ = v20
                                    case "hello":
                                        v21 := GenS3{F0: 127, F1: v16, F2: int16(v16)}
                                        _ = v21
                                    case "foo bar":
                                        v22 := []int{(-(fi14))}
                                        _ = v22
                                    default:
                                        if false {
                                            v23 := []int{}
                                            var v24 uint32 = 2147483647
                                            if false {
                                                var v25 bool = ((p2 * (^(p2))) <= p2)
                                                switch ("\n\t" + "\n\t") {
                                                case "a":
                                                    v26 := GenS3{F0: 42, F1: (p1 * (v16 >> ((v6) & 63))), F2: 42}
                                                    _ = v26
                                                case "hello":
                                                    v27 := GenS4{F0: 722, F1: (!((!v25) || v25))}
                                                    _ = v27
                                                case "\n\t":
                                                    v25 = v25
                                                default:
                                                    v28 := GenS2{F0: ((!false) && v25), F1: 4294967295, F2: 667}
                                                    _ = v28
                                                }
                                                v29 := GenS5{F0: v19, F1: "a"}
                                                v24 += 0
                                                _ = v25
                                                _ = v29
                                            } else {
                                                var v30 int64 = -725
                                                _ = v30
                                            }
                                            switch fi14 {
                                            case -1:
                                                switch "a" {
                                                case "":
                                                    if true {
                                                        v6 = uint(p1)
                                                        if len(v23) >= 2 {
                                                            v23 = v23[1:len(v23)]
                                                        }
                                                    }
                                                case "a":
                                                    var v31 int32 = (p0 - -1)
                                                    _ = v31
                                                default:
                                                    v16 -= 635
                                                }
                                            case 0:
                                                if (uint32(v19) >= v24) {
                                                    v24 = 2
                                                    switch "a" {
                                                    case "foo bar":
                                                        v32 := GenS6{F0: v19, F1: (v24 * 4294967295)}
                                                        _ = v32
                                                    case "a":
                                                        if (v24 >= v24) {
                                                            v16 -= (p1 + (^(p1)))
                                                            var v33 string = (("" + "\n\t") + "a")
                                                            _ = len(v23)
                                                            _ = v33
                                                        }
                                                    }
                                                    var v34 uint16 = p1
                                                    _ = v34
                                                } else {
                                                    if (!false) {
                                                        var v35 string = "hello"
                                                        fmt.Println(v15, v24)
                                                        if ((true || false) && (!true)) {
                                                            _ = len(v23)
                                                        }
                                                        _ = v35
                                                    } else {
                                                        v16 += (v16 << ((v6) & 63))
                                                        fmt.Println(v24)
                                                        if len(v23) >= 2 {
                                                            v23 = v23[1:len(v23)]
                                                        }
                                                        v24 -= v24
                                                    }
                                                }
                                            default:
                                                var v36 int16 = 831
                                                _ = v36
                                            }
                                            _ = v23
                                            _ = v24
                                        }
                                    }
                                    if (0.0 > (1.5 - 0.1)) {
                                        switch (("foo bar" + ("" + "hello")) + "a") {
                                        case "hello":
                                            p0 = (-(p0))
                                        case "":
                                            var v37 int64 = 88
                                            _ = v37
                                        default:
                                            v38 := []int{916, (fi14 * (fi7 << ((1) & 63))), (fi14 | v15)}
                                            v16 += (^(p1))
                                            v38 = append(v38, 588)
                                            _ = v38
                                        }
                                    } else {
                                        fmt.Println(v19, v6, v15)
                                        switch ("hello" + ("\n\t" + "a")) {
                                        case "foo bar":
                                            if (1 < 128) {
                                                switch (v15 * fi14) {
                                                case -79:
                                                    v39 := []int{fi14, 87, fi14, int(p0)}
                                                    _ = v39
                                                case -9223372036854775808:
                                                    v40 := []int{}
                                                    _ = v40
                                                default:
                                                    v15 += (v15 & v15)
                                                }
                                                v41 := []int{fi7, 9223372036854775807, -1, 2147483647}
                                                v42 := GenS4{F0: v6, F1: (1 > 0)}
                                                _ = v41
                                                _ = v42
                                            }
                                        }
                                        var v43 uint32 = 1
                                        _ = v43
                                    }
                                    fmt.Println(p0, v6, v16)
                                    _ = v19
                                }
                            default:
                                var v44 int = (fi7 - v15)
                                _ = v44
                            }
                            _ = v16
                        } else {
                            var v45 int = v15
                            _ = v45
                        }
                        v46 := []int{}
                        switch "hello" {
                        case "":
                            v47 := GenS5{F0: uint64(p1), F1: (((("a" + "foo bar") + "a") + "a") + ("a" + "foo bar"))}
                            _ = v47
                        case "foo bar":
                            var v48 bool = true
                            _ = v48
                        default:
                            if (("hello" + "\n\t") != "hello") {
                                var v49 int32 = p0
                                _ = v49
                            }
                        }
                        _ = v46
                    }
                    var v50 uint64 = p2
                    _ = fi14
                    _ = v15
                    _ = v50
                }
                _ = fi7
                _ = v8
            }
            _ = v6
        }
        _ = v5
    } else {
        var v51 uint8 = 0
        _ = v51
    }
    v52 := []int{32767, 392}
    var v53 uint16 = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v52
    _ = v53
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = -915
    var v3 int16 = v2
    var v4 uint32 = 0
    v54 := GenS3{F0: -33, F1: 42, F2: v2}
    var v55 string = "\n\t"
    var v56 uint = 348
    v57 := GenS7{F0: 855, F1: 9223372036854775807}
    v58 := v1[v4]
    switch v55 {
    case "a":
        v56 = 2
    }
    for fi59 := 0; fi59 < 1; fi59++ {
        for fi60 := 0; fi60 < 9; fi60++ {
            v61 := v1[(v58 % (((v4 & v4)) | 1))]
            v2 = -32768
            var v62 string = v55
            _ = fi60
            _ = v61
            _ = v62
        }
        _ = fi59
    }
    v55 = (v55 + ("\n\t" + "\n\t"))
    switch -2147483648 {
    case 434:
        var v63 int = 128
        _ = v63
    case 413:
        switch -1 {
        case 1:
            v2 += v2
        case 128:
            if (v55 < v55) {
                var v64 int8 = 33
                for fi65 := 0; fi65 < 6; fi65++ {
                    switch "foo bar" {
                    case "\n\t":
                        for fi66 := 0; fi66 < 19; fi66++ {
                            delete(v1, (^(v58)))
                            v55 = ("hello" + "")
                            _ = fi66
                        }
                    case "a":
                        if (1.5 == (0.0 + ((0.1 - float32(1e10)) * 1.5))) {
                            for fi67 := 0; fi67 < 15; fi67++ {
                                v58 += 0
                                _ = fi67
                            }
                        }
                    }
                    _ = fi65
                }
                _ = v64
            } else {
                for fi68 := 0; fi68 < 14; fi68++ {
                    v69 := GenS2{F0: false, F1: v56, F2: (v4 | v58)}
                    v70 := v1[v58]
                    _ = fi68
                    _ = v69
                    _ = v70
                }
            }
        default:
            var v71 float32 = ((((float32(0.1) / 1e10) * 0.1) - -2.25) * -2.25)
            _ = v71
        }
    case 0:
        v1[445] = (^(v4))
    }
    switch int(v2) {
    case 656:
        v56 = v56
    case -344:
        var v72 int16 = 0
        _ = v72
    default:
        v73 := GenS8{F0: uint16(v4)}
        _ = v73
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v58
}
