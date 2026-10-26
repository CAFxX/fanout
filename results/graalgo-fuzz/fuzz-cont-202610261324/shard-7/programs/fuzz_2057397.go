package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string; F1 int16; F2 bool }
func genFn0() int16 {
    switch -2147483648 {
    case 42:
        v6 := []int{}
        _ = v6
    }
    var v7 int = 128
    for fi8 := 0; fi8 < 5; fi8++ {
        if false {
            for fi9 := 0; fi9 < 8; fi9++ {
                if false {
                    v7 -= fi8
                    v10 := GenS1{F0: v7, F1: true}
                    _ = v10
                }
                v11 := []int{-2147483648, (fi9 * fi9), (v7 | 2147483647)}
                v12 := GenS2{F0: 0}
                _ = fi9
                _ = v11
                _ = v12
            }
        } else {
            for fi13 := 0; fi13 < 10; fi13++ {
                fmt.Println(v7, fi13)
                var v14 int = (v7 | fi13)
                var v15 uint64 = 614
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        var v16 uint8 = 133
        var v17 int32 = 0
        _ = fi8
        _ = v16
        _ = v17
    }
    _ = v7
    return -420
}
func genFn1(p0 int8) uint16 {
    v18 := GenS2{F0: 42}
    switch 0 {
    case 3:
        p0 += p0
    }
    _ = p0
    _ = v18
    return 255
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = (((float64(-2.25) / -2.25) / 1.0) - 0.0)
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]string)
    v5 := make(map[int]int)
    v19 := GenS3{F0: 1, F1: 914}
    var v20 int8 = 42
    fmt.Println(v2, v2)
    if (("\n\t" < ("foo bar" + "a")) || ((!true) || (((2 != 672) || true) || true))) {
        fmt.Println(v2, v20)
        var v21 float32 = -2.25
        _ = v21
    }
    v22 := GenS1{F0: 387, F1: false}
    if true {
        for fi23 := 0; fi23 < 9; fi23++ {
            fmt.Println(v20, v20, fi23)
            var v24 bool = true
            _ = fi23
            _ = v24
        }
        for fi25 := 0; fi25 < 7; fi25++ {
            fmt.Println(v20, fi25, fi25)
            v26 := []int{}
            var v27 string = ((("foo bar" + ((("a" + "foo bar") + ("a" + "\n\t")) + "")) + "hello") + "foo bar")
            _ = fi25
            _ = v26
            _ = v27
        }
    } else {
        for fi28 := 0; fi28 < 18; fi28++ {
            var v29 uint64 = 895
            v30 := GenS2{F0: v29}
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    if true {
        switch -545 {
        case 0:
            v31 := GenS3{F0: -1, F1: 18446744073709551615}
            _ = v31
        case 522:
            if (4294967295 == 2147483647) {
                var v32 uint8 = 41
                _ = v32
            }
        case -951:
            v33 := GenS1{F0: -1, F1: (true && (float64(1e10) <= v2))}
            _ = v33
        default:
            v34 := GenS2{F0: 4294967295}
            _ = v34
        }
        switch 878 {
        case 314:
            v35 := []int{}
            _ = v35
        case 128:
            var v36 int32 = 576
            _ = v36
        case 32767:
            delete(v4, ("hello" + ""))
        default:
            for fi37 := 0; fi37 < 6; fi37++ {
                for fi38 := 0; fi38 < 6; fi38++ {
                    switch fi37 {
                    case -944:
                        var v39 int32 = -2147483648
                        _ = v39
                    default:
                        v20 = 100
                    }
                    v40 := GenS1{F0: fi37, F1: true}
                    if (!true) {
                        if (0 == 511) {
                            var v41 bool = true
                            if true {
                                v41 = false
                                v42 := GenS1{F0: fi38, F1: v41}
                                fmt.Println(v41, v41, fi37)
                                _ = v42
                            }
                            _ = v41
                        } else {
                            v20 = v20
                            if false {
                                fmt.Println(fi38)
                                switch (fi38 * -1) {
                                case 128:
                                    switch fi37 {
                                    case 485:
                                        v2 += 0.0
                                        _ = len(v4)
                                        var v43 uint32 = 1
                                        _ = v43
                                    default:
                                        var v44 uint = 235
                                        _ = v44
                                    }
                                    switch "\n\t" {
                                    case "\n\t":
                                        v45 := GenS1{F0: fi37, F1: false}
                                        _ = v45
                                    default:
                                        if false {
                                            if (0 != -9223372036854775808) {
                                                var v46 uint8 = 42
                                                fmt.Println(v2)
                                                _ = v46
                                            }
                                        } else {
                                            v20 -= -52
                                            v47 := GenS4{F0: 1}
                                            var v48 uint8 = 128
                                            if (uint16(v48) >= genFn1(v20)) {
                                                v48 += 2
                                                v49 := []int{fi38, (^(fi38)), fi38, -128}
                                                var v50 bool = false
                                                v20 = (v20 << ((164) & 63))
                                                fmt.Println(v20)
                                                if false {
                                                    v50 = true
                                                } else {
                                                    v50 = v50
                                                    v51 := GenS3{F0: 0, F1: 255}
                                                    _ = v51
                                                }
                                                _ = v49
                                                _ = v50
                                            } else {
                                                if false {
                                                    v48 = (v48 | (^(v48)))
                                                } else {
                                                    if false {
                                                        v48 -= v48
                                                        var v52 uint = 1
                                                        v53 := GenS4{F0: -358}
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                }
                                                fmt.Println(v20, v20, v20)
                                            }
                                            switch "" {
                                            case "a":
                                                v54 := GenS5{F0: 896}
                                                _ = v54
                                            case "":
                                                var v55 bool = (-2.25 <= 0.0)
                                                _ = v55
                                            case "hello":
                                                var v56 int64 = 42
                                                _ = v56
                                            default:
                                                v20 = 45
                                            }
                                            _ = v47
                                            _ = v48
                                        }
                                    }
                                    fmt.Println(v2)
                                default:
                                    v2 = v2
                                    var v57 int8 = -128
                                    fmt.Println(fi37, v20)
                                    _ = v57
                                }
                            } else {
                                if true {
                                    fmt.Println(fi38, v2, v20)
                                    v58 := GenS1{F0: (-(fi37)), F1: ((true || true) || true)}
                                    v2 = -2.25
                                    _ = v58
                                }
                                if ((1.5 * 1e10) == (-2.25 * 1e10)) {
                                    v2 += v2
                                    var v59 float64 = 1.5
                                    v59 = ((((((1.5 - v59) + 0.1) - 1.5) * (v2 / (v2 + (v59 * 0.0)))) * (((v59 + float64(1.5)) + v59) / v2)) * 0.1)
                                    _ = v59
                                } else {
                                    v20 -= v20
                                }
                                var v60 float64 = float64(1.5)
                                _ = v60
                            }
                            switch "hello" {
                            case "foo bar":
                                var v61 float32 = 1e10
                                _ = v61
                            case "hello":
                                v2 += (1.5 + ((v2 - ((0.0 + ((float64(0.1) / (float64(1e10) + (1e10 - v2))) / v2)) / v2)) / (v2 / ((v2 - v2) * float64(0.1)))))
                            case "":
                                var v62 string = "foo bar"
                                _ = v62
                            default:
                                fmt.Println(fi38)
                            }
                        }
                        v2 += (v2 * v2)
                        var v63 string = "hello"
                        _ = v63
                    } else {
                        v64 := GenS6{F0: "", F1: genFn0(), F2: false}
                        _ = v64
                    }
                    _ = fi38
                    _ = v40
                }
                var v65 uint8 = 9
                v65 = 0
                _ = fi37
                _ = v65
            }
        }
    } else {
        if (("" + ("foo bar" + "foo bar")) >= "hello") {
            for fi66 := 0; fi66 < 2; fi66++ {
                switch "hello" {
                case "foo bar":
                    _ = len(v3)
                case "":
                    switch fi66 {
                    case -9223372036854775808:
                        for fi67 := 0; fi67 < 15; fi67++ {
                            var v68 int16 = 1000
                            v69 := GenS4{F0: -704}
                            v20 -= v20
                            _ = fi67
                            _ = v68
                            _ = v69
                        }
                    case 127:
                        _ = len(v3)
                    }
                default:
                    var v70 uint8 = 0
                    _ = v70
                }
                var v71 uint64 = 920
                _ = fi66
                _ = v71
            }
            var v72 uint = 1
            fmt.Println(v2)
            _ = v72
        }
    }
    for fi73 := 0; fi73 < 20; fi73++ {
        v20 = v20
        for fi74 := 0; fi74 < 19; fi74++ {
            if false {
                switch "a" {
                case "\n\t":
                    v75 := GenS2{F0: 4294967295}
                    _ = v75
                default:
                    var v76 uint = 42
                    _ = v76
                }
            }
            _ = fi74
        }
        _ = fi73
    }
    v20 = (-(v20))
    v2 -= (float64(3.14159) + v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
