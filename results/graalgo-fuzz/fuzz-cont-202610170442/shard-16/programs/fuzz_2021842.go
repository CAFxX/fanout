package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS5 struct { F0 bool; F1 bool }
type GenS6 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS7 struct { F0 uint; F1 int64 }
type GenS8 struct { F0 uint16; F1 string; F2 uint64 }
func genFn0(p0 uint, p1 uint16, p2 bool) int8 {
    v5 := GenS1{F0: (("hello" + "a") + (((("\n\t" + "hello") + "hello") + "\n\t") + ("a" + "a")))}
    p2 = (2 < 2)
    var v6 float32 = 1e10
    var v7 int8 = -128
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return (-(v7))
}
func genFn1(p0 uint32) int8 {
    var v8 uint32 = (^(p0))
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 bool = ((214 == 651) || (-606 <= 1))
        _ = fi9
        _ = v10
    }
    switch -9223372036854775808 {
    case 1:
        v11 := []int{}
        _ = v11
    default:
        fmt.Println(v8, v8)
    }
    if (128 == 127) {
        v12 := GenS1{F0: "foo bar"}
        _ = v12
    } else {
        var v13 int64 = 42
        var v14 float64 = (float64(-2.25) * 1e10)
        _ = v13
        _ = v14
    }
    _ = p0
    _ = v8
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]uint32)
    var v3 float64 = (-2.25 + 0.0)
    var v4 float32 = 1e10
    fmt.Println(v4, v3)
    switch 42 {
    case 2147483647:
        var v15 string = "hello"
        _ = v15
    case 127:
        v3 -= (3.14159 / float64(v4))
    }
    var v16 uint = 42
    var v17 float32 = v4
    v4 = float32(v3)
    var v18 uint = 255
    v18 -= v16
    switch ((("" + (("\n\t" + ("hello" + "hello")) + ("hello" + ("a" + "foo bar")))) + (("foo bar" + "a") + "foo bar")) + "foo bar") {
    case "foo bar":
        for fi19 := 0; fi19 < 3; fi19++ {
            v20 := GenS2{F0: 2147483648}
            var v21 float32 = v4
            var v22 uint = (^(v16))
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
    default:
        switch 32767 {
        case -2147483648:
            if false {
                if false {
                    if (1 >= genFn0(v18, 0, ("foo bar" > ("hello" + ((("foo bar" + "") + "") + ("foo bar" + "a")))))) {
                        var v23 int64 = -9223372036854775808
                        if true {
                            v18 += 2
                            v24 := GenS2{F0: 181}
                            _ = v24
                        }
                        if true {
                            switch "" {
                            case "foo bar":
                                v25 := GenS3{F0: -267, F1: 446}
                                _ = v25
                            case "":
                                v16 = v18
                            }
                            var v26 bool = (2 <= (v18 << ((18446744073709551615) & 63)))
                            v27 := v1[2147483647]
                            _ = v26
                            _ = v27
                        }
                        _ = v23
                    }
                    if (!true) {
                        delete(v2, 183)
                        var v28 int16 = -32768
                        _ = v28
                    }
                    v29 := []int{182, -2147483648, 0, -2147483648}
                    _ = v29
                }
                fmt.Println(v17, v18, v4)
            }
        default:
            fmt.Println(v16, v16, v3)
        }
    }
    var v30 float32 = 0.0
    for fi31 := 0; fi31 < 19; fi31++ {
        var v32 int16 = 0
        v33 := v1[42]
        _ = fi31
        _ = v32
        _ = v33
    }
    switch "hello" {
    case "\n\t":
        switch 32767 {
        case 128:
            for fi34 := 0; fi34 < 13; fi34++ {
                for fi35 := 0; fi35 < 6; fi35++ {
                    v3 += v3
                    v18 = v16
                    _ = fi35
                }
                if false {
                    var v36 int64 = 9223372036854775807
                    for fi37 := 0; fi37 < 5; fi37++ {
                        var v38 int8 = genFn0(v18, 1, (!true))
                        switch "hello" {
                        case "foo bar":
                            v39 := GenS4{F0: 123, F1: false, F2: 1}
                            _ = v39
                        default:
                            if false {
                                v38 = v38
                                var v40 int16 = -134
                                switch "foo bar" {
                                case "\n\t":
                                    v41 := GenS1{F0: "foo bar"}
                                    _ = v41
                                case "foo bar":
                                    v42 := GenS3{F0: 420, F1: -116}
                                    switch -432 {
                                    case -555:
                                        v17 = 1.5
                                    case 781:
                                        v16 += v16
                                        var v43 float32 = 1.5
                                        _ = v43
                                    case 42:
                                        if ((^(v38)) >= int8(v40)) {
                                            v38 = genFn0(140, 614, true)
                                            v30 = v4
                                            v30 -= float32(float64(v17))
                                        }
                                    default:
                                        v44 := []int{9223372036854775807, fi34, (fi37 >> ((v16) & 63)), (fi34 * 121)}
                                        _ = v44
                                    }
                                    _ = v42
                                default:
                                    v30 = (v17 - 3.14159)
                                }
                                v45 := GenS4{F0: 42, F1: (v38 <= (v38 | genFn1(uint32(v18)))), F2: 76}
                                var v46 float64 = 0.0
                                _ = v40
                                _ = v45
                                _ = v46
                            }
                        }
                        _ = fi37
                        _ = v38
                    }
                    v47 := GenS1{F0: ("foo bar" + "")}
                    _ = v36
                    _ = v47
                } else {
                    v48 := GenS5{F0: (true || (((458 > 0) && true) && (22 > 24))), F1: false}
                    var v49 float64 = v3
                    _ = v48
                    _ = v49
                }
                var v50 uint32 = 0
                _ = fi34
                _ = v50
            }
        case 32767:
            v51 := GenS4{F0: uint32(v16), F1: true, F2: uint16(v18)}
            _ = v51
        default:
            v18 = (v18 % ((4294967295) | 1))
        }
    }
    if false {
        if false {
            for fi52 := 0; fi52 < 17; fi52++ {
                for fi53 := 0; fi53 < 15; fi53++ {
                    v54 := []int{(fi52 * fi52)}
                    v18 = (^(v16))
                    if false {
                        var v55 uint16 = 177
                        v56 := GenS5{F0: false, F1: false}
                        v57 := GenS4{F0: 2, F1: false, F2: uint16(v18)}
                        _ = v55
                        _ = v56
                        _ = v57
                    } else {
                        v58 := GenS6{F0: 1, F1: 1, F2: -322}
                        var v59 int8 = -33
                        var v60 int32 = -1
                        _ = v58
                        _ = v59
                        _ = v60
                    }
                    _ = len(v54)
                    switch ("foo bar" + ("foo bar" + ("foo bar" + ""))) {
                    case "a":
                        fmt.Println(v30, v4)
                    case "\n\t":
                        fmt.Println(v30, fi52)
                    }
                    _ = fi53
                    _ = v54
                }
                delete(v1, 959)
                _ = fi52
            }
            if false {
                v61 := GenS7{F0: v18, F1: -1}
                var v62 int = 2147483647
                fmt.Println(v4, v18, v4)
                _ = v61
                _ = v62
            }
            var v63 float32 = float32((-2.25 - ((((0.1 / v3) / (1e10 / 1e10)) + v3) + (v3 / 0.1))))
            _ = v63
        } else {
            v4 -= (v30 + (((-2.25 + 1e10) - v4) + (v17 / (float32(v3) - ((v17 + v30) / v30)))))
            v18 = v16
        }
        v64 := GenS8{F0: 2, F1: ((("" + "") + "hello") + ("foo bar" + ("foo bar" + "\n\t"))), F2: 872}
        _ = v64
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
