package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint16 }
func genFn0(p0 int16) int8 {
    var v7 uint64 = 574
    v8 := []int{-910, -719}
    switch ("foo bar" + (("" + "") + "a")) {
    case "":
        var v9 float32 = (0.0 / 1.5)
        _ = v9
    default:
        v7 = (v7 / (((v7 / ((v7) | 1))) | 1))
    }
    _ = len(v8)
    var v10 uint8 = 42
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    return 82
}

func main() {
    v1 := []uint32{}
    v2 := []uint32{2147483647, 899, 0}
    v3 := make(map[string]string)
    var v4 bool = (!true)
    var v5 float64 = 1.5
    var v6 int8 = -1
    var v11 uint64 = 18446744073709551615
    fmt.Println(v4)
    var v12 uint16 = 32767
    if (true && (-855 < 42)) {
        var v13 int8 = (v6 & 127)
        _ = v13
    } else {
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 uint = 192
            switch "\n\t" {
            case "a":
                for fi16 := 0; fi16 < 2; fi16++ {
                    v17 := GenS1{F0: 42, F1: 65535, F2: v12}
                    switch (fi14 % ((int(v6)) | 1)) {
                    case 1:
                        var v18 uint64 = (^(v11))
                        _ = v18
                    default:
                        var v19 uint16 = (v12 >> ((v15) & 63))
                        _ = v19
                    }
                    _ = fi16
                    _ = v17
                }
            case "foo bar":
                if (!v4) {
                    var v20 float64 = 1e10
                    _ = v20
                } else {
                    v21 := GenS1{F0: -1, F1: v12, F2: v12}
                    _ = v21
                }
            case "":
                if false {
                    v12 += v12
                    var v22 string = ("a" + "hello")
                    _ = len(v2)
                    _ = v22
                } else {
                    v15 = 537
                    var v23 int32 = -304
                    _ = v23
                }
            default:
                var v24 uint16 = v12
                _ = v24
            }
            _ = fi14
            _ = v15
        }
        v12 = (v12 >> ((255) & 63))
    }
    var v25 int8 = (v6 >> ((118) & 63))
    v6 += (v25 ^ 42)
    v26 := GenS1{F0: 0, F1: (v12 ^ (v12 | (v12 & v12))), F2: v12}
    if false {
        fmt.Println(v25, v11, v12)
    } else {
        v12 = (v12 * 255)
        v4 = ((1e10 - (1.5 * 0.0)) == float32(1e10))
    }
    v27 := v3[("hello" + "a")]
    v28 := GenS1{F0: -1, F1: 1, F2: (v12 ^ 657)}
    var v29 int = -16
    var v30 uint = 1
    for fi31 := 0; fi31 < 16; fi31++ {
        v29 += fi31
        _ = fi31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
