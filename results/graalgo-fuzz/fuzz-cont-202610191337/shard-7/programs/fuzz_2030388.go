package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{314, 42, 358, 4294967295}
    var v3 int16 = -250
    var v4 int16 = (v3 ^ v3)
    var v5 float64 = (0.0 * 3.14159)
    fmt.Println(v5, v3, v5)
    if false {
        v1[32768] = 0
    } else {
        var v6 bool = true
        switch "a" {
        case "a":
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
        case "":
            for fi7 := 0; fi7 < 17; fi7++ {
                if v6 {
                    fmt.Println(v6)
                    v6 = (true || ((159 < -550) || (124 >= 1)))
                } else {
                    switch fi7 {
                    case -2147483648:
                        if (v6 && v6) {
                            var v8 float64 = v5
                            _ = v8
                        }
                    case 32768:
                        var v9 int32 = 2147483647
                        _ = v9
                    case -1:
                        fmt.Println(v3, v3)
                    }
                    v4 = -32768
                    v6 = (42 > 1)
                }
                delete(v1, (fi7 & -9223372036854775808))
                v10 := GenS1{F0: 65535, F1: 2147483647}
                _ = fi7
                _ = v10
            }
        }
        _ = v6
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        if false {
            _ = len(v1)
            v1[(fi11 % (((fi11 * (fi11 | fi11))) | 1))] = 2147483647
        }
        v3 -= v4
        _ = fi11
    }
    if true {
        var v12 int64 = 1
        v13 := GenS2{F0: v12}
        var v14 int = 32767
        _ = v12
        _ = v13
        _ = v14
    } else {
        var v15 uint16 = 32767
        _ = v15
    }
    v16 := GenS1{F0: 32767, F1: 0}
    delete(v1, 9223372036854775807)
    v4 = 190
    if true {
        if (-1 < 1) {
            v5 -= v5
        } else {
            switch ((((("foo bar" + ("hello" + "hello")) + "foo bar") + "") + (("\n\t" + "foo bar") + "foo bar")) + "") {
            case "\n\t":
                var v17 uint16 = 32767
                _ = v17
            default:
                v1[42] = 1
            }
            v5 = v5
            v4 -= v4
        }
        v18 := GenS2{F0: 610}
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
