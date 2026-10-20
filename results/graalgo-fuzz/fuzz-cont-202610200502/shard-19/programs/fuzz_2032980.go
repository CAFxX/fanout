package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint }

func main() {
    v1 := []uint32{122, 42, 2147483647}
    var v2 uint16 = 65535
    var v3 uint16 = 65535
    var v4 float32 = 0.0
    var v5 uint64 = 836
    var v6 int16 = 1
    v1 = append(v1, 656)
    var v7 int8 = 54
    v7 = (v7 / ((v7) | 1))
    fmt.Println(v5)
    if false {
        if (42 <= 9223372036854775807) {
            v2 -= 255
            if len(v1) > 0 {
                v1[0] = 2147483648
            }
            v6 += (v6 - v6)
        }
        var v8 int64 = -9223372036854775808
        _ = v8
    }
    v1 = append(v1, 640)
    var v9 uint64 = v5
    v9 += v9
    for fi10 := 0; fi10 < 4; fi10++ {
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 int64 = 115
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    if ((1.5 < 0.1) || true) {
        var v13 uint64 = 4294967295
        _ = len(v1)
        _ = v13
    }
    v5 -= 450
    var v14 uint8 = uint8(v2)
    if (303 > -2147483648) {
        v1 = append(v1, 4294967295)
        switch -758 {
        case -475:
            var v15 float32 = float32((1.5 * (0.1 - 0.0)))
            _ = v15
        default:
            var v16 uint8 = 127
            _ = v16
        }
    }
    switch 32767 {
    case -50:
        if false {
            v6 = v6
            v17 := GenS1{F0: (!false), F1: 489, F2: 42}
            v14 -= (v14 >> ((255) & 63))
            _ = v17
        }
    default:
        if (42 == v2) {
            if (false || false) {
                v3 += v2
                if ((v7 << ((18446744073709551615) & 63)) != v7) {
                    fmt.Println(v2, v4, v2)
                } else {
                    var v18 uint32 = uint32(v14)
                    for fi19 := 0; fi19 < 4; fi19++ {
                        v9 -= (v5 + 909)
                        _ = fi19
                    }
                    if (9223372036854775807 > 42) {
                        switch (("\n\t" + "foo bar") + "foo bar") {
                        case "hello":
                            if (((3.14159 * ((float64(v4) - 3.14159) + (1e10 / 0.1))) * 3.14159) > 0.1) {
                                var v20 string = ("\n\t" + "hello")
                                _ = v20
                            } else {
                                v6 = 840
                                var v21 int64 = 9223372036854775807
                                _ = v21
                            }
                        }
                        v6 = 382
                    }
                    _ = v18
                }
                if true {
                    v22 := GenS1{F0: false, F1: -961, F2: v9}
                    fmt.Println(v9, v3, v2)
                    if true {
                        var v23 int16 = v6
                        var v24 uint64 = uint64(v2)
                        for fi25 := 0; fi25 < 4; fi25++ {
                            var v26 int16 = (v6 / ((v23) | 1))
                            var v27 uint16 = v2
                            if len(v1) > 0 {
                                v1[0] = 510
                            }
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                        _ = v23
                        _ = v24
                    } else {
                        v1 = append(v1, 4294967295)
                        var v28 bool = true
                        _ = v28
                    }
                    _ = v22
                }
            } else {
                for fi29 := 0; fi29 < 4; fi29++ {
                    v3 += 665
                    var v30 uint32 = 151
                    _ = fi29
                    _ = v30
                }
                if false {
                    v31 := GenS2{F0: (!(!true)), F1: 65535}
                    v7 += (v7 - (v7 / ((v7) | 1)))
                    fmt.Println(v7, v6, v5)
                    _ = v31
                }
            }
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    v4 += ((v4 / v4) - 0.0)
    v32 := GenS2{F0: (!false), F1: 65}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
