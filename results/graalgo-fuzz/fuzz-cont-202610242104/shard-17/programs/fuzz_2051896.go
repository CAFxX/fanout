package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint32 }

func main() {
    var v1 int = 1
    var v2 float64 = 1.5
    v3 := []string{}
    v4 := make(map[string]string)
    v5 := make(map[int]string)
    var v6 int64 = -1
    var v7 int8 = 127
    var v8 int8 = v7
    fmt.Println(v1, v7)
    v6 += v6
    v1 += 0
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := GenS1{F0: (v7 | v8), F1: 4294967295}
        _ = fi9
        _ = v10
    }
    switch "hello" {
    case "":
        v11 := GenS2{F0: -32768, F1: 42, F2: 42}
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 int32 = -2147483648
            switch "hello" {
            case "":
                v5[(v1 - v1)] = "\n\t"
            default:
                v8 = (v7 * 100)
            }
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        v2 = v2
        v6 += v6
        _ = fi15
    }
    var v16 string = (((("foo bar" + "foo bar") + ((("foo bar" + "a") + ("hello" + "a")) + "\n\t")) + "foo bar") + ("a" + ""))
    v8 -= v8
    switch ((v16 + "a") + ((((("\n\t" + v16) + v16) + "\n\t") + v16) + ("" + v16))) {
    case "foo bar":
        if true {
            v7 = (v7 >> ((567) & 63))
        }
    case "\n\t":
        v17 := GenS1{F0: 100, F1: 0}
        _ = v17
    default:
        for fi18 := 0; fi18 < 4; fi18++ {
            var v19 int64 = (v6 - v6)
            v6 = (v6 | 0)
            for fi20 := 0; fi20 < 6; fi20++ {
                if true {
                    if (!true) {
                        v1 = 32767
                    }
                    v16 = v16
                } else {
                    fmt.Println(fi20)
                    v1 -= (v1 >> ((2) & 63))
                }
                fmt.Println(fi18)
                _ = fi20
            }
            _ = fi18
            _ = v19
        }
    }
    if false {
        v7 = int8(v1)
    }
    if (!(true && (true || false))) {
        v8 = v8
        v16 = ("a" + ((v16 + ("foo bar" + (v16 + v16))) + ("a" + "")))
        var v21 string = v16
        _ = v21
    }
    v7 = v8
    var v22 int64 = (v6 << ((4294967295) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
