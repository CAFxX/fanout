package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int32 }
type GenS2 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint64) int32 {
    for fi7 := 0; fi7 < 20; fi7++ {
        p0 = 1
        fmt.Println(p0)
        _ = fi7
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        switch fi8 {
        case -9223372036854775808:
            var v9 string = "\n\t"
            _ = v9
        case 2147483647:
            v10 := []int{(^(fi8)), (fi8 - -149), (fi8 - fi8)}
            _ = v10
        default:
            v11 := GenS1{F0: 2, F1: 4294967295, F2: 1}
            _ = v11
        }
        fmt.Println(p0)
        p0 += p0
        _ = fi8
    }
    if (("hello" + "a") <= "") {
        var v12 uint16 = 1
        _ = v12
    }
    _ = p0
    return 42
}
func genFn1(p0 uint32, p1 int32, p2 int16) uint16 {
    v13 := []int{-410, -9223372036854775808}
    if true {
        var v14 uint8 = 62
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        fmt.Println(p0)
        _ = v14
    } else {
        p2 += p2
        for fi15 := 0; fi15 < 17; fi15++ {
            if (true && (0.0 > (3.14159 / (0.1 + (3.14159 - 3.14159))))) {
                for fi16 := 0; fi16 < 4; fi16++ {
                    var v17 uint16 = 65535
                    _ = fi16
                    _ = v17
                }
                var v18 uint8 = 42
                _ = v18
            }
            _ = fi15
        }
        p0 = 0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 65535
}

func main() {
    var v1 float32 = (0.1 * (-2.25 / -2.25))
    var v2 string = "a"
    var v3 int64 = 42
    var v4 uint64 = 2
    var v5 float64 = -2.25
    v6 := []uint32{1, 0}
    v19 := GenS2{F0: 32767, F1: 254}
    fmt.Println(v3, v5, v3)
    var v20 uint64 = v4
    if (601 >= 19) {
        v2 = v2
    } else {
        var v21 uint32 = 576
        v3 -= (^(v3))
        v22 := GenS1{F0: 200, F1: 4294967295, F2: 0}
        _ = v21
        _ = v22
    }
    var v23 uint = 42
    fmt.Println(v4, v1, v4)
    var v24 uint8 = 127
    switch v2 {
    case "a":
        _ = len(v6)
    default:
        var v25 uint8 = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        var v27 uint64 = (v20 + 42)
        v6 = append(v6, 4294967295)
        var v28 uint16 = 2
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
