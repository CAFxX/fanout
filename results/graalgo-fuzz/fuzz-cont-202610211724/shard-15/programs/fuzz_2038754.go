package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint; F2 string }
func genFn0(p0 int16, p1 int8) uint16 {
    v7 := GenS1{F0: (p0 * p0), F1: 820}
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS1{F0: p0, F1: 18446744073709551615}
        p1 = (-(p1))
        _ = fi8
        _ = v9
    }
    var v10 uint = 122
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return 42
}

func main() {
    v1 := []string{"a", "foo bar", ("hello" + "hello")}
    v2 := make(map[int]bool)
    v3 := make(map[string]bool)
    var v4 int16 = -1
    var v5 int16 = (^(v4))
    v6 := make(map[int]uint32)
    var v11 int16 = (v4 * -1)
    v12 := GenS1{F0: -32768, F1: 38}
    var v13 bool = true
    var v14 uint64 = 139
    for fi15 := 0; fi15 < 14; fi15++ {
        v4 += -79
        _ = fi15
    }
    if (false && (v14 == 2)) {
        fmt.Println(v4, v4, v14)
    }
    var v16 int16 = -1
    var v17 uint64 = (^(v14))
    var v18 float32 = (3.14159 / float32(1.5))
    if (!v13) {
        var v19 uint16 = 519
        var v20 string = "foo bar"
        _ = v19
        _ = v20
    } else {
        if v13 {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        for fi21 := 0; fi21 < 19; fi21++ {
            v22 := GenS2{F0: 42, F1: 873, F2: "\n\t"}
            v18 += (v18 / v18)
            _ = fi21
            _ = v22
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if len(v1) > 0 {
        v1[0] = ("\n\t" + ("foo bar" + "foo bar"))
    }
    v23 := v2[-869]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
