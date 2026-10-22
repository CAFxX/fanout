package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint8 }

func main() {
    v1 := []int{127, 32767, -9223372036854775808}
    v2 := make(map[uint32]string)
    v3 := make(map[int]string)
    var v4 string = ""
    switch ("hello" + v4) {
    case "hello":
        var v5 bool = true
        _ = v5
    default:
        v6 := GenS1{F0: 127}
        _ = v6
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 int16 = -303
        _ = fi7
        _ = v8
    }
    v4 = v4
    for fi9 := 0; fi9 < 7; fi9++ {
        v4 = v4
        _ = fi9
    }
    fmt.Println(v4)
    if false {
        if (false || (-128 == 100)) {
            fmt.Println(v4)
            v4 = ""
        }
    } else {
        for fi10 := 0; fi10 < 13; fi10++ {
            v4 = (v4 + (v4 + (v4 + v4)))
            _ = fi10
        }
        v4 = v4
        var v11 bool = false
        _ = v11
    }
    v4 = "foo bar"
    var v12 int32 = -1
    v13 := GenS2{F0: 53, F1: 4294967295, F2: 127}
    var v14 int = -1
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 int64 = 0
        v17 := GenS2{F0: 100, F1: 559, F2: 0}
        var v18 float32 = 0.0
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    var v19 int64 = -9223372036854775808
    fmt.Println(v14, v19)
    for fi20 := 0; fi20 < 7; fi20++ {
        if true {
            var v21 uint32 = 26
            _ = v21
        }
        v22 := GenS2{F0: 119, F1: 18446744073709551615, F2: 40}
        _ = fi20
        _ = v22
    }
    fmt.Println(v12, v19, v4)
    v12 = -742
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
