package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint16; F1 uint; F2 bool }
func genFn0(p0 uint8, p1 int64) int32 {
    for fi4 := 0; fi4 < 2; fi4++ {
        if false {
            if (4294967295 < 4294967295) {
                p0 += 2
            } else {
                var v5 string = ""
                _ = v5
            }
            v6 := []int{}
            p1 = (p1 * p1)
            _ = v6
        } else {
            v7 := GenS1{F0: 65535, F1: p0}
            _ = v7
        }
        if (!true) {
            var v8 uint64 = 278
            _ = v8
        }
        p0 -= (p0 | (p0 % ((255) | 1)))
        _ = fi4
    }
    if true {
        fmt.Println(p1)
        fmt.Println(p1)
    } else {
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 int = (fi9 / ((fi9) | 1))
            v11 := GenS2{F0: 18446744073709551615, F1: p0}
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 int = -624
        if ((v12 - v12) == (v12 | v12)) {
            var v13 uint = 48
            v14 := []int{v12, -629}
            v15 := GenS2{F0: (v13 | 881), F1: p0}
            _ = v13
            _ = v14
            _ = v15
        }
        _ = v12
    }
    if (false || ((-2147483648 != int32(p1)) && true)) {
        var v16 bool = false
        for fi17 := 0; fi17 < 11; fi17++ {
            p0 += (p0 >> ((65535) & 63))
            p1 = p1
            _ = fi17
        }
        switch 42 {
        case 9223372036854775807:
            v16 = (!((-789 > int64(p0)) && (100 > 0)))
        case 32767:
            v18 := []int{-382}
            _ = v18
        default:
            v19 := GenS2{F0: 1, F1: (^(p0))}
            _ = v19
        }
        _ = v16
    } else {
        v20 := GenS1{F0: 1, F1: (p0 ^ (^(p0)))}
        _ = v20
    }
    _ = p0
    _ = p1
    return 150
}

func main() {
    var v1 float64 = 0.1
    var v2 uint8 = 128
    var v3 float64 = 1.5
    var v21 uint16 = 2
    v2 = v2
    fmt.Println(v3, v3, v3)
    v22 := GenS1{F0: 2, F1: v2}
    v2 += 2
    fmt.Println(v3, v2)
    var v23 float32 = ((0.1 / (0.0 - 0.1)) - float32((v3 / 3.14159)))
    var v24 float64 = v3
    v25 := GenS1{F0: 0, F1: (v2 * v2)}
    switch -897 {
    case 42:
        v26 := GenS3{F0: (^(v21)), F1: uint(v21), F2: (!true)}
        _ = v26
    case 32768:
        fmt.Println(v1)
    default:
        v27 := []int{42, 1}
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
