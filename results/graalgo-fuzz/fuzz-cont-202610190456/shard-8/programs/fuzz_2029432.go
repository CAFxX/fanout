package main

import (
    "fmt"
)

func genFn0(p0 uint8, p1 int) bool {
    p1 -= -9223372036854775808
    v6 := []int{(-(p1)), p1}
    _ = p0
    _ = p1
    _ = v6
    return true
}
func genFn1(p0 int16) uint64 {
    var v7 int64 = -1
    for fi8 := 0; fi8 < 8; fi8++ {
        p0 = p0
        _ = fi8
    }
    _ = p0
    _ = v7
    return 2
}

func main() {
    v1 := []string{"", (("\n\t" + "\n\t") + "foo bar"), "a", ""}
    v2 := make(map[string]string)
    var v3 uint16 = 42
    var v4 uint64 = 42
    v5 := make(map[string]uint32)
    var v9 uint = 294
    v4 += (v4 | genFn1(42))
    var v10 int64 = 42
    var v11 uint32 = 2
    switch ("" + "") {
    case "":
        fmt.Println(v10, v11)
    case "foo bar":
        fmt.Println(v3)
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v12 int8 = 68
    fmt.Println(v3)
    if true {
        if true {
            var v13 float64 = 0.0
            var v14 uint8 = 2
            _ = v13
            _ = v14
        }
        switch ("\n\t" + "") {
        case "foo bar":
            if false {
                switch "hello" {
                case "":
                    v12 += -123
                }
            }
        default:
            var v15 int64 = v10
            _ = v15
        }
        var v16 uint16 = (v3 ^ uint16(v9))
        _ = v16
    }
    fmt.Println(v10, v11)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
