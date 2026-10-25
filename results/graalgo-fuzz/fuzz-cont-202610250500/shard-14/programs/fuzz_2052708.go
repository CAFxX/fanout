package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 bool; F1 int16; F2 uint64 }

func main() {
    var v1 int64 = 732
    var v2 uint32 = 500
    var v3 uint32 = (v2 * v2)
    v3 = 4294967295
    switch "a" {
    case "":
        var v4 int32 = 1
        _ = v4
    case "hello":
        v5 := []int{}
        _ = v5
    case "\n\t":
        if (v1 < v1) {
            v6 := []int{42, 128}
            v7 := GenS1{F0: (v2 + v2)}
            _ = v6
            _ = v7
        } else {
            switch "" {
            case "\n\t":
                var v8 uint = 65535
                _ = v8
            case "a":
                v3 += (v3 ^ v3)
            }
        }
    default:
        v2 -= 42
    }
    if false {
        for fi9 := 0; fi9 < 20; fi9++ {
            if ("foo bar" > "hello") {
                for fi10 := 0; fi10 < 6; fi10++ {
                    v11 := []int{}
                    if true {
                        switch "foo bar" {
                        case "foo bar":
                            var v12 int32 = -1
                            _ = v12
                        case "a":
                            fmt.Println(v1)
                        case "hello":
                            if (true && false) {
                                var v13 float64 = (0.1 - 0.1)
                                var v14 int8 = -128
                                v15 := GenS2{F0: 816, F1: -2147483648}
                                v1 += v1
                                _ = v13
                                _ = v14
                                _ = v15
                            }
                        default:
                            switch "hello" {
                            case "":
                                var v16 float32 = 0.1
                                _ = v16
                            case "hello":
                                if false {
                                    var v17 int64 = v1
                                    v17 = v1
                                    var v18 uint64 = 1
                                    _ = v17
                                    _ = v18
                                }
                            }
                        }
                    } else {
                        v1 -= v1
                    }
                    if (598 == 1) {
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                    }
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    if ((3.14159 / 0.1) < 1.5) {
                        if (false || (-700 <= 0)) {
                            v19 := GenS2{F0: -2147483648, F1: 2147483647}
                            v11 = append(v11, int(v1))
                            v1 -= v1
                            _ = v19
                        }
                    } else {
                        v2 = v2
                    }
                    _ = fi10
                    _ = v11
                }
            }
            fmt.Println(v1, v1, v3)
            v20 := GenS3{F0: true, F1: -965, F2: 42}
            _ = fi9
            _ = v20
        }
        if (724 == 1000) {
            var v21 uint8 = 233
            _ = v21
        }
    } else {
        v2 += v2
        if (393 <= 74) {
            var v22 bool = ((42 <= 873) && false)
            _ = v22
        } else {
            for fi23 := 0; fi23 < 19; fi23++ {
                for fi24 := 0; fi24 < 7; fi24++ {
                    v25 := GenS2{F0: 0, F1: -218}
                    switch "hello" {
                    case "hello":
                        var v26 int64 = (v1 - v1)
                        _ = v26
                    }
                    var v27 int64 = v1
                    _ = fi24
                    _ = v25
                    _ = v27
                }
                v28 := []int{128, (fi23 & (fi23 >> ((896) & 63)))}
                v3 -= (v2 | 939)
                _ = fi23
                _ = v28
            }
        }
        for fi29 := 0; fi29 < 5; fi29++ {
            switch "\n\t" {
            case "\n\t":
                v2 += (v2 - 748)
            }
            switch "" {
            case "a":
                var v30 uint = 4294967295
                _ = v30
            default:
                if false {
                    if false {
                        v3 = (v2 ^ (^(v3)))
                        fmt.Println(v1, v3)
                    } else {
                        v1 += -389
                        var v31 uint32 = (v2 | v3)
                        _ = v31
                    }
                } else {
                    v1 += (v1 + (^(v1)))
                }
            }
            _ = fi29
        }
    }
    v2 -= (v3 - v3)
    switch 1 {
    case 865:
        v3 -= 38
    case -1:
        v32 := []int{4294967295, 784, 9223372036854775807, 128}
        _ = v32
    }
    for fi33 := 0; fi33 < 11; fi33++ {
        v3 = v3
        _ = fi33
    }
    switch -51 {
    case 128:
        var v34 float64 = -2.25
        _ = v34
    }
    var v35 int16 = -32768
    for fi36 := 0; fi36 < 14; fi36++ {
        v3 += (^(v2))
        v37 := []int{(-(fi36)), 259, -9223372036854775808}
        _ = len(v37)
        _ = fi36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
}
