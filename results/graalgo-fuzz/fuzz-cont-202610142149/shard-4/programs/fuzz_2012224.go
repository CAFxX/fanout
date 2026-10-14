package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 uint64, p1 string) int {
    for fi5 := 0; fi5 < 2; fi5++ {
        fmt.Println(p1)
        switch fi5 {
        case 32767:
            if true {
                v6 := GenS1{F0: 2147483647, F1: -59, F2: 65535}
                for fi7 := 0; fi7 < 19; fi7++ {
                    v8 := []int{}
                    p0 -= p0
                    _ = fi7
                    _ = v8
                }
                _ = v6
            } else {
                fmt.Println(p0)
                var v9 int32 = -51
                _ = v9
            }
        case 9223372036854775807:
            var v10 uint16 = 1
            _ = v10
        default:
            p0 += (p0 & p0)
        }
        _ = fi5
    }
    var v11 float32 = float32(0.1)
    _ = p0
    _ = p1
    _ = v11
    return -68
}

func main() {
    var v1 uint16 = 42
    var v2 uint32 = 2147483648
    v3 := make(map[uint32]bool)
    v4 := make(map[string]int)
    var v12 string = "\n\t"
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 uint8 = 255
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: 181, F1: -40, F2: uint(v1)}
    var v16 uint16 = (v1 & (^(v1)))
    v17 := []int{4294967295, genFn0(2, "foo bar"), genFn0(18446744073709551615, "a"), -1}
    fmt.Println(v12, v16, v16)
    if (v2 != 42) {
        switch -1 {
        case 32767:
            for fi18 := 0; fi18 < 7; fi18++ {
                fmt.Println(v2, v2)
                for fi19 := 0; fi19 < 20; fi19++ {
                    v12 = "hello"
                    v20 := GenS2{F0: 406}
                    _ = fi19
                    _ = v20
                }
                switch v12 {
                case "foo bar":
                    for fi21 := 0; fi21 < 14; fi21++ {
                        var v22 uint64 = 18446744073709551615
                        v23 := GenS3{F0: 251}
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                case "a":
                    v12 = "hello"
                default:
                    for fi24 := 0; fi24 < 20; fi24++ {
                        var v25 float64 = (0.0 * (1.5 + float64(3.14159)))
                        _ = fi24
                        _ = v25
                    }
                }
                _ = fi18
            }
        default:
            var v26 uint = 4294967295
            _ = v26
        }
        var v27 float64 = float64((float32(((1.5 * (3.14159 - float64(-2.25))) - 0.0)) + (3.14159 + (0.0 + 0.0))))
        _ = v27
    } else {
        v1 += 0
        var v28 uint = 65535
        _ = v28
    }
    switch "" {
    case "hello":
        _ = len(v17)
    case "foo bar":
        v16 = v1
    default:
        if (v2 <= (v2 >> ((18446744073709551615) & 63))) {
            for fi29 := 0; fi29 < 20; fi29++ {
                v12 = v12
                var v30 int16 = 32767
                var v31 uint8 = 137
                _ = fi29
                _ = v30
                _ = v31
            }
            fmt.Println(v16, v1)
            v2 = v2
        } else {
            fmt.Println(v1)
        }
    }
    for fi32 := 0; fi32 < 9; fi32++ {
        v1 = 32767
        var v33 uint8 = 2
        _ = fi32
        _ = v33
    }
    v34 := GenS4{F0: v1}
    v4[""] = -985
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
