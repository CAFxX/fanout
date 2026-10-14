package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int }

func main() {
    var v1 int64 = 0
    var v2 uint = 1
    var v3 int64 = (v1 * v1)
    v4 := make(map[int]uint32)
    v5 := []int{4294967295, 9223372036854775807, -732, -712}
    var v6 uint64 = 1
    if true {
        var v7 uint16 = uint16(v6)
        v8 := GenS1{F0: "hello", F1: ("a" + (("foo bar" + (("" + "foo bar") + "hello")) + "foo bar")), F2: 759}
        switch 954 {
        case -506:
            var v9 uint64 = v6
            _ = v9
        default:
            for fi10 := 0; fi10 < 17; fi10++ {
                v7 -= 633
                _ = fi10
            }
        }
        _ = v7
        _ = v8
    } else {
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 int = -1
            v1 += (v1 + (-(v3)))
            for fi13 := 0; fi13 < 15; fi13++ {
                var v14 int = -932
                if (!(v6 < 1)) {
                    v15 := GenS1{F0: "\n\t", F1: ("\n\t" + "hello"), F2: (fi11 << (((v2 / ((65535) | 1))) & 63))}
                    if ("\n\t" > "foo bar") {
                        v1 = v1
                        fmt.Println(v2, v12)
                        var v16 int32 = 0
                        _ = v16
                    } else {
                        v4[fi13] = 80
                        v6 = (^(v6))
                    }
                    v17 := GenS1{F0: "hello", F1: "hello", F2: 9223372036854775807}
                    _ = v15
                    _ = v17
                }
                _ = fi13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
        v5 = append(v5, 1)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    switch 32767 {
    case -1:
        var v18 uint32 = 1
        _ = v18
    case -576:
        var v19 uint32 = 65535
        _ = v19
    default:
        for fi20 := 0; fi20 < 20; fi20++ {
            v1 = v3
            _ = fi20
        }
    }
    v6 = (v6 & uint64(v2))
    v21 := GenS1{F0: "\n\t", F1: "hello", F2: -2147483648}
    fmt.Println(v6, v2, v1)
    v1 = 42
    if ((178 < 1) && false) {
        v6 = (v6 - v6)
        var v22 int16 = -892
        _ = v22
    } else {
        for fi23 := 0; fi23 < 7; fi23++ {
            fmt.Println(v2, v6)
            _ = fi23
        }
    }
    v2 = (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
