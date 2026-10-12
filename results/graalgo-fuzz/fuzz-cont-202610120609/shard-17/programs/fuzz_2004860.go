package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int8; F2 int64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]uint32)
    v3 := []string{"a", ((("hello" + ("" + "")) + "foo bar") + ("a" + ((("" + "hello") + (("" + "foo bar") + "a")) + ("foo bar" + "a"))))}
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]uint32)
    var v6 float32 = 1e10
    v6 = 0.0
    v7 := GenS1{F0: 922}
    v6 -= v6
    v6 += ((-2.25 / 1.0) * v6)
    switch "\n\t" {
    case "hello":
        var v8 uint8 = 127
        _ = v8
    case "\n\t":
        for fi9 := 0; fi9 < 1; fi9++ {
            switch fi9 {
            case 771:
                var v10 int8 = 100
                _ = v10
            case 133:
                switch "hello" {
                case "foo bar":
                    if true {
                        if (4294967295 < 2) {
                            var v11 int8 = -1
                            _ = v11
                        }
                    }
                case "hello":
                    if false {
                        var v12 string = "hello"
                        _ = v12
                    } else {
                        if (1 < -1) {
                            var v13 string = "hello"
                            switch 9223372036854775807 {
                            case 9223372036854775807:
                                if len(v3) > 0 {
                                    v3[0] = v13
                                }
                            }
                            var v14 float32 = 0.1
                            _ = v13
                            _ = v14
                        } else {
                            var v15 uint = 131
                            _ = v15
                        }
                        v16 := GenS2{F0: 255, F1: -1, F2: 9223372036854775807}
                        if true {
                            v6 += v6
                        }
                        _ = v16
                    }
                }
            default:
                v2[2147483647] = 816
            }
            var v17 float32 = float32(0.0)
            _ = fi9
            _ = v17
        }
    }
    var v18 float64 = float64(v6)
    var v19 int = 2147483647
    var v20 uint8 = 255
    v21 := GenS1{F0: 9223372036854775807}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
