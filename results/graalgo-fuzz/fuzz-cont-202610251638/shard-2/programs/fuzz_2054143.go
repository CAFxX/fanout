package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string; F1 uint32; F2 bool }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int; F1 uint; F2 int8 }
type GenS9 struct { F0 uint32; F1 bool }
type GenS10 struct { F0 int64; F1 uint16 }
func genFn0(p0 uint64, p1 uint32, p2 uint) uint {
    switch "a" {
    case "":
        for fi4 := 0; fi4 < 7; fi4++ {
            p2 += 1
            var v5 string = ((("a" + ("\n\t" + (("a" + "\n\t") + (("\n\t" + ((("" + ("foo bar" + ("" + "hello"))) + "") + "")) + "hello")))) + ("a" + ("\n\t" + "foo bar"))) + "a")
            v6 := []int{}
            _ = fi4
            _ = v5
            _ = v6
        }
    default:
        for fi7 := 0; fi7 < 15; fi7++ {
            p1 += 4294967295
            if false {
                switch "hello" {
                case "\n\t":
                    var v8 uint64 = (p0 / (((p0 >> ((p2) & 63))) | 1))
                    _ = v8
                }
            } else {
                p1 -= p1
            }
            _ = fi7
        }
    }
    v9 := GenS1{F0: 100}
    if false {
        fmt.Println(p0)
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 int32 = -928
            p2 -= (p2 << ((p2) & 63))
            v12 := []int{(-(fi10)), 895, -247}
            _ = fi10
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 9; fi13++ {
            v14 := GenS2{F0: 1}
            _ = fi13
            _ = v14
        }
    }
    v15 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v9
    return p2
}
func genFn1(p0 string, p1 int32, p2 uint8) int {
    for fi16 := 0; fi16 < 1; fi16++ {
        if (!false) {
            v17 := []int{fi16}
            _ = v17
        }
        _ = fi16
    }
    p1 -= int32(p2)
    v18 := []int{127, 32767, 121, 42}
    for fi19 := 0; fi19 < 3; fi19++ {
        switch p0 {
        case "\n\t":
            var v20 string = p0
            _ = v20
        default:
            v21 := GenS2{F0: 1}
            _ = v21
        }
        _ = fi19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return 1
}

func main() {
    var v1 int32 = -1
    v2 := make(map[string]int)
    var v3 string = "a"
    v1 = v1
    for fi22 := 0; fi22 < 9; fi22++ {
        var v23 int64 = 0
        var v24 float64 = (3.14159 * (((-2.25 / 3.14159) * 0.0) + -2.25))
        v3 = "foo bar"
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        for fi26 := 0; fi26 < 4; fi26++ {
            v1 = (v1 - v1)
            v2[""] = (fi25 / ((859) | 1))
            var v27 uint = 255
            var v28 int64 = -1
            var v29 float32 = 3.14159
            if true {
                if true {
                    v30 := GenS3{F0: fi25, F1: -310}
                    var v31 uint16 = 1
                    var v32 uint8 = 226
                    _ = v30
                    _ = v31
                    _ = v32
                } else {
                    v2[("foo bar" + v3)] = fi25
                }
                if (v27 <= (v27 | v27)) {
                    var v33 float32 = (v29 / v29)
                    fmt.Println(v1)
                    _ = v33
                } else {
                    v28 = 9223372036854775807
                    var v34 uint = v27
                    v28 = v28
                    _ = v34
                }
            } else {
                fmt.Println(v1, v1, v3)
                v35 := GenS4{F0: 672}
                if (!(!false)) {
                    v36 := GenS4{F0: 253}
                    _ = v36
                }
                _ = v35
            }
            v37 := []int{fi25, fi26}
            if ("\n\t" >= (v3 + (v3 + "\n\t"))) {
                if ((("a" + (v3 + v3)) + "hello") != v3) {
                    var v38 uint = v27
                    var v39 int32 = (v1 + (-(v1)))
                    switch fi26 {
                    case 420:
                        switch (v3 + (((v3 + v3) + v3) + v3)) {
                        case "":
                            v40 := GenS1{F0: 1}
                            _ = v40
                        case "hello":
                            switch 128 {
                            case 32768:
                                fmt.Println(v39)
                            case -979:
                                if (v28 >= 42) {
                                    v41 := GenS1{F0: -1}
                                    var v42 int32 = (v1 << ((v27) & 63))
                                    _ = v41
                                    _ = v42
                                }
                            }
                        }
                    default:
                        var v43 uint8 = 255
                        _ = v43
                    }
                    _ = v38
                    _ = v39
                }
            }
            var v44 int64 = v28
            v45 := v2[v3]
            v44 = v44
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
            _ = v37
            _ = v44
            _ = v45
        }
        _ = fi25
    }
    if (0 <= genFn0(42, 65535, 42)) {
        if false {
            fmt.Println(v3, v1)
            if (42 != 950) {
                var v46 float64 = (1e10 - (1e10 * 0.1))
                _ = v46
            } else {
                var v47 uint16 = 255
                v48 := GenS5{F0: v3, F1: 711, F2: (!(9223372036854775807 <= 42))}
                _ = v47
                _ = v48
            }
            if true {
                v49 := []int{2147483647, 0}
                _ = v49
            }
        }
        v50 := GenS6{F0: (2 <= 42)}
        v51 := GenS7{F0: 255}
        _ = v50
        _ = v51
    } else {
        if (!((1 == 0) && (!(((3.14159 * 0.1) + (0.1 - -2.25)) != float64(1.5))))) {
            for fi52 := 0; fi52 < 15; fi52++ {
                if (29 >= 65535) {
                    if true {
                        if true {
                            if true {
                                v53 := GenS4{F0: 4294967295}
                                v1 -= (v1 >> ((794) & 63))
                                for fi54 := 0; fi54 < 13; fi54++ {
                                    delete(v2, "foo bar")
                                    v1 += v1
                                    var v55 float64 = 1e10
                                    _ = fi54
                                    _ = v55
                                }
                                _ = v53
                            }
                            v3 = "hello"
                        }
                        var v56 uint = 926
                        v57 := []int{fi52}
                        _ = v56
                        _ = v57
                    }
                    var v58 uint64 = 4294967295
                    v1 += v1
                    _ = v58
                } else {
                    for fi59 := 0; fi59 < 12; fi59++ {
                        if (!(!false)) {
                            var v60 uint = 2
                            var v61 int64 = -4
                            _ = v60
                            _ = v61
                        } else {
                            v62 := GenS8{F0: fi52, F1: 749, F2: -1}
                            var v63 uint = 4294967295
                            _ = v62
                            _ = v63
                        }
                        _ = fi59
                    }
                    v64 := GenS1{F0: 1}
                    _ = v64
                }
                v65 := GenS3{F0: (fi52 | fi52), F1: 1}
                _ = fi52
                _ = v65
            }
            fmt.Println(v3, v3)
            if (1 != 1000) {
                var v66 uint64 = 382
                _ = v66
            } else {
                v67 := GenS9{F0: 2147483647, F1: (false || true)}
                for fi68 := 0; fi68 < 7; fi68++ {
                    switch fi68 {
                    case -150:
                        if (!(80 > 100)) {
                            var v69 float64 = 0.1
                            for fi70 := 0; fi70 < 7; fi70++ {
                                v71 := GenS1{F0: 47}
                                if true {
                                    v72 := []int{genFn1((v3 + ((("hello" + v3) + ("\n\t" + (("hello" + v3) + v3))) + "a")), v1, 154), (fi70 % ((fi68) | 1)), fi68, fi68}
                                    switch (fi68 >> ((248) & 63)) {
                                    case -2147483648:
                                        var v73 float32 = 0.0
                                        _ = v73
                                    case 32767:
                                        v3 = v3
                                        switch v3 {
                                        case "hello":
                                            var v74 int64 = 145
                                            _ = v74
                                        }
                                        fmt.Println(fi70, v3)
                                    }
                                    if true {
                                        v69 = 0.1
                                        v3 = (v3 + "a")
                                    }
                                    _ = v72
                                } else {
                                    var v75 int32 = v1
                                    _ = v75
                                }
                                var v76 uint = 42
                                _ = fi70
                                _ = v71
                                _ = v76
                            }
                            _ = v69
                        } else {
                            var v77 uint8 = 128
                            if true {
                                var v78 uint8 = 42
                                v79 := []int{}
                                _ = v78
                                _ = v79
                            }
                            var v80 string = "hello"
                            _ = v77
                            _ = v80
                        }
                    default:
                        var v81 int = genFn1(v3, 1, 128)
                        _ = v81
                    }
                    v1 -= v1
                    _ = fi68
                }
                _ = v67
            }
        } else {
            v82 := GenS4{F0: 2}
            v1 += 0
            _ = v82
        }
        var v83 int = 128
        _ = v83
    }
    v84 := []int{-1, 128, 4294967295, 718}
    v85 := GenS10{F0: 863, F1: 379}
    v3 = v3
    v86 := GenS2{F0: 127}
    fmt.Println(v1, v3)
    if false {
        var v87 uint16 = 26
        var v88 int64 = 0
        var v89 uint32 = 65535
        _ = v87
        _ = v88
        _ = v89
    }
    fmt.Println(v3)
    switch ("a" + (v3 + v3)) {
    case "\n\t":
        v3 = v3
    case "a":
        v1 -= 2147483647
    }
    fmt.Println(v1, v1)
    if (-2.25 != 1.5) {
        v90 := GenS3{F0: -858, F1: 1000}
        _ = v90
    }
    v3 = ((v3 + "\n\t") + v3)
    fmt.Println(v3, v1)
    for fi91 := 0; fi91 < 19; fi91++ {
        for fi92 := 0; fi92 < 19; fi92++ {
            var v93 int = fi92
            _ = fi92
            _ = v93
        }
        _ = fi91
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v84:", len(v84), v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v84
    _ = v85
    _ = v86
}
