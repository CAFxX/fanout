package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 bool, p1 uint8) int16 {
    if p0 {
        var v4 uint32 = 726
        _ = v4
    }
    fmt.Println(p0, p0)
    var v5 bool = p0
    var v6 float32 = float32(0.1)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 uint32, p1 int8, p2 int16) uint16 {
    var v7 int16 = genFn0(false, 1)
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 16; fi9++ {
            fmt.Println(p2, p1)
            v10 := GenS1{F0: false}
            var v11 uint = 0
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        if true {
            for fi13 := 0; fi13 < 11; fi13++ {
                p1 = (p1 << ((255) & 63))
                v14 := GenS1{F0: false}
                p1 += -1
                _ = fi13
                _ = v14
            }
            fmt.Println(v7, p2, fi12)
        }
        switch ("hello" + "foo bar") {
        case "":
            if ((-2.25 - 0.0) < (0.0 / 3.14159)) {
                var v15 uint = 1
                _ = v15
            } else {
                p0 += (p0 * (^(p0)))
                v16 := GenS2{F0: p1, F1: (p2 / ((int16(p1)) | 1)), F2: 0}
                v17 := []int{}
                _ = v16
                _ = v17
            }
        case "a":
            var v18 int8 = p1
            _ = v18
        default:
            v19 := GenS2{F0: p1, F1: p2, F2: 980}
            _ = v19
        }
        _ = fi12
    }
    v7 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 uint8 = 128
    var v2 uint = 2
    v3 := make(map[string]int)
    if false {
        for fi20 := 0; fi20 < 15; fi20++ {
            v1 += v1
            v2 -= 4294967295
            _ = fi20
        }
        if (65535 <= genFn1(391, -21, genFn0(false, 127))) {
            for fi21 := 0; fi21 < 19; fi21++ {
                fmt.Println(v2, fi21)
                v2 -= 65535
                _ = fi21
            }
            v1 = (v1 & (v1 % ((v1) | 1)))
            delete(v3, (("foo bar" + "a") + ("\n\t" + "")))
        }
        var v22 uint = v2
        _ = v22
    }
    if ((42 == 42) && false) {
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 int64 = -9223372036854775808
            _ = fi23
            _ = v24
        }
        var v25 float32 = (((-2.25 / 3.14159) * 0.0) + 0.0)
        v1 = v1
        _ = v25
    }
    delete(v3, "\n\t")
    v26 := []int{127, 278, 478}
    if true {
        v2 += 18446744073709551615
        var v27 uint = v2
        _ = v27
    }
    var v28 float32 = 1e10
    switch "a" {
    case "":
        fmt.Println(v2, v2)
    default:
        for fi29 := 0; fi29 < 13; fi29++ {
            var v30 uint32 = 65535
            for fi31 := 0; fi31 < 19; fi31++ {
                var v32 float32 = v28
                v33 := GenS3{F0: (v1 - (v1 / ((121) | 1)))}
                v28 -= float32(1e10)
                _ = fi31
                _ = v32
                _ = v33
            }
            _ = fi29
            _ = v30
        }
    }
    v2 -= (v2 - 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v3
}
