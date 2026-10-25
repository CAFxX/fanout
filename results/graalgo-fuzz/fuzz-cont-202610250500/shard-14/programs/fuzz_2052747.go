package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 uint16; F1 bool; F2 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8; F1 string }

func main() {
    v1 := []uint32{2, 451, 2}
    var v2 uint = 853
    var v3 uint8 = 0
    var v4 string = "foo bar"
    var v5 float64 = 0.1
    v6 := make(map[string]string)
    v2 = 51
    switch "a" {
    case "foo bar":
        var v7 uint16 = 106
        _ = v7
    case "hello":
        var v8 string = (v4 + ("" + (v4 + ((v4 + v4) + ((v4 + "a") + "")))))
        _ = v8
    case "\n\t":
        v4 = v4
    default:
        v3 -= (v3 | (v3 | 42))
    }
    fmt.Println(v2, v3)
    v6[v4] = ("" + "a")
    switch 469 {
    case 4294967295:
        if true {
            var v9 float32 = float32(v5)
            _ = v9
        }
    default:
        v10 := GenS1{F0: 65535, F1: 42, F2: (v4 >= ("hello" + v4))}
        _ = v10
    }
    v5 += (v5 / (float64((1.5 * ((3.14159 - 0.1) + float32(0.0)))) / v5))
    var v11 int16 = 444
    v12 := GenS2{F0: v11}
    switch v4 {
    case "":
        v11 = v11
    case "hello":
        fmt.Println(v4)
    case "a":
        v13 := GenS3{F0: 99, F1: (v11 / ((v11) | 1))}
        _ = v13
    }
    var v14 int32 = 0
    for fi15 := 0; fi15 < 10; fi15++ {
        if false {
            fmt.Println(v2, v14, v3)
        } else {
            for fi16 := 0; fi16 < 10; fi16++ {
                if (100 < -1) {
                    v11 -= v11
                    switch -1 {
                    case 128:
                        if false {
                            if true {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v1 = append(v1, 921)
                                switch (fi15 - (fi16 - fi15)) {
                                case 0:
                                    if true {
                                        _ = len(v1)
                                    } else {
                                        var v17 uint8 = 128
                                        _ = v17
                                    }
                                case -132:
                                    if true {
                                        var v18 int32 = -1
                                        if len(v1) > 0 {
                                            v1[0] = 4294967295
                                        }
                                        _ = v18
                                    }
                                case 127:
                                    v19 := GenS3{F0: 1, F1: (v11 | -1)}
                                    v20 := GenS4{F0: 127, F1: false, F2: v14}
                                    fmt.Println(v11, v11, fi15)
                                    _ = v19
                                    _ = v20
                                }
                            }
                            v1 = append(v1, 891)
                            if false {
                                var v21 float32 = ((1.5 / 3.14159) - float32(float64(1e10)))
                                _ = v21
                            } else {
                                if false {
                                    v4 = (("foo bar" + "a") + v4)
                                }
                                v3 -= 2
                                var v22 uint64 = 4294967295
                                _ = v22
                            }
                        }
                    }
                    fmt.Println(v3)
                }
                _ = fi16
            }
        }
        var v23 int64 = -1
        _ = fi15
        _ = v23
    }
    v24 := GenS3{F0: 1, F1: v11}
    v25 := GenS5{F0: false}
    for fi26 := 0; fi26 < 13; fi26++ {
        var v27 int64 = -1
        var v28 int = fi26
        v2 += (^(v2))
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println(v3, v5)
    if (v2 > 1) {
        switch v4 {
        case "a":
            v29 := GenS6{F0: 90, F1: ((("\n\t" + ("\n\t" + v4)) + v4) + v4)}
            _ = v29
        }
        if false {
            v3 = (^(v3))
            for fi30 := 0; fi30 < 10; fi30++ {
                fmt.Println(v5, v2)
                var v31 int8 = -9
                _ = fi30
                _ = v31
            }
        }
    }
    for fi32 := 0; fi32 < 11; fi32++ {
        if (!(!false)) {
            var v33 uint32 = 0
            v34 := GenS2{F0: (v11 + 42)}
            var v35 int64 = -1
            _ = v33
            _ = v34
            _ = v35
        } else {
            v36 := GenS6{F0: 127, F1: (v4 + v4)}
            var v37 uint16 = 0
            v38 := GenS4{F0: uint16(v2), F1: (!false), F2: (v14 >> (((v2 * 65535)) & 63))}
            _ = v36
            _ = v37
            _ = v38
        }
        _ = fi32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
