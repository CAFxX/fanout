package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
func genFn0(p0 string, p1 string, p2 uint64) uint64 {
    switch ((p0 + "\n\t") + p1) {
    case "\n\t":
        p1 = p1
    default:
        var v5 int16 = 42
        _ = v5
    }
    v6 := []int{}
    v7 := GenS1{F0: true, F1: (p2 << ((770) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p2
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 2
    v3 := []int{}
    v4 := make(map[uint32]uint32)
    var v8 bool = false
    v1[896] = "\n\t"
    for fi9 := 0; fi9 < 15; fi9++ {
        if v8 {
            var v10 uint16 = 77
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            _ = v10
        } else {
            v11 := GenS1{F0: v8, F1: (v2 - (v2 >> ((42) & 63)))}
            _ = v11
        }
        _ = fi9
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        for fi13 := 0; fi13 < 5; fi13++ {
            fmt.Println(fi12, v2)
            var v14 int64 = -9223372036854775808
            _ = fi13
            _ = v14
        }
        v15 := GenS1{F0: true, F1: 4294967295}
        var v16 uint32 = 1
        _ = fi12
        _ = v15
        _ = v16
    }
    fmt.Println(v8, v8)
    v17 := GenS1{F0: (-898 >= -2147483648), F1: genFn0("hello", "\n\t", (v2 / ((v2) | 1)))}
    for fi18 := 0; fi18 < 17; fi18++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v19 string = ("" + "\n\t")
        v20 := v4[731]
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println(v2)
    switch ("a" + "a") {
    case "a":
        v2 = genFn0("hello", ((("a" + "hello") + ("\n\t" + "foo bar")) + ("\n\t" + "foo bar")), (v2 % (((^(v2))) | 1)))
    case "hello":
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 float64 = -2.25
            switch "a" {
            case "\n\t":
                v2 = 0
            case "hello":
                if ((v8 && v8) || v8) {
                    v23 := GenS1{F0: v8, F1: (v2 * v2)}
                    _ = v23
                }
            case "foo bar":
                switch (fi21 + (fi21 << ((780) & 63))) {
                case -405:
                    switch "a" {
                    case "foo bar":
                        for fi24 := 0; fi24 < 4; fi24++ {
                            v25 := GenS1{F0: false, F1: (v2 | (v2 / ((v2) | 1)))}
                            _ = fi24
                            _ = v25
                        }
                    default:
                        v2 = (v2 >> ((1) & 63))
                    }
                case -539:
                    var v26 uint8 = 42
                    _ = v26
                case 9223372036854775807:
                    v27 := GenS1{F0: (v8 && (0 >= 0)), F1: (v2 << ((65535) & 63))}
                    _ = v27
                default:
                    v8 = true
                }
            }
            _ = fi21
            _ = v22
        }
    default:
        _ = len(v1)
    }
    fmt.Println(v8)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
