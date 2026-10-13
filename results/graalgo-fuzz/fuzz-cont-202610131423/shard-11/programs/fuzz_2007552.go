package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
func genFn0(p0 uint8, p1 int) uint {
    var v4 bool = false
    v5 := GenS1{F0: 2, F1: (p0 & (p0 * (p0 % ((p0) | 1))))}
    if ((0.1 / 0.1) == 0.1) {
        var v6 int32 = 42
        v6 -= v6
        fmt.Println(p0)
        _ = v6
    }
    var v7 int16 = 42
    v8 := GenS1{F0: 2, F1: p0}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    return 261
}
func genFn1(p0 uint) uint64 {
    if (3.14159 == (3.14159 + (1.5 * 1.5))) {
        p0 = genFn0(128, 32767)
        switch 32768 {
        case 2147483647:
            p0 = (^(p0))
        case -9223372036854775808:
            fmt.Println(p0)
        }
    } else {
        p0 = (p0 * genFn0(42, -75))
    }
    v9 := []int{4294967295, -104, -2147483648, 1}
    var v10 int16 = 42
    fmt.Println(v10, p0)
    _ = p0
    _ = v10
    _ = v9
    return 1
}

func main() {
    var v1 uint32 = 4294967295
    v2 := []string{"\n\t", "", ("hello" + "\n\t"), "hello"}
    var v3 uint32 = (v1 | (v1 & 2))
    if (1e10 < -2.25) {
        v1 = 599
    }
    v3 += (v1 | (^(v1)))
    v3 -= v3
    fmt.Println(v3, v3)
    fmt.Println(v1)
    for fi11 := 0; fi11 < 12; fi11++ {
        v1 -= (^(v3))
        _ = fi11
    }
    fmt.Println(v1, v1)
    if ((-2.25 * 3.14159) == float32(float64(1.5))) {
        fmt.Println(v1)
        if false {
            fmt.Println(v3, v3)
        }
    } else {
        var v12 uint16 = 42
        var v13 int = -544
        v13 += 1
        _ = v12
        _ = v13
    }
    var v14 uint = genFn0(2, 9223372036854775807)
    fmt.Println(v1)
    v1 += v1
    if false {
        if len(v2) > 0 {
            v2[0] = "a"
        }
    }
    fmt.Println(v1, v14)
    switch 32767 {
    case 251:
        if true {
            fmt.Println(v1, v14, v1)
            fmt.Println(v3, v1, v14)
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if true {
                v15 := GenS1{F0: (v1 & 1), F1: 119}
                _ = v15
            }
            var v16 int64 = -1
            _ = v16
        }
    case -2147483648:
        for fi17 := 0; fi17 < 3; fi17++ {
            if true {
                switch (fi17 / (((fi17 - fi17)) | 1)) {
                case 12:
                    v1 += (v3 & (v1 ^ v3))
                default:
                    for fi18 := 0; fi18 < 18; fi18++ {
                        var v19 uint32 = v3
                        v20 := GenS1{F0: (v1 + (^(v3))), F1: 255}
                        var v21 int32 = 0
                        _ = fi18
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                }
                for fi22 := 0; fi22 < 6; fi22++ {
                    v3 += v3
                    _ = fi22
                }
                for fi23 := 0; fi23 < 17; fi23++ {
                    v14 = (v14 ^ genFn0(62, fi23))
                    _ = fi23
                }
            }
            if (!false) {
                var v24 uint16 = 255
                var v25 float64 = (float64((1.5 * float32((1.5 + 1e10)))) * 0.1)
                _ = v24
                _ = v25
            }
            v26 := GenS1{F0: v3, F1: 127}
            _ = fi17
            _ = v26
        }
    case -853:
        v2 = append(v2, "")
    }
    if (!(3.14159 >= ((3.14159 / -2.25) + 1e10))) {
        var v27 float32 = 1e10
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
}
