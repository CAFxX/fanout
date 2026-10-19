package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 uint8; F1 bool }
type GenS7 struct { F0 uint }

func main() {
    var v1 uint8 = 255
    v2 := []int{-9223372036854775808, 32767, -2147483648, 128}
    var v3 float32 = 0.0
    switch "a" {
    case "a":
        if true {
            fmt.Println(v1)
            v4 := GenS1{F0: ""}
            v5 := GenS1{F0: (("a" + ("\n\t" + "a")) + "foo bar")}
            _ = v4
            _ = v5
        } else {
            fmt.Println(v1, v1)
            for fi6 := 0; fi6 < 10; fi6++ {
                for fi7 := 0; fi7 < 14; fi7++ {
                    _ = len(v2)
                    var v8 uint = 613
                    fmt.Println(v8)
                    v9 := GenS2{F0: (v8 % ((v8) | 1)), F1: -614}
                    var v10 string = ("hello" + "\n\t")
                    _ = fi7
                    _ = v10
                    _ = v8
                    _ = v9
                }
                fmt.Println(fi6, v1, v3)
                v11 := GenS3{F0: true}
                _ = fi6
                _ = v11
            }
        }
    case "":
        var v12 uint16 = 890
        _ = v12
    case "foo bar":
        for fi13 := 0; fi13 < 4; fi13++ {
            var v14 uint64 = 0
            for fi15 := 0; fi15 < 13; fi15++ {
                if true {
                    switch "hello" {
                    case "a":
                        var v16 uint32 = 963
                        _ = v16
                    case "foo bar":
                        _ = len(v2)
                    case "hello":
                        v14 = (^(v14))
                    }
                    _ = len(v2)
                    v14 += v14
                }
                var v17 float32 = v3
                v14 -= 42
                _ = fi15
                _ = v17
            }
            for fi18 := 0; fi18 < 1; fi18++ {
                switch "a" {
                case "a":
                    var v19 int = fi18
                    var v20 float32 = 3.14159
                    _ = v19
                    _ = v20
                case "foo bar":
                    v21 := GenS3{F0: false}
                    _ = v21
                }
                switch (("\n\t" + "hello") + "\n\t") {
                case "foo bar":
                    var v22 string = "\n\t"
                    _ = v22
                case "":
                    v23 := GenS4{F0: (^(v1))}
                    _ = v23
                }
                _ = fi18
            }
            _ = fi13
            _ = v14
        }
    default:
        var v24 uint64 = 42
        _ = v24
    }
    v1 = 0
    v1 -= v1
    if false {
        if (42 <= 712) {
            v25 := GenS5{F0: 0, F1: 255}
            _ = v25
        }
        fmt.Println(v1, v1)
        v26 := GenS6{F0: 127, F1: false}
        _ = v26
    }
    v3 += v3
    if (100 > 74) {
        if (42 < -1) {
            v27 := GenS4{F0: 128}
            v28 := GenS6{F0: (v1 + 42), F1: true}
            v1 = (v1 + 127)
            _ = v27
            _ = v28
        }
        v3 = (0.0 - (v3 + float32(3.14159)))
        var v29 float64 = (((3.14159 + 0.0) / 1.0) * 0.1)
        _ = v29
    } else {
        var v30 int32 = 0
        _ = v30
    }
    fmt.Println(v3, v3)
    v2 = append(v2, 128)
    fmt.Println(v3, v1)
    v31 := GenS7{F0: 994}
    v3 = v3
    var v32 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
}
