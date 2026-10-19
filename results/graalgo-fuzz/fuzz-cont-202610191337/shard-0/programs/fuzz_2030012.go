package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
func genFn0(p0 int) uint64 {
    p0 = p0
    var v7 uint8 = 1
    var v8 uint32 = 42
    v9 := []int{}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1() uint32 {
    var v10 string = "foo bar"
    v10 = "hello"
    _ = v10
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 100
    v3 := []string{("\n\t" + "\n\t")}
    var v4 int16 = -32768
    var v5 int32 = 42
    var v6 bool = true
    var v11 string = "a"
    v3 = append(v3, (v11 + ""))
    if (0 > 97) {
        var v12 int16 = 1000
        v13 := v1[42]
        _ = v12
        _ = v13
    }
    var v14 float64 = ((1.5 / 1.5) / (1.5 + -2.25))
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 uint8 = 42
        _ = fi15
        _ = v16
    }
    v1[42] = 0
    _ = len(v1)
    v4 = 0
    fmt.Println(v11)
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 uint64 = 2
        for fi19 := 0; fi19 < 15; fi19++ {
            switch ((v11 + "foo bar") + (v11 + v11)) {
            case "\n\t":
                v11 = ((v11 + ((v11 + v11) + "\n\t")) + ("hello" + "hello"))
            default:
                var v20 uint8 = 1
                _ = v20
            }
            _ = fi19
        }
        v21 := GenS1{F0: 65535, F1: 32768}
        _ = fi17
        _ = v18
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
