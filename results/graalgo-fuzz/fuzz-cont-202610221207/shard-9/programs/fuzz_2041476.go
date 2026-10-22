package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }

func main() {
    var v1 int8 = 116
    v2 := []string{}
    v3 := []string{"", (("" + "") + "hello"), "hello"}
    v1 -= 100
    if (42 == 4294967295) {
        switch -9223372036854775808 {
        case -886:
            var v4 int64 = 1
            _ = v4
        case -405:
            var v5 int32 = -1
            _ = v5
        case 42:
            if true {
                var v6 int32 = -2147483648
                v1 = v1
                fmt.Println(v6)
                _ = v6
            }
        default:
            v1 += v1
        }
        v7 := GenS1{F0: -9223372036854775808}
        _ = v7
    } else {
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        v8 := GenS2{F0: 128}
        var v9 int8 = (v1 - (v1 | (-(v1))))
        _ = v8
        _ = v9
    }
    v1 += v1
    _ = len(v3)
    if true {
        _ = len(v2)
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
        var v10 bool = true
        _ = v10
    }
    if len(v3) > 0 {
        v3[0] = ("foo bar" + "")
    }
    fmt.Println(v1)
    var v11 float32 = float32((0.1 + (0.0 - 3.14159)))
    switch -9223372036854775808 {
    case 4294967295:
        fmt.Println(v1, v11)
    case -940:
        if false {
            v11 += ((3.14159 + v11) + 1e10)
            fmt.Println(v11, v11)
            var v12 int64 = -1
            _ = v12
        }
    case 128:
        var v13 string = ""
        _ = v13
    }
    if (!false) {
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        switch "a" {
        case "\n\t":
            fmt.Println(v1, v1)
        case "hello":
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 int8 = v1
                var v16 float64 = (float64(v11) - 1e10)
                _ = fi14
                _ = v15
                _ = v16
            }
        case "a":
            v11 += float32(0.0)
        default:
            var v17 int8 = v1
            _ = v17
        }
    } else {
        var v18 int16 = -130
        _ = v18
    }
    v19 := GenS3{F0: "hello"}
    var v20 int8 = (-(v1))
    var v21 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
