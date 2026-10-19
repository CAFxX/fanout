package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 uint64 = 42
    var v2 float32 = 3.14159
    var v3 bool = false
    var v4 uint64 = 4294967295
    v5 := make(map[int]bool)
    v1 = (v4 | (v4 >> ((2) & 63)))
    var v6 int64 = 9223372036854775807
    switch "foo bar" {
    case "a":
        v7 := GenS1{F0: 42}
        _ = v7
    default:
        v8 := GenS1{F0: 1}
        _ = v8
    }
    if v3 {
        for fi9 := 0; fi9 < 6; fi9++ {
            v5[fi9] = true
            var v10 float32 = 0.0
            for fi11 := 0; fi11 < 13; fi11++ {
                v1 = (^(v1))
                _ = len(v5)
                _ = fi11
            }
            _ = fi9
            _ = v10
        }
    }
    switch (("\n\t" + "foo bar") + (("\n\t" + ("hello" + "a")) + "foo bar")) {
    case "a":
        v12 := []int{}
        _ = v12
    }
    fmt.Println(v2)
    var v13 uint = 255
    if false {
        var v14 uint = (v13 << ((v13) & 63))
        var v15 int16 = 32767
        _ = v14
        _ = v15
    } else {
        if v3 {
            switch "" {
            case "hello":
                fmt.Println(v13, v13, v3)
            default:
                v2 = v2
            }
            v4 = 2
            fmt.Println(v13, v1, v6)
        } else {
            v2 -= ((0.0 + float32(((1.5 + 1.5) / 3.14159))) + 3.14159)
            v3 = (1 < 171)
        }
        for fi16 := 0; fi16 < 16; fi16++ {
            fmt.Println(v1)
            for fi17 := 0; fi17 < 4; fi17++ {
                fmt.Println(v3, v13)
                if (1e10 < 1e10) {
                    var v18 uint64 = (v4 / ((18446744073709551615) | 1))
                    var v19 int8 = -94
                    fmt.Println(v1, fi17, v1)
                    _ = v18
                    _ = v19
                } else {
                    var v20 uint16 = 2
                    _ = v20
                }
                _ = fi17
            }
            _ = fi16
        }
    }
    var v21 int8 = 127
    if (!v3) {
        if true {
            v5[2147483647] = v3
        }
        var v22 int64 = (v6 - -794)
        _ = v22
    }
    v23 := []int{}
    var v24 int8 = v21
    v13 = (v13 | (v13 + (^(v13))))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
