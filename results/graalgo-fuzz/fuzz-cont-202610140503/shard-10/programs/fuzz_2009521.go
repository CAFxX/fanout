package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int64 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16; F1 int16 }
type GenS8 struct { F0 uint32; F1 uint16; F2 uint }
type GenS9 struct { F0 uint; F1 int }
type GenS10 struct { F0 string; F1 uint64 }
func genFn0(p0 uint) int {
    var v6 uint64 = 1
    var v7 uint8 = 255
    if (!(866 <= 42)) {
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := []int{}
            var v10 string = ("hello" + "foo bar")
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 uint = p0
        if (!true) {
            v6 += 0
            for fi12 := 0; fi12 < 15; fi12++ {
                for fi13 := 0; fi13 < 6; fi13++ {
                    v11 -= v11
                    if false {
                        p0 = (v11 / ((65535) | 1))
                    }
                    _ = fi13
                }
                switch (("hello" + "a") + "hello") {
                case "":
                    if true {
                        var v14 int = fi12
                        var v15 uint16 = 836
                        _ = v14
                        _ = v15
                    } else {
                        var v16 int16 = 992
                        if true {
                            for fi17 := 0; fi17 < 5; fi17++ {
                                switch int(v6) {
                                case 0:
                                    v7 += (v7 - v7)
                                }
                                v18 := []int{fi17}
                                switch (fi17 + 128) {
                                case 0:
                                    var v19 uint8 = (v7 - (v7 % (((^(v7))) | 1)))
                                    _ = v19
                                case -2147483648:
                                    v7 += (v7 - v7)
                                    if false {
                                        switch ("foo bar" + "hello") {
                                        case "":
                                            v20 := GenS1{F0: 65535, F1: -671, F2: -1}
                                            _ = v20
                                        case "a":
                                            if ((42 <= v6) || true) {
                                                if true {
                                                    p0 = uint(v7)
                                                    var v21 bool = false
                                                    var v22 uint8 = (^(v7))
                                                    _ = v21
                                                    _ = v22
                                                } else {
                                                    v11 = v11
                                                    p0 += 255
                                                    var v23 int8 = 0
                                                    _ = v23
                                                }
                                                var v24 int8 = 0
                                                _ = v24
                                            } else {
                                                switch "foo bar" {
                                                case "hello":
                                                    v18 = append(v18, fi12)
                                                case "\n\t":
                                                    v25 := GenS1{F0: 255, F1: 1, F2: 42}
                                                    _ = v25
                                                }
                                                if len(v18) > 0 {
                                                    v18[0] = (^(fi17))
                                                }
                                            }
                                        case "\n\t":
                                            var v26 uint32 = 1
                                            _ = v26
                                        default:
                                            v27 := GenS2{F0: p0, F1: v16}
                                            _ = v27
                                        }
                                        switch (fi17 & fi12) {
                                        case 0:
                                            v28 := GenS3{F0: (v6 * (v6 | uint64(v7))), F1: fi12}
                                            _ = v28
                                        default:
                                            v29 := GenS2{F0: v11, F1: (v16 & -850)}
                                            _ = v29
                                        }
                                        v30 := GenS4{F0: 0}
                                        _ = v30
                                    }
                                case 128:
                                    v31 := GenS5{F0: "foo bar"}
                                    _ = v31
                                }
                                switch ("hello" + "") {
                                case "\n\t":
                                    v18 = append(v18, (fi12 % (((fi12 + 9223372036854775807)) | 1)))
                                }
                                var v32 int64 = -9223372036854775808
                                _ = fi17
                                _ = v18
                                _ = v32
                            }
                            if (!true) {
                                for fi33 := 0; fi33 < 3; fi33++ {
                                    switch (-(fi12)) {
                                    case 42:
                                        if ("foo bar" > ("a" + "hello")) {
                                            var v34 float32 = -2.25
                                            v35 := GenS5{F0: ("a" + "foo bar")}
                                            v34 -= (v34 / (v34 * (3.14159 - float32(((float64(float32((float64(v34) / 1.5))) / ((1.5 * -2.25) / -2.25)) + (3.14159 * 1e10))))))
                                            _ = v34
                                            _ = v35
                                        } else {
                                            fmt.Println(fi12)
                                            fmt.Println(fi33, fi33)
                                            v36 := GenS5{F0: ""}
                                            _ = v36
                                        }
                                    case 128:
                                        v6 = v6
                                    default:
                                        fmt.Println(v6, v6)
                                    }
                                    p0 += v11
                                    v37 := GenS6{F0: (^(v7))}
                                    _ = fi33
                                    _ = v37
                                }
                                switch "" {
                                case "":
                                    v38 := GenS6{F0: v7}
                                    _ = v38
                                }
                            } else {
                                fmt.Println(fi12, v6)
                                var v39 int16 = 1
                                v40 := GenS5{F0: ""}
                                _ = v39
                                _ = v40
                            }
                            var v41 uint16 = uint16(v7)
                            _ = v41
                        } else {
                            var v42 int = (fi12 & (fi12 + fi12))
                            _ = v42
                        }
                        _ = v16
                    }
                case "foo bar":
                    v43 := []int{fi12}
                    _ = v43
                }
                _ = fi12
            }
        } else {
            var v44 bool = true
            _ = v44
        }
        _ = v11
    }
    switch 2147483647 {
    case 0:
        switch "a" {
        case "\n\t":
            v6 = v6
        case "hello":
            for fi45 := 0; fi45 < 13; fi45++ {
                switch "hello" {
                case "":
                    for fi46 := 0; fi46 < 16; fi46++ {
                        if false {
                            if true {
                                var v47 uint = (p0 + 0)
                                v47 = v47
                                var v48 float32 = float32(3.14159)
                                fmt.Println(v48)
                                _ = v47
                                _ = v48
                            } else {
                                v7 = (v7 - v7)
                                if false {
                                    v6 = (v6 % ((4294967295) | 1))
                                } else {
                                    var v49 int = 124
                                    _ = v49
                                }
                            }
                            var v50 uint16 = 228
                            v7 = v7
                            _ = v50
                        } else {
                            if ((v7 / ((127) | 1)) != (^(v7))) {
                                v7 = 0
                            }
                            v51 := []int{(fi45 % (((fi46 - (fi45 | (fi45 / ((fi45) | 1))))) | 1)), fi46, 32768}
                            switch "hello" {
                            case "a":
                                var v52 int16 = 56
                                _ = v52
                            case "foo bar":
                                v7 += (v7 & 0)
                            default:
                                var v53 bool = (false && ((p0 % ((p0) | 1)) <= (p0 % ((p0) | 1))))
                                _ = v53
                            }
                            _ = v51
                        }
                        v54 := GenS7{F0: -1, F1: 626}
                        v7 = (v7 * (v7 >> (((p0 / ((p0) | 1))) & 63)))
                        _ = fi46
                        _ = v54
                    }
                case "hello":
                    v6 = (^(v6))
                case "\n\t":
                    for fi55 := 0; fi55 < 15; fi55++ {
                        v56 := GenS7{F0: int16(fi55), F1: 32767}
                        _ = fi55
                        _ = v56
                    }
                default:
                    switch "a" {
                    case "foo bar":
                        v57 := []int{fi45, (-(fi45)), 9223372036854775807}
                        _ = v57
                    case "a":
                        switch fi45 {
                        case 4294967295:
                            if false {
                                v6 = 1
                                var v58 string = ("" + "\n\t")
                                _ = v58
                            } else {
                                p0 = (p0 << ((18446744073709551615) & 63))
                                v59 := []int{(fi45 ^ 32768), (^(fi45))}
                                _ = v59
                            }
                        default:
                            var v60 int32 = 2147483647
                            _ = v60
                        }
                    }
                }
                v61 := GenS3{F0: (^(v6)), F1: fi45}
                var v62 float32 = -2.25
                _ = fi45
                _ = v61
                _ = v62
            }
        case "a":
            fmt.Println(p0, v6, v6)
        default:
            var v63 uint8 = v7
            _ = v63
        }
    default:
        v7 = v7
    }
    _ = p0
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1(p0 int16) uint8 {
    v64 := []int{}
    p0 = p0
    _ = p0
    _ = v64
    return 185
}

func main() {
    v1 := []uint32{545}
    v2 := []uint32{}
    v3 := []string{"hello", "\n\t"}
    var v4 uint16 = 1
    v5 := []string{"\n\t", "a", "a", "\n\t"}
    v4 = 255
    v4 -= (^(v4))
    for fi65 := 0; fi65 < 8; fi65++ {
        switch fi65 {
        case 84:
            fmt.Println(v4, fi65)
        }
        var v66 uint64 = 0
        var v67 float32 = (1.5 / 3.14159)
        _ = fi65
        _ = v66
        _ = v67
    }
    v4 -= 0
    _ = len(v3)
    if false {
        v2 = append(v2, 933)
    } else {
        for fi68 := 0; fi68 < 7; fi68++ {
            v4 = 65535
            var v69 uint32 = 65535
            if (!((^(v69)) == 502)) {
                if true {
                    for fi70 := 0; fi70 < 20; fi70++ {
                        var v71 float32 = (1e10 - (0.0 * 0.0))
                        var v72 int32 = -114
                        v73 := GenS5{F0: "\n\t"}
                        _ = fi70
                        _ = v71
                        _ = v72
                        _ = v73
                    }
                    switch "a" {
                    case "a":
                        v4 -= (v4 ^ v4)
                    case "":
                        v74 := GenS1{F0: v4, F1: 9223372036854775807, F2: -309}
                        _ = v74
                    }
                } else {
                    var v75 uint32 = v69
                    switch "a" {
                    case "a":
                        v76 := GenS8{F0: (v69 / (((v75 * (v69 >> ((0) & 63)))) | 1)), F1: 0, F2: 0}
                        _ = v76
                    case "":
                        v4 -= (v4 << ((759) & 63))
                    case "foo bar":
                        switch ("hello" + "a") {
                        case "":
                            v77 := GenS1{F0: 65535, F1: -990, F2: 0}
                            _ = v77
                        case "hello":
                            v2 = append(v2, (v69 | 497))
                        case "\n\t":
                            var v78 bool = (2147483648 >= v69)
                            _ = v78
                        default:
                            _ = len(v5)
                        }
                    default:
                        v4 -= (v4 >> ((65535) & 63))
                    }
                    var v79 int8 = -84
                    _ = v75
                    _ = v79
                }
            } else {
                if true {
                    v69 -= v69
                } else {
                    v2 = append(v2, v69)
                }
            }
            _ = fi68
            _ = v69
        }
    }
    switch ("a" + "hello") {
    case "":
        v4 = v4
    case "hello":
        fmt.Println(v4)
    case "a":
        var v80 string = "foo bar"
        _ = v80
    default:
        for fi81 := 0; fi81 < 9; fi81++ {
            v82 := GenS3{F0: 2, F1: (fi81 << ((42) & 63))}
            _ = fi81
            _ = v82
        }
    }
    for fi83 := 0; fi83 < 5; fi83++ {
        if true {
            var v84 int16 = 42
            _ = v84
        } else {
            v4 -= v4
        }
        _ = fi83
    }
    v4 = 0
    v85 := GenS8{F0: 714, F1: v4, F2: 2}
    v86 := GenS9{F0: 0, F1: 470}
    v4 = (v4 * 753)
    fmt.Println(v4)
    for fi87 := 0; fi87 < 3; fi87++ {
        var v88 bool = (0 != 826)
        v4 -= 341
        _ = fi87
        _ = v88
    }
    switch "" {
    case "hello":
        var v89 uint64 = 828
        _ = v89
    }
    for fi90 := 0; fi90 < 13; fi90++ {
        var v91 string = ("\n\t" + "hello")
        if false {
            if false {
                if false {
                    for fi92 := 0; fi92 < 2; fi92++ {
                        v93 := GenS2{F0: 183, F1: 42}
                        switch (v91 + (((((v91 + v91) + (v91 + "")) + v91) + (("hello" + "a") + v91)) + (v91 + v91))) {
                        case "":
                            v94 := GenS4{F0: -9223372036854775808}
                            _ = v94
                        }
                        _ = fi92
                        _ = v93
                    }
                }
                for fi95 := 0; fi95 < 6; fi95++ {
                    var v96 uint32 = 2147483648
                    _ = fi95
                    _ = v96
                }
            } else {
                var v97 bool = true
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v97
            }
            v4 -= 319
            for fi98 := 0; fi98 < 10; fi98++ {
                v4 = v4
                v99 := GenS5{F0: v91}
                v100 := GenS10{F0: "foo bar", F1: 185}
                fmt.Println(fi98)
                var v101 float32 = (1e10 + 1e10)
                _ = fi98
                _ = v100
                _ = v101
                _ = v99
            }
        }
        var v102 int = fi90
        _ = fi90
        _ = v102
        _ = v91
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v85
    _ = v86
}
