package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint8; F1 int64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]string)
    var v3 uint64 = 379
    v4 := make(map[string]uint32)
    var v5 bool = (false && (1 >= 42))
    var v6 bool = (2 < 255)
    fmt.Println(v3)
    v6 = true
    var v7 int16 = -1
    v7 += int16(v3)
    if v6 {
        v5 = (2147483647 != -502)
        for fi8 := 0; fi8 < 7; fi8++ {
            fmt.Println(v5)
            v9 := []int{32768}
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: 0, F1: 9223372036854775807}
        _ = v10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 uint16 = 2
        if v5 {
            var v13 float32 = 1.5
            if (v6 || false) {
                v14 := GenS2{F0: 0, F1: 0}
                _ = v14
            }
            _ = v13
        }
        var v15 float32 = 1.5
        _ = fi11
        _ = v12
        _ = v15
    }
    fmt.Println(v7, v6)
    var v16 uint16 = 631
    var v17 string = (("a" + (((("" + "\n\t") + ("foo bar" + ("foo bar" + ""))) + ((("hello" + (("foo bar" + "foo bar") + "\n\t")) + "foo bar") + "")) + ("\n\t" + "foo bar"))) + ("a" + "\n\t"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
