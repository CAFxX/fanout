package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int16; F2 int8 }
type GenS3 struct { F0 bool; F1 bool; F2 int64 }
func genFn0(p0 int8) uint64 {
    v5 := []int{-9223372036854775808, -883, 636}
    switch "foo bar" {
    case "":
        if true {
            var v6 int = 42
            _ = len(v5)
            _ = v6
        } else {
            fmt.Println(p0)
            for fi7 := 0; fi7 < 2; fi7++ {
                for fi8 := 0; fi8 < 2; fi8++ {
                    var v9 uint8 = 255
                    if len(v5) > 0 {
                        v5[0] = 42
                    }
                    if false {
                        if false {
                            _ = len(v5)
                            if (!true) {
                                var v10 int64 = 9223372036854775807
                                p0 += 42
                                _ = v10
                            }
                        }
                        v11 := GenS1{F0: 128}
                        var v12 bool = false
                        var v13 int16 = 32767
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                    _ = fi8
                    _ = v9
                }
                _ = fi7
            }
            if false {
                v14 := GenS2{F0: 9223372036854775807, F1: 1000, F2: p0}
                fmt.Println(p0)
                _ = v14
            }
        }
    case "foo bar":
        for fi15 := 0; fi15 < 19; fi15++ {
            for fi16 := 0; fi16 < 11; fi16++ {
                var v17 int16 = 651
                switch (fi16 % (((^(fi15))) | 1)) {
                case -602:
                    switch (fi15 + fi15) {
                    case 128:
                        var v18 uint = 996
                        _ = v18
                    default:
                        if ((((((0.1 - (1.5 / 1.0)) + 1.5) / 1e10) * float32((float64(1e10) - 0.0))) * ((1e10 - 0.0) * (1.5 * (3.14159 - ((((1e10 - (1e10 / (3.14159 + float32((-2.25 / (-2.25 * (1e10 * -2.25))))))) / 3.14159) - 0.1) + (float32((float64((0.1 + 1.5)) + ((3.14159 - 1.5) + float64(0.1)))) + (float32(0.1) + (3.14159 + (float32(float64(1.5)) / float32(0.1)))))))))) <= 1.5) {
                            fmt.Println(p0)
                            if ((true && (float64((((1.5 / ((-2.25 + 1.5) / 1.0)) + 1.5) * -2.25)) <= 1.5)) || true) {
                                var v19 int64 = 599
                                _ = v19
                            }
                        } else {
                            v20 := GenS1{F0: 255}
                            p0 = (p0 / ((p0) | 1))
                            switch "" {
                            case "foo bar":
                                if len(v5) > 0 {
                                    v5[0] = fi16
                                }
                            default:
                                v17 = v17
                            }
                            _ = v20
                        }
                    }
                case -708:
                    _ = len(v5)
                default:
                    var v21 string = (("foo bar" + "hello") + "")
                    _ = v21
                }
                v22 := GenS1{F0: 1}
                _ = fi16
                _ = v17
                _ = v22
            }
            var v23 int8 = (p0 * 100)
            if false {
                var v24 uint16 = 32767
                _ = v24
            }
            _ = fi15
            _ = v23
        }
    case "hello":
        var v25 string = "foo bar"
        _ = v25
    }
    _ = p0
    _ = v5
    return 2
}

func main() {
    var v1 uint8 = 128
    var v2 uint64 = 42
    v3 := make(map[uint32]uint32)
    v4 := []uint32{}
    var v26 uint = 206
    for fi27 := 0; fi27 < 7; fi27++ {
        v2 += (v2 ^ (v2 | v2))
        switch (("\n\t" + "\n\t") + "hello") {
        case "\n\t":
            var v28 float64 = (-2.25 - (0.0 / (float64((0.1 / float32(0.1))) - float64(-2.25))))
            _ = v28
        case "hello":
            var v29 bool = (346 <= 639)
            _ = v29
        case "a":
            if (("a" + "\n\t") > (("\n\t" + "hello") + "\n\t")) {
                v2 -= (v2 % ((genFn0(0)) | 1))
                v30 := GenS2{F0: -1, F1: -504, F2: -1}
                var v31 uint16 = 255
                _ = v30
                _ = v31
            } else {
                for fi32 := 0; fi32 < 14; fi32++ {
                    v33 := GenS1{F0: 255}
                    switch ("hello" + "hello") {
                    case "foo bar":
                        switch "foo bar" {
                        case "\n\t":
                            v34 := GenS3{F0: (1.5 <= (3.14159 + 3.14159)), F1: true, F2: -107}
                            _ = v34
                        default:
                            if (-32768 <= -1) {
                                v2 = genFn0(0)
                                var v35 int32 = -1
                                _ = v35
                            } else {
                                var v36 float32 = -2.25
                                _ = v36
                            }
                            delete(v3, 65535)
                        }
                    default:
                        v26 -= 4294967295
                    }
                    v37 := v3[0]
                    var v38 uint16 = 0
                    _ = fi32
                    _ = v33
                    _ = v37
                    _ = v38
                }
                var v39 uint = v26
                _ = v39
            }
        default:
            fmt.Println(fi27)
        }
        var v40 float64 = 0.1
        _ = fi27
        _ = v40
    }
    if false {
        fmt.Println(v2, v2, v1)
        var v41 float32 = 1.5
        _ = v41
    } else {
        v42 := GenS1{F0: (v1 << (((v26 * v26)) & 63))}
        _ = v42
    }
    v2 = (v2 - genFn0(42))
    var v43 int16 = 42
    if ((("hello" + ("a" + "a")) < "\n\t") && ((("hello" + ("foo bar" + ("hello" + "foo bar"))) > "\n\t") || false)) {
        v43 = (v43 % ((v43) | 1))
        v43 = (v43 - 611)
        delete(v3, 2147483647)
    }
    _ = len(v3)
    fmt.Println(v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v43
}
