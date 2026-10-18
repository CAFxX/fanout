package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 int; F2 uint8 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS7 struct { F0 int16; F1 uint64 }
func genFn0(p0 int32, p1 int8) int64 {
    fmt.Println(p1, p0)
    var v4 int8 = 11
    if true {
        p1 += v4
    }
    _ = p0
    _ = p1
    _ = v4
    return 9223372036854775807
}

func main() {
    var v1 bool = true
    v2 := make(map[string]string)
    var v3 uint = 0
    v5 := []int{}
    for fi6 := 0; fi6 < 20; fi6++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi6
    }
    v1 = (127 == 42)
    if v1 {
        var v7 uint64 = 1
        _ = v7
    }
    v3 -= v3
    fmt.Println(v3)
    if v1 {
        var v8 uint32 = 42
        fmt.Println(v1, v1, v8)
        _ = v8
    }
    var v9 int8 = 1
    v9 = v9
    v10 := GenS1{F0: 2147483647}
    for fi11 := 0; fi11 < 10; fi11++ {
        if (!(!v1)) {
            var v12 float64 = 0.0
            var v13 int16 = 32767
            _ = v12
            _ = v13
        }
        if ("" != "\n\t") {
            v14 := GenS1{F0: 42}
            _ = len(v2)
            _ = v14
        }
        _ = fi11
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        for fi16 := 0; fi16 < 4; fi16++ {
            switch "\n\t" {
            case "foo bar":
                v17 := GenS1{F0: -2147483648}
                v5 = append(v5, fi15)
                _ = v17
            }
            var v18 uint32 = 65535
            if (42 <= v9) {
                if true {
                    v19 := GenS1{F0: -2147483648}
                    v3 = v3
                    v5 = append(v5, fi15)
                    var v20 string = "foo bar"
                    switch fi15 {
                    case -2147483648:
                        if v1 {
                            v5 = append(v5, fi15)
                            var v21 string = v20
                            if true {
                                _ = len(v2)
                                v20 = ("" + v21)
                                v9 -= v9
                                v1 = v1
                                var v22 uint16 = 42
                                _ = v22
                            }
                            _ = v21
                        }
                        v9 = 42
                        if false {
                            var v23 string = ((v20 + v20) + v20)
                            _ = v23
                        } else {
                            var v24 bool = (90 <= 0)
                            v18 = (v18 - (v18 << ((v3) & 63)))
                            if v24 {
                                fmt.Println(v3, v20, v9)
                                v9 += int8(fi15)
                            }
                            _ = v24
                        }
                    case -54:
                        switch "" {
                        case "":
                            v9 = (v9 / (((v9 + v9)) | 1))
                            v18 = 2147483647
                            if true {
                                fmt.Println(v9, fi15)
                                var v25 int32 = 461
                                _ = v25
                            }
                            if true {
                                v26 := GenS1{F0: 42}
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = v26
                            } else {
                                fmt.Println(v1)
                                v27 := v2[((v20 + v20) + v20)]
                                _ = v27
                            }
                        case "hello":
                            if v1 {
                                var v28 int16 = 898
                                var v29 int64 = genFn0(-710, 76)
                                v18 = (v18 + 0)
                                v30 := GenS2{F0: (("foo bar" + (v20 + (v20 + ((v20 + v20) + v20)))) > v20), F1: 127, F2: 1}
                                _ = v28
                                _ = v29
                                _ = v30
                            } else {
                                if v1 {
                                    if v1 {
                                        v9 -= (v9 >> ((v3) & 63))
                                    }
                                    v3 -= v3
                                    v5 = append(v5, 781)
                                }
                                if (2147483648 != v18) {
                                    var v31 float32 = 1e10
                                    _ = len(v5)
                                    v9 = int8(v3)
                                    _ = v31
                                }
                            }
                        case "foo bar":
                            var v32 uint16 = uint16(v3)
                            _ = v32
                        }
                    default:
                        if v1 {
                            v3 += (v3 ^ v3)
                        } else {
                            var v33 int = fi16
                            var v34 uint = v3
                            if (false || false) {
                                var v35 int64 = -27
                                _ = v35
                            }
                            if true {
                                v2[v20] = ((("foo bar" + (v20 + v20)) + (v20 + ((v20 + "") + "foo bar"))) + (v20 + v20))
                                v36 := GenS3{F0: 717, F1: "foo bar"}
                                v37 := GenS4{F0: 306, F1: 127}
                                if true {
                                    var v38 uint = v34
                                    v18 = v18
                                    v39 := GenS5{F0: fi15, F1: -1}
                                    _ = v38
                                    _ = v39
                                } else {
                                    if (v1 && (!true)) {
                                        fmt.Println(v9)
                                    } else {
                                        v20 = ((v20 + v20) + v20)
                                        if len(v5) > 0 {
                                            v5[0] = 127
                                        }
                                        var v40 float32 = float32((0.1 - 1.5))
                                        _ = v40
                                    }
                                }
                                _ = v36
                                _ = v37
                            }
                            var v41 bool = v1
                            _ = v33
                            _ = v34
                            _ = v41
                        }
                        var v42 int32 = 2147483647
                        _ = v42
                    }
                    var v43 bool = (!(1 < 32767))
                    _ = v19
                    _ = v20
                    _ = v43
                }
                if (!v1) {
                    var v44 string = ("a" + "")
                    fmt.Println(v9, v1, v44)
                    fmt.Println(v9, v3)
                    _ = v44
                }
            }
            _ = fi16
            _ = v18
        }
        var v45 uint8 = 255
        v2[(("foo bar" + "\n\t") + "\n\t")] = ("a" + "\n\t")
        _ = fi15
        _ = v45
    }
    v46 := GenS6{F0: 65535, F1: 1, F2: int8(v3)}
    v1 = (809 <= 32767)
    var v47 int64 = -9223372036854775808
    for fi48 := 0; fi48 < 3; fi48++ {
        switch "\n\t" {
        case "hello":
            var v49 int64 = (-(v47))
            _ = v49
        case "":
            v1 = (!((float32((-2.25 * ((0.0 * 0.1) - 0.0))) + 0.0) <= 1e10))
        case "\n\t":
            for fi50 := 0; fi50 < 2; fi50++ {
                v5 = append(v5, 4294967295)
                v3 = 1
                _ = fi50
            }
        default:
            fmt.Println(v1)
        }
        for fi51 := 0; fi51 < 4; fi51++ {
            if false {
                v52 := GenS3{F0: 1, F1: "foo bar"}
                _ = v52
            }
            _ = fi51
        }
        var v53 bool = true
        _ = fi48
        _ = v53
    }
    var v54 uint32 = 4294967295
    if (v1 && v1) {
        v3 = 42
    } else {
        var v55 uint = v3
        if true {
            fmt.Println(v55)
            for fi56 := 0; fi56 < 16; fi56++ {
                var v57 float64 = (((0.1 * (-2.25 + (((1.5 - 1e10) / ((1e10 / 1.0) + float64(-2.25))) * 1.5))) - 1.5) / 1.0)
                if (1 != 90) {
                    if v1 {
                        if (false && (!v1)) {
                            v9 = -123
                            for fi58 := 0; fi58 < 2; fi58++ {
                                var v59 int64 = v47
                                if (true || ((true || (((0.1 * (0.0 - 0.1)) - 0.0) <= (0.0 * 3.14159))) && true)) {
                                    if ((fi56 / ((-9223372036854775808) | 1)) < fi58) {
                                        switch (fi56 >> ((v55) & 63)) {
                                        case 970:
                                            v1 = (18446744073709551615 != (v55 * 517))
                                        case 32768:
                                            switch (fi56 % ((fi56) | 1)) {
                                            case -2147483648:
                                                var v60 string = "a"
                                                _ = v60
                                            case 1:
                                                v1 = false
                                            case 32767:
                                                v61 := GenS3{F0: 42, F1: "hello"}
                                                _ = v61
                                            }
                                        case 127:
                                            switch 42 {
                                            case -2147483648:
                                                var v62 uint32 = (^(v54))
                                                _ = v62
                                            case 127:
                                                v63 := GenS4{F0: 4294967295, F1: 0}
                                                _ = v63
                                            }
                                        }
                                        var v64 int32 = -896
                                        v3 -= v55
                                        _ = v64
                                    }
                                    if true {
                                        var v65 int = (fi58 % ((fi56) | 1))
                                        if (((v1 || v1) && false) && v1) {
                                            v3 = v55
                                            var v66 string = ((("hello" + "foo bar") + ("foo bar" + "foo bar")) + ("a" + ("a" + ("hello" + ("" + "hello")))))
                                            v54 -= v54
                                            _ = v66
                                        } else {
                                            fmt.Println(fi58)
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            var v67 float32 = float32(v57)
                                            v68 := v2[("" + "")]
                                            fmt.Println(v54)
                                            v59 -= (v59 ^ (v59 << (((v55 ^ (^(v3)))) & 63)))
                                            delete(v2, ((v68 + (v68 + ("" + v68))) + ("foo bar" + "a")))
                                            _ = v67
                                            _ = v68
                                        }
                                        _ = v65
                                    }
                                }
                                fmt.Println(v55, v47)
                                switch "hello" {
                                case "foo bar":
                                    switch "hello" {
                                    case "":
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    case "\n\t":
                                        v57 = v57
                                    case "foo bar":
                                        v5 = append(v5, (^(fi58)))
                                    default:
                                        v1 = true
                                    }
                                default:
                                    switch "hello" {
                                    case "":
                                        var v69 int64 = 1
                                        _ = v69
                                    }
                                }
                                v70 := GenS7{F0: -158, F1: uint64(v55)}
                                _ = fi58
                                _ = v59
                                _ = v70
                            }
                        }
                    }
                    if (-1 != (fi56 % (((fi56 >> (((v3 ^ (v55 + v55))) & 63))) | 1))) {
                        var v71 float32 = float32(v57)
                        var v72 uint64 = 70
                        v73 := GenS5{F0: (-(fi56)), F1: 0}
                        _ = v71
                        _ = v72
                        _ = v73
                    } else {
                        v3 -= v3
                        var v74 float64 = v57
                        _ = v74
                    }
                } else {
                    v3 -= v3
                    v3 += (v3 * 691)
                    v55 = (v55 | v55)
                }
                _ = fi56
                _ = v57
            }
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v2["hello"] = "hello"
        }
        _ = len(v5)
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v54:", v54)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v46
    _ = v47
    _ = v5
    _ = v54
    _ = v9
}
