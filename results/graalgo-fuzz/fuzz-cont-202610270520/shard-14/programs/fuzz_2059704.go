package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 int32; F2 uint8 }
func genFn0(p0 uint8) uint32 {
    fmt.Println(p0)
    for fi7 := 0; fi7 < 6; fi7++ {
        if (true || true) {
            var v8 float32 = 3.14159
            fmt.Println(p0, fi7)
            _ = v8
        }
        var v9 uint16 = 65535
        var v10 uint16 = (v9 * (v9 ^ v9))
        _ = fi7
        _ = v10
        _ = v9
    }
    var v11 bool = (-2.25 < 3.14159)
    fmt.Println(v11, v11)
    v12 := []int{}
    _ = p0
    _ = v11
    _ = v12
    return 2147483648
}
func genFn1(p0 int16) int16 {
    var v13 int8 = -17
    v14 := []int{128, 42, 0}
    _ = p0
    _ = v13
    _ = v14
    return -481
}

func main() {
    var v1 int16 = 0
    var v2 int16 = (v1 | (v1 % ((v1) | 1)))
    var v3 float64 = 0.0
    var v4 string = "\n\t"
    v5 := []int{}
    var v6 int32 = 1
    v5 = append(v5, -1)
    if (!true) {
        var v15 uint32 = 2147483648
        _ = v15
    } else {
        if (27 < 0) {
            var v16 int16 = (v1 % (((v1 - (v1 | (v2 << ((0) & 63))))) | 1))
            v17 := GenS1{F0: 18446744073709551615, F1: 355}
            if (true || (42 <= -46)) {
                fmt.Println(v3)
            } else {
                switch 1 {
                case -9223372036854775808:
                    v3 = v3
                case 258:
                    if true {
                        var v18 uint8 = 2
                        fmt.Println(v2, v1)
                        _ = v18
                    }
                default:
                    v19 := GenS1{F0: 4294967295, F1: 512}
                    _ = v19
                }
                v6 -= (v6 | v6)
            }
            _ = v16
            _ = v17
        } else {
            var v20 string = ("hello" + (v4 + (v4 + v4)))
            _ = v20
        }
    }
    v5 = append(v5, -2147483648)
    var v21 float64 = float64((0.0 * 1e10))
    var v22 int64 = 0
    v23 := GenS2{F0: int(v1)}
    if (!(!(!false))) {
        var v24 uint8 = 229
        if (!(false || true)) {
            var v25 bool = true
            _ = v25
        }
        _ = v24
    }
    v5 = append(v5, 1)
    v4 = ""
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v26 uint = 1
    v21 -= v21
    if true {
        var v27 float64 = 0.0
        var v28 string = (v4 + ("\n\t" + ((v4 + v4) + ("\n\t" + ((v4 + v4) + v4)))))
        var v29 uint16 = 412
        _ = v27
        _ = v28
        _ = v29
    }
    v30 := GenS3{F0: -625, F1: (v6 & v6), F2: 220}
    if ((0.1 >= v3) || false) {
        var v31 uint32 = 65535
        _ = v31
    } else {
        v32 := GenS2{F0: 32767}
        _ = v32
    }
    if true {
        v33 := GenS3{F0: v22, F1: (v6 + v6), F2: 255}
        v2 += genFn1((v2 | genFn1(1)))
        _ = v33
    }
    if len(v5) > 0 {
        v5[0] = 1
    }
    var v34 uint16 = 212
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
