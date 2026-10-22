package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 string; F1 int16; F2 uint64 }
func genFn0(p0 int8, p1 int, p2 int8) int32 {
    v6 := GenS1{F0: 2147483648, F1: 836}
    if false {
        v7 := GenS1{F0: 2, F1: 42}
        v8 := GenS1{F0: 191, F1: 2}
        p2 += (p2 + p0)
        _ = v7
        _ = v8
    }
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}

func main() {
    var v1 uint64 = 2
    var v2 int16 = 32767
    var v3 uint32 = 65535
    v4 := []int{-919}
    v5 := make(map[string]bool)
    if true {
        var v9 int = 2147483647
        for fi10 := 0; fi10 < 7; fi10++ {
            switch 42 {
            case 128:
                fmt.Println(v9, v2, v2)
            case 42:
                var v11 uint64 = 864
                _ = v11
            case -6:
                switch v9 {
                case 655:
                    for fi12 := 0; fi12 < 6; fi12++ {
                        var v13 int64 = -424
                        v14 := GenS1{F0: (v3 * (^(v3))), F1: (v1 | (v1 * (v1 << ((328) & 63))))}
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                default:
                    var v15 uint32 = v3
                    _ = v15
                }
            }
            _ = fi10
        }
        _ = v9
    }
    v1 = (v1 % ((v1) | 1))
    var v16 int8 = -128
    if len(v4) > 0 {
        v4[0] = -1
    }
    var v17 bool = false
    if v17 {
        v18 := v5["foo bar"]
        _ = v18
    }
    if v17 {
        var v19 int16 = 354
        v3 = v3
        _ = v19
    }
    v20 := GenS1{F0: v3, F1: v1}
    v21 := GenS2{F0: (("" + ("" + "foo bar")) + (("hello" + "") + "")), F1: (v2 / (((v2 ^ v2)) | 1)), F2: 42}
    v22 := GenS1{F0: (v3 / ((v3) | 1)), F1: 225}
    var v23 int = 2147483647
    fmt.Println(v17, v23, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
