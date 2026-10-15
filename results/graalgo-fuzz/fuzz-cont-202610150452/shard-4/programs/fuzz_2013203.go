package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int }
type GenS2 struct { F0 int8; F1 uint64; F2 uint8 }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]uint32)
    var v3 int32 = 2147483647
    var v4 uint8 = 1
    var v5 int64 = -9223372036854775808
    v6 := make(map[int]uint32)
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            switch "\n\t" {
            case "a":
                v9 := GenS1{F0: v5, F1: v1, F2: fi8}
                _ = v9
            }
            var v10 bool = false
            v5 = v5
            _ = fi8
            _ = v10
        }
        if false {
            v11 := GenS2{F0: -104, F1: v1, F2: 0}
            fmt.Println(v4)
            var v12 uint = 255
            _ = v11
            _ = v12
        }
        if true {
            var v13 int16 = 1
            _ = v13
        }
        _ = fi7
    }
    v1 = (v1 + 42)
    for fi14 := 0; fi14 < 15; fi14++ {
        switch (fi14 / ((127) | 1)) {
        case 4294967295:
            v1 = (v1 / (((^(v1))) | 1))
        default:
            fmt.Println(v3)
        }
        _ = fi14
    }
    if (-38 <= 52) {
        for fi15 := 0; fi15 < 16; fi15++ {
            if true {
                delete(v2, fi15)
                v1 -= v1
                var v16 int8 = -1
                _ = v16
            }
            switch ("foo bar" + "a") {
            case "":
                fmt.Println(fi15, v3)
            case "hello":
                if false {
                    switch fi15 {
                    case 4294967295:
                        switch fi15 {
                        case 128:
                            var v17 int = (fi15 * -1)
                            _ = v17
                        }
                    }
                    fmt.Println(fi15, v4)
                    v1 -= 18446744073709551615
                } else {
                    fmt.Println(fi15, v3)
                    switch fi15 {
                    case 33:
                        v2[399] = 0
                    case 295:
                        var v18 bool = (!true)
                        _ = v18
                    }
                    var v19 float64 = 1e10
                    _ = v19
                }
            case "foo bar":
                v20 := GenS2{F0: 42, F1: (^(v1)), F2: 255}
                _ = v20
            }
            _ = len(v2)
            _ = fi15
        }
        var v21 uint64 = (v1 ^ (v1 + v1))
        _ = v21
    }
    var v22 int8 = -1
    for fi23 := 0; fi23 < 10; fi23++ {
        fmt.Println(v3, v1)
        fmt.Println(v3, v3)
        _ = fi23
    }
    if ((true && false) && (false && false)) {
        for fi24 := 0; fi24 < 13; fi24++ {
            var v25 uint32 = 2
            _ = fi24
            _ = v25
        }
        v26 := []int{}
        _ = v26
    }
    fmt.Println(v1, v3, v4)
    fmt.Println(v1, v5)
    switch "" {
    case "\n\t":
        if false {
            delete(v2, 1)
            v5 += v5
        } else {
            v5 += int64(v3)
            delete(v6, 2147483647)
            var v27 int32 = v3
            _ = v27
        }
    default:
        v4 += v4
    }
    v28 := GenS1{F0: (v5 | (v5 * 0)), F1: v1, F2: -397}
    v1 = v1
    for fi29 := 0; fi29 < 4; fi29++ {
        v22 = (v22 + -1)
        v22 -= (v22 ^ -96)
        _ = fi29
    }
    var v30 uint32 = 297
    var v31 uint64 = (^(v1))
    fmt.Println(v4, v1)
    if (false && true) {
        fmt.Println(v1, v1)
    } else {
        v1 = (v31 - (v31 + v31))
        v32 := v6[-242]
        _ = v32
    }
    v30 -= (v30 << ((695) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
