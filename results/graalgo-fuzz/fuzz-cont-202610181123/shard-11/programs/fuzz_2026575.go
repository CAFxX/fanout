package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[uint32]string)
    var v3 int64 = 1
    v4 := make(map[int]uint32)
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 uint8 = 254
        var v7 int8 = int8(v6)
        if true {
            if (2 <= 574) {
                v8 := GenS1{F0: v3, F1: (v7 & -128)}
                v9 := []int{-9223372036854775808, (-(fi5)), fi5}
                if false {
                    v6 -= 1
                    var v10 uint32 = 595
                    _ = v10
                }
                _ = v8
                _ = v9
            }
            for fi11 := 0; fi11 < 9; fi11++ {
                v7 = (v7 & -128)
                v12 := GenS2{F0: v6}
                delete(v4, (fi5 >> ((4294967295) & 63)))
                _ = fi11
                _ = v12
            }
            if false {
                v4[32768] = 4294967295
            } else {
                var v13 bool = (((18446744073709551615 <= 1) || false) && false)
                _ = v13
            }
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    var v14 int8 = 25
    fmt.Println(v3, v14, v1)
    fmt.Println(v14, v1)
    delete(v2, 915)
    v15 := GenS1{F0: 0, F1: 1}
    if (661 < 0) {
        for fi16 := 0; fi16 < 9; fi16++ {
            if (((v3 <= (v3 % ((v3) | 1))) || ((v3 << ((57) & 63)) != (-(v3)))) || (42 == 740)) {
                v4[fi16] = 4294967295
                if (731 >= 1) {
                    for fi17 := 0; fi17 < 14; fi17++ {
                        v14 -= v14
                        v18 := []int{0, (fi17 & fi17)}
                        var v19 int32 = -1
                        fmt.Println(v1, v1)
                        var v20 int64 = v3
                        _ = fi17
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    delete(v4, fi16)
                    v3 -= v3
                }
            }
            switch fi16 {
            case 42:
                fmt.Println(v14, v1)
            default:
                var v21 float32 = (v1 - v1)
                _ = v21
            }
            var v22 float32 = (v1 / v1)
            _ = fi16
            _ = v22
        }
        v23 := GenS2{F0: 0}
        _ = v23
    }
    v2[427] = ""
    if (v1 > 1e10) {
        for fi24 := 0; fi24 < 8; fi24++ {
            var v25 uint32 = 2147483647
            for fi26 := 0; fi26 < 9; fi26++ {
                delete(v2, (v25 * (v25 >> ((361) & 63))))
                v1 = v1
                _ = fi26
            }
            if false {
                var v27 float32 = 1.5
                v3 += v3
                _ = v27
            } else {
                v28 := []int{(fi24 >> ((255) & 63))}
                _ = v28
            }
            _ = fi24
            _ = v25
        }
        if false {
            var v29 int32 = 0
            delete(v2, 58)
            v1 = (v1 - v1)
            _ = v29
        }
        for fi30 := 0; fi30 < 8; fi30++ {
            var v31 int16 = 1000
            v32 := GenS1{F0: (v3 + (-(v3))), F1: v14}
            _ = fi30
            _ = v31
            _ = v32
        }
    } else {
        v33 := GenS1{F0: v3, F1: 0}
        var v34 uint = 295
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 8; fi35++ {
        for fi36 := 0; fi36 < 19; fi36++ {
            v37 := v2[2147483648]
            fmt.Println(fi36, fi35, v3)
            _ = fi36
            _ = v37
        }
        _ = fi35
    }
    v38 := []int{}
    var v39 bool = false
    if ((0 >= 1) && v39) {
        var v40 uint32 = 2
        var v41 bool = (true || (4294967295 < 42))
        fmt.Println(v1, v14)
        _ = v40
        _ = v41
    }
    v1 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
}
