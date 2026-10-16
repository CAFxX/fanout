package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint64; F1 uint; F2 int64 }
func genFn0(p0 bool) uint32 {
    switch "hello" {
    case "":
        if p0 {
            fmt.Println(p0)
            v7 := []int{-9223372036854775808, 4294967295, -608}
            if p0 {
                var v8 uint16 = 1
                for fi9 := 0; fi9 < 9; fi9++ {
                    v8 -= v8
                    for fi10 := 0; fi10 < 6; fi10++ {
                        var v11 string = (("" + ("" + "foo bar")) + "a")
                        v11 = (v11 + "\n\t")
                        _ = fi10
                        _ = v11
                    }
                    _ = fi9
                }
                _ = v8
            }
            _ = v7
        }
    case "foo bar":
        var v12 uint = 347
        _ = v12
    default:
        fmt.Println(p0)
    }
    fmt.Println(p0)
    p0 = p0
    v13 := []int{2147483647, 1}
    v14 := GenS1{F0: 0, F1: 913}
    _ = p0
    _ = v13
    _ = v14
    return 2147483648
}
func genFn1(p0 int, p1 string, p2 int) int16 {
    p1 = p1
    var v15 int16 = -746
    var v16 uint32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return -32768
}

func main() {
    var v1 int32 = 42
    v2 := make(map[uint32]bool)
    var v3 float64 = -2.25
    var v4 float32 = float32((float64(3.14159) - v3))
    var v5 bool = (861 <= 1)
    v6 := make(map[string]string)
    v17 := GenS1{F0: 1, F1: 674}
    v18 := GenS1{F0: 1, F1: 1}
    fmt.Println(v3)
    var v19 int = 812
    v5 = (210 != 0)
    var v20 uint64 = 1
    var v21 uint32 = 2147483647
    v22 := GenS1{F0: -102, F1: 18446744073709551615}
    var v23 int16 = 1
    for fi24 := 0; fi24 < 3; fi24++ {
        fmt.Println(fi24, v5)
        _ = fi24
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        v26 := GenS1{F0: -128, F1: 42}
        if ((!(v23 > v23)) || v5) {
            v27 := []int{v19}
            v28 := GenS2{F0: v20, F1: 740, F2: 1}
            _ = v27
            _ = v28
        } else {
            fmt.Println(v5, fi25, v1)
            v29 := GenS2{F0: (v20 * 42), F1: uint(v21), F2: 1}
            for fi30 := 0; fi30 < 6; fi30++ {
                var v31 uint8 = 255
                _ = fi30
                _ = v31
            }
            _ = v29
        }
        _ = fi25
        _ = v26
    }
    if (25 >= 0) {
        fmt.Println(v21)
        var v32 int = (v19 & 9223372036854775807)
        _ = v32
    }
    v33 := GenS2{F0: v20, F1: 961, F2: 0}
    var v34 uint32 = 0
    v35 := GenS2{F0: v20, F1: 18446744073709551615, F2: -797}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
