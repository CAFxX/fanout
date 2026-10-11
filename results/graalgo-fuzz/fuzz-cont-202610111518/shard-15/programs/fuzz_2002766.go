package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 int32; F2 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 uint64; F2 string }
func genFn0() int8 {
    var v7 int16 = 32767
    v8 := GenS1{F0: 1}
    fmt.Println(v7)
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 float64 = 1.5
    var v2 float64 = v1
    var v3 float64 = 3.14159
    var v4 uint8 = 29
    v5 := make(map[string]bool)
    var v6 uint64 = 0
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 float32 = 1e10
        v2 = float64(3.14159)
        _ = fi9
        _ = v10
    }
    if true {
        var v11 int64 = 65
        fmt.Println(v6)
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 int = (-(fi12))
        _ = fi12
        _ = v13
    }
    v14 := v5["foo bar"]
    v1 = (v1 / (-2.25 + ((-2.25 + (v2 / 1.0)) - v2)))
    for fi15 := 0; fi15 < 7; fi15++ {
        for fi16 := 0; fi16 < 9; fi16++ {
            fmt.Println(v4, v3, v2)
            v4 -= v4
            _ = fi16
        }
        _ = fi15
    }
    if (51 <= 801) {
        fmt.Println(v6, v2)
        v17 := GenS1{F0: 288}
        for fi18 := 0; fi18 < 5; fi18++ {
            v1 = ((float64(((3.14159 * 1e10) / -2.25)) / 1.0) / 1.0)
            v19 := GenS2{F0: 414, F1: 20, F2: ((("" + "\n\t") + "") + "\n\t")}
            _ = fi18
            _ = v19
        }
        _ = v17
    } else {
        v20 := GenS3{F0: -9223372036854775808}
        v1 = v3
        _ = v20
    }
    fmt.Println(v1)
    v5[(("\n\t" + "hello") + (("hello" + "foo bar") + "foo bar"))] = (-1 >= int16(v4))
    v21 := GenS4{F0: 0, F1: v6, F2: "foo bar"}
    v22 := v5["hello"]
    var v23 uint8 = v4
    v14 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
