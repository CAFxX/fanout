package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int8; F1 int8 }
func genFn0(p0 uint8) int {
    var v7 float32 = (float32(-2.25) - float32(0.0))
    if (!false) {
        if (0 <= -32768) {
            var v8 uint16 = 337
            v9 := []int{-340, 32767}
            for fi10 := 0; fi10 < 9; fi10++ {
                switch "foo bar" {
                case "":
                    var v11 uint = 255
                    _ = v11
                default:
                    var v12 float64 = 1e10
                    _ = v12
                }
                for fi13 := 0; fi13 < 1; fi13++ {
                    _ = len(v9)
                    var v14 uint = 991
                    v15 := GenS1{F0: 1, F1: (^(v14))}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                _ = fi10
            }
            _ = v8
            _ = v9
        } else {
            var v16 int64 = 1
            v17 := []int{}
            v18 := GenS1{F0: 1, F1: 4294967295}
            _ = v16
            _ = v17
            _ = v18
        }
        fmt.Println(p0, p0)
        for fi19 := 0; fi19 < 9; fi19++ {
            p0 = p0
            v20 := []int{fi19, fi19, (fi19 + fi19)}
            var v21 int64 = 1
            _ = fi19
            _ = v20
            _ = v21
        }
    } else {
        v22 := GenS1{F0: 4294967295, F1: 60}
        v23 := GenS2{F0: -128, F1: -54}
        v24 := []int{9223372036854775807, -651, 0, 913}
        _ = v22
        _ = v23
        _ = v24
    }
    p0 += p0
    _ = p0
    _ = v7
    return 32767
}
func genFn1() int8 {
    v25 := []int{9223372036854775807, 1}
    if true {
        for fi26 := 0; fi26 < 20; fi26++ {
            var v27 int64 = 1
            if (false && ((true && (672 != 2147483648)) || true)) {
                v27 += v27
            }
            for fi28 := 0; fi28 < 1; fi28++ {
                v27 -= (v27 + (-(v27)))
                v27 -= -9223372036854775808
                v27 = v27
                _ = fi28
            }
            _ = fi26
            _ = v27
        }
    }
    var v29 string = "foo bar"
    _ = v25
    _ = v29
    return 0
}

func main() {
    var v1 uint16 = 255
    var v2 int16 = 0
    v3 := []uint32{2147483648}
    v4 := make(map[uint32]bool)
    var v5 float32 = 3.14159
    var v6 uint8 = 255
    var v30 bool = true
    delete(v4, 2147483648)
    var v31 float64 = -2.25
    v4[0] = (!v30)
    v31 += 1e10
    var v32 uint64 = 853
    v33 := v4[4294967295]
    v33 = v30
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
