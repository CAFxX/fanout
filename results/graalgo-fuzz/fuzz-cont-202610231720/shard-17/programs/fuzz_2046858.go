package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int8; F2 int32 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 int64; F1 int16; F2 int }
func genFn0(p0 uint64, p1 int, p2 int) int8 {
    v4 := []int{2147483647, p1}
    switch -1 {
    case 32767:
        v5 := GenS1{F0: 255}
        _ = v5
    }
    var v6 uint = 1
    var v7 bool = true
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    _ = v7
    return 100
}
func genFn1(p0 uint16) int64 {
    p0 += p0
    for fi8 := 0; fi8 < 10; fi8++ {
        if true {
            switch 29 {
            case 127:
                v9 := []int{(fi8 << ((65535) & 63))}
                _ = v9
            }
            p0 += 765
        }
        p0 += (p0 * p0)
        var v10 int64 = 467
        _ = fi8
        _ = v10
    }
    v11 := []int{32768}
    _ = p0
    _ = v11
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = ((271 != -1) && false)
    v3 := make(map[int]bool)
    if (!((true || ((854 >= 1) || (403 >= 65535))) && v2)) {
        v12 := []int{42, -2147483648, 1}
        var v13 float32 = (-2.25 - 0.0)
        var v14 uint = 4294967295
        _ = v12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        for fi16 := 0; fi16 < 10; fi16++ {
            var v17 int = 32767
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v18 := []int{4294967295}
    if v2 {
        if true {
            v3[534] = false
            v2 = (423 >= 2)
            if (391 >= 65535) {
                v19 := GenS2{F0: 18446744073709551615, F1: 127, F2: 0}
                v2 = (429 <= -2147483648)
                _ = v19
            } else {
                var v20 int8 = -1
                v20 -= genFn0(0, 862, -208)
                var v21 float64 = 1e10
                _ = v20
                _ = v21
            }
        } else {
            v22 := GenS3{F0: 42, F1: 549}
            _ = v22
        }
    } else {
        var v23 uint = 65535
        for fi24 := 0; fi24 < 16; fi24++ {
            if v2 {
                var v25 uint = 240
                for fi26 := 0; fi26 < 13; fi26++ {
                    fmt.Println(fi24, fi24, v23)
                    _ = fi26
                }
                _ = v25
            } else {
                fmt.Println(v2, v23)
            }
            _ = fi24
        }
        _ = v23
    }
    var v27 float64 = 1e10
    var v28 int64 = genFn1(2)
    var v29 int8 = -1
    v30 := GenS4{F0: (v28 & v28), F1: -1, F2: -9223372036854775808}
    fmt.Println(v2, v27, v29)
    v31 := GenS4{F0: v28, F1: 32767, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
