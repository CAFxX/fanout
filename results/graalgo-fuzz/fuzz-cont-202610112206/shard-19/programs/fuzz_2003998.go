package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int16 }
func genFn0(p0 uint64, p1 uint32) int64 {
    p1 = uint32(p0)
    v6 := GenS1{F0: 561, F1: 121}
    _ = p0
    _ = p1
    _ = v6
    return 1
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{}
    v3 := make(map[uint32]string)
    v4 := []string{}
    var v5 int64 = 634
    v5 -= (v5 % ((genFn0(2, 4294967295)) | 1))
    v5 = v5
    v7 := GenS2{F0: 52, F1: -1}
    v5 -= -1
    v8 := GenS1{F0: 65535, F1: 0}
    v1[32767] = "hello"
    switch -262 {
    case -2147483648:
        fmt.Println(v5)
    }
    switch "hello" {
    case "":
        switch (("a" + "\n\t") + ("hello" + ("a" + ("a" + "foo bar")))) {
        case "hello":
            v5 -= v5
        case "\n\t":
            var v9 uint32 = 0
            _ = v9
        case "":
            var v10 int8 = 0
            _ = v10
        }
    case "a":
        var v11 string = "foo bar"
        _ = v11
    case "hello":
        v5 -= v5
    default:
        if false {
            if true {
                var v12 int32 = 42
                _ = v12
            }
            if (!(false && true)) {
                v5 = v5
            } else {
                if (42 <= 1) {
                    var v13 uint8 = 121
                    v14 := GenS2{F0: 489, F1: -281}
                    _ = v13
                    _ = v14
                }
            }
            var v15 float32 = 0.1
            _ = v15
        }
    }
    v5 += (v5 / (((^(v5))) | 1))
    var v16 float32 = (0.1 * 1e10)
    v5 = (v5 << ((4294967295) & 63))
    fmt.Println(v16, v5)
    v5 = 42
    v5 += v5
    v5 = -1
    v5 -= v5
    var v17 string = "hello"
    v4 = append(v4, (v17 + ""))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
