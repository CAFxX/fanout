package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 uint; F2 int32 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int64; F1 int8; F2 string }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint16; F1 string }

func main() {
    var v1 int8 = 1
    var v2 uint16 = 255
    var v3 uint16 = v2
    v2 = v3
    switch 32768 {
    case 32767:
        if ((716 <= uint(v2)) && (501 != 209)) {
            for fi4 := 0; fi4 < 20; fi4++ {
                var v5 uint64 = 18446744073709551615
                _ = fi4
                _ = v5
            }
            fmt.Println(v1, v3, v1)
        } else {
            var v6 int32 = 0
            var v7 float64 = float64((0.1 * 0.1))
            _ = v6
            _ = v7
        }
    }
    fmt.Println(v1, v3)
    v8 := GenS1{F0: 42, F1: 4294967295, F2: 1}
    v3 = v3
    v9 := GenS1{F0: -1, F1: 65535, F2: 841}
    if (-32768 >= 1000) {
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 int64 = 42
            v12 := GenS2{F0: fi10}
            var v13 int64 = v11
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
    } else {
        v14 := GenS1{F0: 9223372036854775807, F1: 1, F2: 0}
        if true {
            v1 += (v1 | (v1 | 100))
            var v15 bool = false
            for fi16 := 0; fi16 < 1; fi16++ {
                var v17 uint8 = 42
                _ = fi16
                _ = v17
            }
            _ = v15
        } else {
            for fi18 := 0; fi18 < 8; fi18++ {
                fmt.Println(v3, v1)
                for fi19 := 0; fi19 < 11; fi19++ {
                    switch "\n\t" {
                    case "a":
                        switch ("a" + ("foo bar" + "hello")) {
                        case "hello":
                            var v20 uint8 = 128
                            _ = v20
                        case "a":
                            if false {
                                v21 := []int{fi18}
                                _ = v21
                            }
                            v1 -= (-(v1))
                        case "foo bar":
                            switch ("foo bar" + ("\n\t" + "\n\t")) {
                            case "hello":
                                fmt.Println(v3, v1)
                            case "\n\t":
                                var v22 bool = false
                                _ = v22
                            case "foo bar":
                                fmt.Println(v1)
                            default:
                                var v23 uint32 = 2147483648
                                _ = v23
                            }
                        default:
                            v24 := GenS1{F0: 0, F1: 174, F2: 4294967295}
                            _ = v24
                        }
                    case "hello":
                        v25 := GenS1{F0: 42, F1: 944, F2: 1}
                        _ = v25
                    case "\n\t":
                        var v26 uint32 = 325
                        _ = v26
                    default:
                        fmt.Println(v3, fi19)
                    }
                    v27 := GenS3{F0: int64(v2), F1: 777, F2: 111}
                    if (!true) {
                        v3 -= (v2 << ((255) & 63))
                        var v28 int = (^(fi18))
                        v29 := []int{fi19}
                        _ = v28
                        _ = v29
                    } else {
                        if false {
                            v30 := []int{(fi19 >> ((255) & 63)), (fi19 / ((fi19) | 1))}
                            _ = v30
                        }
                        fmt.Println(v1, v2)
                        if true {
                            if (456 <= 275) {
                                v31 := []int{(-(fi18))}
                                fmt.Println(v1)
                                _ = v31
                            } else {
                                v32 := GenS3{F0: -9223372036854775808, F1: 370, F2: -46}
                                v33 := []int{(fi19 & fi19), (-(fi18)), fi19, fi19}
                                v2 = v3
                                _ = v32
                                _ = v33
                            }
                            if false {
                                v3 -= v2
                                v34 := []int{fi18, -410, 32768, fi19}
                                _ = v34
                            }
                            v1 = 62
                        }
                    }
                    _ = fi19
                    _ = v27
                }
                var v35 int64 = -1
                _ = fi18
                _ = v35
            }
            if false {
                var v36 uint64 = 4294967295
                var v37 uint = 371
                v38 := GenS2{F0: 9223372036854775807}
                _ = v36
                _ = v37
                _ = v38
            } else {
                var v39 string = ""
                fmt.Println(v3, v39, v1)
                v1 = v1
                _ = v39
            }
        }
        v1 = (v1 + (v1 | v1))
        _ = v14
    }
    if (205 >= 0) {
        v3 -= (v2 / (((v2 * v3)) | 1))
    } else {
        if (87 != 42) {
            if false {
                v3 = v2
                v40 := []int{-860}
                v1 += v1
                _ = v40
            } else {
                for fi41 := 0; fi41 < 9; fi41++ {
                    var v42 float32 = 0.1
                    var v43 float32 = (float32((-2.25 - 1e10)) - float32(1.5))
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                switch (("hello" + "a") + "foo bar") {
                case "hello":
                    switch "" {
                    case "\n\t":
                        for fi44 := 0; fi44 < 8; fi44++ {
                            v45 := GenS1{F0: 9223372036854775807, F1: 572, F2: 0}
                            var v46 float64 = 1e10
                            v47 := []int{-266, fi44}
                            _ = fi44
                            _ = v45
                            _ = v46
                            _ = v47
                        }
                    case "a":
                        v2 = v3
                    case "hello":
                        var v48 int16 = 42
                        _ = v48
                    default:
                        var v49 uint8 = 160
                        _ = v49
                    }
                case "foo bar":
                    var v50 int8 = (v1 - (v1 | (v1 ^ (v1 / (((v1 | v1)) | 1)))))
                    _ = v50
                case "a":
                    for fi51 := 0; fi51 < 14; fi51++ {
                        v52 := []int{42}
                        _ = fi51
                        _ = v52
                    }
                default:
                    for fi53 := 0; fi53 < 7; fi53++ {
                        if true {
                            v54 := []int{-1, int(v1), fi53, fi53}
                            if true {
                                if false {
                                    v55 := GenS4{F0: fi53, F1: -1}
                                    if (1 <= 0) {
                                        v2 = (v2 * 32767)
                                        fmt.Println(v3, v2)
                                        if true {
                                            v2 = v3
                                            if (!(!false)) {
                                                v1 = v1
                                                switch "" {
                                                case "":
                                                    var v56 uint = 313
                                                    _ = v56
                                                case "\n\t":
                                                    v2 = (v3 >> ((18446744073709551615) & 63))
                                                }
                                                switch (-(fi53)) {
                                                case 42:
                                                    v57 := GenS3{F0: 42, F1: 1, F2: 892}
                                                    _ = v57
                                                }
                                            }
                                        }
                                    } else {
                                        var v58 int32 = -757
                                        fmt.Println(v3, v1, v1)
                                        _ = v58
                                    }
                                    if len(v54) >= 2 {
                                        v54 = v54[1:len(v54)]
                                    }
                                    _ = v55
                                } else {
                                    var v59 bool = (fi53 <= fi53)
                                    _ = v59
                                }
                            } else {
                                v60 := GenS5{F0: -761}
                                fmt.Println(v2)
                                v61 := GenS6{F0: (((("\n\t" + "hello") + ("\n\t" + "")) + ("\n\t" + "\n\t")) + "a")}
                                _ = v60
                                _ = v61
                            }
                            _ = v54
                        }
                        _ = fi53
                    }
                }
                if true {
                    var v62 int16 = -265
                    var v63 string = "\n\t"
                    v64 := GenS7{F0: 9223372036854775807, F1: v1, F2: ("hello" + (v63 + ("a" + "\n\t")))}
                    _ = v62
                    _ = v63
                    _ = v64
                } else {
                    for fi65 := 0; fi65 < 15; fi65++ {
                        var v66 int8 = (^(v1))
                        v66 = 44
                        switch "hello" {
                        case "foo bar":
                            var v67 bool = false
                            _ = v67
                        default:
                            v1 += v1
                        }
                        _ = fi65
                        _ = v66
                    }
                }
            }
            fmt.Println(v2, v2)
            switch -386 {
            case -9223372036854775808:
                for fi68 := 0; fi68 < 5; fi68++ {
                    if ((false || false) && false) {
                        v3 = v2
                        if ((65535 < 42) && true) {
                            var v69 uint32 = 2147483648
                            switch (fi68 ^ (fi68 / (((fi68 & fi68)) | 1))) {
                            case 622:
                                switch int(v1) {
                                case 4294967295:
                                    fmt.Println(fi68)
                                }
                            case 128:
                                for fi70 := 0; fi70 < 13; fi70++ {
                                    v71 := GenS7{F0: 9223372036854775807, F1: (v1 - 42), F2: (("a" + "hello") + "foo bar")}
                                    fmt.Println(fi70, v1)
                                    v72 := GenS4{F0: fi70, F1: -32768}
                                    if (112 < 1) {
                                        v69 = 65535
                                        var v73 int64 = -1
                                        _ = v73
                                    }
                                    _ = fi70
                                    _ = v71
                                    _ = v72
                                }
                            default:
                                v2 += 2
                            }
                            _ = v69
                        } else {
                            switch (^(fi68)) {
                            case 0:
                                switch (^(fi68)) {
                                case 511:
                                    fmt.Println(v3)
                                case -9223372036854775808:
                                    v74 := GenS8{F0: "\n\t"}
                                    _ = v74
                                }
                            }
                            v75 := GenS5{F0: 42}
                            v76 := []int{fi68}
                            _ = v75
                            _ = v76
                        }
                    } else {
                        var v77 float64 = 1.5
                        _ = v77
                    }
                    v78 := []int{fi68, (fi68 << ((792) & 63)), fi68}
                    _ = fi68
                    _ = v78
                }
            case 1:
                for fi79 := 0; fi79 < 10; fi79++ {
                    v1 = 127
                    _ = fi79
                }
            default:
                var v80 float64 = (((1.5 + float64(-2.25)) - 3.14159) * 1.5)
                _ = v80
            }
        } else {
            fmt.Println(v2)
            var v81 uint8 = 2
            _ = v81
        }
        switch ((("foo bar" + "a") + "a") + "hello") {
        case "foo bar":
            v3 -= 1
        default:
            if true {
                v82 := GenS7{F0: 0, F1: (v1 * v1), F2: "foo bar"}
                if false {
                    v3 -= 475
                    v83 := GenS2{F0: -865}
                    var v84 int32 = 1
                    _ = v83
                    _ = v84
                } else {
                    var v85 int64 = -825
                    v2 -= v3
                    if (!true) {
                        fmt.Println(v85, v85, v2)
                        v3 = 0
                        v3 -= 2
                    }
                    _ = v85
                }
                for fi86 := 0; fi86 < 3; fi86++ {
                    v87 := []int{fi86, -2147483648, int(v1), (fi86 ^ fi86)}
                    _ = fi86
                    _ = v87
                }
                _ = v82
            } else {
                var v88 bool = (false || true)
                v89 := GenS9{F0: (v2 / ((v3) | 1)), F1: "foo bar"}
                fmt.Println(v1)
                _ = v88
                _ = v89
            }
        }
        for fi90 := 0; fi90 < 3; fi90++ {
            var v91 uint = 2
            _ = fi90
            _ = v91
        }
    }
    for fi92 := 0; fi92 < 17; fi92++ {
        if (-2147483648 < -1) {
            for fi93 := 0; fi93 < 5; fi93++ {
                if false {
                    v1 = (v1 >> ((255) & 63))
                    v94 := []int{fi93, fi93}
                    _ = v94
                }
                _ = fi93
            }
        } else {
            v3 += v3
        }
        _ = fi92
    }
    fmt.Println(v2, v3, v2)
    var v95 string = ("\n\t" + "\n\t")
    v1 = 100
    var v96 int64 = -9223372036854775808
    var v97 uint8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
    _ = v9
    _ = v95
    _ = v96
    _ = v97
}
