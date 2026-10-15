package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint64; F1 bool }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 201
    v3 := make(map[int]int)
    var v4 int = 4294967295
    v5 := []int{v4, (^(v4)), v4}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi6 := 0; fi6 < 1; fi6++ {
        if false {
            v1[(fi6 & (fi6 | v4))] = 42
            v4 = 1
            if ((v4 << ((376) & 63)) <= -590) {
                var v7 float64 = (0.1 / 1.5)
                _ = v7
            }
        }
        _ = fi6
    }
    if (127 == 1) {
        v2 -= (^(v2))
        v2 += v2
    } else {
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 int32 = -2147483648
            if len(v5) > 0 {
                v5[0] = fi8
            }
            var v10 int32 = (v9 + v9)
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    if false {
        var v11 int8 = -17
        for fi12 := 0; fi12 < 13; fi12++ {
            _ = len(v5)
            v2 = (v2 | v2)
            _ = fi12
        }
        var v13 int16 = -32768
        _ = v11
        _ = v13
    } else {
        var v14 uint16 = 32767
        for fi15 := 0; fi15 < 6; fi15++ {
            var v16 float64 = 1e10
            _ = fi15
            _ = v16
        }
        var v17 bool = ((v14 > 2) || (-79 < -123))
        _ = v14
        _ = v17
    }
    v18 := GenS1{F0: v2, F1: 331}
    switch (v4 >> ((4294967295) & 63)) {
    case -9223372036854775808:
        var v19 uint16 = 65535
        _ = v19
    }
    switch 42 {
    case 9223372036854775807:
        switch (v4 + (v4 - v4)) {
        case 4294967295:
            for fi20 := 0; fi20 < 12; fi20++ {
                if true {
                    v21 := GenS1{F0: (v2 | v2), F1: 42}
                    switch (-(v4)) {
                    case -1:
                        v2 -= (v2 << ((2) & 63))
                    }
                    _ = v21
                }
                if ((((-2.25 / ((float32(1e10) + 3.14159) * (0.1 / (((1.5 - -2.25) + ((1e10 * 1.5) / (1e10 / 1.0))) + 3.14159)))) - 1.5) - 1.5) > (0.0 * (0.0 - (0.0 - (0.0 + 1.5))))) {
                    switch (v4 >> ((799) & 63)) {
                    case 1:
                        v4 = (v4 / ((42) | 1))
                    case 734:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case 32768:
                        _ = len(v5)
                    }
                } else {
                    var v22 int64 = -464
                    _ = v22
                }
                _ = fi20
            }
        default:
            switch "hello" {
            case "foo bar":
                v23 := GenS1{F0: v2, F1: 442}
                _ = v23
            case "hello":
                if false {
                    v4 = v4
                    v4 = 128
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                } else {
                    if true {
                        if (56 == 100) {
                            var v24 int16 = 1000
                            v25 := GenS1{F0: (v2 & 555), F1: 4294967295}
                            var v26 uint32 = 2147483648
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        v27 := GenS1{F0: (v2 & (v2 / ((881) | 1))), F1: 2147483647}
                        switch (((("foo bar" + (("" + "hello") + "\n\t")) + "") + "foo bar") + "hello") {
                        case "hello":
                            for fi28 := 0; fi28 < 18; fi28++ {
                                for fi29 := 0; fi29 < 14; fi29++ {
                                    var v30 uint64 = v2
                                    _ = fi29
                                    _ = v30
                                }
                                var v31 uint16 = 255
                                v32 := GenS1{F0: (^(v2)), F1: 417}
                                _ = fi28
                                _ = v31
                                _ = v32
                            }
                        }
                        _ = v27
                    }
                    if (false && (true && false)) {
                        if true {
                            v33 := GenS1{F0: 1, F1: 4294967295}
                            _ = v33
                        } else {
                            v4 = 32768
                            switch "" {
                            case "hello":
                                if false {
                                    var v34 uint8 = 77
                                    for fi35 := 0; fi35 < 2; fi35++ {
                                        v36 := GenS1{F0: v2, F1: 1}
                                        var v37 float32 = float32(float64(1e10))
                                        _ = fi35
                                        _ = v36
                                        _ = v37
                                    }
                                    fmt.Println(v2, v4, v4)
                                    _ = v34
                                } else {
                                    v38 := GenS1{F0: v2, F1: 2147483647}
                                    v39 := GenS1{F0: 2, F1: 698}
                                    if true {
                                        v40 := GenS1{F0: (v2 << ((1) & 63)), F1: 375}
                                        _ = v40
                                    }
                                    _ = v38
                                    _ = v39
                                }
                            case "":
                                v4 = 42
                            case "foo bar":
                                var v41 uint8 = 0
                                _ = v41
                            default:
                                v3[v4] = v4
                            }
                            fmt.Println(v4, v2)
                        }
                    }
                    v42 := GenS1{F0: v2, F1: 1}
                    _ = v42
                }
            }
        }
    default:
        for fi43 := 0; fi43 < 15; fi43++ {
            v2 += v2
            v5 = append(v5, 32768)
            _ = fi43
        }
    }
    var v44 uint32 = 561
    delete(v1, v4)
    switch v4 {
    case -1:
        _ = len(v5)
    }
    v5 = append(v5, (v4 | (v4 % ((128) | 1))))
    v45 := GenS2{F0: (v2 * v2), F1: (true || (-652 < -86))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
}
