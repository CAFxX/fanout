package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16 }

func main() {
    var v1 uint8 = 127
    var v2 int32 = 692
    var v3 int32 = v2
    v4 := []string{"", ""}
    var v5 float32 = (1.5 + 0.1)
    var v6 uint = 481
    if len(v4) > 0 {
        v4[0] = ""
    }
    switch 32768 {
    case -826:
        v5 = (v5 * (v5 - (v5 * ((v5 * float32(-2.25)) * 0.0))))
    case -462:
        var v7 int8 = 28
        _ = v7
    default:
        v1 = (v1 | 1)
    }
    v8 := GenS1{F0: (^(v1))}
    if false {
        v9 := GenS2{F0: 4294967295, F1: true}
        if false {
            var v10 float32 = float32((-2.25 * (0.1 / (0.0 / ((1e10 * float64(1e10)) + (float64(v5) / (0.0 + ((1.5 + 1.5) - 3.14159))))))))
            var v11 string = ""
            _ = v10
            _ = v11
        }
        v2 += v2
        _ = v9
    } else {
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
        switch "\n\t" {
        case "hello":
            fmt.Println(v3, v2)
        }
    }
    var v12 uint64 = 216
    v13 := GenS3{F0: -1, F1: "hello"}
    v14 := GenS4{F0: false, F1: 2147483647}
    if len(v4) > 0 {
        v4[0] = "a"
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        switch "a" {
        case "a":
            v16 := GenS5{F0: int16(v6)}
            _ = v16
        case "hello":
            if true {
                if (!true) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v5 -= v5
                } else {
                    v17 := GenS4{F0: true, F1: 146}
                    _ = v17
                }
                switch fi15 {
                case 2147483647:
                    switch ("a" + ("hello" + "\n\t")) {
                    case "hello":
                        if len(v4) > 0 {
                            v4[0] = "hello"
                        }
                    }
                default:
                    v18 := GenS5{F0: -32768}
                    _ = v18
                }
            }
        default:
            v3 -= -2147483648
        }
        _ = fi15
    }
    v19 := GenS3{F0: -1, F1: "hello"}
    v20 := GenS6{F0: 480}
    var v21 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
