package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint }
type GenS2 struct { F0 int16; F1 uint; F2 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS5 struct { F0 uint; F1 uint32; F2 int64 }
func genFn0(p0 int) int16 {
    for fi4 := 0; fi4 < 2; fi4++ {
        if (false || true) {
            v5 := []int{(p0 + p0)}
            _ = v5
        } else {
            p0 = p0
            fmt.Println(p0, fi4)
            fmt.Println(p0, p0)
        }
        _ = fi4
    }
    var v6 uint16 = 0
    p0 += 659
    if true {
        v7 := GenS1{F0: 9223372036854775807, F1: 42, F2: 584}
        v8 := GenS1{F0: 1, F1: 833, F2: 42}
        v9 := GenS2{F0: -910, F1: 18446744073709551615, F2: 2}
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: 1000, F1: 2, F2: 2}
    _ = p0
    _ = v10
    _ = v6
    return int16(v6)
}
func genFn1(p0 int, p1 int8) int {
    var v11 float64 = 3.14159
    v12 := []int{p0, p0, (p0 << ((841) & 63))}
    v13 := GenS1{F0: 42, F1: 2147483648, F2: 1}
    if (!true) {
        for fi14 := 0; fi14 < 13; fi14++ {
            if len(v12) > 0 {
                v12[0] = p0
            }
            for fi15 := 0; fi15 < 16; fi15++ {
                if true {
                    v16 := GenS2{F0: 42, F1: 42, F2: 32}
                    if true {
                        v12 = append(v12, fi14)
                        var v17 int = (^(p0))
                        _ = v17
                    }
                    _ = v16
                } else {
                    var v18 int = (fi14 ^ 704)
                    p1 -= -32
                    v18 += 32767
                    _ = v18
                }
                v19 := GenS1{F0: 1, F1: 2147483648, F2: 553}
                if false {
                    v20 := GenS2{F0: 952, F1: 255, F2: 42}
                    if (1 == 142) {
                        fmt.Println(p1)
                        fmt.Println(fi14, v11, p0)
                        var v21 string = ("" + ("" + "foo bar"))
                        _ = v21
                    }
                    switch 300 {
                    case -838:
                        fmt.Println(v11)
                    case 0:
                        v22 := GenS3{F0: 42}
                        _ = v22
                    case -394:
                        _ = len(v12)
                        if true {
                            switch p0 {
                            case 128:
                                p1 += -128
                            default:
                                var v23 string = ""
                                _ = v23
                            }
                            var v24 uint8 = 19
                            _ = v24
                        }
                    default:
                        switch 1 {
                        case -9223372036854775808:
                            var v25 int8 = p1
                            _ = v25
                        case 128:
                            _ = len(v12)
                        case 627:
                            fmt.Println(fi14)
                        }
                    }
                    _ = v20
                } else {
                    v26 := GenS1{F0: 42, F1: 1, F2: 333}
                    _ = v26
                }
                _ = fi15
                _ = v19
            }
            var v27 float32 = (0.1 / (float32(0.0) - 1.5))
            _ = fi14
            _ = v27
        }
        p0 = 42
        var v28 int64 = -9223372036854775808
        _ = v28
    }
    var v29 uint32 = 401
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    _ = v29
    return p0
}

func main() {
    var v1 uint32 = 2
    var v2 bool = true
    var v3 bool = (1 > 631)
    switch ("\n\t" + ("hello" + ("\n\t" + "\n\t"))) {
    case "":
        if true {
            v2 = (-128 > -1)
            v30 := GenS3{F0: -1}
            _ = v30
        } else {
            v2 = (!(48 < 134))
            v2 = v3
        }
    }
    v2 = v2
    if v3 {
        v31 := GenS4{F0: 498, F1: 0, F2: -1}
        var v32 bool = (42 < 0)
        v1 = v1
        _ = v31
        _ = v32
    } else {
        switch -9223372036854775808 {
        case 32768:
            var v33 float64 = 1.5
            _ = v33
        case -456:
            v2 = v2
        }
        for fi34 := 0; fi34 < 6; fi34++ {
            for fi35 := 0; fi35 < 5; fi35++ {
                if (true || (!true)) {
                    v36 := GenS1{F0: -348, F1: v1, F2: 1}
                    _ = v36
                }
                v37 := GenS5{F0: 394, F1: (v1 % (((v1 << ((4294967295) & 63))) | 1)), F2: 0}
                var v38 int32 = 42
                _ = fi35
                _ = v37
                _ = v38
            }
            if (!(!(!(((1.5 / (0.0 + -2.25)) <= 0.0) && (2147483647 > (v1 | (v1 / (((v1 << ((1) & 63))) | 1)))))))) {
                if ((((0.1 + ((3.14159 * 1.5) + 0.0)) + -2.25) <= (float32(-2.25) * float32(float64((0.0 * (float32(float64(1.5)) / -2.25)))))) || true) {
                    v2 = true
                }
                v2 = v3
            }
            if v3 {
                for fi39 := 0; fi39 < 8; fi39++ {
                    v40 := []int{genFn1(genFn1(fi39, 42), -128)}
                    v2 = true
                    _ = fi39
                    _ = v40
                }
                v41 := []int{}
                _ = v41
            }
            _ = fi34
        }
        v1 -= 1
    }
    v42 := []int{-781}
    v1 = 630
    var v43 int = 9223372036854775807
    v2 = v3
    v1 += (^(v1))
    fmt.Println(v3, v3, v43)
    if len(v42) >= 2 {
        v42 = v42[1:len(v42)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v43
}
