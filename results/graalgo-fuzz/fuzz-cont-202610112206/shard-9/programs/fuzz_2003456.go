package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint; F1 uint16; F2 uint16 }
func genFn0() string {
    v4 := []int{0}
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := GenS1{F0: 255}
        _ = fi5
        _ = v6
    }
    var v7 string = "a"
    _ = v4
    _ = v7
    return v7
}
func genFn1(p0 int32, p1 int, p2 uint8) uint8 {
    if (65535 <= 65535) {
        var v8 uint32 = 42
        p0 -= p0
        v9 := GenS1{F0: 42}
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 string = "hello"
            v12 := GenS2{F0: 184, F1: p2}
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    var v13 uint32 = 65535
    v14 := []int{p1, (-(p1)), p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 179
}

func main() {
    var v1 int8 = -37
    v2 := make(map[string]uint32)
    var v3 uint32 = 366
    for fi15 := 0; fi15 < 14; fi15++ {
        fmt.Println(fi15, v1)
        switch (fi15 ^ (^(fi15))) {
        case -2147483648:
            if (!true) {
                v3 += 2
                switch fi15 {
                case 9223372036854775807:
                    var v16 uint8 = 42
                    _ = v16
                case 683:
                    switch "foo bar" {
                    case "foo bar":
                        var v17 int32 = -161
                        _ = v17
                    case "\n\t":
                        v1 += 100
                    }
                }
            }
        case 32767:
            var v18 int8 = v1
            _ = v18
        }
        v19 := GenS3{F0: 0, F1: 471, F2: 2}
        _ = fi15
        _ = v19
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 uint64 = 42
        v1 -= (v1 ^ v1)
        v1 -= (v1 + (^(v1)))
        _ = fi20
        _ = v21
    }
    v2["\n\t"] = v3
    v22 := []int{}
    for fi23 := 0; fi23 < 5; fi23++ {
        v1 -= (v1 * 0)
        var v24 uint8 = genFn1(-1, fi23, genFn1(42, (fi23 / (((fi23 * fi23)) | 1)), 42))
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        v1 = v1
        v1 = v1
        _ = fi25
    }
    var v26 int64 = 9223372036854775807
    if len(v22) > 0 {
        v22[0] = 4294967295
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        fmt.Println(v1, v26)
        fmt.Println(fi27, fi27, v3)
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v3
}
