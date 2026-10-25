package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint8 }
func genFn0(p0 int16, p1 bool) uint8 {
    var v7 string = (("\n\t" + "\n\t") + "foo bar")
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 float64 = 0.1
        for fi10 := 0; fi10 < 17; fi10++ {
            fmt.Println(fi8, fi8)
            var v11 uint16 = 2
            var v12 bool = false
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    v13 := []int{-9223372036854775808, 2147483647}
    _ = p0
    _ = p1
    _ = v13
    _ = v7
    return 219
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[string]bool)
    var v3 int = 32767
    var v4 bool = true
    var v5 uint32 = 0
    var v6 uint8 = v1
    var v14 float64 = 0.1
    v15 := GenS1{F0: -32768, F1: v3, F2: (v6 ^ v6)}
    var v16 int = (v3 / (((v3 - 9223372036854775807)) | 1))
    for fi17 := 0; fi17 < 14; fi17++ {
        v18 := GenS1{F0: -1, F1: fi17, F2: (v1 * v1)}
        var v19 uint8 = genFn0(-32768, false)
        _ = fi17
        _ = v18
        _ = v19
    }
    v4 = v4
    var v20 int64 = 0
    var v21 uint8 = genFn0(32767, (-2.25 <= float32(float64(0.1))))
    var v22 uint16 = 920
    var v23 int64 = int64(v21)
    if true {
        switch v16 {
        case 0:
            v24 := []int{0}
            _ = v24
        case 2147483647:
            var v25 float64 = -2.25
            _ = v25
        default:
            v1 = (v6 | v1)
        }
        fmt.Println(v14, v6)
        delete(v2, "a")
    }
    delete(v2, "hello")
    var v26 bool = (((!(0.1 != 0.0)) && false) && v4)
    v21 = (v1 / (((v6 * (v6 * 42))) | 1))
    v27 := GenS1{F0: -1, F1: (-(v16)), F2: (v1 ^ (^(v6)))}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
