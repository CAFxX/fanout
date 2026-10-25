package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32; F1 int32; F2 int }
type GenS5 struct { F0 int; F1 uint16; F2 bool }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 uint; F1 int16; F2 int }
type GenS8 struct { F0 uint64; F1 int }
func genFn0(p0 bool) int8 {
    v4 := []int{}
    v5 := GenS1{F0: 65535, F1: 2}
    _ = p0
    _ = v4
    _ = v5
    return 1
}

func main() {
    v1 := make(map[int]int)
    var v2 string = "foo bar"
    var v3 string = ((v2 + "") + "")
    switch 127 {
    case 2147483647:
        for fi6 := 0; fi6 < 13; fi6++ {
            v7 := []int{fi6, (fi6 / ((fi6) | 1)), (fi6 % (((fi6 << ((2) & 63))) | 1))}
            v8 := GenS2{F0: 0}
            if len(v7) > 0 {
                v7[0] = (fi6 ^ fi6)
            }
            _ = fi6
            _ = v7
            _ = v8
        }
    case 305:
        var v9 uint32 = 139
        _ = v9
    case 32767:
        _ = len(v1)
    default:
        fmt.Println(v2)
    }
    fmt.Println(v3, v2)
    v10 := GenS3{F0: 0}
    if true {
        v2 = v2
    } else {
        for fi11 := 0; fi11 < 2; fi11++ {
            if false {
                if false {
                    if false {
                        v12 := GenS4{F0: 524, F1: 2147483647, F2: -9223372036854775808}
                        v13 := GenS1{F0: 771, F1: 255}
                        fmt.Println(fi11, fi11)
                        _ = v12
                        _ = v13
                    }
                }
                fmt.Println(fi11)
            } else {
                switch (v2 + (v2 + v2)) {
                case "a":
                    switch v3 {
                    case "\n\t":
                        v14 := v1[(fi11 | (fi11 | 670))]
                        _ = v14
                    }
                case "foo bar":
                    for fi15 := 0; fi15 < 12; fi15++ {
                        v2 = v3
                        switch (fi15 * (fi15 * fi15)) {
                        case -277:
                            var v16 uint64 = 0
                            _ = v16
                        case 127:
                            if (-128 != 0) {
                                if false {
                                    v17 := v1[fi15]
                                    v18 := []int{872, (fi15 / ((127) | 1))}
                                    var v19 int16 = 42
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                }
                            }
                        default:
                            v3 = v2
                        }
                        v20 := GenS4{F0: 219, F1: 961, F2: -359}
                        _ = fi15
                        _ = v20
                    }
                default:
                    for fi21 := 0; fi21 < 5; fi21++ {
                        var v22 uint32 = 1
                        var v23 uint32 = v22
                        v24 := []int{fi21, fi21, fi21, -130}
                        var v25 int64 = -1
                        _ = fi21
                        _ = v22
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                }
            }
            _ = fi11
        }
        var v26 string = v3
        _ = v26
    }
    v1[-1] = 0
    v3 = v3
    for fi27 := 0; fi27 < 19; fi27++ {
        v2 = ("\n\t" + ("a" + (v2 + ("\n\t" + v3))))
        if false {
            if true {
                for fi28 := 0; fi28 < 15; fi28++ {
                    switch "a" {
                    case "a":
                        v3 = v3
                    case "":
                        fmt.Println(fi28, v3, fi27)
                    }
                    var v29 uint8 = 255
                    if true {
                        var v30 float32 = ((float32((1e10 / 1e10)) - -2.25) + -2.25)
                        _ = v30
                    } else {
                        if false {
                            var v31 int8 = 127
                            if true {
                                var v32 bool = (2 <= 2147483648)
                                fmt.Println(v29)
                                var v33 float64 = 0.1
                                _ = v32
                                _ = v33
                            }
                            delete(v1, (-(fi27)))
                            if (false && ((!true) && false)) {
                                var v34 uint32 = 1
                                v35 := []int{(fi27 | 566), (fi27 + fi28)}
                                _ = v34
                                _ = v35
                            } else {
                                if (!false) {
                                    if (false || true) {
                                        v36 := GenS4{F0: 0, F1: 1, F2: 42}
                                        _ = v36
                                    }
                                } else {
                                    v37 := GenS5{F0: (fi27 * fi27), F1: 2, F2: ((!(1e10 == (((float64(float32(0.1)) * ((-2.25 / 1.0) - float64(((0.0 * (1.5 / -2.25)) - -2.25)))) - 1e10) / 3.14159))) && false)}
                                    _ = v37
                                }
                                v38 := GenS6{F0: true, F1: 974}
                                var v39 uint64 = 18446744073709551615
                                _ = v38
                                _ = v39
                            }
                            if false {
                                v40 := []int{9223372036854775807}
                                v29 = v29
                                var v41 int32 = -2147483648
                                _ = v40
                                _ = v41
                            }
                            _ = v31
                        }
                        if (894 > 42) {
                            var v42 uint = 565
                            v43 := GenS4{F0: 2, F1: -2147483648, F2: (^(fi27))}
                            switch (v2 + v3) {
                            case "\n\t":
                                v3 = v3
                            case "":
                                var v44 uint32 = 2
                                _ = v44
                            default:
                                v45 := GenS2{F0: v42}
                                _ = v45
                            }
                            var v46 uint = (v42 << ((2) & 63))
                            _ = v42
                            _ = v43
                            _ = v46
                        } else {
                            var v47 int32 = 42
                            v48 := []int{fi28, fi27, 32768}
                            if (genFn0(true) == 0) {
                                var v49 int8 = 0
                                _ = len(v48)
                                _ = v49
                            }
                            _ = v47
                            _ = v48
                        }
                        var v50 uint32 = 2147483647
                        _ = v50
                    }
                    _ = fi28
                    _ = v29
                }
            } else {
                v3 = "foo bar"
                delete(v1, (fi27 - fi27))
                v51 := GenS7{F0: 255, F1: 821, F2: fi27}
                _ = v51
            }
            switch fi27 {
            case -2147483648:
                var v52 string = (v2 + v2)
                _ = v52
            case 1:
                v53 := GenS4{F0: 984, F1: -131, F2: (fi27 | 9223372036854775807)}
                _ = v53
            case -292:
                v54 := GenS4{F0: 26, F1: 1, F2: fi27}
                _ = v54
            }
        } else {
            for fi55 := 0; fi55 < 10; fi55++ {
                fmt.Println(v3, v2)
                fmt.Println(fi27, fi27, v2)
                _ = fi55
            }
        }
        var v56 uint = 592
        _ = fi27
        _ = v56
    }
    if false {
        if (!(!false)) {
            v57 := GenS3{F0: 4294967295}
            _ = v57
        } else {
            fmt.Println(v3, v2)
        }
        v58 := GenS4{F0: 2, F1: 1, F2: -363}
        _ = v58
    } else {
        v59 := GenS6{F0: false, F1: 1}
        fmt.Println(v2)
        var v60 int = -333
        _ = v59
        _ = v60
    }
    v61 := GenS5{F0: -202, F1: 0, F2: (3.14159 <= (((1.5 + 0.1) - 3.14159) + (1.5 * (-2.25 + 1.5))))}
    var v62 int8 = -109
    var v63 float64 = (-2.25 * -2.25)
    v62 += genFn0(true)
    var v64 uint16 = 2
    var v65 int = 4294967295
    switch v65 {
    case 32767:
        for fi66 := 0; fi66 < 20; fi66++ {
            for fi67 := 0; fi67 < 16; fi67++ {
                v3 = v2
                var v68 int8 = v62
                if false {
                    v62 = v62
                    switch (v2 + (v3 + (((v2 + "hello") + "\n\t") + "a"))) {
                    case "\n\t":
                        fmt.Println(v2)
                    }
                    delete(v1, fi66)
                } else {
                    var v69 uint64 = 2
                    _ = v69
                }
                _ = fi67
                _ = v68
            }
            _ = fi66
        }
    default:
        for fi70 := 0; fi70 < 3; fi70++ {
            v2 = (("a" + v3) + (v2 + "hello"))
            switch -2147483648 {
            case 9223372036854775807:
                v71 := GenS1{F0: (v64 / ((65535) | 1)), F1: (^(v64))}
                _ = v71
            case 1:
                if false {
                    v65 -= -545
                    var v72 float64 = (v63 / float64(-2.25))
                    v73 := GenS8{F0: 1, F1: 127}
                    _ = v72
                    _ = v73
                } else {
                    fmt.Println(fi70)
                    v74 := []int{fi70, fi70}
                    v63 += v63
                    _ = v74
                }
            case 353:
                v62 = (v62 >> ((0) & 63))
            default:
                var v75 uint64 = 113
                _ = v75
            }
            _ = fi70
        }
    }
    var v76 float32 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v61
    _ = v62
    _ = v63
    _ = v64
    _ = v65
    _ = v76
}
