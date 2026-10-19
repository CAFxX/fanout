package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int32 }
type GenS2 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint8, p1 int32, p2 int64) uint32 {
    if true {
        p1 = p1
        var v6 uint16 = 65535
        v6 = v6
        _ = v6
    }
    if false {
        var v7 int16 = 1
        _ = v7
    }
    v8 := GenS1{F0: 623, F1: 0, F2: 510}
    p2 -= (p2 + (p2 / ((1) | 1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 0
    v3 := []uint32{2147483647}
    var v4 uint = 18446744073709551615
    var v5 int32 = -1
    v4 -= (v4 * (v4 / ((v4) | 1)))
    switch "" {
    case "":
        v4 += (^(v4))
    case "hello":
        var v9 int32 = 123
        _ = v9
    }
    if true {
        if true {
            switch ("hello" + ("\n\t" + "a")) {
            case "":
                var v10 uint16 = 255
                _ = v10
            }
            v11 := GenS1{F0: 42, F1: 100, F2: v5}
            _ = v11
        } else {
            var v12 int8 = 0
            v4 = (v4 | v4)
            for fi13 := 0; fi13 < 14; fi13++ {
                var v14 int16 = 1
                if false {
                    fmt.Println(v2, v5)
                    _ = len(v1)
                    fmt.Println(v14, v14, v5)
                }
                v15 := GenS1{F0: 565, F1: v12, F2: 2147483647}
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = v12
        }
        if (false && true) {
            v5 += (-(v5))
            if (true || (83 == 55)) {
                v4 += v4
                if true {
                    for fi16 := 0; fi16 < 4; fi16++ {
                        v2 += 1000
                        switch "" {
                        case "":
                            v3 = append(v3, 2147483648)
                        case "\n\t":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case "a":
                            switch fi16 {
                            case -2147483648:
                                fmt.Println(v5)
                            default:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        }
                        _ = fi16
                    }
                    var v17 float64 = ((-2.25 * (float64(((0.1 / -2.25) - (3.14159 - -2.25))) * float64(0.0))) + float64(0.0))
                    _ = v17
                }
                v2 = v2
            }
            var v18 string = "foo bar"
            _ = v18
        } else {
            v2 -= 1
            v5 = v5
            v5 += v5
        }
    } else {
        v4 = v4
    }
    v4 = 76
    fmt.Println(v4, v5)
    v19 := v1[("hello" + (("" + "") + "hello"))]
    v20 := GenS1{F0: 18446744073709551615, F1: -128, F2: 26}
    _ = len(v3)
    v3 = append(v3, 0)
    _ = len(v1)
    v21 := GenS2{F0: (true || v19), F1: 127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
