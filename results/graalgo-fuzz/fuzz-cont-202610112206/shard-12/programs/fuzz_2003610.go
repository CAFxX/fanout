package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int }

func main() {
    var v1 bool = ((0.0 - 1.5) != ((float64(-2.25) + 1e10) - float64((float32((1e10 + (((-2.25 / 0.1) + ((-2.25 + 0.0) - 1e10)) / -2.25))) * float32(1e10)))))
    v2 := []uint32{2}
    var v3 int16 = -1
    v4 := []string{}
    var v5 int64 = -1
    v6 := make(map[uint32]uint32)
    switch "hello" {
    case "a":
        if (42 != 0) {
            if len(v2) > 0 {
                v2[0] = 42
            }
            v1 = true
        } else {
            switch 463 {
            case 127:
                fmt.Println(v5)
            case -336:
                switch -148 {
                case 480:
                    v5 = v5
                default:
                    var v7 uint8 = 255
                    _ = v7
                }
            case -137:
                v8 := GenS1{F0: 349}
                _ = v8
            }
            v1 = true
            if v1 {
                var v9 float32 = 0.0
                var v10 uint = 65535
                _ = v10
                _ = v9
            }
        }
    case "foo bar":
        var v11 int = 32767
        _ = v11
    case "hello":
        for fi12 := 0; fi12 < 17; fi12++ {
            delete(v6, 4294967295)
            _ = fi12
        }
    default:
        switch "a" {
        case "\n\t":
            var v13 uint = 255
            _ = v13
        case "a":
            v5 = v5
        case "hello":
            fmt.Println(v1)
        }
    }
    if len(v4) > 0 {
        v4[0] = ((("" + "\n\t") + ("a" + "foo bar")) + "hello")
    }
    if v1 {
        var v14 int16 = v3
        v3 = 375
        v5 += v5
        _ = v14
    } else {
        fmt.Println(v1)
        var v15 bool = v1
        _ = v15
    }
    v16 := GenS1{F0: 2147483647}
    v6[0] = 1
    switch -914 {
    case 2147483647:
        v17 := v6[1]
        _ = v17
    default:
        var v18 int64 = (v5 * (v5 & -815))
        _ = v18
    }
    var v19 int64 = v5
    var v20 uint = 42
    v21 := GenS1{F0: 1}
    for fi22 := 0; fi22 < 11; fi22++ {
        switch (fi22 + fi22) {
        case 127:
            var v23 bool = v1
            _ = v23
        default:
            v24 := GenS1{F0: -815}
            _ = v24
        }
        if v1 {
            var v25 float64 = ((float64(((3.14159 / (float32(1.5) * ((1e10 + -2.25) / 1.5))) / 3.14159)) / -2.25) - (float64((-2.25 + 0.1)) / 0.1))
            v26 := GenS2{F0: fi22, F1: (^(fi22))}
            v2 = append(v2, 42)
            _ = v25
            _ = v26
        } else {
            v1 = (89 > int8(fi22))
            _ = len(v4)
            v19 += v5
        }
        if v1 {
            for fi27 := 0; fi27 < 9; fi27++ {
                var v28 uint64 = 530
                v29 := GenS1{F0: 2147483647}
                _ = fi27
                _ = v28
                _ = v29
            }
            v20 -= 1
            _ = len(v2)
        } else {
            var v30 int8 = -79
            v31 := GenS1{F0: 0}
            _ = v30
            _ = v31
        }
        _ = fi22
    }
    v1 = (v20 > (v20 << (((v20 << (((v20 - (v20 | (^(v20))))) & 63))) & 63)))
    v20 += (^(v20))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
