package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint64; F1 int32; F2 int }

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{}
    v3 := make(map[uint32]uint32)
    if (!((9223372036854775807 > 1) && false)) {
        v2 = append(v2, 2147483648)
        var v4 int64 = -1
        _ = v4
    } else {
        switch "" {
        case "a":
            for fi5 := 0; fi5 < 9; fi5++ {
                if (false || true) {
                    v6 := GenS1{F0: -9223372036854775808, F1: 42}
                    var v7 int16 = 11
                    if (fi5 > (fi5 % ((4294967295) | 1))) {
                        v7 -= 42
                        var v8 bool = (!true)
                        v9 := GenS1{F0: -1, F1: -1}
                        _ = v8
                        _ = v9
                    } else {
                        for fi10 := 0; fi10 < 16; fi10++ {
                            var v11 int8 = 100
                            switch (fi5 << ((913) & 63)) {
                            case 2147483647:
                                v12 := v1[fi5]
                                _ = v12
                            }
                            _ = fi10
                            _ = v11
                        }
                        delete(v3, 317)
                    }
                    _ = v6
                    _ = v7
                } else {
                    var v13 int32 = 2147483647
                    _ = v13
                }
                _ = fi5
            }
        case "":
            var v14 uint32 = 468
            _ = v14
        default:
            for fi15 := 0; fi15 < 8; fi15++ {
                var v16 int16 = 0
                v16 += (v16 ^ v16)
                v17 := GenS2{F0: 18446744073709551615, F1: -135, F2: (fi15 * 2147483647)}
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        var v18 uint16 = 1
        _ = v18
    }
    var v19 uint16 = 2
    if false {
        v19 = (v19 & (v19 ^ (v19 - (v19 * 1))))
        var v20 int64 = 42
        var v21 int16 = 0
        _ = v20
        _ = v21
    } else {
        var v22 uint = 1
        _ = v22
    }
    if (((0.0 + -2.25) * ((3.14159 / -2.25) - 3.14159)) >= 0.0) {
        v23 := GenS2{F0: 2, F1: -2147483648, F2: 128}
        _ = v23
    }
    var v24 bool = (((207 == 0) || false) && (!(false || (2147483647 == 65535))))
    fmt.Println(v19, v19)
    var v25 int8 = 0
    var v26 float32 = float32((0.0 + (0.1 * 0.0)))
    var v27 uint32 = 4294967295
    if ("hello" > "foo bar") {
        v24 = (0 >= 42)
        v28 := GenS1{F0: -9223372036854775808, F1: 42}
        for fi29 := 0; fi29 < 18; fi29++ {
            for fi30 := 0; fi30 < 18; fi30++ {
                var v31 int16 = -32768
                var v32 int8 = v25
                if v24 {
                    var v33 float64 = 1.5
                    _ = v33
                }
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = fi29
        }
        _ = v28
    } else {
        var v34 float32 = v26
        var v35 int16 = -1
        v24 = (2 > v19)
        _ = v34
        _ = v35
    }
    v27 += (^(v27))
    v24 = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
