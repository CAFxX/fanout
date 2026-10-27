package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
func genFn0(p0 int32, p1 uint8) int {
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := []int{fi6}
        p1 = 42
        v8 := GenS1{F0: 2, F1: -9223372036854775808}
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 float32 = (1e10 + 0.0)
        _ = fi9
        _ = v10
    }
    var v11 int8 = -50
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v11
    return 32767
}
func genFn1(p0 int32, p1 string) uint64 {
    var v12 string = "a"
    v12 = ((v12 + "") + ("hello" + ((v12 + "hello") + p1)))
    _ = p0
    _ = p1
    _ = v12
    return 42
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    v3 := make(map[string]uint32)
    var v4 int8 = 0
    v5 := make(map[int]uint32)
    if true {
        switch "foo bar" {
        case "\n\t":
            _ = len(v3)
        }
        var v13 uint = 18446744073709551615
        _ = v13
    }
    v4 = 105
    if false {
        v4 = -1
        v4 = v4
    }
    var v14 uint8 = 59
    for fi15 := 0; fi15 < 12; fi15++ {
        v4 -= v4
        fmt.Println(v14, v4)
        v16 := []int{fi15}
        _ = fi15
        _ = v16
    }
    v14 = 1
    for fi17 := 0; fi17 < 17; fi17++ {
        fmt.Println(fi17, v14, v4)
        _ = fi17
    }
    if (((float32(0.0) + (1e10 / (1.5 - (0.1 - (0.1 + (((0.0 * 0.1) * (1e10 * float32((1.5 + 0.1)))) / ((0.0 - (1e10 / 1.0)) * -2.25))))))) < float32(0.0)) || true) {
        v18 := GenS1{F0: 42, F1: -1}
        switch ("" + "hello") {
        case "\n\t":
            v4 = (v4 - v4)
        case "hello":
            var v19 int64 = -9223372036854775808
            _ = v19
        }
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
