package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 uint32 }
func genFn0(p0 string) int64 {
    v6 := []int{534, 562}
    if false {
        var v7 uint32 = 1
        p0 = p0
        _ = v7
    }
    if (p0 == ((p0 + p0) + p0)) {
        v8 := GenS1{F0: -270}
        _ = len(v6)
        _ = v8
    }
    v9 := GenS2{F0: 2147483647, F1: 2}
    var v10 int = 515
    _ = p0
    _ = v10
    _ = v6
    _ = v9
    return 481
}

func main() {
    var v1 float32 = float32(0.0)
    v2 := make(map[int]uint32)
    var v3 string = ("foo bar" + ("hello" + ""))
    var v4 uint = 2
    v5 := make(map[int]uint32)
    fmt.Println(v3)
    v3 = "foo bar"
    v11 := GenS1{F0: 42}
    var v12 uint = 18446744073709551615
    if (!false) {
        v13 := v5[-1]
        if true {
            v3 = v3
            var v14 int16 = -32768
            v15 := []int{32767}
            _ = v14
            _ = v15
        } else {
            var v16 int16 = 32767
            _ = v16
        }
        var v17 bool = (!(!true))
        _ = v13
        _ = v17
    }
    v18 := GenS1{F0: -1}
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 uint8 = 163
        _ = fi19
        _ = v20
    }
    if ((-2.25 * 3.14159) <= 1e10) {
        v3 = ((v3 + v3) + (v3 + "foo bar"))
        var v21 int8 = 96
        _ = v21
    }
    var v22 uint8 = 91
    var v23 string = "a"
    for fi24 := 0; fi24 < 15; fi24++ {
        v12 = v4
        _ = fi24
    }
    fmt.Println(v22)
    v25 := GenS2{F0: 0, F1: 1}
    var v26 uint = v4
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
