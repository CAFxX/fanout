package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint64; F1 int8 }
func genFn0() int8 {
    v5 := []int{-2147483648}
    switch 32768 {
    case 0:
        if false {
            var v6 string = "foo bar"
            _ = v6
        } else {
            var v7 uint32 = 42
            switch -185 {
            case -1:
                for fi8 := 0; fi8 < 9; fi8++ {
                    v9 := GenS1{F0: -1, F1: 32767}
                    switch 4294967295 {
                    case 2147483647:
                        var v10 bool = false
                        _ = v10
                    case -142:
                        v7 = (v7 / ((2147483647) | 1))
                    case 114:
                        fmt.Println(v7)
                    }
                    v11 := GenS1{F0: 0, F1: 32767}
                    _ = fi8
                    _ = v11
                    _ = v9
                }
            default:
                v7 -= v7
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v7
        }
    case -533:
        var v12 string = (("" + "hello") + "")
        _ = v12
    default:
        var v13 uint8 = 128
        _ = v13
    }
    var v14 string = ""
    v15 := GenS1{F0: -1, F1: 42}
    var v16 float64 = -2.25
    _ = v14
    _ = v15
    _ = v16
    _ = v5
    return 42
}

func main() {
    var v1 int16 = 32767
    var v2 int16 = (v1 - v1)
    v3 := make(map[uint32]string)
    var v4 string = "foo bar"
    _ = len(v3)
    v1 += v2
    var v17 string = v4
    var v18 int = -2147483648
    switch v4 {
    case "\n\t":
        v1 = v1
    case "hello":
        fmt.Println(v2, v2)
    case "":
        var v19 float64 = (0.0 + (1.5 * 1.5))
        _ = v19
    }
    v20 := []int{(^(v18)), v18}
    v17 = v17
    v2 += v2
    switch (((v4 + "a") + v4) + v17) {
    case "foo bar":
        v21 := v3[0]
        _ = v21
    case "":
        fmt.Println(v4, v4)
    case "\n\t":
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
    }
    v22 := GenS2{F0: 42, F1: 136, F2: 1}
    var v23 int64 = 847
    switch v17 {
    case "a":
        if len(v20) > 0 {
            v20[0] = 128
        }
    }
    var v24 int32 = -664
    var v25 uint16 = 2
    if true {
        var v26 int32 = v24
        fmt.Println(v23, v25, v1)
        _ = v26
    }
    v24 = 873
    v27 := GenS3{F0: 1, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
}
