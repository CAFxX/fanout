package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int64; F1 int; F2 uint }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 uint16; F1 int16; F2 int64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint64; F1 uint16; F2 int32 }
func genFn0(p0 int64, p1 uint64, p2 uint) int16 {
    if false {
        v5 := []int{}
        if true {
            var v6 int32 = int32(p1)
            _ = v6
        } else {
            for fi7 := 0; fi7 < 8; fi7++ {
                if true {
                    _ = len(v5)
                    var v8 int8 = -128
                    _ = v8
                }
                switch fi7 {
                case 42:
                    fmt.Println(p0, fi7, p0)
                case 128:
                    p2 -= (^(p2))
                case 1:
                    fmt.Println(p0, p1, p0)
                default:
                    if true {
                        switch fi7 {
                        case 726:
                            var v9 uint32 = 42
                            _ = v9
                        case -9223372036854775808:
                            v10 := GenS1{F0: false}
                            _ = v10
                        case -1:
                            if (p1 == 18446744073709551615) {
                                var v11 uint64 = 18446744073709551615
                                _ = v11
                            }
                        default:
                            var v12 int8 = -128
                            _ = v12
                        }
                        if true {
                            if len(v5) > 0 {
                                v5[0] = (fi7 ^ int(p2))
                            }
                        } else {
                            _ = len(v5)
                            var v13 float32 = 1.5
                            _ = v13
                        }
                        var v14 int32 = -2147483648
                        _ = v14
                    }
                }
                _ = fi7
            }
            p2 -= p2
        }
        _ = v5
    } else {
        if false {
            if true {
                v15 := []int{4294967295, 0, -2147483648, 0}
                p0 = (-(p0))
                _ = v15
            } else {
                var v16 int16 = 32767
                var v17 int64 = p0
                v18 := GenS1{F0: ((!(true && (2147483647 > 2147483647))) && ((!false) && false))}
                _ = v16
                _ = v17
                _ = v18
            }
            if (false && false) {
                v19 := []int{-1, -828, -9223372036854775808, 2147483647}
                p1 += (^(p1))
                _ = v19
            } else {
                var v20 bool = false
                _ = v20
            }
        }
        p1 += (p1 ^ (p1 * (p1 | 70)))
    }
    var v21 int8 = 42
    p2 -= (p2 / ((p2) | 1))
    v22 := GenS2{F0: (p1 % ((1) | 1)), F1: 89}
    switch -1 {
    case 9223372036854775807:
        var v23 uint8 = 142
        _ = v23
    case 127:
        if ((^(p1)) != 2) {
            v24 := GenS3{F0: (p0 / (((p0 ^ (p0 >> (((p2 * p2)) & 63)))) | 1)), F1: 9223372036854775807, F2: (p2 % ((p2) | 1))}
            var v25 string = "foo bar"
            switch 942 {
            case 32768:
                fmt.Println(v21, p1)
            case 32767:
                if true {
                    v26 := []int{-2147483648}
                    if (p1 >= (^(p1))) {
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                        v27 := GenS3{F0: 816, F1: 4294967295, F2: (p2 - (p2 >> ((p2) & 63)))}
                        if (288 <= 1) {
                            fmt.Println(v21, v21)
                            if false {
                                p0 = p0
                            } else {
                                p0 = p0
                                var v28 int32 = -1
                                _ = v28
                            }
                            var v29 uint16 = 422
                            _ = v29
                        } else {
                            for fi30 := 0; fi30 < 16; fi30++ {
                                for fi31 := 0; fi31 < 19; fi31++ {
                                    v26 = append(v26, 2147483647)
                                    var v32 uint = 261
                                    _ = fi31
                                    _ = v32
                                }
                                _ = fi30
                            }
                            v25 = "a"
                        }
                        _ = v27
                    }
                    _ = v26
                } else {
                    v21 = (v21 + v21)
                    v21 = int8(p2)
                    var v33 string = "foo bar"
                    _ = v33
                }
            case 1:
                var v34 uint8 = 255
                _ = v34
            }
            _ = v24
            _ = v25
        }
    case -284:
        for fi35 := 0; fi35 < 2; fi35++ {
            for fi36 := 0; fi36 < 20; fi36++ {
                fmt.Println(p2, v21, fi35)
                fmt.Println(p1, p0)
                _ = fi36
            }
            _ = fi35
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    return -196
}
func genFn1(p0 uint32, p1 int32, p2 int) int64 {
    v37 := []int{p2, 274, (p2 << ((690) & 63))}
    fmt.Println(p1)
    var v38 uint32 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v37
    _ = v38
    return 9223372036854775807
}

func main() {
    var v1 int = 1
    var v2 string = ("foo bar" + "hello")
    var v3 int32 = 0
    var v4 uint = 42
    fmt.Println(v2, v4)
    if (false && (-2.25 < float64(1e10))) {
        v1 -= 44
        v4 = 4294967295
        v3 = 0
    }
    v39 := GenS1{F0: true}
    if (false && false) {
        var v40 uint = (v4 >> (((^(v4))) & 63))
        _ = v40
    } else {
        if (!false) {
            var v41 bool = true
            for fi42 := 0; fi42 < 19; fi42++ {
                var v43 int32 = (v3 % (((v3 & (^(v3)))) | 1))
                _ = fi42
                _ = v43
            }
            _ = v41
        }
    }
    switch v2 {
    case "":
        var v44 float32 = 3.14159
        _ = v44
    default:
        var v45 uint32 = 746
        _ = v45
    }
    switch "foo bar" {
    case "":
        v46 := GenS4{F0: true, F1: 32767}
        _ = v46
    case "hello":
        v47 := GenS3{F0: 1, F1: v1, F2: v4}
        _ = v47
    case "foo bar":
        var v48 int = (v1 + -2147483648)
        _ = v48
    }
    switch v2 {
    case "foo bar":
        v2 = "a"
    case "":
        v49 := []int{(v1 | v1), (v1 * (v1 + v1)), (v1 + v1), v1}
        _ = v49
    case "hello":
        v50 := GenS5{F0: 42, F1: -2, F2: -9223372036854775808}
        _ = v50
    default:
        v51 := GenS6{F0: 978}
        _ = v51
    }
    for fi52 := 0; fi52 < 10; fi52++ {
        switch (v2 + "") {
        case "foo bar":
            switch (fi52 ^ 1) {
            case -2147483648:
                fmt.Println(v4)
            case 32767:
                v53 := []int{fi52, (fi52 & fi52), (v1 & 1)}
                _ = v53
            }
        case "":
            v54 := []int{}
            _ = v54
        }
        if false {
            v55 := GenS2{F0: 0, F1: 100}
            v4 -= v4
            _ = v55
        } else {
            fmt.Println(fi52)
        }
        _ = fi52
    }
    if true {
        var v56 uint64 = 905
        v57 := []int{v1, (v1 | v1), v1, 32768}
        _ = v56
        _ = v57
    } else {
        v1 += v1
        var v58 string = ""
        _ = v58
    }
    switch "a" {
    case "\n\t":
        switch (v1 + (-(v1))) {
        case 32768:
            v1 += v1
        case 32767:
            var v59 int32 = v3
            _ = v59
        case 128:
            var v60 bool = (!(3.14159 < 1.5))
            _ = v60
        }
    case "hello":
        var v61 uint64 = 0
        _ = v61
    }
    v62 := []int{int(v3), v1, v1}
    for fi63 := 0; fi63 < 8; fi63++ {
        var v64 int32 = -1
        v65 := GenS5{F0: 65535, F1: 1, F2: genFn1(168, -2147483648, fi63)}
        v66 := GenS1{F0: false}
        _ = fi63
        _ = v64
        _ = v65
        _ = v66
    }
    v4 = v4
    if (false && (true || true)) {
        if len(v62) > 0 {
            v62[0] = v1
        }
        v67 := GenS7{F0: 587, F1: 1, F2: (v3 * (v3 >> ((v4) & 63)))}
        _ = v67
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v62:", len(v62), v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v62
}
