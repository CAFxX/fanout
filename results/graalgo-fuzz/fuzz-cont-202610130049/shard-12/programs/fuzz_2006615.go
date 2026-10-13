package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int32 }

func main() {
    var v1 float64 = 1.5
    var v2 uint64 = 4294967295
    var v3 uint = 65535
    v4 := []string{}
    v1 -= v1
    fmt.Println(v2, v3)
    for fi5 := 0; fi5 < 2; fi5++ {
        for fi6 := 0; fi6 < 19; fi6++ {
            var v7 bool = false
            switch "foo bar" {
            case "foo bar":
                _ = len(v4)
            case "":
                var v8 int32 = -1
                _ = v8
            case "a":
                switch ("\n\t" + "foo bar") {
                case "foo bar":
                    var v9 int = 2147483647
                    _ = v9
                case "hello":
                    v4 = append(v4, "foo bar")
                }
            }
            var v10 int16 = -1
            _ = fi6
            _ = v10
            _ = v7
        }
        v2 -= 613
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
        _ = fi5
    }
    fmt.Println(v1, v3, v2)
    var v11 float32 = (1.5 * (((0.1 - (0.0 / 1.0)) / 1e10) * 3.14159))
    fmt.Println(v1)
    for fi12 := 0; fi12 < 10; fi12++ {
        fmt.Println(fi12)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if (!true) {
            switch ("foo bar" + ("" + "")) {
            case "hello":
                var v13 uint64 = v2
                _ = v13
            case "a":
                for fi14 := 0; fi14 < 9; fi14++ {
                    switch fi14 {
                    case -888:
                        v11 = v11
                    case 2147483647:
                        var v15 int16 = 1000
                        _ = v15
                    }
                    fmt.Println(fi14, fi14, fi14)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    fmt.Println(fi12, v1)
                    if (-1 < -93) {
                        var v16 float32 = v11
                        _ = v16
                    } else {
                        v4 = append(v4, "a")
                        v17 := GenS1{F0: -32768, F1: v3}
                        v18 := GenS2{F0: 1}
                        _ = v17
                        _ = v18
                    }
                    _ = fi14
                }
            case "\n\t":
                switch "foo bar" {
                case "foo bar":
                    switch fi12 {
                    case 766:
                        fmt.Println(v1)
                    case 32768:
                        fmt.Println(v3)
                    case -878:
                        var v19 int32 = -1
                        _ = v19
                    default:
                        v20 := GenS3{F0: 4294967295, F1: (v3 >> ((v3) & 63))}
                        _ = v20
                    }
                case "hello":
                    for fi21 := 0; fi21 < 11; fi21++ {
                        v3 = v3
                        _ = fi21
                    }
                case "a":
                    v22 := GenS4{F0: 138}
                    _ = v22
                default:
                    if false {
                        for fi23 := 0; fi23 < 3; fi23++ {
                            v24 := GenS5{F0: 32767, F1: 783}
                            if len(v4) > 0 {
                                v4[0] = "hello"
                            }
                            var v25 uint8 = 139
                            v25 = (^(v25))
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        v3 = (v3 | (v3 % ((v3) | 1)))
                        fmt.Println(fi12, v2, v2)
                    }
                }
            default:
                var v26 float64 = 1e10
                _ = v26
            }
        } else {
            var v27 bool = ((v3 / ((536) | 1)) > 593)
            fmt.Println(v2)
            _ = v27
        }
        _ = fi12
    }
    v28 := GenS6{F0: 708}
    v1 = float64(((v11 * 3.14159) * 1e10))
    v4 = append(v4, ("hello" + (("hello" + ((("foo bar" + "hello") + "\n\t") + "")) + "hello")))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v29 := GenS7{F0: -2147483648}
    _ = len(v4)
    var v30 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
