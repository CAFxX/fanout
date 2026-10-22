package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint16; F1 string }
func genFn0(p0 uint16, p1 uint16, p2 int8) bool {
    v4 := []int{9223372036854775807, 32768}
    if true {
        v5 := GenS1{F0: 1}
        v6 := GenS1{F0: -902}
        if false {
            var v7 int64 = 920
            if (1 != 255) {
                fmt.Println(v7)
            }
            p0 = 255
            _ = v7
        }
        _ = v5
        _ = v6
    }
    if false {
        v8 := GenS1{F0: 0}
        var v9 int64 = -1
        _ = v8
        _ = v9
    } else {
        if (!false) {
            var v10 int32 = -189
            if (false || (((((-(v10)) < v10) || (127 <= 32768)) && false) || true)) {
                v10 = 0
                var v11 int = 4294967295
                _ = v11
            }
            fmt.Println(p2, p0, p0)
            _ = v10
        }
    }
    var v12 int64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v4
    return false
}
func genFn1(p0 int, p1 uint) uint32 {
    var v13 float64 = float64(1.5)
    v13 = 1.5
    _ = p0
    _ = p1
    _ = v13
    return 1
}

func main() {
    var v1 int8 = 127
    v2 := make(map[uint32]int)
    var v3 int64 = -69
    v1 = (v1 >> ((18446744073709551615) & 63))
    switch 32768 {
    case 4294967295:
        v14 := GenS2{F0: 42, F1: 255, F2: 0}
        _ = v14
    }
    _ = len(v2)
    for fi15 := 0; fi15 < 2; fi15++ {
        for fi16 := 0; fi16 < 14; fi16++ {
            v17 := GenS3{F0: 651, F1: "foo bar"}
            _ = fi16
            _ = v17
        }
        v3 = 239
        _ = fi15
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        v19 := GenS3{F0: 42, F1: ""}
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        v1 -= v1
        _ = fi20
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        var v22 bool = true
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: int32(v1)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
}
