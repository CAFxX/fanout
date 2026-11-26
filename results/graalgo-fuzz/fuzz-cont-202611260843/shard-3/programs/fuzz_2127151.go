package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 bool }
type GenS2 struct { F0 int16; F1 uint32 }
func genFn0(p0 uint16) string {
    p0 = p0
    if (42 < 225) {
        v5 := GenS1{F0: 4294967295, F1: 255, F2: true}
        p0 += (p0 - 65535)
        var v6 uint32 = 489
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 uint8 = 215
            _ = fi8
            _ = v9
        }
        fmt.Println(fi7, fi7)
        _ = fi7
    }
    _ = p0
    return "\n\t"
}
func genFn1(p0 int32) string {
    switch -920 {
    case -152:
        if true {
            v10 := []int{}
            if len(v10) > 0 {
                v10[0] = 245
            }
            _ = v10
        } else {
            p0 += 0
            for fi11 := 0; fi11 < 5; fi11++ {
                p0 = p0
                _ = fi11
            }
        }
    default:
        p0 = (p0 | p0)
    }
    var v12 bool = true
    var v13 uint64 = 674
    v14 := []int{}
    v12 = (-9223372036854775808 == 191)
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return ""
}

func main() {
    var v1 int16 = 1000
    var v2 float32 = float32(1e10)
    var v3 uint16 = 65535
    var v4 uint8 = 197
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 uint64 = 2
        if true {
            v17 := GenS1{F0: uint(v3), F1: uint8(v3), F2: (1 != 42)}
            if true {
                v18 := []int{4294967295, (fi15 | (fi15 / (((fi15 / ((fi15) | 1))) | 1)))}
                switch "" {
                case "\n\t":
                    v18 = append(v18, fi15)
                case "a":
                    v4 -= v4
                }
                if (1e10 <= 1e10) {
                    var v19 uint16 = v3
                    v19 = (v3 % (((v3 % ((v3) | 1))) | 1))
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                    _ = v19
                }
                _ = v18
            }
            v2 = 0.0
            _ = v17
        } else {
            v1 = 1000
            var v20 int16 = (v1 * v1)
            v21 := []int{fi15, 32767, (fi15 >> ((71) & 63))}
            _ = v20
            _ = v21
        }
        var v22 uint32 = 2
        _ = fi15
        _ = v16
        _ = v22
    }
    var v23 uint64 = 2
    v23 = (v23 + v23)
    var v24 int64 = 42
    var v25 bool = (("hello" > ("\n\t" + "")) || true)
    v26 := []int{-9223372036854775808}
    v27 := GenS1{F0: 264, F1: (v4 * 108), F2: (265 <= 0)}
    for fi28 := 0; fi28 < 9; fi28++ {
        fmt.Println(v25)
        _ = fi28
    }
    fmt.Println(v25)
    v3 = (^(v3))
    var v29 int64 = v24
    if (v25 && v25) {
        v30 := GenS2{F0: (^(v1)), F1: 42}
        var v31 int16 = (v1 % (((-(v1))) | 1))
        var v32 int = 0
        _ = v30
        _ = v31
        _ = v32
    }
    if ((-53 >= 33) && v25) {
        if len(v26) >= 2 {
            v26 = v26[1:len(v26)]
        }
    }
    fmt.Println(v23)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v4
}
