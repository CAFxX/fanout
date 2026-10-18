package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint; F2 string }
type GenS3 struct { F0 int8; F1 int64; F2 int }
type GenS4 struct { F0 uint; F1 uint; F2 uint }
type GenS5 struct { F0 uint32; F1 uint16; F2 int64 }

func main() {
    v1 := []uint32{2, 1}
    v2 := []int{-9223372036854775808, 813, 9223372036854775807, 994}
    v3 := make(map[string]uint32)
    for fi4 := 0; fi4 < 18; fi4++ {
        _ = len(v1)
        var v5 uint32 = 2147483648
        var v6 string = "foo bar"
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 float32 = 0.1
    v7 = ((v7 + (v7 - v7)) - v7)
    if true {
        var v8 uint32 = 968
        for fi9 := 0; fi9 < 19; fi9++ {
            v8 = 42
            _ = fi9
        }
        _ = v8
    }
    if ((!(((v7 * v7) - -2.25) <= v7)) && false) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v10 float64 = 1.5
        v7 = v7
        _ = v10
    }
    if (!true) {
        for fi11 := 0; fi11 < 2; fi11++ {
            switch 32768 {
            case 559:
                v3[("a" + ("foo bar" + "foo bar"))] = 636
            case 4294967295:
                if true {
                    delete(v3, "\n\t")
                    if (!true) {
                        if true {
                            for fi12 := 0; fi12 < 9; fi12++ {
                                var v13 uint8 = 114
                                v14 := GenS1{F0: 42, F1: 42}
                                _ = fi12
                                _ = v13
                                _ = v14
                            }
                            var v15 uint16 = 32767
                            if false {
                                var v16 int8 = 100
                                for fi17 := 0; fi17 < 12; fi17++ {
                                    var v18 int64 = 9223372036854775807
                                    _ = fi17
                                    _ = v18
                                }
                                v16 += v16
                                _ = v16
                            } else {
                                v19 := GenS1{F0: 996, F1: 0}
                                if false {
                                    v7 += float32((((1.5 * float64(((0.1 + v7) + 1.5))) * 1.5) + 1.5))
                                    fmt.Println(v15, v15)
                                    for fi20 := 0; fi20 < 11; fi20++ {
                                        v1 = append(v1, 2147483648)
                                        v21 := GenS2{F0: 118, F1: 4294967295, F2: (("a" + "a") + "hello")}
                                        var v22 uint8 = 0
                                        switch fi11 {
                                        case 9223372036854775807:
                                            v7 -= v7
                                            delete(v3, "\n\t")
                                            fmt.Println(fi20, v22, fi20)
                                        default:
                                            switch -75 {
                                            case 32767:
                                                v15 = 65535
                                            case 128:
                                                if false {
                                                    v23 := GenS2{F0: (v22 << ((0) & 63)), F1: 0, F2: ("foo bar" + "\n\t")}
                                                    v24 := GenS3{F0: -1, F1: -353, F2: 0}
                                                    v1 = append(v1, 2147483647)
                                                    _ = v23
                                                    _ = v24
                                                }
                                            }
                                        }
                                        var v25 int32 = -252
                                        _ = fi20
                                        _ = v21
                                        _ = v22
                                        _ = v25
                                    }
                                } else {
                                    switch "" {
                                    case "foo bar":
                                        var v26 int32 = 0
                                        _ = v26
                                    case "\n\t":
                                        v7 = (v7 / v7)
                                    }
                                    fmt.Println(v7, fi11, fi11)
                                }
                                v27 := GenS4{F0: 626, F1: 4294967295, F2: 136}
                                _ = v19
                                _ = v27
                            }
                            _ = v15
                        } else {
                            var v28 uint64 = 42
                            v3["hello"] = 0
                            for fi29 := 0; fi29 < 4; fi29++ {
                                fmt.Println(v28)
                                _ = fi29
                            }
                            _ = v28
                        }
                        for fi30 := 0; fi30 < 16; fi30++ {
                            switch (fi30 >> ((872) & 63)) {
                            case -307:
                                v31 := GenS2{F0: 255, F1: 0, F2: ("a" + "")}
                                _ = v31
                            }
                            v32 := GenS3{F0: 127, F1: -439, F2: fi30}
                            _ = fi30
                            _ = v32
                        }
                    }
                } else {
                    for fi33 := 0; fi33 < 18; fi33++ {
                        v7 = (0.0 / 1.0)
                        v7 -= 0.0
                        _ = fi33
                    }
                }
            case 720:
                if (1e10 != float32(1.5)) {
                    _ = len(v1)
                    v7 = 3.14159
                }
            }
            _ = fi11
        }
        if true {
            v7 -= float32(1e10)
            v7 = 0.1
            switch -1 {
            case -30:
                for fi34 := 0; fi34 < 8; fi34++ {
                    if false {
                        v7 -= v7
                        fmt.Println(fi34)
                        var v35 uint16 = 2
                        _ = v35
                    }
                    v36 := GenS2{F0: 1, F1: 42, F2: "foo bar"}
                    _ = fi34
                    _ = v36
                }
            case -2147483648:
                for fi37 := 0; fi37 < 8; fi37++ {
                    v38 := v3[("" + ("a" + ("" + "")))]
                    _ = fi37
                    _ = v38
                }
            default:
                v39 := GenS5{F0: 2147483648, F1: 65535, F2: -9223372036854775808}
                _ = v39
            }
        }
    } else {
        fmt.Println(v7)
        for fi40 := 0; fi40 < 6; fi40++ {
            if false {
                v7 -= ((float32(0.0) + 1e10) * 1e10)
                delete(v3, "hello")
                v7 += -2.25
            } else {
                var v41 int64 = 42
                if (!(!false)) {
                    if false {
                        v2 = append(v2, (-(fi40)))
                        v42 := GenS3{F0: -128, F1: v41, F2: fi40}
                        _ = v42
                    } else {
                        var v43 bool = (42 >= 32767)
                        var v44 bool = v43
                        v41 -= (v41 ^ (v41 - v41))
                        _ = v43
                        _ = v44
                    }
                }
                v7 = v7
                _ = v41
            }
            var v45 int32 = 139
            var v46 uint16 = 816
            _ = fi40
            _ = v45
            _ = v46
        }
        fmt.Println(v7)
    }
    _ = len(v3)
    v47 := GenS1{F0: 1, F1: 42}
    v7 = v7
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v47
    _ = v7
}
