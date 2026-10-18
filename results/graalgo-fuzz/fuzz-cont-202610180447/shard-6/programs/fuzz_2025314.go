package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int16 = 414
    var v2 uint64 = 4294967295
    v3 := []uint32{2, 4294967295, 2147483648}
    var v4 float64 = 3.14159
    var v5 float64 = float64(((3.14159 - (((3.14159 + 1e10) / 1.0) / float32(v4))) / 1.5))
    var v6 uint = 768
    v7 := GenS1{F0: uint16(v2)}
    if false {
        var v8 int16 = v1
        _ = v8
    }
    v9 := GenS1{F0: 909}
    fmt.Println(v4, v1)
    var v10 int32 = 1
    var v11 string = "\n\t"
    switch 32767 {
    case 0:
        var v12 int8 = 100
        _ = v12
    default:
        fmt.Println(v5, v6)
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        fmt.Println(v10, v4)
        var v14 uint = v6
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        v3 = append(v3, 68)
        _ = fi15
    }
    var v16 bool = (255 <= 255)
    var v17 string = v11
    var v18 int64 = 826
    var v19 float64 = v5
    if true {
        fmt.Println(v19, v6)
        if (v1 <= -1) {
            _ = len(v3)
            for fi20 := 0; fi20 < 5; fi20++ {
                v21 := GenS1{F0: 0}
                if len(v3) > 0 {
                    v3[0] = 65535
                }
                _ = fi20
                _ = v21
            }
            v19 = v4
        }
    } else {
        v2 += 966
        if len(v3) > 0 {
            v3[0] = 2147483648
        }
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        var v23 float64 = v5
        switch (^(fi22)) {
        case 32767:
            var v24 float32 = float32((v5 / (v5 - (((float64((1e10 + (-2.25 - 0.1))) + float64(1e10)) + (v4 - (-2.25 - (float64(1.5) + (v19 * v23))))) / (v5 * v19)))))
            _ = v24
        case 1:
            if true {
                if true {
                    v25 := GenS1{F0: 255}
                    _ = v25
                }
            } else {
                for fi26 := 0; fi26 < 20; fi26++ {
                    var v27 uint8 = 224
                    _ = fi26
                    _ = v27
                }
                var v28 int = -43
                _ = v28
            }
        default:
            var v29 int64 = v18
            _ = v29
        }
        _ = fi22
        _ = v23
    }
    fmt.Println(v17)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
