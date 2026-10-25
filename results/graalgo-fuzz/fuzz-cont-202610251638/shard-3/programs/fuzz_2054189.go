package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() int64 {
    v5 := []int{-191, 432}
    if len(v5) > 0 {
        v5[0] = -169
    }
    _ = v5
    return 70
}
func genFn1() int16 {
    var v6 float64 = 0.1
    v6 = (v6 - (v6 * v6))
    switch (("\n\t" + "\n\t") + "a") {
    case "a":
        switch 1 {
        case 846:
            switch ("" + "") {
            case "\n\t":
                v7 := []int{32768, 2147483647}
                _ = v7
            }
        }
    case "":
        var v8 int = 2147483647
        _ = v8
    default:
        v6 = (v6 * (0.1 + v6))
    }
    _ = v6
    return 869
}

func main() {
    var v1 uint64 = 808
    var v2 uint64 = v1
    var v3 string = "a"
    v4 := make(map[uint32]string)
    v3 = ((("foo bar" + "a") + v3) + ("hello" + v3))
    v9 := []int{-105, 405, -1}
    v10 := GenS1{F0: false}
    switch 664 {
    case -2147483648:
        v11 := GenS1{F0: (255 != 42)}
        _ = v11
    case -9223372036854775808:
        var v12 int = -12
        _ = v12
    }
    var v13 int32 = 1
    fmt.Println(v3, v1, v3)
    switch 127 {
    case 4294967295:
        _ = len(v9)
    default:
        if len(v9) > 0 {
            v9[0] = -9223372036854775808
        }
    }
    fmt.Println(v13, v3)
    var v14 int32 = 0
    var v15 float32 = -2.25
    for fi16 := 0; fi16 < 8; fi16++ {
        v4[0] = ((v3 + (v3 + v3)) + ("\n\t" + "\n\t"))
        _ = fi16
    }
    v2 += v1
    var v17 bool = true
    switch "foo bar" {
    case "foo bar":
        for fi18 := 0; fi18 < 11; fi18++ {
            fmt.Println(fi18, v3, v3)
            v9 = append(v9, (fi18 / (((fi18 * 32767)) | 1)))
            _ = fi18
        }
    case "a":
        for fi19 := 0; fi19 < 16; fi19++ {
            fmt.Println(v15)
            v20 := v4[2147483648]
            var v21 bool = ((0.0 == 1e10) || false)
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    v2 = (v1 * v1)
    if v17 {
        v22 := GenS1{F0: v17}
        _ = v22
    }
    var v23 float32 = 1e10
    v24 := GenS1{F0: v17}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v9
}
