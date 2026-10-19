package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint }
func genFn0(p0 string, p1 int32, p2 uint8) string {
    switch 128 {
    case -571:
        v7 := GenS1{F0: (p2 * p2), F1: -1, F2: 492}
        _ = v7
    case 128:
        if ("\n\t" != "a") {
            v8 := GenS1{F0: p2, F1: 42, F2: 65535}
            _ = v8
        }
    case 127:
        var v9 float64 = (float64(((0.1 / (0.1 * 0.1)) / 0.1)) * ((-2.25 - 1.5) - (1e10 - 1.5)))
        _ = v9
    default:
        var v10 int = 127
        _ = v10
    }
    p0 = "hello"
    if (255 == 32767) {
        var v11 int8 = 0
        if false {
            switch ((("hello" + "") + "hello") + p0) {
            case "a":
                p2 = 255
            default:
                switch p0 {
                case "":
                    for fi12 := 0; fi12 < 16; fi12++ {
                        fmt.Println(p2, p2, v11)
                        _ = fi12
                    }
                case "a":
                    fmt.Println(p1, p2)
                }
            }
            var v13 float64 = ((0.1 / -2.25) + float64(-2.25))
            _ = v13
        } else {
            v11 -= v11
            p0 = "a"
        }
        for fi14 := 0; fi14 < 10; fi14++ {
            if (0 > (p2 - (p2 >> ((2) & 63)))) {
                fmt.Println(p1)
                v15 := GenS2{F0: 894, F1: 2, F2: 1}
                _ = v15
            } else {
                v16 := GenS2{F0: 465, F1: 139, F2: 287}
                var v17 uint = 857
                for fi18 := 0; fi18 < 6; fi18++ {
                    v19 := []int{(fi18 - fi14), 32767, (fi18 & fi14)}
                    _ = fi18
                    _ = v19
                }
                _ = v16
                _ = v17
            }
            _ = fi14
        }
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return "foo bar"
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[int]uint32)
    v3 := []uint32{65535}
    v4 := []uint32{203}
    v5 := make(map[int]bool)
    var v6 uint8 = 255
    var v20 uint64 = 4294967295
    var v21 uint32 = 42
    fmt.Println(v20, v21)
    switch -9223372036854775808 {
    case -9223372036854775808:
        v3 = append(v3, 2)
    }
    v22 := v5[129]
    fmt.Println(v6, v6)
    v23 := GenS2{F0: 2147483647, F1: (^(v21)), F2: 1}
    v1 += v1
    v21 -= (v21 ^ v21)
    v24 := GenS2{F0: 4294967295, F1: v21, F2: uint(v21)}
    fmt.Println(v20, v6, v21)
    if (!(false && (!(-2147483648 != 2147483647)))) {
        var v25 float32 = -2.25
        _ = len(v2)
        _ = v25
    }
    delete(v5, 42)
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
