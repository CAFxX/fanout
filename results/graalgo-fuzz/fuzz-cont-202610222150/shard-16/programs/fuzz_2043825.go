package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int8; F1 int32; F2 uint32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint16; F1 int8; F2 string }
type GenS10 struct { F0 int16; F1 uint }
func genFn0() uint16 {
    var v4 int64 = 42
    v4 = -1
    _ = v4
    return 907
}
func genFn1(p0 int32, p1 uint32) int16 {
    switch 9223372036854775807 {
    case -9223372036854775808:
        if ((true || true) && (!false)) {
            if ("\n\t" <= "a") {
                switch "" {
                case "\n\t":
                    var v5 uint16 = 2
                    _ = v5
                case "hello":
                    for fi6 := 0; fi6 < 2; fi6++ {
                        if ((0.1 * (float64(3.14159) * -2.25)) >= -2.25) {
                            v7 := GenS1{F0: 451}
                            _ = v7
                        } else {
                            v8 := GenS1{F0: 18446744073709551615}
                            v9 := GenS2{F0: fi6}
                            _ = v8
                            _ = v9
                        }
                        var v10 uint32 = p1
                        _ = fi6
                        _ = v10
                    }
                default:
                    switch "hello" {
                    case "\n\t":
                        v11 := []int{42, 4294967295, -1, 128}
                        _ = v11
                    case "a":
                        v12 := GenS3{F0: -32768}
                        _ = v12
                    case "hello":
                        var v13 int32 = (p0 / (((p0 | p0)) | 1))
                        _ = v13
                    }
                }
                for fi14 := 0; fi14 < 15; fi14++ {
                    p1 = p1
                    p0 = -2147483648
                    _ = fi14
                }
            }
            if (false && (-393 != 2147483647)) {
                for fi15 := 0; fi15 < 1; fi15++ {
                    for fi16 := 0; fi16 < 14; fi16++ {
                        fmt.Println(fi15, fi15)
                        _ = fi16
                    }
                    _ = fi15
                }
                p1 = (p1 - 2147483647)
            } else {
                var v17 uint16 = 0
                _ = v17
            }
        }
    }
    if (((652 != 0) || (!true)) || false) {
        v18 := []int{-2147483648}
        fmt.Println(p1)
        _ = v18
    }
    p0 += 2147483647
    v19 := []int{-2147483648, 1, 32768}
    _ = p0
    _ = p1
    _ = v19
    return 217
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 int64 = 42
    v20 := GenS3{F0: 1}
    if (-9223372036854775808 != v3) {
        switch "a" {
        case "":
            v3 = -857
        default:
            if false {
                switch (("" + "") + ("foo bar" + "a")) {
                case "a":
                    fmt.Println(v3)
                default:
                    for fi21 := 0; fi21 < 15; fi21++ {
                        var v22 int64 = v3
                        _ = fi21
                        _ = v22
                    }
                }
                v3 = -9223372036854775808
            } else {
                if ((-1 == v3) && (65535 > 42)) {
                    v23 := GenS2{F0: 32767}
                    v24 := []int{9223372036854775807, 4294967295, -30, 270}
                    _ = v23
                    _ = v24
                } else {
                    fmt.Println(v3)
                    if (!((0.1 / 1e10) > float32((0.1 / (-2.25 + (0.0 / -2.25)))))) {
                        v3 = (v3 & v3)
                        if (!(true || false)) {
                            v25 := GenS4{F0: (v3 % ((v3) | 1))}
                            _ = v25
                        }
                        for fi26 := 0; fi26 < 4; fi26++ {
                            switch fi26 {
                            case 0:
                                switch "foo bar" {
                                case "hello":
                                    var v27 int8 = 61
                                    _ = v27
                                case "":
                                    if (true && ((!((0 <= 42) && (!false))) || (((-(fi26)) == (fi26 ^ fi26)) && true))) {
                                        if (!true) {
                                            fmt.Println(fi26, fi26)
                                        } else {
                                            if (genFn1(-2147483648, 2147483647) != 123) {
                                                for fi28 := 0; fi28 < 4; fi28++ {
                                                    fmt.Println(fi26, v3)
                                                    _ = fi28
                                                }
                                                v3 = -9223372036854775808
                                            }
                                            var v29 int16 = int16(v3)
                                            var v30 uint = 395
                                            _ = v29
                                            _ = v30
                                        }
                                    }
                                case "a":
                                    var v31 uint = 18446744073709551615
                                    _ = v31
                                default:
                                    v3 -= (v3 | (v3 & v3))
                                }
                            case 2147483647:
                                fmt.Println(fi26)
                            }
                            _ = fi26
                        }
                    } else {
                        v32 := []int{799, 633, 4294967295}
                        _ = v32
                    }
                }
                var v33 uint64 = 0
                _ = v33
            }
        }
        if false {
            switch "\n\t" {
            case "":
                v34 := []int{128, -719, 1}
                _ = v34
            case "hello":
                v35 := GenS3{F0: 1}
                _ = v35
            }
            var v36 string = (("a" + "foo bar") + ("hello" + "a"))
            _ = v36
        }
        v37 := GenS3{F0: 854}
        _ = v37
    } else {
        v38 := GenS1{F0: 255}
        v39 := GenS4{F0: v3}
        _ = v38
        _ = v39
    }
    if false {
        if false {
            var v40 int16 = 42
            _ = v40
        } else {
            v41 := GenS5{F0: 0}
            _ = v41
        }
        v42 := GenS2{F0: 918}
        _ = v42
    } else {
        if false {
            v3 = v3
        }
        v3 -= (v3 ^ (v3 << ((1) & 63)))
    }
    for fi43 := 0; fi43 < 7; fi43++ {
        v3 -= (v3 << ((65535) & 63))
        _ = fi43
    }
    v44 := []int{0, 597}
    var v45 int = -529
    var v46 uint32 = 1
    for fi47 := 0; fi47 < 3; fi47++ {
        var v48 int64 = (v3 & v3)
        for fi49 := 0; fi49 < 10; fi49++ {
            if (!(1.5 > (1e10 - (float32(1.5) + (3.14159 - (0.1 + float32(float64((-2.25 / 1.5))))))))) {
                v50 := v2[(("" + ("" + ("hello" + "foo bar"))) + "\n\t")]
                _ = v50
            }
            _ = fi49
        }
        v45 -= 4294967295
        _ = fi47
        _ = v48
    }
    var v51 float32 = -2.25
    for fi52 := 0; fi52 < 8; fi52++ {
        var v53 int = 32768
        _ = fi52
        _ = v53
    }
    var v54 uint32 = v46
    switch (v45 ^ v45) {
    case -2147483648:
        for fi55 := 0; fi55 < 20; fi55++ {
            for fi56 := 0; fi56 < 12; fi56++ {
                v57 := GenS4{F0: (v3 & 1)}
                delete(v2, ("hello" + "a"))
                switch ("\n\t" + "\n\t") {
                case "hello":
                    var v58 bool = (-1 <= -1)
                    _ = v58
                case "\n\t":
                    fmt.Println(v45, v51, fi56)
                }
                if (v54 >= 42) {
                    var v59 uint = 344
                    if false {
                        var v60 int32 = -1
                        _ = v60
                    } else {
                        v45 += -111
                    }
                    _ = len(v1)
                    _ = v59
                }
                if ("foo bar" != "hello") {
                    v54 -= (v54 & 2)
                    var v61 int64 = (v3 | (v3 + v3))
                    v51 -= v51
                    _ = v61
                } else {
                    v45 += fi56
                    if (!(v51 > ((v51 - ((float32((0.1 / (-2.25 + 0.0))) - 3.14159) + v51)) - -2.25))) {
                        fmt.Println(fi56, v51, fi55)
                    } else {
                        switch (fi55 * v45) {
                        case 0:
                            fmt.Println(v54, v54, v45)
                        case -299:
                            if ((fi56 - (-(fi56))) <= fi56) {
                                v62 := GenS6{F0: 0}
                                switch "" {
                                case "a":
                                    v3 += v3
                                }
                                var v63 uint16 = 196
                                _ = v62
                                _ = v63
                            } else {
                                v54 = 549
                                v64 := GenS3{F0: 32767}
                                switch fi55 {
                                case 0:
                                    _ = len(v1)
                                case 1:
                                    switch 1 {
                                    case -9223372036854775808:
                                        if true {
                                            var v65 float32 = (v51 + (v51 + (0.1 / (v51 * v51))))
                                            v66 := GenS3{F0: 719}
                                            if false {
                                                fmt.Println(v54)
                                                if true {
                                                    v3 -= (v3 & v3)
                                                    v67 := GenS7{F0: -66, F1: -1, F2: (v54 | (v46 - v54))}
                                                    _ = v67
                                                }
                                                v3 -= v3
                                                delete(v2, "foo bar")
                                                v51 -= (1e10 * 0.0)
                                                v51 -= (v65 / (1.5 - v65))
                                                v46 += v46
                                            } else {
                                                var v68 float32 = v51
                                                _ = v68
                                            }
                                            _ = v65
                                            _ = v66
                                        }
                                    case -18:
                                        v69 := GenS6{F0: 353}
                                        _ = v69
                                    default:
                                        v70 := GenS8{F0: ("" + ("" + (("foo bar" + "") + "hello")))}
                                        _ = v70
                                    }
                                case 2147483647:
                                    if true {
                                        var v71 uint16 = 18
                                        v72 := GenS6{F0: 0}
                                        switch "\n\t" {
                                        case "hello":
                                            v45 += (fi55 ^ 32768)
                                        case "":
                                            var v73 bool = false
                                            _ = v73
                                        default:
                                            if true {
                                                var v74 uint32 = (v54 - 4294967295)
                                                var v75 float32 = v51
                                                _ = v74
                                                _ = v75
                                            }
                                        }
                                        _ = v71
                                        _ = v72
                                    }
                                }
                                _ = v64
                            }
                        }
                    }
                    if (("a" + ("a" + ("a" + "\n\t"))) == "foo bar") {
                        v46 = (v54 << ((42) & 63))
                        v3 = v3
                        var v76 uint16 = 0
                        _ = v76
                    } else {
                        v46 += v54
                    }
                    var v77 uint64 = 42
                    if false {
                        if true {
                            var v78 uint = 2
                            _ = v78
                        } else {
                            var v79 uint8 = 0
                            if false {
                                v80 := GenS4{F0: v3}
                                v46 = (^(v46))
                                var v81 bool = false
                                _ = v80
                                _ = v81
                            }
                            _ = v79
                        }
                        _ = len(v44)
                        if false {
                            _ = len(v44)
                            v82 := GenS9{F0: 1, F1: -88, F2: "a"}
                            _ = v82
                        } else {
                            fmt.Println(v51, v51, fi55)
                        }
                    }
                    if false {
                        var v83 uint = 65535
                        _ = v83
                    } else {
                        v84 := GenS10{F0: 1000, F1: 2}
                        var v85 uint = 65535
                        var v86 int16 = 1000
                        _ = v84
                        _ = v85
                        _ = v86
                    }
                    v3 -= (^(v3))
                    switch fi56 {
                    case -2147483648:
                        switch fi55 {
                        case 122:
                            if true {
                                var v87 float64 = ((-2.25 - 1e10) * ((3.14159 + 0.0) * (float64(0.1) * float64(v51))))
                                var v88 int8 = -1
                                fmt.Println(v77)
                                _ = v87
                                _ = v88
                            } else {
                                v77 -= (v77 - v77)
                                var v89 int32 = -301
                                _ = v89
                            }
                        }
                    case 300:
                        v46 = (v46 >> ((42) & 63))
                    }
                    var v90 int = -2147483648
                    var v91 uint32 = v54
                    var v92 int16 = -979
                    v93 := GenS2{F0: v45}
                    if false {
                        v77 += v77
                    }
                    v90 = (v90 * (-(fi55)))
                    _ = v77
                    _ = v90
                    _ = v91
                    _ = v92
                    _ = v93
                }
                _ = fi56
                _ = v57
            }
            _ = fi55
        }
    case -9223372036854775808:
        var v94 int = v45
        _ = v94
    }
    v95 := GenS1{F0: 308}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v51:", v51)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v44
    _ = v45
    _ = v46
    _ = v51
    _ = v54
    _ = v95
}
