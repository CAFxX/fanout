package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
func genFn0() int16 {
    var v5 int8 = 100
    v6 := []int{-1, 9223372036854775807}
    var v7 int64 = 258
    var v8 uint = 770
    v9 := GenS1{F0: 0, F1: 64}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{"foo bar", ("hello" + ("foo bar" + "hello")), "foo bar"}
    var v3 int = 0
    var v4 float32 = 1.5
    var v10 uint16 = 255
    v10 += 933
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 bool = false
        _ = fi11
        _ = v12
    }
    v13 := v1[(-(v3))]
    v14 := GenS1{F0: 959, F1: 62}
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 int64 = 272
        var v17 float64 = 0.0
        _ = fi15
        _ = v16
        _ = v17
    }
    switch v3 {
    case 32768:
        _ = len(v2)
    default:
        v4 = v4
    }
    var v18 uint8 = 127
    for fi19 := 0; fi19 < 20; fi19++ {
        v20 := GenS1{F0: -623, F1: -1}
        v13 -= (fi19 & v13)
        _ = fi19
        _ = v20
    }
    v13 += v13
    v1[(v13 & v3)] = v3
    if (!false) {
        v13 = v3
        v21 := GenS1{F0: 32767, F1: -61}
        switch "" {
        case "":
            var v22 uint16 = v10
            _ = v22
        case "a":
            if ((!(false && true)) || true) {
                var v23 uint32 = 1
                _ = v23
            } else {
                fmt.Println(v18, v10, v18)
                v24 := GenS1{F0: 32767, F1: -15}
                var v25 uint32 = 65535
                _ = v24
                _ = v25
            }
        }
        _ = v21
    }
    v10 += v10
    if false {
        switch "hello" {
        case "\n\t":
            _ = len(v2)
        case "a":
            switch "\n\t" {
            case "a":
                switch -184 {
                case 2147483647:
                    var v26 float32 = (((0.1 + (-2.25 + v4)) * (v4 + 1e10)) - v4)
                    _ = v26
                case 997:
                    _ = len(v2)
                default:
                    v27 := GenS1{F0: -202, F1: 100}
                    _ = v27
                }
            }
        case "hello":
            v1[v13] = (v3 % ((v13) | 1))
        default:
            fmt.Println(v13)
        }
    }
    _ = len(v2)
    if (100 != 641) {
        v28 := GenS1{F0: 1000, F1: 0}
        switch ("a" + ("foo bar" + "\n\t")) {
        case "\n\t":
            var v29 string = ("a" + "hello")
            _ = v29
        }
        _ = v28
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v30 := GenS1{F0: -32768, F1: 100}
        var v31 int8 = -1
        _ = v30
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
