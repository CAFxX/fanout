package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint }
func genFn0(p0 string, p1 uint16, p2 int64) uint32 {
    fmt.Println(p2, p1)
    fmt.Println(p2, p1, p0)
    var v4 float32 = 3.14159
    v5 := GenS1{F0: 4294967295, F1: 980, F2: (^(p1))}
    for fi6 := 0; fi6 < 7; fi6++ {
        p2 -= (p2 | (p2 ^ (p2 | (p2 * 316))))
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 int16 = 42
            var v9 float32 = float32(((3.14159 / 0.1) / 1.5))
            _ = fi7
            _ = v8
            _ = v9
        }
        v4 -= v4
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 0
}

func main() {
    var v1 int = -2147483648
    var v2 string = "hello"
    v3 := []uint32{}
    v2 = v2
    _ = len(v3)
    if true {
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 uint = 42
            v11 += 2
            _ = fi10
            _ = v11
        }
        v12 := GenS1{F0: 2, F1: 2147483647, F2: 2}
        _ = v12
    } else {
        switch v2 {
        case "":
            v13 := GenS1{F0: 548, F1: 1, F2: 0}
            _ = v13
        case "foo bar":
            v2 = v2
        case "\n\t":
            for fi14 := 0; fi14 < 16; fi14++ {
                switch fi14 {
                case 32767:
                    var v15 uint = 255
                    _ = v15
                case -9223372036854775808:
                    v1 += (fi14 - (v1 * -2147483648))
                case 178:
                    v2 = "\n\t"
                default:
                    if false {
                        v16 := GenS1{F0: 4294967295, F1: -295, F2: 279}
                        v3 = append(v3, 42)
                        v17 := GenS2{F0: 586, F1: 949, F2: 2}
                        _ = v16
                        _ = v17
                    } else {
                        if len(v3) > 0 {
                            v3[0] = 75
                        }
                        v2 = v2
                        for fi18 := 0; fi18 < 12; fi18++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi18
                        }
                    }
                }
                if (!(2147483648 <= 1)) {
                    switch ("\n\t" + v2) {
                    case "a":
                        var v19 float64 = 1.5
                        _ = v19
                    case "\n\t":
                        for fi20 := 0; fi20 < 6; fi20++ {
                            v21 := GenS1{F0: 42, F1: 2147483647, F2: 442}
                            v2 = "hello"
                            fmt.Println(v2)
                            _ = fi20
                            _ = v21
                        }
                    }
                }
                _ = fi14
            }
        default:
            var v22 int64 = 1
            _ = v22
        }
    }
    var v23 int = -2147483648
    fmt.Println(v23)
    var v24 float32 = 1e10
    fmt.Println(v1)
    var v25 int16 = -1
    if (true && false) {
        v24 = (0.1 * v24)
        v26 := GenS2{F0: 0, F1: 182, F2: 358}
        _ = v26
    }
    v1 += (v23 - 9223372036854775807)
    v27 := GenS1{F0: 18446744073709551615, F1: 2147483647, F2: 1}
    var v28 int16 = v25
    for fi29 := 0; fi29 < 9; fi29++ {
        v25 += -785
        v3 = append(v3, 148)
        for fi30 := 0; fi30 < 8; fi30++ {
            var v31 uint64 = 18446744073709551615
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    for fi32 := 0; fi32 < 5; fi32++ {
        fmt.Println(v2, v24, v24)
        var v33 int32 = -315
        _ = fi32
        _ = v33
    }
    v3 = append(v3, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v3
}
