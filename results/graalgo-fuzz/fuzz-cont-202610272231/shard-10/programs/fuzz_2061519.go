package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int64 }

func main() {
    v1 := []int{-1, 128, 1}
    var v2 int32 = -1
    var v3 float64 = ((3.14159 - 1e10) - 0.1)
    v4 := make(map[uint32]int)
    var v5 int8 = 42
    v6 := GenS1{F0: v5, F1: "foo bar"}
    var v7 uint16 = 804
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            v10 := GenS2{F0: 778}
            _ = fi9
            _ = v10
        }
        v3 = float64(((((-2.25 - 0.1) * ((float32(float64(float32((((v3 * 0.0) * -2.25) / -2.25)))) * ((float32((float64(3.14159) + 0.1)) + 0.0) / 3.14159)) - (0.1 + 0.0))) - 3.14159) * (3.14159 + 0.0)))
        v11 := GenS2{F0: 9223372036854775807}
        _ = fi8
        _ = v11
    }
    var v12 int32 = 0
    delete(v4, 2147483648)
    for fi13 := 0; fi13 < 20; fi13++ {
        fmt.Println(fi13, v5, v2)
        _ = fi13
    }
    v2 += 234
    switch 546 {
    case 2147483647:
        v3 += (v3 - v3)
    case 127:
        if false {
            _ = len(v1)
        } else {
            if (-2.25 < (v3 / ((1e10 / 1.5) - v3))) {
                var v14 uint = 255
                v4[111] = 1
                _ = v14
            } else {
                v7 -= 503
                v5 = v5
                if true {
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                    v15 := GenS2{F0: -672}
                    fmt.Println(v12)
                    _ = v15
                }
            }
            var v16 int64 = 1
            if true {
                v5 = (v5 * v5)
                fmt.Println(v16, v2)
            }
            _ = v16
        }
    case -195:
        switch 0 {
        case 127:
            v17 := GenS1{F0: (v5 | v5), F1: "hello"}
            _ = v17
        case 32768:
            for fi18 := 0; fi18 < 10; fi18++ {
                var v19 float64 = v3
                _ = fi18
                _ = v19
            }
        case 597:
            switch 128 {
            case -815:
                var v20 int32 = v2
                _ = v20
            case 0:
                var v21 uint16 = v7
                _ = v21
            }
        default:
            if ((0.1 < (0.0 / 0.1)) && true) {
                fmt.Println(v7)
            }
        }
    default:
        fmt.Println(v3)
    }
    v7 = v7
    v22 := GenS1{F0: 0, F1: ""}
    fmt.Println(v5)
    var v23 string = ""
    var v24 float64 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
