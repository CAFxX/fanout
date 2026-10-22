package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 uint64; F1 int16; F2 int }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 string; F1 uint64; F2 int }
type GenS6 struct { F0 int16; F1 uint8; F2 uint }
func genFn0(p0 bool, p1 int) bool {
    var v7 int8 = -101
    var v8 int = (p1 >> ((941) & 63))
    v9 := []int{p1, 1, int(v7), v8}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return p0
}
func genFn1() bool {
    var v10 float64 = 1.5
    fmt.Println(v10)
    _ = v10
    return (!true)
}

func main() {
    var v1 float64 = 0.1
    var v2 int32 = -961
    var v3 float32 = 3.14159
    var v4 int64 = 0
    var v5 string = ((("hello" + "") + "\n\t") + (("foo bar" + ("hello" + "")) + ("foo bar" + "\n\t")))
    var v6 float64 = v1
    v2 = v2
    if false {
        v2 += 42
        var v11 int = 127
        _ = v11
    } else {
        v4 -= (v4 % ((v4) | 1))
        v12 := GenS1{F0: 65535, F1: (v4 & 1)}
        fmt.Println(v4, v2, v2)
        _ = v12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS1{F0: 65535, F1: (-(v4))}
        v15 := []int{fi13, fi13}
        fmt.Println(v5, v1, v5)
        _ = fi13
        _ = v14
        _ = v15
    }
    switch 32768 {
    case 32768:
        if (1 != v4) {
            v3 += -2.25
            var v16 uint32 = 4294967295
            fmt.Println(v16, v2)
            _ = v16
        }
    case -2147483648:
        v6 -= 0.1
    case -9223372036854775808:
        for fi17 := 0; fi17 < 11; fi17++ {
            var v18 int8 = 100
            _ = fi17
            _ = v18
        }
    }
    fmt.Println(v4, v5, v3)
    if false {
        switch 127 {
        case 128:
            v19 := GenS2{F0: v4, F1: 4294967295}
            _ = v19
        case 2147483647:
            v20 := []int{2147483647, -1, -2147483648}
            _ = v20
        case 32767:
            fmt.Println(v5, v3)
        }
        if (!(v2 != -654)) {
            if true {
                v5 = v5
            }
            for fi21 := 0; fi21 < 12; fi21++ {
                var v22 float32 = 1e10
                if (0 >= 127) {
                    v4 -= 0
                } else {
                    var v23 bool = (1 == 1)
                    v24 := []int{fi21, -718}
                    _ = v23
                    _ = v24
                }
                var v25 float64 = (0.0 - ((float64(v22) / v1) - v6))
                _ = fi21
                _ = v22
                _ = v25
            }
            v2 = v2
        } else {
            var v26 int16 = 421
            _ = v26
        }
        switch 4294967295 {
        case 2147483647:
            v27 := GenS1{F0: 388, F1: -9223372036854775808}
            _ = v27
        case 0:
            switch -2147483648 {
            case -9223372036854775808:
                v28 := []int{94}
                _ = v28
            default:
                for fi29 := 0; fi29 < 17; fi29++ {
                    fmt.Println(v2)
                    fmt.Println(v3, fi29, v1)
                    if (!true) {
                        if false {
                            if (0 >= 338) {
                                v30 := GenS2{F0: -9223372036854775808, F1: 18446744073709551615}
                                var v31 uint = 0
                                if (fi29 == -2147483648) {
                                    var v32 uint = v31
                                    var v33 uint = (v32 >> (((v31 | 42)) & 63))
                                    _ = v32
                                    _ = v33
                                } else {
                                    for fi34 := 0; fi34 < 13; fi34++ {
                                        var v35 int8 = -128
                                        v36 := []int{fi34, fi34, -855, fi34}
                                        if true {
                                            fmt.Println(v3, v31, v4)
                                            var v37 bool = (4294967295 == 2147483647)
                                            if v37 {
                                                var v38 uint = v31
                                                switch fi34 {
                                                case -9223372036854775808:
                                                    fmt.Println(v35, fi29)
                                                case 953:
                                                    var v39 int64 = v4
                                                    _ = v39
                                                case 36:
                                                    v40 := GenS3{F0: 1, F1: -971, F2: fi34}
                                                    _ = v40
                                                default:
                                                    if false {
                                                        v2 = (v2 % ((v2) | 1))
                                                        var v41 uint32 = 2147483647
                                                        v36 = append(v36, 4294967295)
                                                        _ = v41
                                                    }
                                                }
                                                _ = v38
                                            } else {
                                                v4 = v4
                                                if len(v36) > 0 {
                                                    v36[0] = -1
                                                }
                                                fmt.Println(v6, v1, v6)
                                                v37 = v37
                                                switch v5 {
                                                case "\n\t":
                                                    v42 := GenS2{F0: -988, F1: 1}
                                                    _ = v42
                                                case "":
                                                    v5 = v5
                                                default:
                                                    fmt.Println(v2, v37)
                                                }
                                            }
                                            _ = v37
                                        }
                                        _ = fi34
                                        _ = v35
                                        _ = v36
                                    }
                                }
                                _ = v30
                                _ = v31
                            }
                            v43 := GenS4{F0: -32, F1: false}
                            _ = v43
                        }
                    }
                    _ = fi29
                }
            }
        case -791:
            for fi44 := 0; fi44 < 3; fi44++ {
                v1 = (float64(float32(((v1 - 1e10) / v6))) - v1)
                var v45 int32 = v2
                _ = fi44
                _ = v45
            }
        }
    }
    for fi46 := 0; fi46 < 7; fi46++ {
        v5 = ((v5 + v5) + v5)
        v47 := []int{(-(fi46)), (-(fi46)), fi46, (fi46 % ((42) | 1))}
        if ((0 > (-(v4))) && (!(!(1e10 > (0.0 * (v3 / float32(((0.0 + ((((1e10 + (3.14159 * v1)) + v6) + (v6 - v1)) / v1)) * 0.0)))))))) {
            v5 = v5
            switch 32767 {
            case 32767:
                var v48 float64 = 3.14159
                _ = v48
            case 792:
                for fi49 := 0; fi49 < 16; fi49++ {
                    var v50 uint32 = 45
                    v50 = 65535
                    fmt.Println(v2, v2)
                    _ = fi49
                    _ = v50
                }
            }
        } else {
            v5 = (v5 + "hello")
            if ((v4 - v4) <= (^(v4))) {
                if len(v47) > 0 {
                    v47[0] = (fi46 - fi46)
                }
            } else {
                v47 = append(v47, fi46)
                v51 := GenS2{F0: v4, F1: 1}
                _ = v51
            }
            var v52 uint64 = 0
            _ = v52
        }
        _ = fi46
        _ = v47
    }
    if false {
        if (!true) {
            var v53 uint64 = 324
            for fi54 := 0; fi54 < 4; fi54++ {
                v55 := []int{32768}
                if true {
                    v3 -= v3
                    v56 := GenS3{F0: v53, F1: 1000, F2: 127}
                    var v57 int64 = (v4 + (-(v4)))
                    _ = v56
                    _ = v57
                } else {
                    var v58 uint8 = 100
                    _ = v58
                }
                _ = fi54
                _ = v55
            }
            _ = v53
        }
    } else {
        var v59 int8 = -78
        _ = v59
    }
    v60 := GenS1{F0: 4294967295, F1: (v4 << ((721) & 63))}
    v61 := GenS5{F0: v5, F1: 42, F2: -1}
    if (376 > 708) {
        v3 += float32((1.5 / 1.5))
        if (255 >= 254) {
            switch (v5 + "hello") {
            case "":
                v3 -= v3
            case "foo bar":
                for fi62 := 0; fi62 < 7; fi62++ {
                    for fi63 := 0; fi63 < 1; fi63++ {
                        v64 := []int{fi62}
                        _ = fi63
                        _ = v64
                    }
                    for fi65 := 0; fi65 < 20; fi65++ {
                        v66 := []int{(fi65 + 32768), int(v2), 0, (fi62 ^ -1)}
                        v67 := GenS1{F0: 42, F1: 42}
                        _ = fi65
                        _ = v66
                        _ = v67
                    }
                    _ = fi62
                }
            }
            var v68 float32 = (v3 - v3)
            _ = v68
        } else {
            for fi69 := 0; fi69 < 6; fi69++ {
                v6 = v1
                switch v5 {
                case "":
                    for fi70 := 0; fi70 < 2; fi70++ {
                        if (v2 < 1) {
                            var v71 uint = 313
                            _ = v71
                        }
                        if true {
                            var v72 uint16 = 737
                            if (!false) {
                                v73 := GenS3{F0: 2, F1: 157, F2: (fi70 - fi69)}
                                v74 := GenS3{F0: 4294967295, F1: 42, F2: (fi70 >> ((955) & 63))}
                                switch fi70 {
                                case -9223372036854775808:
                                    var v75 uint8 = 128
                                    _ = v75
                                case 499:
                                    fmt.Println(v1, fi70, v2)
                                case 128:
                                    v76 := GenS5{F0: v5, F1: 18446744073709551615, F2: (fi70 & fi69)}
                                    _ = v76
                                }
                                _ = v73
                                _ = v74
                            }
                            var v77 int = (^(fi69))
                            fmt.Println(v3)
                            _ = v72
                            _ = v77
                        }
                        fmt.Println(fi69)
                        _ = fi70
                    }
                case "a":
                    switch v5 {
                    case "":
                        v2 += (v2 % ((v2) | 1))
                    default:
                        v78 := []int{fi69, (^(fi69)), 432, fi69}
                        _ = v78
                    }
                case "foo bar":
                    v1 = 1.5
                }
                _ = fi69
            }
            v79 := []int{-1, 9223372036854775807}
            if len(v79) >= 2 {
                v79 = v79[1:len(v79)]
            }
            _ = v79
        }
    }
    v80 := GenS6{F0: 32767, F1: 55, F2: 349}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v61
    _ = v80
}
