package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint8 }
func genFn0() uint8 {
    v7 := GenS1{F0: -731, F1: 65535}
    var v8 float64 = ((0.1 - 1e10) - 0.0)
    var v9 uint8 = 127
    v10 := []int{}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return v9
}

func main() {
    var v1 int = 32768
    var v2 int16 = -1
    var v3 int = (-(v1))
    var v4 uint16 = 337
    v5 := []uint32{}
    var v6 uint64 = 642
    v11 := GenS1{F0: 1, F1: 2147483647}
    if ((642 != 192) || true) {
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 uint32 = 217
            _ = fi12
            _ = v13
        }
        v14 := GenS1{F0: 0, F1: 2147483648}
        v3 += 4294967295
        _ = v14
    } else {
        for fi15 := 0; fi15 < 1; fi15++ {
            v6 -= v6
            _ = fi15
        }
        switch 32 {
        case 32767:
            if (false || true) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            } else {
                if ((0.0 - 0.0) >= 0.1) {
                    if false {
                        var v16 uint8 = 2
                        v2 = 1000
                        fmt.Println(v1, v6, v6)
                        _ = v16
                    } else {
                        v17 := GenS1{F0: 0, F1: 458}
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 20; fi18++ {
                        var v19 uint = 18446744073709551615
                        v3 = int(v6)
                        fmt.Println(v1)
                        _ = fi18
                        _ = v19
                    }
                }
                var v20 string = (("foo bar" + "\n\t") + "foo bar")
                _ = v20
            }
        }
    }
    switch "foo bar" {
    case "a":
        var v21 int16 = v2
        _ = v21
    }
    v22 := GenS1{F0: -493, F1: 2147483647}
    var v23 string = (("a" + (("\n\t" + "foo bar") + "foo bar")) + "a")
    switch (v23 + v23) {
    case "hello":
        if len(v5) > 0 {
            v5[0] = 2147483648
        }
    case "":
        v4 = (v4 & 365)
    case "\n\t":
        fmt.Println(v1)
    }
    var v24 bool = false
    _ = len(v5)
    for fi25 := 0; fi25 < 17; fi25++ {
        if len(v5) > 0 {
            v5[0] = 4294967295
        }
        _ = fi25
    }
    _ = len(v5)
    _ = len(v5)
    for fi26 := 0; fi26 < 16; fi26++ {
        v27 := GenS2{F0: 255}
        _ = fi26
        _ = v27
    }
    v1 -= (^(v3))
    switch "\n\t" {
    case "\n\t":
        for fi28 := 0; fi28 < 12; fi28++ {
            for fi29 := 0; fi29 < 13; fi29++ {
                fmt.Println(v23, fi28, v24)
                v2 = 42
                if len(v5) > 0 {
                    v5[0] = 2147483647
                }
                _ = fi29
            }
            var v30 bool = (genFn0() == 156)
            var v31 uint64 = (v6 & v6)
            _ = fi28
            _ = v30
            _ = v31
        }
    default:
        fmt.Println(v23)
    }
    v32 := GenS1{F0: 895, F1: 406}
    var v33 float32 = (0.0 - ((3.14159 / (float32(1e10) + -2.25)) * 0.0))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
