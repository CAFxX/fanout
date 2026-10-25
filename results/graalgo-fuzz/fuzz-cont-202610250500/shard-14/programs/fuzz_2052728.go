package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint64 }
func genFn0(p0 int, p1 uint64) string {
    fmt.Println(p1, p0)
    v5 := GenS1{F0: true}
    _ = p0
    _ = p1
    _ = v5
    return "\n\t"
}

func main() {
    var v1 uint = 2
    var v2 string = ("" + "")
    var v3 bool = (false || false)
    v4 := make(map[int]bool)
    switch v2 {
    case "foo bar":
        if (1000 != -978) {
            v1 = 65535
            switch "a" {
            case "\n\t":
                var v6 int32 = 0
                _ = v6
            case "a":
                v7 := GenS2{F0: (v2 + "hello"), F1: 42}
                _ = v7
            default:
                fmt.Println(v3, v2)
            }
        } else {
            v3 = (v3 && true)
            v8 := GenS2{F0: v2, F1: 42}
            _ = v8
        }
    }
    v9 := GenS1{F0: (!(!v3))}
    v10 := []int{128, -1, 9223372036854775807}
    v11 := GenS2{F0: (v2 + v2), F1: 18446744073709551615}
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 uint = v1
        _ = fi12
        _ = v13
    }
    v14 := GenS2{F0: v2, F1: 1}
    var v15 int = -605
    v1 += v1
    delete(v4, (v15 & int(v1)))
    var v16 uint64 = 2
    v17 := GenS2{F0: "\n\t", F1: 810}
    for fi18 := 0; fi18 < 10; fi18++ {
        v3 = (-2.25 > 1.5)
        var v19 uint32 = 0
        for fi20 := 0; fi20 < 4; fi20++ {
            v2 = "\n\t"
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        _ = len(v10)
        _ = fi21
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        v23 := GenS2{F0: ("hello" + "hello"), F1: 18446744073709551615}
        _ = fi22
        _ = v23
    }
    var v24 float64 = 3.14159
    switch "" {
    case "hello":
        if v3 {
            if (false && (!(-1 < 115))) {
                if (447 > 0) {
                    for fi25 := 0; fi25 < 3; fi25++ {
                        if false {
                            var v26 uint = v1
                            v1 = (^(v26))
                            _ = v26
                        }
                        v27 := GenS2{F0: v2, F1: v16}
                        _ = fi25
                        _ = v27
                    }
                    v28 := GenS1{F0: ((v15 << ((2) & 63)) != (v15 | v15))}
                    v24 = v24
                    _ = v28
                }
                var v29 int64 = 0
                for fi30 := 0; fi30 < 12; fi30++ {
                    v31 := GenS2{F0: ("\n\t" + v2), F1: 42}
                    _ = fi30
                    _ = v31
                }
                _ = v29
            }
        } else {
            fmt.Println(v2, v3, v15)
        }
    case "\n\t":
        var v32 string = ((v2 + v2) + v2)
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v9
}
