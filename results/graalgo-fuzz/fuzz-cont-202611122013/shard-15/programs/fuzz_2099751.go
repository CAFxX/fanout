package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int32; F1 int32; F2 bool }
type GenS3 struct { F0 int32; F1 string; F2 int32 }

func main() {
    v1 := make(map[string]string)
    var v2 float32 = 3.14159
    var v3 int64 = 9223372036854775807
    var v4 int16 = 304
    var v5 string = "a"
    v6 := []string{}
    v7 := GenS1{F0: -1, F1: 167}
    for fi8 := 0; fi8 < 14; fi8++ {
        v1["a"] = v5
        _ = fi8
    }
    switch v5 {
    case "hello":
        v2 = (v2 / 0.1)
    case "\n\t":
        var v9 uint8 = 121
        _ = v9
    case "foo bar":
        v2 -= (1.5 / v2)
    default:
        v4 = 1000
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 bool = (4294967295 == 42)
        var v12 int64 = v3
        _ = fi10
        _ = v11
        _ = v12
    }
    v3 += -1
    v13 := GenS2{F0: 0, F1: -1, F2: false}
    if true {
        v5 = v5
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 uint = 423
            v3 = 9223372036854775807
            _ = fi14
            _ = v15
        }
    } else {
        if (!false) {
            v2 -= v2
        } else {
            fmt.Println(v4, v4)
            var v16 uint32 = 946
            _ = v16
        }
    }
    v4 += v4
    fmt.Println(v4, v3)
    switch v5 {
    case "a":
        var v17 uint16 = 1
        _ = v17
    default:
        v4 += (v4 ^ -32768)
    }
    if false {
        switch -790 {
        case 0:
            if true {
                if (329 == 255) {
                    var v18 int64 = (v3 + (v3 | (v3 % ((v3) | 1))))
                    _ = v18
                } else {
                    v19 := GenS3{F0: 1, F1: v5, F2: 0}
                    if true {
                        switch ("a" + "\n\t") {
                        case "\n\t":
                            switch (v5 + "foo bar") {
                            case "":
                                for fi20 := 0; fi20 < 4; fi20++ {
                                    v21 := v1[v5]
                                    for fi22 := 0; fi22 < 18; fi22++ {
                                        fmt.Println(v5, v3, v3)
                                        if len(v6) > 0 {
                                            v6[0] = v21
                                        }
                                        _ = fi22
                                    }
                                    fmt.Println(v2, v4)
                                    _ = fi20
                                    _ = v21
                                }
                            case "foo bar":
                                var v23 int64 = v3
                                _ = v23
                            case "\n\t":
                                fmt.Println(v2, v5, v5)
                            default:
                                _ = len(v6)
                            }
                        default:
                            for fi24 := 0; fi24 < 13; fi24++ {
                                v25 := v1[((("hello" + v5) + "foo bar") + ("foo bar" + v5))]
                                _ = fi24
                                _ = v25
                            }
                        }
                        if (32767 == 65535) {
                            switch (("a" + (v5 + v5)) + "a") {
                            case "\n\t":
                                delete(v1, v5)
                            case "a":
                                v2 += float32(0.0)
                            case "":
                                v26 := GenS3{F0: 42, F1: v5, F2: 0}
                                _ = v26
                            }
                            v4 = (^(v4))
                            fmt.Println(v3, v4, v4)
                        } else {
                            for fi27 := 0; fi27 < 7; fi27++ {
                                v1[""] = v5
                                _ = fi27
                            }
                            v28 := GenS2{F0: -671, F1: -1, F2: (42 == 255)}
                            v5 = ((v5 + "") + v5)
                            _ = v28
                        }
                    }
                    if (false || true) {
                        var v29 float32 = v2
                        _ = v29
                    } else {
                        var v30 int = -2147483648
                        _ = v30
                    }
                    _ = v19
                }
            }
        case 1:
            switch v5 {
            case "":
                var v31 float32 = ((float32(0.1) + v2) + float32(1.5))
                _ = v31
            case "a":
                fmt.Println(v4, v5)
            default:
                if (2147483647 > 0) {
                    if true {
                        fmt.Println(v2, v5, v5)
                        for fi32 := 0; fi32 < 6; fi32++ {
                            fmt.Println(v4, fi32, v2)
                            _ = fi32
                        }
                    }
                    var v33 uint8 = 42
                    v5 = ((v5 + v5) + v5)
                    _ = v33
                }
            }
        case 483:
            var v34 string = "foo bar"
            _ = v34
        default:
            if (!false) {
                v3 -= (^(v3))
                if len(v6) > 0 {
                    v6[0] = (v5 + v5)
                }
            } else {
                delete(v1, v5)
            }
        }
        var v35 float64 = 3.14159
        for fi36 := 0; fi36 < 2; fi36++ {
            v37 := GenS3{F0: -78, F1: ((v5 + v5) + ("hello" + (v5 + (v5 + v5)))), F2: 1}
            var v38 int = fi36
            _ = fi36
            _ = v37
            _ = v38
        }
        _ = v35
    } else {
        switch 42 {
        case 128:
            if (1 != 1) {
                var v39 float64 = 0.1
                var v40 int64 = (v3 >> ((18446744073709551615) & 63))
                _ = v39
                _ = v40
            }
        case 32767:
            fmt.Println(v2, v2, v5)
        default:
            var v41 uint32 = 2
            _ = v41
        }
        v3 = 0
        v4 += (v4 - 32767)
    }
    var v42 uint8 = 42
    if true {
        fmt.Println(v3, v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v7
}
