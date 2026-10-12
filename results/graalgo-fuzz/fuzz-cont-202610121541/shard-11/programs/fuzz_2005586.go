package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }
func genFn0() uint16 {
    var v5 uint8 = 255
    var v6 bool = true
    _ = v5
    _ = v6
    return 428
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = "foo bar"
    v3 := []uint32{120}
    var v4 int32 = 250
    _ = len(v3)
    v7 := GenS1{F0: 0, F1: 0}
    var v8 int16 = 42
    v2 = v2
    v8 = v8
    for fi9 := 0; fi9 < 7; fi9++ {
        if false {
            if (0 < -128) {
                for fi10 := 0; fi10 < 8; fi10++ {
                    _ = len(v1)
                    _ = fi10
                }
                v2 = (v2 + v2)
                if true {
                    v8 = v8
                    switch "\n\t" {
                    case "hello":
                        var v11 float64 = 0.1
                        _ = v11
                    case "foo bar":
                        fmt.Println(v2, v4)
                    default:
                        switch fi9 {
                        case -9223372036854775808:
                            if false {
                                var v12 string = "hello"
                                _ = v12
                            } else {
                                var v13 int64 = 1
                                var v14 int16 = (v8 % (((v8 * v8)) | 1))
                                _ = v13
                                _ = v14
                            }
                        }
                    }
                    switch 32767 {
                    case -762:
                        v4 = (v4 * 832)
                    default:
                        if ((!true) || (v2 != v2)) {
                            for fi15 := 0; fi15 < 12; fi15++ {
                                var v16 string = v2
                                _ = fi15
                                _ = v16
                            }
                            for fi17 := 0; fi17 < 2; fi17++ {
                                fmt.Println(fi9, v8, v2)
                                switch fi17 {
                                case 4294967295:
                                    switch fi17 {
                                    case -982:
                                        var v18 float64 = ((3.14159 + (1.5 + 1e10)) / 1e10)
                                        _ = v18
                                    }
                                case 1:
                                    switch (v2 + "hello") {
                                    case "hello":
                                        _ = len(v1)
                                    case "":
                                        v2 = v2
                                    case "\n\t":
                                        v19 := GenS1{F0: 262, F1: 961}
                                        _ = v19
                                    default:
                                        var v20 int64 = 42
                                        _ = v20
                                    }
                                }
                                _ = fi17
                            }
                            fmt.Println(v4, v4)
                        }
                    }
                } else {
                    for fi21 := 0; fi21 < 4; fi21++ {
                        if len(v3) > 0 {
                            v3[0] = 0
                        }
                        _ = fi21
                    }
                    switch "foo bar" {
                    case "":
                        var v22 uint16 = 65535
                        _ = v22
                    }
                    if true {
                        fmt.Println(v8)
                        v23 := GenS2{F0: 2}
                        var v24 int16 = (v8 >> ((4294967295) & 63))
                        _ = v23
                        _ = v24
                    } else {
                        fmt.Println(v4)
                        if (1 == 100) {
                            v4 += (v4 ^ v4)
                            if false {
                                fmt.Println(v4, fi9, v8)
                            } else {
                                var v25 int8 = 0
                                v8 -= v8
                                _ = v25
                            }
                            v2 = v2
                        }
                    }
                }
            } else {
                for fi26 := 0; fi26 < 1; fi26++ {
                    v27 := GenS1{F0: 18446744073709551615, F1: 209}
                    fmt.Println(fi26, fi9, v8)
                    _ = fi26
                    _ = v27
                }
            }
            if false {
                v28 := v1[2147483647]
                _ = v28
            }
        }
        var v29 int8 = -128
        var v30 int8 = -128
        _ = fi9
        _ = v29
        _ = v30
    }
    if true {
        v31 := GenS2{F0: 2}
        _ = v31
    }
    fmt.Println(v8, v4)
    var v32 bool = (true || false)
    switch 32767 {
    case -57:
        v4 = (v4 + v4)
    case 32767:
        if v32 {
            delete(v1, 353)
            v4 = -1
        }
    default:
        for fi33 := 0; fi33 < 7; fi33++ {
            v2 = (v2 + v2)
            _ = fi33
        }
    }
    var v34 string = v2
    _ = len(v3)
    v4 += (v4 + v4)
    v35 := GenS3{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v35
    _ = v4
    _ = v7
    _ = v8
}
