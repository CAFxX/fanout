package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 int64; F2 int }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 int32; F1 int8; F2 int64 }
func genFn0(p0 bool, p1 uint) int32 {
    if false {
        if (18446744073709551615 <= 18446744073709551615) {
            var v6 int64 = -1
            p1 += (p1 | 18446744073709551615)
            fmt.Println(v6)
            _ = v6
        } else {
            if (true || (-661 == 0)) {
                v7 := GenS1{F0: 622}
                _ = v7
            } else {
                p0 = (p1 > p1)
            }
        }
        p1 -= p1
        p1 += 470
    } else {
        var v8 int = 1
        if false {
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 float64 = float64((1e10 * (float32((-2.25 / 1.5)) - 3.14159)))
                v11 := GenS2{F0: fi9}
                _ = fi9
                _ = v10
                _ = v11
            }
            if p0 {
                v12 := []int{-21}
                var v13 uint64 = 18446744073709551615
                _ = v12
                _ = v13
            }
            v14 := []int{}
            _ = v14
        }
        v15 := GenS1{F0: 65535}
        _ = v15
        _ = v8
    }
    var v16 int32 = 175
    fmt.Println(v16)
    var v17 uint32 = 646
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return 644
}

func main() {
    var v1 uint64 = 1
    var v2 bool = false
    var v3 uint = 0
    v4 := []string{}
    v5 := make(map[string]string)
    if (true || false) {
        v3 += v3
        v18 := GenS3{F0: -1, F1: -63, F2: 42}
        _ = v18
    }
    fmt.Println(v1, v1, v1)
    for fi19 := 0; fi19 < 13; fi19++ {
        v20 := GenS3{F0: -32768, F1: 0, F2: (fi19 * fi19)}
        _ = fi19
        _ = v20
    }
    fmt.Println(v1, v3, v1)
    v21 := GenS4{F0: ((v3 - (v3 ^ v3)) != (v3 ^ (v3 & v3))), F1: 0}
    var v22 int = 1
    for fi23 := 0; fi23 < 17; fi23++ {
        fmt.Println(v1, v1)
        var v24 uint = (v3 * uint(v1))
        v22 = v22
        _ = fi23
        _ = v24
    }
    if v2 {
        v25 := GenS1{F0: uint16(v1)}
        var v26 uint8 = 0
        _ = v25
        _ = v26
    } else {
        var v27 string = ("foo bar" + "foo bar")
        if v2 {
            for fi28 := 0; fi28 < 4; fi28++ {
                for fi29 := 0; fi29 < 12; fi29++ {
                    if len(v4) > 0 {
                        v4[0] = v27
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi29
                }
                v30 := GenS5{F0: -1, F1: 42, F2: 42}
                _ = fi28
                _ = v30
            }
        }
        v2 = (!v2)
        _ = v27
    }
    var v31 int8 = 100
    var v32 string = "a"
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch v22 {
    case -2147483648:
        var v33 int8 = 0
        _ = v33
    case 32767:
        var v34 int16 = 362
        _ = v34
    case -1:
        var v35 uint64 = (v1 << (((v3 % (((v3 << (((^(v3))) & 63))) | 1))) & 63))
        _ = v35
    default:
        fmt.Println(v22)
    }
    for fi36 := 0; fi36 < 19; fi36++ {
        _ = len(v4)
        var v37 uint = v3
        _ = fi36
        _ = v37
    }
    v2 = true
    var v38 uint8 = 1
    v1 -= (v1 >> ((v3) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v38
    _ = v4
    _ = v5
}
