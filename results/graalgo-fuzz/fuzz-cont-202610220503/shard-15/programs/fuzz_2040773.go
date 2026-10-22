package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []uint32{64, 2147483648, 2147483648}
    var v3 int32 = 42
    fmt.Println(v3)
    if len(v2) > 0 {
        v2[0] = 158
    }
    v3 += 1
    v1 += v1
    var v4 int32 = v3
    fmt.Println(v1, v3, v3)
    var v5 float64 = float64(0.0)
    switch 447 {
    case 42:
        if (42 < 42) {
            fmt.Println(v4, v1, v3)
        }
    case 9223372036854775807:
        v6 := GenS1{F0: v1}
        _ = v6
    case -64:
        if false {
            v4 -= (v3 % (((v3 / ((0) | 1))) | 1))
        } else {
            var v7 string = ""
            switch "foo bar" {
            case "foo bar":
                v4 = v4
            case "\n\t":
                if false {
                    v1 = (-(v1))
                    if (v7 == v7) {
                        var v8 int32 = v4
                        fmt.Println(v5)
                        fmt.Println(v7, v4, v4)
                        _ = v8
                    } else {
                        v9 := GenS1{F0: (-(v1))}
                        if false {
                            var v10 uint16 = 0
                            v1 -= v1
                            _ = v10
                        }
                        switch -413 {
                        case 4294967295:
                            for fi11 := 0; fi11 < 4; fi11++ {
                                for fi12 := 0; fi12 < 7; fi12++ {
                                    if false {
                                        var v13 int = (fi12 + fi12)
                                        v14 := GenS1{F0: 42}
                                        v1 = -9223372036854775808
                                        _ = v13
                                        _ = v14
                                    } else {
                                        var v15 int64 = 42
                                        fmt.Println(v3)
                                        v4 = v4
                                        _ = v15
                                    }
                                    _ = len(v2)
                                    _ = fi12
                                }
                                v3 = (v4 / ((92) | 1))
                                var v16 int = fi11
                                _ = fi11
                                _ = v16
                            }
                        case 128:
                            var v17 uint8 = 1
                            _ = v17
                        case -160:
                            v5 = v5
                        }
                        _ = v9
                    }
                    switch 0 {
                    case 4294967295:
                        var v18 float64 = ((((float64(0.1) / v5) * 1e10) * (float64(0.0) + (-2.25 / ((3.14159 / 1.5) - (1.5 - v5))))) - (v5 - v5))
                        _ = v18
                    }
                } else {
                    switch -9223372036854775808 {
                    case 718:
                        var v19 float64 = 0.1
                        _ = v19
                    case -262:
                        var v20 int8 = 0
                        _ = v20
                    case 9223372036854775807:
                        var v21 bool = (!((!(13 == 100)) || (0 != 128)))
                        _ = v21
                    }
                    var v22 uint64 = 42
                    _ = v22
                }
            case "":
                v7 = v7
            }
            _ = v7
        }
    default:
        for fi23 := 0; fi23 < 18; fi23++ {
            var v24 int32 = v3
            _ = fi23
            _ = v24
        }
    }
    v3 = v4
    if true {
        v5 = float64(float32(v5))
        for fi25 := 0; fi25 < 18; fi25++ {
            fmt.Println(v4, v5)
            for fi26 := 0; fi26 < 11; fi26++ {
                v27 := GenS1{F0: v1}
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
    }
    v5 -= v5
    v28 := GenS2{F0: v1, F1: (v1 & v1)}
    fmt.Println(v3)
    for fi29 := 0; fi29 < 20; fi29++ {
        for fi30 := 0; fi30 < 15; fi30++ {
            v31 := GenS3{F0: 282}
            _ = fi30
            _ = v31
        }
        if true {
            if (!false) {
                var v32 uint64 = 70
                fmt.Println(fi29, v3)
                _ = v32
            }
            var v33 int = 32768
            for fi34 := 0; fi34 < 12; fi34++ {
                if (v1 != v1) {
                    _ = len(v2)
                    var v35 int = 4294967295
                    _ = v35
                } else {
                    fmt.Println(v33, v33)
                    v4 -= -2147483648
                }
                switch "a" {
                case "":
                    var v36 float64 = v5
                    _ = v36
                }
                fmt.Println(fi34, v1)
                _ = fi34
            }
            _ = v33
        }
        _ = fi29
    }
    v2 = append(v2, 1)
    var v37 int64 = (v1 ^ v1)
    v38 := GenS4{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
