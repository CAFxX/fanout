package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int32 }
type GenS2 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 uint64; F1 uint32; F2 bool }
type GenS5 struct { F0 uint16; F1 int64; F2 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 string }
func genFn0(p0 string, p1 uint8) uint64 {
    switch 128 {
    case 939:
        p1 = p1
    case -65:
        var v5 uint32 = 996
        _ = v5
    }
    for fi6 := 0; fi6 < 9; fi6++ {
        if true {
            var v7 uint16 = 65535
            switch fi6 {
            case 114:
                if (4294967295 < 0) {
                    v8 := []int{fi6}
                    if len(v8) > 0 {
                        v8[0] = 0
                    }
                    if len(v8) > 0 {
                        v8[0] = fi6
                    }
                    _ = v8
                } else {
                    p1 += 255
                    for fi9 := 0; fi9 < 7; fi9++ {
                        v10 := []int{(fi9 - 32767), fi9, (fi6 % ((fi9) | 1))}
                        if true {
                            var v11 uint64 = 4294967295
                            fmt.Println(p1, v11)
                            var v12 string = ""
                            _ = v11
                            _ = v12
                        }
                        _ = fi9
                        _ = v10
                    }
                }
            }
            switch ("" + p0) {
            case "foo bar":
                v13 := GenS1{F0: (p0 + p0), F1: -32768, F2: -2147483648}
                _ = v13
            case "a":
                v14 := GenS1{F0: ((p0 + p0) + (p0 + p0)), F1: -32768, F2: 2147483647}
                _ = v14
            case "\n\t":
                var v15 bool = true
                _ = v15
            default:
                fmt.Println(p1, fi6)
            }
            _ = v7
        } else {
            var v16 int = (fi6 & (fi6 & (fi6 - (fi6 * fi6))))
            v17 := []int{fi6, fi6, (v16 / ((fi6) | 1))}
            _ = v16
            _ = v17
        }
        p1 = (p1 & 255)
        fmt.Println(p0, fi6, p1)
        _ = fi6
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        if (true || true) {
            fmt.Println(fi18)
        } else {
            if true {
                p0 = ((("" + p0) + p0) + ((p0 + (p0 + ("\n\t" + p0))) + p0))
                switch ("foo bar" + (p0 + "foo bar")) {
                case "":
                    p0 = p0
                case "hello":
                    p0 = p0
                default:
                    if (1.5 >= 1e10) {
                        var v19 bool = false
                        _ = v19
                    }
                }
            }
            var v20 float64 = 1.5
            for fi21 := 0; fi21 < 15; fi21++ {
                v22 := GenS2{F0: 4294967295, F1: 0, F2: 224}
                switch "a" {
                case "":
                    p1 = (p1 % (((p1 ^ (p1 * (p1 * 71)))) | 1))
                case "hello":
                    var v23 uint8 = (p1 - p1)
                    _ = v23
                default:
                    v20 = -2.25
                }
                var v24 int8 = 66
                _ = fi21
                _ = v22
                _ = v24
            }
            _ = v20
        }
        for fi25 := 0; fi25 < 6; fi25++ {
            switch ((p0 + p0) + p0) {
            case "\n\t":
                if (!true) {
                    fmt.Println(fi25, p0)
                    if true {
                        var v26 int8 = 0
                        fmt.Println(v26)
                        var v27 int16 = -424
                        _ = v26
                        _ = v27
                    }
                    p0 = "\n\t"
                } else {
                    p0 = (((("foo bar" + p0) + ((p0 + ((p0 + (((((p0 + "hello") + (p0 + ((p0 + ("a" + p0)) + p0))) + ((p0 + p0) + p0)) + "foo bar") + (("a" + "\n\t") + p0))) + p0)) + p0)) + p0) + (p0 + p0))
                    p0 = "hello"
                }
            case "hello":
                if false {
                    p0 = p0
                    v28 := GenS3{F0: true, F1: p0}
                    _ = v28
                } else {
                    if false {
                        if false {
                            var v29 uint = 54
                            _ = v29
                        } else {
                            var v30 uint32 = 65535
                            if false {
                                var v31 int8 = 127
                                v32 := GenS2{F0: 0, F1: int16(v31), F2: p1}
                                _ = v31
                                _ = v32
                            } else {
                                p0 = "hello"
                                v33 := GenS3{F0: ((fi18 - fi18) == fi18), F1: (("" + p0) + ((((("" + ("" + "a")) + (p0 + "\n\t")) + "a") + (p0 + (p0 + ""))) + "a"))}
                                _ = v33
                            }
                            var v34 int64 = 56
                            _ = v30
                            _ = v34
                        }
                    }
                    p1 += (^(p1))
                }
            case "":
                v35 := GenS4{F0: 506, F1: 4294967295, F2: ((false && true) && (true && (1 < 42)))}
                _ = v35
            }
            _ = fi25
        }
        for fi36 := 0; fi36 < 3; fi36++ {
            v37 := GenS5{F0: 42, F1: -872, F2: fi36}
            v38 := GenS5{F0: 1, F1: 1, F2: 32767}
            var v39 uint64 = 765
            if false {
                v40 := []int{}
                v39 -= (^(v39))
                if len(v40) > 0 {
                    v40[0] = fi36
                }
                _ = v40
            }
            _ = fi36
            _ = v37
            _ = v38
            _ = v39
        }
        _ = fi18
    }
    var v41 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = v41
    return v41
}
func genFn1(p0 int16, p1 uint32, p2 uint8) int32 {
    var v42 uint32 = (^(p1))
    p0 = p0
    for fi43 := 0; fi43 < 7; fi43++ {
        switch "\n\t" {
        case "":
            v44 := GenS6{F0: true}
            _ = v44
        default:
            v45 := GenS2{F0: 4294967295, F1: int16(p1), F2: (p2 & p2)}
            _ = v45
        }
        fmt.Println(v42, v42)
        _ = fi43
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v42
    return -2147483648
}

func main() {
    var v1 string = "a"
    var v2 uint8 = 128
    v3 := make(map[uint32]int)
    var v4 int16 = 447
    _ = len(v3)
    v1 = v1
    v46 := []int{-574}
    for fi47 := 0; fi47 < 6; fi47++ {
        if true {
            var v48 string = "foo bar"
            v49 := GenS1{F0: (v1 + v1), F1: 1, F2: 1}
            var v50 string = v48
            _ = v48
            _ = v49
            _ = v50
        } else {
            var v51 float32 = 1e10
            _ = v51
        }
        fmt.Println(fi47, v4)
        _ = fi47
    }
    if false {
        v4 += 42
        v52 := GenS6{F0: (-1 > 1)}
        _ = v52
    }
    var v53 int = -833
    switch (v53 * v53) {
    case 2147483647:
        switch (v53 * (v53 ^ (v53 | v53))) {
        case -198:
            var v54 string = ((v1 + ((v1 + v1) + v1)) + v1)
            _ = v54
        case 127:
            for fi55 := 0; fi55 < 20; fi55++ {
                v56 := GenS2{F0: 4294967295, F1: (v4 % ((v4) | 1)), F2: v2}
                var v57 float32 = 3.14159
                _ = fi55
                _ = v56
                _ = v57
            }
        }
    case 4294967295:
        var v58 float32 = ((((0.0 + 0.1) - (0.0 / -2.25)) - 3.14159) * 1e10)
        _ = v58
    case 0:
        var v59 bool = (true || false)
        _ = v59
    }
    fmt.Println(v2, v1)
    for fi60 := 0; fi60 < 15; fi60++ {
        var v61 float64 = (-2.25 / (-2.25 + -2.25))
        if len(v46) > 0 {
            v46[0] = 792
        }
        _ = fi60
        _ = v61
    }
    var v62 int = (v53 << ((18446744073709551615) & 63))
    if false {
        switch v53 {
        case 32768:
            var v63 int64 = -223
            _ = v63
        case 970:
            for fi64 := 0; fi64 < 18; fi64++ {
                var v65 uint = 42
                _ = fi64
                _ = v65
            }
        case 42:
            delete(v3, 288)
        }
        var v66 string = (v1 + v1)
        if len(v46) >= 2 {
            v46 = v46[1:len(v46)]
        }
        _ = v66
    } else {
        var v67 uint = 18446744073709551615
        if false {
            switch (v62 | (-(v62))) {
            case 131:
                v68 := GenS7{F0: 994}
                _ = v68
            case 42:
                switch -957 {
                case -493:
                    if false {
                        v69 := GenS8{F0: true}
                        switch v1 {
                        case "hello":
                            v70 := v3[165]
                            _ = v70
                        case "\n\t":
                            v71 := GenS9{F0: 42}
                            _ = v71
                        case "foo bar":
                            v67 -= (^(v67))
                        default:
                            v72 := GenS10{F0: v67}
                            _ = v72
                        }
                        var v73 uint = 255
                        _ = v69
                        _ = v73
                    }
                case 4294967295:
                    v74 := GenS11{F0: "foo bar"}
                    _ = v74
                default:
                    var v75 int32 = genFn1(v4, 2147483648, 0)
                    _ = v75
                }
            default:
                switch (v1 + (v1 + v1)) {
                case "\n\t":
                    v53 += 4294967295
                case "":
                    _ = len(v3)
                default:
                    if true {
                        var v76 float64 = -2.25
                        v2 += (v2 | (^(v2)))
                        for fi77 := 0; fi77 < 2; fi77++ {
                            for fi78 := 0; fi78 < 6; fi78++ {
                                var v79 float64 = float64(3.14159)
                                var v80 uint64 = 4294967295
                                switch v1 {
                                case "a":
                                    var v81 int32 = 0
                                    _ = v81
                                case "":
                                    fmt.Println(v4)
                                case "\n\t":
                                    v62 = (v53 % ((v53) | 1))
                                }
                                _ = fi78
                                _ = v79
                                _ = v80
                            }
                            _ = fi77
                        }
                        _ = v76
                    } else {
                        switch ((v1 + v1) + "a") {
                        case "a":
                            var v82 int16 = int16(v62)
                            _ = v82
                        case "\n\t":
                            v83 := GenS5{F0: 0, F1: 9223372036854775807, F2: 677}
                            _ = v83
                        default:
                            var v84 int8 = -56
                            _ = v84
                        }
                        var v85 int16 = v4
                        var v86 int = v62
                        _ = v85
                        _ = v86
                    }
                }
            }
            delete(v3, 1)
            var v87 int64 = 0
            _ = v87
        } else {
            var v88 bool = false
            _ = v88
        }
        var v89 uint = v67
        _ = v67
        _ = v89
    }
    var v90 float64 = 0.1
    fmt.Println(v62)
    delete(v3, 412)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v53:", v53)
    fmt.Println("v62:", v62)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v53
    _ = v62
    _ = v90
}
