package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 int32; F2 uint32 }
type GenS5 struct { F0 int8; F1 int; F2 int16 }
type GenS6 struct { F0 uint8; F1 bool }
func genFn0() string {
    v5 := []int{-9223372036854775808, -2147483648, 32768}
    var v6 bool = (!(false && (("" + "\n\t") < "hello")))
    v7 := GenS1{F0: 1}
    var v8 int8 = 95
    switch "\n\t" {
    case "foo bar":
        v9 := GenS1{F0: (v8 * v8)}
        _ = v9
    case "a":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "\n\t":
        v10 := GenS1{F0: v8}
        _ = v10
    default:
        if len(v5) > 0 {
            v5[0] = 1
        }
    }
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return "\n\t"
}

func main() {
    v1 := []uint32{2, 2147483648}
    var v2 int16 = -960
    v3 := make(map[uint32]uint32)
    var v4 uint64 = 18446744073709551615
    var v11 uint8 = 108
    v12 := GenS1{F0: 110}
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    v11 = v11
    v2 = v2
    var v13 int = 4294967295
    var v14 int8 = -116
    switch ("hello" + (("" + "hello") + ("a" + (((("" + "hello") + "\n\t") + ("hello" + ("" + "\n\t"))) + ("a" + "\n\t"))))) {
    case "a":
        v15 := GenS1{F0: int8(v4)}
        _ = v15
    case "\n\t":
        if true {
            if ((1.5 >= (-2.25 * float64((float32(((1.5 / 3.14159) * 0.0)) + (0.1 / 1.5))))) || true) {
                v13 = 32768
                for fi16 := 0; fi16 < 10; fi16++ {
                    v2 = v2
                    for fi17 := 0; fi17 < 18; fi17++ {
                        if ((32768 >= 42) && true) {
                            var v18 uint64 = (v4 * (v4 | (v4 % ((v4) | 1))))
                            switch ("a" + "a") {
                            case "a":
                                var v19 uint = 604
                                switch "a" {
                                case "":
                                    var v20 string = "a"
                                    _ = v20
                                case "foo bar":
                                    v13 = fi17
                                case "a":
                                    v21 := GenS1{F0: v14}
                                    _ = v21
                                default:
                                    v19 -= 2
                                }
                                _ = v19
                            case "foo bar":
                                v18 = 18446744073709551615
                            }
                            var v22 uint32 = 233
                            _ = v18
                            _ = v22
                        } else {
                            if true {
                                fmt.Println(v14)
                                switch (fi16 & (fi16 - 42)) {
                                case 127:
                                    v23 := GenS2{F0: v11}
                                    _ = v23
                                default:
                                    if false {
                                        var v24 uint16 = 820
                                        if ((v14 | v14) < (v14 & (v14 % (((v14 >> ((1) & 63))) | 1)))) {
                                            fmt.Println(v13)
                                            v25 := GenS2{F0: 127}
                                            var v26 int = (-(fi17))
                                            _ = v25
                                            _ = v26
                                        } else {
                                            var v27 int32 = 0
                                            var v28 string = "\n\t"
                                            switch fi16 {
                                            case 9223372036854775807:
                                                var v29 int8 = v14
                                                var v30 uint8 = 0
                                                v31 := GenS3{F0: v27}
                                                _ = v29
                                                _ = v30
                                                _ = v31
                                            case 127:
                                                var v32 int8 = v14
                                                _ = v32
                                            case -924:
                                                delete(v3, 4294967295)
                                            }
                                            _ = v27
                                            _ = v28
                                        }
                                        var v33 uint64 = uint64(v11)
                                        _ = v24
                                        _ = v33
                                    }
                                }
                            } else {
                                v2 = -1
                                if (272 == 2147483648) {
                                    v34 := v3[4294967295]
                                    v34 = (v34 / ((v34) | 1))
                                    var v35 int32 = -650
                                    _ = v34
                                    _ = v35
                                } else {
                                    switch fi16 {
                                    case -9223372036854775808:
                                        v36 := GenS1{F0: 0}
                                        _ = v36
                                    }
                                    var v37 uint64 = (v4 % ((v4) | 1))
                                    v13 = fi17
                                    _ = v37
                                }
                                v38 := GenS2{F0: (^(v11))}
                                _ = v38
                            }
                            var v39 uint16 = 1
                            if false {
                                switch fi17 {
                                case 314:
                                    var v40 string = "a"
                                    _ = v40
                                case 42:
                                    v11 += 2
                                case 2147483647:
                                    v2 -= 32767
                                default:
                                    v13 -= -2147483648
                                }
                                var v41 uint64 = v4
                                _ = v41
                            }
                            _ = v39
                        }
                        delete(v3, 294)
                        var v42 uint = 2
                        _ = fi17
                        _ = v42
                    }
                    _ = fi16
                }
            }
            var v43 int64 = 0
            _ = v43
        } else {
            var v44 uint32 = 2147483647
            if ((v11 ^ (v11 % (((v11 << ((18446744073709551615) & 63))) | 1))) < 249) {
                fmt.Println(v44, v11, v13)
                var v45 int8 = 127
                if (true && ((!(0 <= -841)) || true)) {
                    v46 := GenS4{F0: v14, F1: -929, F2: v44}
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(v2)
                    }
                    _ = v46
                }
                _ = v45
            } else {
                for fi47 := 0; fi47 < 16; fi47++ {
                    v14 = v14
                    if (546 <= 963) {
                        v13 = fi47
                        v48 := GenS5{F0: v14, F1: -737, F2: (v2 | v2)}
                        _ = v48
                    } else {
                        var v49 int64 = 436
                        if true {
                            v49 = (v49 | v49)
                            for fi50 := 0; fi50 < 8; fi50++ {
                                v13 = fi50
                                v51 := GenS3{F0: -2147483648}
                                var v52 int8 = (v14 * v14)
                                _ = fi50
                                _ = v51
                                _ = v52
                            }
                        }
                        for fi53 := 0; fi53 < 13; fi53++ {
                            var v54 uint64 = (v4 % (((v4 << ((202) & 63))) | 1))
                            if (176 <= 32767) {
                                var v55 float32 = float32(1.5)
                                var v56 int8 = -128
                                _ = v55
                                _ = v56
                            }
                            v14 += (^(v14))
                            _ = fi53
                            _ = v54
                        }
                        _ = v49
                    }
                    v2 -= int16(v44)
                    _ = fi47
                }
                switch "foo bar" {
                case "foo bar":
                    if (!(3.14159 <= (-2.25 * 0.0))) {
                        var v57 int32 = 610
                        for fi58 := 0; fi58 < 9; fi58++ {
                            v59 := GenS6{F0: (v11 % (((^(v11))) | 1)), F1: false}
                            _ = fi58
                            _ = v59
                        }
                        var v60 int8 = (v14 ^ -1)
                        _ = v57
                        _ = v60
                    }
                case "":
                    var v61 uint8 = 159
                    _ = v61
                case "hello":
                    for fi62 := 0; fi62 < 5; fi62++ {
                        var v63 bool = ("hello" >= (((("hello" + "") + "a") + "a") + "hello"))
                        var v64 uint16 = 0
                        _ = fi62
                        _ = v63
                        _ = v64
                    }
                }
            }
            _ = v44
        }
    default:
        delete(v3, 42)
    }
    v65 := GenS3{F0: 0}
    v66 := GenS5{F0: (v14 << ((2) & 63)), F1: (v13 + v13), F2: v2}
    if ((false || false) && true) {
        v13 += (v13 / ((v13) | 1))
    }
    var v67 float32 = float32((1.5 + 3.14159))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v65
    _ = v66
    _ = v67
}
