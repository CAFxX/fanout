package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS3 struct { F0 int32; F1 int32; F2 uint64 }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8; F1 uint64; F2 uint8 }

func main() {
    var v1 int32 = -1
    var v2 float64 = 1.5
    var v3 int8 = -7
    var v4 int32 = (v1 ^ int32(v3))
    var v5 uint8 = 255
    v6 := make(map[string]uint32)
    v2 -= (float64(1e10) - v2)
    if (813 == 255) {
        if ("\n\t" <= ("a" + "hello")) {
            switch -9223372036854775808 {
            case 762:
                var v7 uint16 = 566
                _ = v7
            case 0:
                for fi8 := 0; fi8 < 10; fi8++ {
                    if (2 > 2) {
                        switch ((("a" + ("\n\t" + "a")) + "hello") + "a") {
                        case "foo bar":
                            v3 -= 62
                        case "\n\t":
                            var v9 int = fi8
                            _ = v9
                        case "":
                            for fi10 := 0; fi10 < 16; fi10++ {
                                var v11 int = (fi10 << ((1) & 63))
                                _ = fi10
                                _ = v11
                            }
                        default:
                            if (!false) {
                                v6["foo bar"] = 942
                                v6["foo bar"] = 2147483647
                                v12 := GenS1{F0: 226, F1: 654}
                                _ = v12
                            }
                        }
                    }
                    _ = fi8
                }
            case -127:
                v2 -= v2
            }
            var v13 float32 = ((1e10 * (float32((1e10 - v2)) * 0.0)) - float32(0.1))
            v5 = (v5 - v5)
            _ = v13
        } else {
            v5 -= v5
            v6[""] = 1
        }
        if (!true) {
            delete(v6, ((("foo bar" + "a") + "\n\t") + "hello"))
            var v14 int = 9223372036854775807
            var v15 uint64 = 0
            _ = v14
            _ = v15
        } else {
            v16 := GenS1{F0: uint32(v5), F1: 4294967295}
            if false {
                v17 := GenS1{F0: 831, F1: 18446744073709551615}
                _ = v17
            }
            _ = v16
        }
        v18 := GenS1{F0: 120, F1: 787}
        _ = v18
    } else {
        switch 0 {
        case -721:
            v19 := GenS1{F0: 4294967295, F1: 255}
            _ = v19
        }
        for fi20 := 0; fi20 < 14; fi20++ {
            var v21 bool = (-1 <= 1)
            v21 = false
            _ = fi20
            _ = v21
        }
        if (1 <= 42) {
            switch ((("" + "\n\t") + ("" + "foo bar")) + "hello") {
            case "\n\t":
                if false {
                    v2 = v2
                    v4 = v4
                } else {
                    fmt.Println(v1)
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(v2, v1)
                    case "foo bar":
                        var v22 float64 = v2
                        _ = v22
                    case "\n\t":
                        switch ("hello" + (("a" + ("a" + "hello")) + "hello")) {
                        case "a":
                            var v23 int16 = 1000
                            _ = v23
                        }
                    default:
                        fmt.Println(v4, v5)
                    }
                    v24 := []int{944, 0}
                    _ = v24
                }
            case "a":
                var v25 uint32 = 42
                _ = v25
            case "hello":
                var v26 float64 = float64(0.1)
                _ = v26
            }
            v27 := GenS2{F0: 994, F1: 127, F2: (v3 << ((132) & 63))}
            v6[((("hello" + "a") + ("" + "\n\t")) + "hello")] = 4294967295
            _ = v27
        }
    }
    if false {
        v28 := GenS3{F0: (v1 ^ 304), F1: (v1 - 0), F2: 787}
        for fi29 := 0; fi29 < 7; fi29++ {
            if false {
                v3 = (v3 - (v3 + (^(v3))))
                fmt.Println(v3)
            }
            v30 := GenS4{F0: (fi29 ^ (fi29 % (((fi29 ^ (fi29 / ((fi29) | 1)))) | 1))), F1: -9223372036854775808}
            _ = fi29
            _ = v30
        }
        v5 = 204
        _ = v28
    }
    v31 := GenS5{F0: 2147483647}
    if (185 != -514) {
        v32 := GenS6{F0: (^(v3)), F1: 777, F2: (^(v5))}
        for fi33 := 0; fi33 < 11; fi33++ {
            delete(v6, "foo bar")
            _ = fi33
        }
        _ = v32
    } else {
        if true {
            if false {
                v4 = (v4 + 42)
                fmt.Println(v5, v2, v5)
            }
            v5 = (^(v5))
        }
        if ((("foo bar" > "\n\t") || false) || (v1 <= 0)) {
            var v34 string = "\n\t"
            _ = v34
        }
    }
    v35 := GenS5{F0: 2}
    var v36 string = "foo bar"
    fmt.Println(v1, v2)
    v1 = v4
    var v37 uint8 = v5
    v5 = (v37 << ((18446744073709551615) & 63))
    v6[v36] = 42
    var v38 int32 = v4
    if true {
        v2 = v2
    } else {
        v4 += (v4 * (v38 * v1))
    }
    var v39 float64 = ((v2 / (v2 / (float64(float32((float64(-2.25) + v2))) - (v2 / float64((((0.1 - 3.14159) / 1.0) * 3.14159)))))) - v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
