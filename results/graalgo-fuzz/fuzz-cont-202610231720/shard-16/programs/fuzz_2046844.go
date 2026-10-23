package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 uint16 }
func genFn0(p0 bool, p1 int16, p2 int64) bool {
    switch 4294967295 {
    case 406:
        v4 := GenS1{F0: (p2 ^ p2)}
        _ = v4
    case 32767:
        if (!p0) {
            v5 := []int{-1, 128}
            var v6 int8 = 74
            _ = v5
            _ = v6
        }
    }
    v7 := []int{-135, -1, -796, 127}
    v8 := GenS2{F0: 100}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p0
}

func main() {
    var v1 string = ""
    var v2 int = 2147483647
    var v3 int8 = 100
    if (127 == -2147483648) {
        v2 += v2
        var v9 uint = 65535
        for fi10 := 0; fi10 < 7; fi10++ {
            switch fi10 {
            case 128:
                v11 := GenS3{F0: false, F1: 1}
                _ = v11
            case 4294967295:
                v12 := []int{}
                _ = v12
            }
            if (!(765 > 65535)) {
                var v13 uint32 = 323
                var v14 bool = false
                v3 = v3
                _ = v13
                _ = v14
            } else {
                var v15 int = (fi10 & 221)
                v16 := []int{fi10, (fi10 << (((v9 ^ v9)) & 63))}
                fmt.Println(v9)
                _ = v15
                _ = v16
            }
            var v17 int32 = 0
            _ = fi10
            _ = v17
        }
        _ = v9
    }
    if true {
        v1 = "\n\t"
        fmt.Println(v2, v3, v3)
    }
    switch v1 {
    case "a":
        var v18 int32 = 2147483647
        _ = v18
    case "hello":
        var v19 int8 = 42
        _ = v19
    default:
        v20 := []int{(v2 - 207)}
        _ = v20
    }
    var v21 string = ((v1 + (v1 + ((v1 + v1) + v1))) + "\n\t")
    fmt.Println(v2, v2)
    var v22 float32 = 0.0
    v23 := []int{127, (^(v2)), v2}
    var v24 uint = 964
    v3 += (-(v3))
    var v25 int64 = 1
    fmt.Println(v2)
    for fi26 := 0; fi26 < 8; fi26++ {
        v23 = append(v23, v2)
        for fi27 := 0; fi27 < 8; fi27++ {
            var v28 float64 = ((((0.1 * 0.1) / (1e10 - ((0.0 * 1.5) + ((float64((0.1 + 1e10)) - 3.14159) * 0.0)))) * 1e10) + 0.1)
            _ = fi27
            _ = v28
        }
        fmt.Println(v3, v24, v21)
        _ = fi26
    }
    v21 = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
