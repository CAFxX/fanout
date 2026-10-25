package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint8; F2 bool }
type GenS3 struct { F0 string; F1 int8; F2 uint }
type GenS4 struct { F0 int64; F1 bool; F2 bool }
type GenS5 struct { F0 uint8; F1 bool; F2 uint64 }
func genFn0(p0 uint, p1 int64) uint16 {
    for fi4 := 0; fi4 < 17; fi4++ {
        switch "" {
        case "":
            var v5 int = (-(fi4))
            _ = v5
        case "a":
            var v6 int8 = 79
            _ = v6
        case "\n\t":
            var v7 uint8 = 59
            _ = v7
        }
        p0 += (p0 / ((p0) | 1))
        _ = fi4
    }
    switch "a" {
    case "":
        v8 := GenS1{F0: 42}
        _ = v8
    case "foo bar":
        v9 := GenS2{F0: 4294967295, F1: 255, F2: false}
        _ = v9
    case "hello":
        p1 -= 1
    default:
        switch 32767 {
        case 124:
            var v10 uint64 = 4294967295
            _ = v10
        case 32768:
            p1 += 470
        default:
            p0 = p0
        }
    }
    v11 := GenS3{F0: ((("" + ("hello" + ("a" + ("" + "foo bar")))) + (("a" + (("" + ((("hello" + "\n\t") + "") + ("foo bar" + "hello"))) + "hello")) + "foo bar")) + "hello"), F1: 1, F2: 751}
    _ = p0
    _ = p1
    _ = v11
    return 0
}

func main() {
    v1 := []string{(("hello" + "hello") + ""), "hello", "a", "foo bar"}
    var v2 uint8 = 42
    v3 := make(map[uint32]uint32)
    var v12 int = 4294967295
    switch ("foo bar" + "") {
    case "a":
        var v13 int8 = 1
        _ = v13
    case "foo bar":
        fmt.Println(v2)
    case "\n\t":
        v14 := GenS2{F0: 272, F1: (v2 * v2), F2: ("" > "a")}
        _ = v14
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        fmt.Println(fi15, fi15, v12)
        _ = fi15
    }
    v16 := GenS4{F0: -9223372036854775808, F1: false, F2: false}
    v2 += v2
    switch ("\n\t" + ("a" + ("\n\t" + "a"))) {
    case "\n\t":
        var v17 int32 = 2147483647
        _ = v17
    }
    v3[661] = 2147483647
    var v18 int64 = 0
    delete(v3, 336)
    v19 := GenS5{F0: 255, F1: (0 == 2147483647), F2: 2}
    v20 := GenS5{F0: v2, F1: ("\n\t" < ""), F2: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
