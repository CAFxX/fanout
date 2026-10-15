package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int; F1 int64; F2 string }

func main() {
    v1 := []int{725, 2147483647, -9223372036854775808}
    v2 := []string{"\n\t"}
    v3 := make(map[int]bool)
    var v4 int = 42
    v5 := make(map[uint32]bool)
    fmt.Println(v4)
    v4 = (v4 + (v4 >> ((0) & 63)))
    var v6 uint = 2
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 float64 = 0.0
        switch "" {
        case "foo bar":
            v8 = ((0.1 / v8) / 3.14159)
        case "hello":
            v6 = v6
        case "a":
            var v9 int64 = -9223372036854775808
            _ = v9
        default:
            if (!true) {
                for fi10 := 0; fi10 < 7; fi10++ {
                    v11 := GenS1{F0: 800}
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                    _ = fi10
                    _ = v11
                }
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
                var v12 int64 = 9223372036854775807
                _ = v12
            }
        }
        for fi13 := 0; fi13 < 12; fi13++ {
            fmt.Println(v4, v8, fi7)
            v14 := GenS1{F0: 2147483648}
            var v15 int32 = 2147483647
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi7
        _ = v8
    }
    v16 := GenS2{F0: ("" + ""), F1: 42, F2: 624}
    v6 = (v6 & 42)
    var v17 float64 = -2.25
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 string = (((("foo bar" + "foo bar") + "hello") + "foo bar") + "a")
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        if (0 == 2) {
            fmt.Println(v4, v17)
            if false {
                for fi21 := 0; fi21 < 16; fi21++ {
                    switch "\n\t" {
                    case "\n\t":
                        var v22 uint32 = 2147483648
                        _ = v22
                    default:
                        v4 = fi20
                    }
                    _ = fi21
                }
            } else {
                v23 := GenS3{F0: -641, F1: -1, F2: ""}
                _ = v23
            }
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi20
    }
    var v24 int16 = int16(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
