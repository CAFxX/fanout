package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int64 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 string; F1 int64; F2 uint8 }
func genFn0(p0 string, p1 uint16, p2 uint64) int64 {
    var v4 int16 = 0
    v5 := GenS1{F0: 255, F1: v4, F2: -1}
    var v6 int64 = -432
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return (^(v6))
}

func main() {
    var v1 uint32 = 527
    v2 := []uint32{v1, (v1 * v1), (^(v1)), v1}
    v3 := make(map[uint32]uint32)
    _ = len(v2)
    for fi7 := 0; fi7 < 17; fi7++ {
        v1 -= (v1 - (v1 | 2147483648))
        switch ("hello" + "a") {
        case "":
            var v8 uint64 = 18446744073709551615
            _ = v8
        default:
            v1 = 2
        }
        _ = fi7
    }
    var v9 uint = 343
    v9 = v9
    v1 = 0
    if (true && false) {
        v1 = 42
        for fi10 := 0; fi10 < 2; fi10++ {
            v2 = append(v2, v1)
            for fi11 := 0; fi11 < 2; fi11++ {
                var v12 int = (fi11 * fi11)
                var v13 uint16 = 42
                var v14 int = (v12 | -2147483648)
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            _ = fi10
        }
    } else {
        var v15 string = "hello"
        if (!false) {
            v15 = "a"
        }
        v1 += 649
        _ = v15
    }
    if true {
        var v16 int32 = 0
        _ = v16
    } else {
        for fi17 := 0; fi17 < 4; fi17++ {
            for fi18 := 0; fi18 < 16; fi18++ {
                v1 += v1
                var v19 int8 = 1
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
        for fi20 := 0; fi20 < 18; fi20++ {
            var v21 uint16 = 2
            _ = fi20
            _ = v21
        }
    }
    fmt.Println(v1)
    fmt.Println(v9)
    v22 := GenS2{F0: 9223372036854775807, F1: (-32768 > int16(v1))}
    if true {
        v23 := GenS2{F0: 9223372036854775807, F1: true}
        _ = len(v2)
        _ = v23
    } else {
        for fi24 := 0; fi24 < 15; fi24++ {
            var v25 int = fi24
            v26 := GenS2{F0: -273, F1: true}
            if true {
                v27 := GenS3{F0: "\n\t", F1: 9223372036854775807, F2: 255}
                var v28 uint16 = 42
                _ = v27
                _ = v28
            }
            _ = fi24
            _ = v25
            _ = v26
        }
    }
    switch "foo bar" {
    case "":
        for fi29 := 0; fi29 < 7; fi29++ {
            for fi30 := 0; fi30 < 12; fi30++ {
                switch "hello" {
                case "":
                    var v31 bool = false
                    switch "" {
                    case "":
                        var v32 float64 = (0.0 + 1.5)
                        _ = v32
                    }
                    _ = v31
                case "a":
                    var v33 uint32 = v1
                    _ = v33
                }
                _ = fi30
            }
            v1 = (v1 ^ 1)
            _ = fi29
        }
    case "hello":
        var v34 float32 = 1.5
        _ = v34
    case "a":
        var v35 string = "hello"
        _ = v35
    default:
        v9 = (v9 + 26)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v9
}
