package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint }
func genFn0(p0 int32) int16 {
    if false {
        var v5 uint8 = 0
        var v6 int16 = -32768
        _ = v5
        _ = v6
    }
    if false {
        var v7 bool = true
        for fi8 := 0; fi8 < 6; fi8++ {
            p0 -= 1
            var v9 int8 = 19
            var v10 int32 = p0
            _ = fi8
            _ = v10
            _ = v9
        }
        if (false && v7) {
            var v11 int16 = 32767
            _ = v11
        }
        _ = v7
    }
    var v12 uint = 230
    var v13 uint16 = 843
    p0 -= p0
    _ = p0
    _ = v12
    _ = v13
    return int16(v13)
}
func genFn1(p0 string, p1 uint64) bool {
    switch "hello" {
    case "foo bar":
        p0 = ("foo bar" + "foo bar")
    case "hello":
        var v14 float64 = float64(3.14159)
        _ = v14
    case "":
        var v15 uint16 = 1
        _ = v15
    default:
        var v16 uint8 = 2
        _ = v16
    }
    v17 := []int{-965, -942}
    p1 += 42
    var v18 uint8 = uint8(p1)
    for fi19 := 0; fi19 < 8; fi19++ {
        v18 += uint8(p1)
        v17 = append(v17, (^(fi19)))
        _ = fi19
    }
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return true
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 0
    v3 := make(map[string]int)
    var v4 int16 = -797
    var v20 uint64 = 898
    var v21 uint64 = v20
    v21 = v20
    var v22 int = 128
    var v23 bool = (!false)
    v4 -= genFn0(-2147483648)
    delete(v1, "a")
    for fi24 := 0; fi24 < 4; fi24++ {
        v20 = (v21 << (((^(v2))) & 63))
        v23 = v23
        _ = fi24
    }
    var v25 uint16 = 2
    _ = len(v1)
    v26 := GenS1{F0: 148, F1: v2, F2: v2}
    v27 := []int{773, v22}
    if false {
        if len(v27) >= 2 {
            v27 = v27[1:len(v27)]
        }
        var v28 int16 = (v4 >> (((v2 & uint(v25))) & 63))
        v29 := GenS1{F0: 32767, F1: 0, F2: (v2 << ((v2) & 63))}
        _ = v28
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
