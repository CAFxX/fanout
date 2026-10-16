package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 2
    var v4 int = 32768
    v5 := []string{("foo bar" + (("\n\t" + "") + "a")), "\n\t"}
    v6 := []string{"hello", ("foo bar" + "\n\t"), "hello"}
    fmt.Println(v4, v4)
    v4 = 128
    switch 9223372036854775807 {
    case 42:
        var v7 uint64 = 557
        _ = v7
    case 9223372036854775807:
        if false {
            if (v4 == -984) {
                v8 := v2[(v3 & (v3 - (v3 >> ((65535) & 63))))]
                var v9 uint = 0
                _ = v8
                _ = v9
            }
            switch v4 {
            case 42:
                v3 += (v3 | v3)
            default:
                if (100 == 0) {
                    var v10 string = "a"
                    _ = v10
                } else {
                    fmt.Println(v3, v3)
                    var v11 int8 = 127
                    _ = v11
                }
            }
        }
    case 299:
        v12 := GenS1{F0: 42}
        _ = v12
    default:
        if ("foo bar" > ("\n\t" + "hello")) {
            var v13 string = ("a" + ("a" + (("" + "\n\t") + "\n\t")))
            fmt.Println(v13, v4, v13)
            _ = v13
        }
    }
    var v14 float32 = 3.14159
    _ = len(v2)
    v2[v3] = (v3 & (v3 / (((^(v3))) | 1)))
    var v15 uint64 = 42
    _ = len(v2)
    v14 -= v14
    var v16 uint = 352
    var v17 uint8 = 208
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
