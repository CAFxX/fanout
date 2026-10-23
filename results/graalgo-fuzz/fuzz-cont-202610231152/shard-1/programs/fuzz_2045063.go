package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 uint }

func main() {
    v1 := make(map[int]string)
    var v2 bool = true
    var v3 uint64 = 4294967295
    var v4 int16 = 0
    var v5 uint32 = 103
    v6 := make(map[int]bool)
    if false {
        if v2 {
            _ = len(v1)
            var v7 float32 = 1e10
            _ = v7
        } else {
            var v8 uint32 = 351
            _ = v8
        }
        v4 = (v4 << ((992) & 63))
    }
    switch ("hello" + ("hello" + "hello")) {
    case "\n\t":
        v9 := GenS1{F0: v5}
        _ = v9
    case "a":
        fmt.Println(v2, v5)
    }
    v3 = (v3 * (v3 & (v3 << ((255) & 63))))
    if (!(127 > -108)) {
        v10 := GenS1{F0: (v5 | (^(v5)))}
        for fi11 := 0; fi11 < 12; fi11++ {
            for fi12 := 0; fi12 < 9; fi12++ {
                v3 += v3
                v13 := GenS2{F0: -1}
                if (-2.25 != float64(((0.0 / 3.14159) / 1.0))) {
                    var v14 int16 = 1000
                    _ = v14
                } else {
                    var v15 float64 = (0.0 * 3.14159)
                    var v16 int8 = -1
                    v2 = v2
                    _ = v15
                    _ = v16
                }
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
        _ = v10
    } else {
        v17 := GenS3{F0: 128, F1: 18446744073709551615}
        v1[2147483647] = ("hello" + (("a" + "hello") + ("hello" + "a")))
        _ = v17
    }
    v18 := []int{32768}
    _ = len(v18)
    v19 := GenS1{F0: (^(v5))}
    v20 := v1[-843]
    v21 := GenS3{F0: 89, F1: 0}
    if (v2 && v2) {
        _ = len(v1)
        _ = len(v18)
        var v22 int8 = -1
        _ = v22
    }
    v20 = v20
    var v23 uint8 = 1
    if v2 {
        v5 = (^(v5))
        var v24 int64 = -614
        _ = v24
    }
    var v25 uint16 = 32767
    if (992 <= 9223372036854775807) {
        for fi26 := 0; fi26 < 4; fi26++ {
            v27 := GenS3{F0: fi26, F1: 18446744073709551615}
            _ = fi26
            _ = v27
        }
        if (1 > 0) {
            v28 := GenS2{F0: -62}
            v4 += 963
            _ = v28
        }
        switch "foo bar" {
        case "hello":
            var v29 string = "a"
            _ = v29
        case "":
            var v30 int8 = 1
            _ = v30
        case "\n\t":
            for fi31 := 0; fi31 < 1; fi31++ {
                _ = len(v18)
                _ = fi31
            }
        }
    }
    v25 = (v25 << ((582) & 63))
    var v32 int32 = int32(v25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v25
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
