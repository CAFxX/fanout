package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int32; F2 int64 }

func main() {
    v1 := []int{0, 423, 4294967295}
    var v2 float64 = ((0.0 - 3.14159) - 1.5)
    var v3 uint32 = 790
    v4 := GenS1{F0: (("hello" + ((("hello" + "\n\t") + (("\n\t" + "\n\t") + "a")) + "a")) + "")}
    switch "hello" {
    case "foo bar":
        switch 127 {
        case 32767:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    case "\n\t":
        var v5 int64 = 9223372036854775807
        _ = v5
    default:
        v3 -= 754
    }
    switch 32768 {
    case -2147483648:
        v6 := GenS1{F0: ("hello" + "hello")}
        _ = v6
    }
    switch 854 {
    case -9223372036854775808:
        _ = len(v1)
    }
    v2 += -2.25
    if (!(true || (2 >= 341))) {
        v2 = v2
        var v7 int64 = -9223372036854775808
        v2 -= v2
        _ = v7
    }
    switch "foo bar" {
    case "":
        fmt.Println(v2, v3)
    case "\n\t":
        for fi8 := 0; fi8 < 18; fi8++ {
            v1 = append(v1, 127)
            _ = fi8
        }
    case "hello":
        if (-1 == -1) {
            var v9 int32 = -262
            var v10 uint64 = 4294967295
            v11 := GenS1{F0: "\n\t"}
            _ = v10
            _ = v11
            _ = v9
        }
    }
    var v12 uint = 255
    v2 = 1.5
    var v13 float64 = v2
    for fi14 := 0; fi14 < 20; fi14++ {
        v15 := GenS1{F0: ("a" + "")}
        v2 -= (v13 - ((v13 * v13) / float64(1e10)))
        v1 = append(v1, (fi14 - 42))
        _ = fi14
        _ = v15
    }
    var v16 bool = true
    _ = len(v1)
    v17 := GenS2{F0: 1, F1: -693, F2: 0}
    for fi18 := 0; fi18 < 20; fi18++ {
        fmt.Println(v12, v2, v16)
        _ = fi18
    }
    v13 -= 0.1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
