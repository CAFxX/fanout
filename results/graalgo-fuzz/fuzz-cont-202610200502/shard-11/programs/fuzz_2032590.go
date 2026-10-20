package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int8; F2 uint }
type GenS3 struct { F0 uint16; F1 uint }

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = 42
    var v3 bool = (true || (120 > 1))
    var v4 uint8 = 1
    v5 := []string{("\n\t" + "")}
    v6 := GenS1{F0: 817}
    v3 = (-66 > 42)
    if len(v5) > 0 {
        v5[0] = ("\n\t" + ("foo bar" + "\n\t"))
    }
    for fi7 := 0; fi7 < 9; fi7++ {
        _ = len(v5)
        var v8 uint = uint(v4)
        v8 = (v8 % ((v8) | 1))
        _ = fi7
        _ = v8
    }
    v5 = append(v5, "a")
    _ = len(v5)
    switch (("a" + ("" + (("" + "a") + ""))) + ("\n\t" + "\n\t")) {
    case "":
        var v9 uint32 = 65535
        _ = v9
    case "a":
        var v10 uint32 = 2147483647
        _ = v10
    case "hello":
        if true {
            v11 := GenS1{F0: -2147483648}
            for fi12 := 0; fi12 < 9; fi12++ {
                v4 += v4
                delete(v1, "")
                var v13 uint64 = 0
                _ = fi12
                _ = v13
            }
            fmt.Println(v4, v4, v3)
            _ = v11
        }
    }
    delete(v1, "a")
    if true {
        var v14 int = 4294967295
        switch 32768 {
        case 128:
            switch (("a" + "\n\t") + ("a" + "a")) {
            case "\n\t":
                if (!((v4 + 2) > v4)) {
                    fmt.Println(v3, v4)
                    var v15 uint32 = 4294967295
                    v4 -= 149
                    _ = v15
                } else {
                    if v3 {
                        fmt.Println(v14, v2, v3)
                    } else {
                        v3 = (0 != 118)
                        var v16 bool = v3
                        _ = v16
                    }
                    var v17 uint32 = 0
                    _ = v17
                }
            default:
                var v18 bool = v3
                _ = v18
            }
        default:
            v1[""] = 2
        }
        var v19 int64 = v2
        _ = v14
        _ = v19
    }
    v20 := GenS2{F0: 1, F1: 100, F2: 4294967295}
    for fi21 := 0; fi21 < 18; fi21++ {
        for fi22 := 0; fi22 < 8; fi22++ {
            var v23 int32 = -1
            _ = len(v1)
            _ = fi22
            _ = v23
        }
        v24 := GenS3{F0: 747, F1: 4294967295}
        _ = fi21
        _ = v24
    }
    var v25 uint8 = (v4 * (v4 & v4))
    v25 = (v25 & (v25 | (v4 >> ((42) & 63))))
    v26 := GenS1{F0: -1}
    var v27 int64 = v2
    v28 := GenS2{F0: -128, F1: -1, F2: 255}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
