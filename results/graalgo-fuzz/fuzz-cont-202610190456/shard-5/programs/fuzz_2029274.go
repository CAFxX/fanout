package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint32; F2 bool }
type GenS3 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 0
    var v3 int8 = 1
    v4 := []uint32{}
    var v5 string = ("\n\t" + "foo bar")
    var v6 float32 = (1.5 * 1.5)
    v5 = (v5 + ((v5 + "\n\t") + ("foo bar" + v5)))
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := GenS1{F0: 1000, F1: -1, F2: 0}
            v5 = v5
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v10 uint8 = 42
    fmt.Println(v10, v10, v3)
    var v11 uint16 = (^(v2))
    _ = len(v1)
    switch ("" + "hello") {
    case "a":
        var v12 int32 = -602
        _ = v12
    case "\n\t":
        var v13 int8 = 1
        _ = v13
    }
    switch (("a" + v5) + v5) {
    case "a":
        v4 = append(v4, 699)
    }
    fmt.Println(v3, v11, v6)
    v1[4294967295] = 345
    v11 = v11
    var v14 string = v5
    switch ((((v14 + ((v14 + "a") + (((v14 + v14) + "") + v14))) + v5) + v14) + "\n\t") {
    case "":
        var v15 int64 = 1
        _ = v15
    case "foo bar":
        v4 = append(v4, 4294967295)
    case "\n\t":
        var v16 uint32 = 487
        _ = v16
    default:
        fmt.Println(v11)
    }
    if (!(1 == v11)) {
        var v17 uint64 = 2
        v18 := v1[65535]
        _ = v17
        _ = v18
    }
    if (!(int(v10) >= 404)) {
        if true {
            _ = len(v1)
            var v19 float32 = float32(-2.25)
            _ = v19
        } else {
            switch 128 {
            case 9223372036854775807:
                for fi20 := 0; fi20 < 14; fi20++ {
                    v21 := GenS2{F0: 0, F1: 4294967295, F2: (1.5 == (((-2.25 / v6) + v6) * v6))}
                    _ = fi20
                    _ = v21
                }
            case -24:
                var v22 uint = 2
                _ = v22
            default:
                if (((v3 + v3) <= (v3 << ((170) & 63))) || false) {
                    if false {
                        v23 := GenS1{F0: 1, F1: 1, F2: 1}
                        var v24 uint32 = 52
                        _ = v23
                        _ = v24
                    } else {
                        v25 := GenS3{F0: 1}
                        _ = v25
                    }
                }
            }
            v6 += -2.25
            fmt.Println(v5)
        }
    }
    v26 := v1[91]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
