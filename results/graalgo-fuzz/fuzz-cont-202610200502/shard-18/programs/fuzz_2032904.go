package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint64 }

func main() {
    var v1 float32 = ((float32(0.0) + 1.5) / (0.1 * 1.5))
    v2 := make(map[uint32]int)
    var v3 float64 = 1e10
    var v4 bool = true
    var v5 int32 = -1
    for fi6 := 0; fi6 < 14; fi6++ {
        switch fi6 {
        case -663:
            v1 -= v1
        case 0:
            var v7 uint = 18446744073709551615
            _ = v7
        case 128:
            fmt.Println(v4)
        default:
            v4 = (18446744073709551615 > 321)
        }
        fmt.Println(v4, v3)
        v3 -= (float64(((((3.14159 - 3.14159) / (-2.25 * v1)) - 0.1) + v1)) - -2.25)
        _ = fi6
    }
    if true {
        var v8 int8 = -15
        v3 -= 1.5
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 bool = false
            v11 := GenS1{F0: 32767}
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v8
    }
    switch "\n\t" {
    case "foo bar":
        var v12 int16 = 42
        _ = v12
    case "a":
        var v13 int16 = 271
        _ = v13
    case "":
        if v4 {
            if (!true) {
                switch -2147483648 {
                case -912:
                    v14 := GenS1{F0: 128}
                    _ = v14
                case -169:
                    fmt.Println(v5)
                default:
                    var v15 float32 = (v1 / 1.5)
                    _ = v15
                }
            }
            switch 143 {
            case 128:
                switch "hello" {
                case "\n\t":
                    var v16 uint8 = 0
                    _ = v16
                case "":
                    for fi17 := 0; fi17 < 12; fi17++ {
                        for fi18 := 0; fi18 < 12; fi18++ {
                            switch -237 {
                            case 1:
                                var v19 int64 = 0
                                _ = v19
                            case -9223372036854775808:
                                var v20 bool = (v4 && (42 >= 9223372036854775807))
                                _ = v20
                            }
                            v4 = (!v4)
                            _ = fi18
                        }
                        _ = fi17
                    }
                }
            case -1:
                v21 := GenS1{F0: 0}
                _ = v21
            case 32767:
                v5 = (v5 & 0)
            default:
                if true {
                    if (v4 || false) {
                        if (v3 > (v3 * (1.5 * 0.0))) {
                            v22 := []int{1, -1, 32767, 32768}
                            _ = v22
                        }
                    } else {
                        v1 = v1
                    }
                    var v23 float32 = v1
                    v4 = v4
                    _ = v23
                } else {
                    var v24 string = ((("" + "hello") + "") + ((("\n\t" + "") + "hello") + "a"))
                    _ = v24
                }
            }
        }
    default:
        var v25 int64 = 42
        _ = v25
    }
    v1 += (1e10 * (v1 / ((v1 + v1) - v1)))
    v3 = v3
    for fi26 := 0; fi26 < 14; fi26++ {
        v27 := GenS1{F0: (fi26 % (((^(fi26))) | 1))}
        var v28 int16 = -32768
        _ = fi26
        _ = v27
        _ = v28
    }
    v29 := GenS1{F0: 0}
    var v30 int8 = 127
    v31 := GenS2{F0: v4, F1: 4294967295}
    fmt.Println(v30, v5, v30)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
