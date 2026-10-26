package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint; F1 int32; F2 int }
type GenS3 struct { F0 uint64 }
func genFn0() uint32 {
    v4 := GenS1{F0: -1, F1: 42, F2: 42}
    var v5 int64 = 536
    v6 := GenS1{F0: 42, F1: (^(v5)), F2: 661}
    v5 -= (v5 + 230)
    _ = v4
    _ = v5
    _ = v6
    return 4294967295
}

func main() {
    v1 := []string{(("a" + "foo bar") + ("hello" + (("foo bar" + ("\n\t" + "a")) + ("\n\t" + ("\n\t" + "\n\t"))))), "\n\t", ""}
    var v2 int = -1
    var v3 uint8 = 255
    var v7 uint = 65535
    switch "a" {
    case "a":
        fmt.Println(v2, v2)
    default:
        var v8 int8 = 100
        _ = v8
    }
    v9 := GenS1{F0: 42, F1: 0, F2: 255}
    if (!(-2.25 >= 3.14159)) {
        var v10 int16 = -978
        _ = v10
    }
    v7 -= (v7 ^ (v7 % ((v7) | 1)))
    v3 += v3
    switch "hello" {
    case "foo bar":
        v7 = v7
    case "\n\t":
        v3 -= 255
    }
    var v11 uint64 = 18446744073709551615
    var v12 int64 = -974
    fmt.Println(v11, v7)
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := GenS2{F0: 4294967295, F1: 42, F2: fi13}
        v15 := GenS2{F0: 42, F1: -1, F2: fi13}
        _ = fi13
        _ = v14
        _ = v15
    }
    v11 = v11
    if ((true && (false && (!((1.5 * (1.5 - 0.0)) < (1.5 + (((-2.25 * (0.1 + (float64(0.1) + 1.5))) * float64(3.14159)) - -2.25)))))) || (3.14159 != ((-2.25 + 1.5) / 0.1))) {
        if len(v1) > 0 {
            v1[0] = ("hello" + "\n\t")
        }
    }
    if false {
        v16 := GenS3{F0: v11}
        _ = v16
    } else {
        v1 = append(v1, "a")
        if true {
            var v17 int = (v2 + (v2 ^ (^(v2))))
            _ = v17
        }
    }
    if true {
        v1 = append(v1, "a")
        v12 = 9223372036854775807
    }
    v18 := GenS3{F0: (^(v11))}
    var v19 bool = true
    switch 9223372036854775807 {
    case 9223372036854775807:
        for fi20 := 0; fi20 < 17; fi20++ {
            fmt.Println(fi20, v3)
            var v21 int32 = 2147483647
            _ = fi20
            _ = v21
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v7
    _ = v9
}
