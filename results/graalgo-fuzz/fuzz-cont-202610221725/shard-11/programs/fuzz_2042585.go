package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint; F1 int16; F2 int16 }

func main() {
    var v1 uint8 = 205
    v2 := []uint32{610, 0, 245, 65535}
    var v3 uint32 = 42
    v4 := make(map[int]int)
    v5 := make(map[string]string)
    for fi6 := 0; fi6 < 3; fi6++ {
        v1 += v1
        _ = fi6
    }
    if (true || (-1 < -742)) {
        v3 -= (v3 & v3)
        var v7 int32 = 42
        _ = v7
    }
    v1 = (v1 >> ((65535) & 63))
    v3 += 652
    v8 := GenS1{F0: 840, F1: 18446744073709551615}
    var v9 uint16 = 42
    v5[("a" + (("" + (("foo bar" + "foo bar") + ((("hello" + "\n\t") + (("foo bar" + "\n\t") + (("hello" + "foo bar") + "a"))) + "hello"))) + "\n\t"))] = (("\n\t" + "\n\t") + "\n\t")
    if false {
        var v10 uint = 878
        _ = v10
    } else {
        if false {
            var v11 uint8 = uint8(v3)
            v3 += (v3 >> ((4294967295) & 63))
            _ = v11
        } else {
            v12 := GenS2{F0: 4294967295, F1: -32768, F2: 0}
            _ = v12
        }
        var v13 uint16 = v9
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
