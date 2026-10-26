package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 uint32; F2 int8 }
func genFn0() uint16 {
    v4 := []int{}
    var v5 int8 = -121
    for fi6 := 0; fi6 < 18; fi6++ {
        if (true && (((1.5 + 0.0) - float32(-2.25)) > 0.0)) {
            var v7 float64 = 0.0
            var v8 uint16 = 823
            fmt.Println(v8, fi6, fi6)
            _ = v7
            _ = v8
        } else {
            v9 := GenS1{F0: 42, F1: 0}
            _ = v9
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi10 := 0; fi10 < 18; fi10++ {
            v5 = 42
            var v11 int16 = 838
            if len(v4) > 0 {
                v4[0] = fi6
            }
            _ = fi10
            _ = v11
        }
        _ = fi6
    }
    v5 -= (v5 + 127)
    _ = v4
    _ = v5
    return 255
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = 0
    var v3 uint64 = 1
    for fi12 := 0; fi12 < 13; fi12++ {
        v13 := v1[4294967295]
        v3 += (^(v3))
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        switch 1 {
        case 42:
            if (("\n\t" + (((("" + "hello") + "hello") + ("" + "a")) + "a")) < "hello") {
                if false {
                    for fi15 := 0; fi15 < 20; fi15++ {
                        fmt.Println(fi14)
                        if ((!(false || true)) && ((false && false) && true)) {
                            v16 := []int{fi15, int(v2), 300, (fi14 << ((65535) & 63))}
                            if (!(false && ((910 == 474) || false))) {
                                v3 -= v3
                                var v17 string = "foo bar"
                                _ = v17
                            } else {
                                v3 = v3
                            }
                            _ = v16
                        }
                        fmt.Println(v2)
                        _ = fi15
                    }
                    var v18 uint32 = 414
                    for fi19 := 0; fi19 < 7; fi19++ {
                        switch "foo bar" {
                        case "hello":
                            v20 := v1[(v18 + v18)]
                            _ = v20
                        default:
                            if (!(!false)) {
                                if (1.5 <= 1e10) {
                                    if (323 < 930) {
                                        v21 := GenS1{F0: 86, F1: 255}
                                        v3 -= (v3 + (v3 * (v3 >> ((42) & 63))))
                                        switch (("hello" + "foo bar") + "foo bar") {
                                        case "foo bar":
                                            v3 += 4294967295
                                        case "hello":
                                            if false {
                                                v18 -= (v18 - v18)
                                                v18 += 434
                                                v3 = v3
                                            } else {
                                                if (!(948 < -162)) {
                                                    if (1.5 <= 0.0) {
                                                        if (true || true) {
                                                            fmt.Println(v18, fi14)
                                                            delete(v1, (v18 * 1))
                                                        } else {
                                                            var v22 bool = ((1e10 / ((((-2.25 + 0.0) - 1.5) / 0.1) + (0.1 / 1.0))) > (0.1 - (-2.25 + 0.0)))
                                                            _ = v22
                                                        }
                                                        v2 += 127
                                                        v23 := GenS2{F0: -128, F1: -668}
                                                        _ = v23
                                                    }
                                                }
                                                v18 -= v18
                                            }
                                        case "\n\t":
                                            v24 := GenS2{F0: (v2 << ((1) & 63)), F1: 0}
                                            _ = v24
                                        default:
                                            v18 = v18
                                        }
                                        _ = v21
                                    }
                                    switch "a" {
                                    case "":
                                        if true {
                                            fmt.Println(fi19, v18)
                                            v25 := v1[(v18 / ((v18) | 1))]
                                            var v26 uint8 = 255
                                            _ = v25
                                            _ = v26
                                        } else {
                                            if true {
                                                if false {
                                                    v27 := v1[2147483648]
                                                    _ = v27
                                                }
                                            } else {
                                                if false {
                                                    switch 42 {
                                                    case 127:
                                                        v2 = (-(v2))
                                                    default:
                                                        switch "hello" {
                                                        case "foo bar":
                                                            switch "" {
                                                            case "a":
                                                                _ = len(v1)
                                                            case "\n\t":
                                                                fmt.Println(fi19, fi14)
                                                            default:
                                                                v2 -= v2
                                                            }
                                                        case "\n\t":
                                                            v2 -= (v2 / ((v2) | 1))
                                                        case "":
                                                            var v28 float64 = float64(0.1)
                                                            _ = v28
                                                        }
                                                    }
                                                    if false {
                                                        fmt.Println(v3)
                                                    } else {
                                                        var v29 string = (("\n\t" + "\n\t") + "\n\t")
                                                        var v30 bool = (3.14159 <= 3.14159)
                                                        _ = v29
                                                        _ = v30
                                                    }
                                                } else {
                                                    v31 := GenS1{F0: v3, F1: 693}
                                                    _ = v31
                                                }
                                            }
                                            switch fi19 {
                                            case 301:
                                                var v32 uint16 = genFn0()
                                                _ = v32
                                            default:
                                                fmt.Println(v2, v18, fi19)
                                            }
                                        }
                                    case "\n\t":
                                        var v33 uint32 = v18
                                        _ = v33
                                    default:
                                        v3 -= 42
                                        var v34 string = (("foo bar" + (("foo bar" + ("a" + "a")) + "a")) + (("foo bar" + "\n\t") + ("a" + "\n\t")))
                                        _ = v34
                                    }
                                    if true {
                                        if true {
                                            v35 := GenS3{F0: 1}
                                            v36 := GenS4{F0: 2147483647, F1: v18, F2: v2}
                                            v3 += v3
                                            _ = v35
                                            _ = v36
                                        }
                                        if false {
                                            v18 = (v18 - (v18 - (v18 - v18)))
                                            var v37 uint8 = 42
                                            _ = v37
                                        } else {
                                            var v38 uint16 = 546
                                            var v39 uint16 = genFn0()
                                            _ = v38
                                            _ = v39
                                        }
                                    }
                                }
                            } else {
                                var v40 bool = true
                                _ = v40
                            }
                        }
                        var v41 uint16 = 149
                        var v42 int32 = 1
                        _ = fi19
                        _ = v41
                        _ = v42
                    }
                    _ = v18
                } else {
                    var v43 int64 = 1
                    var v44 int = fi14
                    _ = v43
                    _ = v44
                }
            }
        case 128:
            v3 = (^(v3))
        case 201:
            v2 = v2
        default:
            if true {
                v1[2147483647] = (fi14 * fi14)
                fmt.Println(v3)
            } else {
                var v45 int32 = 0
                _ = v45
            }
        }
        v1[2147483647] = (fi14 & 4294967295)
        v3 = v3
        _ = fi14
    }
    var v46 int32 = 42
    v3 -= (v3 / ((411) | 1))
    var v47 uint32 = 4294967295
    if (!false) {
        var v48 string = "foo bar"
        for fi49 := 0; fi49 < 18; fi49++ {
            var v50 int64 = 1
            _ = fi49
            _ = v50
        }
        _ = v48
    } else {
        var v51 bool = false
        _ = v51
    }
    v52 := GenS2{F0: -1, F1: 1}
    _ = len(v1)
    var v53 uint16 = genFn0()
    var v54 int8 = (v2 * v2)
    v55 := GenS2{F0: -44, F1: -1}
    fmt.Println(v3, v53, v53)
    var v56 uint = 4294967295
    v3 += v3
    v53 += 1
    v57 := GenS1{F0: 95, F1: 821}
    v58 := []int{int(v56), 876}
    switch "\n\t" {
    case "a":
        switch -906 {
        case 9223372036854775807:
            var v59 float64 = float64(((float32(1e10) - 0.1) + (-2.25 / (-2.25 * -2.25))))
            _ = v59
        case -732:
            var v60 float32 = 1.5
            _ = v60
        }
    case "foo bar":
        var v61 float64 = 0.1
        _ = v61
    case "":
        switch 128 {
        case 0:
            var v62 int16 = 0
            _ = v62
        default:
            v54 += int8(v46)
        }
    default:
        v2 = (v2 + v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v58:", len(v58), v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v46
    _ = v47
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v58
}
