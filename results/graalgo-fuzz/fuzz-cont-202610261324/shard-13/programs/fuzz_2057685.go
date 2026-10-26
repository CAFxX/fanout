package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 uint32; F1 uint16 }
func genFn0(p0 int32) uint64 {
    p0 -= p0
    var v6 bool = false
    for fi7 := 0; fi7 < 7; fi7++ {
        p0 = (^(p0))
        if true {
            if v6 {
                var v8 uint64 = 2
                _ = v8
            } else {
                fmt.Println(fi7)
            }
            v9 := GenS1{F0: int64(fi7), F1: 931}
            _ = v9
        }
        _ = fi7
    }
    p0 = (p0 % ((-1) | 1))
    _ = p0
    _ = v6
    return 2
}

func main() {
    var v1 string = "foo bar"
    var v2 int64 = 1
    var v3 float32 = ((0.0 + 0.1) - ((3.14159 * ((1.5 + 1e10) + (1e10 / 0.1))) - ((1e10 + (float32(float64(-2.25)) - -2.25)) / 1.5)))
    var v4 float64 = (float64((float32(0.0) * ((v3 / ((float32(float64((0.0 + v3))) * v3) / v3)) * (v3 * v3)))) - (3.14159 / 1e10))
    var v5 int64 = v2
    var v10 float32 = 3.14159
    v10 += v10
    v11 := GenS1{F0: (-(v5)), F1: 506}
    v12 := GenS2{F0: 18446744073709551615, F1: 200, F2: (v2 + (v5 | (v2 * (v5 % ((v5) | 1)))))}
    v13 := []int{90, 127, -356}
    for fi14 := 0; fi14 < 13; fi14++ {
        if ((fi14 >> ((4294967295) & 63)) != fi14) {
            fmt.Println(v3)
            for fi15 := 0; fi15 < 18; fi15++ {
                v10 += v3
                if len(v13) > 0 {
                    v13[0] = fi15
                }
                _ = fi15
            }
            var v16 int16 = -32768
            _ = v16
        } else {
            v1 = ""
            fmt.Println(v4, fi14, v2)
        }
        for fi17 := 0; fi17 < 6; fi17++ {
            if (!(false || false)) {
                if false {
                    v18 := GenS1{F0: -1, F1: 255}
                    _ = v18
                } else {
                    v3 -= (0.0 - v10)
                    _ = len(v13)
                    v4 = (1e10 / (1.5 + v4))
                    var v19 uint32 = 65535
                    if len(v13) > 0 {
                        v13[0] = (fi17 & 9223372036854775807)
                    }
                    _ = v19
                }
            }
            v20 := GenS1{F0: (v5 >> ((18446744073709551615) & 63)), F1: 1}
            fmt.Println(v5, v5, v5)
            _ = fi17
            _ = v20
        }
        v1 = (v1 + v1)
        _ = fi14
    }
    fmt.Println(v3, v1)
    for fi21 := 0; fi21 < 17; fi21++ {
        if (255 <= 0) {
            _ = len(v13)
            v22 := GenS2{F0: 0, F1: 440, F2: v2}
            _ = v22
        } else {
            var v23 int = fi21
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            _ = v23
        }
        if (-1 <= -32768) {
            if true {
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                v5 = v2
            } else {
                for fi24 := 0; fi24 < 2; fi24++ {
                    var v25 int16 = -804
                    var v26 int = -981
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
            _ = len(v13)
            for fi27 := 0; fi27 < 5; fi27++ {
                if true {
                    v10 -= ((v10 + -2.25) / v3)
                }
                v1 = ((v1 + (("" + v1) + "\n\t")) + "")
                var v28 int32 = -935
                _ = fi27
                _ = v28
            }
        }
        var v29 float32 = 1.5
        _ = fi21
        _ = v29
    }
    v30 := GenS3{F0: 225, F1: v1}
    var v31 uint = 42
    fmt.Println(v31)
    _ = len(v13)
    for fi32 := 0; fi32 < 2; fi32++ {
        fmt.Println(v10, v31)
        for fi33 := 0; fi33 < 6; fi33++ {
            var v34 int32 = -2147483648
            _ = fi33
            _ = v34
        }
        _ = fi32
    }
    var v35 uint64 = 0
    for fi36 := 0; fi36 < 10; fi36++ {
        fmt.Println(v5, v3)
        for fi37 := 0; fi37 < 1; fi37++ {
            switch -967 {
            case 2147483647:
                v38 := GenS2{F0: v31, F1: v35, F2: 42}
                _ = v38
            case 43:
                if (!(292 > -32768)) {
                    if true {
                        var v39 uint64 = (v35 >> (((v31 % ((18446744073709551615) | 1))) & 63))
                        if (42 != -5) {
                            v3 += v10
                            if false {
                                v40 := GenS4{F0: 0, F1: 255}
                                _ = v40
                            }
                        }
                        _ = v39
                    }
                    fmt.Println(fi36, fi36)
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    if ((false && (v35 > (^(v35)))) || false) {
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        _ = len(v13)
                        v2 -= -9223372036854775808
                    } else {
                        var v41 uint8 = uint8(v31)
                        _ = v41
                    }
                } else {
                    v1 = "\n\t"
                    v13 = append(v13, (fi37 ^ (fi36 - fi36)))
                }
            case -2147483648:
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
            default:
                v42 := GenS2{F0: v31, F1: 254, F2: (v2 << ((v31) & 63))}
                _ = v42
            }
            var v43 int = (fi37 | 128)
            _ = fi37
            _ = v43
        }
        _ = fi36
    }
    var v44 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v4
    _ = v44
    _ = v5
}
