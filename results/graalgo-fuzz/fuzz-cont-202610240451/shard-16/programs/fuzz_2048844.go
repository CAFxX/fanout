package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint }

func main() {
    var v1 uint16 = 1
    v2 := make(map[uint32]int)
    var v3 uint64 = 18446744073709551615
    v4 := make(map[int]uint32)
    v5 := make(map[string]bool)
    var v6 uint16 = v1
    v3 += v3
    v7 := GenS1{F0: ("a" + "a"), F1: 32768, F2: 1}
    v8 := GenS1{F0: "foo bar", F1: 127, F2: 0}
    _ = len(v5)
    switch "\n\t" {
    case "foo bar":
        var v9 int8 = 0
        _ = v9
    case "hello":
        v3 -= v3
    default:
        if (-2.25 > ((0.1 / (0.1 + 0.1)) - 3.14159)) {
            if (v3 == (v3 & (^(v3)))) {
                v10 := GenS2{F0: -1}
                for fi11 := 0; fi11 < 8; fi11++ {
                    var v12 int64 = 1
                    if ((v3 - 42) > 0) {
                        for fi13 := 0; fi13 < 19; fi13++ {
                            v6 = v6
                            if (v3 != (v3 / ((v3) | 1))) {
                                var v14 bool = (!(fi11 < fi11))
                                _ = v14
                            } else {
                                v15 := []int{(fi11 * -820), fi13}
                                var v16 uint64 = 0
                                var v17 uint64 = 18446744073709551615
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                            v3 = (v3 + v3)
                            _ = fi13
                        }
                        for fi18 := 0; fi18 < 20; fi18++ {
                            if (!(!false)) {
                                switch ("\n\t" + "\n\t") {
                                case "a":
                                    v6 -= (^(v1))
                                default:
                                    if (false && true) {
                                        fmt.Println(v3, v6)
                                        v3 -= 696
                                        var v19 int8 = 119
                                        var v20 uint = 251
                                        _ = len(v4)
                                        _ = v19
                                        _ = v20
                                    } else {
                                        if true {
                                            v6 += (v6 & v1)
                                            fmt.Println(v3)
                                        }
                                        v21 := GenS1{F0: "a", F1: (-(fi18)), F2: 42}
                                        _ = v21
                                    }
                                }
                            }
                            if false {
                                var v22 string = "a"
                                var v23 int64 = (v12 * 9223372036854775807)
                                v24 := GenS2{F0: (-(v12))}
                                v23 += 9223372036854775807
                                _ = v22
                                _ = v23
                                _ = v24
                            } else {
                                var v25 int32 = -910
                                var v26 uint16 = v6
                                v25 += (v25 | v25)
                                _ = v25
                                _ = v26
                            }
                            var v27 uint32 = 457
                            _ = fi18
                            _ = v27
                        }
                    }
                    _ = fi11
                    _ = v12
                }
                _ = v10
            } else {
                v6 += (v1 * v6)
            }
            switch "hello" {
            case "hello":
                v28 := GenS3{F0: 2147483647, F1: 1, F2: 468}
                _ = v28
            case "\n\t":
                var v29 uint8 = 127
                _ = v29
            default:
                v6 = 42
            }
            for fi30 := 0; fi30 < 8; fi30++ {
                var v31 string = "foo bar"
                _ = fi30
                _ = v31
            }
        }
    }
    v3 += (v3 | v3)
    var v32 uint64 = 47
    v3 = (v32 / (((^(v32))) | 1))
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
