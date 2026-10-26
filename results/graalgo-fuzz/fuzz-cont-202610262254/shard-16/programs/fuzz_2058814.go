package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 uint8 = 255
    v2 := make(map[int]int)
    v3 := make(map[string]uint32)
    v4 := []uint32{2, 962, 183}
    var v5 uint8 = (v1 >> ((0) & 63))
    var v6 float64 = 1e10
    v7 := v2[989]
    v6 = 1e10
    var v8 uint8 = 1
    switch v7 {
    case 344:
        for fi9 := 0; fi9 < 15; fi9++ {
            delete(v2, (v7 | (fi9 / (((^(fi9))) | 1))))
            if ((0.1 / v6) != float64(0.1)) {
                if true {
                    v5 = v1
                    fmt.Println(fi9)
                }
            } else {
                switch v7 {
                case 0:
                    var v10 int16 = 1000
                    _ = v10
                case -519:
                    for fi11 := 0; fi11 < 12; fi11++ {
                        var v12 int64 = 428
                        fmt.Println(fi11)
                        _ = fi11
                        _ = v12
                    }
                default:
                    v5 -= 2
                }
                _ = len(v4)
            }
            for fi13 := 0; fi13 < 6; fi13++ {
                var v14 int16 = -752
                _ = fi13
                _ = v14
            }
            _ = fi9
        }
    case -9223372036854775808:
        var v15 int64 = -142
        _ = v15
    case -31:
        v16 := GenS1{F0: -357}
        _ = v16
    }
    delete(v2, v7)
    var v17 uint8 = v8
    var v18 bool = (65535 == 721)
    var v19 uint8 = (v5 / ((v8) | 1))
    var v20 int32 = 1
    v21 := v2[1]
    delete(v2, 32768)
    for fi22 := 0; fi22 < 8; fi22++ {
        v18 = (!v18)
        _ = fi22
    }
    v1 = (v19 + 206)
    var v23 uint = 702
    var v24 uint16 = 773
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
