package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint64; F1 int32; F2 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16 }
func genFn0(p0 int16, p1 uint16) int {
    var v4 uint32 = 1
    v5 := []int{-313, -1}
    p0 += -476
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 uint8 = 1
    var v2 int16 = -451
    v3 := make(map[int]bool)
    v6 := GenS1{F0: -9223372036854775808, F1: "", F2: 0}
    if (0 <= 1) {
        var v7 int32 = 42
        delete(v3, genFn0((v2 / (((v2 + v2)) | 1)), 32767))
        switch 127 {
        case 32768:
            var v8 uint64 = 237
            _ = v8
        default:
            var v9 int32 = v7
            _ = v9
        }
        _ = v7
    }
    v10 := GenS2{F0: 65535, F1: 166}
    if (540 < 22) {
        v1 = v1
        for fi11 := 0; fi11 < 10; fi11++ {
            for fi12 := 0; fi12 < 18; fi12++ {
                var v13 string = "\n\t"
                delete(v3, (-(fi12)))
                _ = fi12
                _ = v13
            }
            v2 = (v2 - (^(v2)))
            switch fi11 {
            case -9223372036854775808:
                v14 := GenS3{F0: 2, F1: 2147483647, F2: 0}
                _ = v14
            case 433:
                var v15 float32 = 1e10
                _ = v15
            case -103:
                _ = len(v3)
            }
            _ = fi11
        }
        var v16 string = "a"
        _ = v16
    } else {
        v17 := []int{127, genFn0(-32768, 450), -1, 32767}
        _ = v17
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 int64 = -9223372036854775808
        v1 -= (^(v1))
        _ = fi18
        _ = v19
    }
    if (!(42 != 1)) {
        v2 = (v2 & v2)
        _ = len(v3)
    } else {
        v1 += v1
        var v20 uint64 = 435
        _ = v20
    }
    if false {
        v2 += (v2 * (v2 << ((4294967295) & 63)))
        v21 := GenS2{F0: 2147483647, F1: (v1 | 174)}
        _ = v21
    } else {
        v2 = 0
        v2 += int16(v1)
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v1)
        case "hello":
            for fi22 := 0; fi22 < 20; fi22++ {
                for fi23 := 0; fi23 < 13; fi23++ {
                    if (0 > 13) {
                        var v24 uint64 = uint64(v1)
                        _ = v24
                    }
                    _ = fi23
                }
                v2 -= v2
                _ = fi22
            }
        }
    }
    if (-1 != 1) {
        var v25 uint32 = 42
        v25 = (v25 ^ (^(v25)))
        v26 := GenS4{F0: 42}
        _ = v25
        _ = v26
    } else {
        if (false && false) {
            var v27 int8 = 72
            _ = v27
        }
        var v28 int32 = 42
        _ = v28
    }
    for fi29 := 0; fi29 < 15; fi29++ {
        var v30 uint64 = 18446744073709551615
        var v31 uint16 = 1
        _ = fi29
        _ = v30
        _ = v31
    }
    v32 := v3[-371]
    var v33 uint8 = (v1 - (v1 % (((v1 - (v1 & (v1 ^ v1)))) | 1)))
    fmt.Println(v1, v33)
    fmt.Println(v1)
    switch ("\n\t" + ("" + "a")) {
    case "hello":
        v34 := GenS5{F0: 255}
        _ = v34
    case "a":
        fmt.Println(v33, v32)
    case "":
        var v35 int8 = 107
        _ = v35
    default:
        v32 = v32
    }
    switch -9223372036854775808 {
    case 790:
        for fi36 := 0; fi36 < 15; fi36++ {
            v2 = (-(v2))
            fmt.Println(v33, v32, fi36)
            _ = fi36
        }
    case 0:
        for fi37 := 0; fi37 < 15; fi37++ {
            v38 := []int{fi37}
            _ = fi37
            _ = v38
        }
    case 42:
        if v32 {
            _ = len(v3)
        } else {
            var v39 uint8 = 128
            var v40 int32 = -480
            _ = v39
            _ = v40
        }
    }
    if (742 <= 730) {
        fmt.Println(v1, v33)
    }
    v41 := GenS2{F0: uint32(v33), F1: v33}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v41
    _ = v6
}
