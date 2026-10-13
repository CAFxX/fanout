package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int16) int32 {
    p0 = 42
    var v7 int64 = -9223372036854775808
    _ = p0
    _ = v7
    return -1
}

func main() {
    var v1 bool = true
    var v2 int16 = -860
    v3 := []uint32{65535, 292, 2147483648, 4294967295}
    v4 := make(map[string]int)
    v5 := []uint32{2, 819}
    v6 := make(map[int]int)
    if (831 < 254) {
        var v8 int8 = 1
        _ = len(v3)
        _ = v8
    }
    var v9 int = 4294967295
    for fi10 := 0; fi10 < 17; fi10++ {
        v2 = 1000
        _ = fi10
    }
    var v11 uint16 = 977
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        var v12 float64 = (float64((1e10 * (0.0 * 1e10))) + ((-2.25 / (1e10 - 3.14159)) / (1.5 - 0.0)))
        _ = v12
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        v4["hello"] = fi13
        _ = fi13
    }
    var v14 float32 = float32(-2.25)
    fmt.Println(v1, v11)
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 float64 = float64((v14 + 1.5))
        var v17 int = -2147483648
        if (v14 <= 0.1) {
            fmt.Println(v17)
        } else {
            var v18 bool = (v16 == float64(v14))
            fmt.Println(fi15)
            fmt.Println(v17, v14, v2)
            _ = v18
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    switch 32767 {
    case 416:
        for fi19 := 0; fi19 < 19; fi19++ {
            v20 := GenS1{F0: 255}
            if v1 {
                v21 := GenS1{F0: 65535}
                _ = v21
            } else {
                for fi22 := 0; fi22 < 13; fi22++ {
                    if ((v1 && v1) || false) {
                        var v23 float64 = 1e10
                        _ = v23
                    } else {
                        var v24 uint = 237
                        _ = len(v5)
                        var v25 int64 = -856
                        _ = v24
                        _ = v25
                    }
                    var v26 float64 = -2.25
                    _ = fi22
                    _ = v26
                }
            }
            var v27 uint32 = 1
            _ = fi19
            _ = v20
            _ = v27
        }
    default:
        _ = len(v6)
    }
    var v28 string = "a"
    v29 := GenS1{F0: 1}
    var v30 uint8 = 128
    if true {
        v1 = true
    } else {
        for fi31 := 0; fi31 < 6; fi31++ {
            var v32 uint8 = 180
            var v33 string = v28
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
