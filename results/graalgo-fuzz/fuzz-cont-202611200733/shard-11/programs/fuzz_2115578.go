package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }

func main() {
    var v1 string = "foo bar"
    var v2 uint32 = 65535
    var v3 uint8 = 127
    var v4 uint16 = 607
    v5 := []int{32768, 128}
    v6 := make(map[uint32]string)
    if (209 > 300) {
        v7 := GenS1{F0: v2, F1: false}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if true {
            v1 = (v1 + v1)
            v4 -= (^(v4))
            _ = len(v6)
        } else {
            if (0 >= 2147483647) {
                v8 := GenS1{F0: 62, F1: (v3 != v3)}
                _ = v8
            }
        }
        _ = v7
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(v4, v3)
        _ = fi9
    }
    switch 42 {
    case 575:
        switch v1 {
        case "\n\t":
            if false {
                var v10 uint64 = 194
                var v11 int = -311
                if false {
                    v12 := GenS1{F0: v2, F1: true}
                    v13 := GenS1{F0: v2, F1: (true || false)}
                    var v14 uint32 = v2
                    _ = v12
                    _ = v13
                    _ = v14
                }
                _ = v10
                _ = v11
            } else {
                if false {
                    var v15 uint = 1
                    _ = v15
                } else {
                    v6[v2] = ""
                }
                v16 := GenS1{F0: 2, F1: (float32(0.0) < ((1e10 - 0.1) - 0.1))}
                _ = v16
            }
        case "":
            var v17 uint8 = v3
            _ = v17
        case "a":
            v18 := GenS1{F0: (v2 ^ 109), F1: (!(!true))}
            _ = v18
        default:
            switch "a" {
            case "":
                var v19 int8 = -12
                _ = v19
            case "hello":
                var v20 uint32 = (v2 >> ((690) & 63))
                _ = v20
            }
        }
    case 2147483647:
        v6[v2] = "hello"
    default:
        switch 712 {
        case 0:
            for fi21 := 0; fi21 < 7; fi21++ {
                var v22 float64 = 0.1
                v22 += 0.0
                var v23 int64 = -952
                _ = fi21
                _ = v22
                _ = v23
            }
        }
    }
    v5 = append(v5, 406)
    _ = len(v6)
    if false {
        switch v1 {
        case "":
            _ = len(v5)
        case "a":
            if (!(!(((^(v4)) < v4) && (1 == 42)))) {
                var v24 uint16 = v4
                _ = v24
            } else {
                v25 := GenS1{F0: v2, F1: (!(((2 <= 265) || true) && (!(1 != 0))))}
                _ = v25
            }
        case "foo bar":
            fmt.Println(v2, v2, v2)
        }
        var v26 uint8 = 77
        _ = v26
    } else {
        for fi27 := 0; fi27 < 1; fi27++ {
            v3 += v3
            if (368 != 65535) {
                v28 := GenS1{F0: 4294967295, F1: (0 == (v3 - (v3 | 42)))}
                _ = v28
            }
            var v29 int16 = 25
            _ = fi27
            _ = v29
        }
    }
    if false {
        switch 32767 {
        case 378:
            switch 2147483647 {
            case 2147483647:
                fmt.Println(v2, v4)
            default:
                for fi30 := 0; fi30 < 8; fi30++ {
                    _ = len(v5)
                    _ = fi30
                }
            }
        case 42:
            var v31 int32 = 42
            _ = v31
        }
        v4 += (v4 & v4)
    }
    for fi32 := 0; fi32 < 2; fi32++ {
        fmt.Println(v3, v2)
        switch "a" {
        case "\n\t":
            var v33 float32 = float32(0.0)
            _ = v33
        case "a":
            v4 = 0
        }
        _ = fi32
    }
    v2 = (v2 % ((v2) | 1))
    v1 = (("\n\t" + v1) + ("hello" + "a"))
    var v34 float64 = 1.5
    if false {
        _ = len(v6)
        v35 := GenS1{F0: (v2 / ((v2) | 1)), F1: true}
        _ = v35
    }
    v2 += 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
