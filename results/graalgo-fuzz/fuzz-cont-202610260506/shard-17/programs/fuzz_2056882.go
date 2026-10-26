package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint16 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int16 }
func genFn0() int16 {
    var v6 float64 = 1.5
    var v7 int32 = -569
    switch 128 {
    case 32767:
        fmt.Println(v6, v6)
    default:
        var v8 uint = 255
        _ = v8
    }
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 int32, p1 string, p2 int8) uint64 {
    for fi9 := 0; fi9 < 15; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 uint16 = 546
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    if ((9223372036854775807 <= -204) || true) {
        p0 -= p0
    }
    var v12 uint8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 2
}

func main() {
    var v1 string = (("a" + "\n\t") + "\n\t")
    var v2 int32 = 42
    var v3 int = 0
    var v4 bool = (v1 < "\n\t")
    var v5 int64 = -267
    var v13 int64 = v5
    v14 := GenS1{F0: (v2 * (^(v2))), F1: ("" + v1), F2: 13}
    v2 += v2
    if false {
        var v15 int8 = -43
        _ = v15
    } else {
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 bool = true
            _ = fi16
            _ = v17
        }
    }
    v18 := GenS2{F0: 739, F1: 0}
    if true {
        if (!(false || v4)) {
            v19 := GenS1{F0: 42, F1: v1, F2: 0}
            for fi20 := 0; fi20 < 10; fi20++ {
                var v21 uint8 = 1
                switch 328 {
                case -9223372036854775808:
                    v3 = (fi20 * -2147483648)
                default:
                    v1 = v1
                }
                _ = fi20
                _ = v21
            }
            _ = v19
        }
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        v23 := []int{v3, fi22}
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 17; fi24++ {
        for fi25 := 0; fi25 < 1; fi25++ {
            v13 += 9223372036854775807
            v5 = v13
            _ = fi25
        }
        v5 += (-(v5))
        _ = fi24
    }
    if false {
        var v26 uint8 = 42
        switch v3 {
        case 361:
            if v4 {
                var v27 float32 = ((0.0 / 1.5) * 1e10)
                v28 := GenS1{F0: (v2 & (v2 * v2)), F1: v1, F2: 625}
                fmt.Println(v3)
                _ = v27
                _ = v28
            } else {
                fmt.Println(v26, v3, v1)
                switch (v1 + "a") {
                case "":
                    if v4 {
                        if (v4 && false) {
                            var v29 uint32 = 2147483648
                            _ = v29
                        }
                        v30 := GenS3{F0: (v3 * v3), F1: v5}
                        _ = v30
                    } else {
                        v31 := []int{}
                        fmt.Println(v26, v3)
                        _ = v31
                    }
                default:
                    var v32 int64 = (v13 % (((v13 / (((v5 + v13)) | 1))) | 1))
                    _ = v32
                }
                if false {
                    switch v3 {
                    case 32768:
                        for fi33 := 0; fi33 < 12; fi33++ {
                            v1 = "foo bar"
                            v34 := []int{(-(v3)), 127, (fi33 / ((fi33) | 1)), (fi33 ^ int(v5))}
                            _ = fi33
                            _ = v34
                        }
                    case 128:
                        switch v1 {
                        case "a":
                            v35 := GenS4{F0: int16(v2)}
                            _ = v35
                        }
                    case 9223372036854775807:
                        v3 -= v3
                    }
                    switch (v3 / (((^(v3))) | 1)) {
                    case -1:
                        for fi36 := 0; fi36 < 3; fi36++ {
                            fmt.Println(fi36, v1)
                            v37 := GenS1{F0: (v2 + (v2 >> ((166) & 63))), F1: v1, F2: 255}
                            _ = fi36
                            _ = v37
                        }
                    }
                } else {
                    if ((v3 < (-(v3))) && (v1 > v1)) {
                        v38 := []int{-1, (v3 % ((v3) | 1)), -1}
                        if (131 <= -580) {
                            for fi39 := 0; fi39 < 16; fi39++ {
                                for fi40 := 0; fi40 < 12; fi40++ {
                                    switch 1 {
                                    case 42:
                                        var v41 float64 = (((0.0 + (0.0 / 1.0)) * 0.0) * float64(-2.25))
                                        _ = v41
                                    default:
                                        v42 := GenS2{F0: genFn1(-1, "foo bar", 42), F1: 65535}
                                        _ = v42
                                    }
                                    fmt.Println(v2, v26)
                                    v2 = 42
                                    _ = fi40
                                }
                                _ = fi39
                            }
                            v43 := GenS4{F0: -512}
                            _ = v43
                        }
                        _ = v38
                    } else {
                        var v44 int = 1
                        _ = v44
                    }
                }
            }
        default:
            v1 = "a"
        }
        var v45 uint64 = 42
        _ = v26
        _ = v45
    }
    var v46 float32 = ((1e10 * 1.5) / (-2.25 * 3.14159))
    v47 := GenS1{F0: (-(v2)), F1: (v1 + v1), F2: 2}
    v3 = (v3 / ((378) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
}
