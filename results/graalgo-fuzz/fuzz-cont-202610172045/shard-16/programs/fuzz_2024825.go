package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int32 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8 }

func main() {
    v1 := []uint32{1}
    var v2 int64 = -982
    var v3 int8 = 1
    var v4 uint = 348
    var v5 int8 = (^(v3))
    v1 = append(v1, 0)
    v2 += v2
    v6 := GenS1{F0: 2147483648, F1: 42, F2: -922}
    for fi7 := 0; fi7 < 2; fi7++ {
        v1 = append(v1, 2147483647)
        var v8 uint = 1
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: -114, F1: "\n\t"}
    for fi10 := 0; fi10 < 19; fi10++ {
        switch "a" {
        case "foo bar":
            v11 := GenS3{F0: 65535}
            _ = v11
        case "\n\t":
            switch "" {
            case "hello":
                if false {
                    switch "hello" {
                    case "hello":
                        fmt.Println(v2, v5)
                    case "":
                        var v12 float32 = 1e10
                        _ = v12
                    default:
                        fmt.Println(v3, v5, v2)
                    }
                } else {
                    v2 += (v2 ^ v2)
                    var v13 float32 = 0.1
                    _ = v13
                }
            case "\n\t":
                if (!(65535 != 963)) {
                    fmt.Println(v4, fi10)
                    fmt.Println(v2)
                    for fi14 := 0; fi14 < 10; fi14++ {
                        fmt.Println(fi14, v5)
                        _ = fi14
                    }
                } else {
                    var v15 uint32 = 2147483647
                    var v16 int = fi10
                    v17 := GenS2{F0: 1000, F1: ("\n\t" + "hello")}
                    _ = v15
                    _ = v16
                    _ = v17
                }
            case "a":
                var v18 int32 = 42
                _ = v18
            }
        case "a":
            var v19 bool = (!((-2.25 * 0.1) == 0.0))
            _ = v19
        default:
            v1 = append(v1, 1)
        }
        var v20 int32 = -944
        _ = fi10
        _ = v20
    }
    switch ("" + ("hello" + ("foo bar" + "foo bar"))) {
    case "\n\t":
        fmt.Println(v4, v5)
    }
    v21 := GenS3{F0: 891}
    switch ("a" + "\n\t") {
    case "a":
        v5 = -36
    case "foo bar":
        var v22 float64 = 3.14159
        _ = v22
    default:
        if true {
            v4 += (v4 * 1)
            for fi23 := 0; fi23 < 8; fi23++ {
                v24 := GenS4{F0: (v3 >> (((v4 >> ((42) & 63))) & 63))}
                _ = fi23
                _ = v24
            }
            v25 := GenS2{F0: -1, F1: "\n\t"}
            _ = v25
        }
    }
    v26 := GenS5{F0: v3}
    switch 9223372036854775807 {
    case 178:
        _ = len(v1)
    case 4294967295:
        var v27 bool = (v4 < v4)
        _ = v27
    }
    var v28 uint = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
