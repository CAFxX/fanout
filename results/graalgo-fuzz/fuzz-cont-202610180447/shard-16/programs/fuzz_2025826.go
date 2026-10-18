package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint64 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 bool; F1 uint16 }

func main() {
    var v1 uint16 = 2
    var v2 string = "a"
    v3 := make(map[string]uint32)
    var v4 float64 = (1.5 / (0.1 - 1e10))
    var v5 uint = 18446744073709551615
    var v6 uint32 = 42
    var v7 uint16 = (v1 << ((329) & 63))
    v6 += v6
    fmt.Println(v2)
    v4 = (v4 - float64(0.0))
    v8 := GenS1{F0: -961, F1: uint(v7), F2: 1}
    if (!(!true)) {
        if (2 != v5) {
            for fi9 := 0; fi9 < 20; fi9++ {
                fmt.Println(v5, v5, v1)
                _ = fi9
            }
            v10 := []int{168, -2147483648, 1, 2147483647}
            switch "hello" {
            case "hello":
                v6 = (v6 / ((v6) | 1))
            case "":
                switch ("hello" + "a") {
                case "hello":
                    var v11 uint16 = 42
                    _ = v11
                case "foo bar":
                    for fi12 := 0; fi12 < 12; fi12++ {
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        var v13 int16 = 0
                        var v14 uint32 = 42
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                case "\n\t":
                    for fi15 := 0; fi15 < 13; fi15++ {
                        var v16 int16 = -1
                        if ((9223372036854775807 > 9223372036854775807) || true) {
                            fmt.Println(v16, v1)
                            switch (fi15 >> ((v5) & 63)) {
                            case -9223372036854775808:
                                _ = len(v3)
                            case -1:
                                if true {
                                    v16 -= 1
                                    switch (((v2 + "a") + v2) + v2) {
                                    case "foo bar":
                                        v2 = v2
                                    case "":
                                        var v17 int = fi15
                                        _ = v17
                                    default:
                                        fmt.Println(v2)
                                    }
                                    if len(v10) > 0 {
                                        v10[0] = (fi15 % (((fi15 * (fi15 / ((32768) | 1)))) | 1))
                                    }
                                } else {
                                    _ = len(v10)
                                    v18 := GenS1{F0: 9223372036854775807, F1: (v5 << ((v5) & 63)), F2: 633}
                                    _ = v18
                                }
                            }
                            var v19 uint = 4294967295
                            _ = v19
                        }
                        fmt.Println(v2)
                        _ = fi15
                        _ = v16
                    }
                }
            case "foo bar":
                if len(v10) > 0 {
                    v10[0] = -1
                }
            }
            _ = v10
        } else {
            var v20 int = 9223372036854775807
            switch v2 {
            case "":
                v21 := GenS1{F0: -628, F1: v5, F2: 2}
                _ = v21
            case "a":
                var v22 uint64 = 42
                _ = v22
            case "foo bar":
                var v23 float64 = (v4 - v4)
                _ = v23
            default:
                for fi24 := 0; fi24 < 1; fi24++ {
                    for fi25 := 0; fi25 < 10; fi25++ {
                        v26 := GenS1{F0: -800, F1: (v5 << ((v5) & 63)), F2: 42}
                        v3["a"] = (^(v6))
                        v6 -= (v6 - 0)
                        _ = fi25
                        _ = v26
                    }
                    _ = fi24
                }
            }
            _ = v20
        }
        if true {
            var v27 int16 = int16(v1)
            v27 = v27
            var v28 int16 = v27
            _ = v27
            _ = v28
        }
        v7 = (v1 ^ 32767)
    } else {
        var v29 uint = (v5 >> ((947) & 63))
        _ = v29
    }
    switch v2 {
    case "hello":
        switch "a" {
        case "a":
            v1 = 2
        case "":
            v7 = (v1 | uint16(v5))
        default:
            v30 := GenS2{F0: v2, F1: 733}
            _ = v30
        }
    }
    var v31 bool = ((v2 >= ((v2 + v2) + ("\n\t" + (v2 + v2)))) || true)
    switch -1 {
    case 32768:
        v31 = ((v4 - ((float64(0.0) / float64(1e10)) * -2.25)) >= -2.25)
    }
    v6 += (v6 >> (((^(v5))) & 63))
    var v32 uint = (v5 >> (((v5 % (((v5 % ((65535) | 1))) | 1))) & 63))
    for fi33 := 0; fi33 < 2; fi33++ {
        v6 = (v6 - (v6 / ((v6) | 1)))
        fmt.Println(v6, v7)
        _ = fi33
    }
    v34 := GenS3{F0: ((!true) && (v31 || (v4 < (float64(1e10) + v4)))), F1: 1}
    var v35 int8 = -128
    v2 = "a"
    for fi36 := 0; fi36 < 8; fi36++ {
        var v37 int8 = v35
        v38 := GenS2{F0: "hello", F1: -1}
        _ = fi36
        _ = v37
        _ = v38
    }
    v39 := []int{32767}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
