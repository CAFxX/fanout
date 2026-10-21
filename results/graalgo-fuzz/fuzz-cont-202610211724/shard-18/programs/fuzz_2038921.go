package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 int; F1 int16; F2 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 int8, p1 uint, p2 int32) uint32 {
    for fi4 := 0; fi4 < 2; fi4++ {
        var v5 uint = (p1 << ((p1) & 63))
        var v6 float32 = 0.0
        v6 += (1.5 + (v6 / (1e10 / 0.1)))
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := []int{0, -2147483648, 72}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 4294967295
}
func genFn1(p0 int16) uint64 {
    p0 = 891
    v8 := GenS1{F0: 1, F1: 18446744073709551615}
    _ = p0
    _ = v8
    return 4294967295
}

func main() {
    var v1 uint16 = 0
    var v2 uint64 = 408
    v3 := make(map[string]string)
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS1{F0: (v2 * (v2 & genFn1(-809))), F1: 42}
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: 1, F1: 550}
    if true {
        var v12 uint64 = (v2 ^ (v2 + (v2 >> ((65535) & 63))))
        v13 := GenS1{F0: v12, F1: uint(v1)}
        _ = v12
        _ = v13
    } else {
        v14 := GenS3{F0: -2147483648, F1: 708, F2: 0}
        _ = v14
    }
    if false {
        switch 204 {
        case -9223372036854775808:
            var v15 int = -689
            _ = v15
        }
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 string = "\n\t"
            fmt.Println(fi16, fi16)
            var v18 float64 = 3.14159
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    var v19 string = "a"
    v19 = v19
    if (!((v1 & (v1 << ((65535) & 63))) > v1)) {
        v20 := GenS3{F0: -984, F1: 42, F2: 70}
        v3[((v19 + ((v19 + v19) + v19)) + ((v19 + v19) + v19))] = (v19 + "")
        _ = v20
    } else {
        v21 := GenS4{F0: -1}
        var v22 uint8 = 255
        _ = v21
        _ = v22
    }
    v19 = (v19 + v19)
    for fi23 := 0; fi23 < 11; fi23++ {
        v3["\n\t"] = v19
        _ = fi23
    }
    if false {
        if false {
            v19 = ""
        } else {
            v24 := []int{-2147483648, 2147483647}
            _ = v24
        }
    } else {
        if ((!false) || true) {
            var v25 uint64 = (^(v2))
            v26 := v3[v19]
            _ = v25
            _ = v26
        }
        var v27 uint8 = 228
        if false {
            var v28 int16 = -32768
            for fi29 := 0; fi29 < 18; fi29++ {
                v30 := v3[v19]
                _ = fi29
                _ = v30
            }
            _ = v28
        } else {
            for fi31 := 0; fi31 < 19; fi31++ {
                if (!true) {
                    var v32 int64 = -705
                    _ = v32
                }
                if false {
                    v33 := GenS5{F0: (v2 / ((168) | 1))}
                    _ = v33
                }
                _ = fi31
            }
            v34 := GenS4{F0: 0}
            _ = v34
        }
        _ = v27
    }
    var v35 uint64 = 504
    var v36 uint16 = (v1 >> ((136) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v36
}
