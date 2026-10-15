package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint }
type GenS2 struct { F0 string; F1 int32; F2 uint16 }

func main() {
    var v1 int = 706
    var v2 int64 = 42
    var v3 int32 = 42
    var v4 uint = 1
    var v5 uint = v4
    v3 -= -2147483648
    var v6 uint64 = 1
    switch int(v4) {
    case 0:
        fmt.Println(v4)
    case 127:
        v7 := GenS1{F0: "", F1: (("hello" + "hello") + "a"), F2: (v5 | (v5 << (((v5 * (v5 >> (((v5 + 158)) & 63)))) & 63)))}
        _ = v7
    case 2147483647:
        var v8 uint64 = v6
        _ = v8
    }
    v9 := GenS1{F0: "foo bar", F1: "hello", F2: v5}
    v6 += (v6 ^ 668)
    v5 += 18446744073709551615
    if false {
        v6 -= (v6 << ((316) & 63))
    } else {
        switch "foo bar" {
        case "\n\t":
            v2 -= int64(v3)
        default:
            var v10 int = (v1 | (v1 | v1))
            _ = v10
        }
        if false {
            if (-2.25 >= (-2.25 / 1.5)) {
                for fi11 := 0; fi11 < 17; fi11++ {
                    v5 += 760
                    _ = fi11
                }
                fmt.Println(v1, v2, v3)
            }
            var v12 uint16 = 2
            _ = v12
        }
        v1 = v1
    }
    fmt.Println(v6)
    v3 += v3
    v13 := []int{}
    var v14 int16 = 0
    for fi15 := 0; fi15 < 8; fi15++ {
        for fi16 := 0; fi16 < 14; fi16++ {
            v17 := GenS2{F0: "hello", F1: v3, F2: 2}
            _ = fi16
            _ = v17
        }
        var v18 uint64 = v6
        switch ("" + "") {
        case "hello":
            v6 = 42
        case "foo bar":
            v2 = v2
        case "\n\t":
            v13 = append(v13, (v1 >> (((^(v5))) & 63)))
        default:
            if len(v13) > 0 {
                v13[0] = fi15
            }
        }
        _ = fi15
        _ = v18
    }
    v19 := GenS2{F0: (((("hello" + "hello") + (("\n\t" + ("foo bar" + "")) + "hello")) + (("foo bar" + "foo bar") + "hello")) + "\n\t"), F1: v3, F2: 1}
    if false {
        var v20 int64 = (v2 + int64(v6))
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
