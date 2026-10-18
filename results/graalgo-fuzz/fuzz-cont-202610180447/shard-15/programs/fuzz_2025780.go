package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int; F1 int8; F2 int8 }
type GenS4 struct { F0 int16; F1 bool; F2 bool }
type GenS5 struct { F0 uint8; F1 int8; F2 string }
func genFn0(p0 int, p1 bool, p2 int64) uint16 {
    v4 := GenS1{F0: 513}
    v5 := []int{9223372036854775807, (p0 * p0)}
    _ = len(v5)
    for fi6 := 0; fi6 < 16; fi6++ {
        switch 32768 {
        case 0:
            var v7 float64 = -2.25
            _ = v7
        case 4294967295:
            for fi8 := 0; fi8 < 10; fi8++ {
                v9 := GenS1{F0: 1}
                _ = fi8
                _ = v9
            }
        case 589:
            p2 = (p2 % ((42) | 1))
        default:
            fmt.Println(p0, p1, p1)
        }
        _ = fi6
    }
    fmt.Println(p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 32767
}
func genFn1(p0 string, p1 uint16, p2 int64) int32 {
    var v10 int8 = 118
    v11 := GenS2{F0: -32768, F1: p0}
    var v12 uint16 = (p1 & 65535)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 423
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 42
    var v3 int8 = -21
    switch 127 {
    case 9223372036854775807:
        v2 += (v2 << ((733) & 63))
    case 42:
        var v13 bool = (!true)
        _ = v13
    }
    v14 := GenS1{F0: -1}
    v15 := []int{}
    if false {
        var v16 int = -9223372036854775808
        _ = v16
    }
    fmt.Println(v3, v2)
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 int8 = v3
        if true {
            v19 := GenS2{F0: 1, F1: (("foo bar" + "\n\t") + ("hello" + "foo bar"))}
            _ = v19
        } else {
            v3 -= v18
            v18 += (^(v18))
            if ((((3.14159 + 1e10) - -2.25) > (3.14159 * (1e10 - 0.1))) || (-32768 < 629)) {
                for fi20 := 0; fi20 < 6; fi20++ {
                    switch "a" {
                    case "a":
                        v21 := GenS3{F0: fi20, F1: (v3 - 55), F2: (v3 + (^(v3)))}
                        _ = v21
                    }
                    if ((32767 >= 1) && true) {
                        var v22 float32 = (1e10 - 1e10)
                        var v23 bool = false
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        v24 := GenS3{F0: 1, F1: v18, F2: (v3 * 1)}
                        switch "foo bar" {
                        case "":
                            if (!((^(fi17)) < fi20)) {
                                var v25 int64 = 974
                                switch (fi20 / ((-1) | 1)) {
                                case -435:
                                    _ = len(v15)
                                    fmt.Println(v22, v18)
                                    v23 = (100 >= (v3 % ((-1) | 1)))
                                case -953:
                                    var v26 uint8 = 127
                                    _ = v26
                                case -960:
                                    v27 := GenS2{F0: -32768, F1: "a"}
                                    _ = v27
                                default:
                                    v23 = (785 > genFn0(fi17, v23, v25))
                                }
                                if v23 {
                                    var v28 uint = 65535
                                    fmt.Println(fi17, v25)
                                    _ = v28
                                } else {
                                    var v29 uint16 = 344
                                    delete(v1, ("a" + "a"))
                                    v3 = -128
                                    _ = v29
                                }
                                _ = v25
                            } else {
                                v30 := GenS4{F0: 570, F1: (((false || (1 != 26)) || v23) && true), F2: (!(695 < 629))}
                                v31 := GenS5{F0: 42, F1: (v18 - (v18 << ((65535) & 63))), F2: ((("" + "") + "") + (("a" + "a") + ""))}
                                _ = v30
                                _ = v31
                            }
                        case "\n\t":
                            var v32 uint8 = 80
                            _ = v32
                        }
                        _ = v22
                        _ = v23
                        _ = v24
                    } else {
                        var v33 uint8 = 128
                        switch "\n\t" {
                        case "\n\t":
                            if (!(true || true)) {
                                v2 = (v2 % ((18446744073709551615) | 1))
                                v34 := GenS2{F0: -32768, F1: "\n\t"}
                                v35 := v1[(((("hello" + "hello") + ("a" + "foo bar")) + "\n\t") + ("foo bar" + "foo bar"))]
                                v35 = (^(v35))
                                _ = v34
                                _ = v35
                            } else {
                                _ = len(v15)
                                if ((v33 - v33) <= v33) {
                                    v18 -= -120
                                } else {
                                    fmt.Println(fi20, fi17, v18)
                                    var v36 uint16 = uint16(v2)
                                    _ = v36
                                }
                            }
                            if false {
                                delete(v1, "hello")
                                v37 := GenS1{F0: 749}
                                _ = v37
                            }
                        case "":
                            var v38 uint64 = (v2 >> ((0) & 63))
                            _ = v38
                        }
                        v2 -= v2
                        v3 -= (v3 ^ -1)
                        _ = v33
                    }
                    _ = fi20
                }
                delete(v1, "a")
                fmt.Println(v18, fi17, v2)
            }
        }
        var v39 bool = (int8(v2) <= -1)
        _ = fi17
        _ = v18
        _ = v39
    }
    fmt.Println(v2)
    if len(v15) > 0 {
        v15[0] = 332
    }
    var v40 uint32 = 1
    _ = len(v15)
    for fi41 := 0; fi41 < 1; fi41++ {
        var v42 uint32 = 122
        v43 := GenS4{F0: 0, F1: false, F2: (0 <= 133)}
        _ = fi41
        _ = v42
        _ = v43
    }
    v2 += 595
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v40
}
