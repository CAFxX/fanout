package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 string }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    v3 := make(map[int]string)
    var v4 string = (("\n\t" + (("a" + "\n\t") + (("hello" + "") + "\n\t"))) + ("\n\t" + "foo bar"))
    if true {
        var v5 uint16 = 65535
        _ = v5
    }
    v4 = (("foo bar" + v4) + v4)
    fmt.Println(v4)
    var v6 uint = 36
    var v7 int64 = 42
    if (!(32767 < -874)) {
        var v8 uint8 = 42
        switch 32768 {
        case -9223372036854775808:
            if true {
                v6 += v6
            } else {
                v9 := GenS1{F0: -156, F1: -9223372036854775808, F2: v4}
                var v10 uint64 = 2
                _ = v10
                _ = v9
            }
        }
        var v11 int64 = v7
        _ = v11
        _ = v8
    } else {
        var v12 int = 1
        v6 = (v6 / (((v6 >> (((v6 + (v6 + v6))) & 63))) | 1))
        _ = v12
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 float32 = (-2.25 / float32(1.5))
        v7 += v7
        _ = fi13
        _ = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
