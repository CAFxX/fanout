package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 uint; F2 int64 }
type GenS5 struct { F0 int64; F1 int; F2 int16 }
func genFn0(p0 string, p1 int16, p2 int32) uint8 {
    v4 := []int{}
    var v5 uint = 65535
    for fi6 := 0; fi6 < 16; fi6++ {
        if (-9223372036854775808 < -1) {
            fmt.Println(v5, v5, p0)
            fmt.Println(fi6, v5, p0)
        }
        switch int(v5) {
        case -685:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        if false {
            switch fi6 {
            case -34:
                var v7 uint16 = 32767
                _ = v7
            case 32767:
                var v8 uint = (v5 + 971)
                _ = v8
            default:
                var v9 uint32 = 2147483647
                _ = v9
            }
            var v10 uint8 = 0
            _ = v10
        }
        _ = fi6
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 191
}
func genFn1(p0 int32) uint16 {
    v11 := GenS1{F0: 42, F1: 127}
    var v12 uint32 = 304
    var v13 int8 = -32
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 255
}

func main() {
    var v1 int16 = 350
    var v2 float32 = 0.1
    var v3 uint = 448
    var v14 uint16 = 397
    v15 := GenS2{F0: -9223372036854775808, F1: ("hello" + ("foo bar" + (("\n\t" + ("a" + "")) + (((("" + "\n\t") + "\n\t") + ("\n\t" + "\n\t")) + "foo bar"))))}
    v1 += 771
    var v16 uint64 = 104
    if (!(-128 >= 1)) {
        v16 = 875
        fmt.Println(v1, v2)
        var v17 bool = false
        _ = v17
    } else {
        var v18 int32 = 2147483647
        _ = v18
    }
    v1 += (v1 + 0)
    for fi19 := 0; fi19 < 12; fi19++ {
        switch "hello" {
        case "hello":
            v20 := GenS2{F0: -349, F1: "\n\t"}
            _ = v20
        default:
            if false {
                v21 := GenS2{F0: -315, F1: "hello"}
                _ = v21
            } else {
                if (true || false) {
                    v22 := GenS1{F0: (v16 << (((v3 / (((v3 - v3)) | 1))) & 63)), F1: 42}
                    if true {
                        if ((v1 & (v1 - v1)) < -1) {
                            for fi23 := 0; fi23 < 13; fi23++ {
                                var v24 float64 = ((3.14159 * 0.0) - (3.14159 - 0.0))
                                var v25 int32 = 0
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                            var v26 float64 = 3.14159
                            v27 := []int{0}
                            _ = v26
                            _ = v27
                        }
                        v14 = 32767
                    } else {
                        for fi28 := 0; fi28 < 19; fi28++ {
                            if (!false) {
                                v14 = v14
                            } else {
                                fmt.Println(fi28, v16, v14)
                            }
                            _ = fi28
                        }
                        for fi29 := 0; fi29 < 1; fi29++ {
                            v14 -= (v14 + v14)
                            v1 = (v1 ^ (v1 | 0))
                            _ = fi29
                        }
                    }
                    _ = v22
                }
                if (("a" + "foo bar") <= "") {
                    v16 = (v16 & 1)
                    var v30 int64 = -646
                    for fi31 := 0; fi31 < 7; fi31++ {
                        if (true || (!true)) {
                            v2 = 0.1
                            if true {
                                v32 := []int{(fi19 & fi19)}
                                var v33 uint8 = genFn0("foo bar", (^(v1)), 42)
                                _ = v32
                                _ = v33
                            }
                            if (false || true) {
                                v1 = (v1 - v1)
                                v34 := []int{fi31}
                                _ = v34
                            }
                            v35 := []int{fi31, -201, (fi31 >> (((v3 << ((v3) & 63))) & 63)), (fi19 - fi31)}
                            v16 -= (^(v16))
                            fmt.Println(fi31, fi19)
                            if ((464 != (v14 / (((v14 * genFn1(-502))) | 1))) && (false && (0.0 <= 0.0))) {
                                var v36 int32 = 1
                                if true {
                                    var v37 uint8 = 42
                                    _ = v37
                                }
                                _ = v36
                            }
                            if len(v35) >= 2 {
                                v35 = v35[1:len(v35)]
                            }
                            _ = v35
                        } else {
                            v1 = v1
                            if (((^(v3)) <= (v3 + v3)) || true) {
                                var v38 uint32 = 174
                                if (!false) {
                                    v39 := []int{(^(fi31)), fi19, (fi19 | fi31), 32767}
                                    v3 = 49
                                    fmt.Println(v14, v1)
                                    fmt.Println(v3, v14)
                                    var v40 uint64 = 2
                                    switch (fi31 / ((fi31) | 1)) {
                                    case 32767:
                                        fmt.Println(v16, v38)
                                    case 9223372036854775807:
                                        fmt.Println(v30, v14)
                                    case -1:
                                        var v41 uint8 = 42
                                        _ = v41
                                    }
                                    _ = v39
                                    _ = v40
                                } else {
                                    switch ("hello" + "") {
                                    case "a":
                                        v42 := GenS3{F0: (v3 | v3)}
                                        _ = v42
                                    case "foo bar":
                                        var v43 int32 = -518
                                        _ = v43
                                    }
                                    if ((^(v16)) <= v16) {
                                        fmt.Println(v14)
                                        var v44 int = (^(fi31))
                                        _ = v44
                                    } else {
                                        if true {
                                            if ((v16 >> (((v3 + v3)) & 63)) < v16) {
                                                var v45 int64 = 1
                                                v46 := []int{fi19}
                                                if (v3 > (v3 >> ((v3) & 63))) {
                                                    var v47 uint8 = 1
                                                    _ = v47
                                                } else {
                                                    v48 := GenS4{F0: v38, F1: 415, F2: v30}
                                                    _ = v48
                                                }
                                                _ = v45
                                                _ = v46
                                            }
                                            if true {
                                                v49 := GenS1{F0: (v16 + 0), F1: -1}
                                                fmt.Println(v38)
                                                _ = v49
                                            }
                                        }
                                        var v50 string = ("" + "a")
                                        switch fi19 {
                                        case 2147483647:
                                            v16 += 18446744073709551615
                                        case 4294967295:
                                            switch (fi31 % (((fi31 ^ (^(fi19)))) | 1)) {
                                            case 2147483647:
                                                v51 := GenS5{F0: 9223372036854775807, F1: 9223372036854775807, F2: 0}
                                                _ = v51
                                            default:
                                                if (0.1 <= 0.0) {
                                                    var v52 float64 = -2.25
                                                    var v53 int8 = 0
                                                    _ = v52
                                                    _ = v53
                                                } else {
                                                    v30 += (v30 & v30)
                                                }
                                            }
                                        default:
                                            v14 = (v14 * (v14 << ((v3) & 63)))
                                        }
                                        _ = v50
                                    }
                                }
                                _ = v38
                            } else {
                                var v54 int64 = v30
                                _ = v54
                            }
                            if true {
                                v55 := GenS1{F0: v16, F1: -128}
                                _ = v55
                            }
                        }
                        switch (("a" + "foo bar") + "") {
                        case "foo bar":
                            v2 = 1e10
                            v56 := []int{fi19, (-(fi31)), (fi19 + (^(fi19))), int(v16)}
                            if false {
                                var v57 bool = (15 != (v3 + v3))
                                if len(v56) > 0 {
                                    v56[0] = (-(fi19))
                                }
                                _ = v57
                            } else {
                                var v58 int16 = 943
                                if (v16 == (^(v16))) {
                                    v2 = 1e10
                                }
                                _ = v58
                            }
                            _ = v56
                        }
                        _ = fi31
                    }
                    _ = v30
                } else {
                    if false {
                        v59 := GenS2{F0: 42, F1: "foo bar"}
                        v60 := []int{fi19, fi19, (fi19 & fi19), fi19}
                        _ = v59
                        _ = v60
                    } else {
                        var v61 bool = ((v1 & 42) < 1000)
                        if (!v61) {
                            if (1 > int64(v1)) {
                                if (!false) {
                                    var v62 uint32 = 2
                                    var v63 int8 = 13
                                    _ = v62
                                    _ = v63
                                }
                            } else {
                                var v64 uint16 = 728
                                _ = v64
                            }
                            var v65 int8 = 1
                            _ = v65
                        }
                        for fi66 := 0; fi66 < 11; fi66++ {
                            var v67 string = (("foo bar" + "\n\t") + "\n\t")
                            _ = fi66
                            _ = v67
                        }
                        _ = v61
                    }
                    var v68 uint = (v3 >> ((65535) & 63))
                    switch "\n\t" {
                    case "\n\t":
                        v69 := []int{(fi19 * (fi19 + (^(fi19)))), (fi19 / ((4294967295) | 1))}
                        _ = v69
                    default:
                        for fi70 := 0; fi70 < 14; fi70++ {
                            v3 += v68
                            v71 := []int{(fi19 & fi19), fi19}
                            _ = fi70
                            _ = v71
                        }
                    }
                    _ = v68
                }
                fmt.Println(v3, v16, v3)
            }
        }
        v72 := []int{}
        _ = fi19
        _ = v72
    }
    fmt.Println(v14)
    v73 := []int{}
    fmt.Println(v1, v3, v14)
    var v74 float64 = 3.14159
    if (false && true) {
        var v75 uint = (v3 ^ (^(v3)))
        if (v1 > v1) {
            v16 += 774
        }
        _ = v75
    }
    var v76 float32 = float32(v74)
    var v77 int32 = -2147483648
    fmt.Println(v74)
    v77 = 137
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v73:", len(v73), v73)
    fmt.Println("v74:", v74)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v73
    _ = v74
    _ = v76
    _ = v77
}
