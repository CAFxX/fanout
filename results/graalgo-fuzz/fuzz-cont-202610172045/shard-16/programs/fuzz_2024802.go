package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64 }
func genFn0(p0 int8, p1 int32) int8 {
    v4 := GenS1{F0: true}
    var v5 uint = 18446744073709551615
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{"hello", ("\n\t" + "foo bar"), "hello"}
    v3 := make(map[int]string)
    v6 := v3[-791]
    fmt.Println(v6)
    fmt.Println(v6)
    v6 = v6
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            v6 = v6
            v6 = (v6 + "hello")
            _ = len(v2)
            _ = fi8
        }
        fmt.Println(fi7, v6)
        _ = fi7
    }
    switch 583 {
    case 4294967295:
        v9 := GenS2{F0: (!(true && (!false))), F1: 255}
        _ = v9
    default:
        v6 = (((v6 + v6) + "foo bar") + v6)
    }
    switch 587 {
    case 9223372036854775807:
        if false {
            if (!true) {
                if true {
                    fmt.Println(v6)
                    for fi10 := 0; fi10 < 3; fi10++ {
                        fmt.Println(v6)
                        if (0 == 255) {
                            var v11 int32 = 913
                            _ = v11
                        } else {
                            var v12 uint16 = 2
                            _ = v12
                        }
                        _ = fi10
                    }
                }
                v13 := v1[2147483647]
                _ = v13
            }
        }
    case 42:
        var v14 int32 = -1
        _ = v14
    }
    var v15 float32 = 0.0
    if ((!(false && false)) || ((0 == 365) && true)) {
        var v16 bool = false
        if true {
            var v17 float64 = 0.0
            _ = v17
        } else {
            v15 = v15
        }
        _ = v16
    } else {
        v15 = (-2.25 / float32(-2.25))
    }
    if false {
        var v18 int32 = 262
        if len(v2) > 0 {
            v2[0] = v6
        }
        fmt.Println(v15)
        _ = v18
    } else {
        v15 = (float32(0.0) + (v15 + (v15 - float32(1.5))))
        _ = len(v2)
        for fi19 := 0; fi19 < 17; fi19++ {
            var v20 string = v6
            var v21 int64 = 63
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    var v22 int8 = -1
    delete(v3, 2147483647)
    v23 := GenS3{F0: true}
    fmt.Println(v22, v22, v22)
    v24 := GenS4{F0: -356}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v6
}
