package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS5 struct { F0 uint; F1 int64; F2 int16 }
type GenS6 struct { F0 uint8 }
func genFn0(p0 uint64) uint8 {
    v6 := GenS1{F0: -1, F1: 2}
    p0 -= p0
    _ = p0
    _ = v6
    return 127
}

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = true
    v3 := make(map[uint32]int)
    var v4 uint16 = 65535
    var v5 uint8 = 127
    v7 := GenS2{F0: 541, F1: (v4 % ((v4) | 1))}
    if (false || (v2 || (v2 || (-1 <= -515)))) {
        for fi8 := 0; fi8 < 11; fi8++ {
            v3[2] = (fi8 * 128)
            _ = fi8
        }
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 uint16 = (v4 >> ((4294967295) & 63))
            _ = fi9
            _ = v10
        }
        delete(v1, 33)
    } else {
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 string = "hello"
            _ = fi11
            _ = v12
        }
        v13 := GenS1{F0: -719, F1: (v4 | (v4 * (^(v4))))}
        _ = v13
    }
    fmt.Println(v5, v4)
    for fi14 := 0; fi14 < 8; fi14++ {
        if v2 {
            var v15 uint32 = 236
            var v16 int8 = -83
            v17 := []int{(fi14 - (^(fi14))), fi14}
            _ = v15
            _ = v16
            _ = v17
        } else {
            v18 := GenS3{F0: -1}
            var v19 uint = 18446744073709551615
            _ = v18
            _ = v19
        }
        var v20 int16 = -32768
        _ = fi14
        _ = v20
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        v22 := []int{}
        v5 += 2
        for fi23 := 0; fi23 < 6; fi23++ {
            if (true || ((42 >= 42) && (42 >= (v4 | v4)))) {
                var v24 string = "\n\t"
                if v2 {
                    var v25 float32 = -2.25
                    _ = v25
                } else {
                    v2 = (v2 && (!(v24 >= v24)))
                    var v26 uint16 = v4
                    v4 = (v4 | (v26 ^ (v4 ^ (^(v26)))))
                    _ = v26
                }
                switch (((v24 + v24) + v24) + v24) {
                case "hello":
                    v24 = "foo bar"
                case "\n\t":
                    v5 += 156
                default:
                    delete(v1, 492)
                }
                _ = v24
            } else {
                var v27 uint16 = v4
                v5 = (v5 >> ((2) & 63))
                v28 := GenS4{F0: 1, F1: 2, F2: 255}
                _ = v27
                _ = v28
            }
            v2 = true
            v5 += v5
            _ = fi23
        }
        _ = fi21
        _ = v22
    }
    v29 := GenS1{F0: -2147483648, F1: v4}
    var v30 int64 = -9223372036854775808
    v2 = (!(572 < 1))
    v31 := GenS4{F0: 105, F1: 577, F2: 65535}
    var v32 uint32 = 2147483647
    v33 := GenS5{F0: 65535, F1: v30, F2: -32768}
    fmt.Println(v2, v5)
    for fi34 := 0; fi34 < 9; fi34++ {
        v35 := GenS6{F0: uint8(v32)}
        _ = fi34
        _ = v35
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        fmt.Println(fi36)
        var v37 uint = 665
        if v2 {
            var v38 bool = true
            _ = v38
        } else {
            fmt.Println(fi36)
            v5 += genFn0(0)
        }
        _ = fi36
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
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
    _ = v7
}
