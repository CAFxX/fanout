package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
func genFn0(p0 int32, p1 uint32, p2 int8) uint {
    var v4 float64 = 3.14159
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 4294967295
}

func main() {
    var v1 uint = 0
    var v2 uint = v1
    v3 := []string{"a", "a", ""}
    if (!false) {
        _ = len(v3)
        fmt.Println(v2, v1)
        var v5 uint16 = 65535
        _ = v5
    } else {
        for fi6 := 0; fi6 < 8; fi6++ {
            fmt.Println(v1)
            fmt.Println(fi6)
            var v7 uint = v2
            _ = fi6
            _ = v7
        }
        if (!false) {
            if false {
                switch ((("" + ("hello" + "a")) + "") + "") {
                case "\n\t":
                    _ = len(v3)
                }
                for fi8 := 0; fi8 < 8; fi8++ {
                    v1 = (v1 | genFn0(884, 381, 22))
                    v9 := GenS1{F0: -9223372036854775808, F1: false}
                    _ = fi8
                    _ = v9
                }
            } else {
                var v10 uint = (v2 & v2)
                for fi11 := 0; fi11 < 15; fi11++ {
                    var v12 int = fi11
                    _ = len(v3)
                    _ = fi11
                    _ = v12
                }
                _ = v10
            }
            v3 = append(v3, ("a" + "\n\t"))
        }
    }
    if false {
        v2 = (v1 ^ v1)
        fmt.Println(v2, v1)
        for fi13 := 0; fi13 < 2; fi13++ {
            v2 = v2
            fmt.Println(v1, v1, fi13)
            _ = fi13
        }
    } else {
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 uint32 = 129
            _ = fi14
            _ = v15
        }
        var v16 string = ("foo bar" + "a")
        _ = v16
    }
    var v17 float64 = 0.0
    v3 = append(v3, ("foo bar" + "foo bar"))
    fmt.Println(v1, v1, v1)
    var v18 int64 = 1
    if (!(71 == 255)) {
        var v19 int = -72
        v19 -= v19
        if false {
            if false {
                var v20 uint = 2
                _ = v20
            }
            v19 -= (v19 & (v19 + (v19 - v19)))
        }
        _ = v19
    }
    v1 -= 65535
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
