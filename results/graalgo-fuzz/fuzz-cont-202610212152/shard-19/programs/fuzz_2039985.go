package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint, p1 uint8, p2 uint32) int8 {
    p2 += (p2 | (p2 * (p2 << (((p0 << ((p0) & 63))) & 63))))
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    return -128
}

func main() {
    var v1 bool = false
    v2 := []int{}
    v3 := make(map[string]uint32)
    v4 := []string{"a", "a"}
    fmt.Println(v1)
    if (!(v1 && true)) {
        v1 = v1
        if false {
            switch "hello" {
            case "\n\t":
                fmt.Println(v1)
            default:
                var v5 int64 = 63
                _ = v5
            }
        } else {
            delete(v3, "a")
            fmt.Println(v1)
            if v1 {
                var v6 int8 = 127
                _ = v6
            }
        }
    }
    var v7 int64 = 411
    var v8 int16 = 416
    fmt.Println(v7, v8, v8)
    v8 -= v8
    var v9 int16 = 32767
    switch -279 {
    case 1:
        for fi10 := 0; fi10 < 20; fi10++ {
            v7 = 42
            fmt.Println(v9, v7, fi10)
            _ = fi10
        }
    }
    switch -901 {
    case 703:
        v1 = ((0 > 659) && v1)
    case 500:
        switch -9223372036854775808 {
        case 978:
            v9 += (v8 - v9)
        default:
            if v1 {
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
            } else {
                for fi11 := 0; fi11 < 5; fi11++ {
                    if v1 {
                        var v12 float64 = ((-2.25 + 1.5) - (1.5 / 1.0))
                        for fi13 := 0; fi13 < 13; fi13++ {
                            delete(v3, "hello")
                            _ = fi13
                        }
                        _ = v12
                    } else {
                        v14 := v3[("foo bar" + "")]
                        _ = v14
                    }
                    _ = fi11
                }
            }
        }
    case 772:
        v15 := GenS1{F0: 52}
        _ = v15
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
