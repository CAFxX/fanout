package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint }
func genFn0(p0 string, p1 int, p2 string) uint {
    v6 := GenS1{F0: p1, F1: -32768, F2: 654}
    var v7 int16 = 32767
    var v8 uint16 = 255
    var v9 int16 = (v7 + -32768)
    v10 := []int{714, p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[int]uint32)
    var v3 int16 = -296
    v4 := make(map[int]bool)
    v5 := make(map[int]int)
    v5[-927] = 32767
    v1 = v1
    var v11 uint8 = 255
    if ((1.5 - 0.0) == (-2.25 / (float64(-2.25) - 0.1))) {
        for fi12 := 0; fi12 < 20; fi12++ {
            if false {
                fmt.Println(v3, v1, v1)
                var v13 uint = 245
                v11 -= v11
                _ = v13
            } else {
                if true {
                    _ = len(v4)
                    var v14 int = fi12
                    v1 = 2
                    _ = v14
                }
                fmt.Println(v1, v3, v11)
            }
            _ = fi12
        }
    }
    var v15 uint = 2
    if (true && (42 > 127)) {
        for fi16 := 0; fi16 < 15; fi16++ {
            for fi17 := 0; fi17 < 17; fi17++ {
                fmt.Println(v1)
                v18 := []int{fi16, (fi16 - -440)}
                v11 = (v11 << ((v15) & 63))
                switch "a" {
                case "a":
                    var v19 int16 = (^(v3))
                    _ = v19
                default:
                    var v20 int = 639
                    switch "" {
                    case "hello":
                        var v21 bool = (-2.25 < 0.1)
                        _ = v21
                    case "":
                        v11 = v11
                        var v22 uint8 = 128
                        _ = v22
                    }
                    _ = v20
                }
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
    } else {
        fmt.Println(v1, v11)
        var v23 int16 = v3
        _ = v23
    }
    _ = len(v2)
    v24 := GenS1{F0: 4294967295, F1: v3, F2: 18446744073709551615}
    v25 := []int{1, 9223372036854775807, 32768}
    var v26 uint8 = (v11 | (v11 * v11))
    for fi27 := 0; fi27 < 20; fi27++ {
        var v28 float32 = 0.1
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
