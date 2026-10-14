package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 uint32 }
func genFn0() int32 {
    var v6 uint8 = 127
    v7 := GenS1{F0: true, F1: 565}
    var v8 uint32 = 0
    for fi9 := 0; fi9 < 6; fi9++ {
        v8 += 42
        switch 32768 {
        case 4294967295:
            v10 := []int{0}
            _ = v10
        default:
            v6 += v6
        }
        _ = fi9
    }
    switch -9223372036854775808 {
    case 556:
        v11 := GenS2{F0: 960, F1: v8, F2: v6}
        _ = v11
    case 948:
        v6 -= (v6 << ((4294967295) & 63))
    case 0:
        if ((!true) || (uint(v8) >= 482)) {
            v12 := GenS2{F0: -1, F1: v8, F2: 128}
            var v13 bool = false
            _ = v12
            _ = v13
        }
    }
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := []int{-1, -2147483648, 128, -79}
    var v2 uint32 = 42
    v3 := []uint32{(^(v2))}
    var v4 uint16 = 42
    v5 := make(map[uint32]string)
    var v14 float32 = (-2.25 / float32((1.5 + (0.1 * (1e10 - float64(0.1))))))
    fmt.Println(v14, v14)
    var v15 float64 = (((-2.25 / 0.1) + (0.0 / 0.1)) * 1.5)
    var v16 int8 = int8(v2)
    v17 := GenS1{F0: true, F1: 0}
    var v18 int32 = 628
    for fi19 := 0; fi19 < 16; fi19++ {
        if false {
            var v20 int16 = -1
            v4 -= v4
            v20 = v20
            _ = v20
        }
        v2 -= 2147483647
        _ = fi19
    }
    var v21 bool = true
    var v22 int64 = 42
    v16 = -128
    v23 := GenS3{F0: 830, F1: int8(v22)}
    for fi24 := 0; fi24 < 4; fi24++ {
        if len(v3) > 0 {
            v3[0] = 65535
        }
        for fi25 := 0; fi25 < 20; fi25++ {
            if (v21 || true) {
                var v26 uint8 = 0
                switch "a" {
                case "hello":
                    if (v26 != v26) {
                        v27 := GenS1{F0: true, F1: 1000}
                        var v28 uint16 = v4
                        var v29 uint8 = v26
                        _ = len(v5)
                        _ = v27
                        _ = v28
                        _ = v29
                    } else {
                        var v30 int64 = (v22 + v22)
                        _ = v30
                    }
                }
                if (!v21) {
                    var v31 int = int(v18)
                    switch v31 {
                    case 42:
                        if v21 {
                            _ = len(v3)
                        } else {
                            v32 := GenS3{F0: 233, F1: (^(v16))}
                            _ = v32
                        }
                    case 127:
                        fmt.Println(v14)
                    case 32767:
                        v33 := GenS1{F0: (float32(v15) > (0.0 * ((float32(float64(0.0)) - (v14 * (v14 * (v14 + (0.0 - 0.1))))) + ((v14 / (v14 * v14)) * (v14 * float32(float64(v14))))))), F1: -556}
                        _ = v33
                    default:
                        v2 += (^(v2))
                    }
                    v34 := GenS4{F0: (v2 + v2)}
                    _ = v31
                    _ = v34
                }
                _ = v26
            }
            v35 := GenS4{F0: 0}
            v1 = append(v1, (fi25 / ((2147483647) | 1)))
            _ = fi25
            _ = v35
        }
        v36 := GenS4{F0: v2}
        _ = fi24
        _ = v36
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
