package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 uint = 0
    var v2 int32 = 42
    var v3 int16 = -32768
    var v4 uint32 = 2147483648
    v5 := []int{927, int(v2), 2147483647}
    var v6 uint = (v1 + (^(v1)))
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint64 = 0
        var v9 bool = true
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 uint8 = 2
            var v12 int16 = v3
            var v13 int32 = (v2 & -890)
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    switch 0 {
    case -1:
        v14 := GenS1{F0: (v6 ^ v6)}
        _ = v14
    case 128:
        var v15 float64 = float64(-2.25)
        _ = v15
    default:
        v3 = v3
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        fmt.Println(fi16, v3, v2)
        switch "a" {
        case "hello":
            v5 = append(v5, fi16)
        default:
            if false {
                if true {
                    v3 -= v3
                    if false {
                        for fi17 := 0; fi17 < 6; fi17++ {
                            v18 := GenS2{F0: 2147483647}
                            _ = fi17
                            _ = v18
                        }
                        if (1.5 >= -2.25) {
                            v19 := GenS3{F0: (v2 / (((v2 & (-(v2)))) | 1)), F1: v3, F2: 810}
                            for fi20 := 0; fi20 < 5; fi20++ {
                                var v21 int = (fi16 - (fi16 ^ -842))
                                _ = fi20
                                _ = v21
                            }
                            _ = v19
                        }
                    }
                    var v22 int = (fi16 - fi16)
                    _ = v22
                }
                var v23 float64 = (float64(-2.25) / (0.1 * (3.14159 - (0.0 + -2.25))))
                fmt.Println(v1, v3, v2)
                _ = v23
            }
        }
        for fi24 := 0; fi24 < 9; fi24++ {
            v1 = 65535
            fmt.Println(v1, v4, fi16)
            _ = fi24
        }
        _ = fi16
    }
    var v25 int = -2147483648
    fmt.Println(v6, v6)
    var v26 float32 = (-2.25 + 1.5)
    fmt.Println(v1, v3)
    switch "foo bar" {
    case "foo bar":
        if (false || (-95 != 0)) {
            var v27 int8 = 21
            _ = v27
        }
    case "hello":
        fmt.Println(v6)
    case "":
        if (((v3 >> (((v1 ^ 2)) & 63)) <= (v3 - (v3 + (v3 - v3)))) && true) {
            switch "" {
            case "\n\t":
                fmt.Println(v1)
            case "a":
                switch "hello" {
                case "foo bar":
                    v5 = append(v5, v25)
                case "a":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case "hello":
                    v25 = (v25 * v25)
                default:
                    v3 += v3
                }
            }
            var v28 string = (("hello" + "\n\t") + "foo bar")
            var v29 int32 = (-(v2))
            _ = v28
            _ = v29
        } else {
            fmt.Println(v2, v3)
            v25 = (v25 * v25)
        }
    }
    if true {
        var v30 float32 = float32(0.0)
        var v31 uint64 = 305
        v25 = 32767
        _ = v30
        _ = v31
    }
    var v32 uint16 = 255
    var v33 int = (v25 >> ((v1) & 63))
    v25 = v33
    switch v25 {
    case 32767:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case 128:
        v4 = (v4 ^ v4)
    default:
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v3, v26)
        case "a":
            fmt.Println(v25)
        default:
            v26 += v26
        }
    }
    for fi34 := 0; fi34 < 14; fi34++ {
        fmt.Println(v26, v2, v26)
        v35 := GenS4{F0: v32}
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
