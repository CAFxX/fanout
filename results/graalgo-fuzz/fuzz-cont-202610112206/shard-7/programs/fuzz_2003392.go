package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint8 }
func genFn0(p0 bool, p1 string) uint64 {
    v4 := GenS1{F0: 172, F1: -644, F2: 668}
    p1 = p1
    p0 = false
    p1 = ((p1 + "a") + (p1 + "\n\t"))
    _ = p0
    _ = p1
    _ = v4
    return 1
}
func genFn1(p0 bool, p1 uint16, p2 uint64) int64 {
    for fi5 := 0; fi5 < 10; fi5++ {
        fmt.Println(fi5, fi5, fi5)
        _ = fi5
    }
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            fmt.Println(fi7, fi6)
            if (p1 <= (p1 >> ((108) & 63))) {
                p2 -= 4294967295
                var v8 uint8 = 1
                v9 := []int{}
                fmt.Println(p2, p2)
                p2 += p2
                p2 = (p2 & p2)
                _ = v8
                _ = v9
            }
            if (true || true) {
                var v10 uint32 = 2147483648
                _ = v10
            } else {
                if (false || p0) {
                    if (0 != 0) {
                        var v11 float32 = (1e10 - float32(1.5))
                        var v12 int64 = 0
                        _ = v11
                        _ = v12
                    }
                }
            }
            switch "a" {
            case "hello":
                p2 = (p2 % (((p2 / (((^(p2))) | 1))) | 1))
            case "\n\t":
                v13 := []int{fi6, fi7}
                _ = v13
            }
            _ = fi7
        }
        var v14 int = (fi6 ^ fi6)
        _ = fi6
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    return -367
}

func main() {
    var v1 int = 4294967295
    var v2 float64 = 0.0
    var v3 int16 = -508
    fmt.Println(v2)
    switch v1 {
    case 0:
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 uint64 = 0
            var v17 float64 = (0.1 * 0.1)
            if false {
                var v18 int64 = -9223372036854775808
                v19 := []int{fi15, fi15}
                _ = v18
                _ = v19
            } else {
                var v20 uint = 4294967295
                fmt.Println(v20, v1)
                _ = v20
            }
            _ = fi15
            _ = v16
            _ = v17
        }
    case 650:
        v21 := []int{v1, v1, (v1 * (v1 / (((v1 - v1)) | 1)))}
        _ = v21
    }
    v22 := GenS1{F0: 762, F1: v3, F2: 2147483648}
    v23 := GenS1{F0: 409, F1: (^(v3)), F2: 2147483648}
    v24 := []int{v1, 32768, (v1 & -401), (v1 ^ (v1 % (((v1 ^ (v1 & (v1 % ((v1) | 1))))) | 1)))}
    fmt.Println(v1, v3, v2)
    fmt.Println(v2, v3, v2)
    for fi25 := 0; fi25 < 13; fi25++ {
        for fi26 := 0; fi26 < 9; fi26++ {
            var v27 uint8 = 1
            if len(v24) > 0 {
                v24[0] = (fi26 | (fi26 / ((fi25) | 1)))
            }
            _ = fi26
            _ = v27
        }
        if true {
            var v28 int32 = -1
            v29 := GenS1{F0: 65535, F1: v3, F2: 42}
            _ = v28
            _ = v29
        } else {
            var v30 bool = (128 < 2)
            v3 = (v3 / (((v3 ^ (-(v3)))) | 1))
            var v31 int32 = 1
            _ = v30
            _ = v31
        }
        var v32 float32 = 3.14159
        _ = fi25
        _ = v32
    }
    v33 := GenS1{F0: 32, F1: (v3 & v3), F2: 4294967295}
    var v34 int = 4294967295
    if (-784 <= 0) {
        v2 = ((v2 + float64(0.1)) + (((1e10 - (v2 - (float64((3.14159 / ((0.0 - 0.0) + 1e10))) - 3.14159))) + v2) * (v2 / (v2 - (v2 * float64(0.1))))))
        if len(v24) >= 2 {
            v24 = v24[1:len(v24)]
        }
        if len(v24) >= 2 {
            v24 = v24[1:len(v24)]
        }
    }
    v35 := GenS2{F0: 143, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v33
    _ = v34
    _ = v35
}
