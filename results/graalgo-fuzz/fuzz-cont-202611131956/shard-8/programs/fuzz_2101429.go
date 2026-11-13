package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int32 }
type GenS2 struct { F0 uint16; F1 uint; F2 int32 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 int32 = 189
    v2 := []uint32{58, 2147483647}
    var v3 float64 = ((((0.0 + ((1.5 * 0.0) / 1.0)) - float64(((3.14159 - ((3.14159 / ((((((1e10 * (0.1 + -2.25)) - (float32(0.0) / 3.14159)) / 1.0) + 1e10) + (3.14159 + -2.25)) / 0.1)) - (1.5 / 1.0))) + 0.0))) + 1e10) + 1.5)
    v4 := []string{"hello"}
    var v5 uint = 127
    var v6 bool = ((!(1 > 988)) || (!true))
    if (-1 < 0) {
        var v7 int32 = -652
        var v8 int32 = -482
        _ = v7
        _ = v8
    }
    v3 += (v3 + (v3 + (((1e10 - 1.5) + float64(1e10)) / v3)))
    switch 127 {
    case 127:
        switch ("hello" + ("a" + "")) {
        case "a":
            var v9 int16 = 1000
            _ = v9
        case "hello":
            if (v5 != v5) {
                for fi10 := 0; fi10 < 7; fi10++ {
                    var v11 uint = (v5 - (^(v5)))
                    _ = fi10
                    _ = v11
                }
                var v12 uint16 = 0
                v13 := GenS1{F0: 1000, F1: -377, F2: (v1 * 42)}
                _ = v12
                _ = v13
            }
        case "foo bar":
            v5 = (v5 - v5)
        }
    default:
        v3 = 0.1
    }
    v6 = v6
    if v6 {
        for fi14 := 0; fi14 < 17; fi14++ {
            fmt.Println(v6, v3, v5)
            _ = fi14
        }
        _ = len(v2)
    } else {
        for fi15 := 0; fi15 < 19; fi15++ {
            for fi16 := 0; fi16 < 18; fi16++ {
                var v17 int8 = 127
                var v18 float64 = (((-2.25 - (v3 - v3)) - v3) + v3)
                _ = fi16
                _ = v17
                _ = v18
            }
            var v19 int = fi15
            v20 := GenS2{F0: 0, F1: v5, F2: (^(v1))}
            _ = fi15
            _ = v19
            _ = v20
        }
    }
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    switch -1 {
    case 42:
        for fi21 := 0; fi21 < 8; fi21++ {
            v5 = v5
            v22 := GenS3{F0: 0}
            v3 = v3
            _ = fi21
            _ = v22
        }
    default:
        fmt.Println(v5, v5)
    }
    v6 = (598 >= 1000)
    v3 = v3
    fmt.Println(v1, v5, v6)
    if v6 {
        if v6 {
            v3 = 0.0
        } else {
            if true {
                v6 = false
            } else {
                v1 += (v1 ^ 95)
            }
            v2 = append(v2, 37)
            if false {
                for fi23 := 0; fi23 < 7; fi23++ {
                    v5 += v5
                    v4 = append(v4, ("hello" + ((("hello" + "foo bar") + "\n\t") + "")))
                    _ = fi23
                }
                var v24 float32 = ((-2.25 * (0.1 / (((1e10 - 0.0) + -2.25) * 3.14159))) - 1e10)
                for fi25 := 0; fi25 < 5; fi25++ {
                    var v26 bool = (v6 && (!(!false)))
                    var v27 uint16 = 0
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                _ = v24
            } else {
                v1 = 2147483647
                switch -9223372036854775808 {
                case 941:
                    if ((3.14159 / float32(v3)) < 1e10) {
                        var v28 bool = false
                        v29 := GenS3{F0: 255}
                        switch 9223372036854775807 {
                        case 4294967295:
                            fmt.Println(v3, v28)
                        case -362:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        default:
                            v30 := GenS2{F0: 65535, F1: (v5 * (v5 / (((v5 + 14)) | 1))), F2: v1}
                            _ = v30
                        }
                        _ = v28
                        _ = v29
                    }
                case -859:
                    v2 = append(v2, 2147483648)
                }
            }
        }
        v31 := GenS1{F0: 712, F1: 1000, F2: v1}
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
