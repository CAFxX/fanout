package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint32 }
func genFn0() uint {
    var v6 bool = true
    var v7 float64 = 1e10
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1(p0 string, p1 int8, p2 uint32) int64 {
    p1 += p1
    if false {
        p2 = p2
    } else {
        v8 := GenS1{F0: 4294967295, F1: 32767, F2: p2}
        var v9 float64 = 1e10
        switch "hello" {
        case "hello":
            p1 = (p1 - 127)
        case "":
            switch 1 {
            case 4294967295:
                v10 := GenS1{F0: int(p1), F1: -9223372036854775808, F2: p2}
                _ = v10
            }
        case "a":
            if (true || false) {
                var v11 int64 = 42
                _ = v11
            }
        }
        _ = v8
        _ = v9
    }
    var v12 int8 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 9223372036854775807
}

func main() {
    var v1 float64 = (1.5 / 0.1)
    var v2 uint32 = 42
    v3 := []int{-9223372036854775808, -1}
    var v4 uint8 = 1
    v5 := make(map[int]uint32)
    v4 = (v4 << ((240) & 63))
    if (true && (!(true || (false || ((637 > 42) || false))))) {
        var v13 uint32 = (v2 & (v2 ^ (^(v2))))
        _ = v13
    }
    switch ("hello" + "\n\t") {
    case "hello":
        v14 := GenS1{F0: 759, F1: 42, F2: (v2 >> ((42) & 63))}
        _ = v14
    case "":
        var v15 int32 = 1
        _ = v15
    case "a":
        if (!true) {
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                for fi16 := 0; fi16 < 16; fi16++ {
                    var v17 uint = 0
                    var v18 uint16 = 255
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                if (-732 <= 0) {
                    v1 = (v1 * v1)
                    var v19 uint16 = 42
                    _ = v19
                }
            } else {
                var v20 float64 = -2.25
                _ = v20
            }
        } else {
            if (42 >= 42) {
                v2 = v2
                v1 += v1
                v21 := v5[997]
                _ = v21
            }
        }
    }
    v4 += 26
    switch "\n\t" {
    case "foo bar":
        _ = len(v3)
    default:
        switch 42 {
        case -491:
            v22 := GenS1{F0: 32767, F1: -9223372036854775808, F2: (v2 & (^(v2)))}
            _ = v22
        default:
            v23 := GenS1{F0: 42, F1: 42, F2: (v2 * (v2 << ((308) & 63)))}
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        v4 = (^(v4))
        _ = fi24
    }
    var v25 int = 127
    var v26 float64 = v1
    fmt.Println(v2, v25, v4)
    var v27 int64 = 773
    v28 := GenS1{F0: 128, F1: -53, F2: v2}
    var v29 uint8 = 1
    delete(v5, (v25 % (((-(v25))) | 1)))
    if true {
        fmt.Println(v4, v26)
        v30 := v5[-389]
        _ = v30
    }
    if true {
        v31 := v5[(v25 % ((-980) | 1))]
        var v32 string = "\n\t"
        fmt.Println(v25, v32)
        _ = v31
        _ = v32
    }
    fmt.Println(v2)
    for fi33 := 0; fi33 < 10; fi33++ {
        var v34 int64 = v27
        _ = fi33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
