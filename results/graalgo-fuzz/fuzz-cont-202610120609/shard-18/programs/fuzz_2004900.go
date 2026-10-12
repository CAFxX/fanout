package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int32 }

func main() {
    var v1 uint16 = 715
    var v2 float64 = float64((1e10 / (1.5 / 1.5)))
    var v3 string = "hello"
    var v4 uint64 = 2
    var v5 int16 = -84
    v6 := []int{}
    v7 := GenS1{F0: (-(v5)), F1: 127, F2: -398}
    v6 = append(v6, -56)
    for fi8 := 0; fi8 < 4; fi8++ {
        v1 -= (v1 & 1)
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 uint32 = 42
            if true {
                v6 = append(v6, fi9)
                if false {
                    if (fi8 != (fi9 - 42)) {
                        switch (fi9 << ((440) & 63)) {
                        case 32767:
                            fmt.Println(v3, fi8)
                        case 937:
                            if ((^(v5)) <= v5) {
                                v11 := GenS1{F0: v5, F1: 127, F2: 2147483647}
                                _ = v11
                            }
                        case 127:
                            v1 -= v1
                        }
                        _ = len(v6)
                        v4 += v4
                        if false {
                            if true {
                                fmt.Println(v3)
                                fmt.Println(fi9, fi9)
                            }
                            if len(v6) > 0 {
                                v6[0] = -9223372036854775808
                            }
                        }
                        var v12 uint16 = 660
                        _ = v12
                    } else {
                        var v13 string = ((v3 + "\n\t") + ((v3 + v3) + ("foo bar" + v3)))
                        _ = v13
                    }
                    var v14 uint32 = 2
                    fmt.Println(v5, fi8)
                    _ = v14
                }
            } else {
                v1 -= 185
                if ((1 > (v1 - (^(v1)))) || true) {
                    if (((float32(v2) / -2.25) - 0.0) < (float32(v2) - (-2.25 * -2.25))) {
                        var v15 int = (fi8 >> ((2) & 63))
                        switch (-(fi8)) {
                        case -9223372036854775808:
                            var v16 int32 = 42
                            _ = v16
                        case 32767:
                            var v17 int16 = (v5 >> ((278) & 63))
                            _ = v17
                        case 128:
                            v18 := GenS1{F0: v5, F1: -1, F2: 0}
                            _ = v18
                        default:
                            v4 -= (v4 & v4)
                        }
                        v3 = (v3 + (v3 + v3))
                        _ = v15
                    } else {
                        v2 = (0.0 / v2)
                    }
                    switch fi8 {
                    case 0:
                        var v19 uint = 178
                        _ = v19
                    }
                    switch (-(fi9)) {
                    case 341:
                        if true {
                            var v20 bool = (!true)
                            _ = v20
                        }
                    case 2147483647:
                        var v21 int64 = -47
                        _ = v21
                    default:
                        switch (v3 + v3) {
                        case "foo bar":
                            switch fi9 {
                            case 2147483647:
                                var v22 float32 = 1e10
                                _ = v22
                            case 9223372036854775807:
                                v1 += 65535
                            case -2147483648:
                                var v23 int32 = 0
                                v10 -= v10
                                v24 := GenS1{F0: (v5 - v5), F1: 22, F2: (v23 >> ((65535) & 63))}
                                _ = v23
                                _ = v24
                            }
                        default:
                            var v25 uint16 = 2
                            _ = v25
                        }
                    }
                    v6 = append(v6, fi8)
                    var v26 int8 = -128
                    if false {
                        var v27 int16 = (v5 | v5)
                        if len(v6) > 0 {
                            v6[0] = (fi8 >> ((4294967295) & 63))
                        }
                        _ = v27
                    }
                    _ = v26
                }
            }
            v10 = v10
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v4 = 18446744073709551615
    switch v3 {
    case "hello":
        var v28 int8 = -128
        _ = v28
    case "foo bar":
        v29 := GenS1{F0: (v5 - v5), F1: 15, F2: 1}
        _ = v29
    case "\n\t":
        var v30 uint16 = v1
        _ = v30
    }
    fmt.Println(v3, v3)
    v3 = v3
    switch (v3 + "\n\t") {
    case "":
        fmt.Println(v2, v3, v3)
    case "a":
        if (316 < uint(v1)) {
            switch -1 {
            case 103:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            default:
                fmt.Println(v4, v1)
            }
            var v31 uint8 = 150
            _ = v31
        }
    case "hello":
        if false {
            v3 = v3
            for fi32 := 0; fi32 < 1; fi32++ {
                var v33 int8 = 127
                _ = fi32
                _ = v33
            }
        } else {
            v2 = float64((-2.25 + -2.25))
            var v34 int8 = 0
            _ = v34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
