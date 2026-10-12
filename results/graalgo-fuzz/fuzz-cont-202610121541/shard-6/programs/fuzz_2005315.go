package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 int32; F2 uint }
func genFn0() int {
    var v4 int8 = 127
    var v5 float32 = -2.25
    v6 := GenS1{F0: 2147483648, F1: 6, F2: 42}
    for fi7 := 0; fi7 < 2; fi7++ {
        switch (-(fi7)) {
        case 128:
            switch ("" + "foo bar") {
            case "a":
                v8 := []int{}
                _ = v8
            }
        }
        switch fi7 {
        case 521:
            v9 := GenS1{F0: 2147483647, F1: (v4 | 100), F2: 1}
            _ = v9
        case 0:
            v4 -= v4
        }
        var v10 bool = (((-1 < 1) || false) && false)
        _ = fi7
        _ = v10
    }
    _ = v4
    _ = v5
    _ = v6
    return 127
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint = 2
    var v3 int8 = 0
    v11 := []int{9223372036854775807, 899, -76, int(v2)}
    switch ("" + "foo bar") {
    case "foo bar":
        var v12 int16 = -926
        _ = v12
    default:
        if ("hello" >= ("foo bar" + "hello")) {
            if (65535 != 1) {
                switch "" {
                case "foo bar":
                    var v13 uint64 = 4294967295
                    _ = v13
                case "":
                    for fi14 := 0; fi14 < 5; fi14++ {
                        v15 := GenS1{F0: 0, F1: 49, F2: 1000}
                        for fi16 := 0; fi16 < 15; fi16++ {
                            v17 := GenS2{F0: -1}
                            if ((v3 | (v3 ^ (v3 % ((v3) | 1)))) != -128) {
                                if (!(0 > 42)) {
                                    v3 -= (v3 + v3)
                                    if (1e10 < 1e10) {
                                        fmt.Println(fi14)
                                    } else {
                                        v2 += 0
                                        var v18 bool = ((fi16 >> (((v2 & v2)) & 63)) < (fi14 ^ fi16))
                                        v2 = 995
                                        _ = v18
                                    }
                                }
                            } else {
                                if ((("foo bar" + (("" + ("a" + "a")) + ("hello" + "foo bar"))) + "foo bar") < "") {
                                    v2 = (v2 % (((^(v2))) | 1))
                                    v19 := GenS1{F0: 636, F1: 116, F2: 1}
                                    v3 = v3
                                    _ = v19
                                } else {
                                    if (!false) {
                                        v2 -= v2
                                    } else {
                                        v3 += -1
                                        switch (((("" + ((("" + (((("a" + ("a" + "foo bar")) + "a") + "hello") + "foo bar")) + "") + "")) + ("foo bar" + ("foo bar" + ("" + ("hello" + ""))))) + "a") + "") {
                                        case "hello":
                                            var v20 float64 = (float64(0.1) / float64(3.14159))
                                            _ = v20
                                        case "":
                                            switch (fi14 | -1) {
                                            case 32768:
                                                v3 += (v3 % ((42) | 1))
                                            case 127:
                                                v2 += 930
                                            default:
                                                v21 := GenS2{F0: -2147483648}
                                                _ = v21
                                            }
                                        }
                                    }
                                    var v22 bool = false
                                    _ = v22
                                }
                            }
                            _ = fi16
                            _ = v17
                        }
                        _ = fi14
                        _ = v15
                    }
                default:
                    for fi23 := 0; fi23 < 7; fi23++ {
                        if len(v11) > 0 {
                            v11[0] = 1
                        }
                        _ = fi23
                    }
                }
            }
            switch "hello" {
            case "":
                fmt.Println(v3, v3)
            default:
                var v24 bool = (float64(1e10) >= ((-2.25 + (0.0 / 1.0)) / (1e10 / (0.1 - 3.14159))))
                _ = v24
            }
        } else {
            var v25 uint8 = 42
            v26 := GenS1{F0: 2, F1: -58, F2: 1000}
            _ = v25
            _ = v26
        }
    }
    v3 += (v3 / ((v3) | 1))
    var v27 int16 = 0
    if (68 != -622) {
        var v28 int8 = 42
        v27 = 32767
        var v29 uint = (v2 - (v2 % (((v2 & 255)) | 1)))
        _ = v28
        _ = v29
    } else {
        if true {
            var v30 float32 = (float32(float64((1.5 / 1.0))) + (float32((3.14159 - -2.25)) - -2.25))
            if len(v11) > 0 {
                v11[0] = 926
            }
            delete(v1, 65535)
            _ = v30
        } else {
            var v31 uint8 = 127
            _ = v31
        }
    }
    v2 = v2
    fmt.Println(v3, v3, v2)
    var v32 uint = v2
    if (-9223372036854775808 >= 322) {
        v3 = v3
        fmt.Println(v32, v27, v3)
        v33 := GenS2{F0: 42}
        _ = v33
    } else {
        for fi34 := 0; fi34 < 7; fi34++ {
            fmt.Println(v2, fi34, v2)
            var v35 int = -9223372036854775808
            _ = fi34
            _ = v35
        }
    }
    var v36 uint = 885
    v37 := GenS3{F0: v27, F1: -782, F2: (v32 + (v36 + (v32 + 2)))}
    v32 = (^(v2))
    fmt.Println(v32, v2, v2)
    fmt.Println(v32)
    v38 := GenS3{F0: v27, F1: -782, F2: v2}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v36
    _ = v37
    _ = v38
}
