package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int8 }
func genFn0() int8 {
    var v7 int32 = 1
    switch "foo bar" {
    case "":
        v7 = (v7 ^ v7)
    case "\n\t":
        v7 = v7
    }
    v8 := GenS1{F0: 100, F1: 255}
    fmt.Println(v7)
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 string = ""
    v3 := []string{v2, (("hello" + v2) + "a"), v2, v2}
    var v4 int64 = 1
    v5 := []string{("hello" + v2), v2, ("hello" + "hello"), ((v2 + v2) + ("hello" + "\n\t"))}
    var v6 uint64 = 2
    var v9 int16 = -32768
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 int = (-(fi10))
        fmt.Println(v6, v4, v6)
        for fi12 := 0; fi12 < 18; fi12++ {
            fmt.Println(fi12, v11, v4)
            v3 = append(v3, v2)
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    fmt.Println(v9)
    fmt.Println(v6, v4, v9)
    if (32767 == 491) {
        var v13 float32 = float32(3.14159)
        var v14 string = v2
        _ = v13
        _ = v14
    } else {
        v6 = (^(v6))
    }
    if len(v3) > 0 {
        v3[0] = (v2 + v2)
    }
    var v15 float64 = 0.1
    var v16 bool = true
    var v17 float32 = float32(v15)
    var v18 int16 = v9
    v19 := GenS2{F0: 0, F1: 42}
    v20 := GenS2{F0: 1, F1: 32767}
    if false {
        v9 += v9
        if (42 <= 255) {
            if ((v4 << ((511) & 63)) <= (^(v4))) {
                v21 := GenS1{F0: -80, F1: uint16(v6)}
                _ = v21
            } else {
                var v22 string = (v2 + ("\n\t" + "hello"))
                var v23 int16 = v18
                switch 483 {
                case 0:
                    _ = len(v1)
                }
                _ = v22
                _ = v23
            }
            var v24 int64 = -302
            var v25 float32 = v17
            _ = v24
            _ = v25
        } else {
            var v26 string = "foo bar"
            _ = v26
        }
        v15 += v15
    } else {
        for fi27 := 0; fi27 < 3; fi27++ {
            if v16 {
                for fi28 := 0; fi28 < 16; fi28++ {
                    v18 -= v9
                    v2 = v2
                    switch v2 {
                    case "a":
                        v4 -= (^(v4))
                    }
                    _ = fi28
                }
            }
            _ = fi27
        }
        v16 = (genFn0() >= genFn0())
        v29 := GenS3{F0: -86}
        _ = v29
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
