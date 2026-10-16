package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint64; F1 bool; F2 uint16 }
func genFn0(p0 int8, p1 uint16) int32 {
    p1 = (p1 + p1)
    v6 := []int{191}
    _ = p0
    _ = p1
    _ = v6
    return 0
}

func main() {
    var v1 uint64 = 905
    var v2 uint = 4294967295
    var v3 uint32 = 2147483648
    var v4 int32 = 2147483647
    v5 := make(map[int]string)
    for fi7 := 0; fi7 < 16; fi7++ {
        switch fi7 {
        case 783:
            var v8 float32 = ((((((0.0 / (-2.25 / (3.14159 * -2.25))) / 3.14159) * (1e10 / (3.14159 - 0.1))) + (0.0 - ((1e10 - 3.14159) / float32(1e10)))) / ((-2.25 + ((1e10 * 1.5) / (0.1 * 0.1))) * 1.5)) / float32(1e10))
            _ = v8
        default:
            v3 -= v3
        }
        _ = fi7
    }
    var v9 uint = (v2 >> (((v2 ^ 741)) & 63))
    v1 = 0
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 float64 = 0.0
            v13 := []int{fi11, (fi11 / (((fi11 + fi10)) | 1))}
            var v14 uint8 = 191
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        delete(v5, (fi10 << (((v2 - (v9 % (((v2 / ((v2) | 1))) | 1)))) & 63)))
        var v15 uint = (v9 << (((v9 << (((v9 * v9)) & 63))) & 63))
        _ = fi10
        _ = v15
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        v4 += -1
        for fi17 := 0; fi17 < 15; fi17++ {
            if ("hello" <= "") {
                var v18 string = "\n\t"
                _ = v18
            }
            var v19 float32 = ((-2.25 * ((-2.25 + ((1.5 + 1.5) + 0.1)) + 3.14159)) / 1.0)
            v20 := GenS1{F0: uint16(v9), F1: 0}
            v2 = (v9 | v9)
            v21 := GenS2{F0: (v1 ^ 0), F1: false, F2: 265}
            _ = fi17
            _ = v19
            _ = v20
            _ = v21
        }
        delete(v5, fi16)
        _ = fi16
    }
    fmt.Println(v3, v3, v9)
    fmt.Println(v9)
    v3 = 1
    for fi22 := 0; fi22 < 16; fi22++ {
        var v23 uint8 = 155
        if true {
            if false {
                var v24 uint32 = v3
                _ = v24
            }
        }
        if false {
            fmt.Println(v2, v3)
            fmt.Println(v9, v1, fi22)
            var v25 int16 = 318
            _ = v25
        }
        _ = fi22
        _ = v23
    }
    v9 += v2
    if (false && true) {
        v26 := GenS1{F0: 1, F1: 255}
        var v27 bool = (((0 == int8(v1)) || false) || (v3 <= uint32(v1)))
        v1 = (^(v1))
        _ = v26
        _ = v27
    } else {
        v28 := v5[821]
        v29 := v5[1]
        _ = v28
        _ = v29
    }
    fmt.Println(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
