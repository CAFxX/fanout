package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32 }

func main() {
    v1 := []string{"a", "a"}
    v2 := []string{"\n\t", "a", "a"}
    v3 := make(map[uint32]string)
    v4 := GenS1{F0: 127, F1: 1, F2: 2147483647}
    v5 := v3[569]
    if (128 < -721) {
        var v6 float32 = 3.14159
        var v7 int = -374
        _ = v6
        _ = v7
    }
    v3[2] = (v5 + ("hello" + v5))
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 uint8 = 104
        v5 = v5
        _ = fi8
        _ = v9
    }
    if (!false) {
        for fi10 := 0; fi10 < 15; fi10++ {
            v2 = append(v2, "\n\t")
            _ = fi10
        }
    } else {
        v11 := v3[2147483647]
        fmt.Println(v11)
        var v12 string = "\n\t"
        _ = v11
        _ = v12
    }
    switch 128 {
    case 503:
        var v13 float64 = float64(0.1)
        _ = v13
    }
    if ("foo bar" > (v5 + v5)) {
        v5 = (v5 + ((v5 + ("" + "hello")) + v5))
        var v14 float32 = 3.14159
        _ = v14
    } else {
        if (float32(1e10) == 3.14159) {
            v5 = (v5 + v5)
            var v15 string = v5
            _ = v15
        }
        v16 := GenS2{F0: 0}
        v17 := v3[42]
        _ = v16
        _ = v17
    }
    var v18 uint8 = 42
    for fi19 := 0; fi19 < 8; fi19++ {
        var v20 int8 = -1
        v21 := v3[2147483648]
        v20 = v20
        _ = fi19
        _ = v20
        _ = v21
    }
    v22 := GenS2{F0: 0}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch -1 {
    case 345:
        for fi23 := 0; fi23 < 6; fi23++ {
            v18 = v18
            _ = fi23
        }
    case -2147483648:
        if true {
            v5 = ("" + v5)
            fmt.Println(v5, v18)
        }
    default:
        switch -2147483648 {
        case -839:
            _ = len(v3)
        case 126:
            var v24 int8 = -1
            _ = v24
        default:
            if (!true) {
                v25 := GenS3{F0: 168}
                _ = v25
            } else {
                var v26 int16 = -32768
                _ = v26
            }
        }
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        var v28 float32 = 1e10
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 6; fi29++ {
        if true {
            v5 = (v5 + "")
            v30 := GenS4{F0: 824}
            var v31 int8 = 24
            _ = v30
            _ = v31
        }
        var v32 int8 = 1
        _ = fi29
        _ = v32
    }
    v3[0] = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
