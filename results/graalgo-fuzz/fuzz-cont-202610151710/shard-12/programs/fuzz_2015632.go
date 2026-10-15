package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 int16; F1 int16; F2 int16 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 uint8; F1 uint16; F2 int }
type GenS6 struct { F0 uint8 }

func main() {
    v1 := []string{(("foo bar" + "hello") + ""), "\n\t", "a", "\n\t"}
    v2 := make(map[int]string)
    var v3 uint = 1
    v3 = 4294967295
    for fi4 := 0; fi4 < 6; fi4++ {
        for fi5 := 0; fi5 < 19; fi5++ {
            v6 := GenS1{F0: 42}
            fmt.Println(fi4, fi4, v3)
            v3 += v3
            _ = fi5
            _ = v6
        }
        var v7 float64 = 0.0
        _ = fi4
        _ = v7
    }
    v3 = 1
    v3 = v3
    v8 := GenS2{F0: false, F1: (v3 * 2)}
    if true {
        fmt.Println(v3)
        v9 := GenS2{F0: (1 != 2147483647), F1: (v3 ^ v3)}
        switch "\n\t" {
        case "":
            for fi10 := 0; fi10 < 14; fi10++ {
                var v11 int8 = 100
                _ = fi10
                _ = v11
            }
        }
        _ = v9
    } else {
        for fi12 := 0; fi12 < 5; fi12++ {
            for fi13 := 0; fi13 < 1; fi13++ {
                _ = len(v2)
                switch "a" {
                case "hello":
                    v1 = append(v1, "hello")
                case "\n\t":
                    delete(v2, -158)
                case "foo bar":
                    if true {
                        _ = len(v2)
                        v14 := GenS1{F0: 1}
                        delete(v2, 127)
                        _ = v14
                    } else {
                        switch (-(fi13)) {
                        case 0:
                            switch (fi12 << ((v3) & 63)) {
                            case 128:
                                _ = len(v1)
                            default:
                                v15 := GenS1{F0: 42}
                                _ = v15
                            }
                            v16 := GenS3{F0: -553, F1: -1, F2: 1000}
                            _ = v16
                        case -951:
                            v3 -= 567
                        }
                    }
                default:
                    if (true || (!false)) {
                        v17 := GenS4{F0: "", F1: 32767}
                        v3 += (^(v3))
                        _ = v17
                    } else {
                        v18 := GenS5{F0: 42, F1: 1, F2: 898}
                        v19 := GenS5{F0: uint8(v3), F1: 120, F2: (^(fi13))}
                        _ = v18
                        _ = v19
                    }
                }
                _ = fi13
            }
            _ = fi12
        }
    }
    v20 := GenS6{F0: 2}
    if true {
        switch 32768 {
        case 4294967295:
            delete(v2, -776)
        case 1:
            switch "foo bar" {
            case "hello":
                v3 = v3
            case "":
                v3 = (v3 + 428)
            }
        case 32767:
            v21 := GenS3{F0: 676, F1: -32768, F2: 1000}
            _ = v21
        default:
            if (-1 != -802) {
                fmt.Println(v3)
                for fi22 := 0; fi22 < 17; fi22++ {
                    v3 -= 255
                    var v23 float64 = ((0.1 / 1.0) + 1e10)
                    for fi24 := 0; fi24 < 2; fi24++ {
                        v25 := GenS6{F0: 128}
                        v23 += v23
                        _ = fi24
                        _ = v25
                    }
                    _ = fi22
                    _ = v23
                }
                v3 += v3
            } else {
                v3 -= 88
            }
        }
    } else {
        v26 := GenS6{F0: 42}
        _ = v26
    }
    _ = len(v2)
    var v27 int16 = -606
    v3 = (v3 | 255)
    v3 = v3
    if false {
        v27 += 42
        var v28 string = "hello"
        _ = v28
    }
    if (!true) {
        fmt.Println(v27, v27)
        v29 := GenS1{F0: -9223372036854775808}
        _ = v29
    }
    v27 -= v27
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v8
}
