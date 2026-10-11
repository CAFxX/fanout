package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS2 struct { F0 int32; F1 int16; F2 uint16 }
func genFn0(p0 int64) bool {
    v4 := []int{}
    _ = len(v4)
    p0 = p0
    _ = p0
    _ = v4
    return true
}
func genFn1(p0 uint, p1 int) uint8 {
    v5 := GenS1{F0: 14, F1: 239, F2: 4294967295}
    var v6 bool = (475 <= 1)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 128
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{}
    v3 := make(map[uint32]int)
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int = fi7
        v8 += (^(fi7))
        _ = len(v1)
        _ = fi7
        _ = v8
    }
    var v9 int16 = -251
    var v10 int = 32767
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v9, v10)
    switch (v10 / ((127) | 1)) {
    case -2147483648:
        v11 := v3[1]
        _ = v11
    case 891:
        v10 += v10
    case 32767:
        _ = len(v2)
    }
    v12 := GenS1{F0: 65535, F1: v9, F2: 0}
    delete(v3, 4294967295)
    var v13 uint = 4294967295
    if (0.0 == 3.14159) {
        fmt.Println(v13, v10, v10)
    }
    var v14 uint = v13
    v10 += v10
    var v15 int64 = 9223372036854775807
    var v16 uint32 = 978
    switch v10 {
    case -1:
        switch (v10 - v10) {
        case 32767:
            fmt.Println(v9, v16, v15)
        case -1:
            switch 127 {
            case 2147483647:
                v13 += v14
            default:
                fmt.Println(v13, v16, v9)
            }
        default:
            for fi17 := 0; fi17 < 9; fi17++ {
                v18 := GenS1{F0: 32767, F1: (v9 % ((int16(v15)) | 1)), F2: v16}
                var v19 uint64 = 1
                v20 := GenS2{F0: -110, F1: v9, F2: 65535}
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
        }
    case -2147483648:
        v16 += (^(v16))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v9
}
