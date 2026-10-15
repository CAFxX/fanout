package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 string; F1 int8; F2 int32 }
func genFn0(p0 int16, p1 bool, p2 int8) int16 {
    v6 := []int{}
    var v7 bool = p1
    var v8 int64 = -919
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return p0
}

func main() {
    v1 := []uint32{}
    var v2 int16 = 346
    var v3 int32 = 620
    v4 := []uint32{65535, 2147483647, 545}
    v5 := []string{("foo bar" + "hello"), "hello", "a"}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS1{F0: "foo bar"}
        if len(v4) > 0 {
            v4[0] = 42
        }
        var v11 string = (("\n\t" + "a") + "foo bar")
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: "foo bar"}
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    if ((true || false) && (false && false)) {
        v2 += (-(v2))
    }
    if true {
        v13 := GenS2{F0: true, F1: 2}
        if true {
            if false {
                v3 = v3
                if (!false) {
                    switch -1 {
                    case 32768:
                        for fi14 := 0; fi14 < 10; fi14++ {
                            var v15 int32 = (v3 / (((v3 - (v3 ^ v3))) | 1))
                            _ = fi14
                            _ = v15
                        }
                    default:
                        var v16 int32 = (^(v3))
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 1; fi17++ {
                        switch (fi17 | fi17) {
                        case -9223372036854775808:
                            var v18 int32 = (v3 << ((632) & 63))
                            _ = v18
                        case 32768:
                            v19 := GenS1{F0: "foo bar"}
                            _ = v19
                        case 345:
                            if true {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            } else {
                                v3 -= (v3 + -2147483648)
                                v20 := GenS2{F0: false, F1: 1}
                                for fi21 := 0; fi21 < 20; fi21++ {
                                    var v22 uint32 = 486
                                    var v23 int32 = (v3 >> ((0) & 63))
                                    _ = fi21
                                    _ = v22
                                    _ = v23
                                }
                                _ = v20
                            }
                        }
                        _ = fi17
                    }
                    v3 = 1
                }
                _ = len(v5)
            }
            fmt.Println(v2, v3)
        }
        fmt.Println(v2)
        _ = v13
    }
    v2 -= (-(v2))
    if len(v5) > 0 {
        v5[0] = (("\n\t" + "a") + "foo bar")
    }
    var v24 float32 = float32(-2.25)
    for fi25 := 0; fi25 < 3; fi25++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi26 := 0; fi26 < 11; fi26++ {
            _ = len(v4)
            var v27 string = ("hello" + "")
            fmt.Println(v27)
            if (v3 < v3) {
                v3 = (v3 * v3)
            } else {
                v28 := GenS1{F0: (v27 + v27)}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v29 := GenS3{F0: -175, F1: 42}
                _ = v28
                _ = v29
            }
            v27 = "foo bar"
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    var v30 float64 = 1.5
    v3 -= -2147483648
    for fi31 := 0; fi31 < 7; fi31++ {
        if (424 >= 2147483647) {
            if (false && (((-(v2)) <= -1) && true)) {
                _ = len(v5)
            } else {
                v32 := GenS1{F0: (((("foo bar" + "\n\t") + "") + "hello") + "a")}
                v33 := GenS4{F0: 255}
                _ = v32
                _ = v33
            }
            for fi34 := 0; fi34 < 12; fi34++ {
                var v35 uint = 568
                _ = fi34
                _ = v35
            }
            var v36 float32 = 1e10
            _ = v36
        } else {
            if (true && (2 == 18446744073709551615)) {
                fmt.Println(v30, v2, v2)
                var v37 uint = 1
                _ = v37
            } else {
                var v38 bool = (137 > 1)
                if (false && false) {
                    v38 = v38
                    var v39 int16 = -553
                    _ = v39
                } else {
                    v40 := GenS5{F0: 1, F1: 206}
                    if (215 != 2147483647) {
                        v41 := GenS6{F0: "", F1: 1, F2: (^(v3))}
                        fmt.Println(fi31, fi31)
                        _ = v41
                    }
                    if len(v4) > 0 {
                        v4[0] = 42
                    }
                    _ = v40
                }
                v24 += (v24 - 0.0)
                _ = v38
            }
            var v42 uint16 = 42
            _ = v42
        }
        _ = fi31
    }
    if (!(((-9223372036854775808 == 1) || (1 != 42)) && ("foo bar" <= (("foo bar" + ("" + "a")) + "a")))) {
        v24 -= v24
        var v43 string = ("\n\t" + ("hello" + ("hello" + ("" + "foo bar"))))
        if (1e10 != v24) {
            var v44 uint16 = 209
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v44
        }
        _ = v43
    }
    v45 := GenS4{F0: 32767}
    for fi46 := 0; fi46 < 17; fi46++ {
        switch fi46 {
        case 777:
            fmt.Println(v30, v30, v30)
        case 479:
            fmt.Println(v2, v2, v3)
        default:
            v2 = genFn0((v2 / ((genFn0(491, true, -1)) | 1)), (!(true && (!true))), 42)
        }
        if (v24 >= ((v24 / (3.14159 / 3.14159)) / v24)) {
            var v47 uint64 = 522
            _ = v47
        } else {
            if len(v1) > 0 {
                v1[0] = 42
            }
        }
        var v48 uint64 = 18446744073709551615
        _ = fi46
        _ = v48
    }
    for fi49 := 0; fi49 < 4; fi49++ {
        var v50 int16 = v2
        if ((genFn0((v2 * v2), ((fi49 - fi49) == fi49), 127) >= v50) && false) {
            v3 += (v3 << ((375) & 63))
            fmt.Println(v24, v30)
            v51 := GenS3{F0: -584, F1: genFn0(genFn0((-(v2)), true, -111), (990 == 0), 42)}
            _ = v51
        }
        _ = fi49
        _ = v50
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v45
    _ = v5
}
