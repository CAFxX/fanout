package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint8; F1 int; F2 uint }
type GenS3 struct { F0 uint8; F1 int16; F2 string }
type GenS4 struct { F0 string; F1 int8; F2 int8 }
func genFn0(p0 bool, p1 bool) int {
    p1 = (!(true || true))
    v5 := GenS1{F0: -213, F1: -264}
    v6 := GenS1{F0: 2147483647, F1: 0}
    for fi7 := 0; fi7 < 8; fi7++ {
        p0 = (65535 >= 0)
        _ = fi7
    }
    var v8 int8 = -1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v8
    return 162
}
func genFn1(p0 uint8, p1 string) string {
    v9 := GenS1{F0: -1, F1: 1}
    var v10 int64 = -805
    fmt.Println(v10, p0)
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return p1
}

func main() {
    var v1 string = "foo bar"
    var v2 int = 387
    v3 := make(map[uint32]string)
    var v4 string = v1
    var v11 bool = true
    v12 := GenS1{F0: 612, F1: -9223372036854775808}
    var v13 int8 = 95
    var v14 int8 = v13
    v15 := GenS2{F0: 128, F1: -373, F2: 1}
    v16 := GenS1{F0: -2147483648, F1: -1}
    var v17 int64 = 0
    v18 := GenS3{F0: 1, F1: 1000, F2: (v1 + (v1 + "a"))}
    var v19 int16 = -509
    switch genFn0(true, (!v11)) {
    case 841:
        if true {
            var v20 int64 = v17
            var v21 uint16 = 255
            var v22 int = (v2 + (v2 / (((v2 % ((0) | 1))) | 1)))
            _ = v20
            _ = v21
            _ = v22
        } else {
            switch (v2 >> ((562) & 63)) {
            case -2147483648:
                for fi23 := 0; fi23 < 15; fi23++ {
                    v24 := GenS4{F0: v1, F1: (v14 * (^(v14))), F2: int8(v17)}
                    if false {
                        if v11 {
                            v1 = (((v4 + "foo bar") + "foo bar") + ((v4 + ((((v4 + "hello") + ("" + v4)) + (v4 + v1)) + "foo bar")) + v1))
                        } else {
                            v19 = (v19 + v19)
                            var v25 bool = v11
                            var v26 float32 = float32(3.14159)
                            _ = v25
                            _ = v26
                        }
                    }
                    _ = fi23
                    _ = v24
                }
            case -319:
                var v27 float64 = 1e10
                _ = v27
            default:
                var v28 int = v2
                _ = v28
            }
            var v29 uint32 = 4294967295
            var v30 uint16 = 159
            _ = v29
            _ = v30
        }
    default:
        var v31 int64 = (v17 * v17)
        _ = v31
    }
    for fi32 := 0; fi32 < 4; fi32++ {
        v33 := GenS2{F0: 226, F1: (v2 % (((v2 + 128)) | 1)), F2: 255}
        _ = fi32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
