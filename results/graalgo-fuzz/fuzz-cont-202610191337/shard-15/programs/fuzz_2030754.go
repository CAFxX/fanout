package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint; F1 int32 }
func genFn0() int {
    var v7 uint = 1
    switch 4294967295 {
    case 32767:
        var v8 uint16 = 0
        _ = v8
    case 4294967295:
        var v9 uint64 = 42
        _ = v9
    }
    v7 = v7
    if true {
        if (!(!false)) {
            switch 4294967295 {
            case 42:
                v10 := []int{653, 128, 128, 9223372036854775807}
                _ = v10
            default:
                v7 -= v7
            }
            var v11 uint = (^(v7))
            v11 += (v11 >> ((349) & 63))
            _ = v11
        }
    } else {
        if (float32(((0.0 - 1e10) - 3.14159)) < (1.5 + 3.14159)) {
            for fi12 := 0; fi12 < 7; fi12++ {
                v7 += v7
                var v13 uint16 = 491
                _ = fi12
                _ = v13
            }
            var v14 uint = v7
            for fi15 := 0; fi15 < 9; fi15++ {
                if true {
                    switch "" {
                    case "":
                        switch ((("\n\t" + "foo bar") + "\n\t") + ((("" + "") + "") + ("hello" + ((((("\n\t" + ("hello" + (("foo bar" + "a") + ""))) + ((((("hello" + ("\n\t" + "\n\t")) + "hello") + ("\n\t" + "hello")) + "\n\t") + ("a" + (("foo bar" + "a") + "\n\t")))) + ("\n\t" + "hello")) + "a") + "\n\t")))) {
                        case "a":
                            var v16 int32 = -2147483648
                            _ = v16
                        case "\n\t":
                            for fi17 := 0; fi17 < 6; fi17++ {
                                if false {
                                    var v18 uint32 = 441
                                    v7 = v14
                                    v7 += 874
                                    v7 = v14
                                    var v19 int = (fi15 + fi17)
                                    _ = v18
                                    _ = v19
                                } else {
                                    var v20 int32 = 2147483647
                                    v7 = 0
                                    _ = v20
                                }
                                _ = fi17
                            }
                        default:
                            for fi21 := 0; fi21 < 13; fi21++ {
                                var v22 uint32 = 1
                                v23 := []int{}
                                var v24 uint8 = 127
                                var v25 uint8 = (v24 | (v24 << ((v14) & 63)))
                                if (2147483647 < (fi21 + -9223372036854775808)) {
                                    fmt.Println(v22, v22, v7)
                                } else {
                                    v7 = 2
                                    var v26 int8 = -126
                                    if true {
                                        v24 += v25
                                    } else {
                                        v24 += 42
                                        fmt.Println(v7)
                                        fmt.Println(v24, fi21)
                                        if (int8(v22) < (v26 ^ 1)) {
                                            if true {
                                                var v27 uint32 = v22
                                                fmt.Println(v24, v14, v25)
                                                _ = v27
                                            } else {
                                                var v28 int64 = -1
                                                v29 := GenS1{F0: (v22 % (((^(v22))) | 1)), F1: 1000}
                                                v30 := GenS1{F0: v22, F1: 1000}
                                                _ = v28
                                                _ = v29
                                                _ = v30
                                            }
                                            v31 := GenS1{F0: 857, F1: -1}
                                            _ = v31
                                        }
                                    }
                                    _ = v26
                                }
                                if (-2.25 > 3.14159) {
                                    if false {
                                        if len(v23) >= 2 {
                                            v23 = v23[1:len(v23)]
                                        }
                                    } else {
                                        fmt.Println(fi21, v14, v25)
                                        fmt.Println(v7, v7)
                                    }
                                }
                                v32 := GenS2{F0: 107, F1: 455, F2: (v22 ^ v22)}
                                v14 = 818
                                _ = fi21
                                _ = v22
                                _ = v23
                                _ = v24
                                _ = v25
                                _ = v32
                            }
                        }
                    }
                } else {
                    v33 := GenS3{F0: -1}
                    var v34 uint = (v14 - (v7 / ((v14) | 1)))
                    for fi35 := 0; fi35 < 2; fi35++ {
                        var v36 float32 = 3.14159
                        v7 += (^(v7))
                        if false {
                            v37 := GenS3{F0: -7}
                            fmt.Println(v14, v34, v7)
                            switch 42 {
                            case 9223372036854775807:
                                fmt.Println(v14, fi35, v7)
                            default:
                                var v38 float64 = ((-2.25 * -2.25) / (1.5 + 0.1))
                                _ = v38
                            }
                            v7 = v34
                            _ = v37
                        } else {
                            v14 = (v34 / ((18446744073709551615) | 1))
                            v39 := GenS4{F0: 9223372036854775807}
                            var v40 uint = v34
                            _ = v39
                            _ = v40
                        }
                        var v41 float64 = (1e10 * 1.5)
                        v42 := GenS5{F0: 42}
                        _ = fi35
                        _ = v36
                        _ = v41
                        _ = v42
                    }
                    _ = v33
                    _ = v34
                }
                v14 += v7
                _ = fi15
            }
            _ = v14
        } else {
            fmt.Println(v7)
            if (false && false) {
                if ((-1 > 1) || true) {
                    var v43 uint32 = 0
                    _ = v43
                } else {
                    fmt.Println(v7)
                    v44 := GenS6{F0: (v7 >> ((42) & 63)), F1: 1}
                    _ = v44
                }
                v7 -= (v7 - 796)
                v45 := GenS1{F0: 2, F1: -1}
                _ = v45
            } else {
                v7 = (v7 >> ((980) & 63))
                fmt.Println(v7)
                v46 := GenS3{F0: 42}
                _ = v46
            }
        }
        var v47 uint32 = 374
        _ = v47
    }
    switch ("\n\t" + "") {
    case "hello":
        v7 = v7
    case "":
        v48 := []int{0, -947}
        _ = v48
    case "a":
        v49 := []int{-2147483648}
        _ = v49
    default:
        v50 := []int{127, 128, 759, -466}
        _ = v50
    }
    _ = v7
    return 321
}
func genFn1(p0 int) uint8 {
    var v51 uint32 = 2
    v52 := []int{42}
    v53 := GenS3{F0: -128}
    _ = len(v52)
    _ = p0
    _ = v51
    _ = v52
    _ = v53
    return 127
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{42}
    var v3 uint32 = 95
    v4 := make(map[uint32]int)
    var v5 int32 = -965
    var v6 float64 = 1e10
    v5 += v5
    v54 := GenS2{F0: 42, F1: 42, F2: (v3 - 42)}
    var v55 uint16 = 75
    fmt.Println(v3, v5)
    v56 := GenS1{F0: (v3 | v3), F1: -971}
    var v57 uint8 = 255
    var v58 float32 = 0.0
    v58 = (0.1 * v58)
    var v59 uint8 = (^(v57))
    var v60 string = "a"
    switch 32768 {
    case -146:
        v3 -= 1
    case -524:
        var v61 int64 = 9223372036854775807
        _ = v61
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v60
}
