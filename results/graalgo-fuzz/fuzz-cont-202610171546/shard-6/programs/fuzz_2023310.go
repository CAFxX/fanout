package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 int64; F2 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS6 struct { F0 int32; F1 uint16 }
type GenS7 struct { F0 int32; F1 uint32; F2 uint8 }
func genFn0(p0 uint8, p1 int64) uint64 {
    v7 := GenS1{F0: 973, F1: 715, F2: 213}
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 2; fi9++ {
            switch fi8 {
            case 32767:
                if true {
                    p0 = (p0 & p0)
                    if (1 >= 0) {
                        switch "" {
                        case "":
                            var v10 int = fi8
                            _ = v10
                        case "a":
                            if (false || ("hello" != "foo bar")) {
                                var v11 uint16 = 2
                                if (19 == 42) {
                                    v12 := GenS2{F0: true}
                                    fmt.Println(p0, p1)
                                    if true {
                                        v13 := GenS3{F0: (p0 - 128), F1: p1, F2: 896}
                                        var v14 int32 = 386
                                        _ = v13
                                        _ = v14
                                    } else {
                                        var v15 float32 = ((3.14159 / 0.1) / 1.0)
                                        p0 = uint8(v11)
                                        fmt.Println(fi8, v11, fi9)
                                        _ = v15
                                    }
                                    _ = v12
                                }
                                v11 = 255
                                _ = v11
                            } else {
                                if false {
                                    p1 = 111
                                }
                                var v16 string = (("hello" + ("a" + "")) + "hello")
                                if false {
                                    v17 := GenS4{F0: 18446744073709551615}
                                    v18 := GenS4{F0: 1}
                                    fmt.Println(fi9)
                                    if ((fi9 / ((fi9) | 1)) > fi8) {
                                        fmt.Println(fi8)
                                        v19 := GenS2{F0: false}
                                        switch 1 {
                                        case 127:
                                            var v20 uint64 = 1
                                            _ = v20
                                        case 0:
                                            if (-519 < 42) {
                                                p1 += p1
                                            } else {
                                                v21 := GenS5{F0: 177, F1: 4294967295, F2: 128}
                                                _ = v21
                                            }
                                        }
                                        _ = v19
                                    }
                                    _ = v17
                                    _ = v18
                                }
                                var v22 int16 = -498
                                _ = v16
                                _ = v22
                            }
                            fmt.Println(p1)
                            switch "a" {
                            case "a":
                                v23 := GenS6{F0: 2147483647, F1: 952}
                                _ = v23
                            case "hello":
                                p0 = 255
                            }
                        }
                        v24 := GenS6{F0: 2147483647, F1: 42}
                        _ = v24
                    }
                    var v25 int8 = int8(fi9)
                    _ = v25
                } else {
                    fmt.Println(p0, fi9, fi8)
                }
            case 1:
                v26 := GenS2{F0: false}
                _ = v26
            }
            v27 := GenS2{F0: (32767 == 32767)}
            _ = fi9
            _ = v27
        }
        switch (fi8 + (fi8 | (fi8 * 128))) {
        case 1:
            var v28 float32 = float32(1.5)
            _ = v28
        case 775:
            if (0 <= 301) {
                if (780 > 0) {
                    switch 42 {
                    case -321:
                        if (!true) {
                            switch "" {
                            case "foo bar":
                                switch 32767 {
                                case -1:
                                    var v29 int32 = -63
                                    _ = v29
                                }
                            default:
                                for fi30 := 0; fi30 < 10; fi30++ {
                                    fmt.Println(p0)
                                    _ = fi30
                                }
                            }
                            fmt.Println(p0, fi8)
                            var v31 uint16 = 1
                            _ = v31
                        }
                    case -2147483648:
                        switch -566 {
                        case -97:
                            switch ("a" + "a") {
                            case "":
                                p1 = (p1 + p1)
                            default:
                                var v32 uint16 = 32767
                                _ = v32
                            }
                        case 142:
                            var v33 int16 = 0
                            _ = v33
                        case -2147483648:
                            if false {
                                for fi34 := 0; fi34 < 15; fi34++ {
                                    var v35 uint32 = 2147483648
                                    _ = fi34
                                    _ = v35
                                }
                                p0 += 92
                            }
                        default:
                            p1 -= (p1 * (p1 | -718))
                        }
                    case 4294967295:
                        p1 = p1
                    default:
                        var v36 int16 = -147
                        _ = v36
                    }
                    v37 := []int{fi8, (fi8 & fi8)}
                    _ = v37
                } else {
                    p1 -= (p1 << ((1) & 63))
                    p0 = p0
                }
            } else {
                fmt.Println(p0, p0, p1)
                v38 := GenS4{F0: 289}
                _ = v38
            }
        case -1:
            var v39 uint8 = 42
            _ = v39
        default:
            var v40 float32 = 3.14159
            _ = v40
        }
        switch "a" {
        case "foo bar":
            for fi41 := 0; fi41 < 16; fi41++ {
                v42 := []int{-1}
                _ = fi41
                _ = v42
            }
        case "a":
            v43 := []int{fi8, -288, (fi8 % ((-564) | 1)), fi8}
            _ = v43
        default:
            v44 := GenS4{F0: 65535}
            _ = v44
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v7
    return 18446744073709551615
}
func genFn1() int32 {
    var v45 float32 = (3.14159 * 1.5)
    fmt.Println(v45)
    v45 -= 0.1
    var v46 uint = 345
    if false {
        v47 := []int{}
        for fi48 := 0; fi48 < 3; fi48++ {
            v45 -= v45
            v45 = ((1.5 - v45) + float32(0.0))
            v46 = v46
            _ = fi48
        }
        v46 = v46
        _ = v47
    }
    _ = v45
    _ = v46
    return 1
}

func main() {
    var v1 bool = (0.1 <= (0.1 / (3.14159 / -2.25)))
    v2 := make(map[uint32]string)
    v3 := make(map[string]int)
    var v4 uint8 = 1
    v5 := make(map[uint32]string)
    var v6 uint16 = 0
    switch -2147483648 {
    case -567:
        var v49 uint32 = 2147483648
        _ = v49
    default:
        for fi50 := 0; fi50 < 19; fi50++ {
            var v51 uint8 = (v4 - (v4 & 128))
            if (!((-2147483648 == 2147483647) && v1)) {
                for fi52 := 0; fi52 < 15; fi52++ {
                    v3[("\n\t" + "")] = fi50
                    delete(v2, 4294967295)
                    _ = fi52
                }
                v53 := []int{}
                v51 = v51
                _ = v53
            } else {
                v6 = v6
                _ = len(v2)
                if v1 {
                    if v1 {
                        for fi54 := 0; fi54 < 16; fi54++ {
                            v51 += v51
                            _ = fi54
                        }
                    }
                    var v55 int8 = 100
                    delete(v2, 1)
                    _ = v55
                }
            }
            switch "a" {
            case "foo bar":
                v5[uint32(v4)] = "\n\t"
            case "":
                v51 -= v51
            default:
                v56 := GenS2{F0: (42 != 42)}
                _ = v56
            }
            _ = fi50
            _ = v51
        }
    }
    switch "foo bar" {
    case "foo bar":
        for fi57 := 0; fi57 < 11; fi57++ {
            v58 := []int{(fi57 * 2147483647)}
            var v59 uint = 1
            v60 := GenS3{F0: v4, F1: 886, F2: 700}
            _ = fi57
            _ = v58
            _ = v59
            _ = v60
        }
    case "a":
        if ((0 > 4294967295) && (-309 < 32767)) {
            if (!(!true)) {
                var v61 int8 = -1
                _ = v61
            } else {
                var v62 uint = 18446744073709551615
                v63 := GenS1{F0: v6, F1: v6, F2: 2}
                _ = v62
                _ = v63
            }
            v64 := []int{615, 1}
            v6 += (^(v6))
            _ = v64
        }
    case "\n\t":
        if true {
            v6 = 643
            v6 = 516
            switch -748 {
            case 177:
                v65 := GenS4{F0: 439}
                _ = v65
            default:
                v1 = v1
            }
        } else {
            v66 := GenS6{F0: genFn1(), F1: v6}
            v3[""] = 0
            switch ("a" + ("foo bar" + "a")) {
            case "\n\t":
                var v67 string = ("foo bar" + "foo bar")
                _ = v67
            case "a":
                v68 := v5[216]
                _ = v68
            case "foo bar":
                v69 := GenS2{F0: true}
                _ = v69
            }
            _ = v66
        }
    }
    v70 := v3["a"]
    v1 = v1
    v6 += 21
    v4 -= v4
    var v71 uint = uint(v6)
    switch "\n\t" {
    case "":
        v4 = 2
    default:
        v72 := GenS7{F0: genFn1(), F1: 2, F2: v4}
        _ = v72
    }
    for fi73 := 0; fi73 < 19; fi73++ {
        if true {
            v1 = (2147483647 == -2147483648)
            switch "a" {
            case "a":
                var v74 uint = v71
                _ = v74
            }
        } else {
            v75 := GenS4{F0: 18446744073709551615}
            switch 2147483647 {
            case 2147483647:
                if ((v71 % (((v71 ^ 0)) | 1)) > 2) {
                    fmt.Println(v1, v1)
                    for fi76 := 0; fi76 < 6; fi76++ {
                        var v77 float32 = (1e10 / 1e10)
                        _ = fi76
                        _ = v77
                    }
                }
            case 307:
                var v78 int16 = int16(v6)
                _ = v78
            case 128:
                switch 320 {
                case 529:
                    v79 := v2[4294967295]
                    _ = v79
                }
            default:
                v80 := GenS7{F0: 0, F1: 323, F2: 247}
                _ = v80
            }
            if (!(!false)) {
                v81 := []int{v70, (v70 - (fi73 + fi73)), v70, (fi73 / ((v70) | 1))}
                _ = v81
            } else {
                fmt.Println(v6, v1, v71)
                v82 := []int{v70, 0}
                v83 := GenS4{F0: (v71 ^ (v71 >> ((v71) & 63)))}
                _ = v82
                _ = v83
            }
            _ = v75
        }
        var v84 uint8 = v4
        _ = fi73
        _ = v84
    }
    var v85 float64 = 0.0
    var v86 float64 = (v85 * (float64(1e10) + v85))
    var v87 string = "a"
    if false {
        var v88 int64 = 0
        v89 := []int{v70}
        _ = v88
        _ = v89
    }
    fmt.Println(v86)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v70
    _ = v71
    _ = v85
    _ = v86
    _ = v87
}
