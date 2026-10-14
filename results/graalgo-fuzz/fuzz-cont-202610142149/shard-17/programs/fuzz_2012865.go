package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 uint64; F1 uint16; F2 int16 }
func genFn0(p0 int16) uint64 {
    if (!false) {
        v4 := []int{1, -663, -96, -1}
        _ = v4
    } else {
        v5 := []int{}
        p0 = 32767
        var v6 float64 = (float64(1e10) * 1e10)
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 255}
    _ = p0
    _ = v7
    return 500
}
func genFn1(p0 int64, p1 int32) uint32 {
    p0 -= p0
    v8 := GenS2{F0: p1}
    var v9 uint16 = 255
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int = 42
    var v2 uint64 = 0
    var v3 int = v1
    v10 := GenS1{F0: 0}
    v11 := GenS3{F0: 1}
    v1 = v3
    var v12 float64 = 1.5
    if ((false || false) || true) {
        var v13 uint32 = 4294967295
        switch ("" + ("" + "hello")) {
        case "foo bar":
            v12 = v12
        case "hello":
            var v14 int = (v1 << ((255) & 63))
            _ = v14
        }
        v3 += v3
        _ = v13
    } else {
        v15 := GenS4{F0: genFn1(9223372036854775807, 1), F1: 2}
        fmt.Println(v2, v1, v12)
        if (false && true) {
            if true {
                v12 -= 0.0
                fmt.Println(v1, v2)
            }
        }
        _ = v15
    }
    switch "\n\t" {
    case "foo bar":
        switch v1 {
        case 128:
            for fi16 := 0; fi16 < 11; fi16++ {
                v17 := []int{}
                for fi18 := 0; fi18 < 9; fi18++ {
                    _ = len(v17)
                    v12 = 1e10
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                    _ = fi18
                }
                _ = fi16
                _ = v17
            }
        }
    }
    v19 := []int{v3, -58, (v3 + 1), (v1 >> ((2) & 63))}
    for fi20 := 0; fi20 < 7; fi20++ {
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 uint32 = 597
            fmt.Println(fi21, v22, v22)
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        switch (v3 ^ 32768) {
        case 0:
            fmt.Println(v12, v12)
        case 4294967295:
            if (503 == 588) {
                for fi24 := 0; fi24 < 10; fi24++ {
                    v25 := GenS5{F0: (v2 & v2), F1: 1, F2: -32768}
                    var v26 uint8 = 76
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                var v27 int8 = 100
                _ = v27
            } else {
                if (!false) {
                    var v28 uint = 146
                    _ = v28
                } else {
                    var v29 float32 = -2.25
                    _ = v29
                }
            }
        case 2147483647:
            v1 = (v3 ^ -59)
        }
        fmt.Println(v2)
        _ = fi23
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        v12 = ((v12 - ((float64(1e10) / (v12 - -2.25)) - (0.0 + (v12 - 1e10)))) * ((v12 + (v12 + v12)) / ((-2.25 + ((v12 + 0.1) / (v12 * (0.1 * ((1e10 * ((v12 + v12) + (float64(0.0) / (0.0 + ((((((0.1 * (float64(1e10) * (((0.0 * v12) + 1.5) / v12))) * v12) + -2.25) + (v12 / v12)) / (0.1 + (v12 - (float64((((-2.25 * 3.14159) + 1e10) + 1.5)) + v12)))) + float64(1e10)))))) - (v12 / (0.0 - 3.14159))))))) / 3.14159)))
        _ = fi30
    }
    v1 += (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
}
