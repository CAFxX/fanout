package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int64 }
func genFn0(p0 int16) int8 {
    var v6 string = ("\n\t" + (("hello" + "") + "\n\t"))
    var v7 float32 = 3.14159
    v7 = (float32(1e10) + ((v7 + v7) / v7))
    _ = p0
    _ = v6
    _ = v7
    return 100
}

func main() {
    var v1 float32 = (float32(-2.25) * 0.1)
    var v2 uint32 = 2147483647
    var v3 int16 = 365
    var v4 string = (("" + "a") + "\n\t")
    var v5 uint64 = 1
    var v8 float32 = float32(((((((float64(((v1 * -2.25) * (3.14159 / v1))) - float64(((v1 * float32((1e10 + (3.14159 - 3.14159)))) * 3.14159))) * 0.0) / 3.14159) / (-2.25 - ((1.5 / (float64(3.14159) + float64(float32(float64(1e10))))) / 1.0))) - 1.5) - ((-2.25 + 1.5) - (((1e10 * ((0.0 - 1e10) - 0.1)) / 1e10) * float64(v1)))))
    v9 := GenS1{F0: v3, F1: (v2 % (((v2 - v2)) | 1))}
    fmt.Println(v2, v1)
    if (!false) {
        for fi10 := 0; fi10 < 15; fi10++ {
            for fi11 := 0; fi11 < 13; fi11++ {
                v5 = v5
                v12 := []int{(fi11 % ((fi10) | 1)), 32768, 32767, fi11}
                var v13 int64 = 1
                _ = fi11
                _ = v12
                _ = v13
            }
            if true {
                var v14 uint16 = 893
                if true {
                    v15 := []int{32768, (fi10 << ((656) & 63)), fi10}
                    v16 := GenS1{F0: v3, F1: 2}
                    _ = v15
                    _ = v16
                } else {
                    if true {
                        v17 := []int{-885}
                        var v18 int16 = (v3 % (((v3 * (v3 % ((v3) | 1)))) | 1))
                        v5 = 4294967295
                        _ = v17
                        _ = v18
                    }
                    var v19 string = v4
                    _ = v19
                }
                _ = v14
            } else {
                v20 := []int{}
                _ = v20
            }
            _ = fi10
        }
        v3 -= -949
    }
    switch 4294967295 {
    case -9223372036854775808:
        v5 += v5
    }
    switch 128 {
    case 624:
        if (3 > 1) {
            v21 := GenS1{F0: 352, F1: (v2 ^ 164)}
            _ = v21
        } else {
            v22 := []int{32768, -9223372036854775808, 503}
            _ = v22
        }
    case -48:
        var v23 int = 0
        _ = v23
    default:
        v24 := []int{128}
        _ = v24
    }
    v1 = ((v8 - 3.14159) / v1)
    for fi25 := 0; fi25 < 6; fi25++ {
        v26 := GenS1{F0: (v3 & int16(v2)), F1: v2}
        if true {
            v27 := GenS1{F0: (v3 * (v3 - 1)), F1: (v2 << ((18446744073709551615) & 63))}
            v28 := GenS1{F0: (v3 << ((172) & 63)), F1: 2}
            _ = v27
            _ = v28
        }
        _ = fi25
        _ = v26
    }
    v29 := GenS1{F0: 0, F1: (v2 | (v2 >> ((uint(v2)) & 63)))}
    v2 = (v2 | (v2 | v2))
    var v30 string = v4
    var v31 bool = (903 == 42)
    for fi32 := 0; fi32 < 20; fi32++ {
        for fi33 := 0; fi33 < 5; fi33++ {
            v34 := []int{(fi33 % (((fi32 + (fi32 ^ int(v3)))) | 1)), int(v3), -460, fi33}
            _ = fi33
            _ = v34
        }
        v35 := GenS2{F0: -334}
        _ = fi32
        _ = v35
    }
    fmt.Println(v4, v30)
    fmt.Println(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
