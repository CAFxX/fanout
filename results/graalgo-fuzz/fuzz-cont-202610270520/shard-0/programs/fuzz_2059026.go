package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int64 }
func genFn0(p0 int, p1 uint) int64 {
    switch 0 {
    case 128:
        var v4 float32 = float32((1e10 / float64(((0.1 + 0.1) + 3.14159))))
        _ = v4
    }
    var v5 uint64 = 389
    v6 := []int{}
    var v7 int64 = 42
    for fi8 := 0; fi8 < 7; fi8++ {
        for fi9 := 0; fi9 < 14; fi9++ {
            if (v5 == (^(v5))) {
                if (((p1 | 779) > (p1 - (p1 >> ((65535) & 63)))) && true) {
                    fmt.Println(p1)
                }
                v5 = (v5 - (^(v5)))
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v10 := GenS1{F0: (v5 * v5)}
                _ = v10
            }
            v11 := GenS1{F0: v5}
            _ = fi9
            _ = v11
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return v7
}
func genFn1(p0 string, p1 uint, p2 uint32) int16 {
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 uint32 = p2
        _ = fi12
        _ = v13
    }
    p1 = p1
    p2 = (^(p2))
    var v14 uint32 = (p2 - (p2 * p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 0
}

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[string]bool)
    var v3 int8 = 0
    v15 := []int{-31, 128, -9223372036854775808}
    var v16 uint = 4294967295
    v17 := v2[""]
    if (uint8(v1) != 246) {
        var v18 int16 = 1
        v17 = (v17 && v17)
        _ = v18
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        fmt.Println(fi19, v3)
        var v20 uint8 = 127
        _ = fi19
        _ = v20
    }
    var v21 uint = (v16 / ((v16) | 1))
    var v22 int64 = 9223372036854775807
    fmt.Println(v3, v1, v17)
    var v23 uint16 = 65535
    var v24 uint64 = 0
    var v25 bool = false
    for fi26 := 0; fi26 < 7; fi26++ {
        fmt.Println(v25, v22, v1)
        for fi27 := 0; fi27 < 5; fi27++ {
            var v28 uint8 = 1
            var v29 int8 = 42
            v30 := GenS2{F0: (v24 / (((v24 & v24)) | 1)), F1: genFn0((fi27 - 9223372036854775807), 413)}
            if v17 {
                var v31 bool = (1 > 152)
                v32 := GenS1{F0: (v24 + v24)}
                _ = v31
                _ = v32
            } else {
                var v33 uint8 = v28
                v34 := GenS2{F0: 4294967295, F1: v22}
                _ = v33
                _ = v34
            }
            _ = fi27
            _ = v28
            _ = v29
            _ = v30
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
