package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint32 }
func genFn0(p0 uint, p1 int32, p2 bool) string {
    v7 := []int{128}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    p1 += p1
    if p2 {
        v8 := GenS1{F0: 9223372036854775807, F1: 1000, F2: uint32(p0)}
        _ = v8
    } else {
        p0 -= p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ("foo bar" + (("" + "foo bar") + "hello"))
}

func main() {
    v1 := make(map[uint32]int)
    var v2 string = "foo bar"
    var v3 float32 = 0.0
    var v4 int = 364
    var v5 uint32 = 868
    v6 := make(map[uint32]string)
    if false {
        var v9 uint = 42
        _ = v9
    } else {
        v3 -= float32((0.0 - (0.0 + ((-2.25 + (1.5 + float64(v3))) + 3.14159))))
    }
    if false {
        v10 := []int{(v4 >> ((42) & 63))}
        _ = v10
    } else {
        switch (v2 + ((v2 + v2) + "\n\t")) {
        case "foo bar":
            for fi11 := 0; fi11 < 14; fi11++ {
                fmt.Println(v4, v3, v4)
                _ = fi11
            }
        case "hello":
            var v12 uint8 = 42
            _ = v12
        case "\n\t":
            var v13 int64 = -1
            _ = v13
        default:
            v2 = "foo bar"
        }
    }
    switch (((v2 + v2) + "hello") + v2) {
    case "":
        var v14 int = (v4 - 491)
        _ = v14
    case "foo bar":
        v5 += (v5 >> ((930) & 63))
    }
    if (!(!((v2 > "\n\t") && true))) {
        v5 = v5
        var v15 int64 = 92
        _ = v15
    } else {
        delete(v1, (v5 + 329))
        switch v4 {
        case 480:
            var v16 float32 = v3
            _ = v16
        default:
            fmt.Println(v3)
        }
    }
    var v17 uint32 = v5
    v17 = (v5 >> ((0) & 63))
    fmt.Println(v2)
    switch v2 {
    case "hello":
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 int8 = 42
            _ = fi18
            _ = v19
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
