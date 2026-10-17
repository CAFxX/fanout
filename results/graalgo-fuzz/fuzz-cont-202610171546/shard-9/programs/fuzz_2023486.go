package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int64 }
func genFn0(p0 int) int64 {
    fmt.Println(p0)
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint8 = 128
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        v10 := GenS1{F0: 65535, F1: 255, F2: -114}
        var v11 int32 = -1
        if (p0 < -68) {
            if (!(776 < 306)) {
                p0 += (fi9 | (^(fi9)))
            } else {
                p0 = 4294967295
                v12 := GenS1{F0: 65535, F1: 42, F2: 42}
                _ = v12
            }
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        switch fi13 {
        case 0:
            var v14 bool = (false && (!(0.1 > 0.1)))
            _ = v14
        case -1:
            p0 = (p0 + p0)
        case 4294967295:
            var v15 uint16 = 373
            _ = v15
        default:
            for fi16 := 0; fi16 < 3; fi16++ {
                p0 -= (p0 + (p0 >> ((1) & 63)))
                _ = fi16
            }
        }
        switch "a" {
        case "hello":
            v17 := GenS1{F0: 2, F1: 0, F2: -950}
            _ = v17
        }
        _ = fi13
    }
    _ = p0
    return -1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]uint32)
    var v3 int64 = 42
    v4 := make(map[uint32]uint32)
    var v5 string = "hello"
    v6 := make(map[uint32]string)
    fmt.Println(v3, v3)
    v18 := GenS1{F0: 42, F1: 355, F2: -944}
    v3 = v3
    var v19 float64 = (1.5 / -2.25)
    var v20 uint16 = 32767
    fmt.Println(v3, v3, v19)
    var v21 uint32 = 675
    v22 := []int{32768}
    for fi23 := 0; fi23 < 15; fi23++ {
        v3 = genFn0(fi23)
        _ = fi23
    }
    fmt.Println(v20, v20)
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    var v24 int16 = 1
    v22 = append(v22, 128)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
