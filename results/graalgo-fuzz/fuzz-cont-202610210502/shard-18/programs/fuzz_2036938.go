package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS4 struct { F0 int64; F1 uint; F2 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 int; F1 int32; F2 uint32 }

func main() {
    v1 := make(map[int]int)
    var v2 int8 = 1
    var v3 bool = false
    var v4 float64 = (3.14159 - -2.25)
    var v5 int16 = -32768
    v6 := make(map[string]string)
    if (!v3) {
        v7 := GenS1{F0: 18446744073709551615}
        if v3 {
            v8 := GenS1{F0: 42}
            v9 := GenS1{F0: 0}
            if (4294967295 == 0) {
                v10 := v1[-1]
                v11 := GenS1{F0: 18446744073709551615}
                v2 -= v2
                _ = v10
                _ = v11
            }
            _ = v8
            _ = v9
        } else {
            var v12 uint16 = 32767
            for fi13 := 0; fi13 < 20; fi13++ {
                var v14 uint16 = v12
                v15 := GenS2{F0: ((float32(v4) > (((float32((v4 - 0.0)) * 0.0) - float32(v4)) + 1.5)) && (913 > 42))}
                fmt.Println(v5, v5)
                _ = fi13
                _ = v14
                _ = v15
            }
            v16 := GenS1{F0: 42}
            _ = v12
            _ = v16
        }
        _ = v7
    }
    if (4294967295 >= 580) {
        for fi17 := 0; fi17 < 20; fi17++ {
            v18 := []int{(fi17 + 193), fi17, (-(fi17))}
            _ = fi17
            _ = v18
        }
    }
    v4 += 0.1
    switch "" {
    case "hello":
        if false {
            v19 := []int{127, -2147483648, 909}
            _ = v19
        } else {
            switch "" {
            case "hello":
                switch 0 {
                case 989:
                    fmt.Println(v5, v2)
                default:
                    if v3 {
                        var v20 uint64 = 4294967295
                        var v21 uint8 = 127
                        for fi22 := 0; fi22 < 20; fi22++ {
                            for fi23 := 0; fi23 < 10; fi23++ {
                                v24 := []int{(fi22 * -129)}
                                var v25 uint16 = uint16(v20)
                                v26 := GenS3{F0: (^(v20)), F1: -9223372036854775808, F2: v25}
                                _ = fi23
                                _ = v24
                                _ = v25
                                _ = v26
                            }
                            v5 = v5
                            if (0.0 <= 1.5) {
                                v20 = v20
                                v4 -= (3.14159 * v4)
                                for fi27 := 0; fi27 < 1; fi27++ {
                                    fmt.Println(fi27, v5, fi22)
                                    v2 -= (v2 + 42)
                                    if (0 != -493) {
                                        v4 += -2.25
                                        v6["hello"] = (("a" + (("hello" + "") + "\n\t")) + (("\n\t" + "a") + ("foo bar" + "hello")))
                                        v28 := []int{(fi27 & fi27), fi27, fi22}
                                        _ = v28
                                    }
                                    var v29 uint8 = v21
                                    _ = fi27
                                    _ = v29
                                }
                            }
                            _ = fi22
                        }
                        _ = v20
                        _ = v21
                    }
                }
            case "":
                v30 := []int{-622, 2147483647, 869}
                _ = v30
            case "\n\t":
                if (v3 || true) {
                    var v31 int64 = -9223372036854775808
                    _ = v31
                } else {
                    if ((v4 > -2.25) && v3) {
                        var v32 uint64 = 0
                        _ = len(v1)
                        _ = v32
                    } else {
                        for fi33 := 0; fi33 < 9; fi33++ {
                            var v34 uint8 = 212
                            if (fi33 != (^(fi33))) {
                                switch fi33 {
                                case 32767:
                                    for fi35 := 0; fi35 < 20; fi35++ {
                                        _ = len(v1)
                                        v4 -= v4
                                        if (v5 >= (v5 + v5)) {
                                            var v36 uint16 = 537
                                            if (v36 >= (v36 ^ (v36 | (v36 >> ((0) & 63))))) {
                                                v37 := GenS2{F0: true}
                                                _ = v37
                                            }
                                            _ = v36
                                        } else {
                                            if false {
                                                _ = len(v6)
                                                if (0.0 <= 0.0) {
                                                    var v38 uint = 376
                                                    v38 = (v38 % (((v38 * (v38 ^ (^(v38))))) | 1))
                                                    var v39 float32 = 1e10
                                                    _ = v38
                                                    _ = v39
                                                }
                                                v3 = (42 >= 0)
                                                fmt.Println(fi35, v5)
                                            }
                                            v40 := []int{}
                                            v2 = v2
                                            if false {
                                                v41 := GenS1{F0: 2}
                                                var v42 uint64 = 42
                                                if ("hello" != (("a" + (("\n\t" + ("hello" + ("hello" + "foo bar"))) + "foo bar")) + ("foo bar" + "\n\t"))) {
                                                    if false {
                                                        v43 := GenS3{F0: (v42 & 4294967295), F1: 9223372036854775807, F2: 65535}
                                                        var v44 int = 42
                                                        fmt.Println(v2, v42)
                                                        _ = v43
                                                        _ = v44
                                                    } else {
                                                        v45 := GenS4{F0: -9223372036854775808, F1: 4294967295, F2: 2147483647}
                                                        _ = len(v6)
                                                        _ = v45
                                                    }
                                                    v46 := GenS2{F0: (v3 || true)}
                                                    var v47 float64 = v4
                                                    _ = v46
                                                    _ = v47
                                                } else {
                                                    if v3 {
                                                        if (2 <= 2147483648) {
                                                            v3 = v3
                                                            var v48 string = "\n\t"
                                                            var v49 int = fi35
                                                            _ = v48
                                                            _ = v49
                                                        } else {
                                                            var v50 uint64 = (v42 >> ((865) & 63))
                                                            if false {
                                                                if len(v40) >= 2 {
                                                                    v40 = v40[1:len(v40)]
                                                                }
                                                                var v51 uint = 255
                                                                fmt.Println(v3, fi33, v2)
                                                                if true {
                                                                    var v52 uint8 = v34
                                                                    v51 += 338
                                                                    v4 = v4
                                                                    _ = v52
                                                                } else {
                                                                    var v53 int16 = 609
                                                                    if v3 {
                                                                        var v54 int = fi35
                                                                        _ = v54
                                                                    }
                                                                    fmt.Println(v51, v34, v2)
                                                                    _ = v53
                                                                }
                                                                v34 = (^(v34))
                                                                _ = v51
                                                            } else {
                                                                var v55 uint16 = 32767
                                                                _ = v55
                                                            }
                                                            _ = v50
                                                        }
                                                    } else {
                                                        switch (fi35 << ((0) & 63)) {
                                                        case -1:
                                                            var v56 int = 668
                                                            _ = v56
                                                        default:
                                                            var v57 uint8 = 1
                                                            _ = v57
                                                        }
                                                        switch fi35 {
                                                        case 127:
                                                            v42 -= (v42 % ((v42) | 1))
                                                        }
                                                    }
                                                    v5 = v5
                                                    if v3 {
                                                        v4 += (v4 / float64(1e10))
                                                        v34 += v34
                                                    }
                                                }
                                                var v58 bool = false
                                                if ((v58 && v3) || (float32((float64((1.5 * 0.0)) / (v4 + v4))) >= 1e10)) {
                                                    delete(v1, (fi35 & (^(fi35))))
                                                }
                                                _ = v41
                                                _ = v42
                                                _ = v58
                                            } else {
                                                if v3 {
                                                    fmt.Println(v3)
                                                    v6["hello"] = ("\n\t" + ("foo bar" + ("" + "")))
                                                }
                                                if v3 {
                                                    v40 = append(v40, (fi35 * fi35))
                                                } else {
                                                    if (!(("\n\t" + "") <= "hello")) {
                                                        switch "foo bar" {
                                                        case "":
                                                            if v3 {
                                                                var v59 float64 = (v4 * v4)
                                                                _ = v59
                                                            }
                                                        case "a":
                                                            v34 -= (v34 / ((v34) | 1))
                                                        }
                                                        var v60 string = ("\n\t" + "a")
                                                        _ = v60
                                                    } else {
                                                        if len(v40) >= 2 {
                                                            v40 = v40[1:len(v40)]
                                                        }
                                                        delete(v6, ((("" + ("\n\t" + "\n\t")) + ("hello" + "")) + "hello"))
                                                    }
                                                    var v61 float32 = (0.1 + 1.5)
                                                    if ("foo bar" >= "hello") {
                                                        delete(v1, fi35)
                                                        v62 := GenS5{F0: (-(v5))}
                                                        _ = v62
                                                    }
                                                    _ = v61
                                                }
                                                v63 := v6["foo bar"]
                                                _ = v63
                                            }
                                            var v64 int8 = int8(v5)
                                            _ = v40
                                            _ = v64
                                        }
                                        _ = fi35
                                    }
                                case 4294967295:
                                    v5 += -559
                                case -9223372036854775808:
                                    if (!(!(v5 != v5))) {
                                        for fi65 := 0; fi65 < 16; fi65++ {
                                            v2 = v2
                                            var v66 bool = ("a" < "")
                                            var v67 float64 = float64(float32(v4))
                                            _ = fi65
                                            _ = v66
                                            _ = v67
                                        }
                                    }
                                default:
                                    v68 := []int{-980}
                                    _ = v68
                                }
                                if (!((v2 & v2) <= (v2 & (v2 ^ v2)))) {
                                    v69 := []int{652, fi33, (fi33 - 838)}
                                    var v70 int64 = 9223372036854775807
                                    var v71 uint64 = 0
                                    _ = v69
                                    _ = v70
                                    _ = v71
                                }
                                v3 = true
                            } else {
                                var v72 bool = (200 >= 255)
                                _ = v72
                            }
                            _ = fi33
                            _ = v34
                        }
                        var v73 string = "a"
                        v5 = v5
                        _ = v73
                    }
                    fmt.Println(v5)
                    v74 := []int{127}
                    _ = v74
                }
            default:
                v1[-9223372036854775808] = -9223372036854775808
            }
            if ((^(v2)) != v2) {
                for fi75 := 0; fi75 < 9; fi75++ {
                    fmt.Println(v2)
                    _ = fi75
                }
            }
        }
    case "foo bar":
        var v76 uint64 = 234
        _ = v76
    default:
        var v77 bool = (!(!((128 != 128) && v3)))
        _ = v77
    }
    v78 := v1[895]
    v4 = (((v4 + ((((v4 / 3.14159) + float64(((0.0 * 1e10) / 1.5))) / v4) + (v4 / v4))) * (v4 * float64(1.5))) + 1e10)
    var v79 uint64 = 18446744073709551615
    v80 := GenS3{F0: v79, F1: 0, F2: 0}
    v81 := GenS6{F0: 84, F1: 4294967295}
    if (true && (!v3)) {
        var v82 uint16 = 32767
        fmt.Println(v4)
        v78 -= (v78 / (((v78 << ((319) & 63))) | 1))
        _ = v82
    }
    var v83 uint8 = 42
    var v84 int32 = -2147483648
    var v85 uint16 = uint16(v83)
    var v86 int16 = 1000
    for fi87 := 0; fi87 < 13; fi87++ {
        fmt.Println(v4, v79)
        switch "\n\t" {
        case "hello":
            _ = len(v6)
        case "foo bar":
            v4 += (((3.14159 + v4) * v4) / ((v4 / v4) * v4))
        default:
            for fi88 := 0; fi88 < 4; fi88++ {
                var v89 int16 = (v5 ^ 32767)
                var v90 bool = false
                v6[("\n\t" + "")] = ""
                _ = fi88
                _ = v89
                _ = v90
            }
        }
        for fi91 := 0; fi91 < 11; fi91++ {
            v92 := GenS7{F0: -2147483648, F1: (v84 / ((v84) | 1)), F2: 6}
            _ = fi91
            _ = v92
        }
        _ = fi87
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v78
    _ = v79
    _ = v80
    _ = v81
    _ = v83
    _ = v84
    _ = v85
    _ = v86
}
