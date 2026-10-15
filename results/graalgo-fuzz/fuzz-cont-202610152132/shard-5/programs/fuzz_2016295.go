package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int8 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]bool)
    v3 := make(map[string]uint32)
    v4 := make(map[int]bool)
    switch 9223372036854775807 {
    case 0:
        var v5 uint = 2
        _ = v5
    default:
        switch "" {
        case "hello":
            for fi6 := 0; fi6 < 9; fi6++ {
                fmt.Println(fi6)
                for fi7 := 0; fi7 < 5; fi7++ {
                    v8 := GenS1{F0: 1}
                    if false {
                        var v9 string = "a"
                        var v10 uint = 441
                        _ = v10
                        _ = v9
                    }
                    var v11 int32 = 0
                    _ = fi7
                    _ = v11
                    _ = v8
                }
                _ = fi6
            }
        case "foo bar":
            if (1 < 123) {
                v12 := GenS1{F0: 214}
                var v13 int32 = 1
                var v14 int32 = v13
                _ = v12
                _ = v13
                _ = v14
            }
        }
    }
    if false {
        v15 := GenS1{F0: -451}
        var v16 bool = false
        _ = v15
        _ = v16
    } else {
        var v17 float64 = 1.5
        v17 = -2.25
        var v18 uint64 = 42
        _ = v17
        _ = v18
    }
    var v19 int = 2147483647
    v19 += v19
    if (0 >= 2147483648) {
        var v20 uint32 = 42
        _ = v20
    } else {
        if (1e10 >= (1e10 * float32(((0.0 * -2.25) * float64(1.5))))) {
            v21 := GenS2{F0: 65535, F1: 32767}
            _ = v21
        } else {
            v19 = (v19 / ((v19) | 1))
            v22 := GenS3{F0: -33}
            v19 = (v19 * (v19 * v19))
            _ = v22
        }
        var v23 float64 = 3.14159
        for fi24 := 0; fi24 < 11; fi24++ {
            for fi25 := 0; fi25 < 11; fi25++ {
                fmt.Println(fi25, fi25, fi24)
                v23 = v23
                var v26 uint64 = 40
                _ = fi25
                _ = v26
            }
            var v27 int8 = 100
            _ = fi24
            _ = v27
        }
        _ = v23
    }
    var v28 float32 = (((1e10 / (1.5 + -2.25)) * 0.0) / 0.1)
    v29 := GenS3{F0: -43}
    fmt.Println(v28, v28)
    var v30 uint16 = 32767
    var v31 int = 32767
    var v32 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
