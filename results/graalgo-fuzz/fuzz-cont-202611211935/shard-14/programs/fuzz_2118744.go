package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int8; F1 bool }

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 32767
    v3 := []int{32768, 9223372036854775807, 0}
    var v4 uint64 = 4294967295
    v5 := GenS1{F0: -1}
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := GenS2{F0: 255, F1: -668, F2: 957}
        switch "a" {
        case "\n\t":
            var v8 float64 = 3.14159
            _ = v8
        case "foo bar":
            var v9 uint8 = 42
            _ = v9
        case "hello":
            if false {
                if true {
                    if true {
                        if len(v3) > 0 {
                            v3[0] = (^(fi6))
                        }
                    }
                    var v10 uint32 = 2147483648
                    _ = v10
                }
                var v11 uint32 = 702
                _ = v11
            }
        default:
            if false {
                v12 := v1[-2147483648]
                v1[-2147483648] = 128
                _ = v12
            } else {
                var v13 uint64 = v4
                for fi14 := 0; fi14 < 19; fi14++ {
                    _ = len(v1)
                    var v15 string = (("\n\t" + "hello") + (("hello" + "") + ""))
                    _ = fi14
                    _ = v15
                }
                v3 = append(v3, 416)
                _ = v13
            }
        }
        var v16 int8 = 127
        _ = fi6
        _ = v16
        _ = v7
    }
    _ = len(v1)
    switch "a" {
    case "":
        switch 32768 {
        case 42:
            v17 := GenS3{F0: 0, F1: 2147483647, F2: 512}
            _ = v17
        case 4294967295:
            for fi18 := 0; fi18 < 12; fi18++ {
                if (!(!true)) {
                    v19 := GenS3{F0: 42, F1: 75, F2: (v4 % (((v4 % ((826) | 1))) | 1))}
                    _ = v19
                } else {
                    if ((("hello" + "a") + "foo bar") <= ((("" + "") + "\n\t") + "hello")) {
                        var v20 uint = 1
                        _ = v20
                    } else {
                        for fi21 := 0; fi21 < 13; fi21++ {
                            if false {
                                v1[(fi18 >> ((18446744073709551615) & 63))] = (^(fi18))
                                v22 := v1[(^(fi18))]
                                v23 := GenS4{F0: int8(v4), F1: false}
                                _ = v22
                                _ = v23
                            }
                            if true {
                                fmt.Println(v4, v2)
                                fmt.Println(fi18, fi18, fi18)
                                var v24 uint = 65535
                                _ = v24
                            }
                            var v25 float32 = 3.14159
                            _ = fi21
                            _ = v25
                        }
                        v2 = v2
                    }
                }
                _ = fi18
            }
        }
    case "foo bar":
        for fi26 := 0; fi26 < 3; fi26++ {
            fmt.Println(v2, v2, v4)
            var v27 uint16 = 2
            _ = fi26
            _ = v27
        }
    case "\n\t":
        for fi28 := 0; fi28 < 8; fi28++ {
            var v29 uint64 = (v4 * v4)
            for fi30 := 0; fi30 < 8; fi30++ {
                var v31 int64 = 9223372036854775807
                _ = fi30
                _ = v31
            }
            _ = fi28
            _ = v29
        }
    }
    var v32 float64 = (0.0 - (-2.25 + 1e10))
    for fi33 := 0; fi33 < 4; fi33++ {
        var v34 uint = 0
        _ = fi33
        _ = v34
    }
    v1[4294967295] = -169
    for fi35 := 0; fi35 < 14; fi35++ {
        v32 -= v32
        _ = fi35
    }
    switch 127 {
    case 1:
        v2 -= (v2 & (v2 * (^(v2))))
    case 127:
        if len(v3) > 0 {
            v3[0] = -1
        }
    }
    fmt.Println(v2, v2, v32)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
