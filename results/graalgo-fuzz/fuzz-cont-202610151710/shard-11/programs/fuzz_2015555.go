package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS2 struct { F0 int64; F1 int16; F2 int16 }
type GenS3 struct { F0 uint }

func main() {
    var v1 int32 = 0
    var v2 int32 = (^(v1))
    var v3 int = -697
    v4 := []int{v3, 42, -1}
    var v5 uint64 = 4294967295
    if (false || true) {
        var v6 float64 = ((float64(3.14159) - 0.1) + -2.25)
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 uint8 = 218
            _ = fi7
            _ = v8
        }
        _ = v6
    } else {
        v9 := GenS1{F0: 2147483648, F1: 1, F2: -32768}
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        fmt.Println(fi10)
        v11 := GenS1{F0: 65535, F1: 635, F2: 0}
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        for fi13 := 0; fi13 < 2; fi13++ {
            switch 128 {
            case 32768:
                if (v5 < (^(v5))) {
                    if len(v4) > 0 {
                        v4[0] = -348
                    }
                } else {
                    switch 127 {
                    case 0:
                        var v14 uint64 = v5
                        _ = v14
                    case 9223372036854775807:
                        var v15 uint64 = v5
                        _ = v15
                    default:
                        if (v5 == 18446744073709551615) {
                            var v16 uint8 = uint8(v5)
                            if (890 == 875) {
                                var v17 uint32 = 1
                                var v18 int16 = 0
                                if ((true && false) || true) {
                                    var v19 uint = 255
                                    v19 -= 4294967295
                                    v18 += -1
                                    _ = v19
                                } else {
                                    if (true || (v3 != 630)) {
                                        v20 := GenS1{F0: v17, F1: 2, F2: 1}
                                        v21 := GenS2{F0: 9223372036854775807, F1: 32767, F2: (v18 | (v18 >> ((uint(v17)) & 63)))}
                                        fmt.Println(v16, v3)
                                        _ = v20
                                        _ = v21
                                    }
                                }
                                v22 := GenS2{F0: -1, F1: v18, F2: (v18 ^ v18)}
                                v23 := GenS2{F0: 42, F1: -1, F2: v18}
                                _ = v17
                                _ = v18
                                _ = v22
                                _ = v23
                            }
                            var v24 int32 = (v1 >> ((42) & 63))
                            _ = v16
                            _ = v24
                        }
                    }
                    fmt.Println(v1)
                }
            case 32767:
                v5 -= (v5 + v5)
            }
            _ = fi13
        }
        _ = fi12
    }
    var v25 uint8 = 2
    fmt.Println(v2)
    if true {
        v26 := GenS2{F0: 9223372036854775807, F1: -1, F2: 32767}
        _ = v26
    }
    v1 -= v1
    fmt.Println(v5, v25)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v25 = (v25 >> ((42) & 63))
    if true {
        v27 := GenS1{F0: 65535, F1: 32767, F2: 42}
        var v28 float64 = 3.14159
        _ = v27
        _ = v28
    } else {
        v2 = (v1 << ((2) & 63))
        if true {
            v29 := GenS3{F0: 42}
            fmt.Println(v25)
            v4 = append(v4, v3)
            _ = v29
        }
    }
    var v30 uint16 = 761
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
