package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16; F1 string }

func main() {
    v1 := make(map[int]string)
    var v2 float64 = ((0.0 / (1e10 * (-2.25 - 0.1))) * (0.1 / 1e10))
    v3 := []int{-74, 2147483647, -2147483648, -9223372036854775808}
    switch -1 {
    case 1:
        v2 -= v2
    }
    v4 := GenS1{F0: 1, F1: 65535}
    for fi5 := 0; fi5 < 2; fi5++ {
        fmt.Println(v2)
        fmt.Println(v2)
        _ = fi5
    }
    if true {
        v2 = 1.5
        if true {
            v2 += v2
            switch "foo bar" {
            case "":
                v2 += v2
            }
        } else {
            switch -566 {
            case 606:
                switch 85 {
                case 610:
                    v2 += v2
                case 127:
                    v2 += v2
                }
            case -428:
                v6 := GenS1{F0: 100, F1: 4294967295}
                _ = v6
            case 2147483647:
                fmt.Println(v2)
            default:
                for fi7 := 0; fi7 < 17; fi7++ {
                    v2 -= v2
                    _ = fi7
                }
            }
        }
    } else {
        delete(v1, 127)
        for fi8 := 0; fi8 < 11; fi8++ {
            if (18446744073709551615 <= 1) {
                v9 := GenS2{F0: 0, F1: fi8}
                for fi10 := 0; fi10 < 7; fi10++ {
                    v2 -= float64((0.1 * ((float32(0.0) * float32(v2)) - 1e10)))
                    _ = fi10
                }
                v11 := GenS2{F0: 796, F1: fi8}
                _ = v11
                _ = v9
            } else {
                delete(v1, fi8)
            }
            _ = fi8
        }
    }
    v12 := GenS1{F0: -128, F1: 4294967295}
    if false {
        v13 := GenS3{F0: 167, F1: 1, F2: 0}
        if true {
            var v14 int16 = 32767
            for fi15 := 0; fi15 < 7; fi15++ {
                var v16 uint32 = 59
                v17 := GenS4{F0: -2147483648}
                v2 += v2
                _ = fi15
                _ = v16
                _ = v17
            }
            switch ("\n\t" + (("a" + (("\n\t" + (("hello" + "a") + "a")) + (((("" + "\n\t") + ("" + "hello")) + ("foo bar" + ("foo bar" + (("a" + "a") + "foo bar")))) + ""))) + "a")) {
            case "\n\t":
                for fi18 := 0; fi18 < 13; fi18++ {
                    v14 = 0
                    v14 = v14
                    _ = fi18
                }
            case "foo bar":
                fmt.Println(v2)
            default:
                if (!(0 > -9223372036854775808)) {
                    var v19 int = 5
                    _ = v19
                } else {
                    var v20 int64 = 0
                    for fi21 := 0; fi21 < 1; fi21++ {
                        for fi22 := 0; fi22 < 20; fi22++ {
                            v14 -= (v14 ^ (v14 * (v14 & v14)))
                            v2 -= (v2 * v2)
                            _ = fi22
                        }
                        for fi23 := 0; fi23 < 3; fi23++ {
                            fmt.Println(v2)
                            v14 += int16(v20)
                            var v24 uint8 = 0
                            _ = fi23
                            _ = v24
                        }
                        _ = fi21
                    }
                    v20 += v20
                    _ = v20
                }
            }
            _ = v14
        }
        _ = v13
    }
    var v25 uint16 = 42
    v25 -= v25
    var v26 uint32 = 2147483647
    if (false || (!false)) {
        var v27 int = -1
        for fi28 := 0; fi28 < 15; fi28++ {
            _ = len(v3)
            _ = fi28
        }
        var v29 float32 = (0.1 + 1.5)
        _ = v27
        _ = v29
    } else {
        v30 := GenS5{F0: 1, F1: "\n\t"}
        v2 += v2
        fmt.Println(v25, v26)
        _ = v30
    }
    switch "" {
    case "\n\t":
        if len(v3) > 0 {
            v3[0] = -976
        }
    case "a":
        for fi31 := 0; fi31 < 10; fi31++ {
            switch fi31 {
            case 128:
                v32 := GenS2{F0: 42, F1: fi31}
                _ = v32
            default:
                var v33 int16 = -1
                _ = v33
            }
            _ = fi31
        }
    case "hello":
        fmt.Println(v25, v26)
    }
    v25 = 2
    delete(v1, 0)
    var v34 uint32 = v26
    v35 := GenS4{F0: 2147483647}
    v1[-930] = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
