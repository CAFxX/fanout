package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint8 }

func main() {
    v1 := make(map[int]string)
    v2 := []int{-452, -9223372036854775808, 32767}
    var v3 string = ("hello" + ("\n\t" + "\n\t"))
    var v4 bool = (1000 >= 642)
    var v5 int64 = -49
    var v6 uint8 = 35
    if (true || true) {
        for fi7 := 0; fi7 < 19; fi7++ {
            v3 = v3
            var v8 uint16 = 2
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: 42}
        v3 = "foo bar"
        _ = v9
    }
    if len(v2) > 0 {
        v2[0] = 1
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        switch (-(fi10)) {
        case 32768:
            if len(v2) > 0 {
                v2[0] = (fi10 - fi10)
            }
        }
        if v4 {
            v11 := GenS1{F0: 128}
            if v4 {
                fmt.Println(v5)
                switch (v3 + "hello") {
                case "foo bar":
                    v3 = v3
                case "hello":
                    var v12 uint = 367
                    _ = v12
                case "a":
                    v6 = (v6 << ((244) & 63))
                default:
                    v13 := GenS1{F0: 1}
                    _ = v13
                }
            }
            _ = v11
        } else {
            if len(v2) > 0 {
                v2[0] = 0
            }
            for fi14 := 0; fi14 < 15; fi14++ {
                if (0 < 145) {
                    if v4 {
                        if (-1 != 2147483647) {
                            if (42 == -32768) {
                                if len(v2) > 0 {
                                    v2[0] = (fi14 * (^(fi14)))
                                }
                            }
                            var v15 uint16 = 916
                            var v16 int8 = -128
                            _ = v15
                            _ = v16
                        }
                        v17 := GenS1{F0: fi10}
                        v5 = (v5 / (((v5 ^ (v5 / ((v5) | 1)))) | 1))
                        _ = v17
                    }
                    var v18 int = (fi10 >> ((255) & 63))
                    if false {
                        switch "" {
                        case "hello":
                            switch (fi10 - fi10) {
                            case 92:
                                switch fi14 {
                                case 9223372036854775807:
                                    v18 = -956
                                case 32767:
                                    if ((true && (!(-2.25 < (0.1 + (float64(1e10) / (0.1 / 1.5)))))) || (!v4)) {
                                        var v19 uint8 = 2
                                        _ = v19
                                    } else {
                                        v18 = -668
                                        fmt.Println(fi14, v18, v18)
                                    }
                                }
                            }
                        case "\n\t":
                            v20 := GenS2{F0: 324, F1: 988, F2: 2}
                            v5 = (v5 * (-(v5)))
                            var v21 int = (fi14 - 870)
                            _ = v20
                            _ = v21
                        case "a":
                            var v22 string = ""
                            _ = v22
                        default:
                            if (v6 > v6) {
                                var v23 string = v3
                                _ = v23
                            }
                        }
                    }
                    _ = v18
                } else {
                    var v24 uint64 = 0
                    if len(v2) > 0 {
                        v2[0] = -2147483648
                    }
                    _ = v24
                }
                _ = len(v2)
                var v25 uint64 = 1
                _ = fi14
                _ = v25
            }
        }
        switch (fi10 ^ fi10) {
        case 741:
            var v26 uint16 = 2
            _ = v26
        case 32768:
            if true {
                v6 -= (^(v6))
            } else {
                var v27 uint32 = 104
                v28 := v1[504]
                fmt.Println(v27)
                _ = v27
                _ = v28
            }
        }
        _ = fi10
    }
    v4 = true
    var v29 uint8 = (v6 * v6)
    v3 = v3
    v30 := GenS3{F0: 52, F1: -1, F2: (^(v29))}
    v29 = v6
    var v31 bool = (v4 && (!(42 <= 1)))
    v29 -= (v29 ^ v29)
    v6 = 255
    for fi32 := 0; fi32 < 8; fi32++ {
        v29 += (v29 >> ((255) & 63))
        var v33 float64 = (1e10 + 0.1)
        v3 = "hello"
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 7; fi34++ {
        _ = len(v1)
        var v35 uint64 = 42
        v36 := GenS1{F0: (-(fi34))}
        _ = fi34
        _ = v35
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
