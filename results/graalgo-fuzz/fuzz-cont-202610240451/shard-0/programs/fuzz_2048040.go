package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int }
type GenS2 struct { F0 uint32 }
func genFn0() uint64 {
    var v4 int64 = 0
    v5 := []int{}
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 string, p1 bool, p2 int32) uint32 {
    p1 = ((32767 != 255) && p1)
    v6 := []int{42, 881, 42}
    switch 32767 {
    case -2147483648:
        fmt.Println(p0, p2, p1)
    case 795:
        if p1 {
            if len(v6) > 0 {
                v6[0] = 42
            }
            v6 = append(v6, 2147483647)
        }
    default:
        switch 32767 {
        case 42:
            switch "hello" {
            case "\n\t":
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case "":
                for fi7 := 0; fi7 < 9; fi7++ {
                    for fi8 := 0; fi8 < 19; fi8++ {
                        p1 = p1
                        fmt.Println(p1)
                        if (605 <= 308) {
                            p2 += (^(p2))
                            v9 := GenS1{F0: (3.14159 != ((0.1 - (0.1 + (3.14159 + 3.14159))) / float32(1e10))), F1: -1, F2: fi7}
                            _ = len(v6)
                            _ = v9
                        } else {
                            p2 = (p2 * (p2 << ((415) & 63)))
                        }
                        _ = fi8
                    }
                    _ = fi7
                }
            case "a":
                var v10 int32 = p2
                _ = v10
            default:
                if p1 {
                    fmt.Println(p2, p0, p2)
                }
            }
        case -332:
            var v11 uint8 = 128
            _ = v11
        case 0:
            p0 = p0
        default:
            var v12 uint16 = 2
            _ = v12
        }
    }
    var v13 int64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v6
    return 42
}

func main() {
    v1 := []string{"", "a"}
    var v2 int16 = -32768
    var v3 bool = true
    _ = len(v1)
    var v14 uint32 = 1
    v15 := GenS2{F0: (v14 ^ 1)}
    var v16 int32 = -2147483648
    v14 = v14
    var v17 int64 = -1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        var v18 uint64 = 592
        _ = v18
    }
    if ((45 != genFn0()) || false) {
        v16 += v16
    }
    if (false && v3) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        for fi19 := 0; fi19 < 12; fi19++ {
            fmt.Println(v2)
            _ = fi19
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
