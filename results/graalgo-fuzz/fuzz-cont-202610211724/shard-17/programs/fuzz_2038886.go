package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 int = -1
    var v2 float64 = -2.25
    var v3 int16 = -418
    v4 := []uint32{99, 160, 65535}
    if false {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v5 := GenS1{F0: 9223372036854775807, F1: 960, F2: 255}
        _ = v5
    }
    v2 = ((((v2 * float64(0.1)) * v2) - (float64(((1.5 + float32(v2)) * (-2.25 * 0.0))) - v2)) - (float64(0.0) - 1e10))
    for fi6 := 0; fi6 < 14; fi6++ {
        if false {
            fmt.Println(v3, v3)
            var v7 uint64 = 961
            switch -31 {
            case -2147483648:
                for fi8 := 0; fi8 < 4; fi8++ {
                    if len(v4) > 0 {
                        v4[0] = 503
                    }
                    _ = fi8
                }
            case 128:
                var v9 float64 = v2
                _ = v9
            }
            _ = v7
        } else {
            var v10 string = ("foo bar" + "a")
            switch (fi6 + (v1 >> ((65535) & 63))) {
            case 358:
                var v11 int32 = 1
                _ = v11
            }
            v3 = v3
            _ = v10
        }
        var v12 uint8 = 1
        switch ("\n\t" + "\n\t") {
        case "foo bar":
            var v13 int = (fi6 / ((2147483647) | 1))
            _ = v13
        case "\n\t":
            fmt.Println(v12, v3, v12)
        case "":
            fmt.Println(v12, fi6)
        default:
            v2 -= v2
        }
        _ = fi6
        _ = v12
    }
    var v14 string = ("foo bar" + "")
    v3 -= (-(v3))
    for fi15 := 0; fi15 < 4; fi15++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v2 = v2
        _ = fi15
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        v2 = v2
        v14 = (((v14 + v14) + v14) + ((("a" + v14) + "hello") + (v14 + (v14 + "foo bar"))))
        v17 := GenS1{F0: 1, F1: 575, F2: 65535}
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: 790, F1: 2, F2: 1}
    v1 += v1
    v14 = ((v14 + v14) + "")
    var v19 int64 = 311
    if false {
        v3 += (v3 & (v3 + (v3 | v3)))
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        if (!(v14 <= v14)) {
            v21 := GenS2{F0: 1}
            var v22 uint32 = 429
            for fi23 := 0; fi23 < 2; fi23++ {
                v22 = 901
                if false {
                    if false {
                        var v24 float32 = 1.5
                        v14 = "foo bar"
                        var v25 int8 = -1
                        _ = v24
                        _ = v25
                    } else {
                        var v26 int16 = v3
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        fmt.Println(v14, fi23)
                        _ = v26
                    }
                } else {
                    var v27 bool = (937 > 22)
                    v14 = ((v14 + "foo bar") + ("a" + v14))
                    v3 += v3
                    _ = v27
                }
                _ = fi23
            }
            _ = v21
            _ = v22
        }
        _ = len(v4)
        var v28 uint8 = 42
        _ = fi20
        _ = v28
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        for fi30 := 0; fi30 < 20; fi30++ {
            switch v14 {
            case "\n\t":
                if true {
                    if (false || true) {
                        var v31 float32 = -2.25
                        v2 = (v2 - v2)
                        fmt.Println(v14, fi29)
                        _ = v31
                    } else {
                        fmt.Println(v3, v14)
                    }
                    var v32 int64 = (v19 ^ 42)
                    _ = v32
                } else {
                    v33 := GenS1{F0: 42, F1: 2147483648, F2: 255}
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v34 int8 = -104
                    _ = v33
                    _ = v34
                }
                v35 := GenS1{F0: (v19 / ((v19) | 1)), F1: 65535, F2: 42}
                _ = v35
            case "hello":
                v19 = (^(v19))
            default:
                var v36 bool = true
                _ = v36
            }
            _ = fi30
        }
        if len(v4) > 0 {
            v4[0] = 673
        }
        _ = fi29
    }
    var v37 float32 = 0.0
    v38 := GenS2{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}
