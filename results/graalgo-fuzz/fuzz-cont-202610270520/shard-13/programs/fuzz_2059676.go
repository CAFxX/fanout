package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint8 }
func genFn0(p0 uint32) uint8 {
    v6 := []int{-972, 1, 32767, 1}
    v7 := GenS1{F0: 42, F1: 559, F2: 127}
    p0 = (p0 - p0)
    v8 := GenS1{F0: 42, F1: 4294967295, F2: 255}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 127
}

func main() {
    var v1 uint8 = 1
    v2 := []uint32{4294967295, 2147483648, 1, 4294967295}
    v3 := make(map[uint32]uint32)
    v4 := []int{-1}
    var v5 string = ("foo bar" + "a")
    switch 32768 {
    case -875:
        var v9 float32 = (0.1 + -2.25)
        _ = v9
    }
    var v10 uint16 = 276
    v5 = ""
    var v11 int8 = 100
    _ = len(v3)
    v12 := v3[288]
    switch v5 {
    case "":
        var v13 uint = 0
        _ = v13
    default:
        v10 += v10
    }
    if false {
        var v14 int8 = (^(v11))
        v10 = v10
        var v15 float64 = -2.25
        _ = v14
        _ = v15
    } else {
        if true {
            v10 -= 42
        }
        v1 = v1
        var v16 uint16 = (v10 * 177)
        _ = v16
    }
    v12 = v12
    delete(v3, v12)
    for fi17 := 0; fi17 < 16; fi17++ {
        if true {
            v1 = (v1 ^ (v1 % (((v1 * (v1 % ((v1) | 1)))) | 1)))
        } else {
            if (((v11 & v11) > int8(v10)) || true) {
                v18 := v3[(^(v12))]
                if false {
                    if false {
                        v3[uint32(v1)] = uint32(v10)
                    }
                    switch 4294967295 {
                    case -2147483648:
                        var v19 float64 = (((0.0 - -2.25) + (1.5 * float64(float32(3.14159)))) * 0.1)
                        _ = v19
                    case 9223372036854775807:
                        if len(v2) > 0 {
                            v2[0] = 2
                        }
                    default:
                        for fi20 := 0; fi20 < 7; fi20++ {
                            v21 := v3[(v12 | (v18 & (v18 ^ v18)))]
                            v22 := GenS1{F0: 2, F1: (fi17 & 42), F2: (^(v1))}
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                    }
                }
                var v23 float32 = (1.5 * 3.14159)
                _ = v18
                _ = v23
            } else {
                v24 := GenS1{F0: 0, F1: 2147483647, F2: v1}
                fmt.Println(v11, v11)
                _ = v24
            }
        }
        _ = fi17
    }
    var v25 float32 = float32(0.0)
    var v26 uint8 = (v1 % ((v1) | 1))
    v27 := GenS1{F0: 255, F1: -819, F2: (v1 + v26)}
    for fi28 := 0; fi28 < 17; fi28++ {
        if (false && ((v12 & 448) < v12)) {
            delete(v3, (v12 << ((42) & 63)))
            for fi29 := 0; fi29 < 3; fi29++ {
                fmt.Println(v5, v11)
                var v30 uint16 = v10
                fmt.Println(fi28)
                _ = fi29
                _ = v30
            }
        }
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
