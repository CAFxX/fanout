package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS2 struct { F0 int16; F1 int16; F2 uint8 }

func main() {
    var v1 uint = 4294967295
    var v2 float32 = float32(1.5)
    var v3 float64 = 3.14159
    v4 := make(map[int]string)
    var v5 float64 = -2.25
    var v6 float64 = (v5 / v3)
    var v7 string = "a"
    var v8 float64 = v6
    fmt.Println(v8, v2, v3)
    if false {
        v9 := GenS1{F0: 2147483648, F1: 100, F2: 42}
        fmt.Println(v5)
        var v10 string = "hello"
        _ = v10
        _ = v9
    } else {
        if true {
            fmt.Println(v1)
            v3 -= 1e10
        } else {
            var v11 uint = v1
            v3 = float64(v2)
            _ = v11
        }
    }
    v6 = v8
    if false {
        v3 -= v8
        var v12 int8 = 100
        _ = v12
    } else {
        var v13 uint16 = 679
        for fi14 := 0; fi14 < 16; fi14++ {
            for fi15 := 0; fi15 < 10; fi15++ {
                if (false && true) {
                    v6 = -2.25
                }
                v6 = v8
                v4[(fi15 << (((^(v1))) & 63))] = (v7 + v7)
                _ = fi15
            }
            _ = fi14
        }
        switch -9223372036854775808 {
        case 2147483647:
            v5 = 0.1
        default:
            v16 := GenS1{F0: 0, F1: -94, F2: uint32(v13)}
            _ = v16
        }
        _ = v13
    }
    v5 = float64((3.14159 - 0.1))
    v8 = v8
    if true {
        fmt.Println(v5, v7, v8)
        for fi17 := 0; fi17 < 7; fi17++ {
            if ((^(fi17)) == fi17) {
                v18 := GenS1{F0: 1, F1: -1, F2: 423}
                v6 -= 0.0
                _ = v18
            }
            _ = fi17
        }
        v19 := GenS2{F0: 0, F1: 0, F2: 127}
        _ = v19
    } else {
        v20 := GenS2{F0: 32767, F1: 0, F2: 42}
        _ = v20
    }
    if true {
        if false {
            v8 += v8
            var v21 float32 = 1.5
            v3 += 1e10
            _ = v21
        }
    }
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
