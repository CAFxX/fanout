package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 int; F2 uint64 }
type GenS5 struct { F0 uint16; F1 int32 }
func genFn0(p0 uint16) bool {
    if (42 < 32767) {
        if (18446744073709551615 == 277) {
            p0 -= 42
            var v6 uint8 = 159
            _ = v6
        }
    }
    var v7 float32 = -2.25
    v8 := []int{360, 32767}
    var v9 int8 = -1
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (true || ((uint64(p0) > 42) || true))
}

func main() {
    var v1 uint32 = 1
    var v2 float64 = (-2.25 - 1.5)
    var v3 uint8 = 255
    var v4 uint = 1
    var v5 bool = false
    switch 1 {
    case 128:
        var v10 uint32 = (v1 % ((v1) | 1))
        _ = v10
    default:
        var v11 float32 = (3.14159 * float32(((v2 + v2) * 0.1)))
        _ = v11
    }
    var v12 int16 = 42
    if false {
        v13 := GenS1{F0: 128}
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        v15 := []int{}
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        _ = fi14
        _ = v15
    }
    if v5 {
        var v16 int16 = (v12 / ((v12) | 1))
        _ = v16
    } else {
        if (!true) {
            if false {
                for fi17 := 0; fi17 < 14; fi17++ {
                    var v18 uint64 = 4294967295
                    if (9223372036854775807 >= -9223372036854775808) {
                        if (v5 && true) {
                            var v19 string = (("foo bar" + "a") + "foo bar")
                            if v5 {
                                var v20 uint8 = 42
                                v12 = 42
                                _ = v20
                            } else {
                                v21 := GenS1{F0: v3}
                                v22 := []int{(fi17 >> ((4294967295) & 63)), fi17, fi17}
                                var v23 bool = (((!true) || (1 >= 1000)) || v5)
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                            fmt.Println(v3)
                            _ = v19
                        } else {
                            fmt.Println(v4, v3, v3)
                            var v24 uint16 = 878
                            _ = v24
                        }
                        var v25 int32 = -1
                        var v26 int8 = int8(v4)
                        _ = v25
                        _ = v26
                    } else {
                        v27 := GenS1{F0: 255}
                        for fi28 := 0; fi28 < 18; fi28++ {
                            v1 -= 2147483648
                            _ = fi28
                        }
                        v29 := GenS1{F0: v3}
                        _ = v27
                        _ = v29
                    }
                    var v30 int64 = 9223372036854775807
                    _ = fi17
                    _ = v18
                    _ = v30
                }
                v31 := []int{128, 4294967295, 440, 127}
                v32 := GenS2{F0: 1, F1: 42}
                _ = v31
                _ = v32
            }
        }
        v33 := []int{}
        if v5 {
            v34 := GenS2{F0: 1, F1: -34}
            var v35 int16 = v12
            _ = v34
            _ = v35
        }
        _ = v33
    }
    v3 += (^(v3))
    for fi36 := 0; fi36 < 7; fi36++ {
        for fi37 := 0; fi37 < 2; fi37++ {
            v1 += 42
            v12 = (v12 / (((v12 >> (((v4 + v4)) & 63))) | 1))
            fmt.Println(fi37, v2, fi36)
            var v38 uint64 = 0
            _ = fi37
            _ = v38
        }
        var v39 float64 = v2
        var v40 int32 = 1
        _ = fi36
        _ = v39
        _ = v40
    }
    fmt.Println(v2, v1)
    var v41 float32 = 1.5
    var v42 uint16 = 42
    for fi43 := 0; fi43 < 20; fi43++ {
        v1 = 710
        _ = fi43
    }
    v1 = v1
    v42 = (^(v42))
    if (v5 && true) {
        v44 := GenS2{F0: 18446744073709551615, F1: 15}
        var v45 int64 = 230
        var v46 uint8 = v3
        _ = v44
        _ = v45
        _ = v46
    } else {
        if (!v5) {
            v12 = 42
            var v47 uint16 = (^(v42))
            fmt.Println(v2, v5, v42)
            _ = v47
        }
        switch -754 {
        case 32767:
            if true {
                v42 = (v42 / ((v42) | 1))
                for fi48 := 0; fi48 < 20; fi48++ {
                    for fi49 := 0; fi49 < 17; fi49++ {
                        switch (fi48 + 127) {
                        case 4294967295:
                            switch (^(fi49)) {
                            case 0:
                                v50 := GenS3{F0: false}
                                _ = v50
                            case 32768:
                                fmt.Println(v42)
                            default:
                                if (4294967295 >= 18446744073709551615) {
                                    var v51 uint = (v4 >> ((218) & 63))
                                    v1 = (v1 + (v1 << (((v51 >> ((42) & 63))) & 63)))
                                    _ = v51
                                }
                            }
                        default:
                            v41 = (0.1 + 1e10)
                        }
                        _ = fi49
                    }
                    v4 = (v4 * 42)
                    _ = fi48
                }
            }
        case 42:
            v52 := GenS1{F0: 2}
            _ = v52
        default:
            if v5 {
                fmt.Println(v3, v42, v12)
                switch -2147483648 {
                case -1:
                    fmt.Println(v12, v4)
                case -9223372036854775808:
                    if (791 < 496) {
                        var v53 int16 = v12
                        for fi54 := 0; fi54 < 16; fi54++ {
                            switch ("hello" + ("a" + "foo bar")) {
                            case "\n\t":
                                if (!((v41 != 3.14159) && false)) {
                                    if (2 != (^(v42))) {
                                        for fi55 := 0; fi55 < 16; fi55++ {
                                            if (-2.25 != ((float32(-2.25) / v41) / ((v41 + 0.1) / 1e10))) {
                                                if v5 {
                                                    v56 := []int{fi54, (fi54 << ((344) & 63))}
                                                    var v57 uint8 = v3
                                                    _ = v56
                                                    _ = v57
                                                }
                                                var v58 float32 = (float32(v2) + float32((v2 + (((v2 - v2) - 0.0) * (v2 / 0.1)))))
                                                var v59 int64 = -918
                                                switch int(v53) {
                                                case -971:
                                                    var v60 uint64 = 1
                                                    _ = v60
                                                case 42:
                                                    v58 = (v58 - (v58 * float32(float64(((v58 + v41) + v58)))))
                                                case 128:
                                                    var v61 int64 = v59
                                                    _ = v61
                                                }
                                                v62 := []int{(fi55 % (((fi55 & (fi55 * -9223372036854775808))) | 1)), fi55, fi55}
                                                _ = v58
                                                _ = v59
                                                _ = v62
                                            } else {
                                                fmt.Println(v5)
                                                v4 = uint(v42)
                                            }
                                            if ((v1 + v1) > v1) {
                                                switch fi55 {
                                                case 1:
                                                    v63 := []int{(fi55 * 9223372036854775807)}
                                                    v63 = append(v63, fi54)
                                                    _ = v63
                                                case -980:
                                                    if false {
                                                        switch ("\n\t" + "") {
                                                        case "foo bar":
                                                            var v64 uint = v4
                                                            var v65 int8 = 65
                                                            _ = v64
                                                            _ = v65
                                                        }
                                                        v66 := GenS4{F0: 18446744073709551615, F1: (fi55 / ((32767) | 1)), F2: 4294967295}
                                                        _ = v66
                                                    }
                                                    v67 := []int{}
                                                    var v68 int8 = -1
                                                    _ = v67
                                                    _ = v68
                                                case 128:
                                                    fmt.Println(v2, v4, v2)
                                                }
                                                var v69 uint64 = 4294967295
                                                _ = v69
                                            }
                                            _ = fi55
                                        }
                                    }
                                    v70 := []int{}
                                    if v5 {
                                        switch fi54 {
                                        case 2147483647:
                                            var v71 string = ("\n\t" + ("\n\t" + "\n\t"))
                                            _ = v71
                                        case 1:
                                            if (-419 > 327) {
                                                var v72 uint64 = 18446744073709551615
                                                _ = v72
                                            }
                                        default:
                                            var v73 string = "a"
                                            _ = v73
                                        }
                                        v41 = 1.5
                                        var v74 uint = v4
                                        _ = v74
                                    }
                                    _ = v70
                                } else {
                                    for fi75 := 0; fi75 < 9; fi75++ {
                                        if true {
                                            var v76 string = ""
                                            _ = v76
                                        }
                                        fmt.Println(v12)
                                        _ = fi75
                                    }
                                }
                            case "foo bar":
                                fmt.Println(v12, fi54)
                            case "hello":
                                for fi77 := 0; fi77 < 1; fi77++ {
                                    var v78 uint = 65535
                                    _ = fi77
                                    _ = v78
                                }
                            }
                            for fi79 := 0; fi79 < 3; fi79++ {
                                if v5 {
                                    if v5 {
                                        v80 := []int{-1}
                                        switch 2147483647 {
                                        case -2147483648:
                                            fmt.Println(fi79)
                                        case -72:
                                            if v5 {
                                                v53 += v53
                                                if (((v4 | (v4 - (^(v4)))) >= v4) && (!v5)) {
                                                    v80 = append(v80, fi54)
                                                } else {
                                                    v80 = append(v80, fi54)
                                                    var v81 float32 = 3.14159
                                                    v2 += 3.14159
                                                    _ = v81
                                                }
                                            }
                                        case 32768:
                                            var v82 int64 = -1
                                            _ = v82
                                        }
                                        var v83 int16 = 856
                                        v84 := GenS2{F0: 4294967295, F1: -128}
                                        _ = v80
                                        _ = v83
                                        _ = v84
                                    }
                                }
                                v4 += 0
                                _ = fi79
                            }
                            _ = fi54
                        }
                        switch "" {
                        case "\n\t":
                            for fi85 := 0; fi85 < 11; fi85++ {
                                v12 = (v53 + v12)
                                _ = fi85
                            }
                        }
                        _ = v53
                    }
                case 601:
                    v86 := []int{}
                    _ = v86
                default:
                    v87 := GenS1{F0: v3}
                    _ = v87
                }
                for fi88 := 0; fi88 < 3; fi88++ {
                    fmt.Println(v3)
                    switch "" {
                    case "":
                        v2 += v2
                    case "hello":
                        v3 = uint8(v1)
                    default:
                        var v89 int16 = v12
                        _ = v89
                    }
                    v42 = uint16(v1)
                    _ = fi88
                }
            } else {
                v90 := GenS4{F0: (^(v4)), F1: 0, F2: 42}
                switch ("hello" + ((("\n\t" + "a") + "foo bar") + "foo bar")) {
                case "":
                    var v91 uint16 = (v42 ^ (v42 % (((v42 & v42)) | 1)))
                    _ = v91
                default:
                    v92 := GenS5{F0: (v42 - v42), F1: 236}
                    _ = v92
                }
                _ = v90
            }
        }
    }
    v2 += -2.25
    if v5 {
        if (!v5) {
            v93 := []int{int(v12), 32767, 4294967295}
            _ = v93
        } else {
            v5 = (!v5)
            var v94 uint64 = 1
            _ = v94
        }
        v95 := []int{2147483647, 214}
        switch "a" {
        case "a":
            fmt.Println(v42)
        case "\n\t":
            if (v5 || (false && (false || true))) {
                if false {
                    var v96 uint64 = uint64(v42)
                    _ = v96
                } else {
                    switch "a" {
                    case "":
                        if v5 {
                            if v5 {
                                v4 = uint(v1)
                                for fi97 := 0; fi97 < 9; fi97++ {
                                    var v98 float64 = (((float64((v41 * ((1.5 * (v41 + v41)) + float32((1e10 / (v2 / ((((((((float64(-2.25) * v2) * v2) - 1.5) + 1.5) / (v2 - v2)) + ((v2 + v2) / (-2.25 - v2))) * 0.1) / 1e10))))))) * (1e10 - v2)) * 3.14159) - (1.5 - v2))
                                    var v99 int8 = -126
                                    _ = fi97
                                    _ = v98
                                    _ = v99
                                }
                            }
                            for fi100 := 0; fi100 < 9; fi100++ {
                                v12 = 1
                                _ = fi100
                            }
                            if v5 {
                                if len(v95) >= 2 {
                                    v95 = v95[1:len(v95)]
                                }
                            }
                        } else {
                            v41 += 0.1
                        }
                    case "hello":
                        fmt.Println(v42, v41)
                    case "foo bar":
                        v5 = (2 < (v3 & (v3 >> (((v4 & (v4 / ((18446744073709551615) | 1)))) & 63))))
                    }
                }
                switch 128 {
                case 668:
                    v12 += (v12 << ((uint(v42)) & 63))
                case 32767:
                    v95 = append(v95, 9223372036854775807)
                default:
                    fmt.Println(v2, v42, v12)
                }
            }
        default:
            for fi101 := 0; fi101 < 13; fi101++ {
                var v102 string = "\n\t"
                var v103 int = fi101
                _ = fi101
                _ = v102
                _ = v103
            }
        }
        _ = v95
    }
    if ((^(v4)) == (v4 - (v4 ^ 18446744073709551615))) {
        var v104 float64 = v2
        v3 = v3
        var v105 uint64 = 215
        _ = v104
        _ = v105
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
