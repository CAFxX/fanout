package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int32 }
type GenS2 struct { F0 int8; F1 bool }

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{896}
    v3 := make(map[string]string)
    var v4 uint16 = 2
    if false {
        var v5 float64 = 1.5
        if (0 != int8(v4)) {
            var v6 int64 = 1
            if len(v2) > 0 {
                v2[0] = 2147483648
            }
            var v7 int8 = 1
            _ = v6
            _ = v7
        } else {
            var v8 int = -2147483648
            v9 := GenS1{F0: 2147483647, F1: "", F2: 42}
            v10 := GenS2{F0: 100, F1: ((!true) || (!false))}
            _ = v10
            _ = v8
            _ = v9
        }
        var v11 float32 = (0.0 - 1.5)
        _ = v11
        _ = v5
    } else {
        var v12 float64 = (float64(-2.25) - -2.25)
        _ = v12
    }
    var v13 uint32 = 983
    _ = len(v2)
    var v14 uint64 = 462
    v4 = (v4 & (v4 | v4))
    v4 += v4
    var v15 uint8 = 55
    v14 += (v14 | (v14 & (v14 << ((507) & 63))))
    fmt.Println(v13)
    var v16 int64 = 0
    if (v13 >= uint32(v14)) {
        fmt.Println(v15, v14)
    } else {
        if false {
            if (!true) {
                v4 = v4
                v4 -= v4
                v15 -= v15
            }
        }
        v17 := GenS2{F0: 127, F1: false}
        var v18 int8 = -1
        _ = v17
        _ = v18
    }
    if true {
        v15 = v15
        if (!false) {
            if true {
                v14 += v14
                var v19 string = "a"
                v4 = (v4 ^ v4)
                _ = v19
            } else {
                var v20 int16 = -32768
                _ = v20
            }
            if true {
                for fi21 := 0; fi21 < 13; fi21++ {
                    v15 = (^(v15))
                    _ = fi21
                }
                fmt.Println(v16)
            }
            for fi22 := 0; fi22 < 19; fi22++ {
                v23 := GenS1{F0: (v13 + uint32(v14)), F1: ("\n\t" + ""), F2: 2147483647}
                v4 = v4
                _ = len(v2)
                _ = fi22
                _ = v23
            }
        } else {
            v13 -= uint32(v4)
        }
    }
    if true {
        var v24 int = 1
        _ = len(v3)
        if (9223372036854775807 != (v16 >> ((18446744073709551615) & 63))) {
            var v25 int64 = 0
            _ = v25
        }
        _ = v24
    } else {
        var v26 int16 = -1
        _ = v26
    }
    var v27 uint16 = (v4 ^ 452)
    for fi28 := 0; fi28 < 15; fi28++ {
        switch "foo bar" {
        case "hello":
            v15 += (v15 << ((768) & 63))
        }
        v4 += v4
        _ = fi28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
