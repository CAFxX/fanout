package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint16 }
type GenS2 struct { F0 int64; F1 int; F2 uint8 }

func main() {
    var v1 float64 = 1.5
    var v2 int32 = -98
    var v3 int64 = -1
    var v4 string = ""
    v5 := make(map[int]bool)
    var v6 uint8 = 2
    v7 := []int{-856, 2147483647}
    v5[128] = false
    v3 += (^(v3))
    v3 = v3
    if false {
        if (true && (((0.0 - -2.25) + v1) <= v1)) {
            v6 = 42
        }
        if true {
            v4 = ""
            v8 := v5[463]
            _ = v8
        } else {
            for fi9 := 0; fi9 < 5; fi9++ {
                for fi10 := 0; fi10 < 10; fi10++ {
                    var v11 int = fi9
                    v2 = v2
                    v12 := GenS1{F0: 0, F1: "hello", F2: 337}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                if (103 == int8(v6)) {
                    v4 = ("" + ("\n\t" + v4))
                    var v13 uint32 = 0
                    _ = v13
                }
                _ = fi9
            }
        }
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 uint64 = 54
            _ = fi14
            _ = v15
        }
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v1 = v1
    if false {
        var v16 float32 = -2.25
        var v17 int8 = 100
        var v18 int8 = 1
        _ = v16
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: 0, F1: v4, F2: 486}
    var v20 uint64 = 359
    if true {
        v21 := GenS2{F0: (v3 * (v3 / ((v3) | 1))), F1: -609, F2: v6}
        var v22 uint64 = (v20 << ((202) & 63))
        v2 -= 593
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        for fi24 := 0; fi24 < 14; fi24++ {
            switch fi24 {
            case -813:
                fmt.Println(fi23, v20, v2)
            case -809:
                v2 = v2
            default:
                _ = len(v5)
            }
            fmt.Println(v6)
            _ = fi24
        }
        _ = fi23
    }
    var v25 string = ("\n\t" + v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
