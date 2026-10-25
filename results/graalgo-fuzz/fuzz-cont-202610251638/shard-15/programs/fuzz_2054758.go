package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint16 }
func genFn0() int32 {
    for fi4 := 0; fi4 < 16; fi4++ {
        for fi5 := 0; fi5 < 11; fi5++ {
            fmt.Println(fi5, fi5)
            fmt.Println(fi4)
            _ = fi5
        }
        _ = fi4
    }
    var v6 uint64 = 326
    for fi7 := 0; fi7 < 1; fi7++ {
        v6 = v6
        var v8 int64 = 1
        v6 -= 0
        _ = fi7
        _ = v8
    }
    if ((-2.25 / 0.1) == 0.1) {
        var v9 int16 = 42
        var v10 int64 = 638
        _ = v10
        _ = v9
    } else {
        v11 := GenS1{F0: 42, F1: 603}
        _ = v11
    }
    _ = v6
    return 2147483647
}
func genFn1() int {
    v12 := []int{}
    var v13 uint = 517
    v14 := GenS1{F0: 1000, F1: 133}
    _ = v12
    _ = v13
    _ = v14
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 1
    v3 := make(map[uint32]int)
    fmt.Println(v2)
    var v15 uint64 = 709
    v16 := []int{0, genFn1(), 2147483647}
    var v17 uint = 4294967295
    v2 -= (v2 % ((6) | 1))
    if (!true) {
        v17 = (v17 ^ (v17 >> ((v17) & 63)))
    }
    v18 := GenS2{F0: 2, F1: 42, F2: 208}
    if false {
        switch -2147483648 {
        case 32768:
            for fi19 := 0; fi19 < 4; fi19++ {
                if (((v2 % (((v2 | v2)) | 1)) == v2) || (404 > 42)) {
                    for fi20 := 0; fi20 < 1; fi20++ {
                        var v21 uint = (v17 | (v17 ^ v17))
                        v22 := GenS2{F0: 127, F1: 255, F2: 242}
                        v23 := GenS1{F0: 42, F1: 1}
                        fmt.Println(v2, fi20, v17)
                        _ = fi20
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    v15 = v15
                }
                fmt.Println(fi19)
                _ = fi19
            }
        }
        v24 := GenS2{F0: 127, F1: 42, F2: 32767}
        v2 = genFn0()
        _ = v24
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
