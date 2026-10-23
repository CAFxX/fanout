package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int32 }

func main() {
    var v1 float32 = (1.5 * (1.5 + (0.0 - 3.14159)))
    v2 := make(map[string]string)
    var v3 uint64 = 18446744073709551615
    var v4 float32 = (v1 + (1e10 + v1))
    var v5 int64 = -9223372036854775808
    v6 := make(map[int]bool)
    var v7 uint = 745
    var v8 uint16 = 1
    var v9 int64 = 853
    switch -832 {
    case 634:
        v10 := []int{1, -494, 128}
        _ = v10
    case 1:
        if false {
            v2["foo bar"] = "a"
            var v11 int8 = -88
            _ = v11
        }
    default:
        if (2147483647 > int(v8)) {
            v12 := []int{32767, -1, 1}
            _ = v12
        }
    }
    switch 42 {
    case 32767:
        v7 -= v7
    case -582:
        var v13 int64 = v5
        _ = v13
    case 9223372036854775807:
        v14 := GenS1{F0: v3, F1: -66}
        _ = v14
    default:
        v2[(("foo bar" + ("hello" + ((("" + ("" + ("a" + ("" + "foo bar")))) + ("" + (("a" + "a") + "a"))) + "hello"))) + "hello")] = ("foo bar" + "foo bar")
    }
    if ((true && true) && false) {
        v15 := []int{779}
        _ = v15
    }
    if true {
        if (-9223372036854775808 > 42) {
            var v16 uint = (v7 << ((v7) & 63))
            v2["hello"] = "a"
            _ = v16
        }
    }
    v17 := v6[-1]
    v5 = (v5 >> (((v7 << ((42) & 63))) & 63))
    for fi18 := 0; fi18 < 7; fi18++ {
        if v17 {
            for fi19 := 0; fi19 < 10; fi19++ {
                v20 := GenS1{F0: v3, F1: 56}
                switch (fi18 - (fi19 << ((v7) & 63))) {
                case 0:
                    switch fi18 {
                    case 128:
                        fmt.Println(v3, v3)
                        var v21 string = "\n\t"
                        var v22 uint64 = uint64(v7)
                        _ = v21
                        _ = v22
                    default:
                        v23 := GenS1{F0: (^(v3)), F1: -7}
                        _ = v23
                    }
                case -9223372036854775808:
                    v9 = v5
                default:
                    if v17 {
                        var v24 string = ("foo bar" + (("" + ("" + ("hello" + ((("foo bar" + "hello") + "a") + "a")))) + "\n\t"))
                        v25 := GenS2{F0: uint8(v7), F1: 1, F2: 2147483647}
                        _ = v24
                        _ = v25
                    }
                }
                _ = fi19
                _ = v20
            }
            fmt.Println(v9, v4)
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
