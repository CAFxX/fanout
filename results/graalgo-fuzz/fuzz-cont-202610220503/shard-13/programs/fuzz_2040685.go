package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int8) int64 {
    v6 := []int{-1, -22, 0}
    var v7 uint = 18446744073709551615
    if true {
        for fi8 := 0; fi8 < 3; fi8++ {
            for fi9 := 0; fi9 < 7; fi9++ {
                fmt.Println(fi8)
                var v10 float64 = ((0.1 * 3.14159) / 1.0)
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    }
    _ = p0
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 int16, p1 int64) int64 {
    p0 = p0
    if (1e10 < float64(0.0)) {
        v11 := []int{}
        _ = v11
    }
    var v12 int64 = p1
    v13 := GenS1{F0: p0}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return v12
}

func main() {
    v1 := []int{32767}
    var v2 int = 1
    var v3 int8 = 1
    v4 := make(map[int]string)
    var v5 int16 = 1
    var v14 uint8 = 42
    if false {
        var v15 int = 32768
        var v16 uint32 = 0
        _ = v15
        _ = v16
    }
    if (353 <= 4294967295) {
        var v17 uint16 = 405
        _ = v17
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 float64 = 3.14159
        fmt.Println(fi18)
        for fi20 := 0; fi20 < 5; fi20++ {
            var v21 uint16 = 32767
            _ = fi20
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    var v22 bool = (547 <= 65535)
    v2 -= v2
    var v23 string = "foo bar"
    for fi24 := 0; fi24 < 8; fi24++ {
        fmt.Println(v23, v3, v23)
        _ = fi24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
