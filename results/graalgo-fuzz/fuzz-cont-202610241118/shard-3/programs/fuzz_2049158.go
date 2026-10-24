package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 uint8; F2 int64 }
type GenS4 struct { F0 uint16; F1 int8; F2 int }
type GenS5 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS6 struct { F0 uint8; F1 string; F2 int16 }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 int16; F1 string; F2 bool }
type GenS9 struct { F0 int8; F1 int16 }
type GenS10 struct { F0 int16; F1 int8; F2 uint64 }
type GenS11 struct { F0 int32; F1 uint8; F2 uint32 }

func main() {
    var v1 string = (("\n\t" + ("hello" + "foo bar")) + "a")
    var v2 int32 = 750
    var v3 uint8 = 42
    v4 := make(map[int]bool)
    var v5 uint64 = 2
    if (!true) {
        switch 365 {
        case 4294967295:
            for fi6 := 0; fi6 < 6; fi6++ {
                v7 := GenS1{F0: (fi6 / (((fi6 * -674)) | 1)), F1: 42}
                v1 = ((v1 + (v1 + (v1 + ((v1 + v1) + (v1 + v1))))) + (v1 + v1))
                _ = fi6
                _ = v7
            }
        default:
            v8 := []int{153}
            _ = v8
        }
        v9 := []int{9223372036854775807, 4294967295}
        _ = v9
    } else {
        if (1 <= 2) {
            if true {
                var v10 int = 864
                if false {
                    v11 := GenS1{F0: (v10 ^ int(v5)), F1: 42}
                    if true {
                        var v12 string = v1
                        _ = v12
                    } else {
                        var v13 float32 = (0.0 + 3.14159)
                        v2 = (v2 * 575)
                        if (1 == 221) {
                            fmt.Println(v2, v2, v10)
                            var v14 bool = (!(true && (true || (796 != 32767))))
                            _ = v14
                        } else {
                            v15 := []int{v10, (v10 << ((uint(v3)) & 63)), (^(v10)), v10}
                            _ = v15
                        }
                        _ = v13
                    }
                    _ = v11
                } else {
                    v16 := []int{(v10 % ((2147483647) | 1)), v10, v10, v10}
                    _ = v16
                }
                v17 := GenS1{F0: v10, F1: 446}
                _ = v10
                _ = v17
            } else {
                v2 = -117
                if (3.14159 != 1e10) {
                    if false {
                        if (false && false) {
                            fmt.Println(v3, v5)
                            if (!false) {
                                var v18 int16 = 579
                                switch v1 {
                                case "":
                                    v19 := GenS1{F0: 531, F1: 454}
                                    _ = v19
                                case "\n\t":
                                    v18 = v18
                                }
                                var v20 uint16 = 65535
                                _ = v18
                                _ = v20
                            }
                        } else {
                            v21 := GenS2{F0: 1}
                            _ = v21
                        }
                        var v22 float64 = (float64(1e10) / -2.25)
                        for fi23 := 0; fi23 < 20; fi23++ {
                            v24 := []int{42}
                            if false {
                                var v25 uint8 = v3
                                switch 32768 {
                                case 128:
                                    switch ("\n\t" + v1) {
                                    case "hello":
                                        var v26 int = (fi23 ^ 0)
                                        _ = v26
                                    case "\n\t":
                                        v3 = 42
                                    }
                                default:
                                    v1 = "a"
                                }
                                _ = v25
                            }
                            switch 2147483647 {
                            case 32767:
                                v27 := v4[fi23]
                                _ = v27
                            case 127:
                                switch (fi23 + (fi23 % ((int(v3)) | 1))) {
                                case 736:
                                    for fi28 := 0; fi28 < 10; fi28++ {
                                        v29 := GenS3{F0: v1, F1: v3, F2: 575}
                                        _ = fi28
                                        _ = v29
                                    }
                                case 9223372036854775807:
                                    var v30 int16 = 191
                                    _ = v30
                                case -175:
                                    for fi31 := 0; fi31 < 6; fi31++ {
                                        v32 := GenS2{F0: 32767}
                                        _ = fi31
                                        _ = v32
                                    }
                                default:
                                    v22 = (v22 / v22)
                                }
                            }
                            _ = fi23
                            _ = v24
                        }
                        _ = v22
                    } else {
                        v33 := GenS2{F0: 32767}
                        _ = v33
                    }
                    if ((true && false) && false) {
                        fmt.Println(v3, v2)
                        v3 = v3
                        var v34 uint64 = 502
                        _ = v34
                    } else {
                        v35 := GenS4{F0: 2, F1: 100, F2: -185}
                        _ = v35
                    }
                }
                for fi36 := 0; fi36 < 3; fi36++ {
                    for fi37 := 0; fi37 < 14; fi37++ {
                        fmt.Println(v1)
                        _ = fi37
                    }
                    v1 = v1
                    _ = fi36
                }
            }
            fmt.Println(v1, v1)
            v3 = (v3 % (((^(v3))) | 1))
        }
        if false {
            v1 = ("hello" + "\n\t")
            v38 := GenS4{F0: 113, F1: -1, F2: 1}
            _ = v38
        } else {
            for fi39 := 0; fi39 < 17; fi39++ {
                v40 := []int{(fi39 | (-(fi39)))}
                for fi41 := 0; fi41 < 5; fi41++ {
                    var v42 bool = false
                    v43 := v4[fi41]
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                var v44 uint32 = 2147483648
                _ = fi39
                _ = v40
                _ = v44
            }
            if true {
                v45 := GenS4{F0: 64, F1: int8(v3), F2: 32767}
                if (true || (900 >= 844)) {
                    if true {
                        fmt.Println(v3)
                        fmt.Println(v1, v1, v5)
                    } else {
                        switch 4294967295 {
                        case -79:
                            for fi46 := 0; fi46 < 14; fi46++ {
                                v47 := GenS5{F0: (v3 ^ 255), F1: (v2 - (v2 ^ v2)), F2: 65535}
                                _ = fi46
                                _ = v47
                            }
                        case -9223372036854775808:
                            v48 := []int{int(v2), 0, 32767, -2147483648}
                            _ = v48
                        default:
                            v3 = v3
                        }
                        var v49 float32 = 3.14159
                        var v50 uint = uint(v3)
                        _ = v49
                        _ = v50
                    }
                    var v51 bool = false
                    _ = v51
                } else {
                    var v52 int8 = 42
                    v52 = (v52 - (v52 + v52))
                    _ = v52
                }
                if (!(v5 > 4294967295)) {
                    fmt.Println(v2, v3, v2)
                }
                _ = v45
            }
            v53 := []int{}
            _ = v53
        }
        var v54 float32 = 3.14159
        _ = v54
    }
    switch -2147483648 {
    case 275:
        fmt.Println(v2, v5, v1)
    }
    v55 := GenS6{F0: (v3 - v3), F1: (v1 + v1), F2: -32768}
    if (!(1 >= 797)) {
        var v56 uint32 = uint32(v3)
        var v57 uint32 = v56
        _ = v56
        _ = v57
    } else {
        if false {
            if false {
                v5 = v5
                if (!false) {
                    fmt.Println(v1, v5)
                } else {
                    fmt.Println(v2, v1)
                    v58 := GenS3{F0: ("hello" + v1), F1: (v3 << ((2) & 63)), F2: -9223372036854775808}
                    v59 := []int{-2147483648, -1}
                    _ = v58
                    _ = v59
                }
            } else {
                fmt.Println(v3, v5)
                v4[21] = ((float64(1.5) / 1.0) <= 0.0)
            }
        }
        switch 9223372036854775807 {
        case 32768:
            var v60 int16 = 336
            _ = v60
        case 83:
            if false {
                v61 := []int{32767, -2147483648}
                _ = v61
            } else {
                var v62 uint32 = 213
                for fi63 := 0; fi63 < 3; fi63++ {
                    for fi64 := 0; fi64 < 13; fi64++ {
                        v3 += 27
                        var v65 int = 4294967295
                        _ = fi64
                        _ = v65
                    }
                    _ = fi63
                }
                v5 -= (v5 / ((213) | 1))
                _ = v62
            }
        default:
            v66 := GenS6{F0: v3, F1: v1, F2: -394}
            _ = v66
        }
        switch 113 {
        case 9223372036854775807:
            for fi67 := 0; fi67 < 7; fi67++ {
                v68 := []int{}
                v69 := v4[fi67]
                delete(v4, 32768)
                _ = fi67
                _ = v68
                _ = v69
            }
        case 128:
            v70 := []int{-1, 127}
            _ = v70
        default:
            switch 1 {
            case 42:
                if false {
                    v71 := GenS7{F0: 2147483647, F1: 127}
                    _ = v71
                }
            }
        }
    }
    var v72 int8 = 100
    switch 4294967295 {
    case -1:
        switch ("" + "a") {
        case "\n\t":
            v5 += 507
        case "hello":
            if false {
                var v73 int16 = 95
                if true {
                    var v74 uint32 = 213
                    v74 = 0
                    v75 := GenS2{F0: v73}
                    _ = v74
                    _ = v75
                }
                _ = v73
            }
        case "a":
            for fi76 := 0; fi76 < 9; fi76++ {
                if (3.14159 < float32((1e10 + (1e10 - 1e10)))) {
                    v77 := GenS6{F0: v3, F1: v1, F2: -32768}
                    delete(v4, fi76)
                    _ = v77
                } else {
                    fmt.Println(fi76, v5, v72)
                    v5 -= (v5 * v5)
                    switch (v1 + v1) {
                    case "":
                        v72 = (v72 >> ((406) & 63))
                    }
                }
                _ = fi76
            }
        }
    }
    v5 = v5
    for fi78 := 0; fi78 < 10; fi78++ {
        fmt.Println(v1, v3, v2)
        v3 = (v3 + (v3 << ((374) & 63)))
        _ = fi78
    }
    v79 := GenS7{F0: 901, F1: 162}
    fmt.Println(v5, v1, v72)
    var v80 int8 = v72
    _ = len(v4)
    fmt.Println(v5, v80)
    if true {
        for fi81 := 0; fi81 < 16; fi81++ {
            fmt.Println(v3, v72)
            fmt.Println(v2, v80, v1)
            switch v1 {
            case "\n\t":
                switch fi81 {
                case 32768:
                    v82 := GenS8{F0: 1, F1: v1, F2: true}
                    _ = v82
                }
            }
            _ = fi81
        }
    } else {
        var v83 float64 = (1e10 / 1e10)
        _ = v83
    }
    if false {
        v3 += (v3 << ((65535) & 63))
        v1 = v1
    } else {
        var v84 int32 = int32(v72)
        for fi85 := 0; fi85 < 1; fi85++ {
            v5 = (v5 << ((0) & 63))
            if false {
                for fi86 := 0; fi86 < 5; fi86++ {
                    v87 := GenS9{F0: 100, F1: 32767}
                    _ = fi86
                    _ = v87
                }
                v88 := GenS10{F0: 32767, F1: (v72 % ((v80) | 1)), F2: (v5 % (((v5 << ((899) & 63))) | 1))}
                _ = v88
            }
            _ = fi85
        }
        _ = v84
    }
    v89 := v4[127]
    v90 := GenS11{F0: v2, F1: (^(v3)), F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v72:", v72)
    fmt.Println("v80:", v80)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v72
    _ = v79
    _ = v80
    _ = v89
    _ = v90
}
