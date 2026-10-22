package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int; F1 uint16; F2 bool }
type GenS3 struct { F0 int }
func genFn0() int64 {
    switch 225 {
    case 4294967295:
        var v4 float32 = 0.1
        _ = v4
    }
    if (0.0 < ((0.0 - float32(((float64(0.1) * 1.5) * -2.25))) + (3.14159 + -2.25))) {
        switch -924 {
        case -469:
            var v5 uint8 = 224
            _ = v5
        }
        v6 := GenS1{F0: 65535, F1: 4294967295}
        if (9223372036854775807 != 211) {
            for fi7 := 0; fi7 < 3; fi7++ {
                var v8 uint64 = 0
                _ = fi7
                _ = v8
            }
            for fi9 := 0; fi9 < 2; fi9++ {
                for fi10 := 0; fi10 < 15; fi10++ {
                    if (true && true) {
                        var v11 float64 = (((1e10 / 0.1) + 3.14159) - 1.5)
                        v12 := []int{127}
                        var v13 int32 = 512
                        _ = v11
                        _ = v12
                        _ = v13
                    } else {
                        v14 := []int{fi10, 128}
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        _ = v14
                    }
                    switch "" {
                    case "a":
                        var v15 int32 = 924
                        _ = v15
                    default:
                        v16 := GenS1{F0: 541, F1: 1}
                        _ = v16
                    }
                    fmt.Println(fi9, fi9)
                    v17 := GenS1{F0: 2147483647, F1: 0}
                    _ = fi10
                    _ = v17
                }
                _ = fi9
            }
            var v18 float64 = 3.14159
            _ = v18
        }
        _ = v6
    }
    var v19 bool = (65535 >= 1)
    _ = v19
    return -181
}

func main() {
    var v1 uint16 = 547
    v2 := []int{128}
    var v3 int64 = -1
    v3 = (v3 - (v3 << ((4294967295) & 63)))
    v20 := GenS2{F0: -707, F1: v1, F2: ((v1 % ((65535) | 1)) != v1)}
    switch 577 {
    case 902:
        var v21 float32 = float32((1.5 - 1e10))
        _ = v21
    case -2147483648:
        var v22 int = -2147483648
        _ = v22
    case 128:
        var v23 float64 = 1.5
        _ = v23
    }
    v1 -= v1
    var v24 bool = (61 > 107)
    for fi25 := 0; fi25 < 16; fi25++ {
        v1 = (v1 - (v1 + v1))
        v3 = genFn0()
        _ = fi25
    }
    var v26 float32 = 3.14159
    v27 := GenS1{F0: 213, F1: 832}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if v24 {
        if ((1e10 - v26) >= (float32(0.0) / float32(-2.25))) {
            fmt.Println(v3, v26, v1)
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        var v28 uint64 = 1
        v29 := GenS2{F0: 377, F1: v1, F2: (48 > 42)}
        v24 = (v1 < 882)
        _ = v28
        _ = v29
    }
    if (9223372036854775807 == 42) {
        var v30 uint = 255
        if len(v2) > 0 {
            v2[0] = 1
        }
        if (!(!v24)) {
            var v31 int64 = 1
            var v32 uint32 = 65535
            if (v32 >= v32) {
                v33 := GenS3{F0: 9223372036854775807}
                v24 = v24
                v2 = append(v2, 1)
                _ = v33
            } else {
                var v34 uint32 = v32
                _ = v34
            }
            _ = v31
            _ = v32
        }
        _ = v30
    } else {
        var v35 float32 = 3.14159
        _ = len(v2)
        var v36 float32 = ((3.14159 + ((0.0 * float32(1.5)) * float32(((0.0 * -2.25) - 0.1)))) - (v35 + 0.0))
        _ = v35
        _ = v36
    }
    v26 = (float32((0.0 - 0.1)) / (float32(float64((v26 / v26))) / v26))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v26
    _ = v27
    _ = v3
}
