package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 int; F1 uint8 }
func genFn0() uint {
    var v4 bool = true
    v4 = (109 > 1)
    v4 = (239 == 0)
    _ = v4
    return 42
}
func genFn1(p0 uint64, p1 string, p2 int16) int8 {
    for fi5 := 0; fi5 < 15; fi5++ {
        if (uint8(p0) != 207) {
            p2 += p2
            var v6 float32 = -2.25
            var v7 int8 = 126
            _ = v6
            _ = v7
        } else {
            if false {
                p0 -= (p0 | (p0 | (p0 >> ((65535) & 63))))
                fmt.Println(p0, fi5, p1)
                if (2 == p0) {
                    v8 := GenS1{F0: 1, F1: (p2 & p2)}
                    _ = v8
                }
            } else {
                v9 := []int{32767}
                p1 = p1
                p2 = p2
                _ = v9
            }
        }
        var v10 int = fi5
        for fi11 := 0; fi11 < 14; fi11++ {
            fmt.Println(v10, fi5)
            p1 = (p1 + "foo bar")
            _ = fi11
        }
        _ = fi5
        _ = v10
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    return 84
}

func main() {
    var v1 int = 507
    var v2 float32 = float32(float64(0.1))
    v3 := make(map[uint32]string)
    for fi12 := 0; fi12 < 16; fi12++ {
        if false {
            var v13 uint16 = 903
            var v14 float32 = ((((v2 / ((float32(1.5) / 1.0) - v2)) / (v2 - v2)) * 0.0) * v2)
            _ = v13
            _ = v14
        } else {
            v1 = fi12
        }
        v2 = v2
        v2 += v2
        _ = fi12
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        delete(v3, 0)
        var v16 int = (v1 | fi15)
        switch "a" {
        case "hello":
            v2 -= float32((3.14159 - 1e10))
        case "a":
            for fi17 := 0; fi17 < 15; fi17++ {
                if true {
                    fmt.Println(fi17, v1)
                    if false {
                        v18 := GenS2{F0: -1, F1: 376}
                        var v19 float32 = v2
                        v16 = v16
                        _ = v18
                        _ = v19
                    }
                    fmt.Println(fi17, v1)
                }
                _ = len(v3)
                _ = fi17
            }
        case "\n\t":
            v3[456] = ""
        }
        _ = fi15
        _ = v16
    }
    var v20 uint16 = 42
    v21 := []int{42, v1}
    delete(v3, 2)
    for fi22 := 0; fi22 < 10; fi22++ {
        if false {
            v23 := GenS3{F0: 2, F1: 42}
            if (((v1 | (v1 - 32767)) > (fi22 + -1)) || true) {
                v1 = (fi22 >> ((genFn0()) & 63))
                v24 := GenS4{F0: v1, F1: 62}
                v20 = v20
                _ = v24
            } else {
                fmt.Println(fi22, v2, v20)
            }
            _ = v23
        }
        _ = fi22
    }
    v1 -= v1
    fmt.Println(v2, v2)
    fmt.Println(v1, v1, v20)
    var v25 int16 = int16(v20)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
}
