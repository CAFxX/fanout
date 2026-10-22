package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{392, 4294967295, 280, 2}
    var v3 uint16 = 42
    var v4 uint64 = 18446744073709551615
    v3 += v3
    v3 -= (v3 >> ((0) & 63))
    switch 4294967295 {
    case 32767:
        for fi5 := 0; fi5 < 19; fi5++ {
            fmt.Println(v3, fi5, v4)
            _ = fi5
        }
    case -534:
        switch "hello" {
        case "\n\t":
            for fi6 := 0; fi6 < 7; fi6++ {
                var v7 bool = true
                switch "\n\t" {
                case "":
                    fmt.Println(v3, v3, v7)
                }
                if (!(-2.25 == 1e10)) {
                    v3 -= v3
                    v8 := GenS1{F0: 42}
                    _ = v8
                } else {
                    var v9 int8 = 100
                    v3 = (v3 / ((32767) | 1))
                    v2 = append(v2, 1)
                    _ = v9
                }
                _ = fi6
                _ = v7
            }
        case "hello":
            var v10 int8 = int8(v4)
            _ = v10
        default:
            fmt.Println(v3, v3)
        }
    case -774:
        var v11 uint = 2
        _ = v11
    default:
        var v12 int = -985
        _ = v12
    }
    if true {
        var v13 int64 = 9223372036854775807
        fmt.Println(v3)
        _ = v13
    }
    v4 = 2
    switch "hello" {
    case "foo bar":
        fmt.Println(v3, v3)
    }
    var v14 uint64 = v4
    _ = len(v2)
    fmt.Println(v3)
    if (42 < 128) {
        _ = len(v1)
        v3 = 833
        v14 = (v4 & (v14 - (v4 - v14)))
    } else {
        if (42 < 316) {
            if false {
                v14 = 976
                if (false || (false || (((42 != 261) || true) || true))) {
                    var v15 int8 = 1
                    var v16 int16 = 42
                    _ = v15
                    _ = v16
                }
                for fi17 := 0; fi17 < 19; fi17++ {
                    var v18 float64 = 1.5
                    for fi19 := 0; fi19 < 17; fi19++ {
                        fmt.Println(v18)
                        v14 += 4294967295
                        _ = fi19
                    }
                    v14 = uint64(v3)
                    _ = fi17
                    _ = v18
                }
            } else {
                var v20 string = "\n\t"
                var v21 int8 = 1
                _ = v20
                _ = v21
            }
            var v22 uint8 = 0
            var v23 uint = 2
            _ = v22
            _ = v23
        }
    }
    switch ("\n\t" + "") {
    case "\n\t":
        if false {
            var v24 int32 = -2147483648
            v3 = 42
            var v25 uint = 521
            _ = v24
            _ = v25
        }
    case "hello":
        v26 := GenS1{F0: uint(v3)}
        _ = v26
    case "foo bar":
        var v27 string = "\n\t"
        _ = v27
    default:
        v4 = (^(v14))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
