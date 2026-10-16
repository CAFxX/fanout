package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := make(map[string]string)
    var v2 uint = 0
    v3 := make(map[string]uint32)
    var v4 int16 = 1000
    var v5 uint = (v2 | 0)
    v1[(("hello" + "") + "hello")] = (("\n\t" + (("\n\t" + "\n\t") + "a")) + "hello")
    fmt.Println(v4)
    for fi6 := 0; fi6 < 14; fi6++ {
        fmt.Println(v5, v4, fi6)
        var v7 uint16 = uint16(v2)
        v2 += 102
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: -848}
    if (true || (4294967295 <= 0)) {
        switch "hello" {
        case "":
            if (!false) {
                var v9 int64 = 104
                v9 = (v9 >> ((v2) & 63))
                fmt.Println(v9, v5)
                _ = v9
            }
        case "hello":
            for fi10 := 0; fi10 < 3; fi10++ {
                v5 -= 0
                v11 := GenS2{F0: "", F1: 660}
                _ = fi10
                _ = v11
            }
        case "a":
            for fi12 := 0; fi12 < 1; fi12++ {
                switch "" {
                case "\n\t":
                    for fi13 := 0; fi13 < 12; fi13++ {
                        var v14 uint8 = 255
                        v15 := GenS1{F0: 9223372036854775807}
                        if ((^(v14)) < v14) {
                            v16 := GenS3{F0: 32767}
                            delete(v3, ("a" + "a"))
                            _ = v16
                        } else {
                            var v17 float64 = (((0.0 / (3.14159 * (float64(float32(1.5)) / 1e10))) - 3.14159) / 0.1)
                            _ = v17
                        }
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                default:
                    var v18 float64 = (0.0 * -2.25)
                    _ = v18
                }
                _ = fi12
            }
        }
        fmt.Println(v5)
        v2 = 4294967295
    }
    v4 = v4
    v2 -= v5
    var v19 uint32 = 2147483648
    v2 = v2
    v4 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
