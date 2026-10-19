package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint32 }
func genFn0(p0 uint64, p1 int8, p2 int16) uint32 {
    v7 := []int{-854, 32768, -716, 4294967295}
    p1 = (p1 - (p1 & 0))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint32 = 0
    v3 := make(map[uint32]string)
    var v4 float64 = 1.5
    var v5 float32 = 0.0
    v6 := make(map[uint32]bool)
    var v8 float32 = float32((((-2.25 - v4) + v4) + v4))
    fmt.Println(v2, v5, v5)
    var v9 int32 = (v1 % ((v1) | 1))
    if ((1.5 + v4) != v4) {
        delete(v6, (v2 << ((0) & 63)))
        v2 -= 2147483647
        v4 = (v4 - float64(1.5))
    }
    switch 128 {
    case -9223372036854775808:
        fmt.Println(v9)
    default:
        if ((!true) || true) {
            v10 := v6[0]
            fmt.Println(v1, v5)
            _ = len(v6)
            _ = v10
        } else {
            _ = len(v3)
            fmt.Println(v1)
        }
    }
    var v11 uint32 = 1
    if (127 > 127) {
        v12 := GenS1{F0: 1, F1: v9, F2: 42}
        _ = v12
    }
    v9 = (v1 & v1)
    v13 := GenS1{F0: 127, F1: v9, F2: 2}
    if true {
        v14 := GenS2{F0: 121}
        v5 = v5
        _ = v14
    }
    v15 := GenS2{F0: 18446744073709551615}
    v16 := GenS3{F0: (v1 % (((v1 - (^(v1)))) | 1)), F1: v11}
    switch "" {
    case "a":
        if (("a" > "") || true) {
            fmt.Println(v5)
        } else {
            fmt.Println(v5)
        }
    case "":
        for fi17 := 0; fi17 < 6; fi17++ {
            for fi18 := 0; fi18 < 19; fi18++ {
                var v19 float64 = (-2.25 + ((v4 + 1.5) / float64(float32((1e10 * 1e10)))))
                fmt.Println(v9, v4)
                v8 -= (v8 + (float32(0.1) + v8))
                fmt.Println(v8, v2)
                v20 := []int{(fi18 & (fi18 & fi18))}
                _ = fi18
                _ = v19
                _ = v20
            }
            switch fi17 {
            case -340:
                var v21 uint32 = v2
                _ = v21
            case 174:
                v22 := v3[(v11 - 42)]
                _ = v22
            case -508:
                v23 := v6[632]
                _ = v23
            default:
                fmt.Println(fi17, v9)
            }
            _ = fi17
        }
    }
    v4 = (v4 + 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
