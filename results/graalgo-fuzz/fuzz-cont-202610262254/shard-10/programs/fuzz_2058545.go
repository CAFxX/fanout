package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 string; F1 string; F2 int16 }
func genFn0(p0 int8, p1 bool) uint {
    if (0 >= -113) {
        switch "\n\t" {
        case "foo bar":
            v4 := []int{127, -9223372036854775808, 1}
            _ = v4
        case "hello":
            v5 := GenS1{F0: 143, F1: p1}
            _ = v5
        case "":
            p0 -= p0
        default:
            p0 -= p0
        }
    }
    var v6 int = 2147483647
    p0 = 66
    _ = p0
    _ = p1
    _ = v6
    return 0
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[string]int)
    var v3 uint32 = 799
    var v7 uint32 = v3
    if (false && (false && (-9223372036854775808 <= -9223372036854775808))) {
        v8 := GenS2{F0: (("foo bar" + "foo bar") + "foo bar"), F1: (("a" + "\n\t") + (("\n\t" + ("hello" + "foo bar")) + ("foo bar" + ("" + "foo bar")))), F2: -791}
        _ = v8
    } else {
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 int64 = -1
            _ = fi9
            _ = v10
        }
    }
    var v11 uint64 = v1
    var v12 bool = false
    var v13 uint32 = 0
    v13 = (^(v13))
    var v14 int64 = 0
    v15 := GenS2{F0: "", F1: ("hello" + "\n\t"), F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v7
}
