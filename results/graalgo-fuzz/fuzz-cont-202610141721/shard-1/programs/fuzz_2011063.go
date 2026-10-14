package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 90
    v3 := make(map[int]uint32)
    v4 := make(map[string]string)
    var v5 uint8 = 42
    v2 += v2
    v5 = 1
    var v6 string = ""
    fmt.Println(v6)
    var v7 int = 32768
    delete(v4, "foo bar")
    if (!(0 != 1)) {
        v2 = v2
        if true {
            v2 += v2
            if (!(v5 == 42)) {
                var v8 string = "foo bar"
                fmt.Println(v7, v6)
                v9 := GenS1{F0: (v5 | 1), F1: v5}
                _ = v8
                _ = v9
            } else {
                var v10 float64 = 1.5
                _ = v10
            }
        }
        var v11 uint8 = (v5 - (^(v5)))
        _ = v11
    } else {
        v5 = (v5 & (^(v5)))
        v12 := v1[v6]
        _ = v12
    }
    var v13 int32 = 0
    var v14 float64 = 1.5
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 int64 = -9223372036854775808
        _ = fi15
        _ = v16
    }
    if (42 == -677) {
        v17 := GenS1{F0: v5, F1: (v5 ^ v5)}
        v2 += (v2 >> ((0) & 63))
        _ = v17
    }
    switch (v7 >> ((42) & 63)) {
    case 42:
        for fi18 := 0; fi18 < 16; fi18++ {
            if true {
                var v19 int64 = 0
                if (false && ((v5 * 42) > (v5 / ((28) | 1)))) {
                    for fi20 := 0; fi20 < 8; fi20++ {
                        v1["\n\t"] = 2147483648
                        _ = fi20
                    }
                    v13 = (v13 | (v13 / (((-(v13))) | 1)))
                }
                _ = v19
            } else {
                v6 = ((v6 + "foo bar") + ("foo bar" + "a"))
            }
            _ = fi18
        }
    default:
        v2 = (v2 >> ((65535) & 63))
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 uint16 = 255
        switch fi21 {
        case -734:
            if ((!true) && true) {
                if false {
                    v23 := GenS1{F0: (v5 << ((2) & 63)), F1: v5}
                    _ = v23
                }
            } else {
                var v24 int16 = 0
                v5 += 188
                _ = v24
            }
        }
        _ = fi21
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
