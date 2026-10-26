package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int16 }
type GenS2 struct { F0 int16; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 int }
func genFn0(p0 int) uint16 {
    if false {
        var v5 float64 = float64(1.5)
        _ = v5
    }
    var v6 uint32 = 536
    var v7 uint32 = v6
    v8 := []int{}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 int8, p1 uint64, p2 uint32) string {
    p1 = p1
    var v9 bool = false
    v10 := GenS1{F0: (127 == 0), F1: 354, F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return "foo bar"
}

func main() {
    v1 := []string{"foo bar", "a"}
    var v2 int32 = 1
    var v3 float64 = 0.0
    var v4 int8 = 0
    var v11 uint = 167
    var v12 uint16 = 660
    fmt.Println(v4, v4)
    if true {
        if true {
            v11 += v11
        } else {
            v12 += v12
            v13 := GenS2{F0: -32768, F1: (-(v2)), F2: 2}
            _ = v13
        }
        var v14 int16 = 0
        if len(v1) > 0 {
            v1[0] = ("\n\t" + ("\n\t" + ("a" + "")))
        }
        _ = v14
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v15 bool = (!true)
        if v15 {
            v16 := GenS1{F0: v15, F1: -242, F2: 42}
            var v17 int16 = 872
            _ = v16
            _ = v17
        }
        _ = v15
    }
    switch "foo bar" {
    case "a":
        fmt.Println(v3, v11, v2)
    default:
        for fi18 := 0; fi18 < 13; fi18++ {
            var v19 uint16 = (v12 ^ (^(v12)))
            fmt.Println(v3, fi18, fi18)
            v20 := GenS3{F0: (v11 + v11)}
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    switch ("\n\t" + "") {
    case "foo bar":
        v1 = append(v1, ("hello" + ("\n\t" + ("hello" + ("\n\t" + "")))))
    default:
        v21 := GenS4{F0: 0, F1: 127}
        _ = v21
    }
    v11 += 18446744073709551615
    var v22 uint8 = 236
    v23 := GenS1{F0: true, F1: 0, F2: 32767}
    if false {
        var v24 float64 = v3
        _ = v24
    } else {
        var v25 int32 = (v2 ^ v2)
        var v26 float32 = 1.5
        var v27 int32 = 123
        _ = v25
        _ = v26
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
