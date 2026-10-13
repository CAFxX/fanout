package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint }
func genFn0(p0 int, p1 uint64, p2 bool) uint32 {
    fmt.Println(p2, p1)
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = -2.25
    var v3 int32 = 835
    var v4 uint = 42
    fmt.Println(v4, v3, v4)
    var v5 uint8 = 194
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 4; fi7++ {
            v8 := GenS1{F0: (v4 | 1)}
            _ = fi7
            _ = v8
        }
        var v9 float64 = -2.25
        var v10 bool = true
        _ = fi6
        _ = v10
        _ = v9
    }
    if (false || true) {
        if false {
            var v11 int32 = (v3 / (((v3 & v3)) | 1))
            fmt.Println(v5)
            fmt.Println(v3, v11, v3)
            _ = v11
        }
        var v12 float64 = 1e10
        _ = v12
    }
    switch 2147483647 {
    case 247:
        var v13 float64 = 1.5
        _ = v13
    default:
        switch "" {
        case "":
            v14 := GenS1{F0: (^(v4))}
            _ = v14
        case "a":
            if (true || (false || (-2.25 < -2.25))) {
                v15 := GenS2{F0: v4, F1: int8(v4), F2: 358}
                v16 := GenS3{F0: 397, F1: 616, F2: 255}
                _ = v15
                _ = v16
            }
        }
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        if (178 >= 587) {
            switch ("foo bar" + "\n\t") {
            case "foo bar":
                v18 := []int{fi17}
                _ = v18
            default:
                v5 = 1
            }
            v19 := GenS3{F0: 2, F1: 4294967295, F2: 18446744073709551615}
            var v20 uint = 255
            _ = v19
            _ = v20
        } else {
            var v21 float64 = ((1e10 / 1.0) * 0.1)
            switch 0 {
            case 457:
                var v22 int32 = v3
                _ = v22
            case 4294967295:
                var v23 int16 = 556
                _ = v23
            }
            if (591 == 0) {
                v24 := GenS2{F0: (v4 & v4), F1: -128, F2: genFn0(fi17, 2, true)}
                _ = v24
            } else {
                switch -1 {
                case 0:
                    v5 = (^(v5))
                case -47:
                    for fi25 := 0; fi25 < 10; fi25++ {
                        v26 := []int{(^(fi25))}
                        _ = fi25
                        _ = v26
                    }
                case 32767:
                    var v27 uint8 = (v5 ^ 0)
                    _ = v27
                }
                v5 = (v5 & (v5 >> (((v4 ^ v4)) & 63)))
            }
            _ = v21
        }
        v28 := GenS1{F0: 0}
        var v29 string = ""
        _ = fi17
        _ = v28
        _ = v29
    }
    fmt.Println(v3)
    v30 := GenS1{F0: (v4 >> (((v4 / ((v4) | 1))) & 63))}
    for fi31 := 0; fi31 < 19; fi31++ {
        v4 = v4
        _ = fi31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
