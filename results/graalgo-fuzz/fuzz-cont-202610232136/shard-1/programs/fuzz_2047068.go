package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS6 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int64; F1 int16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int64; F1 uint }
func genFn0(p0 uint32, p1 uint8, p2 int16) uint16 {
    var v7 int64 = -1
    v8 := GenS1{F0: 0}
    p1 -= (p1 & (p1 + (p1 >> ((4294967295) & 63))))
    var v9 string = "\n\t"
    v10 := GenS2{F0: (p0 * 0), F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 int8, p1 string) int64 {
    switch 127 {
    case -874:
        var v11 int8 = (p0 & (^(p0)))
        _ = v11
    case 1:
        p0 -= (p0 * p0)
    }
    switch p1 {
    case "foo bar":
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    return 9223372036854775807
}

func main() {
    v1 := []string{}
    var v2 int = 0
    var v3 int8 = -128
    v4 := make(map[uint32]uint32)
    var v5 bool = ((1 <= 347) || true)
    var v6 uint16 = 860
    var v12 string = "foo bar"
    if true {
        var v13 uint8 = 2
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        v15 := GenS3{F0: 4294967295}
        if false {
            switch (v2 << ((4294967295) & 63)) {
            case -1:
                fmt.Println(v3)
            case 1:
                if (!v5) {
                    switch v12 {
                    case "foo bar":
                        v3 -= v3
                    case "":
                        if true {
                            for fi16 := 0; fi16 < 6; fi16++ {
                                var v17 float32 = ((3.14159 / -2.25) + 3.14159)
                                _ = fi16
                                _ = v17
                            }
                            v6 = v6
                            v18 := GenS4{F0: (v6 >> ((65535) & 63))}
                            _ = v18
                        }
                    default:
                        v1 = append(v1, "")
                    }
                } else {
                    v5 = true
                }
            case 32767:
                v19 := GenS4{F0: (v6 * (v6 / (((^(v6))) | 1)))}
                _ = v19
            default:
                v20 := GenS5{F0: v6, F1: 255, F2: 581}
                _ = v20
            }
            if (!(0 == 18446744073709551615)) {
                delete(v4, 497)
            } else {
                var v21 int64 = 42
                if true {
                    v22 := GenS6{F0: 1, F1: 118, F2: 659}
                    _ = v22
                }
                if v5 {
                    var v23 float32 = ((0.0 + float32(3.14159)) + -2.25)
                    if (!v5) {
                        v6 += v6
                        v24 := GenS5{F0: v6, F1: 255, F2: 797}
                        _ = v24
                    }
                    _ = v23
                }
                _ = v21
            }
            for fi25 := 0; fi25 < 13; fi25++ {
                fmt.Println(v2, fi14, fi14)
                _ = fi25
            }
        }
        _ = fi14
        _ = v15
    }
    if (!(v2 >= (v2 ^ v2))) {
        for fi26 := 0; fi26 < 10; fi26++ {
            v12 = v12
            _ = fi26
        }
    } else {
        for fi27 := 0; fi27 < 1; fi27++ {
            var v28 int = (-(fi27))
            if true {
                for fi29 := 0; fi29 < 15; fi29++ {
                    v30 := GenS2{F0: 458, F1: (v28 & fi29)}
                    v6 += v6
                    if (!v5) {
                        var v31 float64 = 3.14159
                        if ((v3 >> ((562) & 63)) <= (v3 << ((4294967295) & 63))) {
                            var v32 uint = 4294967295
                            _ = v32
                        } else {
                            if true {
                                v33 := GenS7{F0: (v2 / ((fi29) | 1))}
                                _ = v33
                            }
                            var v34 uint32 = 4294967295
                            v35 := GenS2{F0: (v34 << ((255) & 63)), F1: int(v6)}
                            fmt.Println(fi27, v5)
                            _ = v34
                            _ = v35
                        }
                        v28 += v28
                        _ = v31
                    } else {
                        if (!(102 > (v3 ^ (v3 * (v3 | v3))))) {
                            var v36 uint64 = 2
                            if (true || v5) {
                                v37 := GenS2{F0: 2147483647, F1: -10}
                                fmt.Println(v12, v2)
                                _ = v37
                            } else {
                                if false {
                                    if len(v1) > 0 {
                                        v1[0] = ("" + v12)
                                    }
                                    if v5 {
                                        if false {
                                            v38 := GenS1{F0: 1}
                                            fmt.Println(fi29, fi27, fi27)
                                            fmt.Println(v36)
                                            _ = v38
                                        } else {
                                            v36 = 784
                                        }
                                        var v39 uint64 = v36
                                        if false {
                                            var v40 int16 = 0
                                            if false {
                                                fmt.Println(v36, v40)
                                                switch 32768 {
                                                case -2147483648:
                                                    v41 := GenS8{F0: 9223372036854775807, F1: (v40 - (v40 % (((-(v40))) | 1)))}
                                                    _ = v41
                                                case 32768:
                                                    v42 := GenS2{F0: 219, F1: 42}
                                                    _ = v42
                                                }
                                            }
                                            v5 = (((1.5 * ((-2.25 * 1e10) - 0.1)) * 3.14159) < (1.5 * (3.14159 + 1.5)))
                                            _ = v40
                                        }
                                        _ = v39
                                    } else {
                                        switch (v28 + 0) {
                                        case -2147483648:
                                            var v43 string = v12
                                            _ = v43
                                        case -589:
                                            switch "a" {
                                            case "foo bar":
                                                var v44 float64 = float64((((-2.25 / 0.1) / (float32(float64(1e10)) - 0.0)) - 1.5))
                                                _ = v44
                                            }
                                        default:
                                            v2 -= v2
                                        }
                                        v2 += (-(v28))
                                        if v5 {
                                            var v45 int8 = (v3 ^ -1)
                                            _ = len(v4)
                                            _ = v45
                                        }
                                    }
                                    v46 := GenS9{F0: v6}
                                    _ = v46
                                }
                                v12 = v12
                                v28 = fi27
                            }
                            if (3.14159 == ((1e10 * float64(1.5)) * 3.14159)) {
                                fmt.Println(v28, v28, fi29)
                                var v47 int64 = int64(v3)
                                _ = v47
                            }
                            _ = v36
                        }
                        v48 := v4[0]
                        v49 := GenS1{F0: uint(v6)}
                        _ = v48
                        _ = v49
                    }
                    _ = fi29
                    _ = v30
                }
                var v50 float32 = 0.0
                _ = v50
            } else {
                v2 += -1
            }
            v5 = (!v5)
            _ = fi27
            _ = v28
        }
        for fi51 := 0; fi51 < 7; fi51++ {
            fmt.Println(v6, v5)
            fmt.Println(fi51)
            _ = fi51
        }
    }
    var v52 uint32 = 0
    v4[v52] = (^(v52))
    switch v12 {
    case "a":
        if true {
            fmt.Println(v2, v2)
        }
    }
    for fi53 := 0; fi53 < 12; fi53++ {
        switch ("a" + ("" + "")) {
        case "":
            v5 = v5
        case "hello":
            for fi54 := 0; fi54 < 1; fi54++ {
                fmt.Println(v5, v2, fi53)
                var v55 uint32 = v52
                _ = fi54
                _ = v55
            }
        case "a":
            v5 = true
        default:
            switch (^(v2)) {
            case 127:
                v56 := GenS10{F0: genFn1(v3, ""), F1: 56}
                _ = v56
            default:
                var v57 int64 = 937
                _ = v57
            }
        }
        v58 := GenS5{F0: 1, F1: 18446744073709551615, F2: 0}
        _ = fi53
        _ = v58
    }
    v59 := v4[v52]
    var v60 uint8 = 38
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v59
    _ = v6
    _ = v60
}
