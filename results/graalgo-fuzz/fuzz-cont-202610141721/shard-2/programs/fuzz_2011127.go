package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint64) uint32 {
    v5 := GenS1{F0: 4294967295, F1: 42}
    var v6 uint64 = p0
    var v7 uint8 = uint8(p0)
    v7 = v7
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 uint8 = 42
    var v2 uint = 2
    var v3 uint32 = 1
    var v4 int16 = -32768
    v3 += (^(v3))
    if true {
        fmt.Println(v4, v4, v1)
    }
    var v8 bool = true
    var v9 int8 = -1
    for fi10 := 0; fi10 < 4; fi10++ {
        if v8 {
            if v8 {
                v1 += 1
            }
            if true {
                for fi11 := 0; fi11 < 2; fi11++ {
                    var v12 uint64 = 0
                    if (((uint64(v3) >= (v12 ^ (v12 % ((uint64(v3)) | 1)))) && (float32((((float64((0.1 * -2.25)) + 0.1) / 1e10) / 1.5)) > 1e10)) && ((v9 <= (v9 & (v9 % ((v9) | 1)))) || ((65535 < 42) && (!((v2 + (v2 >> ((v2) & 63))) == (^(v2))))))) {
                        v4 -= 48
                    }
                    v13 := GenS2{F0: (v2 | (v2 % ((v2) | 1)))}
                    v14 := []int{fi11, (fi11 % ((fi11) | 1)), fi11}
                    var v15 uint8 = v1
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                    _ = v15
                }
            }
            var v16 uint64 = 904
            _ = v16
        } else {
            if v8 {
                for fi17 := 0; fi17 < 13; fi17++ {
                    fmt.Println(v8)
                    v3 = (v3 / ((65535) | 1))
                    var v18 int16 = (^(v4))
                    var v19 int16 = (v4 | (v4 % ((1000) | 1)))
                    var v20 uint = 0
                    v1 = v1
                    _ = fi17
                    _ = v18
                    _ = v19
                    _ = v20
                }
                if v8 {
                    fmt.Println(v2, v3)
                    fmt.Println(v2, v2)
                } else {
                    var v21 int16 = v4
                    _ = v21
                }
            } else {
                var v22 uint16 = uint16(v3)
                for fi23 := 0; fi23 < 17; fi23++ {
                    if ((v8 || false) || ((v4 >> ((65535) & 63)) >= (v4 * -136))) {
                        v24 := GenS2{F0: 595}
                        _ = v24
                    } else {
                        v9 = 1
                    }
                    _ = fi23
                }
                v25 := []int{fi10}
                _ = v22
                _ = v25
            }
        }
        _ = fi10
    }
    fmt.Println(v3, v8)
    v26 := []int{-153, -484}
    v27 := GenS2{F0: v2}
    var v28 uint8 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
