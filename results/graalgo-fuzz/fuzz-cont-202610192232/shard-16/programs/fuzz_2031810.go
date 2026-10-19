package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint }
func genFn0(p0 int8, p1 int8, p2 int16) int8 {
    var v7 float64 = 1.5
    var v8 float32 = -2.25
    var v9 string = ("hello" + "")
    v10 := []int{1, -439, 61, 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return p0
}
func genFn1(p0 int, p1 uint32) int8 {
    v11 := GenS1{F0: (p0 * p0), F1: 1, F2: 42}
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := GenS1{F0: fi12, F1: 757, F2: 2}
        var v14 int64 = -1
        _ = fi12
        _ = v13
        _ = v14
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v11
    return 100
}

func main() {
    var v1 int = -842
    var v2 uint = 2
    v3 := make(map[uint32]uint32)
    var v4 bool = ((!(("\n\t" + "") < "hello")) && (!(true || false)))
    var v5 int = 0
    var v6 int32 = 42
    var v15 int16 = -1
    v1 -= v1
    v6 = (v6 * (^(v6)))
    if (false || (!true)) {
        v16 := GenS1{F0: (v5 - v5), F1: v2, F2: v2}
        _ = v16
    }
    fmt.Println(v2, v15, v6)
    if true {
        for fi17 := 0; fi17 < 6; fi17++ {
            for fi18 := 0; fi18 < 16; fi18++ {
                v6 += (v6 + 42)
                _ = fi18
            }
            var v19 uint = (v2 - v2)
            _ = fi17
            _ = v19
        }
        fmt.Println(v1, v4)
        var v20 bool = false
        _ = v20
    }
    v4 = (!true)
    for fi21 := 0; fi21 < 12; fi21++ {
        v22 := GenS1{F0: 4294967295, F1: (^(v2)), F2: 255}
        var v23 bool = false
        _ = fi21
        _ = v22
        _ = v23
    }
    v5 = 127
    v24 := GenS1{F0: (v1 * 4294967295), F1: (v2 & v2), F2: (^(v2))}
    switch ("hello" + "a") {
    case "foo bar":
        v25 := []int{}
        _ = v25
    default:
        v5 += (v5 + 733)
    }
    var v26 int = v1
    var v27 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
