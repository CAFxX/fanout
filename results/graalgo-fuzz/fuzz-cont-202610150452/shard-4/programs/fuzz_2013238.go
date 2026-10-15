package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 uint16; F2 uint }
func genFn0(p0 uint32, p1 int16, p2 uint64) bool {
    v4 := []int{int(p0), 127}
    p0 = (p0 * (p0 ^ 2))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return true
}

func main() {
    v1 := []string{((("hello" + "") + "") + "hello"), "a", "hello"}
    var v2 bool = (42 >= 36)
    v3 := make(map[uint32]string)
    v5 := GenS1{F0: 634}
    fmt.Println(v2)
    if v2 {
        v2 = (!((!false) && (768 == 32767)))
    }
    var v6 int32 = -2147483648
    if v2 {
        v2 = v2
    } else {
        v6 -= v6
        if (v2 && (0 != -1)) {
            for fi7 := 0; fi7 < 18; fi7++ {
                v8 := GenS1{F0: 42}
                _ = fi7
                _ = v8
            }
            var v9 float64 = 1.5
            _ = v9
        }
        v10 := GenS1{F0: -1}
        _ = v10
    }
    switch "\n\t" {
    case "a":
        var v11 uint64 = 1
        _ = v11
    case "\n\t":
        v12 := GenS1{F0: 42}
        _ = v12
    case "":
        v13 := v3[1]
        _ = v13
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        v2 = (1 != 42)
        v6 += v6
        _ = fi14
    }
    v15 := GenS2{F0: -179, F1: 42, F2: 18446744073709551615}
    if v2 {
        var v16 string = "hello"
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
