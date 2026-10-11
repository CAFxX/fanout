package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
func genFn0(p0 int32, p1 int8, p2 int) string {
    p1 = p1
    fmt.Println(p0)
    p2 = (^(p2))
    _ = p0
    _ = p1
    _ = p2
    return "a"
}

func main() {
    v1 := []uint32{}
    var v2 uint = 24
    var v3 uint8 = 1
    var v4 int8 = 0
    var v5 int16 = 42
    v6 := make(map[uint32]uint32)
    var v7 string = "\n\t"
    fmt.Println(v7, v4, v7)
    var v8 int16 = (v5 ^ -1)
    delete(v6, 65535)
    var v9 int32 = int32(v4)
    switch ("foo bar" + ((v7 + (v7 + v7)) + v7)) {
    case "foo bar":
        v10 := GenS1{F0: 0, F1: 2147483647}
        _ = v10
    case "\n\t":
        if false {
            v7 = (v7 + "hello")
            for fi11 := 0; fi11 < 2; fi11++ {
                switch v7 {
                case "hello":
                    switch (fi11 & (fi11 << ((v2) & 63))) {
                    case -9223372036854775808:
                        v9 = (v9 << ((v2) & 63))
                    case 127:
                        var v12 int16 = v8
                        _ = v12
                    }
                }
                _ = fi11
            }
        } else {
            fmt.Println(v5, v4, v9)
            v3 = (v3 ^ (v3 | (v3 ^ 2)))
            var v13 float32 = 1.5
            _ = v13
        }
    }
    if ((v8 * (v5 + v5)) >= 1000) {
        v5 = (v8 - v5)
    } else {
        v14 := GenS1{F0: 42, F1: -350}
        v4 -= (^(v4))
        _ = v14
    }
    fmt.Println(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
