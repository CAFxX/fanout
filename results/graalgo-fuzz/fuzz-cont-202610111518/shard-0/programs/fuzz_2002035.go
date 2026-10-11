package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
func genFn0(p0 string) int {
    var v4 uint = 4294967295
    var v5 uint8 = 250
    _ = p0
    _ = v4
    _ = v5
    return 32768
}
func genFn1(p0 uint64) int32 {
    if (421 != 414) {
        fmt.Println(p0)
    }
    fmt.Println(p0)
    _ = p0
    return 1
}

func main() {
    var v1 float32 = -2.25
    var v2 int8 = -99
    var v3 string = ((("a" + ("foo bar" + "hello")) + ("" + ((("\n\t" + "a") + "hello") + ("hello" + "a")))) + "a")
    var v6 uint32 = 2147483647
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 int8 = v2
        var v9 uint8 = 0
        _ = fi7
        _ = v8
        _ = v9
    }
    v3 = v3
    v10 := []int{9223372036854775807, -9223372036854775808, -605, 42}
    fmt.Println(v1)
    v3 = ((((v3 + "a") + (v3 + "foo bar")) + v3) + v3)
    v11 := GenS1{F0: 42, F1: 42}
    switch -9223372036854775808 {
    case -204:
        if true {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
        } else {
            v12 := GenS1{F0: 760, F1: 4294967295}
            switch ((v3 + v3) + v3) {
            case "foo bar":
                v13 := GenS1{F0: 255, F1: 65535}
                _ = v13
            case "":
                v14 := GenS1{F0: 65535, F1: 760}
                _ = v14
            default:
                var v15 uint32 = v6
                _ = v15
            }
            _ = v12
        }
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        var v17 uint16 = 28
        if (!true) {
            fmt.Println(v1, v17)
        }
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        v19 := GenS1{F0: 319, F1: 348}
        v3 = v3
        _ = fi18
        _ = v19
    }
    fmt.Println(v2)
    v20 := GenS1{F0: 0, F1: 46}
    fmt.Println(v3, v1)
    v2 -= -103
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v6
}
