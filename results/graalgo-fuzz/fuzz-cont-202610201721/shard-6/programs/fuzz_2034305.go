package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]string)
    var v2 int = -9223372036854775808
    var v3 int = v2
    v4 := []int{0, (^(v3)), (v3 / ((v2) | 1)), (v3 & 128)}
    v5 := []uint32{504, 2147483647}
    var v6 uint = 240
    v2 -= v2
    var v7 int64 = -1
    v6 -= 0
    v2 -= 850
    switch v2 {
    case 2147483647:
        for fi8 := 0; fi8 < 7; fi8++ {
            fmt.Println(v3, fi8)
            for fi9 := 0; fi9 < 18; fi9++ {
                v6 = 203
                _ = fi9
            }
            _ = fi8
        }
    case 32767:
        delete(v1, v3)
    }
    _ = len(v1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v10 int16 = 1000
    if (false && false) {
        var v11 bool = (false && true)
        _ = v11
    } else {
        var v12 int64 = v7
        _ = v12
    }
    fmt.Println(v6, v7)
    if false {
        v2 -= 32768
        var v13 uint = (v6 % (((v6 - v6)) | 1))
        v10 = v10
        _ = v13
    } else {
        v14 := v1[v2]
        _ = v14
    }
    v6 = (v6 - v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
