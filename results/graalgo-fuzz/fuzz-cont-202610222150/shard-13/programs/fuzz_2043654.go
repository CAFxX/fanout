package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }

func main() {
    v1 := []uint32{1, 2, 2147483647, 2147483648}
    v2 := make(map[uint32]uint32)
    v3 := []string{"a", "", "a", ("foo bar" + "\n\t")}
    var v4 bool = (299 < 886)
    v5 := make(map[int]uint32)
    var v6 int32 = -2147483648
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := v2[2]
        var v9 int64 = -325
        v3 = append(v3, "")
        _ = fi7
        _ = v8
        _ = v9
    }
    if v4 {
        if (12 <= 576) {
            fmt.Println(v4, v4)
            v4 = true
        }
        if v4 {
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
        } else {
            fmt.Println(v4)
            fmt.Println(v6, v4)
        }
    }
    v10 := GenS1{F0: 127, F1: v4}
    v4 = v4
    v6 -= (v6 ^ v6)
    v4 = v4
    v4 = (0 > 2)
    fmt.Println(v4, v4)
    for fi11 := 0; fi11 < 1; fi11++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi12 := 0; fi12 < 3; fi12++ {
            v6 -= v6
            v13 := GenS1{F0: 45, F1: v4}
            if len(v1) > 0 {
                v1[0] = 248
            }
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    var v14 uint16 = 812
    switch "\n\t" {
    case "hello":
        var v15 float32 = 0.1
        _ = v15
    case "foo bar":
        fmt.Println(v4)
    default:
        switch 32768 {
        case 127:
            var v16 float32 = 0.0
            _ = v16
        case -1:
            if (0 >= 18446744073709551615) {
                v4 = (false && v4)
                if (559 >= 1) {
                    fmt.Println(v4, v4)
                }
            }
        case 32768:
            for fi17 := 0; fi17 < 11; fi17++ {
                var v18 string = "\n\t"
                _ = fi17
                _ = v18
            }
        }
    }
    if false {
        var v19 uint = 2
        for fi20 := 0; fi20 < 16; fi20++ {
            v21 := GenS1{F0: 64, F1: (("foo bar" + ("\n\t" + "")) < ("foo bar" + "\n\t"))}
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
