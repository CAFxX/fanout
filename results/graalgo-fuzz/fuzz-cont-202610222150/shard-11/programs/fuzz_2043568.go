package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 int8; F2 uint8 }
func genFn0(p0 uint) int64 {
    p0 -= p0
    v5 := GenS1{F0: 1}
    p0 += p0
    if false {
        p0 = 0
        switch ("\n\t" + (("\n\t" + (((("" + "foo bar") + "a") + "\n\t") + (("" + "") + ("\n\t" + "foo bar")))) + ("" + ""))) {
        case "":
            var v6 int32 = 42
            _ = v6
        }
        v7 := []int{128, 888}
        _ = v7
    } else {
        var v8 uint64 = 18446744073709551615
        var v9 int16 = -282
        var v10 string = ("" + "")
        _ = v10
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v5
    return 209
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 65
    var v3 uint64 = 42
    var v4 int8 = (v2 % ((-118) | 1))
    v2 = v4
    v3 = v3
    v4 = v2
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 bool = true
        switch fi11 {
        case -338:
            v3 = v3
        default:
            var v13 float64 = (0.1 - (0.0 * (-2.25 * 0.1)))
            _ = v13
        }
        fmt.Println(v3, v2)
        _ = fi11
        _ = v12
    }
    v14 := GenS2{F0: "", F1: 1}
    v15 := GenS3{F0: 1}
    fmt.Println(v2, v3)
    v16 := GenS4{F0: -824, F1: v4, F2: 128}
    if true {
        var v17 int8 = v4
        _ = v17
    }
    v18 := GenS2{F0: ("\n\t" + ("\n\t" + "foo bar")), F1: -1}
    var v19 bool = false
    v4 = (-(v2))
    for fi20 := 0; fi20 < 15; fi20++ {
        var v21 string = ""
        _ = fi20
        _ = v21
    }
    v1[676] = 0
    var v22 bool = v19
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
