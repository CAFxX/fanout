package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 bool; F1 bool; F2 bool }
func genFn0() uint16 {
    v7 := []int{}
    v8 := GenS1{F0: -271, F1: 1}
    _ = v7
    _ = v8
    return 32767
}

func main() {
    var v1 uint16 = 809
    v2 := make(map[uint32]int)
    var v3 uint = 1
    v4 := make(map[uint32]bool)
    var v5 int16 = 32767
    v6 := []string{"hello"}
    for fi9 := 0; fi9 < 20; fi9++ {
        for fi10 := 0; fi10 < 20; fi10++ {
            switch "\n\t" {
            case "a":
                v11 := GenS1{F0: v5, F1: -940}
                _ = v11
            case "":
                var v12 uint64 = 42
                _ = v12
            case "foo bar":
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v3 = (^(v3))
                var v13 uint32 = 2147483648
                _ = v13
            default:
                var v14 float64 = 1e10
                _ = v14
            }
            switch ("\n\t" + "foo bar") {
            case "":
                if (!true) {
                    fmt.Println(fi10, v1, fi10)
                }
            default:
                v15 := GenS2{F0: false, F1: 255}
                _ = v15
            }
            v1 -= genFn0()
            _ = fi10
        }
        _ = fi9
    }
    v16 := GenS2{F0: true, F1: (v1 << (((v3 + v3)) & 63))}
    if true {
        var v17 int32 = 42
        _ = v17
    }
    v5 = v5
    switch 4294967295 {
    case -96:
        v18 := GenS3{F0: 2147483647, F1: 367}
        _ = v18
    }
    _ = len(v2)
    var v19 string = "foo bar"
    v20 := v4[2147483647]
    v6 = append(v6, "foo bar")
    for fi21 := 0; fi21 < 6; fi21++ {
        v22 := GenS4{F0: true, F1: (1 > 2147483647), F2: true}
        v19 = v19
        v20 = v20
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
