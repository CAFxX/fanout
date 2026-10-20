package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int; F1 uint16; F2 uint }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool; F1 int }
type GenS8 struct { F0 int32; F1 int8; F2 uint64 }
type GenS9 struct { F0 string; F1 uint64; F2 int64 }
func genFn0() string {
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 uint8 = 255
            v8 += (^(v8))
            _ = fi7
            _ = v8
        }
        var v9 uint64 = 2
        for fi10 := 0; fi10 < 12; fi10++ {
            v11 := GenS1{F0: 1, F1: 619}
            switch -2147483648 {
            case 145:
                if true {
                    v12 := GenS2{F0: 632, F1: fi6}
                    _ = v12
                }
            case 591:
                var v13 int16 = 676
                _ = v13
            case 127:
                v14 := GenS3{F0: -1, F1: 2, F2: 0}
                _ = v14
            }
            var v15 float64 = (float64(-2.25) - 0.0)
            _ = fi10
            _ = v11
            _ = v15
        }
        _ = fi6
        _ = v9
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        v17 := []int{(fi16 * (fi16 + (fi16 << ((18446744073709551615) & 63)))), 128}
        for fi18 := 0; fi18 < 8; fi18++ {
            if (!(true && (18446744073709551615 == 1))) {
                var v19 bool = false
                var v20 int8 = -63
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                _ = v19
                _ = v20
            }
            if ((false && false) && ((!false) || ("foo bar" <= "\n\t"))) {
                fmt.Println(fi18)
                var v21 int16 = 747
                var v22 string = (("hello" + "a") + "")
                _ = v21
                _ = v22
            } else {
                var v23 int64 = 9223372036854775807
                _ = v23
            }
            _ = fi18
        }
        var v24 uint = 77
        _ = fi16
        _ = v17
        _ = v24
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        for fi26 := 0; fi26 < 16; fi26++ {
            var v27 int32 = -2147483648
            v28 := []int{fi25, 32768}
            v28 = append(v28, (fi25 << ((65535) & 63)))
            fmt.Println(fi26, fi25)
            _ = fi26
            _ = v27
            _ = v28
        }
        var v29 int = (fi25 ^ fi25)
        for fi30 := 0; fi30 < 4; fi30++ {
            switch ("hello" + "a") {
            case "hello":
                v31 := []int{(fi30 >> ((42) & 63)), fi25}
                _ = v31
            default:
                if false {
                    var v32 uint32 = 813
                    _ = v32
                }
            }
            if false {
                v33 := GenS3{F0: fi25, F1: 255, F2: 0}
                var v34 uint16 = 255
                v35 := GenS4{F0: -477, F1: (!(318 >= 2))}
                _ = v33
                _ = v34
                _ = v35
            }
            switch fi30 {
            case 127:
                switch (fi25 - fi25) {
                case 9223372036854775807:
                    v36 := []int{v29, (fi30 % ((fi30) | 1)), (fi30 + (v29 ^ (fi30 % (((fi25 + (-(fi30)))) | 1))))}
                    _ = v36
                case 621:
                    v29 += (^(fi30))
                default:
                    v37 := GenS3{F0: 9223372036854775807, F1: 65535, F2: 18446744073709551615}
                    _ = v37
                }
            default:
                switch "foo bar" {
                case "a":
                    var v38 int16 = 677
                    _ = v38
                case "":
                    fmt.Println(v29, fi25, v29)
                case "foo bar":
                    fmt.Println(v29)
                default:
                    switch "foo bar" {
                    case "a":
                        if false {
                            if false {
                                var v39 uint32 = 2147483647
                                v29 -= v29
                                _ = v39
                            } else {
                                if false {
                                    v29 -= fi25
                                } else {
                                    fmt.Println(fi25, fi25)
                                    v40 := []int{}
                                    _ = v40
                                }
                                if false {
                                    v41 := []int{4294967295, v29, (v29 | (v29 >> ((4294967295) & 63)))}
                                    v29 -= fi25
                                    _ = len(v41)
                                    _ = v41
                                }
                                v29 -= (fi30 * (fi30 & (fi30 >> ((307) & 63))))
                            }
                        }
                    case "\n\t":
                        var v42 string = ((("" + ((("a" + "\n\t") + "") + "")) + (("a" + "hello") + "")) + (("" + ("foo bar" + ((((("" + "foo bar") + "hello") + "a") + "hello") + (("\n\t" + ((("" + ("a" + "\n\t")) + "") + "")) + "hello")))) + ""))
                        _ = v42
                    case "foo bar":
                        v43 := []int{}
                        _ = v43
                    default:
                        v44 := []int{(fi25 & (v29 | v29)), fi30, (fi25 ^ (fi25 / (((fi30 << ((18446744073709551615) & 63))) | 1))), v29}
                        if ("foo bar" == "") {
                            if true {
                                if len(v44) > 0 {
                                    v44[0] = -245
                                }
                            } else {
                                if ("foo bar" == ((("a" + "") + "hello") + "foo bar")) {
                                    v45 := GenS5{F0: 2, F1: true}
                                    switch "hello" {
                                    case "a":
                                        var v46 uint8 = 2
                                        _ = v46
                                    case "\n\t":
                                        v47 := GenS3{F0: -291, F1: 65535, F2: 4294967295}
                                        _ = v47
                                    default:
                                        if false {
                                            v48 := GenS1{F0: 65535, F1: 0}
                                            if (true && true) {
                                                fmt.Println(fi30, fi30, v29)
                                                if true {
                                                    v29 = (v29 / ((fi30) | 1))
                                                } else {
                                                    var v49 uint = 4294967295
                                                    _ = v49
                                                }
                                                if len(v44) >= 2 {
                                                    v44 = v44[1:len(v44)]
                                                }
                                                var v50 uint32 = 2
                                                v50 -= (v50 + (^(v50)))
                                                fmt.Println(v50)
                                                _ = v50
                                            }
                                            switch "\n\t" {
                                            case "hello":
                                                v51 := GenS2{F0: 2147483648, F1: fi25}
                                                _ = v51
                                            case "\n\t":
                                                switch "a" {
                                                case "a":
                                                    var v52 int8 = 0
                                                    _ = v52
                                                default:
                                                    if true {
                                                        v29 = fi25
                                                    } else {
                                                        fmt.Println(v29, v29, fi30)
                                                    }
                                                }
                                            }
                                            _ = v48
                                        }
                                    }
                                    _ = v45
                                } else {
                                    v29 += 32768
                                }
                                fmt.Println(fi25, fi30)
                                var v53 int = (fi30 & 1)
                                _ = v53
                            }
                        }
                        if true {
                            if (158 > 2) {
                                if false {
                                    if len(v44) >= 2 {
                                        v44 = v44[1:len(v44)]
                                    }
                                    if false {
                                        var v54 float64 = 0.0
                                        _ = v54
                                    } else {
                                        var v55 int8 = 42
                                        _ = v55
                                    }
                                    v56 := GenS2{F0: 0, F1: fi30}
                                    _ = v56
                                } else {
                                    var v57 float64 = -2.25
                                    _ = v57
                                }
                                var v58 uint = 4294967295
                                var v59 uint64 = 4294967295
                                _ = v58
                                _ = v59
                            } else {
                                var v60 uint32 = 2147483647
                                var v61 int = -2147483648
                                v62 := GenS6{F0: (255 < 42)}
                                _ = v60
                                _ = v61
                                _ = v62
                            }
                            if len(v44) > 0 {
                                v44[0] = -2147483648
                            }
                        }
                        v63 := GenS2{F0: 161, F1: (fi30 | (fi30 - -590))}
                        _ = v44
                        _ = v63
                    }
                }
            }
            _ = fi30
        }
        _ = fi25
        _ = v29
    }
    return "foo bar"
}
func genFn1(p0 uint64, p1 bool) uint32 {
    for fi64 := 0; fi64 < 9; fi64++ {
        switch "" {
        case "\n\t":
            var v65 int16 = 914
            _ = v65
        }
        _ = fi64
    }
    fmt.Println(p0, p1)
    p1 = true
    var v66 string = "a"
    v67 := GenS7{F0: p1, F1: -9223372036854775808}
    _ = p0
    _ = p1
    _ = v66
    _ = v67
    return 586
}

func main() {
    var v1 int32 = 1
    var v2 int16 = 32767
    v3 := []uint32{515, 653, 42}
    v4 := make(map[int]int)
    var v5 int16 = v2
    var v68 bool = true
    if (!(!(!(65535 != 2147483648)))) {
        v68 = (v68 && ("foo bar" < "\n\t"))
        for fi69 := 0; fi69 < 11; fi69++ {
            v2 = v5
            _ = fi69
        }
        switch "foo bar" {
        case "":
            v70 := GenS8{F0: v1, F1: -128, F2: 2}
            _ = v70
        case "\n\t":
            switch 4294967295 {
            case -1:
                fmt.Println(v2, v1, v68)
            case 613:
                var v71 uint64 = 42
                _ = v71
            case -443:
                v72 := GenS8{F0: v1, F1: 1, F2: 860}
                _ = v72
            default:
                if v68 {
                    switch "" {
                    case "a":
                        var v73 int16 = v2
                        _ = v73
                    case "\n\t":
                        var v74 int16 = v5
                        _ = v74
                    default:
                        switch "a" {
                        case "\n\t":
                            if (!((32767 < 1) || v68)) {
                                fmt.Println(v5)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v75 uint64 = 18446744073709551615
                                _ = v75
                            } else {
                                v1 -= v1
                                var v76 int32 = v1
                                _ = v76
                            }
                        case "foo bar":
                            v77 := GenS5{F0: 32767, F1: true}
                            _ = v77
                        }
                    }
                } else {
                    _ = len(v4)
                }
            }
        case "foo bar":
            switch 128 {
            case 128:
                if ((v5 * v2) != v5) {
                    var v78 int = -759
                    fmt.Println(v1)
                    v79 := GenS9{F0: "hello", F1: 145, F2: 9223372036854775807}
                    _ = v78
                    _ = v79
                }
            }
        }
    } else {
        switch ("" + (("" + "\n\t") + "hello")) {
        case "a":
            v68 = v68
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println(v1, v68)
    var v80 int64 = 444
    switch int(v80) {
    case 127:
        fmt.Println(v68)
    case -922:
        var v81 int8 = 31
        _ = v81
    default:
        if v68 {
            v68 = (0 < 415)
        } else {
            for fi82 := 0; fi82 < 17; fi82++ {
                v83 := GenS7{F0: true, F1: (fi82 >> ((255) & 63))}
                _ = fi82
                _ = v83
            }
        }
    }
    if false {
        v84 := v4[32768]
        v5 = (v5 + -1)
        v3 = append(v3, 707)
        _ = v84
    }
    v5 -= v5
    switch ("hello" + "hello") {
    case "hello":
        v2 = -162
    case "a":
        v3 = append(v3, 0)
    }
    var v85 int64 = (v80 & v80)
    for fi86 := 0; fi86 < 19; fi86++ {
        for fi87 := 0; fi87 < 18; fi87++ {
            v88 := GenS3{F0: (fi87 & fi87), F1: 314, F2: 2}
            _ = fi87
            _ = v88
        }
        _ = fi86
    }
    fmt.Println(v68, v1, v80)
    var v89 bool = true
    _ = len(v4)
    fmt.Println(v80, v89)
    fmt.Println(v5)
    fmt.Println(v89)
    switch -2147483648 {
    case -9223372036854775808:
        v89 = (!v89)
    case 127:
        var v90 uint8 = 42
        _ = v90
    }
    switch "foo bar" {
    case "hello":
        switch 2147483647 {
        case 2147483647:
            for fi91 := 0; fi91 < 9; fi91++ {
                v92 := GenS3{F0: fi91, F1: 665, F2: 0}
                if false {
                    fmt.Println(v2)
                }
                _ = fi91
                _ = v92
            }
        default:
            var v93 uint = 18446744073709551615
            _ = v93
        }
    case "":
        v94 := GenS6{F0: false}
        _ = v94
    case "foo bar":
        var v95 float32 = ((((-2.25 * (1e10 - 0.0)) + 3.14159) - 1.5) + 3.14159)
        _ = v95
    default:
        v5 -= 32767
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v68:", v68)
    fmt.Println("v80:", v80)
    fmt.Println("v85:", v85)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v68
    _ = v80
    _ = v85
    _ = v89
}
