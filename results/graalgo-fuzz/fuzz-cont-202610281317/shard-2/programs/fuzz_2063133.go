package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 int16 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]uint32)
    var v3 int8 = 10
    var v4 int32 = 1
    var v5 uint8 = 13
    switch "hello" {
    case "a":
        for fi6 := 0; fi6 < 19; fi6++ {
            v7 := GenS1{F0: v3}
            v5 += v5
            if (!((true && (0.0 <= (1e10 * 3.14159))) || ((v5 == (^(v5))) && (!false)))) {
                v8 := []int{(fi6 / ((-9223372036854775808) | 1))}
                var v9 int = fi6
                _ = v8
                _ = v9
            } else {
                v10 := GenS2{F0: -881}
                if (!false) {
                    var v11 uint64 = 456
                    _ = v11
                } else {
                    var v12 int64 = -284
                    v5 = (v5 / ((209) | 1))
                    v4 += (v4 << ((18446744073709551615) & 63))
                    _ = v12
                }
                for fi13 := 0; fi13 < 19; fi13++ {
                    if ((!((-(v3)) != (v3 + v3))) && (false && true)) {
                        v14 := GenS3{F0: v5}
                        _ = v14
                    }
                    switch "" {
                    case "\n\t":
                        if (!true) {
                            switch "foo bar" {
                            case "hello":
                                v15 := GenS3{F0: (v5 ^ v5)}
                                _ = v15
                            default:
                                if false {
                                    if true {
                                        var v16 uint64 = 42
                                        _ = v16
                                    } else {
                                        v4 -= (v4 + (v4 << ((253) & 63)))
                                        var v17 uint64 = 18446744073709551615
                                        v18 := GenS2{F0: -1}
                                        _ = v17
                                        _ = v18
                                    }
                                    v19 := GenS3{F0: v5}
                                    _ = len(v2)
                                    v20 := GenS4{F0: -1, F1: 4294967295}
                                    _ = v19
                                    _ = v20
                                }
                            }
                        }
                    case "hello":
                        switch 127 {
                        case 32767:
                            v21 := v2[42]
                            _ = v21
                        case -2:
                            v22 := GenS5{F0: 42}
                            _ = v22
                        default:
                            var v23 int32 = (v4 | int32(v5))
                            _ = v23
                        }
                    }
                    _ = fi13
                }
                _ = v10
            }
            _ = fi6
            _ = v7
        }
    default:
        v3 = v3
    }
    var v24 int = 2147483647
    var v25 int32 = (v4 ^ (v4 & v4))
    v26 := []int{593}
    for fi27 := 0; fi27 < 1; fi27++ {
        v28 := GenS5{F0: 0}
        var v29 int16 = 1000
        _ = fi27
        _ = v28
        _ = v29
    }
    fmt.Println(v25, v5, v5)
    fmt.Println(v25, v4)
    v24 += (-(v24))
    fmt.Println(v24)
    var v30 int64 = -1
    var v31 int16 = 42
    v32 := v1[v24]
    var v33 uint32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
