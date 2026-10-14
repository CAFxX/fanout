package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
func genFn0(p0 uint32, p1 string, p2 uint64) int8 {
    fmt.Println(p2, p1, p1)
    switch 42 {
    case 781:
        var v4 int = 1
        _ = v4
    }
    var v5 int16 = 0
    var v6 uint8 = 128
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: 42, F1: -1}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 0
}
func genFn1() uint64 {
    var v9 int16 = -243
    v10 := []int{128, 123, 127, 4294967295}
    if false {
        var v11 uint = 1
        var v12 uint = (v11 * v11)
        var v13 uint8 = 42
        _ = v11
        _ = v12
        _ = v13
    } else {
        fmt.Println(v9)
        v9 -= (v9 / ((0) | 1))
        v9 = 42
    }
    _ = v10
    _ = v9
    return 4294967295
}

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{4294967295, 4, 1, 0}
    var v3 uint8 = 2
    var v14 uint32 = 4294967295
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v15 := GenS1{F0: 32767, F1: 423}
    fmt.Println(v14)
    if (v14 != 2147483647) {
        for fi16 := 0; fi16 < 14; fi16++ {
            for fi17 := 0; fi17 < 16; fi17++ {
                var v18 uint64 = 4294967295
                if len(v2) > 0 {
                    v2[0] = 42
                }
                _ = fi17
                _ = v18
            }
            var v19 int16 = 1
            _ = fi16
            _ = v19
        }
        if len(v2) > 0 {
            v2[0] = (v14 + 2)
        }
    } else {
        var v20 uint64 = 1
        var v21 float32 = -2.25
        var v22 uint8 = (^(v3))
        _ = v20
        _ = v21
        _ = v22
    }
    if len(v2) > 0 {
        v2[0] = 975
    }
    v23 := GenS1{F0: 1, F1: 0}
    if (v3 < (^(v3))) {
        var v24 uint32 = v14
        var v25 float64 = (1.5 + 0.1)
        for fi26 := 0; fi26 < 16; fi26++ {
            var v27 int8 = 127
            switch fi26 {
            case 32768:
                v14 = 1
            case -2147483648:
                var v28 int16 = 32767
                _ = v28
            case 4294967295:
                var v29 uint8 = (^(v3))
                _ = v29
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi26
            _ = v27
        }
        _ = v24
        _ = v25
    }
    var v30 uint16 = 255
    var v31 int8 = 127
    var v32 uint8 = v3
    v14 -= v14
    fmt.Println(v1, v30, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
