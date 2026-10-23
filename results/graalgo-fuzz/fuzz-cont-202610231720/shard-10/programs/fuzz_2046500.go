package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint8; F1 bool }

func main() {
    v1 := make(map[string]string)
    var v2 int8 = 42
    v3 := make(map[int]uint32)
    var v4 float64 = 0.0
    v5 := []uint32{2147483647}
    var v6 int32 = -936
    switch ("hello" + "") {
    case "foo bar":
        if (float64(0.0) == 0.1) {
            switch "" {
            case "\n\t":
                switch 1 {
                case 0:
                    v2 = (^(v2))
                case 32768:
                    for fi7 := 0; fi7 < 6; fi7++ {
                        var v8 float64 = (v4 - 1.5)
                        _ = fi7
                        _ = v8
                    }
                }
            case "foo bar":
                v2 = (^(v2))
            case "a":
                v9 := GenS1{F0: 0, F1: 42, F2: 0}
                _ = v9
            }
        } else {
            var v10 uint32 = 4294967295
            _ = v10
        }
    case "a":
        _ = len(v5)
    case "\n\t":
        switch "" {
        case "a":
            for fi11 := 0; fi11 < 14; fi11++ {
                switch "hello" {
                case "hello":
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v4, v2, v6)
                    default:
                        for fi12 := 0; fi12 < 19; fi12++ {
                            v3[(-(fi12))] = 65535
                            _ = fi12
                        }
                    }
                case "\n\t":
                    v2 = 86
                default:
                    _ = len(v3)
                }
                switch (fi11 / (((^(fi11))) | 1)) {
                case -886:
                    var v13 uint64 = 42
                    _ = v13
                case -898:
                    v14 := GenS1{F0: 0, F1: 812, F2: 255}
                    _ = v14
                }
                var v15 int64 = int64(v2)
                _ = fi11
                _ = v15
            }
        }
    }
    var v16 int32 = v6
    for fi17 := 0; fi17 < 1; fi17++ {
        var v18 int64 = -619
        var v19 string = ("hello" + "hello")
        if false {
            switch fi17 {
            case -2147483648:
                switch "a" {
                case "a":
                    switch 9223372036854775807 {
                    case 32768:
                        v20 := GenS1{F0: 65535, F1: 194, F2: 255}
                        _ = v20
                    case 384:
                        v16 = 42
                    case 32767:
                        v21 := GenS2{F0: 128, F1: false}
                        _ = v21
                    }
                case "hello":
                    delete(v1, "a")
                default:
                    v22 := GenS1{F0: 112, F1: 231, F2: 455}
                    _ = v22
                }
            default:
                var v23 uint64 = 613
                _ = v23
            }
            fmt.Println(v2, v19, v4)
            var v24 uint32 = 949
            _ = v24
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    var v25 int32 = v16
    v26 := GenS1{F0: 4294967295, F1: 309, F2: 32767}
    var v27 uint32 = 0
    if false {
        var v28 uint8 = 0
        v29 := v3[32768]
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        var v31 int8 = v2
        var v32 float32 = 0.1
        _ = fi30
        _ = v31
        _ = v32
    }
    v33 := GenS1{F0: 65535, F1: 226, F2: 0}
    var v34 string = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
