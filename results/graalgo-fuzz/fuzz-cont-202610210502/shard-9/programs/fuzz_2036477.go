package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 uint8; F2 int64 }
func genFn0(p0 int) int32 {
    for fi5 := 0; fi5 < 8; fi5++ {
        p0 = (p0 ^ p0)
        _ = fi5
    }
    p0 = (p0 * (p0 << ((2) & 63)))
    if (-1 > 42) {
        switch (("a" + "a") + ("hello" + "foo bar")) {
        case "hello":
            p0 = p0
        }
    } else {
        if (42 < 100) {
            var v6 string = "\n\t"
            for fi7 := 0; fi7 < 5; fi7++ {
                p0 -= 127
                if ((0 <= 114) && false) {
                    var v8 uint16 = 577
                    var v9 string = (v6 + (v6 + v6))
                    if false {
                        var v10 uint32 = 4294967295
                        var v11 float64 = 3.14159
                        _ = v10
                        _ = v11
                    }
                    _ = v8
                    _ = v9
                }
                var v12 uint = 4294967295
                _ = fi7
                _ = v12
            }
            _ = v6
        }
        var v13 uint32 = 2147483648
        v14 := []int{(p0 - p0), 32768, p0}
        _ = v13
        _ = v14
    }
    v15 := []int{p0, (p0 | p0), (p0 & -9223372036854775808), (p0 & p0)}
    p0 = p0
    _ = p0
    _ = v15
    return -359
}

func main() {
    v1 := []string{"hello", "hello", "", "\n\t"}
    var v2 uint8 = 2
    var v3 uint8 = 0
    v4 := make(map[string]int)
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 int = (fi16 / ((fi16) | 1))
        _ = fi16
        _ = v17
    }
    var v18 uint8 = 1
    switch ("" + "hello") {
    case "a":
        for fi19 := 0; fi19 < 8; fi19++ {
            fmt.Println(v3)
            _ = fi19
        }
    case "hello":
        v3 = (v2 + (v3 << ((18446744073709551615) & 63)))
    default:
        v18 = 42
    }
    fmt.Println(v2)
    v20 := GenS1{F0: 42}
    v21 := GenS1{F0: -513}
    for fi22 := 0; fi22 < 17; fi22++ {
        v23 := GenS2{F0: 42, F1: v2, F2: 9223372036854775807}
        v24 := v4[""]
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        var v26 int16 = 1
        _ = fi25
        _ = v26
    }
    _ = len(v4)
    var v27 float32 = 1.5
    v28 := GenS1{F0: 2147483647}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
