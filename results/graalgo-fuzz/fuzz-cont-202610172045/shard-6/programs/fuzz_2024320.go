package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint32 }
type GenS2 struct { F0 uint; F1 bool; F2 int8 }
func genFn0(p0 int, p1 uint64, p2 int) uint {
    switch (p0 + (p0 - p0)) {
    case 9223372036854775807:
        var v6 bool = true
        _ = v6
    case 0:
        var v7 uint = 579
        _ = v7
    default:
        v8 := []int{(p0 >> ((0) & 63))}
        _ = v8
    }
    p2 -= p0
    var v9 uint64 = p1
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 uint8 = 127
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[string]int)
    var v3 int32 = 0
    v4 := make(map[uint32]bool)
    var v5 int64 = 0
    for fi12 := 0; fi12 < 12; fi12++ {
        if (!true) {
            if ((-1 > 0) || (42 <= 1)) {
                v3 = v3
                for fi13 := 0; fi13 < 5; fi13++ {
                    if false {
                        if ((v3 != v3) && true) {
                            var v14 float32 = (0.1 + 0.0)
                            v15 := GenS1{F0: 0, F1: int(v5), F2: 2}
                            v16 := GenS1{F0: -32768, F1: -930, F2: 289}
                            delete(v4, 0)
                            var v17 uint8 = (v1 | v1)
                            _ = v14
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                    } else {
                        if (-74 != 0) {
                            v18 := GenS2{F0: 535, F1: false, F2: -1}
                            if false {
                                v19 := []int{(fi13 % (((fi13 & fi12)) | 1)), 32768}
                                var v20 uint = 2
                                v21 := GenS1{F0: -1, F1: (-(fi12)), F2: 0}
                                _ = v19
                                _ = v20
                                _ = v21
                            } else {
                                switch (fi13 % (((fi13 + (fi12 - int(v1)))) | 1)) {
                                case 128:
                                    var v22 float64 = 0.1
                                    var v23 int8 = -59
                                    _ = v22
                                    _ = v23
                                case -9223372036854775808:
                                    if (!((float32(float64((1.5 * 3.14159))) - (-2.25 / 0.1)) != -2.25)) {
                                        v5 = (v5 + (v5 << ((962) & 63)))
                                    }
                                case 9223372036854775807:
                                    _ = len(v4)
                                default:
                                    var v24 bool = true
                                    _ = v24
                                }
                                v25 := []int{(fi12 % ((fi13) | 1)), (fi12 - fi12), (fi12 << ((genFn0(fi12, 420, (fi12 << ((255) & 63)))) & 63)), fi13}
                                _ = v25
                            }
                            _ = v18
                        }
                        v26 := GenS1{F0: -557, F1: (-(fi13)), F2: 65535}
                        _ = v26
                    }
                    fmt.Println(fi12)
                    if false {
                        var v27 uint64 = 2
                        _ = v27
                    } else {
                        var v28 float32 = (3.14159 / 1e10)
                        _ = v28
                    }
                    _ = fi13
                }
            }
        } else {
            v3 -= 2147483647
        }
        _ = fi12
    }
    v29 := GenS2{F0: 4294967295, F1: ((77 > -67) || (true || false)), F2: 42}
    var v30 bool = true
    var v31 uint = 4294967295
    fmt.Println(v3, v30, v1)
    var v32 uint64 = 0
    switch "a" {
    case "\n\t":
        v30 = (v31 < (^(v31)))
    case "hello":
        fmt.Println(v3, v5, v1)
    }
    if true {
        v30 = true
    }
    var v33 uint32 = 2147483648
    v31 = v31
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
