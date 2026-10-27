package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 int8 }
func genFn0(p0 int64, p1 int16, p2 string) int {
    if (1 <= 32767) {
        fmt.Println(p0, p1)
        switch p2 {
        case "foo bar":
            p0 -= (p0 & p0)
        case "\n\t":
            if false {
                var v6 uint8 = 42
                for fi7 := 0; fi7 < 9; fi7++ {
                    v6 += (v6 & 94)
                    var v8 string = (("a" + p2) + "a")
                    v6 = v6
                    _ = fi7
                    _ = v8
                }
                v9 := []int{}
                _ = v6
                _ = v9
            } else {
                v10 := GenS1{F0: p1}
                _ = v10
            }
        default:
            var v11 bool = true
            _ = v11
        }
        v12 := []int{1, -1, 7}
        _ = v12
    }
    v13 := []int{-292, 128}
    var v14 int16 = -32768
    var v15 int64 = (p0 | -655)
    for fi16 := 0; fi16 < 7; fi16++ {
        var v17 uint32 = 0
        var v18 int64 = (p0 << ((42) & 63))
        var v19 uint64 = 2
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return 199
}

func main() {
    var v1 uint8 = 128
    var v2 float64 = 0.1
    var v3 int32 = 2147483647
    var v4 uint8 = 128
    v5 := make(map[int]uint32)
    v20 := v5[32768]
    switch "a" {
    case "\n\t":
        if (!false) {
            switch ("a" + "") {
            case "foo bar":
                switch "foo bar" {
                case "hello":
                    switch "foo bar" {
                    case "":
                        var v21 string = ((("" + "hello") + ("hello" + ("hello" + "foo bar"))) + "a")
                        _ = v21
                    case "\n\t":
                        var v22 int64 = -9223372036854775808
                        _ = v22
                    }
                case "\n\t":
                    switch ("foo bar" + "hello") {
                    case "foo bar":
                        v23 := GenS1{F0: int16(v3)}
                        _ = v23
                    }
                case "foo bar":
                    v20 += (v20 | v20)
                }
            case "a":
                switch (("\n\t" + "hello") + ("" + "foo bar")) {
                case "foo bar":
                    var v24 float64 = v2
                    _ = v24
                case "\n\t":
                    v25 := GenS2{F0: 1, F1: 585, F2: 0}
                    _ = v25
                }
            default:
                var v26 float64 = (v2 - 0.0)
                _ = v26
            }
            v1 += 255
        } else {
            _ = len(v5)
            v27 := v5[128]
            _ = v27
        }
    case "a":
        for fi28 := 0; fi28 < 6; fi28++ {
            var v29 int64 = -243
            var v30 uint8 = (^(v4))
            _ = fi28
            _ = v29
            _ = v30
        }
    case "":
        if (!true) {
            var v31 int64 = 9223372036854775807
            v32 := []int{42, int(v3)}
            switch ("foo bar" + "hello") {
            case "a":
                var v33 string = "hello"
                _ = v33
            case "\n\t":
                v34 := GenS3{F0: (((("hello" + "\n\t") + ("foo bar" + "hello")) + "hello") + "hello")}
                _ = v34
            case "foo bar":
                v35 := GenS1{F0: 1000}
                _ = v35
            default:
                v36 := GenS4{F0: v3, F1: -128}
                _ = v36
            }
            _ = v31
            _ = v32
        } else {
            switch -513 {
            case 0:
                var v37 bool = false
                _ = v37
            default:
                var v38 uint16 = 2
                _ = v38
            }
            v4 = v4
            if false {
                var v39 int8 = 42
                var v40 string = "hello"
                _ = v39
                _ = v40
            } else {
                var v41 int = genFn0(779, -1, ("hello" + ""))
                _ = v41
            }
        }
    }
    v1 = v1
    switch (("hello" + ("" + "\n\t")) + "\n\t") {
    case "\n\t":
        switch ("a" + (("" + ("\n\t" + "\n\t")) + ("\n\t" + ("" + "a")))) {
        case "a":
            delete(v5, 9223372036854775807)
        }
    default:
        for fi42 := 0; fi42 < 6; fi42++ {
            if ("" < "a") {
                var v43 uint64 = 4294967295
                fmt.Println(v4)
                _ = len(v5)
                _ = v43
            } else {
                for fi44 := 0; fi44 < 12; fi44++ {
                    if (v2 <= -2.25) {
                        v45 := GenS1{F0: 0}
                        var v46 int8 = 1
                        switch fi42 {
                        case 42:
                            v47 := []int{fi42, (^(fi42)), (fi44 / (((fi42 / (((fi42 & (-(fi44)))) | 1))) | 1)), fi44}
                            _ = v47
                        }
                        _ = v45
                        _ = v46
                    }
                    var v48 float64 = (-2.25 + v2)
                    var v49 int32 = 893
                    _ = fi44
                    _ = v48
                    _ = v49
                }
                v50 := []int{genFn0(-1, 32767, ""), fi42, -971, fi42}
                var v51 uint = 1
                _ = v50
                _ = v51
            }
            _ = fi42
        }
    }
    if (124 != 88) {
        switch ("hello" + ("\n\t" + "foo bar")) {
        case "":
            if (332 != 65535) {
                var v52 uint64 = 903
                for fi53 := 0; fi53 < 12; fi53++ {
                    for fi54 := 0; fi54 < 6; fi54++ {
                        v1 -= 113
                        _ = fi54
                    }
                    for fi55 := 0; fi55 < 16; fi55++ {
                        v2 = (v2 - ((((3.14159 * -2.25) / v2) * ((((float64(1e10) * v2) - 1e10) / v2) / ((-2.25 / -2.25) + float64(0.0)))) - v2))
                        _ = fi55
                    }
                    v2 += (v2 * (v2 / v2))
                    _ = fi53
                }
                if false {
                    var v56 int8 = 0
                    for fi57 := 0; fi57 < 4; fi57++ {
                        if (false && (!true)) {
                            var v58 uint = 580
                            delete(v5, (fi57 + genFn0(-140, 1000, "\n\t")))
                            _ = v58
                        }
                        if false {
                            fmt.Println(v4, v20, v1)
                            v4 += (v1 & v1)
                            for fi59 := 0; fi59 < 5; fi59++ {
                                var v60 uint32 = (v20 & 65535)
                                _ = fi59
                                _ = v60
                            }
                        } else {
                            fmt.Println(v20)
                            var v61 bool = false
                            _ = v61
                        }
                        _ = fi57
                    }
                    _ = v56
                } else {
                    v20 += (v20 & (v20 + v20))
                    for fi62 := 0; fi62 < 8; fi62++ {
                        v52 += v52
                        v52 = v52
                        var v63 int64 = int64(fi62)
                        _ = fi62
                        _ = v63
                    }
                }
                _ = v52
            } else {
                if false {
                    var v64 float64 = 0.1
                    _ = len(v5)
                    var v65 float64 = v64
                    _ = v64
                    _ = v65
                }
                v66 := v5[2147483647]
                var v67 uint32 = 1
                _ = v66
                _ = v67
            }
        case "foo bar":
            v1 -= 2
        case "hello":
            var v68 int64 = -9223372036854775808
            _ = v68
        }
        v3 += -27
    } else {
        v69 := GenS1{F0: 1}
        for fi70 := 0; fi70 < 12; fi70++ {
            v2 += (1.5 * (v2 * float64(0.0)))
            _ = fi70
        }
        _ = v69
    }
    fmt.Println(v1)
    v71 := []int{-964}
    if true {
        if (0 <= -32768) {
            var v72 uint = 717
            if (42 > uint16(v20)) {
                v73 := GenS4{F0: (v3 >> ((740) & 63)), F1: int8(v1)}
                _ = v73
            } else {
                var v74 string = ("\n\t" + ("hello" + "a"))
                if (v1 == 0) {
                    v74 = ("" + ((v74 + v74) + "foo bar"))
                    switch 32767 {
                    case 1:
                        v4 = 2
                    }
                    var v75 int32 = int32(v72)
                    _ = v75
                }
                _ = v74
            }
            _ = v72
        } else {
            if (uint64(v20) == 442) {
                v1 = v1
            }
            _ = len(v71)
        }
        for fi76 := 0; fi76 < 19; fi76++ {
            if false {
                fmt.Println(v1, v3, v1)
                var v77 uint8 = (v4 >> ((4294967295) & 63))
                _ = v77
            } else {
                if len(v71) > 0 {
                    v71[0] = (fi76 - fi76)
                }
            }
            var v78 float64 = 1e10
            v5[genFn0(42, 787, "hello")] = (v20 / (((v20 << ((1) & 63))) | 1))
            _ = fi76
            _ = v78
        }
    } else {
        for fi79 := 0; fi79 < 10; fi79++ {
            var v80 bool = (!false)
            switch (fi79 | -259) {
            case -858:
                for fi81 := 0; fi81 < 4; fi81++ {
                    v80 = (18446744073709551615 >= 2)
                    v82 := GenS3{F0: "a"}
                    v1 = 0
                    _ = fi81
                    _ = v82
                }
            case 1:
                v2 = ((v2 + (float64(0.1) - v2)) - 0.1)
            case 128:
                for fi83 := 0; fi83 < 12; fi83++ {
                    v84 := GenS1{F0: 1000}
                    var v85 uint16 = 42
                    var v86 float32 = 1.5
                    v3 = (^(v3))
                    _ = fi83
                    _ = v84
                    _ = v85
                    _ = v86
                }
            default:
                if false {
                    v2 += v2
                    switch (fi79 + 32768) {
                    case -9223372036854775808:
                        v87 := GenS3{F0: ""}
                        _ = v87
                    }
                    fmt.Println(fi79)
                } else {
                    var v88 bool = ((v1 * (v4 - v1)) == (v4 ^ (v1 << ((659) & 63))))
                    _ = v88
                }
            }
            for fi89 := 0; fi89 < 14; fi89++ {
                var v90 uint16 = 433
                fmt.Println(v2)
                _ = fi89
                _ = v90
            }
            _ = fi79
            _ = v80
        }
    }
    if len(v71) >= 2 {
        v71 = v71[1:len(v71)]
    }
    fmt.Println(v3)
    fmt.Println(v4)
    v71 = append(v71, 2147483647)
    v2 -= v2
    v4 += 196
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v71:", len(v71), v71)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v71
}
