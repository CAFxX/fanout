package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint32, p1 int, p2 int64) uint8 {
    v6 := []int{p1, p1}
    fmt.Println(p2, p2)
    v7 := GenS1{F0: -68}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 2
}

func main() {
    v1 := []uint32{791, 42, 65535}
    var v2 uint16 = 1
    v3 := make(map[string]string)
    v4 := []string{("a" + "a"), "foo bar"}
    var v5 uint8 = 128
    v8 := GenS1{F0: -32768}
    v5 = 0
    var v9 int16 = 1
    v10 := GenS2{F0: 0, F1: v2, F2: 1}
    v4 = append(v4, "\n\t")
    v5 = 127
    if (false && (0 > 9223372036854775807)) {
        var v11 uint32 = 4294967295
        v3[((("" + ((("foo bar" + "\n\t") + "a") + "hello")) + ("hello" + "a")) + ("hello" + "hello"))] = "\n\t"
        _ = v11
    } else {
        switch "a" {
        case "\n\t":
            if (((("hello" + "hello") + "hello") + "a") < "\n\t") {
                switch "" {
                case "hello":
                    v12 := GenS3{F0: v9}
                    _ = v12
                default:
                    v13 := GenS3{F0: (v9 >> ((4294967295) & 63))}
                    _ = v13
                }
                var v14 uint8 = 42
                _ = v14
            } else {
                var v15 int32 = 0
                var v16 float32 = 3.14159
                _ = v15
                _ = v16
            }
        case "":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "foo bar":
            if false {
                v2 = (^(v2))
                var v17 string = "hello"
                v5 += v5
                _ = v17
            }
        }
    }
    if ((!false) || ((2147483648 > 2) || (3.14159 == 0.1))) {
        v18 := GenS1{F0: v9}
        _ = v18
    }
    v9 -= v9
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
