package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint64 }
func genFn0() bool {
    if true {
        for fi6 := 0; fi6 < 12; fi6++ {
            var v7 uint64 = 42
            v8 := []int{(fi6 & (fi6 << ((1) & 63))), (fi6 ^ 2147483647), (fi6 | -642), (fi6 & (fi6 / (((fi6 - fi6)) | 1)))}
            if true {
                v7 = 2
                if (false || (128 <= 236)) {
                    if len(v8) > 0 {
                        v8[0] = fi6
                    }
                    var v9 int = 42
                    v10 := GenS1{F0: ((42 == 0) && (128 < 82))}
                    _ = v10
                    _ = v9
                }
                var v11 float32 = ((-2.25 / 1.0) * ((float32(-2.25) * 1.5) - -2.25))
                _ = v11
            } else {
                switch fi6 {
                case -2147483648:
                    v7 += (v7 * (^(v7)))
                case 32768:
                    switch (-(fi6)) {
                    case 576:
                        switch -264 {
                        case 4294967295:
                            var v12 int32 = 893
                            _ = v12
                        case 127:
                            v7 = v7
                        case 9223372036854775807:
                            switch fi6 {
                            case -9223372036854775808:
                                v7 = (v7 ^ (^(v7)))
                            case -61:
                                var v13 string = ("a" + "foo bar")
                                _ = v13
                            case 4294967295:
                                for fi14 := 0; fi14 < 6; fi14++ {
                                    var v15 float64 = 0.0
                                    var v16 uint8 = 126
                                    _ = fi14
                                    _ = v15
                                    _ = v16
                                }
                            default:
                                var v17 uint32 = 65535
                                _ = v17
                            }
                        default:
                            var v18 int64 = 728
                            _ = v18
                        }
                    case 989:
                        v7 = (v7 / ((v7) | 1))
                    default:
                        for fi19 := 0; fi19 < 1; fi19++ {
                            v7 = (^(v7))
                            _ = fi19
                        }
                    }
                case 42:
                    v7 = 875
                }
                if false {
                    var v20 int8 = 0
                    v7 = (v7 | 345)
                    _ = v20
                }
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        switch "\n\t" {
        case "":
            switch 2147483647 {
            case 4294967295:
                var v21 uint32 = 0
                _ = v21
            default:
                var v22 uint32 = 2147483647
                _ = v22
            }
        case "hello":
            var v23 float64 = 1.5
            _ = v23
        case "a":
            v24 := []int{128, -1, -1, -1}
            _ = v24
        }
        var v25 uint32 = 2147483647
        _ = v25
    } else {
        var v26 int32 = -199
        v26 -= -2147483648
        fmt.Println(v26)
        _ = v26
    }
    var v27 bool = false
    v28 := GenS2{F0: 166, F1: 4294967295}
    _ = v27
    _ = v28
    return (-9223372036854775808 != -9223372036854775808)
}

func main() {
    var v1 int = 98
    v2 := make(map[string]uint32)
    var v3 int = (v1 & v1)
    v4 := make(map[string]int)
    var v5 int64 = -910
    fmt.Println(v3)
    fmt.Println(v3, v3, v1)
    var v29 uint32 = 42
    fmt.Println(v5, v3, v5)
    if (127 != 1) {
        var v30 string = "foo bar"
        _ = v30
    }
    v31 := []int{-1}
    var v32 float32 = 1e10
    fmt.Println(v29, v29, v29)
    v33 := v4[""]
    var v34 uint32 = 2147483648
    fmt.Println(v34)
    var v35 int64 = -440
    v36 := v4["foo bar"]
    var v37 int32 = 2147483647
    v38 := GenS1{F0: (174 == 728)}
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
