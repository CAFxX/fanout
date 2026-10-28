package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint; F2 int16 }
type GenS3 struct { F0 uint16; F1 int8; F2 int64 }
type GenS4 struct { F0 string; F1 uint8; F2 int64 }

func main() {
    var v1 int = -1
    v2 := make(map[int]string)
    var v3 string = "a"
    var v4 int16 = 1
    v5 := make(map[string]string)
    if (42 == 42) {
        switch (v3 + "hello") {
        case "foo bar":
            v1 -= (-(v1))
        case "hello":
            switch v3 {
            case "foo bar":
                v6 := GenS1{F0: 42, F1: 2}
                _ = v6
            }
        default:
            fmt.Println(v4, v3)
        }
        fmt.Println(v3, v1)
    } else {
        var v7 float32 = (0.0 / ((0.1 - -2.25) - 0.0))
        _ = v7
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        switch fi8 {
        case 32768:
            v9 := v5[((v3 + "hello") + v3)]
            _ = v9
        default:
            v10 := v5["a"]
            _ = v10
        }
        v11 := GenS1{F0: 65535, F1: 2147483648}
        _ = fi8
        _ = v11
    }
    delete(v2, (v1 * v1))
    if (v4 <= (v4 / ((v4) | 1))) {
        v12 := GenS2{F0: 304, F1: 81, F2: v4}
        _ = v12
    } else {
        for fi13 := 0; fi13 < 14; fi13++ {
            v2[1] = v3
            v1 = (^(v1))
            v14 := GenS1{F0: 517, F1: 65535}
            _ = fi13
            _ = v14
        }
    }
    var v15 float32 = float32(((((1.5 / (0.0 - 1.5)) / (0.0 + 1.5)) - 1.5) + 0.0))
    v2[(v1 | (v1 + 42))] = "\n\t"
    if false {
        v4 = 42
    }
    var v16 uint32 = 2
    switch v1 {
    case 812:
        for fi17 := 0; fi17 < 12; fi17++ {
            fmt.Println(v1, v1)
            switch (v3 + "hello") {
            case "\n\t":
                v18 := GenS3{F0: 42, F1: 100, F2: 42}
                _ = v18
            case "foo bar":
                var v19 int32 = 1
                _ = v19
            }
            for fi20 := 0; fi20 < 20; fi20++ {
                v21 := []int{9223372036854775807}
                _ = len(v5)
                switch (v3 + v3) {
                case "":
                    v16 = 0
                case "a":
                    if false {
                        v22 := GenS1{F0: (v16 >> ((4294967295) & 63)), F1: v16}
                        v3 = (("foo bar" + "foo bar") + v3)
                        if (2147483648 != v16) {
                            v23 := GenS2{F0: -458, F1: 4294967295, F2: 498}
                            _ = v23
                        } else {
                            v24 := GenS2{F0: -115, F1: 0, F2: int16(fi20)}
                            _ = v24
                        }
                        _ = v22
                    }
                default:
                    var v25 uint = 42
                    _ = v25
                }
                _ = len(v21)
                v26 := GenS4{F0: (v3 + (v3 + v3)), F1: 2, F2: -9223372036854775808}
                _ = fi20
                _ = v21
                _ = v26
            }
            _ = fi17
        }
    case 32767:
        fmt.Println(v3)
    default:
        var v27 int8 = 100
        _ = v27
    }
    v28 := []int{2147483647}
    fmt.Println(v4, v15)
    switch (v3 + "a") {
    case "a":
        fmt.Println(v16, v16, v1)
    case "\n\t":
        fmt.Println(v3)
    case "":
        for fi29 := 0; fi29 < 5; fi29++ {
            for fi30 := 0; fi30 < 7; fi30++ {
                switch "a" {
                case "":
                    v3 = "foo bar"
                case "foo bar":
                    v31 := GenS1{F0: v16, F1: 381}
                    _ = v31
                case "\n\t":
                    v16 = (v16 - 2147483648)
                }
                v32 := GenS2{F0: 830, F1: 74, F2: (v4 >> ((4294967295) & 63))}
                _ = fi30
                _ = v32
            }
            _ = fi29
        }
    default:
        var v33 int32 = 841
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
