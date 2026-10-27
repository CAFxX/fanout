package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int32, p1 uint) uint64 {
    v7 := GenS1{F0: 32767, F1: 822}
    v8 := GenS1{F0: 741, F1: 284}
    v9 := GenS2{F0: -1, F1: 18446744073709551615}
    v10 := GenS2{F0: -1, F1: 42}
    var v11 string = "hello"
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return 43
}

func main() {
    var v1 int16 = -1
    v2 := make(map[string]string)
    var v3 int64 = 610
    v4 := []string{((("" + "\n\t") + "a") + (("" + (((("foo bar" + (("a" + ((("\n\t" + "") + "hello") + "foo bar")) + (("foo bar" + ((("\n\t" + "") + "\n\t") + "hello")) + ("hello" + (("a" + "foo bar") + "foo bar"))))) + ("" + "")) + ("\n\t" + "")) + "hello")) + "a")), "\n\t", "\n\t", (("foo bar" + "a") + "")}
    v5 := make(map[uint32]uint32)
    var v6 uint32 = 873
    if true {
        fmt.Println(v6)
    }
    v1 = -252
    fmt.Println(v1, v6, v3)
    var v12 int64 = (v3 ^ v3)
    v3 -= (v3 | (v12 - (v3 >> ((405) & 63))))
    v13 := GenS2{F0: v3, F1: 1}
    for fi14 := 0; fi14 < 19; fi14++ {
        v1 += v1
        v3 += int64(fi14)
        _ = fi14
    }
    v1 = (v1 ^ (v1 >> ((2) & 63)))
    v3 = (v12 / ((v12) | 1))
    v15 := GenS2{F0: 715, F1: 18446744073709551615}
    for fi16 := 0; fi16 < 5; fi16++ {
        v17 := GenS2{F0: v12, F1: 1}
        _ = fi16
        _ = v17
    }
    switch "foo bar" {
    case "a":
        var v18 int64 = (v12 * -1)
        _ = v18
    case "foo bar":
        _ = len(v5)
    default:
        for fi19 := 0; fi19 < 7; fi19++ {
            if false {
                fmt.Println(fi19, v12, v3)
                v20 := GenS3{F0: 643}
                var v21 float64 = (3.14159 - ((0.1 * (1e10 + 1e10)) * ((0.0 - -2.25) / (0.0 + float64(0.1)))))
                _ = v20
                _ = v21
            } else {
                v1 = v1
                v6 = (v6 ^ 4294967295)
            }
            fmt.Println(v6)
            v3 = v3
            _ = fi19
        }
    }
    delete(v2, "a")
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
