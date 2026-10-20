package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint64, p1 int32, p2 uint32) uint32 {
    v7 := []int{793, -1}
    for fi8 := 0; fi8 < 19; fi8++ {
        fmt.Println(p1)
        _ = fi8
    }
    fmt.Println(p1, p2, p2)
    v9 := GenS1{F0: -128}
    var v10 string = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v9
    return p2
}

func main() {
    v1 := make(map[int]int)
    var v2 string = "foo bar"
    var v3 uint64 = 2
    var v4 uint32 = 2147483648
    var v5 float32 = (-2.25 + (((3.14159 + 1e10) + 1e10) * 0.0))
    v6 := make(map[uint32]int)
    var v11 uint16 = 896
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS1{F0: 1}
        _ = fi12
        _ = v13
    }
    var v14 float32 = float32(0.1)
    v15 := GenS2{F0: v3}
    v16 := v1[400]
    for fi17 := 0; fi17 < 11; fi17++ {
        if (0 == 1) {
            fmt.Println(v2, v14)
        } else {
            var v18 bool = (2147483647 <= -2147483648)
            _ = v18
        }
        switch v2 {
        case "hello":
            v11 = 243
        case "\n\t":
            for fi19 := 0; fi19 < 8; fi19++ {
                fmt.Println(v16, v14, v14)
                _ = fi19
            }
        case "a":
            var v20 uint16 = 255
            _ = v20
        }
        var v21 uint16 = (v11 >> ((1) & 63))
        _ = fi17
        _ = v21
    }
    v22 := []int{v16, v16, (^(v16)), v16}
    var v23 int64 = -9223372036854775808
    if true {
        var v24 bool = ((^(v4)) != 2147483647)
        switch (v16 >> ((194) & 63)) {
        case -586:
            for fi25 := 0; fi25 < 7; fi25++ {
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
                _ = fi25
            }
        default:
            var v26 string = v2
            _ = v26
        }
        _ = v24
    }
    v4 = v4
    for fi27 := 0; fi27 < 13; fi27++ {
        for fi28 := 0; fi28 < 4; fi28++ {
            var v29 int8 = 100
            _ = fi28
            _ = v29
        }
        v11 -= v11
        _ = fi27
    }
    fmt.Println(v14, v14, v23)
    for fi30 := 0; fi30 < 12; fi30++ {
        var v31 int8 = 57
        fmt.Println(v4, v5)
        v3 = v3
        _ = fi30
        _ = v31
    }
    var v32 int64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
