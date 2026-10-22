package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int16; F1 int16; F2 bool }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[string]string)
    v3 := make(map[int]int)
    if false {
        if false {
            v1 -= 2
        } else {
            for fi4 := 0; fi4 < 6; fi4++ {
                for fi5 := 0; fi5 < 10; fi5++ {
                    if (!true) {
                        switch "\n\t" {
                        case "foo bar":
                            v6 := GenS1{F0: 791, F1: 2147483647}
                            _ = v6
                        case "a":
                            v7 := GenS2{F0: 32767, F1: 1000, F2: (961 == 312)}
                            v1 = 18446744073709551615
                            _ = v7
                        case "\n\t":
                            delete(v2, "\n\t")
                        default:
                            var v8 bool = true
                            _ = v8
                        }
                    } else {
                        v9 := GenS2{F0: -32768, F1: 42, F2: false}
                        _ = v9
                    }
                    v1 = 255
                    var v10 uint64 = 923
                    _ = fi5
                    _ = v10
                }
                _ = fi4
            }
            v1 = v1
        }
        if (0 == 42) {
            v11 := GenS2{F0: 816, F1: 527, F2: (255 < 116)}
            var v12 int32 = 0
            _ = v11
            _ = v12
        } else {
            var v13 string = ("\n\t" + ("\n\t" + ("\n\t" + "a")))
            var v14 uint16 = 255
            switch 9223372036854775807 {
            case -654:
                fmt.Println(v14, v14)
            default:
                v15 := []int{312, 32767, 9223372036854775807, -163}
                _ = v15
            }
            _ = v13
            _ = v14
        }
    }
    if true {
        v1 += (^(v1))
    } else {
        var v16 float32 = float32(1.5)
        _ = v16
    }
    fmt.Println(v1)
    v1 -= (v1 - v1)
    var v17 float32 = 3.14159
    var v18 int64 = 42
    var v19 uint64 = 42
    if (((v17 + 0.0) + (float32(float64(v17)) + float32(0.0))) == ((v17 / (((float32(3.14159) + (v17 + 0.0)) + float32(0.1)) - 3.14159)) - 1.5)) {
        var v20 float32 = ((-2.25 / float32(0.1)) - v17)
        fmt.Println(v19, v17)
        v17 -= (1e10 + v17)
        _ = v20
    }
    var v21 int64 = (v18 >> ((42) & 63))
    fmt.Println(v1)
    v19 += 534
    switch "\n\t" {
    case "\n\t":
        v22 := []int{-422, 765, 127}
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
}
