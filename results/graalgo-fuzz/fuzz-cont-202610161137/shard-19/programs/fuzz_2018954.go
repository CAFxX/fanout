package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int8; F1 int32; F2 int16 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    var v4 uint64 = 557
    _ = len(v3)
    fmt.Println(v4)
    fmt.Println(v4)
    if (false || (646 != 620)) {
        var v5 uint = 0
        if ((2 == 2147483647) || (!true)) {
            v6 := GenS1{F0: v4, F1: 2}
            _ = v6
        } else {
            if true {
                fmt.Println(v4, v5)
                v5 += 177
            } else {
                v5 -= 221
            }
            v7 := []int{int(v5), 32767, 42}
            _ = v7
        }
        _ = v5
    } else {
        if ("foo bar" == "\n\t") {
            v8 := GenS2{F0: 42, F1: -1, F2: 1}
            v4 = v4
            _ = v8
        } else {
            if true {
                fmt.Println(v4)
                var v9 uint8 = 0
                _ = v9
            }
            var v10 uint8 = 0
            _ = v10
        }
        v11 := GenS2{F0: -128, F1: 2147483647, F2: -891}
        _ = v11
    }
    v4 -= (v4 - v4)
    var v12 int8 = -128
    fmt.Println(v4)
    _ = len(v1)
    fmt.Println(v4, v12)
    v1["\n\t"] = "foo bar"
    var v13 string = ("\n\t" + "\n\t")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
