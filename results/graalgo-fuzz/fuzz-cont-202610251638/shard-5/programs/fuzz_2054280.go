package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int16; F1 int8 }

func main() {
    var v1 uint16 = 23
    var v2 int8 = 0
    var v3 uint64 = 1
    var v4 string = "a"
    for fi5 := 0; fi5 < 12; fi5++ {
        if (!((441 == 976) || (2 <= 255))) {
            v6 := []int{(fi5 % (((fi5 ^ 127)) | 1)), (fi5 & fi5), fi5, (fi5 >> ((0) & 63))}
            fmt.Println(fi5)
            v6 = append(v6, fi5)
            _ = v6
        }
        _ = fi5
    }
    switch 934 {
    case 4294967295:
        if (!true) {
            v7 := []int{-487, 128}
            switch -307 {
            case 32768:
                if false {
                    switch v4 {
                    case "foo bar":
                        var v8 int16 = -1
                        _ = v8
                    }
                    fmt.Println(v3, v2)
                    var v9 int = -2147483648
                    _ = v9
                }
            }
            _ = v7
        }
    case 693:
        v10 := []int{-531, -655}
        _ = v10
    case 0:
        var v11 uint = 42
        _ = v11
    default:
        var v12 uint32 = 42
        _ = v12
    }
    var v13 int64 = 1
    v14 := GenS1{F0: ("" + v4), F1: v3}
    fmt.Println(v13)
    var v15 int64 = v13
    v2 = v2
    var v16 int = 42
    v3 = (^(v3))
    var v17 uint8 = 255
    if (3.14159 <= 3.14159) {
        var v18 uint8 = 2
        if false {
            v19 := GenS1{F0: "a", F1: (v3 ^ v3)}
            var v20 uint64 = v3
            _ = v19
            _ = v20
        } else {
            v2 -= v2
            for fi21 := 0; fi21 < 20; fi21++ {
                v22 := GenS2{F0: -1, F1: v2}
                _ = fi21
                _ = v22
            }
            v23 := []int{int(v13), (v16 | 127), v16, v16}
            _ = v23
        }
        _ = v18
    } else {
        var v24 float64 = (-2.25 * 0.0)
        v2 = (v2 ^ v2)
        _ = v24
    }
    var v25 uint16 = (v1 % (((^(v1))) | 1))
    var v26 uint64 = 1
    var v27 uint32 = uint32(v25)
    v15 = v13
    v28 := []int{v16}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
