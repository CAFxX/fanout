package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
func genFn0(p0 int32) uint16 {
    var v4 uint16 = 42
    v5 := []int{0, 472}
    fmt.Println(p0)
    var v6 uint = 0
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int16) int16 {
    var v7 int = 286
    var v8 uint8 = 255
    fmt.Println(v8, p0)
    switch (v7 & v7) {
    case 884:
        v8 = 140
    default:
        v9 := GenS1{F0: 65535}
        _ = v9
    }
    _ = p0
    _ = v7
    _ = v8
    return p0
}

func main() {
    v1 := make(map[int]string)
    var v2 string = "a"
    var v3 int8 = -117
    if (v3 < v3) {
        if true {
            var v10 uint16 = 861
            for fi11 := 0; fi11 < 6; fi11++ {
                v12 := GenS1{F0: 255}
                _ = fi11
                _ = v12
            }
            v3 -= v3
            _ = v10
        }
        v13 := GenS1{F0: 471}
        v2 = v2
        _ = v13
    } else {
        var v14 float32 = 1e10
        _ = v14
    }
    v2 = "a"
    var v15 uint16 = 1
    for fi16 := 0; fi16 < 5; fi16++ {
        v17 := GenS1{F0: 4294967295}
        fmt.Println(v15, fi16)
        v18 := GenS1{F0: 42}
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 uint8 = 127
    switch 1 {
    case 32768:
        _ = len(v1)
    case 4294967295:
        var v20 int = -1
        _ = v20
    default:
        var v21 float32 = 1.5
        _ = v21
    }
    var v22 int8 = (v3 / (((v3 ^ v3)) | 1))
    v2 = v2
    for fi23 := 0; fi23 < 19; fi23++ {
        v24 := GenS2{F0: fi23}
        _ = fi23
        _ = v24
    }
    var v25 int = int(v3)
    var v26 int = v25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
}
