package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 uint16 = 42
    v2 := make(map[string]bool)
    v3 := []string{"foo bar", ("hello" + "a"), "hello"}
    var v4 int16 = -32768
    v5 := make(map[uint32]int)
    v6 := []uint32{42, 626, 42}
    fmt.Println(v4, v1)
    v4 -= v4
    v1 = (v1 + (v1 * 1))
    v1 -= (v1 * v1)
    if (!(!false)) {
        if (9223372036854775807 > 2147483647) {
            if (!true) {
                var v7 bool = (v4 > v4)
                fmt.Println(v4, v7)
                _ = v7
            } else {
                v8 := v5[uint32(v1)]
                _ = v8
            }
            fmt.Println(v4, v1)
        }
        var v9 uint8 = 127
        _ = v9
    } else {
        v10 := GenS1{F0: (v4 / ((v4) | 1))}
        _ = v10
    }
    v1 += 32767
    if (!(((false || (1.5 < 3.14159)) && true) || false)) {
        var v11 bool = false
        _ = v11
    } else {
        var v12 uint = 65535
        var v13 float32 = (1.5 / (-2.25 * float32(1.5)))
        switch "a" {
        case "hello":
            var v14 uint8 = 255
            _ = v14
        }
        _ = v12
        _ = v13
    }
    v15 := GenS1{F0: (v4 * (^(v4)))}
    if (42 < 65535) {
        var v16 uint32 = 2147483647
        _ = v16
    } else {
        switch "\n\t" {
        case "":
            _ = len(v3)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
