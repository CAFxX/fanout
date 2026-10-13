package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 int8; F1 uint }
type GenS6 struct { F0 int8; F1 bool; F2 string }
type GenS7 struct { F0 int32 }

func main() {
    var v1 uint16 = 255
    v2 := make(map[uint32]int)
    v3 := []string{"a", "a", "foo bar"}
    v4 := GenS1{F0: 1, F1: 806, F2: -464}
    if true {
        for fi5 := 0; fi5 < 15; fi5++ {
            fmt.Println(v1)
            fmt.Println(fi5)
            _ = fi5
        }
    } else {
        switch -448 {
        case 127:
            if true {
                var v6 string = ("" + ("hello" + "hello"))
                _ = v6
            } else {
                var v7 int64 = 820
                v7 = v7
                _ = v7
            }
        default:
            v8 := GenS1{F0: 128, F1: 0, F2: -1}
            _ = v8
        }
    }
    var v9 uint8 = 127
    if true {
        if (-852 < 57) {
            for fi10 := 0; fi10 < 18; fi10++ {
                v11 := GenS2{F0: 1000, F1: 0}
                _ = fi10
                _ = v11
            }
        }
        fmt.Println(v1)
    } else {
        var v12 bool = ((4294967295 > 2) && (-855 <= 32767))
        _ = v12
    }
    v13 := GenS1{F0: v9, F1: uint32(v1), F2: 1}
    if true {
        var v14 string = "foo bar"
        if (!(954 < 786)) {
            v15 := GenS1{F0: 1, F1: 2, F2: -1}
            if true {
                if (-2147483648 <= 1) {
                    v1 -= (v1 ^ v1)
                    for fi16 := 0; fi16 < 4; fi16++ {
                        var v17 int8 = 116
                        v1 += 17
                        var v18 bool = (1.5 != 3.14159)
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    v14 = v14
                } else {
                    v19 := GenS1{F0: v9, F1: 0, F2: 1}
                    for fi20 := 0; fi20 < 19; fi20++ {
                        var v21 int16 = 1
                        for fi22 := 0; fi22 < 8; fi22++ {
                            if (v1 != v1) {
                                switch fi22 {
                                case -9223372036854775808:
                                    if false {
                                        if len(v3) > 0 {
                                            v3[0] = v14
                                        }
                                        fmt.Println(v1, v1, v21)
                                    }
                                }
                                v23 := GenS3{F0: 690, F1: v21}
                                _ = v23
                            } else {
                                if ((fi22 != (-(fi20))) && (!true)) {
                                    if (4294967295 > 240) {
                                        v2[65535] = 0
                                        v24 := GenS3{F0: 44, F1: (v21 >> ((4294967295) & 63))}
                                        _ = v24
                                    }
                                    v14 = (v14 + v14)
                                    v25 := GenS4{F0: 80, F1: -630}
                                    _ = v25
                                }
                                switch fi20 {
                                case 9223372036854775807:
                                    var v26 uint8 = (^(v9))
                                    _ = v26
                                default:
                                    fmt.Println(v1)
                                }
                                var v27 float32 = ((3.14159 / 1.0) - float32(((float64((1.5 / 3.14159)) * (0.0 + (float64(float32((1e10 + float64(-2.25)))) / -2.25))) * (float64(0.0) / 1.5))))
                                _ = v27
                            }
                            v28 := GenS5{F0: 0, F1: 0}
                            _ = fi22
                            _ = v28
                        }
                        v21 = 42
                        _ = fi20
                        _ = v21
                    }
                    v29 := v2[60]
                    _ = v19
                    _ = v29
                }
                var v30 uint8 = v9
                _ = v30
            } else {
                if true {
                    delete(v2, 392)
                    _ = len(v2)
                    switch 42 {
                    case -2147483648:
                        v31 := GenS3{F0: 4294967295, F1: 32767}
                        _ = v31
                    case 32767:
                        fmt.Println(v1)
                    case -807:
                        for fi32 := 0; fi32 < 11; fi32++ {
                            v33 := GenS6{F0: 0, F1: false, F2: v14}
                            _ = fi32
                            _ = v33
                        }
                    }
                }
            }
            _ = v15
        } else {
            var v34 uint64 = 573
            if true {
                v35 := GenS7{F0: 0}
                _ = v35
            } else {
                var v36 int16 = -380
                var v37 uint8 = v9
                _ = v36
                _ = v37
            }
            v34 = 987
            _ = v34
        }
        _ = v14
    }
    var v38 uint8 = v9
    for fi39 := 0; fi39 < 10; fi39++ {
        var v40 float64 = float64(1.5)
        _ = fi39
        _ = v40
    }
    fmt.Println(v9, v38)
    var v41 int8 = -47
    var v42 int8 = (v41 >> ((1) & 63))
    var v43 int = 2147483647
    var v44 uint64 = 302
    var v45 int64 = -968
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v9
}
