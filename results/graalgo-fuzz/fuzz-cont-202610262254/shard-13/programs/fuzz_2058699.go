package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int64, p1 uint64) uint16 {
    if true {
        for fi4 := 0; fi4 < 7; fi4++ {
            var v5 float32 = 3.14159
            for fi6 := 0; fi6 < 17; fi6++ {
                var v7 uint16 = 2
                _ = fi6
                _ = v7
            }
            v8 := []int{fi4}
            _ = fi4
            _ = v5
            _ = v8
        }
        switch (("\n\t" + "a") + "foo bar") {
        case "a":
            var v9 bool = (181 <= 255)
            _ = v9
        case "":
            p0 -= 25
        default:
            var v10 bool = (false && false)
            _ = v10
        }
    }
    v11 := GenS1{F0: 184, F1: 0}
    p0 = 0
    _ = p0
    _ = p1
    _ = v11
    return 312
}

func main() {
    var v1 float32 = 0.1
    var v2 float64 = 1e10
    var v3 uint32 = 1
    v12 := []int{int(v3), 32767, -156}
    for fi13 := 0; fi13 < 6; fi13++ {
        v14 := GenS2{F0: 65535}
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: 623}
    if (946 <= 1) {
        if true {
            v3 -= v3
        } else {
            var v16 uint32 = (v3 / (((v3 & v3)) | 1))
            var v17 float64 = v2
            var v18 string = ("hello" + "a")
            _ = v16
            _ = v17
            _ = v18
        }
        v2 = v2
    } else {
        var v19 string = ""
        _ = v19
    }
    v20 := GenS1{F0: 130, F1: 9223372036854775807}
    v21 := GenS1{F0: 2, F1: 128}
    v12 = append(v12, -1)
    var v22 uint64 = 546
    v22 = 0
    var v23 float32 = v1
    var v24 float32 = (float32(v2) * (1e10 - v1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
