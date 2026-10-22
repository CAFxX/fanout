package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int16 }
type GenS2 struct { F0 uint64; F1 bool }
func genFn0() int32 {
    var v4 bool = (!false)
    if (true && (3.14159 < (1.5 - (1.5 + 1.5)))) {
        var v5 bool = (!(v4 && (483 >= 213)))
        if v5 {
            var v6 uint32 = 147
            for fi7 := 0; fi7 < 5; fi7++ {
                if true {
                    for fi8 := 0; fi8 < 3; fi8++ {
                        v9 := []int{128, -9223372036854775808}
                        v10 := GenS1{F0: 65535, F1: -128, F2: 691}
                        if v4 {
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            var v11 int8 = -106
                            _ = v11
                        }
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                } else {
                    v5 = ((v5 || v4) || (fi7 == fi7))
                    v12 := []int{(fi7 - (^(fi7))), -561, 716}
                    _ = v12
                }
                _ = fi7
            }
            var v13 float32 = 0.1
            _ = v13
            _ = v6
        }
        _ = v5
    } else {
        for fi14 := 0; fi14 < 4; fi14++ {
            if (fi14 <= fi14) {
                v4 = (v4 || true)
            } else {
                fmt.Println(v4, v4)
                v15 := []int{fi14, (-(fi14)), fi14, fi14}
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                _ = v15
            }
            fmt.Println(fi14, v4)
            _ = fi14
        }
        v4 = (v4 || (!((127 > 0) && (v4 && false))))
        v4 = false
    }
    switch 127 {
    case 702:
        var v16 uint32 = 42
        _ = v16
    case -9223372036854775808:
        v17 := GenS2{F0: 4294967295, F1: false}
        _ = v17
    }
    _ = v4
    return -277
}

func main() {
    var v1 uint16 = 32767
    var v2 int8 = 0
    v3 := []uint32{2}
    var v18 int32 = 42
    v18 += genFn0()
    fmt.Println(v18, v1)
    var v19 float64 = 3.14159
    var v20 int16 = 42
    var v21 float64 = (v19 / ((1.5 - (v19 * (v19 + v19))) - v19))
    fmt.Println(v19, v2)
    var v22 uint64 = 538
    v23 := GenS2{F0: (^(v22)), F1: (!(!true))}
    var v24 uint8 = 128
    v18 = genFn0()
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
