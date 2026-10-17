package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int }
type GenS2 struct { F0 uint64; F1 uint32 }

func main() {
    v1 := []int{-9223372036854775808, 4294967295}
    var v2 uint32 = 4294967295
    v3 := make(map[int]bool)
    var v4 int32 = 42
    var v5 uint8 = 199
    switch "hello" {
    case "\n\t":
        if false {
            var v6 bool = false
            var v7 int8 = 42
            _ = len(v3)
            _ = v6
            _ = v7
        }
    case "foo bar":
        switch 167 {
        case 32767:
            var v8 int16 = 0
            _ = v8
        case -107:
            v3[-321] = false
        case -551:
            fmt.Println(v4, v2, v4)
        }
    case "hello":
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 uint16 = 1
            _ = fi9
            _ = v10
        }
    }
    v2 -= (v2 | 65535)
    var v11 int16 = 32767
    var v12 int = -9223372036854775808
    var v13 uint32 = (v2 * (v2 << ((0) & 63)))
    v2 = 0
    var v14 int = (v12 << ((0) & 63))
    var v15 float32 = 3.14159
    if (true || (("foo bar" > ("hello" + "")) && false)) {
        if false {
            if true {
                var v16 int32 = (^(v4))
                _ = v16
            } else {
                fmt.Println(v11, v14)
                if (!(3.14159 == v15)) {
                    v13 += 2147483648
                    v14 -= v14
                } else {
                    v12 -= v14
                }
            }
            fmt.Println(v5)
            fmt.Println(v14)
        }
        v17 := GenS1{F0: (^(v4)), F1: ((("" + "a") + "foo bar") + "a"), F2: v14}
        if false {
            v18 := GenS2{F0: 1, F1: (^(v13))}
            var v19 uint32 = (v2 - 1)
            var v20 uint8 = (v5 + (v5 + (^(v5))))
            _ = v18
            _ = v19
            _ = v20
        }
        _ = v17
    } else {
        fmt.Println(v14)
        if false {
            for fi21 := 0; fi21 < 16; fi21++ {
                v11 += v11
                switch ("a" + "hello") {
                case "foo bar":
                    var v22 int64 = 42
                    _ = v22
                case "a":
                    fmt.Println(v5)
                }
                v5 += 6
                _ = fi21
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
