package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 int32; F1 int; F2 uint8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int8; F1 int8; F2 int16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint16; F1 bool }
func genFn0(p0 int8) uint {
    if true {
        var v6 uint = 18446744073709551615
        _ = v6
    } else {
        var v7 int = 127
        v8 := GenS1{F0: 1, F1: "foo bar"}
        v7 = 2147483647
        _ = v7
        _ = v8
    }
    switch "\n\t" {
    case "":
        switch ("a" + (("\n\t" + "hello") + "a")) {
        case "\n\t":
            p0 += (p0 >> ((42) & 63))
        }
    case "foo bar":
        v9 := GenS2{F0: 2147483647}
        _ = v9
    }
    switch "hello" {
    case "a":
        var v10 uint32 = 138
        _ = v10
    case "hello":
        v11 := []int{}
        _ = v11
    case "":
        if true {
            for fi12 := 0; fi12 < 19; fi12++ {
                for fi13 := 0; fi13 < 9; fi13++ {
                    p0 = p0
                    if ((0 < 0) || (((1 != 21) && false) && false)) {
                        fmt.Println(fi13, fi12)
                    } else {
                        switch ("a" + ("a" + "")) {
                        case "a":
                            v14 := GenS1{F0: -9223372036854775808, F1: "a"}
                            if (231 != 127) {
                                v15 := GenS3{F0: 0, F1: (-(fi13))}
                                p0 = p0
                                v16 := GenS1{F0: -1, F1: ("foo bar" + ("\n\t" + "hello"))}
                                _ = v15
                                _ = v16
                            }
                            _ = v14
                        case "hello":
                            fmt.Println(p0)
                        case "":
                            v17 := GenS3{F0: 444, F1: (fi13 - (fi12 ^ -9223372036854775808))}
                            _ = v17
                        default:
                            v18 := []int{(fi12 + fi12), fi12, 771, fi13}
                            _ = v18
                        }
                    }
                    _ = fi13
                }
                if false {
                    if (!false) {
                        p0 = (p0 % ((p0) | 1))
                        p0 = (p0 % ((1) | 1))
                    }
                    for fi19 := 0; fi19 < 2; fi19++ {
                        if false {
                            var v20 uint32 = 600
                            fmt.Println(fi19, v20)
                            v21 := GenS2{F0: 0}
                            _ = v20
                            _ = v21
                        } else {
                            v22 := []int{(fi12 / (((^(fi12))) | 1)), (fi19 >> ((18446744073709551615) & 63))}
                            v23 := GenS4{F0: -186}
                            _ = v22
                            _ = v23
                        }
                        v24 := []int{fi12, fi19, fi12}
                        _ = fi19
                        _ = v24
                    }
                } else {
                    v25 := GenS2{F0: 1}
                    if true {
                        v26 := GenS5{F0: (0 <= 255), F1: 736}
                        fmt.Println(p0)
                        _ = v26
                    } else {
                        p0 = (p0 + p0)
                        for fi27 := 0; fi27 < 16; fi27++ {
                            var v28 int = (fi12 ^ (fi12 & (^(fi12))))
                            _ = fi27
                            _ = v28
                        }
                        fmt.Println(fi12, fi12)
                    }
                    _ = v25
                }
                _ = fi12
            }
            v29 := GenS5{F0: (1 == 201), F1: -209}
            _ = v29
        }
    default:
        p0 = 1
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        for fi31 := 0; fi31 < 11; fi31++ {
            if (("foo bar" + "") < ("foo bar" + "foo bar")) {
                fmt.Println(p0, fi31, fi31)
                v32 := []int{fi31, (fi30 % ((fi31) | 1))}
                p0 -= 42
                _ = v32
            }
            v33 := []int{(fi30 * fi31), fi31}
            _ = fi31
            _ = v33
        }
        if ((0.0 * -2.25) <= (0.0 - ((((((0.1 * 0.1) * 3.14159) - 0.0) + float64(1e10)) + float64(1e10)) - ((1e10 * ((0.0 * 0.1) * 1e10)) * 1.5)))) {
            switch "foo bar" {
            case "":
                for fi34 := 0; fi34 < 19; fi34++ {
                    var v35 int32 = -2147483648
                    _ = fi34
                    _ = v35
                }
            case "hello":
                var v36 int16 = 307
                _ = v36
            }
            v37 := GenS6{F0: 42, F1: (fi30 % (((^(fi30))) | 1)), F2: 42}
            _ = v37
        } else {
            if false {
                v38 := GenS7{F0: -1}
                p0 = 1
                p0 = (p0 * -1)
                _ = v38
            }
            if (!((4294967295 >= 18446744073709551615) || (1 >= -2147483648))) {
                p0 = p0
                v39 := []int{fi30}
                _ = v39
            } else {
                v40 := []int{(-(fi30)), fi30}
                for fi41 := 0; fi41 < 11; fi41++ {
                    if len(v40) > 0 {
                        v40[0] = fi41
                    }
                    var v42 int64 = 1
                    _ = fi41
                    _ = v42
                }
                _ = v40
            }
            for fi43 := 0; fi43 < 9; fi43++ {
                var v44 uint16 = 861
                _ = fi43
                _ = v44
            }
        }
        _ = fi30
    }
    p0 = p0
    _ = p0
    return 65535
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = -145
    var v3 int64 = 42
    var v4 uint = 450
    var v5 int16 = -1
    v2 = (-(v2))
    var v45 uint32 = 320
    switch 0 {
    case 1:
        var v46 int16 = (v5 << ((2) & 63))
        _ = v46
    case 0:
        var v47 float64 = 0.0
        _ = v47
    default:
        switch -2147483648 {
        case -602:
            if ("a" <= "") {
                if ((!false) || (0.0 <= float64(0.0))) {
                    v3 = v3
                    v48 := GenS7{F0: v2}
                    var v49 int16 = 920
                    _ = v48
                    _ = v49
                }
            } else {
                var v50 uint16 = 65535
                fmt.Println(v2, v5, v50)
                var v51 int = 127
                _ = v50
                _ = v51
            }
        default:
            fmt.Println(v45, v4)
        }
    }
    var v52 uint = (v4 ^ (v4 | (v4 + (v4 >> (((v4 & 0)) & 63)))))
    var v53 int16 = v5
    v54 := GenS8{F0: 0, F1: -1, F2: v5}
    v53 = (v53 << (((v52 / (((v4 ^ genFn0(int8(v3)))) | 1))) & 63))
    for fi55 := 0; fi55 < 2; fi55++ {
        v5 = 895
        _ = fi55
    }
    if false {
        v56 := GenS9{F0: 524}
        _ = v56
    }
    var v57 uint32 = v45
    for fi58 := 0; fi58 < 5; fi58++ {
        for fi59 := 0; fi59 < 9; fi59++ {
            v60 := []int{(fi58 << ((v52) & 63)), int(v53)}
            if (true && ((!(4294967295 > 11)) && true)) {
                switch (fi59 - (fi59 | (^(fi59)))) {
                case 4294967295:
                    if true {
                        var v61 uint16 = 0
                        _ = v61
                    }
                }
                var v62 float32 = 1.5
                _ = v62
            } else {
                delete(v1, "hello")
                if (!false) {
                    var v63 int64 = v3
                    var v64 int16 = (v53 * v5)
                    var v65 string = ("\n\t" + "a")
                    _ = v63
                    _ = v64
                    _ = v65
                }
                delete(v1, "foo bar")
            }
            if (348 == v57) {
                v66 := GenS2{F0: (v2 >> ((v4) & 63))}
                if len(v60) > 0 {
                    v60[0] = (fi58 >> ((genFn0(100)) & 63))
                }
                if false {
                    var v67 int64 = (v3 ^ v3)
                    switch (("a" + "a") + "hello") {
                    case "\n\t":
                        switch fi58 {
                        case -1:
                            var v68 int32 = (-(v2))
                            _ = v68
                        case 127:
                            delete(v1, "a")
                        default:
                            if (("foo bar" + "a") > ("\n\t" + "a")) {
                                fmt.Println(v57, v2, v67)
                                var v69 int16 = (v5 & -827)
                                if false {
                                    if len(v60) >= 2 {
                                        v60 = v60[1:len(v60)]
                                    }
                                    v70 := GenS10{F0: 444, F1: ((v2 ^ -1) <= (^(v2)))}
                                    var v71 int8 = -11
                                    _ = v70
                                    _ = v71
                                }
                                _ = v69
                            } else {
                                var v72 int16 = v53
                                v73 := GenS4{F0: v2}
                                v52 += (v4 | genFn0(-128))
                                fmt.Println(v4, v52)
                                v57 = 2147483648
                                v3 += (v3 | (v3 * v3))
                                _ = v72
                                _ = v73
                            }
                            if ((!(true || false)) && (!false)) {
                                if false {
                                    if ((fi59 / ((fi59) | 1)) < fi59) {
                                        var v74 int = (fi59 >> ((196) & 63))
                                        _ = v74
                                    } else {
                                        if len(v60) > 0 {
                                            v60[0] = fi59
                                        }
                                        v1["a"] = "\n\t"
                                    }
                                    v57 += v57
                                    switch fi58 {
                                    case -373:
                                        var v75 uint32 = v57
                                        _ = v75
                                    case -470:
                                        var v76 uint32 = v57
                                        _ = v76
                                    default:
                                        var v77 string = "a"
                                        _ = v77
                                    }
                                }
                                var v78 int32 = v2
                                v4 = (v52 % (((v4 | (v52 + v52))) | 1))
                                _ = v78
                            }
                            fmt.Println(v52)
                            if true {
                                fmt.Println(v2, v57)
                            } else {
                                if len(v60) >= 2 {
                                    v60 = v60[1:len(v60)]
                                }
                            }
                            v45 += v57
                            var v79 uint64 = 2
                            var v80 string = "hello"
                            fmt.Println(v45, v4, fi58)
                            _ = v79
                            _ = v80
                        }
                    case "foo bar":
                        v45 = v57
                    case "a":
                        if false {
                            _ = len(v60)
                        }
                    }
                    _ = v67
                }
                _ = v66
            }
            _ = fi59
            _ = v60
        }
        _ = fi58
    }
    var v81 uint32 = (v57 | v57)
    v82 := GenS3{F0: 2, F1: 299}
    var v83 float32 = float32(1e10)
    var v84 uint8 = 121
    for fi85 := 0; fi85 < 2; fi85++ {
        var v86 string = "hello"
        v52 = (v4 * uint(v81))
        _ = fi85
        _ = v86
    }
    v87 := v1[("hello" + (("\n\t" + "hello") + "\n\t"))]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v57:", v57)
    fmt.Println("v81:", v81)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v57
    _ = v81
    _ = v82
    _ = v83
    _ = v84
    _ = v87
}
