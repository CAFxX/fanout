package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint64 = 42
    var v2 uint = 0
    var v3 int = int(v1)
    v4 := []int{v3}
    v1 -= v1
    for fi5 := 0; fi5 < 10; fi5++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi5
    }
    v1 += v1
    if (false && true) {
        v2 += v2
        v6 := GenS1{F0: "a"}
        v1 = 18446744073709551615
        _ = v6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            switch -355 {
            case 127:
                _ = len(v4)
                if false {
                    switch "" {
                    case "a":
                        v2 = v2
                    case "":
                        var v9 uint64 = 0
                        _ = v9
                    case "foo bar":
                        var v10 int16 = 818
                        _ = v10
                    default:
                        if false {
                            v3 = fi8
                        }
                    }
                } else {
                    var v11 int32 = -639
                    v3 = (fi7 | -9223372036854775808)
                    _ = v11
                }
                v12 := GenS2{F0: (v2 % ((v2) | 1))}
                _ = v12
            default:
                var v13 float32 = (0.1 + -2.25)
                _ = v13
            }
            v3 = (-(fi8))
            _ = fi8
        }
        _ = fi7
    }
    fmt.Println(v1, v2)
    if false {
        var v14 float64 = (float64((1e10 * float32(-2.25))) * 3.14159)
        switch "\n\t" {
        case "\n\t":
            if len(v4) > 0 {
                v4[0] = v3
            }
        }
        switch -999 {
        case 4294967295:
            var v15 uint8 = 5
            _ = v15
        case 32767:
            if len(v4) > 0 {
                v4[0] = v3
            }
        case 658:
            v14 = (((float64(0.0) * 3.14159) + v14) / (v14 + v14))
        }
        _ = v14
    } else {
        switch v3 {
        case 757:
            var v16 uint64 = (^(v1))
            _ = v16
        }
        v3 = -91
        switch (v3 / ((127) | 1)) {
        case 536:
            v2 += (v2 * v2)
        case -253:
            v3 += (v3 ^ (v3 % ((2147483647) | 1)))
        default:
            if false {
                var v17 float32 = (-2.25 * ((3.14159 * 1.5) - 0.1))
                _ = v17
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
