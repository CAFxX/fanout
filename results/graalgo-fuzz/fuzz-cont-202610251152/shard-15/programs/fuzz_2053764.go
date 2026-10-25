package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint32; F1 uint; F2 uint64 }
func genFn0(p0 int32, p1 int16) uint32 {
    for fi7 := 0; fi7 < 3; fi7++ {
        fmt.Println(p1)
        _ = fi7
    }
    fmt.Println(p0, p0)
    var v8 int = 407
    _ = p0
    _ = p1
    _ = v8
    return 1
}
func genFn1() uint16 {
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 uint8 = 127
        _ = fi9
        _ = v10
    }
    v11 := []int{32767, 330, 128, 0}
    var v12 int8 = -1
    _ = v11
    _ = v12
    return 1
}

func main() {
    var v1 uint32 = 42
    var v2 string = ((("hello" + "\n\t") + ("foo bar" + "a")) + "\n\t")
    var v3 float32 = float32((float64((1e10 / 1.0)) - (-2.25 - 1e10)))
    var v4 uint64 = 4294967295
    var v5 int32 = 0
    var v6 float32 = float32(-2.25)
    var v13 int = -370
    var v14 uint = 65535
    var v15 bool = true
    v2 = v2
    v14 = 18446744073709551615
    v16 := []int{(v13 & (v13 & (v13 / ((v13) | 1))))}
    fmt.Println(v6, v6, v2)
    var v17 int = v13
    v15 = (3.14159 == (float64(1.5) + 3.14159))
    if true {
        var v18 int32 = (v5 / ((int32(v4)) | 1))
        v5 = (v5 - (v5 + (^(v5))))
        v19 := GenS1{F0: 127, F1: uint8(v4)}
        _ = v18
        _ = v19
    }
    var v20 float32 = float32((3.14159 / (1e10 / 3.14159)))
    v15 = true
    v21 := GenS2{F0: (!v15), F1: (v5 & 1)}
    if len(v16) > 0 {
        v16[0] = v17
    }
    v22 := GenS2{F0: ((v14 - v14) > v14), F1: (v5 * v5)}
    if true {
        var v23 int8 = 42
        v16 = append(v16, 9223372036854775807)
        if (v14 <= (v14 % ((v14) | 1))) {
            v14 += (v14 - v14)
            if v15 {
                var v24 string = ("\n\t" + v2)
                _ = v24
            } else {
                v25 := GenS3{F0: uint32(v14), F1: v14, F2: (v4 * v4)}
                _ = v25
            }
        } else {
            if (!v15) {
                v26 := GenS1{F0: v23, F1: 127}
                _ = v26
            } else {
                var v27 uint = v14
                _ = v27
            }
        }
        _ = v23
    }
    v5 += v5
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
