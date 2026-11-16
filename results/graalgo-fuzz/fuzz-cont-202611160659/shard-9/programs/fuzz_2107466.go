package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint; F1 int8; F2 string }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint) uint32 {
    p0 -= p0
    v7 := []int{}
    var v8 uint8 = 1
    for fi9 := 0; fi9 < 15; fi9++ {
        p0 -= (p0 / ((p0) | 1))
        if true {
            p0 = 2
            var v10 int = 42
            _ = v10
        }
        var v11 int8 = 41
        _ = fi9
        _ = v11
    }
    _ = p0
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := []uint32{42}
    v2 := []int{}
    var v3 int64 = -9223372036854775808
    v4 := make(map[string]string)
    v5 := make(map[string]string)
    var v6 int = 127
    if (false && ((65535 != 129) && ((42 == 2147483647) || (v3 > (v3 - v3))))) {
        v3 -= 42
    } else {
        _ = len(v4)
    }
    var v12 uint = 0
    switch "hello" {
    case "a":
        if (247 >= 255) {
            fmt.Println(v12)
            v13 := GenS1{F0: v3, F1: -1}
            v14 := GenS2{F0: v12, F1: -1, F2: ("foo bar" + ("a" + (("" + ("a" + "hello")) + (("" + (("" + (("hello" + "\n\t") + "")) + ("foo bar" + (("a" + "hello") + "hello")))) + ("\n\t" + "")))))}
            _ = v13
            _ = v14
        }
    case "hello":
        v15 := v4["foo bar"]
        _ = v15
    case "\n\t":
        if false {
            var v16 uint16 = 295
            if ((!true) && ((-(v3)) == (^(v3)))) {
                if (0 >= 1) {
                    v12 -= (^(v12))
                } else {
                    var v17 string = ""
                    for fi18 := 0; fi18 < 19; fi18++ {
                        v16 = v16
                        _ = fi18
                    }
                    _ = v17
                }
            } else {
                fmt.Println(v3, v12, v3)
                fmt.Println(v16, v6, v16)
            }
            _ = v16
        }
    default:
        fmt.Println(v6, v12, v12)
    }
    _ = len(v5)
    if len(v2) > 0 {
        v2[0] = v6
    }
    var v19 uint32 = 42
    v20 := GenS3{F0: 1}
    var v21 uint8 = 67
    var v22 uint8 = 227
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
