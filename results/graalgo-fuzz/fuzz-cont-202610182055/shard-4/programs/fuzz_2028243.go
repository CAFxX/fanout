package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 uint64; F1 uint; F2 int }
type GenS5 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS6 struct { F0 int16; F1 uint32 }
type GenS7 struct { F0 uint64; F1 int8; F2 int }
func genFn0(p0 string) uint32 {
    fmt.Println(p0)
    var v4 uint = 597
    v4 = (v4 & (^(v4)))
    v5 := GenS1{F0: uint32(v4)}
    _ = p0
    _ = v4
    _ = v5
    return 2147483647
}
func genFn1(p0 uint16) uint32 {
    if (p0 < 42) {
        var v6 int32 = -1
        v7 := GenS1{F0: genFn0((("hello" + "a") + "foo bar"))}
        _ = v6
        _ = v7
    } else {
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 float32 = 1.5
            _ = fi8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        p0 = 1
        _ = fi10
    }
    p0 += (p0 ^ (^(p0)))
    p0 = 1
    var v11 uint32 = 76
    _ = p0
    _ = v11
    return (v11 / ((778) | 1))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 0
    v3 := make(map[int]int)
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 uint8 = 1
        v14 := GenS2{F0: (("" + ("\n\t" + "")) + "\n\t"), F1: 12}
        var v15 int64 = 0
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 bool = (false || (1.5 >= float64(3.14159)))
    var v17 int32 = 42
    v17 = 797
    v18 := GenS3{F0: -32768, F1: 18446744073709551615}
    if (2 > 255) {
        if ((0.1 >= (((float64(1e10) - 1e10) + (0.0 * 1.5)) / (0.1 - 1e10))) && v16) {
            v19 := GenS4{F0: 2, F1: 4294967295, F2: 9223372036854775807}
            v1[42] = v2
            _ = v19
        } else {
            fmt.Println(v17, v16, v2)
        }
        var v20 uint32 = (v2 << ((454) & 63))
        v17 -= v17
        _ = v20
    } else {
        v21 := GenS1{F0: v2}
        for fi22 := 0; fi22 < 17; fi22++ {
            for fi23 := 0; fi23 < 5; fi23++ {
                v24 := GenS5{F0: -173, F1: 116, F2: 127}
                v25 := GenS6{F0: 1, F1: v2}
                v16 = v16
                if v16 {
                    v26 := []int{(fi23 & (fi23 | fi22)), fi22, 127, fi23}
                    _ = v26
                }
                if false {
                    v27 := v1[fi22]
                    _ = v27
                } else {
                    v2 = v2
                    fmt.Println(v17, fi23, v16)
                }
                v17 = (v17 | v17)
                _ = fi23
                _ = v24
                _ = v25
            }
            v17 += v17
            v2 += v2
            _ = fi22
        }
        _ = v21
    }
    if v16 {
        v3[-297] = 9223372036854775807
        fmt.Println(v2, v17, v17)
        if ((!true) || v16) {
            _ = len(v3)
            var v28 uint8 = 128
            _ = v28
        } else {
            var v29 int64 = 42
            var v30 int = 32767
            _ = v29
            _ = v30
        }
    }
    fmt.Println(v2, v16)
    v31 := GenS6{F0: 1000, F1: v2}
    v2 -= v2
    if false {
        var v32 int64 = -321
        if false {
            fmt.Println(v2)
            delete(v3, -602)
            for fi33 := 0; fi33 < 3; fi33++ {
                v17 += (v17 ^ (v17 * v17))
                if v16 {
                    var v34 int32 = (v17 / (((v17 >> ((255) & 63))) | 1))
                    if v16 {
                        var v35 uint64 = 42
                        _ = v35
                    }
                    var v36 bool = (true && (v32 == v32))
                    _ = v34
                    _ = v36
                } else {
                    v37 := []int{fi33}
                    var v38 uint16 = 42
                    fmt.Println(v2)
                    _ = v37
                    _ = v38
                }
                _ = fi33
            }
        } else {
            delete(v1, 4294967295)
            v32 += 172
            var v39 string = ((("\n\t" + "hello") + "foo bar") + "a")
            _ = v39
        }
        _ = v32
    }
    v40 := GenS7{F0: 4294967295, F1: 3, F2: -2147483648}
    v16 = v16
    for fi41 := 0; fi41 < 18; fi41++ {
        fmt.Println(fi41)
        v42 := []int{-505, -2147483648, 32767}
        _ = fi41
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v40
}
