package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int16 }
func genFn0() int32 {
    var v6 uint64 = 18446744073709551615
    var v7 uint = 0
    v7 += v7
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 bool = (42 < 32767)
    v2 := []string{""}
    var v3 uint32 = 656
    var v4 uint8 = 255
    var v5 uint64 = 355
    v5 = uint64(v4)
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 int16 = 32767
            var v11 float64 = (1e10 + 1.5)
            v12 := GenS1{F0: genFn0(), F1: 42}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        _ = fi8
    }
    var v13 uint64 = 901
    v13 = v13
    fmt.Println(v5)
    if (((0.1 / (-2.25 / 1.0)) / 1.0) <= -2.25) {
        v4 = (v4 >> ((65535) & 63))
        v14 := GenS1{F0: 42, F1: 127}
        var v15 bool = (!v1)
        _ = v14
        _ = v15
    }
    if ((v1 && v1) && v1) {
        v2 = append(v2, "foo bar")
        var v16 float32 = 0.1
        for fi17 := 0; fi17 < 4; fi17++ {
            if true {
                switch (^(fi17)) {
                case 42:
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                default:
                    v4 -= v4
                }
                if len(v2) > 0 {
                    v2[0] = ("\n\t" + "")
                }
            }
            if (380 > 567) {
                v18 := GenS1{F0: 0, F1: 100}
                _ = v18
            } else {
                v19 := GenS2{F0: ((v16 + float32(0.0)) >= ((v16 / v16) - v16))}
                _ = v19
            }
            _ = fi17
        }
        _ = v16
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        v21 := GenS3{F0: v3, F1: 1000}
        var v22 int16 = -1
        _ = fi20
        _ = v21
        _ = v22
    }
    v5 = v13
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
