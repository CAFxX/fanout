package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
func genFn0(p0 uint8, p1 uint16) uint {
    var v5 uint16 = (^(p1))
    var v6 int16 = -240
    if ((p1 << ((42) & 63)) == (^(v5))) {
        switch "foo bar" {
        case "":
            switch (("foo bar" + "") + "foo bar") {
            case "a":
                var v7 uint64 = 1
                _ = v7
            case "\n\t":
                switch "a" {
                case "":
                    fmt.Println(v5)
                case "a":
                    var v8 uint8 = (p0 * p0)
                    _ = v8
                default:
                    for fi9 := 0; fi9 < 6; fi9++ {
                        v10 := []int{(fi9 / ((fi9) | 1)), (fi9 & fi9), (-(fi9))}
                        _ = fi9
                        _ = v10
                    }
                }
            case "foo bar":
                v11 := []int{}
                _ = v11
            }
        default:
            var v12 string = "hello"
            _ = v12
        }
        p1 = (v5 ^ 166)
    } else {
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 int = 127
            p1 += (p1 % ((p1) | 1))
            _ = fi13
            _ = v14
        }
        var v15 uint16 = (^(v5))
        _ = v15
    }
    v6 -= v6
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 float64 = (1e10 + 1e10)
    v2 := []string{}
    v3 := []uint32{0, 65535}
    var v4 float64 = v1
    var v16 string = "hello"
    var v17 bool = true
    fmt.Println(v1, v17)
    var v18 int = 32767
    switch v16 {
    case "\n\t":
        var v19 int64 = 42
        _ = v19
    }
    v18 += v18
    v20 := GenS1{F0: 1, F1: 1}
    fmt.Println(v18)
    if (9223372036854775807 == 214) {
        for fi21 := 0; fi21 < 19; fi21++ {
            var v22 int8 = 0
            _ = fi21
            _ = v22
        }
        v23 := GenS1{F0: 704, F1: 42}
        var v24 uint16 = 363
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        fmt.Println(v17, v17, v1)
        v16 = v16
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
