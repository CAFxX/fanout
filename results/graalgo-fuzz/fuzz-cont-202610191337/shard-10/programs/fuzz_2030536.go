package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 string }
type GenS2 struct { F0 int64; F1 uint16 }

func main() {
    var v1 uint32 = 42
    var v2 int32 = 2147483647
    v3 := make(map[string]string)
    v4 := []uint32{(v1 ^ 2), (^(v1))}
    v5 := []uint32{v1, v1, v1, 42}
    v1 = (v1 * (v1 % ((1) | 1)))
    var v6 bool = (true || (1 <= 2))
    var v7 float64 = (float64(1.5) - 0.0)
    v2 = int32(v1)
    switch "a" {
    case "a":
        v7 = v7
    case "foo bar":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 int64 = 274
            if (((!(732 >= 917)) || false) && true) {
                for fi10 := 0; fi10 < 17; fi10++ {
                    switch "" {
                    case "a":
                        v1 = 0
                    }
                    _ = fi10
                }
                v6 = (v7 <= v7)
                _ = len(v4)
            }
            _ = fi8
            _ = v9
        }
    }
    v6 = ((1 >= 4294967295) || false)
    for fi11 := 0; fi11 < 6; fi11++ {
        fmt.Println(fi11, v1)
        fmt.Println(v7, v6, fi11)
        v3["a"] = ("a" + "\n\t")
        _ = fi11
    }
    v1 += 2147483648
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := GenS1{F0: 255, F1: int64(fi12), F2: ("" + "")}
        _ = fi12
        _ = v13
    }
    var v14 uint8 = 255
    for fi15 := 0; fi15 < 3; fi15++ {
        v16 := GenS1{F0: 819, F1: 531, F2: "\n\t"}
        var v17 float64 = 3.14159
        _ = len(v3)
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        if (!true) {
            _ = len(v5)
            var v19 int64 = 9223372036854775807
            _ = v19
        } else {
            for fi20 := 0; fi20 < 2; fi20++ {
                v3["foo bar"] = "foo bar"
                _ = fi20
            }
            var v21 uint16 = 1
            v22 := GenS2{F0: -1, F1: (^(v21))}
            _ = v21
            _ = v22
        }
        _ = fi18
    }
    var v23 uint8 = v14
    v24 := GenS1{F0: 508, F1: 488, F2: ("hello" + "a")}
    switch "a" {
    case "":
        fmt.Println(v14, v6, v1)
    case "hello":
        switch -817 {
        case 127:
            fmt.Println(v7, v7)
        case 42:
            var v25 float64 = v7
            _ = v25
        default:
            fmt.Println(v7, v1)
        }
    case "\n\t":
        var v26 string = "\n\t"
        _ = v26
    default:
        v27 := GenS1{F0: uint(v1), F1: 9223372036854775807, F2: ((("foo bar" + ("" + "\n\t")) + "") + "")}
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
