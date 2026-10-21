package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint16; F2 uint32 }
type GenS4 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint; F1 int16 }
type GenS9 struct { F0 uint8; F1 int64; F2 uint8 }

func main() {
    var v1 uint8 = 2
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]uint32)
    var v4 string = "\n\t"
    var v5 float64 = -2.25
    fmt.Println(v4, v5)
    v4 = v4
    switch 9223372036854775807 {
    case 42:
        v6 := GenS1{F0: v1, F1: 2, F2: v1}
        _ = v6
    }
    var v7 string = v4
    v1 += 255
    var v8 int32 = 1
    v9 := GenS2{F0: 42}
    switch 32768 {
    case 9223372036854775807:
        var v10 int64 = 387
        _ = v10
    case -2147483648:
        v11 := GenS1{F0: (v1 * (v1 & v1)), F1: 2, F2: (v1 | (v1 | v1))}
        _ = v11
    default:
        if true {
            v4 = (v4 + "foo bar")
            for fi12 := 0; fi12 < 12; fi12++ {
                fmt.Println(v8, v7)
                if (v5 != (v5 + v5)) {
                    v8 = (v8 ^ (v8 / (((v8 >> ((4294967295) & 63))) | 1)))
                }
                if ((fi12 + -370) >= -9223372036854775808) {
                    switch v4 {
                    case "\n\t":
                        v8 = (^(v8))
                    case "hello":
                        var v13 bool = false
                        _ = v13
                    case "a":
                        for fi14 := 0; fi14 < 10; fi14++ {
                            v15 := GenS2{F0: 4294967295}
                            fmt.Println(v1, v1, v1)
                            v8 -= v8
                            _ = fi14
                            _ = v15
                        }
                    }
                    v16 := []int{-958}
                    switch fi12 {
                    case -937:
                        if (!false) {
                            if true {
                                var v17 bool = (true || (69 <= -1))
                                v4 = v7
                                _ = v17
                            }
                            var v18 float32 = (3.14159 * float32(v5))
                            _ = v18
                        }
                    default:
                        var v19 uint = uint(v1)
                        _ = v19
                    }
                    _ = v16
                } else {
                    for fi20 := 0; fi20 < 9; fi20++ {
                        v21 := GenS3{F0: fi20, F1: 42, F2: 1}
                        if true {
                            switch fi20 {
                            case 127:
                                v22 := []int{fi12, 193, fi12}
                                _ = v22
                            case 576:
                                switch ((v7 + v4) + v7) {
                                case "foo bar":
                                    v1 = v1
                                case "hello":
                                    v23 := []int{469, (fi20 | fi12)}
                                    _ = v23
                                }
                            case -2147483648:
                                if false {
                                    switch v4 {
                                    case "":
                                        v24 := GenS2{F0: fi12}
                                        _ = v24
                                    default:
                                        var v25 bool = (!false)
                                        _ = v25
                                    }
                                    fmt.Println(v1, fi20)
                                    v7 = v7
                                }
                                v26 := v3[122]
                                if (v1 < v1) {
                                    var v27 string = "foo bar"
                                    fmt.Println(v7, fi20)
                                    switch v4 {
                                    case "hello":
                                        v7 = v4
                                        switch (v7 + v7) {
                                        case "":
                                            var v28 float32 = float32(float64(float32(((v5 / 1e10) * v5))))
                                            _ = v28
                                        case "hello":
                                            switch v7 {
                                            case "\n\t":
                                                if false {
                                                    if true {
                                                        if false {
                                                            v27 = ""
                                                            if true {
                                                                if true {
                                                                    switch fi12 {
                                                                    case 127:
                                                                        fmt.Println(v27, v27, v8)
                                                                    case 881:
                                                                        switch (v7 + (v4 + (v7 + v4))) {
                                                                        case "a":
                                                                            v29 := GenS3{F0: fi12, F1: uint16(v26), F2: 0}
                                                                            _ = v29
                                                                        case "\n\t":
                                                                            _ = len(v3)
                                                                            v30 := []int{(fi12 / (((fi12 % ((fi20) | 1))) | 1)), (fi20 & fi20), (fi20 & (fi12 << ((643) & 63)))}
                                                                            var v31 uint64 = 0
                                                                            _ = v30
                                                                            _ = v31
                                                                        }
                                                                    default:
                                                                        v7 = ((v7 + "hello") + v4)
                                                                    }
                                                                    v1 -= v1
                                                                } else {
                                                                    v32 := GenS4{F0: v26, F1: v26, F2: 42}
                                                                    var v33 bool = ((!false) || false)
                                                                    var v34 int32 = 0
                                                                    _ = v32
                                                                    _ = v33
                                                                    _ = v34
                                                                }
                                                                if (false && false) {
                                                                    v5 = (float64((3.14159 + 1.5)) - 0.0)
                                                                    switch v7 {
                                                                    case "hello":
                                                                        if false {
                                                                            v7 = (v27 + "foo bar")
                                                                            if true {
                                                                                v5 = ((float64(float32((-2.25 + (v5 - v5)))) / v5) / (v5 * v5))
                                                                                var v35 bool = false
                                                                                v8 = v8
                                                                                _ = v35
                                                                            }
                                                                        }
                                                                    case "\n\t":
                                                                        v1 = (v1 | (v1 | 243))
                                                                    case "a":
                                                                        fmt.Println(v1)
                                                                    default:
                                                                        v36 := GenS1{F0: v1, F1: 625, F2: (v1 | 127)}
                                                                        _ = v36
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            fmt.Println(v7, fi20, v26)
                                                            var v37 uint32 = 0
                                                            _ = v37
                                                        }
                                                        v27 = "\n\t"
                                                        fmt.Println(v8, v27)
                                                    } else {
                                                        var v38 int32 = v8
                                                        if false {
                                                            v39 := GenS3{F0: fi12, F1: 1, F2: 1}
                                                            _ = v39
                                                        } else {
                                                            v7 = "a"
                                                            switch fi20 {
                                                            case 42:
                                                                var v40 bool = true
                                                                _ = v40
                                                            case 2147483647:
                                                                v41 := GenS5{F0: 0, F1: 42}
                                                                _ = v41
                                                            case -2147483648:
                                                                v42 := []int{}
                                                                _ = v42
                                                            }
                                                            switch (("foo bar" + v7) + "foo bar") {
                                                            case "foo bar":
                                                                v5 = ((v5 * v5) * v5)
                                                            case "\n\t":
                                                                v43 := []int{32767, (fi12 / ((2147483647) | 1))}
                                                                _ = v43
                                                            }
                                                        }
                                                        _ = v38
                                                    }
                                                    v7 = v27
                                                    v44 := GenS6{F0: v27, F1: 0}
                                                    v45 := []int{}
                                                    v46 := GenS2{F0: (fi20 * (fi12 ^ fi12))}
                                                    if len(v45) >= 2 {
                                                        v45 = v45[1:len(v45)]
                                                    }
                                                    _ = v44
                                                    _ = v45
                                                    _ = v46
                                                } else {
                                                    var v47 int16 = 1000
                                                    v7 = "foo bar"
                                                    v5 += v5
                                                    _ = v47
                                                }
                                            case "foo bar":
                                                v48 := GenS7{F0: v26}
                                                _ = v48
                                            default:
                                                if false {
                                                    if true {
                                                        v49 := GenS3{F0: fi20, F1: 255, F2: (v26 | 4294967295)}
                                                        var v50 uint64 = 814
                                                        if (!false) {
                                                            fmt.Println(fi20, v27, fi12)
                                                        }
                                                        _ = v49
                                                        _ = v50
                                                    } else {
                                                        if true {
                                                            v51 := v2[(^(v26))]
                                                            var v52 string = ("\n\t" + ((v7 + (v7 + (v7 + v7))) + v27))
                                                            var v53 float32 = 0.0
                                                            _ = v51
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                        v54 := GenS2{F0: (fi12 % (((fi20 - (fi12 - (fi20 + 4294967295)))) | 1))}
                                                        _ = v54
                                                    }
                                                    v2[(v26 | (^(v26)))] = (fi20 | (fi12 - -897))
                                                    var v55 float32 = 0.1
                                                    var v56 int8 = -74
                                                    fmt.Println(v8, v8, v27)
                                                    _ = v55
                                                    _ = v56
                                                }
                                                var v57 bool = false
                                                var v58 bool = v57
                                                _ = v57
                                                _ = v58
                                            }
                                            v27 = v4
                                            var v59 uint32 = v26
                                            v60 := GenS6{F0: v27, F1: (v1 + (^(v1)))}
                                            if true {
                                                if (-1 <= 581) {
                                                    switch v7 {
                                                    case "":
                                                        v61 := GenS8{F0: 42, F1: -32768}
                                                        _ = v61
                                                    case "hello":
                                                        if true {
                                                            var v62 int8 = -1
                                                            v5 += v5
                                                            _ = len(v2)
                                                            _ = v62
                                                        }
                                                    }
                                                    v27 = v27
                                                    switch "foo bar" {
                                                    case "":
                                                        v63 := GenS2{F0: fi20}
                                                        _ = v63
                                                    default:
                                                        v64 := v3[(v26 >> ((811) & 63))]
                                                        v65 := []int{-2147483648, 1, fi12}
                                                        _ = v64
                                                        _ = v65
                                                    }
                                                }
                                            }
                                            _ = v59
                                            _ = v60
                                        case "a":
                                            fmt.Println(v1, v8)
                                        default:
                                            if true {
                                                if false {
                                                    var v66 int64 = 42
                                                    if false {
                                                        var v67 uint8 = v1
                                                        _ = v67
                                                    }
                                                    v5 += v5
                                                    _ = v66
                                                }
                                            }
                                        }
                                    case "\n\t":
                                        delete(v2, v26)
                                    case "foo bar":
                                        v1 -= (^(v1))
                                    }
                                    v7 = v4
                                    v7 = "\n\t"
                                    _ = v27
                                }
                                v4 = v7
                                switch fi12 {
                                case -9223372036854775808:
                                    fmt.Println(v8, fi20)
                                }
                                v68 := GenS9{F0: 40, F1: 42, F2: (v1 * v1)}
                                v69 := v2[65535]
                                _ = v26
                                _ = v68
                                _ = v69
                            }
                            var v70 string = (v7 + "")
                            _ = v70
                        } else {
                            v4 = "hello"
                        }
                        _ = fi20
                        _ = v21
                    }
                    if true {
                        switch ("\n\t" + ((v4 + ("foo bar" + (v4 + "foo bar"))) + "")) {
                        case "hello":
                            if (false || false) {
                                if (2 == 18446744073709551615) {
                                    switch v7 {
                                    case "foo bar":
                                        var v71 bool = (-1 >= 977)
                                        _ = v71
                                    }
                                }
                            }
                        case "a":
                            v5 = v5
                        default:
                            v4 = v4
                        }
                        var v72 int16 = 0
                        v4 = v4
                        _ = v72
                    } else {
                        var v73 uint64 = 42
                        _ = v73
                    }
                    v74 := GenS1{F0: 255, F1: 4294967295, F2: (v1 | (v1 >> ((0) & 63)))}
                    _ = v74
                }
                _ = fi12
            }
            for fi75 := 0; fi75 < 3; fi75++ {
                v5 = (v5 / ((3.14159 + 0.0) / (float64(-2.25) - 0.0)))
                v1 += 213
                var v76 uint16 = 42
                _ = fi75
                _ = v76
            }
        }
    }
    v7 = (v7 + v4)
    if (4294967295 != -22) {
        var v77 int8 = 5
        if (2147483648 == 1) {
            var v78 int64 = -10
            fmt.Println(v78)
            _ = v78
        }
        _ = v77
    }
    fmt.Println(v4, v7, v5)
    var v79 uint8 = (v1 << ((185) & 63))
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v79:", v79)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v79
    _ = v8
    _ = v9
}
