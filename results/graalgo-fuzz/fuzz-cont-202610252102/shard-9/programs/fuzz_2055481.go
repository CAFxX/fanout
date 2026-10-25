package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int }
type GenS2 struct { F0 string }

func main() {
    var v1 int64 = 42
    v2 := make(map[int]uint32)
    var v3 uint16 = 255
    var v4 uint8 = 42
    v5 := v2[-411]
    v4 -= (v4 & (v4 & (v4 * v4)))
    v3 -= uint16(v5)
    v6 := GenS1{F0: 835, F1: (v3 * v3), F2: 127}
    var v7 string = "\n\t"
    for fi8 := 0; fi8 < 15; fi8++ {
        if (!(0.1 > ((((1.5 + (0.1 - 1e10)) + float64((1.5 / (-2.25 + -2.25)))) * 0.1) / (1.5 + ((((0.0 - 0.1) * 3.14159) - -2.25) / (0.0 + -2.25)))))) {
            if false {
                for fi9 := 0; fi9 < 11; fi9++ {
                    if false {
                        var v10 int32 = -307
                        _ = v10
                    }
                    v11 := GenS2{F0: "a"}
                    _ = fi9
                    _ = v11
                }
                v12 := v2[fi8]
                v13 := []int{}
                _ = v12
                _ = v13
            } else {
                v4 += 127
            }
            if (!(!((!true) || (!(v7 > (("\n\t" + v7) + (((v7 + "\n\t") + "hello") + "hello"))))))) {
                v14 := v2[(fi8 + fi8)]
                var v15 uint = 42
                _ = v14
                _ = v15
            }
        }
        switch v7 {
        case "a":
            var v16 uint64 = 1
            _ = v16
        }
        _ = fi8
    }
    v17 := GenS2{F0: v7}
    v3 = (v3 / (((v3 & uint16(v5))) | 1))
    if (18446744073709551615 >= 797) {
        switch (v7 + "foo bar") {
        case "foo bar":
            delete(v2, -264)
        default:
            if false {
                switch "a" {
                case "foo bar":
                    v18 := []int{1, 127}
                    _ = v18
                case "\n\t":
                    var v19 uint16 = v3
                    _ = v19
                }
                fmt.Println(v3, v5)
            }
        }
        var v20 uint8 = (v4 / ((29) | 1))
        _ = v20
    } else {
        v1 -= v1
        v21 := []int{-51, 127, 42}
        _ = v21
    }
    v22 := []int{531, -826, 671, 42}
    for fi23 := 0; fi23 < 20; fi23++ {
        v1 -= v1
        for fi24 := 0; fi24 < 15; fi24++ {
            v25 := GenS2{F0: v7}
            v7 = v7
            var v26 int8 = 127
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = len(v2)
        _ = fi23
    }
    var v27 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
