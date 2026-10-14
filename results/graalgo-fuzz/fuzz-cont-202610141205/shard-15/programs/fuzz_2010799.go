package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS2 struct { F0 int32; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint64; F1 bool; F2 int }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 int16; F1 uint8; F2 string }
type GenS6 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS7 struct { F0 uint; F1 uint64; F2 uint16 }
func genFn0() bool {
    var v5 int8 = -1
    switch 2147483647 {
    case -2147483648:
        var v6 uint = 0
        _ = v6
    case 0:
        fmt.Println(v5)
    case -930:
        switch 32767 {
        case 327:
            var v7 int8 = 55
            _ = v7
        }
    default:
        v8 := GenS1{F0: 42, F1: 1, F2: 635}
        _ = v8
    }
    _ = v5
    return ((true || (42 < 42)) && ((!(false && (-9223372036854775808 <= 9223372036854775807))) && true))
}
func genFn1(p0 uint64) int16 {
    p0 = p0
    if ("\n\t" <= "") {
        var v9 bool = (651 == 531)
        v10 := []int{-9223372036854775808}
        v10 = append(v10, 128)
        _ = v10
        _ = v9
    } else {
        v11 := GenS2{F0: 0, F1: 0, F2: 42}
        p0 = (p0 & (^(p0)))
        p0 -= p0
        _ = v11
    }
    var v12 int32 = 0
    if (uint8(p0) <= 90) {
        v12 = 1
        var v13 uint16 = 305
        for fi14 := 0; fi14 < 7; fi14++ {
            v15 := GenS3{F0: p0, F1: false, F2: (fi14 ^ 127)}
            fmt.Println(v12, p0, fi14)
            v16 := []int{fi14, 32768, fi14}
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v13
    } else {
        v17 := GenS4{F0: 2, F1: -9223372036854775808}
        p0 = (p0 / (((p0 - p0)) | 1))
        _ = v17
    }
    _ = p0
    _ = v12
    return 1
}

func main() {
    var v1 int = -534
    v2 := []uint32{}
    var v3 string = ("hello" + ("hello" + ""))
    var v4 uint64 = 0
    switch v1 {
    case 0:
        for fi18 := 0; fi18 < 15; fi18++ {
            for fi19 := 0; fi19 < 14; fi19++ {
                v20 := GenS5{F0: genFn1((^(v4))), F1: 188, F2: v3}
                _ = fi19
                _ = v20
            }
            _ = len(v2)
            if (-32768 <= 0) {
                fmt.Println(v3, v1, v4)
                v21 := GenS2{F0: -2147483648, F1: 42, F2: 0}
                v3 = ("hello" + ((((v3 + ((("" + v3) + v3) + (v3 + v3))) + (v3 + "hello")) + v3) + (("foo bar" + "\n\t") + ("\n\t" + ((((v3 + ((v3 + v3) + (((v3 + v3) + "foo bar") + (v3 + v3)))) + "a") + (v3 + v3)) + v3)))))
                _ = v21
            } else {
                for fi22 := 0; fi22 < 10; fi22++ {
                    v2 = append(v2, 1)
                    if false {
                        if (!true) {
                            if false {
                                v23 := GenS4{F0: 2, F1: 42}
                                _ = v23
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = 2
                                }
                                v2 = append(v2, 42)
                            }
                        }
                        if len(v2) > 0 {
                            v2[0] = 337
                        }
                    } else {
                        if (true || true) {
                            v4 = v4
                            fmt.Println(v3, v3, fi22)
                            var v24 int = (^(fi22))
                            _ = v24
                        } else {
                            v1 -= fi18
                            fmt.Println(v1, v4)
                            v25 := GenS6{F0: 1000, F1: (v4 ^ v4), F2: 2}
                            _ = v25
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v26 int64 = -9223372036854775808
                        _ = v26
                    }
                    _ = fi22
                }
                v1 -= -689
            }
            _ = fi18
        }
    case 4294967295:
        var v27 uint8 = 127
        _ = v27
    case 183:
        var v28 uint = 42
        _ = v28
    }
    if (2147483648 >= 2147483647) {
        for fi29 := 0; fi29 < 14; fi29++ {
            switch fi29 {
            case -1:
                fmt.Println(v3)
            case -2147483648:
                v4 -= v4
            }
            v3 = ((v3 + (v3 + (v3 + v3))) + (((v3 + v3) + "foo bar") + v3))
            if false {
                v1 = (fi29 >> ((917) & 63))
            } else {
                v30 := GenS7{F0: 18446744073709551615, F1: (v4 * 900), F2: 498}
                v2 = append(v2, 65535)
                _ = v30
            }
            _ = fi29
        }
    } else {
        for fi31 := 0; fi31 < 17; fi31++ {
            fmt.Println(v3, v4, v3)
            v32 := GenS7{F0: 0, F1: v4, F2: 255}
            var v33 bool = false
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    v4 = (v4 & v4)
    var v34 int8 = -128
    var v35 float64 = 1e10
    var v36 string = v3
    if ((((v3 + "\n\t") + "\n\t") + (("a" + "a") + v3)) >= v36) {
        var v37 int8 = v34
        _ = v37
    }
    var v38 uint16 = 0
    var v39 uint = 4294967295
    v39 = (v39 >> ((v39) & 63))
    v2 = append(v2, 2147483647)
    fmt.Println(v39, v1)
    var v40 float32 = (float32(0.0) - 1.5)
    v4 = 520
    var v41 int64 = -9223372036854775808
    v42 := GenS2{F0: -710, F1: v41, F2: (^(v38))}
    v43 := GenS3{F0: v4, F1: true, F2: v1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
}
