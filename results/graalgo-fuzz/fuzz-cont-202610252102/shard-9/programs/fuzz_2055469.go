package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int }
func genFn0() int16 {
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 string = "\n\t"
        if true {
            fmt.Println(v8)
        }
        _ = fi7
        _ = v8
    }
    var v9 int32 = 974
    var v10 int16 = 276
    _ = v10
    _ = v9
    return v10
}

func main() {
    var v1 uint = 91
    var v2 uint = (v1 | (v1 + 442))
    var v3 int8 = 42
    v4 := make(map[string]string)
    var v5 bool = false
    v6 := []string{"\n\t", ((("" + "\n\t") + ("a" + ("\n\t" + "\n\t"))) + "a"), "\n\t"}
    v11 := GenS1{F0: ((!false) || v5)}
    var v12 uint8 = 255
    v1 -= (v2 - (^(v2)))
    var v13 uint64 = 420
    fmt.Println(v3, v1)
    v14 := GenS2{F0: 2147483647, F1: -9223372036854775808}
    v1 -= (v1 * uint(v12))
    var v15 uint16 = 255
    if (791 < 42) {
        v2 = (v1 | 2)
    } else {
        v2 -= 65535
        for fi16 := 0; fi16 < 2; fi16++ {
            if v5 {
                v13 = (v13 >> (((v1 * 65535)) & 63))
                v5 = (genFn0() > 1)
            }
            var v17 int16 = -247
            _ = fi16
            _ = v17
        }
    }
    if len(v6) > 0 {
        v6[0] = ""
    }
    if ((2147483647 <= 1) && true) {
        v3 += v3
        v15 += 337
        switch 128 {
        case 0:
            var v18 uint16 = (v15 & 926)
            _ = v18
        case -1:
            var v19 float64 = 1.5
            _ = v19
        case 93:
            fmt.Println(v5)
        default:
            for fi20 := 0; fi20 < 5; fi20++ {
                v2 += (v2 << (((v2 ^ v1)) & 63))
                _ = fi20
            }
        }
    } else {
        var v21 string = ("foo bar" + "hello")
        var v22 float32 = 0.1
        if (!(v22 >= v22)) {
            switch -114 {
            case -2147483648:
                var v23 int = 32768
                _ = v23
            case 4294967295:
                v22 += 1e10
            default:
                var v24 int8 = v3
                _ = v24
            }
        } else {
            v15 += v15
        }
        _ = v21
        _ = v22
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        var v26 uint64 = 545
        v3 = v3
        _ = fi25
        _ = v26
    }
    delete(v4, "foo bar")
    for fi27 := 0; fi27 < 7; fi27++ {
        fmt.Println(v1, v1)
        _ = fi27
    }
    switch ("foo bar" + "") {
    case "\n\t":
        var v28 uint16 = (v15 >> ((v1) & 63))
        _ = v28
    case "a":
        _ = len(v4)
    }
    v6 = append(v6, "a")
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
