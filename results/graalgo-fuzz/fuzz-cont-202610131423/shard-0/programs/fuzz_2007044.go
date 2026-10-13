package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int16 }
func genFn0() uint32 {
    v5 := []int{32767}
    if false {
        var v6 uint8 = 0
        _ = v6
    }
    _ = v5
    return 1
}
func genFn1(p0 uint16) int {
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := []int{fi7, (fi7 << ((4294967295) & 63)), fi7, (fi7 & fi7)}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        switch fi9 {
        case -1:
            switch 46 {
            case 32768:
                p0 = p0
            case -9223372036854775808:
                fmt.Println(fi9, p0)
            case 390:
                p0 = p0
            default:
                fmt.Println(p0, fi9)
            }
        }
        for fi10 := 0; fi10 < 14; fi10++ {
            v11 := GenS1{F0: 589}
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    fmt.Println(p0)
    for fi12 := 0; fi12 < 2; fi12++ {
        if ((fi12 * (fi12 >> ((4294967295) & 63))) == (^(fi12))) {
            p0 += (p0 + (p0 * (p0 >> ((0) & 63))))
        } else {
            v13 := []int{127, (^(fi12)), fi12, 9223372036854775807}
            for fi14 := 0; fi14 < 19; fi14++ {
                p0 -= (p0 - p0)
                _ = fi14
            }
            var v15 string = ((("hello" + ("" + ((("foo bar" + "") + "\n\t") + "a"))) + "\n\t") + (((("" + "foo bar") + (("foo bar" + "\n\t") + "a")) + "\n\t") + ""))
            _ = v13
            _ = v15
        }
        if true {
            v16 := []int{356, fi12}
            if len(v16) > 0 {
                v16[0] = fi12
            }
            fmt.Println(fi12, p0)
            _ = v16
        }
        p0 += (p0 ^ p0)
        _ = fi12
    }
    p0 += 1
    _ = p0
    return 9223372036854775807
}

func main() {
    v1 := []uint32{1, 31, 785, 63}
    v2 := make(map[int]bool)
    var v3 int = 0
    v4 := []int{}
    v17 := GenS2{F0: 1000, F1: 42}
    var v18 int32 = -2147483648
    v3 = -217
    var v19 float32 = 1e10
    if (3.14159 == (((1e10 / 1.0) / (0.1 / 1.5)) + (v19 + (v19 + 0.1)))) {
        for fi20 := 0; fi20 < 6; fi20++ {
            var v21 int8 = int8(fi20)
            _ = fi20
            _ = v21
        }
        fmt.Println(v19, v3, v19)
    } else {
        var v22 string = "foo bar"
        _ = v22
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    delete(v2, v3)
    for fi23 := 0; fi23 < 12; fi23++ {
        switch ("hello" + "a") {
        case "\n\t":
            if (((484 >= 9223372036854775807) && false) && (!(v18 >= 752))) {
                v24 := GenS3{F0: -1}
                for fi25 := 0; fi25 < 20; fi25++ {
                    v18 += v18
                    if (128 == 1) {
                        v26 := GenS2{F0: -32768, F1: v18}
                        var v27 int8 = -6
                        _ = v26
                        _ = v27
                    }
                    _ = fi25
                }
                _ = v24
            } else {
                for fi28 := 0; fi28 < 9; fi28++ {
                    v29 := GenS2{F0: -123, F1: (^(v18))}
                    v2[v3] = (183 <= 65535)
                    v18 = 0
                    _ = fi28
                    _ = v29
                }
            }
        case "foo bar":
            if false {
                fmt.Println(v19, v3, v19)
                var v30 float64 = 3.14159
                _ = v30
            }
        }
        v31 := GenS1{F0: 42}
        fmt.Println(v18, fi23, v3)
        _ = fi23
        _ = v31
    }
    var v32 uint8 = 102
    v33 := GenS2{F0: 0, F1: v18}
    v4 = append(v4, (v3 & v3))
    switch "a" {
    case "\n\t":
        v2[(v3 & -1)] = (false || ((v3 + (-(v3))) < v3))
    default:
        v18 = v18
    }
    if (!false) {
        for fi34 := 0; fi34 < 13; fi34++ {
            var v35 float64 = 3.14159
            var v36 int64 = 970
            var v37 uint8 = (^(v32))
            _ = fi34
            _ = v35
            _ = v36
            _ = v37
        }
        v38 := GenS3{F0: -32768}
        _ = v38
    } else {
        v39 := GenS3{F0: 32767}
        for fi40 := 0; fi40 < 10; fi40++ {
            for fi41 := 0; fi41 < 17; fi41++ {
                v19 = -2.25
                if true {
                    v42 := v2[fi41]
                    if (!true) {
                        fmt.Println(v42, v19, fi41)
                    } else {
                        v42 = ((18446744073709551615 >= 0) || v42)
                        v2[fi41] = (!v42)
                        if true {
                            _ = len(v2)
                        } else {
                            v2[v3] = v42
                            if (0 > -395) {
                                _ = len(v2)
                                fmt.Println(fi40, v19)
                                v42 = false
                            } else {
                                v42 = false
                                if v42 {
                                    v42 = true
                                }
                                v19 = (3.14159 - (1.5 + v19))
                                if (0 == 1000) {
                                    v42 = v42
                                }
                            }
                        }
                    }
                    v2[(v3 >> ((885) & 63))] = (float64(3.14159) <= 0.1)
                    _ = v42
                }
                var v43 float32 = v19
                _ = fi41
                _ = v43
            }
            fmt.Println(v19)
            if false {
                v19 -= float32(-2.25)
            } else {
                switch (("\n\t" + "a") + "foo bar") {
                case "a":
                    for fi44 := 0; fi44 < 7; fi44++ {
                        if (true && true) {
                            fmt.Println(v19)
                        }
                        fmt.Println(fi44, fi40, v32)
                        v19 = 0.1
                        _ = fi44
                    }
                case "hello":
                    if false {
                        v45 := GenS3{F0: 32767}
                        _ = v45
                    } else {
                        var v46 string = ""
                        fmt.Println(v3, v3, v18)
                        for fi47 := 0; fi47 < 8; fi47++ {
                            if false {
                                v32 = v32
                            }
                            v3 += 1
                            var v48 float64 = 1.5
                            _ = fi47
                            _ = v48
                        }
                        _ = v46
                    }
                case "foo bar":
                    var v49 int = genFn1(65535)
                    _ = v49
                default:
                    for fi50 := 0; fi50 < 3; fi50++ {
                        if false {
                            v3 = (fi50 & 0)
                            v18 = v18
                        } else {
                            fmt.Println(v19, v19)
                            var v51 uint16 = 42
                            _ = len(v2)
                            _ = v51
                        }
                        v18 += (v18 << ((42) & 63))
                        _ = fi50
                    }
                }
                var v52 uint64 = uint64(v32)
                delete(v2, 42)
                _ = v52
            }
            _ = fi40
        }
        switch "\n\t" {
        case "a":
            if ((v32 & 255) > v32) {
                for fi53 := 0; fi53 < 18; fi53++ {
                    for fi54 := 0; fi54 < 19; fi54++ {
                        v32 = v32
                        if false {
                            v18 = v18
                            v19 = float32(((0.1 * -2.25) / 0.1))
                            switch "\n\t" {
                            case "\n\t":
                                v55 := GenS1{F0: 2}
                                _ = v55
                            case "foo bar":
                                v32 = 203
                            default:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        }
                        _ = fi54
                    }
                    v56 := GenS1{F0: 867}
                    _ = fi53
                    _ = v56
                }
                var v57 bool = true
                _ = v57
            } else {
                v58 := GenS1{F0: genFn0()}
                if (((("" + ("" + "a")) + ("\n\t" + "a")) + "\n\t") < ("foo bar" + "\n\t")) {
                    var v59 float32 = ((v19 * float32((1.5 / 1.0))) + -2.25)
                    if (1 <= -128) {
                        for fi60 := 0; fi60 < 20; fi60++ {
                            v19 += v59
                            var v61 int = fi60
                            _ = fi60
                            _ = v61
                        }
                    }
                    _ = v59
                }
                if false {
                    v62 := GenS2{F0: -32768, F1: (v18 ^ v18)}
                    _ = len(v4)
                    var v63 int32 = (v18 ^ v18)
                    _ = v62
                    _ = v63
                } else {
                    v64 := v2[v3]
                    _ = v64
                }
                _ = v58
            }
        case "foo bar":
            switch "" {
            case "foo bar":
                for fi65 := 0; fi65 < 16; fi65++ {
                    fmt.Println(fi65, v19, v3)
                    if true {
                        v32 += 65
                        if true {
                            var v66 uint16 = uint16(v32)
                            _ = v66
                        }
                    }
                    _ = fi65
                }
            default:
                v67 := GenS2{F0: 1000, F1: (v18 * v18)}
                _ = v67
            }
        default:
            v68 := GenS2{F0: 32767, F1: (-(v18))}
            _ = v68
        }
        _ = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
