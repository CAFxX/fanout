package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 int64; F1 int16; F2 string }
type GenS7 struct { F0 uint; F1 int }

func main() {
    var v1 float64 = (0.0 + 1.5)
    v2 := make(map[int]uint32)
    var v3 float32 = 1e10
    v4 := make(map[uint32]uint32)
    if (!((-112 > 32768) && false)) {
        v1 += 1.5
        switch "hello" {
        case "a":
            v5 := []int{4294967295, 128}
            _ = v5
        case "foo bar":
            fmt.Println(v3, v1)
        case "\n\t":
            var v6 uint64 = 18446744073709551615
            _ = v6
        default:
            v1 = 1.5
        }
    } else {
        if true {
            var v7 uint64 = 2
            _ = v7
        } else {
            if ((!(!(9223372036854775807 <= 4294967295))) || true) {
                v3 -= 0.1
                for fi8 := 0; fi8 < 15; fi8++ {
                    switch (((("" + ("a" + "a")) + "foo bar") + "hello") + "foo bar") {
                    case "a":
                        var v9 int = fi8
                        _ = v9
                    }
                    v10 := GenS1{F0: 9223372036854775807}
                    _ = fi8
                    _ = v10
                }
            } else {
                v1 -= float64(((-2.25 / (v3 - v3)) * (((-2.25 / -2.25) / 0.1) + 0.1)))
            }
            v1 = 3.14159
            v11 := GenS2{F0: 522, F1: -9223372036854775808}
            _ = v11
        }
        v3 += 0.0
    }
    if (2147483647 != 485) {
        var v12 uint64 = 158
        v1 = (3.14159 * -2.25)
        _ = v12
    }
    if (2 >= 32767) {
        fmt.Println(v3, v1)
    }
    if true {
        if true {
            if true {
                switch "" {
                case "a":
                    _ = len(v4)
                }
                v13 := v2[32768]
                _ = v13
            } else {
                v1 += (v1 + v1)
                for fi14 := 0; fi14 < 2; fi14++ {
                    switch fi14 {
                    case 4294967295:
                        _ = len(v2)
                    case 1:
                        var v15 int8 = 0
                        _ = v15
                    case 32768:
                        var v16 bool = (true && true)
                        _ = v16
                    }
                    _ = fi14
                }
                v3 -= float32((0.1 + float64(float32(v1))))
            }
        } else {
            v17 := GenS3{F0: "a", F1: 1}
            fmt.Println(v3)
            v3 -= v3
            _ = v17
        }
        v3 += (float32(((v1 * v1) / v1)) + ((3.14159 / (v3 * 1e10)) - (v3 / (0.1 - float32(v1)))))
        if true {
            v18 := GenS4{F0: 785}
            _ = v18
        } else {
            var v19 float32 = (-2.25 - v3)
            v20 := GenS5{F0: 171, F1: "a"}
            _ = v19
            _ = v20
        }
    }
    v3 = v3
    var v21 int16 = 32767
    var v22 uint = 255
    var v23 bool = false
    for fi24 := 0; fi24 < 18; fi24++ {
        v25 := GenS6{F0: -1, F1: (v21 | -33), F2: "\n\t"}
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        for fi27 := 0; fi27 < 10; fi27++ {
            fmt.Println(v23, v21, v22)
            _ = fi27
        }
        _ = fi26
    }
    if (18446744073709551615 != 1) {
        for fi28 := 0; fi28 < 12; fi28++ {
            var v29 int8 = 42
            for fi30 := 0; fi30 < 4; fi30++ {
                fmt.Println(v29)
                if v23 {
                    v21 -= v21
                    v31 := GenS7{F0: v22, F1: (fi30 << ((v22) & 63))}
                    v22 = v22
                    _ = v31
                }
                fmt.Println(fi28)
                _ = fi30
            }
            _ = fi28
            _ = v29
        }
        var v32 int8 = 1
        _ = v32
    }
    v21 = 316
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
