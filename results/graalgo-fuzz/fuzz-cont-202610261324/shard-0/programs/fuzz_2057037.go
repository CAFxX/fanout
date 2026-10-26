package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 int; F2 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string; F1 uint16; F2 uint16 }
type GenS6 struct { F0 int16; F1 int64; F2 int16 }
type GenS7 struct { F0 int32; F1 uint; F2 string }
type GenS8 struct { F0 bool; F1 uint16; F2 uint8 }
func genFn0(p0 int8) int8 {
    var v5 int64 = 9223372036854775807
    if ("a" < "hello") {
        switch -66 {
        case 4294967295:
            v6 := GenS1{F0: 401, F1: 1, F2: 27}
            _ = v6
        default:
            switch 32767 {
            case -362:
                var v7 float32 = -2.25
                _ = v7
            case 737:
                switch "\n\t" {
                case "":
                    var v8 bool = ((!(0 == 602)) && true)
                    _ = v8
                case "foo bar":
                    var v9 uint16 = 0
                    _ = v9
                default:
                    var v10 int32 = 783
                    _ = v10
                }
            case 127:
                v11 := GenS1{F0: -764, F1: -1, F2: 1}
                _ = v11
            default:
                v5 -= v5
            }
        }
        if false {
            for fi12 := 0; fi12 < 9; fi12++ {
                switch "a" {
                case "foo bar":
                    if (!((fi12 << ((42) & 63)) < fi12)) {
                        v13 := []int{(fi12 % ((fi12) | 1)), 0, (fi12 | fi12), fi12}
                        _ = v13
                    }
                default:
                    v14 := GenS1{F0: fi12, F1: 42, F2: 577}
                    _ = v14
                }
                v15 := []int{(fi12 | (fi12 & -921))}
                var v16 uint16 = 32767
                _ = fi12
                _ = v15
                _ = v16
            }
            var v17 uint64 = 1
            _ = v17
        }
        if (false || (2147483647 == 9223372036854775807)) {
            v5 -= 404
        }
    }
    p0 -= (p0 & (-(p0)))
    for fi18 := 0; fi18 < 17; fi18++ {
        v19 := GenS2{F0: -410}
        _ = fi18
        _ = v19
    }
    if true {
        if (42 > 485) {
            v20 := []int{-1, 0, 1, 127}
            for fi21 := 0; fi21 < 1; fi21++ {
                v5 = v5
                _ = fi21
            }
            _ = v20
        }
        var v22 float64 = 1e10
        _ = v22
    } else {
        v23 := []int{}
        _ = v23
    }
    _ = p0
    _ = v5
    return (^(p0))
}
func genFn1(p0 uint16) uint64 {
    for fi24 := 0; fi24 < 5; fi24++ {
        if false {
            v25 := []int{fi24}
            p0 = (p0 << ((18446744073709551615) & 63))
            _ = v25
        } else {
            var v26 uint16 = p0
            var v27 int = 349
            v28 := []int{v27}
            _ = v26
            _ = v27
            _ = v28
        }
        _ = fi24
    }
    var v29 int16 = 0
    v30 := []int{42, 128, 1, 832}
    if len(v30) > 0 {
        v30[0] = 281
    }
    switch ("" + "") {
    case "a":
        v31 := GenS1{F0: -1, F1: -1, F2: 646}
        _ = v31
    default:
        var v32 bool = (!(0 <= (v29 ^ v29)))
        _ = v32
    }
    _ = p0
    _ = v29
    _ = v30
    return 4294967295
}

func main() {
    var v1 uint8 = 42
    var v2 uint = 2
    var v3 uint32 = 65535
    v4 := make(map[int]string)
    v3 = v3
    if ((^(v2)) < uint(v3)) {
        var v33 int16 = -81
        _ = v33
    }
    v34 := GenS2{F0: int(v1)}
    var v35 int32 = -355
    v3 = 2147483648
    fmt.Println(v2)
    var v36 uint64 = genFn1(42)
    fmt.Println(v35)
    for fi37 := 0; fi37 < 15; fi37++ {
        v36 += (v36 ^ (v36 | genFn1(uint16(v1))))
        switch "hello" {
        case "":
            _ = len(v4)
        case "foo bar":
            switch "a" {
            case "a":
                for fi38 := 0; fi38 < 13; fi38++ {
                    var v39 int16 = -436
                    _ = fi38
                    _ = v39
                }
            case "":
                v35 = v35
            case "hello":
                var v40 int = fi37
                _ = v40
            default:
                _ = len(v4)
            }
        case "\n\t":
            v41 := GenS1{F0: fi37, F1: v35, F2: (v2 & v2)}
            _ = v41
        }
        v42 := []int{(fi37 % ((fi37) | 1)), (fi37 % ((-9223372036854775808) | 1)), fi37, int(v3)}
        _ = fi37
        _ = v42
    }
    var v43 float32 = 1.5
    v3 = 2147483648
    switch -9223372036854775808 {
    case -395:
        v44 := GenS2{F0: -11}
        _ = v44
    case -736:
        v45 := GenS3{F0: v3, F1: int(v2), F2: v35}
        _ = v45
    }
    v46 := GenS4{F0: v3}
    switch (("a" + "a") + "a") {
    case "foo bar":
        v47 := GenS4{F0: v3}
        _ = v47
    case "hello":
        v48 := GenS3{F0: 4294967295, F1: -78, F2: v35}
        _ = v48
    default:
        switch "hello" {
        case "foo bar":
            var v49 string = "\n\t"
            _ = v49
        case "a":
            v4[-9223372036854775808] = "foo bar"
        }
    }
    if true {
        v2 = (v2 / (((v2 | 0)) | 1))
        v1 = v1
    }
    if (!((^(v36)) != 859)) {
        v50 := GenS3{F0: v3, F1: 1, F2: v35}
        if (((v43 + 1.5) / 3.14159) >= (v43 + v43)) {
            v43 = 1.5
        }
        switch ("a" + ("" + "hello")) {
        case "\n\t":
            v51 := v4[9223372036854775807]
            _ = v51
        case "foo bar":
            if (v35 < v35) {
                if false {
                    for fi52 := 0; fi52 < 5; fi52++ {
                        if false {
                            fmt.Println(v2, v2)
                            v53 := GenS5{F0: "", F1: 2, F2: 0}
                            _ = v53
                        } else {
                            v4[685] = "hello"
                            switch fi52 {
                            case -9223372036854775808:
                                v54 := []int{fi52, int(v36), (fi52 ^ (fi52 - fi52)), fi52}
                                _ = v54
                            case 75:
                                v36 = (^(v36))
                            }
                        }
                        if false {
                            if (("foo bar" + "\n\t") < "") {
                                v36 = (v36 | (v36 % (((^(v36))) | 1)))
                            }
                            fmt.Println(v3, v2)
                        } else {
                            if (!("" < ("\n\t" + ("" + "a")))) {
                                for fi55 := 0; fi55 < 20; fi55++ {
                                    fmt.Println(v43)
                                    if (genFn1(1) > v36) {
                                        v43 = v43
                                        v56 := GenS5{F0: ("\n\t" + "\n\t"), F1: 0, F2: 0}
                                        var v57 int32 = -2147483648
                                        switch (("foo bar" + "") + ("hello" + ("hello" + ("hello" + "hello")))) {
                                        case "hello":
                                            fmt.Println(v57)
                                        default:
                                            v58 := []int{188}
                                            _ = v58
                                        }
                                        _ = v56
                                        _ = v57
                                    } else {
                                        var v59 int32 = (v35 % ((v35) | 1))
                                        if ((^(v1)) < (v1 + v1)) {
                                            v60 := GenS4{F0: (v3 & (v3 | (v3 | (v3 | (^(v3))))))}
                                            v61 := GenS6{F0: 32767, F1: int64(fi55), F2: 0}
                                            var v62 uint16 = 65535
                                            _ = v60
                                            _ = v61
                                            _ = v62
                                        }
                                        _ = v59
                                    }
                                    _ = fi55
                                }
                            }
                        }
                        var v63 int32 = v35
                        _ = fi52
                        _ = v63
                    }
                    v36 = (^(v36))
                    var v64 uint32 = v3
                    _ = v64
                }
                v43 = ((float32((3.14159 + (((0.1 / (-2.25 + (float64(v43) * 0.0))) / 3.14159) + -2.25))) / (v43 / v43)) - float32(0.0))
                fmt.Println(v35, v35, v43)
            } else {
                v65 := []int{101, 127, int(v1), 127}
                _ = v65
            }
        default:
            v66 := GenS7{F0: (v35 + 42), F1: (v2 - v2), F2: "a"}
            _ = v66
        }
        _ = v50
    } else {
        v3 += 389
    }
    if false {
        for fi67 := 0; fi67 < 1; fi67++ {
            v4[2147483647] = "hello"
            _ = fi67
        }
    } else {
        var v68 uint32 = (v3 & uint32(v2))
        v2 = v2
        switch int(v1) {
        case 127:
            v43 = float32((3.14159 - 0.0))
        case 128:
            switch -1 {
            case 128:
                v69 := GenS8{F0: true, F1: 65535, F2: uint8(v68)}
                _ = v69
            case 22:
                var v70 int64 = 42
                _ = v70
            default:
                v4[32768] = "a"
            }
        default:
            v71 := GenS5{F0: "a", F1: 255, F2: 1}
            _ = v71
        }
        _ = v68
    }
    if (-732 <= 9223372036854775807) {
        for fi72 := 0; fi72 < 19; fi72++ {
            v2 = (v2 % ((2) | 1))
            var v73 int16 = -403
            _ = fi72
            _ = v73
        }
        for fi74 := 0; fi74 < 12; fi74++ {
            var v75 uint8 = 128
            var v76 uint32 = (v3 | v3)
            var v77 int = (fi74 | (-(fi74)))
            _ = fi74
            _ = v75
            _ = v76
            _ = v77
        }
    } else {
        var v78 uint64 = uint64(v3)
        switch -9223372036854775808 {
        case 9223372036854775807:
            if false {
                v1 = v1
            } else {
                if (0 != int(v36)) {
                    fmt.Println(v78, v78, v43)
                    v3 = (v3 >> ((uint(v3)) & 63))
                    var v79 bool = (2147483647 > v35)
                    _ = v79
                } else {
                    var v80 int8 = 42
                    var v81 uint16 = 1
                    v81 -= 2
                    _ = v80
                    _ = v81
                }
                v82 := []int{int(v3), 787, 42, 178}
                _ = v82
            }
        case -996:
            v36 = (v78 | genFn1(uint16(v2)))
        default:
            v43 += v43
        }
        _ = v78
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v43
    _ = v46
}
