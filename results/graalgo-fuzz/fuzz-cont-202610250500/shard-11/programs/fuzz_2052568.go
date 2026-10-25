package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint16 }
type GenS2 struct { F0 int16; F1 int16 }
func genFn0(p0 bool) int16 {
    p0 = (!p0)
    p0 = p0
    var v5 uint8 = 60
    _ = p0
    _ = v5
    return 1
}
func genFn1(p0 int, p1 int32, p2 uint8) int32 {
    var v6 int32 = (p1 << ((18446744073709551615) & 63))
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (p1 & (-(v6)))
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (!((!(false && false)) && true))
    v3 := make(map[uint32]string)
    v4 := make(map[string]int)
    if v2 {
        v8 := GenS1{F0: "\n\t", F1: 1, F2: 0}
        _ = v8
    } else {
        for fi9 := 0; fi9 < 5; fi9++ {
            if (!((-2147483648 > 42) || (0.0 < (1e10 - -2.25)))) {
                v10 := []int{(-(fi9)), (fi9 + fi9), fi9, (fi9 | (-(fi9)))}
                _ = v10
            } else {
                for fi11 := 0; fi11 < 1; fi11++ {
                    fmt.Println(fi11, fi11, fi9)
                    v2 = (32767 >= 42)
                    switch ("hello" + "hello") {
                    case "":
                        v2 = v2
                    case "a":
                        fmt.Println(fi9)
                    case "\n\t":
                        switch "a" {
                        case "a":
                            fmt.Println(fi11)
                        }
                        var v12 int16 = -163
                        var v13 int16 = v12
                        _ = v12
                        _ = v13
                    default:
                        delete(v4, "\n\t")
                    }
                    v2 = (0 == 2)
                    _ = fi11
                }
                switch 127 {
                case -9223372036854775808:
                    var v14 uint16 = 586
                    _ = v14
                default:
                    v2 = v2
                }
                v15 := []int{9223372036854775807, 42, (fi9 / ((fi9) | 1)), (fi9 | (fi9 + (fi9 ^ (^(fi9)))))}
                _ = v15
            }
            for fi16 := 0; fi16 < 19; fi16++ {
                var v17 int32 = 1
                _ = fi16
                _ = v17
            }
            v18 := GenS2{F0: -32768, F1: 32767}
            _ = fi9
            _ = v18
        }
        v2 = (v2 && (-379 < 9223372036854775807))
        var v19 uint64 = 42
        _ = v19
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        v21 := v3[42]
        v2 = (!v2)
        _ = fi20
        _ = v21
    }
    if v2 {
        v2 = true
    }
    switch ("hello" + "\n\t") {
    case "a":
        var v22 uint32 = 42
        _ = v22
    case "":
        v23 := GenS2{F0: -1, F1: -32768}
        _ = v23
    case "\n\t":
        v2 = false
    default:
        var v24 uint8 = 1
        _ = v24
    }
    v25 := []int{127, 9223372036854775807}
    v26 := GenS1{F0: "foo bar", F1: -631, F2: 1}
    v2 = v2
    for fi27 := 0; fi27 < 7; fi27++ {
        v2 = (!v2)
        fmt.Println(v2)
        v2 = false
        _ = fi27
    }
    var v28 int64 = 1
    var v29 int64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
