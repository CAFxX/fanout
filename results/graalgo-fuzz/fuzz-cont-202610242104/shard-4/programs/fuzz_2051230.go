package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 bool; F2 uint64 }

func main() {
    var v1 float32 = 0.1
    v2 := []int{}
    var v3 int32 = -2147483648
    var v4 int64 = 42
    v5 := GenS1{F0: true}
    var v6 int = -263
    fmt.Println(v3)
    v1 = 1e10
    for fi7 := 0; fi7 < 6; fi7++ {
        switch -319 {
        case 127:
            var v8 uint16 = 32767
            _ = v8
        }
        switch "a" {
        case "a":
            fmt.Println(v1, v4, v4)
        case "hello":
            var v9 float32 = (v1 * 0.1)
            _ = v9
        case "foo bar":
            var v10 uint16 = 930
            _ = v10
        }
        _ = fi7
    }
    v11 := GenS1{F0: ("hello" >= "")}
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := GenS1{F0: true}
        _ = fi12
        _ = v13
    }
    if (!((1.5 != ((v1 + (((v1 * 0.1) * (-2.25 + v1)) / float32(1e10))) + v1)) || (883 >= 42))) {
        v3 -= -1
        fmt.Println(v3, v6)
        var v14 string = "hello"
        _ = v14
    }
    if true {
        v15 := GenS1{F0: (2147483647 <= 2147483647)}
        v6 += (v6 % ((v6) | 1))
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 bool = false
            v17 = false
            var v18 uint32 = 2
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = v15
    } else {
        v19 := GenS1{F0: (v3 < v3)}
        v1 = float32(((-2.25 - 1.5) - (1e10 / (((0.0 - ((0.1 - -2.25) + (float64(v1) / 1.0))) * 1e10) / 0.1))))
        _ = v19
    }
    v3 = (v3 / ((-2147483648) | 1))
    if false {
        v2 = append(v2, -2147483648)
        switch ("" + ((("hello" + ("a" + ("\n\t" + ""))) + ("hello" + "a")) + "\n\t")) {
        case "":
            v4 += v4
        case "\n\t":
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                for fi20 := 0; fi20 < 8; fi20++ {
                    switch ("foo bar" + "") {
                    case "hello":
                        var v21 int64 = (^(v4))
                        _ = v21
                    case "\n\t":
                        v6 = (v6 / ((fi20) | 1))
                    default:
                        var v22 string = "\n\t"
                        _ = v22
                    }
                    if (v6 != fi20) {
                        v4 = 155
                        _ = len(v2)
                    }
                    fmt.Println(v3, v4, fi20)
                    _ = fi20
                }
            }
        case "foo bar":
            switch "\n\t" {
            case "":
                var v23 int32 = (v3 >> ((921) & 63))
                _ = v23
            }
        }
        var v24 uint64 = 3
        _ = v24
    } else {
        var v25 int16 = 42
        v25 -= 1000
        _ = v25
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        switch 2147483647 {
        case 1:
            if false {
                v3 -= (-(v3))
            } else {
                _ = len(v2)
            }
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi26
    }
    v27 := GenS2{F0: "\n\t", F1: (-270 >= 1), F2: 818}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
