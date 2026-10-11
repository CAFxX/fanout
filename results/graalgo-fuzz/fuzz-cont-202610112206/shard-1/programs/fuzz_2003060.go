package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int; F2 string }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 string; F1 uint8; F2 int8 }
type GenS5 struct { F0 int16; F1 int8; F2 int16 }
type GenS6 struct { F0 uint8; F1 int16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 bool; F1 uint16 }
func genFn0(p0 int8, p1 int16, p2 uint32) int {
    for fi5 := 0; fi5 < 11; fi5++ {
        if (!true) {
            v6 := GenS1{F0: p0}
            fmt.Println(fi5, fi5)
            _ = v6
        } else {
            switch fi5 {
            case 32767:
                for fi7 := 0; fi7 < 20; fi7++ {
                    p1 -= (p1 >> ((18446744073709551615) & 63))
                    _ = fi7
                }
            case 32768:
                if (false && false) {
                    var v8 string = "\n\t"
                    if (300 == 255) {
                        var v9 int16 = (p1 * (p1 + p1))
                        _ = v9
                    }
                    p1 += int16(p2)
                    _ = v8
                }
            case -2147483648:
                var v10 int64 = 38
                _ = v10
            }
            for fi11 := 0; fi11 < 11; fi11++ {
                v12 := []int{fi11, fi11, fi5, -126}
                p1 = p1
                _ = fi11
                _ = v12
            }
        }
        p2 -= p2
        _ = fi5
    }
    v13 := GenS1{F0: (p0 - (^(p0)))}
    p0 -= (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 1
}
func genFn1(p0 uint, p1 int16) uint32 {
    fmt.Println(p0, p1)
    if (true && (188 >= 128)) {
        p1 += 1000
        for fi14 := 0; fi14 < 7; fi14++ {
            p1 = -32768
            var v15 uint32 = 830
            v16 := []int{genFn0(-1, (p1 / (((p1 % (((p1 >> ((p0) & 63))) | 1))) | 1)), (v15 * v15)), 128, (fi14 << (((p0 | 4294967295)) & 63))}
            _ = fi14
            _ = v15
            _ = v16
        }
    } else {
        var v17 int16 = 0
        v18 := []int{-9223372036854775808, 879, -1, -9223372036854775808}
        var v19 int = genFn0(100, int16(p0), 0)
        _ = v17
        _ = v18
        _ = v19
    }
    fmt.Println(p1, p1)
    var v20 int8 = 20
    _ = p0
    _ = p1
    _ = v20
    return 43
}

func main() {
    v1 := []int{42, 127, 32768}
    var v2 uint32 = 487
    var v3 int = 9223372036854775807
    var v4 uint16 = 671
    for fi21 := 0; fi21 < 12; fi21++ {
        for fi22 := 0; fi22 < 2; fi22++ {
            var v23 uint32 = 861
            if true {
                switch v3 {
                case -709:
                    fmt.Println(v3, v23, v2)
                case 4294967295:
                    v24 := GenS1{F0: 42}
                    _ = v24
                case 0:
                    fmt.Println(fi22, v2)
                default:
                    v2 += (v2 >> ((517) & 63))
                }
                v4 -= 255
                v4 = v4
            } else {
                v25 := GenS1{F0: 82}
                var v26 int = (v3 * 131)
                _ = v25
                _ = v26
            }
            v27 := GenS2{F0: 42, F1: (^(v3)), F2: "\n\t"}
            _ = fi22
            _ = v23
            _ = v27
        }
        if false {
            var v28 uint = 4294967295
            switch ("hello" + "foo bar") {
            case "\n\t":
                if (1 == 609) {
                    for fi29 := 0; fi29 < 2; fi29++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        switch ("\n\t" + "") {
                        case "foo bar":
                            if true {
                                var v30 bool = false
                                v31 := GenS3{F0: 612, F1: (2147483647 != -498)}
                                _ = v30
                                _ = v31
                            }
                        case "\n\t":
                            var v32 uint64 = 2
                            _ = v32
                        }
                        v2 = genFn1(v28, -767)
                        _ = fi29
                    }
                    v33 := GenS1{F0: 42}
                    for fi34 := 0; fi34 < 11; fi34++ {
                        v4 = (^(v4))
                        if false {
                            fmt.Println(fi34, v3, v4)
                        } else {
                            v2 = 2147483647
                            var v35 float64 = (0.0 / ((float64((3.14159 - 1.5)) * ((1e10 / -2.25) - ((0.0 + 0.1) - 1.5))) - 0.1))
                            _ = len(v1)
                            v4 = v4
                            fmt.Println(v35, v4)
                            _ = v35
                        }
                        _ = fi34
                    }
                    _ = v33
                } else {
                    v28 -= (v28 ^ (v28 - (v28 >> (((^(v28))) & 63))))
                    v28 += 42
                    v36 := GenS4{F0: "hello", F1: 2, F2: 110}
                    _ = v36
                }
            case "hello":
                switch "" {
                case "a":
                    v2 = 1
                case "foo bar":
                    v4 = (v4 - 0)
                case "":
                    fmt.Println(v28, v28)
                }
            default:
                if false {
                    for fi37 := 0; fi37 < 6; fi37++ {
                        switch (v3 + fi21) {
                        case 128:
                            var v38 uint8 = 0
                            _ = v38
                        case 4294967295:
                            var v39 string = "a"
                            _ = v39
                        }
                        _ = fi37
                    }
                    var v40 uint32 = v2
                    if (!(true || false)) {
                        if true {
                            if (!false) {
                                _ = len(v1)
                            }
                            for fi41 := 0; fi41 < 6; fi41++ {
                                v28 -= 719
                                switch (fi21 & (fi41 | fi41)) {
                                case -2147483648:
                                    v3 -= (fi21 / ((genFn0(127, int16(v2), v2)) | 1))
                                case 35:
                                    v42 := GenS2{F0: -76, F1: v3, F2: "hello"}
                                    _ = v42
                                }
                                v40 -= v2
                                _ = fi41
                            }
                            for fi43 := 0; fi43 < 12; fi43++ {
                                var v44 string = "foo bar"
                                var v45 bool = false
                                _ = len(v1)
                                _ = fi43
                                _ = v44
                                _ = v45
                            }
                        } else {
                            var v46 string = "hello"
                            if len(v1) > 0 {
                                v1[0] = 435
                            }
                            _ = v46
                        }
                    } else {
                        switch 32768 {
                        case 42:
                            if (1.5 != -2.25) {
                                fmt.Println(v3, v40)
                            }
                        case -1:
                            var v47 int32 = 42
                            _ = v47
                        default:
                            var v48 uint32 = (^(v40))
                            _ = v48
                        }
                    }
                    _ = v40
                }
            }
            v3 = int(v28)
            _ = v28
        }
        v49 := GenS5{F0: -906, F1: -54, F2: -32768}
        _ = fi21
        _ = v49
    }
    var v50 int16 = -1
    v50 = (v50 + v50)
    v2 = (v2 ^ 1)
    fmt.Println(v3)
    var v51 int32 = -2147483648
    var v52 bool = true
    _ = len(v1)
    v50 = (v50 >> ((943) & 63))
    v3 -= v3
    if v52 {
        switch (^(v3)) {
        case 1:
            v53 := GenS6{F0: 127, F1: (v50 | v50)}
            _ = v53
        case 437:
            v51 += v51
        default:
            v54 := GenS7{F0: v2}
            _ = v54
        }
        if (652 > (v4 & 719)) {
            var v55 uint32 = (v2 - genFn1(2, 1000))
            var v56 uint = 18446744073709551615
            v57 := GenS8{F0: (!v52), F1: (v4 * (v4 % ((v4) | 1)))}
            _ = v55
            _ = v56
            _ = v57
        } else {
            if (v50 <= v50) {
                _ = len(v1)
                var v58 uint8 = 1
                _ = v58
            }
        }
        for fi59 := 0; fi59 < 5; fi59++ {
            var v60 float64 = (0.1 + (float64(((float32(((1e10 * 0.1) - ((((3.14159 - (1.5 - 1e10)) / (0.1 * (3.14159 / 1.0))) - ((0.1 / ((((3.14159 / ((((((-2.25 / 1e10) + (0.1 * 0.0)) - ((3.14159 - (0.0 * 0.0)) + (0.0 - (1e10 + 0.0)))) / float64(((-2.25 + 1e10) * 1.5))) + (3.14159 / 1e10)) + 0.1)) - 1.5) - -2.25) + ((-2.25 + 1e10) * 3.14159))) - (-2.25 + 3.14159))) - (0.0 / 1.0)))) / -2.25) + 1.5)) + 0.0))
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch v3 {
            case 127:
                for fi61 := 0; fi61 < 3; fi61++ {
                    if v52 {
                        v1 = append(v1, fi61)
                        v4 = v4
                    } else {
                        v51 += v51
                        v60 = v60
                        v60 += v60
                    }
                    _ = fi61
                }
            case -1:
                if v52 {
                    var v62 float64 = v60
                    v60 = ((v60 / 1.5) / (v62 * v62))
                    _ = v62
                } else {
                    var v63 bool = v52
                    _ = v63
                }
            case 1:
                fmt.Println(v51, v4, v2)
            default:
                v64 := GenS3{F0: 4294967295, F1: (!(!(!v52)))}
                _ = v64
            }
            _ = fi59
            _ = v60
        }
    }
    switch ("" + "a") {
    case "\n\t":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "a":
        v50 -= -1
    case "":
        for fi65 := 0; fi65 < 8; fi65++ {
            v51 = v51
            _ = len(v1)
            switch genFn0(127, (v50 - v50), v2) {
            case 127:
                if v52 {
                    _ = len(v1)
                    switch "\n\t" {
                    case "\n\t":
                        v66 := GenS7{F0: 0}
                        _ = v66
                    case "a":
                        if v52 {
                            var v67 float32 = (float32(0.1) / 1.5)
                            for fi68 := 0; fi68 < 3; fi68++ {
                                var v69 float32 = (v67 * (1.5 / 1.0))
                                fmt.Println(v69, fi65)
                                _ = fi68
                                _ = v69
                            }
                            if true {
                                switch genFn0(int8(v4), (-(v50)), uint32(v4)) {
                                case -2147483648:
                                    v4 = v4
                                case 42:
                                    if v52 {
                                        v3 -= (fi65 & int(v51))
                                    }
                                case 32767:
                                    v70 := GenS7{F0: (v2 + (v2 - 2))}
                                    _ = v70
                                }
                            } else {
                                switch genFn0(127, v50, (^(v2))) {
                                case -927:
                                    fmt.Println(fi65, v67, v50)
                                default:
                                    v1 = append(v1, (^(fi65)))
                                }
                                fmt.Println(v51)
                                switch fi65 {
                                case 2147483647:
                                    fmt.Println(v50)
                                default:
                                    v50 += (v50 % ((-597) | 1))
                                }
                            }
                            _ = v67
                        } else {
                            if v52 {
                                v71 := GenS7{F0: v2}
                                var v72 int64 = 42
                                switch genFn0(100, 32767, v2) {
                                case -63:
                                    var v73 uint16 = 1
                                    _ = v73
                                case 9223372036854775807:
                                    var v74 uint32 = (v2 | (v2 * (v2 & (v2 | (v2 * v2)))))
                                    _ = v74
                                case 677:
                                    var v75 uint = 757
                                    _ = v75
                                }
                                _ = v71
                                _ = v72
                            } else {
                                _ = len(v1)
                            }
                            switch ("\n\t" + "\n\t") {
                            case "\n\t":
                                var v76 string = "foo bar"
                                _ = v76
                            case "":
                                fmt.Println(v4, v50, v52)
                            case "a":
                                fmt.Println(v51, v50, fi65)
                            }
                        }
                    default:
                        v52 = (true || ((v50 % ((v50) | 1)) > (v50 << ((65535) & 63))))
                    }
                    fmt.Println(v2, v3, v3)
                } else {
                    var v77 int64 = 436
                    _ = v77
                }
            default:
                var v78 uint64 = 1
                _ = v78
            }
            _ = fi65
        }
    }
    var v79 string = ("a" + (((("hello" + "\n\t") + "") + ("\n\t" + "\n\t")) + ("hello" + (("\n\t" + "foo bar") + ""))))
    v80 := GenS1{F0: -26}
    v50 -= v50
    v1 = append(v1, (v3 * 32767))
    v81 := GenS7{F0: 65535}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
    _ = v52
    _ = v79
    _ = v80
    _ = v81
}
