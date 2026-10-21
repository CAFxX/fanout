package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 uint; F2 uint32 }
func genFn0(p0 uint16, p1 int) int8 {
    if false {
        fmt.Println(p1, p0)
        v6 := GenS1{F0: "\n\t"}
        var v7 int32 = 964
        _ = v6
        _ = v7
    } else {
        p0 = (p0 | p0)
        if false {
            var v8 bool = (true && true)
            fmt.Println(v8, p1)
            if (4294967295 != 186) {
                fmt.Println(v8)
            } else {
                var v9 uint8 = 42
                fmt.Println(v9, v9)
                v9 = (v9 - (v9 - v9))
                _ = v9
            }
            _ = v8
        } else {
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 uint64 = 1
                _ = fi10
                _ = v11
            }
            p1 = p1
        }
        for fi12 := 0; fi12 < 12; fi12++ {
            for fi13 := 0; fi13 < 19; fi13++ {
                v14 := []int{(fi13 << ((255) & 63)), (fi13 | (fi12 - (^(fi12))))}
                _ = fi13
                _ = v14
            }
            if false {
                if (987 < 0) {
                    if true {
                        for fi15 := 0; fi15 < 18; fi15++ {
                            p1 += (fi15 * (^(fi15)))
                            var v16 int32 = 1
                            fmt.Println(p1, fi15)
                            v16 = (^(v16))
                            _ = fi15
                            _ = v16
                        }
                        p1 += (p1 - 42)
                        var v17 int8 = -33
                        _ = v17
                    }
                }
                fmt.Println(fi12, p1, p0)
            }
            _ = fi12
        }
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        v19 := []int{}
        p1 = (p1 - (^(fi18)))
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        switch 9223372036854775807 {
        case -690:
            fmt.Println(fi20, p0, fi20)
        case 933:
            if true {
                p1 -= p1
                p1 -= (^(p1))
                p1 -= fi20
            }
        default:
            v21 := GenS2{F0: 18446744073709551615, F1: 18446744073709551615, F2: 4294967295}
            _ = v21
        }
        _ = fi20
    }
    if (127 > 62) {
        var v22 float32 = 0.1
        var v23 int16 = 285
        var v24 uint64 = 18446744073709551615
        _ = v22
        _ = v23
        _ = v24
    } else {
        if (p0 != 2) {
            var v25 string = "\n\t"
            v26 := []int{9223372036854775807, (^(p1)), 32768, 950}
            _ = v25
            _ = v26
        } else {
            var v27 uint = 65535
            p1 = (p1 % ((p1) | 1))
            var v28 bool = false
            _ = v27
            _ = v28
        }
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 uint32 = 0
    v2 := []int{-382, 127, 42}
    v3 := make(map[string]int)
    v4 := []string{("" + "hello"), "", "foo bar"}
    var v5 uint16 = 1
    fmt.Println(v5)
    var v29 bool = (float64(float32(1.5)) <= (3.14159 * (1e10 + 3.14159)))
    var v30 uint = 4294967295
    v5 -= (^(v5))
    v29 = v29
    fmt.Println(v1, v29)
    fmt.Println(v1)
    v31 := GenS1{F0: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
