package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint; F1 bool }
func genFn0() int16 {
    switch 42 {
    case 128:
        switch 32767 {
        case 0:
            v6 := GenS1{F0: 42, F1: false}
            _ = v6
        case 128:
            v7 := []int{-776, 78}
            _ = v7
        default:
            v8 := GenS1{F0: 2147483647, F1: (2147483647 <= -1)}
            _ = v8
        }
    case 32767:
        var v9 int = 552
        _ = v9
    case 2147483647:
        if (78 == 2147483647) {
            v10 := GenS1{F0: 730, F1: false}
            if true {
                v11 := []int{127, -676}
                var v12 uint32 = 2147483647
                _ = v11
                _ = v12
            }
            var v13 uint8 = 42
            _ = v10
            _ = v13
        } else {
            for fi14 := 0; fi14 < 18; fi14++ {
                var v15 uint16 = 0
                var v16 float64 = 3.14159
                for fi17 := 0; fi17 < 14; fi17++ {
                    if (42 <= -1) {
                        v16 -= 3.14159
                        v18 := []int{(^(fi14)), 1, fi14, fi14}
                        _ = v18
                    }
                    var v19 int8 = 42
                    var v20 uint8 = 127
                    _ = fi17
                    _ = v19
                    _ = v20
                }
                _ = fi14
                _ = v15
                _ = v16
            }
        }
    default:
        var v21 float64 = 1e10
        _ = v21
    }
    if false {
        if ((!(49 >= 42)) || (!((true || false) || true))) {
            var v22 float32 = float32(3.14159)
            v23 := GenS2{F0: 4294967295, F1: (0 < 2147483648)}
            for fi24 := 0; fi24 < 13; fi24++ {
                v22 = 1e10
                if ((float32(((1e10 - 0.1) - ((1.5 - (((float64(3.14159) + (3.14159 / 3.14159)) * 0.1) + 0.1)) - 1e10))) * v22) == (v22 * 1.5)) {
                    fmt.Println(v22)
                }
                fmt.Println(v22)
                _ = fi24
            }
            _ = v22
            _ = v23
        } else {
            var v25 int64 = 1
            _ = v25
        }
        var v26 uint16 = 986
        _ = v26
    }
    return 32767
}

func main() {
    v1 := []uint32{42}
    v2 := []string{""}
    v3 := make(map[int]bool)
    var v4 bool = (true && false)
    var v5 uint32 = 2
    v4 = false
    v27 := GenS2{F0: 4294967295, F1: (42 >= 575)}
    if v4 {
        _ = len(v3)
    } else {
        fmt.Println(v4, v5)
    }
    var v28 bool = false
    if (-1 > -703) {
        v29 := v3[128]
        _ = v29
    }
    var v30 int64 = 0
    var v31 int16 = 341
    var v32 int64 = v30
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
