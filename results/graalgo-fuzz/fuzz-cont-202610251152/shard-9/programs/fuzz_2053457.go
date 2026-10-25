package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint }
type GenS2 struct { F0 string }
func genFn0(p0 uint, p1 uint, p2 int64) int32 {
    v5 := []int{-1, 0, 0, -1}
    switch "\n\t" {
    case "":
        p0 = 42
    case "\n\t":
        fmt.Println(p1, p2)
    case "foo bar":
        v6 := GenS1{F0: 652, F1: 821, F2: 568}
        _ = v6
    default:
        v5 = append(v5, 2147483647)
    }
    p0 -= (p1 * 65535)
    if false {
        var v7 float32 = 1e10
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -2147483648
}
func genFn1(p0 uint) string {
    var v8 int16 = 0
    var v9 int32 = 691
    _ = p0
    _ = v8
    _ = v9
    return ("hello" + "a")
}

func main() {
    v1 := []uint32{4294967295, 2147483648}
    v2 := []string{(("hello" + "foo bar") + "a")}
    var v3 string = ""
    var v4 uint64 = 1
    for fi10 := 0; fi10 < 14; fi10++ {
        v1 = append(v1, 2147483647)
        _ = fi10
    }
    var v11 int32 = 1
    if true {
        v3 = v3
    } else {
        v11 -= v11
    }
    if ((365 <= 687) || (v3 < "")) {
        v11 -= (v11 % ((v11) | 1))
        var v12 float64 = 3.14159
        v4 = v4
        _ = v12
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        fmt.Println(v4)
        var v14 int64 = 42
        _ = fi13
        _ = v14
    }
    switch 1 {
    case 127:
        v1 = append(v1, 2147483648)
    case 42:
        v11 = genFn0(2, 18446744073709551615, 581)
    case 2:
        if (false || (true || (true && false))) {
            v15 := GenS1{F0: -2147483648, F1: 131, F2: 197}
            fmt.Println(v4, v4)
            _ = v15
        } else {
            var v16 int32 = v11
            v17 := GenS1{F0: -221, F1: 1, F2: 255}
            switch 9223372036854775807 {
            case 127:
                if true {
                    v18 := GenS1{F0: genFn0(255, 1, -963), F1: 4294967295, F2: 18446744073709551615}
                    v19 := GenS2{F0: v3}
                    _ = v18
                    _ = v19
                }
            case 42:
                if (!(-25 != 0)) {
                    var v20 int32 = (v11 >> ((688) & 63))
                    _ = v20
                }
            case 0:
                v16 += genFn0(255, 18446744073709551615, -922)
            }
            _ = v16
            _ = v17
        }
    default:
        v3 = v3
    }
    v11 -= (^(v11))
    fmt.Println(v11)
    fmt.Println(v4)
    var v21 bool = false
    _ = len(v1)
    v4 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
