package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
func genFn0(p0 uint32) string {
    v4 := []int{729, 264, -9223372036854775808, 138}
    switch "foo bar" {
    case "\n\t":
        p0 += 2147483647
    case "a":
        p0 = p0
    default:
        var v5 int16 = 876
        _ = v5
    }
    _ = p0
    _ = v4
    return "a"
}

func main() {
    v1 := []int{-635, -568}
    var v2 int32 = -1
    var v3 int64 = -1
    if len(v1) > 0 {
        v1[0] = -2147483648
    }
    var v6 uint = 137
    v6 -= 310
    v2 += v2
    var v7 int = 127
    _ = len(v1)
    if true {
        fmt.Println(v7, v2)
        var v8 string = "hello"
        _ = v8
    } else {
        var v9 int = (-(v7))
        var v10 uint16 = 1
        _ = v10
        _ = v9
    }
    v3 -= (v3 & (v3 & (^(v3))))
    for fi11 := 0; fi11 < 9; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            _ = len(v1)
            var v13 int16 = 0
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    if (true || (v6 == v6)) {
        switch ("hello" + "\n\t") {
        case "":
            v7 += (^(v7))
        case "a":
            if len(v1) > 0 {
                v1[0] = 128
            }
        }
    }
    v2 = -2147483648
    switch "a" {
    case "\n\t":
        var v14 int = (v7 % ((v7) | 1))
        _ = v14
    case "":
        if (!true) {
            var v15 uint16 = 1
            _ = v15
        }
    case "hello":
        v6 = (v6 * 2)
    }
    switch v7 {
    case 764:
        if (!true) {
            if len(v1) > 0 {
                v1[0] = (v7 / ((-1) | 1))
            }
            v16 := GenS1{F0: 42, F1: ""}
            _ = v16
        }
    default:
        for fi17 := 0; fi17 < 1; fi17++ {
            v3 += v3
            v6 += v6
            var v18 uint64 = 1
            _ = fi17
            _ = v18
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
