package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 uint32 }

func main() {
    var v1 float64 = float64(3.14159)
    var v2 uint8 = 107
    v3 := []int{128}
    var v4 string = "a"
    var v5 uint = 0
    v6 := []uint32{}
    var v7 uint8 = (v2 << ((v5) & 63))
    v1 = 0.1
    switch (v4 + (v4 + (v4 + v4))) {
    case "":
        if false {
            v8 := GenS1{F0: 42, F1: "\n\t", F2: 32767}
            _ = v8
        } else {
            var v9 uint8 = (v2 - (v7 * v7))
            var v10 uint16 = 255
            var v11 uint16 = v10
            _ = v10
            _ = v11
            _ = v9
        }
    case "foo bar":
        var v12 string = v4
        _ = v12
    case "a":
        for fi13 := 0; fi13 < 4; fi13++ {
            var v14 uint64 = 42
            if (true && false) {
                if len(v6) > 0 {
                    v6[0] = 2147483648
                }
                v7 = 255
                v4 = (v4 + "\n\t")
            } else {
                for fi15 := 0; fi15 < 3; fi15++ {
                    var v16 int8 = -23
                    switch ((((v4 + "a") + "\n\t") + "") + v4) {
                    case "hello":
                        var v17 int64 = 1
                        _ = v17
                    case "":
                        v18 := GenS2{F0: uint64(v5)}
                        _ = v18
                    case "a":
                        var v19 int64 = 9223372036854775807
                        _ = v19
                    }
                    if true {
                        fmt.Println(v2, fi15, fi13)
                    } else {
                        v7 = (v2 >> ((v5) & 63))
                        var v20 string = (v4 + v4)
                        var v21 float32 = float32(v1)
                        _ = v20
                        _ = v21
                    }
                    _ = fi15
                    _ = v16
                }
                for fi22 := 0; fi22 < 16; fi22++ {
                    if false {
                        v14 = (v14 ^ (v14 / (((^(v14))) | 1)))
                        if len(v3) > 0 {
                            v3[0] = fi13
                        }
                        var v23 uint16 = 1
                        _ = v23
                    } else {
                        if ((!false) && false) {
                            var v24 int16 = 42
                            switch (("" + (v4 + "")) + v4) {
                            case "a":
                                v24 = (v24 ^ 156)
                            case "\n\t":
                                switch fi13 {
                                case 0:
                                    if true {
                                        var v25 uint64 = (v14 << (((v5 & (v5 ^ (v5 << ((v5) & 63))))) & 63))
                                        v1 += v1
                                        _ = v25
                                    } else {
                                        switch v4 {
                                        case "\n\t":
                                            var v26 uint16 = 745
                                            _ = v26
                                        default:
                                            v27 := GenS2{F0: v14}
                                            _ = v27
                                        }
                                        var v28 int16 = (v24 ^ 453)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v29 uint8 = v7
                                        _ = v28
                                        _ = v29
                                    }
                                default:
                                    v3 = append(v3, (fi13 & (fi13 ^ 2147483647)))
                                }
                            default:
                                fmt.Println(v24, v7, v7)
                                v3 = append(v3, (fi13 / (((fi22 >> ((4294967295) & 63))) | 1)))
                            }
                            v14 = (^(v14))
                            _ = v24
                        } else {
                            fmt.Println(v4)
                        }
                    }
                    _ = fi22
                }
                for fi30 := 0; fi30 < 14; fi30++ {
                    fmt.Println(v5)
                    _ = len(v3)
                    var v31 int = fi13
                    _ = fi30
                    _ = v31
                }
            }
            _ = fi13
            _ = v14
        }
    }
    var v32 int8 = 100
    if len(v3) > 0 {
        v3[0] = 32767
    }
    if false {
        v32 = (v32 + (v32 % (((v32 / (((v32 + (v32 ^ 42))) | 1))) | 1)))
    } else {
        if true {
            for fi33 := 0; fi33 < 13; fi33++ {
                var v34 int32 = 638
                for fi35 := 0; fi35 < 8; fi35++ {
                    var v36 bool = (42 > 0)
                    v37 := GenS1{F0: 42, F1: v4, F2: -21}
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                if len(v6) > 0 {
                    v6[0] = 0
                }
                _ = fi33
                _ = v34
            }
            fmt.Println(v2)
        }
        if false {
            var v38 int64 = 42
            _ = v38
        } else {
            if true {
                var v39 float32 = 1e10
                _ = v39
            }
            v40 := GenS3{F0: 1, F1: 35}
            _ = v40
        }
    }
    fmt.Println(v4, v32)
    var v41 uint16 = 0
    var v42 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v7
}
