package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int, p1 int) uint {
    var v7 uint8 = 2
    var v8 int64 = 439
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 255
}

func main() {
    var v1 uint32 = 1
    var v2 uint16 = 255
    var v3 uint8 = 82
    var v4 uint32 = v1
    v5 := make(map[string]uint32)
    var v6 uint = 593
    _ = len(v5)
    _ = len(v5)
    var v9 uint64 = uint64(v3)
    _ = len(v5)
    v4 -= uint32(v2)
    if false {
        v10 := GenS1{F0: 0, F1: -1}
        if true {
            var v11 int32 = -1
            if ((!false) && true) {
                v12 := GenS1{F0: 9223372036854775807, F1: v11}
                var v13 float64 = -2.25
                _ = v12
                _ = v13
            } else {
                _ = len(v5)
                switch 2147483647 {
                case 1:
                    if (986 != -584) {
                        fmt.Println(v2, v3)
                        v14 := GenS1{F0: 1, F1: 229}
                        fmt.Println(v11)
                        _ = v14
                    } else {
                        v4 += 65535
                        v6 = 18446744073709551615
                    }
                case -2147483648:
                    var v15 uint8 = v3
                    _ = v15
                case 9223372036854775807:
                    if true {
                        switch "hello" {
                        case "foo bar":
                            var v16 string = "a"
                            _ = v16
                        case "hello":
                            for fi17 := 0; fi17 < 18; fi17++ {
                                v18 := GenS2{F0: (^(v2))}
                                _ = fi17
                                _ = v18
                            }
                        case "\n\t":
                            var v19 uint64 = 1
                            _ = v19
                        }
                        if (true && false) {
                            var v20 int16 = 0
                            if (4294967295 != 32767) {
                                var v21 uint = (^(v6))
                                var v22 uint = (^(v21))
                                _ = v21
                                _ = v22
                            } else {
                                var v23 uint16 = (v2 / ((uint16(v6)) | 1))
                                var v24 int16 = 620
                                _ = v23
                                _ = v24
                            }
                            var v25 uint8 = v3
                            _ = v20
                            _ = v25
                        } else {
                            v26 := []int{2147483647}
                            fmt.Println(v6, v2, v4)
                            var v27 int8 = 0
                            _ = v26
                            _ = v27
                        }
                    }
                }
            }
            _ = v11
        }
        _ = v10
    }
    switch "foo bar" {
    case "hello":
        var v28 int64 = 1
        _ = v28
    case "\n\t":
        var v29 uint64 = 1
        _ = v29
    default:
        var v30 int64 = -9223372036854775808
        _ = v30
    }
    v31 := v5[(("" + "foo bar") + "hello")]
    v3 -= v3
    v5["foo bar"] = (v4 << ((65535) & 63))
    v32 := GenS1{F0: 0, F1: 0}
    fmt.Println(v2, v2, v4)
    v1 += (v1 >> ((18446744073709551615) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
