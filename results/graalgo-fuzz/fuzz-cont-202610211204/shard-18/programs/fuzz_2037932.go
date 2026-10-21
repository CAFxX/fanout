package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int64, p1 int) int32 {
    switch ("" + (("a" + "\n\t") + "foo bar")) {
    case "\n\t":
        v4 := GenS1{F0: 122}
        _ = v4
    case "hello":
        p0 = p0
    default:
        var v5 int64 = p0
        _ = v5
    }
    var v6 int64 = (^(p0))
    fmt.Println(p0, v6)
    _ = p0
    _ = p1
    _ = v6
    return 195
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = -128
    var v3 uint16 = 42
    v7 := []int{}
    if (-1 == -2147483648) {
        _ = len(v7)
    } else {
        switch "hello" {
        case "":
            if true {
                if false {
                    fmt.Println(v2, v3)
                    v2 = -128
                }
            }
        default:
            switch ("" + ("\n\t" + "foo bar")) {
            case "\n\t":
                var v8 float32 = 1.5
                _ = v8
            }
        }
        switch "hello" {
        case "foo bar":
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        case "\n\t":
            fmt.Println(v2)
        }
        fmt.Println(v2, v3)
    }
    var v9 uint64 = 884
    var v10 string = "hello"
    var v11 uint8 = 96
    var v12 int32 = 2147483647
    v9 = v9
    var v13 uint16 = (v3 >> ((912) & 63))
    var v14 int64 = 0
    var v15 int = 127
    var v16 int16 = int16(v12)
    for fi17 := 0; fi17 < 1; fi17++ {
        v2 += -49
        var v18 uint32 = uint32(v9)
        v19 := GenS1{F0: 1}
        _ = fi17
        _ = v18
        _ = v19
    }
    if len(v7) > 0 {
        v7[0] = v15
    }
    var v20 uint8 = v11
    for fi21 := 0; fi21 < 16; fi21++ {
        var v22 float64 = ((float64(0.1) + (1.5 * (float64((1e10 / -2.25)) + 0.1))) + 3.14159)
        switch ((v10 + "foo bar") + v10) {
        case "a":
            v10 = (v10 + ((v10 + v10) + ""))
        }
        _ = fi21
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v7
    _ = v9
}
