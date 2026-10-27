package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint32 }
func genFn0(p0 int16, p1 bool) int {
    p1 = false
    if ((p0 != (-(p0))) || false) {
        v7 := []int{}
        _ = v7
    } else {
        v8 := []int{-579, -2147483648}
        for fi9 := 0; fi9 < 1; fi9++ {
            for fi10 := 0; fi10 < 4; fi10++ {
                switch ("foo bar" + (("a" + (("foo bar" + ((("hello" + "a") + ("a" + "hello")) + "foo bar")) + "a")) + "\n\t")) {
                case "a":
                    var v11 uint8 = 1
                    _ = v11
                case "foo bar":
                    if p1 {
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        if len(v8) > 0 {
                            v8[0] = (fi10 >> ((4294967295) & 63))
                        }
                    } else {
                        p0 = (^(p0))
                    }
                default:
                    v12 := GenS1{F0: 255, F1: 282}
                    if len(v8) > 0 {
                        v8[0] = -9223372036854775808
                    }
                    _ = len(v8)
                    _ = v12
                }
                var v13 int8 = 127
                _ = fi10
                _ = v13
            }
            p1 = (!true)
            _ = fi9
        }
        v14 := GenS2{F0: 9223372036854775807, F1: 42}
        _ = v14
        _ = v8
    }
    var v15 int32 = -687
    _ = p0
    _ = p1
    _ = v15
    return 0
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]uint32)
    var v3 uint = 863
    v4 := make(map[uint32]string)
    v5 := make(map[int]bool)
    var v6 float64 = float64((((((-2.25 + (1.5 - 3.14159)) / 0.1) / 1e10) + -2.25) * 1.5))
    var v16 bool = false
    var v17 int32 = (v1 % ((v1) | 1))
    var v18 float32 = (1.5 + (1.5 * (((1e10 * (0.0 + 0.0)) / 1.5) + 1e10)))
    v1 = -628
    var v19 uint32 = 65535
    fmt.Println(v16)
    fmt.Println(v3, v18, v18)
    v20 := []int{-884}
    var v21 string = (("hello" + "") + "\n\t")
    for fi22 := 0; fi22 < 10; fi22++ {
        fmt.Println(v3)
        v19 = (^(v19))
        _ = fi22
    }
    v21 = "hello"
    switch v21 {
    case "foo bar":
        fmt.Println(v21, v6)
    case "hello":
        var v23 uint16 = 42
        _ = v23
    case "\n\t":
        switch -1 {
        case 128:
            switch -444 {
            case 940:
                v24 := GenS2{F0: 1, F1: 42}
                _ = v24
            }
        }
    }
    v18 = v18
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
