package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 string }

func main() {
    var v1 int32 = -251
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    v4 := []int{}
    v1 -= 663
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 float64 = (float64(float32(0.1)) - 1.5)
        _ = fi5
        _ = v6
    }
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    fmt.Println(v1)
    fmt.Println(v1)
    if (127 > 225) {
        fmt.Println(v1)
        fmt.Println(v1)
    }
    var v7 float64 = 0.1
    v7 += 1e10
    var v8 int64 = -1
    v9 := GenS1{F0: 116}
    v8 = (v8 << ((519) & 63))
    v8 = (v8 << ((18446744073709551615) & 63))
    var v10 uint8 = 42
    v7 = ((-2.25 - ((v7 * float64(0.1)) / ((v7 + (v7 + v7)) * 0.1))) - v7)
    v11 := GenS2{F0: "foo bar", F1: ""}
    var v12 int16 = -32768
    v1 = int32(v12)
    switch 9223372036854775807 {
    case 128:
        v13 := GenS2{F0: "\n\t", F1: "a"}
        _ = v13
    case 0:
        v10 = v10
    case -422:
        v14 := GenS2{F0: "foo bar", F1: "hello"}
        _ = v14
    default:
        switch "" {
        case "hello":
            if len(v4) > 0 {
                v4[0] = 0
            }
        case "foo bar":
            for fi15 := 0; fi15 < 12; fi15++ {
                var v16 string = ""
                v16 = (v16 + (v16 + ""))
                var v17 uint = 18446744073709551615
                _ = fi15
                _ = v16
                _ = v17
            }
        case "\n\t":
            v18 := GenS1{F0: v10}
            _ = v18
        default:
            if true {
                if true {
                    if false {
                        var v19 int = 128
                        _ = len(v4)
                        _ = v19
                    } else {
                        v8 = v8
                    }
                    switch "hello" {
                    case "a":
                        switch ("\n\t" + "hello") {
                        case "\n\t":
                            v7 -= float64(1e10)
                        default:
                            v3[381] = (int64(v12) > 1)
                        }
                    }
                    for fi20 := 0; fi20 < 16; fi20++ {
                        for fi21 := 0; fi21 < 4; fi21++ {
                            if true {
                                var v22 bool = false
                                v10 -= (v10 ^ 42)
                                _ = v22
                            }
                            v7 -= ((v7 / (v7 / (v7 * 3.14159))) / (1e10 / (v7 + 1e10)))
                            if (2 <= 333) {
                                var v23 uint16 = 593
                                fmt.Println(fi21, v7)
                                var v24 string = "hello"
                                var v25 string = (v24 + (("hello" + v24) + "a"))
                                v8 = (^(v8))
                                _ = v23
                                _ = v24
                                _ = v25
                            } else {
                                v1 += v1
                                var v26 int16 = 1
                                _ = v26
                            }
                            if false {
                                v10 += (v10 >> ((18446744073709551615) & 63))
                            } else {
                                var v27 bool = false
                                _ = v27
                            }
                            v12 -= v12
                            delete(v2, int(v8))
                            _ = fi21
                        }
                        fmt.Println(v1, v10)
                        _ = fi20
                    }
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
