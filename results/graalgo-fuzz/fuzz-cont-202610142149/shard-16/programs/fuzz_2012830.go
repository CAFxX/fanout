package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int }
type GenS2 struct { F0 bool }

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = ((0.1 < 0.0) || (1 > 1))
    v3 := []uint32{341}
    var v4 uint8 = 128
    v5 := []uint32{2147483647}
    var v6 int16 = -245
    if true {
        v7 := GenS1{F0: 0, F1: (("hello" + "hello") + ((("\n\t" + "\n\t") + ("\n\t" + "foo bar")) + "a")), F2: -260}
        var v8 uint16 = 42
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v7
        _ = v8
    }
    fmt.Println(v4, v6, v4)
    var v9 int32 = 0
    v10 := GenS2{F0: (57 <= 588)}
    v6 = v6
    if (!true) {
        fmt.Println(v4, v6, v2)
    }
    v1[417] = "hello"
    v11 := v1[1]
    if (true || true) {
        if ((498 > 589) || (0 > -128)) {
            fmt.Println(v4, v11, v4)
            fmt.Println(v11, v4, v4)
            if (v2 || ((true || (313 == 875)) && v2)) {
                v2 = (9223372036854775807 != 893)
                v6 -= -488
                _ = len(v1)
            } else {
                fmt.Println(v2, v9)
                fmt.Println(v6, v2)
                fmt.Println(v11, v2, v6)
            }
        } else {
            var v12 int = 32768
            for fi13 := 0; fi13 < 9; fi13++ {
                fmt.Println(fi13, v6, v4)
                _ = fi13
            }
            fmt.Println(v2, v9)
            _ = v12
        }
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        v3 = append(v3, 4294967295)
        _ = fi14
    }
    v15 := GenS2{F0: (2 >= 436)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
