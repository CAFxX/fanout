package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 string }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32; F1 int64; F2 uint }
type GenS6 struct { F0 uint32; F1 uint8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 bool; F1 bool; F2 uint16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int; F1 uint }
type GenS11 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 int32; F1 bool }
type GenS14 struct { F0 uint16; F1 uint8; F2 uint8 }
func genFn0(p0 uint32, p1 uint8) uint8 {
    v6 := GenS1{F0: -328, F1: -428, F2: "a"}
    v7 := GenS2{F0: -2147483648, F1: (p1 << ((uint(p1)) & 63))}
    for fi8 := 0; fi8 < 8; fi8++ {
        if false {
            if (65535 > 42) {
                fmt.Println(p1)
                var v9 float32 = float32(0.1)
                v10 := GenS3{F0: 42, F1: false}
                _ = v10
                _ = v9
            }
            if (1 == 2147483647) {
                p0 = p0
            } else {
                v11 := GenS3{F0: (p1 - (p1 ^ (p1 + p1))), F1: false}
                v12 := []int{fi8, fi8}
                _ = v11
                _ = v12
            }
        }
        p1 += (p1 % ((p1) | 1))
        _ = fi8
    }
    var v13 uint16 = 799
    _ = p0
    _ = p1
    _ = v13
    _ = v6
    _ = v7
    return p1
}
func genFn1() uint16 {
    for fi14 := 0; fi14 < 13; fi14++ {
        fmt.Println(fi14)
        _ = fi14
    }
    var v15 int16 = 0
    for fi16 := 0; fi16 < 3; fi16++ {
        if (true || (42 < 2)) {
            for fi17 := 0; fi17 < 9; fi17++ {
                v15 = (v15 + v15)
                if (343 <= 237) {
                    fmt.Println(fi16)
                    var v18 string = "foo bar"
                    v19 := GenS4{F0: 42}
                    _ = v18
                    _ = v19
                }
                var v20 bool = (!(!false))
                _ = fi17
                _ = v20
            }
            if (1 != 4294967295) {
                v21 := []int{(fi16 / (((fi16 << ((1) & 63))) | 1))}
                var v22 int16 = 1000
                _ = v21
                _ = v22
            }
        } else {
            v15 = 42
        }
        v23 := GenS1{F0: 42, F1: (fi16 ^ fi16), F2: ("foo bar" + "foo bar")}
        for fi24 := 0; fi24 < 19; fi24++ {
            v25 := GenS3{F0: 0, F1: true}
            _ = fi24
            _ = v25
        }
        _ = fi16
        _ = v23
    }
    v26 := []int{-1, -9223372036854775808, 4294967295, -2147483648}
    v27 := GenS5{F0: 1, F1: 907, F2: 2}
    _ = v15
    _ = v26
    _ = v27
    return 184
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []int{467, 4294967295}
    v3 := make(map[int]int)
    v4 := []string{"a", ("a" + "a")}
    v5 := []string{}
    for fi28 := 0; fi28 < 2; fi28++ {
        v1 = (^(v1))
        for fi29 := 0; fi29 < 14; fi29++ {
            var v30 int16 = 32767
            _ = fi29
            _ = v30
        }
        for fi31 := 0; fi31 < 7; fi31++ {
            var v32 uint32 = 188
            v32 += (^(v32))
            var v33 uint8 = 2
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi28
    }
    v2 = append(v2, -2147483648)
    v34 := v3[32767]
    var v35 uint16 = 0
    v36 := GenS1{F0: 1, F1: v34, F2: ("\n\t" + "\n\t")}
    for fi37 := 0; fi37 < 16; fi37++ {
        if (65535 < 65535) {
            switch 0 {
            case -2147483648:
                if true {
                    v38 := GenS4{F0: 564}
                    var v39 uint8 = 0
                    var v40 uint8 = v39
                    _ = v38
                    _ = v39
                    _ = v40
                }
            case -1:
                if ((!false) || (296 < -793)) {
                    for fi41 := 0; fi41 < 20; fi41++ {
                        v42 := GenS1{F0: 262, F1: (v34 / ((fi41) | 1)), F2: "foo bar"}
                        _ = fi41
                        _ = v42
                    }
                    v43 := GenS6{F0: 42, F1: 128}
                    _ = v43
                } else {
                    v34 += -2147483648
                    if true {
                        var v44 uint = 65535
                        _ = v44
                    }
                }
            default:
                if len(v2) > 0 {
                    v2[0] = 1
                }
            }
            v45 := GenS7{F0: 128}
            _ = v45
        } else {
            if (((!false) || true) || (-1 <= 1000)) {
                var v46 uint16 = v35
                _ = len(v3)
                _ = v46
            } else {
                switch "\n\t" {
                case "hello":
                    for fi47 := 0; fi47 < 18; fi47++ {
                        v3[v34] = fi37
                        _ = fi47
                    }
                default:
                    fmt.Println(v1, v35)
                }
                var v48 float32 = (3.14159 * -2.25)
                if true {
                    var v49 string = ""
                    fmt.Println(v34, fi37)
                    switch v49 {
                    case "hello":
                        for fi50 := 0; fi50 < 2; fi50++ {
                            v51 := GenS8{F0: true, F1: true, F2: v35}
                            fmt.Println(v34)
                            if true {
                                var v52 uint32 = 353
                                v53 := GenS9{F0: 387}
                                v52 = (v52 >> ((1) & 63))
                                _ = v52
                                _ = v53
                            } else {
                                if ((-(v1)) >= v1) {
                                    v54 := GenS10{F0: fi37, F1: 0}
                                    v55 := GenS11{F0: 731, F1: 4294967295, F2: 345}
                                    _ = v54
                                    _ = v55
                                }
                            }
                            _ = fi50
                            _ = v51
                        }
                    case "":
                        v56 := GenS2{F0: 42, F1: 1}
                        _ = v56
                    case "a":
                        if ((v35 | (v35 >> ((2) & 63))) >= 65535) {
                            v34 = fi37
                            var v57 float64 = float64(v48)
                            _ = v57
                        }
                    default:
                        if (true && (true || true)) {
                            v58 := GenS3{F0: 82, F1: (4294967295 != 42)}
                            _ = v58
                        } else {
                            var v59 bool = false
                            _ = v59
                        }
                    }
                    _ = v49
                } else {
                    var v60 int64 = v1
                    var v61 uint = 2
                    v62 := GenS2{F0: 42, F1: 127}
                    _ = v60
                    _ = v61
                    _ = v62
                }
                _ = v48
            }
            var v63 int64 = v1
            v1 = v63
            _ = v63
        }
        v64 := GenS12{F0: true}
        _ = fi37
        _ = v64
    }
    if (103 < 100) {
        v35 -= v35
        if (v34 == 1) {
            fmt.Println(v35, v1, v1)
            if false {
                var v65 uint64 = 4294967295
                fmt.Println(v35)
                v1 += (v1 + (v1 & v1))
                _ = v65
            } else {
                for fi66 := 0; fi66 < 6; fi66++ {
                    if (1.5 >= (0.0 + 1.5)) {
                        v3[(v34 ^ 42)] = 682
                        fmt.Println(v1)
                        var v67 int16 = -1
                        _ = v67
                    } else {
                        var v68 int32 = 0
                        v3[v34] = v34
                        fmt.Println(v35, v68)
                        _ = v68
                    }
                    v69 := GenS8{F0: true, F1: ((v35 | v35) <= genFn1()), F2: 1}
                    if (689 != 1) {
                        if true {
                            v70 := GenS13{F0: -2147483648, F1: true}
                            var v71 uint32 = 2147483648
                            _ = v70
                            _ = v71
                        }
                        for fi72 := 0; fi72 < 5; fi72++ {
                            var v73 uint64 = 2
                            _ = len(v5)
                            _ = fi72
                            _ = v73
                        }
                    } else {
                        if ((!true) || false) {
                            if false {
                                for fi74 := 0; fi74 < 19; fi74++ {
                                    var v75 uint8 = 255
                                    _ = fi74
                                    _ = v75
                                }
                            } else {
                                v1 -= v1
                            }
                            for fi76 := 0; fi76 < 5; fi76++ {
                                delete(v3, int(v1))
                                v35 += (v35 << ((255) & 63))
                                if len(v4) > 0 {
                                    v4[0] = ("foo bar" + (("\n\t" + "a") + "hello"))
                                }
                                _ = fi76
                            }
                            for fi77 := 0; fi77 < 6; fi77++ {
                                v78 := GenS14{F0: 708, F1: 0, F2: 255}
                                v35 = 918
                                _ = fi77
                                _ = v78
                            }
                        } else {
                            v3[v34] = 0
                            for fi79 := 0; fi79 < 2; fi79++ {
                                v35 = (v35 * (v35 & 32767))
                                _ = fi79
                            }
                            var v80 float32 = ((-2.25 + 0.0) / 1.5)
                            _ = v80
                        }
                        var v81 uint32 = 2
                        if (float32((float64((-2.25 / 1.5)) / ((0.1 - (1e10 + 1.5)) / 1.5))) >= 1.5) {
                            for fi82 := 0; fi82 < 5; fi82++ {
                                if true {
                                    var v83 uint16 = 2
                                    v84 := GenS12{F0: (((v35 | genFn1()) >= (^(v35))) && true)}
                                    var v85 uint8 = 127
                                    _ = v83
                                    _ = v84
                                    _ = v85
                                }
                                _ = fi82
                            }
                            var v86 float32 = 0.1
                            for fi87 := 0; fi87 < 15; fi87++ {
                                v4 = append(v4, "")
                                if (false || (-2.25 <= v86)) {
                                    var v88 uint = 4294967295
                                    _ = v88
                                }
                                _ = fi87
                            }
                            _ = v86
                        } else {
                            fmt.Println(v1, v81)
                            _ = len(v4)
                            switch (^(v34)) {
                            case -399:
                                v34 = (v34 << ((0) & 63))
                            case -1:
                                var v89 float64 = 3.14159
                                _ = v89
                            case 32768:
                                for fi90 := 0; fi90 < 16; fi90++ {
                                    var v91 int32 = 1
                                    v35 = (v35 - 761)
                                    var v92 int64 = v1
                                    _ = fi90
                                    _ = v91
                                    _ = v92
                                }
                            }
                        }
                        _ = v81
                    }
                    _ = fi66
                    _ = v69
                }
                for fi93 := 0; fi93 < 19; fi93++ {
                    delete(v3, v34)
                    _ = fi93
                }
            }
            v1 -= 9223372036854775807
        }
        v94 := GenS14{F0: v35, F1: 1, F2: 51}
        _ = v94
    }
    var v95 uint16 = (v35 & genFn1())
    for fi96 := 0; fi96 < 4; fi96++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi96
    }
    for fi97 := 0; fi97 < 10; fi97++ {
        var v98 uint64 = 18446744073709551615
        if true {
            v95 -= v95
            var v99 bool = (true && true)
            switch ((("\n\t" + (("" + "a") + (("hello" + "hello") + (("a" + "a") + "hello")))) + "") + ((("" + "foo bar") + (("foo bar" + "a") + "\n\t")) + "foo bar")) {
            case "foo bar":
                v3[v34] = (fi97 ^ v34)
            default:
                delete(v3, (v34 & (v34 % ((653) | 1))))
            }
            _ = v99
        } else {
            var v100 uint32 = 919
            _ = v100
        }
        _ = fi97
        _ = v98
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v95
}
