package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint }
type GenS2 struct { F0 bool; F1 int16; F2 int32 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int64; F1 int64; F2 uint16 }
type GenS5 struct { F0 uint8; F1 uint; F2 int }

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 42
    var v3 int = 1
    v4 := []string{("" + "\n\t"), ("hello" + "foo bar"), ""}
    var v5 int8 = -1
    var v6 uint16 = v2
    switch "foo bar" {
    case "a":
        v6 -= v6
    case "foo bar":
        if false {
            if (!true) {
                var v7 float64 = -2.25
                _ = v7
            }
            v8 := GenS1{F0: int(v5), F1: 42, F2: 42}
            v9 := GenS1{F0: (v3 + 32767), F1: 2, F2: 255}
            _ = v8
            _ = v9
        } else {
            switch "hello" {
            case "hello":
                fmt.Println(v6, v5)
            case "":
                if ((0 != 958) && (!(4294967295 > 2147483647))) {
                    for fi10 := 0; fi10 < 14; fi10++ {
                        switch v3 {
                        case 2147483647:
                            switch (("foo bar" + "") + "") {
                            case "\n\t":
                                switch "foo bar" {
                                case "foo bar":
                                    var v11 uint8 = 255
                                    _ = v11
                                case "\n\t":
                                    v12 := GenS1{F0: v3, F1: 2, F2: 0}
                                    _ = v12
                                default:
                                    delete(v1, (fi10 & v3))
                                }
                            case "foo bar":
                                v2 -= (v2 & 130)
                            }
                        case 4294967295:
                            delete(v1, 4294967295)
                        case -521:
                            var v13 int64 = 1
                            _ = v13
                        default:
                            fmt.Println(v3, v6)
                        }
                        _ = fi10
                    }
                } else {
                    var v14 uint16 = v6
                    switch v3 {
                    case 32767:
                        if true {
                            fmt.Println(v5, v6)
                            _ = len(v1)
                            v15 := GenS1{F0: int(v5), F1: 299, F2: 2}
                            _ = v15
                        }
                    default:
                        var v16 int8 = (v5 >> ((65535) & 63))
                        _ = v16
                    }
                    v17 := GenS2{F0: (!false), F1: 1000, F2: -1}
                    _ = v14
                    _ = v17
                }
            case "\n\t":
                if len(v4) > 0 {
                    v4[0] = ""
                }
            }
            if true {
                var v18 uint8 = 127
                _ = v18
            } else {
                v19 := v1[616]
                fmt.Println(v19, v2, v2)
                v20 := GenS3{F0: 2, F1: -568}
                _ = v19
                _ = v20
            }
        }
    }
    switch ("hello" + "foo bar") {
    case "hello":
        if (968 <= -1) {
            v6 += (v2 & (v2 % ((v2) | 1)))
        } else {
            v21 := GenS4{F0: 1, F1: -9223372036854775808, F2: (v2 ^ v2)}
            _ = v21
        }
    case "\n\t":
        if (2 > v6) {
            switch 4294967295 {
            case -17:
                var v22 int32 = 995
                _ = v22
            case -2147483648:
                if true {
                    var v23 uint8 = 1
                    v6 = v6
                    _ = v23
                } else {
                    fmt.Println(v3, v3)
                }
            case 32768:
                fmt.Println(v5, v2)
            default:
                fmt.Println(v2, v2, v6)
            }
            var v24 float64 = float64(1.5)
            _ = v24
        }
    default:
        var v25 float32 = 3.14159
        _ = v25
    }
    v26 := GenS5{F0: 0, F1: 4294967295, F2: 166}
    var v27 int32 = 0
    v27 -= 42
    var v28 uint = 255
    v5 = v5
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v29 float64 = 0.1
    for fi30 := 0; fi30 < 4; fi30++ {
        var v31 float64 = ((float64(((0.0 - 1e10) - 1.5)) - (v29 - float64((float32(v29) / 1e10)))) - (1.5 * (1e10 / (v29 + v29))))
        if len(v4) > 0 {
            v4[0] = "a"
        }
        v3 += fi30
        _ = fi30
        _ = v31
    }
    v32 := GenS5{F0: 2, F1: (v28 / (((v28 * v28)) | 1)), F2: (v3 + (v3 / (((v3 - (v3 ^ 282))) | 1)))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
