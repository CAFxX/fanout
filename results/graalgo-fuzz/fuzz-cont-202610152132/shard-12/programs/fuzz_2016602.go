package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int8 }
func genFn0(p0 int16, p1 uint64, p2 int8) uint16 {
    if (false && (65535 < 577)) {
        p0 = 42
        if false {
            var v7 uint64 = (p1 / (((p1 >> ((325) & 63))) | 1))
            for fi8 := 0; fi8 < 15; fi8++ {
                p1 = 42
                _ = fi8
            }
            _ = v7
        } else {
            var v9 float64 = (1.5 / (float64(3.14159) - ((1.5 * 3.14159) + float64(1.5))))
            var v10 int64 = -9223372036854775808
            _ = v10
            _ = v9
        }
        var v11 int32 = -1
        _ = v11
    } else {
        v12 := GenS1{F0: -2147483648}
        v13 := []int{-452, 32768, 4294967295}
        v14 := GenS1{F0: -1}
        _ = v12
        _ = v13
        _ = v14
    }
    if (true && true) {
        var v15 uint8 = 2
        p2 = p2
        _ = v15
    }
    v16 := GenS1{F0: -2147483648}
    switch "\n\t" {
    case "foo bar":
        for fi17 := 0; fi17 < 7; fi17++ {
            switch (^(fi17)) {
            case -1000:
                v18 := []int{(fi17 * fi17), 0}
                _ = v18
            case -1:
                v19 := []int{-1}
                _ = v19
            }
            _ = fi17
        }
    case "\n\t":
        for fi20 := 0; fi20 < 14; fi20++ {
            for fi21 := 0; fi21 < 16; fi21++ {
                fmt.Println(p0, p1, p0)
                if (0.0 < (0.0 + 3.14159)) {
                    var v22 string = "foo bar"
                    _ = v22
                }
                p0 += 780
                _ = fi21
            }
            _ = fi20
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 255
}

func main() {
    var v1 int = 9223372036854775807
    v2 := []uint32{2147483647, 2147483648}
    var v3 float64 = 0.0
    v4 := make(map[int]int)
    v5 := make(map[uint32]bool)
    var v6 uint8 = 2
    var v23 float64 = 1.5
    v2 = append(v2, 649)
    for fi24 := 0; fi24 < 13; fi24++ {
        var v25 float32 = 1.5
        _ = fi24
        _ = v25
    }
    v26 := GenS1{F0: 2147483647}
    v3 += ((v23 * ((v23 / (((((3.14159 - 0.1) / 1.5) - v23) - 0.1) / v23)) + -2.25)) + (-2.25 + ((((float64((float32((float64(0.1) - (v3 - 1e10))) * (((3.14159 + float32(v3)) * ((float32(v3) - 1e10) - -2.25)) + ((3.14159 * 3.14159) + 1.5)))) * (((((v3 * v3) * (v23 / 1.5)) / (1.5 / 0.1)) - v23) / v23)) - -2.25) * float64(((0.1 / float32(float64(-2.25))) * -2.25))) + (0.1 / float64((0.1 / -2.25))))))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v27 uint8 = v6
    v5[691] = (2147483647 == 0)
    v27 -= v27
    v28 := GenS1{F0: -61}
    v29 := v4[(v1 & v1)]
    v30 := GenS2{F0: 4294967295, F1: 2, F2: 4294967295}
    var v31 string = "\n\t"
    v32 := GenS1{F0: int32(v1)}
    switch "\n\t" {
    case "\n\t":
        v33 := GenS3{F0: -101}
        _ = v33
    }
    v1 = v1
    var v34 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
