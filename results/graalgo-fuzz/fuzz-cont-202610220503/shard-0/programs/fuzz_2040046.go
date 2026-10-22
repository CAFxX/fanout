package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 bool; F1 uint64; F2 bool }
type GenS3 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS4 struct { F0 uint16; F1 bool; F2 int64 }
type GenS5 struct { F0 uint32 }
func genFn0() string {
    var v7 uint8 = 1
    v7 -= 1
    v7 = v7
    if true {
        var v8 bool = (32768 <= 128)
        _ = v8
    }
    _ = v7
    return "\n\t"
}

func main() {
    var v1 string = "foo bar"
    var v2 bool = (42 > 541)
    var v3 int32 = 196
    var v4 int = 9223372036854775807
    v5 := make(map[int]bool)
    var v6 int32 = v3
    for fi9 := 0; fi9 < 18; fi9++ {
        switch (v4 % ((1) | 1)) {
        case 32767:
            for fi10 := 0; fi10 < 15; fi10++ {
                if (fi10 != (v4 | fi10)) {
                    var v11 uint64 = 509
                    if (!(!(!false))) {
                        if ((42 > 430) || ((-2.25 - -2.25) >= (1.5 * float64(1e10)))) {
                            if ((!v2) || v2) {
                                var v12 int16 = -277
                                var v13 uint8 = 127
                                v13 += v13
                                v14 := v5[fi10]
                                _ = v12
                                _ = v13
                                _ = v14
                            } else {
                                var v15 int64 = 42
                                _ = v15
                            }
                            var v16 int32 = 2147483647
                            v17 := []int{fi9, v4}
                            var v18 string = (v1 + v1)
                            if (v2 && ((!v2) && v2)) {
                                v3 = v16
                            }
                            v3 = (v6 ^ v6)
                            var v19 float64 = 0.1
                            _ = v16
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                        v20 := []int{(-(fi9))}
                        var v21 int32 = (v3 | (v3 - (-(v3))))
                        _ = v20
                        _ = v21
                    } else {
                        v2 = v2
                        v11 += (v11 * v11)
                        var v22 float64 = (1e10 / 3.14159)
                        _ = v22
                    }
                    if ((fi9 & v4) <= 9223372036854775807) {
                        v11 -= v11
                    } else {
                        fmt.Println(v2, fi10)
                        v23 := GenS1{F0: 1, F1: -532}
                        _ = v23
                    }
                    _ = v11
                } else {
                    v2 = true
                }
                var v24 uint8 = 91
                var v25 uint8 = v24
                _ = fi10
                _ = v24
                _ = v25
            }
        case -2147483648:
            v5[(fi9 & (fi9 % (((v4 & (v4 / ((371) | 1)))) | 1)))] = (!v2)
        default:
            v3 -= (v3 - (v3 + 2147483647))
        }
        var v26 float64 = 0.0
        _ = fi9
        _ = v26
    }
    var v27 int64 = -1
    v3 = int32(v27)
    v28 := GenS1{F0: 0, F1: 0}
    var v29 uint16 = 42
    v30 := GenS1{F0: 1, F1: -24}
    for fi31 := 0; fi31 < 6; fi31++ {
        for fi32 := 0; fi32 < 15; fi32++ {
            if v2 {
                var v33 float64 = 1e10
                fmt.Println(fi32)
                if ((!((v1 < v1) && false)) && ((23 != 48) || v2)) {
                    v2 = v2
                } else {
                    v34 := GenS2{F0: v2, F1: 772, F2: (((v33 / v33) > v33) && v2)}
                    v35 := GenS3{F0: 1, F1: int16(v27), F2: 2147483648}
                    v6 = v6
                    _ = v34
                    _ = v35
                }
                fmt.Println(fi31)
                _ = v33
            } else {
                v5[fi31] = v2
                if v2 {
                    v1 = v1
                    v5[0] = (1.5 > -2.25)
                } else {
                    if (((-2.25 * (3.14159 * -2.25)) / (-2.25 + (((1.5 * ((0.0 / 1.0) + -2.25)) - 3.14159) + 0.0))) == 0.0) {
                        v36 := GenS3{F0: 437, F1: -1, F2: 2}
                        switch ((v1 + v1) + ("\n\t" + v1)) {
                        case "a":
                            switch ((((v1 + v1) + ((((v1 + v1) + v1) + v1) + (v1 + ((v1 + "\n\t") + (v1 + "hello"))))) + v1) + "foo bar") {
                            case "\n\t":
                                _ = len(v5)
                            case "a":
                                var v37 uint32 = 0
                                _ = v37
                            case "hello":
                                var v38 float32 = (-2.25 * 1.5)
                                _ = v38
                            default:
                                var v39 float64 = (0.0 + -2.25)
                                _ = v39
                            }
                            v3 += v3
                            v40 := GenS4{F0: (v29 / (((v29 ^ (v29 % (((v29 + 42)) | 1)))) | 1)), F1: (false || (!(0.1 == 1.5))), F2: v27}
                            _ = v40
                        case "":
                            var v41 bool = v2
                            var v42 uint = 305
                            v2 = (!(v29 < 0))
                            v43 := []int{fi31, v4, fi31}
                            _ = v41
                            _ = v42
                            _ = v43
                        }
                        v27 = v27
                        _ = v36
                    }
                    fmt.Println(fi32, v2)
                }
            }
            var v44 float64 = -2.25
            if (!(100 <= 127)) {
                v29 += v29
                v45 := GenS3{F0: 18446744073709551615, F1: 197, F2: 1}
                _ = v45
            } else {
                fmt.Println(v2, v44)
                if (-1 == 42) {
                    if true {
                        v46 := GenS5{F0: 2}
                        fmt.Println(v3, v4, v6)
                        v47 := GenS4{F0: (v29 >> ((65535) & 63)), F1: v2, F2: v27}
                        _ = v46
                        _ = v47
                    } else {
                        if (1 != 32767) {
                            var v48 string = (v1 + ("a" + v1))
                            _ = v48
                        } else {
                            v2 = v2
                            v1 = ("hello" + "")
                            fmt.Println(v27, v1, v27)
                        }
                        var v49 uint64 = 4294967295
                        var v50 int32 = v6
                        _ = v49
                        _ = v50
                    }
                    v51 := GenS2{F0: ((v29 ^ (v29 >> ((1) & 63))) < (v29 << ((42) & 63))), F1: 346, F2: false}
                    if v2 {
                        if (false || (v2 && (((v27 >> ((296) & 63)) < v27) && true))) {
                            if (417 == 2) {
                                var v52 int8 = -65
                                switch "foo bar" {
                                case "":
                                    var v53 int = (fi31 / ((fi32) | 1))
                                    _ = v53
                                }
                                var v54 float64 = 0.0
                                _ = v52
                                _ = v54
                            }
                            v29 -= v29
                            v55 := GenS4{F0: (v29 << ((500) & 63)), F1: ((0 <= 937) || v2), F2: v27}
                            fmt.Println(v4, v6, v4)
                            _ = v55
                        }
                        v6 = (-(v3))
                    } else {
                        var v56 uint64 = 124
                        v4 = (-(fi32))
                        switch (v1 + v1) {
                        case "\n\t":
                            var v57 uint = 18446744073709551615
                            _ = v57
                        case "foo bar":
                            var v58 int8 = 100
                            var v59 uint16 = (v29 ^ 32767)
                            v60 := GenS2{F0: (v2 && (v58 > v58)), F1: (v56 / ((v56) | 1)), F2: v2}
                            _ = v58
                            _ = v59
                            _ = v60
                        }
                        _ = v56
                    }
                    _ = v51
                }
            }
            _ = fi32
            _ = v44
        }
        var v61 int8 = 100
        _ = fi31
        _ = v61
    }
    for fi62 := 0; fi62 < 20; fi62++ {
        if false {
            v29 = (v29 << ((1) & 63))
        }
        _ = fi62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
