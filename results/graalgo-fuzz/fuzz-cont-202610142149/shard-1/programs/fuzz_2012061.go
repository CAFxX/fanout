package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 uint64 }
func genFn0(p0 uint64) int64 {
    p0 -= p0
    p0 = p0
    if false {
        if false {
            var v6 bool = (((3.14159 / 1.0) + 3.14159) > ((((-2.25 - 0.0) * 1e10) / 3.14159) - (1.5 - 1.5)))
            v7 := GenS1{F0: 437}
            fmt.Println(v6)
            _ = v6
            _ = v7
        } else {
            var v8 uint32 = 171
            v9 := []int{}
            _ = v8
            _ = v9
        }
    }
    _ = p0
    return -51
}
func genFn1(p0 bool, p1 uint16, p2 int32) int {
    p1 = (p1 >> ((909) & 63))
    p2 -= p2
    p1 = 2
    p2 -= p2
    p2 += 1
    _ = p0
    _ = p1
    _ = p2
    return 32768
}

func main() {
    var v1 float64 = -2.25
    var v2 uint32 = 0
    var v3 int = 128
    var v4 int64 = -1
    var v5 float32 = ((float32(v1) * float32(((float64(float32(v1)) * ((-2.25 + 0.0) / (((float64(float32(0.1)) - (v1 + 3.14159)) - v1) / (3.14159 / v1)))) - float64(0.1)))) + 0.1)
    switch ("" + "\n\t") {
    case "\n\t":
        if (!(-128 >= -128)) {
            v4 += (v4 ^ 753)
            v1 += (v1 * -2.25)
        }
    default:
        v10 := []int{(v3 - genFn1((!false), 65535, 1)), v3}
        _ = v10
    }
    v11 := []int{v3, 290, (v3 / ((genFn1((1 <= 42), 65535, -1)) | 1))}
    if (!(true && true)) {
        v12 := GenS1{F0: (v2 | (v2 << ((42) & 63)))}
        fmt.Println(v1)
        _ = v12
    }
    v13 := GenS2{F0: 42}
    v4 = (v4 | (v4 | (v4 + (v4 ^ 42))))
    switch (v3 % ((-711) | 1)) {
    case 2147483647:
        switch v3 {
        case -9223372036854775808:
            v1 = v1
        case 328:
            if (v1 != 0.0) {
                v5 -= 0.1
            }
        default:
            v5 = v5
        }
    case -2147483648:
        v14 := GenS2{F0: (v3 & (v3 * v3))}
        _ = v14
    default:
        if true {
            switch (v3 * (v3 / ((v3) | 1))) {
            case -264:
                var v15 int16 = -32768
                _ = v15
            case -1000:
                var v16 uint16 = 255
                _ = v16
            }
            var v17 uint = 65535
            _ = v17
        } else {
            var v18 float32 = (v5 + v5)
            _ = v18
        }
    }
    if ("\n\t" >= "") {
        _ = len(v11)
        switch -1 {
        case 32768:
            v4 = (-(v4))
        case 42:
            _ = len(v11)
        case -861:
            _ = len(v11)
        default:
            fmt.Println(v3, v2)
        }
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        var v20 int32 = -369
        _ = fi19
        _ = v20
    }
    v4 = genFn0(551)
    v21 := GenS3{F0: 217, F1: 2}
    fmt.Println(v1, v1, v4)
    if len(v11) > 0 {
        v11[0] = (v3 / ((v3) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
