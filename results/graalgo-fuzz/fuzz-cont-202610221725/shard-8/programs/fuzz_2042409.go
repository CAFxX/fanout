package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 int64; F1 uint8; F2 int64 }
type GenS5 struct { F0 int8; F1 uint64 }
type GenS6 struct { F0 int32; F1 uint; F2 int64 }
type GenS7 struct { F0 uint32 }
func genFn0(p0 uint8, p1 uint8) uint16 {
    if (255 <= 65535) {
        v6 := GenS1{F0: true}
        _ = v6
    }
    v7 := GenS2{F0: ("\n\t" + ""), F1: 18446744073709551615}
    p1 = (p1 ^ 42)
    switch -192 {
    case 2147483647:
        v8 := []int{0}
        _ = v8
    case 1:
        var v9 int = 42
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 527
}
func genFn1(p0 uint8) uint32 {
    p0 -= p0
    p0 += p0
    if true {
        if (p0 <= (p0 * (p0 / ((p0) | 1)))) {
            fmt.Println(p0)
            for fi10 := 0; fi10 < 13; fi10++ {
                var v11 uint32 = 65535
                v11 = v11
                _ = fi10
                _ = v11
            }
        } else {
            p0 -= (p0 + p0)
            fmt.Println(p0)
            p0 -= p0
        }
        for fi12 := 0; fi12 < 11; fi12++ {
            fmt.Println(fi12, fi12)
            fmt.Println(fi12, fi12)
            var v13 int64 = 9223372036854775807
            _ = fi12
            _ = v13
        }
        if false {
            p0 = (^(p0))
            switch "a" {
            case "a":
                if (9223372036854775807 != 0) {
                    var v14 uint64 = 811
                    v15 := GenS3{F0: -668, F1: 2147483648}
                    _ = v14
                    _ = v15
                }
            case "hello":
                p0 += (p0 >> ((42) & 63))
            case "":
                fmt.Println(p0)
            default:
                if (-1 >= -1) {
                    switch 4294967295 {
                    case -9223372036854775808:
                        p0 = p0
                    }
                } else {
                    var v16 int16 = 0
                    _ = v16
                }
            }
        } else {
            v17 := GenS4{F0: 42, F1: p0, F2: 0}
            v18 := []int{127, 128, 113}
            v18 = append(v18, int(p0))
            _ = v17
            _ = v18
        }
    } else {
        switch 2147483647 {
        case 42:
            v19 := GenS5{F0: -24, F1: uint64(p0)}
            _ = v19
        case 1:
            p0 = 1
        default:
            p0 = p0
        }
        var v20 uint = 0
        var v21 uint = (^(v20))
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        for fi23 := 0; fi23 < 1; fi23++ {
            if (0 < 928) {
                var v24 uint8 = (^(p0))
                fmt.Println(fi23)
                v24 = v24
                _ = v24
            } else {
                v25 := []int{}
                var v26 uint8 = p0
                var v27 int32 = 1
                _ = len(v25)
                var v28 string = "\n\t"
                _ = v25
                _ = v26
                _ = v27
                _ = v28
            }
            var v29 int32 = 42
            _ = fi23
            _ = v29
        }
        _ = fi22
    }
    _ = p0
    return 42
}

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{}
    v3 := []string{}
    var v4 int64 = 42
    v5 := []uint32{2, 2147483647, 980, 811}
    for fi30 := 0; fi30 < 10; fi30++ {
        switch fi30 {
        case 128:
            v31 := GenS6{F0: 0, F1: 42, F2: (-(v4))}
            _ = v31
        case -9223372036854775808:
            if (42 <= 128) {
                v32 := GenS4{F0: v4, F1: 2, F2: (v4 ^ -1)}
                for fi33 := 0; fi33 < 3; fi33++ {
                    fmt.Println(v4, fi30)
                    var v34 int32 = -1
                    _ = fi33
                    _ = v34
                }
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = v32
            }
        case 127:
            var v35 uint32 = 2147483648
            _ = v35
        default:
            v36 := v1[fi30]
            _ = v36
        }
        v37 := GenS3{F0: 512, F1: 4294967295}
        _ = fi30
        _ = v37
    }
    for fi38 := 0; fi38 < 16; fi38++ {
        var v39 int16 = 1
        v4 = 1
        v40 := v1[-9223372036854775808]
        _ = fi38
        _ = v39
        _ = v40
    }
    v4 = v4
    v4 -= 42
    var v41 uint32 = 1
    v4 = 42
    v42 := GenS7{F0: v41}
    v41 = (^(v41))
    v5 = append(v5, 176)
    v43 := GenS6{F0: 2147483647, F1: 181, F2: (v4 * (v4 & v4))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
