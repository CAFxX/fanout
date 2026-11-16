package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 int32; F2 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32 }

func main() {
    var v1 float32 = 1.5
    var v2 int32 = -748
    v3 := []string{"a", "foo bar", "", ("hello" + ((("foo bar" + ("a" + "")) + "hello") + ("\n\t" + "foo bar")))}
    var v4 int64 = -1
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    v5 := GenS1{F0: 42, F1: "a", F2: 100}
    fmt.Println(v4, v4)
    v1 = ((((v1 * (((((v1 / ((1.5 - v1) + (1.5 * v1))) - v1) / 1.0) / (((1e10 * v1) - 1e10) - (1.5 / v1))) + -2.25)) + 1.5) - v1) / (1.5 * ((1.5 * float32(0.1)) + (v1 * (-2.25 - 0.1)))))
    v6 := GenS2{F0: false}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = v2
    fmt.Println(v4, v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (true || false) {
        if true {
            switch "hello" {
            case "\n\t":
                var v7 uint32 = 2147483647
                _ = v7
            case "a":
                if len(v3) > 0 {
                    v3[0] = "foo bar"
                }
            case "hello":
                var v8 uint16 = 1
                _ = v8
            default:
                if (!(1 != 42)) {
                    switch 0 {
                    case 0:
                        v9 := GenS2{F0: false}
                        _ = v9
                    default:
                        var v10 uint32 = 2147483647
                        _ = v10
                    }
                    fmt.Println(v4, v1, v2)
                    switch 32768 {
                    case 614:
                        v11 := GenS3{F0: 32767, F1: v2, F2: -392}
                        _ = v11
                    default:
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                    }
                }
            }
            var v12 int16 = 1
            if (v1 != 0.0) {
                var v13 uint64 = 85
                _ = len(v3)
                v14 := GenS2{F0: (!(("" == "") && false))}
                _ = v13
                _ = v14
            } else {
                fmt.Println(v1)
            }
            _ = v12
        }
        _ = len(v3)
    } else {
        fmt.Println(v4, v4, v2)
        if len(v3) > 0 {
            v3[0] = ("hello" + ("" + "hello"))
        }
    }
    fmt.Println(v4, v2, v4)
    for fi15 := 0; fi15 < 12; fi15++ {
        v16 := GenS4{F0: 127}
        if true {
            switch "\n\t" {
            case "\n\t":
                v17 := GenS4{F0: 128}
                _ = v17
            case "hello":
                v18 := GenS4{F0: 1}
                _ = v18
            case "a":
                fmt.Println(v1)
            }
        }
        v19 := GenS5{F0: (v2 / ((v2) | 1))}
        _ = fi15
        _ = v16
        _ = v19
    }
    var v20 int = -2147483648
    switch "hello" {
    case "":
        v21 := GenS5{F0: (v2 & -70)}
        _ = v21
    }
    v22 := GenS2{F0: false}
    if true {
        var v23 int64 = (v4 >> ((223) & 63))
        v2 = (v2 / ((v2) | 1))
        fmt.Println(v2, v4)
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
