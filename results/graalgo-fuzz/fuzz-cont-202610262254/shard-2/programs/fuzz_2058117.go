package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]int)
    var v3 uint64 = 657
    v4 := make(map[string]string)
    switch "" {
    case "":
        var v5 uint32 = 2
        _ = v5
    case "\n\t":
        v6 := GenS1{F0: -1, F1: -2147483648}
        _ = v6
    default:
        if ((91 != 100) || true) {
            v7 := []int{}
            v3 += (v3 ^ v3)
            _ = v7
        }
    }
    var v8 uint8 = 0
    switch "foo bar" {
    case "hello":
        switch ("hello" + ("foo bar" + "hello")) {
        case "foo bar":
            var v9 uint8 = (v8 >> ((4294967295) & 63))
            _ = v9
        case "hello":
            v10 := v1[692]
            _ = v10
        default:
            switch ("foo bar" + "a") {
            case "":
                var v11 uint8 = 2
                _ = v11
            case "a":
                v12 := []int{-472}
                _ = v12
            case "\n\t":
                for fi13 := 0; fi13 < 2; fi13++ {
                    v14 := []int{}
                    fmt.Println(fi13, v8)
                    _ = fi13
                    _ = v14
                }
            default:
                var v15 string = (("foo bar" + "a") + ("a" + "hello"))
                _ = v15
            }
        }
    }
    if true {
        fmt.Println(v8, v3)
    } else {
        fmt.Println(v3, v3)
        var v16 float64 = -2.25
        _ = v16
    }
    v8 = (v8 * (^(v8)))
    var v17 int8 = -1
    switch 4294967295 {
    case -494:
        if false {
            v17 = (v17 + 127)
        }
    case 128:
        for fi18 := 0; fi18 < 6; fi18++ {
            var v19 int32 = -1
            _ = fi18
            _ = v19
        }
    case 0:
        v20 := []int{128, int(v3)}
        _ = v20
    }
    var v21 int64 = 9223372036854775807
    var v22 uint8 = (^(v8))
    v23 := v1[-1]
    fmt.Println(v21, v8)
    var v24 uint32 = 1
    v25 := GenS1{F0: v17, F1: -1}
    var v26 uint8 = (^(v22))
    var v27 uint = 4294967295
    v28 := GenS1{F0: (^(v17)), F1: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v8
}
