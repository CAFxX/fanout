package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int64 }

func main() {
    v1 := []uint32{529}
    v2 := make(map[int]uint32)
    v3 := make(map[int]string)
    v4 := make(map[uint32]int)
    var v5 uint8 = 127
    v6 := []uint32{0, 556}
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int8 = -128
        v5 -= v5
        _ = fi7
        _ = v8
    }
    v5 -= v5
    if false {
        if (!false) {
            fmt.Println(v5)
            if true {
                var v9 uint16 = 255
                v9 = v9
                v9 -= 397
                _ = v9
            } else {
                if len(v6) > 0 {
                    v6[0] = 491
                }
            }
        } else {
            v5 = 1
            if (0 > 0) {
                for fi10 := 0; fi10 < 10; fi10++ {
                    for fi11 := 0; fi11 < 6; fi11++ {
                        fmt.Println(fi10)
                        if len(v6) > 0 {
                            v6[0] = 422
                        }
                        _ = fi11
                    }
                    fmt.Println(fi10, fi10)
                    v5 = v5
                    _ = fi10
                }
            }
            v12 := GenS1{F0: 65535, F1: 116, F2: -9223372036854775808}
            _ = v12
        }
    } else {
        v5 = v5
        switch -860 {
        case 2147483647:
            if false {
                switch "" {
                case "a":
                    switch "" {
                    case "hello":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case "":
                        v5 = 220
                    default:
                        for fi13 := 0; fi13 < 15; fi13++ {
                            v5 = (v5 & 2)
                            var v14 int32 = -2147483648
                            var v15 bool = false
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    }
                case "foo bar":
                    v5 = 127
                default:
                    var v16 uint64 = 147
                    _ = v16
                }
            }
        case 640:
            if (1 >= 42) {
                _ = len(v1)
                var v17 float64 = ((1.5 * float64(3.14159)) * 0.1)
                if len(v6) > 0 {
                    v6[0] = 65535
                }
                _ = v17
            } else {
                v5 = (v5 | v5)
            }
        }
        if ("" >= "a") {
            _ = len(v4)
        }
    }
    delete(v3, -989)
    if (282 <= 42) {
        fmt.Println(v5)
        if ((!false) || (1 == -1)) {
            var v18 int32 = 2147483647
            _ = v18
        } else {
            var v19 int32 = 37
            var v20 float32 = float32(1.5)
            if true {
                for fi21 := 0; fi21 < 3; fi21++ {
                    fmt.Println(v5, v19, v5)
                    v5 = (v5 % ((v5) | 1))
                    _ = fi21
                }
                fmt.Println(v19, v20)
                var v22 int32 = (v19 / (((v19 >> ((255) & 63))) | 1))
                _ = v22
            } else {
                v20 -= (v20 - v20)
                v19 = (v19 | v19)
            }
            _ = v19
            _ = v20
        }
        var v23 bool = (true || true)
        _ = v23
    }
    switch 2147483647 {
    case 377:
        var v24 int32 = 42
        _ = v24
    case 16:
        var v25 bool = (524 <= int(v5))
        _ = v25
    }
    fmt.Println(v5)
    v4[321] = 128
    v26 := v4[661]
    for fi27 := 0; fi27 < 6; fi27++ {
        fmt.Println(v5)
        _ = fi27
    }
    delete(v3, v26)
    v28 := GenS1{F0: 2, F1: 1, F2: -9223372036854775808}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
