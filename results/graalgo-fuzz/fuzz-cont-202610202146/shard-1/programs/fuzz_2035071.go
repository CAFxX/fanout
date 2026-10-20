package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 float64 = 1.5
    v2 := make(map[string]string)
    var v3 uint = 967
    var v4 string = "a"
    v5 := []uint32{323, 65535, 1, 997}
    var v6 uint = 618
    v6 = v6
    var v7 uint64 = 0
    if false {
        if (255 > 210) {
            fmt.Println(v3)
            var v8 uint64 = v7
            _ = v8
        } else {
            v9 := GenS1{F0: 261, F1: 1, F2: 0}
            _ = v9
        }
    } else {
        v7 -= v7
    }
    v1 = (v1 * ((v1 + (v1 * v1)) + (v1 * -2.25)))
    v10 := GenS1{F0: 66, F1: -1, F2: -1}
    _ = len(v2)
    _ = len(v2)
    if ((v7 / ((v7) | 1)) > v7) {
        v1 -= v1
        var v11 bool = false
        var v12 uint32 = 688
        _ = v11
        _ = v12
    }
    var v13 uint64 = v7
    if false {
        var v14 uint16 = 2
        for fi15 := 0; fi15 < 9; fi15++ {
            v16 := GenS2{F0: 128}
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    v17 := GenS1{F0: 32767, F1: -863, F2: 991}
    if true {
        var v18 bool = true
        delete(v2, v4)
        for fi19 := 0; fi19 < 11; fi19++ {
            var v20 uint16 = 0
            _ = fi19
            _ = v20
        }
        _ = v18
    } else {
        var v21 uint = v6
        if true {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = len(v5)
        } else {
            v22 := GenS1{F0: 255, F1: 554, F2: 128}
            _ = v22
        }
        _ = v21
    }
    v1 = float64(1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
