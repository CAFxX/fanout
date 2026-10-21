package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
func genFn0(p0 int8, p1 uint32, p2 int64) uint32 {
    v6 := GenS1{F0: (p0 * (p0 ^ -1)), F1: -32768}
    p1 = (p1 * (p1 & 205))
    var v7 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[string]bool)
    var v3 uint16 = 1
    v4 := []int{}
    v5 := make(map[int]string)
    var v8 string = ("foo bar" + ("" + (("" + "") + "")))
    switch v8 {
    case "a":
        fmt.Println(v1, v8, v3)
    case "\n\t":
        if false {
            if len(v4) > 0 {
                v4[0] = -22
            }
        }
    case "foo bar":
        var v9 int16 = 55
        _ = v9
    }
    v8 = (v8 + "hello")
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 bool = (1 <= -32768)
        fmt.Println(fi10, v3)
        v8 = ""
        _ = fi10
        _ = v11
    }
    v1 = v1
    if true {
        if (0 == -9) {
            v3 = (v3 / (((v3 & (v3 & v3))) | 1))
            var v12 uint64 = (v1 - v1)
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        for fi14 := 0; fi14 < 12; fi14++ {
            switch fi13 {
            case -2147483648:
                switch fi13 {
                case 1:
                    var v15 uint16 = (v3 / (((v3 / ((255) | 1))) | 1))
                    _ = v15
                case 9223372036854775807:
                    switch v8 {
                    case "a":
                        var v16 uint = 255
                        _ = v16
                    case "":
                        var v17 uint32 = 133
                        _ = v17
                    case "foo bar":
                        var v18 string = v8
                        _ = v18
                    }
                case 32767:
                    fmt.Println(v8, v8, v8)
                }
            case 128:
                v8 = ((v8 + v8) + v8)
            case 4294967295:
                if true {
                    switch int(v1) {
                    case 128:
                        _ = len(v5)
                    default:
                        var v19 int = (fi14 - 0)
                        _ = v19
                    }
                    fmt.Println(v3)
                    switch v8 {
                    case "":
                        v20 := v5[fi14]
                        _ = v20
                    }
                } else {
                    var v21 int16 = -801
                    _ = v21
                }
            }
            var v22 string = v8
            _ = fi14
            _ = v22
        }
        _ = fi13
    }
    if (0 < 0) {
        if (!false) {
            v23 := GenS1{F0: 118, F1: 604}
            switch v8 {
            case "\n\t":
                if false {
                    v1 += (^(v1))
                    v24 := GenS1{F0: 127, F1: 42}
                    switch 874 {
                    case -2147483648:
                        var v25 int32 = -1
                        _ = v25
                    case 0:
                        v3 -= (v3 >> ((18446744073709551615) & 63))
                    case 4294967295:
                        v26 := GenS1{F0: 127, F1: 1}
                        _ = v26
                    default:
                        v8 = (v8 + "hello")
                    }
                    _ = v24
                }
            case "foo bar":
                if false {
                    fmt.Println(v8, v3)
                    if ((!false) || (3.14159 <= float64(float32(0.0)))) {
                        v1 -= 42
                    }
                    v8 = "hello"
                }
            default:
                v3 += 65535
            }
            v8 = "hello"
            _ = v23
        }
    }
    var v27 int = 42
    v3 = v3
    var v28 int32 = 838
    for fi29 := 0; fi29 < 2; fi29++ {
        fmt.Println(v1)
        _ = fi29
    }
    v8 = ("\n\t" + "a")
    v4 = append(v4, (v27 >> ((uint(v1)) & 63)))
    if len(v4) > 0 {
        v4[0] = (v27 >> ((18446744073709551615) & 63))
    }
    if false {
        v1 += (^(v1))
    } else {
        v4 = append(v4, v27)
        switch -9223372036854775808 {
        case 351:
            for fi30 := 0; fi30 < 1; fi30++ {
                fmt.Println(v28, v27, v3)
                fmt.Println(v27)
                v31 := GenS1{F0: 100, F1: 0}
                _ = fi30
                _ = v31
            }
        case 199:
            if len(v4) > 0 {
                v4[0] = (v27 | 2147483647)
            }
        case 2147483647:
            var v32 uint = 18446744073709551615
            _ = v32
        default:
            fmt.Println(v1, v28)
        }
    }
    fmt.Println(v27, v8)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
