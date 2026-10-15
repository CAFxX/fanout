package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 bool }

func main() {
    var v1 bool = false
    var v2 uint16 = 42
    var v3 float64 = 1e10
    v4 := []uint32{}
    v5 := []string{"\n\t"}
    v3 = -2.25
    if v1 {
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
        var v6 uint8 = 128
        _ = v6
    }
    var v7 bool = (!(42 > 28))
    switch (("" + "\n\t") + ("\n\t" + ("" + ""))) {
    case "hello":
        v8 := GenS1{F0: 1, F1: 700, F2: v1}
        _ = v8
    default:
        fmt.Println(v1)
    }
    v2 -= 2
    fmt.Println(v3, v1)
    v2 = (v2 / ((v2) | 1))
    v1 = v1
    v1 = false
    v9 := GenS1{F0: 42, F1: 2147483647, F2: (!v1)}
    var v10 int8 = -66
    var v11 uint = 18446744073709551615
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 uint8 = 2
        if v7 {
            if v7 {
                var v14 uint16 = v2
                var v15 uint16 = (v14 | 2)
                v7 = (0 <= (-(fi12)))
                _ = v14
                _ = v15
            } else {
                switch (("" + ("hello" + "hello")) + "a") {
                case "hello":
                    fmt.Println(v11)
                default:
                    v11 = v11
                }
            }
        } else {
            v13 += (v13 >> ((v11) & 63))
            var v16 uint16 = 32767
            v16 = (v16 + v2)
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    if v1 {
        if (!(0 > v11)) {
            v17 := GenS1{F0: v10, F1: -557, F2: (!(!v1))}
            var v18 uint8 = 0
            _ = v17
            _ = v18
        }
    }
    switch ("\n\t" + "foo bar") {
    case "":
        switch 32768 {
        case 285:
            var v19 int16 = 421
            _ = v19
        case 500:
            fmt.Println(v3, v10)
        }
    default:
        v4 = append(v4, 0)
    }
    var v20 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
