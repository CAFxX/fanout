package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
func genFn0(p0 int64) int8 {
    v6 := GenS1{F0: -1}
    switch (("hello" + "\n\t") + "a") {
    case "hello":
        if (!(255 <= 1)) {
            var v7 int16 = 1000
            _ = v7
        } else {
            var v8 uint = 4294967295
            fmt.Println(v8)
            _ = v8
        }
    case "":
        p0 = p0
    case "foo bar":
        p0 = p0
    default:
        var v9 string = "foo bar"
        _ = v9
    }
    var v10 int = 2147483647
    v11 := GenS2{F0: -1}
    _ = p0
    _ = v10
    _ = v11
    _ = v6
    return 1
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{76}
    var v3 int16 = 32767
    var v4 string = (((("hello" + "") + "a") + "a") + "")
    var v5 int32 = -1
    fmt.Println(v4)
    var v12 bool = true
    v4 = v4
    var v13 int32 = (v5 ^ v5)
    for fi14 := 0; fi14 < 8; fi14++ {
        switch ((v4 + (v4 + v4)) + v4) {
        case "a":
            v15 := GenS2{F0: -1}
            _ = v15
        }
        _ = fi14
    }
    switch "foo bar" {
    case "":
        for fi16 := 0; fi16 < 3; fi16++ {
            if v12 {
                var v17 bool = (!v12)
                if len(v2) > 0 {
                    v2[0] = 65535
                }
                var v18 string = v4
                _ = v17
                _ = v18
            } else {
                var v19 int16 = (v3 % ((v3) | 1))
                _ = v19
            }
            _ = fi16
        }
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        fmt.Println(v4, v4)
        var v21 int16 = (v3 + v3)
        v5 = 708
        _ = fi20
        _ = v21
    }
    var v22 float64 = (1e10 + (0.1 / 1.0))
    v22 -= (v22 * 1e10)
    v3 = (v3 + (v3 & (v3 >> ((255) & 63))))
    fmt.Println(v4)
    if false {
        fmt.Println(v3, v5, v12)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
