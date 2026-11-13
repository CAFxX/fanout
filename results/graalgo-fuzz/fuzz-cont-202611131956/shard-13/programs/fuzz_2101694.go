package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int32 }
type GenS2 struct { F0 int32; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint64 }
func genFn0(p0 bool) int64 {
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := []int{-1, (-(fi6)), (fi6 - fi6), -959}
        var v8 string = "a"
        p0 = (65535 == 255)
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 255, F1: 172, F2: 314}
    v10 := []int{}
    p0 = p0
    v10 = append(v10, 32768)
    _ = p0
    _ = v10
    _ = v9
    return 597
}
func genFn1(p0 uint32) bool {
    p0 = (p0 ^ (p0 + (p0 - (p0 * (p0 * (p0 / ((p0) | 1)))))))
    p0 = (^(p0))
    v11 := GenS1{F0: 65535, F1: 2, F2: 0}
    _ = p0
    _ = v11
    return ((!(!(2 <= 167))) || false)
}

func main() {
    var v1 int = 2147483647
    v2 := []uint32{4294967295, 1, 776, 0}
    var v3 uint8 = 1
    var v4 float64 = 3.14159
    var v5 int32 = 973
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v3 += (v3 % ((v3) | 1))
    }
    if (true || (v4 != (v4 + 1.5))) {
        var v12 uint8 = v3
        v3 -= (v3 % (((v12 >> ((0) & 63))) | 1))
        if true {
            if true {
                for fi13 := 0; fi13 < 8; fi13++ {
                    v3 = v12
                    v5 += 1
                    _ = fi13
                }
            }
        }
        _ = v12
    }
    v14 := GenS1{F0: 0, F1: 4294967295, F2: v5}
    fmt.Println(v5)
    fmt.Println(v1, v1)
    if true {
        v5 = (v5 + (v5 ^ -2147483648))
        _ = len(v2)
        v4 += 1e10
    } else {
        v15 := GenS1{F0: 122, F1: 492, F2: (^(v5))}
        if true {
            if (!false) {
                if false {
                    fmt.Println(v4)
                    var v16 bool = (-9223372036854775808 > 0)
                    _ = v16
                }
                var v17 int = (v1 & -897)
                _ = v17
            }
            if true {
                v1 += 32768
                if (943 <= 1) {
                    v2 = append(v2, 1)
                } else {
                    var v18 string = "foo bar"
                    _ = v18
                }
                v5 += -816
            } else {
                v5 -= (-(v5))
                v5 += (v5 << ((661) & 63))
                switch v1 {
                case 219:
                    var v19 int32 = (v5 & (v5 * v5))
                    _ = v19
                default:
                    v20 := GenS2{F0: (v5 + (-(v5))), F1: 0, F2: 100}
                    _ = v20
                }
            }
            var v21 int = v1
            _ = v21
        } else {
            v1 = v1
        }
        _ = v15
    }
    v22 := GenS1{F0: 2, F1: 42, F2: v5}
    v5 -= v5
    v23 := GenS3{F0: v3, F1: 0}
    fmt.Println(v4, v4)
    for fi24 := 0; fi24 < 20; fi24++ {
        var v25 float64 = 0.0
        _ = fi24
        _ = v25
    }
    v26 := GenS2{F0: -1, F1: 270, F2: -125}
    v4 += float64(((float32(float64(1e10)) + ((1.5 / 1.5) * 1.5)) - -2.25))
    v4 -= 3.14159
    v5 -= -201
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
