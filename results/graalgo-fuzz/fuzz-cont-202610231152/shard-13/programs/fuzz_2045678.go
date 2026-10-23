package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 string; F2 int8 }
type GenS3 struct { F0 uint32; F1 uint64 }

func main() {
    var v1 bool = false
    v2 := make(map[uint32]bool)
    var v3 string = "\n\t"
    var v4 int64 = -758
    var v5 string = v3
    v6 := []int{4294967295}
    switch -9223372036854775808 {
    case 1:
        switch 4294967295 {
        case 42:
            var v7 float64 = float64(3.14159)
            _ = v7
        case -2147483648:
            v8 := v2[2]
            _ = v8
        case 1:
            fmt.Println(v1, v1)
        }
    }
    v9 := GenS1{F0: 1}
    if (-771 <= -822) {
        v5 = v5
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        switch ("foo bar" + "") {
        case "foo bar":
            var v11 int8 = 42
            _ = v11
        case "hello":
            switch v5 {
            case "foo bar":
                var v12 int32 = -429
                _ = v12
            default:
                switch v3 {
                case "foo bar":
                    fmt.Println(fi10)
                case "\n\t":
                    var v13 int = (fi10 - (fi10 / (((fi10 & fi10)) | 1)))
                    _ = v13
                case "a":
                    if (32767 != 1) {
                        switch "hello" {
                        case "foo bar":
                            var v14 uint32 = 42
                            _ = v14
                        case "hello":
                            for fi15 := 0; fi15 < 8; fi15++ {
                                if false {
                                    var v16 uint64 = 2
                                    _ = v16
                                } else {
                                    var v17 float32 = -2.25
                                    v18 := GenS2{F0: 1000, F1: v5, F2: 0}
                                    _ = v17
                                    _ = v18
                                }
                                _ = fi15
                            }
                        case "\n\t":
                            var v19 int = fi10
                            _ = v19
                        default:
                            fmt.Println(v3, v3, v1)
                        }
                        if v1 {
                            if ((0.1 * 1.5) < 0.1) {
                                for fi20 := 0; fi20 < 7; fi20++ {
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v21 int64 = (v4 + (v4 + v4))
                                        _ = v21
                                    case "":
                                        if (((float32((1e10 + 1e10)) + 0.0) / (((-2.25 + 1e10) / (((0.1 + 1.5) / 1e10) + 0.1)) - 0.0)) <= -2.25) {
                                            var v22 uint8 = 2
                                            var v23 bool = v1
                                            v4 = v4
                                            fmt.Println(v3, fi10)
                                            _ = v22
                                            _ = v23
                                        }
                                    case "hello":
                                        v5 = "foo bar"
                                        v3 = v3
                                        if (!(!(990 < 2))) {
                                            v24 := GenS3{F0: 42, F1: 0}
                                            switch fi10 {
                                            case 1:
                                                var v25 int16 = -32768
                                                _ = v25
                                            }
                                            var v26 uint = 255
                                            _ = v24
                                            _ = v26
                                        }
                                    default:
                                        v27 := GenS2{F0: -796, F1: (v5 + (v3 + v5)), F2: 0}
                                        _ = v27
                                    }
                                    _ = fi20
                                }
                            }
                            if true {
                                v1 = (-26 != 127)
                            } else {
                                var v28 string = (v5 + v3)
                                for fi29 := 0; fi29 < 11; fi29++ {
                                    if false {
                                        var v30 int = (fi10 * (fi10 + 956))
                                        _ = v30
                                    }
                                    _ = fi29
                                }
                                var v31 int64 = v4
                                _ = v28
                                _ = v31
                            }
                            var v32 float32 = 1e10
                            _ = v32
                        } else {
                            fmt.Println(v1, v1)
                        }
                        var v33 int = (fi10 + 0)
                        _ = v33
                    }
                }
            }
        default:
            if v1 {
                if false {
                    if (127 > 128) {
                        for fi34 := 0; fi34 < 16; fi34++ {
                            v4 += v4
                            _ = fi34
                        }
                    }
                    if len(v6) > 0 {
                        v6[0] = fi10
                    }
                }
                for fi35 := 0; fi35 < 8; fi35++ {
                    v2[0] = (560 > 2)
                    _ = fi35
                }
                v6 = append(v6, -731)
            }
        }
        switch 128 {
        case 42:
            v1 = (int32(fi10) <= -121)
        }
        v4 += 9223372036854775807
        _ = fi10
    }
    v5 = v5
    v36 := GenS2{F0: -32768, F1: "hello", F2: 0}
    v5 = v3
    if len(v6) > 0 {
        v6[0] = -2147483648
    }
    v37 := GenS1{F0: 0}
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
