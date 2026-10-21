package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint64; F1 int32; F2 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16; F1 uint; F2 int }
type GenS6 struct { F0 int8; F1 int; F2 string }
type GenS7 struct { F0 uint16; F1 bool }
func genFn0(p0 int8, p1 int32) int {
    var v6 int8 = (p0 ^ (p0 * (p0 ^ (p0 * 127))))
    fmt.Println(p1, p1, v6)
    _ = p0
    _ = p1
    _ = v6
    return 0
}
func genFn1(p0 uint8, p1 uint64, p2 int32) uint {
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: 65535, F1: (p1 & 0)}
        _ = fi7
        _ = v8
    }
    var v9 float32 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 255
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    var v3 int32 = -252
    var v4 int32 = v3
    var v5 int = 9223372036854775807
    if true {
        if (false && true) {
            if false {
                var v10 uint64 = 1
                fmt.Println(v4, v4, v4)
                if false {
                    var v11 int8 = -38
                    _ = v11
                } else {
                    var v12 float64 = 3.14159
                    switch "hello" {
                    case "hello":
                        v4 += -2147483648
                    default:
                        if true {
                            v12 -= v12
                        }
                    }
                    _ = v12
                }
                _ = v10
            }
            if true {
                v13 := GenS1{F0: 42, F1: 437}
                v14 := GenS2{F0: 42, F1: 42, F2: 1}
                for fi15 := 0; fi15 < 18; fi15++ {
                    if (0.0 != (-2.25 / 1.0)) {
                        v16 := []int{(fi15 ^ (fi15 / ((-763) | 1)))}
                        var v17 int32 = v4
                        _ = v16
                        _ = v17
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            for fi18 := 0; fi18 < 12; fi18++ {
                                v5 = (fi18 - fi18)
                                if (false && ((false && false) && (147 >= 0))) {
                                    var v19 int16 = 42
                                    v3 = (v4 * 42)
                                    _ = v19
                                } else {
                                    if true {
                                        v20 := GenS3{F0: 42, F1: v3, F2: 42}
                                        var v21 uint32 = 4294967295
                                        _ = v20
                                        _ = v21
                                    } else {
                                        if false {
                                            var v22 int = (fi18 ^ (fi15 >> ((42) & 63)))
                                            fmt.Println(v5, fi15, v22)
                                            _ = v22
                                        } else {
                                            _ = len(v1)
                                            if false {
                                                fmt.Println(fi15)
                                            }
                                            fmt.Println(fi18, v3, fi15)
                                        }
                                        fmt.Println(fi18, fi18)
                                    }
                                    fmt.Println(fi15)
                                }
                                v4 = (v3 * -2147483648)
                                var v23 uint64 = 4294967295
                                var v24 int = v5
                                v4 = 1
                                _ = fi18
                                _ = v23
                                _ = v24
                            }
                        default:
                            switch (fi15 >> ((2) & 63)) {
                            case -838:
                                v3 = v3
                            case 32767:
                                for fi25 := 0; fi25 < 11; fi25++ {
                                    var v26 uint16 = 421
                                    var v27 uint = 42
                                    if true {
                                        v5 -= genFn0(-1, 0)
                                    } else {
                                        v28 := GenS4{F0: (v3 ^ 2147483647)}
                                        _ = v28
                                    }
                                    v29 := GenS5{F0: -614, F1: 622, F2: fi15}
                                    v3 += (v4 / ((-2147483648) | 1))
                                    var v30 int = genFn0(127, 1)
                                    v31 := GenS6{F0: 1, F1: (v30 & genFn0(13, 2147483647)), F2: "foo bar"}
                                    var v32 string = "a"
                                    switch (fi25 % (((^(fi15))) | 1)) {
                                    case 128:
                                        switch v5 {
                                        case 32768:
                                            switch v32 {
                                            case "":
                                                v30 -= v30
                                            }
                                        }
                                        var v33 uint8 = 197
                                        v32 = ((((v32 + "") + ((((("\n\t" + ("foo bar" + "")) + v32) + "a") + "hello") + ("a" + ((v32 + (v32 + v32)) + (v32 + (("hello" + "\n\t") + (("foo bar" + ((v32 + "\n\t") + "\n\t")) + "foo bar"))))))) + ("hello" + (v32 + v32))) + v32)
                                        v3 -= v4
                                        var v34 int16 = 42
                                        _ = v33
                                        _ = v34
                                    default:
                                        if (42 <= 42) {
                                            if false {
                                                var v35 int16 = 1000
                                                _ = v35
                                            } else {
                                                v30 += 336
                                                fmt.Println(v26)
                                            }
                                        }
                                    }
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                }
                            case -634:
                                if (127 > 100) {
                                    var v36 uint = 255
                                    var v37 int32 = v4
                                    v38 := GenS1{F0: (v36 - 148), F1: 1}
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                }
                            }
                        }
                    }
                    _ = fi15
                }
                _ = v13
                _ = v14
            } else {
                var v39 uint8 = 255
                var v40 uint8 = 27
                _ = v39
                _ = v40
            }
            var v41 float64 = ((((0.1 - (-2.25 + 3.14159)) + float64(3.14159)) * -2.25) * -2.25)
            _ = v41
        } else {
            v4 -= v3
        }
        switch (v5 + 0) {
        case 348:
            fmt.Println(v5, v5)
        case 65:
            v42 := GenS6{F0: 0, F1: -2147483648, F2: ("\n\t" + ("a" + "foo bar"))}
            _ = v42
        default:
            var v43 uint32 = 2
            _ = v43
        }
    } else {
        var v44 int16 = 32767
        var v45 int16 = (v44 + (v44 / ((-908) | 1)))
        _ = v44
        _ = v45
    }
    if ((true || true) || (!(!(false || (0 != 2147483647))))) {
        fmt.Println(v4)
    } else {
        var v46 uint = genFn1(1, 2, (v4 * v4))
        _ = v46
    }
    fmt.Println(v5, v4)
    if (true && ((((float64((float32(1e10) * -2.25)) + 1e10) * 0.0) != (((((0.1 + -2.25) + (1.5 + (3.14159 + (0.0 - 0.1)))) * 1.5) * -2.25) / ((float64(0.0) * 1e10) + -2.25))) || false)) {
        if true {
            fmt.Println(v4, v4, v3)
            _ = len(v2)
            var v47 uint32 = 65535
            _ = v47
        }
        v48 := []int{(v5 ^ int(v3))}
        _ = v48
    }
    v4 = v3
    for fi49 := 0; fi49 < 12; fi49++ {
        for fi50 := 0; fi50 < 6; fi50++ {
            var v51 int32 = v3
            fmt.Println(fi49, fi50)
            if true {
                v5 -= 0
            } else {
                v52 := []int{fi49, fi49, (fi50 << ((42) & 63)), fi49}
                var v53 uint = 1
                var v54 int32 = (v51 + v51)
                v1[(v5 % ((32767) | 1))] = 742
                _ = v52
                _ = v53
                _ = v54
            }
            var v55 string = (("foo bar" + "a") + "hello")
            if false {
                v1[v5] = 2147483648
                switch v55 {
                case "hello":
                    v3 = 42
                case "\n\t":
                    v56 := GenS4{F0: 42}
                    _ = v56
                case "":
                    if false {
                        var v57 uint64 = 192
                        _ = v57
                    } else {
                        switch fi50 {
                        case 0:
                            switch v55 {
                            case "hello":
                                v3 += -1
                            default:
                                if true {
                                    v58 := v1[(fi49 * fi50)]
                                    fmt.Println(v5, v55, v58)
                                    _ = v58
                                }
                            }
                        case 9223372036854775807:
                            var v59 string = (v55 + "a")
                            fmt.Println(v59)
                            _ = v59
                        }
                        var v60 int8 = -63
                        var v61 uint16 = 42
                        _ = v60
                        _ = v61
                    }
                default:
                    v3 = v51
                }
                var v62 uint16 = 0
                _ = v62
            }
            if false {
                switch 127 {
                case 127:
                    var v63 string = v55
                    _ = v63
                case 2147483647:
                    v64 := []int{(v5 * (v5 + fi50)), fi50}
                    _ = v64
                case -925:
                    v65 := GenS7{F0: 42, F1: true}
                    _ = v65
                }
                fmt.Println(fi50, v55, v51)
            } else {
                v5 = (-(fi49))
                v5 -= -855
                if ((v55 + v55) != v55) {
                    switch "\n\t" {
                    case "a":
                        var v66 float64 = (3.14159 / 1.5)
                        _ = v66
                    }
                    _ = len(v1)
                }
            }
            var v67 float64 = 1e10
            _ = fi50
            _ = v51
            _ = v55
            _ = v67
        }
        v68 := GenS6{F0: -81, F1: v5, F2: "a"}
        _ = fi49
        _ = v68
    }
    if false {
        v69 := v1[v5]
        _ = v69
    }
    for fi70 := 0; fi70 < 4; fi70++ {
        switch genFn0(-1, v3) {
        case -1:
            for fi71 := 0; fi71 < 4; fi71++ {
                if true {
                    switch "" {
                    case "hello":
                        var v72 int64 = -9223372036854775808
                        _ = v72
                    case "\n\t":
                        v5 = (fi71 - (fi70 + (fi70 - v5)))
                    }
                    var v73 int16 = 1
                    _ = v73
                } else {
                    var v74 uint = 4294967295
                    fmt.Println(v4, fi70)
                    _ = v74
                }
                v5 = v5
                fmt.Println(v4, fi70, fi71)
                _ = fi71
            }
        default:
            var v75 uint32 = 377
            _ = v75
        }
        var v76 int32 = (^(v4))
        v77 := []int{(v5 | (fi70 / (((v5 ^ fi70)) | 1))), (v5 + genFn0(42, -1)), (v5 | 0), fi70}
        _ = fi70
        _ = v76
        _ = v77
    }
    if (857 <= 1) {
        var v78 uint64 = 1
        if false {
            var v79 bool = true
            _ = v79
        } else {
            v78 -= v78
            if (127 == -1) {
                v80 := []int{9223372036854775807}
                if true {
                    fmt.Println(v5, v5, v4)
                    v5 = v5
                }
                v80 = append(v80, (-(v5)))
                _ = v80
            }
            fmt.Println(v5)
        }
        _ = v78
    }
    switch (("hello" + "a") + "hello") {
    case "foo bar":
        v4 = int32(v5)
    case "\n\t":
        switch "hello" {
        case "a":
            if (1 >= 1) {
                var v81 uint32 = 2147483647
                _ = v81
            } else {
                v3 = v4
            }
        default:
            fmt.Println(v4, v3)
        }
    default:
        v82 := GenS7{F0: 255, F1: (((float32((float64(-2.25) + float64(1.5))) * float32(0.0)) * (((1e10 * -2.25) + 1.5) * 0.1)) >= float32(3.14159))}
        _ = v82
    }
    v3 += -288
    fmt.Println(v3, v4, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
