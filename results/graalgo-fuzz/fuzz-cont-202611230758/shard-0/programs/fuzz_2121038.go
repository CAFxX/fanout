package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int16; F1 string; F2 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 int16; F2 uint }
type GenS6 struct { F0 int16 }
func genFn0(p0 int, p1 uint) int16 {
    var v4 int64 = 0
    var v5 int64 = (v4 + v4)
    var v6 string = ("a" + ("\n\t" + "hello"))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 32767
}

func main() {
    var v1 string = "foo bar"
    var v2 bool = true
    var v3 bool = (42 >= 2147483648)
    var v7 uint = 79
    var v8 bool = (255 >= v7)
    v7 = (v7 >> ((v7) & 63))
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 uint16 = 361
        _ = fi9
        _ = v10
    }
    if true {
        var v11 uint = v7
        _ = v11
    } else {
        var v12 int64 = 1
        var v13 uint8 = 128
        v13 = 138
        _ = v12
        _ = v13
    }
    var v14 float64 = float64(0.0)
    v8 = false
    var v15 uint8 = 1
    var v16 uint32 = 2147483647
    if true {
        v17 := GenS1{F0: 9}
        _ = v17
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        switch v1 {
        case "foo bar":
            switch v1 {
            case "a":
                switch fi18 {
                case 1:
                    if (-1 <= -128) {
                        fmt.Println(v7)
                        v16 += v16
                        fmt.Println(v1, v2)
                    }
                case 9223372036854775807:
                    if (!v2) {
                        var v19 int8 = 39
                        var v20 int64 = 42
                        var v21 float64 = v14
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                default:
                    switch ((("a" + "a") + "\n\t") + (v1 + (v1 + "foo bar"))) {
                    case "a":
                        v22 := GenS2{F0: v1, F1: uint16(v15)}
                        _ = v22
                    case "":
                        for fi23 := 0; fi23 < 9; fi23++ {
                            fmt.Println(v7, v8)
                            var v24 uint16 = 814
                            _ = fi23
                            _ = v24
                        }
                    case "foo bar":
                        v14 += 0.0
                    }
                }
            case "":
                var v25 string = "\n\t"
                _ = v25
            case "hello":
                if (v8 && (!v8)) {
                    var v26 float32 = 1e10
                    for fi27 := 0; fi27 < 9; fi27++ {
                        if (-739 > int32(v16)) {
                            var v28 float32 = v26
                            v29 := GenS2{F0: v1, F1: 65}
                            _ = v28
                            _ = v29
                        }
                        v26 = (float32((v14 / 0.1)) + v26)
                        var v30 float32 = v26
                        _ = fi27
                        _ = v30
                    }
                    v31 := GenS3{F0: -1, F1: (v1 + ("\n\t" + v1)), F2: -1}
                    _ = v26
                    _ = v31
                }
            }
        case "a":
            fmt.Println(fi18, v15)
        case "\n\t":
            var v32 uint8 = (v15 ^ 128)
            _ = v32
        }
        if (v7 < v7) {
            if v2 {
                var v33 string = (v1 + v1)
                _ = v33
            } else {
                if v8 {
                    for fi34 := 0; fi34 < 20; fi34++ {
                        fmt.Println(v3, fi18, v3)
                        v14 += ((float64((3.14159 + float32(v14))) * 0.0) / v14)
                        fmt.Println(v15)
                        var v35 uint64 = 2
                        _ = fi34
                        _ = v35
                    }
                } else {
                    v7 = (^(v7))
                    var v36 float32 = float32((3.14159 + v14))
                    _ = v36
                }
            }
        } else {
            fmt.Println(v3, v14, v15)
            v37 := []int{1, 557, (fi18 + fi18), fi18}
            fmt.Println(fi18, fi18, v15)
            _ = v37
        }
        switch (-(fi18)) {
        case 128:
            var v38 string = v1
            _ = v38
        default:
            v39 := GenS3{F0: genFn0((fi18 - fi18), (^(v7))), F1: "foo bar", F2: 850}
            _ = v39
        }
        _ = fi18
    }
    switch ("foo bar" + ("a" + "foo bar")) {
    case "foo bar":
        v14 += 1.5
    case "\n\t":
        var v40 float64 = v14
        _ = v40
    default:
        for fi41 := 0; fi41 < 18; fi41++ {
            for fi42 := 0; fi42 < 12; fi42++ {
                var v43 int32 = 0
                var v44 int64 = 0
                _ = fi42
                _ = v43
                _ = v44
            }
            if v2 {
                for fi45 := 0; fi45 < 18; fi45++ {
                    v8 = (!false)
                    switch fi45 {
                    case 42:
                        if (42 >= 2147483647) {
                            if (4294967295 == 4294967295) {
                                v46 := []int{0, fi41, fi45, fi45}
                                if true {
                                    v8 = v2
                                } else {
                                    if len(v46) >= 2 {
                                        v46 = v46[1:len(v46)]
                                    }
                                    if ((!v8) && false) {
                                        if ((v14 <= ((v14 * (v14 - float64((((3.14159 + (0.1 - (((0.0 - 1e10) - 1e10) + -2.25))) * 3.14159) / 3.14159)))) * v14)) && (1 > 732)) {
                                            if len(v46) > 0 {
                                                v46[0] = (fi45 - fi41)
                                            }
                                            var v47 uint16 = 2
                                            v2 = false
                                            _ = v47
                                        }
                                        if ((v7 * v7) < (v7 - 418)) {
                                            v15 += v15
                                            switch v1 {
                                            case "foo bar":
                                                v14 += v14
                                            case "hello":
                                                v48 := GenS1{F0: 42}
                                                _ = v48
                                            }
                                        } else {
                                            var v49 int = fi45
                                            v46 = append(v46, (^(v49)))
                                            v50 := GenS4{F0: v15}
                                            _ = v49
                                            _ = v50
                                        }
                                        if (((v2 || v8) || v3) && true) {
                                            v15 -= v15
                                        }
                                    } else {
                                        var v51 int16 = -966
                                        _ = v51
                                    }
                                }
                                var v52 float64 = (0.1 + 1e10)
                                fmt.Println(v3)
                                var v53 string = ""
                                var v54 uint32 = uint32(v7)
                                if (!((1e10 * (-2.25 - 0.0)) > 1.5)) {
                                    v16 = 1
                                    v55 := GenS4{F0: (^(v15))}
                                    v46 = append(v46, fi41)
                                    _ = v55
                                }
                                var v56 uint8 = v15
                                _ = v46
                                _ = v52
                                _ = v53
                                _ = v54
                                _ = v56
                            } else {
                                v15 = v15
                            }
                            var v57 uint = v7
                            _ = v57
                        }
                    case 2147483647:
                        v58 := []int{}
                        _ = v58
                    }
                    if v2 {
                        var v59 bool = false
                        switch v1 {
                        case "hello":
                            fmt.Println(v3, fi41)
                        case "a":
                            var v60 uint16 = 255
                            _ = v60
                        case "\n\t":
                            switch (-(fi41)) {
                            case -31:
                                v1 = "hello"
                            case 4294967295:
                                fmt.Println(v59, v59, v1)
                            }
                        default:
                            var v61 float64 = (1e10 / v14)
                            _ = v61
                        }
                        _ = v59
                    }
                    _ = fi45
                }
                if false {
                    if v3 {
                        var v62 uint16 = 255
                        v63 := GenS3{F0: -1, F1: v1, F2: int16(v7)}
                        fmt.Println(v15)
                        _ = v62
                        _ = v63
                    } else {
                        v3 = (42 >= 1)
                        v64 := GenS4{F0: uint8(v16)}
                        var v65 int32 = 0
                        _ = v64
                        _ = v65
                    }
                    for fi66 := 0; fi66 < 17; fi66++ {
                        v67 := GenS5{F0: 65535, F1: -1, F2: 546}
                        var v68 bool = v8
                        if v3 {
                            var v69 bool = v8
                            _ = v69
                        } else {
                            fmt.Println(v16, v1)
                        }
                        _ = fi66
                        _ = v67
                        _ = v68
                    }
                    if (768 > 761) {
                        v8 = (-961 >= 9223372036854775807)
                        for fi70 := 0; fi70 < 13; fi70++ {
                            v71 := []int{(^(fi41)), 9223372036854775807, int(v15)}
                            switch -926 {
                            case 32767:
                                var v72 string = "\n\t"
                                _ = v72
                            case -9223372036854775808:
                                v8 = v2
                            }
                            var v73 uint32 = 1
                            _ = fi70
                            _ = v71
                            _ = v73
                        }
                        if v8 {
                            v8 = (!(!(false && true)))
                        }
                    }
                } else {
                    if (!(int32(v16) <= -185)) {
                        v74 := []int{fi41, int(v7), fi41}
                        v8 = false
                        _ = v74
                    }
                    switch v1 {
                    case "":
                        fmt.Println(v8, v2)
                    }
                    var v75 bool = (42 >= 9223372036854775807)
                    _ = v75
                }
                v1 = v1
            }
            v76 := GenS6{F0: 482}
            _ = fi41
            _ = v76
        }
    }
    var v77 uint = (v7 << ((2) & 63))
    v7 += v7
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v77:", v77)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v7
    _ = v77
    _ = v8
}
