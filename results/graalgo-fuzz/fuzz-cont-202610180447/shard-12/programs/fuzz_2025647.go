package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 int64; F2 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 bool; F1 bool }
type GenS6 struct { F0 string; F1 uint64 }
func genFn0() uint8 {
    if false {
        var v5 bool = (0 == 42)
        v5 = true
        v6 := GenS1{F0: 0}
        _ = v5
        _ = v6
    } else {
        if false {
            v7 := []int{-1}
            for fi8 := 0; fi8 < 3; fi8++ {
                var v9 uint16 = 42
                _ = fi8
                _ = v9
            }
            _ = v7
        } else {
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 int32 = -1
                v11 = (v11 % (((v11 << ((4294967295) & 63))) | 1))
                _ = fi10
                _ = v11
            }
            if true {
                var v12 float64 = 0.0
                v13 := GenS1{F0: 100}
                _ = v12
                _ = v13
            }
        }
        if (true && (18446744073709551615 <= 0)) {
            if false {
                var v14 int16 = -875
                switch 32767 {
                case 127:
                    for fi15 := 0; fi15 < 14; fi15++ {
                        if (-2.25 <= 0.0) {
                            if (((3.14159 * 1e10) / float64(float32(float64((1.5 - (0.1 + 1.5)))))) > float64((1e10 * 3.14159))) {
                                switch 976 {
                                case 4294967295:
                                    v16 := []int{}
                                    _ = v16
                                default:
                                    v14 = 32767
                                }
                            }
                            if true {
                                switch ("foo bar" + "") {
                                case "":
                                    v17 := []int{(fi15 << ((4294967295) & 63))}
                                    _ = v17
                                case "hello":
                                    var v18 uint8 = 18
                                    _ = v18
                                }
                                v14 = 1
                                switch fi15 {
                                case -345:
                                    v14 += v14
                                case 42:
                                    v19 := GenS1{F0: 127}
                                    _ = v19
                                }
                            }
                        } else {
                            v14 = (v14 << ((1) & 63))
                            v14 = (v14 - v14)
                        }
                        v20 := []int{fi15, (fi15 << ((1) & 63))}
                        for fi21 := 0; fi21 < 20; fi21++ {
                            var v22 int32 = -214
                            _ = fi21
                            _ = v22
                        }
                        _ = fi15
                        _ = v20
                    }
                default:
                    if true {
                        v23 := GenS1{F0: -1}
                        _ = v23
                    }
                }
                _ = v14
            }
            switch (("\n\t" + "hello") + "hello") {
            case "\n\t":
                var v24 uint64 = 2
                _ = v24
            default:
                v25 := GenS1{F0: 0}
                _ = v25
            }
            var v26 float64 = (3.14159 / (3.14159 / (-2.25 / 1.0)))
            _ = v26
        }
        if true {
            if true {
                for fi27 := 0; fi27 < 1; fi27++ {
                    if false {
                        if false {
                            fmt.Println(fi27)
                            var v28 int8 = -1
                            var v29 int64 = 1
                            _ = v28
                            _ = v29
                        }
                        if true {
                            if true {
                                var v30 uint = 790
                                _ = v30
                            }
                            for fi31 := 0; fi31 < 5; fi31++ {
                                var v32 int32 = -1
                                _ = fi31
                                _ = v32
                            }
                        }
                    } else {
                        v33 := []int{fi27, (fi27 * -699), (-(fi27))}
                        var v34 int64 = 938
                        _ = v33
                        _ = v34
                    }
                    var v35 bool = true
                    var v36 uint = 0
                    _ = fi27
                    _ = v35
                    _ = v36
                }
            }
            var v37 uint32 = 128
            var v38 bool = true
            _ = v37
            _ = v38
        } else {
            if ((((-2.25 - (0.0 * (1e10 + ((3.14159 * ((3.14159 * 1.5) - float32(float64(3.14159)))) - (3.14159 - float32(0.0)))))) + ((1e10 / 1.5) / -2.25)) == (0.1 * (0.1 / 1.0))) && (!(true && false))) {
                for fi39 := 0; fi39 < 17; fi39++ {
                    var v40 uint = 647
                    _ = fi39
                    _ = v40
                }
                var v41 int = 4294967295
                _ = v41
            }
        }
    }
    v42 := GenS2{F0: (104 > 2), F1: -1, F2: 2}
    switch "a" {
    case "a":
        v43 := GenS3{F0: (true && (!false))}
        _ = v43
    case "":
        v44 := []int{-9223372036854775808, 42, -1, 9223372036854775807}
        _ = v44
    case "hello":
        v45 := GenS2{F0: true, F1: 820, F2: 4294967295}
        _ = v45
    default:
        v46 := GenS4{F0: 32768, F1: (2 >= 42)}
        _ = v46
    }
    _ = v42
    return 37
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[string]bool)
    var v3 bool = (-511 >= -105)
    var v4 float64 = float64(0.0)
    if v3 {
        var v47 int = -9223372036854775808
        _ = v47
    }
    v48 := GenS5{F0: (!(1 != -7)), F1: (968 != 42)}
    if true {
        if v3 {
            var v49 uint = 782
            for fi50 := 0; fi50 < 20; fi50++ {
                for fi51 := 0; fi51 < 4; fi51++ {
                    fmt.Println(fi50, v49)
                    delete(v2, "foo bar")
                    _ = fi51
                }
                v1 += (v1 << ((v49) & 63))
                v52 := GenS6{F0: "a", F1: 396}
                _ = fi50
                _ = v52
            }
            _ = len(v2)
            _ = v49
        } else {
            v4 += float64(-2.25)
            v53 := []int{1, -976}
            _ = v53
        }
        if false {
            fmt.Println(v3, v3, v1)
            v3 = v3
            fmt.Println(v1, v3)
        } else {
            for fi54 := 0; fi54 < 3; fi54++ {
                fmt.Println(v4, v4, v4)
                v4 = 0.0
                _ = fi54
            }
            if v3 {
                switch 1 {
                case 265:
                    v55 := GenS1{F0: 98}
                    _ = v55
                case 4294967295:
                    for fi56 := 0; fi56 < 11; fi56++ {
                        var v57 uint16 = 0
                        _ = fi56
                        _ = v57
                    }
                case 0:
                    v58 := GenS5{F0: (!(-2.25 >= float32(1.5))), F1: v3}
                    _ = v58
                }
                _ = len(v2)
                v59 := []int{1, -635, 32768, 32768}
                _ = v59
            }
            fmt.Println(v1, v3)
        }
        switch 32768 {
        case -689:
            var v60 bool = ((v3 || (true && (v4 > v4))) || false)
            _ = v60
        case 2147483647:
            fmt.Println(v4, v1)
        case -527:
            v61 := []int{726, -789, -2147483648}
            _ = v61
        }
    } else {
        switch "hello" {
        case "foo bar":
            var v62 float64 = float64(-2.25)
            _ = v62
        default:
            switch 32768 {
            case 428:
                for fi63 := 0; fi63 < 8; fi63++ {
                    var v64 bool = true
                    _ = fi63
                    _ = v64
                }
            case -242:
                v65 := v2[""]
                _ = v65
            }
        }
        fmt.Println(v4, v1)
    }
    v4 = (1e10 * v4)
    if (42 >= 2) {
        if (!v3) {
            v3 = (0 == 2)
            var v66 int8 = 127
            _ = v66
        }
    } else {
        v1 -= 42
        if ((!v3) || v3) {
            v4 += v4
        } else {
            if (v3 || (true || v3)) {
                v67 := GenS1{F0: -128}
                _ = v67
            } else {
                if v3 {
                    v68 := GenS2{F0: false, F1: 1, F2: 664}
                    var v69 uint64 = 0
                    fmt.Println(v3, v4)
                    _ = v68
                    _ = v69
                }
            }
            if false {
                for fi70 := 0; fi70 < 10; fi70++ {
                    var v71 float64 = 3.14159
                    v1 = (v1 / ((0) | 1))
                    _ = fi70
                    _ = v71
                }
            } else {
                v4 = float64(1e10)
                v72 := []int{}
                _ = v72
            }
            fmt.Println(v1)
        }
        delete(v2, "\n\t")
    }
    if (0 >= 18446744073709551615) {
        v73 := GenS4{F0: 935, F1: true}
        _ = v73
    }
    var v74 uint32 = 2
    var v75 uint32 = (^(v74))
    var v76 uint = 65535
    for fi77 := 0; fi77 < 1; fi77++ {
        v78 := GenS5{F0: ((v4 >= -2.25) || v3), F1: false}
        _ = fi77
        _ = v78
    }
    var v79 float64 = float64(0.0)
    if (!true) {
        var v80 uint32 = (v75 + v75)
        v3 = false
        _ = v80
    } else {
        if (!true) {
            var v81 float64 = (1e10 / (((v4 - ((3.14159 + 0.0) - v4)) + ((-2.25 / v4) + 1.5)) - 0.1))
            var v82 uint64 = 174
            _ = v81
            _ = v82
        }
        v83 := GenS4{F0: 128, F1: v3}
        v84 := []int{32768}
        _ = v83
        _ = v84
    }
    var v85 float32 = 0.0
    fmt.Println(v74, v76, v74)
    if (v3 && (("a" + "hello") >= "foo bar")) {
        v86 := GenS1{F0: -57}
        fmt.Println(v79, v74, v74)
        _ = v86
    }
    if false {
        v75 += v74
        fmt.Println(v76)
        for fi87 := 0; fi87 < 13; fi87++ {
            fmt.Println(v1)
            delete(v2, "\n\t")
            _ = fi87
        }
    }
    v1 += (v1 ^ genFn0())
    for fi88 := 0; fi88 < 12; fi88++ {
        var v89 string = ("hello" + "foo bar")
        v90 := []int{-574, 32767}
        switch (^(fi88)) {
        case 2147483647:
            var v91 int16 = 537
            _ = v91
        case 4294967295:
            for fi92 := 0; fi92 < 8; fi92++ {
                v75 = 65535
                if (53 <= 9223372036854775807) {
                    v79 += float64((3.14159 * v85))
                }
                if v3 {
                    var v93 uint16 = 42
                    var v94 string = v89
                    _ = v93
                    _ = v94
                } else {
                    v95 := GenS3{F0: v3}
                    _ = v95
                }
                _ = fi92
            }
        case -2147483648:
            var v96 int32 = -868
            _ = v96
        }
        _ = fi88
        _ = v89
        _ = v90
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v79:", v79)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v74
    _ = v75
    _ = v76
    _ = v79
    _ = v85
}
