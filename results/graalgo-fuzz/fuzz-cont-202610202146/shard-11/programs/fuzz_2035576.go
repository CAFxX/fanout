package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
func genFn0(p0 uint8, p1 int64, p2 uint) uint32 {
    for fi6 := 0; fi6 < 8; fi6++ {
        p1 -= p1
        _ = fi6
    }
    p1 += -783
    var v7 uint32 = 4294967295
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return v7
}

func main() {
    var v1 int16 = 1
    var v2 float32 = 0.0
    v3 := make(map[string]string)
    var v4 int64 = 584
    var v5 float32 = 3.14159
    for fi9 := 0; fi9 < 17; fi9++ {
        if false {
            v4 -= v4
        } else {
            if false {
                for fi10 := 0; fi10 < 13; fi10++ {
                    v11 := GenS1{F0: 4294967295, F1: 2}
                    var v12 bool = false
                    v13 := v3[("foo bar" + ("hello" + "a"))]
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
                fmt.Println(fi9, v5)
                fmt.Println(fi9, v5, v2)
            } else {
                v4 = -72
                var v14 int8 = -1
                v2 = v5
                _ = v14
            }
            v15 := GenS1{F0: 801, F1: genFn0(128, v4, 42)}
            _ = v15
        }
        delete(v3, "hello")
        var v16 uint64 = 18446744073709551615
        _ = fi9
        _ = v16
    }
    switch "a" {
    case "\n\t":
        var v17 uint16 = 65535
        _ = v17
    case "foo bar":
        v1 = v1
    case "":
        v4 = int64(v1)
    default:
        fmt.Println(v5)
    }
    var v18 int8 = 1
    v5 -= (v5 - 0.0)
    switch 32767 {
    case 2147483647:
        v3["hello"] = ("\n\t" + ("a" + ""))
    default:
        var v19 float64 = -2.25
        _ = v19
    }
    var v20 int32 = -2147483648
    var v21 int8 = -66
    var v22 uint = 18446744073709551615
    fmt.Println(v4, v5, v20)
    v18 -= int8(v22)
    v20 = (v20 - -2147483648)
    var v23 int64 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
