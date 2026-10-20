package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int8 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int32) int8 {
    for fi7 := 0; fi7 < 15; fi7++ {
        p0 += (p0 - -357)
        _ = fi7
    }
    p0 = (p0 ^ (^(p0)))
    if true {
        v8 := []int{32767, -1}
        _ = v8
    }
    var v9 uint16 = 0
    _ = p0
    _ = v9
    return 100
}

func main() {
    var v1 int8 = -1
    v2 := make(map[string]bool)
    v3 := []uint32{0, 401, 1, 212}
    v4 := []uint32{}
    v5 := make(map[int]bool)
    var v6 string = "\n\t"
    var v10 uint16 = 609
    var v11 int = 32767
    switch v6 {
    case "":
        var v12 uint = 945
        _ = v12
    default:
        var v13 int64 = 9223372036854775807
        _ = v13
    }
    var v14 string = v6
    var v15 int8 = -1
    var v16 uint16 = 1
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 int8 = v1
        _ = fi17
        _ = v18
    }
    if true {
        var v19 uint8 = 42
        _ = v19
    }
    var v20 int8 = int8(v10)
    v5[(^(v11))] = true
    var v21 int64 = 644
    if false {
        fmt.Println(v20, v1)
        var v22 float32 = 1.5
        _ = v22
    } else {
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 bool = true
            _ = fi23
            _ = v24
        }
    }
    if false {
        for fi25 := 0; fi25 < 17; fi25++ {
            for fi26 := 0; fi26 < 13; fi26++ {
                if (42 >= 42) {
                    fmt.Println(v11)
                    v27 := GenS1{F0: 255, F1: 0, F2: 106}
                    _ = v27
                } else {
                    _ = len(v5)
                    switch v11 {
                    case -249:
                        var v28 uint = 255
                        _ = v28
                    case 9223372036854775807:
                        v10 += (v10 & (v10 | v16))
                    }
                    v6 = (((v14 + v14) + v6) + v14)
                }
                fmt.Println(v21, v20)
                fmt.Println(fi25)
                _ = fi26
            }
            _ = fi25
        }
        v21 = v21
        v29 := GenS2{F0: -2147483648}
        _ = v29
    }
    v30 := GenS1{F0: 74, F1: -478, F2: 95}
    for fi31 := 0; fi31 < 9; fi31++ {
        v32 := v5[v11]
        var v33 float32 = 0.1
        _ = fi31
        _ = v32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
