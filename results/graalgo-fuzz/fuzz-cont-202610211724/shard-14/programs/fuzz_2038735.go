package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint }

func main() {
    var v1 int = 42
    v2 := make(map[string]uint32)
    var v3 uint32 = 4294967295
    switch (v1 << ((2) & 63)) {
    case 127:
        var v4 uint = 2
        _ = v4
    case 128:
        var v5 int64 = -1
        _ = v5
    case 2147483647:
        var v6 float32 = 0.1
        _ = v6
    }
    fmt.Println(v1, v3)
    var v7 int32 = -114
    v8 := []int{v1, int(v3), (v1 * (v1 + (v1 ^ (v1 >> ((255) & 63))))), (v1 >> ((2) & 63))}
    v7 = 128
    v2["hello"] = 241
    delete(v2, (("hello" + ("hello" + "")) + ("hello" + "")))
    v2[("hello" + "hello")] = v3
    switch "\n\t" {
    case "\n\t":
        if false {
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 int64 = 115
                _ = len(v2)
                _ = fi9
                _ = v10
            }
            for fi11 := 0; fi11 < 19; fi11++ {
                for fi12 := 0; fi12 < 20; fi12++ {
                    if (0.0 == (1e10 / float32(1e10))) {
                        _ = len(v2)
                        v8 = append(v8, (fi12 * -364))
                        var v13 int8 = 100
                        _ = v13
                    } else {
                        var v14 int = v1
                        _ = v14
                    }
                    var v15 bool = false
                    switch v1 {
                    case -2147483648:
                        var v16 uint16 = 42
                        _ = v16
                    case 0:
                        _ = len(v8)
                    case 2147483647:
                        v17 := GenS1{F0: 968}
                        _ = v17
                    }
                    v1 = 32767
                    v7 = v7
                    _ = fi12
                    _ = v15
                }
                _ = fi11
            }
            fmt.Println(v7, v3)
        }
    case "foo bar":
        switch "a" {
        case "foo bar":
            v18 := GenS1{F0: v1}
            _ = v18
        case "a":
            var v19 uint = 65535
            _ = v19
        default:
            var v20 int = (v1 | -9223372036854775808)
            _ = v20
        }
    case "hello":
        var v21 int32 = v7
        _ = v21
    }
    v3 += v3
    v22 := GenS2{F0: 127, F1: 128}
    v23 := GenS3{F0: 42}
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 int32 = v7
        if false {
            var v26 int16 = -814
            var v27 float64 = (1.5 - (1e10 - (0.1 * -2.25)))
            v25 = v25
            _ = v26
            _ = v27
        }
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v7
    _ = v8
}
