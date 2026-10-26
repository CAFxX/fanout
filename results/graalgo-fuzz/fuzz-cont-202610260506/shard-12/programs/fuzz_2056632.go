package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 bool; F1 uint16; F2 int64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint16; F1 bool }
type GenS10 struct { F0 uint; F1 uint64; F2 int8 }
type GenS11 struct { F0 uint64; F1 int }
func genFn0(p0 int16, p1 int) string {
    var v6 uint8 = 159
    if true {
        var v7 int8 = 8
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v6
    return "\n\t"
}
func genFn1(p0 int8) int {
    v8 := []int{4294967295, 127, -2147483648}
    var v9 uint16 = 65535
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 17; fi11++ {
            fmt.Println(v9, v9, p0)
            _ = fi11
        }
        fmt.Println(v9, p0, fi10)
        if (128 < 127) {
            if (true || true) {
                for fi12 := 0; fi12 < 9; fi12++ {
                    var v13 float64 = -2.25
                    p0 = (p0 - (p0 - -40))
                    var v14 int8 = (p0 | (^(p0)))
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            } else {
                fmt.Println(v9, v9, p0)
            }
            fmt.Println(fi10, fi10)
            for fi15 := 0; fi15 < 10; fi15++ {
                var v16 int32 = -1
                p0 = (^(p0))
                v16 = (v16 << ((58) & 63))
                _ = fi15
                _ = v16
            }
        } else {
            var v17 string = "a"
            _ = v17
        }
        _ = fi10
    }
    v9 = v9
    _ = p0
    _ = v8
    _ = v9
    return 519
}

func main() {
    var v1 float64 = 3.14159
    var v2 bool = (-727 >= 848)
    var v3 int32 = 0
    var v4 float32 = 1.5
    var v5 int64 = -675
    switch "\n\t" {
    case "foo bar":
        v18 := []int{-9223372036854775808, -2147483648, 4294967295, 2147483647}
        _ = v18
    case "a":
        var v19 int32 = 42
        _ = v19
    }
    var v20 float64 = 1e10
    v21 := GenS1{F0: 42}
    var v22 int64 = v5
    for fi23 := 0; fi23 < 9; fi23++ {
        for fi24 := 0; fi24 < 11; fi24++ {
            var v25 float32 = (0.1 - v4)
            v1 = 0.0
            v26 := GenS2{F0: 383, F1: 1}
            v22 += v22
            var v27 uint64 = 0
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 10; fi28++ {
            v29 := GenS2{F0: -658, F1: 42}
            fmt.Println(v2, v3, v2)
            _ = fi28
            _ = v29
        }
        _ = fi23
    }
    v30 := GenS3{F0: "\n\t"}
    var v31 bool = (42 == -9223372036854775808)
    v1 += float64(v4)
    if v31 {
        for fi32 := 0; fi32 < 17; fi32++ {
            for fi33 := 0; fi33 < 5; fi33++ {
                var v34 int8 = 127
                v20 = v1
                v34 = (v34 + (v34 + v34))
                _ = fi33
                _ = v34
            }
            _ = fi32
        }
        if v31 {
            v4 = v4
            v31 = (!true)
        }
        v20 -= ((float64(v4) * (((v1 * v20) - 0.1) / v20)) / v20)
    } else {
        var v35 uint = 4294967295
        var v36 uint8 = 0
        v37 := []int{-2147483648}
        _ = v35
        _ = v36
        _ = v37
    }
    switch "a" {
    case "":
        for fi38 := 0; fi38 < 3; fi38++ {
            v39 := GenS2{F0: 1000, F1: 127}
            _ = fi38
            _ = v39
        }
    case "a":
        v40 := []int{4294967295, 127}
        _ = v40
    }
    if (42 <= 65535) {
        if v2 {
            if false {
                var v41 uint16 = 83
                var v42 int8 = -77
                var v43 int = -9223372036854775808
                _ = v41
                _ = v42
                _ = v43
            } else {
                v4 = v4
                if (255 > 127) {
                    var v44 uint8 = 0
                    if (4294967295 != 2) {
                        v45 := GenS4{F0: ("hello" + "a")}
                        v46 := GenS5{F0: 42}
                        _ = v45
                        _ = v46
                    } else {
                        for fi47 := 0; fi47 < 14; fi47++ {
                            v48 := []int{genFn1(0)}
                            var v49 string = "a"
                            _ = fi47
                            _ = v48
                            _ = v49
                        }
                    }
                    _ = v44
                } else {
                    var v50 int8 = -104
                    switch "" {
                    case "a":
                        var v51 int32 = 2147483647
                        _ = v51
                    case "\n\t":
                        v52 := []int{}
                        _ = v52
                    }
                    var v53 int64 = (v22 - (v22 & (v5 | v22)))
                    _ = v50
                    _ = v53
                }
                v2 = v31
            }
            fmt.Println(v3, v2, v2)
            var v54 uint = 18446744073709551615
            _ = v54
        } else {
            if false {
                v55 := GenS5{F0: 18446744073709551615}
                for fi56 := 0; fi56 < 15; fi56++ {
                    v1 = v1
                    _ = fi56
                }
                v57 := GenS5{F0: 4294967295}
                _ = v55
                _ = v57
            } else {
                var v58 uint = 65535
                fmt.Println(v22, v31)
                _ = v58
            }
            for fi59 := 0; fi59 < 11; fi59++ {
                switch (fi59 / ((fi59) | 1)) {
                case -307:
                    fmt.Println(v20)
                case -736:
                    for fi60 := 0; fi60 < 18; fi60++ {
                        v61 := GenS6{F0: -128}
                        v3 -= (v3 >> ((0) & 63))
                        _ = fi60
                        _ = v61
                    }
                default:
                    v22 += 9223372036854775807
                }
                for fi62 := 0; fi62 < 19; fi62++ {
                    v31 = v2
                    v63 := []int{genFn1(42)}
                    v22 += v5
                    _ = fi62
                    _ = v63
                }
                v64 := GenS7{F0: v2, F1: 2, F2: v22}
                _ = fi59
                _ = v64
            }
        }
        if true {
            fmt.Println(v3, v5, v4)
            for fi65 := 0; fi65 < 20; fi65++ {
                v66 := []int{fi65}
                if v2 {
                    switch genFn1(42) {
                    case 32767:
                        switch ("\n\t" + "foo bar") {
                        case "hello":
                            var v67 uint64 = 4294967295
                            _ = v67
                        case "":
                            switch "hello" {
                            case "hello":
                                if v31 {
                                    v68 := GenS4{F0: "\n\t"}
                                    if v2 {
                                        for fi69 := 0; fi69 < 5; fi69++ {
                                            v70 := GenS2{F0: 1000, F1: 128}
                                            _ = fi69
                                            _ = v70
                                        }
                                        if (-1 < (v5 >> ((2) & 63))) {
                                            if len(v66) >= 2 {
                                                v66 = v66[1:len(v66)]
                                            }
                                            if (!v2) {
                                                v22 += 42
                                                fmt.Println(v1)
                                                var v71 uint64 = 4294967295
                                                _ = v71
                                            }
                                        } else {
                                            switch "foo bar" {
                                            case "foo bar":
                                                var v72 uint32 = 65535
                                                _ = v72
                                            }
                                            v22 += v22
                                            _ = len(v66)
                                        }
                                        v31 = v31
                                    }
                                    if v2 {
                                        v73 := GenS8{F0: 934}
                                        _ = v73
                                    } else {
                                        for fi74 := 0; fi74 < 2; fi74++ {
                                            v2 = false
                                            _ = fi74
                                        }
                                        for fi75 := 0; fi75 < 10; fi75++ {
                                            if v2 {
                                                v22 = (^(v5))
                                                v22 = 606
                                                fmt.Println(v20, v22)
                                            } else {
                                                if len(v66) > 0 {
                                                    v66[0] = (fi75 / ((fi75) | 1))
                                                }
                                            }
                                            if (v1 <= 0.0) {
                                                v1 -= (v1 / (v1 / 1.5))
                                            }
                                            _ = fi75
                                        }
                                        var v76 uint64 = 0
                                        _ = v76
                                    }
                                    _ = v68
                                } else {
                                    if (65535 != 479) {
                                        var v77 uint = 653
                                        var v78 uint8 = 127
                                        for fi79 := 0; fi79 < 10; fi79++ {
                                            v2 = true
                                            v3 -= v3
                                            _ = fi79
                                        }
                                        _ = v77
                                        _ = v78
                                    } else {
                                        var v80 uint8 = 17
                                        v4 += 3.14159
                                        var v81 int8 = -128
                                        _ = v80
                                        _ = v81
                                    }
                                    var v82 int64 = 1
                                    v83 := GenS9{F0: 65535, F1: (98 <= 4294967295)}
                                    _ = v82
                                    _ = v83
                                }
                            default:
                                v31 = (0.1 >= (((((v4 / v4) + ((v4 * (v4 + v4)) * v4)) - (v4 - (v4 * v4))) - ((float32(float64(3.14159)) * (float32((float64((float32(0.1) + float32(v20))) + (v1 - v20))) + (3.14159 / (((-2.25 - -2.25) / (v4 + 1e10)) * v4)))) / v4)) * v4))
                            }
                        case "\n\t":
                            v84 := GenS10{F0: 510, F1: 18446744073709551615, F2: 100}
                            _ = v84
                        default:
                            var v85 int8 = 0
                            _ = v85
                        }
                    }
                    v86 := GenS11{F0: 42, F1: fi65}
                    v1 -= (1.5 / v1)
                    _ = v86
                }
                _ = fi65
                _ = v66
            }
            if v31 {
                v87 := GenS1{F0: 255}
                var v88 uint64 = 25
                if (1.5 > v20) {
                    v5 = v5
                    for fi89 := 0; fi89 < 9; fi89++ {
                        for fi90 := 0; fi90 < 14; fi90++ {
                            v22 = -703
                            v2 = (-9223372036854775808 <= v5)
                            _ = fi90
                        }
                        var v91 int64 = (v5 << ((865) & 63))
                        v92 := GenS5{F0: 336}
                        _ = fi89
                        _ = v91
                        _ = v92
                    }
                }
                _ = v87
                _ = v88
            }
        }
        fmt.Println(v1, v20, v31)
    }
    var v93 uint64 = 922
    v4 = (1e10 + (1.5 / float32((v20 * ((v1 / v1) / 1.0)))))
    v22 = (v5 / ((0) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v93
}
