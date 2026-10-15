package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int16 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []uint32{0, 486, 169}
    v3 := make(map[int]int)
    v4 := []string{"a"}
    var v5 uint = 65535
    for fi6 := 0; fi6 < 1; fi6++ {
        v3[fi6] = fi6
        _ = fi6
    }
    var v7 bool = true
    v8 := GenS1{F0: 579, F1: 0, F2: -1}
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 uint16 = 65535
            v11 -= 497
            _ = fi10
            _ = v11
        }
        var v12 int = -552
        _ = fi9
        _ = v12
    }
    var v13 int8 = -72
    if ((32767 > -136) && true) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        switch 128 {
        case -1:
            var v14 int64 = 9223372036854775807
            _ = v14
        }
        var v15 string = "foo bar"
        _ = v15
    }
    v13 -= (v13 + 0)
    for fi16 := 0; fi16 < 8; fi16++ {
        v5 -= 2
        _ = fi16
    }
    v17 := GenS1{F0: 0, F1: 32767, F2: 190}
    var v18 int = 32767
    switch v18 {
    case 0:
        v1 = (v1 % ((v1) | 1))
    case 42:
        var v19 uint8 = 2
        _ = v19
    case 127:
        switch ("\n\t" + "foo bar") {
        case "hello":
            var v20 uint64 = 367
            _ = v20
        default:
            var v21 uint64 = 1
            _ = v21
        }
    }
    if (v7 || (279 != 395)) {
        v22 := v3[v18]
        v3[-2147483648] = v18
        _ = v22
    }
    if v7 {
        _ = len(v3)
    } else {
        _ = len(v2)
    }
    var v23 uint8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
