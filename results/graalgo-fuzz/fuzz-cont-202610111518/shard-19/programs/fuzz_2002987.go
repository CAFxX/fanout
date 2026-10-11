package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 string; F1 int; F2 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 string; F1 uint8 }
type GenS10 struct { F0 string; F1 int }
func genFn0(p0 int16, p1 int64) uint {
    if (794 <= 0) {
        p0 = (p0 % (((p0 & (p0 - 1))) | 1))
    } else {
        v5 := GenS1{F0: 540}
        _ = v5
    }
    var v6 int16 = int16(p1)
    fmt.Println(p0)
    switch "foo bar" {
    case "foo bar":
        if (-2147483648 >= 2147483647) {
            v7 := GenS1{F0: 740}
            _ = v7
        } else {
            switch "\n\t" {
            case "a":
                switch "foo bar" {
                case "hello":
                    v8 := GenS1{F0: 42}
                    _ = v8
                }
            case "hello":
                var v9 int32 = 42
                _ = v9
            default:
                for fi10 := 0; fi10 < 2; fi10++ {
                    v11 := GenS2{F0: 65535, F1: (^(p0))}
                    switch "\n\t" {
                    case "foo bar":
                        v12 := []int{fi10, (fi10 / ((fi10) | 1)), (fi10 >> ((65535) & 63)), (fi10 | 513)}
                        _ = v12
                    default:
                        v6 = (-(v6))
                    }
                    v6 = v6
                    _ = fi10
                    _ = v11
                }
            }
        }
    case "\n\t":
        switch 0 {
        case -2147483648:
            var v13 uint32 = 443
            _ = v13
        case 0:
            p0 = (v6 - (v6 / ((p0) | 1)))
        case 128:
            fmt.Println(p1, p1, p1)
        default:
            switch "\n\t" {
            case "\n\t":
                if true {
                    v14 := GenS1{F0: 42}
                    v15 := []int{-646, 128, 32767, 127}
                    v16 := GenS3{F0: (("hello" + "") + ("hello" + "")), F1: -9223372036854775808, F2: 260}
                    _ = v14
                    _ = v15
                    _ = v16
                } else {
                    switch 0 {
                    case -303:
                        for fi17 := 0; fi17 < 1; fi17++ {
                            for fi18 := 0; fi18 < 10; fi18++ {
                                p0 = -1
                                if true {
                                    fmt.Println(fi17, fi18)
                                    var v19 float64 = 0.0
                                    fmt.Println(p0, fi18)
                                    _ = v19
                                } else {
                                    p1 += p1
                                    v6 += v6
                                    fmt.Println(p1)
                                    fmt.Println(fi18, p0)
                                }
                                _ = fi18
                            }
                            var v20 float64 = -2.25
                            _ = fi17
                            _ = v20
                        }
                    case 127:
                        var v21 int8 = 0
                        _ = v21
                    default:
                        if false {
                            if (878 >= 808) {
                                for fi22 := 0; fi22 < 7; fi22++ {
                                    v23 := GenS2{F0: 2, F1: 42}
                                    var v24 uint16 = 397
                                    _ = fi22
                                    _ = v23
                                    _ = v24
                                }
                            }
                            for fi25 := 0; fi25 < 3; fi25++ {
                                var v26 int16 = v6
                                var v27 float32 = (1.5 + (-2.25 + (1e10 - 0.1)))
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                            for fi28 := 0; fi28 < 8; fi28++ {
                                if (2147483647 <= 2147483648) {
                                    v29 := GenS4{F0: 81}
                                    _ = v29
                                }
                                for fi30 := 0; fi30 < 1; fi30++ {
                                    switch 32767 {
                                    case 936:
                                        var v31 bool = ((255 != 159) || ((("hello" + "foo bar") != "foo bar") && (float64(-2.25) == (-2.25 * 1e10))))
                                        _ = v31
                                    case -1:
                                        fmt.Println(fi30)
                                    }
                                    fmt.Println(p0)
                                    v32 := []int{fi28, 127, fi30}
                                    _ = fi30
                                    _ = v32
                                }
                                _ = fi28
                            }
                        }
                    }
                    for fi33 := 0; fi33 < 7; fi33++ {
                        for fi34 := 0; fi34 < 20; fi34++ {
                            fmt.Println(fi34, fi33)
                            _ = fi34
                        }
                        for fi35 := 0; fi35 < 3; fi35++ {
                            var v36 string = "\n\t"
                            if false {
                                if false {
                                    v37 := GenS4{F0: 545}
                                    _ = v37
                                } else {
                                    v36 = v36
                                }
                            }
                            _ = fi35
                            _ = v36
                        }
                        _ = fi33
                    }
                }
            case "hello":
                p0 = p0
            case "":
                if (true || false) {
                    v38 := []int{-2147483648}
                    _ = v38
                }
            }
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return 1
}
func genFn1(p0 uint64, p1 int, p2 uint32) uint8 {
    var v39 float64 = (float64(((3.14159 + 1.5) * (-2.25 / 3.14159))) + 0.0)
    fmt.Println(p1, v39)
    v39 = v39
    _ = p0
    _ = p1
    _ = p2
    _ = v39
    return 255
}

func main() {
    v1 := []uint32{}
    var v2 uint = 0
    var v3 uint64 = 2
    v4 := []string{("a" + ("\n\t" + "foo bar")), "a"}
    v2 = v2
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    v40 := GenS4{F0: 619}
    v1 = append(v1, 2147483647)
    switch ("\n\t" + ("hello" + "foo bar")) {
    case "":
        var v41 uint16 = 1
        _ = v41
    case "hello":
        var v42 int16 = 42
        _ = v42
    case "foo bar":
        v4 = append(v4, ("hello" + "a"))
    }
    var v43 int64 = -9223372036854775808
    v44 := GenS5{F0: 255, F1: 501, F2: 42}
    for fi45 := 0; fi45 < 3; fi45++ {
        var v46 uint16 = 878
        fmt.Println(fi45, v46)
        _ = fi45
        _ = v46
    }
    v47 := GenS2{F0: 1, F1: 1}
    for fi48 := 0; fi48 < 20; fi48++ {
        v3 = v3
        if (42 != v3) {
            v3 += 872
            if false {
                if (false || true) {
                    switch "" {
                    case "":
                        if len(v4) > 0 {
                            v4[0] = ("a" + ("a" + ("a" + (("\n\t" + "") + "a"))))
                        }
                    }
                } else {
                    v49 := GenS6{F0: ""}
                    v50 := GenS3{F0: (("\n\t" + "a") + ("\n\t" + ("\n\t" + ("hello" + "")))), F1: 9223372036854775807, F2: (v3 * (v3 - v3))}
                    _ = v49
                    _ = v50
                }
                v51 := GenS5{F0: genFn1(v3, (fi48 << ((v2) & 63)), 42), F1: 1000, F2: 65535}
                if false {
                    var v52 uint32 = 0
                    _ = v52
                }
                _ = v51
            }
            if false {
                if (1.5 == 3.14159) {
                    v53 := GenS7{F0: (v3 & 42), F1: 0, F2: 471}
                    _ = v53
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            } else {
                var v54 uint8 = 37
                _ = v54
            }
        } else {
            if (694 >= 828) {
                v2 = genFn0(262, v43)
                for fi55 := 0; fi55 < 4; fi55++ {
                    switch "" {
                    case "\n\t":
                        v3 += v3
                    case "foo bar":
                        if (("foo bar" + ((("hello" + "a") + "a") + (("\n\t" + "a") + "\n\t"))) >= "a") {
                            if true {
                                v2 -= (^(v2))
                                v43 += v43
                                v56 := GenS3{F0: "hello", F1: fi48, F2: 42}
                                _ = v56
                            } else {
                                if len(v4) > 0 {
                                    v4[0] = (("" + ("\n\t" + "foo bar")) + "foo bar")
                                }
                                switch "\n\t" {
                                case "a":
                                    switch "foo bar" {
                                    case "foo bar":
                                        switch (fi55 / (((fi55 * 2147483647)) | 1)) {
                                        case -1:
                                            v2 -= (v2 & 760)
                                        case 32767:
                                            v43 = v43
                                        }
                                    case "\n\t":
                                        if false {
                                            if ((1 != 2) || true) {
                                                v43 = -9223372036854775808
                                                var v57 int32 = 1
                                                _ = v57
                                            }
                                            switch (("\n\t" + "\n\t") + "foo bar") {
                                            case "a":
                                                fmt.Println(v3, v3, v3)
                                            case "foo bar":
                                                var v58 int32 = 424
                                                _ = v58
                                            case "\n\t":
                                                v1 = append(v1, 1)
                                            }
                                        }
                                        var v59 uint = v2
                                        _ = v59
                                    case "hello":
                                        if (!(1 == -306)) {
                                            var v60 bool = false
                                            if (!(255 < uint(v3))) {
                                                var v61 uint8 = 2
                                                v62 := GenS5{F0: 1, F1: -1, F2: 42}
                                                v43 = (v43 * 42)
                                                _ = v61
                                                _ = v62
                                            } else {
                                                var v63 int8 = 42
                                                _ = v63
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = v60
                                        } else {
                                            if ((1.5 - 0.1) < float32((float64(((0.0 - 1e10) + float32(float64(float32(0.1))))) - (3.14159 - ((1.5 / ((0.1 / 1.0) + 1.5)) - (1.5 + 0.1)))))) {
                                                v64 := GenS8{F0: (^(v2))}
                                                _ = v64
                                            } else {
                                                v43 += v43
                                                v3 += (v3 | (v3 >> (((v2 ^ v2)) & 63)))
                                                if (v3 != v3) {
                                                    var v65 uint8 = 0
                                                    v3 += (v3 ^ v3)
                                                    fmt.Println(v43)
                                                    fmt.Println(v3, v65, v43)
                                                    _ = v65
                                                } else {
                                                    if (false || false) {
                                                        var v66 float32 = (3.14159 + ((-2.25 + 3.14159) + 3.14159))
                                                        var v67 uint8 = genFn1(v3, 1, 89)
                                                        if (0 >= (^(fi48))) {
                                                            if true {
                                                                v2 = v2
                                                                fmt.Println(v3, v43)
                                                                var v68 int16 = 32767
                                                                _ = v68
                                                            }
                                                            if true {
                                                                _ = len(v1)
                                                                v69 := GenS9{F0: "", F1: genFn1((v3 | v3), -9223372036854775808, 602)}
                                                                switch ("\n\t" + "") {
                                                                case "hello":
                                                                    v66 = v66
                                                                case "foo bar":
                                                                    v70 := GenS10{F0: "foo bar", F1: fi55}
                                                                    _ = v70
                                                                case "\n\t":
                                                                    var v71 uint64 = v3
                                                                    _ = v71
                                                                }
                                                                _ = v69
                                                            }
                                                        } else {
                                                            v2 = v2
                                                            var v72 int8 = 0
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            _ = v72
                                                        }
                                                        _ = v66
                                                        _ = v67
                                                    } else {
                                                        _ = len(v4)
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        var v73 uint = (^(v2))
                                                        v73 = v73
                                                        _ = v73
                                                    }
                                                    if true {
                                                        v3 -= v3
                                                        v3 = uint64(v2)
                                                        v3 -= 600
                                                    } else {
                                                        v74 := GenS10{F0: (("\n\t" + ("a" + "a")) + "a"), F1: (-(fi48))}
                                                        _ = v74
                                                    }
                                                }
                                            }
                                            if false {
                                                var v75 float64 = (1e10 * 0.1)
                                                _ = v75
                                            }
                                        }
                                    }
                                case "\n\t":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                if ((v3 << ((v2) & 63)) <= v3) {
                                    v76 := GenS4{F0: 42}
                                    if false {
                                        var v77 uint16 = 65535
                                        _ = v77
                                    } else {
                                        v2 = (v2 * 838)
                                        var v78 float64 = 1.5
                                        v79 := GenS3{F0: (("" + ("" + "hello")) + "\n\t"), F1: (-(fi55)), F2: 4294967295}
                                        _ = v78
                                        _ = v79
                                    }
                                    _ = v76
                                } else {
                                    v2 -= v2
                                }
                            }
                            _ = len(v4)
                        } else {
                            v3 = (v3 / (((v3 % (((v3 * (v3 - (v3 ^ 2)))) | 1))) | 1))
                            v80 := GenS9{F0: "\n\t", F1: 1}
                            v81 := GenS4{F0: 2147483648}
                            _ = v80
                            _ = v81
                        }
                    case "a":
                        v43 -= 42
                    }
                    var v82 float64 = -2.25
                    v83 := GenS3{F0: "a", F1: -976, F2: v3}
                    _ = fi55
                    _ = v82
                    _ = v83
                }
                var v84 uint16 = 461
                _ = v84
            } else {
                if (false || true) {
                    v3 = v3
                    var v85 float32 = 3.14159
                    _ = v85
                } else {
                    v2 -= (v2 + v2)
                }
                switch 32767 {
                case 2147483647:
                    var v86 int8 = 0
                    _ = v86
                case 1:
                    for fi87 := 0; fi87 < 14; fi87++ {
                        v43 -= (v43 ^ (v43 >> (((^(v2))) & 63)))
                        v2 = (v2 << ((v2) & 63))
                        var v88 bool = (!(-1 >= -1))
                        _ = fi87
                        _ = v88
                    }
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                v2 += (v2 % ((v2) | 1))
            }
            if true {
                switch ("\n\t" + ((("foo bar" + ("\n\t" + "foo bar")) + "a") + "\n\t")) {
                case "":
                    var v89 uint64 = 1
                    _ = v89
                case "a":
                    fmt.Println(v43)
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v90 uint64 = 4294967295
                _ = v90
            }
            v91 := GenS5{F0: 2, F1: 1, F2: 65535}
            _ = v91
        }
        _ = fi48
    }
    _ = len(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v43
    _ = v44
    _ = v47
}
