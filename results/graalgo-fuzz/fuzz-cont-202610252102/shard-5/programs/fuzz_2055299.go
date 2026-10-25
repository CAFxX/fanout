package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int8; F1 int64; F2 int32 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 int8 = 127
    v2 := []int{-2147483648, 286}
    var v3 int64 = 1
    var v4 string = "hello"
    var v5 uint8 = 19
    v6 := []int{-2147483648, 2147483647}
    var v7 int16 = 0
    var v8 int32 = 884
    var v9 int64 = 9223372036854775807
    switch v4 {
    case "foo bar":
        switch 2147483647 {
        case 127:
            v9 = v9
        case -9223372036854775808:
            var v10 uint64 = 4294967295
            _ = v10
        case 9223372036854775807:
            v11 := GenS1{F0: 0, F1: 42}
            _ = v11
        }
    case "":
        v9 = -9223372036854775808
    default:
        v7 = 32767
    }
    var v12 bool = (false && true)
    v1 += (v1 - v1)
    for fi13 := 0; fi13 < 5; fi13++ {
        _ = len(v6)
        switch (fi13 | (fi13 / ((fi13) | 1))) {
        case 42:
            var v14 uint = 4294967295
            _ = v14
        default:
            for fi15 := 0; fi15 < 17; fi15++ {
                _ = len(v2)
                _ = fi15
            }
        }
        v4 = ((v4 + v4) + v4)
        _ = fi13
    }
    var v16 uint16 = 0
    fmt.Println(v7, v12)
    var v17 float32 = 3.14159
    if false {
        v9 += (v9 >> ((1) & 63))
        if v12 {
            v17 = float32(1.5)
            for fi18 := 0; fi18 < 15; fi18++ {
                if len(v2) > 0 {
                    v2[0] = int(v7)
                }
                _ = fi18
            }
            v19 := GenS1{F0: 1, F1: v16}
            _ = v19
        }
    } else {
        if ((v17 - ((0.0 + 1e10) - (v17 / ((v17 * v17) * (((-2.25 - -2.25) / 1e10) * v17))))) == v17) {
            v4 = v4
            v20 := GenS2{F0: v1, F1: -15, F2: v8}
            switch 2147483647 {
            case 124:
                for fi21 := 0; fi21 < 10; fi21++ {
                    v9 = 813
                    _ = fi21
                }
            case 128:
                var v22 uint64 = 42
                _ = v22
            }
            _ = v20
        } else {
            v9 = 42
            if v12 {
                v6 = append(v6, 32768)
                fmt.Println(v16, v5, v16)
            } else {
                switch "hello" {
                case "a":
                    for fi23 := 0; fi23 < 2; fi23++ {
                        for fi24 := 0; fi24 < 12; fi24++ {
                            _ = len(v6)
                            v4 = v4
                            _ = fi24
                        }
                        var v25 bool = (!false)
                        if false {
                            if (v12 && (!v25)) {
                                v9 = 42
                            }
                            var v26 uint16 = 255
                            _ = v26
                        }
                        _ = fi23
                        _ = v25
                    }
                default:
                    for fi27 := 0; fi27 < 18; fi27++ {
                        v16 = (v16 & v16)
                        _ = fi27
                    }
                }
                var v28 uint8 = v5
                _ = v28
            }
            if v12 {
                var v29 int16 = v7
                v30 := GenS1{F0: 42, F1: (v16 / ((v16) | 1))}
                _ = v29
                _ = v30
            }
        }
        if (v16 > v16) {
            fmt.Println(v5, v16)
            v31 := GenS3{F0: 2147483647}
            v1 -= -84
            _ = v31
        }
        fmt.Println(v16)
    }
    if len(v6) > 0 {
        v6[0] = 4294967295
    }
    if false {
        v7 -= (v7 * v7)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
