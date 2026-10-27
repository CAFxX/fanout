package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 bool; F1 int16 }
type GenS7 struct { F0 uint64; F1 uint }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]string)
    var v3 bool = ((32767 == 32767) && false)
    v4 := make(map[int]string)
    var v5 float32 = (float32(-2.25) + float32(3.14159))
    for fi6 := 0; fi6 < 6; fi6++ {
        switch (fi6 / ((fi6) | 1)) {
        case 32767:
            v7 := GenS1{F0: 65535, F1: "\n\t"}
            _ = v7
        case 833:
            var v8 int16 = 32767
            _ = v8
        case 32768:
            var v9 uint32 = 974
            _ = v9
        default:
            v10 := GenS1{F0: 32767, F1: ""}
            _ = v10
        }
        for fi11 := 0; fi11 < 16; fi11++ {
            if v3 {
                var v12 int16 = 32767
                if true {
                    v12 = 1000
                    if (0 > 402) {
                        if v3 {
                            var v13 float32 = (-2.25 / 1e10)
                            if (v3 && (442 <= -1)) {
                                if v3 {
                                    if (-1 <= -128) {
                                        var v14 uint8 = 127
                                        v15 := GenS2{F0: 9223372036854775807, F1: 1}
                                        _ = v14
                                        _ = v15
                                    }
                                    v12 -= 620
                                } else {
                                    v13 += float32(3.14159)
                                    switch fi6 {
                                    case 32768:
                                        var v16 int64 = -1
                                        _ = v16
                                    case 915:
                                        var v17 uint16 = 563
                                        _ = v17
                                    case 830:
                                        var v18 int16 = (v12 / (((^(v12))) | 1))
                                        _ = v18
                                    default:
                                        v19 := GenS3{F0: -1, F1: 25}
                                        _ = v19
                                    }
                                }
                                v12 -= v12
                                var v20 string = ""
                                var v21 bool = v3
                                _ = v20
                                _ = v21
                            }
                            v5 -= (v5 / ((v5 * 1.5) - (v13 + float32(-2.25))))
                            v1[1] = v3
                            _ = v13
                        } else {
                            if ((296 > 255) || false) {
                                v5 = v5
                                if ((((!false) || v3) || v3) && (true && (81 <= 105))) {
                                    v3 = false
                                    if false {
                                        switch (fi6 ^ fi11) {
                                        case -401:
                                            fmt.Println(v12, v5)
                                        }
                                        var v22 string = ("hello" + "hello")
                                        _ = v22
                                    }
                                } else {
                                    var v23 int16 = (v12 ^ v12)
                                    fmt.Println(fi11, v5)
                                    v3 = v3
                                    _ = v23
                                }
                                switch fi11 {
                                case 127:
                                    v24 := GenS1{F0: 478, F1: "a"}
                                    _ = v24
                                case 167:
                                    if v3 {
                                        var v25 uint32 = 42
                                        if (726 <= v25) {
                                            var v26 int16 = (v12 << ((2) & 63))
                                            switch fi11 {
                                            case 116:
                                                v27 := GenS1{F0: 849, F1: ""}
                                                _ = v27
                                            }
                                            _ = v26
                                        } else {
                                            v28 := GenS3{F0: -1, F1: 2}
                                            v29 := v4[fi11]
                                            v5 = 0.1
                                            _ = v28
                                            _ = v29
                                        }
                                        _ = v25
                                    }
                                default:
                                    v30 := GenS2{F0: -478, F1: -71}
                                    _ = v30
                                }
                            } else {
                                v12 = -32768
                            }
                            fmt.Println(v5, v3, v5)
                        }
                        var v31 uint = 572
                        var v32 uint16 = 0
                        _ = v31
                        _ = v32
                    }
                }
                _ = len(v1)
                _ = v12
            }
            v33 := GenS2{F0: -659, F1: -128}
            v3 = (109 != 14)
            _ = fi11
            _ = v33
        }
        var v34 uint = 42
        _ = fi6
        _ = v34
    }
    fmt.Println(v3, v3)
    v3 = (9223372036854775807 > 0)
    v3 = ((float32(0.0) * -2.25) <= 3.14159)
    for fi35 := 0; fi35 < 2; fi35++ {
        for fi36 := 0; fi36 < 20; fi36++ {
            v3 = true
            v37 := v1[4294967295]
            _ = len(v4)
            _ = fi36
            _ = v37
        }
        _ = fi35
    }
    v38 := GenS3{F0: -974, F1: 42}
    v3 = false
    for fi39 := 0; fi39 < 8; fi39++ {
        if (18446744073709551615 >= 541) {
            v3 = v3
        } else {
            var v40 float64 = -2.25
            var v41 int8 = 1
            for fi42 := 0; fi42 < 16; fi42++ {
                var v43 uint64 = 1
                _ = fi42
                _ = v43
            }
            _ = v40
            _ = v41
        }
        var v44 bool = v3
        _ = fi39
        _ = v44
    }
    var v45 string = ""
    if v3 {
        for fi46 := 0; fi46 < 5; fi46++ {
            for fi47 := 0; fi47 < 5; fi47++ {
                if v3 {
                    if v3 {
                        if (65535 <= 723) {
                            v48 := GenS3{F0: -1, F1: 127}
                            if (0.1 < v5) {
                                fmt.Println(v3, fi47)
                            } else {
                                v3 = false
                                var v49 float32 = float32((1.5 / 1.5))
                                v50 := GenS4{F0: 255}
                                var v51 bool = v3
                                _ = v49
                                _ = v50
                                _ = v51
                            }
                            if (338 < 0) {
                                if (1 == -32768) {
                                    fmt.Println(v3)
                                    switch "foo bar" {
                                    case "hello":
                                        v52 := GenS1{F0: 255, F1: (v45 + "\n\t")}
                                        _ = v52
                                    }
                                }
                            }
                            _ = v48
                        }
                        v53 := []int{}
                        var v54 int = fi46
                        _ = v53
                        _ = v54
                    } else {
                        var v55 uint32 = 4294967295
                        var v56 int16 = 0
                        _ = v55
                        _ = v56
                    }
                }
                v57 := GenS5{F0: ((v45 + "foo bar") + "\n\t"), F1: (2 > 65535)}
                var v58 int64 = -839
                var v59 uint8 = 127
                delete(v4, (^(fi47)))
                var v60 int32 = -1
                _ = fi47
                _ = v57
                _ = v58
                _ = v59
                _ = v60
            }
            for fi61 := 0; fi61 < 20; fi61++ {
                v62 := []int{370}
                var v63 uint8 = 26
                v3 = false
                _ = fi61
                _ = v62
                _ = v63
            }
            v64 := []int{(fi46 / ((fi46) | 1)), (-(fi46)), 4294967295, fi46}
            _ = fi46
            _ = v64
        }
    } else {
        fmt.Println(v45, v5, v3)
        switch -90 {
        case -609:
            fmt.Println(v3, v5, v45)
        case 526:
            _ = len(v4)
        default:
            if (9223372036854775807 < -1) {
                if (v5 > float32(float64(((v5 / (float32(0.1) + v5)) * (v5 + v5))))) {
                    var v65 uint8 = 81
                    _ = v65
                }
            } else {
                var v66 int8 = -1
                switch -997 {
                case 0:
                    if (-1 != 0) {
                        for fi67 := 0; fi67 < 10; fi67++ {
                            for fi68 := 0; fi68 < 17; fi68++ {
                                var v69 uint64 = 42
                                v70 := GenS5{F0: "hello", F1: false}
                                if (v3 || false) {
                                    v3 = false
                                    var v71 int16 = 1
                                    _ = v71
                                } else {
                                    v72 := GenS6{F0: (4294967295 > 42), F1: int16(v69)}
                                    var v73 int16 = int16(v66)
                                    v74 := []int{-2147483648, -9223372036854775808, (-(fi67))}
                                    _ = v72
                                    _ = v73
                                    _ = v74
                                }
                                _ = fi68
                                _ = v69
                                _ = v70
                            }
                            fmt.Println(v3, v3, v45)
                            _ = fi67
                        }
                    } else {
                        v75 := GenS7{F0: 4294967295, F1: 65535}
                        _ = v75
                    }
                case 385:
                    if v3 {
                        for fi76 := 0; fi76 < 4; fi76++ {
                            var v77 uint = 2
                            _ = fi76
                            _ = v77
                        }
                    }
                default:
                    if (v3 || (false || ((360 > 0) || true))) {
                        fmt.Println(v5, v66, v66)
                    }
                }
                var v78 string = "a"
                _ = v66
                _ = v78
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v45
    _ = v5
}
