package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 int8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int; F1 uint16; F2 int }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int32; F1 int32; F2 int16 }
type GenS11 struct { F0 string; F1 uint8 }
func genFn0(p0 uint32, p1 uint, p2 int32) bool {
    if (828 < 2147483647) {
        for fi7 := 0; fi7 < 4; fi7++ {
            v8 := []int{fi7}
            p1 += (p1 / (((p1 & 42)) | 1))
            for fi9 := 0; fi9 < 12; fi9++ {
                switch (-(fi7)) {
                case 1000:
                    var v10 uint8 = 255
                    _ = v10
                case 1:
                    var v11 int16 = 1
                    _ = v11
                case -895:
                    p2 = (p2 & (-(p2)))
                }
                _ = fi9
            }
            _ = fi7
            _ = v8
        }
        for fi12 := 0; fi12 < 14; fi12++ {
            var v13 uint = p1
            var v14 int64 = 594
            _ = fi12
            _ = v13
            _ = v14
        }
    } else {
        v15 := GenS1{F0: 128}
        _ = v15
    }
    var v16 float32 = (-2.25 * 1.5)
    if true {
        v16 = float32(1.5)
        v17 := GenS1{F0: 1}
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return (!true)
}
func genFn1() int32 {
    v18 := GenS2{F0: "foo bar", F1: 274, F2: 228}
    var v19 uint8 = 8
    var v20 float64 = float64(0.1)
    v20 = v20
    if (v20 < v20) {
        v21 := []int{-9223372036854775808, -2147483648, 4294967295, 683}
        for fi22 := 0; fi22 < 20; fi22++ {
            for fi23 := 0; fi23 < 15; fi23++ {
                var v24 int = (fi23 + fi23)
                v25 := GenS3{F0: 2147483647}
                if len(v21) > 0 {
                    v21[0] = (fi23 / ((fi22) | 1))
                }
                _ = fi23
                _ = v24
                _ = v25
            }
            if true {
                if len(v21) >= 2 {
                    v21 = v21[1:len(v21)]
                }
            } else {
                if len(v21) >= 2 {
                    v21 = v21[1:len(v21)]
                }
                var v26 bool = (65535 >= 2147483647)
                _ = v26
            }
            _ = fi22
        }
        v21 = append(v21, -2147483648)
        _ = v21
    } else {
        v27 := []int{2147483647, 42, 354, 718}
        var v28 int = -785
        v29 := GenS4{F0: -1}
        _ = v27
        _ = v28
        _ = v29
    }
    _ = v18
    _ = v19
    _ = v20
    return -824
}

func main() {
    var v1 int64 = -878
    var v2 int64 = -1
    var v3 uint64 = 259
    v4 := make(map[string]uint32)
    v5 := make(map[string]uint32)
    var v6 string = ((("a" + "hello") + "") + ("foo bar" + "hello"))
    var v30 string = ""
    v30 = ""
    v6 = ((v30 + v30) + v30)
    fmt.Println(v2, v6)
    for fi31 := 0; fi31 < 8; fi31++ {
        v32 := GenS2{F0: (((v6 + (v30 + "\n\t")) + v6) + "hello"), F1: 0, F2: 1}
        delete(v5, (v30 + (v6 + v30)))
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 5; fi33++ {
        v30 = v6
        if ((1 >= 1) && false) {
            var v34 float64 = ((3.14159 * 0.1) / 1.0)
            v35 := []int{(fi33 % ((fi33) | 1)), (fi33 - (^(fi33))), -1, (fi33 / ((int(v3)) | 1))}
            v35 = append(v35, (fi33 | fi33))
            _ = v34
            _ = v35
        }
        fmt.Println(v3)
        _ = fi33
    }
    v36 := GenS5{F0: (v6 + "a"), F1: 122}
    if false {
        for fi37 := 0; fi37 < 8; fi37++ {
            _ = len(v5)
            for fi38 := 0; fi38 < 16; fi38++ {
                fmt.Println(fi38)
                if false {
                    v39 := GenS2{F0: "a", F1: 786, F2: 255}
                    _ = v39
                }
                if true {
                    var v40 int32 = 1
                    v41 := []int{fi38, (fi37 * fi37), (fi37 / ((fi37) | 1))}
                    _ = v40
                    _ = v41
                }
                _ = fi38
            }
            _ = fi37
        }
        for fi42 := 0; fi42 < 16; fi42++ {
            if false {
                var v43 int32 = 44
                v44 := []int{fi42, (fi42 >> ((795) & 63)), (fi42 >> ((881) & 63)), fi42}
                if (true || false) {
                    v45 := GenS6{F0: 100}
                    _ = v45
                } else {
                    for fi46 := 0; fi46 < 12; fi46++ {
                        var v47 uint32 = 2147483648
                        v48 := GenS7{F0: fi46}
                        if true {
                            fmt.Println(v6)
                            v49 := GenS8{F0: (-(fi46)), F1: 610, F2: fi46}
                            _ = v49
                        }
                        var v50 uint32 = 633
                        v50 -= v47
                        if true {
                            v51 := v4[""]
                            v52 := GenS7{F0: fi46}
                            if len(v44) >= 2 {
                                v44 = v44[1:len(v44)]
                            }
                            _ = v51
                            _ = v52
                        }
                        switch 60 {
                        case -931:
                            fmt.Println(v50, v30, fi42)
                        case 42:
                            var v53 int = fi46
                            _ = v53
                        default:
                            if len(v44) > 0 {
                                v44[0] = fi46
                            }
                        }
                        v30 = ((v6 + ((v30 + "") + v6)) + v6)
                        _ = fi46
                        _ = v47
                        _ = v48
                        _ = v50
                    }
                }
                _ = v43
                _ = v44
            } else {
                if (true && true) {
                    if (0 >= 0) {
                        for fi54 := 0; fi54 < 13; fi54++ {
                            v55 := v5[v6]
                            v56 := GenS9{F0: (true && true)}
                            switch v30 {
                            case "hello":
                                var v57 int16 = -378
                                _ = v57
                            default:
                                v58 := GenS1{F0: 109}
                                _ = v58
                            }
                            _ = fi54
                            _ = v55
                            _ = v56
                        }
                        for fi59 := 0; fi59 < 11; fi59++ {
                            if true {
                                v60 := GenS10{F0: 932, F1: genFn1(), F2: 0}
                                _ = v60
                            }
                            var v61 int16 = 577
                            fmt.Println(v1)
                            _ = fi59
                            _ = v61
                        }
                        var v62 int8 = 82
                        _ = v62
                    }
                    delete(v4, v6)
                }
                var v63 bool = true
                switch v30 {
                case "a":
                    fmt.Println(fi42)
                case "hello":
                    for fi64 := 0; fi64 < 15; fi64++ {
                        switch v6 {
                        case "":
                            delete(v5, "foo bar")
                        case "hello":
                            v65 := GenS4{F0: -128}
                            _ = v65
                        }
                        _ = fi64
                    }
                case "":
                    v2 -= v2
                }
                _ = v63
            }
            for fi66 := 0; fi66 < 8; fi66++ {
                if true {
                    v6 = ((v30 + v6) + ("" + v6))
                }
                if true {
                    v4["hello"] = 65535
                    var v67 uint64 = (v3 ^ (v3 + 0))
                    var v68 float32 = 0.1
                    if ((!((fi66 + fi42) != fi42)) || (v68 <= ((1e10 * (v68 + ((v68 + v68) * (v68 / 0.1)))) * v68))) {
                        v69 := GenS8{F0: fi42, F1: 42, F2: 4294967295}
                        if false {
                            var v70 float64 = ((0.1 + (0.0 * 0.0)) / ((-2.25 * (1.5 * -2.25)) - (float64(v68) + 3.14159)))
                            fmt.Println(fi42, v70)
                            var v71 int = fi42
                            _ = v70
                            _ = v71
                        }
                        _ = v69
                    }
                    var v72 float64 = float64(1.5)
                    _ = v67
                    _ = v68
                    _ = v72
                } else {
                    if (!false) {
                        v3 = v3
                    }
                }
                _ = fi66
            }
            v2 = v1
            _ = fi42
        }
    } else {
        for fi73 := 0; fi73 < 2; fi73++ {
            var v74 float64 = (float64(0.1) + -2.25)
            if (true || ((!true) || true)) {
                v75 := GenS9{F0: (-724 > genFn1())}
                if false {
                    v76 := GenS11{F0: v6, F1: 2}
                    v77 := GenS4{F0: 1}
                    _ = v76
                    _ = v77
                }
                for fi78 := 0; fi78 < 2; fi78++ {
                    fmt.Println(v6)
                    v79 := v5[((v30 + "\n\t") + v30)]
                    _ = fi78
                    _ = v79
                }
                _ = v75
            } else {
                switch v6 {
                case "a":
                    var v80 uint8 = 0
                    _ = v80
                case "\n\t":
                    if false {
                        v74 += v74
                        var v81 uint8 = 128
                        v82 := []int{}
                        _ = v81
                        _ = v82
                    }
                default:
                    fmt.Println(v74)
                }
                for fi83 := 0; fi83 < 20; fi83++ {
                    var v84 uint16 = 799
                    _ = fi83
                    _ = v84
                }
            }
            _ = fi73
            _ = v74
        }
        fmt.Println(v1, v2, v6)
        v30 = v6
    }
    for fi85 := 0; fi85 < 2; fi85++ {
        v86 := GenS9{F0: false}
        v30 = (v30 + "\n\t")
        _ = fi85
        _ = v86
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
