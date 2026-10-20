package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{978, 627}
    var v3 int32 = 42
    v4 := []int{1}
    v5 := []int{634, 4294967295, 32767}
    var v6 int8 = 75
    v7 := GenS1{F0: "a"}
    v6 = (v6 - v6)
    v3 += v3
    switch "hello" {
    case "":
        v8 := GenS1{F0: ""}
        _ = v8
    case "foo bar":
        for fi9 := 0; fi9 < 19; fi9++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v10 uint8 = 57
            _ = fi9
            _ = v10
        }
    default:
        var v11 int8 = -49
        _ = v11
    }
    var v12 int64 = -1
    if (false || (0.1 != 1.5)) {
        v12 = -1
        var v13 uint64 = 2
        _ = v13
    }
    v14 := v1[128]
    if (1e10 <= 3.14159) {
        fmt.Println(v14, v12)
    } else {
        if true {
            v3 -= v3
            var v15 float64 = -2.25
            v6 = v6
            _ = v15
        } else {
            for fi16 := 0; fi16 < 4; fi16++ {
                _ = len(v4)
                for fi17 := 0; fi17 < 16; fi17++ {
                    var v18 string = ("" + "hello")
                    if (true && true) {
                        v18 = "\n\t"
                        var v19 string = "hello"
                        var v20 float64 = 3.14159
                        _ = v19
                        _ = v20
                    } else {
                        v18 = ""
                        var v21 int64 = 1
                        _ = v21
                    }
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
            v3 = v3
        }
        v3 = (-(v3))
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v22 := GenS2{F0: false, F1: 246}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
