package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string }

func main() {
    var v1 uint8 = 117
    var v2 int32 = 2147483647
    var v3 int16 = -32768
    v4 := make(map[uint32]int)
    if (638 >= 631) {
        switch ("hello" + "\n\t") {
        case "":
            fmt.Println(v2, v2, v2)
        case "hello":
            v1 -= v1
        case "foo bar":
            v4[11] = 32768
        default:
            v3 = -1
        }
    } else {
        v2 = (v2 | (v2 << ((90) & 63)))
        var v5 int8 = 127
        var v6 int8 = 42
        _ = v5
        _ = v6
    }
    var v7 int16 = v3
    v3 -= (v7 + 1000)
    if true {
        var v8 uint32 = 2147483648
        fmt.Println(v7)
        _ = v8
    }
    if (0 >= 65535) {
        v7 = v3
        v9 := GenS1{F0: -860, F1: -1, F2: 9223372036854775807}
        _ = v9
    }
    var v10 uint8 = (v1 + (v1 << ((2) & 63)))
    var v11 string = ("" + "")
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 int = fi12
        v14 := v4[382]
        _ = fi12
        _ = v13
        _ = v14
    }
    v3 += v7
    switch -255 {
    case -2147483648:
        for fi15 := 0; fi15 < 2; fi15++ {
            if ((!(v11 >= ((("a" + "") + (v11 + "")) + ""))) || false) {
                v16 := []int{(fi15 % (((fi15 + (fi15 ^ (fi15 - 2147483647)))) | 1)), fi15, (fi15 + (fi15 ^ fi15)), int(v2)}
                v17 := GenS2{F0: 1, F1: 255}
                _ = v16
                _ = v17
            }
            v3 -= (v3 << ((1) & 63))
            for fi18 := 0; fi18 < 11; fi18++ {
                v19 := GenS3{F0: v2}
                var v20 uint8 = (v10 >> ((769) & 63))
                v21 := GenS3{F0: int32(v3)}
                v2 = v2
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            _ = fi15
        }
    default:
        v22 := GenS3{F0: (v2 & (v2 - (v2 * v2)))}
        _ = v22
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        for fi24 := 0; fi24 < 1; fi24++ {
            var v25 uint = 330
            var v26 int64 = 1
            v27 := GenS1{F0: int16(v25), F1: v26, F2: fi23}
            v3 += v3
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
        _ = fi23
    }
    v28 := GenS4{F0: v11}
    for fi29 := 0; fi29 < 20; fi29++ {
        var v30 uint8 = v10
        if true {
            var v31 float32 = 3.14159
            var v32 int = (fi29 + fi29)
            _ = v31
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    var v33 uint64 = 2
    v34 := []int{189, int(v10), 32768, 42}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v7
}
