package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 uint8 }

func main() {
    var v1 string = "hello"
    var v2 uint = 943
    var v3 float64 = float64(1e10)
    v4 := []string{}
    switch 557 {
    case 42:
        var v5 uint = 1
        _ = v5
    case -370:
        if false {
            fmt.Println(v1, v1, v2)
        }
    }
    switch -9223372036854775808 {
    case 42:
        for fi6 := 0; fi6 < 1; fi6++ {
            v7 := GenS1{F0: -32768, F1: -2147483648, F2: 1000}
            for fi8 := 0; fi8 < 18; fi8++ {
                var v9 int32 = 0
                v2 -= v2
                v9 -= v9
                if false {
                    fmt.Println(fi6, fi8)
                    var v10 int16 = int16(v9)
                    _ = v10
                } else {
                    var v11 int16 = int16(v9)
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    case -1:
        switch "foo bar" {
        case "a":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "hello":
            v1 = (v1 + ((v1 + "foo bar") + ((((v1 + ("a" + v1)) + (((v1 + v1) + "hello") + v1)) + v1) + v1)))
        default:
            if (100 != 0) {
                if len(v4) > 0 {
                    v4[0] = v1
                }
            } else {
                v1 = ("hello" + "hello")
            }
        }
    case 1:
        if false {
            for fi12 := 0; fi12 < 18; fi12++ {
                switch 1 {
                case -1:
                    v13 := GenS1{F0: -32768, F1: 2147483647, F2: -61}
                    _ = v13
                default:
                    v14 := GenS2{F0: 1}
                    _ = v14
                }
                _ = fi12
            }
        }
    }
    v1 = "a"
    v2 -= v2
    var v15 string = ""
    v3 += v3
    v15 = (v15 + (v1 + v15))
    for fi16 := 0; fi16 < 9; fi16++ {
        v3 += (((v3 + (0.0 - float64((float32(1.5) * 0.0)))) / 1.5) + v3)
        if (0 > 255) {
            v15 = (v15 + v1)
            if false {
                if false {
                    for fi17 := 0; fi17 < 5; fi17++ {
                        if (!(((!(2 == 32767)) && false) || (-480 >= -1))) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        } else {
                            switch v1 {
                            case "\n\t":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            default:
                                v4 = append(v4, v1)
                            }
                        }
                        v18 := GenS2{F0: -163}
                        var v19 uint64 = 18446744073709551615
                        v1 = "hello"
                        if (false && (fi16 > (fi17 | -952))) {
                            fmt.Println(v2, v2)
                            var v20 string = ("" + "a")
                            _ = v20
                        }
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                }
                switch fi16 {
                case 196:
                    switch v1 {
                    case "foo bar":
                        var v21 int16 = 32767
                        _ = v21
                    default:
                        v22 := GenS3{F0: (v15 + "a"), F1: 42}
                        _ = v22
                    }
                case 369:
                    v15 = v15
                }
            }
        } else {
            var v23 bool = true
            v2 = v2
            _ = v23
        }
        _ = fi16
    }
    var v24 int8 = -1
    var v25 float64 = (v3 + v3)
    var v26 float64 = (0.1 * (v3 / (1.5 - (v3 / v3))))
    fmt.Println(v26, v3)
    var v27 uint32 = 823
    var v28 uint64 = 42
    var v29 uint32 = (v27 % ((uint32(v28)) | 1))
    var v30 int32 = -2147483648
    if true {
        for fi31 := 0; fi31 < 18; fi31++ {
            var v32 int = (fi31 << ((v2) & 63))
            for fi33 := 0; fi33 < 6; fi33++ {
                if (9223372036854775807 > 42) {
                    fmt.Println(v15, v27, v26)
                    fmt.Println(v1, v15)
                } else {
                    var v34 uint = 96
                    _ = v34
                }
                var v35 float64 = -2.25
                _ = fi33
                _ = v35
            }
            var v36 int64 = 1
            _ = fi31
            _ = v32
            _ = v36
        }
    } else {
        if (v3 > float64(1e10)) {
            fmt.Println(v25)
        }
        var v37 int32 = (v30 / ((int32(v28)) | 1))
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
