package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int8 }
func genFn0(p0 int32) int8 {
    v6 := []int{0, -9223372036854775808, -861, 0}
    for fi7 := 0; fi7 < 12; fi7++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        fmt.Println(fi7, p0)
        _ = fi7
    }
    if (!(true && ((true || false) || true))) {
        if true {
            p0 -= 42
            if true {
                v8 := GenS1{F0: (1 <= 384), F1: true, F2: 117}
                _ = v8
            } else {
                fmt.Println(p0)
            }
        } else {
            var v9 uint8 = 0
            v9 += (v9 - v9)
            _ = v9
        }
        if (42 < 1000) {
            v10 := GenS1{F0: true, F1: false, F2: -128}
            switch 32768 {
            case 32767:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            default:
                v11 := GenS1{F0: (-1 == 1), F1: (-1 <= -9223372036854775808), F2: -97}
                _ = v11
            }
            p0 = p0
            _ = v10
        } else {
            v12 := GenS1{F0: (32767 != 1), F1: (127 >= 128), F2: 86}
            var v13 uint8 = 1
            _ = v12
            _ = v13
        }
        var v14 int16 = -547
        _ = v14
    }
    p0 = p0
    p0 = p0
    _ = p0
    _ = v6
    return 127
}

func main() {
    v1 := []string{"", ("foo bar" + "hello"), "\n\t"}
    var v2 int8 = 42
    var v3 int8 = (v2 >> ((42) & 63))
    var v4 uint = 496
    var v5 uint = (^(v4))
    switch "foo bar" {
    case "\n\t":
        v3 -= 1
    }
    var v15 uint64 = 18446744073709551615
    if len(v1) > 0 {
        v1[0] = ""
    }
    v15 = 42
    v1 = append(v1, "foo bar")
    v16 := GenS1{F0: true, F1: false, F2: v3}
    v15 = 700
    fmt.Println(v4, v3, v15)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
