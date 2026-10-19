package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 int64 }
func genFn0(p0 int64) uint8 {
    p0 -= p0
    fmt.Println(p0)
    fmt.Println(p0)
    v5 := GenS1{F0: (p0 % (((p0 << ((410) & 63))) | 1))}
    var v6 float32 = 3.14159
    _ = p0
    _ = v5
    _ = v6
    return 4
}
func genFn1() uint {
    switch 42 {
    case 32768:
        var v7 uint16 = 65535
        _ = v7
    }
    switch "" {
    case "hello":
        switch 2147483647 {
        case -9223372036854775808:
            v8 := GenS1{F0: 0}
            _ = v8
        case -928:
            for fi9 := 0; fi9 < 2; fi9++ {
                var v10 int = (fi9 | fi9)
                _ = fi9
                _ = v10
            }
        }
    case "foo bar":
        if (!(true && (true && true))) {
            var v11 uint = 255
            v12 := GenS2{F0: "a", F1: 32767}
            var v13 uint = v11
            _ = v11
            _ = v12
            _ = v13
        } else {
            if true {
                for fi14 := 0; fi14 < 18; fi14++ {
                    var v15 bool = false
                    for fi16 := 0; fi16 < 10; fi16++ {
                        fmt.Println(v15, fi14, v15)
                        v17 := GenS1{F0: 1}
                        if (v15 && (-1 == 42)) {
                            v15 = v15
                        } else {
                            v18 := GenS3{F0: -1, F1: fi14}
                            _ = v18
                        }
                        v19 := []int{fi16, 42, (fi14 % (((fi16 | fi14)) | 1)), -369}
                        _ = fi16
                        _ = v17
                        _ = v19
                    }
                    fmt.Println(fi14, v15)
                    _ = fi14
                    _ = v15
                }
            } else {
                var v20 int64 = -475
                _ = v20
            }
            switch -158 {
            case 32768:
                if (32767 <= 1) {
                    var v21 uint16 = 592
                    for fi22 := 0; fi22 < 5; fi22++ {
                        v21 -= v21
                        v23 := []int{fi22}
                        if false {
                            switch "" {
                            case "a":
                                var v24 uint8 = 144
                                _ = v24
                            default:
                                fmt.Println(fi22)
                            }
                            v25 := GenS1{F0: -9223372036854775808}
                            if false {
                                _ = len(v23)
                                if (-9223372036854775808 >= 165) {
                                    fmt.Println(v21, fi22)
                                    var v26 uint8 = 2
                                    _ = v26
                                } else {
                                    if len(v23) >= 2 {
                                        v23 = v23[1:len(v23)]
                                    }
                                    if true {
                                        var v27 float32 = (3.14159 + 0.1)
                                        if (float64(float32((float64(v27) + (1.5 - float64(((v27 * 0.0) / (((((v27 + (float32((1.5 * (0.0 / 1.5))) / -2.25)) - 1.5) - v27) * 0.1) * ((1e10 / (1.5 + 0.1)) / 0.1)))))))) < (float64((v27 * ((float32(1e10) + v27) + float32((0.0 * -2.25))))) + 3.14159)) {
                                            var v28 float32 = v27
                                            var v29 int = (fi22 & fi22)
                                            _ = v28
                                            _ = v29
                                        } else {
                                            var v30 uint32 = 1
                                            v21 += 483
                                            v31 := GenS4{F0: (fi22 ^ fi22)}
                                            _ = v30
                                            _ = v31
                                        }
                                        var v32 string = "\n\t"
                                        _ = v27
                                        _ = v32
                                    }
                                }
                            }
                            _ = v25
                        }
                        _ = fi22
                        _ = v23
                    }
                    _ = v21
                } else {
                    switch "" {
                    case "":
                        v33 := []int{-701, 0}
                        _ = v33
                    case "a":
                        var v34 int32 = 42
                        _ = v34
                    }
                    var v35 uint32 = 2147483648
                    _ = v35
                }
            }
            var v36 uint8 = 211
            _ = v36
        }
    case "a":
        for fi37 := 0; fi37 < 12; fi37++ {
            var v38 uint32 = 2147483647
            v39 := GenS5{F0: "a", F1: -528}
            var v40 uint64 = 4294967295
            _ = fi37
            _ = v38
            _ = v39
            _ = v40
        }
    default:
        var v41 uint = 4294967295
        _ = v41
    }
    return 601
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int32 = 1
    v3 := make(map[uint32]uint32)
    v4 := []string{"\n\t"}
    switch ((((("a" + (("hello" + "foo bar") + ("a" + ("" + ("hello" + "a"))))) + "hello") + "a") + "a") + "\n\t") {
    case "\n\t":
        var v42 uint8 = 128
        _ = v42
    case "hello":
        fmt.Println(v2)
    case "foo bar":
        if (false && ((0 != 255) && false)) {
            var v43 float64 = 0.0
            for fi44 := 0; fi44 < 14; fi44++ {
                v4 = append(v4, "")
                v1 = -493
                _ = fi44
            }
            v45 := GenS4{F0: 561}
            _ = v43
            _ = v45
        }
    }
    if (((v1 << ((853) & 63)) >= v1) && false) {
        for fi46 := 0; fi46 < 16; fi46++ {
            v1 += v1
            var v47 int16 = -218
            _ = fi46
            _ = v47
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v48 int32 = (v2 >> ((1) & 63))
    v48 += v2
    var v49 uint64 = 2
    var v50 float64 = 1e10
    switch 42 {
    case -9223372036854775808:
        v48 = (-(v2))
    case -443:
        v51 := GenS3{F0: 730, F1: 0}
        _ = v51
    }
    var v52 int32 = (v2 * 2147483647)
    fmt.Println(v48, v52)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v50
    _ = v52
}
