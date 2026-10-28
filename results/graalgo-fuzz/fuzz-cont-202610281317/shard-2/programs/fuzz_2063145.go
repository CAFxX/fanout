package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 int8; F1 string; F2 uint8 }
type GenS4 struct { F0 int8; F1 int; F2 int16 }
func genFn0(p0 uint, p1 int8, p2 bool) uint16 {
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: ("hello" + "a"), F1: fi7}
        _ = fi7
        _ = v8
    }
    fmt.Println(p2, p2, p2)
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1(p0 int, p1 int16) uint {
    var v9 uint16 = 255
    p1 = p1
    v10 := GenS1{F0: "\n\t", F1: (p0 << ((255) & 63))}
    var v11 uint8 = 93
    v11 += v11
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 42
}

func main() {
    var v1 int8 = -128
    var v2 float32 = -2.25
    v3 := make(map[int]bool)
    var v4 int32 = -2147483648
    var v5 uint32 = 993
    v6 := make(map[uint32]bool)
    if true {
        v2 = ((1e10 / v2) / (v2 + float32((3.14159 - 0.1))))
        if true {
            v12 := GenS2{F0: v5, F1: uint(v5)}
            fmt.Println(v1, v4, v4)
            var v13 uint16 = 2
            _ = v12
            _ = v13
        } else {
            var v14 float64 = 0.0
            _ = v14
        }
    }
    v15 := v3[414]
    if v15 {
        for fi16 := 0; fi16 < 2; fi16++ {
            var v17 int64 = 1
            _ = fi16
            _ = v17
        }
        if (799 <= 0) {
            fmt.Println(v4, v1)
        } else {
            if v15 {
                var v18 int32 = int32(v5)
                _ = v18
            }
            v19 := []int{839, 32768, 281, -2147483648}
            v4 = -2147483648
            _ = v19
        }
        switch ((("a" + "hello") + (("\n\t" + ("" + "")) + ((("" + ("" + "hello")) + "\n\t") + ("foo bar" + ("\n\t" + "foo bar"))))) + "a") {
        case "foo bar":
            switch 1 {
            case 127:
                if (v1 != (^(v1))) {
                    v5 += v5
                } else {
                    switch "\n\t" {
                    case "":
                        if ((113 <= 33) && v15) {
                            v5 += (v5 & v5)
                            fmt.Println(v5, v1)
                        }
                    case "hello":
                        if v15 {
                            if ((v2 < ((v2 * (v2 + 1e10)) * v2)) && v15) {
                                v20 := []int{}
                                var v21 float64 = 0.1
                                _ = v20
                                _ = v21
                            }
                        }
                    case "a":
                        switch "foo bar" {
                        case "a":
                            if true {
                                for fi22 := 0; fi22 < 5; fi22++ {
                                    for fi23 := 0; fi23 < 3; fi23++ {
                                        switch ((("hello" + ("foo bar" + "")) + "foo bar") + "a") {
                                        case "hello":
                                            if v15 {
                                                v15 = v15
                                            } else {
                                                if (v15 && (328 != 1)) {
                                                    var v24 uint = 291
                                                    _ = v24
                                                } else {
                                                    v1 = (v1 % (((v1 & v1)) | 1))
                                                    if (v4 > 42) {
                                                        v25 := GenS2{F0: 2, F1: 65535}
                                                        if v15 {
                                                            v26 := GenS1{F0: "hello", F1: fi23}
                                                            _ = v26
                                                        } else {
                                                            v27 := []int{42, (fi23 * 1)}
                                                            v1 = 118
                                                            if v15 {
                                                                var v28 int16 = 42
                                                                v5 = v5
                                                                _ = v28
                                                            }
                                                            _ = v27
                                                        }
                                                        v2 += (0.1 / v2)
                                                        _ = v25
                                                    }
                                                }
                                                if (0 == 72) {
                                                    var v29 uint = 18446744073709551615
                                                    var v30 uint8 = 125
                                                    _ = v29
                                                    _ = v30
                                                }
                                                fmt.Println(fi22, v1)
                                            }
                                        }
                                        _ = fi23
                                    }
                                    _ = fi22
                                }
                                for fi31 := 0; fi31 < 16; fi31++ {
                                    for fi32 := 0; fi32 < 3; fi32++ {
                                        var v33 float32 = v2
                                        _ = fi32
                                        _ = v33
                                    }
                                    v34 := GenS3{F0: 100, F1: ("" + ("" + "")), F2: 2}
                                    _ = fi31
                                    _ = v34
                                }
                            }
                        case "foo bar":
                            var v35 int8 = v1
                            _ = v35
                        default:
                            var v36 float64 = 1.5
                            _ = v36
                        }
                    default:
                        var v37 int = 128
                        _ = v37
                    }
                    var v38 int8 = (v1 | v1)
                    _ = v38
                }
            default:
                fmt.Println(v1, v2)
            }
        case "hello":
            if (!false) {
                var v39 uint16 = 17
                _ = v39
            } else {
                var v40 int = int(v4)
                v41 := GenS3{F0: v1, F1: "\n\t", F2: 2}
                for fi42 := 0; fi42 < 8; fi42++ {
                    v43 := []int{}
                    var v44 int16 = 280
                    _ = len(v6)
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                _ = v40
                _ = v41
            }
        case "":
            fmt.Println(v4, v4, v1)
        }
    }
    v45 := GenS3{F0: (v1 - (v1 << ((uint(v5)) & 63))), F1: "foo bar", F2: 128}
    switch "a" {
    case "hello":
        var v46 uint32 = v5
        _ = v46
    case "":
        var v47 bool = (v2 == 1.5)
        _ = v47
    default:
        for fi48 := 0; fi48 < 2; fi48++ {
            v49 := []int{fi48, fi48}
            for fi50 := 0; fi50 < 13; fi50++ {
                v49 = append(v49, 1)
                switch (("foo bar" + "foo bar") + ("a" + "hello")) {
                case "a":
                    var v51 string = ""
                    _ = v51
                case "hello":
                    var v52 int8 = v1
                    _ = v52
                }
                _ = fi50
            }
            switch "a" {
            case "":
                fmt.Println(v2, v1, v15)
            default:
                v5 -= v5
            }
            _ = fi48
            _ = v49
        }
    }
    for fi53 := 0; fi53 < 3; fi53++ {
        v54 := []int{(fi53 & 127), (fi53 % ((fi53) | 1))}
        v55 := GenS3{F0: v1, F1: "a", F2: 255}
        var v56 int64 = 9223372036854775807
        _ = fi53
        _ = v54
        _ = v55
        _ = v56
    }
    var v57 int8 = 100
    v58 := GenS3{F0: (-(v1)), F1: (("\n\t" + "foo bar") + "a"), F2: 1}
    v59 := GenS1{F0: ("a" + "a"), F1: 2147483647}
    v60 := GenS4{F0: (v1 >> ((71) & 63)), F1: 9223372036854775807, F2: 32767}
    var v61 uint = 89
    fmt.Println(v15, v15, v1)
    var v62 uint32 = (v5 / ((4294967295) | 1))
    v15 = (!(!(!v15)))
    var v63 uint8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v62
    _ = v63
}
