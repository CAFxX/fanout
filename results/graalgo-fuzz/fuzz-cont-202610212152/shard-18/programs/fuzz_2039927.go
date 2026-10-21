package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int64 }
type GenS2 struct { F0 uint64; F1 int }
func genFn0(p0 int8, p1 int32) uint32 {
    var v7 string = (("" + (("" + "foo bar") + (((("foo bar" + ("hello" + "\n\t")) + "hello") + "\n\t") + "hello"))) + "\n\t")
    v8 := []int{-1, -9223372036854775808, -356}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 634
}

func main() {
    v1 := []int{32767, -546, 127, 2147483647}
    v2 := []uint32{0, 65535}
    v3 := make(map[string]int)
    var v4 int16 = 896
    var v5 uint16 = 1
    var v6 string = "a"
    if true {
        if (42 > 9223372036854775807) {
            v4 -= v4
            delete(v3, "hello")
        }
        v4 -= v4
        var v9 bool = (2 != genFn0(1, 2147483647))
        _ = v9
    } else {
        var v10 string = v6
        _ = v10
    }
    fmt.Println(v5, v5, v6)
    switch v6 {
    case "foo bar":
        v6 = ((v6 + v6) + "a")
    case "\n\t":
        switch 1 {
        case 1:
            fmt.Println(v5)
        }
    case "":
        for fi11 := 0; fi11 < 1; fi11++ {
            for fi12 := 0; fi12 < 13; fi12++ {
                var v13 string = v6
                fmt.Println(fi11)
                var v14 int8 = 113
                fmt.Println(v14)
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi11
        }
    default:
        v15 := GenS1{F0: -9223372036854775808, F1: false, F2: -195}
        _ = v15
    }
    for fi16 := 0; fi16 < 20; fi16++ {
        v4 += (v4 >> ((4294967295) & 63))
        switch ("a" + v6) {
        case "\n\t":
            v17 := GenS2{F0: 440, F1: fi16}
            _ = v17
        case "hello":
            for fi18 := 0; fi18 < 13; fi18++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v1 = append(v1, (-(fi18)))
                _ = fi18
            }
        default:
            switch v6 {
            case "\n\t":
                fmt.Println(v6)
            }
        }
        _ = fi16
    }
    if len(v2) > 0 {
        v2[0] = uint32(v5)
    }
    v6 = (((v6 + v6) + "a") + "hello")
    var v19 int = 2147483647
    v1 = append(v1, (v19 * (v19 + -2147483648)))
    switch v6 {
    case "a":
        v20 := GenS2{F0: 635, F1: (v19 * (^(v19)))}
        _ = v20
    default:
        v19 = (v19 / ((128) | 1))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
