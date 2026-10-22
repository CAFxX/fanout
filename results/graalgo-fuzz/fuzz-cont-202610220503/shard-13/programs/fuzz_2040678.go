package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = (1.5 * (3.14159 * ((-2.25 / 1.5) / 1e10)))
    v3 := []int{32768, 1, 42, 9223372036854775807}
    v1[42] = "foo bar"
    v4 := GenS1{F0: 1, F1: 4294967295}
    var v5 int16 = 32767
    v2 += (v2 / (v2 + ((v2 + (3.14159 + (-2.25 / 3.14159))) * float32((3.14159 - (0.0 / (1.5 * 1e10)))))))
    v2 = v2
    switch -2147483648 {
    case -363:
        switch "\n\t" {
        case "hello":
            if (103 == 507) {
                for fi6 := 0; fi6 < 6; fi6++ {
                    var v7 int = (fi6 + -2147483648)
                    v8 := GenS2{F0: "foo bar", F1: 2}
                    fmt.Println(v7, v2, v5)
                    _ = fi6
                    _ = v7
                    _ = v8
                }
                var v9 string = "foo bar"
                switch (("\n\t" + v9) + "a") {
                case "\n\t":
                    switch 43 {
                    case -1:
                        v9 = v9
                    case 308:
                        v10 := GenS3{F0: 728, F1: 194}
                        _ = v10
                    case 128:
                        v11 := GenS2{F0: ("" + v9), F1: 1}
                        _ = v11
                    }
                case "":
                    if (true && false) {
                        var v12 int32 = 42
                        _ = v12
                    } else {
                        if ((212 < 118) && (((-2.25 * -2.25) * float64(0.1)) == (0.1 - 1e10))) {
                            var v13 int8 = -122
                            delete(v1, 2147483648)
                            if len(v3) > 0 {
                                v3[0] = 4294967295
                            }
                            _ = v13
                        }
                        for fi14 := 0; fi14 < 15; fi14++ {
                            fmt.Println(v2, v2, v2)
                            v1[2] = ""
                            _ = fi14
                        }
                        fmt.Println(v5, v9)
                    }
                case "hello":
                    v5 -= (v5 << ((4294967295) & 63))
                default:
                    var v15 float32 = 0.0
                    _ = v15
                }
                _ = v9
            }
        case "\n\t":
            v16 := GenS1{F0: -1, F1: 127}
            _ = v16
        case "":
            _ = len(v1)
        }
    case -9223372036854775808:
        v2 = float32(1.5)
    case -2147483648:
        for fi17 := 0; fi17 < 1; fi17++ {
            switch "" {
            case "\n\t":
                var v18 string = "\n\t"
                _ = v18
            case "a":
                v5 += (-(v5))
            default:
                fmt.Println(fi17, fi17)
            }
            switch "\n\t" {
            case "\n\t":
                fmt.Println(fi17, v2, v5)
            case "foo bar":
                var v19 int8 = 42
                _ = v19
            }
            _ = fi17
        }
    default:
        v2 -= float32(-2.25)
    }
    v20 := GenS2{F0: "\n\t", F1: 127}
    var v21 uint16 = 42
    if true {
        if true {
            fmt.Println(v5)
        } else {
            for fi22 := 0; fi22 < 20; fi22++ {
                for fi23 := 0; fi23 < 4; fi23++ {
                    v5 = v5
                    _ = fi23
                }
                _ = fi22
            }
        }
        v2 = float32(((float64(v2) * -2.25) * 1e10))
    }
    var v24 int8 = 105
    for fi25 := 0; fi25 < 8; fi25++ {
        if false {
            var v26 string = "a"
            v2 = v2
            v5 = 0
            _ = v26
        }
        v2 += v2
        _ = fi25
    }
    for fi27 := 0; fi27 < 4; fi27++ {
        if true {
            switch "\n\t" {
            case "":
                for fi28 := 0; fi28 < 7; fi28++ {
                    v29 := GenS4{F0: v5}
                    v24 -= v24
                    _ = fi28
                    _ = v29
                }
            }
            for fi30 := 0; fi30 < 14; fi30++ {
                if ("" < ("foo bar" + "\n\t")) {
                    if (false && ("" == "foo bar")) {
                        v31 := v1[1]
                        switch (fi27 + fi27) {
                        case -99:
                            v32 := GenS2{F0: v31, F1: 1}
                            _ = v32
                        }
                        var v33 float64 = float64(v2)
                        _ = v31
                        _ = v33
                    }
                } else {
                    var v34 int64 = -9223372036854775808
                    switch "" {
                    case "foo bar":
                        var v35 uint16 = 0
                        _ = v35
                    case "":
                        v2 -= v2
                    case "\n\t":
                        fmt.Println(fi30, fi30, v21)
                        v21 -= (v21 ^ (^(v21)))
                        _ = len(v1)
                    default:
                        var v36 float32 = (float32(-2.25) - 1.5)
                        _ = v36
                    }
                    _ = v34
                }
                _ = fi30
            }
        }
        _ = fi27
    }
    v2 = ((-2.25 + float32((1.5 * 1e10))) / v2)
    if len(v3) > 0 {
        v3[0] = 477
    }
    var v37 float32 = v2
    if true {
        if true {
            var v38 uint16 = (v21 & 739)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v38
        }
        var v39 float32 = v37
        fmt.Println(v2, v21, v5)
        _ = v39
    } else {
        var v40 int32 = 2147483647
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
