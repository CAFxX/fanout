package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int16 }
type GenS2 struct { F0 uint32; F1 int64 }
func genFn0(p0 bool, p1 uint64, p2 uint64) uint8 {
    p1 -= (p1 % (((p2 * 42)) | 1))
    var v5 uint16 = 885
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return uint8(v5)
}

func main() {
    var v1 int16 = 1
    var v2 int = 1
    v3 := []uint32{4294967295, 65535, 790, 0}
    v4 := make(map[int]string)
    v2 = 128
    v2 = (v2 % ((v2) | 1))
    if false {
        if false {
            switch "hello" {
            case "a":
                v6 := GenS1{F0: ("" + "foo bar"), F1: 276, F2: (v1 / ((v1) | 1))}
                _ = v6
            case "\n\t":
                v7 := GenS1{F0: "\n\t", F1: 30, F2: 1000}
                _ = v7
            default:
                if true {
                    var v8 bool = (false && true)
                    var v9 uint32 = 2147483648
                    _ = v8
                    _ = v9
                }
            }
        }
        v1 -= (-(v1))
        fmt.Println(v1, v1)
    } else {
        fmt.Println(v1, v2)
        delete(v4, (v2 & 128))
    }
    switch (("foo bar" + "") + "foo bar") {
    case "a":
        var v10 string = ("\n\t" + "a")
        _ = v10
    case "\n\t":
        if (2147483647 > -579) {
            v2 -= v2
        }
    }
    var v11 uint8 = genFn0(false, 1, 0)
    var v12 int32 = int32(v2)
    delete(v4, (v2 % ((v2) | 1)))
    v13 := GenS2{F0: 4294967295, F1: 358}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
