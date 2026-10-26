package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 int32; F2 uint64 }
type GenS6 struct { F0 int; F1 uint64; F2 uint32 }
func genFn0(p0 uint16, p1 bool) uint {
    v6 := []int{32768}
    p0 += (p0 / ((2) | 1))
    var v7 uint = 255
    var v8 uint = v7
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return (v7 >> ((v8) & 63))
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = 2147483647
    var v3 uint64 = 464
    var v4 int64 = -9223372036854775808
    var v5 float32 = float32(1.5)
    v9 := GenS1{F0: (v4 & 0), F1: 2, F2: 383}
    for fi10 := 0; fi10 < 1; fi10++ {
        fmt.Println(fi10)
        _ = fi10
    }
    var v11 int16 = -32768
    v12 := GenS2{F0: 127}
    var v13 int = -331
    if false {
        delete(v1, 2147483647)
        switch (v13 + v13) {
        case 128:
            fmt.Println(v3, v4, v13)
        case 32768:
            var v14 uint32 = 4294967295
            _ = v14
        case 2147483647:
            fmt.Println(v13)
        default:
            v11 -= (v11 - v11)
        }
        for fi15 := 0; fi15 < 18; fi15++ {
            v4 += -335
            _ = fi15
        }
    } else {
        var v16 uint16 = 255
        _ = v16
    }
    if false {
        v17 := []int{v13, (v13 / ((746) | 1)), (-(v13)), (v13 % (((v13 + (-(v13)))) | 1))}
        v5 += v5
        switch ("foo bar" + "hello") {
        case "foo bar":
            for fi18 := 0; fi18 < 11; fi18++ {
                v2 = -150
                var v19 uint = genFn0(65535, true)
                _ = fi18
                _ = v19
            }
        case "hello":
            var v20 uint = 0
            _ = v20
        case "a":
            v13 = (v13 * (v13 - v13))
        }
        _ = v17
    } else {
        var v21 bool = true
        _ = v21
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        fmt.Println(v4, v2)
        for fi23 := 0; fi23 < 10; fi23++ {
            fmt.Println(v2)
            v24 := GenS3{F0: (fi22 / ((v13) | 1))}
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    v11 = (v11 >> ((357) & 63))
    for fi25 := 0; fi25 < 13; fi25++ {
        switch (fi25 ^ (fi25 | v13)) {
        case 128:
            v3 = v3
        case 835:
            var v26 int32 = v2
            _ = v26
        case 667:
            var v27 uint8 = 255
            _ = v27
        default:
            v28 := GenS4{F0: 4294967295}
            _ = v28
        }
        v4 -= (-(v4))
        _ = fi25
    }
    v11 = 1000
    if (("hello" + "a") > "foo bar") {
        if (!false) {
            var v29 int8 = 31
            if true {
                var v30 float32 = 1e10
                _ = v30
            }
            _ = v29
        }
        v5 += v5
        v3 = 0
    } else {
        var v31 string = "\n\t"
        _ = v31
    }
    switch "" {
    case "\n\t":
        v13 -= v13
    case "a":
        var v32 uint32 = 628
        _ = v32
    case "foo bar":
        v33 := GenS5{F0: -128, F1: v2, F2: v3}
        _ = v33
    }
    if ((^(v2)) == (v2 << ((65535) & 63))) {
        for fi34 := 0; fi34 < 3; fi34++ {
            fmt.Println(v13, v5, v5)
            if (v2 == v2) {
                v35 := GenS6{F0: fi34, F1: 336, F2: 1}
                _ = v35
            }
            _ = fi34
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
