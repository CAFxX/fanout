package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint }
type GenS2 struct { F0 uint; F1 int8 }

func main() {
    v1 := []string{"foo bar", ("" + (("\n\t" + "a") + "hello"))}
    v2 := []int{-771, 9223372036854775807, 0}
    var v3 uint32 = 65535
    var v4 uint8 = 127
    v1 = append(v1, "a")
    for fi5 := 0; fi5 < 19; fi5++ {
        if false {
            var v6 uint8 = (^(v4))
            if true {
                v7 := GenS1{F0: 42, F1: uint(v6), F2: 971}
                _ = v7
            } else {
                if len(v2) > 0 {
                    v2[0] = 127
                }
                v8 := GenS1{F0: -568, F1: 766, F2: 0}
                _ = v8
            }
            _ = v6
        }
        v3 -= 1
        v9 := GenS1{F0: 91, F1: 0, F2: 870}
        _ = fi5
        _ = v9
    }
    switch -9223372036854775808 {
    case -274:
        var v10 uint64 = 168
        _ = v10
    default:
        v4 = (v4 - (v4 * 235))
    }
    if ((!true) && true) {
        var v11 int16 = 0
        v11 -= (v11 & 0)
        _ = v11
    }
    if (1000 < 805) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        v4 = (v4 ^ 0)
    }
    v12 := GenS1{F0: 1, F1: 989, F2: 255}
    v4 -= v4
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := GenS1{F0: 1, F1: uint(v3), F2: 255}
        v1 = append(v1, ("\n\t" + (("" + (("a" + "hello") + ("a" + "hello"))) + "\n\t")))
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        v16 := GenS2{F0: 874, F1: 127}
        for fi17 := 0; fi17 < 10; fi17++ {
            if len(v1) > 0 {
                v1[0] = "a"
            }
            var v18 uint = 42
            _ = fi17
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    v2 = append(v2, 42)
    v19 := GenS1{F0: 42, F1: 18446744073709551615, F2: 1}
    var v20 int64 = 518
    var v21 int16 = 242
    for fi22 := 0; fi22 < 4; fi22++ {
        v4 = v4
        v4 = (^(v4))
        _ = fi22
    }
    v23 := GenS1{F0: v20, F1: 657, F2: 99}
    v4 = 42
    for fi24 := 0; fi24 < 17; fi24++ {
        v21 = (v21 + (v21 ^ v21))
        v25 := GenS1{F0: v20, F1: 2, F2: 1}
        v2 = append(v2, int(v20))
        _ = fi24
        _ = v25
    }
    switch 127 {
    case -562:
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
    default:
        var v26 int8 = 1
        _ = v26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
}
