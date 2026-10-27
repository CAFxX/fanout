package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int16) int32 {
    p0 += -410
    v4 := []int{-1, -846, 42}
    var v5 uint = 158
    _ = len(v4)
    var v6 uint32 = 2147483647
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return -833
}

func main() {
    v1 := []string{"a", "hello", "", ""}
    var v2 uint8 = 127
    var v3 uint = 255
    v1 = append(v1, "foo bar")
    var v7 bool = (1.5 <= 1.5)
    var v8 int = -515
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch "" {
    case "":
        var v9 int8 = -128
        _ = v9
    default:
        v10 := GenS1{F0: "foo bar"}
        _ = v10
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        fmt.Println(v2, v3)
        _ = fi11
    }
    v7 = v7
    if (1 >= 320) {
        var v12 int16 = 0
        v1 = append(v1, "\n\t")
        for fi13 := 0; fi13 < 9; fi13++ {
            v12 = 1000
            _ = fi13
        }
        _ = v12
    } else {
        for fi14 := 0; fi14 < 12; fi14++ {
            if (false && false) {
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
                if len(v1) > 0 {
                    v1[0] = ("" + "")
                }
            }
            _ = fi14
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
