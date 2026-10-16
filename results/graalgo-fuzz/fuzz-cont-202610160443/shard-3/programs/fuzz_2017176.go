package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS3 struct { F0 uint; F1 uint16; F2 string }
func genFn0(p0 int32, p1 uint8, p2 int8) uint32 {
    p1 = 2
    v5 := GenS1{F0: ("a" + "hello")}
    p1 = p1
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 2
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[uint32]int)
    var v3 bool = (!true)
    v4 := []string{"hello", "hello", "hello"}
    v6 := GenS2{F0: 2, F1: genFn0(42, 2, -1), F2: 2}
    fmt.Println(v3, v3)
    v7 := GenS2{F0: 1, F1: genFn0(-1, 162, 127), F2: 128}
    for fi8 := 0; fi8 < 17; fi8++ {
        fmt.Println(v3, v1, v3)
        v9 := GenS3{F0: 255, F1: 2, F2: "a"}
        _ = fi8
        _ = v9
    }
    var v10 int8 = 0
    var v11 uint = 18446744073709551615
    switch "a" {
    case "\n\t":
        fmt.Println(v1, v10)
    case "":
        if false {
            fmt.Println(v10, v1)
        } else {
            v11 += 361
        }
    case "a":
        v12 := GenS3{F0: 42, F1: 42, F2: ((("hello" + ("a" + "hello")) + "") + (("\n\t" + "\n\t") + "\n\t"))}
        _ = v12
    default:
        delete(v2, 65535)
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        if false {
            var v14 int16 = (^(v1))
            if v3 {
                if v3 {
                    switch (fi13 & 522) {
                    case 42:
                        v3 = ((v10 % ((int8(v11)) | 1)) < (v10 + v10))
                    }
                    if (false || (!true)) {
                        v3 = ((v3 || ((true || (("foo bar" >= "foo bar") || (!(255 == 85)))) || false)) && ((v10 | v10) <= (v10 | 84)))
                        var v15 uint = (v11 * (v11 + v11))
                        v10 = v10
                        _ = v15
                    } else {
                        v10 -= v10
                    }
                }
                v16 := GenS3{F0: (v11 * 42), F1: 121, F2: "hello"}
                _ = v16
            }
            _ = v14
        }
        if (-1 >= 1) {
            v17 := GenS2{F0: 204, F1: 2147483648, F2: 128}
            var v18 uint16 = 32767
            if len(v4) > 0 {
                v4[0] = ""
            }
            _ = v17
            _ = v18
        }
        _ = fi13
    }
    v19 := GenS1{F0: ""}
    v11 = (v11 * 9)
    v11 = (v11 - v11)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
