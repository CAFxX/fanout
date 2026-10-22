package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int32; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 uint32; F2 int }
func genFn0(p0 uint32, p1 uint32, p2 string) uint32 {
    v4 := []int{0}
    if ("hello" < (p2 + ("\n\t" + "foo bar"))) {
        fmt.Println(p2)
        switch "\n\t" {
        case "a":
            for fi5 := 0; fi5 < 9; fi5++ {
                for fi6 := 0; fi6 < 11; fi6++ {
                    v7 := GenS1{F0: "hello", F1: 513}
                    _ = fi6
                    _ = v7
                }
                v8 := GenS1{F0: p2, F1: 18446744073709551615}
                var v9 uint16 = 255
                _ = fi5
                _ = v8
                _ = v9
            }
        case "hello":
            p2 = (p2 + p2)
        default:
            var v10 string = ((("a" + "hello") + p2) + (p2 + ("foo bar" + p2)))
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 float64 = 0.0
        p0 = p0
        _ = fi11
        _ = v12
    }
    p1 = p1
    var v13 uint = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v4
    return (p1 << ((4294967295) & 63))
}

func main() {
    v1 := []int{}
    var v2 uint64 = 18446744073709551615
    var v3 bool = true
    for fi14 := 0; fi14 < 16; fi14++ {
        for fi15 := 0; fi15 < 15; fi15++ {
            v16 := GenS1{F0: "\n\t", F1: v2}
            fmt.Println(v2, fi14, fi14)
            _ = fi15
            _ = v16
        }
        v2 = v2
        var v17 float64 = -2.25
        _ = fi14
        _ = v17
    }
    if v3 {
        var v18 int8 = -37
        switch 9223372036854775807 {
        case 128:
            if (42 != 65535) {
                if len(v1) > 0 {
                    v1[0] = 811
                }
                v3 = v3
                v2 -= (v2 - 42)
            }
        case -2147483648:
            v18 = -128
        case 1:
            fmt.Println(v3, v3)
        default:
            fmt.Println(v3)
        }
        v2 -= (^(v2))
        _ = v18
    }
    v19 := GenS2{F0: 1, F1: true, F2: 1}
    v20 := GenS1{F0: ("\n\t" + "hello"), F1: (v2 << ((840) & 63))}
    var v21 uint64 = v2
    v22 := GenS3{F0: (v2 / ((v21) | 1))}
    v23 := GenS1{F0: "hello", F1: (^(v21))}
    var v24 string = "\n\t"
    v25 := GenS3{F0: (v21 << ((422) & 63))}
    var v26 uint64 = (^(v2))
    for fi27 := 0; fi27 < 10; fi27++ {
        switch v24 {
        case "":
            v28 := GenS4{F0: 255, F1: 2147483648, F2: fi27}
            _ = v28
        }
        v2 -= v26
        _ = fi27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
