package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int8 }
type GenS2 struct { F0 int; F1 int64; F2 int16 }
type GenS3 struct { F0 uint32; F1 bool; F2 uint64 }
func genFn0(p0 uint) int16 {
    p0 = (^(p0))
    p0 = 18446744073709551615
    if true {
        v6 := GenS1{F0: 4294967295, F1: "a", F2: 42}
        v7 := []int{}
        var v8 string = (("\n\t" + ("a" + "")) + "")
        _ = v6
        _ = v7
        _ = v8
    } else {
        var v9 int8 = 55
        var v10 uint32 = 764
        for fi11 := 0; fi11 < 7; fi11++ {
            v9 = -30
            v12 := GenS2{F0: (fi11 >> ((97) & 63)), F1: 9223372036854775807, F2: 1000}
            switch "" {
            case "foo bar":
                var v13 uint = p0
                _ = v13
            case "hello":
                switch "\n\t" {
                case "a":
                    var v14 uint = (^(p0))
                    _ = v14
                }
            default:
                if (false || (!((^(v10)) >= (v10 / (((v10 - v10)) | 1))))) {
                    v15 := GenS2{F0: 128, F1: -970, F2: 32767}
                    var v16 uint32 = v10
                    _ = v15
                    _ = v16
                } else {
                    var v17 string = "a"
                    var v18 float64 = 3.14159
                    var v19 int32 = -1
                    _ = v17
                    _ = v18
                    _ = v19
                }
            }
            _ = fi11
            _ = v12
        }
        _ = v10
        _ = v9
    }
    _ = p0
    return -1
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[string]uint32)
    var v3 uint32 = 48
    v4 := make(map[int]uint32)
    v5 := []uint32{(v3 - 4294967295), 65535}
    switch -294 {
    case 9223372036854775807:
        if (32767 <= 0) {
            switch 32767 {
            case -338:
                v1 = v1
            case 127:
                fmt.Println(v1)
            }
            var v20 int64 = 0
            fmt.Println(v3)
            _ = v20
        } else {
            v3 = (^(v3))
        }
    case -2147483648:
        if true {
            fmt.Println(v3, v1)
        }
    case 32768:
        v1 = (v1 ^ v1)
    }
    fmt.Println(v3, v3)
    v21 := GenS3{F0: (v3 / ((65535) | 1)), F1: false, F2: 0}
    fmt.Println(v3, v3)
    var v22 string = ("a" + ((("\n\t" + (("\n\t" + ((("" + "foo bar") + "hello") + "")) + ("hello" + ("foo bar" + "foo bar")))) + "hello") + "\n\t"))
    switch "" {
    case "foo bar":
        v23 := GenS2{F0: 32767, F1: 159, F2: -1}
        _ = v23
    case "hello":
        fmt.Println(v1)
    case "\n\t":
        fmt.Println(v22, v22, v3)
    }
    var v24 uint32 = (v3 + v3)
    v5 = append(v5, 42)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
