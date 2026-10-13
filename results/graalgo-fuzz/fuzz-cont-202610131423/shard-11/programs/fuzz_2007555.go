package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]string)
    var v3 float64 = ((((0.0 + -2.25) - ((1e10 / 1e10) / 1.0)) + 1e10) * (0.1 / 1e10))
    v3 = float64((0.0 - -2.25))
    v3 = ((v3 / (v3 - (1.5 * v3))) * 1e10)
    fmt.Println(v3)
    v3 = float64(0.1)
    v3 = (-2.25 / v3)
    v4 := GenS1{F0: -1, F1: 128, F2: 2}
    var v5 int = 597
    for fi6 := 0; fi6 < 4; fi6++ {
        v3 += float64((((0.1 - float32(v3)) + 1e10) - 0.0))
        if false {
            v7 := GenS1{F0: -276, F1: 139, F2: 0}
            for fi8 := 0; fi8 < 17; fi8++ {
                if true {
                    v3 -= v3
                }
                var v9 bool = false
                v10 := []int{(fi6 * (^(v5)))}
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = v7
        } else {
            var v11 bool = ("\n\t" == "\n\t")
            var v12 int8 = 1
            v13 := []int{42, v5}
            _ = v11
            _ = v12
            _ = v13
        }
        var v14 uint = 255
        _ = fi6
        _ = v14
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 uint8 = 42
        fmt.Println(fi15)
        v17 := GenS1{F0: -1, F1: (v16 | v16), F2: 2}
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        switch (("" + "hello") + "") {
        case "":
            var v19 uint8 = 40
            _ = v19
        default:
            _ = len(v1)
        }
        _ = fi18
    }
    var v20 int16 = -149
    v21 := []int{}
    fmt.Println(v5, v3)
    for fi22 := 0; fi22 < 19; fi22++ {
        v5 += (fi22 * 546)
        v1[65535] = ((v3 == ((v3 + (1.5 + v3)) + 1e10)) && (false && (v5 >= (^(fi22)))))
        v5 -= v5
        _ = fi22
    }
    var v23 int16 = -1
    if len(v21) > 0 {
        v21[0] = (v5 << ((0) & 63))
    }
    var v24 float64 = (1.5 - (v3 * v3))
    var v25 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
