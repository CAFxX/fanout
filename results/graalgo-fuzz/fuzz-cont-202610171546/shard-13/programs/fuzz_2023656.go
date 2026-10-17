package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint32; F1 uint; F2 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 int; F2 string }
type GenS6 struct { F0 bool; F1 uint32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int32; F1 uint8 }
type GenS9 struct { F0 string }
func genFn0(p0 int16, p1 string, p2 bool) uint64 {
    var v7 uint16 = 32767
    var v8 string = ""
    v9 := []int{}
    p0 = 0
    if (600 == 42) {
        fmt.Println(p0)
        switch -2147483648 {
        case 32767:
            var v10 string = p1
            _ = v10
        case 332:
            fmt.Println(p2, p0)
        case 0:
            switch "\n\t" {
            case "foo bar":
                v7 = 428
            case "hello":
                if p2 {
                    if (false && true) {
                        fmt.Println(p2, v8, p1)
                        var v11 int16 = p0
                        if p2 {
                            switch 127 {
                            case -2147483648:
                                v12 := GenS1{F0: 2, F1: 42}
                                _ = v12
                            case -439:
                                fmt.Println(v11)
                            case -9223372036854775808:
                                switch -948 {
                                case -9223372036854775808:
                                    for fi13 := 0; fi13 < 1; fi13++ {
                                        switch (fi13 - 127) {
                                        case -9223372036854775808:
                                            v8 = (("\n\t" + p1) + p1)
                                        case 0:
                                            v11 -= v11
                                        default:
                                            v11 -= -239
                                        }
                                        if ((127 != 127) && true) {
                                            p1 = v8
                                        } else {
                                            if true {
                                                if true {
                                                    p0 = p0
                                                    v8 = "a"
                                                    var v14 int32 = 2147483647
                                                    _ = v14
                                                }
                                            }
                                        }
                                        v15 := GenS2{F0: p2, F1: v11}
                                        _ = fi13
                                        _ = v15
                                    }
                                case 2147483647:
                                    p2 = (v7 < 1)
                                default:
                                    var v16 string = "hello"
                                    _ = v16
                                }
                            default:
                                var v17 uint64 = 42
                                _ = v17
                            }
                        }
                        _ = v11
                    }
                } else {
                    var v18 int16 = 32767
                    for fi19 := 0; fi19 < 16; fi19++ {
                        var v20 int32 = 2147483647
                        v21 := GenS3{F0: 412, F1: 42, F2: ("a" + ((v8 + p1) + v8))}
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    _ = v18
                }
            case "\n\t":
                var v22 uint64 = 4294967295
                _ = v22
            }
        }
        for fi23 := 0; fi23 < 19; fi23++ {
            p2 = (65535 < 255)
            if (26 != v7) {
                p0 -= p0
                v9 = append(v9, (fi23 % ((fi23) | 1)))
                var v24 int = -845
                _ = v24
            }
            if false {
                v8 = ("hello" + ("foo bar" + v8))
                _ = len(v9)
            }
            _ = fi23
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}
func genFn1(p0 bool) int {
    p0 = (2 < 42)
    v25 := []int{128, 0, 4294967295}
    fmt.Println(p0)
    _ = p0
    _ = v25
    return -527
}

func main() {
    var v1 uint = 2
    v2 := []string{"a"}
    v3 := []uint32{}
    var v4 int8 = 127
    v5 := []uint32{}
    v6 := make(map[string]bool)
    if true {
        v26 := GenS4{F0: 2}
        v4 = (v4 ^ v4)
        if (v1 < (^(v1))) {
            switch ("\n\t" + "foo bar") {
            case "\n\t":
                v1 = 42
            case "a":
                switch 2147483647 {
                case 128:
                    fmt.Println(v4, v4)
                }
            case "":
                var v27 uint32 = 65535
                _ = v27
            default:
                v1 = v1
            }
        }
        _ = v26
    }
    for fi28 := 0; fi28 < 6; fi28++ {
        v29 := GenS1{F0: 2, F1: 42}
        _ = fi28
        _ = v29
    }
    var v30 float64 = 0.0
    var v31 float64 = 3.14159
    v6[""] = (-671 < -1)
    if (0 < 2) {
        for fi32 := 0; fi32 < 16; fi32++ {
            v33 := GenS5{F0: 239, F1: fi32, F2: ((("foo bar" + ("a" + ("" + "hello"))) + "foo bar") + ("" + ("a" + (("foo bar" + "\n\t") + ""))))}
            v34 := GenS6{F0: false, F1: 4294967295}
            v35 := GenS3{F0: 119, F1: (v1 ^ 4294967295), F2: "a"}
            _ = fi32
            _ = v33
            _ = v34
            _ = v35
        }
        var v36 uint = v1
        _ = len(v6)
        _ = v36
    }
    var v37 uint = 2
    _ = len(v6)
    var v38 uint8 = 127
    var v39 bool = (float64(0.1) < v31)
    v31 += v30
    for fi40 := 0; fi40 < 5; fi40++ {
        if (((0.0 - 1.5) / float32((-2.25 - (float64(1.5) * float64(0.1))))) > -2.25) {
            var v41 uint32 = 42
            var v42 int32 = 2147483647
            switch (fi40 << (((^(v37))) & 63)) {
            case 9223372036854775807:
                v38 = v38
            default:
                var v43 uint32 = (v41 / ((v41) | 1))
                _ = v43
            }
            _ = v41
            _ = v42
        }
        var v44 uint8 = 132
        switch (((("" + ("\n\t" + ("hello" + ""))) + ("hello" + ("\n\t" + "\n\t"))) + "") + "foo bar") {
        case "hello":
            for fi45 := 0; fi45 < 7; fi45++ {
                if true {
                    v46 := GenS4{F0: 32767}
                    if v39 {
                        if v39 {
                            var v47 int = (fi45 + (fi40 / (((fi45 / (((fi40 / (((fi40 & (fi40 & 4294967295))) | 1))) | 1))) | 1)))
                            v44 -= v44
                            _ = v47
                        }
                        if true {
                            _ = len(v6)
                            var v48 uint32 = 1
                            _ = v48
                        } else {
                            var v49 int16 = -661
                            _ = v49
                        }
                        if ((v30 * 1e10) == (-2.25 + float64(float32(float64(-2.25))))) {
                            if v39 {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        } else {
                            var v50 int32 = 689
                            v38 += 127
                            _ = v50
                        }
                    } else {
                        var v51 int8 = (v4 / (((v4 << ((v1) & 63))) | 1))
                        v52 := GenS7{F0: 617}
                        _ = v51
                        _ = v52
                    }
                    _ = len(v3)
                    _ = v46
                } else {
                    delete(v6, "a")
                    var v53 string = "a"
                    var v54 float64 = float64(3.14159)
                    _ = v53
                    _ = v54
                }
                if (v39 || ("hello" < "")) {
                    var v55 bool = (v39 && true)
                    if (v39 || false) {
                        v56 := GenS8{F0: 2147483647, F1: (v38 / ((v38) | 1))}
                        v57 := GenS2{F0: false, F1: 32767}
                        var v58 string = "a"
                        _ = v56
                        _ = v57
                        _ = v58
                    }
                    _ = v55
                } else {
                    var v59 bool = true
                    if true {
                        var v60 bool = true
                        var v61 string = "foo bar"
                        _ = v60
                        _ = v61
                    } else {
                        v37 = (^(v37))
                        v4 = v4
                        fmt.Println(fi40, v44, v4)
                    }
                    var v62 bool = (v38 >= v38)
                    _ = v59
                    _ = v62
                }
                _ = fi45
            }
        case "a":
            var v63 int = (fi40 + fi40)
            _ = v63
        }
        _ = fi40
        _ = v44
    }
    v64 := GenS7{F0: 4294967295}
    var v65 uint = 42
    if true {
        var v66 uint32 = 42
        for fi67 := 0; fi67 < 16; fi67++ {
            v65 = (^(v1))
            var v68 uint = 18446744073709551615
            var v69 float32 = 0.0
            _ = fi67
            _ = v68
            _ = v69
        }
        v70 := GenS3{F0: (v66 >> ((v1) & 63)), F1: v1, F2: ""}
        _ = v66
        _ = v70
    }
    if v39 {
        if (42 <= 42) {
            v71 := GenS5{F0: (v38 * 2), F1: genFn1(v39), F2: ("\n\t" + "a")}
            v72 := GenS9{F0: "hello"}
            var v73 uint32 = 479
            _ = v71
            _ = v72
            _ = v73
        } else {
            var v74 uint32 = 1
            var v75 int64 = 0
            _ = v74
            _ = v75
        }
    } else {
        switch -512 {
        case -2147483648:
            var v76 uint64 = 952
            _ = v76
        default:
            switch 460 {
            case 32767:
                var v77 uint32 = 2147483647
                _ = v77
            case 127:
                v78 := v6["\n\t"]
                _ = v78
            case 499:
                var v79 int32 = 1
                _ = v79
            }
        }
        var v80 uint8 = 127
        _ = v80
    }
    v39 = (v39 && (v39 && ((-(v4)) < v4)))
    for fi81 := 0; fi81 < 7; fi81++ {
        var v82 bool = true
        v65 -= (v37 * v65)
        _ = fi81
        _ = v82
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v65
}
