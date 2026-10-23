package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int32; F1 int32; F2 uint }
type GenS4 struct { F0 uint64; F1 uint64; F2 uint8 }
func genFn0(p0 int16, p1 int64, p2 int32) int {
    var v6 int8 = 42
    p1 = (p1 + (p1 / ((543) | 1)))
    var v7 uint32 = 1
    switch "a" {
    case "":
        p2 -= (p2 | 2147483647)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -662
}
func genFn1(p0 string, p1 uint64, p2 uint32) uint {
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 16; fi9++ {
            if false {
                v10 := GenS1{F0: false, F1: -1}
                var v11 string = ("hello" + p0)
                _ = v10
                _ = v11
            } else {
                if (!true) {
                    v12 := []int{}
                    _ = v12
                }
                p0 = p0
                fmt.Println(fi8)
            }
            _ = fi9
        }
        for fi13 := 0; fi13 < 11; fi13++ {
            if (65535 < 136) {
                if true {
                    v14 := GenS1{F0: false, F1: 100}
                    v15 := GenS1{F0: true, F1: 0}
                    _ = v14
                    _ = v15
                }
            }
            p1 += p1
            if (!true) {
                v16 := []int{int(p1), fi8, genFn0(1000, 0, -2147483648)}
                p1 += (p1 | 4294967295)
                v16 = append(v16, (fi8 + (fi8 << ((18446744073709551615) & 63))))
                _ = v16
            } else {
                var v17 int = -1
                v18 := GenS2{F0: -195, F1: 32767}
                _ = v17
                _ = v18
            }
            _ = fi13
        }
        _ = fi8
    }
    p0 = (("foo bar" + (((p0 + (p0 + p0)) + p0) + ((p0 + p0) + p0))) + p0)
    v19 := []int{genFn0(0, 42, 635)}
    if len(v19) > 0 {
        v19[0] = -1
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 float64 = -2.25
        var v22 uint = 712
        _ = fi20
        _ = v21
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 65535
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = (996 > 2147483648)
    var v3 float64 = 1e10
    v4 := make(map[int]int)
    v5 := make(map[uint32]int)
    fmt.Println(v3, v2)
    v23 := []int{-2147483648}
    for fi24 := 0; fi24 < 19; fi24++ {
        var v25 bool = v2
        if ((((-499 <= 195) || (v2 || false)) && true) && true) {
            var v26 float32 = 0.0
            for fi27 := 0; fi27 < 20; fi27++ {
                v28 := GenS1{F0: (v25 && v2), F1: 44}
                v26 = 1e10
                fmt.Println(v3, v25, fi24)
                _ = fi27
                _ = v28
            }
            _ = v26
        } else {
            if true {
                fmt.Println(v25, fi24)
            }
        }
        switch (fi24 * (fi24 % ((genFn0(694, 42, -708)) | 1))) {
        case -9223372036854775808:
            v29 := GenS3{F0: 1, F1: 1, F2: 65535}
            _ = v29
        case 227:
            var v30 int32 = 42
            _ = v30
        case -2147483648:
            if ("" != "\n\t") {
                v23 = append(v23, -9223372036854775808)
                var v31 int8 = 1
                _ = v31
            } else {
                v32 := v5[42]
                v33 := GenS2{F0: -1, F1: 390}
                _ = v32
                _ = v33
            }
        default:
            switch (fi24 & (fi24 / ((32768) | 1))) {
            case 858:
                var v34 uint16 = 162
                _ = v34
            case 319:
                var v35 uint8 = 0
                _ = v35
            case 127:
                v36 := v4[(fi24 | fi24)]
                _ = v36
            }
        }
        _ = fi24
        _ = v25
    }
    if v2 {
        var v37 int = genFn0(32767, 42, 2147483647)
        _ = v37
    } else {
        _ = len(v23)
        v3 = (1.5 - v3)
    }
    var v38 uint8 = 42
    _ = len(v1)
    var v39 int64 = 9223372036854775807
    v38 += v38
    var v40 bool = true
    for fi41 := 0; fi41 < 18; fi41++ {
        var v42 bool = (!v40)
        var v43 int8 = 1
        v38 = (v38 >> ((1) & 63))
        _ = fi41
        _ = v42
        _ = v43
    }
    v44 := GenS2{F0: -1, F1: 0}
    var v45 bool = (v2 && (517 > 42))
    v46 := GenS4{F0: 0, F1: 920, F2: (v38 << ((88) & 63))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
