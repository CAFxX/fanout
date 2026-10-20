package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 int32; F2 int32 }

func main() {
    var v1 uint16 = 0
    v2 := []uint32{2, 197, 2147483648}
    v3 := []int{}
    if false {
        v4 := GenS1{F0: 255}
        var v5 string = ""
        _ = v4
        _ = v5
    } else {
        if (0.1 <= (1e10 / 0.1)) {
            for fi6 := 0; fi6 < 14; fi6++ {
                fmt.Println(v1, v1)
                _ = fi6
            }
            v1 += 0
            var v7 uint = 507
            _ = v7
        } else {
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 int32 = 285
                _ = fi8
                _ = v9
            }
            switch 910 {
            case 0:
                v10 := GenS2{F0: 100}
                _ = v10
            }
        }
    }
    v11 := GenS1{F0: 255}
    var v12 int64 = 0
    v1 -= 0
    var v13 uint8 = 2
    v14 := GenS1{F0: uint8(v1)}
    var v15 int16 = -412
    fmt.Println(v15, v13, v12)
    var v16 float64 = 0.1
    if ((v12 ^ (v12 << ((936) & 63))) <= v12) {
        if (1e10 <= 3.14159) {
            if (-1 == -95) {
                switch 127 {
                case 32768:
                    v16 = 1e10
                case 604:
                    v17 := GenS1{F0: 2}
                    _ = v17
                default:
                    v18 := GenS1{F0: 42}
                    _ = v18
                }
            }
            switch 551 {
            case 805:
                for fi19 := 0; fi19 < 16; fi19++ {
                    switch "\n\t" {
                    case "foo bar":
                        v12 -= -9223372036854775808
                    default:
                        var v20 uint64 = 2
                        _ = v20
                    }
                    _ = len(v2)
                    if true {
                        if true {
                            var v21 uint64 = 729
                            _ = v21
                        }
                        switch (fi19 ^ (fi19 & (fi19 + (fi19 - (fi19 * 128))))) {
                        case -291:
                            v22 := GenS1{F0: (^(v13))}
                            _ = v22
                        default:
                            if true {
                                if (!true) {
                                    var v23 uint = 506
                                    v12 = (v12 + v12)
                                    _ = v23
                                }
                            }
                        }
                        v24 := GenS3{F0: v15, F1: -244, F2: -715}
                        _ = v24
                    }
                    _ = fi19
                }
            case -583:
                if true {
                    if (false && (false && ((^(v12)) > v12))) {
                        var v25 uint = 42
                        _ = v25
                    }
                } else {
                    if true {
                        v2 = append(v2, 4294967295)
                    } else {
                        v16 -= 1e10
                        v26 := GenS3{F0: 1, F1: -2147483648, F2: 2147483647}
                        if false {
                            switch 977 {
                            case 128:
                                if false {
                                    v16 = (v16 / ((v16 + float64(3.14159)) / float64(1e10)))
                                }
                            default:
                                v1 -= v1
                            }
                            switch -479 {
                            case 1:
                                fmt.Println(v12, v13, v12)
                            }
                        }
                        _ = v26
                    }
                }
            default:
                v12 = int64(v15)
            }
        }
        v13 = 2
    }
    fmt.Println(v1, v1, v16)
    if true {
        v12 -= v12
    } else {
        if (!true) {
            var v27 uint64 = 18446744073709551615
            fmt.Println(v15, v1, v13)
            var v28 bool = false
            _ = v27
            _ = v28
        }
        fmt.Println(v12)
    }
    var v29 float32 = (((float32(v16) * 0.1) - 0.0) / 1.5)
    for fi30 := 0; fi30 < 6; fi30++ {
        for fi31 := 0; fi31 < 11; fi31++ {
            var v32 uint32 = 42
            fmt.Println(v32, v16, fi31)
            var v33 int32 = -1
            _ = fi31
            _ = v32
            _ = v33
        }
        fmt.Println(v16, v13)
        if ((-(v12)) != (^(v12))) {
            v34 := GenS1{F0: 145}
            for fi35 := 0; fi35 < 12; fi35++ {
                v12 = -461
                _ = fi35
            }
            _ = v34
        } else {
            if true {
                var v36 uint32 = 2147483648
                _ = v36
            }
            v16 = v16
            v15 = v15
        }
        _ = fi30
    }
    fmt.Println(v13)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v29
    _ = v3
}
