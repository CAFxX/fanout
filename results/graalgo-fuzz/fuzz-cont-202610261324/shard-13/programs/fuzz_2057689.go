package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 int; F1 string; F2 uint16 }
type GenS9 struct { F0 int64; F1 int16; F2 bool }
func genFn0(p0 bool) int16 {
    v7 := GenS1{F0: 127, F1: 2, F2: 127}
    fmt.Println(p0)
    for fi8 := 0; fi8 < 3; fi8++ {
        p0 = (!p0)
        v9 := GenS1{F0: 131, F1: 128, F2: 1}
        p0 = (!false)
        _ = fi8
        _ = v9
    }
    var v10 uint = 1
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS2{F0: 1, F1: 120}
        if false {
            for fi13 := 0; fi13 < 13; fi13++ {
                p0 = true
                var v14 uint16 = 42
                p0 = false
                _ = fi13
                _ = v14
            }
            if p0 {
                v15 := GenS3{F0: fi11, F1: -2147483648}
                for fi16 := 0; fi16 < 20; fi16++ {
                    switch fi11 {
                    case 0:
                        if (-128 == -1) {
                            var v17 int = fi11
                            p0 = ((42 >= 9223372036854775807) || p0)
                            _ = v17
                        } else {
                            p0 = ((fi16 & (fi16 >> ((v10) & 63))) == (fi11 / ((fi16) | 1)))
                            v18 := GenS4{F0: fi11}
                            v19 := []int{127, (-(fi11)), (^(fi11)), (fi16 + (fi11 % ((fi11) | 1)))}
                            fmt.Println(v10)
                            _ = v18
                            _ = v19
                        }
                        if false {
                            p0 = p0
                            v10 = v10
                        } else {
                            if ("\n\t" < ("\n\t" + "\n\t")) {
                                if (-128 == -1) {
                                    var v20 int8 = 127
                                    var v21 int = fi11
                                    var v22 float32 = float32(float64(3.14159))
                                    v22 += v22
                                    if (905 >= v10) {
                                        v23 := GenS1{F0: 255, F1: 243, F2: 1}
                                        v24 := []int{(fi11 >> ((v10) & 63)), (v21 / ((42) | 1)), (fi16 - (fi16 - 2147483647)), (v21 & fi16)}
                                        v20 -= 34
                                        v24 = append(v24, fi11)
                                        var v25 bool = (-338 == 1)
                                        v26 := GenS5{F0: (fi11 | v21)}
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                        _ = v26
                                    }
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                }
                                if true {
                                    var v27 uint32 = 65535
                                    var v28 int64 = -613
                                    if false {
                                        v29 := GenS6{F0: 4294967295}
                                        _ = v29
                                    }
                                    _ = v27
                                    _ = v28
                                }
                                if (-2147483648 != 1) {
                                    v30 := GenS7{F0: (p0 && (!p0)), F1: -1}
                                    _ = v30
                                }
                                if (0 != int8(fi16)) {
                                    switch "a" {
                                    case "foo bar":
                                        var v31 int32 = -2147483648
                                        _ = v31
                                    }
                                    if (42 <= 9223372036854775807) {
                                        p0 = ("a" > "foo bar")
                                        v10 += 255
                                    }
                                    v32 := []int{fi16, (fi11 | 499), 622, (fi16 + (fi11 | (fi16 + (fi11 / ((fi16) | 1)))))}
                                    fmt.Println(v10)
                                    fmt.Println(p0)
                                    v33 := GenS6{F0: fi11}
                                    _ = v32
                                    _ = v33
                                } else {
                                    fmt.Println(v10)
                                    var v34 int32 = 0
                                    v35 := GenS8{F0: (fi11 / ((-449) | 1)), F1: "hello", F2: 2}
                                    _ = v34
                                    _ = v35
                                }
                                v36 := []int{(fi16 << (((v10 | 4294967295)) & 63)), int(v10), fi11, 2147483647}
                                _ = v36
                            } else {
                                switch ((("foo bar" + "hello") + "foo bar") + "") {
                                case "foo bar":
                                    v37 := []int{fi16}
                                    _ = v37
                                }
                                var v38 int64 = 0
                                _ = v38
                            }
                            v39 := []int{128, fi16, fi11, 2147483647}
                            var v40 uint64 = 767
                            _ = v39
                            _ = v40
                        }
                    case 735:
                        if false {
                            v10 = v10
                            v41 := []int{(fi16 - (^(fi16)))}
                            _ = v41
                        } else {
                            v42 := GenS4{F0: (fi16 << (((v10 ^ 528)) & 63))}
                            fmt.Println(fi11, fi11)
                            _ = v42
                        }
                    case 32767:
                        var v43 uint32 = 42
                        _ = v43
                    default:
                        var v44 string = ("\n\t" + "\n\t")
                        _ = v44
                    }
                    v10 = 2
                    _ = fi16
                }
                _ = v15
            }
        }
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = v10
    _ = v7
    return 330
}

func main() {
    var v1 uint = 757
    var v2 bool = (-2147483648 >= int32(v1))
    var v3 float64 = 1.5
    var v4 uint8 = 42
    var v5 uint = (v1 >> ((18446744073709551615) & 63))
    var v6 string = ""
    v4 = 128
    for fi45 := 0; fi45 < 8; fi45++ {
        fmt.Println(v2, v2)
        var v46 int8 = -128
        _ = fi45
        _ = v46
    }
    v5 = (v5 | (v5 & 899))
    v47 := []int{42, 127}
    switch 128 {
    case 423:
        for fi48 := 0; fi48 < 1; fi48++ {
            v3 -= (float64((0.0 - -2.25)) / (v3 + v3))
            v49 := GenS5{F0: 816}
            _ = fi48
            _ = v49
        }
    }
    if v2 {
        if len(v47) > 0 {
            v47[0] = 32768
        }
        _ = len(v47)
    } else {
        if v2 {
            if v2 {
                if false {
                    if (float32(((float64(0.0) / float64(1e10)) + ((3.14159 / v3) + 0.0))) == 1e10) {
                        v1 -= 0
                        switch "a" {
                        case "a":
                            var v50 bool = (v6 < v6)
                            _ = v50
                        case "foo bar":
                            v51 := GenS6{F0: -9223372036854775808}
                            _ = v51
                        case "hello":
                            var v52 uint16 = 42
                            _ = v52
                        }
                        for fi53 := 0; fi53 < 10; fi53++ {
                            for fi54 := 0; fi54 < 16; fi54++ {
                                switch v6 {
                                case "hello":
                                    var v55 uint64 = uint64(v5)
                                    _ = v55
                                case "foo bar":
                                    if true {
                                        v5 -= (v1 * v5)
                                    }
                                case "a":
                                    v1 = (^(v1))
                                }
                                _ = fi54
                            }
                            _ = fi53
                        }
                    } else {
                        v47 = append(v47, -9223372036854775808)
                    }
                    v56 := GenS8{F0: -1, F1: ((((((("foo bar" + (v6 + v6)) + "hello") + v6) + v6) + (v6 + v6)) + v6) + ((v6 + v6) + v6)), F2: 1}
                    fmt.Println(v5)
                    _ = v56
                } else {
                    v57 := GenS4{F0: 341}
                    fmt.Println(v4)
                    _ = v57
                }
                var v58 int64 = 1
                _ = v58
            } else {
                var v59 float64 = float64(0.1)
                switch 128 {
                case -79:
                    var v60 int = -2147483648
                    _ = v60
                }
                if (!v2) {
                    if (v2 && v2) {
                        var v61 bool = v2
                        _ = v61
                    }
                } else {
                    v5 = (v1 ^ 895)
                }
                _ = v59
            }
        } else {
            var v62 bool = true
            _ = v62
        }
        var v63 bool = (!false)
        _ = v63
    }
    v64 := GenS2{F0: 64, F1: 972}
    v6 = (v6 + (((v6 + "a") + "\n\t") + v6))
    v6 = ""
    var v65 uint64 = 18446744073709551615
    v5 = v1
    for fi66 := 0; fi66 < 18; fi66++ {
        var v67 float32 = (3.14159 / 3.14159)
        v68 := GenS9{F0: 50, F1: genFn0(v2), F2: false}
        _ = fi66
        _ = v67
        _ = v68
    }
    _ = len(v47)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v64
    _ = v65
}
