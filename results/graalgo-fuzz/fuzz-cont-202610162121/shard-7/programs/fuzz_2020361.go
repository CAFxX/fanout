package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 int16; F1 bool }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []string{"foo bar", "", ""}
    v3 := []int{-677, 268, 128, 2147483647}
    var v4 int16 = 0
    var v5 int64 = 805
    var v6 float64 = ((0.0 + (0.0 * 0.0)) * 0.1)
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int64 = v5
        switch 4294967295 {
        case -208:
            v1 -= 1
        case 127:
            for fi9 := 0; fi9 < 20; fi9++ {
                var v10 int16 = v4
                v11 := GenS1{F0: 42}
                v10 = (v10 << ((4294967295) & 63))
                _ = fi9
                _ = v10
                _ = v11
            }
        case 9223372036854775807:
            v6 += float64(1.5)
        default:
            v2 = append(v2, "a")
        }
        _ = fi7
        _ = v8
    }
    if (-1 == -2147483648) {
        if true {
            if (4294967295 <= 0) {
                v12 := GenS2{F0: "\n\t"}
                for fi13 := 0; fi13 < 15; fi13++ {
                    var v14 int32 = 227
                    var v15 uint8 = 100
                    var v16 uint16 = 807
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
                for fi17 := 0; fi17 < 9; fi17++ {
                    v5 = (v5 - -703)
                    var v18 float32 = (-2.25 + 1e10)
                    var v19 uint16 = 0
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = v12
            }
            var v20 uint8 = 102
            _ = v20
        }
        v4 = v4
        for fi21 := 0; fi21 < 1; fi21++ {
            fmt.Println(fi21)
            v5 += (v5 << ((42) & 63))
            _ = fi21
        }
    }
    var v22 bool = true
    v22 = ((100 == 0) && v22)
    if false {
        for fi23 := 0; fi23 < 2; fi23++ {
            if ((3.14159 < 3.14159) || ((!true) || v22)) {
                v24 := GenS3{F0: 882, F1: (v1 >> ((65535) & 63))}
                v6 += 1.5
                _ = v24
            }
            _ = fi23
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v6 = (((v6 + ((v6 / v6) - (v6 + v6))) - v6) + (v6 + (1.5 / 1.0)))
        v22 = v22
    }
    var v25 int64 = (v5 / ((v5) | 1))
    for fi26 := 0; fi26 < 17; fi26++ {
        v1 += int64(fi26)
        _ = fi26
    }
    v27 := GenS4{F0: (v4 << ((18446744073709551615) & 63)), F1: (v4 == v4)}
    if true {
        v2 = append(v2, "\n\t")
    } else {
        v22 = ((((v4 | v4) > 1) || v22) || (484 != 4294967295))
    }
    if (v22 && (55 == 2)) {
        v28 := GenS4{F0: -763, F1: true}
        _ = v28
    }
    _ = len(v2)
    v29 := GenS2{F0: "\n\t"}
    v1 -= v1
    var v30 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
