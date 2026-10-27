package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }

func main() {
    var v1 uint = 65535
    var v2 int = -401
    var v3 uint8 = 128
    var v4 int16 = 1
    v5 := make(map[int]int)
    v2 = (v2 + 764)
    if (42 > -1) {
        v1 += v1
    }
    var v6 string = "\n\t"
    if ("\n\t" <= v6) {
        fmt.Println(v4)
        switch ((((("a" + v6) + (v6 + (("" + v6) + v6))) + (v6 + (v6 + v6))) + v6) + ("a" + (v6 + "foo bar"))) {
        case "\n\t":
            for fi7 := 0; fi7 < 2; fi7++ {
                v6 = v6
                var v8 int32 = 1
                var v9 string = "foo bar"
                _ = fi7
                _ = v8
                _ = v9
            }
        default:
            v10 := v5[127]
            _ = v10
        }
        v6 = (((v6 + (v6 + "hello")) + "") + ("hello" + (("foo bar" + ("a" + v6)) + v6)))
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        v2 -= 32768
        _ = fi11
    }
    var v12 float32 = 1.5
    v2 = 127
    v13 := GenS1{F0: (v6 + v6), F1: 42}
    v14 := []int{v2, v2}
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 uint64 = 0
        _ = fi15
        _ = v16
    }
    var v17 uint16 = 42
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 int32 = -1
        v5[(-(v2))] = (fi18 | (fi18 * (v2 << (((v1 & v1)) & 63))))
        var v20 float32 = v12
        _ = fi18
        _ = v19
        _ = v20
    }
    v4 = (v4 & (v4 << ((18446744073709551615) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
