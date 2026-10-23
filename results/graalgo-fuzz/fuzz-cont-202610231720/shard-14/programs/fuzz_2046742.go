package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint64 }

func main() {
    var v1 int16 = -1
    var v2 string = "a"
    var v3 int64 = -941
    v4 := make(map[int]int)
    v5 := []string{v2, ("\n\t" + v2)}
    v6 := GenS1{F0: "\n\t", F1: 65535, F2: 2}
    v3 = v3
    var v7 uint8 = 1
    var v8 uint16 = 2
    var v9 float32 = 0.0
    switch v2 {
    case "hello":
        v3 -= v3
    case "a":
        var v10 int64 = v3
        _ = v10
    default:
        var v11 string = ""
        _ = v11
    }
    v1 += (v1 / ((v1) | 1))
    v2 = ""
    v9 = (1e10 * v9)
    var v12 int16 = v1
    for fi13 := 0; fi13 < 7; fi13++ {
        for fi14 := 0; fi14 < 2; fi14++ {
            switch v2 {
            case "a":
                v9 += (v9 - v9)
            case "hello":
                v12 += (v1 / (((v1 * v12)) | 1))
            default:
                var v15 float64 = 0.0
                _ = v15
            }
            delete(v4, (fi14 ^ fi13))
            if false {
                var v16 bool = true
                _ = v16
            }
            v12 += v1
            _ = fi14
        }
        _ = fi13
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        v9 = v9
        if (((v12 >> ((1) & 63)) > (v12 + v1)) || (false && true)) {
            var v18 float32 = v9
            _ = v18
        } else {
            var v19 uint = 838
            v9 += v9
            _ = v19
        }
        _ = fi17
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        for fi21 := 0; fi21 < 7; fi21++ {
            v4[(fi21 ^ (fi21 % (((^(fi20))) | 1)))] = (fi21 * 128)
            if false {
                v8 = 0
            }
            _ = fi21
        }
        _ = fi20
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        v3 = (v3 / ((781) | 1))
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
