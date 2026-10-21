package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 bool; F2 uint8 }
type GenS3 struct { F0 uint16 }
func genFn0() bool {
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS1{F0: 127}
        var v8 bool = (300 > -385)
        _ = fi6
        _ = v7
        _ = v8
    }
    switch 657 {
    case 766:
        v9 := []int{9223372036854775807}
        _ = v9
    case 9223372036854775807:
        var v10 uint32 = 65535
        _ = v10
    case 32767:
        var v11 float64 = 1e10
        _ = v11
    }
    var v12 int16 = 42
    _ = v12
    return false
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 55
    var v3 float32 = -2.25
    var v4 uint16 = 2
    v5 := []uint32{65535, 65535}
    switch "\n\t" {
    case "":
        fmt.Println(v2, v3)
    case "hello":
        for fi13 := 0; fi13 < 20; fi13++ {
            v14 := GenS1{F0: 127}
            _ = fi13
            _ = v14
        }
    case "foo bar":
        v15 := GenS2{F0: (-(v2)), F1: true, F2: uint8(v4)}
        _ = v15
    default:
        if (414 < 255) {
            var v16 uint16 = (v4 | v4)
            _ = v16
        } else {
            var v17 uint32 = 483
            _ = v17
        }
    }
    v3 = 3.14159
    var v18 uint8 = uint8(v4)
    for fi19 := 0; fi19 < 2; fi19++ {
        v20 := GenS3{F0: 2}
        var v21 int = 128
        v22 := v1[(("" + "a") + "\n\t")]
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        var v24 int64 = -9223372036854775808
        _ = fi23
        _ = v24
    }
    _ = len(v5)
    var v25 int16 = -1
    var v26 int64 = 0
    var v27 int64 = (v26 | v26)
    var v28 bool = (2 == 4294967295)
    switch 4294967295 {
    case -2147483648:
        var v29 uint32 = 0
        _ = v29
    case 9223372036854775807:
        for fi30 := 0; fi30 < 20; fi30++ {
            fmt.Println(v18, fi30)
            _ = len(v5)
            _ = fi30
        }
    case -9223372036854775808:
        var v31 int64 = v26
        _ = v31
    }
    fmt.Println(v25, v4)
    v18 = v18
    if (224 >= 887) {
        fmt.Println(v3, v27, v18)
        fmt.Println(v18)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
