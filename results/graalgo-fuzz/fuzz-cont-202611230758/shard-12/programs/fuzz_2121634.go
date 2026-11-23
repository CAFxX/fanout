package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint16; F1 uint16; F2 int16 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 2147483647
    v3 := make(map[int]bool)
    var v4 uint16 = 352
    for fi5 := 0; fi5 < 19; fi5++ {
        v4 = 42
        var v6 int16 = -551
        _ = fi5
        _ = v6
    }
    switch 100 {
    case -543:
        v4 -= (v4 * v4)
    case -215:
        if false {
            var v7 uint = 1
            for fi8 := 0; fi8 < 11; fi8++ {
                v9 := v1[fi8]
                v10 := GenS1{F0: v7, F1: (fi8 / ((fi8) | 1))}
                _ = fi8
                _ = v10
                _ = v9
            }
            var v11 int64 = 673
            _ = v11
            _ = v7
        } else {
            fmt.Println(v4, v2)
            var v12 int16 = -1
            _ = v12
        }
    case -582:
        if true {
            v13 := v3[42]
            var v14 uint16 = v4
            fmt.Println(v4)
            _ = v13
            _ = v14
        } else {
            v15 := GenS2{F0: 835, F1: (v4 | 255), F2: -1}
            v2 += 410
            if true {
                var v16 uint16 = 255
                v16 = v16
                _ = v16
            }
            _ = v15
        }
    }
    fmt.Println(v2, v4)
    var v17 int = 0
    fmt.Println(v4)
    v17 = v17
    var v18 float64 = (-2.25 - 1.5)
    for fi19 := 0; fi19 < 18; fi19++ {
        v4 = 42
        _ = fi19
    }
    var v20 int16 = -1
    fmt.Println(v18)
    var v21 int32 = 1
    if false {
        fmt.Println(v21, v20)
        fmt.Println(v20, v18, v17)
        fmt.Println(v2, v21)
    } else {
        switch "a" {
        case "hello":
            fmt.Println(v18)
        case "a":
            if (-73 >= v21) {
                v21 -= v21
            } else {
                var v22 uint = 0
                for fi23 := 0; fi23 < 16; fi23++ {
                    var v24 float64 = ((v18 / v18) * v18)
                    _ = fi23
                    _ = v24
                }
                _ = v22
            }
        case "\n\t":
            if (!(false && false)) {
                _ = len(v1)
            }
        }
        v25 := v3[(v17 & 127)]
        switch "hello" {
        case "foo bar":
            var v26 string = "\n\t"
            _ = v26
        default:
            fmt.Println(v2, v4, v18)
        }
        _ = v25
    }
    fmt.Println(v20)
    _ = len(v1)
    if ((v21 % (((v21 ^ v21)) | 1)) > v21) {
        v27 := []int{v17, (v17 >> ((4294967295) & 63)), v17}
        v28 := GenS1{F0: 42, F1: (v17 | (v17 - (v17 + 782)))}
        var v29 uint32 = v2
        _ = v27
        _ = v28
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
