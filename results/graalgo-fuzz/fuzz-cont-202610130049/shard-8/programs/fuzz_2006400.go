package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 int16; F2 uint }

func main() {
    v1 := []int{-2147483648}
    v2 := make(map[uint32]int)
    var v3 float64 = 0.1
    var v4 string = "a"
    var v5 float64 = 1e10
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 string = v4
        var v8 int32 = -484
        _ = fi6
        _ = v7
        _ = v8
    }
    if true {
        v4 = v4
        switch 32768 {
        case -522:
            delete(v2, 2147483647)
        default:
            delete(v2, 1)
        }
        var v9 uint64 = 18446744073709551615
        _ = v9
    } else {
        var v10 uint8 = 37
        fmt.Println(v5, v10)
        fmt.Println(v10, v5, v3)
        _ = v10
    }
    var v11 uint16 = 1
    v3 -= 0.0
    var v12 float32 = ((1.5 * 1.5) * 1e10)
    if (127 <= 796) {
        v3 = 0.0
        for fi13 := 0; fi13 < 1; fi13++ {
            var v14 int = (fi13 + int(v11))
            _ = fi13
            _ = v14
        }
    }
    switch 128 {
    case 4294967295:
        if (-1 < -1) {
            v15 := GenS1{F0: (v11 | v11), F1: 4294967295}
            _ = v15
        }
    case -1:
        fmt.Println(v5, v5, v4)
    case 128:
        delete(v2, 267)
    default:
        var v16 int64 = 42
        _ = v16
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        for fi18 := 0; fi18 < 12; fi18++ {
            var v19 int8 = 42
            fmt.Println(v5, v12, v3)
            fmt.Println(v3, v4, fi18)
            if true {
                fmt.Println(fi18, fi18, v12)
                switch fi18 {
                case -184:
                    _ = len(v2)
                case -86:
                    if (false || false) {
                        var v20 string = v4
                        var v21 float32 = (v12 - v12)
                        fmt.Println(v19)
                        var v22 int32 = 982
                        if false {
                            if true {
                                var v23 string = v4
                                if true {
                                    v3 += v5
                                    v4 = ("" + v20)
                                    v4 = ((v23 + "hello") + "")
                                } else {
                                    fmt.Println(fi18, v3)
                                    fmt.Println(v12, v21, v4)
                                }
                                _ = v23
                            }
                        }
                        _ = v20
                        _ = v21
                        _ = v22
                    } else {
                        if true {
                            var v24 uint = 65535
                            _ = v24
                        }
                        if true {
                            delete(v2, 283)
                            fmt.Println(v19)
                            v11 = 0
                        } else {
                            v4 = v4
                        }
                        v25 := GenS2{F0: false}
                        _ = v25
                    }
                default:
                    if true {
                        var v26 int8 = (v19 + v19)
                        v3 = -2.25
                        var v27 int8 = 119
                        _ = v26
                        _ = v27
                    }
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = fi18
            _ = v19
        }
        switch "foo bar" {
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "hello":
            switch v4 {
            case "hello":
                _ = len(v2)
            default:
                switch (fi17 + fi17) {
                case 948:
                    _ = len(v1)
                case 42:
                    fmt.Println(v5, v3, v4)
                }
            }
        }
        if ((0 <= 821) && false) {
            v28 := GenS1{F0: (v11 + (v11 % ((65535) | 1))), F1: 0}
            fmt.Println(v5, v5)
            _ = v28
        }
        _ = fi17
    }
    if true {
        switch -9223372036854775808 {
        case 807:
            fmt.Println(v11)
        }
        v3 -= ((0.1 - ((-2.25 + float64(0.0)) * v5)) + v3)
    }
    delete(v2, 2147483647)
    if ((((v4 + ("\n\t" + (("foo bar" + "a") + "hello"))) + ("\n\t" + (v4 + v4))) + "a") != "") {
        v29 := GenS3{F0: 0, F1: -55, F2: 18446744073709551615}
        _ = v29
    }
    delete(v2, 168)
    _ = len(v1)
    for fi30 := 0; fi30 < 8; fi30++ {
        delete(v2, 2147483648)
        var v31 int32 = -541
        _ = fi30
        _ = v31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
