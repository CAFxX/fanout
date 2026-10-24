package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 bool }

func main() {
    var v1 uint64 = 984
    var v2 uint64 = v1
    var v3 float32 = (0.1 / 1.5)
    var v4 int64 = 9223372036854775807
    var v5 int16 = -265
    v5 -= v5
    fmt.Println(v5, v4, v2)
    for fi6 := 0; fi6 < 8; fi6++ {
        v5 += v5
        v3 += (v3 * float32(1.5))
        _ = fi6
    }
    v7 := []int{32768}
    v3 += v3
    var v8 int = 128
    fmt.Println(v4, v3)
    if true {
        fmt.Println(v3)
        if true {
            switch -9223372036854775808 {
            case 128:
                if ("hello" != "a") {
                    var v9 int16 = -794
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 11; fi10++ {
                        if true {
                            var v11 uint = 65535
                            var v12 uint = (^(v11))
                            _ = v11
                            _ = v12
                        }
                        var v13 uint8 = 127
                        _ = fi10
                        _ = v13
                    }
                    fmt.Println(v3, v1, v8)
                }
            case 275:
                if len(v7) > 0 {
                    v7[0] = 0
                }
            case -2147483648:
                if len(v7) > 0 {
                    v7[0] = (v8 + v8)
                }
            default:
                if true {
                    for fi14 := 0; fi14 < 13; fi14++ {
                        fmt.Println(v2, v3, v3)
                        _ = fi14
                    }
                } else {
                    fmt.Println(v2, v3, v4)
                    for fi15 := 0; fi15 < 4; fi15++ {
                        v16 := GenS1{F0: 377}
                        switch fi15 {
                        case 0:
                            for fi17 := 0; fi17 < 15; fi17++ {
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                v18 := GenS1{F0: 870}
                                v4 += (v4 | v4)
                                _ = fi17
                                _ = v18
                            }
                        case 4294967295:
                            v1 = v1
                        case -2147483648:
                            if false {
                                v19 := GenS1{F0: 702}
                                if true {
                                    v7 = append(v7, v8)
                                }
                                _ = v19
                            }
                        default:
                            for fi20 := 0; fi20 < 9; fi20++ {
                                var v21 int = fi15
                                _ = fi20
                                _ = v21
                            }
                        }
                        fmt.Println(v3, fi15)
                        _ = fi15
                        _ = v16
                    }
                    v7 = append(v7, v8)
                }
            }
        } else {
            var v22 int32 = 0
            _ = v22
        }
        if (true || false) {
            v4 -= (v4 << ((255) & 63))
        }
    } else {
        fmt.Println(v1, v2)
        v5 += v5
        if (false && (v1 == v2)) {
            switch (-(v8)) {
            case -2147483648:
                fmt.Println(v4, v8, v8)
            case -1:
                switch "" {
                case "foo bar":
                    for fi23 := 0; fi23 < 4; fi23++ {
                        v24 := GenS2{F0: (v4 / ((v4) | 1)), F1: ((v1 + v1) != (^(v1)))}
                        switch (v8 * fi23) {
                        case 158:
                            fmt.Println(v2)
                        case 127:
                            fmt.Println(v5, v5)
                        case -9223372036854775808:
                            var v25 float32 = float32(0.0)
                            _ = v25
                        }
                        _ = fi23
                        _ = v24
                    }
                }
            case 32767:
                fmt.Println(v8)
            }
            switch ("hello" + "hello") {
            case "foo bar":
                fmt.Println(v3, v8)
            case "a":
                var v26 int = 2147483647
                _ = v26
            default:
                var v27 bool = false
                _ = v27
            }
            v28 := GenS1{F0: 2}
            _ = v28
        }
    }
    v8 -= v8
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v29 int8 = 127
    if true {
        var v30 uint16 = 947
        var v31 int = (v8 ^ (v8 / ((v8) | 1)))
        v29 += v29
        _ = v30
        _ = v31
    } else {
        switch ((((("hello" + "a") + "hello") + "\n\t") + "foo bar") + ("" + "")) {
        case "hello":
            v7 = append(v7, (-(v8)))
        case "foo bar":
            switch 951 {
            case 127:
                if len(v7) > 0 {
                    v7[0] = (v8 - (v8 | 636))
                }
            default:
                var v32 float32 = float32(((0.0 - 1e10) / float64(v3)))
                _ = v32
            }
        case "a":
            _ = len(v7)
        default:
            var v33 int = v8
            _ = v33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
