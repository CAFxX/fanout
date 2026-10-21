package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int32 }
type GenS2 struct { F0 int32 }
func genFn0() bool {
    switch -954 {
    case -929:
        for fi5 := 0; fi5 < 9; fi5++ {
            var v6 int16 = 42
            var v7 int64 = 0
            _ = fi5
            _ = v6
            _ = v7
        }
    case -441:
        v8 := GenS1{F0: 32767, F1: -1, F2: -2147483648}
        _ = v8
    default:
        var v9 uint32 = 65535
        _ = v9
    }
    var v10 float32 = float32((float64(3.14159) + float64(float32(((1.5 - (0.1 * (0.0 + 0.0))) * -2.25)))))
    fmt.Println(v10)
    switch "\n\t" {
    case "foo bar":
        v11 := GenS1{F0: 32768, F1: 42, F2: -2147483648}
        _ = v11
    case "hello":
        v12 := []int{}
        _ = v12
    case "\n\t":
        v10 += (v10 * v10)
    default:
        var v13 uint32 = 65535
        _ = v13
    }
    _ = v10
    return true
}
func genFn1() uint16 {
    if (true || (93 >= 245)) {
        var v14 int8 = 11
        var v15 float64 = (0.0 + (-2.25 * float64(1.5)))
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        var v17 uint16 = 32767
        if (111 < 38) {
            v18 := []int{fi16, fi16}
            var v19 bool = true
            for fi20 := 0; fi20 < 17; fi20++ {
                v21 := GenS2{F0: -609}
                var v22 uint8 = 1
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = v18
            _ = v19
        }
        switch fi16 {
        case 32768:
            if false {
                for fi23 := 0; fi23 < 17; fi23++ {
                    fmt.Println(fi23, fi23, fi16)
                    _ = fi23
                }
                v24 := []int{fi16, fi16}
                _ = v24
            }
        case 42:
            v17 = (^(v17))
        case 128:
            var v25 uint = 1
            _ = v25
        default:
            v26 := []int{}
            _ = v26
        }
        _ = fi16
        _ = v17
    }
    var v27 int = -412
    switch 515 {
    case 0:
        v27 = -10
    case 9223372036854775807:
        fmt.Println(v27)
    case -669:
        v28 := GenS2{F0: -2147483648}
        _ = v28
    }
    _ = v27
    return 255
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = (((((0.0 * float64(1e10)) + (0.1 - 0.0)) + (1e10 / -2.25)) / ((-2.25 / ((float64(3.14159) * 0.0) + -2.25)) - 3.14159)) / float64((0.1 + -2.25)))
    v3 := []int{-177}
    var v4 int = -216
    v29 := GenS2{F0: 0}
    v30 := GenS2{F0: 0}
    v4 += (v4 * (^(v4)))
    if (255 == 1) {
        var v31 int8 = 1
        v32 := GenS1{F0: -1, F1: v31, F2: 42}
        _ = v31
        _ = v32
    } else {
        var v33 uint32 = 120
        var v34 float64 = -2.25
        _ = v33
        _ = v34
    }
    switch v4 {
    case -1:
        var v35 uint32 = 1
        _ = v35
    case 4294967295:
        var v36 uint32 = 4294967295
        _ = v36
    default:
        for fi37 := 0; fi37 < 2; fi37++ {
            var v38 uint8 = 128
            _ = fi37
            _ = v38
        }
    }
    v39 := GenS2{F0: 330}
    var v40 int8 = -128
    v41 := GenS2{F0: -2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
