package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint }
func genFn0(p0 string, p1 bool) bool {
    fmt.Println(p0)
    v4 := []int{}
    p0 = ""
    _ = p0
    _ = p1
    _ = v4
    return p1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]int)
    switch -902 {
    case -1:
        var v5 bool = true
        _ = v5
    case 2147483647:
        for fi6 := 0; fi6 < 13; fi6++ {
            var v7 float32 = float32(0.1)
            switch 4294967295 {
            case 496:
                switch ("" + ((((("hello" + "\n\t") + "") + ("" + "\n\t")) + "hello") + "hello")) {
                case "":
                    if false {
                        var v8 uint16 = 65535
                        _ = v8
                    }
                }
            default:
                v9 := v2[585]
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    case 9223372036854775807:
        var v10 uint = 0
        _ = v10
    default:
        var v11 string = ("" + (((("a" + ("" + "")) + "foo bar") + "hello") + "a"))
        _ = v11
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        switch "" {
        case "a":
            for fi13 := 0; fi13 < 15; fi13++ {
                v14 := GenS1{F0: 127}
                fmt.Println(fi13, fi12)
                var v15 float32 = 3.14159
                v15 -= v15
                if false {
                    var v16 string = "hello"
                    _ = v16
                }
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        switch "\n\t" {
        case "foo bar":
            var v17 uint32 = 4294967295
            _ = v17
        case "":
            fmt.Println(fi12)
        case "hello":
            fmt.Println(fi12)
        default:
            switch "foo bar" {
            case "":
                v18 := GenS1{F0: -1}
                _ = v18
            }
        }
        v19 := GenS1{F0: 100}
        _ = fi12
        _ = v19
    }
    var v20 string = "\n\t"
    var v21 int8 = 127
    v22 := GenS2{F0: (-9223372036854775808 != -9223372036854775808), F1: 150, F2: 255}
    var v23 int16 = -617
    fmt.Println(v20)
    v24 := GenS1{F0: 100}
    switch 2147483647 {
    case 1:
        for fi25 := 0; fi25 < 3; fi25++ {
            var v26 int16 = (v23 ^ v23)
            if (!false) {
                for fi27 := 0; fi27 < 16; fi27++ {
                    fmt.Println(fi25, v20, v26)
                    var v28 int = 1
                    switch v20 {
                    case "foo bar":
                        var v29 int64 = 600
                        _ = v29
                    case "\n\t":
                        v30 := []int{(fi27 + (fi25 << ((18446744073709551615) & 63)))}
                        _ = v30
                    default:
                        v31 := []int{128}
                        _ = v31
                    }
                    _ = fi27
                    _ = v28
                }
            }
            fmt.Println(v21, v26, v23)
            _ = fi25
            _ = v26
        }
    case 32767:
        v32 := []int{-9223372036854775808, -105, -188, 1}
        _ = v32
    default:
        var v33 string = v20
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
