package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 string; F1 uint32; F2 int }

func main() {
    var v1 float32 = 3.14159
    var v2 float32 = 0.1
    var v3 int8 = -128
    v4 := make(map[string]int)
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]bool)
    var v7 uint64 = 42
    var v8 int = 127
    fmt.Println(v8, v1, v7)
    v9 := GenS1{F0: v7, F1: 0}
    for fi10 := 0; fi10 < 18; fi10++ {
        v2 = 1e10
        for fi11 := 0; fi11 < 5; fi11++ {
            _ = len(v6)
            delete(v6, ((("a" + ("foo bar" + "hello")) + "\n\t") + ""))
            _ = fi11
        }
        _ = fi10
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        if false {
            var v13 uint32 = 438
            v1 = float32(((-2.25 + (1e10 + -2.25)) - 1.5))
            fmt.Println(v1)
            _ = v13
        }
        for fi14 := 0; fi14 < 16; fi14++ {
            v15 := []int{-653}
            v3 = int8(v7)
            if false {
                v16 := GenS2{F0: "\n\t", F1: 2147483648, F2: fi14}
                _ = v16
            }
            _ = fi14
            _ = v15
        }
        fmt.Println(v3)
        _ = fi12
    }
    v1 = v1
    var v17 float32 = 3.14159
    var v18 float64 = -2.25
    fmt.Println(v18, v17, v2)
    if false {
        if (1e10 == (v18 - (0.0 * v18))) {
            var v19 uint8 = 223
            fmt.Println(v7)
            _ = v19
        } else {
            v7 = v7
            if true {
                v20 := GenS2{F0: ((("foo bar" + (("a" + (("" + "\n\t") + ((("" + "\n\t") + ("foo bar" + ("" + "a"))) + "foo bar"))) + "hello")) + "\n\t") + (("foo bar" + "a") + "hello")), F1: 2147483648, F2: (v8 << ((uint(v7)) & 63))}
                v21 := v5[2147483647]
                _ = v20
                _ = v21
            }
        }
        var v22 int32 = 12
        _ = v22
    }
    if (!true) {
        for fi23 := 0; fi23 < 15; fi23++ {
            switch "a" {
            case "hello":
                v3 = (v3 / (((v3 * v3)) | 1))
            default:
                switch 284 {
                case -2147483648:
                    v24 := GenS2{F0: "\n\t", F1: 1, F2: v8}
                    _ = v24
                }
            }
            _ = fi23
        }
        var v25 float32 = (v17 - (v2 / 3.14159))
        var v26 string = "a"
        _ = v25
        _ = v26
    } else {
        if true {
            fmt.Println(v2, v2)
            for fi27 := 0; fi27 < 20; fi27++ {
                var v28 float64 = 1.5
                v4[""] = v8
                var v29 int64 = -9223372036854775808
                _ = fi27
                _ = v28
                _ = v29
            }
            var v30 int64 = 0
            _ = v30
        }
    }
    var v31 bool = false
    v2 += 3.14159
    var v32 float32 = v1
    v32 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
