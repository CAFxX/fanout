package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = -281
    v3 := make(map[uint32]bool)
    var v4 uint = 2
    if (2147483648 <= 0) {
        var v5 float32 = 3.14159
        _ = v5
    }
    v2 += (v2 * v2)
    for fi6 := 0; fi6 < 10; fi6++ {
        switch "a" {
        case "foo bar":
            switch fi6 {
            case 128:
                var v7 uint = (v4 & 65535)
                _ = v7
            case 127:
                var v8 int64 = -9223372036854775808
                _ = v8
            }
        case "":
            switch fi6 {
            case 127:
                switch fi6 {
                case -1:
                    v9 := GenS1{F0: 970, F1: uint16(v4)}
                    _ = v9
                case -9223372036854775808:
                    var v10 int16 = -157
                    _ = v10
                default:
                    v4 = v4
                }
            default:
                if (float64((((0.0 - 0.1) - -2.25) * ((0.1 * (((1.5 + 3.14159) - ((3.14159 + 0.0) + 3.14159)) / (1.5 * float32(float64(1e10))))) - 0.0))) > (-2.25 + 0.1)) {
                    v2 = (-(v2))
                    switch (fi6 * 2147483647) {
                    case 127:
                        if (0 == 389) {
                            var v11 int = fi6
                            fmt.Println(v4, v4, v11)
                            _ = v11
                        }
                    case -9223372036854775808:
                        fmt.Println(v4)
                    case 1:
                        v4 -= 255
                    default:
                        var v12 int8 = -106
                        _ = v12
                    }
                }
            }
        }
        if (0 <= 565) {
            if true {
                fmt.Println(fi6)
                v13 := []int{fi6, fi6, 42, 32768}
                _ = v13
            }
        }
        _ = fi6
    }
    var v14 int32 = 355
    var v15 uint64 = 444
    for fi16 := 0; fi16 < 7; fi16++ {
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 float64 = (float64((1.5 / 1e10)) / 1.5)
            var v19 uint8 = 42
            v20 := []int{(fi17 - (fi16 + (-(fi17)))), (fi17 * (fi16 + (-(fi16)))), fi17, (^(fi17))}
            var v21 uint32 = 0
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
            _ = v21
        }
        for fi22 := 0; fi22 < 3; fi22++ {
            if false {
                v2 += 1000
            } else {
                switch "foo bar" {
                case "\n\t":
                    var v23 int = fi22
                    _ = v23
                }
                v14 += v14
                if true {
                    fmt.Println(fi22, v15, v14)
                } else {
                    v24 := v3[65535]
                    fmt.Println(v2, fi22, fi22)
                    _ = v24
                }
                v4 = (v4 & (v4 << ((v4) & 63)))
                v25 := []int{fi16, fi16}
                _ = v25
            }
            _ = fi22
        }
        v2 += (^(v2))
        _ = fi16
    }
    var v26 bool = true
    switch "\n\t" {
    case "":
        for fi27 := 0; fi27 < 10; fi27++ {
            v28 := []int{0, (fi27 / ((fi27) | 1)), int(v2)}
            fmt.Println(v15, v2)
            _ = fi27
            _ = v28
        }
    case "foo bar":
        if (!false) {
            v15 -= 985
        }
    case "a":
        var v29 int16 = -32768
        _ = v29
    default:
        var v30 int64 = 9223372036854775807
        _ = v30
    }
    var v31 int = -968
    var v32 float32 = 1e10
    fmt.Println(v4, v31, v2)
    var v33 int16 = (v2 % (((v2 >> ((v4) & 63))) | 1))
    var v34 float32 = ((v32 - v32) * 3.14159)
    var v35 uint = 414
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
