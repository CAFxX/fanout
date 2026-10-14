package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int16; F1 int16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    v3 := []string{(("" + "foo bar") + "\n\t"), "hello", "hello"}
    v4 := []uint32{529}
    switch "\n\t" {
    case "a":
        switch 909 {
        case 42:
            if false {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v5 := GenS1{F0: 2, F1: 0}
                _ = v5
            } else {
                switch "foo bar" {
                case "":
                    if (42 == 65535) {
                        var v6 uint64 = 18446744073709551615
                        _ = v6
                    }
                case "a":
                    for fi7 := 0; fi7 < 16; fi7++ {
                        var v8 int16 = -32768
                        _ = fi7
                        _ = v8
                    }
                }
                switch 32767 {
                case 2147483647:
                    v9 := GenS1{F0: 0, F1: 42}
                    _ = v9
                default:
                    v10 := GenS1{F0: 4294967295, F1: -1}
                    _ = v10
                }
            }
        case -9223372036854775808:
            if true {
                var v11 float64 = 0.0
                _ = v11
            } else {
                var v12 uint32 = 65535
                fmt.Println(v12)
                _ = v12
            }
        case 122:
            var v13 float32 = 0.0
            _ = v13
        }
    default:
        if (!true) {
            v14 := GenS2{F0: 316, F1: -32768}
            _ = v14
        }
    }
    v3 = append(v3, "")
    var v15 float32 = (0.1 / 1.0)
    v15 -= 3.14159
    v15 += 1e10
    v15 = 3.14159
    _ = len(v1)
    var v16 uint = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
