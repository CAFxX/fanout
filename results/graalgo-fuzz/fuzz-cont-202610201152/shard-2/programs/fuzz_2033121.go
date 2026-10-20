package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int16; F1 uint; F2 uint8 }
type GenS4 struct { F0 int32; F1 uint16 }
func genFn0(p0 uint16, p1 int64, p2 uint16) uint8 {
    var v5 uint16 = 65535
    switch "" {
    case "":
        if (((-855 == -379) && (!(1e10 > (float64(1.5) * (((0.1 + float64(0.1)) / 1.5) / (-2.25 / 3.14159)))))) && true) {
            switch "\n\t" {
            case "a":
                var v6 uint = 65535
                _ = v6
            case "":
                v7 := []int{942}
                _ = v7
            case "foo bar":
                var v8 uint32 = 2
                _ = v8
            default:
                p1 = p1
            }
        } else {
            var v9 uint32 = 42
            _ = v9
        }
    default:
        v10 := []int{}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}
func genFn1() bool {
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 float64 = 1.5
        var v13 uint8 = genFn0(1, -1, 2)
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 uint8 = genFn0(454, -1, 0)
    _ = v14
    return true
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[string]bool)
    var v3 string = "foo bar"
    v4 := make(map[uint32]int)
    v15 := GenS1{F0: -9223372036854775808, F1: 437}
    v3 = "\n\t"
    switch 339 {
    case 128:
        if true {
            var v16 int32 = 372
            v17 := []int{785, 32767, 42}
            if (v3 != (v3 + v3)) {
                for fi18 := 0; fi18 < 16; fi18++ {
                    var v19 float32 = 1.5
                    var v20 float32 = (v19 - v19)
                    var v21 bool = false
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                }
            }
            _ = v16
            _ = v17
        }
    case 9223372036854775807:
        v4[664] = int(v1)
    case 533:
        v1 -= (v1 - (v1 / (((^(v1))) | 1)))
    }
    var v22 string = (v3 + ((v3 + v3) + v3))
    v23 := []int{-1, 266, 32767}
    var v24 float32 = 3.14159
    switch 9223372036854775807 {
    case 128:
        var v25 uint8 = genFn0(215, 989, 1)
        _ = v25
    case 32768:
        fmt.Println(v1, v22)
    case 24:
        v26 := GenS2{F0: -171, F1: -1}
        _ = v26
    default:
        if true {
            v27 := GenS3{F0: 806, F1: 255, F2: 2}
            _ = v27
        } else {
            fmt.Println(v3)
            if (!true) {
                v28 := GenS3{F0: -32768, F1: 18446744073709551615, F2: 195}
                if (((false && (255 != 611)) || true) || (811 == 42)) {
                    if ((v1 != 163) || (18446744073709551615 >= 42)) {
                        if ((v1 >> ((18446744073709551615) & 63)) > (^(v1))) {
                            var v29 int = 128
                            v22 = ((v3 + "\n\t") + v3)
                            v24 -= v24
                            _ = v29
                        }
                    }
                    var v30 int32 = 42
                    _ = v30
                }
                v31 := GenS4{F0: 318, F1: 256}
                _ = v28
                _ = v31
            }
        }
    }
    v22 = v22
    var v32 uint32 = 2
    if (false || true) {
        var v33 string = (v3 + "\n\t")
        _ = v33
    }
    var v34 int16 = 32767
    fmt.Println(v1, v32, v24)
    var v35 int32 = 1
    for fi36 := 0; fi36 < 14; fi36++ {
        if len(v23) > 0 {
            v23[0] = (-(fi36))
        }
        _ = fi36
    }
    v22 = v22
    v37 := GenS2{F0: 1, F1: v35}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v32
    _ = v34
    _ = v35
    _ = v37
    _ = v4
}
