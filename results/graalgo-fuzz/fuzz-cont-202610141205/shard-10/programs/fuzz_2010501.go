package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int64; F1 bool }

func main() {
    var v1 float32 = ((((0.1 - 1e10) / 0.1) + 1.5) * 0.1)
    v2 := make(map[int]bool)
    var v3 int16 = 42
    var v4 float32 = 0.1
    fmt.Println(v4, v1, v4)
    var v5 int64 = 42
    for fi6 := 0; fi6 < 17; fi6++ {
        if true {
            if false {
                var v7 float64 = 1.5
                var v8 uint32 = 2147483647
                if (true || (128 != (fi6 * (fi6 % (((^(fi6))) | 1))))) {
                    switch "" {
                    case "\n\t":
                        v9 := []int{(fi6 & (-(fi6))), fi6, -419}
                        _ = v9
                    case "hello":
                        var v10 float64 = (1.5 / v7)
                        _ = v10
                    case "a":
                        var v11 float32 = (v1 * 1.5)
                        _ = v11
                    }
                    v4 += v1
                }
                _ = v7
                _ = v8
            }
            fmt.Println(v5, v4, v4)
            var v12 float64 = float64(v4)
            _ = v12
        }
        var v13 string = ""
        _ = fi6
        _ = v13
    }
    var v14 uint = 191
    var v15 uint8 = 127
    if false {
        var v16 int32 = 1
        for fi17 := 0; fi17 < 16; fi17++ {
            v18 := GenS1{F0: 42, F1: 1}
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 9; fi19++ {
            _ = len(v2)
            var v20 float64 = 1e10
            _ = fi19
            _ = v20
        }
        _ = v16
    } else {
        var v21 uint64 = 18446744073709551615
        v3 -= v3
        _ = v21
    }
    v3 -= (v3 ^ (v3 / ((v3) | 1)))
    v22 := GenS1{F0: 2147483647, F1: 0}
    for fi23 := 0; fi23 < 11; fi23++ {
        v15 = v15
        var v24 uint16 = 2
        var v25 int8 = 0
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 float64 = 1.5
    v27 := GenS1{F0: 1, F1: v15}
    var v28 uint16 = 32767
    if false {
        switch 9223372036854775807 {
        case -220:
            v29 := GenS1{F0: 42, F1: v15}
            _ = v29
        default:
            v30 := GenS1{F0: 505, F1: (v15 << ((v14) & 63))}
            _ = v30
        }
    } else {
        v1 -= 1e10
        v26 = v26
        if ((^(v3)) >= v3) {
            v31 := []int{32767, 127}
            var v32 bool = (!((v14 != v14) && (1 == 0)))
            _ = v31
            _ = v32
        } else {
            var v33 uint64 = 2
            for fi34 := 0; fi34 < 16; fi34++ {
                fmt.Println(v15)
                for fi35 := 0; fi35 < 6; fi35++ {
                    v36 := GenS2{F0: (v5 >> (((v14 - 4294967295)) & 63)), F1: (!false)}
                    _ = fi35
                    _ = v36
                }
                switch fi34 {
                case 42:
                    if false {
                        if (v4 > float32(v26)) {
                            for fi37 := 0; fi37 < 7; fi37++ {
                                v38 := GenS2{F0: -9223372036854775808, F1: (1e10 <= v26)}
                                _ = fi37
                                _ = v38
                            }
                        }
                        var v39 uint32 = 0
                        _ = v39
                    } else {
                        if (!false) {
                            _ = len(v2)
                            var v40 int64 = (v5 & (^(v5)))
                            _ = v40
                        } else {
                            v1 -= 3.14159
                            fmt.Println(v33)
                            v1 = float32(v26)
                        }
                        for fi41 := 0; fi41 < 1; fi41++ {
                            v42 := GenS1{F0: 0, F1: uint8(v28)}
                            _ = fi41
                            _ = v42
                        }
                        v43 := []int{(^(fi34))}
                        _ = v43
                    }
                case -9223372036854775808:
                    v33 += uint64(v14)
                default:
                    v1 += v1
                }
                _ = fi34
            }
            _ = v33
        }
    }
    v4 += v1
    for fi44 := 0; fi44 < 9; fi44++ {
        var v45 int = (fi44 + (fi44 - fi44))
        _ = fi44
        _ = v45
    }
    v46 := v2[-9223372036854775808]
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v46
    _ = v5
}
