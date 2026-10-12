package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 uint64, p1 uint64) bool {
    p1 -= (p0 * p0)
    var v4 uint32 = 42
    for fi5 := 0; fi5 < 19; fi5++ {
        if ((!((0.0 + (0.1 + ((float64((1.5 / 3.14159)) + 1.5) + (float64((((0.1 * 0.1) + 1.5) + (0.1 * ((1.5 + 3.14159) + float32(((0.1 / 0.1) * ((((-2.25 + float64((3.14159 - 0.1))) - ((float64(3.14159) - 1.5) - (0.0 * (float64((1e10 / 1.0)) + ((3.14159 * 0.0) * ((0.0 - -2.25) - 0.1)))))) * 1.5) - 1.5))))))) + 3.14159)))) != 3.14159)) && true) {
            v6 := []int{-654, 353}
            _ = v6
        } else {
            for fi7 := 0; fi7 < 7; fi7++ {
                v8 := GenS1{F0: -1}
                var v9 int8 = -7
                _ = fi7
                _ = v8
                _ = v9
            }
            for fi10 := 0; fi10 < 11; fi10++ {
                v11 := []int{(fi10 << ((0) & 63)), fi10, -9223372036854775808}
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 1; fi12++ {
                v13 := []int{(^(fi5)), (fi5 % ((fi12) | 1))}
                fmt.Println(p0, p0, fi5)
                _ = fi12
                _ = v13
            }
        }
        v14 := []int{(fi5 + (fi5 % (((fi5 + (fi5 / (((-(fi5))) | 1)))) | 1)))}
        if ((5 > 161) && (false && ((false && true) || true))) {
            var v15 uint8 = 186
            if true {
                v14 = append(v14, fi5)
                for fi16 := 0; fi16 < 6; fi16++ {
                    if true {
                        v4 = 2
                        var v17 int32 = -292
                        _ = v17
                    }
                    fmt.Println(v15, fi5, fi5)
                    v4 -= (v4 & 1)
                    _ = fi16
                }
            } else {
                v4 -= v4
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
            }
            _ = v15
        } else {
            fmt.Println(p1, fi5)
        }
        _ = fi5
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v4
    return false
}

func main() {
    var v1 uint = 771
    v2 := []uint32{1, 2147483648, 0}
    var v3 bool = true
    fmt.Println(v1, v1)
    if ((((((v3 || (995 == 657)) && v3) && true) || (72 != 127)) || false) || v3) {
        var v18 uint8 = 128
        _ = v18
    } else {
        for fi19 := 0; fi19 < 11; fi19++ {
            for fi20 := 0; fi20 < 4; fi20++ {
                if v3 {
                    var v21 float32 = 3.14159
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v2 = append(v2, 992)
                    _ = v21
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi20
            }
            _ = fi19
        }
        var v22 int = 4294967295
        if len(v2) > 0 {
            v2[0] = 587
        }
        _ = v22
    }
    v3 = v3
    var v23 int = 1
    v3 = v3
    var v24 uint16 = 255
    v24 += (v24 - (v24 ^ (v24 - v24)))
    v1 += (v1 & (v1 / ((v1) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
