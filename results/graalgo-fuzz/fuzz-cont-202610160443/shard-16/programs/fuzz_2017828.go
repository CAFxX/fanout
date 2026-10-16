package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int; F2 uint64 }
type GenS4 struct { F0 int16; F1 uint32; F2 uint64 }
func genFn0(p0 uint32, p1 int8) uint64 {
    v6 := []int{9223372036854775807, 32767}
    var v7 string = ("hello" + "a")
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 uint64, p1 int16, p2 uint64) uint64 {
    v8 := GenS1{F0: 9223372036854775807, F1: 32767, F2: -1}
    if (false && (1 < 218)) {
        var v9 uint32 = 65535
        for fi10 := 0; fi10 < 6; fi10++ {
            switch (fi10 % ((-9223372036854775808) | 1)) {
            case 32767:
                fmt.Println(p1, fi10, p2)
            case 719:
                v11 := GenS1{F0: 1, F1: 65535, F2: -114}
                _ = v11
            case 127:
                if (-1 >= 9223372036854775807) {
                    v12 := []int{fi10, fi10, fi10, fi10}
                    p2 -= (p0 * (p0 * (p2 - genFn0(677, 41))))
                    if ((p0 << ((42) & 63)) <= (p2 / (((p0 * (p2 - p2))) | 1))) {
                        v13 := GenS1{F0: -9223372036854775808, F1: 304, F2: 127}
                        p2 += p2
                        _ = v13
                    } else {
                        var v14 uint8 = 84
                        v14 += 42
                        _ = len(v12)
                        _ = v14
                    }
                    _ = v12
                }
            default:
                var v15 int = (fi10 - (fi10 * fi10))
                _ = v15
            }
            _ = fi10
        }
        _ = v9
    } else {
        p1 = p1
        var v16 uint16 = 226
        _ = v16
    }
    v17 := GenS2{F0: 127}
    v18 := GenS2{F0: 220}
    if (0.0 >= 1e10) {
        v19 := GenS3{F0: 2, F1: -2147483648, F2: genFn0(2147483647, 22)}
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 uint64 = p2
            _ = fi20
            _ = v21
        }
        switch "\n\t" {
        case "a":
            fmt.Println(p1, p0, p1)
        case "":
            var v22 int8 = int8(p2)
            _ = v22
        case "\n\t":
            v23 := []int{-1}
            _ = v23
        default:
            p2 -= (p0 * genFn0(65535, -128))
        }
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    _ = v8
    return (p2 % ((p0) | 1))
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    var v3 uint8 = 91
    var v4 uint16 = 32767
    var v5 int64 = -95
    if false {
        fmt.Println(v4, v4, v4)
        v24 := v1[2]
        _ = v24
    }
    fmt.Println(v4, v4)
    for fi25 := 0; fi25 < 17; fi25++ {
        v3 -= 2
        fmt.Println(fi25, fi25)
        var v26 int32 = 1
        _ = fi25
        _ = v26
    }
    v27 := GenS4{F0: 42, F1: 2147483648, F2: 4294967295}
    var v28 string = "hello"
    var v29 int = 2147483647
    v29 -= (-(v29))
    v5 = -1
    fmt.Println(v3)
    switch "hello" {
    case "\n\t":
        if true {
            v3 -= (v3 * v3)
            if false {
                v2["\n\t"] = v29
                var v30 int32 = 42
                _ = v30
            }
        }
    case "a":
        v31 := v2["\n\t"]
        _ = v31
    default:
        for fi32 := 0; fi32 < 6; fi32++ {
            v3 += 2
            _ = fi32
        }
    }
    v29 = -313
    if (v3 != v3) {
        _ = len(v1)
    } else {
        v4 += uint16(v3)
        var v33 string = v28
        fmt.Println(v5)
        _ = v33
    }
    var v34 uint64 = 458
    v5 = (v5 & (v5 & v5))
    for fi35 := 0; fi35 < 17; fi35++ {
        v36 := []int{(fi35 >> ((18446744073709551615) & 63)), fi35}
        var v37 float64 = (1.5 - float64(((0.0 / (-2.25 - 1e10)) + 0.1)))
        var v38 int64 = (v5 % (((v5 * -1)) | 1))
        _ = fi35
        _ = v36
        _ = v37
        _ = v38
    }
    v29 -= (v29 + 187)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
