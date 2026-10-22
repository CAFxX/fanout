package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint64; F2 string }
type GenS3 struct { F0 int8; F1 uint16; F2 bool }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint8; F1 int64; F2 uint }
type GenS9 struct { F0 uint64; F1 uint8 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    var v3 uint16 = 0
    var v4 int = 32767
    var v5 uint8 = 127
    v5 += v5
    v6 := GenS1{F0: 9223372036854775807, F1: 982, F2: -128}
    for fi7 := 0; fi7 < 15; fi7++ {
        _ = len(v2)
        fmt.Println(v5, v3, fi7)
        _ = fi7
    }
    v4 += -367
    v2[(v4 - (v4 << ((991) & 63)))] = 128
    var v8 int16 = 1
    v3 -= v3
    var v9 int = (^(v4))
    v5 += v5
    for fi10 := 0; fi10 < 11; fi10++ {
        if (true && false) {
            v8 = v8
        } else {
            v11 := []int{v9, (-(v4)), (v9 | (v9 ^ (v9 << ((552) & 63))))}
            v12 := v2[-2147483648]
            _ = v11
            _ = v12
        }
        var v13 float64 = (1.5 + (1.5 - 0.0))
        _ = fi10
        _ = v13
    }
    v8 = v8
    var v14 int8 = -20
    var v15 int16 = 560
    for fi16 := 0; fi16 < 20; fi16++ {
        for fi17 := 0; fi17 < 17; fi17++ {
            v9 += fi17
            var v18 bool = true
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 11; fi19++ {
            v3 = v3
            if true {
                var v20 float32 = (3.14159 - 1.5)
                fmt.Println(v14, v14, v15)
                if false {
                    if (v14 == 1) {
                        v9 = (v9 / (((v4 + (^(v4)))) | 1))
                        delete(v1, "hello")
                        var v21 uint8 = (^(v5))
                        v9 = (fi19 & fi19)
                        _ = v21
                    }
                    v3 = v3
                    if (((1e10 + 1.5) / 1.0) >= ((0.0 + 1e10) / 3.14159)) {
                        var v22 float64 = 0.1
                        _ = v22
                    } else {
                        if true {
                            var v23 uint8 = v5
                            _ = v23
                        } else {
                            v24 := GenS1{F0: -903, F1: 216, F2: 42}
                            var v25 int = (fi16 >> ((2) & 63))
                            _ = v24
                            _ = v25
                        }
                        if true {
                            if ((377 == 18446744073709551615) && true) {
                                if true {
                                    var v26 float64 = 1.5
                                    v27 := v1["foo bar"]
                                    switch "" {
                                    case "hello":
                                        delete(v1, ("\n\t" + "foo bar"))
                                    }
                                    switch ("foo bar" + "a") {
                                    case "":
                                        if true {
                                            fmt.Println(fi19, v20, v4)
                                            var v28 string = "hello"
                                            v29 := GenS1{F0: 62, F1: 2, F2: -6}
                                            _ = v28
                                            _ = v29
                                        }
                                    case "a":
                                        fmt.Println(v9)
                                    }
                                    _ = v26
                                    _ = v27
                                } else {
                                    switch "\n\t" {
                                    case "a":
                                        var v30 string = "foo bar"
                                        _ = v30
                                    }
                                    v31 := GenS1{F0: 42, F1: 1, F2: int8(v9)}
                                    var v32 float64 = 1e10
                                    v33 := GenS2{F0: 1, F1: 1, F2: "hello"}
                                    switch (fi16 ^ (^(v9))) {
                                    case 275:
                                        v34 := []int{-2147483648, v9, 127, fi16}
                                        _ = v34
                                    case 2147483647:
                                        if false {
                                            var v35 int32 = 2147483647
                                            v36 := []int{}
                                            _ = v35
                                            _ = v36
                                        } else {
                                            if false {
                                                fmt.Println(v32)
                                                var v37 uint8 = v5
                                                v38 := GenS3{F0: (^(v14)), F1: v3, F2: true}
                                                var v39 uint32 = 0
                                                _ = v37
                                                _ = v38
                                                _ = v39
                                            }
                                            var v40 int16 = v8
                                            v41 := []int{127, fi16, int(v15)}
                                            _ = v40
                                            _ = v41
                                        }
                                    case -724:
                                        v42 := GenS4{F0: (^(v14)), F1: int(v14)}
                                        _ = v42
                                    default:
                                        if (!((v5 << ((742) & 63)) > 128)) {
                                            v43 := GenS5{F0: 939, F1: -1}
                                            v44 := []int{}
                                            v2[int(v15)] = 525
                                            _ = v43
                                            _ = v44
                                        } else {
                                            var v45 uint16 = v3
                                            v9 = (fi16 >> ((255) & 63))
                                            _ = v45
                                        }
                                    }
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                }
                            } else {
                                var v46 float32 = v20
                                fmt.Println(v9)
                                fmt.Println(v5, v8)
                                _ = v46
                            }
                        } else {
                            var v47 float32 = v20
                            if (!false) {
                                var v48 float64 = -2.25
                                var v49 uint8 = (^(v5))
                                _ = v48
                                _ = v49
                            } else {
                                v50 := GenS6{F0: ("\n\t" + ("" + "foo bar")), F1: 2}
                                var v51 int16 = (v8 << ((1) & 63))
                                if (-562 < 679) {
                                    var v52 int8 = -21
                                    _ = v52
                                }
                                _ = v50
                                _ = v51
                            }
                            var v53 int = fi19
                            v15 = -674
                            var v54 bool = true
                            _ = v47
                            _ = v53
                            _ = v54
                        }
                    }
                    var v55 int64 = 706
                    switch ((("a" + (("" + "\n\t") + "")) + ("" + "a")) + "hello") {
                    case "":
                        v55 += -210
                    case "hello":
                        _ = len(v2)
                    case "\n\t":
                        var v56 uint8 = (v5 - (v5 >> ((193) & 63)))
                        _ = v56
                    }
                    var v57 uint64 = 18446744073709551615
                    if (-2.25 == ((v20 + (3.14159 + (float32(float64(0.1)) / (1e10 / v20)))) + -2.25)) {
                        v3 -= uint16(v5)
                        v9 = v4
                        v14 -= (v14 & 1)
                    } else {
                        v15 += v8
                        var v58 float32 = ((1e10 / ((v20 * (1e10 - 1.5)) - 0.1)) * (v20 + 3.14159))
                        v20 += ((v20 + v20) * float32(1.5))
                        _ = v58
                    }
                    _ = v55
                    _ = v57
                } else {
                    var v59 int8 = v14
                    if false {
                        v1["foo bar"] = 0
                        v2[int(v3)] = (v9 + fi19)
                        if (v8 >= v8) {
                            var v60 int16 = (v15 + 42)
                            v61 := []int{v9, (^(v9)), (fi19 ^ (fi19 << ((0) & 63)))}
                            _ = v60
                            _ = v61
                        } else {
                            switch "a" {
                            case "\n\t":
                                var v62 int32 = 42
                                _ = v62
                            default:
                                v63 := []int{(v4 << ((2) & 63)), fi16}
                                _ = v63
                            }
                            var v64 int8 = v59
                            v65 := GenS7{F0: v8}
                            _ = v64
                            _ = v65
                        }
                    } else {
                        var v66 uint16 = v3
                        var v67 int16 = (v8 >> ((323) & 63))
                        var v68 uint8 = uint8(v3)
                        _ = v66
                        _ = v67
                        _ = v68
                    }
                    if (!((1e10 == 1e10) && true)) {
                        var v69 uint8 = (v5 ^ 0)
                        _ = v69
                    } else {
                        fmt.Println(v15, v15)
                        v70 := []int{v4}
                        var v71 uint64 = 0
                        v2[(fi16 << ((2) & 63))] = 1
                        var v72 uint32 = 941
                        _ = v70
                        _ = v71
                        _ = v72
                    }
                    _ = v59
                }
                v73 := []int{(fi19 >> ((423) & 63)), v9}
                var v74 uint16 = 645
                _ = v20
                _ = v73
                _ = v74
            } else {
                var v75 uint16 = v3
                if false {
                    v76 := GenS7{F0: (v15 - -698)}
                    var v77 float32 = float32(1.5)
                    var v78 int = v9
                    _ = v76
                    _ = v77
                    _ = v78
                }
                var v79 int8 = 127
                _ = v75
                _ = v79
            }
            fmt.Println(v8, v15, v4)
            switch (v4 & (v9 % ((2147483647) | 1))) {
            case 128:
                v80 := []int{(fi16 * fi19), v9, 891}
                _ = v80
            }
            if true {
                switch 32767 {
                case 4294967295:
                    switch "a" {
                    case "\n\t":
                        var v81 uint8 = v5
                        _ = v81
                    case "a":
                        v82 := []int{}
                        _ = v82
                    }
                case 1:
                    v83 := GenS5{F0: 42, F1: 42}
                    _ = v83
                }
                if (!(v3 >= v3)) {
                    if false {
                        fmt.Println(v3)
                    }
                    if (2147483647 >= -1) {
                        var v84 float64 = 0.1
                        v85 := GenS8{F0: v5, F1: 42, F2: 65535}
                        var v86 float32 = -2.25
                        _ = v84
                        _ = v85
                        _ = v86
                    } else {
                        var v87 int8 = 100
                        v2[128] = -2147483648
                        var v88 int = fi16
                        _ = v87
                        _ = v88
                    }
                } else {
                    if (true && false) {
                        v89 := GenS2{F0: 42, F1: 4294967295, F2: "a"}
                        _ = v89
                    }
                    var v90 uint16 = v3
                    var v91 uint8 = (v5 / ((v5) | 1))
                    switch (fi16 | (v9 % ((v9) | 1))) {
                    case 21:
                        var v92 int16 = (v15 - -146)
                        _ = v92
                    }
                    v93 := []int{}
                    _ = v90
                    _ = v91
                    _ = v93
                }
            } else {
                v15 = v8
                fmt.Println(v4, fi16, fi16)
                v94 := GenS9{F0: 2, F1: v5}
                _ = v94
            }
            _ = fi19
        }
        _ = fi16
    }
    v15 -= 0
    fmt.Println(v8)
    v1[("hello" + (("\n\t" + "") + "hello"))] = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
