package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int32 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint; F1 uint32; F2 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string; F1 uint16; F2 uint32 }
type GenS6 struct { F0 uint; F1 uint16; F2 bool }

func main() {
    var v1 int = -1
    v2 := make(map[string]bool)
    var v3 string = "\n\t"
    v4 := make(map[string]bool)
    var v5 uint16 = 2
    fmt.Println(v5, v3, v5)
    if (false || (true && (-53 <= 42))) {
        fmt.Println(v1, v3, v5)
        v3 = ("\n\t" + v3)
    }
    _ = len(v2)
    v6 := GenS1{F0: 704, F1: v1, F2: 112}
    v7 := GenS2{F0: (-9223372036854775808 != 128), F1: v5}
    var v8 float64 = float64((float32(float64(1.5)) / (1e10 / 1.0)))
    fmt.Println(v1, v1, v5)
    if false {
        v9 := GenS3{F0: 1, F1: 2, F2: 1}
        var v10 uint32 = 0
        _ = v10
        _ = v9
    }
    var v11 int = (v1 - (v1 * -1))
    var v12 int32 = 828
    if false {
        switch (("foo bar" + v3) + v3) {
        case "a":
            switch v11 {
            case -1:
                for fi13 := 0; fi13 < 17; fi13++ {
                    for fi14 := 0; fi14 < 14; fi14++ {
                        var v15 uint = 4294967295
                        _ = fi14
                        _ = v15
                    }
                    var v16 int8 = -125
                    fmt.Println(v3, v8, v12)
                    _ = fi13
                    _ = v16
                }
            default:
                v17 := GenS4{F0: 0}
                _ = v17
            }
        case "hello":
            switch (v1 + (v11 & v1)) {
            case 2147483647:
                v5 = (v5 % ((65535) | 1))
            case 32767:
                switch -680 {
                case -9223372036854775808:
                    switch -754 {
                    case -1:
                        for fi18 := 0; fi18 < 20; fi18++ {
                            var v19 int64 = 0
                            for fi20 := 0; fi20 < 12; fi20++ {
                                if (0.1 < 3.14159) {
                                    var v21 float32 = 3.14159
                                    var v22 int8 = -97
                                    var v23 uint16 = 877
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                }
                                _ = fi20
                            }
                            for fi24 := 0; fi24 < 12; fi24++ {
                                v8 += 0.1
                                switch v11 {
                                case -9223372036854775808:
                                    fmt.Println(v8, v3, v11)
                                case 1:
                                    v19 += -1
                                }
                                var v25 int = 128
                                _ = fi24
                                _ = v25
                            }
                            _ = fi18
                            _ = v19
                        }
                    case -766:
                        var v26 uint16 = 376
                        _ = v26
                    }
                case 1:
                    switch 2147483647 {
                    case 32767:
                        var v27 uint64 = 727
                        _ = v27
                    case -9223372036854775808:
                        v28 := GenS1{F0: 42, F1: v1, F2: 0}
                        _ = v28
                    case 627:
                        var v29 uint64 = 558
                        _ = v29
                    }
                default:
                    v30 := GenS5{F0: "\n\t", F1: (v5 + 88), F2: 2147483647}
                    _ = v30
                }
            default:
                v31 := []int{}
                _ = v31
            }
        case "\n\t":
            var v32 uint32 = 387
            _ = v32
        }
        v12 = (^(v12))
        var v33 float32 = 1.5
        _ = v33
    }
    for fi34 := 0; fi34 < 2; fi34++ {
        var v35 uint64 = 564
        v8 += v8
        v36 := GenS6{F0: 2, F1: v5, F2: false}
        _ = fi34
        _ = v35
        _ = v36
    }
    var v37 string = v3
    for fi38 := 0; fi38 < 5; fi38++ {
        fmt.Println(v5)
        _ = len(v2)
        v2["hello"] = true
        _ = fi38
    }
    fmt.Println(v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
