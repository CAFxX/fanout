package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 int; F2 uint32 }

func main() {
    var v1 string = ("" + "foo bar")
    var v2 float32 = 1.5
    v3 := make(map[string]uint32)
    var v4 int8 = -2
    v5 := make(map[string]string)
    var v6 uint32 = 2147483647
    var v7 string = "foo bar"
    v8 := GenS1{F0: 4294967295, F1: v6}
    switch 368 {
    case 776:
        v5[v1] = v7
    case -1:
        var v9 bool = (397 != 255)
        _ = v9
    case 4294967295:
        v7 = ((v7 + v1) + "a")
    default:
        v10 := GenS2{F0: 1}
        _ = v10
    }
    v4 += v4
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 float64 = 1.5
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 bool = true
            if v14 {
                v15 := []int{(fi13 % (((fi13 ^ (fi13 >> ((810) & 63)))) | 1)), (fi11 % (((fi11 / ((-261) | 1))) | 1)), fi13, (fi13 + fi13)}
                v14 = v14
                var v16 int = 834
                _ = v15
                _ = v16
            }
            var v17 int64 = -1
            _ = fi13
            _ = v14
            _ = v17
        }
        _ = fi11
        _ = v12
    }
    delete(v3, ((v1 + v7) + v7))
    for fi18 := 0; fi18 < 9; fi18++ {
        v2 -= (v2 - float32(-2.25))
        v19 := GenS1{F0: 149, F1: (v6 ^ v6)}
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        if ((fi20 ^ fi20) <= fi20) {
            v1 = "a"
            var v21 uint = 0
            var v22 uint = (v21 * (v21 ^ v21))
            _ = v21
            _ = v22
        } else {
            v2 += 1.5
            v23 := GenS3{F0: (v6 % ((2147483648) | 1)), F1: (fi20 % ((4294967295) | 1)), F2: (v6 >> ((299) & 63))}
            _ = v23
        }
        if (!true) {
            var v24 float32 = v2
            switch "" {
            case "":
                for fi25 := 0; fi25 < 13; fi25++ {
                    if (false || (!true)) {
                        var v26 int8 = v4
                        var v27 float64 = 3.14159
                        v6 -= (v6 % (((^(v6))) | 1))
                        _ = v26
                        _ = v27
                    } else {
                        v28 := GenS1{F0: 18446744073709551615, F1: v6}
                        var v29 uint64 = 18446744073709551615
                        v24 = v24
                        _ = v28
                        _ = v29
                    }
                    v3[(v7 + ("foo bar" + v7))] = 65535
                    _ = fi25
                }
            default:
                var v30 bool = ((((((v24 + v2) - ((((-2.25 + 1.5) / (v24 - v2)) - 1e10) + v2)) + ((v2 + (float32(1.5) + v24)) * float32((1.5 + 0.1)))) + 0.0) - (((v24 + 1e10) / ((v24 / float32(3.14159)) - (v24 - 1e10))) + v24)) < (0.0 / v24))
                _ = v30
            }
            v1 = v7
            _ = v24
        } else {
            fmt.Println(fi20)
            v31 := GenS1{F0: 2, F1: v6}
            _ = v31
        }
        _ = fi20
    }
    v32 := []int{444, -94}
    switch v1 {
    case "\n\t":
        if true {
            v6 += v6
            for fi33 := 0; fi33 < 19; fi33++ {
                _ = len(v5)
                _ = fi33
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
