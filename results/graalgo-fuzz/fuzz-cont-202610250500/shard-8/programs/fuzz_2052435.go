package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint16; F1 uint }
func genFn0() int8 {
    v7 := []int{}
    var v8 string = "hello"
    v8 = (v8 + v8)
    var v9 string = ""
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 16
    var v3 int16 = -18
    v4 := []int{897, -533, 2147483647, 32768}
    var v5 int = 128
    var v6 uint16 = 623
    var v10 uint8 = 127
    var v11 int32 = -1
    v3 = (-(v3))
    var v12 int32 = (v11 * (v11 & (v11 | v11)))
    var v13 float64 = 0.1
    v14 := GenS1{F0: (v2 << ((18446744073709551615) & 63)), F1: 2}
    if (-1 <= -324) {
        delete(v1, 4294967295)
        var v15 float32 = ((float32(float64(1e10)) / 3.14159) * 3.14159)
        _ = v15
    }
    v10 += v10
    if (v2 != 369) {
        for fi16 := 0; fi16 < 16; fi16++ {
            v17 := GenS1{F0: 18446744073709551615, F1: 971}
            v13 -= 3.14159
            _ = fi16
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 uint32 = 1
        for fi20 := 0; fi20 < 18; fi20++ {
            if len(v4) > 0 {
                v4[0] = int(v19)
            }
            if false {
                if false {
                    switch "" {
                    case "\n\t":
                        var v21 int = 9223372036854775807
                        _ = v21
                    case "foo bar":
                        if false {
                            v22 := GenS1{F0: 989, F1: 2147483647}
                            _ = v22
                        }
                    case "a":
                        switch v5 {
                        case -702:
                            var v23 int64 = 889
                            _ = v23
                        default:
                            v24 := GenS2{F0: (v6 & v6), F1: 84}
                            _ = v24
                        }
                    default:
                        var v25 int8 = -1
                        _ = v25
                    }
                    var v26 float32 = 1.5
                    _ = v26
                }
            }
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    var v27 bool = false
    switch (v5 * (v5 * (v5 + (v5 * (v5 % ((int(v12)) | 1)))))) {
    case 42:
        var v28 int = 32768
        _ = v28
    }
    if ((!(v27 || (!v27))) && ((v12 * 42) >= -1)) {
        for fi29 := 0; fi29 < 11; fi29++ {
            v10 = (^(v10))
            var v30 int8 = -92
            _ = fi29
            _ = v30
        }
        fmt.Println(v6, v6, v12)
    }
    for fi31 := 0; fi31 < 7; fi31++ {
        v13 += v13
        _ = fi31
    }
    v32 := v1[4294967295]
    v3 = (v3 << ((933) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
