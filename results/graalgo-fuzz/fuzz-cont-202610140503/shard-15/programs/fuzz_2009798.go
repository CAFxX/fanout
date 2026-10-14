package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int32; F1 bool; F2 int64 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 string; F1 int32; F2 uint32 }
type GenS5 struct { F0 int8; F1 uint }
type GenS6 struct { F0 string }
func genFn0(p0 uint8, p1 int16) int64 {
    v5 := []int{}
    var v6 int32 = 2147483647
    v7 := GenS1{F0: 1, F1: 1}
    p1 = p1
    p0 += (^(p0))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return -864
}
func genFn1(p0 bool, p1 bool, p2 int32) uint {
    if p1 {
        p2 -= p2
        for fi8 := 0; fi8 < 12; fi8++ {
            v9 := []int{fi8, -9223372036854775808, (-(fi8)), 2147483647}
            switch fi8 {
            case 0:
                fmt.Println(p0)
            case 121:
                for fi10 := 0; fi10 < 14; fi10++ {
                    var v11 int8 = -61
                    v12 := GenS1{F0: 65535, F1: 0}
                    p0 = p0
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            case 9223372036854775807:
                var v13 uint = 42
                _ = v13
            default:
                switch (^(fi8)) {
                case 2147483647:
                    v14 := GenS2{F0: p2, F1: ((p2 / ((p2) | 1)) == -2147483648), F2: -9223372036854775808}
                    _ = v14
                default:
                    var v15 uint32 = 0
                    _ = v15
                }
            }
            _ = fi8
            _ = v9
        }
        switch (("hello" + ("hello" + (("\n\t" + "\n\t") + ""))) + (("a" + ("\n\t" + "a")) + ((("\n\t" + ("foo bar" + "a")) + ("hello" + "a")) + "foo bar"))) {
        case "hello":
            var v16 bool = p1
            _ = v16
        default:
            var v17 string = "foo bar"
            _ = v17
        }
    }
    switch 32768 {
    case 652:
        p2 = -1
    case 1:
        fmt.Println(p0, p2, p2)
    case -41:
        v18 := GenS3{F0: "hello", F1: 1}
        _ = v18
    default:
        v19 := []int{0, 32768, 32768}
        _ = v19
    }
    if p1 {
        var v20 int8 = -63
        v21 := []int{2147483647, 76, 127, 128}
        v20 += (-(v20))
        _ = v20
        _ = v21
    } else {
        p0 = p0
    }
    p1 = false
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 int16 = -253
    var v2 int64 = -9223372036854775808
    v3 := make(map[string]uint32)
    var v4 int32 = -30
    v22 := []int{477, -968, -967, 32767}
    if ((v4 % ((1) | 1)) == v4) {
        v23 := v3[""]
        v1 = 225
        _ = v23
    }
    v24 := v3["foo bar"]
    _ = len(v22)
    v25 := GenS4{F0: "hello", F1: int32(v24), F2: 755}
    if (!false) {
        for fi26 := 0; fi26 < 12; fi26++ {
            fmt.Println(v2)
            if (v4 == (v4 + -1)) {
                if (2 != 659) {
                    v4 -= (v4 | v4)
                    v1 = int16(v2)
                }
            }
            _ = fi26
        }
    }
    v2 = v2
    v1 += (v1 + -119)
    for fi27 := 0; fi27 < 18; fi27++ {
        if false {
            v24 += (v24 >> ((2) & 63))
            fmt.Println(v24, v4, fi27)
        } else {
            if (true && (57 == 514)) {
                var v28 float64 = -2.25
                v29 := GenS2{F0: (v4 | int32(v2)), F1: false, F2: (v2 & (^(v2)))}
                _ = v28
                _ = v29
            }
            fmt.Println(v4)
        }
        _ = fi27
    }
    if ((343 <= 42) && (0.0 < 0.0)) {
        _ = len(v3)
    }
    v24 += v24
    if len(v22) > 0 {
        v22[0] = 42
    }
    fmt.Println(v1, v1)
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    if false {
        switch 32768 {
        case -9223372036854775808:
            v1 = (v1 * 1000)
        case 0:
            if false {
                if false {
                    for fi30 := 0; fi30 < 14; fi30++ {
                        var v31 float32 = 0.0
                        var v32 float64 = 0.1
                        switch fi30 {
                        case 32768:
                            for fi33 := 0; fi33 < 3; fi33++ {
                                v1 = 834
                                _ = fi33
                            }
                        }
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                    delete(v3, "a")
                    var v34 uint = 65535
                    _ = v34
                } else {
                    for fi35 := 0; fi35 < 4; fi35++ {
                        if (!(true && false)) {
                            v36 := GenS1{F0: 0, F1: 306}
                            _ = v36
                        } else {
                            v4 += -739
                            fmt.Println(v4, v2, v1)
                        }
                        switch ("a" + "a") {
                        case "hello":
                            if false {
                                v24 -= 584
                                v37 := GenS5{F0: 100, F1: 1}
                                _ = v37
                            } else {
                                var v38 uint64 = 1
                                _ = v38
                            }
                        }
                        var v39 float32 = (float32(0.1) / 1.0)
                        _ = fi35
                        _ = v39
                    }
                    if true {
                        for fi40 := 0; fi40 < 18; fi40++ {
                            v2 = v2
                            _ = fi40
                        }
                    }
                }
                for fi41 := 0; fi41 < 12; fi41++ {
                    if false {
                        if (!true) {
                            v24 -= (^(v24))
                            v42 := GenS6{F0: "\n\t"}
                            _ = v42
                        } else {
                            fmt.Println(v24, v24)
                            fmt.Println(v1)
                        }
                        v1 = (v1 ^ -32768)
                    }
                    if (!(1 == 2)) {
                        v22 = append(v22, (fi41 % ((604) | 1)))
                    }
                    switch fi41 {
                    case 0:
                        fmt.Println(v1, v4)
                    case 32767:
                        switch fi41 {
                        case 32768:
                            v3[("foo bar" + "foo bar")] = (v24 << ((416) & 63))
                        case -1:
                            v43 := GenS5{F0: 100, F1: 255}
                            _ = v43
                        case -770:
                            v44 := GenS1{F0: (v24 >> ((0) & 63)), F1: 255}
                            _ = v44
                        }
                    case 475:
                        v45 := GenS4{F0: "hello", F1: -1, F2: v24}
                        _ = v45
                    }
                    _ = fi41
                }
            } else {
                v46 := v3[("" + ("hello" + "a"))]
                _ = v46
            }
        case 320:
            var v47 int64 = genFn0(1, 827)
            _ = v47
        }
    } else {
        fmt.Println(v4)
        var v48 bool = true
        fmt.Println(v2, v4)
        _ = v48
    }
    delete(v3, "foo bar")
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
