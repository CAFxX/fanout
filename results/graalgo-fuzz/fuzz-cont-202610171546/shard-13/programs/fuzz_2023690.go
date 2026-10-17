package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 int16 }
func genFn0(p0 int, p1 int16) uint8 {
    p1 = p1
    if false {
        v7 := []int{p0, 9223372036854775807, 32768}
        _ = len(v7)
        fmt.Println(p1, p0)
        _ = v7
    }
    v8 := GenS1{F0: 939}
    v9 := []int{404}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 165
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = true
    var v3 int = -506
    var v4 int8 = 1
    var v5 int16 = -32768
    v6 := make(map[uint32]int)
    for fi10 := 0; fi10 < 5; fi10++ {
        for fi11 := 0; fi11 < 19; fi11++ {
            v12 := GenS1{F0: -1}
            var v13 float32 = 3.14159
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 string = "foo bar"
        var v16 float64 = (((-2.25 / 1.0) - 1e10) * 1.5)
        var v17 int32 = 2147483647
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: -1}
    var v19 uint16 = 0
    v20 := GenS1{F0: 9223372036854775807}
    if v2 {
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 uint8 = 47
            var v23 int8 = -88
            var v24 uint16 = 32767
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
    } else {
        v25 := GenS2{F0: 613, F1: v5}
        switch ((("foo bar" + "\n\t") + "hello") + "foo bar") {
        case "":
            if ((((-2.25 * (-2.25 / 1.0)) - (-2.25 + 1.5)) / 3.14159) >= (0.1 / (0.1 / (1e10 * (1.5 / 3.14159))))) {
                v26 := GenS2{F0: int32(v3), F1: 1000}
                v1[2147483648] = 1
                _ = v26
            }
        case "hello":
            var v27 uint = 351
            _ = v27
        case "\n\t":
            if (v3 <= 9223372036854775807) {
                if true {
                    v28 := GenS2{F0: 0, F1: -1}
                    _ = v28
                } else {
                    var v29 bool = (true || (!(42 == v3)))
                    v30 := GenS2{F0: 42, F1: (v5 & (v5 << ((65535) & 63)))}
                    v19 = v19
                    _ = v29
                    _ = v30
                }
            }
        }
        var v31 float64 = ((float64(float32(float64(1e10))) - (1.5 - (1.5 + 1e10))) - ((1.5 * 0.1) * 1.5))
        _ = v25
        _ = v31
    }
    var v32 uint16 = 32767
    var v33 int64 = -911
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
