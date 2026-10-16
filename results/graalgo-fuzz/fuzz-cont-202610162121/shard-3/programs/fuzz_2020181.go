package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint8; F1 int8; F2 int8 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 bool; F1 uint8; F2 uint64 }
func genFn0(p0 int) uint {
    p0 += (p0 % (((p0 | (p0 / ((127) | 1)))) | 1))
    v4 := []int{}
    fmt.Println(p0)
    for fi5 := 0; fi5 < 3; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            p0 = 32767
            fmt.Println(fi5)
            var v7 uint64 = 42
            _ = fi6
            _ = v7
        }
        fmt.Println(fi5)
        switch "" {
        case "":
            v8 := GenS1{F0: 65535, F1: 736, F2: 2147483647}
            _ = v8
        case "foo bar":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            p0 = (p0 * 32767)
        }
        _ = fi5
    }
    _ = p0
    _ = v4
    return 323
}
func genFn1(p0 uint32) int16 {
    fmt.Println(p0)
    var v9 string = "hello"
    fmt.Println(v9, p0)
    var v10 uint64 = 346
    var v11 int32 = -2147483648
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return 32767
}

func main() {
    var v1 uint = 0
    var v2 uint64 = 2
    var v3 int32 = 42
    if ("" != "a") {
        v1 -= v1
    } else {
        var v12 int8 = 42
        if (!false) {
            v2 += v2
            var v13 uint64 = (v2 >> ((v1) & 63))
            _ = v13
        }
        switch 32768 {
        case -623:
            fmt.Println(v2, v3, v3)
        case 997:
            if true {
                v14 := GenS2{F0: 0, F1: v12, F2: (v12 ^ v12)}
                _ = v14
            } else {
                v15 := GenS3{F0: int32(v12), F1: 42}
                v3 = (v3 / (((-(v3))) | 1))
                _ = v15
            }
        }
        _ = v12
    }
    v2 += 702
    v16 := []int{}
    v1 = (v1 & v1)
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    v2 += v2
    fmt.Println(v3, v1, v3)
    var v17 string = ("\n\t" + "\n\t")
    if (2 != 0) {
        var v18 int16 = -656
        switch 128 {
        case 128:
            if (((v17 + ((v17 + v17) + ("hello" + (v17 + "a")))) + ("hello" + ("\n\t" + ((v17 + ("a" + ("foo bar" + "a"))) + v17)))) < (v17 + (("a" + (v17 + v17)) + ((v17 + v17) + ((("\n\t" + v17) + "") + ""))))) {
                switch (v17 + "\n\t") {
                case "":
                    var v19 int32 = v3
                    _ = v19
                case "foo bar":
                    for fi20 := 0; fi20 < 12; fi20++ {
                        v21 := GenS1{F0: 4294967295, F1: 42, F2: 2147483648}
                        _ = fi20
                        _ = v21
                    }
                default:
                    var v22 int16 = -765
                    _ = v22
                }
                if true {
                    v2 = (v2 / ((570) | 1))
                    if (1 > -689) {
                        v2 = (v2 % ((v2) | 1))
                        var v23 int64 = 9223372036854775807
                        v24 := GenS4{F0: ((1 == 2147483647) && (float64(0.0) < 0.1)), F1: 127, F2: (v2 - 950)}
                        _ = v23
                        _ = v24
                    }
                    if true {
                        if false {
                            v25 := GenS4{F0: false, F1: 128, F2: (v2 / (((v2 ^ 0)) | 1))}
                            _ = v25
                        } else {
                            switch "\n\t" {
                            case "\n\t":
                                if len(v16) >= 2 {
                                    v16 = v16[1:len(v16)]
                                }
                            case "foo bar":
                                var v26 uint = genFn0(-9223372036854775808)
                                _ = v26
                            }
                            for fi27 := 0; fi27 < 14; fi27++ {
                                v16 = append(v16, -9223372036854775808)
                                _ = fi27
                            }
                            v28 := GenS3{F0: (v3 ^ v3), F1: (v18 % (((-(v18))) | 1))}
                            _ = v28
                        }
                    } else {
                        for fi29 := 0; fi29 < 14; fi29++ {
                            var v30 uint32 = 683
                            if true {
                                var v31 int = 32768
                                v17 = "foo bar"
                                _ = v31
                            } else {
                                v32 := GenS3{F0: (v3 | 42), F1: (v18 - v18)}
                                _ = v32
                            }
                            _ = fi29
                            _ = v30
                        }
                        var v33 string = ((v17 + "hello") + v17)
                        _ = v33
                    }
                }
            }
        case 1:
            var v34 uint32 = 21
            _ = v34
        }
        _ = v18
    } else {
        var v35 float32 = -2.25
        v36 := GenS1{F0: 0, F1: 32767, F2: 4294967295}
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 4; fi37++ {
        v17 = v17
        var v38 float32 = 1e10
        _ = fi37
        _ = v38
    }
    var v39 int8 = 100
    if false {
        v39 = v39
        var v40 float32 = ((-2.25 + 1e10) + (float32(-2.25) - float32(3.14159)))
        v39 += v39
        _ = v40
    }
    var v41 string = v17
    fmt.Println(v2, v2, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v39
    _ = v41
}
