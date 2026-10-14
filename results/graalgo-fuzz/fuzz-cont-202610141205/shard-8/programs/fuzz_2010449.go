package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint8; F2 uint64 }

func main() {
    v1 := []int{-82, 9223372036854775807, 107}
    v2 := make(map[string]bool)
    v3 := make(map[int]string)
    var v4 uint8 = 14
    v5 := GenS1{F0: ("" + "foo bar")}
    var v6 int8 = -17
    v7 := GenS1{F0: "a"}
    v6 -= v6
    v6 = (v6 % ((42) | 1))
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 bool = (1 <= 377)
        v4 += (v4 - v4)
        v4 = 42
        _ = fi8
        _ = v9
    }
    fmt.Println(v4, v6)
    var v10 int64 = 42
    v11 := GenS2{F0: 127, F1: 2, F2: 2}
    switch (("hello" + ("" + "")) + (("foo bar" + "\n\t") + "a")) {
    case "hello":
        switch 32768 {
        case 657:
            fmt.Println(v10, v10)
        case 127:
            if true {
                for fi12 := 0; fi12 < 9; fi12++ {
                    v6 -= (v6 >> ((79) & 63))
                    _ = fi12
                }
                if false {
                    v10 -= (v10 / ((-1) | 1))
                    v2["foo bar"] = ((true && (-2147483648 <= 42)) && (!(v10 < (v10 - v10))))
                    var v13 int = 0
                    _ = v13
                }
                fmt.Println(v4)
            }
        }
    case "foo bar":
        delete(v2, "a")
    case "":
        if (32767 > 512) {
            v14 := GenS2{F0: 32768, F1: (v4 & v4), F2: 18446744073709551615}
            if (65535 >= 18446744073709551615) {
                var v15 float64 = 3.14159
                _ = v15
            } else {
                var v16 int16 = 851
                v10 = v10
                _ = v16
            }
            switch "\n\t" {
            case "":
                var v17 int = 296
                _ = v17
            default:
                var v18 string = "\n\t"
                _ = v18
            }
            _ = v14
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
