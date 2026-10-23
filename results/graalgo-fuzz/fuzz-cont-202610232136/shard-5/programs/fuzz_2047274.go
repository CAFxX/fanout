package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int8 }
type GenS2 struct { F0 int32; F1 int8; F2 bool }
type GenS3 struct { F0 uint8; F1 uint64 }

func main() {
    var v1 int32 = 42
    var v2 uint = 4294967295
    var v3 uint16 = 20
    v4 := []int{32767, -960, 133}
    if false {
        switch 32768 {
        case 9223372036854775807:
            v1 = v1
        case 32768:
            v5 := GenS1{F0: (v1 * v1), F1: 32767, F2: -18}
            _ = v5
        default:
            for fi6 := 0; fi6 < 17; fi6++ {
                var v7 int16 = -1
                _ = fi6
                _ = v7
            }
        }
    }
    switch "a" {
    case "hello":
        var v8 int = -9223372036854775808
        _ = v8
    case "\n\t":
        var v9 float32 = (-2.25 / 1.0)
        _ = v9
    }
    var v10 uint = (v2 << (((v2 - (v2 >> (((^(v2))) & 63)))) & 63))
    v1 -= -249
    v4 = append(v4, 0)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v11 bool = false
    if false {
        var v12 uint16 = v3
        switch 210 {
        case -645:
            for fi13 := 0; fi13 < 16; fi13++ {
                fmt.Println(v12)
                _ = fi13
            }
        }
        v14 := GenS1{F0: v1, F1: 2, F2: 2}
        _ = v12
        _ = v14
    } else {
        v11 = true
        switch "\n\t" {
        case "":
            var v15 int16 = -32768
            _ = v15
        case "a":
            fmt.Println(v11, v2)
        case "\n\t":
            fmt.Println(v1, v2, v1)
        }
        var v16 int64 = 42
        _ = v16
    }
    if false {
        switch 1 {
        case 0:
            switch "foo bar" {
            case "hello":
                for fi17 := 0; fi17 < 16; fi17++ {
                    if len(v4) > 0 {
                        v4[0] = (fi17 % ((32768) | 1))
                    }
                    _ = fi17
                }
            case "":
                v18 := GenS1{F0: v1, F1: 929, F2: 42}
                _ = v18
            default:
                switch -615 {
                case 598:
                    switch ("a" + "foo bar") {
                    case "hello":
                        v19 := GenS1{F0: v1, F1: 1, F2: 0}
                        _ = v19
                    }
                case -1:
                    var v20 int16 = -1
                    _ = v20
                default:
                    v4 = append(v4, 9223372036854775807)
                }
            }
        case -1:
            for fi21 := 0; fi21 < 3; fi21++ {
                if true {
                    fmt.Println(v10, v3, v3)
                    v22 := GenS1{F0: (v1 >> (((v2 % (((v2 & v10)) | 1))) & 63)), F1: (^(v3)), F2: 55}
                    if false {
                        switch -1 {
                        case 128:
                            v11 = (128 == 109)
                        default:
                            v4 = append(v4, 4294967295)
                        }
                        fmt.Println(fi21, v11)
                    }
                    _ = v22
                } else {
                    fmt.Println(fi21, v10, fi21)
                }
                _ = fi21
            }
        case -2147483648:
            v2 -= 255
        }
        if (((((("hello" + "hello") + "a") + (("" + ("foo bar" + "\n\t")) + "")) + ("hello" + ("a" + ("\n\t" + ("" + "foo bar"))))) + "a") == "a") {
            for fi23 := 0; fi23 < 8; fi23++ {
                fmt.Println(v1)
                for fi24 := 0; fi24 < 8; fi24++ {
                    switch "a" {
                    case "":
                        v4 = append(v4, fi23)
                    default:
                        v10 -= v10
                    }
                    v25 := GenS1{F0: (v1 | v1), F1: 1, F2: -80}
                    _ = fi24
                    _ = v25
                }
                var v26 int64 = 153
                _ = fi23
                _ = v26
            }
            v27 := GenS1{F0: (v1 << ((v2) & 63)), F1: 42, F2: 114}
            var v28 uint32 = 42
            _ = v27
            _ = v28
        }
    }
    if (1000 < 563) {
        v29 := GenS2{F0: 2147483647, F1: -77, F2: ((!v11) && v11)}
        var v30 int32 = (v1 / ((v1) | 1))
        _ = v29
        _ = v30
    }
    var v31 float64 = (0.0 + (3.14159 / (float64(float32(1e10)) * -2.25)))
    switch "\n\t" {
    case "":
        fmt.Println(v31, v11, v31)
    case "a":
        v2 = v2
    case "foo bar":
        _ = len(v4)
    default:
        if v11 {
            switch "\n\t" {
            case "\n\t":
                var v32 uint8 = 255
                _ = v32
            case "hello":
                var v33 uint64 = 0
                _ = v33
            default:
                v31 = ((v31 + ((1.5 + float64((3.14159 - 1e10))) - v31)) - (0.1 + v31))
            }
            v2 = v2
            if v11 {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v2 += 0
            } else {
                v31 -= (v31 / v31)
                v34 := GenS2{F0: v1, F1: -128, F2: (1 < uint8(v2))}
                var v35 int8 = -128
                _ = v34
                _ = v35
            }
        } else {
            var v36 int = -125
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 17; fi37++ {
        for fi38 := 0; fi38 < 3; fi38++ {
            if true {
                fmt.Println(v31, v3, v11)
                var v39 int32 = (v1 | v1)
                _ = v39
            }
            if (42 > 1) {
                v3 += v3
                var v40 float64 = 1.5
                if len(v4) > 0 {
                    v4[0] = (fi38 & (fi38 << ((v2) & 63)))
                }
                _ = v40
            }
            _ = len(v4)
            var v41 int16 = 32767
            _ = fi38
            _ = v41
        }
        _ = fi37
    }
    v10 = v2
    for fi42 := 0; fi42 < 7; fi42++ {
        v43 := GenS3{F0: 128, F1: 1}
        _ = fi42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v4
}
