package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int; F1 int64 }
func genFn0(p0 int16, p1 int64) uint16 {
    v7 := []int{-9223372036854775808}
    if (199 < 2147483647) {
        p0 = 968
        fmt.Println(p0, p1)
        var v8 uint16 = 1
        _ = v8
    } else {
        p1 += -717
    }
    switch "" {
    case "":
        for fi9 := 0; fi9 < 4; fi9++ {
            if ((-1 >= 100) || (128 == 2)) {
                switch "foo bar" {
                case "hello":
                    p1 = p1
                default:
                    fmt.Println(p1)
                }
            }
            p1 = 62
            v7 = append(v7, fi9)
            _ = fi9
        }
    default:
        switch ("hello" + ("a" + "a")) {
        case "hello":
            var v10 uint8 = 0
            _ = v10
        case "\n\t":
            if true {
                if len(v7) > 0 {
                    v7[0] = 4294967295
                }
                var v11 uint32 = 0
                _ = v11
            }
        default:
            for fi12 := 0; fi12 < 10; fi12++ {
                var v13 int64 = 42
                var v14 uint64 = 42
                _ = fi12
                _ = v13
                _ = v14
            }
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 32767
}

func main() {
    var v1 int16 = 289
    v2 := []int{9223372036854775807, 32767, 32767, 1}
    v3 := []uint32{}
    var v4 string = ("hello" + "foo bar")
    var v5 float32 = (3.14159 * 1.5)
    v6 := make(map[int]int)
    var v15 int16 = -995
    fmt.Println(v4, v15, v15)
    _ = len(v2)
    for fi16 := 0; fi16 < 6; fi16++ {
        if len(v2) > 0 {
            v2[0] = (fi16 + fi16)
        }
        switch (fi16 * 42) {
        case 2147483647:
            v17 := GenS1{F0: 118, F1: 911}
            _ = v17
        case 42:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            v18 := GenS1{F0: (v15 % ((1000) | 1)), F1: 2}
            _ = v18
        }
        _ = fi16
    }
    fmt.Println(v15)
    var v19 float64 = -2.25
    fmt.Println(v15)
    v6[9223372036854775807] = 128
    switch -9223372036854775808 {
    case 0:
        if false {
            v1 += (v1 >> ((617) & 63))
            var v20 bool = ((56 >= 100) && false)
            var v21 uint = 1
            _ = v20
            _ = v21
        }
    }
    var v22 float32 = ((3.14159 - (((1e10 / ((0.0 + v5) - 0.1)) / (v5 - (0.1 + v5))) - v5)) + (v5 / (v5 - v5)))
    v22 = (1.5 / v22)
    v1 = v15
    v23 := GenS2{F0: -972, F1: 60}
    _ = len(v6)
    v19 = (v19 / v19)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
