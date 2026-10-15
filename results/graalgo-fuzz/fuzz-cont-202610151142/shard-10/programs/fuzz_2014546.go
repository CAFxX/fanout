package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 int8; F1 int; F2 uint32 }
func genFn0(p0 bool, p1 int8) bool {
    for fi4 := 0; fi4 < 13; fi4++ {
        p1 += p1
        p1 += (p1 ^ (p1 - (p1 << ((18446744073709551615) & 63))))
        _ = fi4
    }
    if p0 {
        v5 := []int{32768}
        if len(v5) > 0 {
            v5[0] = -39
        }
        _ = v5
    }
    _ = p0
    _ = p1
    return p0
}

func main() {
    var v1 float32 = ((((((0.1 / -2.25) + 1e10) / 1e10) * (3.14159 * (-2.25 / 1.5))) / 3.14159) / 1.5)
    var v2 int8 = -84
    var v3 uint64 = 375
    v6 := []int{-2147483648, -1}
    v7 := GenS1{F0: 4294967295, F1: 230, F2: 2}
    var v8 bool = true
    var v9 int64 = 1
    var v10 int16 = 1
    for fi11 := 0; fi11 < 9; fi11++ {
        v10 = (v10 | 1000)
        for fi12 := 0; fi12 < 13; fi12++ {
            _ = len(v6)
            v10 = -32768
            var v13 uint16 = 65535
            _ = fi12
            _ = v13
        }
        if (v8 || v8) {
            switch "hello" {
            case "hello":
                v2 += 1
            case "":
                if v8 {
                    var v14 string = ""
                    for fi15 := 0; fi15 < 3; fi15++ {
                        v8 = (255 > 1)
                        var v16 int = (fi15 + -78)
                        _ = fi15
                        _ = v16
                    }
                    var v17 int8 = (v2 * (v2 / (((v2 / (((^(v2))) | 1))) | 1)))
                    _ = v14
                    _ = v17
                }
            default:
                fmt.Println(v2)
            }
            if (!v8) {
                var v18 uint64 = v3
                v19 := GenS2{F0: (fi11 + fi11), F1: v9}
                _ = v18
                _ = v19
            }
        } else {
            if len(v6) > 0 {
                v6[0] = 4294967295
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v20 := GenS2{F0: fi11, F1: v9}
            _ = v20
        }
        _ = fi11
    }
    v21 := GenS2{F0: -2147483648, F1: v9}
    if (0 >= 32767) {
        v2 -= v2
        var v22 string = "\n\t"
        _ = v22
    }
    if (65535 == 612) {
        v23 := GenS2{F0: -2147483648, F1: v9}
        v1 = (v1 + (v1 / ((v1 - (v1 * v1)) * v1)))
        v6 = append(v6, 1)
        _ = v23
    }
    v24 := GenS1{F0: 42, F1: 1, F2: 42}
    var v25 uint64 = (^(v3))
    v9 -= (v9 & v9)
    if v8 {
        for fi26 := 0; fi26 < 11; fi26++ {
            fmt.Println(v9, v25)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi26
        }
    } else {
        v2 += (v2 / ((100) | 1))
    }
    v27 := GenS3{F0: (v2 - -1), F1: 4294967295, F2: 397}
    switch 127 {
    case -897:
        var v28 int16 = 42
        _ = v28
    case -658:
        var v29 uint64 = v3
        _ = v29
    default:
        fmt.Println(v2, v25)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
