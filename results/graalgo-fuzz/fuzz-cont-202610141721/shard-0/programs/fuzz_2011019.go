package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int16 }
func genFn0(p0 int8) int8 {
    v4 := GenS1{F0: 193, F1: -9223372036854775808, F2: 32767}
    p0 = (^(p0))
    if ((false && (42 == 127)) || ((float32(((-2.25 + (1.5 - (float64(1.5) / -2.25))) / -2.25)) == -2.25) && (!true))) {
        fmt.Println(p0)
    } else {
        p0 = (p0 * 100)
        var v5 uint64 = 2
        _ = v5
    }
    var v6 bool = ((1.5 * -2.25) == (((0.1 - ((0.1 + 0.1) / 1.0)) + (-2.25 * 1e10)) + ((((-2.25 * 3.14159) / 1.5) + 0.1) - (-2.25 / (1.5 * -2.25)))))
    v6 = v6
    _ = p0
    _ = v4
    _ = v6
    return (^(p0))
}

func main() {
    v1 := []int{266, -911}
    var v2 bool = (0.1 > -2.25)
    v3 := make(map[string]int)
    for fi7 := 0; fi7 < 9; fi7++ {
        v2 = (509 != 2147483647)
        var v8 string = "a"
        v8 = v8
        _ = fi7
        _ = v8
    }
    v2 = ((18446744073709551615 == 1) || v2)
    for fi9 := 0; fi9 < 12; fi9++ {
        v2 = v2
        _ = fi9
    }
    var v10 uint64 = 42
    _ = len(v3)
    for fi11 := 0; fi11 < 2; fi11++ {
        for fi12 := 0; fi12 < 8; fi12++ {
            v1 = append(v1, 42)
            v10 = v10
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = len(v1)
            v2 = (30 == 128)
            _ = fi12
        }
        switch (fi11 / (((-(fi11))) | 1)) {
        case 686:
            for fi13 := 0; fi13 < 5; fi13++ {
                _ = len(v3)
                v10 = 940
                var v14 float32 = (-2.25 / 1.0)
                _ = fi13
                _ = v14
            }
        case 2147483647:
            v2 = (-1 <= 1)
        case 215:
            if true {
                if (-129 != 9223372036854775807) {
                    if ((v2 && v2) || (!false)) {
                        if ("hello" > ("a" + "\n\t")) {
                            var v15 bool = true
                            _ = v15
                        }
                        for fi16 := 0; fi16 < 12; fi16++ {
                            v2 = (!(294 <= 2147483647))
                            if (65535 >= 2147483647) {
                                var v17 int8 = 1
                                _ = v17
                            } else {
                                v18 := GenS1{F0: 65535, F1: 9223372036854775807, F2: 42}
                                v2 = (genFn0(42) > 100)
                                v19 := v3[""]
                                _ = v18
                                _ = v19
                            }
                            var v20 int32 = 0
                            _ = fi16
                            _ = v20
                        }
                        _ = len(v3)
                    }
                    v10 = (v10 - (^(v10)))
                    switch fi11 {
                    case 42:
                        v10 -= 4294967295
                    case 127:
                        var v21 int64 = 1
                        _ = v21
                    default:
                        var v22 int = fi11
                        _ = v22
                    }
                }
                v2 = ((float32((-2.25 * 1e10)) / float32(-2.25)) >= ((-2.25 - 0.0) - 0.0))
                v10 = 18446744073709551615
            }
        default:
            v10 = v10
        }
        v23 := GenS1{F0: 2, F1: -9223372036854775808, F2: -32768}
        _ = fi11
        _ = v23
    }
    if true {
        v2 = v2
    }
    v2 = (-51 > 921)
    fmt.Println(v10, v2)
    v3[""] = 128
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 uint32 = 43
        _ = fi24
        _ = v25
    }
    switch -9223372036854775808 {
    case -647:
        var v26 float32 = (0.0 * 1.5)
        _ = v26
    default:
        if ((1.5 - ((float32(-2.25) + 3.14159) + (float32(0.1) * 1.5))) <= 3.14159) {
            v1 = append(v1, 132)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
}
