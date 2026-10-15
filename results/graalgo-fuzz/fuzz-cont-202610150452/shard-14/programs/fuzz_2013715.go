package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint16) uint8 {
    v4 := GenS1{F0: (^(p0))}
    p0 = p0
    v5 := []int{-9223372036854775808, 127}
    _ = p0
    _ = v4
    _ = v5
    return 2
}
func genFn1(p0 uint8) int32 {
    v6 := []int{32767, -2147483648, 285}
    p0 = (^(p0))
    p0 = genFn0(255)
    _ = p0
    _ = v6
    return 1
}

func main() {
    var v1 float64 = float64((0.0 / float32((float64((((1e10 - (1.5 / float32((1e10 * (-2.25 / 1.0))))) * (0.0 / (1.5 - -2.25))) + 0.1)) + (0.1 * ((3.14159 * 0.1) / (1.5 / 1e10)))))))
    var v2 int32 = -2147483648
    v3 := []string{"hello", "\n\t", "\n\t"}
    v7 := GenS1{F0: 32767}
    var v8 int64 = 1
    for fi9 := 0; fi9 < 6; fi9++ {
        if false {
            v10 := GenS1{F0: 0}
            _ = v10
        }
        if false {
            if len(v3) > 0 {
                v3[0] = ("foo bar" + "hello")
            }
            if (42 == 0) {
                var v11 int16 = 1000
                v12 := GenS1{F0: 32767}
                _ = v11
                _ = v12
            }
        } else {
            if true {
                switch fi9 {
                case 4294967295:
                    v1 = 3.14159
                default:
                    switch ("foo bar" + ("foo bar" + ("foo bar" + "\n\t"))) {
                    case "hello":
                        for fi13 := 0; fi13 < 15; fi13++ {
                            var v14 uint = 509
                            _ = fi13
                            _ = v14
                        }
                    case "a":
                        switch fi9 {
                        case -1:
                            v15 := GenS2{F0: 1}
                            _ = v15
                        case 32767:
                            v2 = v2
                        default:
                            v8 += v8
                        }
                    }
                }
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi16 := 0; fi16 < 20; fi16++ {
                if (32767 >= 218) {
                    v2 += v2
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                v2 = v2
                switch (("\n\t" + (((("" + (("hello" + "foo bar") + ("foo bar" + "\n\t"))) + "foo bar") + "") + ("a" + ((("a" + ("a" + "")) + ((("a" + ("\n\t" + (("\n\t" + "foo bar") + ("hello" + ("\n\t" + ("" + "")))))) + ("foo bar" + ("\n\t" + ""))) + "a")) + "")))) + ("hello" + ("foo bar" + "hello"))) {
                case "\n\t":
                    if (false && false) {
                        var v17 float32 = 3.14159
                        if (1 >= genFn0(640)) {
                            switch "\n\t" {
                            case "\n\t":
                                v18 := GenS1{F0: 2}
                                _ = v18
                            default:
                                if (1 < 383) {
                                    switch 1 {
                                    case -9223372036854775808:
                                        v19 := GenS3{F0: 42}
                                        _ = v19
                                    case -748:
                                        var v20 bool = false
                                        _ = v20
                                    case 1:
                                        var v21 uint64 = 1
                                        _ = v21
                                    }
                                    var v22 float64 = v1
                                    _ = v22
                                } else {
                                    var v23 int32 = v2
                                    _ = v23
                                }
                            }
                            fmt.Println(v1, v1)
                            v24 := GenS2{F0: 2147483647}
                            _ = v24
                        }
                        v3 = append(v3, "hello")
                        _ = v17
                    } else {
                        var v25 string = "foo bar"
                        var v26 bool = false
                        v27 := GenS4{F0: 355, F1: 0}
                        switch (-(fi16)) {
                        case 9223372036854775807:
                            _ = len(v3)
                        case 127:
                            v3 = append(v3, v25)
                        default:
                            v2 -= v2
                        }
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                }
                var v28 uint32 = 784
                _ = fi16
                _ = v28
            }
        }
        _ = fi9
    }
    v1 = float64(float32(v1))
    for fi29 := 0; fi29 < 15; fi29++ {
        if (!((0 <= 980) || false)) {
            v30 := GenS4{F0: 0, F1: 100}
            var v31 string = "foo bar"
            _ = v30
            _ = v31
        } else {
            switch ((("foo bar" + "a") + ("hello" + "hello")) + "foo bar") {
            case "hello":
                for fi32 := 0; fi32 < 18; fi32++ {
                    var v33 int64 = 0
                    _ = fi32
                    _ = v33
                }
            default:
                var v34 int = fi29
                _ = v34
            }
            fmt.Println(v8)
            if true {
                var v35 float64 = v1
                v36 := GenS3{F0: 0}
                _ = v35
                _ = v36
            } else {
                var v37 int = fi29
                _ = v37
            }
        }
        _ = fi29
    }
    var v38 int32 = (^(v2))
    var v39 uint16 = 255
    for fi40 := 0; fi40 < 19; fi40++ {
        if true {
            var v41 uint8 = 0
            v2 = v38
            _ = v41
        }
        for fi42 := 0; fi42 < 3; fi42++ {
            var v43 int16 = 1
            var v44 uint8 = 255
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = fi40
    }
    if true {
        switch "\n\t" {
        case "\n\t":
            if true {
                v45 := GenS1{F0: (v39 / ((v39) | 1))}
                _ = len(v3)
                _ = v45
            } else {
                v8 += v8
            }
        case "a":
            var v46 uint64 = 18446744073709551615
            _ = v46
        case "":
            _ = len(v3)
        }
        for fi47 := 0; fi47 < 16; fi47++ {
            v1 = -2.25
            fmt.Println(v2, v38)
            _ = fi47
        }
    } else {
        v8 = (v8 + (v8 + (v8 % ((886) | 1))))
        v39 = v39
        if true {
            if true {
                switch -610 {
                case 42:
                    v48 := GenS2{F0: 2}
                    _ = v48
                case -345:
                    _ = len(v3)
                case 0:
                    if (!(!false)) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v49 uint16 = (v39 ^ v39)
                        switch "hello" {
                        case "foo bar":
                            v39 -= 1
                        case "a":
                            if false {
                                v39 = v39
                            } else {
                                v2 -= genFn1(108)
                                v50 := GenS2{F0: 2147483647}
                                _ = v50
                            }
                        case "hello":
                            var v51 uint32 = 2147483647
                            _ = v51
                        }
                        _ = v49
                    } else {
                        v52 := GenS3{F0: 488}
                        if (-2.25 == (0.0 - (1.5 + 1.5))) {
                            var v53 int8 = 8
                            _ = v53
                        } else {
                            v54 := GenS2{F0: 65535}
                            if (-64 >= 1) {
                                if len(v3) > 0 {
                                    v3[0] = ""
                                }
                                v55 := GenS5{F0: 2147483648}
                                var v56 int = 943
                                _ = v55
                                _ = v56
                            } else {
                                v57 := GenS4{F0: 266, F1: -102}
                                v2 += genFn1(2)
                                for fi58 := 0; fi58 < 5; fi58++ {
                                    v59 := GenS3{F0: v39}
                                    var v60 uint = 42
                                    if ((367 < -721) || false) {
                                        v1 = -2.25
                                    }
                                    _ = fi58
                                    _ = v59
                                    _ = v60
                                }
                                _ = v57
                            }
                            fmt.Println(v1, v39)
                            _ = v54
                        }
                        _ = v52
                    }
                }
                v38 = v2
            } else {
                fmt.Println(v38)
            }
        } else {
            fmt.Println(v1, v8, v2)
            fmt.Println(v2)
            switch "hello" {
            case "a":
                var v61 int = 4294967295
                _ = v61
            case "foo bar":
                var v62 string = "\n\t"
                _ = v62
            default:
                v39 -= (v39 % ((1) | 1))
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v7
    _ = v8
}
