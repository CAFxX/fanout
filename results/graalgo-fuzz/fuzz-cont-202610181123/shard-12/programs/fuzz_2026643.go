package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 bool }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int32, p1 uint16) uint {
    if false {
        p1 += 0
        if (!false) {
            var v5 uint64 = 4294967295
            _ = v5
        } else {
            fmt.Println(p1)
            p1 = (p1 / (((p1 << ((2) & 63))) | 1))
        }
    } else {
        v6 := []int{4294967295, 0, -398, 641}
        v6 = append(v6, -810)
        _ = v6
    }
    v7 := GenS1{F0: 65535, F1: -1}
    fmt.Println(p1, p0)
    p0 = (p0 ^ (p0 ^ -823))
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 127
    var v2 int32 = -539
    var v3 int16 = -1
    v4 := []int{0, 338}
    switch ("\n\t" + "a") {
    case "foo bar":
        _ = len(v4)
    case "hello":
        if true {
            v2 = (v2 | (-(v2)))
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    case "\n\t":
        if true {
            v1 -= v1
            v4 = append(v4, int(v2))
            switch "hello" {
            case "\n\t":
                fmt.Println(v2, v1)
            case "hello":
                v8 := GenS2{F0: 32767, F1: 2147483648, F2: true}
                _ = v8
            }
        } else {
            for fi9 := 0; fi9 < 6; fi9++ {
                for fi10 := 0; fi10 < 17; fi10++ {
                    var v11 uint64 = 2
                    _ = fi10
                    _ = v11
                }
                v12 := GenS1{F0: 1, F1: v3}
                _ = fi9
                _ = v12
            }
        }
    default:
        var v13 string = ("\n\t" + ((("\n\t" + ("hello" + "a")) + "a") + (("" + "hello") + ("" + "\n\t"))))
        _ = v13
    }
    v2 = 386
    v4 = append(v4, 32767)
    fmt.Println(v1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 = -295
    var v14 float32 = 0.0
    v15 := GenS3{F0: (v1 & v1)}
    if len(v4) > 0 {
        v4[0] = -2147483648
    }
    v14 = v14
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 int16 = v3
        var v18 bool = (((fi16 <= fi16) || false) && true)
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
