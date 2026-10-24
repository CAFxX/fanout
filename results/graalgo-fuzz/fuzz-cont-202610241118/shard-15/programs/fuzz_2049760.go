package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int32 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 uint16; F1 uint64; F2 string }

func main() {
    v1 := []string{"\n\t", "hello", "a", ""}
    v2 := make(map[string]string)
    var v3 uint8 = 255
    if ((62 != 2) && ((((-128 != 127) && ((0.0 * ((0.1 - float64(0.1)) * 0.1)) != 0.0)) || false) && (916 != -737))) {
        var v4 bool = (65535 != 1)
        _ = v4
    }
    fmt.Println(v3)
    fmt.Println(v3)
    switch 1 {
    case -619:
        switch 32768 {
        case -894:
            v1 = append(v1, "foo bar")
        case -9223372036854775808:
            v5 := GenS1{F0: 1, F1: 321, F2: 614}
            _ = v5
        default:
            v6 := GenS2{F0: false, F1: 255}
            _ = v6
        }
    case 42:
        v7 := GenS2{F0: (!true), F1: 0}
        _ = v7
    }
    v3 = 33
    if (65535 >= 255) {
        if (42 < -1) {
            v3 -= (v3 + (v3 >> ((4294967295) & 63)))
        }
        v8 := GenS3{F0: 0, F1: 94}
        _ = v8
    }
    v3 = (v3 * (v3 * v3))
    var v9 int64 = -1
    v9 = v9
    var v10 int8 = -100
    switch "a" {
    case "a":
        var v11 float64 = 0.1
        _ = v11
    case "\n\t":
        v9 = v9
    case "foo bar":
        if len(v1) > 0 {
            v1[0] = "a"
        }
    default:
        v12 := GenS4{F0: 985, F1: 813, F2: "a"}
        _ = v12
    }
    v9 = v9
    var v13 int64 = -9223372036854775808
    if false {
        var v14 uint8 = v3
        var v15 int = 9223372036854775807
        for fi16 := 0; fi16 < 17; fi16++ {
            if true {
                var v17 uint = 255
                for fi18 := 0; fi18 < 4; fi18++ {
                    var v19 uint16 = 950
                    _ = fi18
                    _ = v19
                }
                for fi20 := 0; fi20 < 8; fi20++ {
                    switch ("\n\t" + "\n\t") {
                    case "hello":
                        var v21 string = ""
                        _ = v21
                    case "\n\t":
                        if ((v17 - (v17 % (((v17 * 42)) | 1))) != v17) {
                            fmt.Println(v15, v9)
                        }
                    }
                    fmt.Println(v14, v3, v14)
                    var v22 int = -623
                    _ = fi20
                    _ = v22
                }
                _ = v17
            } else {
                v15 = (fi16 % ((v15) | 1))
                switch "\n\t" {
                case "":
                    var v23 int = -193
                    _ = v23
                case "\n\t":
                    var v24 uint32 = 70
                    _ = v24
                case "a":
                    switch v15 {
                    case 32768:
                        var v25 int = v15
                        _ = v25
                    case -247:
                        v14 += v3
                    case 460:
                        var v26 float32 = (0.0 - (((1.5 - (((0.1 * 3.14159) * 1e10) / (0.1 + -2.25))) / 1.5) - (-2.25 + 0.0)))
                        _ = v26
                    }
                }
                var v27 uint16 = 731
                _ = v27
            }
            switch 204 {
            case 52:
                _ = len(v2)
            case -983:
                v10 = -1
            }
            v9 += v13
            _ = fi16
        }
        _ = v14
        _ = v15
    }
    switch "a" {
    case "a":
        v1 = append(v1, "")
    default:
        v10 += (v10 - -94)
    }
    var v28 uint8 = (^(v3))
    v13 -= v13
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v28
    _ = v3
    _ = v9
}
