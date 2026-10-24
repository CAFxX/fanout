package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint16 }
func genFn0(p0 int8, p1 int8) uint16 {
    v6 := []int{1}
    var v7 int = 4294967295
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 982
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 2
    var v3 int8 = 70
    v4 := []int{32768}
    v5 := make(map[uint32]string)
    for fi8 := 0; fi8 < 20; fi8++ {
        switch "foo bar" {
        case "foo bar":
            delete(v5, 4294967295)
        default:
            v9 := GenS1{F0: 406, F1: -1}
            _ = v9
        }
        if len(v4) > 0 {
            v4[0] = (fi8 ^ fi8)
        }
        _ = fi8
    }
    if ((false || true) || (18446744073709551615 <= 255)) {
        v10 := GenS2{F0: 18446744073709551615, F1: 127, F2: 2}
        _ = v10
    } else {
        v11 := GenS2{F0: 2, F1: 42, F2: 2}
        switch 2147483647 {
        case 32768:
            var v12 uint16 = genFn0(v3, (v3 & (v3 << ((255) & 63))))
            _ = v12
        }
        v3 = int8(v2)
        _ = v11
    }
    v2 = v2
    var v13 float32 = -2.25
    for fi14 := 0; fi14 < 1; fi14++ {
        v3 = -49
        _ = fi14
    }
    if (992 < 42) {
        if (false && false) {
            delete(v5, 65535)
            if true {
                v13 = v13
            }
        }
    }
    if true {
        fmt.Println(v13)
        var v15 uint8 = 2
        fmt.Println(v13, v2)
        _ = v15
    }
    v2 += 255
    if (2147483647 < 813) {
        switch 9223372036854775807 {
        case 2147483647:
            var v16 int64 = -9223372036854775808
            _ = v16
        case 1:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case -9223372036854775808:
            if ("a" >= "foo bar") {
                var v17 uint16 = 65535
                fmt.Println(v13)
                _ = v17
            } else {
                v2 = 206
            }
        }
        v18 := GenS1{F0: -723, F1: 42}
        _ = v18
    }
    var v19 int32 = 1
    for fi20 := 0; fi20 < 15; fi20++ {
        fmt.Println(v19, v3)
        for fi21 := 0; fi21 < 10; fi21++ {
            v22 := GenS2{F0: 502, F1: v2, F2: 42}
            switch (fi21 ^ fi20) {
            case -406:
                var v23 int64 = 1
                _ = v23
            case 4294967295:
                v24 := v1[(fi21 | -1)]
                _ = v24
            }
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    var v25 uint64 = 134
    var v26 int8 = -127
    fmt.Println(v2, v13, v25)
    var v27 int16 = 786
    for fi28 := 0; fi28 < 2; fi28++ {
        v4 = append(v4, (fi28 - (fi28 & 42)))
        var v29 bool = true
        _ = fi28
        _ = v29
    }
    v30 := GenS2{F0: v25, F1: 2, F2: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
