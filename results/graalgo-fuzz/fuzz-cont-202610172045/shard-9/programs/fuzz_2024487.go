package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 bool }
func genFn0(p0 int32) uint16 {
    var v5 float64 = -2.25
    p0 += p0
    for fi6 := 0; fi6 < 8; fi6++ {
        if false {
            if (!true) {
                v5 += 0.0
            }
            v7 := []int{fi6, fi6}
            _ = v7
        }
        if (1 >= 0) {
            v8 := GenS1{F0: 773, F1: 255}
            if true {
                switch ("foo bar" + "a") {
                case "\n\t":
                    var v9 int = fi6
                    _ = v9
                default:
                    p0 -= 1
                }
                v10 := GenS2{F0: 1, F1: fi6}
                p0 = 745
                _ = v10
            } else {
                v5 += v5
                v5 += v5
            }
            var v11 bool = (9223372036854775807 != fi6)
            _ = v11
            _ = v8
        }
        v12 := []int{fi6, fi6, 4294967295}
        _ = fi6
        _ = v12
    }
    _ = p0
    _ = v5
    return 2
}

func main() {
    var v1 float32 = 1.5
    var v2 float64 = 3.14159
    var v3 int64 = -1
    v4 := []int{128, 1, 541}
    var v13 float32 = v1
    var v14 uint16 = 1
    v1 -= 1e10
    var v15 float32 = (v1 - 1e10)
    v3 = 9223372036854775807
    var v16 uint = 42
    v15 = v1
    v17 := GenS3{F0: (!true)}
    for fi18 := 0; fi18 < 7; fi18++ {
        var v19 float32 = v1
        switch "hello" {
        case "":
            switch (fi18 * -1) {
            case 42:
                fmt.Println(v13, v16)
            case 9223372036854775807:
                if true {
                    if (v14 != genFn0(-918)) {
                        _ = len(v4)
                        switch "foo bar" {
                        case "hello":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case "":
                            var v20 int64 = (-(v3))
                            _ = v20
                        case "\n\t":
                            v2 += v2
                        }
                    } else {
                        switch "foo bar" {
                        case "\n\t":
                            for fi21 := 0; fi21 < 5; fi21++ {
                                v4 = append(v4, fi21)
                                _ = fi21
                            }
                        case "":
                            v22 := GenS3{F0: ((true || (65535 == 42)) || (227 > 85))}
                            _ = v22
                        case "hello":
                            var v23 int = fi18
                            _ = v23
                        }
                        var v24 float64 = v2
                        _ = v24
                    }
                } else {
                    var v25 uint32 = 1
                    _ = v25
                }
            default:
                for fi26 := 0; fi26 < 2; fi26++ {
                    switch fi18 {
                    case 4294967295:
                        _ = len(v4)
                    case 623:
                        v19 = (((((v13 + v13) / (v13 * 0.0)) * 1e10) - 3.14159) * (((v15 / (v19 - (v15 * v15))) * float32((float64((-2.25 * v19)) - (0.0 * 0.0)))) * (v19 - 1.5)))
                    }
                    _ = fi26
                }
            }
        case "hello":
            switch fi18 {
            case -2147483648:
                fmt.Println(v3)
            case 119:
                var v27 uint64 = 1
                _ = v27
            default:
                for fi28 := 0; fi28 < 4; fi28++ {
                    var v29 int16 = 32767
                    v29 = v29
                    _ = fi28
                    _ = v29
                }
            }
        case "\n\t":
            v1 -= v19
        default:
            v13 = 1.5
        }
        _ = fi18
        _ = v19
    }
    var v30 uint16 = (v14 / (((^(v14))) | 1))
    v14 = v30
    switch (("foo bar" + "hello") + ("\n\t" + "foo bar")) {
    case "":
        fmt.Println(v2)
    case "foo bar":
        fmt.Println(v15, v16, v2)
    case "\n\t":
        var v31 int8 = 83
        _ = v31
    default:
        v4 = append(v4, -2147483648)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
