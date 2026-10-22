package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint8; F1 int8; F2 int8 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint16) int {
    var v4 string = "a"
    v5 := GenS1{F0: 42, F1: 394, F2: 2}
    _ = p0
    _ = v4
    _ = v5
    return 524
}
func genFn1(p0 uint32) bool {
    var v6 uint8 = 2
    switch "" {
    case "a":
        for fi7 := 0; fi7 < 6; fi7++ {
            v8 := []int{genFn0(65535), -2147483648, 0, (fi7 / ((526) | 1))}
            _ = fi7
            _ = v8
        }
    case "":
        fmt.Println(p0, p0)
    }
    _ = p0
    _ = v6
    return false
}

func main() {
    v1 := []int{}
    var v2 bool = ((4294967295 <= 65535) && false)
    var v3 uint32 = 703
    v3 += (v3 * v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v9 float32 = ((float32(float64(((-2.25 - 1e10) - 0.0))) + float32(0.1)) / 1e10)
    var v10 int8 = -20
    _ = len(v1)
    v10 = (v10 - (v10 & -101))
    fmt.Println(v2)
    v1 = append(v1, -1)
    fmt.Println(v10, v9)
    v9 = -2.25
    switch ("" + "a") {
    case "a":
        if (v2 && (float32(1.5) <= 0.1)) {
            if ((v3 << ((42) & 63)) == 4294967295) {
                fmt.Println(v3, v3, v9)
            }
            v11 := GenS2{F0: 83, F1: 127, F2: (^(v10))}
            _ = v11
        }
    case "":
        switch "foo bar" {
        case "foo bar":
            var v12 int64 = -1
            _ = v12
        case "a":
            if (936 > 0) {
                v1 = append(v1, -510)
            }
        case "\n\t":
            v10 += (v10 + v10)
        }
    case "foo bar":
        v9 = 1.5
    }
    v13 := GenS3{F0: 42}
    if (!(("\n\t" != "") && true)) {
        for fi14 := 0; fi14 < 18; fi14++ {
            v9 -= float32((1e10 / 1e10))
            _ = fi14
        }
        _ = len(v1)
        v2 = v2
    } else {
        v15 := GenS2{F0: 255, F1: (v10 >> ((0) & 63)), F2: v10}
        v16 := GenS2{F0: 36, F1: 127, F2: v10}
        fmt.Println(v2)
        _ = v15
        _ = v16
    }
    var v17 uint16 = 174
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}
