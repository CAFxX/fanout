package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int32; F1 int8; F2 bool }

func main() {
    var v1 float64 = ((3.14159 + -2.25) / 1.0)
    var v2 string = "\n\t"
    var v3 uint = 255
    var v4 uint16 = 42
    var v5 uint = v3
    v6 := []string{"a", ("\n\t" + v2), ("foo bar" + v2), v2}
    v3 += (v3 % ((v3) | 1))
    fmt.Println(v4)
    _ = len(v6)
    if false {
        v7 := GenS1{F0: -1}
        switch 128 {
        case -9223372036854775808:
            v3 += (v3 - (v3 / (((v5 >> ((v3) & 63))) | 1)))
        case 811:
            var v8 bool = false
            _ = v8
        default:
            var v9 uint64 = 4294967295
            _ = v9
        }
        _ = v7
    }
    var v10 float32 = 1.5
    v2 = v2
    var v11 float64 = ((float64(float32(1.5)) - float64(v10)) / 3.14159)
    v4 = (v4 << (((v5 >> (((v5 - (v3 | v3))) & 63))) & 63))
    v12 := GenS1{F0: -718}
    if (2147483648 >= 2147483648) {
        v13 := GenS2{F0: 255, F1: "a"}
        v3 = (v3 << (((v5 ^ 4294967295)) & 63))
        _ = v13
    } else {
        v10 = (v10 + v10)
        var v14 bool = ((v5 != (^(v5))) || false)
        if v14 {
            var v15 uint16 = 1
            switch 879 {
            case -9223372036854775808:
                v4 -= v15
            case 4294967295:
                v5 = 0
            case 0:
                var v16 int = -813
                _ = v16
            }
            _ = v15
        }
        _ = v14
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        fmt.Println(fi17)
        if true {
            v5 += v3
        } else {
            var v18 bool = ((!false) || (v1 < float64((float32((1e10 - v11)) + (float32(1e10) * (v10 * 1e10))))))
            v3 = v3
            if v18 {
                fmt.Println(v1, v18)
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v19 := GenS2{F0: (v5 ^ v5), F1: v2}
                if (1 >= 42) {
                    _ = len(v6)
                    for fi20 := 0; fi20 < 14; fi20++ {
                        v21 := GenS2{F0: v3, F1: ((v2 + v2) + "foo bar")}
                        _ = fi20
                        _ = v21
                    }
                } else {
                    v22 := GenS2{F0: 0, F1: (v2 + v2)}
                    v10 = v10
                    _ = v22
                }
                _ = v19
            }
            _ = v18
        }
        for fi23 := 0; fi23 < 17; fi23++ {
            v24 := GenS3{F0: 791, F1: 0, F2: false}
            if ((v10 * (v10 + v10)) != v10) {
                fmt.Println(v4)
            }
            v6 = append(v6, v2)
            _ = fi23
            _ = v24
        }
        _ = fi17
    }
    v25 := GenS2{F0: v3, F1: v2}
    if false {
        v10 += v10
        if ((v2 >= ((v2 + ((v2 + ("a" + v2)) + "\n\t")) + (v2 + v2))) || true) {
            var v26 uint64 = uint64(v4)
            _ = v26
        } else {
            var v27 float32 = 1e10
            _ = v27
        }
        var v28 bool = (!(!false))
        _ = v28
    }
    if (26 == 255) {
        if false {
            _ = len(v6)
            v3 = 403
        }
        v10 = v10
    }
    fmt.Println(v3)
    v29 := GenS1{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
