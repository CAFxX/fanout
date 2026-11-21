package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
func genFn0(p0 uint, p1 string) uint8 {
    var v7 float64 = float64(3.14159)
    v8 := []int{2147483647}
    fmt.Println(p0)
    v9 := GenS1{F0: (p0 / ((p0) | 1)), F1: -1}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 133
}

func main() {
    var v1 bool = (42 <= 9223372036854775807)
    var v2 uint64 = 42
    v3 := []string{"hello", (((("foo bar" + ("hello" + "")) + ((("\n\t" + "\n\t") + "hello") + ("\n\t" + ("" + "\n\t")))) + "foo bar") + ("\n\t" + "")), "foo bar"}
    v4 := make(map[int]int)
    v5 := make(map[uint32]uint32)
    var v6 int32 = 1
    for fi10 := 0; fi10 < 2; fi10++ {
        fmt.Println(fi10, v6)
        _ = fi10
    }
    switch "\n\t" {
    case "foo bar":
        v1 = v1
    default:
        fmt.Println(v1, v1)
    }
    var v11 string = "hello"
    switch 372 {
    case 417:
        v6 = v6
    case 128:
        for fi12 := 0; fi12 < 2; fi12++ {
            if (((0 == 327) && true) || v1) {
                if len(v3) > 0 {
                    v3[0] = v11
                }
                if v1 {
                    v3 = append(v3, v11)
                    v2 = (v2 | (^(v2)))
                }
            } else {
                v2 -= 906
                delete(v4, -416)
            }
            fmt.Println(fi12, v11)
            _ = fi12
        }
    }
    var v13 uint32 = 805
    _ = len(v4)
    if true {
        v14 := GenS1{F0: 2, F1: 0}
        _ = v14
    } else {
        _ = len(v5)
        var v15 float32 = (1e10 * (0.1 * 1.5))
        v16 := GenS1{F0: 435, F1: -1}
        _ = v15
        _ = v16
    }
    if len(v3) > 0 {
        v3[0] = (v11 + v11)
    }
    switch "a" {
    case "\n\t":
        v13 = v13
    case "foo bar":
        v3 = append(v3, "\n\t")
    default:
        switch 9223372036854775807 {
        case -2147483648:
            v2 += v2
        case 2147483647:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 0:
            var v17 float32 = -2.25
            _ = v17
        default:
            var v18 int32 = 436
            _ = v18
        }
    }
    v1 = ((192 > 1) || false)
    v2 += (v2 % (((v2 | 418)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
