package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 int8; F2 uint32 }
type GenS7 struct { F0 uint16; F1 uint16 }
type GenS8 struct { F0 uint }
func genFn0(p0 uint32, p1 int32) uint {
    if false {
        p1 -= p1
    }
    var v6 string = (("a" + "") + "a")
    var v7 uint64 = 546
    var v8 int16 = -32768
    var v9 int64 = 0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 4294967295
}
func genFn1(p0 uint64, p1 int32, p2 uint8) uint64 {
    switch 325 {
    case 1:
        if (true || true) {
            p0 = 18446744073709551615
        }
    case 9223372036854775807:
        v10 := GenS1{F0: 681}
        _ = v10
    case 32767:
        v11 := GenS1{F0: -188}
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        if (-127 >= 9223372036854775807) {
            fmt.Println(p0, fi12)
            if ((!true) || false) {
                var v13 int = (fi12 * fi12)
                var v14 int8 = 100
                _ = v13
                _ = v14
            } else {
                p1 += p1
            }
        }
        var v15 int64 = 33
        _ = fi12
        _ = v15
    }
    v16 := GenS2{F0: -9223372036854775808, F1: -1}
    var v17 float32 = 1.5
    v18 := GenS3{F0: 42, F1: -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    _ = v18
    return (p0 << ((18446744073709551615) & 63))
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{("" + "\n\t"), ("foo bar" + ("a" + "foo bar")), "foo bar"}
    v3 := make(map[uint32]bool)
    var v4 int64 = 1
    var v5 int16 = 32767
    for fi19 := 0; fi19 < 18; fi19++ {
        v2 = append(v2, "")
        var v20 int64 = v4
        if (!(fi19 != 2147483647)) {
            v5 = 1
            var v21 float32 = 0.0
            var v22 float64 = 3.14159
            _ = v21
            _ = v22
        } else {
            v4 = (v20 - (v4 & v4))
            v20 -= (^(v20))
            if (87 <= 65535) {
                if false {
                    v23 := GenS3{F0: 733, F1: -1}
                    var v24 int16 = -804
                    _ = v23
                    _ = v24
                } else {
                    switch (fi19 ^ fi19) {
                    case 885:
                        for fi25 := 0; fi25 < 15; fi25++ {
                            v20 -= (^(v4))
                            if ((0.1 == -2.25) || (!true)) {
                                var v26 int16 = (^(v5))
                                _ = v26
                            }
                            v27 := GenS4{F0: (v5 % ((v5) | 1)), F1: 1}
                            v4 = -9223372036854775808
                            var v28 int8 = 100
                            var v29 int16 = 1
                            _ = fi25
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                    default:
                        var v30 int32 = -41
                        _ = v30
                    }
                    var v31 uint = 65535
                    v32 := GenS4{F0: v5, F1: 2147483647}
                    _ = v31
                    _ = v32
                }
                var v33 int64 = v20
                _ = v33
            }
        }
        _ = fi19
        _ = v20
    }
    var v34 int16 = v5
    fmt.Println(v5, v34, v34)
    v34 = 0
    if (!true) {
        if true {
            var v35 int8 = -128
            v4 -= (v4 - (v4 & (-(v4))))
            for fi36 := 0; fi36 < 15; fi36++ {
                switch 935 {
                case 127:
                    if ((false || (true && false)) && (826 == 2)) {
                        var v37 string = "foo bar"
                        fmt.Println(v4)
                        for fi38 := 0; fi38 < 3; fi38++ {
                            var v39 uint16 = 65535
                            if (((343 < -1) && true) && false) {
                                v3[587] = true
                                if ((((!true) || (-2147483648 >= 2147483647)) && (genFn1(4294967295, 312, uint8(v39)) < 4294967295)) && (v39 <= v39)) {
                                    v1[2147483647] = fi38
                                    fmt.Println(v37, v4, v39)
                                } else {
                                    fmt.Println(v4, v39)
                                    _ = len(v3)
                                    v40 := v1[2147483647]
                                    _ = len(v2)
                                    switch fi36 {
                                    case 1:
                                        if len(v2) > 0 {
                                            v2[0] = ""
                                        }
                                    }
                                    _ = v40
                                }
                            }
                            fmt.Println(fi36, fi38, v34)
                            _ = fi38
                            _ = v39
                        }
                        _ = v37
                    }
                case 128:
                    fmt.Println(fi36)
                default:
                    var v41 uint32 = 65535
                    _ = v41
                }
                for fi42 := 0; fi42 < 15; fi42++ {
                    fmt.Println(v4, v35, v4)
                    v43 := GenS5{F0: v34}
                    v4 -= -1
                    _ = fi42
                    _ = v43
                }
                _ = fi36
            }
            _ = v35
        }
        for fi44 := 0; fi44 < 3; fi44++ {
            var v45 uint16 = 921
            var v46 float32 = 1e10
            for fi47 := 0; fi47 < 20; fi47++ {
                if false {
                    if true {
                        fmt.Println(v34, v4, v46)
                    }
                    var v48 int16 = 1
                    _ = v48
                }
                var v49 uint8 = 255
                if true {
                    v50 := GenS4{F0: v5, F1: -2147483648}
                    v51 := GenS3{F0: 123, F1: v4}
                    _ = v50
                    _ = v51
                }
                _ = fi47
                _ = v49
            }
            _ = fi44
            _ = v45
            _ = v46
        }
        v5 -= 720
    } else {
        if (!true) {
            if (!false) {
                if (-267 > 1) {
                    for fi52 := 0; fi52 < 2; fi52++ {
                        v5 -= (^(v5))
                        fmt.Println(v4, v4)
                        _ = fi52
                    }
                    switch "a" {
                    case "hello":
                        var v53 float32 = ((float32((1.5 - 1e10)) / 1.5) / 0.1)
                        _ = v53
                    case "":
                        v54 := GenS6{F0: 42, F1: 100, F2: 0}
                        _ = v54
                    default:
                        var v55 int = 128
                        _ = v55
                    }
                } else {
                    v5 = v5
                    v56 := GenS4{F0: (v34 & (v5 * -1)), F1: 0}
                    _ = v56
                }
                var v57 bool = true
                if (0 >= -1) {
                    var v58 uint = 255
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v59 := GenS4{F0: v34, F1: int32(v58)}
                    _ = v58
                    _ = v59
                }
                _ = v57
            } else {
                for fi60 := 0; fi60 < 9; fi60++ {
                    if true {
                        v4 = v4
                        v61 := GenS7{F0: 1, F1: 255}
                        if (1.5 >= (-2.25 - 1.5)) {
                            var v62 bool = (!true)
                            v63 := GenS5{F0: v34}
                            _ = v62
                            _ = v63
                        } else {
                            switch (fi60 << ((2) & 63)) {
                            case -360:
                                var v64 uint8 = 127
                                _ = v64
                            default:
                                if true {
                                    _ = len(v2)
                                } else {
                                    var v65 float64 = (1e10 / 0.1)
                                    var v66 int16 = v34
                                    fmt.Println(v34, v5, v65)
                                    _ = v65
                                    _ = v66
                                }
                            }
                        }
                        _ = v61
                    } else {
                        v4 += v4
                        switch "foo bar" {
                        case "a":
                            v5 -= (v5 / (((v5 - (v5 / ((v34) | 1)))) | 1))
                        case "foo bar":
                            if false {
                                var v67 uint32 = 2147483647
                                var v68 uint64 = 2
                                _ = v67
                                _ = v68
                            }
                        }
                        fmt.Println(v4, fi60, v34)
                    }
                    _ = fi60
                }
            }
            for fi69 := 0; fi69 < 18; fi69++ {
                var v70 uint = 18446744073709551615
                fmt.Println(fi69, v70)
                _ = fi69
                _ = v70
            }
        }
    }
    if false {
        var v71 uint8 = 42
        v71 += (v71 & 1)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v71
    }
    var v72 int64 = v4
    v5 = v5
    if true {
        v73 := GenS8{F0: 255}
        switch -499 {
        case 235:
            if (!true) {
                v5 = (-(v5))
            } else {
                v74 := GenS6{F0: v5, F1: 27, F2: 648}
                _ = v74
            }
        case 1:
            fmt.Println(v5)
        default:
            v75 := GenS8{F0: 42}
            _ = v75
        }
        v72 += v72
        _ = v73
    }
    delete(v3, 42)
    if ((1e10 >= 1e10) && true) {
        for fi76 := 0; fi76 < 9; fi76++ {
            v72 -= (v72 - (-(v72)))
            _ = fi76
        }
        for fi77 := 0; fi77 < 3; fi77++ {
            var v78 uint = 478
            _ = fi77
            _ = v78
        }
        v72 -= -1
    }
    var v79 string = "\n\t"
    var v80 float64 = (1.5 - (3.14159 + 3.14159))
    v80 = (-2.25 - 1e10)
    fmt.Println(v5)
    switch "a" {
    case "\n\t":
        v2 = append(v2, v79)
    default:
        v4 = -1
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v72:", v72)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v72
    _ = v79
    _ = v80
}
