package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 uint; F2 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint; F1 int }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int8 }

func main() {
    var v1 int64 = 0
    v2 := make(map[string]string)
    v3 := []int{0}
    var v4 bool = false
    v5 := []string{((("a" + "") + "") + "hello")}
    var v6 uint32 = 1
    for fi7 := 0; fi7 < 7; fi7++ {
        if len(v5) > 0 {
            v5[0] = "hello"
        }
        v8 := GenS1{F0: -2147483648, F1: 1, F2: -1}
        _ = fi7
        _ = v8
    }
    var v9 float32 = -2.25
    v9 -= v9
    if (!true) {
        if (-2147483648 != -2147483648) {
            v4 = true
            _ = len(v5)
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v2["a"] = "foo bar"
            v4 = (!v4)
        }
        var v10 string = "a"
        _ = v10
    }
    var v11 int16 = 1000
    for fi12 := 0; fi12 < 1; fi12++ {
        v4 = true
        fmt.Println(v11, v6, v6)
        _ = fi12
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        v6 = (^(v6))
        for fi14 := 0; fi14 < 11; fi14++ {
            _ = len(v2)
            _ = fi14
        }
        if v4 {
            v15 := GenS1{F0: 0, F1: -1, F2: -1}
            v4 = (42 <= 234)
            switch "foo bar" {
            case "a":
                if false {
                    var v16 string = "\n\t"
                    switch -121 {
                    case 4294967295:
                        v4 = false
                    case 2147483647:
                        v17 := GenS1{F0: -1, F1: 1, F2: 42}
                        _ = v17
                    case -879:
                        var v18 bool = (-282 > v11)
                        _ = v18
                    default:
                        switch fi13 {
                        case 108:
                            switch fi13 {
                            case 127:
                                v4 = ((!(255 < 52)) || (0.1 >= 3.14159))
                            case 9223372036854775807:
                                v19 := GenS2{F0: 128, F1: 401, F2: 96}
                                _ = v19
                            case 2147483647:
                                for fi20 := 0; fi20 < 8; fi20++ {
                                    if (v1 != v1) {
                                        switch v16 {
                                        case "\n\t":
                                            var v21 uint16 = 294
                                            _ = v21
                                        case "a":
                                            fmt.Println(v11)
                                            var v22 int32 = 42
                                            var v23 int64 = (v1 >> ((0) & 63))
                                            _ = v22
                                            _ = v23
                                        case "":
                                            v24 := GenS2{F0: 78, F1: 2, F2: 2}
                                            _ = v24
                                        default:
                                            v25 := GenS3{F0: 4294967295}
                                            _ = v25
                                        }
                                        v11 = -32768
                                        if (!((v11 / ((v11) | 1)) > 1000)) {
                                            fmt.Println(v4, v11)
                                            var v26 int8 = 100
                                            var v27 uint32 = 410
                                            _ = v26
                                            _ = v27
                                        } else {
                                            v3 = append(v3, 128)
                                            if (255 < 42) {
                                                if (((1.5 + 0.0) + 1.5) >= 1e10) {
                                                    var v28 string = ("hello" + "hello")
                                                    v4 = (v4 && (v4 && ((float32((-2.25 / 1.0)) == -2.25) && v4)))
                                                    _ = v28
                                                }
                                                if v4 {
                                                    if ((v16 + v16) <= v16) {
                                                        if v4 {
                                                            var v29 int32 = -2147483648
                                                            _ = v29
                                                        }
                                                    }
                                                } else {
                                                    switch v16 {
                                                    case "\n\t":
                                                        if true {
                                                            switch fi20 {
                                                            case -9223372036854775808:
                                                                _ = len(v3)
                                                            case 32768:
                                                                var v30 int8 = 100
                                                                v30 = v30
                                                                if v4 {
                                                                    v4 = v4
                                                                    if true {
                                                                        switch fi13 {
                                                                        case 1:
                                                                            var v31 int = 1
                                                                            _ = v31
                                                                        case -963:
                                                                            switch "foo bar" {
                                                                            case "a":
                                                                                var v32 int = fi20
                                                                                _ = v32
                                                                            case "":
                                                                                switch (v16 + (("hello" + "a") + v16)) {
                                                                                case "foo bar":
                                                                                    fmt.Println(v4, v4, v9)
                                                                                case "a":
                                                                                    v11 -= v11
                                                                                }
                                                                            default:
                                                                                fmt.Println(v16, fi20, v30)
                                                                            }
                                                                        case 0:
                                                                            v3 = append(v3, fi13)
                                                                        }
                                                                    } else {
                                                                        var v33 uint16 = 142
                                                                        v16 = v16
                                                                        if true {
                                                                            v16 = ((((v16 + v16) + "hello") + (v16 + "\n\t")) + v16)
                                                                            v34 := GenS2{F0: 128, F1: 4294967295, F2: uint16(v6)}
                                                                            v2["hello"] = ("hello" + v16)
                                                                            _ = v34
                                                                        } else {
                                                                            v35 := GenS4{F0: fi13, F1: uint(v6), F2: 2147483647}
                                                                            _ = v35
                                                                        }
                                                                        var v36 uint8 = 36
                                                                        if ((!false) && (2 <= 92)) {
                                                                            v30 = 42
                                                                        } else {
                                                                            var v37 uint8 = v36
                                                                            switch (fi20 - (fi20 & (fi13 / ((4294967295) | 1)))) {
                                                                            case -286:
                                                                                v38 := GenS5{F0: ""}
                                                                                _ = v38
                                                                            case -114:
                                                                                fmt.Println(v11)
                                                                            case 4294967295:
                                                                                v2[v16] = v16
                                                                            default:
                                                                                var v39 float64 = (-2.25 + float64((v9 / 1e10)))
                                                                                _ = v39
                                                                            }
                                                                            _ = v37
                                                                        }
                                                                        _ = v33
                                                                        _ = v36
                                                                    }
                                                                } else {
                                                                    var v40 uint16 = 1
                                                                    var v41 bool = true
                                                                    var v42 int64 = 934
                                                                    if ((!(!(v42 == v42))) && (v1 > v1)) {
                                                                        switch v16 {
                                                                        case "foo bar":
                                                                            if true {
                                                                                var v43 uint64 = 18446744073709551615
                                                                                _ = v43
                                                                            }
                                                                            switch fi13 {
                                                                            case 127:
                                                                                if v4 {
                                                                                    v6 = 2147483647
                                                                                    v11 = (v11 * 32767)
                                                                                    switch v16 {
                                                                                    case "foo bar":
                                                                                        v1 = (v1 ^ (v1 & v1))
                                                                                    default:
                                                                                        v1 -= -9223372036854775808
                                                                                    }
                                                                                    var v44 uint32 = (v6 ^ v6)
                                                                                    var v45 float64 = 0.0
                                                                                    _ = v44
                                                                                    _ = v45
                                                                                } else {
                                                                                    var v46 string = v16
                                                                                    switch ("" + "\n\t") {
                                                                                    case "a":
                                                                                        delete(v2, v46)
                                                                                        v2[v46] = "\n\t"
                                                                                        if (!(v41 && v41)) {
                                                                                            v47 := GenS5{F0: v46}
                                                                                            if v41 {
                                                                                                if v4 {
                                                                                                    v1 = (v42 - v42)
                                                                                                } else {
                                                                                                    var v48 uint8 = 42
                                                                                                    v49 := GenS1{F0: -174, F1: (v30 >> ((65535) & 63)), F2: (v30 + -33)}
                                                                                                    _ = v48
                                                                                                    _ = v49
                                                                                                }
                                                                                                if (!false) {
                                                                                                    var v50 uint64 = 2
                                                                                                    _ = v50
                                                                                                }
                                                                                            }
                                                                                            if (-658 <= -2147483648) {
                                                                                                fmt.Println(v30)
                                                                                            }
                                                                                            _ = v47
                                                                                        } else {
                                                                                            var v51 uint32 = 2147483648
                                                                                            v30 = (-(v30))
                                                                                            _ = v51
                                                                                        }
                                                                                    case "\n\t":
                                                                                        var v52 uint16 = v40
                                                                                        _ = v52
                                                                                    case "hello":
                                                                                        if v4 {
                                                                                            v40 += v40
                                                                                            var v53 int32 = -839
                                                                                            var v54 int64 = int64(v40)
                                                                                            var v55 float32 = 0.0
                                                                                            v56 := GenS3{F0: 2}
                                                                                            v54 = (v1 + v54)
                                                                                            v57 := GenS6{F0: 2, F1: 473}
                                                                                            _ = v53
                                                                                            _ = v54
                                                                                            _ = v55
                                                                                            _ = v56
                                                                                            _ = v57
                                                                                        }
                                                                                    default:
                                                                                        v1 -= -947
                                                                                    }
                                                                                    if (!(v4 && (v42 == v1))) {
                                                                                        if (!(1 >= 100)) {
                                                                                            var v58 string = v16
                                                                                            if v4 {
                                                                                                var v59 uint64 = uint64(v6)
                                                                                                _ = v59
                                                                                            }
                                                                                            _ = v58
                                                                                        }
                                                                                        v60 := GenS2{F0: 128, F1: 42, F2: v40}
                                                                                        _ = v60
                                                                                    } else {
                                                                                        var v61 bool = (((0 < 2) && (v9 != 0.0)) && (v30 == (^(v30))))
                                                                                        _ = v61
                                                                                    }
                                                                                    _ = v46
                                                                                }
                                                                            }
                                                                        case "hello":
                                                                            v62 := GenS7{F0: v11}
                                                                            _ = v62
                                                                        case "":
                                                                            if len(v5) > 0 {
                                                                                v5[0] = v16
                                                                            }
                                                                        default:
                                                                            v11 = (v11 - (v11 & v11))
                                                                        }
                                                                        switch (fi13 % (((fi20 / ((fi13) | 1))) | 1)) {
                                                                        case 127:
                                                                            v40 -= (v40 + 880)
                                                                        case 0:
                                                                            v1 = (v42 >> ((4294967295) & 63))
                                                                        case 128:
                                                                            var v63 int8 = (v30 & v30)
                                                                            _ = v63
                                                                        }
                                                                        if len(v5) > 0 {
                                                                            v5[0] = v16
                                                                        }
                                                                    } else {
                                                                        v64 := GenS8{F0: fi20}
                                                                        v65 := GenS4{F0: (fi20 << ((373) & 63)), F1: 4294967295, F2: 1}
                                                                        _ = v64
                                                                        _ = v65
                                                                    }
                                                                    _ = v40
                                                                    _ = v41
                                                                    _ = v42
                                                                }
                                                                _ = v30
                                                            case 127:
                                                                if v4 {
                                                                    v6 += v6
                                                                    v16 = ""
                                                                }
                                                            default:
                                                                var v66 string = ("hello" + ("" + v16))
                                                                _ = v66
                                                            }
                                                        } else {
                                                            var v67 int64 = v1
                                                            var v68 string = v16
                                                            switch fi20 {
                                                            case 1:
                                                                v69 := GenS9{F0: -1}
                                                                _ = v69
                                                            case 4294967295:
                                                                switch (fi13 & fi20) {
                                                                case -415:
                                                                    v16 = ""
                                                                case 9223372036854775807:
                                                                    v70 := GenS9{F0: 0}
                                                                    _ = v70
                                                                case -32:
                                                                    v3 = append(v3, fi13)
                                                                }
                                                            case -940:
                                                                v71 := GenS5{F0: (v68 + "hello")}
                                                                _ = v71
                                                            }
                                                            _ = v67
                                                            _ = v68
                                                        }
                                                    case "hello":
                                                        v9 = ((v9 - v9) / v9)
                                                    case "":
                                                        fmt.Println(v1, v1, v4)
                                                        if v4 {
                                                            var v72 bool = (1 > 1)
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            _ = v72
                                                        }
                                                    }
                                                }
                                                fmt.Println(v11, v11, v1)
                                            } else {
                                                fmt.Println(v1, v1)
                                            }
                                        }
                                    }
                                    _ = fi20
                                }
                            }
                        case 9223372036854775807:
                            v2[(((v16 + v16) + ("a" + "")) + v16)] = v16
                        case 329:
                            v73 := GenS7{F0: v11}
                            _ = v73
                        default:
                            for fi74 := 0; fi74 < 8; fi74++ {
                                v11 -= (v11 / (((v11 % ((v11) | 1))) | 1))
                                var v75 int8 = 0
                                var v76 int = fi13
                                v6 += v6
                                fmt.Println(fi13)
                                _ = fi74
                                _ = v75
                                _ = v76
                            }
                        }
                    }
                    var v77 bool = false
                    _ = v16
                    _ = v77
                }
            }
            _ = v15
        } else {
            v9 = 1e10
            fmt.Println(v4)
        }
        _ = fi13
    }
    v78 := GenS10{F0: 100}
    v9 = v9
    v11 = 1000
    var v79 int = 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v79:", v79)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v78
    _ = v79
    _ = v9
}
