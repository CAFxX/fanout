package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 string }
type GenS2 struct { F0 uint; F1 uint16; F2 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 int64 }
func genFn0() uint64 {
    for fi5 := 0; fi5 < 18; fi5++ {
        if (!((!false) || false)) {
            switch (fi5 << ((770) & 63)) {
            case 642:
                var v6 uint16 = 255
                _ = v6
            case 0:
                fmt.Println(fi5)
            default:
                var v7 int16 = 0
                _ = v7
            }
            var v8 uint16 = 65535
            _ = v8
        }
        for fi9 := 0; fi9 < 10; fi9++ {
            v10 := GenS1{F0: "foo bar", F1: (^(fi9)), F2: ("" + "hello")}
            switch "a" {
            case "a":
                v11 := []int{fi9}
                _ = v11
            default:
                v12 := GenS1{F0: "a", F1: (fi9 % (((fi9 >> ((1) & 63))) | 1)), F2: (("a" + "\n\t") + "a")}
                _ = v12
            }
            _ = fi9
            _ = v10
        }
        _ = fi5
    }
    var v13 bool = true
    fmt.Println(v13)
    v14 := GenS2{F0: 2, F1: 0, F2: 765}
    _ = v13
    _ = v14
    return 0
}

func main() {
    var v1 float32 = (-2.25 + (-2.25 + float32(0.0)))
    v2 := []uint32{1, 1, 2147483648, 61}
    var v3 int64 = 136
    v4 := make(map[string]uint32)
    switch 32767 {
    case 431:
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 float64 = (((0.0 / 1e10) + 1.5) / 1.0)
            _ = fi15
            _ = v16
        }
    case -149:
        fmt.Println(v1)
    case 42:
        v1 += v1
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println(v3, v3)
    if true {
        switch 32768 {
        case 2147483647:
            fmt.Println(v3)
        case -1:
            delete(v4, "hello")
        case 42:
            fmt.Println(v1, v3)
        default:
            var v17 string = "hello"
            _ = v17
        }
        var v18 int32 = 400
        var v19 float64 = (float64(0.0) / 0.1)
        _ = v18
        _ = v19
    } else {
        v20 := GenS3{F0: 606}
        fmt.Println(v1, v3)
        v1 += v1
        _ = v20
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        fmt.Println(v1)
        _ = fi21
    }
    v22 := GenS2{F0: 258, F1: 411, F2: 1}
    v23 := GenS3{F0: -9223372036854775808}
    var v24 uint8 = 2
    v25 := GenS4{F0: "", F1: (v3 * 0)}
    v24 = v24
    fmt.Println(v1)
    switch 9223372036854775807 {
    case 904:
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
    case 566:
        var v26 uint32 = 345
        _ = v26
    case 9223372036854775807:
        fmt.Println(v3, v3, v1)
    }
    _ = len(v4)
    if (!true) {
        fmt.Println(v1)
        var v27 string = "foo bar"
        _ = v27
    }
    for fi28 := 0; fi28 < 8; fi28++ {
        var v29 int8 = 46
        v2 = append(v2, 4294967295)
        _ = fi28
        _ = v29
    }
    v4[(("a" + "hello") + ("hello" + "foo bar"))] = 754
    v24 = (^(v24))
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
