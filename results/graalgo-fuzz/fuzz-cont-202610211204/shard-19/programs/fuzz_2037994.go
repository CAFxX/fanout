package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int8; F1 int; F2 uint64 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 155
    v3 := make(map[int]string)
    var v4 int32 = int32(v2)
    for fi5 := 0; fi5 < 4; fi5++ {
        if (!true) {
            v2 += v2
            v4 = 0
            var v6 int32 = -2147483648
            _ = v6
        } else {
            if (!true) {
                delete(v1, 128)
            } else {
                v4 = int32(fi5)
            }
            v7 := GenS1{F0: 758, F1: -1}
            _ = v7
        }
        _ = len(v3)
        _ = fi5
    }
    fmt.Println(v2)
    fmt.Println(v2)
    v4 += v4
    var v8 int8 = -128
    var v9 uint64 = 42
    v9 = v9
    if (!true) {
        v2 += 344
        v8 -= (v8 | v8)
    } else {
        v2 = v2
        v8 -= -128
    }
    v4 += 0
    v10 := []int{}
    if (!true) {
        v4 -= (v4 << ((707) & 63))
        switch ("hello" + "\n\t") {
        case "hello":
            var v11 int32 = (-(v4))
            _ = v11
        case "foo bar":
            switch "hello" {
            case "":
                fmt.Println(v8, v4, v2)
            case "\n\t":
                var v12 uint32 = v2
                _ = v12
            case "a":
                if (-32768 <= -728) {
                    var v13 string = ("\n\t" + (("a" + ("\n\t" + "hello")) + ("" + ("" + ("hello" + "a")))))
                    _ = v13
                }
            }
        }
    }
    var v14 uint8 = 23
    var v15 bool = true
    v16 := GenS2{F0: (v8 * (^(v8))), F1: 9223372036854775807, F2: 18446744073709551615}
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
