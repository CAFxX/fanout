package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint64; F1 uint }

func main() {
    v1 := []string{"\n\t", "\n\t", "foo bar", ""}
    v2 := []string{"foo bar", ""}
    v3 := make(map[int]int)
    v4 := make(map[uint32]bool)
    v2 = append(v2, "hello")
    var v5 bool = (982 != 0)
    v3[-43] = 32768
    v6 := v4[658]
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            v9 := GenS1{F0: 2147483647, F1: 42}
            v6 = false
            var v10 int8 = 1
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := GenS1{F0: -2147483648, F1: 100}
        if v5 {
            if (!v6) {
                v5 = (42 <= 809)
            } else {
                _ = len(v4)
            }
            switch ("foo bar" + ("foo bar" + "foo bar")) {
            case "foo bar":
                switch (fi7 & (fi7 | 42)) {
                case 127:
                    if (0.1 != 3.14159) {
                        _ = len(v3)
                    }
                case 4294967295:
                    for fi12 := 0; fi12 < 5; fi12++ {
                        v6 = false
                        v5 = (v6 || v5)
                        _ = fi12
                    }
                case 505:
                    fmt.Println(fi7, v6, fi7)
                }
            case "":
                _ = len(v4)
            case "a":
                fmt.Println(v5, v6, v5)
            }
        }
        _ = fi7
        _ = v11
    }
    if (4294967295 > 18446744073709551615) {
        var v13 bool = (0 > 255)
        var v14 float32 = (1.5 / 1.5)
        for fi15 := 0; fi15 < 17; fi15++ {
            for fi16 := 0; fi16 < 14; fi16++ {
                var v17 uint8 = 255
                switch fi15 {
                case -1:
                    var v18 uint = 18446744073709551615
                    _ = v18
                default:
                    if true {
                        if (1000 >= -32768) {
                            if (128 > (fi15 - fi15)) {
                                v3[fi15] = fi16
                                var v19 int64 = -83
                                _ = v19
                            }
                            v5 = v13
                            if true {
                                v20 := GenS2{F0: -1, F1: 203}
                                v21 := v4[65535]
                                _ = v20
                                _ = v21
                            }
                        }
                    } else {
                        _ = len(v2)
                        delete(v4, 1)
                    }
                }
                if true {
                    if true {
                        v22 := GenS3{F0: -1, F1: 0}
                        _ = v22
                    } else {
                        var v23 int32 = -2147483648
                        var v24 int = (^(fi15))
                        _ = v23
                        _ = v24
                    }
                    var v25 uint = 187
                    _ = v25
                } else {
                    v26 := GenS4{F0: 659, F1: 255}
                    fmt.Println(v17, v17, v14)
                    v5 = false
                    _ = v26
                }
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
        _ = v13
        _ = v14
    } else {
        v5 = v5
    }
    var v27 uint16 = 42
    var v28 int8 = 127
    v3[0] = 32768
    v6 = (((!v5) || v6) && v6)
    v6 = (801 > 18446744073709551615)
    switch 288 {
    case 127:
        fmt.Println(v6)
    case 32767:
        v29 := GenS2{F0: -9223372036854775808, F1: 341}
        _ = v29
    case 42:
        v28 -= -1
    default:
        fmt.Println(v28)
    }
    v27 = (^(v27))
    var v30 int16 = 42
    _ = len(v4)
    fmt.Println(v27)
    switch ("hello" + "a") {
    case "a":
        var v31 int8 = v28
        _ = v31
    case "foo bar":
        v32 := GenS1{F0: int32(v28), F1: (^(v28))}
        _ = v32
    case "hello":
        v33 := GenS4{F0: 42, F1: 18446744073709551615}
        _ = v33
    }
    switch "foo bar" {
    case "\n\t":
        v34 := GenS2{F0: 141, F1: 398}
        _ = v34
    case "a":
        for fi35 := 0; fi35 < 5; fi35++ {
            var v36 float64 = 0.0
            v37 := GenS2{F0: 1, F1: 2147483647}
            v5 = false
            _ = fi35
            _ = v36
            _ = v37
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
