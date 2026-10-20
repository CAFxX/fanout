package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 string; F1 uint32 }

func main() {
    var v1 int16 = 1000
    var v2 float64 = (1.5 + (-2.25 / float64(1.5)))
    var v3 bool = false
    var v4 float32 = 0.1
    v5 := []uint32{2147483647}
    if (!v3) {
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 uint32 = 4294967295
            var v8 string = ("\n\t" + ("foo bar" + "a"))
            _ = fi6
            _ = v7
            _ = v8
        }
    } else {
        v1 += -531
        fmt.Println(v1, v2, v3)
        v9 := GenS1{F0: 65535, F1: 0}
        _ = v9
    }
    v10 := GenS1{F0: 2, F1: 42}
    v11 := GenS1{F0: 2, F1: 0}
    if v3 {
        var v12 int = -9223372036854775808
        v13 := GenS1{F0: 0, F1: 1}
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 int8 = 100
            if (0 < 2) {
                v16 := GenS1{F0: 65535, F1: v15}
                if (!v3) {
                    for fi17 := 0; fi17 < 14; fi17++ {
                        _ = len(v5)
                        _ = fi17
                    }
                    switch (fi14 + fi14) {
                    case 921:
                        switch "foo bar" {
                        case "foo bar":
                            v4 = float32(v2)
                        case "hello":
                            v5 = append(v5, 4294967295)
                        case "a":
                            var v18 int8 = v15
                            _ = v18
                        }
                    default:
                        v19 := GenS2{F0: (v15 >> ((18446744073709551615) & 63)), F1: 65535}
                        _ = v19
                    }
                    v5 = append(v5, 42)
                }
                var v20 bool = v3
                _ = v16
                _ = v20
            }
            _ = fi14
            _ = v15
        }
        var v21 string = "foo bar"
        var v22 int16 = v1
        _ = v21
        _ = v22
    }
    if v3 {
        if v3 {
            if v3 {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v23 := GenS1{F0: 743, F1: -48}
                v24 := GenS3{F0: "\n\t", F1: 1}
                _ = v23
                _ = v24
            }
            fmt.Println(v1)
            fmt.Println(v4, v1)
        }
        var v25 int64 = 1
        _ = v25
    }
    var v26 uint = 691
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, 2147483647)
    v5 = append(v5, 785)
    var v27 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
