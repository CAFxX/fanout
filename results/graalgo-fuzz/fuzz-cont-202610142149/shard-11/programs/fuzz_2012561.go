package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int; F1 int32; F2 uint32 }
func genFn0() int16 {
    var v6 string = "foo bar"
    var v7 int32 = -1
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := []string{"\n\t"}
    var v2 float64 = (1e10 - 0.1)
    var v3 uint8 = 42
    v4 := make(map[uint32]string)
    var v5 uint16 = 863
    v8 := GenS1{F0: (v5 ^ v5), F1: -384, F2: 18446744073709551615}
    v5 += (v5 >> ((2) & 63))
    v5 -= uint16(v3)
    switch -441 {
    case 4294967295:
        var v9 bool = (v5 > 34)
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 uint8 = 1
        v2 -= v2
        _ = fi10
        _ = v11
    }
    fmt.Println(v2)
    if false {
        var v12 uint = 690
        if (127 < 42) {
            v13 := GenS1{F0: v5, F1: -535, F2: (v12 - 672)}
            _ = v13
        }
        v12 += (^(v12))
        _ = v12
    }
    v5 -= 255
    switch "hello" {
    case "a":
        var v14 float64 = ((0.1 - v2) / float64(((1.5 - 1.5) - (3.14159 * (((3.14159 / 1e10) * ((3.14159 * (float32(v2) + -2.25)) - 0.0)) - 1e10)))))
        _ = v14
    default:
        v2 = (((v2 + v2) * v2) * (float64(((float32(3.14159) / 1.5) - 3.14159)) + 3.14159))
    }
    switch ("foo bar" + "") {
    case "a":
        switch "\n\t" {
        case "\n\t":
            v2 -= (((float64(float32(-2.25)) / v2) * (v2 - float64(0.1))) / v2)
        }
    case "hello":
        for fi15 := 0; fi15 < 7; fi15++ {
            switch "" {
            case "hello":
                var v16 uint8 = v3
                _ = v16
            case "foo bar":
                if false {
                    var v17 bool = false
                    _ = v17
                }
            case "\n\t":
                var v18 int32 = 2147483647
                _ = v18
            default:
                if (42 > 1) {
                    v3 = v3
                    v3 = v3
                }
            }
            _ = fi15
        }
    case "foo bar":
        var v19 bool = (301 > 18446744073709551615)
        _ = v19
    }
    var v20 uint8 = (v3 & (^(v3)))
    if len(v1) > 0 {
        v1[0] = (("" + "") + "hello")
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        if (757 <= 42) {
            if false {
                v22 := GenS1{F0: v5, F1: (fi21 / (((-(fi21))) | 1)), F2: 42}
                _ = v22
            } else {
                switch "" {
                case "a":
                    var v23 bool = true
                    _ = v23
                case "\n\t":
                    if false {
                        var v24 int64 = 680
                        _ = v24
                    } else {
                        var v25 int = fi21
                        var v26 int16 = 1000
                        v25 = fi21
                        _ = v25
                        _ = v26
                    }
                case "hello":
                    if (0 < 255) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if false {
                            switch (fi21 ^ fi21) {
                            case 4294967295:
                                var v27 float64 = v2
                                _ = v27
                            case 2147483647:
                                var v28 uint64 = 540
                                _ = v28
                            }
                            v29 := GenS2{F0: 0, F1: 32767}
                            _ = v29
                        } else {
                            switch "\n\t" {
                            case "foo bar":
                                switch "a" {
                                case "hello":
                                    var v30 float32 = 1e10
                                    _ = v30
                                case "a":
                                    v2 -= (v2 * v2)
                                default:
                                    var v31 int = 32767
                                    _ = v31
                                }
                            case "":
                                var v32 int16 = 672
                                _ = v32
                            case "a":
                                v33 := GenS3{F0: 4294967295, F1: 2147483647, F2: 65535}
                                _ = v33
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v34 bool = false
                            _ = v34
                        }
                    } else {
                        v35 := v4[4294967295]
                        var v36 string = v35
                        v37 := GenS2{F0: -1, F1: -32768}
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                default:
                    var v38 int16 = -1
                    _ = v38
                }
            }
        } else {
            var v39 int = 128
            fmt.Println(fi21, v20)
            _ = v39
        }
        for fi40 := 0; fi40 < 17; fi40++ {
            v4[497] = "hello"
            v5 = v5
            _ = fi40
        }
        _ = fi21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
