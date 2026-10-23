package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint64 }

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 476
    var v3 int32 = 247
    var v4 uint16 = 530
    v5 := []uint32{916, 411, 457}
    _ = len(v1)
    var v6 uint = 1
    var v7 uint = (v6 & 1)
    if false {
        switch (("foo bar" + "\n\t") + ("a" + "hello")) {
        case "a":
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 bool = (!((2 <= 18446744073709551615) && false))
                _ = fi8
                _ = v9
            }
        case "foo bar":
            var v10 float64 = float64(0.1)
            _ = v10
        default:
            for fi11 := 0; fi11 < 13; fi11++ {
                _ = len(v5)
                _ = fi11
            }
        }
        v12 := GenS1{F0: 2, F1: -1, F2: 961}
        if false {
            switch 0 {
            case -997:
                var v13 string = (((("" + "hello") + "hello") + ("a" + "hello")) + "foo bar")
                _ = v13
            default:
                _ = len(v5)
            }
        } else {
            v4 = (v4 + 2)
        }
        _ = v12
    }
    v2 = 572
    if (!(!(true || (v3 >= (v3 << ((v7) & 63)))))) {
        v3 -= 42
    }
    var v14 float64 = 0.1
    var v15 int64 = 42
    var v16 uint64 = 291
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
