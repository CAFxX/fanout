package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }

func main() {
    var v1 uint64 = 595
    v2 := make(map[string]int)
    v3 := []int{}
    var v4 float32 = (1.5 * (1e10 + 1e10))
    var v5 bool = false
    fmt.Println(v1)
    switch "\n\t" {
    case "\n\t":
        var v6 uint8 = 128
        _ = v6
    default:
        for fi7 := 0; fi7 < 10; fi7++ {
            if (!false) {
                var v8 uint16 = 983
                var v9 int16 = 1
                _ = v8
                _ = v9
            }
            var v10 uint32 = 4294967295
            v11 := GenS1{F0: 42}
            _ = fi7
            _ = v10
            _ = v11
        }
    }
    v12 := GenS1{F0: 0}
    for fi13 := 0; fi13 < 20; fi13++ {
        delete(v2, "\n\t")
        v3 = append(v3, -9223372036854775808)
        var v14 int = -1
        _ = fi13
        _ = v14
    }
    if (3.14159 != (v4 + v4)) {
        v15 := GenS1{F0: -1}
        v1 = v1
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 uint32 = 437
            _ = fi16
            _ = v17
        }
        _ = v15
    }
    _ = len(v2)
    for fi18 := 0; fi18 < 17; fi18++ {
        if (1 != 2147483647) {
            var v19 int = 420
            for fi20 := 0; fi20 < 18; fi20++ {
                if v5 {
                    v4 -= float32(1.5)
                    fmt.Println(fi18, v19, fi18)
                    fmt.Println(fi18)
                    switch -2147483648 {
                    case -290:
                        if v5 {
                            v21 := GenS1{F0: 669}
                            v22 := GenS1{F0: 42}
                            var v23 uint8 = 98
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                    case 209:
                        v4 += float32(((((0.0 * 0.0) * 3.14159) - float64(v4)) - ((float64((0.0 + (v4 / float32(((1e10 / 1.0) / 3.14159))))) + 3.14159) + 0.1)))
                    case -2147483648:
                        v1 += 0
                    default:
                        v4 += (((float32((((3.14159 * float64(v4)) * 0.1) - float64(1e10))) * (float32(3.14159) + v4)) / 1e10) - (1.5 + (v4 + ((((((v4 / (0.1 * 3.14159)) - v4) * (v4 - v4)) - (v4 * (v4 * (1.5 * (v4 * float32(1.5)))))) * ((1e10 - 1e10) * -2.25)) + 1.5))))
                    }
                    fmt.Println(v19, v1)
                    fmt.Println(v1)
                    fmt.Println(v19, v5)
                    v24 := GenS1{F0: -480}
                    v25 := GenS2{F0: ""}
                    switch (v19 | (fi20 - (fi20 + (fi20 * (^(fi20)))))) {
                    case 4294967295:
                        fmt.Println(v1)
                    default:
                        fmt.Println(fi18)
                    }
                    _ = v24
                    _ = v25
                } else {
                    var v26 uint16 = 0
                    _ = v26
                }
                if v5 {
                    v27 := GenS1{F0: 42}
                    _ = v27
                }
                var v28 bool = ((v1 | (^(v1))) <= (v1 ^ v1))
                _ = fi20
                _ = v28
            }
            for fi29 := 0; fi29 < 4; fi29++ {
                v1 = (v1 | v1)
                _ = fi29
            }
            _ = v19
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
