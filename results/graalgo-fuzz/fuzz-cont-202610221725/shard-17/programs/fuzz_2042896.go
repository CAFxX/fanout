package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 int64; F2 int16 }
func genFn0(p0 int64, p1 uint) int16 {
    if true {
        var v7 uint64 = 1
        var v8 uint32 = 42
        v9 := GenS1{F0: 125}
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        p1 += p1
        p0 -= 512
        var v11 int64 = p0
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    return 32767
}

func main() {
    var v1 string = "a"
    v2 := make(map[uint32]string)
    var v3 uint8 = 42
    v4 := make(map[uint32]uint32)
    var v5 uint16 = 255
    var v6 float32 = (0.0 * (1.5 / (0.1 / 0.1)))
    var v12 int8 = -128
    v6 += v6
    if true {
        v4[2] = uint32(v5)
        v1 = (v1 + ((v1 + (v1 + v1)) + "foo bar"))
    }
    var v13 string = v1
    var v14 float32 = 0.1
    var v15 uint = 255
    fmt.Println(v14)
    v16 := []int{128, 157, -2147483648}
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 int = -2147483648
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
        _ = fi17
        _ = v18
    }
    v19 := GenS1{F0: -16}
    var v20 string = "\n\t"
    if false {
        var v21 string = "\n\t"
        fmt.Println(v3, v5)
        v14 += (v6 / (v6 / (v6 - v14)))
        _ = v21
    }
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    var v22 int = 128
    v15 = v15
    for fi23 := 0; fi23 < 9; fi23++ {
        for fi24 := 0; fi24 < 18; fi24++ {
            if (fi24 < fi23) {
                v22 -= (v22 >> ((42) & 63))
            } else {
                fmt.Println(v20, v5, v22)
                var v25 uint64 = 533
                var v26 int8 = v12
                v27 := v2[332]
                _ = v25
                _ = v26
                _ = v27
            }
            var v28 string = v20
            if ((((!false) && true) || (1e10 >= 0.1)) || false) {
                if (!true) {
                    var v29 float64 = 3.14159
                    var v30 uint32 = 2
                    _ = v29
                    _ = v30
                } else {
                    fmt.Println(v14)
                    v14 -= 0.1
                    v3 = 42
                }
                v12 -= 100
            } else {
                v31 := GenS2{F0: v15, F1: -584, F2: int16(v15)}
                v32 := GenS1{F0: (v12 + (v12 % (((v12 & int8(v22))) | 1)))}
                v6 += (v6 / (((0.1 - v6) * v14) + v6))
                fmt.Println(v28, v13)
                var v33 int = (-(fi24))
                _ = v31
                _ = v32
                _ = v33
            }
            _ = fi24
            _ = v28
        }
        _ = fi23
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
