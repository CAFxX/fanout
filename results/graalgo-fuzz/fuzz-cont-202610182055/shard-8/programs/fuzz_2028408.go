package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint8 }
type GenS2 struct { F0 int64; F1 int; F2 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32 }
func genFn0(p0 int8, p1 int32, p2 uint16) uint64 {
    p1 += -2147483648
    fmt.Println(p0)
    p2 -= 32767
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1(p0 bool) uint {
    v4 := GenS1{F0: 1, F1: 2, F2: 42}
    v5 := GenS1{F0: -1, F1: 1, F2: 216}
    for fi6 := 0; fi6 < 19; fi6++ {
        if false {
            switch "\n\t" {
            case "foo bar":
                if false {
                    v7 := []int{(fi6 >> ((2) & 63)), -158, -2147483648}
                    if p0 {
                        for fi8 := 0; fi8 < 9; fi8++ {
                            if p0 {
                                var v9 int16 = -1
                                _ = v9
                            } else {
                                v10 := GenS1{F0: 32767, F1: 42, F2: 128}
                                if (-2147483648 >= -680) {
                                    fmt.Println(p0, fi6)
                                } else {
                                    v7 = append(v7, fi8)
                                    fmt.Println(p0, fi6, fi6)
                                }
                                _ = v10
                            }
                            v11 := GenS1{F0: 0, F1: 42, F2: 42}
                            var v12 string = "hello"
                            if (p0 || (p0 && p0)) {
                                v12 = v12
                            }
                            v7 = append(v7, -9223372036854775808)
                            _ = fi8
                            _ = v11
                            _ = v12
                        }
                    }
                    fmt.Println(p0, fi6)
                    _ = v7
                }
            default:
                var v13 float32 = 3.14159
                _ = v13
            }
        } else {
            v14 := []int{32768, (fi6 * (fi6 << ((1) & 63)))}
            for fi15 := 0; fi15 < 5; fi15++ {
                v14 = append(v14, (fi15 * fi6))
                _ = fi15
            }
            if (("a" + "hello") <= "\n\t") {
                for fi16 := 0; fi16 < 17; fi16++ {
                    switch (^(fi16)) {
                    case 32768:
                        switch (("" + ((("a" + "\n\t") + (("hello" + "hello") + ("foo bar" + ("" + "")))) + ("hello" + "\n\t"))) + "") {
                        case "hello":
                            switch fi16 {
                            case -2147483648:
                                v17 := GenS1{F0: 1, F1: 497, F2: 42}
                                _ = v17
                            default:
                                var v18 uint16 = 732
                                _ = v18
                            }
                        case "a":
                            v19 := GenS1{F0: -32768, F1: 1, F2: 255}
                            _ = v19
                        case "\n\t":
                            _ = len(v14)
                        }
                        v14 = append(v14, fi6)
                        var v20 uint = 255
                        _ = v20
                    case -868:
                        p0 = p0
                    }
                    _ = len(v14)
                    _ = fi16
                }
            } else {
                var v21 string = "foo bar"
                v14 = append(v14, fi6)
                _ = v21
            }
            _ = v14
        }
        var v22 bool = ("hello" == ("hello" + (("" + "foo bar") + ("foo bar" + ""))))
        _ = fi6
        _ = v22
    }
    _ = p0
    _ = v4
    _ = v5
    return 0
}

func main() {
    var v1 uint32 = 42
    var v2 float64 = ((0.1 + (0.1 - 3.14159)) + 0.1)
    var v3 uint16 = 65535
    var v23 uint32 = 2
    if true {
        v24 := []int{0}
        _ = v24
    }
    v23 -= v23
    if (!true) {
        var v25 uint16 = 65535
        _ = v25
    } else {
        if true {
            if (false && false) {
                v26 := GenS1{F0: 1000, F1: 451, F2: 127}
                _ = v26
            }
        }
    }
    v3 = 653
    var v27 string = (((("hello" + ("\n\t" + ("foo bar" + ("hello" + "a")))) + "a") + "a") + "hello")
    if false {
        fmt.Println(v3, v27)
        var v28 int16 = -175
        v29 := GenS2{F0: -514, F1: 4294967295, F2: 762}
        _ = v28
        _ = v29
    } else {
        v30 := []int{4294967295, 32768}
        _ = v30
    }
    var v31 uint64 = 4294967295
    if false {
        if (0.1 > (float32(-2.25) - (0.1 / 3.14159))) {
            var v32 uint32 = v23
            fmt.Println(v32, v27, v32)
            for fi33 := 0; fi33 < 2; fi33++ {
                v3 += v3
                var v34 uint = 18446744073709551615
                v23 += (v32 / (((v23 % ((4294967295) | 1))) | 1))
                _ = fi33
                _ = v34
            }
            _ = v32
        } else {
            v35 := GenS1{F0: -274, F1: genFn1(true), F2: 255}
            _ = v35
        }
    } else {
        if ((!(v3 == v3)) && ((v23 ^ v1) <= uint32(v31))) {
            switch "" {
            case "a":
                switch (v27 + v27) {
                case "a":
                    var v36 int16 = -284
                    _ = v36
                default:
                    v37 := GenS2{F0: 851, F1: -1, F2: -1}
                    _ = v37
                }
            case "\n\t":
                var v38 float64 = -2.25
                _ = v38
            }
            var v39 uint16 = (v3 & 734)
            _ = v39
        } else {
            if false {
                fmt.Println(v1, v31)
            } else {
                v40 := GenS1{F0: 1000, F1: 534, F2: 0}
                v3 = 65535
                v41 := []int{1}
                _ = v40
                _ = v41
            }
            switch v27 {
            case "a":
                if (((v31 <= v31) || true) || (-599 <= -9223372036854775808)) {
                    v2 += v2
                    v42 := GenS2{F0: 42, F1: 42, F2: 1}
                    v31 = (v31 ^ (v31 * 2))
                    _ = v42
                }
            case "hello":
                v2 = v2
            case "\n\t":
                v3 = v3
            default:
                v27 = ("hello" + v27)
            }
            v1 = (v23 - (v1 >> ((18446744073709551615) & 63)))
        }
        fmt.Println(v31, v1)
    }
    for fi43 := 0; fi43 < 9; fi43++ {
        var v44 float32 = (1.5 * -2.25)
        _ = fi43
        _ = v44
    }
    v1 -= v23
    v45 := []int{-614}
    var v46 uint32 = v1
    if (true && true) {
        v2 = (((((v2 + v2) / 0.1) / v2) - v2) + v2)
        for fi47 := 0; fi47 < 7; fi47++ {
            v1 -= (v46 ^ 877)
            if len(v45) >= 2 {
                v45 = v45[1:len(v45)]
            }
            _ = fi47
        }
        fmt.Println(v31, v31)
    } else {
        v48 := GenS3{F0: 9223372036854775807}
        v45 = append(v45, 4294967295)
        _ = v48
    }
    v49 := GenS4{F0: 793}
    fmt.Println(v3)
    fmt.Println(v46)
    fmt.Println(v1, v46, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v31
    _ = v45
    _ = v46
    _ = v49
}
