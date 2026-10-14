package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 int; F2 int32 }
type GenS4 struct { F0 uint16; F1 uint; F2 int32 }
type GenS5 struct { F0 bool; F1 int32 }
func genFn0(p0 int32, p1 int64) int16 {
    fmt.Println(p0, p1)
    var v4 float64 = (-2.25 / 1.0)
    if (("\n\t" + "foo bar") >= "foo bar") {
        if true {
            p1 = (p1 * (p1 * p1))
            var v5 bool = false
            v6 := []int{}
            _ = v5
            _ = v6
        } else {
            p1 = p1
            p0 = 0
        }
    } else {
        var v7 int32 = (p0 << ((1) & 63))
        _ = v7
    }
    v8 := []int{128, 7, 32768}
    _ = p0
    _ = p1
    _ = v4
    _ = v8
    return 883
}
func genFn1(p0 int64) uint {
    if false {
        var v9 uint32 = 0
        var v10 uint16 = 65535
        _ = v10
        _ = v9
    }
    switch 663 {
    case -9223372036854775808:
        var v11 int8 = -128
        _ = v11
    }
    _ = p0
    return 104
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 2
    var v3 float32 = (float32((0.1 + 1e10)) / ((float32(-2.25) + ((1.5 / 1.0) + (3.14159 + -2.25))) + 1.5))
    if true {
        v12 := GenS1{F0: -1}
        if (84 >= -128) {
            if true {
                var v13 bool = (18446744073709551615 <= 18446744073709551615)
                if v13 {
                    for fi14 := 0; fi14 < 18; fi14++ {
                        v2 = (v2 / ((2147483647) | 1))
                        if false {
                            v15 := GenS2{F0: v2}
                            v16 := []int{(fi14 * -9223372036854775808), 127}
                            var v17 uint32 = v2
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        _ = fi14
                    }
                }
                delete(v1, "")
                _ = v13
            }
        }
        delete(v1, "hello")
        _ = v12
    } else {
        var v18 int = -9223372036854775808
        _ = v18
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 int16 = 42
        if true {
            fmt.Println(v3, v20, v3)
        }
        _ = fi19
        _ = v20
    }
    var v21 uint32 = (v2 & (v2 ^ (^(v2))))
    fmt.Println(v3, v2, v21)
    if (!(!false)) {
        v3 -= float32(1.5)
        v22 := GenS2{F0: (v21 - (v21 >> ((18446744073709551615) & 63)))}
        _ = v22
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        var v24 int32 = 2147483647
        _ = fi23
        _ = v24
    }
    v2 = (v2 | (^(v21)))
    v3 -= v3
    v25 := GenS2{F0: (v21 + (v2 - (v21 ^ (v21 / ((v21) | 1)))))}
    for fi26 := 0; fi26 < 3; fi26++ {
        if (2 >= 1) {
            if (2 <= genFn1(329)) {
                for fi27 := 0; fi27 < 13; fi27++ {
                    var v28 int8 = 42
                    switch fi27 {
                    case 9223372036854775807:
                        fmt.Println(v3)
                    case 42:
                        var v29 int16 = 32767
                        _ = v29
                    default:
                        fmt.Println(fi27, v21, fi27)
                    }
                    switch (fi27 * fi26) {
                    case 977:
                        v30 := GenS1{F0: 42}
                        var v31 string = "hello"
                        fmt.Println(fi27, v31)
                        _ = v30
                        _ = v31
                    }
                    _ = fi27
                    _ = v28
                }
                if true {
                    v32 := GenS3{F0: 2147483647, F1: fi26, F2: -903}
                    for fi33 := 0; fi33 < 19; fi33++ {
                        fmt.Println(fi33)
                        v21 = (^(v2))
                        _ = fi33
                    }
                    _ = v32
                } else {
                    switch -9223372036854775808 {
                    case 9223372036854775807:
                        v34 := []int{-1, fi26, 2147483647}
                        _ = v34
                    case 2147483647:
                        for fi35 := 0; fi35 < 12; fi35++ {
                            var v36 bool = (false || ((genFn0(2147483647, 9223372036854775807) <= 0) && true))
                            var v37 float64 = (1.5 * 0.1)
                            switch fi35 {
                            case 2147483647:
                                if (-640 <= -1) {
                                    var v38 string = ("hello" + "")
                                    var v39 int64 = 771
                                    v1[v38] = (v38 + v38)
                                    _ = v38
                                    _ = v39
                                } else {
                                    var v40 int8 = 103
                                    var v41 uint64 = 42
                                    _ = v40
                                    _ = v41
                                }
                            case -2147483648:
                                if false {
                                    v42 := v1[("hello" + "")]
                                    v3 = float32(1e10)
                                    switch (fi26 % (((fi35 + (fi35 << ((genFn1(1)) & 63)))) | 1)) {
                                    case 0:
                                        v43 := GenS1{F0: 32767}
                                        _ = v43
                                    }
                                    _ = v42
                                } else {
                                    fmt.Println(v21)
                                }
                            }
                            _ = fi35
                            _ = v36
                            _ = v37
                        }
                    case 911:
                        v44 := GenS3{F0: 2147483647, F1: fi26, F2: 42}
                        _ = v44
                    }
                }
                var v45 int64 = -524
                _ = v45
            }
        } else {
            _ = len(v1)
            for fi46 := 0; fi46 < 2; fi46++ {
                if true {
                    v47 := []int{(fi26 & (fi46 / ((0) | 1))), 1, fi46, -9223372036854775808}
                    v21 = (^(v2))
                    _ = v47
                }
                v2 -= 0
                _ = fi46
            }
        }
        v48 := []int{128, (fi26 / ((fi26) | 1)), fi26, fi26}
        _ = fi26
        _ = v48
    }
    for fi49 := 0; fi49 < 16; fi49++ {
        if (("foo bar" < "") && false) {
            fmt.Println(v21, fi49)
            v1[((("\n\t" + "\n\t") + ("\n\t" + "hello")) + "\n\t")] = (("\n\t" + ("foo bar" + (("a" + ("" + "")) + "hello"))) + "")
        } else {
            for fi50 := 0; fi50 < 20; fi50++ {
                v51 := GenS4{F0: 87, F1: 65535, F2: 1}
                if false {
                    v52 := GenS2{F0: 829}
                    if (162 != 32767) {
                        var v53 uint16 = 42
                        var v54 string = ("" + ("\n\t" + "a"))
                        switch fi50 {
                        case 128:
                            var v55 int = fi49
                            _ = v55
                        case 323:
                            var v56 uint8 = 1
                            _ = v56
                        default:
                            var v57 int8 = 100
                            _ = v57
                        }
                        v54 = (v54 + v54)
                        _ = v53
                        _ = v54
                    } else {
                        var v58 int16 = 675
                        var v59 uint32 = 932
                        _ = v58
                        _ = v59
                    }
                    _ = v52
                } else {
                    switch "" {
                    case "a":
                        v60 := GenS4{F0: 5, F1: 1, F2: 42}
                        _ = v60
                    default:
                        var v61 uint64 = 1
                        _ = v61
                    }
                    if true {
                        fmt.Println(v2)
                    }
                }
                _ = fi50
                _ = v51
            }
        }
        var v62 int16 = -65
        if true {
            var v63 int8 = -69
            fmt.Println(v3)
            _ = v63
        }
        _ = fi49
        _ = v62
    }
    v2 -= (v21 * (^(v2)))
    v64 := GenS2{F0: (v2 | (v2 / ((915) | 1)))}
    fmt.Println(v21, v21, v2)
    if ((1e10 < (v3 / float32(((-2.25 / 1.5) + 0.1)))) || true) {
        v65 := GenS5{F0: (42 < -867), F1: 42}
        var v66 bool = (953 != -527)
        _ = v65
        _ = v66
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v64
}
