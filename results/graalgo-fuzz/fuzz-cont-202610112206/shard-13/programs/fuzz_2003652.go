package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 int64; F2 int64 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint32; F1 int32; F2 int32 }
type GenS8 struct { F0 int32; F1 int16 }
type GenS9 struct { F0 int16; F1 int16; F2 int }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    var v3 int64 = 531
    var v4 bool = false
    var v5 int32 = 2147483647
    var v6 uint16 = 42
    for fi7 := 0; fi7 < 8; fi7++ {
        if (!(!(!(!v4)))) {
            for fi8 := 0; fi8 < 8; fi8++ {
                fmt.Println(fi7, fi8)
                if true {
                    var v9 int64 = v3
                    _ = v9
                } else {
                    var v10 bool = v4
                    switch (("hello" + ((("a" + ("" + "a")) + ("" + "")) + "hello")) + "hello") {
                    case "a":
                        v11 := []int{fi8, (^(fi7)), (fi8 | (fi7 * (fi7 % ((fi8) | 1))))}
                        _ = v11
                    }
                    _ = v10
                }
                _ = fi8
            }
            v12 := GenS1{F0: (^(v6))}
            for fi13 := 0; fi13 < 10; fi13++ {
                v14 := GenS1{F0: (v6 * 275)}
                _ = fi13
                _ = v14
            }
            _ = v12
        }
        v4 = v4
        _ = fi7
    }
    v5 = (v5 >> ((221) & 63))
    v5 = v5
    switch ("a" + "a") {
    case "":
        v15 := GenS2{F0: (v5 + v5), F1: (v3 & (v3 + (v3 % (((v3 ^ 9223372036854775807)) | 1)))), F2: -9223372036854775808}
        _ = v15
    }
    v4 = (0 <= 0)
    v16 := []int{-1}
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 uint = 674
        if v4 {
            for fi19 := 0; fi19 < 7; fi19++ {
                v20 := v1["hello"]
                _ = fi19
                _ = v20
            }
            if v4 {
                if len(v16) > 0 {
                    v16[0] = (fi17 | (fi17 ^ -2147483648))
                }
                if (fi17 >= (fi17 << ((v18) & 63))) {
                    _ = len(v2)
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    var v21 uint16 = (v6 << (((v18 << ((v18) & 63))) & 63))
                    _ = v21
                } else {
                    switch (("\n\t" + "") + "a") {
                    case "a":
                        fmt.Println(v5, v6)
                    case "foo bar":
                        var v22 bool = false
                        _ = v22
                    case "":
                        switch (fi17 - -2147483648) {
                        case -969:
                            for fi23 := 0; fi23 < 1; fi23++ {
                                if ((128 < 2) || false) {
                                    v24 := GenS3{F0: -86, F1: (fi17 * -515)}
                                    var v25 float32 = 3.14159
                                    switch "" {
                                    case "foo bar":
                                        var v26 uint32 = 4294967295
                                        _ = v26
                                    case "hello":
                                        var v27 int8 = 100
                                        _ = v27
                                    case "a":
                                        var v28 uint16 = uint16(v18)
                                        _ = v28
                                    }
                                    _ = len(v16)
                                    if (false || (!false)) {
                                        if v4 {
                                            if (v6 < v6) {
                                                var v29 bool = (3.14159 > 0.0)
                                                v25 = v25
                                                v5 = v5
                                                _ = v29
                                            } else {
                                                fmt.Println(v6, v25)
                                            }
                                            switch "hello" {
                                            case "foo bar":
                                                if true {
                                                    v30 := GenS4{F0: 238}
                                                    _ = v30
                                                }
                                            case "":
                                                v31 := GenS5{F0: 2}
                                                _ = v31
                                            case "hello":
                                                var v32 string = "\n\t"
                                                _ = v32
                                            default:
                                                v4 = (v4 && true)
                                                var v33 int = fi23
                                                v16 = append(v16, v33)
                                                switch "foo bar" {
                                                case "":
                                                    switch -2147483648 {
                                                    case -972:
                                                        v5 = v5
                                                    case 1:
                                                        v33 = (-(v33))
                                                    case 2147483647:
                                                        v34 := v2[971]
                                                        _ = v34
                                                    }
                                                case "a":
                                                    if len(v16) > 0 {
                                                        v16[0] = fi23
                                                    }
                                                case "\n\t":
                                                    v16 = append(v16, -2147483648)
                                                }
                                                if (!v4) {
                                                    v35 := GenS5{F0: v6}
                                                    v36 := v1["foo bar"]
                                                    v37 := GenS5{F0: (^(v6))}
                                                    _ = v35
                                                    _ = v36
                                                    _ = v37
                                                } else {
                                                    if (fi23 == fi17) {
                                                        v38 := GenS6{F0: fi23}
                                                        _ = v38
                                                    }
                                                    if v4 {
                                                        var v39 uint8 = uint8(v18)
                                                        v16 = append(v16, 4294967295)
                                                        _ = v39
                                                    }
                                                }
                                                var v40 uint8 = 255
                                                _ = v33
                                                _ = v40
                                            }
                                        } else {
                                            var v41 int32 = (^(v5))
                                            if true {
                                                if (!((!(100 != 0)) || (false && false))) {
                                                    fmt.Println(v25, fi23, v41)
                                                    v16 = append(v16, int(v18))
                                                    v42 := GenS7{F0: 319, F1: v41, F2: v5}
                                                    _ = v42
                                                } else {
                                                    var v43 int8 = -128
                                                    switch 42 {
                                                    case 127:
                                                        v18 -= v18
                                                    }
                                                    _ = v43
                                                }
                                            } else {
                                                fmt.Println(v25, fi17, v5)
                                                v44 := GenS6{F0: -9223372036854775808}
                                                _ = v44
                                            }
                                            var v45 uint32 = 4294967295
                                            v46 := GenS8{F0: -100, F1: 42}
                                            _ = v41
                                            _ = v45
                                            _ = v46
                                        }
                                        var v47 uint64 = 638
                                        _ = v47
                                    } else {
                                        v6 = v6
                                    }
                                    _ = v24
                                    _ = v25
                                } else {
                                    fmt.Println(fi17, v5, v3)
                                    switch "a" {
                                    case "a":
                                        if false {
                                            var v48 uint8 = 19
                                            v49 := GenS9{F0: 32767, F1: 766, F2: fi17}
                                            v4 = (v4 && (v4 && ((-128 < 97) && (-2.25 <= 0.1))))
                                            var v50 uint = uint(v48)
                                            _ = v48
                                            _ = v49
                                            _ = v50
                                        }
                                    case "\n\t":
                                        switch fi17 {
                                        case 42:
                                            v5 = v5
                                        }
                                    default:
                                        var v51 string = "hello"
                                        _ = v51
                                    }
                                    _ = len(v2)
                                }
                                fmt.Println(fi17, v3)
                                v52 := GenS5{F0: uint16(v18)}
                                _ = fi23
                                _ = v52
                            }
                        case 1:
                            var v53 float64 = -2.25
                            _ = v53
                        default:
                            var v54 int32 = (v5 + (v5 / ((0) | 1)))
                            _ = v54
                        }
                    }
                    for fi55 := 0; fi55 < 5; fi55++ {
                        v4 = (255 < 255)
                        var v56 uint = 893
                        if (!v4) {
                            v3 = (v3 | -349)
                        }
                        _ = fi55
                        _ = v56
                    }
                }
            }
            var v57 bool = (!(-2147483648 > (fi17 - fi17)))
            _ = v57
        } else {
            v6 = v6
            for fi58 := 0; fi58 < 20; fi58++ {
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                if ("foo bar" != ("" + ((("a" + "foo bar") + "a") + ""))) {
                    v59 := v1[("hello" + "hello")]
                    v60 := GenS5{F0: v6}
                    _ = v59
                    _ = v60
                }
                var v61 float64 = -2.25
                _ = fi58
                _ = v61
            }
            switch "hello" {
            case "hello":
                v16 = append(v16, fi17)
            case "":
                if (!(-32768 > 1)) {
                    v62 := GenS4{F0: 2}
                    for fi63 := 0; fi63 < 17; fi63++ {
                        var v64 uint16 = (v6 << ((42) & 63))
                        _ = len(v16)
                        _ = fi63
                        _ = v64
                    }
                    if v4 {
                        if len(v16) > 0 {
                            v16[0] = -2147483648
                        }
                        v65 := GenS3{F0: 127, F1: (-(fi17))}
                        var v66 int64 = v3
                        _ = v65
                        _ = v66
                    }
                    _ = v62
                } else {
                    var v67 string = "a"
                    var v68 int64 = (v3 - 9223372036854775807)
                    for fi69 := 0; fi69 < 5; fi69++ {
                        switch fi17 {
                        case 1:
                            fmt.Println(v67)
                            _ = len(v2)
                            v68 -= -9223372036854775808
                            if v4 {
                                v18 = (v18 / ((42) | 1))
                            }
                            var v70 uint64 = 42
                            switch (v67 + v67) {
                            case "":
                                v6 += (v6 % ((v6) | 1))
                            case "a":
                                var v71 uint32 = 2
                                _ = v71
                            default:
                                var v72 float32 = 1e10
                                fmt.Println(v70, v70)
                                fmt.Println(v4, v3)
                                delete(v2, 42)
                                v72 += v72
                                _ = v72
                            }
                            _ = v70
                        case 32767:
                            fmt.Println(v4)
                        default:
                            var v73 string = (v67 + (v67 + v67))
                            _ = v73
                        }
                        var v74 float32 = 0.0
                        _ = fi69
                        _ = v74
                    }
                    _ = v67
                    _ = v68
                }
            case "\n\t":
                var v75 int = fi17
                _ = v75
            }
        }
        fmt.Println(v3, v3)
        _ = fi17
        _ = v18
    }
    fmt.Println(v4, v3)
    v76 := GenS10{F0: 399}
    var v77 int = 32768
    fmt.Println(v3, v5)
    var v78 uint32 = 182
    v3 = 671
    v79 := GenS11{F0: (v78 ^ v78)}
    for fi80 := 0; fi80 < 9; fi80++ {
        v5 -= int32(v78)
        _ = fi80
    }
    for fi81 := 0; fi81 < 19; fi81++ {
        for fi82 := 0; fi82 < 2; fi82++ {
            fmt.Println(v3, v3, v77)
            v2[(v78 + 947)] = ((("\n\t" + "") + ("hello" + ("hello" + "foo bar"))) + "foo bar")
            _ = fi82
        }
        switch "" {
        case "a":
            _ = len(v2)
        case "hello":
            fmt.Println(v4, v5)
        case "":
            v78 += v78
        }
        if (v4 && (890 == 1)) {
            for fi83 := 0; fi83 < 10; fi83++ {
                delete(v2, v78)
                if v4 {
                    v6 = (v6 << ((4294967295) & 63))
                    var v84 int64 = v3
                    if v4 {
                        v85 := GenS3{F0: -75, F1: v77}
                        v2[v78] = "\n\t"
                        _ = v85
                    } else {
                        var v86 float32 = 3.14159
                        switch v77 {
                        case 127:
                            if false {
                                v3 -= 0
                                var v87 uint16 = v6
                                _ = len(v16)
                                _ = v87
                            }
                        case 13:
                            fmt.Println(fi83, v84)
                        default:
                            v78 = 2147483648
                        }
                        _ = v86
                    }
                    v88 := GenS10{F0: (^(fi81))}
                    if v4 {
                        var v89 int = int(v84)
                        v77 += fi83
                        _ = v89
                    } else {
                        v77 += fi83
                    }
                    _ = v84
                    _ = v88
                } else {
                    switch "" {
                    case "foo bar":
                        var v90 string = "foo bar"
                        _ = v90
                    case "hello":
                        v5 += -2147483648
                    case "":
                        v3 = v3
                    }
                }
                _ = fi83
            }
            switch "foo bar" {
            case "":
                fmt.Println(v77)
            case "foo bar":
                var v91 int16 = -492
                _ = v91
            }
        }
        _ = fi81
    }
    switch ("hello" + "a") {
    case "\n\t":
        fmt.Println(v5, v77, v77)
    case "hello":
        v6 = (v6 / (((v6 - (v6 - (^(v6))))) | 1))
    case "":
        v4 = false
    default:
        for fi92 := 0; fi92 < 5; fi92++ {
            v78 += v78
            _ = fi92
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
    _ = v77
    _ = v78
    _ = v79
}
