package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 int32; F2 uint64 }

func main() {
    var v1 uint64 = 274
    var v2 float64 = float64(0.1)
    var v3 int = 968
    var v4 int32 = 0
    fmt.Println(v4, v1)
    v5 := GenS1{F0: 512, F1: 255}
    if true {
        for fi6 := 0; fi6 < 17; fi6++ {
            v7 := []int{128, (v3 & 391)}
            v8 := GenS2{F0: 92}
            var v9 int64 = 0
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        var v10 int64 = -9223372036854775808
        var v11 bool = true
        _ = v10
        _ = v11
    }
    v1 += 0
    v3 += v3
    v12 := []int{0, -972}
    for fi13 := 0; fi13 < 3; fi13++ {
        if true {
            v12 = append(v12, (v3 / ((-931) | 1)))
        }
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = fi13
    }
    if true {
        var v14 bool = (!false)
        _ = v14
    } else {
        _ = len(v12)
    }
    if len(v12) > 0 {
        v12[0] = 851
    }
    var v15 uint16 = 566
    v16 := GenS2{F0: 42}
    var v17 int16 = -32768
    v2 -= (1e10 - v2)
    if (!false) {
        v17 = v17
        for fi18 := 0; fi18 < 5; fi18++ {
            if ((!false) || (-9223372036854775808 >= 9223372036854775807)) {
                v2 -= v2
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            } else {
                v2 = ((v2 * 1e10) - (float64(0.1) - (v2 / (v2 + v2))))
                if true {
                    for fi19 := 0; fi19 < 16; fi19++ {
                        switch fi18 {
                        case -9223372036854775808:
                            switch "\n\t" {
                            case "foo bar":
                                var v20 int32 = v4
                                _ = v20
                            case "hello":
                                v17 = v17
                            case "\n\t":
                                var v21 int8 = 127
                                _ = v21
                            default:
                                var v22 int16 = (v17 - v17)
                                _ = v22
                            }
                        }
                        v1 = v1
                        _ = fi19
                    }
                } else {
                    if (-1 != 42) {
                        v23 := GenS3{F0: 42}
                        if ("a" == "hello") {
                            v4 = v4
                        } else {
                            v24 := GenS1{F0: 787, F1: 52}
                            _ = len(v12)
                            if false {
                                var v25 bool = (v15 < (^(v15)))
                                _ = v25
                            } else {
                                for fi26 := 0; fi26 < 11; fi26++ {
                                    v27 := GenS4{F0: 1, F1: int32(v15), F2: v1}
                                    v2 -= v2
                                    v12 = append(v12, v3)
                                    _ = fi26
                                    _ = v27
                                }
                            }
                            _ = v24
                        }
                        v4 -= (v4 + v4)
                        _ = v23
                    }
                    var v28 int8 = 42
                    if false {
                        var v29 int32 = v4
                        _ = v29
                    }
                    _ = v28
                }
                fmt.Println(v17, v15, fi18)
            }
            fmt.Println(v17, v4)
            _ = fi18
        }
        fmt.Println(v2)
    }
    var v30 uint16 = (v15 + 405)
    v15 = (v15 * (v15 * v15))
    if (false && true) {
        var v31 int8 = 100
        var v32 float32 = 1e10
        _ = v31
        _ = v32
    } else {
        if len(v12) > 0 {
            v12[0] = 42
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
