package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 int; F2 string }
type GenS4 struct { F0 int; F1 int; F2 int32 }
type GenS5 struct { F0 int; F1 uint32; F2 int32 }
func genFn0(p0 string, p1 int8) uint32 {
    v6 := []int{42, 32767, -905, 32768}
    v7 := GenS1{F0: 32767, F1: -9223372036854775808, F2: -215}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 int64, p1 int, p2 uint) int32 {
    if true {
        fmt.Println(p1)
        p1 += p1
        v8 := []int{0, p1, (p1 + p1), 32767}
        _ = v8
    } else {
        var v9 int64 = 71
        var v10 uint32 = 857
        if true {
            var v11 int32 = -311
            var v12 uint16 = 879
            _ = v11
            _ = v12
        } else {
            v10 = v10
            var v13 bool = true
            var v14 uint32 = 0
            _ = v13
            _ = v14
        }
        _ = v10
        _ = v9
    }
    if ((((0.1 * 1.5) - 1.5) < -2.25) && (2 >= 308)) {
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 uint8 = 0
            p0 = (p0 + 0)
            _ = fi15
            _ = v16
        }
    }
    p1 += (p1 * 9223372036854775807)
    _ = p0
    _ = p1
    _ = p2
    return -56
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[int]int)
    var v3 int16 = -32768
    var v4 int32 = 888
    v5 := make(map[uint32]int)
    v1 = (0.0 / -2.25)
    v3 += -446
    v3 -= v3
    v3 = v3
    v17 := GenS1{F0: 2, F1: 128, F2: -32768}
    switch "hello" {
    case "a":
        var v18 int = 377
        _ = v18
    }
    var v19 int32 = -1
    if (((v3 ^ v3) >= (^(v3))) || false) {
        v1 += ((v1 / float32(float64(v1))) * v1)
        v19 = genFn1(-785, -2147483648, 1)
        var v20 float32 = (v1 + 1e10)
        _ = v20
    }
    var v21 uint16 = 998
    switch 32767 {
    case 32768:
        switch ("hello" + ("" + "hello")) {
        case "\n\t":
            v1 = (v1 + 3.14159)
        case "foo bar":
            var v22 int = 32768
            _ = v22
        case "":
            if true {
                if false {
                    if false {
                        v1 += float32(0.0)
                        v19 += (v19 & v19)
                        v23 := GenS2{F0: 42}
                        _ = v23
                    }
                } else {
                    if false {
                        if true {
                            for fi24 := 0; fi24 < 6; fi24++ {
                                for fi25 := 0; fi25 < 12; fi25++ {
                                    v5[65535] = (-(fi24))
                                    _ = fi25
                                }
                                _ = fi24
                            }
                            v26 := []int{-1, 32768}
                            v27 := GenS1{F0: 350, F1: 1, F2: v3}
                            _ = v26
                            _ = v27
                        } else {
                            for fi28 := 0; fi28 < 12; fi28++ {
                                if false {
                                    var v29 uint8 = 2
                                    v1 = float32(0.0)
                                    _ = v29
                                } else {
                                    v30 := []int{fi28, fi28}
                                    v3 = v3
                                    _ = v30
                                }
                                _ = fi28
                            }
                        }
                        v31 := GenS3{F0: false, F1: 32768, F2: ""}
                        _ = v31
                    } else {
                        var v32 uint64 = 42
                        v33 := v2[-980]
                        _ = v32
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 12; fi34++ {
                        v5[2147483648] = 9223372036854775807
                        fmt.Println(v21, v21)
                        if (-1 == 0) {
                            v1 = v1
                        } else {
                            for fi35 := 0; fi35 < 18; fi35++ {
                                var v36 string = ("foo bar" + "\n\t")
                                v37 := GenS4{F0: fi34, F1: fi35, F2: v19}
                                v1 += v1
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                            v38 := GenS4{F0: fi34, F1: 812, F2: genFn1(int64(v21), fi34, 762)}
                            _ = v38
                        }
                        _ = fi34
                    }
                }
                v39 := []int{42, 128}
                _ = v39
            } else {
                var v40 uint32 = genFn0((("a" + "hello") + ((((("\n\t" + "foo bar") + "\n\t") + ("a" + (("\n\t" + "a") + ("" + ("" + "hello"))))) + "") + ("" + "a"))), 36)
                var v41 uint16 = (v21 ^ v21)
                for fi42 := 0; fi42 < 8; fi42++ {
                    v4 -= (v4 / ((genFn1(-1, int(v21), 4294967295)) | 1))
                    v21 -= v21
                    fmt.Println(v40, v19, v40)
                    _ = fi42
                }
                _ = v40
                _ = v41
            }
        default:
            v21 = 2
        }
    case 1:
        if false {
            switch ("foo bar" + (("hello" + "a") + "a")) {
            case "a":
                var v43 uint = 1
                _ = v43
            default:
                var v44 uint8 = 255
                _ = v44
            }
            switch 624 {
            case -486:
                v45 := []int{-8, -147, -1}
                _ = v45
            case 0:
                var v46 int16 = (^(v3))
                _ = v46
            case -1:
                switch "a" {
                case "foo bar":
                    if false {
                        if false {
                            v3 = -467
                        }
                    } else {
                        var v47 float64 = 1.5
                        var v48 uint64 = 837
                        _ = v47
                        _ = v48
                    }
                case "a":
                    switch 330 {
                    case 32768:
                        v21 -= (v21 * 0)
                    case 0:
                        var v49 uint64 = 398
                        _ = v49
                    }
                }
            default:
                v19 = -556
            }
            fmt.Println(v3, v21, v19)
        }
    case 588:
        var v50 uint32 = 1
        _ = v50
    }
    if ((!(34 != 41)) && true) {
        if ((v3 | (v3 + -792)) < 1) {
            var v51 int32 = v4
            if (251 != 19) {
                v1 = -2.25
                switch -255 {
                case -502:
                    var v52 bool = ((v21 | 389) > (v21 << ((546) & 63)))
                    _ = v52
                case 725:
                    switch (("hello" + ("" + ("\n\t" + "foo bar"))) + "foo bar") {
                    case "foo bar":
                        var v53 uint64 = 18446744073709551615
                        _ = v53
                    default:
                        for fi54 := 0; fi54 < 8; fi54++ {
                            for fi55 := 0; fi55 < 10; fi55++ {
                                var v56 int64 = 42
                                v51 += v51
                                _ = fi55
                                _ = v56
                            }
                            _ = fi54
                        }
                    }
                case 2147483647:
                    if ((!false) || (true && true)) {
                        for fi57 := 0; fi57 < 16; fi57++ {
                            for fi58 := 0; fi58 < 10; fi58++ {
                                v5[42] = -1
                                fmt.Println(v1, v3, v19)
                                v5[118] = (fi57 - fi58)
                                _ = fi58
                            }
                            _ = fi57
                        }
                        var v59 uint8 = 127
                        v60 := GenS1{F0: (v21 ^ v21), F1: 0, F2: 32767}
                        _ = v59
                        _ = v60
                    }
                default:
                    var v61 uint8 = 60
                    _ = v61
                }
                if (42 <= 42) {
                    v51 -= 2147483647
                    if true {
                        v62 := GenS4{F0: 2147483647, F1: 849, F2: 279}
                        var v63 string = ("foo bar" + "a")
                        _ = v62
                        _ = v63
                    }
                } else {
                    for fi64 := 0; fi64 < 15; fi64++ {
                        if (1 == 4294967295) {
                            for fi65 := 0; fi65 < 11; fi65++ {
                                fmt.Println(v21, v51, v3)
                                _ = fi65
                            }
                            var v66 int32 = (v19 + (-(v51)))
                            _ = v66
                        }
                        v21 += (v21 ^ v21)
                        _ = fi64
                    }
                    delete(v5, 0)
                }
            }
            switch "a" {
            case "":
                switch 32767 {
                case 42:
                    _ = len(v2)
                case 0:
                    fmt.Println(v21)
                }
            case "foo bar":
                var v67 int32 = v4
                _ = v67
            }
            _ = v51
        }
        v68 := GenS5{F0: 127, F1: 4294967295, F2: v4}
        var v69 float32 = 0.1
        _ = v68
        _ = v69
    }
    var v70 uint32 = 1
    v4 = genFn1(9223372036854775807, 9223372036854775807, 2)
    v3 = (v3 + v3)
    for fi71 := 0; fi71 < 14; fi71++ {
        v19 = -1
        _ = fi71
    }
    v1 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v70:", v70)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v70
}
