package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
func genFn0(p0 bool) int {
    var v4 uint64 = 18446744073709551615
    fmt.Println(p0)
    for fi5 := 0; fi5 < 5; fi5++ {
        v4 += (v4 ^ v4)
        var v6 uint8 = 1
        var v7 float64 = 1e10
        _ = fi5
        _ = v6
        _ = v7
    }
    p0 = p0
    _ = p0
    _ = v4
    return 1
}

func main() {
    var v1 uint32 = 0
    var v2 int8 = -128
    var v3 int32 = -1
    v8 := GenS1{F0: -1, F1: 2}
    var v9 bool = false
    var v10 int64 = -1
    for fi11 := 0; fi11 < 12; fi11++ {
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 string = (("\n\t" + "foo bar") + "hello")
            v9 = v9
            _ = fi12
            _ = v13
        }
        var v14 uint64 = 2
        _ = fi11
        _ = v14
    }
    var v15 bool = (((v1 & 2) >= v1) || v9)
    v16 := GenS1{F0: 1, F1: uint16(v1)}
    var v17 uint = 1
    for fi18 := 0; fi18 < 2; fi18++ {
        v15 = (v9 && (!(4294967295 < 4294967295)))
        v19 := GenS1{F0: -467, F1: 32767}
        _ = fi18
        _ = v19
    }
    var v20 float32 = 0.1
    for fi21 := 0; fi21 < 8; fi21++ {
        v9 = (!false)
        var v22 int32 = v3
        switch fi21 {
        case 42:
            var v23 uint = (v17 - v17)
            _ = v23
        default:
            var v24 string = "\n\t"
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v8
    _ = v9
}
