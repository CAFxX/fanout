package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    var v3 int = 4294967295
    v4 := []string{(("foo bar" + ("\n\t" + "foo bar")) + ("foo bar" + "")), "a", "a"}
    fmt.Println(v3)
    switch "a" {
    case "foo bar":
        var v5 float64 = float64(1.5)
        _ = v5
    }
    v3 = (^(v3))
    var v6 int = v3
    fmt.Println(v3)
    var v7 string = ""
    v6 += v3
    if true {
        fmt.Println(v7)
    }
    var v8 int32 = -1
    switch (v3 >> ((575) & 63)) {
    case 4294967295:
        v8 -= -1
    case 2147483647:
        if false {
            if (2 != 65535) {
                var v9 uint = 255
                _ = v9
            }
            for fi10 := 0; fi10 < 16; fi10++ {
                v6 += (v6 / (((v3 - 128)) | 1))
                v3 += (v3 + (v6 & (v6 | 32767)))
                var v11 uint = 2
                _ = fi10
                _ = v11
            }
            v12 := GenS1{F0: 65535, F1: 0, F2: 32767}
            _ = v12
        } else {
            if (0.1 == -2.25) {
                v3 -= (v3 >> ((4294967295) & 63))
                v7 = (((("\n\t" + v7) + v7) + "foo bar") + v7)
            } else {
                v4 = append(v4, ("\n\t" + (((((v7 + "hello") + ((v7 + v7) + (v7 + "foo bar"))) + "") + v7) + v7)))
                v8 += (v8 % (((^(v8))) | 1))
            }
        }
    }
    v8 -= (-(v8))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
