package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint64; F1 bool; F2 uint }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 uint; F1 uint }
func genFn0(p0 int16) int8 {
    var v7 uint = 2
    p0 -= 151
    v7 = (^(v7))
    fmt.Println(p0, v7)
    _ = p0
    _ = v7
    return 100
}

func main() {
    var v1 int16 = 562
    var v2 uint32 = 2
    var v3 int8 = 127
    v4 := make(map[uint32]int)
    v5 := []uint32{(v2 | (v2 * 2))}
    v6 := make(map[uint32]string)
    var v8 int = 128
    for fi9 := 0; fi9 < 9; fi9++ {
        v1 -= -1
        _ = fi9
    }
    v10 := GenS1{F0: (!(42 >= -9223372036854775808)), F1: (v8 + -9223372036854775808)}
    fmt.Println(v3, v1, v2)
    v3 -= genFn0(0)
    v11 := GenS1{F0: true, F1: (v8 | v8)}
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := v4[v2]
        v14 := GenS1{F0: ((-2147483648 < (v13 + v8)) || (("hello" + (("\n\t" + (("\n\t" + ("hello" + "foo bar")) + (("hello" + "a") + ("\n\t" + "\n\t")))) + (("foo bar" + ((("" + ("" + "foo bar")) + ("a" + "")) + ("foo bar" + (((("\n\t" + "foo bar") + "") + "a") + ("" + "hello"))))) + "hello"))) >= (((("" + "foo bar") + ("hello" + "foo bar")) + "foo bar") + "hello"))), F1: -560}
        var v15 uint8 = 2
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    if true {
        var v16 float32 = (-2.25 + -2.25)
        _ = v16
    } else {
        v17 := v6[(v2 - (v2 % ((v2) | 1)))]
        if false {
            for fi18 := 0; fi18 < 10; fi18++ {
                if (!true) {
                    if false {
                        for fi19 := 0; fi19 < 15; fi19++ {
                            var v20 uint64 = 2
                            v21 := v6[v2]
                            fmt.Println(fi18, v17)
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                        fmt.Println(v3, v17, v1)
                    }
                    delete(v4, v2)
                } else {
                    fmt.Println(v8, v17)
                    v3 -= -128
                }
                v22 := GenS1{F0: false, F1: v8}
                if (128 >= 130) {
                    v3 = (^(v3))
                    v23 := GenS2{F0: 2}
                    for fi24 := 0; fi24 < 13; fi24++ {
                        v4[2] = (^(v8))
                        switch 862 {
                        case -994:
                            _ = len(v5)
                        case -9223372036854775808:
                            v25 := v6[(v2 >> ((65535) & 63))]
                            _ = v25
                        case 1:
                            v3 += (v3 * v3)
                        default:
                            if false {
                                var v26 int16 = v1
                                var v27 float64 = -2.25
                                switch v8 {
                                case 4294967295:
                                    if false {
                                        fmt.Println(v17, v17)
                                        v28 := GenS2{F0: 2}
                                        var v29 int32 = 42
                                        _ = v28
                                        _ = v29
                                    }
                                case 42:
                                    var v30 int8 = v3
                                    _ = v30
                                case -9223372036854775808:
                                    var v31 int64 = -613
                                    _ = v31
                                }
                                switch fi24 {
                                case -967:
                                    var v32 int8 = genFn0(v1)
                                    _ = v32
                                case 1:
                                    v26 = v1
                                }
                                _ = v26
                                _ = v27
                            } else {
                                switch v17 {
                                case "a":
                                    var v33 int64 = -300
                                    _ = v33
                                case "foo bar":
                                    if (false && false) {
                                        v34 := GenS3{F0: (fi24 >> ((781) & 63)), F1: (v3 - v3)}
                                        _ = v34
                                    }
                                default:
                                    v8 += v8
                                }
                                if len(v5) > 0 {
                                    v5[0] = (v2 ^ v2)
                                }
                                if false {
                                    v17 = (v17 + "foo bar")
                                } else {
                                    var v35 int16 = (v1 / (((v1 % ((v1) | 1))) | 1))
                                    fmt.Println(fi24, v2, v8)
                                    _ = v35
                                }
                                _ = len(v6)
                                if false {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    v36 := GenS4{F0: 42, F1: (true || ("hello" <= v17)), F2: 18446744073709551615}
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = v36
                                } else {
                                    switch fi18 {
                                    case 32768:
                                        fmt.Println(v17, v2)
                                    default:
                                        v8 += fi24
                                    }
                                    switch "" {
                                    case "\n\t":
                                        v2 -= v2
                                    case "a":
                                        var v37 float64 = -2.25
                                        _ = v37
                                    case "":
                                        if true {
                                            v38 := GenS3{F0: fi18, F1: v3}
                                            if true {
                                                v39 := v6[2147483647]
                                                v8 = (v8 % ((v8) | 1))
                                                _ = v39
                                            }
                                            v40 := GenS5{F0: v1, F1: ("foo bar" + v17)}
                                            _ = v38
                                            _ = v40
                                        } else {
                                            fmt.Println(fi18, fi24)
                                            v1 = v1
                                            if (true || (42 != 0)) {
                                                switch "hello" {
                                                case "\n\t":
                                                    var v41 uint16 = 42
                                                    _ = v41
                                                default:
                                                    v42 := v4[v2]
                                                    _ = v42
                                                }
                                                if true {
                                                    var v43 int64 = -660
                                                    switch v17 {
                                                    case "hello":
                                                        v43 += (v43 * v43)
                                                    }
                                                    switch "foo bar" {
                                                    case "hello":
                                                        v43 -= int64(v2)
                                                    case "a":
                                                        fmt.Println(fi18)
                                                    case "foo bar":
                                                        fmt.Println(v3, v3)
                                                    default:
                                                        if len(v5) > 0 {
                                                            v5[0] = v2
                                                        }
                                                    }
                                                    _ = v43
                                                }
                                                var v44 int = fi24
                                                _ = v44
                                            } else {
                                                v6[(v2 & (v2 & 2))] = (v17 + ((v17 + v17) + ("\n\t" + v17)))
                                            }
                                        }
                                    default:
                                        var v45 int = (fi18 / ((fi24) | 1))
                                        _ = v45
                                    }
                                    v46 := GenS6{F0: 255, F1: 647}
                                    _ = v46
                                }
                            }
                        }
                        v2 = 42
                        _ = fi24
                    }
                    _ = v23
                }
                _ = fi18
                _ = v22
            }
        }
        switch v17 {
        case "\n\t":
            v47 := GenS3{F0: (v8 << ((65535) & 63)), F1: v3}
            _ = v47
        case "foo bar":
            v3 += genFn0(-1)
        case "":
            for fi48 := 0; fi48 < 2; fi48++ {
                v8 = int(v2)
                _ = fi48
            }
        default:
            v17 = v17
        }
        _ = v17
    }
    var v49 bool = false
    switch (-(v8)) {
    case 0:
        if (!(true && false)) {
            for fi50 := 0; fi50 < 6; fi50++ {
                for fi51 := 0; fi51 < 11; fi51++ {
                    var v52 int64 = -1
                    v53 := GenS3{F0: fi50, F1: (-(v3))}
                    v49 = (2 != 1)
                    v54 := v6[v2]
                    v6[(v2 / ((v2) | 1))] = ((v54 + ("\n\t" + v54)) + "\n\t")
                    _ = fi51
                    _ = v52
                    _ = v53
                    _ = v54
                }
                _ = fi50
            }
            v49 = (!(96 == 128))
        }
    case 364:
        var v55 int64 = 0
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v8
}
