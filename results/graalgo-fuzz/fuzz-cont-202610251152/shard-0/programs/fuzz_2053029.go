package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 int; F2 int16 }
func genFn0(p0 int32, p1 int8) uint8 {
    if false {
        var v6 string = ""
        v6 = ("\n\t" + "foo bar")
        _ = v6
    } else {
        p1 = p1
        v7 := []int{-1}
        _ = v7
    }
    var v8 uint16 = 402
    var v9 uint = 2
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 int32, p1 uint32, p2 uint) bool {
    var v10 string = "hello"
    for fi11 := 0; fi11 < 10; fi11++ {
        switch v10 {
        case "foo bar":
            var v12 uint16 = 32767
            _ = v12
        default:
            v13 := []int{9223372036854775807, 32768}
            _ = v13
        }
        _ = fi11
    }
    var v14 int = -675
    var v15 uint = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v14
    _ = v15
    return true
}

func main() {
    var v1 float32 = float32(float64((-2.25 / 0.1)))
    v2 := []int{}
    v3 := []string{}
    v4 := make(map[string]string)
    v5 := make(map[string]bool)
    var v16 uint64 = 0
    for fi17 := 0; fi17 < 17; fi17++ {
        if (!true) {
            v16 -= v16
            v18 := GenS1{F0: 32767}
            var v19 uint64 = (v16 ^ (v16 / ((584) | 1)))
            _ = v18
            _ = v19
        } else {
            v20 := GenS1{F0: 60}
            v1 += v1
            v21 := GenS2{F0: fi17, F1: 9223372036854775807, F2: -397}
            _ = v20
            _ = v21
        }
        fmt.Println(v1)
        _ = fi17
    }
    v1 -= (v1 - (((1e10 * 1e10) / ((v1 / 1.0) * (1e10 - 1.5))) * v1))
    var v22 float64 = 0.1
    v22 -= float64(0.1)
    v23 := GenS1{F0: 65535}
    for fi24 := 0; fi24 < 6; fi24++ {
        var v25 uint = 42
        _ = fi24
        _ = v25
    }
    var v26 int32 = 2147483647
    v27 := GenS1{F0: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
