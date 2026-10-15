package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint32) int8 {
    if false {
        if true {
            fmt.Println(p0)
        }
        var v4 bool = false
        var v5 int64 = 0
        _ = v4
        _ = v5
    } else {
        v6 := GenS1{F0: 65535, F1: 381, F2: 42}
        p0 = p0
        _ = v6
    }
    var v7 int16 = 42
    _ = p0
    _ = v7
    return 102
}
func genFn1(p0 int16) uint64 {
    fmt.Println(p0)
    switch 503 {
    case -2147483648:
        v8 := GenS1{F0: 42, F1: 18446744073709551615, F2: 0}
        _ = v8
    case -882:
        if (!(2147483647 != -1)) {
            p0 += p0
            p0 -= (p0 << ((255) & 63))
        }
    default:
        if (true && (42 < 42)) {
            for fi9 := 0; fi9 < 6; fi9++ {
                for fi10 := 0; fi10 < 16; fi10++ {
                    switch 2147483647 {
                    case 111:
                        p0 += 954
                    default:
                        p0 -= 0
                        v11 := GenS1{F0: 4294967295, F1: 623, F2: 32767}
                        _ = v11
                    }
                    _ = fi10
                }
                v12 := []int{-66, fi9, 2147483647}
                var v13 uint8 = 0
                _ = fi9
                _ = v12
                _ = v13
            }
        }
    }
    var v14 int32 = 0
    p0 = 155
    for fi15 := 0; fi15 < 11; fi15++ {
        v16 := []int{}
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = v14
    return 331
}

func main() {
    var v1 int16 = 106
    v2 := []uint32{42}
    v3 := make(map[int]uint32)
    var v17 int16 = 883
    if ((false || (!true)) || false) {
        var v18 uint8 = 128
        for fi19 := 0; fi19 < 7; fi19++ {
            v1 += v1
            _ = fi19
        }
        v1 = (-(v1))
        _ = v18
    } else {
        v1 = 1000
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        v17 -= 288
        if (42 == 0) {
            v2 = append(v2, 2147483647)
            fmt.Println(v1)
        } else {
            var v21 uint16 = 32767
            _ = v21
        }
        v22 := GenS1{F0: 0, F1: 4294967295, F2: (fi20 - (fi20 / ((fi20) | 1)))}
        _ = fi20
        _ = v22
    }
    switch "" {
    case "hello":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "a":
        for fi23 := 0; fi23 < 6; fi23++ {
            var v24 int = (fi23 % ((9223372036854775807) | 1))
            switch ("hello" + ("hello" + (("" + "a") + "a"))) {
            case "hello":
                v25 := GenS2{F0: (v17 >> ((2) & 63))}
                _ = v25
            case "":
                v1 = 1
            default:
                var v26 int8 = 42
                _ = v26
            }
            var v27 uint = 0
            _ = fi23
            _ = v24
            _ = v27
        }
    }
    v28 := GenS2{F0: (-(v17))}
    v29 := GenS1{F0: 255, F1: 1, F2: 128}
    v30 := v3[-2147483648]
    if true {
        switch 32768 {
        case 32767:
            var v31 float64 = ((float64(0.1) / 0.1) + ((0.0 - 1.5) + 0.0))
            _ = v31
        case 42:
            v17 -= (v1 / (((v17 >> ((2) & 63))) | 1))
        case 349:
            var v32 float64 = 3.14159
            _ = v32
        }
        var v33 uint16 = 0
        _ = v33
    } else {
        v34 := GenS3{F0: v1}
        v3[-277] = 504
        _ = v34
    }
    for fi35 := 0; fi35 < 2; fi35++ {
        for fi36 := 0; fi36 < 6; fi36++ {
            if true {
                v1 = v1
                fmt.Println(fi36, v30)
            }
            _ = fi36
        }
        if ((3.14159 != (1e10 * 1e10)) || true) {
            var v37 int64 = -1
            _ = v37
        }
        _ = len(v3)
        _ = fi35
    }
    v2 = append(v2, (v30 ^ 4294967295))
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v17
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
