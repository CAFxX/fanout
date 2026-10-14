package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int }
type GenS2 struct { F0 uint16; F1 uint64; F2 int }
func genFn0(p0 bool) int64 {
    v6 := []int{}
    p0 = p0
    if (1.5 == 0.0) {
        if len(v6) > 0 {
            v6[0] = 140
        }
    }
    _ = p0
    _ = v6
    return -208
}
func genFn1() int {
    var v7 int16 = -32768
    fmt.Println(v7)
    _ = v7
    return 32767
}

func main() {
    var v1 string = "foo bar"
    v2 := []uint32{417, 42, 4294967295, 1}
    var v3 int8 = 127
    var v4 int = 420
    var v5 float64 = -2.25
    var v8 int32 = 2147483647
    v9 := GenS1{F0: 18446744073709551615, F1: 0, F2: v4}
    var v10 float32 = 0.0
    var v11 float32 = (-2.25 * v10)
    v11 = 0.1
    v3 += v3
    v5 -= v5
    var v12 uint8 = 193
    var v13 float64 = -2.25
    fmt.Println(v12)
    v10 = v11
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 uint16 = 32767
        _ = fi14
        _ = v15
    }
    v2 = append(v2, 2147483647)
    v11 += 0.1
    for fi16 := 0; fi16 < 15; fi16++ {
        v17 := GenS1{F0: 330, F1: 0, F2: fi16}
        _ = fi16
        _ = v17
    }
    if (541 == (v4 / ((-9223372036854775808) | 1))) {
        switch v4 {
        case 42:
            var v18 int32 = int32(v12)
            _ = v18
        case -9223372036854775808:
            switch ("foo bar" + "\n\t") {
            case "":
                if false {
                    v19 := GenS2{F0: 255, F1: 523, F2: (v4 + genFn1())}
                    v11 -= v11
                    var v20 float64 = (((0.1 + ((0.1 * float64(v10)) / 1.0)) + v13) * (v13 - v13))
                    _ = v19
                    _ = v20
                }
            default:
                v12 = v12
            }
        case 0:
            if (42 > v3) {
                fmt.Println(v3, v4, v12)
            } else {
                fmt.Println(v3, v12, v13)
                var v21 uint = 538
                if ((((v8 - 2147483647) == (^(v8))) || (!((3.14159 * v5) == float64(v11)))) && (v3 != (v3 & -119))) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    switch v1 {
                    case "hello":
                        for fi22 := 0; fi22 < 8; fi22++ {
                            fmt.Println(v12, v12)
                            var v23 uint = (v21 ^ v21)
                            for fi24 := 0; fi24 < 18; fi24++ {
                                switch (v1 + ("" + "")) {
                                case "":
                                    fmt.Println(fi24, v12)
                                    v25 := GenS2{F0: 42, F1: 971, F2: (v4 - 2147483647)}
                                    _ = v25
                                }
                                var v26 int16 = 42
                                _ = fi24
                                _ = v26
                            }
                            _ = fi22
                            _ = v23
                        }
                    }
                }
                _ = v21
            }
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        var v27 int64 = 42
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
