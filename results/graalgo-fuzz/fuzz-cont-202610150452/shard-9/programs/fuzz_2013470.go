package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int8 }
type GenS2 struct { F0 int8; F1 int16 }
func genFn0() uint32 {
    var v4 int16 = 0
    if false {
        var v5 int = -1
        _ = v5
    } else {
        switch "" {
        case "a":
            switch "a" {
            case "\n\t":
                v6 := GenS1{F0: -1, F1: 0, F2: 100}
                _ = v6
            case "foo bar":
                v4 += v4
            }
        default:
            v7 := GenS1{F0: 1, F1: 634, F2: -1}
            _ = v7
        }
    }
    v8 := GenS1{F0: 499, F1: -214, F2: 42}
    _ = v4
    _ = v8
    return 863
}

func main() {
    var v1 uint = 2
    v2 := []int{128, 9223372036854775807, 4294967295, 42}
    var v3 float32 = (float32(0.0) - 3.14159)
    fmt.Println(v3)
    var v9 bool = false
    v1 += (v1 & (v1 / ((v1) | 1)))
    v1 = v1
    for fi10 := 0; fi10 < 12; fi10++ {
        if (0 > 0) {
            for fi11 := 0; fi11 < 12; fi11++ {
                if v9 {
                    if (-2.25 > float32(0.0)) {
                        v1 = (^(v1))
                        var v12 uint16 = 597
                        _ = v12
                    }
                } else {
                    if v9 {
                        var v13 float64 = (1e10 - 0.1)
                        _ = len(v2)
                        _ = v13
                    } else {
                        v14 := GenS2{F0: -1, F1: 32767}
                        _ = v14
                    }
                    v2 = append(v2, 127)
                    v15 := GenS2{F0: -1, F1: 413}
                    _ = v15
                }
                var v16 string = "hello"
                _ = fi11
                _ = v16
            }
        }
        if false {
            var v17 int = (-(fi10))
            switch "\n\t" {
            case "a":
                v18 := GenS2{F0: 1, F1: -1}
                _ = v18
            case "foo bar":
                _ = len(v2)
            case "\n\t":
                for fi19 := 0; fi19 < 6; fi19++ {
                    switch v17 {
                    case 887:
                        var v20 uint32 = 2147483648
                        _ = v20
                    case 127:
                        v21 := GenS2{F0: 1, F1: 1000}
                        _ = v21
                    }
                    switch 127 {
                    case -555:
                        switch (fi10 | (fi10 + (v17 << ((v1) & 63)))) {
                        case -175:
                            v22 := GenS2{F0: 127, F1: -32768}
                            _ = v22
                        case 1:
                            if v9 {
                                if true {
                                    switch (fi10 ^ fi10) {
                                    case -744:
                                        if len(v2) > 0 {
                                            v2[0] = 694
                                        }
                                    case 32768:
                                        var v23 string = ("foo bar" + ("hello" + "\n\t"))
                                        _ = v23
                                    }
                                    var v24 string = "hello"
                                    v25 := GenS2{F0: 127, F1: 1}
                                    _ = v24
                                    _ = v25
                                }
                                if (3.14159 < 1e10) {
                                    v26 := GenS2{F0: 42, F1: 0}
                                    _ = v26
                                }
                                fmt.Println(v3, v1, v17)
                            }
                        default:
                            v9 = false
                        }
                    default:
                        v17 += (fi10 % (((fi19 - 237)) | 1))
                    }
                    v17 -= (v17 - (fi19 - fi10))
                    v2 = append(v2, (fi10 % ((v17) | 1)))
                    _ = fi19
                }
            default:
                for fi27 := 0; fi27 < 3; fi27++ {
                    var v28 int8 = 1
                    _ = fi27
                    _ = v28
                }
            }
            _ = v17
        }
        _ = fi10
    }
    var v29 int8 = -1
    _ = len(v2)
    fmt.Println(v9, v9)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v9
}
