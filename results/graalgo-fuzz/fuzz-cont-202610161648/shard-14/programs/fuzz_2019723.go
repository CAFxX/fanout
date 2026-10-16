package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int8; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint16; F1 uint16; F2 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int32; F1 int16; F2 uint16 }
type GenS8 struct { F0 uint8; F1 uint; F2 uint64 }
func genFn0() string {
    if (!false) {
        var v4 int8 = 1
        var v5 int64 = 1
        switch -156 {
        case 9223372036854775807:
            v6 := []int{-684, 4294967295}
            _ = v6
        case 4294967295:
            for fi7 := 0; fi7 < 15; fi7++ {
                for fi8 := 0; fi8 < 12; fi8++ {
                    v5 += -9223372036854775808
                    _ = fi8
                }
                var v9 uint64 = 867
                _ = fi7
                _ = v9
            }
        case 0:
            v4 += v4
        default:
            var v10 int = -879
            _ = v10
        }
        _ = v4
        _ = v5
    }
    if (0 < 100) {
        var v11 int64 = 414
        for fi12 := 0; fi12 < 17; fi12++ {
            switch "a" {
            case "hello":
                var v13 string = "\n\t"
                _ = v13
            default:
                v11 -= v11
            }
            _ = fi12
        }
        var v14 int64 = 42
        _ = v11
        _ = v14
    }
    v15 := GenS1{F0: 813, F1: 737}
    _ = v15
    return ("a" + "hello")
}
func genFn1() int64 {
    if (true || (false || true)) {
        var v16 float32 = 1.5
        if true {
            v16 = v16
        }
        for fi17 := 0; fi17 < 7; fi17++ {
            fmt.Println(v16, v16)
            var v18 int64 = -961
            fmt.Println(v18, fi17, v16)
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        v20 := []int{fi19, (fi19 ^ (-(fi19))), 128}
        v20 = append(v20, (fi19 | 7))
        fmt.Println(fi19)
        _ = fi19
        _ = v20
    }
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 4294967295
    v3 := make(map[int]string)
    v2 = v2
    v21 := GenS1{F0: 547, F1: 42}
    for fi22 := 0; fi22 < 5; fi22++ {
        for fi23 := 0; fi23 < 19; fi23++ {
            if (-638 < -437) {
                switch int(v2) {
                case 1:
                    v24 := v3[(fi23 * (fi22 ^ (fi23 - (-(fi23)))))]
                    _ = v24
                case 9223372036854775807:
                    fmt.Println(v2)
                case -2147483648:
                    switch "foo bar" {
                    case "foo bar":
                        v25 := GenS1{F0: -619, F1: 70}
                        _ = v25
                    case "":
                        var v26 int16 = -992
                        _ = v26
                    default:
                        v2 = v2
                    }
                }
                if false {
                    if (231 < 0) {
                        v2 = v2
                        if ((true || true) && false) {
                            switch "foo bar" {
                            case "foo bar":
                                var v27 uint16 = 1
                                _ = v27
                            default:
                                v2 = (v2 & (^(v2)))
                            }
                            var v28 float32 = 0.1
                            v2 -= v2
                            v29 := []int{(fi22 - (-(fi22))), fi22}
                            var v30 float32 = (v28 * v28)
                            _ = v28
                            _ = v29
                            _ = v30
                        } else {
                            var v31 uint32 = 1
                            _ = v31
                        }
                        var v32 int64 = 779
                        _ = v32
                    }
                    v33 := v1[fi23]
                    _ = v33
                }
            }
            if false {
                v34 := GenS2{F0: 0, F1: (v2 + v2), F2: -56}
                _ = v34
            } else {
                var v35 uint16 = 2
                fmt.Println(fi22)
                if true {
                    if (genFn1() <= 9223372036854775807) {
                        v36 := GenS3{F0: 0, F1: 0, F2: "hello"}
                        v35 = v35
                        v37 := GenS3{F0: 42, F1: 65535, F2: ""}
                        _ = v36
                        _ = v37
                    }
                } else {
                    v38 := GenS4{F0: 347}
                    if true {
                        v39 := GenS5{F0: 1, F1: v35}
                        v35 -= v35
                        switch (fi23 | (^(fi22))) {
                        case 42:
                            v40 := GenS3{F0: 676, F1: (v35 - v35), F2: "\n\t"}
                            _ = v40
                        case 9223372036854775807:
                            v35 = (v35 ^ 405)
                            v35 = v35
                            var v41 string = ("" + "hello")
                            _ = v41
                        }
                        _ = v39
                    }
                    switch "hello" {
                    case "":
                        v42 := GenS2{F0: 28, F1: (v2 - 1), F2: 1}
                        _ = v42
                    case "foo bar":
                        var v43 uint16 = v35
                        _ = v43
                    case "hello":
                        var v44 int8 = 1
                        _ = v44
                    }
                    if (false && (!true)) {
                        if true {
                            switch (fi22 / ((4294967295) | 1)) {
                            case 32768:
                                _ = len(v1)
                            case 9223372036854775807:
                                v35 = 65535
                            }
                            var v45 uint8 = 2
                            _ = v45
                        }
                    }
                    v46 := GenS6{F0: (false && true)}
                    _ = v38
                    _ = v46
                }
                _ = v35
            }
            v1[-889] = 65535
            _ = fi23
        }
        _ = fi22
    }
    var v47 uint8 = 113
    v48 := GenS7{F0: 0, F1: 709, F2: 65535}
    if (!false) {
        for fi49 := 0; fi49 < 15; fi49++ {
            v2 += 42
            var v50 float64 = 1e10
            for fi51 := 0; fi51 < 13; fi51++ {
                var v52 uint64 = (v2 % ((0) | 1))
                switch "a" {
                case "hello":
                    switch ("foo bar" + "foo bar") {
                    case "foo bar":
                        v52 = 441
                    default:
                        v53 := []int{fi51, fi51, fi51, (fi49 + fi49)}
                        _ = v53
                    }
                case "a":
                    switch "" {
                    case "\n\t":
                        v2 += v2
                    case "a":
                        v50 -= (((float64((0.0 / 1.0)) - float64(0.1)) / (v50 * float64((1e10 * -2.25)))) - -2.25)
                    }
                    var v54 bool = false
                    _ = v54
                case "foo bar":
                    v52 -= (v52 % ((v52) | 1))
                default:
                    v47 += 1
                    if true {
                        fmt.Println(v50, fi49)
                    } else {
                        switch "a" {
                        case "\n\t":
                            var v55 int8 = -7
                            _ = v55
                        }
                    }
                    switch "foo bar" {
                    case "":
                        v56 := GenS8{F0: (^(v47)), F1: 0, F2: 42}
                        _ = v56
                    case "foo bar":
                        if false {
                            switch (^(fi51)) {
                            case -2147483648:
                                var v57 int8 = -8
                                _ = v57
                            case 9223372036854775807:
                                var v58 int64 = 9223372036854775807
                                _ = v58
                            case 128:
                                fmt.Println(fi51, v2)
                            }
                            var v59 float32 = 1.5
                            v47 += 1
                            _ = v59
                        }
                    case "a":
                        var v60 int = (fi49 & fi49)
                        _ = v60
                    default:
                        var v61 int8 = -64
                        _ = v61
                    }
                    v52 = (^(v52))
                }
                _ = fi51
                _ = v52
            }
            _ = fi49
            _ = v50
        }
        if false {
            fmt.Println(v2, v47)
            v2 += v2
            var v62 uint64 = (^(v2))
            _ = v62
        }
    }
    switch 9223372036854775807 {
    case 2147483647:
        var v63 float64 = 0.1
        _ = v63
    }
    var v64 float64 = (3.14159 * 3.14159)
    v47 -= 255
    if ((v47 <= (v47 + v47)) && false) {
        _ = len(v1)
        v64 = v64
    }
    var v65 int8 = 1
    v65 = v65
    var v66 uint32 = 395
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v47:", v47)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v47
    _ = v48
    _ = v64
    _ = v65
    _ = v66
}
