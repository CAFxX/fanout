package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
func genFn0() int8 {
    v6 := []int{475, 42}
    v7 := GenS1{F0: 1, F1: 185}
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := []int{-2147483648, 128, -1, 562}
    var v2 uint16 = 509
    var v3 int32 = 218
    v4 := []uint32{361, 2, 995, 2}
    v5 := make(map[string]bool)
    var v8 string = "\n\t"
    var v9 int64 = 848
    for fi10 := 0; fi10 < 11; fi10++ {
        switch "\n\t" {
        case "a":
            if false {
                var v11 int16 = 32767
                _ = v11
            }
        case "":
            if (4294967295 > 42) {
                v2 += (v2 + v2)
                delete(v5, v8)
                if true {
                    var v12 int64 = v9
                    _ = v12
                } else {
                    var v13 int32 = (v3 << ((1) & 63))
                    var v14 uint16 = 82
                    _ = v13
                    _ = v14
                }
            }
        }
        var v15 int16 = 32767
        v3 += 42
        _ = fi10
        _ = v15
    }
    v3 = v3
    _ = len(v5)
    v5[((v8 + v8) + "foo bar")] = (721 <= 2)
    if true {
        var v16 string = ((v8 + v8) + "foo bar")
        _ = v16
    }
    switch (v8 + v8) {
    case "foo bar":
        var v17 int64 = v9
        _ = v17
    default:
        v5[v8] = true
    }
    v9 = (v9 << ((662) & 63))
    if true {
        v2 = v2
    } else {
        switch 32767 {
        case 4294967295:
            for fi18 := 0; fi18 < 1; fi18++ {
                var v19 uint32 = 65535
                _ = fi18
                _ = v19
            }
        case -731:
            fmt.Println(v9)
        case -392:
            if len(v4) > 0 {
                v4[0] = 1
            }
        }
        for fi20 := 0; fi20 < 20; fi20++ {
            v21 := v5[("" + v8)]
            fmt.Println(v2, v8)
            _ = fi20
            _ = v21
        }
        v22 := GenS1{F0: -1, F1: 9223372036854775807}
        _ = v22
    }
    var v23 uint8 = 1
    for fi24 := 0; fi24 < 9; fi24++ {
        v4 = append(v4, 1)
        _ = fi24
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        v26 := GenS1{F0: -32768, F1: -296}
        _ = fi25
        _ = v26
    }
    switch 1 {
    case 42:
        if ((0 == 0) && false) {
            for fi27 := 0; fi27 < 17; fi27++ {
                var v28 int16 = 42
                if (0 < 52) {
                    v8 = "a"
                }
                v23 -= (v23 % ((128) | 1))
                _ = fi27
                _ = v28
            }
            v8 = (((("" + (v8 + v8)) + "hello") + ((v8 + (v8 + ("foo bar" + ("foo bar" + v8)))) + (v8 + ((v8 + ("hello" + v8)) + v8)))) + v8)
        }
    case 4294967295:
        v2 -= v2
    }
    _ = len(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
