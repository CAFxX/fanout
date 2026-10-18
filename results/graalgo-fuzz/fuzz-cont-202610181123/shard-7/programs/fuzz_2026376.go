package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 uint; F2 int16 }
type GenS3 struct { F0 int16; F1 uint32; F2 int32 }
type GenS4 struct { F0 int }

func main() {
    v1 := make(map[int]string)
    v2 := []string{"a", "foo bar"}
    var v3 uint32 = 65535
    v4 := make(map[int]string)
    v5 := GenS1{F0: (-1 >= 4294967295)}
    var v6 int32 = 2147483647
    var v7 uint32 = (^(v3))
    v8 := GenS1{F0: (0 != 255)}
    v3 -= v7
    for fi9 := 0; fi9 < 18; fi9++ {
        fmt.Println(v3, v6)
        v10 := GenS2{F0: true, F1: 955, F2: -32768}
        switch ("hello" + "hello") {
        case "a":
            if true {
                if (!(841 > (v7 >> ((42) & 63)))) {
                    v3 -= (v7 << ((1) & 63))
                }
            }
        case "\n\t":
            v7 -= 65535
        default:
            v2 = append(v2, (((("a" + (("hello" + "foo bar") + "hello")) + "foo bar") + "\n\t") + "foo bar"))
        }
        _ = fi9
        _ = v10
    }
    v7 -= (v7 >> ((767) & 63))
    var v11 int = 0
    v3 = (v3 | 83)
    v7 += v7
    for fi12 := 0; fi12 < 6; fi12++ {
        for fi13 := 0; fi13 < 9; fi13++ {
            v14 := GenS3{F0: -1, F1: (v3 ^ (v7 ^ (v7 - v7))), F2: v6}
            var v15 uint16 = 0
            if (!true) {
                delete(v1, (fi13 & fi12))
            } else {
                v16 := GenS4{F0: 32768}
                _ = v16
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        for fi17 := 0; fi17 < 14; fi17++ {
            switch "\n\t" {
            case "":
                var v18 int16 = int16(fi17)
                _ = v18
            case "a":
                v6 -= v6
            }
            fmt.Println(v7)
            _ = fi17
        }
        fmt.Println(v6, v7, v3)
        _ = fi12
    }
    var v19 int64 = -9223372036854775808
    v3 -= (^(v3))
    var v20 int64 = (v19 * (v19 % ((v19) | 1)))
    v19 += (v19 * 42)
    for fi21 := 0; fi21 < 5; fi21++ {
        v7 = 1
        _ = fi21
    }
    fmt.Println(v19, v11, v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
