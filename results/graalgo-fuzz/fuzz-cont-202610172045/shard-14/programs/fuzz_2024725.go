package main

import (
    "fmt"
)

func genFn0(p0 int8) int64 {
    p0 -= (p0 % (((p0 >> ((664) & 63))) | 1))
    for fi4 := 0; fi4 < 14; fi4++ {
        fmt.Println(fi4)
        _ = fi4
    }
    for fi5 := 0; fi5 < 18; fi5++ {
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 uint32 = 468
            v7 += v7
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    p0 = p0
    _ = p0
    return 1
}

func main() {
    var v1 string = ("foo bar" + "foo bar")
    v2 := make(map[string]uint32)
    var v3 int64 = 1
    v1 = v1
    var v8 int32 = 44
    v8 += 0
    for fi9 := 0; fi9 < 9; fi9++ {
        if false {
            v10 := []int{fi9}
            for fi11 := 0; fi11 < 3; fi11++ {
                v8 -= v8
                if (!true) {
                    fmt.Println(fi11)
                    var v12 uint = 156
                    var v13 uint8 = 134
                    _ = v12
                    _ = v13
                } else {
                    var v14 uint8 = 255
                    v8 += -304
                    _ = v14
                }
                _ = fi11
            }
            _ = v10
        }
        switch fi9 {
        case 127:
            var v15 int = fi9
            _ = v15
        case 876:
            _ = len(v2)
        case 128:
            v8 -= v8
        }
        _ = fi9
    }
    switch 287 {
    case 2147483647:
        fmt.Println(v1, v1)
    case 1:
        switch 0 {
        case -708:
            fmt.Println(v3, v3)
        }
    default:
        var v16 float32 = float32(float64((0.0 * 0.1)))
        _ = v16
    }
    v2[v1] = 2147483647
    v17 := []int{294, 42, 42, 32767}
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 string = "a"
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v8
}
