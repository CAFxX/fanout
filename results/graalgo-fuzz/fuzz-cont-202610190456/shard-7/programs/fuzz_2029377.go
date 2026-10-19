package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
func genFn0(p0 uint64) int8 {
    if (!true) {
        v5 := GenS1{F0: 1, F1: 0, F2: p0}
        v6 := GenS1{F0: 1, F1: 0, F2: p0}
        _ = v5
        _ = v6
    }
    v7 := []int{9223372036854775807}
    var v8 int64 = -9223372036854775808
    _ = p0
    _ = v7
    _ = v8
    return 100
}

func main() {
    var v1 int8 = -1
    var v2 bool = false
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]uint32)
    var v9 uint8 = 0
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 uint32 = 0
        v4[(v11 * (v11 ^ v11))] = 2
        v12 := GenS2{F0: (v1 & v1)}
        _ = fi10
        _ = v11
        _ = v12
    }
    if true {
        switch 4294967295 {
        case 127:
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 string = (("a" + "foo bar") + (("hello" + "") + ("hello" + "hello")))
                v4[2147483647] = 595
                _ = fi13
                _ = v14
            }
        case 0:
            fmt.Println(v9, v2)
        default:
            var v15 uint8 = (v9 / ((v9) | 1))
            _ = v15
        }
        if false {
            v16 := GenS2{F0: v1}
            _ = v16
        } else {
            for fi17 := 0; fi17 < 4; fi17++ {
                if v2 {
                    if (((1e10 + 1e10) * 0.1) == 1e10) {
                        v18 := GenS3{F0: 18446744073709551615}
                        var v19 uint16 = 0
                        var v20 float64 = 0.0
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    delete(v4, 4294967295)
                } else {
                    v21 := GenS2{F0: genFn0(94)}
                    for fi22 := 0; fi22 < 6; fi22++ {
                        v23 := GenS1{F0: 42, F1: -32768, F2: 774}
                        _ = fi22
                        _ = v23
                    }
                    if v2 {
                        var v24 float64 = 0.1
                        if ((-(fi17)) >= (fi17 | -2147483648)) {
                            v25 := GenS1{F0: -574, F1: 32767, F2: 0}
                            _ = v25
                        }
                        _ = v24
                    } else {
                        v1 -= v1
                        v9 -= (v9 + (v9 + 208))
                        v26 := v4[2]
                        _ = v26
                    }
                    _ = v21
                }
                v27 := GenS3{F0: 39}
                _ = fi17
                _ = v27
            }
        }
    }
    if v2 {
        v28 := []int{42, -9223372036854775808, 42}
        for fi29 := 0; fi29 < 15; fi29++ {
            if false {
                fmt.Println(v1, v2)
            }
            _ = fi29
        }
        if len(v28) > 0 {
            v28[0] = 0
        }
        _ = v28
    } else {
        switch 42 {
        case 32767:
            v1 -= genFn0(4294967295)
        default:
            v30 := GenS1{F0: 1, F1: -162, F2: 0}
            _ = v30
        }
    }
    var v31 uint32 = uint32(v9)
    for fi32 := 0; fi32 < 1; fi32++ {
        var v33 float32 = 3.14159
        v34 := GenS2{F0: (v1 << ((555) & 63))}
        fmt.Println(v33, v31, v31)
        _ = fi32
        _ = v33
        _ = v34
    }
    v31 = 608
    v31 -= v31
    if false {
        for fi35 := 0; fi35 < 14; fi35++ {
            var v36 uint32 = (v31 ^ (v31 + (v31 | (v31 << ((255) & 63)))))
            v37 := GenS1{F0: -507, F1: 1000, F2: 0}
            var v38 int8 = v1
            _ = fi35
            _ = v36
            _ = v37
            _ = v38
        }
        v3[416] = 42
    }
    v1 -= (v1 / ((genFn0(181)) | 1))
    var v39 int8 = (v1 >> ((929) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v39
    _ = v4
    _ = v9
}
