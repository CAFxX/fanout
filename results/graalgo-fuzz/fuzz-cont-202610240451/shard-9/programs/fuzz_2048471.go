package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 int16; F2 bool }
type GenS6 struct { F0 int; F1 int8; F2 int64 }
type GenS7 struct { F0 uint; F1 uint32; F2 int64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = -2.25
    v3 := make(map[int]int)
    v4 := make(map[string]bool)
    for fi5 := 0; fi5 < 3; fi5++ {
        for fi6 := 0; fi6 < 2; fi6++ {
            switch ("foo bar" + "\n\t") {
            case "hello":
                v7 := GenS1{F0: ("\n\t" + "foo bar"), F1: -1, F2: 692}
                _ = v7
            default:
                v2 -= 1.5
            }
            v8 := GenS2{F0: 224}
            fmt.Println(fi6)
            _ = fi6
            _ = v8
        }
        switch ("hello" + "a") {
        case "hello":
            v2 = ((v2 / 1.0) / 0.1)
        }
        for fi9 := 0; fi9 < 2; fi9++ {
            fmt.Println(v2)
            _ = fi9
        }
        _ = fi5
    }
    if true {
        v2 += 1.5
        v10 := GenS2{F0: 352}
        for fi11 := 0; fi11 < 4; fi11++ {
            v12 := []int{}
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            for fi13 := 0; fi13 < 5; fi13++ {
                fmt.Println(fi13)
                var v14 uint8 = 127
                if true {
                    if ((((1 >= 100) || true) || true) || true) {
                        v3[(-(fi13))] = (fi13 ^ fi11)
                        var v15 int8 = -1
                        v16 := GenS3{F0: (v14 / ((v14) | 1))}
                        _ = v15
                        _ = v16
                    }
                    switch -1 {
                    case -213:
                        var v17 int16 = -481
                        _ = v17
                    case -698:
                        if (false && true) {
                            if (4294967295 >= 18446744073709551615) {
                                v2 = 3.14159
                            }
                            switch fi11 {
                            case 127:
                                var v18 string = "hello"
                                _ = v18
                            case 483:
                                v4["foo bar"] = (32767 <= 32767)
                            case 765:
                                var v19 uint = 4294967295
                                _ = v19
                            default:
                                delete(v3, fi11)
                                v14 = v14
                                var v20 uint32 = 2147483648
                                _ = v20
                            }
                        } else {
                            switch (fi11 * 914) {
                            case -1:
                                v21 := v1[-1]
                                _ = v21
                            case 127:
                                v22 := GenS1{F0: "a", F1: 2147483647, F2: 1}
                                _ = v22
                            case -640:
                                v23 := GenS1{F0: "hello", F1: 30, F2: 1}
                                _ = v23
                            default:
                                if ("foo bar" >= ("hello" + (("\n\t" + (("" + ("foo bar" + "hello")) + ("\n\t" + "hello"))) + "hello"))) {
                                    fmt.Println(v14)
                                } else {
                                    if (!false) {
                                        v2 -= v2
                                    }
                                    var v24 uint64 = uint64(v14)
                                    var v25 string = ""
                                    _ = v24
                                    _ = v25
                                }
                            }
                        }
                    case 1:
                        if false {
                            v26 := GenS4{F0: 65535}
                            _ = v26
                        }
                    }
                    if true {
                        if false {
                            fmt.Println(fi11, v2)
                            switch (^(fi13)) {
                            case 0:
                                var v27 int = 32768
                                _ = v27
                            case -288:
                                fmt.Println(v2, fi11, fi13)
                            case -9223372036854775808:
                                fmt.Println(fi11, v2, v2)
                            }
                            v2 = 0.1
                        }
                        var v28 uint16 = 0
                        switch fi11 {
                        case -986:
                            v2 -= (((v2 + (float32(-2.25) / v2)) * 1.5) / (0.0 / v2))
                        case -575:
                            switch fi13 {
                            case 128:
                                _ = len(v3)
                            case 25:
                                v29 := GenS2{F0: 2}
                                _ = v29
                            case -9223372036854775808:
                                fmt.Println(v28, v28, v14)
                            default:
                                if (-153 > (-(fi13))) {
                                    if true {
                                        var v30 int32 = 42
                                        delete(v4, ("\n\t" + "foo bar"))
                                        _ = v30
                                    } else {
                                        var v31 float64 = 0.0
                                        _ = v31
                                    }
                                    var v32 int16 = -32768
                                    var v33 uint8 = (^(v14))
                                    var v34 int8 = 100
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                } else {
                                    v3[(fi11 / ((fi11) | 1))] = 4294967295
                                }
                            }
                        }
                        var v35 string = "foo bar"
                        v14 -= 86
                        v36 := GenS5{F0: 0, F1: -32768, F2: (-734 >= -2147483648)}
                        var v37 bool = true
                        _ = v28
                        _ = v35
                        _ = v36
                        _ = v37
                    } else {
                        v2 -= v2
                    }
                }
                _ = fi13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    if true {
        v38 := GenS6{F0: -112, F1: 0, F2: 9223372036854775807}
        _ = v38
    }
    var v39 int64 = 0
    v2 -= v2
    _ = len(v4)
    var v40 int8 = 100
    v2 = (1e10 * v2)
    v41 := []int{}
    for fi42 := 0; fi42 < 11; fi42++ {
        v43 := GenS7{F0: 1, F1: 2, F2: v39}
        _ = fi42
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
