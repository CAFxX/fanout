package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int16; F1 int64; F2 uint32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 255
    var v3 float64 = ((((((1.5 * 3.14159) / 1.0) * (-2.25 * 0.1)) / (-2.25 / (0.0 + (0.1 - -2.25)))) + (1.5 - (float64(1e10) - (-2.25 / 1.0)))) - 1.5)
    var v4 int16 = 0
    var v5 int64 = 0
    var v6 float64 = 1e10
    v7 := GenS1{F0: 1, F1: v2}
    v8 := GenS2{F0: v4, F1: (^(v5)), F2: 42}
    fmt.Println(v5)
    var v9 uint32 = 624
    if true {
        var v10 uint16 = 653
        for fi11 := 0; fi11 < 10; fi11++ {
            if true {
                var v12 int32 = 919
                for fi13 := 0; fi13 < 17; fi13++ {
                    v14 := GenS2{F0: (v4 / ((v4) | 1)), F1: v5, F2: (^(v9))}
                    v15 := []int{fi11, fi13, (-(fi13))}
                    switch -296 {
                    case 532:
                        if len(v15) > 0 {
                            v15[0] = (fi11 >> ((42) & 63))
                        }
                    case 9223372036854775807:
                        switch (^(fi13)) {
                        case 172:
                            v16 := v1[fi11]
                            _ = v16
                        case -2147483648:
                            if ((-2.25 != (float32(-2.25) - ((3.14159 - (1e10 - (-2.25 / float32(v6)))) + (-2.25 / (0.1 / 3.14159))))) || false) {
                                delete(v1, fi11)
                                fmt.Println(v3, v10)
                                _ = len(v15)
                            }
                        case 128:
                            if false {
                                switch (fi13 >> ((v2) & 63)) {
                                case 2147483647:
                                    var v17 float64 = (float64((((float32(v6) + (0.1 + (-2.25 - 3.14159))) * (3.14159 - 1e10)) + (1e10 / -2.25))) * v6)
                                    _ = v17
                                case 154:
                                    var v18 bool = ("a" == "hello")
                                    _ = v18
                                default:
                                    if false {
                                        if true {
                                            v2 -= v2
                                            v19 := GenS2{F0: (v4 ^ 799), F1: v5, F2: v9}
                                            _ = v19
                                        }
                                    }
                                }
                                v9 += 2147483648
                            } else {
                                var v20 int64 = (v5 << ((42) & 63))
                                v1[int(v2)] = uint32(v2)
                                _ = v20
                            }
                        }
                    }
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                _ = v12
            } else {
                switch -975 {
                case -250:
                    for fi21 := 0; fi21 < 15; fi21++ {
                        fmt.Println(v3, v9, v5)
                        var v22 bool = false
                        _ = fi21
                        _ = v22
                    }
                default:
                    switch ("\n\t" + (((("a" + "foo bar") + (("hello" + "a") + ("\n\t" + ""))) + ("hello" + "\n\t")) + "hello")) {
                    case "\n\t":
                        switch (fi11 * fi11) {
                        case -2147483648:
                            v5 = v5
                        case 4294967295:
                            switch (fi11 & 4294967295) {
                            case 354:
                                switch (fi11 % (((fi11 & (fi11 ^ -391))) | 1)) {
                                case 9223372036854775807:
                                    for fi23 := 0; fi23 < 17; fi23++ {
                                        if false {
                                            delete(v1, (fi11 & (fi11 & -2147483648)))
                                            var v24 uint64 = 4294967295
                                            fmt.Println(v6, fi11)
                                            _ = v24
                                        }
                                        fmt.Println(v2, v10, v3)
                                        _ = fi23
                                    }
                                case 1:
                                    for fi25 := 0; fi25 < 1; fi25++ {
                                        var v26 int64 = v5
                                        v3 += (0.1 - v6)
                                        v26 = -9223372036854775808
                                        _ = fi25
                                        _ = v26
                                    }
                                case -351:
                                    switch fi11 {
                                    case 0:
                                        if ((!false) || (0 <= 135)) {
                                            fmt.Println(v10, v10)
                                            v4 -= v4
                                            v10 += 212
                                        }
                                    default:
                                        v9 = (v9 + (v9 | uint32(v10)))
                                    }
                                }
                            case 366:
                                var v27 int16 = 1000
                                _ = v27
                            }
                        case 127:
                            fmt.Println(v6, v6, v9)
                        }
                    default:
                        v9 += v9
                    }
                }
            }
            var v28 int32 = 2147483647
            _ = fi11
            _ = v28
        }
        _ = v10
    } else {
        v3 -= v3
        if false {
            switch -2147483648 {
            case -2147483648:
                var v29 float64 = v6
                _ = v29
            case 42:
                v6 -= v3
            }
            v30 := []int{32768, -609, -912, 128}
            _ = v30
        } else {
            fmt.Println(v6, v4)
            var v31 bool = true
            _ = v31
        }
    }
    var v32 int8 = 42
    v33 := []int{-189, 166}
    for fi34 := 0; fi34 < 1; fi34++ {
        v35 := GenS1{F0: 0, F1: 470}
        _ = fi34
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
