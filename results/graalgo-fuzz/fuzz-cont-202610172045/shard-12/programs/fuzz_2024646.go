package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 uint32; F1 int16; F2 int }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int; F1 string }
func genFn0(p0 string, p1 string) int8 {
    switch 2147483647 {
    case -9223372036854775808:
        var v5 uint32 = 2147483648
        _ = v5
    case -2147483648:
        v6 := GenS1{F0: 4294967295}
        _ = v6
    }
    if (!(-2.25 < float64((((0.1 + (0.0 + 1e10)) / 1.5) / (float32(1e10) + 1e10))))) {
        var v7 float32 = 1e10
        _ = v7
    } else {
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 int8 = 127
            var v10 int = (fi8 + fi8)
            _ = fi8
            _ = v10
            _ = v9
        }
        p0 = p1
    }
    _ = p0
    _ = p1
    return -128
}
func genFn1() uint64 {
    for fi11 := 0; fi11 < 5; fi11++ {
        if false {
            if ((-2.25 - 0.0) != -2.25) {
                v12 := []int{-421}
                fmt.Println(fi11)
                _ = v12
            } else {
                if (false && (("" > "foo bar") && (0 > 65535))) {
                    v13 := GenS1{F0: 0}
                    var v14 int64 = -1
                    var v15 uint64 = 42
                    _ = v13
                    _ = v14
                    _ = v15
                }
            }
        } else {
            var v16 uint8 = 255
            v16 = v16
            v17 := []int{}
            _ = v16
            _ = v17
        }
        if ((128 <= 61) || (920 <= 2)) {
            for fi18 := 0; fi18 < 14; fi18++ {
                if false {
                    switch ("" + (("a" + "\n\t") + ("" + "\n\t"))) {
                    case "a":
                        var v19 uint32 = 850
                        _ = v19
                    case "foo bar":
                        var v20 int32 = -637
                        _ = v20
                    case "":
                        v21 := GenS2{F0: 4294967295, F1: 100}
                        _ = v21
                    }
                    if (!((79 <= 465) || (!false))) {
                        v22 := []int{(fi11 & fi18), -998, (-(fi18)), (fi18 - fi18)}
                        v23 := GenS1{F0: 0}
                        if len(v22) > 0 {
                            v22[0] = fi18
                        }
                        _ = v22
                        _ = v23
                    }
                    v24 := GenS3{F0: fi11}
                    var v25 uint64 = 0
                    v26 := GenS4{F0: false, F1: -1}
                    v25 -= (v25 | 42)
                    var v27 int = (fi18 & fi18)
                    _ = v24
                    _ = v25
                    _ = v26
                    _ = v27
                } else {
                    var v28 uint16 = 255
                    _ = v28
                }
                _ = fi18
            }
            if true {
                if (false || true) {
                    v29 := []int{fi11}
                    _ = v29
                } else {
                    v30 := []int{}
                    v31 := GenS1{F0: 42}
                    v32 := GenS2{F0: 18446744073709551615, F1: 127}
                    _ = v30
                    _ = v31
                    _ = v32
                }
                v33 := []int{}
                var v34 uint = 2
                _ = v33
                _ = v34
            }
            for fi35 := 0; fi35 < 11; fi35++ {
                v36 := []int{-2147483648, 536, fi11, 127}
                fmt.Println(fi11, fi35)
                v36 = append(v36, fi35)
                _ = fi35
                _ = v36
            }
        } else {
            switch "a" {
            case "":
                if false {
                    v37 := []int{127, fi11, fi11, (-(fi11))}
                    var v38 string = "\n\t"
                    var v39 string = v38
                    _ = v37
                    _ = v38
                    _ = v39
                }
            case "a":
                switch fi11 {
                case 2147483647:
                    var v40 string = "hello"
                    _ = v40
                case -588:
                    v41 := []int{fi11, fi11}
                    _ = v41
                case 9223372036854775807:
                    v42 := GenS3{F0: -1}
                    _ = v42
                }
            default:
                v43 := GenS2{F0: 0, F1: 91}
                _ = v43
            }
            if (85 >= 127) {
                if (692 != 0) {
                    fmt.Println(fi11)
                    fmt.Println(fi11)
                } else {
                    fmt.Println(fi11)
                    var v44 uint64 = 2
                    v45 := GenS3{F0: (^(fi11))}
                    _ = v44
                    _ = v45
                }
            } else {
                v46 := []int{}
                _ = v46
            }
        }
        v47 := GenS3{F0: 32767}
        _ = fi11
        _ = v47
    }
    var v48 int = -157
    for fi49 := 0; fi49 < 13; fi49++ {
        if (false && false) {
            fmt.Println(fi49)
            if false {
                v50 := GenS3{F0: (fi49 | v48)}
                _ = v50
            }
            v48 -= 2147483647
        } else {
            v51 := GenS1{F0: 255}
            v52 := []int{(^(v48)), v48}
            if len(v52) > 0 {
                v52[0] = (-(v48))
            }
            _ = v51
            _ = v52
        }
        for fi53 := 0; fi53 < 5; fi53++ {
            if true {
                if (fi49 > v48) {
                    var v54 float64 = 3.14159
                    _ = v54
                }
                v48 += fi49
            }
            var v55 string = "a"
            _ = fi53
            _ = v55
        }
        _ = fi49
    }
    _ = v48
    return 0
}

func main() {
    var v1 int8 = -1
    v2 := []string{"", "hello", "foo bar"}
    v3 := []string{("" + "a")}
    var v4 uint16 = 0
    fmt.Println(v4, v1)
    v4 = 65535
    switch ((("" + "\n\t") + "a") + "") {
    case "\n\t":
        var v56 uint8 = 252
        _ = v56
    case "":
        if true {
            v57 := GenS2{F0: 1, F1: v1}
            var v58 int32 = 29
            _ = v57
            _ = v58
        }
    }
    for fi59 := 0; fi59 < 13; fi59++ {
        var v60 int8 = -61
        v1 += (v1 << ((65535) & 63))
        _ = fi59
        _ = v60
    }
    v1 = (v1 ^ -1)
    var v61 float64 = float64(3.14159)
    var v62 float64 = v61
    for fi63 := 0; fi63 < 8; fi63++ {
        var v64 uint8 = 1
        var v65 string = ""
        _ = fi63
        _ = v64
        _ = v65
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v66 := GenS5{F0: 4294967295, F1: 42, F2: 1}
    v62 = v61
    switch ("hello" + ("foo bar" + "\n\t")) {
    case "foo bar":
        v67 := GenS6{F0: 0}
        _ = v67
    default:
        fmt.Println(v61, v4, v62)
    }
    var v68 string = "foo bar"
    v69 := GenS4{F0: true, F1: 42}
    if (120 == 2147483648) {
        v61 -= 1.5
        for fi70 := 0; fi70 < 11; fi70++ {
            if true {
                var v71 uint64 = 42
                var v72 string = (v68 + "hello")
                _ = v71
                _ = v72
            } else {
                if (!(4294967295 >= 577)) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v73 uint64 = 0
                    var v74 float64 = v62
                    _ = v73
                    _ = v74
                }
            }
            var v75 float32 = (3.14159 - (3.14159 - 3.14159))
            if len(v3) > 0 {
                v3[0] = ""
            }
            _ = fi70
            _ = v75
        }
        v62 += float64(1.5)
    }
    switch 128 {
    case 746:
        for fi76 := 0; fi76 < 12; fi76++ {
            var v77 int64 = 42
            _ = fi76
            _ = v77
        }
    case -90:
        for fi78 := 0; fi78 < 12; fi78++ {
            var v79 int = 1
            var v80 bool = false
            _ = fi78
            _ = v79
            _ = v80
        }
    default:
        for fi81 := 0; fi81 < 19; fi81++ {
            if len(v3) > 0 {
                v3[0] = ((("" + ("foo bar" + v68)) + "foo bar") + (v68 + v68))
            }
            _ = fi81
        }
    }
    v82 := GenS7{F0: 32767, F1: ""}
    var v83 uint16 = (v4 + v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v68:", v68)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v61
    _ = v62
    _ = v66
    _ = v68
    _ = v69
    _ = v82
    _ = v83
}
