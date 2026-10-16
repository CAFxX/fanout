package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 int16; F2 uint32 }
func genFn0() int {
    var v4 int64 = 1
    v5 := GenS1{F0: 42}
    v4 = (v4 << ((65535) & 63))
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 string, p1 int64, p2 int16) uint32 {
    fmt.Println(p2, p1)
    var v6 uint8 = 128
    var v7 uint16 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    v1 := []string{"", "hello", ""}
    var v2 int16 = 1000
    v3 := make(map[int]bool)
    v2 = v2
    switch ((("\n\t" + "a") + "\n\t") + "a") {
    case "hello":
        if false {
            v8 := GenS2{F0: 0, F1: v2, F2: 2147483647}
            v3[0] = (42 <= -2147483648)
            _ = v8
        } else {
            v2 -= v2
            var v9 int64 = -515
            _ = v9
        }
    case "":
        v10 := GenS2{F0: 4294967295, F1: 1, F2: 2}
        _ = v10
    case "foo bar":
        if ("" < (("" + "hello") + "foo bar")) {
            v2 = v2
        }
    }
    switch ("" + "\n\t") {
    case "":
        var v11 bool = (!(false && false))
        _ = v11
    case "a":
        var v12 uint32 = genFn1("", 42, (v2 - v2))
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        v2 = (v2 >> ((255) & 63))
        _ = len(v3)
        var v14 float32 = 0.1
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: 991, F1: -32768, F2: 65535}
    v2 = v2
    v16 := GenS2{F0: 0, F1: v2, F2: 2}
    fmt.Println(v2)
    var v17 bool = (2 > genFn1("\n\t", -1, v2))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
