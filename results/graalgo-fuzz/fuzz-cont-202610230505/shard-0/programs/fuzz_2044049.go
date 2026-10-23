package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint; F1 uint8; F2 uint16 }

func main() {
    var v1 int8 = 127
    var v2 int = -425
    var v3 int32 = -1
    var v4 int32 = (v3 & (v3 >> ((65535) & 63)))
    v5 := make(map[uint32]string)
    var v6 string = (("a" + "foo bar") + "")
    var v7 string = ("" + "foo bar")
    var v8 uint64 = 384
    delete(v5, 0)
    var v9 float64 = (((1e10 * (0.1 + -2.25)) * (3.14159 * 1.5)) * (0.0 * float64(1.5)))
    var v10 uint64 = 1
    v10 -= 23
    var v11 int16 = -32768
    if false {
        switch ((v7 + (v7 + (("a" + v6) + v6))) + v7) {
        case "":
            v12 := []int{v2, v2, v2, v2}
            _ = v12
        case "\n\t":
            switch -405 {
            case -613:
                v13 := GenS1{F0: 42}
                _ = v13
            }
        }
    } else {
        for fi14 := 0; fi14 < 4; fi14++ {
            if (v10 >= 103) {
                v6 = (v6 + "hello")
                v10 = 18446744073709551615
            }
            v9 = v9
            v8 -= 4294967295
            _ = fi14
        }
    }
    fmt.Println(v2)
    v15 := GenS1{F0: 510}
    for fi16 := 0; fi16 < 19; fi16++ {
        v17 := GenS2{F0: v11}
        v7 = "foo bar"
        v11 += v11
        _ = fi16
        _ = v17
    }
    if (!false) {
        switch (v2 & v2) {
        case 32767:
            for fi18 := 0; fi18 < 4; fi18++ {
                if (((v7 + (v6 + v6)) + "a") == v7) {
                    v19 := GenS3{F0: v4, F1: 163}
                    _ = len(v5)
                    _ = v19
                } else {
                    v8 = v10
                    v20 := GenS2{F0: v11}
                    _ = v20
                }
                v1 = -121
                _ = fi18
            }
        case 1:
            for fi21 := 0; fi21 < 8; fi21++ {
                v22 := GenS3{F0: v3, F1: 32767}
                v23 := GenS2{F0: (v11 * v11)}
                _ = fi21
                _ = v22
                _ = v23
            }
        case -267:
            _ = len(v5)
        default:
            fmt.Println(v2)
        }
        v8 = (^(v8))
        fmt.Println(v11, v2)
    } else {
        var v24 int = int(v3)
        v25 := GenS4{F0: 916, F1: 1, F2: 505}
        v24 += v24
        _ = v24
        _ = v25
    }
    var v26 uint32 = 390
    for fi27 := 0; fi27 < 10; fi27++ {
        v4 += (v3 * v3)
        fmt.Println(v7)
        var v28 int = fi27
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
