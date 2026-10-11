package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint16; F1 int8; F2 string }
func genFn0(p0 int32, p1 uint32) uint {
    v5 := []int{4294967295, 4294967295, 32768}
    var v6 bool = (((19 > 18446744073709551615) && (255 >= 2)) && false)
    if true {
        if (!v6) {
            v6 = (("hello" >= "a") && false)
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        } else {
            for fi7 := 0; fi7 < 5; fi7++ {
                fmt.Println(fi7)
                var v8 float64 = 0.1
                _ = fi7
                _ = v8
            }
        }
        var v9 int16 = 42
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 18446744073709551615
}
func genFn1(p0 uint32) int32 {
    v10 := GenS1{F0: 0, F1: 127, F2: 500}
    if (((0.1 + ((float64((3.14159 * (0.0 * 0.0))) * -2.25) - (((-2.25 - (3.14159 - (3.14159 - 3.14159))) + (-2.25 * float64((-2.25 * 1e10)))) + 3.14159))) - float64(-2.25)) == (3.14159 * 1e10)) {
        var v11 uint = 4294967295
        _ = v11
    }
    p0 = 238
    v12 := []int{}
    _ = p0
    _ = v10
    _ = v12
    return -2147483648
}

func main() {
    var v1 uint32 = 2
    var v2 int64 = -410
    var v3 int = -250
    var v4 uint64 = 2
    var v13 float32 = 3.14159
    v14 := []int{0}
    v3 += v3
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    v15 := GenS2{F0: 127, F1: (v4 >> ((455) & 63))}
    if false {
        var v16 uint32 = 2147483648
        _ = v16
    }
    switch (-(v3)) {
    case -188:
        if true {
            if false {
                v17 := GenS3{F0: v4, F1: genFn1(uint32(v4))}
                _ = v17
            }
        } else {
            v18 := GenS3{F0: v4, F1: genFn1((^(v1)))}
            _ = v18
        }
    case 42:
        if false {
            fmt.Println(v13)
            var v19 uint32 = v1
            _ = v19
        } else {
            for fi20 := 0; fi20 < 5; fi20++ {
                var v21 float32 = (float32(-2.25) + ((0.0 - v13) + v13))
                v22 := GenS4{F0: (^(v4)), F1: 486}
                _ = fi20
                _ = v21
                _ = v22
            }
        }
    case -982:
        v2 += v2
    default:
        for fi23 := 0; fi23 < 19; fi23++ {
            var v24 uint16 = 65535
            _ = fi23
            _ = v24
        }
    }
    if false {
        for fi25 := 0; fi25 < 6; fi25++ {
            v1 -= (v1 | 2)
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            _ = fi25
        }
        for fi26 := 0; fi26 < 2; fi26++ {
            var v27 int16 = -32768
            if len(v14) > 0 {
                v14[0] = v3
            }
            _ = fi26
            _ = v27
        }
        for fi28 := 0; fi28 < 1; fi28++ {
            if ((v1 % ((v1) | 1)) >= (v1 % ((v1) | 1))) {
                v14 = append(v14, (fi28 << ((genFn0(-2147483648, 581)) & 63)))
                var v29 bool = true
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                _ = v29
            }
            _ = fi28
        }
    }
    v2 += (v2 >> ((103) & 63))
    v30 := GenS5{F0: 1, F1: 1, F2: ("foo bar" + "\n\t")}
    var v31 uint32 = v1
    var v32 float64 = float64((float32((float64(v13) / 1e10)) / (float32(1e10) - v13)))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
