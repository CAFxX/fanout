package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 uint32; F2 int32 }
type GenS4 struct { F0 string; F1 uint8; F2 int32 }
type GenS5 struct { F0 int; F1 bool; F2 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint32; F1 string }
func genFn0(p0 int8) uint {
    fmt.Println(p0)
    v7 := GenS1{F0: 127, F1: -32768}
    if (p0 >= -128) {
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := []int{(-(fi8)), (fi8 << ((42) & 63)), fi8}
            v10 := GenS1{F0: 219, F1: -32768}
            _ = fi8
            _ = v10
            _ = v9
        }
        fmt.Println(p0)
    } else {
        v11 := []int{42, -2147483648, 1, -878}
        _ = v11
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        switch fi12 {
        case 1:
            fmt.Println(fi12, p0)
        case 513:
            switch (fi12 << ((870) & 63)) {
            case 9223372036854775807:
                if false {
                    switch fi12 {
                    case -2147483648:
                        p0 -= p0
                    case 9223372036854775807:
                        v13 := GenS1{F0: 1, F1: -664}
                        _ = v13
                    case -1:
                        v14 := GenS2{F0: 1}
                        _ = v14
                    default:
                        v15 := []int{9223372036854775807, (^(fi12)), fi12}
                        _ = v15
                    }
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(fi12)
                    default:
                        var v16 uint32 = 4294967295
                        _ = v16
                    }
                    if false {
                        p0 = p0
                    } else {
                        switch -9223372036854775808 {
                        case 128:
                            var v17 float64 = (0.0 * 0.1)
                            _ = v17
                        }
                        v18 := []int{fi12, fi12, 599}
                        switch ("foo bar" + ("a" + "\n\t")) {
                        case "hello":
                            _ = len(v18)
                        case "a":
                            if false {
                                for fi19 := 0; fi19 < 5; fi19++ {
                                    v20 := GenS3{F0: (767 == 42), F1: 42, F2: -2147483648}
                                    _ = fi19
                                    _ = v20
                                }
                                v21 := GenS1{F0: 42, F1: -1}
                                p0 = (p0 | p0)
                                _ = v21
                            } else {
                                for fi22 := 0; fi22 < 20; fi22++ {
                                    p0 = (p0 / (((p0 + p0)) | 1))
                                    _ = fi22
                                }
                                p0 = 127
                                if true {
                                    v23 := GenS1{F0: 8, F1: -94}
                                    _ = v23
                                } else {
                                    fmt.Println(fi12, fi12)
                                    var v24 float64 = (((1e10 * (((0.1 + float64(-2.25)) / (-2.25 - 0.1)) / (1.5 * 1.5))) - 3.14159) + 3.14159)
                                    _ = v24
                                }
                            }
                        }
                        _ = v18
                    }
                }
            }
        case 32768:
            var v25 float64 = (0.0 * 1e10)
            _ = v25
        default:
            v26 := GenS1{F0: 1, F1: -532}
            _ = v26
        }
        for fi27 := 0; fi27 < 12; fi27++ {
            switch "" {
            case "\n\t":
                p0 += p0
            case "hello":
                if ("" > ("\n\t" + "foo bar")) {
                    fmt.Println(fi12, fi12)
                } else {
                    var v28 int64 = 1
                    _ = v28
                }
            default:
                if (!false) {
                    var v29 float64 = -2.25
                    if (true || true) {
                        var v30 uint32 = 0
                        if true {
                            v31 := GenS2{F0: uint8(v30)}
                            v32 := GenS3{F0: ((247 >= -1) && true), F1: (v30 ^ v30), F2: 42}
                            _ = v31
                            _ = v32
                        }
                        v33 := []int{}
                        _ = v30
                        _ = v33
                    } else {
                        fmt.Println(p0)
                        v34 := []int{fi27, fi27}
                        _ = v34
                    }
                    var v35 int32 = 42
                    _ = v29
                    _ = v35
                }
            }
            var v36 int8 = p0
            if true {
                var v37 int = fi27
                fmt.Println(p0)
                _ = v37
            } else {
                var v38 float64 = -2.25
                fmt.Println(p0)
                _ = v38
            }
            _ = fi27
            _ = v36
        }
        _ = fi12
    }
    var v39 float32 = 0.0
    _ = p0
    _ = v39
    _ = v7
    return 987
}
func genFn1(p0 int16, p1 int32) int {
    var v40 string = "\n\t"
    for fi41 := 0; fi41 < 17; fi41++ {
        for fi42 := 0; fi42 < 5; fi42++ {
            p0 += p0
            var v43 float32 = 3.14159
            if true {
                switch (fi42 % (((fi42 % (((-(fi42))) | 1))) | 1)) {
                case -271:
                    if (!(!true)) {
                        var v44 float32 = v43
                        _ = v44
                    } else {
                        var v45 uint64 = 18446744073709551615
                        p0 = p0
                        _ = v45
                    }
                }
                fmt.Println(fi42, fi42)
            } else {
                v46 := GenS1{F0: 122, F1: (p0 >> ((65535) & 63))}
                v47 := []int{32768, (fi42 + (fi41 ^ 128)), fi41}
                v43 = float32((0.1 * 1.5))
                _ = v46
                _ = v47
            }
            switch v40 {
            case "":
                switch fi41 {
                case 264:
                    var v48 uint64 = 923
                    _ = v48
                case -9223372036854775808:
                    var v49 float32 = v43
                    _ = v49
                case 9223372036854775807:
                    var v50 uint32 = 1
                    _ = v50
                default:
                    switch fi41 {
                    case 4294967295:
                        if (510 < 48) {
                            if (-999 == -1) {
                                switch fi41 {
                                case 4294967295:
                                    switch 18 {
                                    case 128:
                                        switch 4294967295 {
                                        case 127:
                                            v40 = ""
                                        case 1:
                                            v51 := GenS2{F0: 42}
                                            _ = v51
                                        default:
                                            var v52 bool = ((42 >= 4294967295) || false)
                                            _ = v52
                                        }
                                    }
                                    if (0.1 != (1e10 / float64(v43))) {
                                        if true {
                                            v53 := []int{2147483647, -9223372036854775808, (fi41 - (fi42 / ((32767) | 1)))}
                                            if len(v53) >= 2 {
                                                v53 = v53[1:len(v53)]
                                            }
                                            switch (("\n\t" + ("foo bar" + (v40 + ("" + v40)))) + v40) {
                                            case "\n\t":
                                                fmt.Println(fi42, p0)
                                            case "":
                                                fmt.Println(v40)
                                            }
                                            _ = v53
                                        }
                                        switch (v40 + "") {
                                        case "foo bar":
                                            if false {
                                                fmt.Println(p1, fi42)
                                            } else {
                                                v54 := []int{(-(fi42)), (fi42 << ((112) & 63)), fi41}
                                                _ = v54
                                            }
                                        }
                                    }
                                    var v55 int16 = 1
                                    _ = v55
                                case -2147483648:
                                    if false {
                                        v56 := GenS3{F0: true, F1: 42, F2: p1}
                                        v57 := GenS4{F0: "a", F1: 42, F2: (p1 & (-(p1)))}
                                        v43 = float32((3.14159 + -2.25))
                                        _ = v56
                                        _ = v57
                                    } else {
                                        if true {
                                            if (0 < 5) {
                                                var v58 int32 = (^(p1))
                                                _ = v58
                                            }
                                            var v59 int8 = -1
                                            _ = v59
                                        }
                                        p0 = (p0 << ((18446744073709551615) & 63))
                                    }
                                case 127:
                                    v60 := GenS5{F0: -9223372036854775808, F1: (538 < 578), F2: 18446744073709551615}
                                    _ = v60
                                }
                            } else {
                                v61 := []int{fi42, (fi41 | (fi42 ^ -878)), fi42, (fi41 + fi41)}
                                _ = v61
                            }
                            v62 := GenS3{F0: false, F1: 4294967295, F2: (p1 / ((p1) | 1))}
                            _ = v62
                        } else {
                            v63 := GenS6{F0: 127}
                            _ = v63
                        }
                    case -189:
                        fmt.Println(p0, fi41)
                    case -2147483648:
                        var v64 int = -2147483648
                        _ = v64
                    default:
                        if true {
                            p1 = (-(p1))
                        } else {
                            var v65 int = fi42
                            v43 = v43
                            _ = v65
                        }
                    }
                }
            }
            _ = fi42
            _ = v43
        }
        v66 := GenS7{F0: 1, F1: ("a" + (v40 + "a"))}
        _ = fi41
        _ = v66
    }
    v40 = "foo bar"
    _ = p0
    _ = p1
    _ = v40
    return 2147483647
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    var v3 float32 = float32(0.1)
    var v4 bool = false
    v5 := []uint32{42, 2, 2147483648, 2147483647}
    var v6 uint32 = 65535
    for fi67 := 0; fi67 < 15; fi67++ {
        if true {
            for fi68 := 0; fi68 < 17; fi68++ {
                var v69 int32 = 42
                v70 := GenS3{F0: (348 >= 255), F1: 65535, F2: v69}
                var v71 uint = 255
                _ = fi68
                _ = v69
                _ = v70
                _ = v71
            }
        } else {
            v6 += v6
        }
        _ = fi67
    }
    v5 = append(v5, (v6 ^ (v6 >> ((genFn0(100)) & 63))))
    v3 = float32((1e10 + 3.14159))
    v4 = true
    delete(v2, v6)
    var v72 int64 = 0
    _ = len(v5)
    v73 := GenS1{F0: 0, F1: -32768}
    v6 = (v6 << ((994) & 63))
    for fi74 := 0; fi74 < 15; fi74++ {
        var v75 float64 = -2.25
        var v76 bool = v4
        _ = fi74
        _ = v75
        _ = v76
    }
    var v77 int32 = 1
    var v78 uint = 42
    if false {
        var v79 uint16 = 42
        delete(v2, v6)
        if ((float32(((-2.25 - -2.25) / ((3.14159 / (1.5 / 3.14159)) / 3.14159))) / v3) == v3) {
            var v80 int16 = -1
            v3 -= (v3 / v3)
            _ = v80
        }
        _ = v79
    }
    if len(v5) > 0 {
        v5[0] = v6
    }
    var v81 int64 = (v72 >> ((4294967295) & 63))
    fmt.Println(v77)
    v3 = (float32(float64(-2.25)) + float32(1.5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v72:", v72)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v72
    _ = v73
    _ = v77
    _ = v78
    _ = v81
}
