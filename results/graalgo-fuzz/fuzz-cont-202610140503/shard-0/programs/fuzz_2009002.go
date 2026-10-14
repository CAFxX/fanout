package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int64; F2 uint16 }
func genFn0(p0 bool) int64 {
    var v4 uint8 = 128
    p0 = (4294967295 >= 0)
    _ = p0
    _ = v4
    return -1
}

func main() {
    v1 := make(map[int]string)
    var v2 int = 32767
    v3 := make(map[int]bool)
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 uint = 65535
        _ = fi5
        _ = v6
    }
    var v7 uint64 = 937
    if true {
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 bool = true
            _ = fi8
            _ = v9
        }
        fmt.Println(v7, v7)
        v1[(v2 * v2)] = ""
    } else {
        if (!(1.5 < 0.1)) {
            for fi10 := 0; fi10 < 14; fi10++ {
                var v11 float64 = 0.0
                _ = fi10
                _ = v11
            }
            var v12 uint = 298
            var v13 int32 = int32(v12)
            _ = v12
            _ = v13
        }
        v14 := []int{v2}
        var v15 uint32 = 65535
        _ = v14
        _ = v15
    }
    v7 += v7
    var v16 int64 = 1
    for fi17 := 0; fi17 < 11; fi17++ {
        if false {
            v18 := GenS1{F0: 1000}
            v7 += 2
            _ = v18
        }
        if false {
            var v19 bool = true
            for fi20 := 0; fi20 < 7; fi20++ {
                v21 := GenS2{F0: -1, F1: v16, F2: 1}
                _ = fi20
                _ = v21
            }
            _ = v19
        }
        _ = fi17
    }
    fmt.Println(v2, v7)
    var v22 int32 = -283
    v7 = (v7 * v7)
    v23 := []int{(v2 & v2)}
    var v24 float64 = float64(((-2.25 / 1.5) + 3.14159))
    v24 = v24
    v22 = 2147483647
    var v25 int64 = genFn0(((^(v16)) < v16))
    for fi26 := 0; fi26 < 14; fi26++ {
        var v27 bool = false
        v28 := GenS1{F0: 1}
        _ = fi26
        _ = v27
        _ = v28
    }
    v2 = (v2 & v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v7
}
