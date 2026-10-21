package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int64; F1 string; F2 int }

func main() {
    var v1 bool = true
    v2 := []int{-2147483648, -9223372036854775808}
    v3 := make(map[uint32]string)
    var v4 int8 = 42
    v5 := make(map[string]int)
    v4 -= (^(v4))
    var v6 uint16 = 1
    v6 = (v6 >> ((18446744073709551615) & 63))
    v6 = (v6 | (v6 * v6))
    v7 := GenS1{F0: 0, F1: ""}
    var v8 int = 32768
    for fi9 := 0; fi9 < 6; fi9++ {
        fmt.Println(v1)
        switch (v8 + v8) {
        case 127:
            switch (fi9 << ((1) & 63)) {
            case 32768:
                if (!(1 >= 1)) {
                    v1 = (1e10 <= 1e10)
                    v4 += (v4 ^ 0)
                } else {
                    fmt.Println(v4)
                    if true {
                        var v10 uint8 = 174
                        v1 = (v6 != 255)
                        var v11 float32 = ((((3.14159 + 1e10) / float32(1.5)) * 0.1) + 0.1)
                        _ = v10
                        _ = v11
                    } else {
                        fmt.Println(v1, fi9)
                        fmt.Println(v8, v6)
                        var v12 uint64 = 4294967295
                        _ = v12
                    }
                    v6 = v6
                }
            case 64:
                v4 += v4
            }
        }
        _ = fi9
    }
    var v13 uint16 = (v6 + v6)
    v2 = append(v2, v8)
    if true {
        fmt.Println(v6)
        delete(v3, 2147483647)
    } else {
        fmt.Println(v6, v8)
    }
    v5[""] = v8
    var v14 float64 = ((-2.25 * 1e10) + 3.14159)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch (((("a" + "foo bar") + (("\n\t" + "hello") + "a")) + (("hello" + "\n\t") + ("hello" + "a"))) + "hello") {
    case "foo bar":
        if v1 {
            fmt.Println(v1)
            switch "foo bar" {
            case "\n\t":
                v1 = true
            case "hello":
                var v15 string = "\n\t"
                _ = v15
            }
            v1 = v1
        }
    case "hello":
        v16 := GenS2{F0: -609, F1: "foo bar", F2: 481}
        _ = v16
    case "\n\t":
        if v1 {
            v13 = v13
        }
    default:
        fmt.Println(v13, v6, v8)
    }
    if (((128 != 254) || (!((v8 - v8) == (v8 ^ (v8 | v8))))) && true) {
        var v17 int64 = 33
        switch (-(v8)) {
        case 749:
            v4 = -21
        case 2147483647:
            v18 := GenS1{F0: (v4 + 100), F1: ""}
            _ = v18
        }
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
