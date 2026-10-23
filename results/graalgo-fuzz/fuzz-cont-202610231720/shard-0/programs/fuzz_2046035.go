package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int16, p1 uint8, p2 uint32) uint64 {
    if true {
        var v7 int64 = -992
        _ = v7
    } else {
        v8 := []int{2147483647, 9223372036854775807, -1, 357}
        _ = v8
    }
    p2 = p2
    var v9 bool = (-2.25 >= -2.25)
    var v10 int = 776
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 2
}
func genFn1() bool {
    var v11 uint8 = 127
    var v12 int16 = -769
    _ = v11
    _ = v12
    return true
}

func main() {
    var v1 uint32 = 0
    v2 := []string{("" + "foo bar"), (("a" + "") + "a")}
    v3 := make(map[uint32]string)
    var v4 int16 = 42
    v5 := make(map[string]int)
    var v6 float32 = (0.0 - (1e10 / 0.1))
    if (false || false) {
        fmt.Println(v4)
    }
    v13 := GenS1{F0: ((false || false) || false)}
    var v14 int = 1
    if (42 <= 255) {
        if true {
            for fi15 := 0; fi15 < 9; fi15++ {
                v6 = 0.1
                v6 += ((float32(-2.25) / (((float32(1e10) - v6) + (0.1 * -2.25)) / -2.25)) / v6)
                _ = fi15
            }
            var v16 uint8 = 1
            _ = v16
        }
        var v17 int64 = 0
        _ = v17
    } else {
        for fi18 := 0; fi18 < 16; fi18++ {
            v6 = v6
            var v19 int16 = v4
            v20 := v3[(^(v1))]
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    if len(v2) > 0 {
        v2[0] = ("" + (("\n\t" + ("a" + "a")) + "hello"))
    }
    v21 := GenS1{F0: false}
    if ((51 != 1) && ((127 > 127) && false)) {
        var v22 int64 = -357
        _ = v22
    } else {
        v1 = (v1 & (v1 ^ (v1 - (^(v1)))))
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 string = "hello"
            v25 := GenS1{F0: ((!false) && (genFn0(v4, 0, 2) < 1))}
            switch (fi23 + fi23) {
            case -776:
                fmt.Println(v24, v24)
            case -9223372036854775808:
                v24 = v24
            }
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    v6 += (v6 * float32(3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
